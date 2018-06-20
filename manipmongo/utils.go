package manipmongo

import (
	"strings"

	"go.aporeto.io/elemental"
	"github.com/globalsign/mgo"
)

// collectionFromIdentity returns the mgo*.Collection associated to the given Identity from the
// given *mgo.Database.
func collectionFromIdentity(db *mgo.Database, identity elemental.Identity, prefix string) *mgo.Collection {

	if prefix != "" {
		return db.C(prefix + "-" + identity.Name)
	}

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

func applyOrdering(order []string, inverted bool) []string {

	var o []string

	for _, key := range order {
		o = append(o, strings.ToLower(invertSortKey(key, inverted)))
	}

	return o
}
