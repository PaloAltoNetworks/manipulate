package maniptest

import (
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
)

// A TestManipulator is an empty manipulator that can be used with ApoMock.
type TestManipulator struct {
}

// RetrieveMany is part of the implementation of the Manipulator interface.
func (*TestManipulator) RetrieveMany(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {
	return nil
}

// Retrieve is part of the implementation of the Manipulator interface.
func (*TestManipulator) Retrieve(context *manipulate.Context, objects ...manipulate.Manipulable) error {
	return nil
}

// Create is part of the implementation of the Manipulator interface.
func (*TestManipulator) Create(context *manipulate.Context, objects ...manipulate.Manipulable) error {
	return nil
}

// Update is part of the implementation of the Manipulator interface.
func (*TestManipulator) Update(context *manipulate.Context, objects ...manipulate.Manipulable) error {
	return nil
}

// Delete is part of the implementation of the Manipulator interface.
func (*TestManipulator) Delete(context *manipulate.Context, objects ...manipulate.Manipulable) error {
	return nil
}

// Count is part of the implementation of the Manipulator interface.
func (*TestManipulator) Count(context *manipulate.Context, identity elemental.Identity) (int, error) {
	return 0, nil
}

// Assign is part of the implementation of the Manipulator interface.
func (*TestManipulator) Assign(context *manipulate.Context, assignation *elemental.Assignation) error {
	return nil
}

// Increment is part of the implementation of the Manipulator interface.
func (*TestManipulator) Increment(context *manipulate.Context, name string, counter string, inc int, filterKeys []string, filterValues []interface{}) error {
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
