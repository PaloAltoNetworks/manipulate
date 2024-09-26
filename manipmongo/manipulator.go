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
	neturl "net/url"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go/log"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/objectid"
	"go.aporeto.io/manipulate/internal/tracing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Use a default 2 min timeout for mongo queries when timeout is not specified in context
const defaultGlobalContextTimeout = 120 * time.Second

// MongoStore represents a MongoDB session.
type mongoManipulator struct {
	dbName              string
	client              *mongo.Client
	sharder             Sharder
	defaultRetryFunc    manipulate.RetryFunc
	forcedReadFilter    bson.D
	attributeEncrypter  elemental.AttributeEncrypter
	explain             map[elemental.Identity]map[elemental.Operation]struct{}
	attributeSpecifiers map[elemental.Identity]elemental.AttributeSpecifiable
}

// New returns a new manipulator backed by MongoDB.
func New(url string, db string, opts ...Option) (manipulate.TransactionalManipulator, error) {

	cfg := newConfig()
	for _, o := range opts {
		o(cfg)
	}

	// Parse the URL to check for authMechanism
	parsedURL, err := neturl.Parse(url)
	if err != nil {
		return nil, fmt.Errorf("invalid mongo URL: %w", err)
	}

	// Ensure there's a '/' before the query parameters if missing and not already present
	if parsedURL.Path == "" && parsedURL.RawQuery != "" {
		url = strings.Replace(url, "?"+parsedURL.RawQuery, "/?"+parsedURL.RawQuery, 1)

		// Re-parse the URL after potential modification
		parsedURL, err = neturl.Parse(url)
		if err != nil {
			return nil, fmt.Errorf("invalid mongo URL after modification: %w", err)
		}
	}

	// Check if authMechanism is present in the query
	queryParams := parsedURL.Query()
	authMechanism := queryParams.Get("authMechanism")

	clientOptions := options.Client().ApplyURI(url).
		SetMaxPoolSize(uint64(cfg.poolLimit)).
		SetConnectTimeout(cfg.connectTimeout)

	// If authMechanism is MONGODB-X509, then we don't specify username etc as those are derived from
	// the certificate. Specifying them regardless would cause mongo-go-driver to pick default authMechanism.
	if authMechanism != "MONGODB-X509" {
		clientOptions.SetAuth(options.Credential{
			Username:   cfg.username,
			Password:   cfg.password,
			AuthSource: cfg.authsource,
		})
	}

	if cfg.tlsConfig != nil {
		clientOptions.SetTLSConfig(cfg.tlsConfig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.connectTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to mongo url using mongo-go-driver'%s': %s", url, err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot ping mongo: %s", err)
	}

	return &mongoManipulator{
		dbName:              db,
		client:              client,
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

	var err error

	c, session, err := m.makeSession(dest.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	if err != nil {
		return err
	}
	defer session.EndSession(mctx.Context())

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
		filter = bson.D{{Key: "$and", Value: append(ands, filter)}}
	}

	// Query building
	findOptions := options.Find()

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
			sortFields = append(sortFields, bson.E{Key: field, Value: 1})
		}
		findOptions.SetSort(sortFields)
	}

	// Fields selection
	if sels := makeFieldsSelector(mctx.Fields(), attrSpec); sels != nil {
		findOptions.SetProjection(sels)
	}

	// Query timing limiting
	updatedFindOptions, err := setMaxTime(mctx.Context(), findOptions)
	if err != nil {
		return err
	}
	findOptions = updatedFindOptions.(*options.FindOptions)

	err = mongo.WithSession(context.Background(), session, func(sessCtx mongo.SessionContext) error {

		cursor, err := c.Find(sessCtx, filter, findOptions)
		if err != nil {
			return err
		}
		defer func() {
			if err := cursor.Close(sessCtx); err != nil {
				log.Error(err)
			}
		}()

		if _, err := RunQuery(
			mctx,
			func() (any, error) {
				if exp := explainIfNeeded(c, filter, dest.Identity(), elemental.OperationRetrieveMany, m.explain); exp != nil {
					if err := exp(); err != nil {
						return nil, manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("retrievemany: unable to explain: %w", err)}
					}
				}
				return nil, cursor.All(sessCtx, dest)
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
	})

	return err
}

func (m *mongoManipulator) Retrieve(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	var err error

	c, session, err := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	if err != nil {
		return err
	}
	defer session.EndSession(mctx.Context())

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
		filter = append(filter, bson.E{Key: "_id", Value: oid})
	} else {
		filter = append(filter, bson.E{Key: "_id", Value: object.Identifier()})
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
		filter = bson.D{{Key: "$and", Value: append(ands, filter)}}
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.retrieve.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()), log.Object("filter", filter))
	defer sp.Finish()

	findOptions := options.FindOne()

	if sels := makeFieldsSelector(mctx.Fields(), attrSpec); sels != nil {
		findOptions.SetProjection(sels)
	}

	updatedFindOptions, err := setMaxTime(mctx.Context(), findOptions)
	if err != nil {
		return err
	}
	findOptions = updatedFindOptions.(*options.FindOneOptions)

	err = mongo.WithSession(context.Background(), session, func(sessCtx mongo.SessionContext) error {

		result := c.FindOne(sessCtx, filter, findOptions)

		if _, err := RunQuery(
			mctx,
			func() (any, error) {
				if exp := explainIfNeeded(c, filter, object.Identity(), elemental.OperationRetrieve, m.explain); exp != nil {
					if err := exp(); err != nil {
						return nil, manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("retrieve: unable to explain: %w", err)}
					}
				}
				return nil, result.Decode(object)
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
	})

	return err
}

