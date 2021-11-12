package cli

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strings"

	"github.com/araddon/dateparse"
	"github.com/globalsign/mgo/bson"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniphttp"
	"go.uber.org/zap"
)

// PrepareManipulator creates the manipulator used to process commands. Currently
// only HTTP manipulator is supported, and a constants.OptionTokenKey field is therefore required.
func prepareManipulator(api string, token string, credsPath string, namespace string, capath string, skip bool, encoding string) (manipulate.Manipulator, error) {

	var tlsConfig *tls.Config

	var enc elemental.EncodingType
	switch encoding {
	case "json":
		enc = elemental.EncodingTypeJSON
	case "msgpack":
		enc = elemental.EncodingTypeMSGPACK
	default:
		return nil, fmt.Errorf("unsuported encoding '%s'. Must be 'json' or 'msgpack'", encoding)
	}

	if tlsConfig == nil {
		rootCAPool, err := prepareAPICACertPool(capath)
		if err != nil {
			return nil, fmt.Errorf("unable to load root ca pool: %s", err)
		}
		tlsConfig = &tls.Config{
			InsecureSkipVerify: skip,
			RootCAs:            rootCAPool,
		}
	}

	// If by then we still don't have an api, we set it to console.
	if api == "" {
		api = "https://api.console.aporeto.com" // TODO: Try to put that in the default values...
	}

	return maniphttp.New(
		context.Background(),
		api,
		maniphttp.OptionNamespace(namespace),
		maniphttp.OptionTLSConfig(tlsConfig),
		maniphttp.OptionEncoding(enc),
		maniphttp.OptionToken(token),
		// maniphttp.OptionSendCredentialsAsCookie("x-aporeto-token"), // TODO: Check why this is necessary
	)
}

// prepareAPICACertPool prepares the API cert pool if not empty.
func prepareAPICACertPool(capath string) (*x509.CertPool, error) {

	if capath == "" {
		if runtime.GOOS == "windows" {
			// use nil as RootCAs on Windows in order to call systemVerify,
			// which will work even if Windows has not cached all its root certs.
			return nil, nil
		}
		return x509.SystemCertPool()
	}

	capool := x509.NewCertPool()
	cadata, err := os.ReadFile(capath)
	if err != nil {
		return nil, err
	}

	capool.AppendCertsFromPEM(cadata)

	return capool, nil
}

// shouldManageSpecification indicates if the attribute should be managed or not.
func shouldManageSpecification(spec elemental.AttributeSpecification) bool {

	if !spec.Exposed {
		return false
	}

	if spec.PrimaryKey {
		return false
	}

	if spec.Autogenerated {
		return false
	}

	if spec.ReadOnly {
		return false
	}

	if spec.Type == "external" {
		return false
	}

	return true
}

// ParametersToURLValues converts the list of `key=value` to url.Values.
func parametersToURLValues(params []string) (url.Values, error) {

	values := url.Values{}
	for _, keyVal := range params {

		kV := strings.SplitN(keyVal, "=", 2)

		if len(kV) != 2 {
			return nil, fmt.Errorf("invalid parameter %s", keyVal)
		}

		values.Add(kV[0], kV[1])
	}

	return values, nil
}

// persistentPreRunE ensure all bindings are done and validate required flags.
func persistentPreRunE(cmd *cobra.Command, args []string) error {

	if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
		return err
	}

	if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
		return err
	}

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	if err := validateOutputParameters(viper.GetString(FlagOutput)); err != nil {
		return err
	}

	return checkRequiredFlags(cmd.Flags())
}

// validateOutputParameters validates output parameters are correct
func validateOutputParameters(output string) error {

	// Check output constraints
	allOutputOptions := []string{
		FlagOutputTable,
		FlagOutputJSON,
		FlagOutputNone,
		FlagOutputDefault,
		FlagOutputYAML,
		FlagOutputTemplate,
	}

	for _, opt := range allOutputOptions {
		if opt == output {
			return nil
		}
	}
	return fmt.Errorf("invalid output %s.", output)
}

// checkRequiredFlags checks if all required flags are set
func checkRequiredFlags(flags *pflag.FlagSet) error {
	requiredError := false
	flagName := ""

	flags.VisitAll(func(flag *pflag.Flag) {
		requiredAnnotation := flag.Annotations[cobra.BashCompOneRequiredFlag]
		if len(requiredAnnotation) == 0 {
			return
		}

		flagRequired := requiredAnnotation[0] == "true"

		if flagRequired && !flag.Changed {
			requiredError = true
			flagName = flag.Name
		}
	})

	if requiredError {
		return errors.New("Required argument `--" + flagName + "` must be passed")
	}

	return nil
}

