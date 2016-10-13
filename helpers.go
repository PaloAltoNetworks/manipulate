package manipulate

import "reflect"

// ConvertArrayToManipulables convert the given array of interface into an array of Manipulable
func ConvertArrayToManipulables(i interface{}) []Manipulable {

	var manipulables []Manipulable
	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			manipulables = append(manipulables, val.Index(i).Interface().(Manipulable))
		}
	}

	return manipulables
}
