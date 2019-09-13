// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package manipmongo

import (
	"context"
	"crypto/tls"
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
)

const defaultGlobalContextTimeout = 60 * time.Second

// MongoStore represents a MongoDB session.
type mongoManipulator struct {
	rootSession        *mgo.Session
	dbName             string
	sharder            Sharder
	defaultRetryFunc   manipulate.RetryFunc
	forcedReadFilter   bson.M
	attributeEncrypter elemental.AttributeEncrypter
	explain            map[elemental.Identity]map[elemental.Operation]struct{}
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
		dbName:             db,
		rootSession:        session,
		sharder:            cfg.sharder,
		defaultRetryFunc:   cfg.defaultRetryFunc,
		forcedReadFilter:   cfg.forcedReadFilter,
		attributeEncrypter: cfg.attributeEncrypter,
		explain:            cfg.explain,
	}, nil
}

func (m *mongoManipulator) RetrieveMany(mctx manipulate.Context, dest elemental.Identifiables) error {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.retrieve_many.%s", dest.Identity().Category))
	defer sp.Finish()

	c, close := m.makeSession(dest.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	filter := bson.M{}

	if f := mctx.Filter(); f != nil {
		filter = compiler.CompileFilter(f)
	}

	if m.sharder != nil {
		sq, err := m.sharder.FilterMany(m, mctx, dest.Identity())
		if err != nil {
			return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("cannot compute sharding filter: %s", err))
		}
		if sq != nil {
			filter = bson.M{"$and": []bson.M{sq, filter}}
		}
	}

	if m.forcedReadFilter != nil {
		filter = bson.M{"$and": []bson.M{m.forcedReadFilter, filter}}
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

	query = query.SetMaxTime(defaultGlobalContextTimeout)
	if d, ok := mctx.Context().Deadline(); ok {
		query = query.SetMaxTime(time.Until(d))
	}

	if _, err := RunQuery(
		mctx,
		func() (interface{}, error) {
			if exp := explainIfNeeded(query, filter, dest.Identity(), elemental.OperationRetrieveMany, m.explain); exp != nil {
				if err := exp(); err != nil {
					return nil, manipulate.NewErrCannotBuildQuery(fmt.Sprintf("retrievemany: unable to explain: %s", err))
				}
			}
			return nil, query.All(dest)
		},
		RetryInfo{
			Operation:        elemental.OperationRetrieveMany,
			Identity:         dest.Identity(),
			defaultRetryFunc: m.defaultRetryFunc,
		},
	); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	// backport all default values that are empty.
	for _, o := range dest.List() {
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}

		if m.attributeEncrypter != nil {
			if a, ok := o.(elemental.AttributeEncryptable); ok {
				if err := a.DecryptAttributes(m.attributeEncrypter); err != nil {
					return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("retrievemany: unable to decrypt attributes: %s", err))
				}
			}
		}
	}

	return nil
}

func (m *mongoManipulator) Retrieve(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	c, close := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	filter := bson.M{}

	if f := mctx.Filter(); f != nil {
		filter = compiler.CompileFilter(f)
	}

	filter["_id"] = object.Identifier()

	if m.sharder != nil {
		sq, err := m.sharder.FilterOne(m, mctx, object)
		if err != nil {
			return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("cannot compute sharding filter: %s", err))
		}
		if sq != nil {
			filter = bson.M{"$and": []bson.M{sq, filter}}
		}
	}

	if m.forcedReadFilter != nil {
		filter = bson.M{"$and": []bson.M{m.forcedReadFilter, filter}}
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.retrieve.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()), log.Object("filter", filter))
	defer sp.Finish()

	query := c.Find(filter)
	if sels := makeFieldsSelector(mctx.Fields()); sels != nil {
		query = query.Select(sels)
	}

	query = query.SetMaxTime(defaultGlobalContextTimeout)
	if d, ok := mctx.Context().Deadline(); ok {
		query = query.SetMaxTime(time.Until(d))
	}

	if _, err := RunQuery(
		mctx,
		func() (interface{}, error) {
			if exp := explainIfNeeded(query, filter, object.Identity(), elemental.OperationRetrieve, m.explain); exp != nil {
				if err := exp(); err != nil {
					return nil, manipulate.NewErrCannotBuildQuery(fmt.Sprintf("retrieve: unable to explain: %s", err))
				}
			}
			return nil, query.One(object)
		},
		RetryInfo{
			Operation:        elemental.OperationRetrieve,
			Identity:         object.Identity(),
			defaultRetryFunc: m.defaultRetryFunc,
		},
	); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	// backport all default values that are empty.
	if a, ok := object.(elemental.AttributeSpecifiable); ok {
		elemental.ResetDefaultForZeroValues(a)
	}

	if m.attributeEncrypter != nil {
		if a, ok := object.(elemental.AttributeEncryptable); ok {
			if err := a.DecryptAttributes(m.attributeEncrypter); err != nil {
				return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("retrieve: unable to decrypt attributes: %s", err))
			}
		}
	}

	return nil
}

