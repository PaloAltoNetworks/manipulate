package manipmemory

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"

	memdb "github.com/hashicorp/go-memdb"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"gopkg.in/mgo.v2/bson"
)

type txnRegistry map[manipulate.TransactionID]*memdb.Txn

// A memoryManipulator is an empty manipulator that can be used with ApoMock.
type memdbManipulator struct {
	db              *memdb.MemDB
	txnRegistry     txnRegistry
	txnRegistryLock *sync.Mutex
}

// NewMemoryManipulatorFromDB will create a new memdb manipulator assuming that the DB
// is already provided at the input. This allow multiple instances of the manipulator
// to operate on the same memdb.
func NewMemoryManipulatorFromDB(db *memdb.MemDB) manipulate.TransactionalManipulator {
	return &memdbManipulator{
		db:              db,
		txnRegistryLock: &sync.Mutex{},
		txnRegistry:     txnRegistry{},
	}

}

// NewMemoryManipulator returns a new TransactionalManipulator backed by memdb. It first
// instantiates the memdb with the provided schema.
func NewMemoryManipulator(schema *memdb.DBSchema) (manipulate.TransactionalManipulator, error) {

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	return NewMemoryManipulatorFromDB(db), nil
}

// RetrieveMany is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) RetrieveMany(mctx manipulate.Context, dest elemental.Identifiables) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	items := map[string]elemental.Identifiable{}

	if err := s.retrieveFromFilter(dest.Identity().Category, mctx.Filter(), &items, true); err != nil {
		return err
	}

	out := reflect.ValueOf(dest).Elem()

	for _, obj := range items {
		out.Set(reflect.Append(out, reflect.ValueOf(obj)))
	}

	return nil
}

// Retrieve is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) Retrieve(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	txn := s.db.Txn(false)

	for _, object := range objects {

		raw, err := txn.First(object.Identity().Category, "id", object.Identifier())
		if err != nil {
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}

		if raw == nil {
			return manipulate.NewErrObjectNotFound("cannot find the object for the given ID")
		}

		reflect.ValueOf(object).Elem().Set(reflect.ValueOf(raw).Elem())
	}

	return nil
}

// Create is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) Create(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	tid := mctx.TransactionID()
	txn := s.txnForID(tid)
	defer txn.Abort()

	for _, object := range objects {

		// In caching scenarios the identifier is already set. Do not insert
		// here. We will get it pre-populated from the master DB.
		if object.Identifier() == "" {
			object.SetIdentifier(bson.NewObjectId().Hex())
		}

		if err := txn.Insert(object.Identity().Category, object); err != nil {
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
	}

	if tid == "" {
		s.commitTxn(txn)
	}

	return nil
}

// Update is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) Update(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	tid := mctx.TransactionID()
	txn := s.txnForID(tid)
	defer txn.Abort()

	for _, object := range objects {
		if err := txn.Insert(object.Identity().Category, object); err != nil {
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
	}

	if tid == "" {
		s.commitTxn(txn)
	}

	return nil
}

// Delete is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) Delete(mctx manipulate.Context, objects ...elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	tid := mctx.TransactionID()
	txn := s.txnForID(tid)
	defer txn.Abort()

	for _, object := range objects {
		if err := txn.Delete(object.Identity().Category, object); err != nil {
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
	}

	if tid == "" {
		s.commitTxn(txn)
	}

	return nil
}

// DeleteMany is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) DeleteMany(mctx manipulate.Context, identity elemental.Identity) error {
	return manipulate.NewErrNotImplemented("DeleteMany not implemented in manipmemory")
}

// Count is part of the implementation of the Manipulator interface. Count is very expensive.
func (s *memdbManipulator) Count(mctx manipulate.Context, identity elemental.Identity) (int, error) {

	txn := s.db.Txn(false)

	iterator, err := txn.Get(identity.Category, "id")
	if err != nil {
		return 0, fmt.Errorf("failed to create iterator for %s: %s", identity.Category, err)
	}

	count := 0

	raw := iterator.Next()
	for raw != nil {
		count = count + 1
		raw = iterator.Next()
	}

	return count, nil

}

// Commit is part of the implementation of the TransactionalManipulator interface.
func (s *memdbManipulator) Commit(id manipulate.TransactionID) error {

	txn := s.registeredTxnWithID(id)

	if txn == nil {
		return manipulate.NewErrCannotCommit("Cannot find transaction " + string(id))
	}

	defer func() { s.unregisterTxn(id) }()

	s.commitTxn(txn)

	return nil
}

