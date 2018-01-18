package manipulate

import (
	"bytes"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
)

func writeString(buffer *bytes.Buffer, str string) {
	if _, err := buffer.WriteString(str); err != nil {
		panic(err)
	}
}

// RetryManipulation will retry the given function that tries a manipulate
// operation at least maxTries if the error is a manipulate.ErrCannotCommunicate.
// You can pass -1 to always retry. The function will retry immediately the first try,
// then after 1s, 2s etc until a try every 5s.
//
// Deprecated: manipulate.RetryManipulation is deprecated. Please switch to manipulate.Retry instead.
func RetryManipulation(manipulation func() error, onRetryFunc func(int), maxTries int) error {

	zap.L().Warn("manipulate.RetryManipulation is deprecated. Please switch to manipulate.Retry")
	return retryManipulation(manipulation, onRetryFunc, nil, maxTries)
}

// Retry will retry the given function that tries a manipulate
// operation at least maxTries if the error is a manipulate.ErrCannotCommunicate.
// You can pass -1 to always retry. The function will retry immediately the first try,
// then after 1s, 2s etc until a try every 5s. If the onRetryFunc is passed and it returns false,
// the retrying process will be interrupted.
func Retry(manipulation func() error, onRetryFunc func(int, error) bool, maxTries int) error {

	return retryManipulation(manipulation, nil, onRetryFunc, maxTries)
}

func retryManipulation(manipulation func() error, onRetryFunc func(int), onRetryCheckFunc func(int, error) bool, maxTries int) error {

	var try int
	var waitTime time.Duration
	var c chan os.Signal

	for {

		err := manipulation()

		// If the error is nil, its a success.
		if err == nil {
			break
		}

		// install the quit handler.
		if c == nil {
			c = make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			defer func() {
				signal.Stop(c)
			}()
		}

		switch err.(type) {

		// If this is a ErrLocked, we retry forever and we don't give the client any choice.
		// but to eventually die.
		case ErrLocked:

			zap.L().Warn("API locked. Retrying in 10s", zap.Error(err))
			select {
			case <-time.After(10 * time.Second):
			case <-c:
				return NewErrDisconnected("Disconnected per signal")
			}

		// If this is a ErrCannotCommunicate, we retry until the maxTries.
		case ErrCannotCommunicate:

			// If we reach the maxtries, return the error
			if maxTries != -1 && try == maxTries {
				return err
			}

			if onRetryFunc != nil {
				onRetryFunc(try)
			}

			if onRetryCheckFunc != nil && !onRetryCheckFunc(try, err) {
				return nil
			}

			// Otherwise wait, increase the time and try again.
			select {
			case <-time.After(waitTime):
			case <-c:
				return NewErrDisconnected("Disconnected per signal")
			}

			if waitTime < 5*time.Second {
				waitTime += 1 * time.Second
			}

			try++
		default:
			return err
		}
	}

	return nil
}
