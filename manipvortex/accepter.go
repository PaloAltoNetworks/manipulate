package manipvortex

import (
	"context"
	"sync"
	"testing"

	"go.aporeto.io/elemental"
)

// An Accepter can be given to manipvortex
// to decide to write or not an object.
type Accepter interface {

	// Accept is called before any write operation to
	// to determine is the given object should be inserted to
	// the downstream manipulator or not. If it returns true,
	// the object is accepted and will be written. If it returns
	// false, the object is ignored.
	// If it returns an error, the error will be forwarder
	// to the caller.
	Accept(context.Context, ...elemental.Identifiable) (bool, error)
}

// A TestAccepter is an Accepter that can be used for
// testing purposes.
type TestAccepter interface {
	Accepter
	MockAccept(t *testing.T, impl func(context.Context, ...elemental.Identifiable) (bool, error))
}

type mockedAccepterMethods struct {
	acceptMock func(context.Context, ...elemental.Identifiable) (bool, error)
}

type testAccepter struct {
	mocks       map[*testing.T]*mockedAccepterMethods
	lock        *sync.Mutex
	currentTest *testing.T
}

// NewTestAccepter returns a new TestAccepter.
func NewTestAccepter() TestAccepter {
	return &testAccepter{
		lock:  &sync.Mutex{},
		mocks: map[*testing.T]*mockedAccepterMethods{},
	}
}

// MockPrefetch sets the mocked implementation of Prefetch.
func (p *testAccepter) MockAccept(t *testing.T, impl func(context.Context, ...elemental.Identifiable) (bool, error)) {
	p.currentMocks(t).acceptMock = impl
}

func (p *testAccepter) Accept(ctx context.Context, i ...elemental.Identifiable) (bool, error) {
	if mock := p.currentMocks(p.currentTest); mock != nil && mock.acceptMock != nil {
		return mock.acceptMock(ctx, i...)
	}

	return true, nil
}

func (p *testAccepter) currentMocks(t *testing.T) *mockedAccepterMethods {

	p.lock.Lock()
	defer p.lock.Unlock()

	mocks := p.mocks[t]

	if mocks == nil {
		mocks = &mockedAccepterMethods{}
		p.mocks[t] = mocks
	}

	p.currentTest = t
	return mocks
}