// Abort is part of the implementation of the TransactionalManipulator interface.
func (s *memdbManipulator) Abort(id manipulate.TransactionID) bool {

	txn := s.registeredTxnWithID(id)
	if txn == nil {
		return false
	}

	txn.Abort()
	s.unregisterTxn(id)

	return true
}

func (s *memdbManipulator) txnForID(id manipulate.TransactionID) *memdb.Txn {

	if id == "" {
		return s.db.Txn(true)
	}

	txn := s.registeredTxnWithID(id)

	if txn == nil {
		txn = s.db.Txn(true)
		s.registerTxn(id, txn)
	}

	return txn
}

func (s *memdbManipulator) commitTxn(t *memdb.Txn) {

	t.Commit()
}

func (s *memdbManipulator) registerTxn(id manipulate.TransactionID, txn *memdb.Txn) {

	s.txnRegistryLock.Lock()
	defer s.txnRegistryLock.Unlock()
	s.txnRegistry[id] = txn
}

func (s *memdbManipulator) unregisterTxn(id manipulate.TransactionID) {

	s.txnRegistryLock.Lock()
	defer s.txnRegistryLock.Unlock()
	delete(s.txnRegistry, id)
}

func (s *memdbManipulator) registeredTxnWithID(id manipulate.TransactionID) *memdb.Txn {

	s.txnRegistryLock.Lock()
	defer s.txnRegistryLock.Unlock()
	b := s.txnRegistry[id]

	return b
}

// RetrieveFromFilter compiles the given manipulate Filter into a mongo filter.
func (s *memdbManipulator) retrieveFromFilter(identity string, f *manipulate.Filter, items *map[string]elemental.Identifiable, fullQuery bool) error {

	if f == nil {
		return s.retrieveIntersection(identity, "id", nil, items, fullQuery)
	}

	if len(f.Operators()) == 0 {
		return nil
	}

	for i, operator := range f.Operators() {

		switch operator {

		case manipulate.AndOperator:

			k := strings.ToLower(f.Keys()[i])

			switch f.Comparators()[i] {

			case manipulate.EqualComparator:

				if err := s.retrieveIntersection(identity, k, f.Values()[i][0], items, fullQuery); err != nil {
					return err
				}

			case manipulate.ContainComparator:

				values := f.Values()[i]

				containItems := map[string]elemental.Identifiable{}

				for _, value := range values {
					valueItems := map[string]elemental.Identifiable{}
					if err := s.retrieveIntersection(identity, k, value, &valueItems, true); err != nil {
						return err
					}
					mergeIn(&containItems, &valueItems)
				}

				intersection(items, &containItems, fullQuery)

			default:
				return manipulate.NewErrCannotExecuteQuery(fmt.Sprintf("invalid comparator for memdb: %d", f.Comparators()[i]))
			}

		case manipulate.AndFilterOperator:

			for _, sub := range f.AndFilters()[i] {
				if err := s.retrieveFromFilter(identity, sub, items, fullQuery); err != nil {
					return err
				}
				fullQuery = false
			}

		case manipulate.OrFilterOperator:

			orItems := map[string]elemental.Identifiable{}

			for _, sub := range f.OrFilters()[i] {
				valueItems := map[string]elemental.Identifiable{}

				if err := s.retrieveFromFilter(identity, sub, &valueItems, true); err != nil {
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

func (s *memdbManipulator) retrieveIntersection(identity string, k string, value interface{}, items *map[string]elemental.Identifiable, fullquery bool) error {

	var iterator memdb.ResultIterator
	var err error

	existingItems := *items

	txn := s.db.Txn(false)

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
		obj := raw.(elemental.Identifiable)
		if _, ok := existingItems[obj.Identifier()]; ok || fullquery {
			combinedItems[obj.Identifier()] = obj
		}
		raw = iterator.Next()
	}

	*items = combinedItems

	return nil
}

func mergeIn(target, source *map[string]elemental.Identifiable) {
	for k, v := range *source {
		(*target)[k] = v
	}
}

func intersection(target, source *map[string]elemental.Identifiable, queryStart bool) {

	combined := map[string]elemental.Identifiable{}

	for k, v := range *source {
		if _, ok := (*target)[k]; ok || queryStart {
			combined[k] = v
		}
	}

	*target = combined
}
