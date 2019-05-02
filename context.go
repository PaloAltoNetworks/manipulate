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
	"context"
	"fmt"
	"net/url"

	"go.aporeto.io/elemental"
)

// ReadConsistency represents the desired consistency of the request.
// Not all driver may implement this.
type ReadConsistency string

// Various values for Consistency
const (
	ReadConsistencyDefault   ReadConsistency = "default"
	ReadConsistencyNearest   ReadConsistency = "nearest"
	ReadConsistencyEventual  ReadConsistency = "eventual"
	ReadConsistencyMonotonic ReadConsistency = "monotonic"
	ReadConsistencyStrong    ReadConsistency = "strong"
)

// WriteConsistency represents the desired consistency of the request.
// Not all driver may implement this.
type WriteConsistency string

// Various values for Consistency
const (
	WriteConsistencyDefault   WriteConsistency = "default"
	WriteConsistencyNone      WriteConsistency = "none"
	WriteConsistencyStrong    WriteConsistency = "strong"
	WriteConsistencyStrongest WriteConsistency = "strongest"
)

// A FinalizerFunc is the type of a function that can be used as a creation finalizer.
// This is only supported by manipulators that generate an ID to let a chance to the user.
// to now the intended ID before actually creating the object.
type FinalizerFunc func(o elemental.Identifiable) error

// A Context holds all information regarding a particular manipulate operation.
type Context interface {
	Count() int
	SetCount(count int)
	Filter() *elemental.Filter
	Finalizer() FinalizerFunc
	Version() int
	TransactionID() TransactionID
	Page() int
	PageSize() int
	Override() bool
	Recursive() bool
	Namespace() string
	Parameters() url.Values
	Parent() elemental.Identifiable
	ExternalTrackingID() string
	ExternalTrackingType() string
	Order() []string
	Context() context.Context
	Derive(...ContextOption) Context
	Fields() []string
	ReadConsistency() ReadConsistency
	WriteConsistency() WriteConsistency
	Messages() []string
	SetMessages([]string)

	fmt.Stringer
}

// NewContext creates a context with the given ContextOption.
func NewContext(ctx context.Context, options ...ContextOption) Context {

	if ctx == nil {
		panic("nil context")
	}

	mctx := &mcontext{
		ctx:              ctx,
		writeConsistency: WriteConsistencyDefault,
		readConsistency:  ReadConsistencyDefault,
	}

	for _, opt := range options {
		opt(mctx)
	}

	return mctx
}

type mcontext struct {
	page                 int
	pageSize             int
	parent               elemental.Identifiable
	countTotal           int
	filter               *elemental.Filter
	parameters           url.Values
	transactionID        TransactionID
	namespace            string
	recursive            bool
	overrideProtection   bool
	createFinalizer      FinalizerFunc
	version              int
	externalTrackingID   string
	externalTrackingType string
	order                []string
	ctx                  context.Context
	fields               []string
	writeConsistency     WriteConsistency
	readConsistency      ReadConsistency
	messages             []string
	idempotencyKey       string
}

// Count returns the count
func (c *mcontext) Count() int { return c.countTotal }

// SetCount returns the total count.
func (c *mcontext) SetCount(count int) { c.countTotal = count }

// Filter returns the filter.
func (c *mcontext) Filter() *elemental.Filter { return c.filter }

// Finalizer returns the finalizer.
func (c *mcontext) Finalizer() FinalizerFunc { return c.createFinalizer }

// Version returns the version.
func (c *mcontext) Version() int { return c.version }

// TransactionID returns the transactionID.
func (c *mcontext) TransactionID() TransactionID { return c.transactionID }

// Page returns the page number.
func (c *mcontext) Page() int { return c.page }

// PageSize returns the pageSize.
func (c *mcontext) PageSize() int { return c.pageSize }

// Override returns the override value.
func (c *mcontext) Override() bool { return c.overrideProtection }

// Recursive returns the recursive value.
func (c *mcontext) Recursive() bool { return c.recursive }

// Namespace returns the namespace value.
func (c *mcontext) Namespace() string { return c.namespace }

// Parameters returns the parameters.
func (c *mcontext) Parameters() url.Values { return c.parameters }

// Parent returns the parent.
func (c *mcontext) Parent() elemental.Identifiable { return c.parent }

// ExternalTrackingID returns the ExternalTrackingID.
func (c *mcontext) ExternalTrackingID() string { return c.externalTrackingID }

// ExternalTrackingType returns the ExternalTrackingType.
func (c *mcontext) ExternalTrackingType() string { return c.externalTrackingType }

// Order returns the Order.
func (c *mcontext) Order() []string { return c.order }

// Fields returns the fields.
func (c *mcontext) Fields() []string { return c.fields }

// WriteConsistency returns the desired write consistency.
func (c *mcontext) WriteConsistency() WriteConsistency { return c.writeConsistency }

// ReadConsistency returns the desired read consistency.
func (c *mcontext) ReadConsistency() ReadConsistency { return c.readConsistency }

// Messages returns the eventual list of messages regarding a manipulation.
func (c *mcontext) Messages() []string { return c.messages }

// SetMessages sets the message in the context.
// You should not need to use this.
func (c *mcontext) SetMessages(messages []string) { c.messages = messages }

// Context returns the internal context.Context.
func (c *mcontext) Context() context.Context { return c.ctx }

// IdempotencyKey returns the idempotency key.
func (c *mcontext) IdempotencyKey() string { return c.idempotencyKey }

// SetIdempotencyKey sets the IdempotencyKey. This is used internally
// by manipulator implementation supporting it.
func (c *mcontext) SetIdempotencyKey(k string) { c.idempotencyKey = k }

// String returns the string representation of the Context.
func (c *mcontext) String() string {

	return fmt.Sprintf("<Context page:%d pagesize:%d filter:%v version:%d>", c.page, c.pageSize, c.filter, c.version)
}

// Derive creates a copy of the context but updates the values of the given options.
func (c *mcontext) Derive(options ...ContextOption) Context {

	copy := &mcontext{
		page:                 c.page,
		pageSize:             c.pageSize,
		parent:               c.parent,
		countTotal:           c.countTotal,
		filter:               c.filter,
		parameters:           c.parameters,
		transactionID:        c.transactionID,
		namespace:            c.namespace,
		recursive:            c.recursive,
		overrideProtection:   c.overrideProtection,
		createFinalizer:      c.createFinalizer,
		version:              c.version,
		externalTrackingID:   c.externalTrackingID,
		externalTrackingType: c.externalTrackingType,
		order:                c.order,
		fields:               c.fields,
		ctx:                  c.ctx,
	}

	for _, opt := range options {
		opt(copy)
	}

	return copy
}
