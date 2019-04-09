// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import (
	"context"

	"go.aporeto.io/elemental"
)

// Manipulator is the interface of a storage backend.
type Manipulator interface {

	// RetrieveMany retrieves the a list of objects with the given elemental.Identity and put them in the given dest.
	RetrieveMany(mctx Context, dest elemental.Identifiables) error

	// Retrieve retrieves one or multiple elemental.Identifiables.
	// In order to be retrievable, the elemental.Identifiable needs to have their Identifier correctly set.
	Retrieve(mctx Context, object elemental.Identifiable) error

	// Create creates a the given elemental.Identifiables.
	Create(mctx Context, object elemental.Identifiable) error

	// Update updates one or multiple elemental.Identifiables.
	// In order to be updatable, the elemental.Identifiable needs to have their Identifier correctly set.
	Update(mctx Context, object elemental.Identifiable) error

	// Delete deletes one or multiple elemental.Identifiables.
	// In order to be deletable, the elemental.Identifiable needs to have their Identifier correctly set.
	Delete(mctx Context, object elemental.Identifiable) error

	// DeleteMany deletes all objects of with the given identity or
	// all the ones matching the filter in the given context.
	DeleteMany(mctx Context, identity elemental.Identity) error

	// Count returns the number of objects with the given identity.
	Count(mctx Context, identity elemental.Identity) (int, error)
}

// A TransactionalManipulator is a Manipulator that handles transactions.
type TransactionalManipulator interface {

	// Commit commits the given TransactionID.
	Commit(id TransactionID) error

	// Abort aborts the give TransactionID. It returns true if
	// a transaction has been effectively aborted, otherwise it returns false.
	Abort(id TransactionID) bool

	Manipulator
}

// A FlushableManipulator is a manipulator that can flush its
// content to somewhere, like a file.
type FlushableManipulator interface {

	// Flush flushes and empties the cache.
	Flush(ctx context.Context) error
}

// A BufferedManipulator is a Manipulator with a local cache
type BufferedManipulator interface {
	FlushableManipulator
	Manipulator
}

// SubscriberStatus is the type of a subscriber status.
type SubscriberStatus int

// Various values of SubscriberEvent.
const (
	SubscriberStatusInitialConnection SubscriberStatus = iota + 1
	SubscriberStatusInitialConnectionFailure
	SubscriberStatusReconnection
	SubscriberStatusReconnectionFailure
	SubscriberStatusDisconnection
	SubscriberStatusFinalDisconnection
	SubscriberStatusTokenRenewal
)

// A Subscriber is the interface to control a push event subscription.
type Subscriber interface {

	// Start connects to the websocket and starts collecting events
	// until the given context is canceled or any non communication error is
	// received. The eventual error will be received in the Errors() channel.
	// If not nil, the given filter will be applied right away.
	Start(context.Context, *elemental.PushFilter)

	// UpdateFilter updates the current filter.
	UpdateFilter(*elemental.PushFilter)

	// Events returns the events channel.
	Events() chan *elemental.Event

	// Errors returns the errors channel.
	Errors() chan error

	// Status returns the status channel.
	Status() chan SubscriberStatus
}

// A TokenManager issues an renew tokens periodically.
type TokenManager interface {

	// Issues isses a new token.
	Issue(context.Context) (string, error)

	// Run runs the token renewal job and published the new token in the
	// given channel.
	Run(ctx context.Context, tokenCh chan string)
}
