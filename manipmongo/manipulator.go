package manipmongo

import (
	"fmt"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/manipmongo/compiler"
	"gopkg.in/mgo.v2/bson"

	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
)

// Logger contains the main logger
var Logger = logrus.New()

var log = Logger.WithField("package", "github.com/aporeto-inc/manipulate/manipmongo")

type transactionsRegistry map[manipulate.TransactionID]map[*mgo.Collection]*mgo.Bulk

// MongoStore represents a MongoDB session.
type mongoManipulator struct {
	session *mgo.Session
	db      *mgo.Database
	dbName  string
	url     string

	transactionsRegistry     transactionsRegistry
	transactionsRegistryLock *sync.Mutex
}

// NewMongoManipulator returns a new TransactionalManipulator backed by MongoDB
func NewMongoManipulator(url string, dbName string) manipulate.TransactionalManipulator {

	session, err := mgo.Dial(url)
	if err != nil {
		log.WithFields(logrus.Fields{
			"url":   url,
			"db":    dbName,
			"error": err.Error(),
		}).Fatal("Cannot connect to mongo.")
	}

	return &mongoManipulator{
		url:                      url,
		dbName:                   dbName,
		transactionsRegistry:     transactionsRegistry{},
		transactionsRegistryLock: &sync.Mutex{},
		session:                  session,
		db:                       session.DB(dbName),
	}
}

func (s *mongoManipulator) RetrieveMany(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	collection := collectionFromIdentity(s.db, identity)

	var query *mgo.Query
	if context.Filter != nil {
		query = collection.Find(compiler.CompileFilter(context.Filter))
	} else {
		query = collection.Find(nil)
	}

	if err := query.All(dest); err != nil {
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	return nil
}

func (s *mongoManipulator) Retrieve(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	collection := collectionFromIdentity(s.db, objects[0].Identity())

	for i := 0; i < len(objects); i++ {

		var query *mgo.Query
		if context.Filter != nil {
			query = collection.Find(compiler.CompileFilter(context.Filter))
		} else {
			query = collection.Find(bson.M{"_id": objects[i].Identifier()})
		}

		if err := query.One(objects[i]); err != nil {
			if err == mgo.ErrNotFound {
				return manipulate.NewErrObjectNotFound("cannot find the object for the given ID")
			}
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
	}

	return nil
}

func (s *mongoManipulator) Create(context *manipulate.Context, children ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	collection := collectionFromIdentity(s.db, children[0].Identity())
	tid := context.TransactionID
	bulk := s.bulkForIDAndCollection(tid, collection)

	for _, child := range children {
		child.SetIdentifier(uuid.NewV4().String())
		bulk.Insert(child)
	}

	if tid == "" {
		if err := s.Commit(tid); err != nil {
			return manipulate.NewErrCannotCommit(err.Error())
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
	bulk := s.bulkForIDAndCollection(tid, collection)

	for i := 0; i < len(objects); i++ {

		filter := bson.M{}
		if context.Filter != nil {
			filter = compiler.CompileFilter(context.Filter)
		}

		if objects[i].Identifier() != "" {
			filter["_id"] = objects[i].Identifier()
		}

		bulk.Update(filter, objects[i])
	}

	if tid == "" {
		if err := s.Commit(tid); err != nil {
			return manipulate.NewErrCannotCommit(err.Error())
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
	bulk := s.bulkForIDAndCollection(tid, collection)

	for i := 0; i < len(objects); i++ {
		bulk.Remove(bson.M{"_id": objects[i].Identifier()})
	}

	if tid == "" {
		if err := s.Commit(tid); err != nil {
			return manipulate.NewErrCannotCommit(err.Error())
		}
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
		query = collection.Find(compiler.CompileFilter(context.Filter))
	} else {
		query = collection.Find(nil)
	}

	c, err := query.Count()
	if err != nil {
		return 0, manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	return c, nil
}

func (s *mongoManipulator) Assign(context *manipulate.Context, assignation *elemental.Assignation) error {
	return fmt.Errorf("Assign is not implemented in mongo")
}

func (s *mongoManipulator) Increment(context *manipulate.Context, identity elemental.Identity, counter string, inc int) error {
	return nil
}

func (s *mongoManipulator) Commit(id manipulate.TransactionID) error {

	defer func() { s.unregisterBulk(id) }()

	bulks := s.registeredTransactionWithID(id)
	if bulks == nil {
		log.WithFields(logrus.Fields{
			"store":         s,
			"transactionID": id,
		}).Error("No batch found for the given transaction.")

		return manipulate.NewErrTransactionNotFound("No batch found for the given transaction.")
	}

	for _, bulk := range bulks {

		log.WithField("bulk", bulk).Debug("Commiting bulks to mongo.")

		if _, err := bulk.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (s *mongoManipulator) Abort(id manipulate.TransactionID) bool {

	if s.registeredTransactionWithID(id) == nil {
		return false
	}

	s.unregisterBulk(id)

	return true
}

func (s *mongoManipulator) bulkForIDAndCollection(id manipulate.TransactionID, collection *mgo.Collection) *mgo.Bulk {

	if id == "" {
		return collection.Bulk()
	}

	bulks := s.registeredTransactionWithID(id)

	if bulks != nil && bulks[collection] != nil {
		return bulks[collection]
	}

	bulk := collection.Bulk()
	s.registerBulk(id, bulk, collection)

	return bulk
}

func (s *mongoManipulator) registerBulk(id manipulate.TransactionID, bulk *mgo.Bulk, collection *mgo.Collection) {

	s.transactionsRegistryLock.Lock()

	if s.transactionsRegistry[id] == nil {
		s.transactionsRegistry[id] = map[*mgo.Collection]*mgo.Bulk{}
	}
	s.transactionsRegistry[id][collection] = bulk

	s.transactionsRegistryLock.Unlock()
}

func (s *mongoManipulator) unregisterBulk(id manipulate.TransactionID) {

	s.transactionsRegistryLock.Lock()
	delete(s.transactionsRegistry, id)
	s.transactionsRegistryLock.Unlock()
}

func (s *mongoManipulator) registeredTransactionWithID(id manipulate.TransactionID) map[*mgo.Collection]*mgo.Bulk {

	s.transactionsRegistryLock.Lock()
	b := s.transactionsRegistry[id]
	s.transactionsRegistryLock.Unlock()

	return b
}
