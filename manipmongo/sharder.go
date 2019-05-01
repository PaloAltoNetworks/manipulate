// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
