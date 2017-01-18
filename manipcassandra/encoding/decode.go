// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package cassandra

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gocql/gocql"
)

// Unmarshal returns an object from the given map[string]interface{}
// This method has the same behavior as the method json.Unmarshal
func Unmarshal(data interface{}, v interface{}) error {

	if (reflect.TypeOf(v).Kind() == reflect.Array || reflect.TypeOf(v).Kind() == reflect.Slice) && (reflect.TypeOf(data).Kind() != reflect.Array && reflect.TypeOf(data).Kind() != reflect.Slice) {
		return fmt.Errorf("The given data should be an array")
	}

	var val reflect.Value

	if reflect.TypeOf(data).Kind() == reflect.Array || reflect.TypeOf(data).Kind() == reflect.Slice {

		elm := reflect.TypeOf(v).Elem().Elem()

		if elm.Kind() == reflect.Ptr {
			val = reflect.Zero(elm.Elem())
		} else {
			val = reflect.Zero(elm)
		}

	} else {
		val = reflect.Indirect(reflect.ValueOf(v))
	}

	structFields := cachedTypeFields(val.Type())

	// Create fields map for faster lookup
	fieldsMap := make(map[string]field)

	for _, field := range structFields {
		fieldsMap[field.name] = field
	}

	if reflect.TypeOf(data).Kind() == reflect.Map {
		unmarshal(val, data.(map[string]interface{}), fieldsMap)
	} else {

		listData := data.([]map[string]interface{})
		list := reflect.ValueOf(v).Elem()

		cap := len(listData)
		dest := reflect.MakeSlice(list.Type(), cap, cap)

		for index := 0; index < len(listData); index++ {

			d := listData[index]
			object := dest.Index(index)
			unmarshal(object, d, fieldsMap)
		}

		list.Set(dest)
	}

	return nil
}

func unmarshal(val reflect.Value, data map[string]interface{}, fieldsMap map[string]field) {

	var defaultTime time.Time

	for key, value := range data {
		if info, ok := fieldsMap[key]; ok {
			structField := fieldByIndex(val, info.index)

			// If we have a gocql.UUID type, we convert it to a string
			if reflect.TypeOf(value).Name() == "UUID" {
				value = value.(gocql.UUID).String()
			}

			if structField.Kind() == reflect.Struct && structField.Type() != reflect.TypeOf(defaultTime) {
				newObject := reflect.New(reflect.TypeOf(structField.Interface())).Interface()
				if err := Unmarshal(value.(map[string]interface{}), newObject); err != nil {
					panic(err)
				}
				structField.Set(reflect.ValueOf(newObject).Elem())

			} else if structField.Kind() == reflect.Ptr {
				newObject := reflect.New(structField.Type().Elem()).Interface()
				if err := Unmarshal(value.(map[string]interface{}), newObject); err != nil {
					panic(err)
				}
				structField.Set(reflect.ValueOf(newObject))

			} else if (structField.Kind() == reflect.Slice || structField.Kind() == reflect.Array) && (structField.Type().Elem().Kind() == reflect.Struct || (structField.Type().Elem().Kind() == reflect.Ptr && structField.Type().Elem().Elem().Kind() == reflect.Struct)) {

				listData := value.([]map[string]interface{})
				count := len(listData)

				structFieldElem := structField.Type().Elem()
				objectType := structFieldElem

				if structFieldElem.Kind() == reflect.Ptr {
					objectType = structFieldElem.Elem()
				}

				dest := reflect.MakeSlice(structField.Type(), count, count)

				for index := 0; index < count; index++ {

					object := reflect.New(objectType).Interface()
					if err := Unmarshal(listData[index], object); err != nil {
						panic(err)
					}

					if structFieldElem.Kind() == reflect.Struct {
						dest.Index(index).Set(reflect.ValueOf(object).Elem())
					} else {
						dest.Index(index).Set(reflect.ValueOf(object))
					}
				}

				structField.Set(dest)

			} else if structField.Type().Name() == reflect.TypeOf(value).Name() {
				structField.Set(reflect.ValueOf(value))
			} else {
				// This is used when having specific types, for instance for an enum
				structField.Set(reflect.ValueOf(value).Convert(structField.Type()))
			}
		}
	}
}
