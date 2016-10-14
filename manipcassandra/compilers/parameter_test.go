package compilers

import (
	"testing"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParameterCompile(t *testing.T) {

	Convey("When I create a new Parameter", t, func() {
		parameter := &manipulate.Parameters{}
		parameter.IfNotExists = true

		parameter2 := &manipulate.Parameters{}
		parameter2.IfNotExists = false

		parameter3 := &manipulate.Parameters{}
		parameter3.IfExists = true

		parameter4 := &manipulate.Parameters{}
		parameter4.UsingTTL = true

		parameter5 := &manipulate.Parameters{}
		parameter5.OrderByAsc = "name"

		parameter6 := &manipulate.Parameters{}
		parameter6.OrderByDesc = "AGE"

		Convey("Then I should get the good values when calling the method compile", func() {
			So(CompileParameters(parameter), ShouldEqual, "IF NOT EXISTS ")
			So(CompileParameters(parameter2), ShouldEqual, "")
			So(CompileParameters(parameter3), ShouldEqual, "IF EXISTS ")
			So(CompileParameters(parameter4), ShouldEqual, "USING TTL ")
			So(CompileParameters(parameter5), ShouldEqual, "ORDER BY name ASC ")
			So(CompileParameters(parameter6), ShouldEqual, "ORDER BY AGE DESC ")
		})
	})
}
