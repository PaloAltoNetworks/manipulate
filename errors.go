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

import "fmt"

// ErrInvalidQuery represents an error due to an invalid query.
type ErrInvalidQuery struct {
	// DueToFilter represents whether the query was invalid likely due to the filter supplied by the client.
	DueToFilter bool
	Err         error
}

// Unwrap unwraps the internal error.
func (err ErrInvalidQuery) Unwrap() error {
	return err.Err
}

func (err ErrInvalidQuery) Error() string {
	return fmt.Sprintf("Query invalid: %s", err.Err)
}

// ErrCannotUnmarshal represents unmarshaling error.
type ErrCannotUnmarshal struct{ Err error }

// NewErrCannotUnmarshal returns a new ErrCannotUnmarshal.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrCannotUnmarshal(message string) ErrCannotUnmarshal {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrCannotUnmarshal{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrCannotUnmarshal) Unwrap() error { return e.Err }

func (e ErrCannotUnmarshal) Error() string { return "Unable to unmarshal data: " + e.Err.Error() }

// IsCannotUnmarshalError returns true if the given error is am ErrCannotUnmarshal.
func IsCannotUnmarshalError(err error) bool {
	_, ok := err.(ErrCannotUnmarshal)
	return ok
}

// ErrCannotMarshal represents marshaling error.
type ErrCannotMarshal struct{ Err error }

// NewErrCannotMarshal returns a new ErrCannotMarshal.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrCannotMarshal(message string) ErrCannotMarshal {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrCannotMarshal{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrCannotMarshal) Unwrap() error { return e.Err }

func (e ErrCannotMarshal) Error() string { return "Unable to marshal data: " + e.Err.Error() }

// IsCannotMarshalError returns true if the given error is am ErrCannotMarshal.
func IsCannotMarshalError(err error) bool {
	_, ok := err.(ErrCannotMarshal)
	return ok
}

// ErrObjectNotFound represents object not found error.
type ErrObjectNotFound struct{ Err error }

// NewErrObjectNotFound returns a new ErrObjectNotFound.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrObjectNotFound(message string) ErrObjectNotFound {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrObjectNotFound{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrObjectNotFound) Unwrap() error { return e.Err }

func (e ErrObjectNotFound) Error() string { return "Object not found: " + e.Err.Error() }

// IsObjectNotFoundError returns true if the given error is am ErrObjectNotFound.
func IsObjectNotFoundError(err error) bool {
	_, ok := err.(ErrObjectNotFound)
	return ok
}

// ErrMultipleObjectsFound represents too many object found error.
type ErrMultipleObjectsFound struct{ Err error }

// NewErrMultipleObjectsFound returns a new ErrMultipleObjectsFound.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrMultipleObjectsFound(message string) ErrMultipleObjectsFound {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrMultipleObjectsFound{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrMultipleObjectsFound) Unwrap() error { return e.Err }

func (e ErrMultipleObjectsFound) Error() string { return "Multiple objects found: " + e.Err.Error() }

// IsMultipleObjectsFoundError returns true if the given error is am ErrMultipleObjectsFound.
func IsMultipleObjectsFoundError(err error) bool {
	_, ok := err.(ErrMultipleObjectsFound)
	return ok
}

// ErrCannotBuildQuery represents query building error.
type ErrCannotBuildQuery struct{ Err error }

// NewErrCannotBuildQuery returns a new ErrCannotBuildQuery.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrCannotBuildQuery(message string) ErrCannotBuildQuery {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrCannotBuildQuery{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrCannotBuildQuery) Unwrap() error { return e.Err }

func (e ErrCannotBuildQuery) Error() string { return "Unable to build query: " + e.Err.Error() }

// IsCannotBuildQueryError returns true if the given error is am ErrCannotBuildQuery.
func IsCannotBuildQueryError(err error) bool {
	_, ok := err.(ErrCannotBuildQuery)
	return ok
}

// ErrCannotExecuteQuery represents query execution error.
type ErrCannotExecuteQuery struct{ Err error }

// NewErrCannotExecuteQuery returns a new ErrCannotExecuteQuery.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrCannotExecuteQuery(message string) ErrCannotExecuteQuery {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrCannotExecuteQuery{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrCannotExecuteQuery) Unwrap() error { return e.Err }

func (e ErrCannotExecuteQuery) Error() string { return "Unable to execute query: " + e.Err.Error() }

// IsCannotExecuteQueryError returns true if the given error is am ErrCannotExecuteQuery.
func IsCannotExecuteQueryError(err error) bool {
	_, ok := err.(ErrCannotExecuteQuery)
	return ok
}

// ErrCannotCommit represents commit execution error.
type ErrCannotCommit struct{ Err error }

// NewErrCannotCommit returns a new ErrCannotCommit.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrCannotCommit(message string) ErrCannotCommit {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrCannotCommit{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrCannotCommit) Unwrap() error { return e.Err }

func (e ErrCannotCommit) Error() string { return "Unable to commit transaction: " + e.Err.Error() }

// IsCannotCommitError returns true if the given error is am ErrCannotCommit.
func IsCannotCommitError(err error) bool {
	_, ok := err.(ErrCannotCommit)
	return ok
}

