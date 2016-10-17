package manipmongo

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/manipmongo/compilers"
	"gopkg.in/mgo.v2/bson"

	log "github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
)

type bulksRegistry map[manipulate.TransactionID]*mgo.Bulk

// MongoStore represents a MongoDB session.
type mongoManipulator struct {
	session *mgo.Session
	db      *mgo.Database
	dbName  string
	url     string

	bulksRegistry     bulksRegistry
	bulksRegistryLock *sync.Mutex
}

// NewMongoManipulator returns a new TransactionalManipulator backed by MongoDB
func NewMongoManipulator(url string, dbName string) manipulate.TransactionalManipulator {

	session, err := mgo.Dial(url)
	if err != nil {
		log.WithFields(log.Fields{
			"package": "manipmongo",
			"url":     url,
			"db":      dbName,
			"error":   err.Error(),
		}).Fatal("Cannot connect to mongo.")
	}

	return &mongoManipulator{
		url:               url,
		dbName:            dbName,
		bulksRegistry:     bulksRegistry{},
		bulksRegistryLock: &sync.Mutex{},
		session:           session,
		db:                session.DB(dbName),
	}
}

func (s *mongoManipulator) Create(context *manipulate.Context, children ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	collection := collectionFromIdentity(s.db, children[0].Identity())
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

func (s *mongoManipulator) Retrieve(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	collection := collectionFromIdentity(s.db, objects[0].Identity())

	for i := 0; i < len(objects); i++ {
		query := collection.Find(bson.M{"_id": objects[i].Identifier()})
		if err := query.One(objects[i]); err != nil {
			return elemental.NewError("Error", err.Error(), "", 5000)
		}
	}

	return nil
}

func (s *mongoManipulator) Update(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	collection := collectionFromIdentity(s.db, objects[0].Identity())
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

func (s *mongoManipulator) Delete(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	collection := collectionFromIdentity(s.db, objects[0].Identity())
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

func (s *mongoManipulator) RetrieveMany(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	collection := collectionFromIdentity(s.db, identity)

	var query *mgo.Query
	if context.Filter != nil {
		query = collection.Find(compilers.CompileFilter(context.Filter))
	} else {
		query = collection.Find(nil)
	}

	if err := query.All(dest); err != nil {
		return elemental.NewError("Error", err.Error(), "", 5000)
	}

	return nil
}

func (s *mongoManipulator) Count(context *manipulate.Context, identity elemental.Identity) (int, error) {

	if context == nil {
		context = manipulate.NewContext()
	}

	collection := collectionFromIdentity(s.db, identity)

	var query *mgo.Query
	if context.Filter != nil {
		query = collection.Find(compilers.CompileFilter(context.Filter))
	} else {
		query = collection.Find(nil)
	}

	c, err := query.Count()
	if err != nil {
		return 0, elemental.NewError("Error", err.Error(), "manipulate", 5000)
	}

	return c, nil
}

func (s *mongoManipulator) Assign(context *manipulate.Context, assignation *elemental.Assignation) error {
	return fmt.Errorf("Assign is not implemented in mongo")
}

func (s *mongoManipulator) Increment(context *manipulate.Context, name string, counter string, inc int, filterKeys []string, filterValues []interface{}) error {
	return fmt.Errorf("Increment is not implemented in mongo")
}

func (s *mongoManipulator) Commit(id manipulate.TransactionID) error {

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

func (s *mongoManipulator) Abort(id manipulate.TransactionID) bool {

	if s.registeredBulkWithID(id) == nil {
		return false
	}

	s.unregisterBulk(id)

	return true
}

func (s *mongoManipulator) bulkForID(id manipulate.TransactionID, collection *mgo.Collection) *mgo.Bulk {

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

func (s *mongoManipulator) commitBulk(b *mgo.Bulk) error {

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

func (s *mongoManipulator) registerBulk(id manipulate.TransactionID, bulk *mgo.Bulk) {

	s.bulksRegistryLock.Lock()
	s.bulksRegistry[id] = bulk
	s.bulksRegistryLock.Unlock()
}

func (s *mongoManipulator) unregisterBulk(id manipulate.TransactionID) {

	s.bulksRegistryLock.Lock()
	delete(s.bulksRegistry, id)
	s.bulksRegistryLock.Unlock()
}

func (s *mongoManipulator) registeredBulkWithID(id manipulate.TransactionID) *mgo.Bulk {

	s.bulksRegistryLock.Lock()
	b := s.bulksRegistry[id]
	s.bulksRegistryLock.Unlock()

	return b
}
