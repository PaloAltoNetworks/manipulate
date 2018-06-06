package push

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
)

func decodeErrors(r io.Reader) error {

	es := []elemental.Error{}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return manipulate.NewErrCannotUnmarshal(fmt.Sprintf("%s: %s", err.Error(), string(data)))
	}

	if err := json.Unmarshal(data, &es); err != nil {
		return manipulate.NewErrCannotUnmarshal(fmt.Sprintf("%s: %s", err.Error(), string(data)))
	}

	errs := elemental.NewErrors()
	for _, e := range es {
		errs = append(errs, e)
	}

	return errs
}

func makeURL(u string, namespace string, password string, recursive bool) string {

	u = strings.Replace(u, "http://", "ws://", 1)
	u = strings.Replace(u, "https://", "wss://", 1)
	u = fmt.Sprintf("%s?token=%s", u, password)

	if namespace != "" {
		u += "&namespace=" + url.QueryEscape(namespace)
	}

	if recursive {
		u = u + "&mode=all"
	}

	return u
}

func isCommError(resp *http.Response) bool {

	if resp == nil {
		return true
	}

	switch resp.StatusCode {
	case http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		return true
	default:
		return false
	}
}
