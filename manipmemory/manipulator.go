package manipmemory

import (
	"reflect"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"

	memdb "github.com/hashicorp/go-memdb"
	uuid "github.com/satori/go.uuid"
)

// Logger contains the main logger.
var Logger = logrus.New()

var log = Logger.WithField("package", "github.com/aporeto-inc/manipulate/manipmemory")

type txnRegistry map[manipulate.TransactionID]*memdb.Txn

// A memoryManipulator is an empty manipulator that can be used with ApoMock.
type memdbManipulator struct {
	db              *memdb.MemDB
	txnRegistry     txnRegistry
	txnRegistryLock *sync.Mutex
}

// NewMemoryManipulator returns a new TransactionalManipulator backed by memdb.
func NewMemoryManipulator(schema *memdb.DBSchema) manipulate.TransactionalManipulator {

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	return &memdbManipulator{
		db:              db,
		txnRegistryLock: &sync.Mutex{},
		txnRegistry:     txnRegistry{},
	}
}

// RetrieveMany is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) RetrieveMany(context *manipulate.Context, dest elemental.ContentIdentifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	txn := s.db.Txn(false)

	index := "id"
	args := []interface{}{}
	if context.Filter != nil {
		index = context.Filter.Keys()[0][0]
		args = context.Filter.Values()[0]
	}

	iterator, err := txn.Get(dest.ContentIdentity().Category, index, args...)

	if err != nil {
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	out := reflect.ValueOf(dest).Elem()

	raw := iterator.Next()
	for raw != nil {
		out.Set(reflect.Append(out, reflect.ValueOf(raw)))
		raw = iterator.Next()
	}

	return nil
}

// Retrieve is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) Retrieve(context *manipulate.Context, objects ...elemental.Identifiable) error {

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
func (s *memdbManipulator) Create(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	tid := context.TransactionID
	txn := s.txnForID(tid)
	defer txn.Abort()

	for _, object := range objects {

		object.SetIdentifier(uuid.NewV4().String())

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
func (s *memdbManipulator) Update(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	tid := context.TransactionID
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
func (s *memdbManipulator) Delete(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	tid := context.TransactionID
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
func (s *memdbManipulator) DeleteMany(context *manipulate.Context, identity elemental.Identity) error {
	return manipulate.NewErrNotImplemented("DeleteMany not implemented in manipmemory")
}

// Count is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) Count(context *manipulate.Context, identity elemental.Identity) (int, error) {

	// out := elemental.IdentifiablesList{}
	// if err := s.RetrieveMany(context, &out); err != nil {
	// 	return -1, err
	// }
	//
	// return len(out), nil

	return 0, nil
}

// Assign is part of the implementation of the Manipulator interface.
func (*memdbManipulator) Assign(context *manipulate.Context, assignation *elemental.Assignation) error {

	return manipulate.NewErrNotImplemented("Assign not implemented in memory manipulator")
}

// Increment is part of the implementation of the Manipulator interface.
func (*memdbManipulator) Increment(context *manipulate.Context, identity elemental.Identity, counter string, inc int) error {

	return manipulate.NewErrNotImplemented("Increment not implemented in memory manipulator")
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

	log.WithField("transaction", t).Debug("Commiting transaction to MemDB.")

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
