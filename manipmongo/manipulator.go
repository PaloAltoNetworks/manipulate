package manipmongo

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/opentracing/opentracing-go/log"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/tracing"
	"go.aporeto.io/manipulate/manipmongo/internal/compiler"
	"go.uber.org/zap"
)

// MongoStore represents a MongoDB session.
type mongoManipulator struct {
	rootSession *mgo.Session
	dbName      string
	sharder     Sharder
}

// New returns a new manipulator backed by MongoDB.
func New(url string, db string, options ...Option) (manipulate.TransactionalManipulator, error) {

	cfg := newConfig()
	for _, o := range options {
		o(cfg)
	}

	dialInfo, err := mgo.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("cannot parse mongo url '%s': %s", url, err)
	}

	dialInfo.Database = db
	dialInfo.PoolLimit = cfg.poolLimit
	dialInfo.Username = cfg.username
	dialInfo.Password = cfg.password
	dialInfo.Source = cfg.authsource
	dialInfo.Timeout = cfg.connectTimeout

	if cfg.tlsConfig != nil {
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), cfg.tlsConfig)
		}
	} else {
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			return net.Dial("tcp", addr.String())
		}
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to mongo url '%s': %s", url, err)
	}

	session.SetSocketTimeout(cfg.socketTimeout)
	session.SetMode(convertReadConsistency(cfg.readConsistency), true)
	session.SetSafe(convertWriteConsistency(cfg.writeConsistency))

	return &mongoManipulator{
		dbName:      db,
		rootSession: session,
		sharder:     cfg.sharder,
	}, nil
}

// NewMongoManipulator returns a new TransactionalManipulator backed by MongoDB
func NewMongoManipulator(connectionString string, dbName string, user string, password string, authsource string, poolLimit int, CAPool *x509.CertPool, clientCerts []tls.Certificate) manipulate.TransactionalManipulator {

	fmt.Println("DEPRECATED: manipmongo.NewMongoManipulator is deprecated in favor of manipmongo.New")

	m, err := New(
		connectionString,
		dbName,
		OptionCredentials(user, password, authsource),
		OptionConnectionPoolLimit(poolLimit),
		OptionTLS(&tls.Config{
			RootCAs:      CAPool,
			Certificates: clientCerts,
		}),
	)

	if err != nil {
		zap.L().Fatal("Unable to connect to mongo",
			zap.String("uri", connectionString),
			zap.String("db", dbName),
			zap.String("username", user),
			zap.Error(err),
		)
	}

	return m
}

func (s *mongoManipulator) RetrieveMany(mctx manipulate.Context, dest elemental.Identifiables) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.retrieve_many.%s", dest.Identity().Category))
	defer sp.Finish()

	c, close := s.makeSession(dest.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	filter := bson.M{}

	if f := mctx.Filter(); f != nil {
		filter = compiler.CompileFilter(f)
	}

	if s.sharder != nil {
		sq := s.sharder.FilterMany(dest.Identity())
		if sq != nil {
			filter = bson.M{"$and": []bson.M{sq, filter}}
		}
	}

	query := c.Find(filter)

	// This makes squall returning a 500 error.
	// we should have an ErrBadRequest or something like this.
	// if mctx.Page > 0 && mctx.PageSize <= 0 {
	// 	return manipulate.NewErrCannotBuildQuery("Invalid pagination information")
	// }

	var inverted bool

	p := mctx.Page()
	ps := mctx.PageSize()
	if p > 0 {
		query = query.Skip((p - 1) * ps).Limit(ps)
	} else if p < 0 {
		query = query.Skip((-p - 1) * ps).Limit(ps)
		inverted = true
	}

	if o := mctx.Order(); len(o) > 0 {
		query = query.Sort(applyOrdering(o, inverted)...)
	} else if orderer, ok := dest.(elemental.DefaultOrderer); ok {
		query = query.Sort(applyOrdering(orderer.DefaultOrder(), inverted)...)
	} else {
		query = query.Sort(invertSortKey("$natural", inverted))
	}

	if sels := makeFieldsSelector(mctx.Fields()); sels != nil {
		query = query.Select(sels)
	}

	if err := query.All(dest); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return handleQueryError(err)
	}

	// backport all default values that are empty.
	for _, o := range dest.List() {
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}
	}

	return nil
}

