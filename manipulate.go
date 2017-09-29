// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import (
	"time"

	"github.com/aporeto-inc/elemental"
)

// EventHandler is the type of event an handler.
type EventHandler func(event *elemental.Event, err error)

// RecoveryHandler is the type of connection recovery handler.
type RecoveryHandler func()

// EventUnsubscriber is the type of the unsubscribe function.
type EventUnsubscriber func() error

// EventFilterUpdater is the type of the function to update a PushFilter.
type EventFilterUpdater func(*elemental.PushFilter) error

// Manipulator is the interface of a storage backend.
type Manipulator interface {
	// RetrieveMany retrieves the a list of objects with the given elemental.Identity and put them in the given dest.
	RetrieveMany(context *Context, dest elemental.ContentIdentifiable) error

	// Retrieve retrieves one or multiple elemental.Identifiables. In order to be retrievable,
	// the elemental.Identifiable needs to have their Identifier correctly set.
	Retrieve(context *Context, objects ...elemental.Identifiable) error

	// Create creates a the given elemental.Identifiables in the given parent elemental.Identifiable.
	Create(context *Context, objects ...elemental.Identifiable) error

	// Update updates one or multiple elemental.Identifiables. In order to be updatable,
	// the elemental.Identifiable needs to have their Identifier correctly set.
	Update(context *Context, objects ...elemental.Identifiable) error

	// Delete deletes one or multiple elemental.Identifiables. In order to be deletable,
	// the elemental.Identifiable needs to have their Identifier correctly set.
	Delete(context *Context, objects ...elemental.Identifiable) error

	// DeleteMany deletes all objects of with the given identity or
	// all the ones matching the filtet in the given context.
	DeleteMany(context *Context, identity elemental.Identity) error

	// Count returns the number of objects with the given identity.
	Count(context *Context, identity elemental.Identity) (int, error)

	// Increment increments the given counter filter by the given increment for the given identity.
	// Filter should be set in the context to decide which element to increment.
	Increment(context *Context, identity elemental.Identity, counter string, inc int) error
}

// A TransactionalManipulator is a Manipulator that handles transactions.
type TransactionalManipulator interface {
	Manipulator

	// Commit commits the given TransactionID.
	Commit(id TransactionID) error

	// Abort aborts the give TransactionID. It returns true if
	// a transaction has been effectively aborted, otherwise it returns false.
	Abort(id TransactionID) bool
}

// An EventManipulator is a manipulator with event capabilities.
type EventManipulator interface {
	Manipulator

	// Subscribe subscribes the given EventHandler for event on the given identities.
	// If you pass nil as identities, then you will receive events for everything.
	Subscribe(filter *elemental.PushFilter, lNamespaces bool, eventHandler EventHandler, recoHandler RecoveryHandler) (EventUnsubscriber, EventFilterUpdater, error)
}

// TokenRetrieveManipulator is the interface of a manipulator that can retrieve a token.
type TokenRetrieveManipulator interface {
	RetrieveToken() error
	Validity() time.Duration
}

// Pinger is a manipulator that has a Ping method
type Pinger interface {
	Manipulator

	// Ping the connection in the given timeout. Returns an error if any
	Ping(timeout time.Duration) error
}
