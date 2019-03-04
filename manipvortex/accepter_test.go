package manipvortex

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
)

func TestTestAccepter(t *testing.T) {

	Convey("Given I create a new TestAccepter", t, func() {

		p := NewTestAccepter()

		Convey("Then it should be initialized", func() {
			So(p, ShouldImplement, (*TestAccepter)(nil))
			So(p.(*testAccepter).lock, ShouldNotBeNil)
			So(p.(*testAccepter).mocks, ShouldNotBeNil)
		})

		Convey("When I call the Accept method without mock", func() {

			ok, err := p.Accept(context.Background(), nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call the Accept method with a mock", func() {

			p.MockAccept(t, func(context.Context, ...elemental.Identifiable) (bool, error) {
				return false, fmt.Errorf("boom")
			})

			ok, err := p.Accept(context.Background(), nil)

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
