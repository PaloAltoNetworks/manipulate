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
	"fmt"
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/manipmongo/internal/compiler"
)

// CompileFilter compiles the given manipulate filter into a raw mongo filter.
func CompileFilter(f *elemental.Filter) bson.M {
	return compiler.CompileFilter(f)
}

// DoesDatabaseExist checks if the database used by the given manipulator exists.
func DoesDatabaseExist(manipulator manipulate.Manipulator) (bool, error) {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to DoesDatabaseExist")
	}

	dbs, err := m.rootSession.DatabaseNames()
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
		panic("you can only pass a mongo manipulator to CreateIndex")
	}

	session := m.rootSession.Copy()
	defer session.Close()

	return session.DB(m.dbName).DropDatabase()
}

// CreateIndex creates multiple mgo.Index for the collection storing info for the given identity using the given manipulator.
func CreateIndex(manipulator manipulate.Manipulator, identity elemental.Identity, indexes ...mgo.Index) error {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to CreateIndex")
	}

	session := m.rootSession.Copy()
	defer session.Close()

	collection := session.DB(m.dbName).C(identity.Name)

	for i, index := range indexes {
		if index.Name == "" {
			index.Name = "index_" + identity.Name + "_" + strconv.Itoa(i)
		}
		if err := collection.EnsureIndex(index); err != nil {
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

	session := m.rootSession.Copy()
	defer session.Close()

	collection := session.DB(m.dbName).C(identity.Name)

	for _, index := range indexes {
		if err := collection.DropIndexName(index); err != nil {
			return err
		}
	}

	return nil
}

// CreateCollection creates a collection using the given mgo.CollectionInfo.
func CreateCollection(manipulator manipulate.Manipulator, identity elemental.Identity, info *mgo.CollectionInfo) error {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to CreateCollection")
	}

	session := m.rootSession.Copy()
	defer session.Close()

	collection := session.DB(m.dbName).C(identity.Name)

	return collection.Create(info)
}

// GetDatabase returns a ready to use mgo.Database. Use at your own risks.
// You are responsible for closing the session by calling the returner close function
func GetDatabase(manipulator manipulate.Manipulator) (*mgo.Database, func(), error) {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a mongo manipulator to GetSession")
	}

	session := m.rootSession.Copy()

	return session.DB(m.dbName), func() { session.Close() }, nil
}

// SetConsistencyMode sets the mongo consistency mode of the mongo session.
func SetConsistencyMode(manipulator manipulate.Manipulator, mode mgo.Mode, refresh bool) {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		panic("you can only pass a Mongo Manipulator to SetConsistencyMode")
	}

	if m.rootSession == nil {
		panic("cannot apply SetConsistencyMode. The root mongo session is not ready")
	}

	m.rootSession.SetMode(mode, refresh)
}
