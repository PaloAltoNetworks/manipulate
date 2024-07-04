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
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/opentracing/opentracing-go/log"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/objectid"
	"go.aporeto.io/manipulate/internal/tracing"

	"go.mongodb.org/mongo-driver/mongo"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const defaultGlobalContextTimeout = 60 * time.Second

// MongoStore represents a MongoDB session.
type mongoManipulator struct {
	dbName              string
	client              *mongo.Client
	database            *mongo.Database
	sharder             Sharder
	defaultRetryFunc    manipulate.RetryFunc
	forcedReadFilter    bson.D
	attributeEncrypter  elemental.AttributeEncrypter
	explain             map[elemental.Identity]map[elemental.Operation]struct{}
	attributeSpecifiers map[elemental.Identity]elemental.AttributeSpecifiable
}

// New returns a new manipulator backed by MongoDB.
func New(url string, db string, options ...Option) (manipulate.TransactionalManipulator, error) {

	cfg := newConfig()
	for _, o := range options {
		o(cfg)
	}

	clientOptions := mongooptions.Client().ApplyURI(url).
		SetAuth(mongooptions.Credential{
			Username:   cfg.username,
			Password:   cfg.password,
			AuthSource: cfg.authsource,
		}).
		SetMaxPoolSize(uint64(cfg.poolLimit)).
		SetConnectTimeout(cfg.connectTimeout)

	if cfg.tlsConfig != nil {
		clientOptions.SetTLSConfig(cfg.tlsConfig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.connectTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to mongo url '%s': %s", url, err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot ping mongo: %s", err)
	}

	return &mongoManipulator{
		dbName:              db,
		client:              client,
		database:            client.Database(db),
		sharder:             cfg.sharder,
		defaultRetryFunc:    cfg.defaultRetryFunc,
		forcedReadFilter:    cfg.forcedReadFilter,
		attributeEncrypter:  cfg.attributeEncrypter,
		explain:             cfg.explain,
		attributeSpecifiers: cfg.attributeSpecifiers,
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

	c, closeFunc := m.makeSession(dest.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer closeFunc(mctx.Context())

	var attrSpec elemental.AttributeSpecifiable
	if m.attributeSpecifiers != nil {
		attrSpec = m.attributeSpecifiers[dest.Identity()]
	}

	var order []string
	if o := mctx.Order(); len(o) > 0 {
		order = applyOrdering(o, attrSpec)
	} else if orderer, ok := dest.(elemental.DefaultOrderer); ok {
		order = applyOrdering(orderer.DefaultOrder(), attrSpec)
	}

	// Filtering
	filter := bson.D{}
	if f := mctx.Filter(); f != nil {
		var opts []CompilerOption
		if attrSpec != nil {
			opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
		}
		filter = CompileFilter(f, opts...)
	}

	var ands []bson.D
	var err error

	if m.sharder != nil {
		sq, err := m.sharder.FilterMany(m, mctx, dest.Identity())
		if err != nil {
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("cannot compute sharding filter: %w", err)}
		}
		if sq != nil {
			ands = append(ands, sq)
		}
	}

	if mctx.Namespace() != "" {
		var opts []CompilerOption
		if attrSpec != nil {
			opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
		}
		f := manipulate.NewNamespaceFilter(mctx.Namespace(), mctx.Recursive())
		if mctx.Propagated() {
			if fp := manipulate.NewPropagationFilter(mctx.Namespace()); fp != nil {
				f = elemental.NewFilterComposer().Or(f, fp).Done()
			}
		}
		ands = append(ands, CompileFilter(f, opts...))
	}

	if m.forcedReadFilter != nil {
		ands = append(ands, m.forcedReadFilter)
	}

	if after := mctx.After(); after != "" {

		if len(order) > 1 {
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("cannot use multiple ordering fields when using 'after'")}
		}

		var o string
		if len(order) == 1 {
			o = order[0]
		}

		f, err := prepareNextFilter(c, o, after)
		if err != nil {
			return err
		}

		ands = append(ands, f)
	}

	if len(ands) > 0 {
		filter = bson.D{{Name: "$and", Value: append(ands, filter)}}
	}

	// Query building
	findOptions := mongooptions.Find()

	// Limiting
	if limit := mctx.Limit(); limit > 0 {
		findOptions.SetLimit(int64(limit))
	} else if pageSize := mctx.PageSize(); pageSize > 0 {
		findOptions.SetLimit(int64(pageSize))
	}

	// Old pagination
	if p := mctx.Page(); p > 0 {
		findOptions.SetSkip(int64((p - 1) * mctx.PageSize()))
	}

	// Ordering
	if len(order) > 0 {
		var sortFields bson.D
		for _, field := range order {
			sortFields = append(sortFields, bson.DocElem{Name: field, Value: 1})
		}
		findOptions.SetSort(sortFields)
	}

	// Fields selection
	if sels := makeFieldsSelector(mctx.Fields(), attrSpec); sels != nil {
		findOptions.SetProjection(sels)
	}

	// Query timing limiting
	if findOptions, err = setQueryMaxTime(mctx.Context(), findOptions); err != nil {
		return err
	}

	cursor, err := c.Find(mctx.Context(), filter, findOptions)
	if err != nil {
		return err
	}
	defer cursor.Close(mctx.Context())

	if _, err := RunQuery(
		mctx,
		func() (any, error) {
			if exp := explainIfNeeded(c, filter, dest.Identity(), elemental.OperationRetrieveMany, m.explain); exp != nil {
				if err := exp(); err != nil {
					return nil, manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("retrievemany: unable to explain: %w", err)}
				}
			}
			return nil, cursor.All(mctx.Context(), dest)
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

	var lastID string

	lst := dest.List()
	for _, o := range lst {

		// backport all default values that are empty.
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}

		// Decrypt attributes if needed.
		if m.attributeEncrypter != nil {
			if a, ok := o.(elemental.AttributeEncryptable); ok {
				if err := a.DecryptAttributes(m.attributeEncrypter); err != nil {
					return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("retrievemany: unable to decrypt attributes: %w", err)}
				}
			}
		}

		lastID = o.Identifier()
	}

	if lastID != "" && (mctx.After() != "" || mctx.Limit() > 0) && len(lst) == mctx.Limit() {
		if lastID != mctx.After() {
			mctx.SetNext(lastID)
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

	c, closeFunc := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer closeFunc(mctx.Context())

	var attrSpec elemental.AttributeSpecifiable
	if m.attributeSpecifiers != nil {
		attrSpec = m.attributeSpecifiers[object.Identity()]
	}

	filter := bson.D{}
	if f := mctx.Filter(); f != nil {
		var opts []CompilerOption
		if attrSpec != nil {
			opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
		}
		filter = CompileFilter(f, opts...)
	}

	if oid, ok := objectid.Parse(object.Identifier()); ok {
		filter = append(filter, bson.DocElem{Name: "_id", Value: oid})
	} else {
		filter = append(filter, bson.DocElem{Name: "_id", Value: object.Identifier()})
	}

	var ands []bson.D

	if m.sharder != nil {
		sq, err := m.sharder.FilterOne(m, mctx, object)
		if err != nil {
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("cannot compute sharding filter: %w", err)}
		}
		if sq != nil {
			ands = append(ands, sq)
		}
	}

	if mctx.Namespace() != "" {
		var opts []CompilerOption
		if attrSpec != nil {
			opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
		}
		f := manipulate.NewNamespaceFilter(mctx.Namespace(), mctx.Recursive())
		if mctx.Propagated() {
			if fp := manipulate.NewPropagationFilter(mctx.Namespace()); fp != nil {
				f = elemental.NewFilterComposer().Or(f, fp).Done()
			}
		}
		ands = append(ands, CompileFilter(f, opts...))
	}

	if m.forcedReadFilter != nil {
		ands = append(ands, m.forcedReadFilter)
	}

	if len(ands) > 0 {
		filter = bson.D{{Name: "$and", Value: append(ands, filter)}}
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.retrieve.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()), log.Object("filter", filter))
	defer sp.Finish()

	findOptions := mongooptions.Find()
	var err error

	if sels := makeFieldsSelector(mctx.Fields(), attrSpec); sels != nil {
		findOptions.SetProjection(sels)
	}

	if findOptions, err = setQueryMaxTime(mctx.Context(), findOptions); err != nil {
		return err
	}

	cursor, err := c.Find(mctx.Context(), filter, findOptions)
	if err != nil {
		return err
	}
	defer cursor.Close(mctx.Context())

	if _, err := RunQuery(
		mctx,
		func() (any, error) {
			if exp := explainIfNeeded(c, filter, object.Identity(), elemental.OperationRetrieve, m.explain); exp != nil {
				if err := exp(); err != nil {
					return nil, manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("retrieve: unable to explain: %w", err)}
				}
			}
			return nil, cursor.Decode(object)
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
				return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("retrieve: unable to decrypt attributes: %w", err)}
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

	c, closeFunc := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer closeFunc(mctx.Context())

	oid := bson.NewObjectId()
	object.SetIdentifier(oid.Hex())

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
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("unable to execute sharder.Shard: %w", err)}
		}
	}

	var encryptable elemental.AttributeEncryptable
	if m.attributeEncrypter != nil {
		if a, ok := object.(elemental.AttributeEncryptable); ok {
			encryptable = a
			if err := a.EncryptAttributes(m.attributeEncrypter); err != nil {
				return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("create: unable to encrypt attributes: %w", err)}
			}
		}
	}

	if operations, upsert := mctx.(opaquer).Opaque()[opaqueKeyUpsert]; upsert {

		object.SetIdentifier("")

		ops, ok := operations.(bson.M)
		if !ok {
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("upsert operations must be of type bson.M")}
		}

		baseOps := bson.M{
			"$set":         object,
			"$setOnInsert": bson.M{"_id": oid},
		}

		if len(ops) > 0 {

			if soi, ok := ops["$setOnInsert"]; ok {
				for k, v := range soi.(bson.M) {
					baseOps["$setOnInsert"].(bson.M)[k] = v
				}
			}

			for k, v := range ops {
				if k == "$setOnInsert" {
					continue
				}
				baseOps[k] = v
			}
		}

		var attrSpec elemental.AttributeSpecifiable
		if m.attributeSpecifiers != nil {
			attrSpec = m.attributeSpecifiers[object.Identity()]
		}

		// Filtering
		filter := bson.D{}
		if f := mctx.Filter(); f != nil {
			var opts []CompilerOption
			if attrSpec != nil {
				opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
			}
			filter = CompileFilter(f, opts...)
		}

		var ands []bson.D

		if m.sharder != nil {
			sq, err := m.sharder.FilterOne(m, mctx, object)
			if err != nil {
				return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("cannot compute sharding filter: %w", err)}
			}
			if sq != nil {
				ands = append(ands, sq)
			}
		}

		if mctx.Namespace() != "" {
			var opts []CompilerOption
			if attrSpec != nil {
				opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
			}
			f := manipulate.NewNamespaceFilter(mctx.Namespace(), mctx.Recursive())
			if mctx.Propagated() {
				if fp := manipulate.NewPropagationFilter(mctx.Namespace()); fp != nil {
					f = elemental.NewFilterComposer().Or(f, fp).Done()
				}
			}
			ands = append(ands, CompileFilter(f, opts...))
		}

		if m.forcedReadFilter != nil {
			ands = append(ands, m.forcedReadFilter)
		}

		if len(ands) > 0 {
			filter = bson.D{{Name: "$and", Value: append(ands, filter)}}
		}

		info, err := RunQuery(
			mctx,
			func() (any, error) { return c.Upsert(filter, baseOps) },
			RetryInfo{
				Operation:        elemental.OperationCreate,
				Identity:         object.Identity(),
				defaultRetryFunc: m.defaultRetryFunc,
			},
		)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}

		switch chinfo := info.(type) {
		case *mgo.ChangeInfo:
			if noid, ok := chinfo.UpsertedId.(bson.ObjectId); ok {
				object.SetIdentifier(noid.Hex())
			}
		}

	} else {
		_, err := RunQuery(
			mctx,
			func() (any, error) { return nil, c.Insert(object) },
			RetryInfo{
				Operation:        elemental.OperationCreate,
				Identity:         object.Identity(),
				defaultRetryFunc: m.defaultRetryFunc,
			},
		)

		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}
	}

	if encryptable != nil {
		if err := encryptable.DecryptAttributes(m.attributeEncrypter); err != nil {
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("create: unable to decrypt attributes: %w", err)}
		}
	}

	if m.sharder != nil {
		if err := m.sharder.OnShardedWrite(m, mctx, elemental.OperationCreate, object); err != nil {
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("unable to execute sharder.OnShardedWrite on create: %w", err)}
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
				return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("update: unable to encrypt attributes: %w", err)}
			}
		}
	}

	c, closeFunc := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer closeFunc(mctx.Context())

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.update.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()))
	defer sp.Finish()

	var filter bson.D
	if oid, ok := objectid.Parse(object.Identifier()); ok {
		filter = append(filter, bson.DocElem{Name: "_id", Value: oid})
	} else {
		filter = append(filter, bson.DocElem{Name: "_id", Value: object.Identifier()})
	}

	var ands []bson.D

	if m.sharder != nil {
		sq, err := m.sharder.FilterOne(m, mctx, object)
		if err != nil {
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("cannot compute sharding filter: %w", err)}
		}
		if sq != nil {
			ands = append(ands, sq)
		}
	}

	if mctx.Namespace() != "" {
		var attrSpec elemental.AttributeSpecifiable
		if m.attributeSpecifiers != nil {
			attrSpec = m.attributeSpecifiers[object.Identity()]
		}
		var opts []CompilerOption
		if attrSpec != nil {
			opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
		}
		f := manipulate.NewNamespaceFilter(mctx.Namespace(), mctx.Recursive())
		if mctx.Propagated() {
			if fp := manipulate.NewPropagationFilter(mctx.Namespace()); fp != nil {
				f = elemental.NewFilterComposer().Or(f, fp).Done()
			}
		}
		ands = append(ands, CompileFilter(f, opts...))
	}

	if m.forcedReadFilter != nil {
		ands = append(ands, m.forcedReadFilter)
	}

	if len(ands) > 0 {
		filter = bson.D{{Name: "$and", Value: append(ands, filter)}}
	}

	if _, err := RunQuery(
		mctx,
		func() (any, error) { return nil, c.Update(filter, bson.M{"$set": object}) },
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
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("update: unable to decrypt attributes: %w", err)}
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

	c, closeFunc := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	defer closeFunc(mctx.Context())

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongobject.delete.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()))
	defer sp.Finish()

	var filter bson.D
	if oid, ok := objectid.Parse(object.Identifier()); ok {
		filter = append(filter, bson.DocElem{Name: "_id", Value: oid})
	} else {
		filter = append(filter, bson.DocElem{Name: "_id", Value: object.Identifier()})
	}

	var ands []bson.D

	if m.sharder != nil {
		sq, err := m.sharder.FilterOne(m, mctx, object)
		if err != nil {
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("cannot compute sharding filter: %w", err)}
		}
		if sq != nil {
			ands = append(ands, sq)
		}
	}

	if mctx.Namespace() != "" {
		var attrSpec elemental.AttributeSpecifiable
		if m.attributeSpecifiers != nil {
			attrSpec = m.attributeSpecifiers[object.Identity()]
		}
		var opts []CompilerOption
		if attrSpec != nil {
			opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
		}
		f := manipulate.NewNamespaceFilter(mctx.Namespace(), mctx.Recursive())
		if mctx.Propagated() {
			if fp := manipulate.NewPropagationFilter(mctx.Namespace()); fp != nil {
				f = elemental.NewFilterComposer().Or(f, fp).Done()
			}
		}
		ands = append(ands, CompileFilter(f, opts...))
	}

	if m.forcedReadFilter != nil {
		ands = append(ands, m.forcedReadFilter)
	}

	if len(ands) > 0 {
		filter = bson.D{{Name: "$and", Value: append(ands, filter)}}
	}

	if _, err := RunQuery(
		mctx,
		func() (any, error) { return nil, c.Remove(filter) },
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
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("unable to execute sharder.OnShardedWrite for delete: %w", err)}
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

	c, closeFunc := m.makeSession(identity, mctx.ReadConsistency(), mctx.WriteConsistency())
	defer closeFunc(mctx.Context())

	var attrSpec elemental.AttributeSpecifiable
	if m.attributeSpecifiers != nil {
		attrSpec = m.attributeSpecifiers[identity]
	}

	// Filtering
	filter := bson.D{}
	if f := mctx.Filter(); f != nil {
		var opts []CompilerOption
		if attrSpec != nil {
			opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
		}
		filter = CompileFilter(f, opts...)
	}

	var ands []bson.D

	if m.sharder != nil {
		sq, err := m.sharder.FilterMany(m, mctx, identity)
		if err != nil {
			return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("cannot compute sharding filter: %w", err)}
		}
		if sq != nil {
			ands = append(ands, sq)
		}
	}

	if mctx.Namespace() != "" {
		var opts []CompilerOption
		if attrSpec != nil {
			opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
		}
		f := manipulate.NewNamespaceFilter(mctx.Namespace(), mctx.Recursive())
		if mctx.Propagated() {
			if fp := manipulate.NewPropagationFilter(mctx.Namespace()); fp != nil {
				f = elemental.NewFilterComposer().Or(f, fp).Done()
			}
		}
		ands = append(ands, CompileFilter(f, opts...))
	}

	if m.forcedReadFilter != nil {
		ands = append(ands, m.forcedReadFilter)
	}

	if len(ands) > 0 {
		filter = bson.D{{Name: "$and", Value: append(ands, filter)}}
	}

	if _, err := RunQuery(
		mctx,
		func() (any, error) { return c.RemoveAll(filter) },
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

	c, closeFunc := m.makeSession(identity, mctx.ReadConsistency(), mctx.WriteConsistency())
	defer closeFunc(mctx.Context())

	var attrSpec elemental.AttributeSpecifiable
	if m.attributeSpecifiers != nil {
		attrSpec = m.attributeSpecifiers[identity]
	}

	// Filtering
	filter := bson.D{}
	if f := mctx.Filter(); f != nil {
		var opts []CompilerOption
		if attrSpec != nil {
			opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
		}
		filter = CompileFilter(f, opts...)
	}

	var ands []bson.D

	if m.sharder != nil {
		sq, err := m.sharder.FilterMany(m, mctx, identity)
		if err != nil {
			return 0, manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("cannot compute sharding filter: %w", err)}
		}
		if sq != nil {
			ands = append(ands, sq)
		}
	}

	if mctx.Namespace() != "" {
		var opts []CompilerOption
		if attrSpec != nil {
			opts = append(opts, CompilerOptionTranslateKeysFromSpec(attrSpec))
		}
		f := manipulate.NewNamespaceFilter(mctx.Namespace(), mctx.Recursive())
		if mctx.Propagated() {
			if fp := manipulate.NewPropagationFilter(mctx.Namespace()); fp != nil {
				f = elemental.NewFilterComposer().Or(f, fp).Done()
			}
		}
		ands = append(ands, CompileFilter(f, opts...))
	}

	if m.forcedReadFilter != nil {
		ands = append(ands, m.forcedReadFilter)
	}

	if len(ands) > 0 {
		filter = bson.D{{Name: "$and", Value: append(ands, filter)}}
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.count.%s", identity.Category))
	defer sp.Finish()

	q := c.Find(filter)

	var err error
	if q, err = setMaxTime(mctx.Context(), q); err != nil {
		return 0, err
	}

	out, err := RunQuery(
		mctx,
		func() (any, error) {
			if exp := explainIfNeeded(q, filter, identity, elemental.OperationInfo, m.explain); exp != nil {
				if err := exp(); err != nil {
					return nil, manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("count: unable to explain: %w", err)}
				}
			}
			return q.Count()
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		errChannel <- m.client.Ping(ctx, nil)
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
) (*mongo.Collection, func(context.Context)) {

	// session := m.rootSession.Copy()

	// if mrc := convertReadConsistency(readConsistency); mrc != -1 {
	// 	session.SetMode(mrc, true)
	// }

	// session.SetSafe(convertWriteConsistency(writeConsistency))

	// return session.DB(m.dbName).C(identity.Name), session.Close

	readConcern := &readconcern.ReadConcern{}
	writeConcern := &writeconcern.WriteConcern{}
	opts := mongooptions.Collection().
		SetReadConcern(readConcern).
		SetWriteConcern(writeConcern)

	session, err := m.client.StartSession()
	if err != nil {
		// Handle error
		return nil, nil
	}

	collection := m.database.Collection(identity.Name, opts)
	return collection, session.EndSession
}
