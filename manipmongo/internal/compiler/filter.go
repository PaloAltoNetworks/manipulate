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
	"reflect"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate/internal/objectid"
)

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

func massageValue(k string, v interface{}) interface{} {

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

func massageValues(key string, values []interface{}) []interface{} {

	k := massageKey(key)
	out := make([]interface{}, len(values))

	for i, v := range values {
		out[i] = massageValue(k, v)
	}

	return out
}

// CompileFilter compiles the given manipulate Filter into a mongo filter.
func CompileFilter(f *elemental.Filter) bson.D {

	if len(f.Operators()) == 0 {
		return bson.D{}
	}

	ands := []bson.D{}

	for i, operator := range f.Operators() {

		switch operator {

		case elemental.AndOperator:

			items := []bson.D{}
			k := massageKey(f.Keys()[i])

			switch f.Comparators()[i] {

			case elemental.EqualComparator:
				v := f.Values()[i][0]
				switch b := v.(type) {
				case bool:
					if b {
						items = append(items, bson.D{{Name: k, Value: bson.M{"$eq": v}}})
					} else {
						items = append(
							items,
							bson.D{{
								Name: "$or",
								Value: []bson.D{
									{{Name: k, Value: bson.D{{Name: "$eq", Value: v}}}},
									{{Name: k, Value: bson.D{{Name: "$exists", Value: false}}}},
								},
							}},
						)
					}
				default:
					items = append(items, bson.D{{Name: k, Value: bson.D{{Name: "$eq", Value: massageValue(k, v)}}}})
				}

			case elemental.NotEqualComparator:
				items = append(items, bson.D{{Name: k, Value: bson.D{{Name: "$ne", Value: massageValue(k, f.Values()[i][0])}}}})

			case elemental.InComparator, elemental.ContainComparator:
				items = append(items, bson.D{{Name: k, Value: bson.D{{Name: "$in", Value: massageValues(k, f.Values()[i])}}}})

			case elemental.NotInComparator, elemental.NotContainComparator:
				items = append(items, bson.D{{Name: k, Value: bson.D{{Name: "$nin", Value: massageValues(k, f.Values()[i])}}}})

			case elemental.GreaterOrEqualComparator:
				items = append(items, bson.D{{Name: k, Value: bson.D{{Name: "$gte", Value: massageValue(k, f.Values()[i][0])}}}})

			case elemental.GreaterComparator:
				items = append(items, bson.D{{Name: k, Value: bson.D{{Name: "$gt", Value: massageValue(k, f.Values()[i][0])}}}})

			case elemental.LesserOrEqualComparator:
				items = append(items, bson.D{{Name: k, Value: bson.D{{Name: "$lte", Value: massageValue(k, f.Values()[i][0])}}}})

			case elemental.LesserComparator:
				items = append(items, bson.D{{Name: k, Value: bson.D{{Name: "$lt", Value: massageValue(k, f.Values()[i][0])}}}})

			case elemental.ExistsComparator:
				items = append(items, bson.D{{Name: k, Value: bson.D{{Name: "$exists", Value: true}}}})

			case elemental.NotExistsComparator:
				items = append(items, bson.D{{Name: k, Value: bson.D{{Name: "$exists", Value: false}}}})

			case elemental.MatchComparator:
				dest := []bson.D{}
				for _, v := range f.Values()[i] {
					dest = append(dest, bson.D{{Name: k, Value: bson.D{{Name: "$regex", Value: v}}}})
				}
				items = append(items, bson.D{{Name: "$or", Value: dest}})
			}

			ands = append(ands, items...)

		case elemental.AndFilterOperator:
			subs := []bson.D{}
			for _, sub := range f.AndFilters()[i] {
				subs = append(subs, CompileFilter(sub))
			}
			ands = append(ands, bson.D{{Name: "$and", Value: subs}})

		case elemental.OrFilterOperator:
			subs := []bson.D{}
			for _, sub := range f.OrFilters()[i] {
				subs = append(subs, CompileFilter(sub))
			}
			ands = append(ands, bson.D{{Name: "$or", Value: subs}})
		}
	}

	return bson.D{
		{
			Name:  "$and",
			Value: ands,
		},
	}
}
