package compiler

import (
	"strings"
	"testing"

	"github.com/aporeto-inc/manipulate"
	"github.com/globalsign/mgo/bson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUtils_compiler(t *testing.T) {

	Convey("Given I have a simple manipulate.Filter", t, func() {

		f := manipulate.NewFilterComposer().WithKey("id").Equals(1).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"_id":{"$eq":1}}]}`)
			})
		})
	})

	Convey("Given I have a simple manipulate.Filter with dots", t, func() {

		f := manipulate.NewFilterComposer().WithKey("X.TOTO.Titu").Equals(1).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"x.TOTO.Titu":{"$eq":1}}]}`)
			})
		})
	})

	Convey("Given I have a simple and manipulate.Filter", t, func() {

		f := manipulate.NewFilterComposer().WithKey("x").Equals(1).WithKey("y").Equals(2).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"x":{"$eq":1}},{"y":{"$eq":2}}]}`)
			})
		})
	})

	Convey("Given I have a simple multiple key and manipulate.Filter", t, func() {

		f := manipulate.NewFilterComposer().WithKey("x").NotEquals(1).WithKey("x").NotEquals(2).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"x":{"$ne":1}},{"x":{"$ne":2}}]}`)
			})
		})
	})

	Convey("Given I have a simple a complex and manipulate.Filter", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("x").Equals(1).
			WithKey("z").Contains("a", "b").
			WithKey("a").GreaterOrEqualThan(1).
			WithKey("b").LesserOrEqualThan(1).
			WithKey("c").GreaterThan(1).
			WithKey("d").LesserThan(1).
			Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"x":{"$eq":1}},{"z":{"$in":["a","b"]}},{"a":{"$gte":1}},{"b":{"$lte":1}},{"c":{"$gt":1}},{"d":{"$lt":1}}]}`)
			})
		})
	})

	Convey("Given I have filter that contains Match", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("x").Matches("$abc^", ".*").
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"$or":[{"x":{"$options":"m","$regex":"$abc^"}},{"x":{"$options":"m","$regex":".*"}}]}]}`)
			})
		})
	})

	Convey("Given I have a composed filters", t, func() {

		f := manipulate.NewFilterComposer().
			WithKey("namespace").Equals("coucou").
			And(
				manipulate.NewFilterComposer().
					WithKey("name").Equals("toto").
					WithKey("surname").Equals("titi").
					Done(),
				manipulate.NewFilterComposer().
					WithKey("color").Equals("blue").
					Or(
						manipulate.NewFilterComposer().
							WithKey("size").Equals("big").
							Done(),
						manipulate.NewFilterComposer().
							WithKey("size").Equals("medium").
							Done(),
						manipulate.NewFilterComposer().
							WithKey("list").NotIn("a", "b", "c").
							Done(),
					).
					Done(),
			).
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(CompileFilter(f))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"namespace":{"$eq":"coucou"}},{"$and":[{"$and":[{"name":{"$eq":"toto"}},{"surname":{"$eq":"titi"}}]},{"$and":[{"color":{"$eq":"blue"}},{"$or":[{"$and":[{"size":{"$eq":"big"}}]},{"$and":[{"size":{"$eq":"medium"}}]},{"$and":[{"list":{"$nin":["a","b","c"]}}]}]}]}]}]}`)
			})
		})
	})
}
