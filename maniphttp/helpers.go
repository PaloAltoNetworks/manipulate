package maniphttp

import (
	"net/http"

	"github.com/aporeto-inc/manipulate"
)

// AddQueryParameters appends each key-value pair from ctx.Parameters.KeyValues
// to a request as query parameters with proper escaping.
func addQueryParameters(req *http.Request, ctx *manipulate.Context) {
	if req == nil || ctx == nil || ctx.Parameters == nil || ctx.Parameters.KeyValues == nil {
		return
	}
	keyValues := ctx.Parameters.KeyValues
	q := req.URL.Query()
	for k, v := range keyValues {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
}
