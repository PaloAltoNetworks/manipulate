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
	"bytes"
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
// Note: the given manipulator must be an HTTP Manipulator or it will panic.
func ExtractEndpoint(manipulator manipulate.Manipulator) string {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractEndpoint")
	}

	return m.url
}

// ExtractNamespace extracts the default namespace from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will panic.
func ExtractNamespace(manipulator manipulate.Manipulator) string {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractNamespace")
	}

	return m.namespace
}

// ExtractTLSConfig returns a copy of the tls config from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will panic.
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
// Note: the given manipulator must be an HTTP Manipulator or it will panic.
func ExtractEncoding(manipulator manipulate.Manipulator) elemental.EncodingType {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractEncoding")
	}

	return m.encoding
}

// SetGlobalHeaders sets the given headers to all requests that will be sent.
// Note: the given manipulator must be an HTTP Manipulator or it will panic.
func SetGlobalHeaders(manipulator manipulate.Manipulator, headers http.Header) {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to SetGlobalHeaders")
	}

	m.globalHeaders = headers
}

// DirectSend allows to send direct bytes using the given manipulator.
// This is only useful in extremely particular scenario, like fuzzing.
// Note: the given manipulator must be an HTTP Manipulator or it will panic.
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

	return m.send(mctx, method, url, bytes.NewReader(body), nil, sp)
}

// BatchCreate is an experimental feature that may eventually be incorporated in the standard interface.
//
// But for now it is not recommended to use it unless you know exactly how
// this works on the server side. This API is NOT considered as stable and may break
// at any time.
func BatchCreate(manipulator manipulate.Manipulator, mctx manipulate.Context, objects ...elemental.Identifiable) (*http.Response, error) {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to BatchCreate")
	}

	if len(objects) == 0 {
		panic("You must pass at least one object to BatchCreate")
	}

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	var encoding elemental.EncodingType
	opaque := mctx.(opaquer).Opaque()

	if v, ok := opaque[opaqueKeyOverrideHeaderContentType]; ok {
		encoding = elemental.EncodingType(v.(string))
	} else {
		encoding = m.encoding
	}

	url := m.getGeneralURL(objects[0], mctx.Version())

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.batchcreate"))
	defer sp.Finish()

	body := bytes.NewBuffer(nil)

	encoder, close := elemental.MakeStreamEncoder(encoding, body)
	defer close()

	for _, o := range objects {
		if err := encoder(o); err != nil {
			return nil, err
		}
	}

	// Whatever is encoding and where it is coming from,
	// we always set a custom encoding by suffixing with '+batch'.
	opaque[opaqueKeyOverrideHeaderContentType] = string(encoding) + "+batch"

	return m.send(mctx, http.MethodPost, url, bytes.NewReader(body.Bytes()), nil, sp)
}
