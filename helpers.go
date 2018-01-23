package manipulate

import (
	"bytes"
	"context"
	"time"

	"go.uber.org/zap"
)

func writeString(buffer *bytes.Buffer, str string) {
	if _, err := buffer.WriteString(str); err != nil {
		panic(err)
	}
}

// Retry will retry the given function that tries a manipulate
// operation at least maxTries if the error is a manipulate.ErrCannotCommunicate.
// You can pass -1 to always retry. The function will retry immediately the first try,
// then after 1s, 2s etc until a try every 5s. If the onRetryFunc is passed and it returns false,
// the retrying process will be interrupted.
func Retry(ctx context.Context, manipulation func() error, onRetryFunc func(int, error) bool, maxTries int) error {

	var try int
	var waitTime time.Duration

	for {

		err := manipulation()

		// If the error is nil, its a success.
		if err == nil {
			break
		}

		switch err.(type) {

		// If this is a ErrLocked, we retry forever and we don't give the client any choice.
		// but to eventually die.
		case ErrLocked:

			zap.L().Warn("API locked. Retrying in 10s", zap.Error(err))
			select {
			case <-time.After(10 * time.Second):
			case <-ctx.Done():
				return NewErrDisconnected("Disconnected per signal")
			}

		// If this is a ErrCannotCommunicate, we retry until the maxTries.
		case ErrCannotCommunicate:

			// If we reach the maxtries, return the error
			if maxTries != -1 && try == maxTries {
				return err
			}

			if onRetryFunc != nil && !onRetryFunc(try, err) {
				return nil
			}

			// Otherwise wait, increase the time and try again.
			select {
			case <-time.After(waitTime):
			case <-ctx.Done():
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
