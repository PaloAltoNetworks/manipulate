package manipulate

import uuid "github.com/satori/go.uuid"

// TransactionID is the type used to define a transcation ID of a store
type TransactionID string

// NewTransactionID returns a new transaction ID.
func NewTransactionID() TransactionID {

	return TransactionID(uuid.Must(uuid.NewV4()).String())
}
