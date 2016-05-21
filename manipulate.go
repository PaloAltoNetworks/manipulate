// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import (
	"reflect"

	"github.com/aporeto-inc/cid/materia/elemental"
)

// ManipulablesList is a list of objects implementing the Manipulable interface.
type ManipulablesList []Manipulable

// Manipulable is the interface of objects that can be manipulated.
type Manipulable interface {
	elemental.Identifiable
	elemental.Validatable
}

// Manipulator is the interface of a storage backend.
type Manipulator interface {
	RetrieveChildren(contexts Contexts, parent Manipulable, identity elemental.Identity, dest interface{}) elemental.Errors
	Retrieve(contexts Contexts, objects ...Manipulable) elemental.Errors
	Create(contexts Contexts, parent Manipulable, objects ...Manipulable) elemental.Errors
	Update(contexts Contexts, objects ...Manipulable) elemental.Errors
	Delete(contexts Contexts, objects ...Manipulable) elemental.Errors
	Count(contexts Contexts, identity elemental.Identity) (int, elemental.Errors)
	Assign(contexts Contexts, parent Manipulable, assignation *elemental.Assignation) elemental.Errors
}

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

// // EventListener is the interface
// type EventListener interface {
// 	NextEvent(elemental.NotificationsChannel) elemental.Errors
// }
