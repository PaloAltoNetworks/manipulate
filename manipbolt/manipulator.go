package manipbolt

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"gopkg.in/mgo.v2/bson"
)

var (
	_ manipulate.Manipulator              = &boltManipulator{}
	_ manipulate.TransactionalManipulator = &boltManipulator{}
	_ manipulate.FlushableManipulator     = &boltManipulator{}
	_ manipulate.BufferedManipulator      = &boltManipulator{}
)

type txnRegistry map[manipulate.TransactionID]storm.Node

type boltManipulator struct {
	db          *storm.DB
	txnRegistry txnRegistry
	manager     elemental.ModelManager
	cfg         *config

	txnRegistryLock sync.RWMutex
	dbLock          sync.RWMutex
}

// New creates a new datastore backed by a boltdb with storm toolkit.
func New(path string, manager elemental.ModelManager, options ...Option) (manipulate.TransactionalManipulator, error) {

	cfg := newConfig()
	for _, opt := range options {
		opt(cfg)
	}

	db, err := storm.Open(path, storm.Codec(cfg.codec))
	if err != nil {
		return nil, manipulate.ErrCannotExecuteQuery{Err: err}
	}

	return &boltManipulator{
		db:          db,
		txnRegistry: txnRegistry{},
		manager:     manager,
		cfg:         cfg,
	}, nil
}

