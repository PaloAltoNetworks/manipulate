package objectid

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Parse parses the given string and returns
// a primitive.ObjectID and true if the given value is valid,
// otherwise it will return an empty primitive.ObjectID and false.
func Parse(s string) (primitive.ObjectID, bool) {

	if len(s) != 24 {
		return primitive.NilObjectID, false
	}

	objectID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		return primitive.NilObjectID, false
	}
	return objectID, true
}
