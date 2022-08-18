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

// Package maniptest contains a Mockable TransactionalManipulator.
// It implements all method of the TransactionalManipulator but do nothing.
//
// Methods can be mocked by using one of the MockXX method.
//
// For example:
//
//	m := maniptest.NewTestManipulator()
//	m.MockCreate(t, func(context *manipulate.Context, objects ...elemental.Identifiable) error {
//	    return elemental.NewError("title", "description", "subject", 43)
//	})
//
// The next calls to the Create method will use the given method, in the context of the given *testing.T.
// If you need to reset the mocked method in the context of the same test, simply do:
//
//	m.MockCreate(t, nil)
package maniptest // import "go.aporeto.io/manipulate/maniptest"
