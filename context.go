// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import "fmt"

// A FinalizerFunc is the type of a function that can be used as a creation finalizer.
// This is only supported by manipulators that generate an ID to let a chance to the user.
// to now the intended ID before actually creating the object.
type FinalizerFunc func(o Manipulable) error

// Context is a structure
type Context struct {
	PageCurrent     int
	PageSize        int
	PageFirst       string
	PageNext        string
	PagePrev        string
	PageLast        string
	Parent          Manipulable
	CountLocal      int
	CountTotal      int
	Filter          *Filter
	Parameters      *Parameters
	Attributes      []string
	TransactionID   TransactionID
	Namespace       string
	Recursive       bool
	CreateFinalizer FinalizerFunc
}

// NewContext returns a new *Context
func NewContext() *Context {

	return &Context{
		PageCurrent: 0,
		PageSize:    0,
		Parameters:  NewParameters(),
		Attributes:  []string{},
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
