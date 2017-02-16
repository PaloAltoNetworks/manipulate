package manipmongo

import (
	"fmt"
	"strconv"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	mgo "gopkg.in/mgo.v2"
)

// DoesDatabaseExist checks if the database used by the given manipulator exists.
func DoesDatabaseExist(manipulator manipulate.Manipulator) (bool, error) {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		return false, fmt.Errorf("You can only pass a Mongo Manipulator to CreateIndex")
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
		panic("You can only pass a Mongo Manipulator to CreateIndex")
	}

	session := m.rootSession.Copy()
	defer session.Close()

	return session.DB(m.dbName).DropDatabase()
}

// CreateIndex creates multiple mgo.Index for the collection storing info for the given identity using the given manipulator.
func CreateIndex(manipulator manipulate.Manipulator, identity elemental.Identity, indexes ...mgo.Index) error {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		return fmt.Errorf("You can only pass a Mongo Manipulator to CreateIndex")
	}

	session := m.rootSession.Copy()
	defer session.Close()

	collection := session.DB(m.dbName).C(identity.Name)

	for i, index := range indexes {
		index.Name = "index_" + identity.Name + "_" + strconv.Itoa(i)
		if err := collection.EnsureIndex(index); err != nil {
			return err
		}
	}

	return nil
}

// CreateCollection creates a collection using the given mgo.CollectionInfo.
func CreateCollection(manipulator manipulate.Manipulator, identity elemental.Identity, info *mgo.CollectionInfo) error {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		return fmt.Errorf("You can only pass a Mongo Manipulator to CreateIndex")
	}

	session := m.rootSession.Copy()
	defer session.Close()

	collection := session.DB(m.dbName).C(identity.Name)

	return collection.Create(info)
}

// GetSession returns a ready to use session. Use at your own risks.
// You are responsible for closing the session by calling the returner close function
func GetSession(manipulator manipulate.Manipulator) (*mgo.Database, func(), error) {

	m, ok := manipulator.(*mongoManipulator)
	if !ok {
		return nil, nil, fmt.Errorf("You can only pass a Mongo Manipulator to CreateIndex")
	}

	session := m.rootSession.Copy()

	return session.DB(m.dbName), func() { session.Close() }, nil
}
