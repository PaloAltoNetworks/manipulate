package manipmongo

import (
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"

	mgo "gopkg.in/mgo.v2"
)

type transaction struct {
	bulks   map[elemental.Identity]*mgo.Bulk
	db      *mgo.Database
	id      manipulate.TransactionID
	lock    *sync.Mutex
	session *mgo.Session
}

func newTransaction(id manipulate.TransactionID, rootSession *mgo.Session, dbName string) *transaction {

	s := rootSession.Copy()

	return &transaction{
		bulks:   map[elemental.Identity]*mgo.Bulk{},
		db:      s.DB(dbName),
		id:      id,
		lock:    &sync.Mutex{},
		session: s,
	}
}

func (t *transaction) closeSession() {

	if t.session == nil {
		return
	}

	t.session.Close()
	t.session = nil
}

func (t *transaction) bulkForIdentity(identity elemental.Identity) *mgo.Bulk {

	t.lock.Lock()
	defer t.lock.Unlock()

	bulk := t.bulks[identity]
	if bulk != nil {
		return bulk
	}

	t.bulks[identity] = t.db.C(identity.Name).Bulk()

	return t.bulks[identity]
}
