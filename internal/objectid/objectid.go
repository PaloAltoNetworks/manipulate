package objectid

import (
	"encoding/hex"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Parse parses the given string and returns
// a primitive.ObjectID and true if the given value is valid,
// otherwise it will return a nil ObjectID and false.
//
// This is copy of the primative.ObjectIDFromHex function, but
// prevents doing double decode in the classic procedure:
// if primative.IsValidObjectID(x) { primative.ObjectIDFromHex(x} }
func Parse(s string) (primitive.ObjectID, bool) {

	if len(s) != 24 {
		return primitive.NilObjectID, false
	}

	d, err := hex.DecodeString(s)
	if err != nil || len(d) != 12 {
		return primitive.NilObjectID, false
	}

	return primitive.ObjectID(d), true
}
