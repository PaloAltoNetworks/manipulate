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

func Test_generateDeleteManyCommandForIdentity(t *testing.T) {

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

		m.MockDelete(t, func(ctx manipulate.Context, object elemental.Identifiable) error {
			return nil
		})

		cmd, err := generateDeleteManyCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I use a valid filter", func() {
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			cmd.Flags().Set("filter", "name == x")
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "you are about to delete 2 tasks. If you are sure, please use --confirm option to delete")
			})
		})

		Convey("When I use an invalid filter", func() {
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			cmd.Flags().Set("filter", "name...<.ds>")
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "unable to parse filter")
			})
		})

		Convey("When I call execute without confirm flag", func() {
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "you are about to delete 2 tasks. If you are sure, please use --confirm option to delete")
			})
		})

		Convey("When I call execute with confirm flag", func() {
			cmd.Flags().Set("confirm", "true")

			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldEqual, nil)
				So(output.String(), ShouldEqual, `617aec75a829de0001da2032
111aec75a829de0001da1111`)
			})
		})
	})

	Convey("Given I generate a delete-many command that returns an error", t, func() {

		cmd, err := generateDeleteManyCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
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

	Convey("Given a manipulator that fails on retrieve-many", t, func() {

		m := maniptest.NewTestManipulator()
		m.MockRetrieveMany(t, func(mctx manipulate.Context, dest elemental.Identifiables) error {
			return fmt.Errorf("retrieve-many boom")
		})

		m.MockDelete(t, func(ctx manipulate.Context, object elemental.Identifiable) error {
			return nil
		})

		cmd, err := generateDeleteManyCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call the function", func() {
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "unable to retrieve tasks: retrieve-many boom")
			})
		})
	})

	Convey("Given a manipulator that fails on delete", t, func() {

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

		m.MockDelete(t, func(ctx manipulate.Context, object elemental.Identifiable) error {
			return fmt.Errorf("delete boom")
		})

		cmd, err := generateDeleteManyCommandForIdentity(testmodel.TaskIdentity, testmodel.Manager(), func(opts ...maniphttp.Option) (manipulate.Manipulator, error) {
			return m, nil
		})

		So(err, ShouldEqual, nil)
		assertCommandAndSetFlags(cmd)

		Convey("When I call the function", func() {
			cmd.Flags().Set("confirm", "true")
			output := bytes.NewBufferString("")
			cmd.SetOut(output)
			err := cmd.Execute()

			Convey("Then I should get a generated command", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "delete boom")
			})
		})
	})

}
