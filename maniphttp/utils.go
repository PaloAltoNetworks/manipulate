package maniphttp

import (
	"compress/gzip"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniphttp/internal/compiler"
)

var gzipReaders = sync.Pool{New: func() interface{} { return &gzip.Reader{} }}

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
		gz := gzipReaders.Get().(*gzip.Reader)
		defer gzipReaders.Put(gz)

		if err := gz.Reset(r.Body); err != nil {
			panic(err)
		}
		defer gz.Close() // nolint

		dataReader = gz

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

func getDefaultTLSConfig() *tls.Config {

	systemCertPoolLock.Lock()
	defer systemCertPoolLock.Unlock()

	if systemCertPool == nil {
		var err error
		if systemCertPool, err = x509.SystemCertPool(); err != nil {
			panic(fmt.Sprintf("Unable to load system root cert pool: %s", err))
		}
	}

	return &tls.Config{RootCAs: systemCertPool}
}

func getDefaultTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func getDefaultClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}
