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
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniphttp/internal/compiler"
	"go.aporeto.io/manipulate/maniphttp/internal/syscall"
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

	if ctx.Propagated() {
		q.Add("propagated", "true")
	}

	req.URL.RawQuery = q.Encode()

	return nil
}

func decodeData(r *http.Response, dest any) (err error) {

	if r.Body == nil {
		return manipulate.ErrCannotUnmarshal{Err: fmt.Errorf("nil reader")}
	}

	var data []byte
	if data, err = io.ReadAll(r.Body); err != nil {
		return manipulate.ErrCannotUnmarshal{Err: fmt.Errorf("unable to read data: %w", err)}
	}

	encoding := elemental.EncodingTypeJSON
	if r.Header.Get("Content-Type") != "" {
		encoding, _, err = elemental.EncodingFromHeaders(r.Header)
		if err != nil {
			return elemental.NewErrors(err)
		}
	}

	if err = elemental.Decode(encoding, data, dest); err != nil {
		return manipulate.ErrCannotUnmarshal{Err: fmt.Errorf("%w. original data:\n%s", err, string(data))}
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

	return &tls.Config{
		RootCAs:            systemCertPool,
		ClientSessionCache: tls.NewLRUClientSessionCache(0),
	}
}

func getDefaultHTTPTransport(url string, disableCompression bool, tcpUserTimeout time.Duration) (*http.Transport, string) {

	dialer := (&net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 30 * time.Second,
		Control:   syscall.MakeDialerControlFunc(tcpUserTimeout),
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
		ForceAttemptHTTP2:     true,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer,
		MaxConnsPerHost:       32,
		MaxIdleConns:          32,
		MaxIdleConnsPerHost:   32,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableCompression:    disableCompression,
	}, outURL
}

func getDefaultClient() *http.Client {
	return &http.Client{
		Timeout: 0, // we manage timeouts with contexts only.
	}
}
