package manipvortex

import (
	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

// CacheMode is the mode of the cache.
type CacheMode int

// Values of CacheMode
const (
	// WriteThrough means that all transactions must be
	// first written the main data store and then to the local
	// memory.
	WriteThrough CacheMode = iota

	// WriteBack means that writes will be committed locally
	// and lazily synced with the main data store. These objects
	// will not be accessible until they are actually synced
	// with the backend since they don't have a unique ID
	// yet.
	WriteBack
)

// Hook is the hook function type that is called by the processors.
type Hook func(method elemental.Operation, mctx manipulate.Context, objects []elemental.Identifiable) (reconcile bool, err error)

// RetrieveManyHook is the type of a hook for retrieve many
type RetrieveManyHook func(m manipulate.Manipulator, mctx manipulate.Context, dest elemental.Identifiables) (reconcile bool, err error)

// ProcessorConfiguration configures the processing details for a specific identity.
type ProcessorConfiguration struct {

	// Identity is the identity of the object that is stored in the DB.
	Identity elemental.Identity

	// Mode is the type of default consistency mode required from the cache.
	// This consistency can be overwritten by manipulate options.
	Mode CacheMode

	// QueueingDuration is the maximum time that an object should be
	// cached if the backend is not responding.
	QueueingDuration time.Duration

	// RetrieveManyHook is a hook function that can be called
	// before a RetrieveMany call. It returns an error and a continue
	// boolean. If the continue false, we can return without any
	// additional calls.
	RetrieveManyHook RetrieveManyHook

	// UpstreamHook is a hook function that can be called before a transaction
	// is performed on the upstream. If the hook returns an error or reconcile
	// equals to fals, the transaction is aborted.
	UpstreamHook Hook

	// LocalHook is a hook function that can be called before a transaction
	// is committed locally. If the hook returns an error or reconcile
	// is false, the transaction will be aborted. The error will be returned.
	LocalHook Hook

	// CommitOnEvent with commit the event in the cache even if a client
	// is subscribed for this event. If left false, it will only commit
	// if no client has subscribed for this event. It will always forward
	// the event to the clients.
	CommitOnEvent bool

	// LazySync will not sync all data of the identity on startup, but
	// will only load data on demand based on the transactions.
	LazySync bool
}
