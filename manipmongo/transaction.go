package manipmongo

import (
	"context"
	"sync"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"github.com/globalsign/mgo"
)

type transaction struct {
	bulks   map[elemental.Identity]*mgo.Bulk
	db      *mgo.Database
	id      manipulate.TransactionID
	lock    *sync.Mutex
	session *mgo.Session
	ctx     context.Context
}

func newTransaction(ctx context.Context, id manipulate.TransactionID, session *mgo.Session, dbName string) *transaction {

	return &transaction{
		bulks:   map[elemental.Identity]*mgo.Bulk{},
		db:      session.DB(dbName),
		id:      id,
		lock:    &sync.Mutex{},
		session: session,
		ctx:     ctx,
	}
}

func (t *transaction) closeSession() {

	if t.session == nil {
		return
	}

	t.session.Close()
	t.session = nil
}

func (t *transaction) bulkForIdentity(identity elemental.Identity, prefix string) *mgo.Bulk {

	t.lock.Lock()
	defer t.lock.Unlock()

	bulk := t.bulks[identity]
	if bulk != nil {
		return bulk
	}

	t.bulks[identity] = collectionFromIdentity(t.db, identity, prefix).Bulk()

	return t.bulks[identity]
}
