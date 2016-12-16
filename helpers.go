package manipulate

import (
	"bytes"
	"net/http"
	"reflect"
)

// ConvertArrayToManipulables convert the given array of interface into an array of Manipulable
func ConvertArrayToManipulables(i interface{}) []Manipulable {

	var manipulables []Manipulable
	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			manipulables = append(manipulables, val.Index(i).Interface().(Manipulable))
		}
	}

	return manipulables
}

// WriteString is a wrapper to buffer.WriteString that panics in
// case of write error.
func WriteString(buffer *bytes.Buffer, str string) {
	if _, err := buffer.WriteString(str); err != nil {
		panic(err)
	}
}

// AddQueryParameters appends each key-value pair from keyValues to a request
// as query parameters with proper escaping.
func AddQueryParameters(req *http.Request, keyValues map[string]string) {
	if req == nil || keyValues == nil {
		return
	}
	q := req.URL.Query()
	for k, v := range keyValues {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
}
