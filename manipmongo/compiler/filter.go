package compiler

import (
	"strings"

	"github.com/aporeto-inc/manipulate"
	"gopkg.in/mgo.v2/bson"
)

// CompileFilter compiles the given manipulate Filter into a mongo filter.
func CompileFilter(f *manipulate.Filter) bson.M {

	filter := bson.M{}
	for index, key := range f.Keys() {

		k := strings.ToLower(key[0])

		switch f.Comparators()[index] {

		case manipulate.EqualComparator:
			filter[k] = f.Values()[index][0]

		case manipulate.NotEqualComparator:
			filter[k] = bson.M{"$ne": f.Values()[index][0]}

		case manipulate.ContainComparator:
			filter[k] = bson.M{"$in": f.Values()[index]}
		}
	}

	return filter
}
