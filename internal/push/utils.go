package push

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/url"
	"strings"
	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
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

const maxBackoff = 8000

func nextBackoff(try int) time.Duration {

	return time.Duration(math.Min(math.Pow(2.0, float64(try))-1, maxBackoff)) * time.Millisecond
}
