package manipulate

import (
	"context"
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
func ContextOptionRecursive() ContextOption {
	return func(c *mcontext) {
		c.recursive = true
	}
}

// ContextOptionVersion sets the version option of the context.
func ContextOptionVersion(v int) ContextOption {
	return func(c *mcontext) {
		c.version = v
	}
}

// ContextOptionVersion sets the override option of the context.
func ContextOptionOverride() ContextOption {
	return func(c *mcontext) {
		c.overrideProtection = true
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

// ContextOptionContext sets the context.Context option of the context.
func ContextOptionContext(ctx context.Context) ContextOption {

	if ctx == nil {
		panic("nil context")
	}

	return func(c *mcontext) {
		c.ctx = ctx
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
