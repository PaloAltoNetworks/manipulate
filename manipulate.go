// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import "github.com/aporeto-inc/elemental"

// ManipulablesList is a list of objects implementing the Manipulable interface.
type ManipulablesList []Manipulable

// Manipulable is the interface of objects that can be manipulated.
type Manipulable interface {
	elemental.Identifiable
	elemental.Validatable
}

// Manipulator is the interface of a storage backend.
type Manipulator interface {
	// RetrieveMany retrieves the a list of objects with the given elemental.Identity and put them in the given dest.
	RetrieveMany(context *Context, identity elemental.Identity, dest interface{}) error

	// Retrieve retrieves one or multiple Manipulables. In order to be retrievable,
	// the Manipulables needs to have their Identifier correctly set.
	Retrieve(context *Context, objects ...Manipulable) error

	// Create creates a the given Manipulables in the given parent Manipulable.
	Create(context *Context, objects ...Manipulable) error

	// Update updates one or multiple Manipulables. In order to be updatable,
	// the Manipulables needs to have their Identifier correctly set.
	Update(context *Context, objects ...Manipulable) error

	// Delete deletes one or multiple Manipulables. In order to be deletable,
	// the Manipulables needs to have their Identifier correctly set.
	Delete(context *Context, objects ...Manipulable) error

	// Count returns the number of objects with the given identity.
	Count(context *Context, identity elemental.Identity) (int, error)

	// Assign is not really used yet.
	Assign(context *Context, assignation *elemental.Assignation) error

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
