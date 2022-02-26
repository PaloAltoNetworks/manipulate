package k6composer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

var mu sync.Mutex
var k6file *os.File

func init() {

	basestr := []string{
		"import http from 'k6/http';",
		"import { group } from 'k6';",
		"import { check } from 'k6';",
		"",
		"export default function () {",
		"",
	}

	var err error
	if k6file, err = os.OpenFile("k6-script.js", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755); err != nil {
		panic(errors.Wrap(err, "failed to create k6 script"))
	}

	if err := K6Writer(strings.Join(basestr, "\n")); err != nil {
		panic(err)
	}
}

// K6Writer setups the k6 script for writing.
func K6Writer(content string) error {

	mu.Lock()
	defer mu.Unlock()

	if _, err := k6file.WriteString(content); err != nil {
		return errors.Wrap(err, "failed to write to k6 script")
	}

	// if err := k6file.Close(); err != nil {
	// 	return errors.Wrap(err, "failed to closed k6 script")
	// }

	return nil
}

// K6Copy copies the http request for scripting.
func K6Copy(req *http.Request) error {

	fmt.Printf("\n======== K6 dump =======\n")

	// Insert the k6 group request.
	url := req.URL
	method := req.Method

	apiPath, err := decodedURI(url)
	if err != nil {
		return errors.Wrap(err, "failed to parse url")
	}

	if err := K6Writer(
		fmt.Sprintf("\n\tgroup('%s %s', function () {",
			method,
			apiPath)); err != nil {
		return errors.Wrap(err, "failed to create request group")
	}

	// Get the headers.
	headers, err := json.MarshalIndent(req.Header, "", "\t\t")
	if err != nil {
		return err
	}
	fmt.Printf("\nK6-Headers:\n%s\n", string(headers))

	if err := k6Headers(headers); err != nil {
		return errors.Wrap(err, "failed to write headers to k6 script")
	}

	switch method {

	case "GET":
		if err := K6Writer(fmt.Sprintf("\n\tvar response = http.get('%s', params);\n\tcheck(response, { 'status is 200': (response) => response.status === 200, });", url)); err != nil {
			return errors.Wrap(err, "failed in k6 GET op")
		}

	case "POST":
		var data []byte
		data, err = io.ReadAll(req.Body)
		if err != nil {
			return err
		}
		fmt.Printf("\nK6-Body:\n%s\n", string(data))

		// Re-write the body to original request object.
		req.Body = ioutil.NopCloser(bytes.NewReader(data))
	}
	fmt.Printf("\n=========================\n\n")

	// Close the k6 request group.
	if err := K6Writer("\n});"); err != nil {
		return errors.Wrap(err, "failed to close request group")
	}

	return nil
}

// k6Headers extracts the headers from the request.
func k6Headers(headers []byte) error {

	if err := K6Writer(fmt.Sprintf("\nvar params = {\n\theaders:%s\n}\n", string(headers))); err != nil {
		return err
	}

	return nil
}

// decodedURL constructs human readable URI by parsing raw url.
func decodedURI(rawURL *url.URL) (string, error) {

	// Parse url for better readability.
	parsedURL, err := url.Parse(rawURL.String())
	if err != nil {
		return "", errors.Wrap(err, "failed to parse the url")
	}

	path, err := url.PathUnescape(parsedURL.Path)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse url path")
	}

	query, err := url.QueryUnescape(parsedURL.RawQuery)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse url query")
	}

	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}

	return path, nil
}
