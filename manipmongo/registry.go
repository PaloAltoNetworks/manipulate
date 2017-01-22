package manipmongo

import (
	"sync"

	"github.com/aporeto-inc/manipulate"
)

type transactionsRegistry struct {
	registry map[manipulate.TransactionID]*transaction
	lock     *sync.Mutex
}

func newTransactionRegistry() *transactionsRegistry {

	return &transactionsRegistry{
		registry: map[manipulate.TransactionID]*transaction{},
		lock:     &sync.Mutex{},
	}
}

func (r *transactionsRegistry) transactionWithID(id manipulate.TransactionID) *transaction {

	r.lock.Lock()
	t := r.registry[id]
	r.lock.Unlock()
	return t
}

func (r *transactionsRegistry) registerTransaction(id manipulate.TransactionID, t *transaction) {

	r.lock.Lock()
	r.registry[id] = t
	r.lock.Unlock()
}

func (r *transactionsRegistry) unregisterTransaction(id manipulate.TransactionID) {

	r.lock.Lock()
	delete(r.registry, id)
	r.lock.Unlock()
}
