package k6composer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var k6file *os.File

// init setups the k6 script for writing.
func init() {
	var err error
	k6file, err = os.Create("k6-script.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer k6file.Close()

}

// K6Copy copies the http request for scripting.
func K6Copy(req *http.Request) error {

	headers, err := json.Marshal(req.Header)
	if err != nil {
		return err
	}

	fmt.Printf("\n======== K6 dump =======\n")
	fmt.Printf("\nK6-Headers:\n%s\n", string(headers))

	if req.Method == "POST" {
		var data []byte
		data, err = io.ReadAll(req.Body)
		if err != nil {
			return err
		}
		fmt.Printf("\nK6-Body:\n%s\n", string(data))

		// re-write the body to original request object.
		req.Body = ioutil.NopCloser(bytes.NewReader(data))
	}
	fmt.Printf("\n=========================\n\n")

	return nil
}
