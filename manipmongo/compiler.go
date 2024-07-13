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
	"reflect"
	"strings"
	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate/internal/objectid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type compilerConfig struct {
	translateKeysFromSpec bool
	attrSpec              elemental.AttributeSpecifiable
}

// CompilerOption represents an option that can be passed to CompileFilter.
type CompilerOption func(*compilerConfig)

// CompilerOptionTranslateKeysFromSpec is an option that will configure the compiler to use the provided elemental.AttributeSpecifiable
// to lookup the BSON field name corresponding to the filter keys.
//
// This option is mostly useful in cases where the exposed attribute name is not the same as the field name that is stored
// in Mongo as sometimes you need to use a short field name to save space.
func CompilerOptionTranslateKeysFromSpec(spec elemental.AttributeSpecifiable) CompilerOption {

	if spec == nil {
		panic("invalid argument: must provide a non-nil elemental.AttributeSpecifiable")
	}

	return func(config *compilerConfig) {
		config.attrSpec = spec
		config.translateKeysFromSpec = true
	}
}

// CompileFilter compiles the given manipulate Filter into a mongo filter.
func CompileFilter(f *elemental.Filter, opts ...CompilerOption) bson.D {

	config := compilerConfig{}
	for _, o := range opts {
		o(&config)
	}

	if len(f.Operators()) == 0 {
		return bson.D{}
	}

	ands := []bson.D{}

	for i, operator := range f.Operators() {

		switch operator {

		case elemental.AndOperator:

			items := []bson.D{}
			k := massageKey(f.Keys()[i])
			if config.translateKeysFromSpec {
				attrSpec := config.attrSpec.SpecificationForAttribute(k)
				if attrSpec.BSONFieldName != "" {
					k = attrSpec.BSONFieldName
				}
			}

			switch f.Comparators()[i] {

			case elemental.EqualComparator:
				v := f.Values()[i][0]
				switch b := v.(type) {
				case bool:
					if b {
						items = append(items, bson.D{{Key: k, Value: primitive.M{"$eq": v}}})
					} else {
						items = append(
							items,
							bson.D{{
								Key: "$or",
								Value: []bson.D{
									{{Key: k, Value: primitive.D{{Key: "$eq", Value: v}}}},
									{{Key: k, Value: primitive.D{{Key: "$exists", Value: false}}}},
								},
							}},
						)
					}
				default:
					items = append(items, bson.D{{Key: k, Value: primitive.D{{Key: "$eq", Value: massageValue(k, v)}}}})
				}

			case elemental.NotEqualComparator:
				items = append(items, bson.D{{Key: k, Value: primitive.D{{Key: "$ne", Value: massageValue(k, f.Values()[i][0])}}}})

			case elemental.InComparator, elemental.ContainComparator:
				items = append(items, bson.D{{Key: k, Value: primitive.D{{Key: "$in", Value: massageValues(k, f.Values()[i])}}}})

			case elemental.NotInComparator, elemental.NotContainComparator:
				items = append(items, bson.D{{Key: k, Value: primitive.D{{Key: "$nin", Value: massageValues(k, f.Values()[i])}}}})

			case elemental.GreaterOrEqualComparator:
				items = append(items, bson.D{{Key: k, Value: primitive.D{{Key: "$gte", Value: massageValue(k, f.Values()[i][0])}}}})

			case elemental.GreaterComparator:
				items = append(items, bson.D{{Key: k, Value: primitive.D{{Key: "$gt", Value: massageValue(k, f.Values()[i][0])}}}})

			case elemental.LesserOrEqualComparator:
				items = append(items, bson.D{{Key: k, Value: primitive.D{{Key: "$lte", Value: massageValue(k, f.Values()[i][0])}}}})

			case elemental.LesserComparator:
				items = append(items, bson.D{{Key: k, Value: primitive.D{{Key: "$lt", Value: massageValue(k, f.Values()[i][0])}}}})

			case elemental.ExistsComparator:
				items = append(items, bson.D{{Key: k, Value: primitive.D{{Key: "$exists", Value: true}}}})

			case elemental.NotExistsComparator:
				items = append(items, bson.D{{Key: k, Value: primitive.D{{Key: "$exists", Value: false}}}})

			case elemental.MatchComparator:
				dest := []bson.D{}
				for _, v := range f.Values()[i] {
					dest = append(dest, bson.D{{Key: k, Value: primitive.D{{Key: "$regex", Value: v}}}})
				}
				items = append(items, bson.D{{Key: "$or", Value: dest}})
			}

			ands = append(ands, items...)

		case elemental.AndFilterOperator:
			subs := []bson.D{}
			for _, sub := range f.AndFilters()[i] {
				subs = append(subs, CompileFilter(sub, opts...))
			}
			ands = append(ands, bson.D{{Key: "$and", Value: subs}})

		case elemental.OrFilterOperator:
			subs := []bson.D{}
			for _, sub := range f.OrFilters()[i] {
				subs = append(subs, CompileFilter(sub, opts...))
			}
			ands = append(ands, bson.D{{Key: "$or", Value: subs}})
		}
	}

	return bson.D{
		{
			Key:   "$and",
			Value: ands,
		},
	}
}

func massageKey(key string) string {

	var k string
	if path := strings.SplitN(key, ".", 2); len(path) > 1 {
		path[0] = strings.ToLower(path[0])
		k = strings.Join(path, ".")
	} else {
		k = strings.ToLower(key)
	}

	if k == "id" {
		k = "_id"
	}

	return k
}

func massageValue(k string, v any) any {

	if reflect.TypeOf(v).Name() == "Duration" {
		return time.Now().Add(v.(time.Duration))
	}

	if k == "_id" {
		switch sv := v.(type) {
		case string:
			oid, ok := objectid.Parse(sv)
			if ok {
				return oid
			}
		}
	}

	return v
}

func massageValues(key string, values []any) []any {

	k := massageKey(key)
	out := make([]any, len(values))

	for i, v := range values {
		out[i] = massageValue(k, v)
	}

	return out
}