func (m *mongoManipulator) Create(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	var err error

	c, session, err := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	if err != nil {
		return err
	}
	defer session.EndSession(mctx.Context())

	oid := primitive.NewObjectID()
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

	err = mongo.WithSession(context.Background(), session, func(sessCtx mongo.SessionContext) error {

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
				filter = bson.D{{Key: "$and", Value: append(ands, filter)}}
			}

			info, err := RunQuery(
				mctx,
				func() (any, error) {
					upsertOptions := options.Update().SetUpsert(true)
					// Perform the upsert operation
					return c.UpdateOne(sessCtx, filter, baseOps, upsertOptions)
				},
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
			case *mongo.UpdateResult:
				if noid, ok := chinfo.UpsertedID.(primitive.ObjectID); ok {
					object.SetIdentifier(noid.Hex())
				}
			}

		} else {
			info, err := RunQuery(
				mctx,
				func() (any, error) { return c.InsertOne(sessCtx, object) },
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
			case *mongo.InsertOneResult:
				if noid, ok := chinfo.InsertedID.(primitive.ObjectID); ok {
					object.SetIdentifier(noid.Hex())
				}
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
	})

	return err
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

	var err error

	c, session, err := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	if err != nil {
		return err
	}
	defer session.EndSession(mctx.Context())

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.update.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()))
	defer sp.Finish()

	var filter bson.D
	if oid, ok := objectid.Parse(object.Identifier()); ok {
		filter = append(filter, bson.E{Key: "_id", Value: oid})
	} else {
		filter = append(filter, bson.E{Key: "_id", Value: object.Identifier()})
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
		filter = bson.D{{Key: "$and", Value: append(ands, filter)}}
	}

	err = mongo.WithSession(context.Background(), session, func(sessCtx mongo.SessionContext) error {

		if _, err := RunQuery(
			mctx,
			func() (any, error) { return c.UpdateOne(sessCtx, filter, bson.M{"$set": object}) },
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
	})

	return err
}

func (m *mongoManipulator) Delete(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	var err error

	c, session, err := m.makeSession(object.Identity(), mctx.ReadConsistency(), mctx.WriteConsistency())
	if err != nil {
		return err
	}
	defer session.EndSession(mctx.Context())

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongobject.delete.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()))
	defer sp.Finish()

	var filter bson.D
	if oid, ok := objectid.Parse(object.Identifier()); ok {
		filter = append(filter, bson.E{Key: "_id", Value: oid})
	} else {
		filter = append(filter, bson.E{Key: "_id", Value: object.Identifier()})
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
		filter = bson.D{{Key: "$and", Value: append(ands, filter)}}
	}

	err = mongo.WithSession(context.Background(), session, func(sessCtx mongo.SessionContext) error {

		if _, err := RunQuery(
			mctx,
			func() (any, error) { return c.DeleteOne(sessCtx, filter) },
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
	})

	return err
}

func (m *mongoManipulator) DeleteMany(mctx manipulate.Context, identity elemental.Identity) error {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.delete_many.%s", identity.Name))
	defer sp.Finish()

	var err error

	c, session, err := m.makeSession(identity, mctx.ReadConsistency(), mctx.WriteConsistency())
	if err != nil {
		return err
	}
	defer session.EndSession(mctx.Context())

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
		filter = bson.D{{Key: "$and", Value: append(ands, filter)}}
	}

	err = mongo.WithSession(context.Background(), session, func(sessCtx mongo.SessionContext) error {

		if _, err := RunQuery(
			mctx,
			func() (any, error) { return c.DeleteMany(sessCtx, filter) },
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
	})

	return err
}

func (m *mongoManipulator) Count(mctx manipulate.Context, identity elemental.Identity) (int, error) {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	var err error

	c, session, err := m.makeSession(identity, mctx.ReadConsistency(), mctx.WriteConsistency())
	if err != nil {
		return 0, err
	}
	defer session.EndSession(mctx.Context())

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
		filter = bson.D{{Key: "$and", Value: append(ands, filter)}}
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("manipmongo.count.%s", identity.Category))
	defer sp.Finish()

	countOptions := options.Count()
	updatedCountOptions, err := setMaxTime(mctx.Context(), countOptions)
	if err != nil {
		return 0, err
	}
	countOptions = updatedCountOptions.(*options.CountOptions)

	var count int64
	err = mongo.WithSession(context.Background(), session, func(sessCtx mongo.SessionContext) error {

		out, err := RunQuery(
			mctx,
			func() (any, error) {
				if exp := explainIfNeeded(c, filter, identity, elemental.OperationInfo, m.explain); exp != nil {
					if err := exp(); err != nil {
						return nil, manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("count: unable to explain: %w", err)}
					}
				}
				return c.CountDocuments(sessCtx, filter, countOptions)
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
			return err
		}

		count = out.(int64)
		return nil
	})

	return int(count), err
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
) (*mongo.Collection, mongo.Session, error) {

	readConcern := convertReadConsistency(readConsistency)
	writeConcern := convertWriteConsistency(writeConsistency)

	sessionOptions := options.Session()
	sessionOptions.SetDefaultReadConcern(readConcern)
	sessionOptions.SetDefaultWriteConcern(writeConcern)
	session, err := m.client.StartSession(sessionOptions)
	if err != nil {
		return nil, nil, err
	}

	database := m.client.Database(m.dbName)
	collection := database.Collection(identity.Name)
	return collection, session, nil
}
