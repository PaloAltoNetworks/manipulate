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
			},
		},
	},
}