// Flush will flush the datastore essentially creating a new one.
// It will wait until all transactions are complete. The caller
// has to ensure no other methods are called on the  db
// unless the current Flush operation is complete.
func (b *boltManipulator) Flush(ctx context.Context) error {

	path := b.getDB().Bolt.Path()

	if err := b.getDB().Close(); err != nil {
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	if err := os.RemoveAll(path); err != nil {
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	db, err := storm.Open(path, storm.Codec(b.cfg.codec))
	if err != nil {
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	b.setDB(db)

	return nil
}

// RetrieveMany retrieves the a list of objects with the given elemental.Identity and put them in the given dest.
func (b *boltManipulator) RetrieveMany(mctx manipulate.Context, dest elemental.Identifiables) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	txn, err := b.getDB().Begin(false)
	if err != nil {
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	if err := b.executeWithFilter(txn, getOp, mctx.Filter(), dest); err != nil {
		txn.Rollback() // nolint: errcheck
		return err
	}

	if err := txn.Rollback(); err != nil {
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	return nil
}

// Retrieve retrieves one or multiple elemental.Identifiables.
// In order to be retrievable, the elemental.Identifiable needs to have their Identifier correctly set.
func (b *boltManipulator) Retrieve(mctx manipulate.Context, object elemental.Identifiable) error {

	txn, err := b.getDB().Begin(false)
	if err != nil {
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	if err := txn.One("ID", object.Identifier(), object); err != nil {
		txn.Rollback() // nolint: errcheck
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	if err := txn.Rollback(); err != nil {
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	return nil
}

// Create creates a the given elemental.Identifiable.
func (b *boltManipulator) Create(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	tid := mctx.TransactionID()
	txn, err := b.txnForID(tid)
	if err != nil {
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	if object.Identifier() == "" {
		object.SetIdentifier(bson.NewObjectId().Hex())
	}

	if err := txn.Save(object); err != nil {
		txn.Rollback() // nolint: errcheck
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	if err := b.commit(tid, txn); err != nil {
		return manipulate.ErrCannotCommit{Err: err}
	}

	return nil
}

// Update updates the given elemental.Identifiable.
// In order to be updatable, the elemental.Identifiable needs to have their Identifier correctly set.
func (b *boltManipulator) Update(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	tid := mctx.TransactionID()
	txn, err := b.txnForID(tid)
	if err != nil {
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	obj := b.manager.Identifiable(object.Identity())
	if err := txn.One("ID", object.Identifier(), obj); err != nil {
		txn.Rollback() // nolint: errcheck
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	if err := txn.Save(object); err != nil {
		txn.Rollback() // nolint: errcheck
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	if err := b.commit(tid, txn); err != nil {
		return manipulate.ErrCannotCommit{Err: err}
	}

	return nil
}

// Delete deletes the given elemental.Identifiable.
// In order to be deletable, the elemental.Identifiable needs to have their Identifier correctly set.
func (b *boltManipulator) Delete(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	tid := mctx.TransactionID()
	txn, err := b.txnForID(tid)
	if err != nil {
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	if err := txn.DeleteStruct(object); err != nil {
		txn.Rollback() // nolint: errcheck
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	if err := b.commit(tid, txn); err != nil {
		return manipulate.ErrCannotCommit{Err: err}
	}

	return nil
}

// DeleteMany deletes all objects of with the given identity or
// all the ones matching the filter in the given context.
func (b *boltManipulator) DeleteMany(mctx manipulate.Context, identity elemental.Identity) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	tid := mctx.TransactionID()
	txn, err := b.txnForID(tid)
	if err != nil {
		return manipulate.ErrCannotExecuteQuery{Err: err}
	}

	if err := b.executeWithFilter(txn, deleteOp, mctx.Filter(), b.manager.Identifiables(identity)); err != nil {
		txn.Rollback() // nolint: errcheck
		return err
	}

	if err := b.commit(tid, txn); err != nil {
		return manipulate.ErrCannotCommit{Err: err}
	}

	return nil
}

// Count returns the number of objects with the given identity.
func (b *boltManipulator) Count(mctx manipulate.Context, identity elemental.Identity) (int, error) {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	txn, err := b.getDB().Begin(false)
	if err != nil {
		return -1, manipulate.ErrCannotExecuteQuery{Err: err}
	}

	objs := b.manager.Identifiables(identity)
	if err := b.executeWithFilter(txn, countOp, mctx.Filter(), objs); err != nil {
		txn.Rollback() // nolint: errcheck
		return -1, err
	}

	if err := txn.Rollback(); err != nil {
		return -1, manipulate.ErrCannotExecuteQuery{Err: err}
	}

	return len(objs.List()), nil
}

// Commit commits the given TransactionID.
func (b *boltManipulator) Commit(id manipulate.TransactionID) error {

	txn := b.registeredTxnWithID(id)
	if txn == nil {
		return manipulate.ErrTransactionNotFound{Err: fmt.Errorf("cannot find transaction: %s", string(id))}
	}

	if err := b.commit("", txn); err != nil {
		return manipulate.ErrCannotCommit{Err: err}
	}

	b.unregisterTxn(id)

	return nil
}

// Abort aborts the give TransactionID. It returns true if
// a transaction has been effectively aborted, otherwise it returns false.
func (b *boltManipulator) Abort(id manipulate.TransactionID) bool {

	txn := b.registeredTxnWithID(id)
	if txn == nil {
		return false
	}

	if err := txn.Rollback(); err != nil {
		// TODO: log error ? interface doesn't support error reporting.
		return false
	}

	b.unregisterTxn(id)

	return true
}

func (b *boltManipulator) executeWithFilter(txn storm.Node, ops operation, f *elemental.Filter, dest elemental.Identifiables) error {

	if dest == nil {
		return manipulate.ErrCannotExecuteQuery{Err: fmt.Errorf("destination cannot be nil")}
	}

	if f == nil {
		return b.executeWithMatcher(txn, ops, q.True(), dest)
	}

	if len(f.Operators()) == 0 {
		return nil
	}

	matcher, err := compileFilter(f)
	if err != nil {
		return err
	}

	return b.executeWithMatcher(txn, ops, matcher, dest)
}

func (b *boltManipulator) executeWithMatcher(txn storm.Node, ops operation, matcher q.Matcher, dest elemental.Identifiables) error {

	// TODO: the ifshort linter shouldn't complain about this
	query := txn.Select(matcher) // nolint: ifshort
	if query == nil {
		return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("bad query")}
	}

	switch ops {

	case getOp, countOp:

		if err := query.Find(dest); err != nil {
			if err == storm.ErrNotFound || err == q.ErrUnknownField {
				return nil
			}

			return manipulate.ErrCannotBuildQuery{Err: err}
		}

	case deleteOp:

		obj := b.manager.Identifiable(dest.Identity())
		if err := query.Delete(obj); err != nil {
			if err == storm.ErrNotFound || err == q.ErrUnknownField {
				return nil
			}

			return manipulate.ErrCannotBuildQuery{Err: err}
		}

	default:
		return manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("invalid operation: %d", ops)}
	}

	return nil
}

func (b *boltManipulator) txnForID(id manipulate.TransactionID) (storm.Node, error) {

	if id == "" {
		return b.getDB().Begin(true)
	}

	if txn := b.registeredTxnWithID(id); txn != nil {
		return txn, nil
	}

	txn, err := b.getDB().Begin(true)
	if err != nil {
		return nil, err
	}

	b.registerTxn(id, txn)

	return txn, nil
}

func (b *boltManipulator) registerTxn(id manipulate.TransactionID, txn storm.Node) {

	b.txnRegistryLock.Lock()
	defer b.txnRegistryLock.Unlock()
	b.txnRegistry[id] = txn
}

func (b *boltManipulator) unregisterTxn(id manipulate.TransactionID) {

	b.txnRegistryLock.Lock()
	defer b.txnRegistryLock.Unlock()
	delete(b.txnRegistry, id)
}

func (b *boltManipulator) registeredTxnWithID(id manipulate.TransactionID) storm.Node {

	b.txnRegistryLock.RLock()
	defer b.txnRegistryLock.RUnlock()
	t := b.txnRegistry[id]

	return t
}

func (b *boltManipulator) getDB() *storm.DB {

	b.dbLock.RLock()
	defer b.dbLock.RUnlock()

	return b.db
}

func (b *boltManipulator) setDB(db *storm.DB) {

	b.dbLock.Lock()
	b.db = db
	b.dbLock.Unlock()
}

func (b *boltManipulator) commit(tid manipulate.TransactionID, txn storm.Node) error {

	if tid != "" {
		return nil
	}

	if err := txn.Commit(); err != nil {
		return txn.Rollback()
	}

	return nil
}
