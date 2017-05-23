package compiler

import (
	"strings"

	"github.com/aporeto-inc/manipulate"
	"gopkg.in/mgo.v2/bson"
)

// CompileFilter compiles the given manipulate Filter into a mongo filter.
func CompileFilter(f *manipulate.Filter) bson.M {

	ands := []bson.M{bson.M{}}

	for index, key := range f.Keys() {

		var k string
		if path := strings.SplitN(key[0], ".", 2); len(path) > 1 {
			path[0] = strings.ToLower(path[0])
			k = strings.Join(path, ".")
		} else {
			k = strings.ToLower(key[0])
		}

		op := f.Operators()[index]

		if k == "id" {
			k = "_id"
		}

		var dst bson.M

		if op == manipulate.OrOperator {
			ands = append(ands, bson.M{})
		}

		dst = ands[len(ands)-1]

		if _, ok := dst["$and"]; !ok {
			dst["$and"] = []bson.M{}
		}

		b := dst["$and"].([]bson.M)

		switch f.Comparators()[index] {

		case manipulate.EqualComparator:
			b = append(b, bson.M{k: bson.M{"$eq": f.Values()[index][0]}})

		case manipulate.NotEqualComparator:
			b = append(b, bson.M{k: bson.M{"$ne": f.Values()[index][0]}})

		case manipulate.ContainComparator:
		case manipulate.InComparator:
			b = append(b, bson.M{k: bson.M{"$in": f.Values()[index]}})

		case manipulate.GreaterComparator:
			b = append(b, bson.M{k: bson.M{"$gte": f.Values()[index][0]}})

		case manipulate.LesserComparator:
			b = append(b, bson.M{k: bson.M{"$lte": f.Values()[index][0]}})
		}

		dst["$and"] = b
	}

	if len(ands) == 1 {
		return ands[0]
	}

	return bson.M{"$or": ands}
}
