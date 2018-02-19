// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/aporeto-inc/elemental"
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
	TrackingSpan         opentracing.Span
	ExternalTrackingID   string
	ExternalTrackingType string
	Order                []string
	ctx                  context.Context
}

// NewContext returns a new *Context
func NewContext(ctx context.Context) *Context {

	return &Context{
		Parameters: url.Values{},
		ctx:        ctx,
	}
}

// String returns the string representation of the Context.
func (c *Context) String() string {

	return fmt.Sprintf("<Context page:%d pagesize:%d filter:%v version:%d>", c.Page, c.PageSize, c.Filter, c.Version)
}

// Done implements the context.Context interface.
func (c *Context) Done() <-chan struct{} { return c.ctx.Done() }

// Err implements the context.Context interface.
func (c *Context) Err() error { return c.ctx.Err() }

// Deadline implements the context.Context interface.
func (c *Context) Deadline() (time.Time, bool) { return c.ctx.Deadline() }

// Value implements the context.Context interface.
func (c *Context) Value(key interface{}) interface{} { return c.ctx.Value(key) }
