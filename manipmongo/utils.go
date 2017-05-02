package manipmongo

import (
	"strings"

	"github.com/aporeto-inc/elemental"

	mgo "gopkg.in/mgo.v2"
)

// collectionFromIdentity returns the mgo*.Collection associated to the given Identity from the
// given *mgo.Database.
func collectionFromIdentity(db *mgo.Database, identity elemental.Identity) *mgo.Collection {

	return db.C(identity.Name)
}

// invertSortKey eventually inverts the given sorting key.
func invertSortKey(k string, revert bool) string {

	if !revert {
		return k
	}

	if strings.HasPrefix(k, "-") {
		return k[1:]
	}

	return "-" + k
}
