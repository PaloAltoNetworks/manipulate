package maniptest

import (
	"sync"
	"testing"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

type mockedMethods struct {
	retrieveManyMock func(mctx manipulate.Context, dest elemental.Identifiables) error
	retrieveMock     func(mctx manipulate.Context, objects ...elemental.Identifiable) error
	createMock       func(mctx manipulate.Context, objects ...elemental.Identifiable) error
	updateMock       func(mctx manipulate.Context, objects ...elemental.Identifiable) error
	deleteMock       func(mctx manipulate.Context, objects ...elemental.Identifiable) error
	deleteManyMock   func(mctx manipulate.Context, identity elemental.Identity) error
	countMock        func(mctx manipulate.Context, identity elemental.Identity) (int, error)
	commitMock       func(id manipulate.TransactionID) error
	abortMock        func(id manipulate.TransactionID) bool
}

// A TestManipulator is the interface of mockable test manipulator.
type TestManipulator interface {
	manipulate.TransactionalManipulator
	MockRetrieveMany(t *testing.T, impl func(mctx manipulate.Context, dest elemental.Identifiables) error)
	MockRetrieve(t *testing.T, impl func(mctx manipulate.Context, objects ...elemental.Identifiable) error)
	MockCreate(t *testing.T, impl func(mctx manipulate.Context, objects ...elemental.Identifiable) error)
	MockUpdate(t *testing.T, impl func(mctx manipulate.Context, objects ...elemental.Identifiable) error)
	MockDelete(t *testing.T, impl func(mctx manipulate.Context, objects ...elemental.Identifiable) error)
	MockDeleteMany(t *testing.T, impl func(mctx manipulate.Context, identity elemental.Identity) error)
	MockCount(t *testing.T, impl func(mctx manipulate.Context, identity elemental.Identity) (int, error))
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

func (m *testManipulator) MockRetrieveMany(t *testing.T, impl func(mctx manipulate.Context, dest elemental.Identifiables) error) {

	m.lock.Lock()
	defer m.lock.Unlock()

	m.currentMocks(t).retrieveManyMock = impl
}

func (m *testManipulator) MockRetrieve(t *testing.T, impl func(mctx manipulate.Context, objects ...elemental.Identifiable) error) {

	m.lock.Lock()
	defer m.lock.Unlock()

	m.currentMocks(t).retrieveMock = impl
}

func (m *testManipulator) MockCreate(t *testing.T, impl func(mctx manipulate.Context, objects ...elemental.Identifiable) error) {

	m.lock.Lock()
	defer m.lock.Unlock()

	m.currentMocks(t).createMock = impl
}

func (m *testManipulator) MockUpdate(t *testing.T, impl func(mctx manipulate.Context, objects ...elemental.Identifiable) error) {

	m.lock.Lock()
	defer m.lock.Unlock()

	m.currentMocks(t).updateMock = impl
}

func (m *testManipulator) MockDelete(t *testing.T, impl func(mctx manipulate.Context, objects ...elemental.Identifiable) error) {

	m.lock.Lock()
	defer m.lock.Unlock()

	m.currentMocks(t).deleteMock = impl
}

func (m *testManipulator) MockDeleteMany(t *testing.T, impl func(mctx manipulate.Context, identity elemental.Identity) error) {

	m.lock.Lock()
	defer m.lock.Unlock()

	m.currentMocks(t).deleteManyMock = impl
}

func (m *testManipulator) MockCount(t *testing.T, impl func(mctx manipulate.Context, identity elemental.Identity) (int, error)) {

	m.lock.Lock()
	defer m.lock.Unlock()

	m.currentMocks(t).countMock = impl
}

func (m *testManipulator) MockCommit(t *testing.T, impl func(id manipulate.TransactionID) error) {

	m.lock.Lock()
	defer m.lock.Unlock()

	m.currentMocks(t).commitMock = impl
}

func (m *testManipulator) MockAbort(t *testing.T, impl func(id manipulate.TransactionID) bool) {

	m.lock.Lock()
	defer m.lock.Unlock()

	m.currentMocks(t).abortMock = impl
}

func (m *testManipulator) RetrieveMany(mctx manipulate.Context, dest elemental.Identifiables) error {

	m.lock.Lock()
	defer m.lock.Unlock()

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.retrieveManyMock != nil {
		return mock.retrieveManyMock(mctx, dest)
	}

	return nil
}

func (m *testManipulator) Retrieve(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	m.lock.Lock()
	defer m.lock.Unlock()

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.retrieveMock != nil {
		return mock.retrieveMock(mctx, objects...)
	}

	return nil
}

func (m *testManipulator) Create(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	m.lock.Lock()
	defer m.lock.Unlock()

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.createMock != nil {
		return mock.createMock(mctx, objects...)
	}

	return nil
}

func (m *testManipulator) Update(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	m.lock.Lock()
	defer m.lock.Unlock()

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.updateMock != nil {
		return mock.updateMock(mctx, objects...)
	}

	return nil
}

func (m *testManipulator) Delete(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	m.lock.Lock()
	defer m.lock.Unlock()

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.deleteMock != nil {
		return mock.deleteMock(mctx, objects...)
	}

	return nil
}

// DeleteMany is part of the implementation of the Manipulator interface.
func (m *testManipulator) DeleteMany(mctx manipulate.Context, identity elemental.Identity) error {

	m.lock.Lock()
	defer m.lock.Unlock()

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.deleteManyMock != nil {
		return mock.deleteManyMock(mctx, identity)
	}

	return nil
}

func (m *testManipulator) Count(mctx manipulate.Context, identity elemental.Identity) (int, error) {

	m.lock.Lock()
	defer m.lock.Unlock()

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.countMock != nil {
		return mock.countMock(mctx, identity)
	}

	return 0, nil
}

func (m *testManipulator) Commit(id manipulate.TransactionID) error {

	m.lock.Lock()
	defer m.lock.Unlock()

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.commitMock != nil {
		return mock.commitMock(id)
	}

	return nil
}

func (m *testManipulator) Abort(id manipulate.TransactionID) bool {

	m.lock.Lock()
	defer m.lock.Unlock()

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.abortMock != nil {
		return mock.abortMock(id)
	}

	return true
}

func (m *testManipulator) currentMocks(t *testing.T) *mockedMethods {

	mocks := m.mocks[t]

	if mocks == nil {
		mocks = &mockedMethods{}
		m.mocks[t] = mocks
	}

	m.currentTest = t
	return mocks
}
