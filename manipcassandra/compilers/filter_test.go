package compilers

import (
	"testing"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMethodNewFilter(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewSimpleFilter("ID", "20", manipulate.EqualOperator)

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE ID = ?")
		})
	})
}

func TestParameterFilterParamCassandraEqualSeparatorOperator(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilter(
			manipulate.NewFilterKeys("ID"),
			manipulate.NewFilterValues("20"),
			manipulate.NewFilterOperators(manipulate.EqualOperator),
		)

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE ID = ?")
		})
	})
}

func TestParameterFilterMultiColumn(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilter(
			manipulate.NewFilterKeys("ID", "name"),
			manipulate.NewFilterValues("20", 0),
			manipulate.NewFilterOperators(manipulate.EqualOperator),
		)

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE (ID,name) = (?,?)")
		})
	})
}

func TestParameterFilterMultiColumnAndValues(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilter(
			manipulate.NewFilterKeys("ID", "name"),
			manipulate.FilterValues{[]interface{}{
				[]interface{}{"20", 0},
				[]interface{}{"60", 122},
			}},
			manipulate.NewFilterOperators(manipulate.EqualOperator),
		)

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE (ID,name) = ((?,?),(?,?))")
		})
	})
}

func TestParameterFilterParamCassandraEqualSeparatorWithSeveralKeysOperator(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilter(
			manipulate.NewFilterKeys("ID").Then("Name"),
			manipulate.NewFilterValues("20").Then("Alexandre"),
			manipulate.NewFilterOperators(manipulate.EqualOperator, manipulate.EqualOperator),
		)

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE ID = ? AND Name = ?")
		})
	})
}

func TestParameterFilterParamCassandraInOperatorSeparator(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilter(
			manipulate.NewFilterKeys("ID"),
			manipulate.NewFilterValues("20"),
			manipulate.NewFilterOperators(manipulate.InOperator),
		)

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE ID IN (?)")
		})
	})
}

func TestParameterFilterParamCassandraInOperatorSeparatorSeveralValues(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilter(
			manipulate.NewFilterKeys("ID"),
			manipulate.NewFilterValues("20", "30"),
			manipulate.NewFilterOperators(manipulate.InOperator),
		)

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE ID IN (?,?)")
		})
	})
}

func TestParameterFilterWithNoFilter(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := &manipulate.Filter{}

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "")
		})
	})
}

func TestParameterFilterWithEverything(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilter(
			manipulate.NewFilterKeys("ID").Then("Name"),
			manipulate.NewFilterValues("20", "30").Then("Alexandre"),
			manipulate.NewFilterOperators(manipulate.InOperator, manipulate.EqualOperator),
		)

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE ID IN (?,?) AND Name = ?")
		})
	})
}
