package manipmongo

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"gopkg.in/mgo.v2/bson"

	log "github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
)

type bulksRegistry map[manipulate.TransactionID]*mgo.Bulk

// MongoStore represents a MongoDB session.
type MongoStore struct {
	session *mgo.Session
	db      *mgo.Database
	dbName  string
	url     string

	bulksRegistry     bulksRegistry
	bulksRegistryLock *sync.Mutex
}

// NewMongoStore returns a new *MongoStore
func NewMongoStore(url string, dbName string) *MongoStore {

	session, err := mgo.Dial(url)
	if err != nil {
		log.WithFields(log.Fields{
			"package": "manipmongo",
			"url":     url,
			"db":      dbName,
			"error":   err.Error(),
		}).Fatal("Cannot connect to mongo.")
	}

	return &MongoStore{
		url:               url,
		dbName:            dbName,
		bulksRegistry:     bulksRegistry{},
		bulksRegistryLock: &sync.Mutex{},
		session:           session,
		db:                session.DB(dbName),
	}
}

// Create is part of the implementation of the Manipulator interface.
func (s *MongoStore) Create(contexts manipulate.Contexts, parent manipulate.Manipulable, children ...manipulate.Manipulable) error {

	collection := collectionFromIdentity(s.db, children[0].Identity())
	context := manipulate.ContextForIndex(contexts, 0)
	tid := context.TransactionID
	bulk := s.bulkForID(tid, collection)

	for _, child := range children {
		child.SetIdentifier(uuid.NewV4().String())
		bulk.Insert(child)
	}

	if tid == "" {
		if err := s.commitBulk(bulk); err != nil {
			return elemental.NewError("Error", err.Error(), "", 5000)
		}
	}

	return nil
}

// Retrieve is part of the implementation of the Manipulator interface.
func (s *MongoStore) Retrieve(contexts manipulate.Contexts, objects ...manipulate.Manipulable) error {

	collection := collectionFromIdentity(s.db, objects[0].Identity())

	for i := 0; i < len(objects); i++ {
		query := collection.Find(bson.M{"_id": objects[i].Identifier()})
		if err := query.One(objects[i]); err != nil {
			return elemental.NewError("Error", err.Error(), "", 5000)
		}
	}

	return nil
}

// Update is part of the implementation of the Manipulator interface.
func (s *MongoStore) Update(contexts manipulate.Contexts, objects ...manipulate.Manipulable) error {

	collection := collectionFromIdentity(s.db, objects[0].Identity())
	context := manipulate.ContextForIndex(contexts, 0)
	tid := context.TransactionID
	bulk := s.bulkForID(tid, collection)

	for i := 0; i < len(objects); i++ {
		bulk.Update(bson.M{"_id": objects[i].Identifier()}, objects[i])
	}

	if tid == "" {
		if err := s.commitBulk(bulk); err != nil {
			return elemental.NewError("Error", err.Error(), "manipulate", 5000)
		}
	}

	return nil
}

// Delete is part of the implementation of the Manipulator interface.
func (s *MongoStore) Delete(contexts manipulate.Contexts, objects ...manipulate.Manipulable) error {

	collection := collectionFromIdentity(s.db, objects[0].Identity())
	context := manipulate.ContextForIndex(contexts, 0)
	tid := context.TransactionID
	bulk := s.bulkForID(tid, collection)

	for i := 0; i < len(objects); i++ {
		bulk.Remove(bson.M{"_id": objects[i].Identifier()})
	}

	if tid == "" {
		if err := s.commitBulk(bulk); err != nil {
			return elemental.NewError("Error", err.Error(), "manipulate", 5000)
		}
	}

	return nil
}

