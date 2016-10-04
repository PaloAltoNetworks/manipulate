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
	CassandraFilterContainsSeparator        = "CONTAINS"
)

// Filter is a filter struct which can be used with Cassandra
type Filter struct {
	Keys           [][]string
	Values         [][]interface{}
	Separators     []string
	AllowFiltering bool
}

// NewFilter return a filter for operation as NAME = 'Alexandre'
func NewFilter(key string, value interface{}, separator string) *Filter {
	filter := &Filter{}

	filter.Keys = [][]string{[]string{key}}
	filter.Separators = []string{separator}
	filter.Values = [][]interface{}{[]interface{}{value}}

	return filter
}

// NewMultipleFilters return a filter for operation as ID = 123 AND Name = Alexandre
// The arg separator could be either a list of string or a string, if only one string is given, the same separator will be used everywhere
func NewMultipleFilters(keys []string, values []interface{}, separators interface{}) *Filter {

	var isSeparatorArray bool

	if reflect.TypeOf(separators).Kind() == reflect.Slice || reflect.TypeOf(separators).Kind() == reflect.Array {
		isSeparatorArray = true
	}

	filter := &Filter{}
	filter.Keys = [][]string{}
	filter.Separators = []string{}
	filter.Values = [][]interface{}{}

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		value := values[i]

		filter.Keys = append(filter.Keys, []string{key})

		if isSeparatorArray {
			filter.Separators = append(filter.Separators, separators.([]string)[i])
		} else {
			filter.Separators = append(filter.Separators, separators.(string))
		}

		filter.Values = append(filter.Values, []interface{}{value})
	}

	return filter
}

// NewCollectionFilter return a filter for operation as (ID,name) = (123,'Alexandre')
func NewCollectionFilter(keys []string, values []interface{}, separator string) *Filter {
	filter := &Filter{}

	filter.Keys = [][]string{keys}
	filter.Separators = []string{separator}
	filter.Values = [][]interface{}{values}

	return filter
}

// Compile returns the string of the current filer
// for instance it could be WHERE ID IN (20,30) AND Name = Alexandre ALLOW FILTERING
func (f *Filter) Compile() interface{} {

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
