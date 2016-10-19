package maniptest

import (
	"fmt"
	"testing"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTestManipulator_Mocking(t *testing.T) {

	Convey("Given I have TestManipulator", t, func() {

		m := NewTestManipulator()

		Convey("When I call RetrieveMany without mock", func() {

			err := m.RetrieveMany(nil, PersonIdentity, nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I mock it to return an error", func() {

				m.MockRetrieveMany(t, func(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {
					return fmt.Errorf("wow such error")
				})

				err := m.RetrieveMany(nil, PersonIdentity, nil)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("When I don't mock it", func() {

				err := m.RetrieveMany(nil, PersonIdentity, nil)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
