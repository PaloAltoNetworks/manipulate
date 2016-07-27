// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package cassandra

import (
	"reflect"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMethodPrimaryFieldsAndValues(t *testing.T) {

	Convey("Given I call the method cassandra.Fields", t, func() {

		person := &Stark{}
		person.Name = "Ned"
		person.ID = "124"
		person.Age = 44
		fields, values, err := PrimaryFieldsAndValues(person)

		Convey("Then I should get the appropriate map", func() {
			So(err, ShouldBeNil)
			So(fields, ShouldResemble, []string{"id", "name"})
			So(values[1], ShouldEqual, "Ned")
			So(values[0], ShouldEqual, "124")
		})
	})
}

func TestMethodPrimaryFieldsAndValuesError(t *testing.T) {

	Convey("Given I call the method cassandra.Fields", t, func() {

		fields, values, err := PrimaryFieldsAndValues(2)

		Convey("Then I should get the appropriate map", func() {
			So(err.Error(), ShouldEqual, `The given interface is not a struct type`)
			So(fields, ShouldBeNil)
			So(values, ShouldBeNil)
		})
	})
}

func TestMethodFields(t *testing.T) {

	Convey("Given I call the method cassandra.Fields", t, func() {

		person := &Person{}
		person.Name = "Alexandre"
		fields, values, err := FieldsAndValues(person)

		Convey("Then I should get the appropriate map", func() {
			So(err, ShouldBeNil)
			So(fields, ShouldResemble, []string{"id", "name", "Age", "siblings"})
			So(values[1], ShouldEqual, "Alexandre")
		})
	})
}

func TestMethodFieldsAndTags(t *testing.T) {

	Convey("Given I call the method cassandra.Fields", t, func() {

		now := time.Now()
		date := &Date{}
		fields, values, err := FieldsAndValues(date)
		now2 := time.Now()

		Convey("Then I should get the appropriate map", func() {
			So(err, ShouldBeNil)
			So(fields, ShouldResemble, []string{"creationDate", "updateDate"})
			So(now.Before(values[0].(time.Time)), ShouldBeTrue)
			So(now2.After(values[0].(time.Time)), ShouldBeTrue)
			So(now.Before(values[1].(time.Time)), ShouldBeTrue)
			So(now2.After(values[1].(time.Time)), ShouldBeTrue)

			date.CreationDate = values[0].(time.Time)
			date.UpdateDate = values[1].(time.Time)
			date.Name = "birthday"

			now3 := time.Now()
			fields, values, err = FieldsAndValues(date)
			now4 := time.Now()
			So(err, ShouldBeNil)
			So(fields, ShouldResemble, []string{"creationDate", "updateDate", "name"})
			So(now.Before(values[0].(time.Time)), ShouldBeTrue)
			So(now3.After(values[0].(time.Time)), ShouldBeTrue)
			So(now4.After(values[0].(time.Time)), ShouldBeTrue)

			So(now.Before(values[1].(time.Time)), ShouldBeTrue)
			So(now3.Before(values[1].(time.Time)), ShouldBeTrue)
			So(now4.After(values[1].(time.Time)), ShouldBeTrue)
			So(values[2], ShouldEqual, "birthday")

		})
	})
}

func TestMethodFieldsError(t *testing.T) {

	Convey("Given I call the method cassandra.Fields with a none struct", t, func() {

		list, values, err := FieldsAndValues(2)

		Convey("Then I should get the appropriate map", func() {
			So(err.Error(), ShouldEqual, `The given interface is not a struct type`)
			So(list, ShouldBeNil)
			So(values, ShouldBeNil)
		})
	})
}

func TestMethodMarshalWithSimpleStruct(t *testing.T) {

	Convey("Given I call the method cassandra.Marshal", t, func() {

		person := &Person{}
		person.Name = "Alexandre"
		person.Age = 2
		person.Address = "Sarralbe"
		person.Country = "France"
		person.ID = "19a65f7e-102f-11e6-8309-f45c89941b79"
		person.Siblings = []string{"Celine", "Dominique"}
		person.language = "French"
		dict, err := Marshal(person)

		Convey("Then I should get the appropriate map", func() {
			So(err, ShouldBeNil)
			So(dict["name"], ShouldEqual, "Alexandre")
			So(dict["Age"], ShouldEqual, 2)
			So(dict["Address"], ShouldBeNil)
			So(dict["id"], ShouldEqual, "19a65f7e-102f-11e6-8309-f45c89941b79")
			So(dict["siblings"], ShouldResemble, []string{"Celine", "Dominique"})
			So(dict["country"], ShouldEqual, "France")
			So(dict["zipcode"], ShouldEqual, "")
			So(dict["language"], ShouldEqual, nil)
		})
	})
}

func TestMethodMarshalWithEmbeededStruct(t *testing.T) {

	Convey("Given I call the method cassandra.Marshal", t, func() {
		person := &God{}
		person.Name = "Alexandre"
		person.Age = 2
		person.Address = "Sarralbe"
		person.Country = "France"
		person.ID = "19a65f7e-102f-11e6-8309-f45c89941b79"
		person.Siblings = []string{"Celine", "Dominique"}
		person.language = "French"
		person.FirstName = "Alexandre"
		person.LastName = "Wilhelm"
		person.X = "coucou"
		dict, err := Marshal(person)

		Convey("Then I should get the appropriate map", func() {
			So(err, ShouldBeNil)
			So(dict["name"], ShouldEqual, "Wilhelm")
			So(dict["firstName"], ShouldEqual, "Alexandre")
			So(dict["Age"], ShouldEqual, 2)
			So(dict["Address"], ShouldBeNil)
			So(dict["id"], ShouldEqual, "19a65f7e-102f-11e6-8309-f45c89941b79")
			So(dict["siblings"], ShouldResemble, []string{"Celine", "Dominique"})
			So(dict["country"], ShouldEqual, "France")
			So(dict["zipcode"], ShouldEqual, "")
			So(dict["language"], ShouldEqual, nil)
			So(dict["X"], ShouldEqual, "coucou")
			So(dict["json_attribute"], ShouldEqual, nil)
		})
	})
}

func TestMethodMarshalWithTwoTimesSameTag(t *testing.T) {

	Convey("Given I call the method cassandra.Marshal", t, func() {

		person := &God{}
		person.Team = "PSG"
		person.TeamName = "PSG"
		dict, err := Marshal(person)

		Convey("Then I should get the appropriate map", func() {
			So(err, ShouldBeNil)
			So(dict["team"], ShouldEqual, nil)
			So(dict["TeamName"], ShouldEqual, nil)
		})
	})
}

func TestMethodMarshalWithSpecialTag(t *testing.T) {

	Convey("Given I call the method cassandra.Marshal", t, func() {

		animal := &Animal{}
		animal.Type = "Dog"
		dict, err := Marshal(animal)

		Convey("Then I should get the appropriate map", func() {
			So(err, ShouldBeNil)
			So(dict["type_animal"], ShouldEqual, "Dog")
		})
	})
}

func TestMethodMarshalWithWrongSpecialTag(t *testing.T) {

	Convey("Given I call the method cassandra.Marshal", t, func() {

		phone := &Phone{}
		phone.Brand = "Apple"
		dict, err := Marshal(phone)

		Convey("Then I should get the appropriate map", func() {
			So(err, ShouldBeNil)
			So(dict["Brand"], ShouldEqual, "Apple")
		})
	})
}

func TestMethodMarshalWithStructInAStruct(t *testing.T) {

	Convey("Given I call the method cassandra.Marshal", t, func() {

		master := &Master{}
		master.Name = "Alexandre"

		animal := Animal{}
		animal.Type = "Dog"

		secondPet := &Animal{}
		secondPet.Type = "Cat"

		master.Pet = animal
		master.SecondPet = secondPet

		dict, err := Marshal(master)
		animalDict, animalErr := Marshal(animal)
		petDict, petErr := Marshal(secondPet)

		Convey("Then I should get the appropriate map", func() {
			So(err, ShouldBeNil)
			So(animalErr, ShouldBeNil)
			So(petErr, ShouldBeNil)
			So(dict["name"], ShouldEqual, "Alexandre")
			So(dict["animal"], ShouldResemble, animalDict)
			So(dict["secondPet"], ShouldResemble, petDict)

		})
	})
}

func TestMethodMarshalWithStructInAStructNotAsigned(t *testing.T) {

	Convey("Given I call the method cassandra.Marshal", t, func() {

		master := &Master{}
		master.Name = "Alexandre"

		dict, err := Marshal(master)

		Convey("Then I should get the appropriate map", func() {
			So(err, ShouldBeNil)
			So(dict["name"], ShouldEqual, "Alexandre")
			So(dict["animal"], ShouldResemble, nil)
			So(dict["secondePet"], ShouldResemble, nil)
		})
	})
}

func TestCacheOfType(t *testing.T) {

	Convey("Given I call the method cassandra.Marshal two times", t, func() {

		person := &Person{}
		person.Name = "Alexandre"
		person.Age = 2
		person.Address = "Sarralbe"
		person.Country = "France"
		person.ID = "19a65f7e-102f-11e6-8309-f45c89941b79"
		person.Siblings = []string{"Celine", "Dominique"}

		Marshal(person)
		Marshal(person)

		structVal := reflect.Indirect(reflect.ValueOf(person))

		Convey("Then a cache should be initialized", func() {
			So(fieldCache.m[structVal.Type()][0].name, ShouldEqual, "id")
			So(fieldCache.m[structVal.Type()][0].omitEmpty, ShouldEqual, false)
			So(fieldCache.m[structVal.Type()][0].quoted, ShouldEqual, false)
			So(fieldCache.m[structVal.Type()][0].tag, ShouldEqual, true)
			So(fieldCache.m[structVal.Type()][0].typ, ShouldEqual, reflect.TypeOf("string"))

			So(fieldCache.m[structVal.Type()][1].name, ShouldEqual, "name")
			So(fieldCache.m[structVal.Type()][1].omitEmpty, ShouldEqual, false)
			So(fieldCache.m[structVal.Type()][1].quoted, ShouldEqual, false)
			So(fieldCache.m[structVal.Type()][1].tag, ShouldEqual, true)
			So(fieldCache.m[structVal.Type()][1].typ, ShouldEqual, reflect.TypeOf("string"))

			So(fieldCache.m[structVal.Type()][2].name, ShouldEqual, "Age")
			So(fieldCache.m[structVal.Type()][2].omitEmpty, ShouldEqual, false)
			So(fieldCache.m[structVal.Type()][2].quoted, ShouldEqual, false)
			So(fieldCache.m[structVal.Type()][2].tag, ShouldEqual, false)
			So(fieldCache.m[structVal.Type()][2].typ, ShouldEqual, reflect.TypeOf(2))

			So(fieldCache.m[structVal.Type()][3].name, ShouldEqual, "siblings")
			So(fieldCache.m[structVal.Type()][3].omitEmpty, ShouldEqual, false)
			So(fieldCache.m[structVal.Type()][3].quoted, ShouldEqual, false)
			So(fieldCache.m[structVal.Type()][3].tag, ShouldEqual, true)
			So(fieldCache.m[structVal.Type()][3].typ, ShouldEqual, reflect.TypeOf([]string{}))

			So(fieldCache.m[structVal.Type()][4].name, ShouldEqual, "zipcode")
			So(fieldCache.m[structVal.Type()][4].omitEmpty, ShouldEqual, true)
			So(fieldCache.m[structVal.Type()][4].quoted, ShouldEqual, false)
			So(fieldCache.m[structVal.Type()][4].tag, ShouldEqual, true)
			So(fieldCache.m[structVal.Type()][4].typ, ShouldEqual, reflect.TypeOf("string"))

			So(fieldCache.m[structVal.Type()][5].name, ShouldEqual, "country")
			So(fieldCache.m[structVal.Type()][5].omitEmpty, ShouldEqual, true)
			So(fieldCache.m[structVal.Type()][5].quoted, ShouldEqual, false)
			So(fieldCache.m[structVal.Type()][5].tag, ShouldEqual, true)
			So(fieldCache.m[structVal.Type()][5].typ, ShouldEqual, reflect.TypeOf("string"))
		})
	})
}

func TestMethodMarshalErrorNoneStruct(t *testing.T) {

	Convey("Given I call the method cassandra.Marshal with a none struct object", t, func() {

		dict, err := Marshal(1)

		Convey("Then I should get the appropriate map", func() {
			So(dict, ShouldBeNil)
			So(err.Error(), ShouldEqual, "The given interface is not a struct type")
		})
	})
}

func TestMethodMarshalWithListOfStructInAStruct(t *testing.T) {

	Convey("Given I call the method cassandra.Marshal", t, func() {

		world := &World{}
		world.Name = "Earth"

		person := Person{}
		person.Name = "Alexandre"

		person1 := Person{}
		person1.Name = "Antoine"

		world.People = []Person{person, person1}

		dict, err := Marshal(world)
		p, pErr := Marshal(person)
		p1, p1Err := Marshal(person1)

		Convey("Then I should get the appropriate map", func() {
			So(err, ShouldBeNil)
			So(pErr, ShouldBeNil)
			So(p1Err, ShouldBeNil)
			So(dict["name"], ShouldEqual, "Earth")
			So(dict["people"].([]map[string]interface{})[0], ShouldResemble, p)
			So(dict["people"].([]map[string]interface{})[1], ShouldResemble, p1)
		})
	})
}
