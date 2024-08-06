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

func Test_generateUpdateCommandForIdentity(t *testing.T) {

	Convey("Given I generate a update command", t, func() {

		task1 := testmodel.NewTask()
		task1.ID = "617aec75a829de0001da2032"
		task1.Name = "task1"

		m := maniptest.NewTestManipulator()
		m.MockRetrieve(t, func(mctx manipulate.Context, object elemental.Identifiable) error {
			object.SetIdentifier(task1.ID)
			object.(*testmodel.Task).Name = task1.Name
			return nil
		})

		m.MockUpdate(t, func(mctx manipulate.Context, object elemental.Identifiable) error {
			return nil
		})

		cmd, err := generateUpdateCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute with a json output", func() {
			cmd.SetArgs([]string{"617aec75a829de0001da2032"})
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldEqual, nil)
				So(output.String(), ShouldEqual, task1.ID)
			})
		})
	})

	Convey("Given a manipulator that returns an error on update", t, func() {

		task1 := testmodel.NewTask()
		task1.ID = "617aec75a829de0001da2032"
		task1.Name = "task1"

		m := maniptest.NewTestManipulator()
		m.MockRetrieve(t, func(mctx manipulate.Context, object elemental.Identifiable) error {
			object.SetIdentifier(task1.ID)
			object.(*testmodel.Task).Name = task1.Name
			return nil
		})

		m.MockUpdate(t, func(mctx manipulate.Context, object elemental.Identifiable) error {
			return fmt.Errorf("update boom")
		})

		cmd, err := generateUpdateCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute with a json output", func() {
			cmd.SetArgs([]string{"617aec75a829de0001da2032"})
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, "unable to update task: update boom")
			})
		})
	})

	Convey("Given a manipulator that returns an error on retrieve", t, func() {

		task1 := testmodel.NewTask()
		task1.ID = "617aec75a829de0001da2032"
		task1.Name = "task1"

		m := maniptest.NewTestManipulator()
		m.MockRetrieve(t, func(mctx manipulate.Context, object elemental.Identifiable) error {
			return fmt.Errorf("retrieve boom")
		})

		m.MockUpdate(t, func(mctx manipulate.Context, object elemental.Identifiable) error {
			return nil
		})

		cmd, err := generateUpdateCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute with a json output", func() {
			cmd.SetArgs([]string{"617aec75a829de0001da2032"})
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, "unable to retrieve task: retrieve boom")
			})
		})
	})

	Convey("Given I generate a update command that returns an error", t, func() {

		cmd, err := generateUpdateCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return nil, fmt.Errorf("boom")
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute", func() {

			cmd.SetArgs([]string{"617aec75a829de0001da2032"})
			err := cmd.Execute()

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, "unable to make manipulator: boom")
			})
		})
	})

}