func (m *mongoManipulator) Create(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	c, close := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	object.SetIdentifier(bson.NewObjectId().Hex())

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.create.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()))
	defer sp.Finish()

	if f := mctx.Finalizer(); f != nil {
		if err := f(object); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}
	}

	if m.sharder != nil {
		if err := m.sharder.Shard(m, mctx, object); err != nil {
			return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("unable to execute sharder.Shard: %s", err))
		}
	}

	var encryptable elemental.AttributeEncryptable
	if m.attributeEncrypter != nil {
		if a, ok := object.(elemental.AttributeEncryptable); ok {
			encryptable = a
			if err := a.EncryptAttributes(m.attributeEncrypter); err != nil {
				return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("create: unable to encrypt attributes: %s", err))
			}
		}
	}

	if _, err := RunQuery(
		mctx,
		func() (interface{}, error) { return nil, c.Insert(object) },
		RetryInfo{
			Operation:        elemental.OperationCreate,
			Identity:         object.Identity(),
			defaultRetryFunc: m.defaultRetryFunc,
		},
	); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if encryptable != nil {
		if err := encryptable.DecryptAttributes(m.attributeEncrypter); err != nil {
			return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("create: unable to decrypt attributes: %s", err))
		}
	}

	if m.sharder != nil {
		if err := m.sharder.OnShardedWrite(m, mctx, elemental.OperationCreate, object); err != nil {
			return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("unable to execute sharder.OnShardedWrite on create: %s", err))
		}
	}

	return nil
}

func (m *mongoManipulator) Update(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	var encryptable elemental.AttributeEncryptable
	if m.attributeEncrypter != nil {
		if a, ok := object.(elemental.AttributeEncryptable); ok {
			encryptable = a
			if err := a.EncryptAttributes(m.attributeEncrypter); err != nil {
				return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("update: unable to encrypt attributes: %s", err))
			}
		}
	}

	c, close := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	var filter bson.M

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.update.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()))
	defer sp.Finish()

	filter = bson.M{"_id": object.Identifier()}
	if m.sharder != nil {
		sq, err := m.sharder.FilterOne(m, mctx, object)
		if err != nil {
			return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("cannot compute sharding filter: %s", err))
		}
		if sq != nil {
			filter = bson.M{"$and": []bson.M{sq, filter}}
		}
	}

	if m.forcedReadFilter != nil {
		filter = bson.M{"$and": []bson.M{m.forcedReadFilter, filter}}
	}

	if _, err := RunQuery(
		mctx,
		func() (interface{}, error) { return nil, c.Update(filter, bson.M{"$set": object}) },
		RetryInfo{
			Operation:        elemental.OperationUpdate,
			Identity:         object.Identity(),
			defaultRetryFunc: m.defaultRetryFunc,
		},
	); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if encryptable != nil {
		if err := encryptable.DecryptAttributes(m.attributeEncrypter); err != nil {
			return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("update: unable to decrypt attributes: %s", err))
		}
	}

	return nil
}

func (m *mongoManipulator) Delete(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	c, close := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	var filter bson.M

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongobject.delete.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()))
	defer sp.Finish()

	filter = bson.M{"_id": object.Identifier()}
	if m.sharder != nil {
		sq, err := m.sharder.FilterOne(m, mctx, object)
		if err != nil {
			return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("cannot compute sharding filter: %s", err))
		}
		if sq != nil {
			filter = bson.M{"$and": []bson.M{sq, filter}}
		}
	}

	if m.forcedReadFilter != nil {
		filter = bson.M{"$and": []bson.M{m.forcedReadFilter, filter}}
	}

	if _, err := RunQuery(
		mctx,
		func() (interface{}, error) { return nil, c.Remove(filter) },
		RetryInfo{
			Operation:        elemental.OperationDelete,
			Identity:         object.Identity(),
			defaultRetryFunc: m.defaultRetryFunc,
		},
	); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if m.sharder != nil {
		if err := m.sharder.OnShardedWrite(m, mctx, elemental.OperationDelete, object); err != nil {
			return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("unable to execute sharder.OnShardedWrite for delete: %s", err))
		}
	}

	// backport all default values that are empty.
	if a, ok := object.(elemental.AttributeSpecifiable); ok {
		elemental.ResetDefaultForZeroValues(a)
	}

	return nil
}

