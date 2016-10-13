package manipmongo

import (
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

	return &MongoStore{
		url:               url,
		dbName:            dbName,
		bulksRegistry:     bulksRegistry{},
		bulksRegistryLock: &sync.Mutex{},
	}
}

// Start starts the session.
func (s *MongoStore) Start() error {

	session, err := mgo.Dial(s.url)
	if err != nil {
		return err
	}

	s.session = session
	s.db = session.DB(s.dbName)

	return nil
}

// Stop stops the session.
func (s *MongoStore) Stop() {

	if s.session != nil {
		s.session.Close()
	}
}

// Create creates a new child Identifiable under the given parent Identifiable in the server.
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

// Retrieve fetchs the given Identifiable from the server.
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

// Update saves the given Identifiable into the server.
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

// Delete deletes the given Identifiable from the server.
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

// RetrieveChildren fetches the children with of given parent identified by the given Identity.
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

// Count count the number of element with the given Identity.
// Not implemented yet
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

// Assign assigns the list of given child Identifiables to the given Identifiable parent in the server.
func (s *MongoStore) Assign(contexts manipulate.Contexts, parent manipulate.Manipulable, assignation *elemental.Assignation) error {

	panic("Not Implemented")
}

// Commit will execute the batch of the given transaction
// The method will return an error if the batch does not succeed
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

// Abort aborts the given transaction ID.
func (s *MongoStore) Abort(id manipulate.TransactionID) bool {

	if s.registeredBulkWithID(id) == nil {
		return false
	}

	s.unregisterBulk(id)

	return true
}

// bulkForID return a mgo.Bulk.
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

// CommitBulk commit the given bulk
// The dev can then do whatever he wants with
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
