// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import (
	"reflect"

	"github.com/aporeto-inc/elemental"
)

// stringInSlice returns true or false if the given string is in the given list
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// getValues returns a list of all of the value of the given object, it will iterate in all array of the given object
func getValues(v []interface{}) []interface{} {

	var values []interface{}

	for i := 0; i < len(v); i++ {

		value := v[i]

		if reflect.ValueOf(value).Kind() == reflect.Array || reflect.ValueOf(value).Kind() == reflect.Slice {
			values = append(values, getValues(value.([]interface{}))...)
		} else {
			return v
		}
	}

	return values
}

func makeError(err string, code int) error {
	return elemental.NewError(
		errorTitles[code],
		err,
		"manipulate",
		code,
	)
}
