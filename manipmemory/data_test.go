package manipmemory

import (
	"github.com/aporeto-inc/elemental"

	memdb "github.com/hashicorp/go-memdb"
)

// PersonIdentity represents the Identity of the object
var PersonIdentity = elemental.Identity{
	Name:     "person",
	Category: "persons",
}

type PersonsList []*Person

func (o PersonsList) ContentIdentity() elemental.Identity {
	return PersonIdentity
}

func (o PersonsList) List() elemental.IdentifiablesList {
	return nil
}

type Person struct {
	ID       string
	Name     string
	Siblings []string
	ZipCode  string
	Country  string
}

func (p *Person) Identifier() string {
	return p.ID
}

// Identity returns the Identity of the object.
func (p *Person) Identity() elemental.Identity {

	return PersonIdentity
}

// SetIdentifier sets the value of the object's unique identifier.
func (p *Person) SetIdentifier(ID string) {
	p.ID = ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (p *Person) Validate() error {
	return nil
}

type NotPerson struct {
	Person
}

// Identity returns the Identity of the object.
func (p *NotPerson) Identity() elemental.Identity {

	return elemental.EmptyIdentity
}

var Schema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"persons": &memdb.TableSchema{
			Name: "persons",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "ID"},
				},
				"Name": &memdb.IndexSchema{
					Name:    "Name",
					Indexer: &memdb.StringFieldIndex{Field: "Name"},
				},
			},
		},
	},
}
