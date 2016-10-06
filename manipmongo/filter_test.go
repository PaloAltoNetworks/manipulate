package manipmongo

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFilter_NewFilter(t *testing.T) {

	Convey("Given I create Filter", t, func() {

		f := NewFilter(nil)

		Convey("Then f should implement the interface FilterCompiler", func() {
			So(f, ShouldImplement, (*manipulate.FilterCompiler)(nil))
		})
	})
}

func TestFilter_Compile(t *testing.T) {

	Convey("Given I create Filter", t, func() {

		f := NewFilter(bson.M{"hello": "world"})

		Convey("When I run Compile", func() {

			out := f.Compile()

			Convey("Then out should have the correct type", func() {
				So(out, ShouldHaveSameTypeAs, bson.M{})
			})

			Convey("Then out['hello'] should be world", func() {
				So(out.(bson.M)["hello"], ShouldEqual, "world")
			})
		})
	})
}
