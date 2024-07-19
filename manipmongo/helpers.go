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
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/backoff"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DoesDatabaseExist checks if the database used by the given manipulator exists.
func DoesDatabaseExist(manipulator manipulate.Manipulator) (bool, error) {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to DoesDatabaseExist")
	}

	dbs, err := m.client.ListDatabaseNames(context.Background(), nil)
	if err != nil {
		return false, err
	}

	for _, db := range dbs {
		if db == m.dbName {
			return true, nil
		}
	}

	return false, nil
}

// DropDatabase drops the entire database used by the given manipulator.
func DropDatabase(manipulator manipulate.Manipulator) error {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to DropDatabase")
	}
	database := m.client.Database(m.dbName)
	return database.Drop(context.Background())
}

// CreateIndex creates multiple index for the collection storing info for the given identity using the given manipulator.
func CreateIndex(manipulator manipulate.Manipulator, identity elemental.Identity, indexes ...mgo.Index) error {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to CreateIndex")
	}

	database := m.client.Database(m.dbName)
	collection := database.Collection(identity.Name)

	indexModels := make([]mongo.IndexModel, len(indexes))
	for i, index := range indexes {
		if index.Name == "" {
			index.Name = "index_" + identity.Name + "_" + strconv.Itoa(i)
		}
		indexModels = append(indexModels, mongo.IndexModel{
			Keys:    index.Key,
			Options: options.Index().SetName(index.Name).SetUnique(index.Unique).SetSparse(index.Sparse).SetExpireAfterSeconds(int32(index.ExpireAfter.Seconds())),
		})
	}
	_, err := collection.Indexes().CreateMany(context.Background(), indexModels)
	if err != nil {
		return fmt.Errorf("unable to ensure indexes: %s", err)
	}

	return nil
}

// EnsureIndex works like create index, but it will delete existing index
// if they changed before creating a new one.
func EnsureIndex(manipulator manipulate.Manipulator, identity elemental.Identity, indexes ...mgo.Index) error {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to CreateIndex")
	}

	database := m.client.Database(m.dbName)
	collection := database.Collection(identity.Name)
	// TODO: mgo driver based code was using below to set strong consistency here. Figure out how
	// to do that with mongo-go-driver on an existing mongo client
	// session := m.rootSession.Copy()
	// session.SetMode(mgo.Strong, false)
	// session.EnsureSafe(&mgo.Safe{})

	for i, index := range indexes {
		if index.Name == "" {
			index.Name = "index_" + identity.Name + "_" + strconv.Itoa(i)
		}
		var keys bson.D
		for _, key := range index.Key {
			keys = append(keys, bson.E{Key: key, Value: 1}) // Use 1 for ascending order
		}
		indexModel := mongo.IndexModel{
			Keys:    keys,
			Options: options.Index().SetName(index.Name).SetUnique(index.Unique).SetSparse(index.Sparse).SetExpireAfterSeconds(int32(index.ExpireAfter.Seconds())),
		}
		_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
		if err != nil {
			if strings.Contains(err.Error(), "already exists with different options") {

				// In case we are changing a TTL we are using colMod instead
				// as per https://docs.mongodb.com/manual/core/index-ttl/#restrictions
				if index.ExpireAfter > 0 {

					modifyCmd := bson.D{
						{Key: "collMod", Value: collection.Name()},
						{Key: "index", Value: bson.M{"name": index.Name, "expireAfterSeconds": int(index.ExpireAfter.Seconds())}},
					}
					database := m.client.Database(m.dbName)
					err := database.RunCommand(context.Background(), modifyCmd).Err()
					if err != nil {
						return fmt.Errorf("cannot update TTL index: %s", err)
					}

				} else {

					_, err := collection.Indexes().DropOne(context.Background(), index.Name)
					if err != nil {
						return fmt.Errorf("cannot delete previous index: %s", err)
					}

					_, err = collection.Indexes().CreateOne(context.Background(), indexModel)
					if err != nil {
						return fmt.Errorf("unable to ensure index after dropping old one '%s': %s", index.Name, err)
					}

				}

				continue
			}

			return fmt.Errorf("unable to ensure index '%s': %s", index.Name, err)
		}
	}

	return nil
}

// DeleteIndex deletes multiple mgo.Index for the collection.
func DeleteIndex(manipulator manipulate.Manipulator, identity elemental.Identity, indexes ...string) error {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to DeleteIndex")
	}

	database := m.client.Database(m.dbName)
	collection := database.Collection(identity.Name)

	for _, index := range indexes {
		_, err := collection.Indexes().DropOne(context.Background(), index)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateCollection creates a collection using the given mgo.CollectionInfo.
func CreateCollection(manipulator manipulate.Manipulator, identity elemental.Identity, info *options.CreateCollectionOptions) error {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to CreateCollection")
	}

	database := m.client.Database(m.dbName)
	err := database.CreateCollection(context.Background(), identity.Name, info)
	if err != nil {
		return fmt.Errorf("unable to create collection '%s': %w", identity.Name, err)
	}

	return nil
}

