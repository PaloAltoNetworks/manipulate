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
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

type testTokenManager struct{}

func (t *testTokenManager) Issue(context.Context) (string, error)        { return "", nil }
func (t *testTokenManager) Run(ctx context.Context, tokenCh chan string) {}

func Test_Options(t *testing.T) {

	Convey("Calling OptionCredentials should work", t, func() {
		m := &httpManipulator{}
		OptionCredentials("user", "password")(m)
		So(m.username, ShouldEqual, "user")
		So(m.password, ShouldEqual, "password")
	})

	Convey("Calling OptionNamespace should work", t, func() {
		m := &httpManipulator{}
		OptionNamespace("ns")(m)
		So(m.namespace, ShouldEqual, "ns")
	})

	Convey("Calling OptionToken should work", t, func() {
		m := &httpManipulator{}
		OptionToken("token")(m)
		So(m.username, ShouldEqual, "Bearer")
		So(m.password, ShouldEqual, "token")
	})

	Convey("Calling OptionTokenManager should work", t, func() {
		m := &httpManipulator{}
		tm := &testTokenManager{}
		OptionTokenManager(tm)(m)
		So(m.tokenManager, ShouldEqual, tm)
	})

	Convey("Calling OptionHTTPClient should work", t, func() {
		m := &httpManipulator{}
		c := &http.Client{}
		OptionHTTPClient(c)(m)
		So(m.client, ShouldEqual, c)
	})

	Convey("Calling OptionHTTPTransport should work", t, func() {
		m := &httpManipulator{}
		t := &http.Transport{}
		OptionHTTPTransport(t)(m)
		So(m.transport, ShouldEqual, t)
	})

	Convey("Calling OptionTLSConfig should work", t, func() {
		m := &httpManipulator{}
		cfg := &tls.Config{}
		OptionTLSConfig(cfg)(m)
		So(m.tlsConfig, ShouldEqual, cfg)
	})

	Convey("Calling OptionAdditonalHeaders should work", t, func() {
		m := &httpManipulator{}
		h := http.Header{}
		OptionAdditonalHeaders(h)(m)
		So(m.globalHeaders, ShouldEqual, h)
	})

	Convey("Calling OptionDisableBuiltInRetry should work", t, func() {
		m := &httpManipulator{}
		OptionDisableBuiltInRetry()(m)
		So(m.disableAutoRetry, ShouldBeTrue)
	})

	Convey("Calling OptionEncoding should work", t, func() {
		m := &httpManipulator{}
		OptionEncoding(elemental.EncodingTypeMSGPACK)(m)
		So(m.encoding, ShouldEqual, elemental.EncodingTypeMSGPACK)
	})

	Convey("Calling OptionDefaultRetryFunc should work", t, func() {
		f := func(manipulate.RetryInfo) error { return nil }
		m := &httpManipulator{}
		OptionDefaultRetryFunc(f)(m)
		So(m.defaultRetryFunc, ShouldEqual, f)
	})

	Convey("Calling OptionDisableCompression should work", t, func() {
		m := &httpManipulator{}
		OptionDisableCompression()(m)
		So(m.disableCompression, ShouldEqual, true)
	})

	Convey("Calling OptionSimulateFailures should work", t, func() {
		m := &httpManipulator{}
		f := map[float64]error{}
		OptionSimulateFailures(f)(m)
		So(m.failureSimulations, ShouldEqual, f)
	})

	Convey("Calling OptionSendCredentialsAsCookie should work", t, func() {
		m := &httpManipulator{}
		OptionSendCredentialsAsCookie("x-token")(m)
		So(m.tokenCookieKey, ShouldEqual, "x-token")
	})

	Convey("Calling OptionForceAttemptHTTP2 should work", t, func() {
		m := &httpManipulator{}
		OptionForceAttemptHTTP2(true)(m)
		So(m.forceAttemptHTTP2, ShouldBeTrue)
	})
}
