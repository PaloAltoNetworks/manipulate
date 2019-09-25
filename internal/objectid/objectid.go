package objectid

import (
	"encoding/hex"

	"github.com/globalsign/mgo/bson"
)

// Parse parses the given string and returns
// a bson.ObjectId and true if the given value is valid,
// otherwise it will return an empty bson.ObjectId and false.
//
// This is copy of the bson.ObjectIdHex function, but
// prevents doing double decode in the classic procedure:
// if bson.IsObjectId(x) { bson.ObjectIdHex(x} }
func Parse(s string) (bson.ObjectId, bool) {

	if len(s) != 24 {
		return bson.ObjectId(""), false
	}

	d, err := hex.DecodeString(s)
	if err != nil || len(d) != 12 {
		return bson.ObjectId(""), false
	}

	return bson.ObjectId(d), true
}
