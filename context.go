// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import (
	"fmt"
	"reflect"
)

// Contexts is a type, this type will be checked and can be only a Context or []*Context
type Contexts interface{}

// Context is a structure
type Context struct {
	PageCurrent   int
	PageSize      int
	PageFirst     string
	PageNext      string
	PagePrev      string
	PageLast      string
	CountLocal    int
	CountTotal    int
	Filter        *Filter
	Parameters    *Parameters
	Attributes    []string
	TransactionID TransactionID
}

// NewContext returns a new *Context
func NewContext() *Context {

	return &Context{
		PageCurrent: 1,
		PageSize:    100,
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

// String returns the string representation of the Context.
func (c *Context) String() string {

	return fmt.Sprintf("<Context page: %d, pagesize: %d> <Filter : %v>", c.PageCurrent, c.PageSize, c.Filter)
}

// ContextForIndex return the Context from the given Contexts
// If you give an array of *Context, the function will return the context of the given index
// Otherwise it will only cast the given Contexts to a *Context
// This method will crash if the given Contexts is not a *Context or a []*Context
func ContextForIndex(c Contexts, index int) *Context {

	if c == nil {
		return NewContext()
	}

	castContextFunction := func(i interface{}) *Context {
		return i.(*Context)
	}

	if reflect.TypeOf(c).Kind() == reflect.Slice || reflect.TypeOf(c).Kind() == reflect.Array {
		return castContextFunction(c.([]*Context)[index])
	}

	return castContextFunction(c)
}