// RetrieveChildren is part of the implementation of the Manipulator interface.
func (s *MongoStore) RetrieveChildren(contexts manipulate.Contexts, parent manipulate.Manipulable, identity elemental.Identity, dest interface{}) error {

	collection := collectionFromIdentity(s.db, identity)
	context := manipulate.ContextForIndex(contexts, 0)

	var query *mgo.Query
	if context.Filter != nil {
		query = collection.Find(context.Filter.Compile())
	} else {
		query = collection.Find(nil)
	}

	if err := query.All(dest); err != nil {
		return elemental.NewError("Error", err.Error(), "", 5000)
	}

	return nil
}

// Count is part of the implementation of the Manipulator interface.
func (s *MongoStore) Count(contexts manipulate.Contexts, identity elemental.Identity) (int, error) {

	collection := collectionFromIdentity(s.db, identity)
	context := manipulate.ContextForIndex(contexts, 0)

	var query *mgo.Query
	if context.Filter != nil {
		query = collection.Find(context.Filter.Compile())
	} else {
		query = collection.Find(nil)
	}

	c, err := query.Count()
	if err != nil {
		return 0, elemental.NewError("Error", err.Error(), "manipulate", 5000)
	}

	return c, nil
}

// Assign is part of the implementation of the Manipulator interface.
func (s *MongoStore) Assign(contexts manipulate.Contexts, parent manipulate.Manipulable, assignation *elemental.Assignation) error {

	panic("Not Implemented")
}

// Increment is part of the implementation of the Manipulator interface.
func (s *MongoStore) Increment(contexts manipulate.Contexts, name string, counter string, inc int, filterKeys []string, filterValues []interface{}) error {
	return fmt.Errorf("Increment is not implemented in mongo")
}

// Commit is part of the implementation of the TransactionalManipulator interface.
func (s *MongoStore) Commit(id manipulate.TransactionID) error {

	defer func() { s.unregisterBulk(id) }()

	if s.registeredBulkWithID(id) == nil {
		log.WithFields(log.Fields{
			"store":         s,
			"transactionID": id,
		}).Error("No batch found for the given transaction.")

		return elemental.NewError("Manipulation Error", "No batch found for the given transaction.", "manipulate", 5001)
	}

	if err := s.commitBulk(s.registeredBulkWithID(id)); err != nil {
		return elemental.NewError("Manipulation Error", "Unable to commit.", "manipulate", 5001)
	}

	return nil
}

// Abort is part of the implementation of the TransactionalManipulator interface.
func (s *MongoStore) Abort(id manipulate.TransactionID) bool {

	if s.registeredBulkWithID(id) == nil {
		return false
	}

	s.unregisterBulk(id)

	return true
}

func (s *MongoStore) bulkForID(id manipulate.TransactionID, collection *mgo.Collection) *mgo.Bulk {

	if id == "" {
		return collection.Bulk()
	}

	bulk := s.registeredBulkWithID(id)

	if bulk == nil {
		bulk = collection.Bulk()
		s.registerBulk(id, bulk)
	}

	return bulk
}

func (s *MongoStore) commitBulk(b *mgo.Bulk) error {

	log.WithFields(log.Fields{
		"bulk": b,
	}).Debug("Commiting bulk to mongo.")

	if _, err := b.Run(); err != nil {

		log.WithFields(log.Fields{
			"bulk":  b,
			"error": err,
		}).Debug("Unable to send bulk command.")

		return err
	}

	return nil
}

func (s *MongoStore) registerBulk(id manipulate.TransactionID, bulk *mgo.Bulk) {

	s.bulksRegistryLock.Lock()
	s.bulksRegistry[id] = bulk
	s.bulksRegistryLock.Unlock()
}

func (s *MongoStore) unregisterBulk(id manipulate.TransactionID) {

	s.bulksRegistryLock.Lock()
	delete(s.bulksRegistry, id)
	s.bulksRegistryLock.Unlock()
}

func (s *MongoStore) registeredBulkWithID(id manipulate.TransactionID) *mgo.Bulk {

	s.bulksRegistryLock.Lock()
	b := s.bulksRegistry[id]
	s.bulksRegistryLock.Unlock()

	return b
}
