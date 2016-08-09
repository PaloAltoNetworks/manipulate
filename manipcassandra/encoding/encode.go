// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package cassandra

import (
	"fmt"
	"reflect"
	"time"
)

// Marshal returns a map[string]interface{} from the given object
// This method has the same behavior as the method json.Marshal
func Marshal(v interface{}) (map[string]interface{}, error) {

	structVal := reflect.Indirect(reflect.ValueOf(v))
	kind := structVal.Kind()

	if kind != reflect.Struct {
		return nil, fmt.Errorf("The given interface is not a struct type")
	}

	structFields := cachedTypeFields(structVal.Type())
	mapVal := make(map[string]interface{}, len(structFields))

	for _, info := range structFields {

		field := fieldByIndex(structVal, info.index)
		kind := field.Kind()

		if kind == reflect.Struct || kind == reflect.Ptr {

			// TODO: decide if we want an empty dict or nil
			if kind == reflect.Ptr && reflect.ValueOf(field.Interface()).IsNil() {
				continue
			}

			// TODO: decide if we want an empty dict or nil
			if kind == reflect.Struct && field.Interface() == reflect.New(field.Type()).Elem().Interface() {
				continue
			}

			dict, err := Marshal(field.Interface())

			if err != nil {
				return nil, err
			}

			mapVal[info.name] = dict

		} else if (kind == reflect.Slice || kind == reflect.Array) && (field.Type().Elem().Kind() == reflect.Struct || (field.Type().Elem().Kind() == reflect.Ptr && field.Type().Elem().Elem().Kind() == reflect.Struct)) {

			count := field.Len()
			objects := []map[string]interface{}{}

			for index := 0; index < count; index++ {
				dict, err := Marshal(field.Index(index).Interface())

				if err != nil {
					return nil, err
				}

				objects = append(objects, dict)
			}

			mapVal[info.name] = objects

		} else {
			mapVal[info.name] = field.Interface()
		}
	}

	return mapVal, nil
}

func fieldByIndex(v reflect.Value, index []int) reflect.Value {

	for _, i := range index {

		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				if v.CanSet() {
					v.Set(reflect.New(v.Type().Elem()))
				} else {
					return reflect.Value{}
				}
			}
			v = v.Elem()
		}
		v = v.Field(i)
	}

	return v
}

// FieldsAndValues returns a list field names and a corresponing list of values
// for the given struct. For details on how the field names are determined please
// see StructToMap.
func FieldsAndValues(val interface{}) ([]string, []interface{}, error) {

	// indirect so function works with both structs and pointers to them
	structVal := reflect.Indirect(reflect.ValueOf(val))
	kind := structVal.Kind()

	if kind != reflect.Struct {
		return nil, nil, fmt.Errorf("The given interface is not a struct type")
	}

	structFields := cachedTypeFields(structVal.Type())
	fields := []string{}
	values := []interface{}{}

	for _, info := range structFields {
		field := fieldByIndex(structVal, info.index)

		if info.omitEmpty && isEmptyValue(field) {
			continue
		}

		kind := field.Kind()
		fields = append(fields, info.name)

		if isEmptyValue(field) && info.autoTimestamp {
			// TODO: should it be calculated directly in cassandra ?
			values = append(values, time.Now())
		} else if info.autoTimestampOverride {
			// TODO: should it be calculated directly in cassandra ?
			values = append(values, time.Now())
		} else if (kind == reflect.Slice || kind == reflect.Array) && (field.Type().Elem().Kind() == reflect.Struct || (field.Type().Elem().Kind() == reflect.Ptr && field.Type().Elem().Elem().Kind() == reflect.Struct)) {

			count := field.Len()
			objects := []map[string]interface{}{}

			for index := 0; index < count; index++ {
				dict, err := Marshal(field.Index(index).Interface())

				if err != nil {
					return nil, nil, err
				}

				objects = append(objects, dict)
			}

			values = append(values, objects)

		} else {
			values = append(values, field.Interface())
		}
	}

	return fields, values, nil
}

// PrimaryFieldsAndValues returns a list field names and a corresponing list of values of th eprimary keys of the object
// for the given struct. For details on how the field names are determined please
// see StructToMap.
func PrimaryFieldsAndValues(val interface{}) ([]string, []interface{}, error) {

	// indirect so function works with both structs and pointers to them
	structVal := reflect.Indirect(reflect.ValueOf(val))
	kind := structVal.Kind()

	if kind != reflect.Struct {
		return nil, nil, fmt.Errorf("The given interface is not a struct type")
	}

	structFields := cachedTypeFields(structVal.Type())
	fields := []string{}
	values := []interface{}{}

	for _, info := range structFields {
		field := fieldByIndex(structVal, info.index)

		if !info.isPrimaryKey || isEmptyPrimaryValue(field) {
			continue
		}

		fields = append(fields, info.name)
		values = append(values, field.Interface())
	}

	return fields, values, nil
}

func isEmptyPrimaryValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	var defaultTime time.Time
	if v.Type() == reflect.TypeOf(defaultTime) {
		return defaultTime.Equal(v.Interface().(time.Time))
	}

	return false
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	var defaultTime time.Time
	if v.Type() == reflect.TypeOf(defaultTime) {
		return defaultTime.Equal(v.Interface().(time.Time))
	}

	return false
}
