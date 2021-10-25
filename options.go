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
type ContextOption func(Context)

// ContextOptionFilter sets the filter.
func ContextOptionFilter(f *elemental.Filter) ContextOption {
	return func(c Context) {
		c.(*mcontext).filter = f
	}
}

// ContextOptionNamespace sets the namespace.
func ContextOptionNamespace(n string) ContextOption {
	return func(c Context) {
		c.(*mcontext).namespace = n
	}
}

// ContextOptionPropagate sets the propagate option of the context.
func ContextOptionPropagate(p bool) ContextOption {
	return func(c Context) {
		c.(*mcontext).propagate = p
	}
}

// ContextOptionRecursive sets the recursive option of the context.
func ContextOptionRecursive(r bool) ContextOption {
	return func(c Context) {
		c.(*mcontext).recursive = r
	}
}

// ContextOptionVersion sets the version option of the context.
func ContextOptionVersion(v int) ContextOption {
	return func(c Context) {
		c.(*mcontext).version = v
	}
}

// ContextOptionOverride sets the override option of the context.
func ContextOptionOverride(o bool) ContextOption {
	return func(c Context) {
		c.(*mcontext).overrideProtection = o
	}
}

// ContextOptionPage sets the pagination option of the context.
func ContextOptionPage(n, size int) ContextOption {
	return func(c Context) {
		c.(*mcontext).page = n
		c.(*mcontext).pageSize = size
	}
}

// ContextOptionAfter sets the lazy pagination option of the context.
func ContextOptionAfter(from string, limit int) ContextOption {
	return func(c Context) {
		c.(*mcontext).after = from
		c.(*mcontext).limit = limit
	}
}

// ContextOptionTracking sets the opentracing tracking option of the context.
func ContextOptionTracking(identifier, typ string) ContextOption {
	return func(c Context) {
		c.(*mcontext).externalTrackingID = identifier
		c.(*mcontext).externalTrackingType = typ
	}
}

// ContextOptionOrder sets the ordering option of the context.
func ContextOptionOrder(orders ...string) ContextOption {
	return func(c Context) {
		c.(*mcontext).order = orders
	}
}

// ContextOptionParameters sets the parameters option of the context.
func ContextOptionParameters(p url.Values) ContextOption {
	return func(c Context) {
		c.(*mcontext).parameters = p
	}
}

// ContextOptionFinalizer sets the create finalizer option of the context.
func ContextOptionFinalizer(f FinalizerFunc) ContextOption {
	return func(c Context) {
		c.(*mcontext).createFinalizer = f
	}
}

// ContextOptionTransactionID sets the parameters option of the context.
func ContextOptionTransactionID(tid TransactionID) ContextOption {
	return func(c Context) {
		c.(*mcontext).transactionID = tid
	}
}

// ContextOptionParent sets the parent option of the context.
func ContextOptionParent(i elemental.Identifiable) ContextOption {
	return func(c Context) {
		c.(*mcontext).parent = i
	}
}

// ContextOptionFields sets the list of fields to include in the reply.
func ContextOptionFields(fields []string) ContextOption {
	return func(c Context) {
		c.(*mcontext).fields = fields
	}
}

// ContextOptionWriteConsistency sets the desired write consistency of the request.
func ContextOptionWriteConsistency(consistency WriteConsistency) ContextOption {
	return func(c Context) {
		c.(*mcontext).writeConsistency = consistency
	}
}

// ContextOptionReadConsistency sets the desired read consistency of the request.
func ContextOptionReadConsistency(consistency ReadConsistency) ContextOption {
	return func(c Context) {
		c.(*mcontext).readConsistency = consistency
	}
}

// ContextOptionCredentials sets user name and password for this context.
func ContextOptionCredentials(username, password string) ContextOption {
	return func(c Context) {
		c.(*mcontext).username = username
		c.(*mcontext).password = password
	}
}

// ContextOptionToken sets the token for this request.
func ContextOptionToken(token string) ContextOption {
	return func(c Context) {
		c.(*mcontext).username = "Bearer"
		c.(*mcontext).password = token
	}
}

// ContextOptionClientIP sets the optional headers for the request.
func ContextOptionClientIP(clientIP string) ContextOption {
	return func(c Context) {
		c.(*mcontext).clientIP = clientIP
	}
}

// ContextOptionRetryFunc sets the retry function.
// This function will be called on every communication error, and will be passed
// the try number and the error. If it itself return an error, retrying will stop and
// that error will be returned from the manipulator operation.
func ContextOptionRetryFunc(f RetryFunc) ContextOption {
	return func(c Context) {
		c.(*mcontext).retryFunc = f
	}
}

// ContextOptionRetryRatio sets the retry ratio.
//
// RetryRatio divides the remaining time unitl context
// deadline to perfrom single retry query down to a minimum
// defined by manipulator implementations (typically 20s).
// The default value is 4.
//
// For example if the context has a timeout of 2m,
// each retry will use a sub context with a timeout
// of 30s.
func ContextOptionRetryRatio(r int64) ContextOption {
	return func(c Context) {
		c.(*mcontext).retryRatio = r
	}
}

// ContextOptionIdempotencyKey sets a custom idempotency key.
func ContextOptionIdempotencyKey(key string) ContextOption {
	return func(c Context) {
		c.(*mcontext).idempotencyKey = key
	}
}

// ContextOptionOpaque sets a opaque data. Their interpretation
// depends on the manipulator implementation.
func ContextOptionOpaque(o map[string]interface{}) ContextOption {
	return func(c Context) {
		c.(*mcontext).opaque = o
	}
}
