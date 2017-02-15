package compiler

import (
	"strings"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUtils_compiler(t *testing.T) {

	Convey("Given I have a simple manipulate.Filter", t, func() {

		f := manipulate.NewFilterComposer().WithKey("x").Equals(1).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"x":1}`)
			})
		})
	})

	Convey("Given I have a simple and manipulate.Filter", t, func() {

		f := manipulate.NewFilterComposer().WithKey("x").Equals(1).AndKey("y").Equals(2).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"x":1,"y":2}`)
			})
		})
	})

	Convey("Given I have a simple a complex and manipulate.Filter", t, func() {

		f := manipulate.NewFilterComposer().WithKey("x").Equals(1).AndKey("z").Contains("a", "b").Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"x":1,"z":{"$in":["a","b"]}}`)
			})
		})
	})

	Convey("Given I have a simple a simple or manipulate.Filter", t, func() {

		f := manipulate.NewFilterComposer().WithKey("x").Equals(1).OrKey("x").Equals(100).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$or":[{"x":1},{"x":100}]}`)
			})
		})
	})

	Convey("Given I have a simple a complex or manipulate.Filter", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("x").Equals(1).
			AndKey("z").Contains("a", "b").
			OrKey("x").Equals(100).
			AndKey("z").Contains("aa", "bb").
			Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$or":[{"x":1,"z":{"$in":["a","b"]}},{"x":100,"z":{"$in":["aa","bb"]}}]}`)
			})
		})
	})
}
