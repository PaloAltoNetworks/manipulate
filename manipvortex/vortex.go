package manipvortex

import (
	"context"

	"go.aporeto.io/manipulate"
)

// Cache is the Vortex interface for interacting with a memory cache.
type Cache interface {

	// ReSync forces a complete resync of the cache.
	ReSync(context.Context) error

	// Flush flushes and empties the cache.
	Flush(ctx context.Context) error

	// Run starts the cache.
	Run(ctx context.Context) error

	manipulate.Manipulator
	manipulate.Subscriber
}
