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

func Test_generateListCommandForIdentity(t *testing.T) {

	Convey("Given I generate a delete-many command", t, func() {

		task1 := testmodel.NewTask()
		task1.ID = "617aec75a829de0001da2032"
		task1.Name = "task1"

		task2 := testmodel.NewTask()
		task2.ID = "111aec75a829de0001da1111"
		task2.Name = "task2"

		m := maniptest.NewTestManipulator()
		m.MockRetrieveMany(t, func(mctx manipulate.Context, dest elemental.Identifiables) error {
			tasks := testmodel.TasksList{
				task1,
				task2,
			}
			*dest.(*testmodel.TasksList) = tasks
			return nil
		})

		cmd, err := generateListCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
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
				So(output.String(), ShouldEqual, `[
  {
    "ID": "617aec75a829de0001da2032",
    "description": "",
    "name": "task1",
    "parentID": "",
    "parentType": "",
    "status": "TODO"
  },
  {
    "ID": "111aec75a829de0001da1111",
    "description": "",
    "name": "task2",
    "parentID": "",
    "parentType": "",
    "status": "TODO"
  }
]`)
			})
		})

		Convey("When I call execute with a valid filter", func() {
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			cmd.Flags().Set("filter", "name == x") // nolint
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldEqual, nil)
				So(output.String(), ShouldEqual, `[
  {
    "ID": "617aec75a829de0001da2032",
    "description": "",
    "name": "task1",
    "parentID": "",
    "parentType": "",
    "status": "TODO"
  },
  {
    "ID": "111aec75a829de0001da1111",
    "description": "",
    "name": "task2",
    "parentID": "",
    "parentType": "",
    "status": "TODO"
  }
]`)
			})
		})

		Convey("When I call execute with an invalid filter", func() {
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

	Convey("Given I generate a delete-many command that returns an error", t, func() {

		cmd, err := generateListCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
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

	Convey("Given I generate a delete-many command and a manipulator that fails", t, func() {

		m := maniptest.NewTestManipulator()
		m.MockRetrieveMany(t, func(mctx manipulate.Context, dest elemental.Identifiables) error {
			return fmt.Errorf("boom")
		})

		cmd, err := generateListCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call execute without filter", func() {
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldEqual, "unable to retrieve all tasks: boom")

			})
		})
	})

}
