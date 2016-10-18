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

		manipulate.WriteString(&buffer, translateOperator(f.Operators()[index]))

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

		manipulate.WriteString(&buffer, " ")
		manipulate.WriteString(&buffer, keyValue)
		manipulate.WriteString(&buffer, " ")
		manipulate.WriteString(&buffer, translateComparator(f.Comparators()[index]))

		if param == "" {
			manipulate.WriteString(&buffer, " ?")
		} else {
			manipulate.WriteString(&buffer, " ")
			manipulate.WriteString(&buffer, param)
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
	manipulate.WriteString(&buffer, "(")

	for i := 0; i < len(v); i++ {

		value := v[i]

		if reflect.ValueOf(value).Kind() == reflect.Array || reflect.ValueOf(value).Kind() == reflect.Slice {
			manipulate.WriteString(&buffer, paramForValues(value.([]interface{})))
		} else {
			manipulate.WriteString(&buffer, "?")
		}

		if i < len(v)-1 {
			manipulate.WriteString(&buffer, ",")
		}
	}

	manipulate.WriteString(&buffer, ")")

	return buffer.String()
}
