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

		k := strings.ToLower(key[0])
		op := f.Operators()[index]

		var dst bson.M

		if op == manipulate.OrOperator {
			ands = append(ands, bson.M{})
		}

		dst = ands[len(ands)-1]

		switch f.Comparators()[index] {

		case manipulate.EqualComparator:
			dst[k] = f.Values()[index][0]

		case manipulate.NotEqualComparator:
			dst[k] = bson.M{"$ne": f.Values()[index][0]}

		case manipulate.ContainComparator:
			dst[k] = bson.M{"$in": f.Values()[index]}

		case manipulate.GreaterComparator:
			dst[k] = bson.M{"$gte": f.Values()[index][0]}

		case manipulate.LesserComparator:
			dst[k] = bson.M{"$lte": f.Values()[index][0]}
		}
	}

	if len(ands) == 1 {
		return ands[0]
	}
	return bson.M{"$or": ands}
}
