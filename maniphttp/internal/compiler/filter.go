package compiler

import (
	"net/url"

	"go.aporeto.io/manipulate"
)

// CompileFilter compiles the given filter into a http query filter.
func CompileFilter(f *manipulate.Filter) (url.Values, error) {
	return url.Values{"q": []string{f.String()}}, nil
}
