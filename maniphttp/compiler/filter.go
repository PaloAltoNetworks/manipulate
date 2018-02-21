package compiler

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/aporeto-inc/manipulate"
)

// CompileFilter compiles the given filter into a cassandra filter.
func CompileFilter(f *manipulate.Filter) (url.Values, error) {

	ret := url.Values{}

	for index, key := range f.Keys() {

		if len(f.Values()[index]) != 1 {
			return nil, fmt.Errorf("Invalid filter. Only single filter value is supported")
		}

		if f.Operators()[index] != manipulate.AndOperator && f.Operators()[index] != manipulate.InitialOperator {
			return nil, fmt.Errorf("Invalid filter. Only AND operator is supported %v", f.Operators()[index])
		}

		switch f.Comparators()[index] {

		case manipulate.EqualComparator:
			ret.Add("tag", fmt.Sprintf("$%s=%v", strings.ToLower(key), f.Values()[index][0]))

		case manipulate.ContainComparator:
			value := fmt.Sprintf("%s", f.Values()[index][0])
			parts := strings.SplitN(value, "=", 2)

			if len(parts) != 2 {
				return nil, fmt.Errorf("Invalid filter. When using Contain, you must provide a valid tag format as value")
			}

			ret.Add("tag", value)

		default:
			return nil, fmt.Errorf("Invalid filter. Only Equals or Contains comparator is supported")
		}
	}

	return ret, nil
}
