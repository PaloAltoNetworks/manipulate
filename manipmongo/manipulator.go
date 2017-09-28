package manipmongo

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"time"

	"go.uber.org/zap"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/tracing"
	"github.com/aporeto-inc/manipulate/manipmongo/compiler"
	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

// MongoStore represents a MongoDB session.
type mongoManipulator struct {
	rootSession  *mgo.Session
	dbName       string
	transactions *transactionsRegistry
}

// NewMongoManipulator returns a new TransactionalManipulator backed by MongoDB
func NewMongoManipulator(connectionString string, dbName string, user string, password string, authsource string, poolLimit int, CAPool *x509.CertPool, clientCerts []tls.Certificate) manipulate.TransactionalManipulator {

	dialInfo, err := mgo.ParseURL(connectionString)
	if err != nil {
		zap.L().Fatal("Unable to create dial information",
			zap.String("uri", connectionString),
			zap.String("db", dbName),
			zap.String("username", user),
			zap.Error(err),
		)
	}

	dialInfo.PoolLimit = poolLimit
	dialInfo.Database = dbName
	dialInfo.Source = authsource
	dialInfo.Username = user
	dialInfo.Password = password
	dialInfo.Timeout = 10 * time.Second
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {

		conn, e := tls.Dial("tcp", addr.String(), &tls.Config{
			RootCAs:      CAPool,
			Certificates: clientCerts,
		})

		if e == nil {
			return conn, nil
		}

		zap.L().Warn("Unable to dial to mongo using TLS. Trying with unencrypted dialing", zap.Error(e))
		return net.Dial("tcp", addr.String())
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		zap.L().Fatal("Cannot connect to mongo",
			zap.String("url", connectionString),
			zap.String("db", dbName),
			zap.String("username", user),
			zap.Error(err),
		)
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

	sp := tracing.StartTrace(context.TrackingSpan, fmt.Sprintf("manipmongo.retrieve_many.%s", dest.ContentIdentity().Category), context)

	session := s.copySession(context)
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

	var inverted bool

	if context.Page > 0 {
		query = query.Skip((context.Page - 1) * context.PageSize).Limit(context.PageSize)
	} else if context.Page < 0 {
		query = query.Skip((-context.Page - 1) * context.PageSize).Limit(context.PageSize)
		inverted = true
	}

	if len(context.Order) > 0 {
		query = query.Sort(applyOrdering(context.Order, inverted)...)
	} else if orderer, ok := dest.(elemental.DefaultOrderer); ok {
		query = query.Sort(applyOrdering(orderer.DefaultOrder(), inverted)...)
	} else {
		query = query.Sort(invertSortKey("$natural", inverted))
	}

	if err := query.All(dest); err != nil {
		tracing.FinishTraceWithError(sp, err)
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	// backport all default values that are empty.
	for _, o := range dest.List() {
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}
	}

	tracing.FinishTrace(sp)

	return nil
}

func (s *mongoManipulator) Retrieve(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, "manipmongo.retrieve", context)
	defer tracing.FinishTrace(sp)

	session := s.copySession(context)
	defer session.Close()

	db := session.DB(s.dbName)
	collection := collectionFromIdentity(db, objects[0].Identity())
	filter := bson.M{}

	if context.Filter != nil {
		filter = compiler.CompileFilter(context.Filter)
	}

	for _, o := range objects {

		subSp := tracing.StartTrace(sp, fmt.Sprintf("manipmongo.retrieve.object.%s", o.Identity().Name), context)
		tracing.SetTag(subSp, "manipmongo.retrieve.object.id", o.Identifier())

		filter["_id"] = o.Identifier()

		if err := collection.Find(filter).One(o); err != nil {

			if err == mgo.ErrNotFound {
				tracing.FinishTrace(subSp)
				return manipulate.NewErrObjectNotFound("cannot find the object for the given ID")
			}

			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}

		// backport all default values that are empty.
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}

		tracing.FinishTrace(subSp)
	}

	return nil
}

