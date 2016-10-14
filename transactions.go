package manipulate

import "github.com/gocql/gocql"

// TransactionID is the type used to define a transcation ID of a store
type TransactionID string

// NewTransactionID returns a new transaction ID.
func NewTransactionID() TransactionID {

	return TransactionID(gocql.TimeUUID().String())
}
