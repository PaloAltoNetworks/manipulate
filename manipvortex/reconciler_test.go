package manipvortex

import (
	"context"
	"fmt"
	"testing"

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

			ok, err := r.Reconcile(manipulate.NewContext(context.Background()), elemental.OperationCreate, nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call the Accept method with a mock", func() {

			r.MockReconcile(t, func(manipulate.Context, elemental.Operation, elemental.Identifiable) (bool, error) {
				return false, fmt.Errorf("boom")
			})

			ok, err := r.Reconcile(manipulate.NewContext(context.Background()), elemental.OperationCreate, nil)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "boom")
			})

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}
