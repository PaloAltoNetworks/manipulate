package manipmongo

import (
	"context"
	"sync"

	"github.com/globalsign/mgo"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

type transaction struct {
	bulks map[elemental.Identity]*mgo.Bulk
	db    *mgo.Database
	id    manipulate.TransactionID
	lock  *sync.Mutex
	ctx   context.Context
}

func newTransaction(ctx context.Context, id manipulate.TransactionID, db *mgo.Database) *transaction {

	return &transaction{
		bulks: map[elemental.Identity]*mgo.Bulk{},
		db:    db,
		id:    id,
		lock:  &sync.Mutex{},
		ctx:   ctx,
	}
}

func (t *transaction) bulkForIdentity(identity elemental.Identity) *mgo.Bulk { // nolint:unparam

	t.lock.Lock()
	defer t.lock.Unlock()

	bulk := t.bulks[identity]
	if bulk != nil {
		return bulk
	}

	t.bulks[identity] = collectionFromIdentity(t.db, identity).Bulk()

	return t.bulks[identity]
}
