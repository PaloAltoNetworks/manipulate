package maniphttp

import (
	"context"
	"crypto/tls"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testTokenManager struct{}

func (t *testTokenManager) Issue(context.Context) (string, error)        { return "", nil }
func (t *testTokenManager) Run(ctx context.Context, tokenCh chan string) {}

func TestManipHttp_Options(t *testing.T) {

	Convey("Calling OptCredentials should work", t, func() {
		m := &httpManipulator{}
		OptCredentials("user", "password")(m)
		So(m.username, ShouldEqual, "user")
		So(m.password, ShouldEqual, "password")
	})

	Convey("Calling OptNamespace should work", t, func() {
		m := &httpManipulator{}
		OptNamespace("ns")(m)
		So(m.namespace, ShouldEqual, "ns")
	})

	Convey("Calling OptToken should work", t, func() {
		m := &httpManipulator{}
		OptToken("token")(m)
		So(m.username, ShouldEqual, "Bearer")
		So(m.password, ShouldEqual, "token")
	})

	Convey("Calling OptTokenManager should work", t, func() {
		m := &httpManipulator{}
		tm := &testTokenManager{}
		OptTokenManager(tm)(m)
		So(m.tokenManager, ShouldEqual, tm)
	})

	Convey("Calling OptHTTPClient should work", t, func() {
		m := &httpManipulator{}
		c := &http.Client{}
		OptHTTPClient(c)(m)
		So(m.client, ShouldEqual, c)
	})

	Convey("Calling OptHTTPTransport should work", t, func() {
		m := &httpManipulator{}
		t := &http.Transport{}
		OptHTTPTransport(t)(m)
		So(m.transport, ShouldEqual, t)
	})

	Convey("Calling OptTLSConfig should work", t, func() {
		m := &httpManipulator{}
		cfg := &tls.Config{}
		OptTLSConfig(cfg)(m)
		So(m.tlsConfig, ShouldEqual, cfg)
	})

	Convey("Calling OptAdditonalHeaders should work", t, func() {
		m := &httpManipulator{}
		h := http.Header{}
		OptAdditonalHeaders(h)(m)
		So(m.globalHeaders, ShouldEqual, h)
	})
}
