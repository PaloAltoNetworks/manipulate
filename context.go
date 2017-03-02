// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import (
	"fmt"

	"github.com/aporeto-inc/elemental"
)

// A FinalizerFunc is the type of a function that can be used as a creation finalizer.
// This is only supported by manipulators that generate an ID to let a chance to the user.
// to now the intended ID before actually creating the object.
type FinalizerFunc func(o elemental.Identifiable) error

// Context is a structure
type Context struct {
	Page               int
	PageSize           int
	Parent             elemental.Identifiable
	CountTotal         int
	Filter             *Filter
	Parameters         *Parameters
	Attributes         []string
	TransactionID      TransactionID
	Namespace          string
	Recursive          bool
	OverrideProtection bool
	CreateFinalizer    FinalizerFunc
}

// NewContext returns a new *Context
func NewContext() *Context {

	return &Context{
		Page:       0,
		PageSize:   0,
		Parameters: NewParameters(),
		Attributes: []string{},
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

	return fmt.Sprintf("<Context page: %d, pagesize: %d> <Filter : %v>", c.Page, c.PageSize, c.Filter)
}
