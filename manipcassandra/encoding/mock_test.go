// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package cassandra

import "time"

type Person struct {
	ID       string `cql:"id"`
	Name     string `cql:"name"`
	Age      int
	Address  string   `cql:"-"`
	Siblings []string `cql:"siblings"`
	ZipCode  string   `cql:"zipcode,omitempty"`
	Country  string   `cql:"country,omitempty"`
	language string   `cql:"language"`
}

type Game struct {
	ID string `cql:"id"`
}

type Animal struct {
	Type string `cql:"type_animal"`
}

type Phone struct {
	Brand string `cql:"ÆÊ™"`
}

type Master struct {
	Name string `cql:"name"`
	Pet  Animal `cql:"animal"`

	SecondPet *Animal `cql:"secondPet"`
}

type God struct {
	Person

	FirstName string `cql:"firstName"`
	LastName  string `cql:"name"`

	Team     string `cql:"team"`
	TeamName string `cql:"team"`

	X string `json:"json_attribute"`
}

type World struct {
	Name   string   `cql:"name"`
	People []Person `cql:"people"`
}

type Planet struct {
	Name   string    `cql:"name"`
	People []*Person `cql:"people"`
}

type Group struct {
	Name  string   `cql:"name"`
	Owner []string `cql:"owner"`
}

type Date struct {
	CreationDate time.Time `cql:"creationDate,autotimestamp"`
	UpdateDate   time.Time `cql:"updateDate,autotimestampoverride"`
	Name         string    `cql:"name,omitempty"`
}

type LannisterName string

type Lannister struct {
	Name LannisterName `cql:"name"`
}

type Stark struct {
	ID       string `cql:"id,primarykey"`
	Name     string `cql:"name,primarykey"`
	Age      int
	Address  string   `cql:"-"`
	Siblings []string `cql:"siblings"`
}
