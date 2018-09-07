package maniphttp

import (
	"compress/gzip"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniphttp/internal/compiler"
)

// AddQueryParameters appends each key-value pair from ctx.Parameters
// to a request as query parameters with proper escaping.
func addQueryParameters(req *http.Request, ctx manipulate.Context) error {

	q := req.URL.Query()

	if f := ctx.Filter(); f != nil {
		query, err := compiler.CompileFilter(f)
		if err != nil {
			return err
		}
		for k, v := range query {
			q[k] = v
		}
	}

	for k, v := range ctx.Parameters() {
		q[k] = v
	}

	for _, order := range ctx.Order() {
		q.Add("order", order)
	}

	if p := ctx.Page(); p != 0 {
		q.Add("page", strconv.Itoa(p))
	}

	if p := ctx.PageSize(); p > 0 {
		q.Add("pagesize", strconv.Itoa(p))
	}

	if ctx.Recursive() {
		q.Add("recursive", "true")
	}

	if ctx.Override() {
		q.Add("override", "true")
	}

	req.URL.RawQuery = q.Encode()

	return nil
}

func decodeData(r *http.Response, dest interface{}) (err error) {

	if r.Body == nil {
		return manipulate.NewErrCannotUnmarshal("nil reader")
	}

	var dataReader io.ReadCloser
	switch r.Header.Get("Content-Encoding") {
	case "gzip":
		dataReader, _ = gzip.NewReader(r.Body)
		defer dataReader.Close() // nolint
	default:
		dataReader = r.Body
	}

	var data []byte

	if data, err = ioutil.ReadAll(dataReader); err != nil {
		return manipulate.NewErrCannotUnmarshal(fmt.Sprintf("unable to read data: %s", err.Error()))
	}

	if err = json.Unmarshal(data, dest); err != nil {
		return manipulate.NewErrCannotUnmarshal(fmt.Sprintf("%s. original data:\n%s", err.Error(), string(data)))
	}

	return nil
}

var systemCertPoolLock sync.Mutex
var systemCertPool *x509.CertPool

func getSystemCertPool() (*x509.CertPool, error) {

	systemCertPoolLock.Lock()
	defer systemCertPoolLock.Unlock()

	if systemCertPool == nil {
		var err error

		if systemCertPool, err = x509.SystemCertPool(); err != nil {
			return nil, err
		}
	}

	return systemCertPool, nil
}
