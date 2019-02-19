package manipvortex

import (
	"context"

	"go.aporeto.io/manipulate"
)

// A BufferedManipulator is a Manipulator with a local cache
type BufferedManipulator interface {

	// ReSync forces a complete resync of the cache.
	ReSync(context.Context) error

	// Flush flushes and empties the cache.
	Flush(ctx context.Context) error

	manipulate.Manipulator
}
