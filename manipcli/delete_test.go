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

func Test_generateDeleteCommandForIdentity(t *testing.T) {

	Convey("Given I generate a delete command", t, func() {

		task1 := testmodel.NewTask()
		task1.ID = "617aec75a829de0001da2032"
		task1.Name = "task1"

		m := maniptest.NewTestManipulator()
		m.MockRetrieveMany(t, func(mctx manipulate.Context, dest elemental.Identifiables) error {
			tasks := testmodel.TasksList{
				task1,
			}
			*dest.(*testmodel.TasksList) = tasks
			return nil
		})

		m.MockDelete(t, func(ctx manipulate.Context, object elemental.Identifiable) error {
			return nil
		})

		cmd, err := generateDeleteCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute", func() {
			cmd.SetArgs([]string{"617aec75a829de0001da2032"})

			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldEqual, nil)
				So(output.String(), ShouldEqual, "617aec75a829de0001da2032")
			})
		})
	})

	Convey("Given I generate a delete command that returns an error", t, func() {

		cmd, err := generateDeleteCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
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

	Convey("Given I generate a delete command and a manipulator that fails on retrieve", t, func() {

		task1 := testmodel.NewTask()
		task1.ID = "617aec75a829de0001da2032"
		task1.Name = "task1"

		m := maniptest.NewTestManipulator()
		m.MockRetrieve(t, func(mctx manipulate.Context, object elemental.Identifiable) error {
			return fmt.Errorf("retrieve boom")
		})

		m.MockDelete(t, func(ctx manipulate.Context, object elemental.Identifiable) error {
			return nil
		})

		cmd, err := generateDeleteCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute", func() {
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

	Convey("Given I generate a delete command and a manipulator that fails on delete", t, func() {

		task1 := testmodel.NewTask()
		task1.ID = "617aec75a829de0001da2032"
		task1.Name = "task1"

		m := maniptest.NewTestManipulator()

		m.MockDelete(t, func(ctx manipulate.Context, object elemental.Identifiable) error {
			return fmt.Errorf("boom")
		})

		cmd, err := generateDeleteCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute", func() {
			cmd.SetArgs([]string{"617aec75a829de0001da2032"})

			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, "unable to delete task: boom")
			})
		})
	})

}
