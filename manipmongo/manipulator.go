package manipmongo

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/tracing"
	"go.aporeto.io/manipulate/manipmongo/internal/compiler"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
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

	session.SetSocketTimeout(60 * time.Second)

	return &mongoManipulator{
		dbName:       dbName,
		rootSession:  session,
		transactions: newTransactionRegistry(),
	}
}

func (s *mongoManipulator) RetrieveMany(mctx *manipulate.Context, dest elemental.Identifiables) error {

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.retrieve_many.%s", dest.Identity().Category))
	defer sp.Finish()

	session := s.rootSession.Copy()
	defer session.Close()

	db := session.DB(s.dbName)
	collection := collectionFromIdentity(db, dest.Identity(), "")
	filter := bson.M{}

	if mctx.Filter != nil {
		filter = compiler.CompileFilter(mctx.Filter)
	}

	query := collection.Find(filter)

	// This makes squall returning a 500 error.
	// we should have an ErrBadRequest or something like this.
	// if mctx.Page > 0 && mctx.PageSize <= 0 {
	// 	return manipulate.NewErrCannotBuildQuery("Invalid pagination information")
	// }

	var inverted bool

	if mctx.Page > 0 {
		query = query.Skip((mctx.Page - 1) * mctx.PageSize).Limit(mctx.PageSize)
	} else if mctx.Page < 0 {
		query = query.Skip((-mctx.Page - 1) * mctx.PageSize).Limit(mctx.PageSize)
		inverted = true
	}

	if len(mctx.Order) > 0 {
		query = query.Sort(applyOrdering(mctx.Order, inverted)...)
	} else if orderer, ok := dest.(elemental.DefaultOrderer); ok {
		query = query.Sort(applyOrdering(orderer.DefaultOrder(), inverted)...)
	} else {
		query = query.Sort(invertSortKey("$natural", inverted))
	}

	if err := query.All(dest); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
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

func (s *mongoManipulator) Retrieve(mctx *manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	session := s.rootSession.Copy()
	defer session.Close()

	db := session.DB(s.dbName)
	collection := collectionFromIdentity(db, objects[0].Identity(), "")
	filter := bson.M{}

	if mctx.Filter != nil {
		filter = compiler.CompileFilter(mctx.Filter)
	}

	for _, o := range objects {

		filter["_id"] = o.Identifier()

		sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.retrieve.object.%s", objects[0].Identity().Name))
		sp.LogFields(log.String("object_id", o.Identifier()), log.Object("filter", filter))
		defer sp.Finish()

		if err := collection.Find(filter).One(o); err != nil {

			sp.SetTag("error", true)

			if err == mgo.ErrNotFound {
				err = manipulate.NewErrObjectNotFound("cannot find the object for the given ID")
				sp.LogFields(log.Error(err))
				return err
			}

			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}

		// backport all default values that are empty.
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}
	}

	return nil
}

func (s *mongoManipulator) Create(mctx *manipulate.Context, children ...elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	transaction, commit := s.retrieveTransaction(mctx)
	bulk := transaction.bulkForIdentity(children[0].Identity(), "")

	for _, child := range children {

		child.SetIdentifier(bson.NewObjectId().Hex())

		sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.create.object.%s", child.Identity().Name))
		sp.LogFields(log.String("object_id", child.Identifier()))
		defer sp.Finish()

		if mctx.CreateFinalizer != nil {
			if err := mctx.CreateFinalizer(child); err != nil {
				sp.SetTag("error", true)
				sp.LogFields(log.Error(err))
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

func (s *mongoManipulator) Update(mctx *manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	transaction, commit := s.retrieveTransaction(mctx)
	bulk := transaction.bulkForIdentity(objects[0].Identity(), "")

	for _, o := range objects {

		sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.update.object.%s", o.Identity().Name))
		sp.LogFields(log.String("object_id", o.Identifier()))
		defer sp.Finish()

		bulk.Update(bson.M{"_id": o.Identifier()}, o)
	}

	if commit {
		return s.Commit(transaction.id)
	}

	return nil
}

func (s *mongoManipulator) Delete(mctx *manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	transaction, commit := s.retrieveTransaction(mctx)
	bulk := transaction.bulkForIdentity(objects[0].Identity(), "")

	for _, o := range objects {

		sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.delete.object.%s", o.Identity().Name))
		sp.LogFields(log.String("object_id", o.Identifier()))
		defer sp.Finish()

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

func (s *mongoManipulator) DeleteMany(mctx *manipulate.Context, identity elemental.Identity) error {

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.delete_many.%s", identity.Name))
	defer sp.Finish()

	transaction, commit := s.retrieveTransaction(mctx)
	bulk := transaction.bulkForIdentity(identity, "")

	bulk.RemoveAll(compiler.CompileFilter(mctx.Filter))

	if commit {
		return s.Commit(transaction.id)
	}

	return nil
}

func (s *mongoManipulator) Count(mctx *manipulate.Context, identity elemental.Identity) (int, error) {

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	session := s.rootSession.Copy()
	defer session.Close()

	db := session.DB(s.dbName)
	collection := collectionFromIdentity(db, identity, "")
	filter := bson.M{}

	if mctx.Filter != nil {
		filter = compiler.CompileFilter(mctx.Filter)
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.count.%s", identity.Category))
	defer sp.Finish()

	c, err := collection.Find(filter).Count()
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return 0, manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	return c, nil
}

func (s *mongoManipulator) Commit(id manipulate.TransactionID) error {

	transaction := s.transactions.transactionWithID(id)
	if transaction == nil {
		return manipulate.NewErrTransactionNotFound("No batch found for the given transaction.")
	}

	sp, _ := opentracing.StartSpanFromContext(transaction.ctx, "manipmongo.commit")
	defer sp.Finish()

	defer func() {
		transaction.closeSession()
		s.transactions.unregisterTransaction(id)
	}()

	for _, bulk := range transaction.bulks {

		if _, err := bulk.Run(); err != nil {

			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))

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

func (s *mongoManipulator) retrieveTransaction(mctx *manipulate.Context) (*transaction, bool) {

	var created bool

	tid := mctx.TransactionID
	if tid == "" {
		tid = manipulate.NewTransactionID()
		created = true
	}

	t := s.transactions.transactionWithID(tid)
	if t != nil {
		return t, created
	}

	t = newTransaction(mctx.Context(), tid, s.rootSession.Copy(), s.dbName)
	s.transactions.registerTransaction(tid, t)

	return t, created
}

func (s *mongoManipulator) Ping(timeout time.Duration) error {

	errChannel := make(chan error, 1)

	go func() {
		errChannel <- s.rootSession.Ping()
	}()

	select {
	case <-time.After(timeout):
		return fmt.Errorf("timeout")
	case err := <-errChannel:
		return err
	}

}
