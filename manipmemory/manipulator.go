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

package manipmemory

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/globalsign/mgo/bson"
	memdb "github.com/hashicorp/go-memdb"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

type txnRegistry map[manipulate.TransactionID]*memdb.Txn

// A memoryManipulator is an empty manipulator that can be used with ApoMock.
type memdbManipulator struct {
	db              *memdb.MemDB
	schema          *memdb.DBSchema
	txnRegistry     txnRegistry
	txnRegistryLock sync.RWMutex
	dbLock          sync.RWMutex
	noCopy          bool
}

// New creates a new datastore backed by a memdb.
func New(c map[string]*IdentitySchema, options ...Option) (manipulate.TransactionalManipulator, error) {

	cfg := newConfig()
	for _, opt := range options {
		opt(cfg)
	}

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{},
	}

	for table, cfg := range c {
		index, err := createSchema(cfg)
		if err != nil {
			return nil, err
		}
		schema.Tables[table] = index
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	return &memdbManipulator{
		schema:      schema,
		db:          db,
		noCopy:      cfg.noCopy,
		txnRegistry: txnRegistry{},
	}, nil
}

// Flush will flush the datastore essentially creating a new one.
func (m *memdbManipulator) Flush(ctx context.Context) error {

	db, err := memdb.NewMemDB(m.schema)
	if err != nil {
		return err
	}

	m.setDB(db)

	return nil
}

// RetrieveMany is part of the implementation of the Manipulator interface.
func (m *memdbManipulator) RetrieveMany(mctx manipulate.Context, dest elemental.Identifiables) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	items := map[string]elemental.Identifiable{}

	if err := m.retrieveFromFilter(dest.Identity().Category, mctx.Filter(), &items, true); err != nil {
		return err
	}

	out := reflect.ValueOf(dest).Elem()

	for _, obj := range items {
		out.Set(reflect.Append(out, reflect.ValueOf(obj)))
	}

	return nil
}

