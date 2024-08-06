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

package manipvortex

import (
	"context"
	"fmt"
	"testing"

	// nolint:revive // Allow dot imports for readability in tests
	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

func TestTestReconciler(t *testing.T) {

	Convey("Given I create a new TestReconciler", t, func() {

		r := NewTestReconciler()

		Convey("Then it should be initialized", func() {
			So(r, ShouldImplement, (*TestReconciler)(nil))
			So(r.(*testReconciler).lock, ShouldNotBeNil)
			So(r.(*testReconciler).mocks, ShouldNotBeNil)
		})

		Convey("When I call the Accept method without mock", func() {

			obj, ok, err := r.Reconcile(manipulate.NewContext(context.Background()), elemental.OperationCreate, nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("The object should be nil", func() {
				So(obj, ShouldBeNil)
			})
		})

		Convey("When I call the Accept method with a mock", func() {

			r.MockReconcile(t, func(_ manipulate.Context, _ elemental.Operation, obj elemental.Identifiable) (elemental.Identifiable, bool, error) {
				return obj, false, fmt.Errorf("boom")
			})

			obj, ok, err := r.Reconcile(manipulate.NewContext(context.Background()), elemental.OperationCreate, nil)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "boom")
			})

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})

			Convey("The object should be nil", func() {
				So(obj, ShouldBeNil)
			})

		})
	})
}
