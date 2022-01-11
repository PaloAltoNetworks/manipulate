package manipcli

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniphttp"
	"go.aporeto.io/manipulate/maniptest"
)

func Test_generateCountCommandForIdentity(t *testing.T) {

	Convey("Given I generate a count command", t, func() {

		m := maniptest.NewTestManipulator()
		m.MockCount(t, func(ctx manipulate.Context, identity elemental.Identity) (int, error) {
			return 2, nil
		})

		cmd, err := generateCountCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute without filter", func() {

			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldEqual, nil)
				So(output.String(), ShouldEqual, "2")
			})
		})

		Convey("When I call execute with valid filter", func() {

			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			cmd.Flags().Set("filter", "name == x") // nolint
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldEqual, nil)
				So(output.String(), ShouldEqual, "2")
			})
		})

		Convey("When I call execute with invalid filter", func() {

			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			cmd.Flags().Set("filter", "name...") // nolint
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "unable to parse filter")
			})
		})
	})

	Convey("Given I generate a count command that returns an error", t, func() {

		cmd, err := generateCountCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return nil, fmt.Errorf("boom")
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute", func() {

			err := cmd.Execute()

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, "unable to make manipulator: boom")
			})
		})
	})

}
