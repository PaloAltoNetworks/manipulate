package manipulate

import (
	"bytes"
	"reflect"
	"time"
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

// RetryManipulation will retry the given function that tries a manipulate
// operation at least maxTries if the error is a manipulate.ErrCannotCommunicate.
// You can pass -1 to always retry. The function will retry immediately the first try,
// then after 1s, 2s etc until a try every 5s.
func RetryManipulation(manipulation func() error, onRetryFunc func(int), maxTries int) error {

	try := 0
	waitTime := 0 * time.Second

	for {

		err := manipulation()

		// If the error is nil, its a success.
		if err == nil {
			break
		}

		// Check the type of the error.
		if _, ok := err.(ErrCannotCommunicate); !ok {
			return err
		}

		// If we reach the maxtries, return the error
		if maxTries != -1 && try == maxTries {
			return err
		}

		if onRetryFunc != nil {
			onRetryFunc(try)
		}

		// Otherwise wait, increase the time and try again.
		<-time.After(waitTime)

		if waitTime < 5*time.Second {
			waitTime += 1 * time.Second
		}

		try++
	}

	return nil
}
