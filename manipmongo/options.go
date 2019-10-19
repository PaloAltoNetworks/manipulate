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
	"crypto/tls"
	"time"

	"github.com/globalsign/mgo/bson"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

// An Option represents a maniphttp.Manipulator option.
type Option func(*config)

type config struct {
	username           string
	password           string
	authsource         string
	tlsConfig          *tls.Config
	poolLimit          int
	connectTimeout     time.Duration
	socketTimeout      time.Duration
	readConsistency    manipulate.ReadConsistency
	writeConsistency   manipulate.WriteConsistency
	sharder            Sharder
	defaultRetryFunc   manipulate.RetryFunc
	forcedReadFilter   bson.M
	attributeEncrypter elemental.AttributeEncrypter
	explain            map[elemental.Identity]map[elemental.Operation]struct{}
}

func newConfig() *config {
	return &config{
		poolLimit:        4096,
		connectTimeout:   10 * time.Second,
		socketTimeout:    60 * time.Second,
		readConsistency:  manipulate.ReadConsistencyDefault,
		writeConsistency: manipulate.WriteConsistencyDefault,
	}
}

// OptionCredentials sets the username and password to use for authentication.
func OptionCredentials(username, password, authsource string) Option {
	return func(c *config) {
		c.username = username
		c.password = password
		c.authsource = authsource
	}
}

// OptionTLS sets the tls configuration for the connection.
func OptionTLS(tlsConfig *tls.Config) Option {
	return func(c *config) {
		c.tlsConfig = tlsConfig
	}
}

// OptionConnectionPoolLimit sets maximum size of the connection pool.
func OptionConnectionPoolLimit(poolLimit int) Option {
	return func(c *config) {
		c.poolLimit = poolLimit
	}
}

// OptionConnectionTimeout sets the connection timeout.
func OptionConnectionTimeout(connectTimeout time.Duration) Option {
	return func(c *config) {
		c.connectTimeout = connectTimeout
	}
}

// OptionSocketTimeout sets the socket timeout.
func OptionSocketTimeout(socketTimeout time.Duration) Option {
	return func(c *config) {
		c.socketTimeout = socketTimeout
	}
}

// OptionDefaultReadConsistencyMode sets the default read consistency mode.
func OptionDefaultReadConsistencyMode(consistency manipulate.ReadConsistency) Option {
	return func(c *config) {
		c.readConsistency = consistency
	}
}

// OptionDefaultWriteConsistencyMode sets the default write consistency mode.
func OptionDefaultWriteConsistencyMode(consistency manipulate.WriteConsistency) Option {
	return func(c *config) {
		c.writeConsistency = consistency
	}
}

// OptionSharder sets the sharder.
func OptionSharder(sharder Sharder) Option {
	return func(c *config) {
		c.sharder = sharder
	}
}

// OptionDefaultRetryFunc sets the default retry func to use
// if manipulate.Context does not have one.
func OptionDefaultRetryFunc(f manipulate.RetryFunc) Option {
	return func(c *config) {
		c.defaultRetryFunc = f
	}
}

// OptionForceReadFilter allows to set a bson.M filter that
// will always reducing the scope of the reads to that filter.
func OptionForceReadFilter(f bson.M) Option {
	return func(c *config) {
		c.forcedReadFilter = f
	}
}

// OptionAttributeEncrypter allows to set an elemental.AttributeEncrypter
// to use to encrypt/decrypt elemental.AttributeEncryptable.
func OptionAttributeEncrypter(enc elemental.AttributeEncrypter) Option {
	return func(c *config) {
		c.attributeEncrypter = enc
	}
}

// OptionExplain allows to tell manipmongo to explain the query before it
// runs it for the given identities on the given operations.
// For example, consider passing:
//      map[elemental.Identity][]elemental.Operation{
//          model.ThisIndentity: []elemental.Operation{elemental.OperationRetrieveMany, elemental.OperationCreate},
//          model.ThatIndentity: []elemental.Operation{}, // or nil
//      }
//
//  This would trigger explanation on retrieveMany and create for model.ThisIndentity
//  and every operation on model.ThatIndentity.
func OptionExplain(explain map[elemental.Identity]map[elemental.Operation]struct{}) Option {
	return func(c *config) {
		c.explain = explain
	}
}

const opaqueKeyUpsert = "manipmongo.upsert"

type opaquer interface {
	Opaque() map[string]interface{}
}

// ContextOptionUpsert tells to use upsert for an Create operation.
// The given operation will be executed for the upsert command.
// You cannot use "$set" which is always set to be the identifier.
// If you do so, ContextOptionUpsert will panic.
// If you use $setOnInsert, you must not set _id. If you do so,
// it will panic.
func ContextOptionUpsert(operations bson.M) manipulate.ContextOption {

	if _, ok := operations["$set"]; ok {
		panic("cannot use $set in upsert operations")
	}

	if soi, ok := operations["$setOnInsert"]; ok {
		for k := range soi.(bson.M) {
			if k == "_id" {
				panic("cannot use $setOnInsert on _id in upsert operations")
			}
		}
	}

	return func(c manipulate.Context) {
		c.(opaquer).Opaque()[opaqueKeyUpsert] = operations
	}
}
