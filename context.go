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
type Context struct {
	Page                 int
	PageSize             int
	Parent               elemental.Identifiable
	CountTotal           int
	Filter               *Filter
	Parameters           url.Values
	TransactionID        TransactionID
	Namespace            string
	Recursive            bool
	OverrideProtection   bool
	CreateFinalizer      FinalizerFunc
	Version              int
	ExternalTrackingID   string
	ExternalTrackingType string
	Order                []string

	ctx context.Context
}

// NewContext returns a new *Context
func NewContext() *Context {

	return &Context{
		Parameters: url.Values{},
	}
}

// NewContextWithFilter returns a new *Context with the given filter.
func NewContextWithFilter(filter *Filter) *Context {

	ctx := NewContext()
	ctx.Filter = filter

	return ctx
}

// NewContextWithTransactionID returns a new *Context with the given transactionID.
func NewContextWithTransactionID(tid TransactionID) *Context {

	ctx := NewContext()
	ctx.TransactionID = tid

	return ctx
}

// WithContext returns a shallow copy of the manipulate.Context
// with it's internal context changed to the given context.Context.
func (c *Context) WithContext(ctx context.Context) *Context {

	if ctx == nil {
		panic("nil context")
	}

	c2 := &Context{}
	*c2 = *c
	c2.ctx = ctx

	return c2
}

// Context returns the internal context.Context of the
// manipulate.Context.
func (c *Context) Context() context.Context {

	if c.ctx != nil {
		return c.ctx
	}

	return context.Background()
}

// String returns the string representation of the Context.
func (c *Context) String() string {

	return fmt.Sprintf("<Context page:%d pagesize:%d filter:%v version:%d>", c.Page, c.PageSize, c.Filter, c.Version)
}
