package compiler

import (
	"github.com/aporeto-inc/manipulate"
	"gopkg.in/mgo.v2/bson"
)

// CompileFilter compiles the given manipulate Filter into a mongo filter.
func CompileFilter(f *manipulate.Filter) bson.M {
	return bson.M{}
}
