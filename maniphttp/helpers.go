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
	"fmt"
	"net/http"
	"strings"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/tracing"
)

// ExtractCredentials extracts the username and password from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractCredentials(manipulator manipulate.Manipulator) (string, string) {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractCredentials")
	}

	m.renewLock.Lock()
	u, p := m.username, m.password
	m.renewLock.Unlock()

	return u, p
}

// ExtractEndpoint extracts the endpoint url from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractEndpoint(manipulator manipulate.Manipulator) string {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractEndpoint")
	}

	return m.url
}

// ExtractNamespace extracts the default namespace from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractNamespace(manipulator manipulate.Manipulator) string {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractNamespace")
	}

	return m.namespace
}

// ExtractTLSConfig returns a copy of the tls config from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractTLSConfig(manipulator manipulate.Manipulator) *tls.Config {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractTLSConfig")
	}

	return &tls.Config{
		RootCAs:            m.tlsConfig.RootCAs,
		Certificates:       m.tlsConfig.Certificates,
		InsecureSkipVerify: m.tlsConfig.InsecureSkipVerify,
		ClientSessionCache: m.tlsConfig.ClientSessionCache,
	} // #nosec
}

// ExtractEncoding returns the encoding used by the given manipulator.
func ExtractEncoding(manipulator manipulate.Manipulator) elemental.EncodingType {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractEncoding")
	}

	return m.encoding
}

// SetGlobalHeaders sets the given headers to all requests that will be sent.
func SetGlobalHeaders(manipulator manipulate.Manipulator, headers http.Header) {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to SetGlobalHeaders")
	}

	m.globalHeaders = headers
}

// DirectSend allows to send direct bytes using the given manipulator.
// This is only useful in extremely particular scenario, like fuzzing.
func DirectSend(manipulator manipulate.Manipulator, mctx manipulate.Context, endpoint string, method string, body []byte) (*http.Response, error) {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to DirectSend")
	}

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	v := m.computeVersion(0, mctx.Version())
	url := m.url + strings.Replace("/"+v+endpoint, "//", "/", -1)

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.directsend"))
	defer sp.Finish()

	return m.send(mctx, method, url, body, nil, sp)
}
