// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compiler

import (
	"strings"
	"testing"

	"github.com/globalsign/mgo/bson"
	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
)

func toMap(in bson.D) bson.M {

	out := make(bson.M, len(in))

	for _, item := range in {

		switch iv := item.Value.(type) {

		case bson.D:
			out[item.Name] = toMap(iv)

		case []bson.D:
			outs := make([]bson.M, len(iv))
			for i, subitem := range iv {
				outs[i] = toMap(subitem)
			}
			out[item.Name] = outs

		default:
			out[item.Name] = item.Value
		}
	}

	return out
}

func TestUtils_compiler(t *testing.T) {

	Convey("Given I have a empty manipulate.Filter", t, func() {

		f := elemental.NewFilterComposer().Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{}`)
			})
		})
	})

	Convey("Given I have a simple manipulate.Filter", t, func() {

		f := elemental.NewFilterComposer().WithKey("id").Equals("5d83e7eedb40280001887565").Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"_id":{"$eq":{"$oid":"5d83e7eedb40280001887565"}}}]}`)
			})
		})
	})

	Convey("Given I have a simple manipulate.Filter with boolean set to true", t, func() {

		f := elemental.NewFilterComposer().WithKey("bool").Equals(true).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"bool":{"$eq":true}}]}`)
			})
		})
	})

	Convey("Given I have a simple manipulate.Filter with boolean set to false", t, func() {

		f := elemental.NewFilterComposer().WithKey("bool").Equals(false).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"$or":[{"bool":{"$eq":false}},{"bool":{"$exists":false}}]}]}`)
			})
		})
	})

	Convey("Given I have a simple manipulate.Filter with dots", t, func() {

		f := elemental.NewFilterComposer().WithKey("X.TOTO.Titu").Equals(1).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"x.TOTO.Titu":{"$eq":1}}]}`)
			})
		})
	})

	Convey("Given I have a simple and manipulate.Filter", t, func() {

		f := elemental.NewFilterComposer().WithKey("x").Equals(1).WithKey("y").Equals(2).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"x":{"$eq":1}},{"y":{"$eq":2}}]}`)
			})
		})
	})

	Convey("Given I have a simple multiple key and manipulate.Filter", t, func() {

		f := elemental.NewFilterComposer().WithKey("x").NotEquals(1).WithKey("x").NotEquals(2).Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"x":{"$ne":1}},{"x":{"$ne":2}}]}`)
			})
		})
	})

	Convey("Given I have a simple a complex and manipulate.Filter", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("x").Equals(1).
			WithKey("z").Contains("a", "b").
			WithKey("a").GreaterOrEqualThan(1).
			WithKey("b").LesserOrEqualThan(1).
			WithKey("c").GreaterThan(1).
			WithKey("d").LesserThan(1).
			Done()

		Convey("When I compile the filter", func() {

			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"x":{"$eq":1}},{"z":{"$in":["a","b"]}},{"a":{"$gte":1}},{"b":{"$lte":1}},{"c":{"$gt":1}},{"d":{"$lt":1}}]}`)
			})
		})
	})

	Convey("Given I have filter that contains Match", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("x").Matches("$abc^", ".*").
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"$or":[{"x":{"$regex":"$abc^"}},{"x":{"$regex":".*"}}]}]}`)
			})
		})
	})

	Convey("Given I have filter that contains Exists", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("x").Exists().
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"x":{"$exists":true}}]}`)
			})
		})
	})

	Convey("Given I have filter that contains NotExists", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("x").NotExists().
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"x":{"$exists":false}}]}`)
			})
		})
	})

	Convey("Given I have a single match on valid ID", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("ID").Equals("5d85727b919e0c397a58e940").
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"_id":{"$eq":{"$oid":"5d85727b919e0c397a58e940"}}}]}`)
			})
		})
	})

	Convey("Given I have a single match on invalid ID", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("ID").Equals("not-object-id").
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"_id":{"$eq":"not-object-id"}}]}`)
			})
		})
	})

	Convey("Given I have a single match on valid id", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("id").Equals("5d85727b919e0c397a58e940").
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"_id":{"$eq":{"$oid":"5d85727b919e0c397a58e940"}}}]}`)
			})
		})
	})

	Convey("Given I have a single match on valid _id", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("_id").Equals("5d85727b919e0c397a58e940").
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"_id":{"$eq":{"$oid":"5d85727b919e0c397a58e940"}}}]}`)
			})
		})
	})

	Convey("Given I have a In on valid ID", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("ID").In("5d85727b919e0c397a58e940", "5d85727b919e0c397a58e941").
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"_id":{"$in":[{"$oid":"5d85727b919e0c397a58e940"},{"$oid":"5d85727b919e0c397a58e941"}]}}]}`)
			})
		})
	})

	Convey("Given I have a In on invalid ID", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("ID").In("not-object-id", "5d85727b919e0c397a58e941").
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"_id":{"$in":["not-object-id",{"$oid":"5d85727b919e0c397a58e941"}]}}]}`)
			})
		})
	})

	Convey("Given I have a NotIn on valid ID", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("ID").NotIn("5d85727b919e0c397a58e940", "5d85727b919e0c397a58e941").
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"_id":{"$nin":[{"$oid":"5d85727b919e0c397a58e940"},{"$oid":"5d85727b919e0c397a58e941"}]}}]}`)
			})
		})
	})

	Convey("Given I have a composed filters", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("namespace").Equals("coucou").
			And(
				elemental.NewFilterComposer().
					WithKey("name").Equals("toto").
					WithKey("surname").Equals("titi").
					Done(),
				elemental.NewFilterComposer().
					WithKey("color").Equals("blue").
					Or(
						elemental.NewFilterComposer().
							WithKey("size").Equals("big").
							Done(),
						elemental.NewFilterComposer().
							WithKey("size").Equals("medium").
							Done(),
						elemental.NewFilterComposer().
							WithKey("list").NotIn("a", "b", "c").
							Done(),
					).
					Done(),
			).
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldEqual, `{"$and":[{"namespace":{"$eq":"coucou"}},{"$and":[{"$and":[{"name":{"$eq":"toto"}},{"surname":{"$eq":"titi"}}]},{"$and":[{"color":{"$eq":"blue"}},{"$or":[{"$and":[{"size":{"$eq":"big"}}]},{"$and":[{"size":{"$eq":"medium"}}]},{"$and":[{"list":{"$nin":["a","b","c"]}}]}]}]}]}]}`)
			})
		})
	})
}
