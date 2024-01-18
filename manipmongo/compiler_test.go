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

package manipmongo

import (
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate/manipmongo/internal"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCompilerOption(t *testing.T) {

	tests := map[string]struct {
		shouldPanic bool
		verify      func(t *testing.T)
	}{
		"CompilerOptionTranslateKeysFromSpec: basic": {
			shouldPanic: false,
			verify: func(t *testing.T) {

				config := &compilerConfig{}
				CompilerOptionTranslateKeysFromSpec(&internal.MockAttributeSpecifiable{})(config)

				if !config.translateKeysFromSpec {
					t.Error("expected 'config.translateKeysFromSpec' to be true, but it wasn't")
				}

				if config.attrSpec == nil {
					t.Fatalf("expected 'config.attrSpecs' to not be nil, but it was")
				}
			},
		},
		"CompilerOptionTranslateKeysFromSpec: nil spec should panic": {
			shouldPanic: true,
			verify: func(t *testing.T) {

				config := &compilerConfig{}
				CompilerOptionTranslateKeysFromSpec(nil)(config)
			},
		},
	}

	for summary, tc := range tests {
		t.Run(summary, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && !tc.shouldPanic {
					t.Errorf("did not expect a panic, but one occurred: %s", err)
				}
			}()

			tc.verify(t)
		})
	}
}

