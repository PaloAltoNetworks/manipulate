package maniptest

import (
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
)

// A TestManipulator is an empty manipulator that can be used with ApoMock.
type TestManipulator struct {
}

// RetrieveChildren is part of the implementation of the Manipulator interface.
func (*TestManipulator) RetrieveChildren(contexts manipulate.Contexts, parent manipulate.Manipulable, identity elemental.Identity, dest interface{}) error {
	return nil
}

// Retrieve is part of the implementation of the Manipulator interface.
func (*TestManipulator) Retrieve(contexts manipulate.Contexts, objects ...manipulate.Manipulable) error {
	return nil
}

// Create is part of the implementation of the Manipulator interface.
func (*TestManipulator) Create(contexts manipulate.Contexts, parent manipulate.Manipulable, objects ...manipulate.Manipulable) error {
	return nil
}

// Update is part of the implementation of the Manipulator interface.
func (*TestManipulator) Update(contexts manipulate.Contexts, objects ...manipulate.Manipulable) error {
	return nil
}

// Delete is part of the implementation of the Manipulator interface.
func (*TestManipulator) Delete(contexts manipulate.Contexts, objects ...manipulate.Manipulable) error {
	return nil
}

// Count is part of the implementation of the Manipulator interface.
func (*TestManipulator) Count(contexts manipulate.Contexts, identity elemental.Identity) (int, error) {
	return 0, nil
}

// Assign is part of the implementation of the Manipulator interface.
func (*TestManipulator) Assign(contexts manipulate.Contexts, parent manipulate.Manipulable, assignation *elemental.Assignation) error {
	return nil
}

// Increment is part of the implementation of the Manipulator interface.
func (*TestManipulator) Increment(contexts manipulate.Contexts, name string, counter string, inc int, filterKeys []string, filterValues []interface{}) error {
	return nil
}

// Commit is part of the implementation of the TransactionalManipulator interface.
func (*TestManipulator) Commit(id manipulate.TransactionID) error {
	return nil
}

// Abort is part of the implementation of the TransactionalManipulator interface.
func (*TestManipulator) Abort(id manipulate.TransactionID) bool {
	return true
}
