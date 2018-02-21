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
				So(v["tag"][0], ShouldEqual, "$name=thename")
				So(v["tag"][1], ShouldEqual, "$id=xxx")
				So(v["tag"][2], ShouldEqual, "yy=zz")
			})
		})
	})

	Convey("Given I create filter with bad comparator", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("name").Equals("thename").
			WithKey("ID").GreaterThan("x").
			WithKey("associatedTags").Contains("yy=zz").
			Done()

		Convey("When I call CompileFilter on it", func() {

			v, err := CompileFilter(f)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then v should be nil", func() {
				So(v, ShouldBeNil)
			})
		})
	})

	Convey("Given I create filter with bad contains value", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("name").Equals("thename").
			WithKey("ID").Equals("x").
			WithKey("associatedTags").Contains("yy").
			Done()

		Convey("When I call CompileFilter on it", func() {

			v, err := CompileFilter(f)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then v should be nil", func() {
				So(v, ShouldBeNil)
			})
		})
	})
}