func TestCompiler_WithCompilerOption(t *testing.T) {

	tests := map[string]struct {
		filter *elemental.Filter
		setup  func(t *testing.T, ctrl *gomock.Controller) []CompilerOption
		opts   []CompilerOption
		want   string
	}{
		"CompilerOptionTranslateKeysFromSpec should lookup filter keys from provided spec": {
			filter: elemental.NewFilterComposer().
				WithKey("field_a").Equals("test_value").
				Done(),
			setup: func(t *testing.T, ctrl *gomock.Controller) []CompilerOption {

				spec := internal.NewMockAttributeSpecifiable(ctrl)
				spec.
					EXPECT().
					SpecificationForAttribute("field_a").
					Return(
						elemental.AttributeSpecification{
							BSONFieldName: "a",
						},
					)

				return []CompilerOption{
					CompilerOptionTranslateKeysFromSpec(spec),
				}
			},
			want: `{"$and":[{"a":{"$eq":"test_value"}}]}`,
		},
		"CompilerOptionTranslateKeysFromSpec should default to the the key name if no entry found in the provided spec": {
			filter: elemental.NewFilterComposer().
				WithKey("field_a").Equals("test_value").
				WithKey("field_b").Equals("test_value").
				Done(),
			setup: func(t *testing.T, ctrl *gomock.Controller) []CompilerOption {

				spec := internal.NewMockAttributeSpecifiable(ctrl)
				spec.
					EXPECT().
					SpecificationForAttribute("field_a").
					Return(
						elemental.AttributeSpecification{
							BSONFieldName: "a",
						},
					)

				spec.
					EXPECT().
					SpecificationForAttribute("field_b").
					Return(
						elemental.AttributeSpecification{

							// notice how this is returning an empty value since there was no spec entry for "field_b"

							BSONFieldName: "",
						},
					)

				return []CompilerOption{
					CompilerOptionTranslateKeysFromSpec(spec),
				}
			},
			want: `{"$and":[{"a":{"$eq":"test_value"}},{"field_b":{"$eq":"test_value"}}]}`,
		},
		"CompilerOptionTranslateKeysFromSpec should be able to handle nested filters": {
			filter: elemental.NewFilterComposer().
				WithKey("field_a").Equals("test_value").
				And(
					elemental.NewFilterComposer().
						WithKey("field_b").Equals("test_value").
						WithKey("field_c").Equals("test_value").
						Done(),
					elemental.NewFilterComposer().
						WithKey("field_d").Equals("test_value").
						Or(
							elemental.NewFilterComposer().
								WithKey("field_e").Equals("test_value").
								Done(),
							elemental.NewFilterComposer().
								WithKey("field_f").Equals("test_value").
								Done(),
							elemental.NewFilterComposer().
								WithKey("field_g").NotIn("test_value_a", "test_value_b", "test_value_c").
								Done(),
						).
						Done(),
				).
				Done(),
			setup: func(t *testing.T, ctrl *gomock.Controller) []CompilerOption {

				spec := internal.NewMockAttributeSpecifiable(ctrl)
				spec.
					EXPECT().
					SpecificationForAttribute("field_a").
					Return(
						elemental.AttributeSpecification{
							BSONFieldName: "a",
						},
					)

				spec.
					EXPECT().
					SpecificationForAttribute("field_b").
					Return(
						elemental.AttributeSpecification{
							BSONFieldName: "b",
						},
					)

				spec.
					EXPECT().
					SpecificationForAttribute("field_c").
					Return(
						elemental.AttributeSpecification{
							BSONFieldName: "c",
						},
					)

				spec.
					EXPECT().
					SpecificationForAttribute("field_d").
					Return(
						elemental.AttributeSpecification{
							BSONFieldName: "d",
						},
					)

				spec.
					EXPECT().
					SpecificationForAttribute("field_e").
					Return(
						elemental.AttributeSpecification{
							BSONFieldName: "e",
						},
					)

				spec.
					EXPECT().
					SpecificationForAttribute("field_f").
					Return(
						elemental.AttributeSpecification{
							BSONFieldName: "f",
						},
					)

				spec.
					EXPECT().
					SpecificationForAttribute("field_g").
					Return(
						elemental.AttributeSpecification{
							BSONFieldName: "g",
						},
					)

				return []CompilerOption{
					CompilerOptionTranslateKeysFromSpec(spec),
				}
			},
			want: `{"$and":[{"a":{"$eq":"test_value"}},{"$and":[{"$and":[{"b":{"$eq":"test_value"}},{"c":{"$eq":"test_value"}}]},{"$and":[{"d":{"$eq":"test_value"}},{"$or":[{"$and":[{"e":{"$eq":"test_value"}}]},{"$and":[{"f":{"$eq":"test_value"}}]},{"$and":[{"g":{"$nin":["test_value_a","test_value_b","test_value_c"]}}]}]}]}]}]}`,
		},
		"CompilerOptionTranslateKeysFromSpec should be able to handle filter with different casing": {
			filter: elemental.NewFilterComposer().
				WithKey("FiElD_A").Equals("test_value").
				Done(),
			setup: func(t *testing.T, ctrl *gomock.Controller) []CompilerOption {

				spec := internal.NewMockAttributeSpecifiable(ctrl)
				spec.
					EXPECT().
					SpecificationForAttribute("field_a").
					Return(
						elemental.AttributeSpecification{
							BSONFieldName: "a",
						},
					)

				return []CompilerOption{
					CompilerOptionTranslateKeysFromSpec(spec),
				}
			},
			want: `{"$and":[{"a":{"$eq":"test_value"}}]}`,
		},
	}

	for summary, tc := range tests {
		t.Run(summary, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			filter := CompileFilter(tc.filter, tc.setup(t, ctrl)...)
			bytes, err := bson.MarshalJSON(toMap(filter))
			if err != nil {
				t.Fatalf("failed to marshall compiled BSON to JSON: %s", err)
			}

			actual := strings.Replace(string(bytes), "\n", "", 1)
			if actual != tc.want {
				t.Errorf("\n"+
					"expected: %s\n"+
					"got: %s\n",
					tc.want,
					actual,
				)
			}
		})
	}
}

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

	Convey("Given I have filter that contains a duration", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("x").Equals(3 * time.Second).
			Done()

		Convey("When I compile the filter", func() {
			b, _ := bson.MarshalJSON(toMap(CompileFilter(f)))

			Convey("Then the bson should be correct", func() {
				So(strings.Replace(string(b), "\n", "", 1), ShouldStartWith, `{"$and":[{"x":{"$eq":{"$date":"`)
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