func (s *mongoManipulator) Retrieve(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	c, close := s.makeSession(objects[0].Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	filter := bson.M{}

	if f := mctx.Filter(); f != nil {
		filter = compiler.CompileFilter(f)
	}

	for _, o := range objects {

		filter["_id"] = o.Identifier()

		if s.sharder != nil {
			sq := s.sharder.FilterOne(o)
			if sq != nil {
				filter = bson.M{"$and": []bson.M{sq, filter}}
			}
		}

		sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.retrieve.object.%s", objects[0].Identity().Name))
		sp.LogFields(log.String("object_id", o.Identifier()), log.Object("filter", filter))
		defer sp.Finish()

		query := c.Find(filter)
		if sels := makeFieldsSelector(mctx.Fields()); sels != nil {
			query = query.Select(sels)
		}

		if err := query.One(o); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return handleQueryError(err)
		}

		// backport all default values that are empty.
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}
	}

	return nil
}

func (s *mongoManipulator) Create(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	c, close := s.makeSession(objects[0].Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	for _, obj := range objects {

		obj.SetIdentifier(bson.NewObjectId().Hex())

		sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.create.object.%s", obj.Identity().Name))
		sp.LogFields(log.String("object_id", obj.Identifier()))
		defer sp.Finish()

		if f := mctx.Finalizer(); f != nil {
			if err := f(obj); err != nil {
				sp.SetTag("error", true)
				sp.LogFields(log.Error(err))
				return err
			}
		}

		if s.sharder != nil {
			s.sharder.Shard(obj)
		}

		if err := c.Insert(obj); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return handleQueryError(err)
		}
	}

	return nil
}

func (s *mongoManipulator) Update(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	c, close := s.makeSession(objects[0].Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	var filter bson.M

	for _, o := range objects {

		sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.update.object.%s", o.Identity().Name))
		sp.LogFields(log.String("object_id", o.Identifier()))
		defer sp.Finish()

		filter = bson.M{"_id": o.Identifier()}
		if s.sharder != nil {
			sq := s.sharder.FilterOne(o)
			if sq != nil {
				filter = bson.M{"$and": []bson.M{sq, filter}}
			}
		}

		if err := c.Update(filter, o); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return handleQueryError(err)
		}
	}

	return nil
}

func (s *mongoManipulator) Delete(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	c, close := s.makeSession(objects[0].Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	var filter bson.M

	for _, o := range objects {

		sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.delete.object.%s", o.Identity().Name))
		sp.LogFields(log.String("object_id", o.Identifier()))
		defer sp.Finish()

		filter = bson.M{"_id": o.Identifier()}
		if s.sharder != nil {
			sq := s.sharder.FilterOne(o)
			if sq != nil {
				filter = bson.M{"$and": []bson.M{sq, filter}}
			}
		}

		if err := c.Remove(filter); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return handleQueryError(err)
		}

		// backport all default values that are empty.
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}
	}

	return nil
}

func (s *mongoManipulator) DeleteMany(mctx manipulate.Context, identity elemental.Identity) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.delete_many.%s", identity.Name))
	defer sp.Finish()

	c, close := s.makeSession(identity, mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	filter := compiler.CompileFilter(mctx.Filter())
	if s.sharder != nil {
		sq := s.sharder.FilterMany(identity)
		if sq != nil {
			filter = bson.M{"$and": []bson.M{sq, filter}}
		}
	}

	if _, err := c.RemoveAll(filter); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return handleQueryError(err)
	}

	return nil
}

func (s *mongoManipulator) Count(mctx manipulate.Context, identity elemental.Identity) (int, error) {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	c, close := s.makeSession(identity, mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	filter := bson.M{}

	if f := mctx.Filter(); f != nil {
		filter = compiler.CompileFilter(f)
	}

	if s.sharder != nil {
		sq := s.sharder.FilterMany(identity)
		if sq != nil {
			filter = bson.M{"$and": []bson.M{sq, filter}}
		}
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.count.%s", identity.Category))
	defer sp.Finish()

	n, err := c.Find(filter).Count()
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return 0, handleQueryError(err)
	}

	return n, nil
}

func (s *mongoManipulator) Commit(id manipulate.TransactionID) error { return nil }

func (s *mongoManipulator) Abort(id manipulate.TransactionID) bool { return true }

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

func (s *mongoManipulator) makeSession(
	identity elemental.Identity,
	readConsistency manipulate.ReadConsistency,
	writeConsistency manipulate.WriteConsistency,
) (*mgo.Collection, func()) {

	session := s.rootSession.Copy()

	if mrc := convertReadConsistency(readConsistency); mrc != -1 {
		session.SetMode(mrc, true)
	}

	session.SetSafe(convertWriteConsistency(writeConsistency))

	return session.DB(s.dbName).C(identity.Name), session.Close
}
