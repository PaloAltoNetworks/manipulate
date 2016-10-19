package maniptest

import (
	"sync"
	"testing"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
)

type mockedMethods struct {
	retrieveManyMock func(context *manipulate.Context, identity elemental.Identity, dest interface{}) error
	retrieveMock     func(context *manipulate.Context, objects ...manipulate.Manipulable) error
	createMock       func(context *manipulate.Context, objects ...manipulate.Manipulable) error
	updateMock       func(context *manipulate.Context, objects ...manipulate.Manipulable) error
	deleteMock       func(context *manipulate.Context, objects ...manipulate.Manipulable) error
	countMock        func(context *manipulate.Context, identity elemental.Identity) (int, error)
	assignMock       func(context *manipulate.Context, assignation *elemental.Assignation) error
	incrementMock    func(context *manipulate.Context, name string, counter string, inc int, filterKeys []string, filterValues []interface{}) error
	commitMock       func(id manipulate.TransactionID) error
	abortMock        func(id manipulate.TransactionID) bool
}

// A TestManipulator is an empty manipulator that can be used with ApoMock.
type TestManipulator struct {
	mocks       map[*testing.T]*mockedMethods
	lock        *sync.Mutex
	currentTest *testing.T
}

// NewTestManipulator returns a new TestManipulator.
func NewTestManipulator() *TestManipulator {
	return &TestManipulator{
		lock:  &sync.Mutex{},
		mocks: map[*testing.T]*mockedMethods{},
	}
}

func (m *TestManipulator) currentMocks(t *testing.T) *mockedMethods {
	m.lock.Lock()
	defer m.lock.Unlock()

	mocks := m.mocks[t]

	if mocks == nil {
		mocks = &mockedMethods{}
		m.mocks[t] = mocks
	}

	m.currentTest = t
	return mocks
}

// MockRetrieveMany mocks RetrieveMany.
func (m *TestManipulator) MockRetrieveMany(t *testing.T, impl func(context *manipulate.Context, identity elemental.Identity, dest interface{}) error) {
	m.currentMocks(t).retrieveManyMock = impl
}

// MockRetrieve mocks Retrieve.
func (m *TestManipulator) MockRetrieve(t *testing.T, impl func(context *manipulate.Context, objects ...manipulate.Manipulable) error) {
	m.currentMocks(t).retrieveMock = impl
}

// MockCreate mocks Create.
func (m *TestManipulator) MockCreate(t *testing.T, impl func(context *manipulate.Context, objects ...manipulate.Manipulable) error) {
	m.currentMocks(t).createMock = impl
}

// MockUpdate mocks Update.
func (m *TestManipulator) MockUpdate(t *testing.T, impl func(context *manipulate.Context, objects ...manipulate.Manipulable) error) {
	m.currentMocks(t).updateMock = impl
}

// MockDelete mocks Delete.
func (m *TestManipulator) MockDelete(t *testing.T, impl func(context *manipulate.Context, objects ...manipulate.Manipulable) error) {
	m.currentMocks(t).deleteMock = impl
}

// MockCount mocks Count.
func (m *TestManipulator) MockCount(t *testing.T, impl func(context *manipulate.Context, identity elemental.Identity) (int, error)) {
	m.currentMocks(t).countMock = impl
}

// MockAssign mocks Assign.
func (m *TestManipulator) MockAssign(t *testing.T, impl func(context *manipulate.Context, assignation *elemental.Assignation) error) {
	m.currentMocks(t).assignMock = impl
}

// MockIncrement mocks Increment.
func (m *TestManipulator) MockIncrement(t *testing.T, impl func(context *manipulate.Context, name string, counter string, inc int, filterKeys []string, filterValues []interface{}) error) {
	m.currentMocks(t).incrementMock = impl
}

// MockCommit mocks Commit.
func (m *TestManipulator) MockCommit(t *testing.T, impl func(id manipulate.TransactionID) error) {
	m.currentMocks(t).commitMock = impl
}

// MockAbort mocks Abort.
func (m *TestManipulator) MockAbort(t *testing.T, impl func(id manipulate.TransactionID) bool) {
	m.currentMocks(t).abortMock = impl
}

// RetrieveMany is part of the implementation of the Manipulator interface.
func (m *TestManipulator) RetrieveMany(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.retrieveManyMock != nil {
		return mock.retrieveManyMock(context, identity, dest)
	}

	return nil
}

// Retrieve is part of the implementation of the Manipulator interface.
func (m *TestManipulator) Retrieve(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.retrieveMock != nil {
		return mock.retrieveMock(context, objects...)
	}

	return nil
}

// Create is part of the implementation of the Manipulator interface.
func (m *TestManipulator) Create(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.createMock != nil {
		return mock.createMock(context, objects...)
	}

	return nil
}

// Update is part of the implementation of the Manipulator interface.
func (m *TestManipulator) Update(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.updateMock != nil {
		return mock.updateMock(context, objects...)
	}

	return nil
}

// Delete is part of the implementation of the Manipulator interface.
func (m *TestManipulator) Delete(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.updateMock != nil {
		return mock.updateMock(context, objects...)
	}

	return nil
}

// Count is part of the implementation of the Manipulator interface.
func (m *TestManipulator) Count(context *manipulate.Context, identity elemental.Identity) (int, error) {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.countMock != nil {
		return mock.countMock(context, identity)
	}

	return 0, nil
}

// Assign is part of the implementation of the Manipulator interface.
func (m *TestManipulator) Assign(context *manipulate.Context, assignation *elemental.Assignation) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.assignMock != nil {
		return mock.assignMock(context, assignation)
	}

	return nil
}

// Increment is part of the implementation of the Manipulator interface.
func (m *TestManipulator) Increment(context *manipulate.Context, name string, counter string, inc int, filterKeys []string, filterValues []interface{}) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.incrementMock != nil {
		return mock.incrementMock(context, name, counter, inc, filterKeys, filterValues)
	}

	return nil
}

// Commit is part of the implementation of the TransactionalManipulator interface.
func (m *TestManipulator) Commit(id manipulate.TransactionID) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.commitMock != nil {
		return mock.commitMock(id)
	}

	return nil
}

// Abort is part of the implementation of the TransactionalManipulator interface.
func (m *TestManipulator) Abort(id manipulate.TransactionID) bool {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.abortMock != nil {
		return mock.abortMock(id)
	}

	return true
}