// GetDatabase returns a ready to use mongo.Database. Use at your own risks.
func GetDatabase(manipulator manipulate.Manipulator) *mongo.Database {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to GetDatabase")
	}

	return m.client.Database(m.dbName)
}

// SetConsistencyMode sets the mongo consistency mode of the mongo session.
func SetConsistencyMode(manipulator manipulate.Manipulator, mode mgo.Mode, refresh bool) {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to SetConsistencyMode")
	}

	// if m.rootSession == nil {
	// 	panic("cannot apply SetConsistencyMode. The root mongo session is not ready")
	// }

	// m.rootSession.SetMode(mode, refresh)

	switch mode {
	case mgo.Strong:
		// Set strong consistency using majority read concern and primary read preference
		opts := options.Client().SetReadConcern(readconcern.Majority()).SetReadPreference(readpref.Primary())
		if refresh {
			var err error
			err = m.client.Disconnect(context.Background())
			if err != nil {
				panic(fmt.Sprintf("unable to disconnect for connecting with new consistency mode: %v", err))
			}
			m.client, err = mongo.Connect(context.Background(), opts)
			if err != nil {
				panic(fmt.Sprintf("unable to reconnect with new consistency mode: %v", err))
			}
		}

	case mgo.Monotonic:
		// Set monotonic consistency using local read concern and nearest read preference
		opts := options.Client().SetReadConcern(readconcern.Local()).SetReadPreference(readpref.Nearest())
		if refresh {
			var err error
			err = m.client.Disconnect(context.Background())
			if err != nil {
				panic(fmt.Sprintf("unable to disconnect for connecting with new consistency mode: %v", err))
			}
			m.client, err = mongo.Connect(context.Background(), opts)
			if err != nil {
				panic(fmt.Sprintf("unable to reconnect with new consistency mode: %v", err))
			}
		}

	case mgo.Eventual:
		// Set eventual consistency using available read concern and secondaryPreferred read preference
		opts := options.Client().SetReadConcern(readconcern.Available()).SetReadPreference(readpref.SecondaryPreferred())
		if refresh {
			var err error
			err = m.client.Disconnect(context.Background())
			if err != nil {
				panic(fmt.Sprintf("unable to disconnect for connecting with new consistency mode: %v", err))
			}
			m.client, err = mongo.Connect(context.Background(), opts)
			if err != nil {
				panic(fmt.Sprintf("unable to reconnect with new consistency mode: %v", err))
			}
		}

	default:
		panic("unsupported consistency mode")
	}
}

// RunQuery runs a function that must run a mongodb operation.
// It will retry in case of failure. This is an advanced helper can
// be used when you get a session from using GetDatabase().
func RunQuery(mctx manipulate.Context, operationFunc func() (any, error), baseRetryInfo RetryInfo) (any, error) {

	var try int

	for {

		out, err := operationFunc()
		if err == nil {
			return out, nil
		}

		err = HandleQueryError(err)
		if !manipulate.IsCannotCommunicateError(err) {
			return out, err
		}

		baseRetryInfo.try = try
		baseRetryInfo.err = err
		baseRetryInfo.mctx = mctx

		if rf := mctx.RetryFunc(); rf != nil {
			if rerr := rf(baseRetryInfo); rerr != nil {
				return nil, rerr
			}
		} else if baseRetryInfo.defaultRetryFunc != nil {
			if rerr := baseRetryInfo.defaultRetryFunc(baseRetryInfo); rerr != nil {
				return nil, rerr
			}
		}

		select {
		case <-mctx.Context().Done():
			return nil, manipulate.ErrCannotExecuteQuery{Err: mctx.Context().Err()}
		default:
		}

		deadline, _ := mctx.Context().Deadline()
		time.Sleep(backoff.NextWithCurve(try, deadline, defaultBackoffCurve))
		try++
	}
}

// SetAttributeEncrypter switch the attribute encrypter of the given mongo manipulator.
// This is only useful in some rare cases like miugration, and it is not go routine safe.
func SetAttributeEncrypter(manipulator manipulate.Manipulator, enc elemental.AttributeEncrypter) {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to SetAttributeEncrypter")
	}

	m.attributeEncrypter = enc
}

// GetAttributeEncrypter returns the attribute encrypter of the given mongo manipulator..
func GetAttributeEncrypter(manipulator manipulate.Manipulator) elemental.AttributeEncrypter {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to GetAttributeEncrypter")
	}

	return m.attributeEncrypter
}

// IsUpsert returns True if the mongo request is an Upsert operation, false otherwise.
func IsUpsert(mctx manipulate.Context) bool {
	_, upsert := mctx.(opaquer).Opaque()[opaqueKeyUpsert]
	return upsert
}

// IsMongoManipulator returns true if this is a mongo manipulator
func IsMongoManipulator(manipulator manipulate.Manipulator) bool {
	_, ok := manipulator.(*mongoManipulator)

	return ok
}
