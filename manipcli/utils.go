package manipcli

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/araddon/dateparse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniphttp"
	"go.uber.org/zap"
)

// ManipulatorMaker returs a function which can create a manipulator based on pflags.
type ManipulatorMaker = func() (manipulate.Manipulator, error)

// ManipulatorMakerFromFlags returns a func that creates a manipulator based on command flags. Command flags are read using viper.
// It needs the following flags: FlagAPI, FlagToken, FlagAppCredentials, FlagNamespace, FlagCACertPath, FlagAPISkipVerify, FlagEncoding
// Use SetCLIFlags to add these flags to your command.
func ManipulatorMakerFromFlags(options ...maniphttp.Option) ManipulatorMaker {

	return func() (manipulate.Manipulator, error) {
		api := viper.GetString(flagAPI)
		token := viper.GetString(flagToken)
		// credsPath := viper.GetString(flagAppCredentials)
		namespace := viper.GetString(flagNamespace)
		capath := viper.GetString(flagCACertPath)
		skip := viper.GetBool(flagAPISkipVerify)
		encoding := viper.GetString(flagEncoding)

		var tlsConfig *tls.Config

		var enc elemental.EncodingType
		switch encoding {
		case "json":
			enc = elemental.EncodingTypeJSON
		case "msgpack":
			enc = elemental.EncodingTypeMSGPACK
		default:
			return nil, fmt.Errorf("unsupported encoding '%s'. Must be 'json' or 'msgpack'", encoding)
		}

		if tlsConfig == nil && capath != "" {
			rootCAPool, err := prepareAPICACertPool(capath)
			if err != nil {
				return nil, fmt.Errorf("unable to load root ca pool: %s", err)
			}
			tlsConfig = &tls.Config{
				InsecureSkipVerify: skip,
				RootCAs:            rootCAPool,
			}
		}

		opts := []maniphttp.Option{
			maniphttp.OptionNamespace(namespace),
			maniphttp.OptionTLSConfig(tlsConfig),
			maniphttp.OptionEncoding(enc),
			maniphttp.OptionToken(token),
		}

		if len(options) > 0 {
			opts = append(opts, options...)
		}

		return maniphttp.New(
			context.Background(),
			api,
			opts...,
		)
	}
}

// ManipulatorFlagSet returns the flagSet required to call ManipulatorFromFlags.
func ManipulatorFlagSet() *pflag.FlagSet {

	set := pflag.NewFlagSet("", pflag.ExitOnError)

	set.StringP(flagAPI, "A", "", "Server API URL.") // default is managed inline.
	set.BoolP(flagAPISkipVerify, "", false, "If set, skip api endpoint verification. This is insecure.")
	set.String(flagCACertPath, "", "Path to the CA to use for validating api endpoint.")
	set.String(flagTrackingID, "", "ID to trace the request. Use this when asked to help debug the system.")
	set.String(flagEncoding, "msgpack", "encoding to use to communicate with the platform. Can be 'msgpack' or 'json'")
	set.StringP(flagNamespace, "n", "", "Namespace to use.")

	return set
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

	if cmd.Root().PersistentPreRunE != nil {
		if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
			return err
		}
	}

	if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
		return err
	}

	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	if err := validateOutputParameters(viper.GetString(flagOutput)); err != nil {
		return err
	}

	return checkRequiredFlags(cmd.Flags())
}

// validateOutputParameters validates output parameters are correct
func validateOutputParameters(output string) error {

	// Check output constraints
	allOutputOptions := []string{
		flagOutputTable,
		flagOutputJSON,
		flagOutputNone,
		flagOutputDefault,
		flagOutputYAML,
		flagOutputTemplate,
	}

	for _, opt := range allOutputOptions {
		if opt == output {
			return nil
		}
	}
	return fmt.Errorf("invalid output %s", output)
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

	if _, err := hex.DecodeString(identifiable.Identifier()); err == nil {
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

		flagName := nameToFlag(spec.Name)

		if !viper.IsSet(flagName) {
			continue
		}

		fv := rv.FieldByName(spec.ConvertedName)

		switch spec.Type {
		case "string":
			fv.SetString(viper.GetString(flagName))

		case "enum":
			fv.SetString(viper.GetString(flagName))

		case "float64":
			fv.SetFloat(viper.GetFloat64(flagName))

		case "boolean":
			fv.SetBool(viper.GetBool(flagName))

		case "integer":
			fv.SetInt(int64(viper.GetInt(flagName)))

		case "time":
			t, err := dateparse.ParseAny(viper.GetString(flagName))
			if err != nil {
				return err
			}
			fv.Set(reflect.ValueOf(t))

		default:
			t := reflect.TypeOf(specifiable.ValueForAttribute(flagName))
			v := reflect.New(t)
			if err := json.Unmarshal([]byte(viper.GetString(flagName)), v.Interface()); err != nil {
				return err
			}
			fv.Set(v.Elem())
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

		flagName := nameToFlag(spec.Name)

		// Register flag based on type
		switch spec.Type {

		case "string":
			cmd.Flags().StringP(flagName, "", "", spec.Description)

		case "enum":
			cmd.Flags().StringP(flagName, "", "", spec.Description)

		case "float64":
			cmd.Flags().Float64P(flagName, "", 0, spec.Description)

		case "boolean":
			cmd.Flags().BoolP(flagName, "", false, spec.Description)

		case "integer":
			cmd.Flags().IntP(flagName, "", 0, spec.Description)

		case "time":
			cmd.Flags().StringP(flagName, "", "", spec.Description)

		default:
			zap.L().Debug("use default type string for attribute", zap.String("attribute", spec.Name), zap.String("identity", identifiable.Identity().Name))
			cmd.Flags().StringP(flagName, "", "", spec.Description)
		}

		if forceRequired && spec.Required {
			if err := cmd.MarkFlagRequired(flagName); err != nil {
				return fmt.Errorf("unable to mark flag %s as required", flagName)
			}
		}
	}

	return nil
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func nameToFlag(name string) string {
	flag := matchFirstCap.ReplaceAllString(name, "${1}-${2}")
	flag = matchAllCap.ReplaceAllString(flag, "${1}-${2}")
	return strings.ToLower(flag)
}
