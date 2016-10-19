package maniptest

import "github.com/aporeto-inc/elemental"

// UserIdentity represents the Identity of the object
var PersonIdentity = elemental.Identity{
	Name:     "person",
	Category: "persons",
}

type Person struct {
	ID       string   `bson:"_id"`
	Name     string   `bson:"name"`
	Siblings []string `bson:"siblings"`
	ZipCode  string   `bson:"zipcode"`
	Country  string   `bson:"country"`
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
func (p *Person) Validate() elemental.Errors {
	return nil
}
