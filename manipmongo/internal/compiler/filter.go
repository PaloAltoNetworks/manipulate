package compiler

import (
	"strings"

	"github.com/aporeto-inc/manipulate"
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
				items = append(items, bson.M{k: bson.M{"$eq": f.Values()[i][0]}})

			case manipulate.NotEqualComparator:
				items = append(items, bson.M{k: bson.M{"$ne": f.Values()[i][0]}})

			case manipulate.InComparator, manipulate.ContainComparator:
				items = append(items, bson.M{k: bson.M{"$in": f.Values()[i]}})

			case manipulate.NotInComparator, manipulate.NotContainComparator:
				items = append(items, bson.M{k: bson.M{"$nin": f.Values()[i]}})

			case manipulate.GreaterComparator:
				items = append(items, bson.M{k: bson.M{"$gte": f.Values()[i][0]}})

			case manipulate.LesserComparator:
				items = append(items, bson.M{k: bson.M{"$lte": f.Values()[i][0]}})

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
