// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maniphttp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.aporeto.io/elemental"
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

	if p := ctx.After(); p != "" {
		q.Add("after", p)
	}

	if l := ctx.Limit(); l > 0 {
		q.Add("limit", strconv.Itoa(l))
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

	var data []byte
	if data, err = ioutil.ReadAll(r.Body); err != nil {
		return manipulate.NewErrCannotUnmarshal(fmt.Sprintf("unable to read data: %s", err.Error()))
	}

	encoding := elemental.EncodingTypeJSON
	if r.Header.Get("Content-Type") != "" {
		encoding, _, err = elemental.EncodingFromHeaders(r.Header)
		if err != nil {
			return elemental.NewErrors(err)
		}
	}

	if err = elemental.Decode(encoding, data, dest); err != nil {
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

func getDefaultTransport(url string) (*http.Transport, string) {

	dialer := (&net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext

	outURL := url
	isUnix := strings.HasPrefix(url, "unix://")
	isUnixTLS := strings.HasPrefix(url, "unixs://")

	if isUnix || isUnixTLS {

		if isUnixTLS {
			outURL = "https://localhost"
		} else {
			outURL = "http://localhost"
		}

		dialer = func(context.Context, string, string) (net.Conn, error) {
			return net.Dial("unix", strings.TrimPrefix(strings.TrimPrefix(url, "unix://"), "unixs://"))
		}
	}

	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   50,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}, outURL
}

func getDefaultClient() *http.Client {
	return &http.Client{
		Timeout: 0, // we manage timeouts with contexts only.
	}
}
