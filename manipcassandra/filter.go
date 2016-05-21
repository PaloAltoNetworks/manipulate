// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import (
	"bytes"
	"reflect"
	"strings"
)

// CassandraFilterEqualSeparator is equal to =
// CassandraFilterEqualOrSuperiorSeparator is equal to >=
// CassandraFilterEqualOrInferiorSeparator is equal to <=
// CassandraFilterEqualOrSuperiorSeparator is equal to IN
const (
	CassandraFilterEqualSeparator           = "="
	CassandraFilterEqualOrSuperiorSeparator = ">="
	CassandraFilterEqualOrInferiorSeparator = "<="
	CassandraFilterInSeparator              = "IN"
)

// Filter is a filter struct which can be used with Cassandra
type Filter struct {
	Keys           [][]string
	Values         [][]interface{}
	Separators     []string
	AllowFiltering bool
}

// Compile returns the string of the current filer
// for instance it could be WHERE ID IN (20,30) AND Name = Alexandre ALLOW FILTERING
func (f *Filter) Compile() string {

	var buffer bytes.Buffer

	for index, key := range f.Keys {

		if index == 0 {
			buffer.WriteString("WHERE")
		}

		if index > 0 {
			buffer.WriteString(" AND")
		}

		var keyValue string

		if len(key) == 1 {
			keyValue = key[0]
		} else {
			keyValue = "(" + strings.Join(key, ",") + ")"
		}

		var param string

		if len(f.Values[index]) > 1 || f.Separators[index] == CassandraFilterInSeparator {
			param = paramForValues(f.Values[index])
		}

		buffer.WriteString(" ")
		buffer.WriteString(keyValue)
		buffer.WriteString(" ")
		buffer.WriteString(f.Separators[index])

		if param == "" {
			buffer.WriteString(" ?")
		} else {
			buffer.WriteString(" ")
			buffer.WriteString(param)
		}
	}

	return buffer.String()
}

// paramForValues create the end of the query WHERE, the ((?,?),(?,?)) part.
// It will iterate in all of the array contained by the given object
func paramForValues(v []interface{}) string {

	var buffer bytes.Buffer
	buffer.WriteString("(")

	for i := 0; i < len(v); i++ {

		value := v[i]

		if reflect.ValueOf(value).Kind() == reflect.Array || reflect.ValueOf(value).Kind() == reflect.Slice {
			buffer.WriteString(paramForValues(value.([]interface{})))
		} else {
			buffer.WriteString("?")
		}

		if i < len(v)-1 {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString(")")

	return buffer.String()
}
