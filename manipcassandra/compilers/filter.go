package compilers

import (
	"bytes"
	"reflect"
	"strings"

	"github.com/aporeto-inc/manipulate"
)

// CompileFilter compiles the given filter into a cassandra filter.
func CompileFilter(f *manipulate.Filter) string {

	var buffer bytes.Buffer

	for index, key := range f.Keys() {

		buffer.WriteString(translateOperator(f.Operators()[index]))

		var keyValue string

		if len(key) == 1 {
			keyValue = key[0]
		} else {
			keyValue = "(" + strings.Join(key, ",") + ")"
		}

		var param string

		if len(f.Values()[index]) > 1 || f.Comparators()[index] == manipulate.InComparator {
			param = paramForValues(f.Values()[index])
		}

		buffer.WriteString(" ")
		buffer.WriteString(keyValue)
		buffer.WriteString(" ")
		buffer.WriteString(translateComparator(f.Comparators()[index]))

		if param == "" {
			buffer.WriteString(" ?")
		} else {
			buffer.WriteString(" ")
			buffer.WriteString(param)
		}
	}

	return buffer.String()
}

func translateComparator(comparator manipulate.FilterComparator) string {

	switch comparator {
	case manipulate.EqualComparator:
		return "="
	case manipulate.GreaterComparator:
		return ">="
	case manipulate.LesserComparator:
		return "<="
	case manipulate.InComparator:
		return "IN"
	case manipulate.ContainComparator:
		return "CONTAINS"
	}

	return ""
}

func translateOperator(operator manipulate.FilterOperator) string {

	switch operator {
	case manipulate.InitialOperator:
		return "WHERE"
	case manipulate.AndOperator:
		return " AND"
	case manipulate.OrOperator:
		return " OR"
	}

	return ""
}

func paramForValues(v []interface{}) string {

	var buffer bytes.Buffer
	buffer.WriteString("(")

	for i := 0; i < len(v); i++ {

		value := v[i]

		if reflect.ValueOf(value).Kind() == reflect.Array || reflect.ValueOf(value).Kind() == reflect.Slice {
			buffer.WriteString(paramForValues(value.([]interface{})))
		} else {
			buffer.WriteString("?")
		}

		if i < len(v)-1 {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString(")")

	return buffer.String()
}
