package manipmemory

import (
	"reflect"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"

	log "github.com/Sirupsen/logrus"
	memdb "github.com/hashicorp/go-memdb"
	uuid "github.com/satori/go.uuid"
)

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
		log.WithFields(log.Fields{
			"package": "manipmemory",
			"error":   err.Error(),
		}).Fatal("Cannot initialize MemDB.")
	}

	return &memdbManipulator{
		db: db,
	}
}

// RetrieveMany is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) RetrieveMany(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {

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

	iterator, err := txn.Get(identity.Category, index, args...)

	if err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
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
func (s *memdbManipulator) Retrieve(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	txn := s.db.Txn(false)

	for _, object := range objects {

		raw, err := txn.First(object.Identity().Category, "id", object.Identifier())

		if err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
		}

		if reflect.ValueOf(object).Kind() != reflect.Ptr {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotUnmarshal)
		}

		reflect.ValueOf(object).Elem().Set(reflect.ValueOf(raw).Elem())
	}

	return nil
}

// Create is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) Create(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	tid := context.TransactionID
	txn := s.txnForID(tid)
	defer txn.Abort()

	for _, object := range objects {
		object.SetIdentifier(uuid.NewV4().String())

		if err := txn.Insert(object.Identity().Category, object); err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
		}
	}

	if tid == "" {
		if err := s.commitTxn(txn); err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotCommit)
		}
	}

	return nil
}

// Update is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) Update(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	tid := context.TransactionID
	txn := s.txnForID(tid)
	defer txn.Abort()

	for _, object := range objects {
		if err := txn.Insert(object.Identity().Category, object); err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
		}
	}

	if tid == "" {
		if err := s.commitTxn(txn); err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotCommit)
		}
	}

	return nil
}

// Delete is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) Delete(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	tid := context.TransactionID
	txn := s.txnForID(tid)
	defer txn.Abort()

	for _, object := range objects {
		if err := txn.Delete(object.Identity().Category, object); err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
		}
	}

	if tid == "" {
		if err := s.commitTxn(txn); err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotCommit)
		}
	}

	return nil
}

// Count is part of the implementation of the Manipulator interface.
func (s *memdbManipulator) Count(context *manipulate.Context, identity elemental.Identity) (int, error) {

	out := manipulate.ManipulablesList{}
	s.RetrieveMany(context, identity, &out)
	return len(out), nil
}

// Assign is part of the implementation of the Manipulator interface.
func (*memdbManipulator) Assign(context *manipulate.Context, assignation *elemental.Assignation) error {
	return nil
}

// Increment is part of the implementation of the Manipulator interface.
func (*memdbManipulator) Increment(context *manipulate.Context, name string, counter string, inc int, filterKeys []string, filterValues []interface{}) error {
	return nil
}

// Commit is part of the implementation of the TransactionalManipulator interface.
func (s *memdbManipulator) Commit(id manipulate.TransactionID) error {

	txn := s.txnForID(id)

	if txn == nil {
		log.WithFields(log.Fields{
			"package":       "manipmemory",
			"store":         s,
			"transactionID": id,
		}).Error("No transaction found for the given transaction ID.")

		return manipulate.NewError("No transaction found for the given transaction ID.", manipulate.ErrCannotCommit)
	}

	defer func() { s.unregisterTxn(id) }()

	s.commitTxn(txn)

	return nil
}

// Abort is part of the implementation of the TransactionalManipulator interface.
func (*memdbManipulator) Abort(id manipulate.TransactionID) bool {
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

func (s *memdbManipulator) commitTxn(t *memdb.Txn) error {

	log.WithFields(log.Fields{
		"transaction": t,
	}).Debug("Commiting transaction to MemDB.")

	t.Commit()

	return nil
}

func (s *memdbManipulator) registerTxn(id manipulate.TransactionID, txn *memdb.Txn) {

	s.txnRegistryLock.Lock()
	s.txnRegistry[id] = txn
	s.txnRegistryLock.Unlock()
}

func (s *memdbManipulator) unregisterTxn(id manipulate.TransactionID) {

	s.txnRegistryLock.Lock()
	delete(s.txnRegistry, id)
	s.txnRegistryLock.Unlock()
}

func (s *memdbManipulator) registeredTxnWithID(id manipulate.TransactionID) *memdb.Txn {

	s.txnRegistryLock.Lock()
	b := s.txnRegistry[id]
	s.txnRegistryLock.Unlock()

	return b
}
