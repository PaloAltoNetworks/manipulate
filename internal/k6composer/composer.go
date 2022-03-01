package k6composer

import (
	"bytes"
	"encoding/base64"
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
		"import encoding from 'k6/encoding';",
		"",
		"export default function () {",
		"",
	}

	var err error
	if k6file, err = os.OpenFile("k6-script.js", os.O_CREATE|os.O_WRONLY, 0755); err != nil {
		panic(errors.Wrap(err, "failed to create k6 script"))
	}

	if err := k6Writer(strings.Join(basestr, "\n")); err != nil {
		panic(err)
	}

}

// k6Writer setups the k6 script for writing.
func k6Writer(content string) error {

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
	reqGroup := []string{}

	apiPath, err := decodedURI(url)
	if err != nil {
		return errors.Wrap(err, "failed to parse url")
	}

	var line string
	reqGroup = append(reqGroup,
		fmt.Sprintf("\n\tgroup('%s %s', function () {", method, apiPath))

	checker := []string{
		"\tcheck(response,",
		"\t\t{ 'status is 200': (response) => response.status === 200, }",
		"\t);",
		"\tif (response.status !== 200) {",
		"\t\tconsole.error(`${response.request.method} ${response.request.url},",
		"\t\tStatus ${response.status},",
		"\t\tError code: ${response.error_code},",
		"\t\tError Msg: ${response.error},",
		"\t\tBody: ${JSON.stringify(response.body)},",
		"\t`);}",
	}

	line, err = k6Headers(req.Header)
	if err != nil {
		return errors.Wrap(err, "failed to write headers to k6 script")
	}
	reqGroup = append(reqGroup, line)

	lines := []string{}
	switch method {

	case "GET":
		lines = []string{
			fmt.Sprintf("\n\tvar response = http.get('%s', params);", url),
		}

	case "POST":
		encodedData, err := b64Body(req)
		if err != nil {
			return errors.Wrap(err, "failed to encode body")
		}

		lines = []string{
			fmt.Sprintf("\n\tvar postdata = '%s'", encodedData),
			fmt.Sprintf("\tvar response = http.post('%s',", url),
			"encoding.b64decode(postdata, 'std', ''),",
			"params);",
		}

	case "DELETE":
		lines = []string{
			fmt.Sprintf("\n\tvar response = http.del('%s'", url),
			", null, params);",
		}

	case "PATCH":

		encodedData, err := b64Body(req)
		if err != nil {
			return errors.Wrap(err, "failed to encode body in PATCH")
		}

		lines = []string{
			fmt.Sprintf("\n\tvar patchdata = '%s'", encodedData),
			fmt.Sprintf("\tvar response = http.patch('%s',", url),
			"encoding.b64decode(patchdata, 'std', ''),",
			"params);",
		}

	default:
		fmt.Printf("unhandled request for method %s", method)
	}

	// Add checkers.
	lines = append(lines, strings.Join(checker, "\n"))

	// Close the k6 request group.
	lines = append(lines, "});")
	reqGroup = append(reqGroup, strings.Join(lines, "\n"))

	fmt.Println(strings.Join(reqGroup, ""))
	fmt.Printf("\n====================\n")

	if err := k6Writer(strings.Join(reqGroup, "")); err != nil {
		return errors.Wrap(err, "failed write the request group")
	}

	return nil
}

// k6Headers extracts the headers from the request.
func k6Headers(headers http.Header) (string, error) {

	hdrs, err := json.MarshalIndent(headers, "", "\t\t")
	if err != nil {
		return "", errors.Wrap(err, "failed to extract headers")
	}

	return fmt.Sprintf("\nvar params = {\n\theaders:%s\n}\n", string(hdrs)), nil
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

//b64Body encodes the request body in base64 encoding.
func b64Body(req *http.Request) (string, error) {

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	// Re-write the body to original request object.
	req.Body = ioutil.NopCloser(bytes.NewReader(data))

	return base64.StdEncoding.EncodeToString(data), nil
}
