package compiler

import (
	"testing"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFilter_CompileFilter(t *testing.T) {

	Convey("Given I create a new Filter", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("name").Equals("thename").
			WithKey("ID").Equals("xxx").
			WithKey("associatedTags").Contains("yy=zz").
			Done()

		Convey("When I call CompileFilter on it", func() {

			v, err := CompileFilter(f)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the filter should be correct", func() {
				So(v.Get("q"), ShouldEqual, `name == "thename" and ID == "xxx" and associatedTags contains ["yy=zz"]`)
			})
		})
	})
}
