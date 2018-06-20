package compiler

import (
	"reflect"
	"strings"
	"time"

	"go.aporeto.io/manipulate"
	"github.com/globalsign/mgo/bson"
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

func massageValue(v interface{}) interface{} {

	if reflect.TypeOf(v).Name() == "Duration" {
		return time.Now().Add(v.(time.Duration))
	}

	return v
}

// CompileFilter compiles the given manipulate Filter into a mongo filter.
func CompileFilter(f *manipulate.Filter) bson.M {

	if len(f.Operators()) == 0 {
		return bson.M{}
	}

	ands := []bson.M{}

	for i, operator := range f.Operators() {

		switch operator {

		case manipulate.AndOperator:

			items := []bson.M{}
			k := massageKey(f.Keys()[i])

			switch f.Comparators()[i] {

			case manipulate.EqualComparator:
				items = append(items, bson.M{k: bson.M{"$eq": massageValue(f.Values()[i][0])}})

			case manipulate.NotEqualComparator:
				items = append(items, bson.M{k: bson.M{"$ne": massageValue(f.Values()[i][0])}})

			case manipulate.InComparator, manipulate.ContainComparator:
				items = append(items, bson.M{k: bson.M{"$in": f.Values()[i]}})

			case manipulate.NotInComparator, manipulate.NotContainComparator:
				items = append(items, bson.M{k: bson.M{"$nin": f.Values()[i]}})

			case manipulate.GreaterOrEqualComparator:
				items = append(items, bson.M{k: bson.M{"$gte": massageValue(f.Values()[i][0])}})

			case manipulate.GreaterComparator:
				items = append(items, bson.M{k: bson.M{"$gt": massageValue(f.Values()[i][0])}})

			case manipulate.LesserOrEqualComparator:
				items = append(items, bson.M{k: bson.M{"$lte": massageValue(f.Values()[i][0])}})

			case manipulate.LesserComparator:
				items = append(items, bson.M{k: bson.M{"$lt": massageValue(f.Values()[i][0])}})

			case manipulate.ExistsComparator:
				items = append(items, bson.M{k: bson.M{"$exists": true}})

			case manipulate.NotExistsComparator:
				items = append(items, bson.M{k: bson.M{"$exists": false}})

			case manipulate.MatchComparator:
				dest := []bson.M{}
				for _, v := range f.Values()[i] {
					dest = append(dest, bson.M{k: bson.M{"$regex": v, "$options": "m"}})
				}
				items = append(items, bson.M{"$or": dest})
			}

			ands = append(ands, items...)

		case manipulate.AndFilterOperator:
			subs := []bson.M{}
			for _, sub := range f.AndFilters()[i] {
				subs = append(subs, CompileFilter(sub))
			}
			ands = append(ands, bson.M{"$and": subs})

		case manipulate.OrFilterOperator:
			subs := []bson.M{}
			for _, sub := range f.OrFilters()[i] {
				subs = append(subs, CompileFilter(sub))
			}
			ands = append(ands, bson.M{"$or": subs})
		}
	}

	return bson.M{"$and": ands}
}
