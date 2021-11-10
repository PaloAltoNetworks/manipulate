package cli

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/template"

	prettyjson "github.com/hokaccha/go-prettyjson"
	"github.com/olekukonko/tablewriter"
	"go.aporeto.io/elemental"
	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"
)

type prepareOutputConfig struct {
	tableCaption string
}

// PrepareOutputOption represents an option that can be passed to PrepareOutputFormat.
type PrepareOutputOption func(*prepareOutputConfig)

// PrepareOutputOptionTableCaption can be used to add a caption to a table output.
func PrepareOutputOptionTableCaption(caption string) PrepareOutputOption {
	return func(config *prepareOutputConfig) {
		config.tableCaption = caption
	}
}

// OutputFormat retains all output information
type OutputFormat struct {
	columns      []string
	formatType   string
	output       string
	template     string
	tableCaption string
}

// prepareOutputFormat returns an OutputFormat structure that contains output information
func prepareOutputFormat(output string, formatType string, columns []string, template string, opts ...PrepareOutputOption) OutputFormat {

	cfg := prepareOutputConfig{}
	for _, o := range opts {
		o(&cfg)
	}

	return OutputFormat{
		columns:      columns,
		formatType:   formatType,
		output:       output,
		template:     template,
		tableCaption: cfg.tableCaption,
	}
}

// FormatEvents prints all events given an output format
func FormatEvents(outputFormat OutputFormat, forceList bool, events ...*elemental.Event) (string, error) {

	objectMaps := make([]map[string]interface{}, 0, len(events))

	for _, event := range events {
		if event.Encoding == elemental.EncodingTypeMSGPACK {
			if err := event.Convert(elemental.EncodingTypeJSON); err != nil {
				return "", err
			}
		}

		var objectMap map[string]interface{}
		if err := Remarshal(event, &objectMap); err == nil {
			objectMaps = append(objectMaps, objectMap)
		}
	}

	output, err := formatMaps(outputFormat, forceList, objectMaps)
	if err != nil {
		return "", err
	}

	return output, nil
}

// FormatObjects prints all identifiable objects given an output format
func FormatObjects(outputFormat OutputFormat, forceList bool, objects ...elemental.Identifiable) (string, error) {

	return FormatObjectsStripped(outputFormat, false, false, forceList, objects...)
}

// FormatObjectsStripped prints all identifiable objects given an output format and eventually strip out some data.
func FormatObjectsStripped(outputFormat OutputFormat, stripReadOnly bool, stripCreationOnly bool, forceList bool, objects ...elemental.Identifiable) (string, error) {

	objectMaps := make([]map[string]interface{}, 0, len(objects))

	for _, object := range objects {
		var objectMap map[string]interface{}
		if err := Remarshal(object, &objectMap); err == nil {
			objectMaps = append(objectMaps, objectMap)
		}

		if ats, ok := object.(elemental.AttributeSpecifiable); ok && stripReadOnly {
			for _, spec := range ats.AttributeSpecifications() {
				if spec.ReadOnly {
					delete(objectMap, spec.Name)
				}
			}
		}

		if ats, ok := object.(elemental.AttributeSpecifiable); ok && stripCreationOnly {
			for _, spec := range ats.AttributeSpecifications() {
				if spec.CreationOnly {
					delete(objectMap, spec.Name)
				}
			}
		}
	}

	output, err := formatMaps(outputFormat, forceList, objectMaps)
	if err != nil {
		return "", err
	}

	return output, nil
}

// FormatRandomMap prints all maps object
func FormatRandomMap(outputFormat OutputFormat, forceList bool, objects ...map[string]interface{}) string {
	output, err := formatMaps(outputFormat, forceList, objects)

	if err != nil {
		zap.L().Error(err.Error())
	}

	return output
}

func formatMaps(outputFormat OutputFormat, forceList bool, objects []map[string]interface{}) (string, error) {

	switch outputFormat.output {
	case FlagOutputNone:
		return formatObjectsInNone(outputFormat, objects...)

	case FlagOutputTable:
		if len(objects) == 1 && !forceList {
			return formatSingleObjectInTable(outputFormat, objects[0])
		}
		return formatObjectsInTable(outputFormat, objects)

	case FlagOutputJSON:
		return formatObjectsWithMarshaler(outputFormat, objects, prettyjson.Marshal)

	case FlagOutputYAML:
		return formatObjectsWithMarshaler(outputFormat, objects, yaml.Marshal)

	case FlagOutputTemplate:
		if len(objects) == 1 && !forceList {
			return FormatSingleObjectWithTemplate(objects[0], outputFormat.template)
		}
		return FormatObjectsWithTemplate(objects, outputFormat.template)

	default:
		panic(fmt.Sprintf("invalid output format '%s'", outputFormat))
	}
}

func formatObjectsInNone(outputFormat OutputFormat, objects ...map[string]interface{}) (string, error) {

	var ids []string // nolint: prealloc

	for _, object := range objects {
		if outputFormat.formatType == FormatTypeCount {
			return fmt.Sprintf("%d", object[FormatTypeCount]), nil
		}

		id, ok := object["ID"].(string)
		if !ok {
			continue
		}

		ids = append(ids, id)
	}

	return strings.Join(ids, "\n"), nil
}

