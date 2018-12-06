package manipmongo

import (
	"github.com/globalsign/mgo/bson"
	"go.aporeto.io/elemental"
)

// A Sharder is the interface of an object that can be use
// to manage sharding of resources.
type Sharder interface {

	// Shard will be call when the shard key needs to be set to
	// the given elemental.Identifiable.
	Shard(elemental.Identifiable)

	// FilterOne returns the filter bit as bson.M that must be
	// used to perform an efficient localized query for a single object.
	//
	// You can return nil which will trigger a broadcast.
	FilterOne(elemental.Identifiable) bson.M

	// FilterMany returns the filter bit as bson.M that must be
	// used to perform an efficient localized query for multiple objects.
	//
	// You can return nil which will trigger a broadcast.
	FilterMany(elemental.Identity) bson.M
}
