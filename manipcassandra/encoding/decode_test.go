// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package cassandra

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMethodUnmarshalWithListOfStructAndPointerArray(t *testing.T) {
	Convey("Given I call the method cassandra.Unmarshal", t, func() {

		person1 := make(map[string]interface{})
		person1["name"] = "Alexandre"
		person1["Age"] = 2
		person1["Address"] = "Sarralbe"
		person1["id"] = "19a65f7e-102f-11e6-8309-f45c89941b79"
		person1["siblings"] = []string{"Celine", "Dominique"}
		person1["country"] = "France"
		person1["language"] = "French"

		person2 := make(map[string]interface{})
		person2["name"] = "Antoine"
		person2["Age"] = 3
		person2["Address"] = "Paris"
		person2["id"] = "29a65f7e-102f-11e6-8309-f45c89941b79"
		person2["siblings"] = []string{"Cyril"}
		person2["country"] = "Allemagne"
		person2["language"] = "German"

		persons := []map[string]interface{}{person1, person2}

		var results []*Person
		err := Unmarshal(persons, &results)

		Convey("Then I should get the appropriate object", func() {
			So(err, ShouldBeNil)
			So(results[0].Name, ShouldEqual, "Alexandre")
			So(results[0].Age, ShouldEqual, 2)
			So(results[0].Address, ShouldEqual, "")
			So(results[0].ID, ShouldEqual, "19a65f7e-102f-11e6-8309-f45c89941b79")
			So(results[0].Siblings, ShouldResemble, []string{"Celine", "Dominique"})
			So(results[0].Country, ShouldEqual, "France")
			So(results[0].ZipCode, ShouldEqual, "")
			So(results[0].language, ShouldEqual, "")

			So(results[1].Name, ShouldEqual, "Antoine")
			So(results[1].Age, ShouldEqual, 3)
			So(results[1].Address, ShouldEqual, "")
			So(results[1].ID, ShouldEqual, "29a65f7e-102f-11e6-8309-f45c89941b79")
			So(results[1].Siblings, ShouldResemble, []string{"Cyril"})
			So(results[1].Country, ShouldEqual, "Allemagne")
			So(results[1].ZipCode, ShouldEqual, "")
			So(results[1].language, ShouldEqual, "")
		})
	})
}

func TestMethodUnmarshalWithListOfStruct(t *testing.T) {
	Convey("Given I call the method cassandra.Unmarshal", t, func() {

		person1 := make(map[string]interface{})
		person1["name"] = "Alexandre"
		person1["Age"] = 2
		person1["Address"] = "Sarralbe"
		person1["id"] = "19a65f7e-102f-11e6-8309-f45c89941b79"
		person1["siblings"] = []string{"Celine", "Dominique"}
		person1["country"] = "France"
		person1["language"] = "French"

		person2 := make(map[string]interface{})
		person2["name"] = "Antoine"
		person2["Age"] = 3
		person2["Address"] = "Paris"
		person2["id"] = "29a65f7e-102f-11e6-8309-f45c89941b79"
		person2["siblings"] = []string{"Cyril"}
		person2["country"] = "Allemagne"
		person2["language"] = "German"

		persons := []map[string]interface{}{person1, person2}

		var results []Person
		err := Unmarshal(persons, &results)

		Convey("Then I should get the appropriate object", func() {
			So(err, ShouldBeNil)
			So(results[0].Name, ShouldEqual, "Alexandre")
			So(results[0].Age, ShouldEqual, 2)
			So(results[0].Address, ShouldEqual, "")
			So(results[0].ID, ShouldEqual, "19a65f7e-102f-11e6-8309-f45c89941b79")
			So(results[0].Siblings, ShouldResemble, []string{"Celine", "Dominique"})
			So(results[0].Country, ShouldEqual, "France")
			So(results[0].ZipCode, ShouldEqual, "")
			So(results[0].language, ShouldEqual, "")

			So(results[1].Name, ShouldEqual, "Antoine")
			So(results[1].Age, ShouldEqual, 3)
			So(results[1].Address, ShouldEqual, "")
			So(results[1].ID, ShouldEqual, "29a65f7e-102f-11e6-8309-f45c89941b79")
			So(results[1].Siblings, ShouldResemble, []string{"Cyril"})
			So(results[1].Country, ShouldEqual, "Allemagne")
			So(results[1].ZipCode, ShouldEqual, "")
			So(results[1].language, ShouldEqual, "")
		})
	})
}