func formatObjectsWithMarshaler(outputFormat OutputFormat, objects []map[string]interface{}, marshal func(interface{}) ([]byte, error)) (string, error) {

	if len(objects) == 0 {
		if outputFormat.output == FlagOutputJSON {
			if outputFormat.formatType == FormatTypeHash || outputFormat.formatType == "" {
				return "{}", nil
			}
			return "[]", nil
		}
		return "", nil
	}

	var target interface{}
	// Print objects as an array only when multiple objects are to be printed
	if len(objects) == 1 && (outputFormat.formatType == FormatTypeHash || outputFormat.formatType == FormatTypeCount || outputFormat.formatType == "") {
		target = objects[0]
	} else {
		target = objects
	}

	output, err := marshal(target)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// FormatObjectsWithTemplate formats the given []map[string]interface{} using given template.
func FormatObjectsWithTemplate(obj []map[string]interface{}, tpl string) (string, error) {

	t, err := template.New("tpl").Parse(tpl)
	if err != nil {
		return "", fmt.Errorf("unable to parse template: %s", err)
	}

	buffer := bytes.NewBuffer(nil)
	if err := t.Execute(buffer, obj); err != nil {
		return "", fmt.Errorf("unable to execute template: %s", err)
	}

	return buffer.String(), nil
}

// FormatSingleObjectWithTemplate formats the given map[string]interface{} using given template.
func FormatSingleObjectWithTemplate(obj map[string]interface{}, tpl string) (string, error) {

	t, err := template.New("tpl").Parse(tpl)
	if err != nil {
		return "", fmt.Errorf("unable to parse template: %s", err)
	}

	buffer := bytes.NewBuffer(nil)
	if err := t.Execute(buffer, obj); err != nil {
		return "", fmt.Errorf("unable to execute template: %s", err)
	}

	return buffer.String(), nil
}

// listFields validates the given "columns" arg with the given object, and
// returns the validated column list.
// If `columns` is empty, it lists all fields from the object and makes "ID"
// frontmost if found.
// If `columns` is non-empty, it filters the columns with what are found in the
// object, keeping the order.
func listFields(object map[string]interface{}, columns []string) []string {

	if len(columns) == 0 {

		keys := make([]string, 0, len(object))

		for k := range object {
			if k != "ID" {
				keys = append(keys, k)
			}
		}

		sort.Strings(keys)

		// Keep "ID" the first field if possible
		if _, ok := object["ID"]; ok {
			keys = append([]string{"ID"}, keys...)
		}

		return keys
	}

	keys := make([]string, 0, len(columns))
	for _, k := range columns {
		if _, ok := object[k]; ok {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)
	return keys
}

func tabulate(header []string, rows [][]string, single bool, caption string) string {

	out := &bytes.Buffer{}

	// colors := make([]tablewriter.Colors, len(header))
	// for i := 0; i < len(header); i++ {
	// 	colors[i] = tablewriter.Color(tablewriter.FgCyanColor, tablewriter.Bold)
	// }

	table := tablewriter.NewWriter(out)
	table.SetHeader(header)
	table.AppendBulk(rows)
	table.SetAutoFormatHeaders(false)
	table.SetHeaderLine(true)
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	// table.SetHeaderColor(colors...)

	if single {
		table.SetColumnAlignment([]int{tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_LEFT})
	}

	if caption != "" {
		table.SetCaption(true, caption)
	}

	table.Render()

	return "\n" + out.String()
}

func formatSingleObjectInTable(outputFormat OutputFormat, object map[string]interface{}) (string, error) {

	fields := listFields(object, outputFormat.columns)
	data := make([][]string, len(fields))

	for fieldIdx, field := range fields {
		data[fieldIdx] = []string{field, fmt.Sprintf("%v", object[field])}
	}

	return tabulate([]string{"property", "value"}, data, true, outputFormat.tableCaption), nil
}

func formatObjectsInTable(outputFormat OutputFormat, objects []map[string]interface{}) (string, error) {

	if len(objects) == 0 {
		return "", nil
	}

	fields := listFields(objects[0], outputFormat.columns)
	data := make([][]string, len(objects))

	for objectIdx, object := range objects {
		objectData := make([]string, len(fields))
		for fieldIdx, field := range fields {
			objectData[fieldIdx] = fmt.Sprintf("%v", object[field])
		}

		data[objectIdx] = objectData
	}

	return tabulate(fields, data, false, outputFormat.tableCaption), nil
}

// Remarshal marshals an object into a JSON string, and unmarshal it back to the
// given object. If object is not given, a new map[string]interface{} is used.
// The object is returned if no error occurs; otherwise nil with the error.
func Remarshal(object interface{}, target interface{}) error {

	if target == nil {
		return fmt.Errorf("Unable to call remarshall on a nil target")
	}

	if encodable, ok := object.(elemental.Encodable); ok {
		buf, err := elemental.Encode(encodable.GetEncoding(), object)
		if err != nil {
			return err
		}
		return elemental.Decode(encodable.GetEncoding(), buf, target)

	}

	buf, err := elemental.Encode(elemental.EncodingTypeMSGPACK, object)
	if err != nil {
		return err
	}

	return elemental.Decode(elemental.EncodingTypeMSGPACK, buf, target)
}