// ErrNotImplemented represents a non implemented function.
type ErrNotImplemented struct{ Err error }

// NewErrNotImplemented returns a new ErrNotImplemented.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrNotImplemented(message string) ErrNotImplemented {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrNotImplemented{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrNotImplemented) Unwrap() error { return e.Err }

func (e ErrNotImplemented) Error() string { return "Not implemented: " + e.Err.Error() }

// IsNotImplementedError returns true if the given error is am ErrNotImplemented.
func IsNotImplementedError(err error) bool {
	_, ok := err.(ErrNotImplemented)
	return ok
}

// ErrCannotCommunicate represents a failure in backend communication.
type ErrCannotCommunicate struct{ Err error }

// NewErrCannotCommunicate returns a new ErrCannotCommunicate.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrCannotCommunicate(message string) ErrCannotCommunicate {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrCannotCommunicate{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrCannotCommunicate) Unwrap() error { return e.Err }

func (e ErrCannotCommunicate) Error() string { return "Cannot communicate: " + e.Err.Error() }

// IsCannotCommunicateError returns true if the given error is am ErrCannotCommunicate.
func IsCannotCommunicateError(err error) bool {
	_, ok := err.(ErrCannotCommunicate)
	return ok
}

// ErrLocked represents the error returned when the server api is locked..
type ErrLocked struct{ Err error }

// NewErrLocked returns a new ErrCannotCommunicate.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrLocked(message string) ErrLocked {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrLocked{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrLocked) Unwrap() error { return e.Err }

func (e ErrLocked) Error() string { return "Cannot communicate: " + e.Err.Error() }

// IsLockedError returns true if the given error is am ErrLocked.
func IsLockedError(err error) bool {
	_, ok := err.(ErrLocked)
	return ok
}

// ErrTransactionNotFound represents a failure to find a transaction.
type ErrTransactionNotFound struct{ Err error }

// NewErrTransactionNotFound returns a new ErrTransactionNotFound.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrTransactionNotFound(message string) ErrTransactionNotFound {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrTransactionNotFound{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrTransactionNotFound) Unwrap() error { return e.Err }

func (e ErrTransactionNotFound) Error() string { return "Transaction not found: " + e.Err.Error() }

// IsTransactionNotFoundError returns true if the given error is am ErrTransactionNotFound.
func IsTransactionNotFoundError(err error) bool {
	_, ok := err.(ErrTransactionNotFound)
	return ok
}

// ErrConstraintViolation represents a failure to find a transaction.
type ErrConstraintViolation struct{ Err error }

// NewErrConstraintViolation returns a new ErrConstraintViolation.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrConstraintViolation(message string) ErrConstraintViolation {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrConstraintViolation{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrConstraintViolation) Unwrap() error { return e.Err }

func (e ErrConstraintViolation) Error() string { return "Constraint violation: " + e.Err.Error() }

// IsConstraintViolationError returns true if the given error is am ErrConstraintViolation.
func IsConstraintViolationError(err error) bool {
	_, ok := err.(ErrConstraintViolation)
	return ok
}

// ErrDisconnected represents an error due user disconnection.
type ErrDisconnected struct{ Err error }

// NewErrDisconnected returns a new ErrDisconnected.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrDisconnected(message string) ErrDisconnected {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrDisconnected{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrDisconnected) Unwrap() error { return e.Err }

func (e ErrDisconnected) Error() string { return "Disconnected: " + e.Err.Error() }

// IsDisconnectedError returns true if the given error is am ErrDisconnected.
func IsDisconnectedError(err error) bool {
	_, ok := err.(ErrDisconnected)
	return ok
}

// ErrTooManyRequests represents the error returned when the server api is locked.
type ErrTooManyRequests struct{ Err error }

// NewErrTooManyRequests returns a new ErrTooManyRequests.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrTooManyRequests(message string) ErrTooManyRequests {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrTooManyRequests{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrTooManyRequests) Unwrap() error { return e.Err }

func (e ErrTooManyRequests) Error() string { return "Too many requests: " + e.Err.Error() }

// IsTooManyRequestsError returns true if the given error is am ErrTooManyRequests.
func IsTooManyRequestsError(err error) bool {
	_, ok := err.(ErrTooManyRequests)
	return ok
}

// ErrTLS represents the error returned when there is a TLS error.
type ErrTLS struct{ Err error }

// NewErrTLS returns a new ErrTLS.
// Deprecated: this method is deprecated and should not be used anymore.
func NewErrTLS(message string) ErrTLS {
	fmt.Println("DEPRECATED: This function is deprecated. Please use simple constructor")
	return ErrTLS{Err: fmt.Errorf("%s", message)}
}

// Unwrap unwraps the internal error.
func (e ErrTLS) Unwrap() error { return e.Err }

func (e ErrTLS) Error() string { return "TLS error: " + e.Err.Error() }

// IsTLSError returns true if the given error is am ErrTLS.
func IsTLSError(err error) bool {
	_, ok := err.(ErrTLS)
	return ok
}
