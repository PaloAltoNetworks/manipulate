package manipmongo

import (
	"github.com/aporeto-inc/elemental"

	mgo "gopkg.in/mgo.v2"
)

// collectionFromIdentity returns the mgo*.Collection associated to the given Identity from the
// given *mgo.Database.
func collectionFromIdentity(db *mgo.Database, identity elemental.Identity) *mgo.Collection {

	return db.C(identity.Category)
}