func TestMethodUnmarshalWithSimpleStruct(t *testing.T) {

	Convey("Given I call the method cassandra.Unmarshal", t, func() {

		dict := make(map[string]interface{})
		dict["name"] = "Alexandre"
		dict["Age"] = 2
		dict["Address"] = "Sarralbe"
		dict["id"] = "19a65f7e-102f-11e6-8309-f45c89941b79"
		dict["siblings"] = []string{"Celine", "Dominique"}
		dict["country"] = "France"
		dict["language"] = "French"

		var person Person
		err := Unmarshal(dict, &person)

		Convey("Then I should get the appropriate object", func() {
			So(err, ShouldBeNil)
			So(person.Name, ShouldEqual, "Alexandre")
			So(person.Age, ShouldEqual, 2)
			So(person.Address, ShouldEqual, "")
			So(person.ID, ShouldEqual, "19a65f7e-102f-11e6-8309-f45c89941b79")
			So(person.Siblings, ShouldResemble, []string{"Celine", "Dominique"})
			So(person.Country, ShouldEqual, "France")
			So(person.ZipCode, ShouldEqual, "")
			So(person.language, ShouldEqual, "")
		})
	})
}

func TestMethodUnmarshalWithSimpleStructWithGocqlUUID(t *testing.T) {

	Convey("Given I call the method cassandra.Unmarshal", t, func() {

		dict := make(map[string]interface{})
		dict["id"] = "19a65f7e-102f-11e6-8309-f45c89941b79"

		var game Game
		err := Unmarshal(dict, &game)

		Convey("Then I should get the appropriate object", func() {
			So(err, ShouldBeNil)
			So(game.ID, ShouldEqual, "19a65f7e-102f-11e6-8309-f45c89941b79")
		})
	})
}

func TestMethodUnmarshalStructInStruct(t *testing.T) {

	Convey("Given I call the method cassandra.Unmarshal", t, func() {

		dictAnimal := make(map[string]interface{})
		dictAnimal["type_animal"] = "Dog"

		dictSecondPet := make(map[string]interface{})
		dictSecondPet["type_animal"] = "Cat"

		dictMaster := make(map[string]interface{})
		dictMaster["name"] = "Alexandre"
		dictMaster["animal"] = dictAnimal
		dictMaster["secondPet"] = dictSecondPet

		var master Master
		err := Unmarshal(dictMaster, &master)

		Convey("Then I should get the appropriate object", func() {
			So(err, ShouldBeNil)
			So(master.Name, ShouldEqual, "Alexandre")
			So(master.Pet.Type, ShouldEqual, "Dog")
			So(master.SecondPet.Type, ShouldEqual, "Cat")
		})
	})
}

func TestMethodUnmarshalErrorType(t *testing.T) {

	Convey("Given I call the method cassandra.Unmarshal", t, func() {

		dictMaster := make(map[string]interface{})
		dictMaster["name"] = "Alexandre"

		var master []Master
		err := Unmarshal(dictMaster, master)

		Convey("Then I should get the error", func() {
			So(err.Error(), ShouldEqual, `The given data should be an array`)
		})
	})
}

