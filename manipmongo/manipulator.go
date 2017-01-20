package manipmongo

import (
	"fmt"

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

// MongoStore represents a MongoDB session.
type mongoManipulator struct {
	rootSession  *mgo.Session
	dbName       string
	url          string
	transactions *transactionsRegistry
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
		url:          url,
		dbName:       dbName,
		rootSession:  session,
		transactions: newTransactionRegistry(),
	}
}

func (s *mongoManipulator) RetrieveMany(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	session := s.rootSession.Copy()
	defer session.Close()

	db := session.DB(s.dbName)
	collection := collectionFromIdentity(db, identity)
	filter := bson.M{}

	if context.Filter != nil {
		filter = compiler.CompileFilter(context.Filter)
	}

	if err := collection.Find(filter).All(dest); err != nil {
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	return nil
}

func (s *mongoManipulator) Retrieve(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	session := s.rootSession.Copy()
	defer session.Close()

	db := session.DB(s.dbName)
	collection := collectionFromIdentity(db, objects[0].Identity())
	filter := bson.M{}

	if context.Filter != nil {
		filter = compiler.CompileFilter(context.Filter)
	}

	for _, o := range objects {

		if o.Identifier() != "" {
			filter["_id"] = o.Identifier()
		}

		if err := collection.Find(filter).One(o); err != nil {

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

	transaction, commit := s.retrieveTransaction(context)
	bulk := transaction.bulkForIdentity(children[0].Identity())

	for _, child := range children {
		child.SetIdentifier(uuid.NewV4().String())
		bulk.Insert(child)
	}

	if commit {
		return s.Commit(transaction.id)
	}

	return nil
}

func (s *mongoManipulator) Update(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	transaction, commit := s.retrieveTransaction(context)
	bulk := transaction.bulkForIdentity(objects[0].Identity())

	for _, o := range objects {
		bulk.Update(bson.M{"_id": o.Identifier()}, o)
	}

	if commit {
		return s.Commit(transaction.id)
	}

	return nil
}

func (s *mongoManipulator) Delete(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	transaction, commit := s.retrieveTransaction(context)
	bulk := transaction.bulkForIdentity(objects[0].Identity())

	for _, o := range objects {
		bulk.Remove(bson.M{"_id": o.Identifier()})
	}

	if commit {
		return s.Commit(transaction.id)
	}

	return nil
}

func (s *mongoManipulator) DeleteMany(context *manipulate.Context, identity elemental.Identity) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	transaction, commit := s.retrieveTransaction(context)
	bulk := transaction.bulkForIdentity(identity)

	bulk.Remove(compiler.CompileFilter(context.Filter))

	if commit {
		return s.Commit(transaction.id)
	}

	return nil
}

func (s *mongoManipulator) Count(context *manipulate.Context, identity elemental.Identity) (int, error) {

	if context == nil {
		context = manipulate.NewContext()
	}

	session := s.rootSession.Copy()
	defer session.Close()

	db := session.DB(s.dbName)
	collection := collectionFromIdentity(db, identity)
	filter := bson.M{}

	if context.Filter != nil {
		filter = compiler.CompileFilter(context.Filter)
	}

	c, err := collection.Find(filter).Count()
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

	transaction := s.transactions.transactionWithID(id)
	if transaction == nil {
		log.WithFields(logrus.Fields{
			"store":         s,
			"transactionID": id,
		}).Error("No batch found for the given transaction.")

		return manipulate.NewErrTransactionNotFound("No batch found for the given transaction.")
	}

	defer func() {
		transaction.closeSession()
		s.transactions.unregisterTransaction(id)
	}()

	for _, bulk := range transaction.bulks {

		log.WithField("bulk", bulk).Debug("Commiting bulks to mongo.")

		if _, err := bulk.Run(); err != nil {

			if mgo.IsDup(err) {
				return manipulate.NewErrConstraintViolation("duplicate key.")
			}

			return manipulate.NewErrCannotCommit(err.Error())
		}
	}

	return nil
}

func (s *mongoManipulator) Abort(id manipulate.TransactionID) bool {

	transaction := s.transactions.transactionWithID(id)
	if transaction == nil {
		return false
	}

	transaction.closeSession()
	s.transactions.unregisterTransaction(id)

	return true
}

func (s *mongoManipulator) retrieveTransaction(context *manipulate.Context) (*transaction, bool) {

	var created bool

	tid := context.TransactionID
	if tid == "" {
		tid = manipulate.NewTransactionID()
		created = true
	}

	t := s.transactions.transactionWithID(tid)
	if t != nil {
		return t, created
	}

	t = newTransaction(tid, s.rootSession, s.dbName)
	s.transactions.registerTransaction(tid, t)

	return t, created
}
