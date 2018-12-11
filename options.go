package manipulate

import (
	"net/url"

	"go.aporeto.io/elemental"
)

// ContextOption represents an option can can be passed to NewContext.
type ContextOption func(*mcontext)

// ContextOptionFilter sets the filter.
func ContextOptionFilter(f *Filter) ContextOption {
	return func(c *mcontext) {
		c.filter = f
	}
}

// ContextOptionNamespace sets the namespace.
func ContextOptionNamespace(n string) ContextOption {
	return func(c *mcontext) {
		c.namespace = n
	}
}

// ContextOptionRecursive sets the recursive option of the context.
func ContextOptionRecursive(r bool) ContextOption {
	return func(c *mcontext) {
		c.recursive = r
	}
}

// ContextOptionVersion sets the version option of the context.
func ContextOptionVersion(v int) ContextOption {
	return func(c *mcontext) {
		c.version = v
	}
}

// ContextOptionOverride sets the override option of the context.
func ContextOptionOverride(o bool) ContextOption {
	return func(c *mcontext) {
		c.overrideProtection = o
	}
}

// ContextOptionPage sets the pagination option of the context.
func ContextOptionPage(n, size int) ContextOption {
	return func(c *mcontext) {
		c.page = n
		c.pageSize = size
	}
}

// ContextOptionTracking sets the opentracing tracking option of the context.
func ContextOptionTracking(identifier, typ string) ContextOption {
	return func(c *mcontext) {
		c.externalTrackingID = identifier
		c.externalTrackingType = typ
	}
}

// ContextOptionOrder sets the ordering option of the context.
func ContextOptionOrder(orders ...string) ContextOption {
	return func(c *mcontext) {
		c.order = orders
	}
}

// ContextOptionParameters sets the parameters option of the context.
func ContextOptionParameters(p url.Values) ContextOption {
	return func(c *mcontext) {
		c.parameters = p
	}
}

// ContextOptionFinalizer sets the create finalizer option of the context.
func ContextOptionFinalizer(f FinalizerFunc) ContextOption {
	return func(c *mcontext) {
		c.createFinalizer = f
	}
}

// ContextOptionTransationID sets the parameters option of the context.
func ContextOptionTransationID(tid TransactionID) ContextOption {
	return func(c *mcontext) {
		c.transactionID = tid
	}
}

// ContextOptionParent sets the parent option of the context.
func ContextOptionParent(i elemental.Identifiable) ContextOption {
	return func(c *mcontext) {
		c.parent = i
	}
}

// ContextOptionFields sets the list of fields to include in the reply.
func ContextOptionFields(fields []string) ContextOption {
	return func(c *mcontext) {
		c.fields = fields
	}
}

// ContextOptionWriteConsistency sets the desired write consistency of the request.
func ContextOptionWriteConsistency(consistency WriteConsistency) ContextOption {
	return func(c *mcontext) {
		c.writeConsistency = consistency
	}
}

// ContextOptionReadConsistency sets the desired read consistency of the request.
func ContextOptionReadConsistency(consistency ReadConsistency) ContextOption {
	return func(c *mcontext) {
		c.readConsistency = consistency
	}
}
