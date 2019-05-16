// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package manipulate

import (
	"net/url"

	"go.aporeto.io/elemental"
)

// ContextOption represents an option can can be passed to NewContext.
type ContextOption func(*mcontext)

// ContextOptionFilter sets the filter.
func ContextOptionFilter(f *elemental.Filter) ContextOption {
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

// ContextOptionCredentials sets user name and password for this context.
func ContextOptionCredentials(username, password string) ContextOption {
	return func(c *mcontext) {
		c.username = username
		c.password = password
	}
}

// ContextOptionToken sets the token for this request.
func ContextOptionToken(token string) ContextOption {
	return func(c *mcontext) {
		c.username = "Bearer"
		c.password = token
	}
}
