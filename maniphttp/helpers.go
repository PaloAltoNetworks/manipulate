package maniphttp

import (
	"net/http"
	"strconv"

	"github.com/aporeto-inc/manipulate"
)

// AddQueryParameters appends each key-value pair from ctx.Parameters.KeyValues
// to a request as query parameters with proper escaping.
func addQueryParameters(req *http.Request, ctx *manipulate.Context) {

	q := req.URL.Query()

	if ctx.Parameters != nil && ctx.Parameters.KeyValues != nil {
		keyValues := ctx.Parameters.KeyValues
		for k, v := range keyValues {
			q.Add(k, v)
		}
	}

	if ctx.Page != 0 {
		q.Add("page", strconv.Itoa(ctx.Page))
	}

	if ctx.PageSize > 0 {
		q.Add("pagesize", strconv.Itoa(ctx.PageSize))
	}

	if ctx.Recursive {
		q.Add("recursive", "true")
	}

	req.URL.RawQuery = q.Encode()
}
