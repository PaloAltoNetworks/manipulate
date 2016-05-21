// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import (
	"testing"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParameterInterfaceImplementations(t *testing.T) {
	var _ manipulate.ParameterCompiler = (*Parameter)(nil)
}

func TestParameterCompile(t *testing.T) {

	Convey("When I create a new Parameter", t, func() {
		parameter := &Parameter{}
		parameter.IfNotExists = true

		parameter2 := &Parameter{}
		parameter2.IfNotExists = false

		parameter3 := &Parameter{}
		parameter3.IfExists = true

		parameter4 := &Parameter{}
		parameter4.UsingTTL = true

		parameter5 := &Parameter{}
		parameter5.OrderByAsc = "name"

		parameter6 := &Parameter{}
		parameter6.OrderByDesc = "AGE"

		Convey("Then I should get the good values when calling the method compile", func() {
			So(parameter.Compile(), ShouldEqual, "IF NOT EXISTS ")
			So(parameter2.Compile(), ShouldEqual, "")
			So(parameter3.Compile(), ShouldEqual, "IF EXISTS ")
			So(parameter4.Compile(), ShouldEqual, "USING TTL ")
			So(parameter5.Compile(), ShouldEqual, "ORDER BY name ASC ")
			So(parameter6.Compile(), ShouldEqual, "ORDER BY AGE DESC ")
		})
	})
}
