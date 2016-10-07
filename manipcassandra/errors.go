// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

const (
	// ManipCassandraDatabaseError represents the an internal db error.
	ManipCassandraDatabaseError = "Database Manipulation Error"
)

const (
	// ErrCannotUnmarshal represents unmarshaling error.
	ErrCannotUnmarshal int = iota + 5000

	// ErrObjectNotFound represents object not found error.
	ErrObjectNotFound

	// ErrCannotSlice represents slicing error.
	ErrCannotSlice

	// ErrCannotCloseIterator represents iterator closing error.
	ErrCannotCloseIterator

	// ErrCannotExecuteBatch represents batch execution error.
	ErrCannotExecuteBatch

	// ErrCannotScan represents scanning error.
	ErrCannotScan

	// ErrCannotExecuteQuery represents query execution error.
	ErrCannotExecuteQuery

	// ErrCannotExtractFieldsAndValues represents field an values extraction error.
	ErrCannotExtractFieldsAndValues

	// ErrCannotExractPrimaryFieldsAndValues represents primary field an values extraction error.
	ErrCannotExractPrimaryFieldsAndValues

	// ErrCannotCommit represents commit execution error.
	ErrCannotCommit
)
