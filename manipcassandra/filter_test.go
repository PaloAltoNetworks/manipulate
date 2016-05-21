// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import (
	"testing"

	"github.com/aporeto-inc/cid/materia/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInterfaceImplementations(t *testing.T) {
	var _ manipulate.FilterCompiler = (*Filter)(nil)
}

func TestParameterFilterParamCassandraFilterEqualSeparator(t *testing.T) {

	Convey("When I create a new Filter", t, func() {
		filter := &Filter{}

		filter.Keys = [][]string{[]string{"ID"}}
		filter.Separators = []string{CassandraFilterEqualSeparator}
		filter.Values = [][]interface{}{[]interface{}{"20"}}

		Convey("Then I should get the good values when calling the method compile", func() {
			So(filter.Compile(), ShouldEqual, "WHERE ID = ?")
		})
	})
}

func TestParameterFilterMultiColumn(t *testing.T) {

	Convey("When I create a new Filter", t, func() {
		filter := &Filter{}

		filter.Keys = [][]string{[]string{"ID", "name"}}
		filter.Separators = []string{CassandraFilterEqualSeparator}
		filter.Values = [][]interface{}{[]interface{}{"20", 0}}

		Convey("Then I should get the good values when calling the method compile", func() {
			So(filter.Compile(), ShouldEqual, "WHERE (ID,name) = (?,?)")
		})
	})
}

func TestParameterFilterMultiColumnAndValues(t *testing.T) {

	Convey("When I create a new Filter", t, func() {
		filter := &Filter{}

		filter.Keys = [][]string{[]string{"ID", "name"}}
		filter.Separators = []string{CassandraFilterEqualSeparator}
		filter.Values = [][]interface{}{[]interface{}{[]interface{}{"20", 0}, []interface{}{"60", 122}}}

		Convey("Then I should get the good values when calling the method compile", func() {
			So(filter.Compile(), ShouldEqual, "WHERE (ID,name) = ((?,?),(?,?))")
		})
	})
}

func TestParameterFilterParamCassandraFilterEqualSeparatorWithSeveralKeys(t *testing.T) {

	Convey("When I create a new Filter", t, func() {
		filter := &Filter{}

		filter.Keys = [][]string{[]string{"ID"}, []string{"Name"}}
		filter.Separators = []string{CassandraFilterEqualSeparator, CassandraFilterEqualSeparator}
		filter.Values = [][]interface{}{[]interface{}{"20"}, []interface{}{"Alexandre"}}

		Convey("Then I should get the good values when calling the method compile", func() {
			So(filter.Compile(), ShouldEqual, "WHERE ID = ? AND Name = ?")
		})
	})
}

func TestParameterFilterParamCassandraFilterInSeparator(t *testing.T) {

	Convey("When I create a new Filter", t, func() {
		filter := &Filter{}

		filter.Keys = [][]string{[]string{"ID"}}
		filter.Separators = []string{CassandraFilterInSeparator}
		filter.Values = [][]interface{}{[]interface{}{"20"}}

		Convey("Then I should get the good values when calling the method compile", func() {
			So(filter.Compile(), ShouldEqual, "WHERE ID IN (?)")
		})
	})
}

func TestParameterFilterParamCassandraFilterInSeparatorSeveralValues(t *testing.T) {

	Convey("When I create a new Filter", t, func() {
		filter := &Filter{}

		filter.Keys = [][]string{[]string{"ID"}}
		filter.Separators = []string{CassandraFilterInSeparator}
		filter.Values = [][]interface{}{[]interface{}{"20", "30"}}

		Convey("Then I should get the good values when calling the method compile", func() {
			So(filter.Compile(), ShouldEqual, "WHERE ID IN (?,?)")
		})
	})
}

func TestParameterFilterWithNoFilter(t *testing.T) {

	Convey("When I create a new Filter", t, func() {
		filter := &Filter{}

		Convey("Then I should get the good values when calling the method compile", func() {
			So(filter.Compile(), ShouldEqual, "")
		})
	})
}

func TestParameterFilterWithEverything(t *testing.T) {

	Convey("When I create a new Filter", t, func() {
		filter := &Filter{}
		filter.AllowFiltering = true

		filter.Keys = [][]string{[]string{"ID"}, []string{"Name"}}
		filter.Separators = []string{CassandraFilterInSeparator, CassandraFilterEqualSeparator}
		filter.Values = [][]interface{}{[]interface{}{"20", "30"}, []interface{}{"Alexandre"}}

		Convey("Then I should get the good values when calling the method compile", func() {
			So(filter.Compile(), ShouldEqual, "WHERE ID IN (?,?) AND Name = ?")
		})
	})
}
