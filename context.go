// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import (
	"context"
	"fmt"
	"net/url"

	"go.aporeto.io/elemental"
)

// A FinalizerFunc is the type of a function that can be used as a creation finalizer.
// This is only supported by manipulators that generate an ID to let a chance to the user.
// to now the intended ID before actually creating the object.
type FinalizerFunc func(o elemental.Identifiable) error

// A Context holds all information regarding a particular manipulate operation.
type Context interface {
	Count() int
	SetCount(count int)
	Filter() *Filter
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

	fmt.Stringer
}

// NewContext creates a context with the given ContextOption.
func NewContext(ctx context.Context, options ...ContextOption) Context {

	if ctx == nil {
		panic("nil context")
	}

	mctx := &mcontext{
		ctx: ctx,
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
	filter               *Filter
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
}

// Count returns the count
func (c *mcontext) Count() int { return c.countTotal }

// SetCount returns the total count.
func (c *mcontext) SetCount(count int) { c.countTotal = count }

// Filter returns the filter.
func (c *mcontext) Filter() *Filter { return c.filter }

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

// Context returns the internal context.Context.
func (c *mcontext) Context() context.Context { return c.ctx }

// String returns the string representation of the Context.
func (c *mcontext) String() string {

	return fmt.Sprintf("<Context page:%d pagesize:%d filter:%v version:%d>", c.page, c.pageSize, c.filter, c.version)
}
