package manipmongo

import "gopkg.in/mgo.v2/bson"

// A Filter is a structure holding filtering information.
type Filter struct {
	data bson.M
}

// NewFilter returns a new Filter.
func NewFilter(data bson.M) *Filter {

	return &Filter{
		data: data,
	}
}

// Compile compiles the filter.
func (f *Filter) Compile() interface{} {

	return f.data
}
