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
	// RetrieveChildren retrieves the children with the given elemental.Identity from the given parent Manipulable
	// and put them in the given dest.
	RetrieveChildren(contexts Contexts, parent Manipulable, identity elemental.Identity, dest interface{}) error

	// Retrieve retrieves one of multiple Manipulables. In order to be retrievable,
	// the Manipulables needs to have their Identifier correctly set.
	Retrieve(contexts Contexts, objects ...Manipulable) error

	// Create creates a the given Manipulables in the given parent Manipulable.
	Create(contexts Contexts, parent Manipulable, objects ...Manipulable) error

	// Update updates one of multiple Manipulables. In order to be updatable,
	// the Manipulables needs to have their Identifier correctly set.
	Update(contexts Contexts, objects ...Manipulable) error

	// Delete deletes one of multiple Manipulables. In order to be deletable,
	// the Manipulables needs to have their Identifier correctly set.
	Delete(contexts Contexts, objects ...Manipulable) error

	// Count returns the number of objects with the given identity.
	Count(contexts Contexts, identity elemental.Identity) (int, error)

	// Assign is not really used yet.
	Assign(contexts Contexts, parent Manipulable, assignation *elemental.Assignation) error

	// Increment is not very cool.
	Increment(contexts Contexts, name string, counter string, inc int, filterKeys []string, filterValues []interface{}) error
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
