// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate_test

import (
	"testing"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

// UserIdentity represents the Identity of the object
var TagIdentity = elemental.Identity{
	Name:     "tag",
	Category: "tag",
}

type Tag struct {
	ID string `cql:"id"`
}

func (t *Tag) Identifier() string {
	return t.ID
}

// Identity returns the Identity of the object.
func (t *Tag) Identity() elemental.Identity {

	return TagIdentity
}

// SetIdentifier sets the value of the object's unique identifier.
func (t *Tag) SetIdentifier(ID string) {
	t.ID = ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (t *Tag) Validate() elemental.Errors {
	return nil
}

func TestMethodConvertPointerArrayToManipulables(t *testing.T) {

	Convey("Given I create a two objects manipulator", t, func() {

		tag := &Tag{}
		tag2 := &Tag{}

		var listTags []*Tag
		listTags = append(listTags, tag, tag2)

		Convey("Then I call the method ConvertArrayToManipulables", func() {

			m := manipulate.ConvertArrayToManipulables(listTags)
			So(m, ShouldHaveSameTypeAs, []manipulate.Manipulable{})
			So(m[0], ShouldEqual, tag)
			So(m[1], ShouldEqual, tag2)
		})
	})
}
