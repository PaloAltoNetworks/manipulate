package sharedcompiler

import (
	"testing"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFilter_CompileFilter(t *testing.T) {

	Convey("Given I create a new Filter", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("name").Equals("thename").
			AndKey("ID").Equals("xxx").
			AndKey("associatedTags").Contains("yy=zz").
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

	Convey("Given I create filter with multiple keys", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("name", "description").Equals("thename", "thedesc").
			AndKey("ID").Equals("xxx").
			AndKey("associatedTags").Contains("yy=zz").
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

	Convey("Given I create filter with multiple values", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("name").Equals("thename", "thedesc").
			AndKey("ID").Equals("xxx").
			AndKey("associatedTags").Contains("yy=zz").
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

	Convey("Given I create filter with or operator", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("name").Equals("thename").
			OrKey("ID").Equals("xxx").
			AndKey("associatedTags").Contains("yy=zz").
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

	Convey("Given I create filter with bad comparator", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("name").Equals("thename").
			AndKey("ID").GreaterThan("x").
			AndKey("associatedTags").Contains("yy=zz").
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
			AndKey("ID").Equals("x").
			AndKey("associatedTags").Contains("yy").
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
