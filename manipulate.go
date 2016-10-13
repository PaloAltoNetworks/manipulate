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
	// RetrieveChildren is cool.
	RetrieveChildren(contexts Contexts, parent Manipulable, identity elemental.Identity, dest interface{}) error

	// Retrieve is cool.
	Retrieve(contexts Contexts, objects ...Manipulable) error

	// Create is cool.
	Create(contexts Contexts, parent Manipulable, objects ...Manipulable) error

	// Update is cool.
	Update(contexts Contexts, objects ...Manipulable) error

	// Delete is cool.
	Delete(contexts Contexts, objects ...Manipulable) error

	// Count is cool.
	Count(contexts Contexts, identity elemental.Identity) (int, error)

	// Assign is cool.
	Assign(contexts Contexts, parent Manipulable, assignation *elemental.Assignation) error

	// Increment is cool
	Increment(contexts Contexts, name string, counter string, inc int, filterKeys []string, filterValues []interface{}) error
}

// A TransactionalManipulator is a Manipulator that handles transactions.
type TransactionalManipulator interface {
	Manipulator

	// Commit is cool.
	Commit(id TransactionID) error

	// Abort is cool.
	Abort(id TransactionID) bool
}
