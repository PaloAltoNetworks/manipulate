// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maniptest

import (
	"context"
	"sync"
	"testing"

	"go.aporeto.io/manipulate"
)

type mockedTokenManagerMethods struct {
	issueMock func(context.Context) (string, error)
	runMock   func(ctx context.Context, tokenCh chan string)
}

// A TestTokenManager is the interface of mockable test manipulator.
type TestTokenManager interface {
	manipulate.TokenManager
	MockIssue(t *testing.T, impl func(context.Context) (string, error))
	MockRun(t *testing.T, impl func(ctx context.Context, tokenCh chan string))
}

// A testTokenManager is an empty TransactionalManipulator that can be easily mocked.
type testTokenManager struct {
	mocks       map[*testing.T]*mockedTokenManagerMethods
	lock        *sync.Mutex
	currentTest *testing.T
}

// NewTestTokenManager returns a new TestTokenManager.
func NewTestTokenManager() TestTokenManager {
	return &testTokenManager{
		lock:  &sync.Mutex{},
		mocks: map[*testing.T]*mockedTokenManagerMethods{},
	}
}

func (m *testTokenManager) MockIssue(t *testing.T, impl func(context.Context) (string, error)) {

	m.lock.Lock()
	defer m.lock.Unlock()

	m.currentMocks(t).issueMock = impl
}

func (m *testTokenManager) MockRun(t *testing.T, impl func(ctx context.Context, tokenCh chan string)) {

	m.lock.Lock()
	defer m.lock.Unlock()

	m.currentMocks(t).runMock = impl
}

func (m *testTokenManager) Issue(ctx context.Context) (string, error) {

	m.lock.Lock()
	defer m.lock.Unlock()

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.issueMock != nil {
		return mock.issueMock(ctx)
	}

	return "", nil
}

func (m *testTokenManager) Run(ctx context.Context, tokenCh chan string) {

	m.lock.Lock()
	defer m.lock.Unlock()

	if mock := m.currentMocks(m.currentTest); mock != nil && mock.runMock != nil {
		mock.runMock(ctx, tokenCh)
	}
}

func (m *testTokenManager) currentMocks(t *testing.T) *mockedTokenManagerMethods {

	mocks := m.mocks[t]

	if mocks == nil {
		mocks = &mockedTokenManagerMethods{}
		m.mocks[t] = mocks
	}

	m.currentTest = t
	return mocks
}
