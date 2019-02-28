package manipvortex

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

// A Prefetcher is used to perform prefetching operations
// in a Vortex manipulator. Vortex will call the Prefetch method
// before performing a Retrieve (elemental.OperationRetrieve) or a
// RetrieveMany (elemental.OperationRetrieveMany) operation to eventually
// lazy load more data than the one requested so they'll be cached in advanced.
type Prefetcher interface {

	// WarmUp is called during Vortex initialization.
	// Implementations can use the given manipulator to retrieve all the
	// objects to add into the Vortex cache before starting.
	WarmUp(context.Context, manipulate.Manipulator, elemental.ModelManager, elemental.Identity) (elemental.Identifiables, error)

	// Prefetch is called before a Retrieve or RetrieveMany operation to perform
	// cache prefetching. The given operation will be either elemental.OperationRetrieve or
	// elemental.OperationRetrieveMany. The requested identity is passed, alongside
	// with the request manipulate.Context allowing to do additional conditional logic based on
	// the initial request. It is a copy of the original context, so you cannot change anything
	// in the original request.
	//
	// If Prefetch returns some elemental.Identifiables, all of them will be added to the local cache
	// before peforming the initial request.
	//
	// If Prefetch returns nil, the original operation continues with not additional processing. According
	// to the requested consistency, the data will be either retrieved locally or from upstream.
	//
	// If prefetch returns an error, the upstream operation will be canceled and the error returned.
	// You can use the provided manipulator to retrieve the needed data.
	Prefetch(context.Context, elemental.Operation, elemental.Identity, manipulate.Manipulator, manipulate.Context) (elemental.Identifiables, error)

	// If the prefetcher uses some internal state
	// if must reset it when this is called.
	Flush()
}

// A DefaultPrefetcher will load everything
// during a warm up.
type defaultPrefetcher struct {
}

// NewDefaultPrefetcher returns a new Prefetcher that
// will simply load everything into memory
func NewDefaultPrefetcher() Prefetcher {
	return &defaultPrefetcher{}
}

func (p *defaultPrefetcher) WarmUp(ctx context.Context, m manipulate.Manipulator, manager elemental.ModelManager, identity elemental.Identity) (elemental.Identifiables, error) {

	out := manager.Identifiables(identity)

	if err := manipulate.IterFunc(
		ctx,
		m,
		manager.Identifiables(identity),
		manipulate.NewContext(ctx, manipulate.ContextOptionRecursive(true)),
		func(block elemental.Identifiables) error {
			out = out.Append(block.List()...)
			return nil
		},
		1000,
	); err != nil {
		return nil, fmt.Errorf("unable to warm up identity '%s': %s", identity.Name, err)
	}

	return out, nil
}

func (p *defaultPrefetcher) Prefetch(context.Context, elemental.Operation, elemental.Identity, manipulate.Manipulator, manipulate.Context) (elemental.Identifiables, error) {
	return nil, nil
}

func (p *defaultPrefetcher) Flush() {}

// A TestPrefetcher is prefetcher that can be used for
// testing purposes.
type TestPrefetcher interface {
	Prefetcher
	MockWarmUp(t *testing.T, impl func(context.Context, manipulate.Manipulator, elemental.ModelManager, elemental.Identity) (elemental.Identifiables, error))
	MockPrefetch(t *testing.T, impl func(context.Context, elemental.Operation, elemental.Identity, manipulate.Manipulator, manipulate.Context) (elemental.Identifiables, error))
	MockFlush(t *testing.T, impl func())
}

type mockedPrefetcherMethods struct {
	warmUpMock   func(context.Context, manipulate.Manipulator, elemental.ModelManager, elemental.Identity) (elemental.Identifiables, error)
	prefetchMock func(context.Context, elemental.Operation, elemental.Identity, manipulate.Manipulator, manipulate.Context) (elemental.Identifiables, error)
	flushMock    func()
}

type testPrefetcher struct {
	mocks       map[*testing.T]*mockedPrefetcherMethods
	lock        *sync.Mutex
	currentTest *testing.T
}

// NewTestPrefetcher returns a new TestPrefetcher.
func NewTestPrefetcher() TestPrefetcher {
	return &testPrefetcher{
		lock:  &sync.Mutex{},
		mocks: map[*testing.T]*mockedPrefetcherMethods{},
	}
}

// MockPrefetch sets the mocked implementation of Prefetch.
func (p *testPrefetcher) MockWarmUp(t *testing.T, impl func(context.Context, manipulate.Manipulator, elemental.ModelManager, elemental.Identity) (elemental.Identifiables, error)) {
	p.currentMocks(t).warmUpMock = impl
}

// MockPrefetch sets the mocked implementation of Prefetch.
func (p *testPrefetcher) MockPrefetch(t *testing.T, impl func(context.Context, elemental.Operation, elemental.Identity, manipulate.Manipulator, manipulate.Context) (elemental.Identifiables, error)) {
	p.currentMocks(t).prefetchMock = impl
}

// MockPrefetch sets the mocked implementation of Prefetch.
func (p *testPrefetcher) MockFlush(t *testing.T, impl func()) {
	p.currentMocks(t).flushMock = impl
}

func (p *testPrefetcher) WarmUp(ctx context.Context, m manipulate.Manipulator, manager elemental.ModelManager, identity elemental.Identity) (elemental.Identifiables, error) {
	if mock := p.currentMocks(p.currentTest); mock != nil && mock.warmUpMock != nil {
		return mock.warmUpMock(ctx, m, manager, identity)
	}

	return nil, nil
}

func (p *testPrefetcher) Prefetch(ctx context.Context, op elemental.Operation, identity elemental.Identity, m manipulate.Manipulator, mctx manipulate.Context) (elemental.Identifiables, error) {

	if mock := p.currentMocks(p.currentTest); mock != nil && mock.prefetchMock != nil {
		return mock.prefetchMock(ctx, op, identity, m, mctx)
	}

	return nil, nil
}

func (p *testPrefetcher) Flush() {
	if mock := p.currentMocks(p.currentTest); mock != nil && mock.flushMock != nil {
		mock.flushMock()
	}
}

func (p *testPrefetcher) currentMocks(t *testing.T) *mockedPrefetcherMethods {

	p.lock.Lock()
	defer p.lock.Unlock()

	mocks := p.mocks[t]

	if mocks == nil {
		mocks = &mockedPrefetcherMethods{}
		p.mocks[t] = mocks
	}

	p.currentTest = t
	return mocks
}
