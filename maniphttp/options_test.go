package maniphttp

import (
	"context"
	"crypto/tls"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
)

type testTokenManager struct{}

func (t *testTokenManager) Issue(context.Context) (string, error)        { return "", nil }
func (t *testTokenManager) Run(ctx context.Context, tokenCh chan string) {}

func TestManipHttp_Optionions(t *testing.T) {

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
}
