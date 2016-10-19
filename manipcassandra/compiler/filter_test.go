package compiler

import (
	"testing"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParameterFilterParamCassandraEqualSeparatorOperator(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilterComposer().WithKey("ID").Equals("20").Done()

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE ID = ?")
		})
	})
}

func TestParameterFilterMultiColumn(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilterComposer().WithKey("ID", "name").Equals("20", 0).Done()

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE (ID,name) = (?,?)")
		})
	})
}

func TestParameterFilterMultiColumnAndValues(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilterComposer().WithKey("ID", "name").Equals([]interface{}{"20", 0}, []interface{}{"60", 122}).Done()

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE (ID,name) = ((?,?),(?,?))")
		})
	})
}

func TestParameterFilterParamCassandraEqualSeparatorWithSeveralKeysOperator(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilterComposer().WithKey("ID").Equals("20").AndKey("Name").Equals("Alexandre").Done()

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE ID = ? AND Name = ?")
		})
	})
}

func TestParameterFilterParamCassandraInOperatorSeparator(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilterComposer().WithKey("ID").In("20").Done()

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE ID IN (?)")
		})
	})
}

func TestParameterFilterParamCassandraInOperatorSeparatorSeveralValues(t *testing.T) {

	Convey("When I create a new Filter", t, func() {

		filter := manipulate.NewFilterComposer().WithKey("ID").In("20", "30").Done()

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

		filter := manipulate.NewFilterComposer().WithKey("ID").In("20", "30").AndKey("Name").Equals("Alexandre").Done()

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileFilter(filter), ShouldEqual, "WHERE ID IN (?,?) AND Name = ?")
		})
	})
}