func TestMethodUnmarshalWithStructAndListOfStruct(t *testing.T) {
	Convey("Given I call the method cassandra.Unmarshal", t, func() {

		person1 := make(map[string]interface{})
		person1["name"] = "Alexandre"
		person1["Age"] = 2
		person1["Address"] = "Sarralbe"
		person1["id"] = "19a65f7e-102f-11e6-8309-f45c89941b79"
		person1["siblings"] = []string{"Celine", "Dominique"}
		person1["country"] = "France"
		person1["language"] = "French"

		person2 := make(map[string]interface{})
		person2["name"] = "Antoine"
		person2["Age"] = 3
		person2["Address"] = "Paris"
		person2["id"] = "29a65f7e-102f-11e6-8309-f45c89941b79"
		person2["siblings"] = []string{"Cyril"}
		person2["country"] = "Allemagne"
		person2["language"] = "German"

		world := make(map[string]interface{})
		world["name"] = "Earth"
		world["people"] = []map[string]interface{}{person1, person2}

		var result World
		err := Unmarshal(world, &result)

		Convey("Then I should get the appropriate object", func() {
			So(err, ShouldBeNil)
			So(result.Name, ShouldEqual, "Earth")
			So(result.People[0].Name, ShouldEqual, "Alexandre")
			So(result.People[0].Age, ShouldEqual, 2)
			So(result.People[0].Address, ShouldEqual, "")
			So(result.People[0].ID, ShouldEqual, "19a65f7e-102f-11e6-8309-f45c89941b79")
			So(result.People[0].Siblings, ShouldResemble, []string{"Celine", "Dominique"})
			So(result.People[0].Country, ShouldEqual, "France")

			So(result.People[1].Name, ShouldEqual, "Antoine")
			So(result.People[1].Age, ShouldEqual, 3)
			So(result.People[1].Address, ShouldEqual, "")
			So(result.People[1].ID, ShouldEqual, "29a65f7e-102f-11e6-8309-f45c89941b79")
			So(result.People[1].Siblings, ShouldResemble, []string{"Cyril"})
			So(result.People[1].Country, ShouldEqual, "Allemagne")

		})
	})
}

func TestMethodUnmarshalWithStructAndListOfPointer(t *testing.T) {
	Convey("Given I call the method cassandra.Unmarshal", t, func() {

		person1 := make(map[string]interface{})
		person1["name"] = "Alexandre"
		person1["Age"] = 2
		person1["Address"] = "Sarralbe"
		person1["id"] = "19a65f7e-102f-11e6-8309-f45c89941b79"
		person1["siblings"] = []string{"Celine", "Dominique"}
		person1["country"] = "France"
		person1["language"] = "French"

		person2 := make(map[string]interface{})
		person2["name"] = "Antoine"
		person2["Age"] = 3
		person2["Address"] = "Paris"
		person2["id"] = "29a65f7e-102f-11e6-8309-f45c89941b79"
		person2["siblings"] = []string{"Cyril"}
		person2["country"] = "Allemagne"
		person2["language"] = "German"

		world := make(map[string]interface{})
		world["name"] = "Earth"
		world["people"] = []map[string]interface{}{person1, person2}

		var result Planet
		err := Unmarshal(world, &result)

		Convey("Then I should get the appropriate object", func() {
			So(err, ShouldBeNil)
			So(result.Name, ShouldEqual, "Earth")
			So(result.People[0].Name, ShouldEqual, "Alexandre")
			So(result.People[0].Age, ShouldEqual, 2)
			So(result.People[0].Address, ShouldEqual, "")
			So(result.People[0].ID, ShouldEqual, "19a65f7e-102f-11e6-8309-f45c89941b79")
			So(result.People[0].Siblings, ShouldResemble, []string{"Celine", "Dominique"})
			So(result.People[0].Country, ShouldEqual, "France")

			So(result.People[1].Name, ShouldEqual, "Antoine")
			So(result.People[1].Age, ShouldEqual, 3)
			So(result.People[1].Address, ShouldEqual, "")
			So(result.People[1].ID, ShouldEqual, "29a65f7e-102f-11e6-8309-f45c89941b79")
			So(result.People[1].Siblings, ShouldResemble, []string{"Cyril"})
			So(result.People[1].Country, ShouldEqual, "Allemagne")

		})
	})
}

func TestMethodUnmarshalWithSpecificType(t *testing.T) {

	Convey("Given I call the method cassandra.Unmarshal", t, func() {

		dict := make(map[string]interface{})
		dict["name"] = "Tyrion"

		var lannister Lannister
		err := Unmarshal(dict, &lannister)

		Convey("Then I should get the appropriate object", func() {
			So(err, ShouldBeNil)
			So(lannister.Name, ShouldEqual, "Tyrion")
		})
	})
}
