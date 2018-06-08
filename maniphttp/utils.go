package maniphttp

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniphttp/internal/compiler"
)

// AddQueryParameters appends each key-value pair from ctx.Parameters
// to a request as query parameters with proper escaping.
func addQueryParameters(req *http.Request, ctx *manipulate.Context) error {

	q := req.URL.Query()

	if ctx.Filter != nil {
		query, err := compiler.CompileFilter(ctx.Filter)
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

func decodeData(dataReader io.Reader, dest interface{}) (err error) {

	var data []byte

	if dataReader == nil {
		return manipulate.NewErrCannotUnmarshal("nil reader")
	}

	if data, err = ioutil.ReadAll(dataReader); err != nil {
		return manipulate.NewErrCannotUnmarshal(fmt.Sprintf("unable to read data: %s", err.Error()))
	}

	if err = json.Unmarshal(data, dest); err != nil {
		return manipulate.NewErrCannotUnmarshal(fmt.Sprintf("%s. original data:\n%s", err.Error(), string(data)))
	}

	return nil
}
