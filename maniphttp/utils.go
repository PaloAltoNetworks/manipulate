package maniphttp

import (
	"net/http"
	"strconv"

	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/sharedcompiler"
)

// AddQueryParameters appends each key-value pair from ctx.Parameters
// to a request as query parameters with proper escaping.
func addQueryParameters(req *http.Request, ctx *manipulate.Context) error {

	q := req.URL.Query()

	if ctx.Filter != nil {
		query, err := sharedcompiler.CompileFilter(ctx.Filter)
		if err != nil {
			return err
		}
		for k, v := range query {
			q[k] = v
		}
	}

	for k, v := range ctx.Parameters {
		q[k] = v
	}

	for _, order := range ctx.Order {
		q.Add("order", order)
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

	if ctx.OverrideProtection {
		q.Add("override", "true")
	}

	req.URL.RawQuery = q.Encode()

	return nil
}