func (s *mongoManipulator) Create(context *manipulate.Context, children ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	transaction, commit := s.retrieveTransaction(context)
	bulk := transaction.bulkForIdentity(children[0].Identity())

	sp := tracing.StartTrace(context.TrackingSpan, "manipmongo.create", context)
	defer tracing.FinishTrace(sp)

	for _, child := range children {

		child.SetIdentifier(bson.NewObjectId().Hex())

		subSp := tracing.StartTrace(sp, fmt.Sprintf("manipmongo.create.object.%s", child.Identity().Name), context)
		tracing.SetTag(subSp, "manipmongo.create.object.id", child.Identifier())

		if context.CreateFinalizer != nil {
			if err := context.CreateFinalizer(child); err != nil {
				tracing.FinishTraceWithError(subSp, err)
				return err
			}
		}

		bulk.Insert(child)
		tracing.FinishTrace(subSp)
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

	sp := tracing.StartTrace(context.TrackingSpan, "manipmongo.update", context)
	defer tracing.FinishTrace(sp)

	transaction, commit := s.retrieveTransaction(context)
	bulk := transaction.bulkForIdentity(objects[0].Identity())

	for _, o := range objects {

		subSp := tracing.StartTrace(sp, fmt.Sprintf("manipmongo.update.object.%s", o.Identity().Name), context)
		tracing.SetTag(subSp, "manipmongo.update.object.id", o.Identifier())

		bulk.Update(bson.M{"_id": o.Identifier()}, o)
		tracing.FinishTrace(subSp)
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

	sp := tracing.StartTrace(context.TrackingSpan, "manipmongo.delete", context)
	defer tracing.FinishTrace(sp)

	transaction, commit := s.retrieveTransaction(context)
	bulk := transaction.bulkForIdentity(objects[0].Identity())

	for _, o := range objects {

		subSp := tracing.StartTrace(sp, fmt.Sprintf("manipmongo.delete.object.%s", o.Identity().Name), context)
		tracing.SetTag(subSp, "manipmongo.delete.object.id", o.Identifier())

		bulk.Remove(bson.M{"_id": o.Identifier()})

		// backport all default values that are empty.
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}

		tracing.FinishTrace(subSp)
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

	sp := tracing.StartTrace(context.TrackingSpan, "manipmongo.delete_many", context)
	defer tracing.FinishTrace(sp)

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

	sp := tracing.StartTrace(context.TrackingSpan, fmt.Sprintf("manipmongo.count.%s", identity.Category), context)

	session := s.copySession(context)
	defer session.Close()

	db := session.DB(s.dbName)
	collection := collectionFromIdentity(db, identity)
	filter := bson.M{}

	if context.Filter != nil {
		filter = compiler.CompileFilter(context.Filter)
	}

	c, err := collection.Find(filter).Count()
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return 0, manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	tracing.FinishTrace(sp)
	return c, nil
}

func (s *mongoManipulator) Increment(context *manipulate.Context, identity elemental.Identity, counter string, inc int) error {
	return nil
}

func (s *mongoManipulator) Commit(id manipulate.TransactionID) error {

	transaction := s.transactions.transactionWithID(id)
	if transaction == nil {
		return manipulate.NewErrTransactionNotFound("No batch found for the given transaction.")
	}

	sp := tracing.StartTrace(transaction.rootTracer, "manipmongo.commit", nil)

	defer func() {
		transaction.closeSession()
		s.transactions.unregisterTransaction(id)
	}()

	for _, bulk := range transaction.bulks {

		if _, err := bulk.Run(); err != nil {

			if mgo.IsDup(err) {
				tracing.FinishTrace(sp)
				return manipulate.NewErrConstraintViolation("duplicate key.")
			}

			tracing.FinishTraceWithError(sp, err)
			return manipulate.NewErrCannotCommit(err.Error())
		}
	}

	tracing.FinishTrace(sp)

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

	t = newTransaction(tid, s.copySession(context), s.dbName, context.TrackingSpan)
	s.transactions.registerTransaction(tid, t)

	return t, created
}

func (s *mongoManipulator) copySession(context *manipulate.Context) *mgo.Session {

	session := s.rootSession.Copy()
	session.SetSocketTimeout(context.Timeout)

	return session
}

func (s *mongoManipulator) Ping(timeout time.Duration) error {
	session := s.rootSession.Copy()
	errChannel := make(chan error, 1)

	go func() {
		errChannel <- session.Ping()
	}()

	select {
	case <-time.After(timeout):
		return fmt.Errorf("Connection Timeout")
	case err := <-errChannel:
		return err
	}

}