func (m *mongoManipulator) DeleteMany(mctx manipulate.Context, identity elemental.Identity) error {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.delete_many.%s", identity.Name))
	defer sp.Finish()

	c, close := m.makeSession(identity, mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	filter := compiler.CompileFilter(mctx.Filter())
	if m.sharder != nil {
		sq, err := m.sharder.FilterMany(m, mctx, identity)
		if err != nil {
			return manipulate.NewErrCannotBuildQuery(fmt.Sprintf("cannot compute sharding filter: %s", err))
		}
		if sq != nil {
			filter = bson.M{"$and": []bson.M{sq, filter}}
		}
	}

	if m.forcedReadFilter != nil {
		filter = bson.M{"$and": []bson.M{m.forcedReadFilter, filter}}
	}

	if _, err := RunQuery(
		mctx,
		func() (interface{}, error) { return c.RemoveAll(filter) },
		RetryInfo{
			Operation:        elemental.OperationDelete, // we miss DeleteMany
			Identity:         identity,
			defaultRetryFunc: m.defaultRetryFunc,
		},
	); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	return nil
}

func (m *mongoManipulator) Count(mctx manipulate.Context, identity elemental.Identity) (int, error) {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	c, close := m.makeSession(identity, mctx.ReadConsistency(), mctx.WriteConsistency())
	defer close()

	filter := bson.M{}

	if f := mctx.Filter(); f != nil {
		filter = compiler.CompileFilter(f)
	}

	if m.sharder != nil {
		sq, err := m.sharder.FilterMany(m, mctx, identity)
		if err != nil {
			return 0, manipulate.NewErrCannotBuildQuery(fmt.Sprintf("cannot compute sharding filter: %s", err))
		}
		if sq != nil {
			filter = bson.M{"$and": []bson.M{sq, filter}}
		}
	}

	if m.forcedReadFilter != nil {
		filter = bson.M{"$and": []bson.M{m.forcedReadFilter, filter}}
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.count.%s", identity.Category))
	defer sp.Finish()

	query := c.Find(filter)

	query = query.SetMaxTime(defaultGlobalContextTimeout)
	if d, ok := mctx.Context().Deadline(); ok {
		query = query.SetMaxTime(time.Until(d))
	}

	out, err := RunQuery(
		mctx,
		func() (interface{}, error) {
			if exp := explainIfNeeded(query, filter, identity, elemental.OperationInfo, m.explain); exp != nil {
				if err := exp(); err != nil {
					return nil, manipulate.NewErrCannotBuildQuery(fmt.Sprintf("count: unable to explain: %s", err))
				}
			}
			return query.Count()
		},
		RetryInfo{
			Operation:        elemental.OperationInfo,
			Identity:         identity,
			defaultRetryFunc: m.defaultRetryFunc,
		},
	)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return 0, err
	}

	return out.(int), nil
}

func (m *mongoManipulator) Commit(id manipulate.TransactionID) error { return nil }

func (m *mongoManipulator) Abort(id manipulate.TransactionID) bool { return true }

func (m *mongoManipulator) Ping(timeout time.Duration) error {

	errChannel := make(chan error, 1)

	go func() {
		errChannel <- m.rootSession.Ping()
	}()

	select {
	case <-time.After(timeout):
		return fmt.Errorf("timeout")
	case err := <-errChannel:
		return err
	}
}

func (m *mongoManipulator) makeSession(
	identity elemental.Identity,
	readConsistency manipulate.ReadConsistency,
	writeConsistency manipulate.WriteConsistency,
) (*mgo.Collection, func()) {

	session := m.rootSession.Copy()

	if mrc := convertReadConsistency(readConsistency); mrc != -1 {
		session.SetMode(mrc, true)
	}

	session.SetSafe(convertWriteConsistency(writeConsistency))

	return session.DB(m.dbName).C(identity.Name), session.Close
}
