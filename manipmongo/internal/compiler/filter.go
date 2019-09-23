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
			if bson.IsObjectIdHex(sv) {
				return bson.ObjectIdHex(sv)
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
func CompileFilter(f *elemental.Filter) bson.M {

	if len(f.Operators()) == 0 {
		return bson.M{}
	}

	ands := []bson.M{}

	for i, operator := range f.Operators() {

		switch operator {

		case elemental.AndOperator:

			items := []bson.M{}
			k := massageKey(f.Keys()[i])

			switch f.Comparators()[i] {

			case elemental.EqualComparator:
				v := f.Values()[i][0]
				switch b := v.(type) {
				case bool:
					if b {
						items = append(items, bson.M{k: bson.M{"$eq": v}})
					} else {
						items = append(
							items,
							bson.M{
								"$or": []bson.M{
									bson.M{k: bson.M{"$eq": v}},
									bson.M{k: bson.M{"$exists": false}},
								},
							},
						)
					}
				default:
					items = append(items, bson.M{k: bson.M{"$eq": massageValue(k, v)}})
				}

			case elemental.NotEqualComparator:
				items = append(items, bson.M{k: bson.M{"$ne": massageValue(k, f.Values()[i][0])}})

			case elemental.InComparator, elemental.ContainComparator:
				items = append(items, bson.M{k: bson.M{"$in": massageValues(k, f.Values()[i])}})

			case elemental.NotInComparator, elemental.NotContainComparator:
				items = append(items, bson.M{k: bson.M{"$nin": massageValues(k, f.Values()[i])}})

			case elemental.GreaterOrEqualComparator:
				items = append(items, bson.M{k: bson.M{"$gte": massageValue(k, f.Values()[i][0])}})

			case elemental.GreaterComparator:
				items = append(items, bson.M{k: bson.M{"$gt": massageValue(k, f.Values()[i][0])}})

			case elemental.LesserOrEqualComparator:
				items = append(items, bson.M{k: bson.M{"$lte": massageValue(k, f.Values()[i][0])}})

			case elemental.LesserComparator:
				items = append(items, bson.M{k: bson.M{"$lt": massageValue(k, f.Values()[i][0])}})

			case elemental.ExistsComparator:
				items = append(items, bson.M{k: bson.M{"$exists": true}})

			case elemental.NotExistsComparator:
				items = append(items, bson.M{k: bson.M{"$exists": false}})

			case elemental.MatchComparator:
				dest := []bson.M{}
				for _, v := range f.Values()[i] {
					dest = append(dest, bson.M{k: bson.M{"$regex": v}})
				}
				items = append(items, bson.M{"$or": dest})
			}

			ands = append(ands, items...)

		case elemental.AndFilterOperator:
			subs := []bson.M{}
			for _, sub := range f.AndFilters()[i] {
				subs = append(subs, CompileFilter(sub))
			}
			ands = append(ands, bson.M{"$and": subs})

		case elemental.OrFilterOperator:
			subs := []bson.M{}
			for _, sub := range f.OrFilters()[i] {
				subs = append(subs, CompileFilter(sub))
			}
			ands = append(ands, bson.M{"$or": subs})
		}
	}

	return bson.M{"$and": ands}
}