// Retrieve is part of the implementation of the Manipulator interface.
func (m *memdbManipulator) Retrieve(mctx manipulate.Context, object elemental.Identifiable) error {

	txn := m.getDB().Txn(false)

	raw, err := txn.First(object.Identity().Category, "id", object.Identifier())
	if err != nil {
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	if raw == nil {
		return manipulate.NewErrObjectNotFound("cannot find the object for the given ID")
	}

	var cp interface{}
	if m.noCopy {
		cp = raw
	} else {
		cp, err = copystructure.Copy(raw)
		if err != nil {
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
	}

	reflect.ValueOf(object).Elem().Set(reflect.ValueOf(cp).Elem())

	return nil
}

// Create is part of the implementation of the Manipulator interface.
func (m *memdbManipulator) Create(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	tid := mctx.TransactionID()
	txn := m.txnForID(tid)
	defer txn.Abort()

	// In caching scenarios the identifier is already set. Do not insert
	// here. We will get it pre-populated from the master DB.
	if object.Identifier() == "" {
		object.SetIdentifier(bson.NewObjectId().Hex())
	}

	var cp interface{}
	if m.noCopy {
		cp = object
	} else {
		var err error
		cp, err = copystructure.Copy(object)
		if err != nil {
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
	}

	if err := txn.Insert(object.Identity().Category, cp); err != nil {
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	if tid == "" {
		txn.Commit()
	}

	return nil
}

// Update is part of the implementation of the Manipulator interface.
func (m *memdbManipulator) Update(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	tid := mctx.TransactionID()
	txn := m.txnForID(tid)
	defer txn.Abort()

	o, err := txn.Get(object.Identity().Category, "id", object.Identifier())
	if err != nil || o.Next() == nil {
		return manipulate.NewErrObjectNotFound("Cannot find object with given ID")
	}

	var cp interface{}
	if m.noCopy {
		cp = object
	} else {
		cp, err = copystructure.Copy(object)
		if err != nil {
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
	}

	// Delete prior to insert to avoid tangling indices
	if err := txn.Delete(object.Identity().Category, object); err != nil {
		if err == memdb.ErrNotFound {
			return manipulate.NewErrObjectNotFound(err.Error())
		}
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	if err := txn.Insert(object.Identity().Category, cp); err != nil {
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	if tid == "" {
		txn.Commit()
	}

	return nil
}

// Delete is part of the implementation of the Manipulator interface.
func (m *memdbManipulator) Delete(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	tid := mctx.TransactionID()
	txn := m.txnForID(tid)
	defer txn.Abort()

	if err := txn.Delete(object.Identity().Category, object); err != nil {
		if err == memdb.ErrNotFound {
			return manipulate.NewErrObjectNotFound(err.Error())
		}
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	if tid == "" {
		txn.Commit()
	}

	return nil
}

// DeleteMany is part of the implementation of the Manipulator interface.
func (m *memdbManipulator) DeleteMany(mctx manipulate.Context, identity elemental.Identity) error {
	return manipulate.NewErrNotImplemented("DeleteMany not implemented in manipmemory")
}

// Count is part of the implementation of the Manipulator interface. Count is very expensive.
func (m *memdbManipulator) Count(mctx manipulate.Context, identity elemental.Identity) (int, error) {

	items := map[string]elemental.Identifiable{}

	if err := m.retrieveFromFilter(identity.Category, mctx.Filter(), &items, true); err != nil {
		return 0, err
	}

	return len(items), nil
}

// Commit is part of the implementation of the TransactionalManipulator interface.
func (m *memdbManipulator) Commit(id manipulate.TransactionID) error {

	txn := m.registeredTxnWithID(id)

	if txn == nil {
		return manipulate.NewErrCannotCommit("Cannot find transaction " + string(id))
	}

	txn.Commit()
	m.unregisterTxn(id)

	return nil
}

// Abort is part of the implementation of the TransactionalManipulator interface.
func (m *memdbManipulator) Abort(id manipulate.TransactionID) bool {

	txn := m.registeredTxnWithID(id)
	if txn == nil {
		return false
	}

	txn.Abort()
	m.unregisterTxn(id)

	return true
}

func (m *memdbManipulator) txnForID(id manipulate.TransactionID) *memdb.Txn {

	if id == "" {
		return m.getDB().Txn(true)
	}

	txn := m.registeredTxnWithID(id)

	if txn == nil {
		txn = m.getDB().Txn(true)
		m.registerTxn(id, txn)
	}

	return txn
}

func (m *memdbManipulator) registerTxn(id manipulate.TransactionID, txn *memdb.Txn) {

	m.txnRegistryLock.Lock()
	defer m.txnRegistryLock.Unlock()
	m.txnRegistry[id] = txn
}

func (m *memdbManipulator) unregisterTxn(id manipulate.TransactionID) {

	m.txnRegistryLock.Lock()
	defer m.txnRegistryLock.Unlock()
	delete(m.txnRegistry, id)
}

func (m *memdbManipulator) registeredTxnWithID(id manipulate.TransactionID) *memdb.Txn {

	m.txnRegistryLock.RLock()
	defer m.txnRegistryLock.RUnlock()
	b := m.txnRegistry[id]

	return b
}

// RetrieveFromFilter compiles the given manipulate Filter into a mongo filter.
func (m *memdbManipulator) retrieveFromFilter(identity string, f *elemental.Filter, items *map[string]elemental.Identifiable, fullQuery bool) error {

	if f == nil {
		return m.retrieveIntersection(identity, "id", nil, items, fullQuery)
	}

	if len(f.Operators()) == 0 {
		return nil
	}

	for i, operator := range f.Operators() {

		switch operator {

		case elemental.AndOperator:

			k := strings.ToLower(f.Keys()[i])

			switch f.Comparators()[i] {

			case elemental.EqualComparator:

				if err := m.retrieveIntersection(identity, k, f.Values()[i][0], items, fullQuery); err != nil {
					return err
				}

			case elemental.MatchComparator:

				values := f.Values()[i]

				for _, v := range values {

					if !strings.HasPrefix(v.(string), "^") {
						return manipulate.NewErrCannotExecuteQuery("Matches filter only works for prefix matching and must always start with a '^'")
					}

					fv := strings.TrimPrefix(v.(string), "^")
					fv = strings.TrimSuffix(fv, "$")

					valueItems := map[string]elemental.Identifiable{}
					if err := m.retrieveIntersection(identity, k+"_prefix", fv, &valueItems, fullQuery); err != nil {
						return err
					}
					mergeIn(items, &valueItems)
				}

			case elemental.ContainComparator:

				values := f.Values()[i]

				containItems := map[string]elemental.Identifiable{}

				for _, value := range values {
					valueItems := map[string]elemental.Identifiable{}
					if err := m.retrieveIntersection(identity, k, value, &valueItems, true); err != nil {
						return err
					}
					mergeIn(&containItems, &valueItems)
				}

				intersection(items, &containItems, fullQuery)

			default:
				return manipulate.NewErrCannotExecuteQuery(fmt.Sprintf("invalid comparator for memdb: %d", f.Comparators()[i]))
			}

		case elemental.AndFilterOperator:

			for _, sub := range f.AndFilters()[i] {
				if err := m.retrieveFromFilter(identity, sub, items, fullQuery); err != nil {
					return err
				}
				fullQuery = false
			}

		case elemental.OrFilterOperator:

			orItems := map[string]elemental.Identifiable{}

			for _, sub := range f.OrFilters()[i] {
				valueItems := map[string]elemental.Identifiable{}

				if err := m.retrieveFromFilter(identity, sub, &valueItems, true); err != nil {
					return err
				}

				mergeIn(&orItems, &valueItems)
			}

			intersection(items, &orItems, fullQuery)

		default:
			return manipulate.NewErrCannotExecuteQuery(fmt.Sprintf("invalid operator for memdb: %d", operator))
		}

		fullQuery = false
	}

	return nil
}

func (m *memdbManipulator) retrieveIntersection(identity string, k string, value interface{}, items *map[string]elemental.Identifiable, fullquery bool) error {

	var iterator memdb.ResultIterator
	var err error

	existingItems := *items

	txn := m.getDB().Txn(false)

	if value == nil {
		iterator, err = txn.Get(identity, k)
	} else {
		iterator, err = txn.Get(identity, k, value)
	}
	if err != nil {
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	combinedItems := map[string]elemental.Identifiable{}

	raw := iterator.Next()

	for raw != nil {

		var o interface{}
		if m.noCopy {
			o = raw
		} else {
			o, err = copystructure.Copy(raw)
			if err != nil {
				return manipulate.NewErrCannotExecuteQuery(err.Error())
			}
		}

		obj, ok := o.(elemental.Identifiable)
		if !ok {
			return manipulate.NewErrCannotExecuteQuery("stored object is not an identifiable")
		}
		if _, ok := existingItems[obj.Identifier()]; ok || fullquery {
			combinedItems[obj.Identifier()] = obj
		}
		raw = iterator.Next()
	}

	*items = combinedItems

	return nil
}

func (m *memdbManipulator) getDB() *memdb.MemDB {

	m.dbLock.RLock()
	defer m.dbLock.RUnlock()

	return m.db
}

func (m *memdbManipulator) setDB(db *memdb.MemDB) {

	m.dbLock.Lock()
	m.db = db
	m.dbLock.Unlock()
}
