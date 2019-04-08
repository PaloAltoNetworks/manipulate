package manipvortex

import (
	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

// Transaction is the event that captures the transaction for later processing. It is
// also the structure stored in the transaction logs.
type Transaction struct {
	Date     time.Time
	mctx     manipulate.Context
	Object   elemental.Identifiable
	Method   elemental.Operation
	Deadline time.Time
}
