package manipcli

import (
	"bytes"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/araddon/dateparse"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.uber.org/zap"
	"k8s.io/helm/pkg/strvals"
)

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

// shouldManageAttribute indicates if the attribute should be managed or not.
func shouldManageAttribute(spec elemental.AttributeSpecification) bool {

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

// readViperFlags reads all vipers flags without prefix
func readViperFlags(identifiable elemental.Identifiable, modelManager elemental.ModelManager, prefix string) error {

	if identifiable == nil {
		return fmt.Errorf("provided identifiable is nil")
	}
	specifiable, ok := identifiable.(elemental.AttributeSpecifiable)
	if !ok {
		return fmt.Errorf("%s is not an AttributeSpecifiable", identifiable.Identity().Name)
	}

	_, err := readViperFlagsWithPrefix(specifiable, modelManager, prefix)
	return err
}

// readViperFlags reads all viper flags and fill the identifiable properties.
// TODO: Make it better and add more types here.
func readViperFlagsWithPrefix(specifiable elemental.AttributeSpecifiable, modelManager elemental.ModelManager, prefix string) (bool, error) {

	var didSet bool

	rv := reflect.ValueOf(specifiable).Elem()

	for _, attrSpec := range specifiable.AttributeSpecifications() {

		if !shouldManageAttribute(attrSpec) {
			continue
		}

		flagName := nameToFlag(attrSpec.Name)
		if prefix != "" {
			flagName = fmt.Sprintf("%s.%s", prefix, flagName)
		}

		fv := rv.FieldByName(attrSpec.ConvertedName)

		switch attrSpec.Type {

		case "string", "enum":
			if !viper.IsSet(flagName) {
				continue
			}
			fv.SetString(viper.GetString(flagName))
			didSet = true

		case "float64":
			if !viper.IsSet(flagName) {
				continue
			}
			fv.SetFloat(viper.GetFloat64(flagName))
			didSet = true

		case "boolean":
			if !viper.IsSet(flagName) {
				continue
			}
			fv.SetBool(viper.GetBool(flagName))
			didSet = true

		case "integer":
			if !viper.IsSet(flagName) {
				continue
			}
			fv.SetInt(viper.GetInt64(flagName))
			didSet = true

		case "time":
			if !viper.IsSet(flagName) {
				continue
			}

			t, err := dateparse.ParseAny(viper.GetString(flagName))
			if err != nil {
				return false, err
			}

			fv.Set(reflect.ValueOf(t))
			didSet = true

		case "ref":

			instance := reflect.New(fv.Type().Elem())
			specifiable, ok := instance.Interface().(elemental.AttributeSpecifiable)

			if !ok {
				if !viper.IsSet(flagName) {
					continue
				}

				rt := reflect.TypeOf(specifiable.ValueForAttribute(attrSpec.Name))
				if rt == nil {
					return false, fmt.Errorf("unable to find attribute %s", attrSpec.Name)
				}

				rv := reflect.New(rt)
				if err := json.Unmarshal([]byte(viper.GetString(flagName)), rv.Interface()); err != nil {
					return false, err
				}

				fv.Set(rv.Elem())
				didSet = true

				continue
			}

			innerDidSet, err := readViperFlagsWithPrefix(specifiable, modelManager, flagName)
			if err != nil {
				return false, err
			}

			if innerDidSet {
				fv.Set(reflect.ValueOf(specifiable))
				didSet = true
			}

		default:
			if !viper.IsSet(flagName) {
				continue
			}

			rt := reflect.TypeOf(specifiable.ValueForAttribute(attrSpec.Name))
			if rt == nil {
				return false, fmt.Errorf("unable to find attribute %s", attrSpec.Name)
			}

			rv := reflect.New(rt)
			if err := json.Unmarshal([]byte(viper.GetString(flagName)), rv.Interface()); err != nil {
				return false, err
			}

			fv.Set(rv.Elem())
			didSet = true
		}
	}

	return didSet, nil
}

// setViperFlags
func setViperFlags(cmd *cobra.Command, identifiable elemental.Identifiable, modelManager elemental.ModelManager, prefix string) error {

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

	return setViperFlagsWithPrefix(cmd, specifiable, modelManager, prefix)
}

// setViperFlagsWithPrefix sets the viper flags to the command according to the identifiable
// TODO: Make it better and add more types here.
func setViperFlagsWithPrefix(cmd *cobra.Command, specifiable elemental.AttributeSpecifiable, modelManager elemental.ModelManager, prefix string) error {

	rv := reflect.ValueOf(specifiable).Elem()

	for _, spec := range specifiable.AttributeSpecifications() {

		if !shouldManageAttribute(spec) {
			continue
		}

		flagName := nameToFlag(spec.Name)
		if prefix != "" {
			flagName = fmt.Sprintf("%s.%s", prefix, flagName)
		}

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

		case "ref":
			fv := rv.FieldByName(spec.ConvertedName)
			instance := reflect.New(fv.Type().Elem())

			specifiable, ok := instance.Interface().(elemental.AttributeSpecifiable)
			if !ok {
				cmd.Flags().StringP(flagName, "", "", spec.Description)
				continue
			}

			err := setViperFlagsWithPrefix(cmd, specifiable, modelManager, flagName)
			if err != nil {
				return err
			}

		default:
			zap.L().Debug("use default type string for attribute", zap.String("attribute", spec.Name))
			cmd.Flags().StringP(flagName, "", "", spec.Description)
		}
	}

	return nil
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllNumber = regexp.MustCompile("([a-z])([0-9]+)")
var matchAllCap = regexp.MustCompile("([a-z])([A-Z])")

func nameToFlag(name string) string {
	flag := matchFirstCap.ReplaceAllString(name, "${1}-${2}")
	flag = matchAllNumber.ReplaceAllString(flag, "${1}-${2}")
	flag = matchAllCap.ReplaceAllString(flag, "${1}-${2}")
	return strings.ToLower(flag)
}

type commonValues struct {
	Namespace string
	API       string
}

type tplValues struct {
	Values map[string]any
	Common commonValues
}

// ReadData reads data from a template coming from a url or a file.
// It can get some additional values from different sources (user input, values file or single values).
func ReadData(
	apiurl string,
	namespace string,
	file string,
	url string,
	valuesFile string,
	values []string,
	printOnly bool,
	mandatory bool,
) (data []byte, err error) {

	if url == "" && file == "" {
		if mandatory {
			return nil, fmt.Errorf("you must pass either %s or %s", flagInputFile, flagInputURL)
		}
		return nil, nil
	}

	if url != "" && file != "" {
		return nil, fmt.Errorf("you cannot set both %s and %s", flagInputFile, flagInputURL)
	}

	var processed bool

	// reading input-url
	if url != "" {
		/* #nosec */
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unable to retrieve data from url %s: %s", url, resp.Status)
		}

		defer resp.Body.Close() // nolint
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		processed = true
	}

	// Reading input-file
	if file != "" && file != "-" {

		data, err = os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		processed = true
	}

	if file == "-" {

		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}

		processed = true
	}

	if !processed && mandatory {
		return nil, fmt.Errorf("you must pass either %s or %s", flagInputFile, flagInputURL)
	}

	if !processed && !mandatory {
		return nil, nil
	}

	if printOnly {
		return data, nil
	}

	templateValues := map[string]any{}

	// reading input-values
	if valuesFile != "" {

		data, err := os.ReadFile(valuesFile)
		if err != nil {
			return nil, err
		}

		if err = yaml.Unmarshal(data, &templateValues); err != nil {
			return nil, err
		}
	}

	for _, value := range values {
		if err := strvals.ParseInto(value, templateValues); err != nil {
			return nil, err
		}
	}

	result, err := renderTemplate(string(data[:]), tplValues{
		Values: templateValues,
		Common: commonValues{
			Namespace: namespace,
			API:       apiurl,
		},
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// readData reads the data from a path or a url or stdin
func readData(mandatory bool) (data []byte, err error) {

	if inputData := viper.GetString(flagInputData); inputData != "" {
		return []byte(inputData), nil
	}

	data, err = ReadData(
		viper.GetString(flagAPI),
		viper.GetString(flagNamespace),
		viper.GetString(flagInputFile),
		viper.GetString(flagInputURL),
		viper.GetString(flagInputValues),
		viper.GetStringSlice(flagInputSet),
		viper.GetBool(flagPrint),
		mandatory,
	)
	if err != nil {
		return nil, err
	}

	if viper.GetBool(flagPrint) || viper.GetBool(flagRender) {
		fmt.Println(string(data))
		flushOutputAndExit(0)
	}

	return data, nil

}

func renderTemplate(content string, values any) ([]byte, error) {

	funcs := sprig.TxtFuncMap()
	funcs["required"] = func(warn string, val any) (any, error) {
		if val == nil {
			return val, fmt.Errorf(warn)
		} else if _, ok := val.(string); ok {
			if val == "" {
				return val, fmt.Errorf(warn)
			}
		}
		return val, nil
	}
	funcs["readFile"] = func(path string) (string, error) {
		data, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}

	tpl, err := template.New("template").Funcs(funcs).Parse(content)
	if err != nil {
		return nil, fmt.Errorf("unable to parse template: %s", err)
	}

	buf := &bytes.Buffer{}
	if err := tpl.Execute(
		buf,
		values,
	); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func openInEditor(
	identifiable elemental.Identifiable,
	editor string,
	message string,
	stripReadOnlyAttribute bool,
	stripCreationOnlyAttribute bool,
	showRawInitialData bool,
) ([]byte, error) {

	if editor == "" {
		editor, _ = os.LookupEnv("EDITOR")
		if editor == "" {
			return nil, fmt.Errorf("you must pass a valid --%s or set EDITOR environment variable", flagEditor)
		}
	}

	file, err := os.CreateTemp(os.TempDir(), "cli")
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name()) // nolint: errcheck

	data, err := generateFileData(
		identifiable,
		message,
		stripReadOnlyAttribute,
		stripCreationOnlyAttribute,
		showRawInitialData,
		outputFormat{
			formatType: formatTypeHash,
			output:     flagOutputYAML,
		},
	)
	if err != nil {
		return nil, err
	}

	if _, err = file.Write([]byte(data)); err != nil {
		return nil, err
	}
	if err = file.Close(); err != nil {
		return nil, err
	}

	var params []string
	switch editor {
	case "atom", "atom-beta", "code", "code-insiders":
		params = append(params, "-w")
	}

	params = append(params, file.Name())
	/* #nosec */
	cmd := exec.Command(editor, params...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err = cmd.Run(); err != nil {
		return nil, err
	}

	newData, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}

	if string(newData) == data {
		return nil, nil
	}

	reJSON, err := yaml.YAMLToJSON(newData)
	if err != nil {
		return nil, err
	}

	return reJSON, nil
}

// generateFileData generate []bytes containing
// a ready to write text file of the given identity.
func generateFileData(
	identifiable elemental.Identifiable,
	message string,
	stripReadOnlyAttribute bool,
	stripCreationOnlyAttribute bool,
	showRawInitialData bool,
	format outputFormat,
) (string, error) {

	if identifiable == nil {
		return "", fmt.Errorf("identifiable is nil")
	}
	initialData, err := formatObjectsStripped(
		format,
		stripReadOnlyAttribute,
		stripCreationOnlyAttribute,
		false,
		identifiable,
	)
	if err != nil {
		return "", err
	}

	sData := ""
	if message != "" {
		sData = fmt.Sprintf("# %s\n\n", message)
	}
	sData = sData + initialData

	if showRawInitialData {
		rawInitialData, e := formatObjects(format, false, identifiable)

		if e != nil {
			return "", e
		}
		rawInitialLines := strings.Split(rawInitialData, "\n")
		rawInitialLines = append(rawInitialLines, "\n\n")
		for i, l := range rawInitialLines {
			rawInitialLines[i] = "# " + l
		}

		sData = fmt.Sprintf("%s\n# Here is a copy of the full original object you are editing:\n#\n%s",
			sData,
			strings.Join(rawInitialLines[:len(rawInitialLines)-3], "\n"))
	}

	return sData, nil
}

// splitParentInfo splits the parent info into name, ID
func splitParentInfo(parent string) (string, string, error) {

	if parent == "" {
		return "", "", fmt.Errorf("no parent information provided")
	}

	data := strings.Split(parent, "/")

	if len(data) != 2 {
		return "", "", fmt.Errorf("invalid parent format, use name or ID")
	}

	return data[0], data[1], nil
}
