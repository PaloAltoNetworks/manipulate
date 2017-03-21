package manipmongo

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/manipmongo/compiler"
	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

// Logger contains the main logger
var Logger = logrus.New()

var log = Logger.WithField("package", "github.com/aporeto-inc/manipulate/manipmongo")

// MongoStore represents a MongoDB session.
type mongoManipulator struct {
	rootSession  *mgo.Session
	dbName       string
	transactions *transactionsRegistry
}

// NewMongoManipulator returns a new TransactionalManipulator backed by MongoDB
func NewMongoManipulator(urls []string, dbName string, user string, password string, authsource string, poolLimit int, CAPool *x509.CertPool, clientCerts []tls.Certificate) manipulate.TransactionalManipulator {

	dialInfo, err := mgo.ParseURL(strings.Join(urls, ","))
	if err != nil {
		log.WithFields(logrus.Fields{
			"urls":     urls,
			"db":       dbName,
			"username": user,
			"error":    err.Error(),
		}).Fatal("Unable to create dial information")
	}

	dialInfo.PoolLimit = poolLimit
	dialInfo.Database = dbName
	dialInfo.Source = authsource
	dialInfo.Username = user
	dialInfo.Password = password
	dialInfo.Timeout = 3 * time.Second
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {

		conn, e := tls.Dial("tcp", addr.String(), &tls.Config{
			RootCAs:      CAPool,
			Certificates: clientCerts,
		})

		if e == nil {
			return conn, nil
		}

		log.WithError(e).Warn("Unable to dial to mongo using TLS. Trying with unencrypted dialing")
		return net.Dial("tcp", addr.String())
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.WithFields(logrus.Fields{
			"urls":     urls,
			"db":       dbName,
			"username": user,
			"error":    err.Error(),
		}).Fatal("Cannot connect to mongo.")
	}

	return &mongoManipulator{
		dbName:       dbName,
		rootSession:  session,
		transactions: newTransactionRegistry(),
	}
}

func (s *mongoManipulator) RetrieveMany(context *manipulate.Context, dest elemental.ContentIdentifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	session := s.rootSession.Copy()
	defer session.Close()

	db := session.DB(s.dbName)
	collection := collectionFromIdentity(db, dest.ContentIdentity())
	filter := bson.M{}

	if context.Filter != nil {
		filter = compiler.CompileFilter(context.Filter)
	}

	query := collection.Find(filter)

	// This makes squall returning a 500 error.
	// we should have an ErrBadRequest or something like this.
	// if context.Page > 0 && context.PageSize <= 0 {
	// 	return manipulate.NewErrCannotBuildQuery("Invalid pagination information")
	// }

	var err error
	if context.Page == 0 || context.PageSize == 0 {

		err = query.All(dest)

	} else if context.Page > 0 {

		skip := (context.Page - 1) * context.PageSize
		err = query.Skip(skip).Limit(context.PageSize).All(dest)

	} else {

		var n int
		n, err = s.Count(context, dest.ContentIdentity())
		if err != nil {
			return err
		}

		page := -context.Page
		skip := n - page*context.PageSize
		limit := context.PageSize

		if skip < 0 {

			maxPage := n / context.PageSize
			balance := n % context.PageSize
			if balance != 0 {
				maxPage++
			}

			// If the use asks or a page we know doesn't exist, we don't even talk to the dabatase.
			if page > maxPage {
				return nil
			}

			// otherwise, we have balance that we need to return.
			skip = 0
			limit = balance
		}

		err = query.Skip(skip).Limit(limit).All(dest)
	}

	if err != nil {
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	// backport all default values that are empty.
	for _, o := range dest.List() {
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}
	}

	return nil
}

func (s *mongoManipulator) Retrieve(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

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

		filter["_id"] = o.Identifier()

		if err := collection.Find(filter).One(o); err != nil {

			if err == mgo.ErrNotFound {
				return manipulate.NewErrObjectNotFound("cannot find the object for the given ID")
			}

			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}

		// backport all default values that are empty.
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}
	}

	return nil
}

func (s *mongoManipulator) Create(context *manipulate.Context, children ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	transaction, commit := s.retrieveTransaction(context)
	bulk := transaction.bulkForIdentity(children[0].Identity())

	for _, child := range children {
		child.SetIdentifier(bson.NewObjectId().Hex())

		if context.CreateFinalizer != nil {
			if err := context.CreateFinalizer(child); err != nil {
				return err
			}
		}

		bulk.Insert(child)
	}

	if commit {
		return s.Commit(transaction.id)
	}

	return nil
}

func (s *mongoManipulator) Update(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

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

func (s *mongoManipulator) Delete(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if context == nil {
		context = manipulate.NewContext()
	}

	transaction, commit := s.retrieveTransaction(context)
	bulk := transaction.bulkForIdentity(objects[0].Identity())

	for _, o := range objects {
		bulk.Remove(bson.M{"_id": o.Identifier()})

		// backport all default values that are empty.
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}
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

	bulk.RemoveAll(compiler.CompileFilter(context.Filter))

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

		log.WithField("bulk", bulk).Debug("Committing bulks to mongo.")

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
