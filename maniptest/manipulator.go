package maniptest

import (
	"sync"
	"testing"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
)

type mockedMethods struct {
	retrieveManyMock func(context *manipulate.Context, identity elemental.Identity, dest interface{}) error
	retrieveMock     func(context *manipulate.Context, objects ...elemental.Identifiable) error
	createMock       func(context *manipulate.Context, objects ...elemental.Identifiable) error
	updateMock       func(context *manipulate.Context, objects ...elemental.Identifiable) error
	deleteMock       func(context *manipulate.Context, objects ...elemental.Identifiable) error
	deleteManyMock   func(context *manipulate.Context, identity elemental.Identity) error
	countMock        func(context *manipulate.Context, identity elemental.Identity) (int, error)
	assignMock       func(context *manipulate.Context, assignation *elemental.Assignation) error
	incrementMock    func(context *manipulate.Context, identity elemental.Identity, counter string, inc int) error
	commitMock       func(id manipulate.TransactionID) error
	abortMock        func(id manipulate.TransactionID) bool
}

// A TestManipulator is the interface of mockable test manipulator.
type TestManipulator interface {
	manipulate.TransactionalManipulator
	MockRetrieveMany(t *testing.T, impl func(ctx *manipulate.Context, identity elemental.Identity, dest interface{}) error)
	MockRetrieve(t *testing.T, impl func(ctx *manipulate.Context, objects ...elemental.Identifiable) error)
	MockCreate(t *testing.T, impl func(ctx *manipulate.Context, objects ...elemental.Identifiable) error)
	MockUpdate(t *testing.T, impl func(ctx *manipulate.Context, objects ...elemental.Identifiable) error)
	MockDelete(t *testing.T, impl func(ctx *manipulate.Context, objects ...elemental.Identifiable) error)
	MockDeleteMany(t *testing.T, impl func(ctx *manipulate.Context, identity elemental.Identity) error)
	MockCount(t *testing.T, impl func(ctx *manipulate.Context, identity elemental.Identity) (int, error))
	MockAssign(t *testing.T, impl func(ctx *manipulate.Context, assignation *elemental.Assignation) error)
	MockIncrement(t *testing.T, impl func(ctx *manipulate.Context, identity elemental.Identity, counter string, inc int) error)
	MockCommit(t *testing.T, impl func(tid manipulate.TransactionID) error)
	MockAbort(t *testing.T, impl func(tid manipulate.TransactionID) bool)
}

// A testManipulator is an empty TransactionalManipulator that can be easily mocked.
type testManipulator struct {
	mocks       map[*testing.T]*mockedMethods
	lock        *sync.Mutex
	currentTest *testing.T
}

// NewTestManipulator returns a new TestManipulator.
func NewTestManipulator() TestManipulator {
	return &testManipulator{
		lock:  &sync.Mutex{},
		mocks: map[*testing.T]*mockedMethods{},
	}
}

func (m *testManipulator) MockRetrieveMany(t *testing.T, impl func(context *manipulate.Context, identity elemental.Identity, dest interface{}) error) {

	m.currentMocks(t).retrieveManyMock = impl
}

func (m *testManipulator) MockRetrieve(t *testing.T, impl func(context *manipulate.Context, objects ...elemental.Identifiable) error) {

	m.currentMocks(t).retrieveMock = impl
}

func (m *testManipulator) MockCreate(t *testing.T, impl func(context *manipulate.Context, objects ...elemental.Identifiable) error) {

	m.currentMocks(t).createMock = impl
}

func (m *testManipulator) MockUpdate(t *testing.T, impl func(context *manipulate.Context, objects ...elemental.Identifiable) error) {

	m.currentMocks(t).updateMock = impl
}

func (m *testManipulator) MockDelete(t *testing.T, impl func(context *manipulate.Context, objects ...elemental.Identifiable) error) {

	m.currentMocks(t).deleteMock = impl
}

func (m *testManipulator) MockDeleteMany(t *testing.T, impl func(context *manipulate.Context, identity elemental.Identity) error) {

	m.currentMocks(t).deleteManyMock = impl
}

func (m *testManipulator) MockCount(t *testing.T, impl func(context *manipulate.Context, identity elemental.Identity) (int, error)) {

	m.currentMocks(t).countMock = impl
}

func (m *testManipulator) MockAssign(t *testing.T, impl func(context *manipulate.Context, assignation *elemental.Assignation) error) {

	m.currentMocks(t).assignMock = impl
}

func (m *testManipulator) MockIncrement(t *testing.T, impl func(context *manipulate.Context, identity elemental.Identity, counter string, inc int) error) {

	m.currentMocks(t).incrementMock = impl
}

func (m *testManipulator) MockCommit(t *testing.T, impl func(id manipulate.TransactionID) error) {

	m.currentMocks(t).commitMock = impl
}

func (m *testManipulator) MockAbort(t *testing.T, impl func(id manipulate.TransactionID) bool) {

	m.currentMocks(t).abortMock = impl
}

func (m *testManipulator) RetrieveMany(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.retrieveManyMock != nil {
		return mock.retrieveManyMock(context, identity, dest)
	}

	return nil
}

func (m *testManipulator) Retrieve(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.retrieveMock != nil {
		return mock.retrieveMock(context, objects...)
	}

	return nil
}

func (m *testManipulator) Create(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.createMock != nil {
		return mock.createMock(context, objects...)
	}

	return nil
}

func (m *testManipulator) Update(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.updateMock != nil {
		return mock.updateMock(context, objects...)
	}

	return nil
}

func (m *testManipulator) Delete(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.deleteMock != nil {
		return mock.deleteMock(context, objects...)
	}

	return nil
}

// DeleteMany is part of the implementation of the Manipulator interface.
func (m *testManipulator) DeleteMany(context *manipulate.Context, identity elemental.Identity) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.deleteManyMock != nil {
		return mock.deleteManyMock(context, identity)
	}

	return nil
}

func (m *testManipulator) Count(context *manipulate.Context, identity elemental.Identity) (int, error) {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.countMock != nil {
		return mock.countMock(context, identity)
	}

	return 0, nil
}

func (m *testManipulator) Assign(context *manipulate.Context, assignation *elemental.Assignation) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.assignMock != nil {
		return mock.assignMock(context, assignation)
	}

	return nil
}

func (m *testManipulator) Increment(context *manipulate.Context, identity elemental.Identity, counter string, inc int) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.incrementMock != nil {
		return mock.incrementMock(context, identity, counter, inc)
	}

	return nil
}

func (m *testManipulator) Commit(id manipulate.TransactionID) error {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.commitMock != nil {
		return mock.commitMock(id)
	}

	return nil
}

func (m *testManipulator) Abort(id manipulate.TransactionID) bool {

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.abortMock != nil {
		return mock.abortMock(id)
	}

	return true
}

func (m *testManipulator) currentMocks(t *testing.T) *mockedMethods {
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
