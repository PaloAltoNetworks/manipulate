package manipcli

import (
	"bytes"
	"fmt"
	"testing"

	// nolint:revive // Allow dot imports for readability in tests
	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniphttp"
	"go.aporeto.io/manipulate/maniptest"
)

func Test_generateCreateCommandForIdentity(t *testing.T) {

	Convey("Given I generate a create command", t, func() {

		m := maniptest.NewTestManipulator()
		m.MockCreate(t, func(mctx manipulate.Context, object elemental.Identifiable) error {
			o := object.(*testmodel.Task)
			o.ID = "617aec75a829de0001da2032"
			return nil
		})

		cmd, err := generateCreateCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute with a json output", func() {

			cmd.Flags().Set("name", "my task") // nolint
			cmd.Flags().Set("output", "json")  // nolint
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldEqual, nil)
				So(output.String(), ShouldEqual, `{
  "ID": "617aec75a829de0001da2032",
  "description": "",
  "name": "my task",
  "parentID": "",
  "parentType": "",
  "status": "TODO"
}`)
			})
		})
	})

	Convey("Given I generate a create command that returns an error", t, func() {

		cmd, err := generateCreateCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return nil, fmt.Errorf("boom")
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute", func() {

			cmd.Flags().Set("name", "my task") // nolint
			err := cmd.Execute()

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, "unable to make manipulator: boom")
			})
		})
	})

	Convey("Given a manipulator that fails", t, func() {

		m := maniptest.NewTestManipulator()
		m.MockCreate(t, func(mctx manipulate.Context, object elemental.Identifiable) error {
			return fmt.Errorf("boom")
		})

		cmd, err := generateCreateCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute with a json output", func() {

			cmd.Flags().Set("name", "my task") // nolint
			cmd.Flags().Set("output", "json")  // nolint
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, "unable to create task: boom")
			})
		})
	})

}
