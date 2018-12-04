package manipulate

import (
	"context"
	"fmt"
	"math"
	"time"

	"go.uber.org/zap"
)

const maxBackoff = 8000

// Retry will retry the given function that performs a manipulate operation if it fails and the error is
// a manipulate.ErrCannotCommunicate.
//
// It will retry with exponential backoff (up to 8s) until the provided context is canceled.
//
// If the onRetryFunc is not nil and returns an error, the retrying process will be interrupted and
// manipulate.Retry will return the provided error.
// The retry function gets the retry number and the error produced by manipulateFunc.
//
// Example:
//
//      ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//      defer cancel()
//
//      if err := manipulate.Retry(
//          ctx,
//          func() error {
//              return m.Create(nil, obj)
//          },
//          func(t int, e error) error {
//              if _, ok := e.(manipulate.ErrLocked); ok {
//                  return errors.New("nah, I don't wanna retry")
//              }
//              return nil
//          },
//      ); err != nil {
//          // do interesting stuff.
//      }
func Retry(ctx context.Context, manipulateFunc func() error, onRetryFunc func(int, error) error) error {

	var try int
	var err, userErr error

	for {

		try++

		err = manipulateFunc()
		if err == nil {
			zap.L().Info("Retry success", zap.Int("attempt", try))
			return nil
		}

		switch err.(type) {

		// If this is a ErrCannotCommunicate or ErrLocked, we retry until the context is canceled.
		case ErrCannotCommunicate, ErrLocked, ErrTooManyRequests:

			// If onRetryFunc is provided we call it and decide what to do.
			if onRetryFunc != nil {
				if userErr = onRetryFunc(try, err); userErr != nil {
					return userErr
				}
			}

			d := nextBackoff(try)
			zap.L().Info("Retry failed...retrying", zap.Int("attempt", try), zap.Error(err), zap.Duration("backoff", d))

			// Otherwise we wait.
			select {
			case <-time.After(d):
			case <-ctx.Done():
				return NewErrDisconnected(fmt.Sprintf("interupted by context: %s. original error: %s", ctx.Err(), err))
			}

		// If it's any other kind of error, we do nothing and we return the error.
		default:
			zap.L().Info("Retry failed...end", zap.Int("attempt", try), zap.Error(err))
			return err
		}
	}
}

func nextBackoff(try int) time.Duration {

	return time.Duration(math.Min(math.Pow(2.0, float64(try))-1, maxBackoff)) * time.Millisecond
}