// retrieveByIDOrByName retrieves an object from its id or name
func retrieveObjectByIDOrByName(
	ctx manipulate.Context,
	manipulator manipulate.Manipulator,
	identity elemental.Identity,
	idOrName string,
	modelManager elemental.ModelManager,
) (elemental.Identifiable, error) {

	identifiable := modelManager.IdentifiableFromString(identity.Name)
	identifiable.SetIdentifier(idOrName)

	if bson.IsObjectIdHex(identifiable.Identifier()) {
		if err := manipulator.Retrieve(ctx, identifiable); err != nil {
			return nil, err
		}
		return identifiable, nil
	}

	// let's try to find the object by name
	var dest elemental.Identifiables
	if _, ok := identifiable.(elemental.SparseIdentifiable); ok {
		dest = modelManager.SparseIdentifiables(identifiable.Identity())
	} else {
		dest = modelManager.Identifiables(identifiable.Identity())
	}

	mctx := ctx.Derive(
		manipulate.ContextOptionFilter(
			elemental.NewFilter().
				WithKey("name").Equals(identifiable.Identifier()).
				Done(),
		),
	)

	if err := manipulator.RetrieveMany(mctx, dest); err != nil {
		return nil, err
	}

	lst := dest.List()

	if len(lst) == 0 {
		return nil, fmt.Errorf("no %s found with id or name %s", identifiable.Identity().Name, idOrName)
	}

	if len(lst) > 1 {
		return nil, fmt.Errorf("more than one %s has been found with id or name %s. Use ID", identifiable.Identity().Name, idOrName)
	}

	return lst[0], nil
}

// readViperFlags reads all viper flags and fill the identifiable properties.
// TODO: Make it better and add more types here.
func readViperFlags(identifiable elemental.Identifiable) error {

	if identifiable == nil {
		return fmt.Errorf("provided identifiable is nil")
	}
	specifiable, ok := identifiable.(elemental.AttributeSpecifiable)
	if !ok {
		return fmt.Errorf("%s is not an AttributeSpecifiable", identifiable.Identity().Name)
	}

	rv := reflect.ValueOf(identifiable).Elem()

	for _, spec := range specifiable.AttributeSpecifications() {

		if !shouldManageSpecification(spec) {
			continue
		}

		if !viper.IsSet(spec.Name) {
			continue
		}

		fv := rv.FieldByName(spec.ConvertedName)

		switch spec.Type {
		case "string":
			fv.SetString(viper.GetString(spec.Name))

		case "enum":
			fv.SetString(viper.GetString(spec.Name))

		case "float64":
			fv.SetFloat(viper.GetFloat64(spec.Name))

		case "boolean":
			fv.SetBool(viper.GetBool(spec.Name))

		case "integer":
			fv.SetInt(int64(viper.GetInt(spec.Name)))

		case "time":
			t, err := dateparse.ParseAny(viper.GetString(spec.Name))
			if err != nil {
				return err
			}
			fv.Set(reflect.ValueOf(t))

		case "list":
			fv.Set(reflect.ValueOf(viper.GetStringSlice(spec.Name)))

		default:
			fmt.Println("use default value for ", spec.Name, spec.Type)
			fv.Set(reflect.ValueOf(viper.GetString(spec.Name)))
		}
	}

	return nil
}

// setViperFlags sets the viper flags to the command according to the identifiable
// TODO: Make it better and add more types here.
func setViperFlags(cmd *cobra.Command, identifiable elemental.Identifiable, forceRequired bool) error {

	if cmd == nil {
		return fmt.Errorf("provided command is nil")
	}

	if identifiable == nil {
		return fmt.Errorf("provided identifiable is nil")
	}

	specifiable, ok := identifiable.(elemental.AttributeSpecifiable)
	if !ok {
		return fmt.Errorf("%s is not an AttributeSpecifiable", identifiable.Identity().Name)
	}

	for _, spec := range specifiable.AttributeSpecifications() {

		if !shouldManageSpecification(spec) {
			continue
		}

		// Register flag based on type
		switch spec.Type {

		case "string":
			cmd.Flags().StringP(spec.Name, "", "", spec.Description)

		case "enum":
			cmd.Flags().StringP(spec.Name, "", "", spec.Description)

		case "float64":
			cmd.Flags().Float64P(spec.Name, "", 0, spec.Description)

		case "boolean":
			cmd.Flags().BoolP(spec.Name, "", false, spec.Description)

		case "integer":
			cmd.Flags().IntP(spec.Name, "", 0, spec.Description)

		case "time":
			cmd.Flags().StringP(spec.Name, "", "", spec.Description)

		case "list":
			cmd.Flags().StringSliceP(spec.Name, "", nil, spec.Description)

		default:
			zap.L().Debug("use default type string for attribute", zap.String("attribute", spec.Name), zap.String("identity", identifiable.Identity().Name))
			cmd.Flags().StringP(spec.Name, "", "", spec.Description)
		}

		if forceRequired && spec.Required {
			if err := cmd.MarkFlagRequired(spec.Name); err != nil {
				return fmt.Errorf("unable to mark flag %s as required", spec.Name)
			}
		}
	}

	return nil
}
