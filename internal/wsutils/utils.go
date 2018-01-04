package wsutils

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/gorilla/websocket"
)

// DecodeErrors decodes the error in the given data.
func decodeErrors(r io.Reader) error {

	es := []elemental.Error{}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return manipulate.NewErrCannotUnmarshal(err.Error())
	}

	if err := json.Unmarshal(data, &es); err != nil {
		return manipulate.NewErrCannotUnmarshal(err.Error())
	}

	errs := elemental.NewErrors()
	for _, e := range es {
		errs = append(errs, e)
	}

	return errs
}

// MakeURL makes the websocket url from the given information.
func MakeURL(u string, endpoint string, namespace string, password string, recursive bool) string {

	u = strings.Replace(u, "http://", "ws://", 1)
	u = strings.Replace(u, "https://", "wss://", 1)
	u = fmt.Sprintf("%s/%s?token=%s", u, endpoint, password)

	if namespace != "" {
		u += "&namespace=" + url.QueryEscape(namespace)
	}

	if recursive {
		u = u + "&mode=all"
	}

	return u
}

// Dial returns a connected websocket.
func Dial(u string, tlsConfig *tls.Config) (*websocket.Conn, error) {

	dialer := &websocket.Dialer{
		Proxy:           http.ProxyFromEnvironment,
		TLSClientConfig: tlsConfig,
	}

	conn, resp, err := dialer.Dial(u, nil)
	// this is a com error.
	if err != nil && resp == nil {
		return nil, manipulate.NewErrCannotCommunicate(err.Error())
	}

	// this has been rejected for a reason. let's decode it.
	if err != nil {
		defer resp.Body.Close() // nolint: errcheck
		return nil, decodeErrors(resp.Body)
	}

	return conn, nil
}
