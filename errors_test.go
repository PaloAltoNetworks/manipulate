package manipulate

import (
	"testing"

	"github.com/aporeto-inc/elemental"
	. "github.com/smartystreets/goconvey/convey"
)

func TestErrors_NewError(t *testing.T) {

	Convey("Given I create a NewError", t, func() {

		err := NewError("this is an error", ErrCannotUnmarshal)

		Convey("Then err should be an elemental.Errors", func() {
			_, ok := err.(elemental.Error)

			So(ok, ShouldBeTrue)
		})

		Convey("Then the error should should be correctly initialized", func() {
			So(err.(elemental.Error).Code, ShouldEqual, ErrCannotUnmarshal)
			So(err.(elemental.Error).Title, ShouldEqual, "Unable to unmarshal data.")
			So(err.(elemental.Error).Description, ShouldEqual, "this is an error")
		})
	})
}
