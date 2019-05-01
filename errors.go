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

// ErrCannotUnmarshal represents unmarshaling error.
type ErrCannotUnmarshal struct{ message string }

// NewErrCannotUnmarshal returns a new ErrCannotUnmarshal.
func NewErrCannotUnmarshal(message string) ErrCannotUnmarshal {
	return ErrCannotUnmarshal{message: message}
}

func (e ErrCannotUnmarshal) Error() string { return "Unable to unmarshal data: " + e.message }

// IsCannotUnmarshalError returns true if the given error is am ErrCannotUnmarshal.
func IsCannotUnmarshalError(err error) bool {
	_, ok := err.(ErrCannotUnmarshal)
	return ok
}

// ErrCannotMarshal represents marshaling error.
type ErrCannotMarshal struct{ message string }

// NewErrCannotMarshal returns a new ErrCannotMarshal.
func NewErrCannotMarshal(message string) ErrCannotMarshal {
	return ErrCannotMarshal{message: message}
}

func (e ErrCannotMarshal) Error() string { return "Unable to marshal data: " + e.message }

// IsCannotMarshalError returns true if the given error is am ErrCannotMarshal.
func IsCannotMarshalError(err error) bool {
	_, ok := err.(ErrCannotMarshal)
	return ok
}

// ErrObjectNotFound represents object not found error.
type ErrObjectNotFound struct{ message string }

// NewErrObjectNotFound returns a new ErrObjectNotFound.
func NewErrObjectNotFound(message string) ErrObjectNotFound {
	return ErrObjectNotFound{message: message}
}

func (e ErrObjectNotFound) Error() string { return "Object not found: " + e.message }

// IsObjectNotFoundError returns true if the given error is am ErrObjectNotFound.
func IsObjectNotFoundError(err error) bool {
	_, ok := err.(ErrObjectNotFound)
	return ok
}

// ErrMultipleObjectsFound represents too many object found error.
type ErrMultipleObjectsFound struct{ message string }

// NewErrMultipleObjectsFound returns a new ErrMultipleObjectsFound.
func NewErrMultipleObjectsFound(message string) ErrMultipleObjectsFound {
	return ErrMultipleObjectsFound{message: message}
}

func (e ErrMultipleObjectsFound) Error() string { return "Multiple objects found: " + e.message }

// IsMultipleObjectsFoundError returns true if the given error is am ErrMultipleObjectsFound.
func IsMultipleObjectsFoundError(err error) bool {
	_, ok := err.(ErrMultipleObjectsFound)
	return ok
}

// ErrCannotBuildQuery represents query building error.
type ErrCannotBuildQuery struct{ message string }

// NewErrCannotBuildQuery returns a new ErrCannotBuildQuery.
func NewErrCannotBuildQuery(message string) ErrCannotBuildQuery {
	return ErrCannotBuildQuery{message: message}
}

func (e ErrCannotBuildQuery) Error() string { return "Unable to build query: " + e.message }

// IsCannotBuildQueryError returns true if the given error is am ErrCannotBuildQuery.
func IsCannotBuildQueryError(err error) bool {
	_, ok := err.(ErrCannotBuildQuery)
	return ok
}

// ErrCannotExecuteQuery represents query execution error.
type ErrCannotExecuteQuery struct{ message string }

// NewErrCannotExecuteQuery returns a new ErrCannotExecuteQuery.
func NewErrCannotExecuteQuery(message string) ErrCannotExecuteQuery {
	return ErrCannotExecuteQuery{message: message}
}

func (e ErrCannotExecuteQuery) Error() string { return "Unable to execute query: " + e.message }

// IsCannotExecuteQueryError returns true if the given error is am ErrCannotExecuteQuery.
func IsCannotExecuteQueryError(err error) bool {
	_, ok := err.(ErrCannotExecuteQuery)
	return ok
}

// ErrCannotCommit represents commit execution error.
type ErrCannotCommit struct{ message string }

// NewErrCannotCommit returns a new ErrCannotCommit.
func NewErrCannotCommit(message string) ErrCannotCommit {
	return ErrCannotCommit{message: message}
}

func (e ErrCannotCommit) Error() string { return "Unable to commit transaction: " + e.message }

// IsCannotCommitError returns true if the given error is am ErrCannotCommit.
func IsCannotCommitError(err error) bool {
	_, ok := err.(ErrCannotCommit)
	return ok
}

// ErrNotImplemented represents a non implemented function.
type ErrNotImplemented struct{ message string }

// NewErrNotImplemented returns a new ErrNotImplemented.
func NewErrNotImplemented(message string) ErrNotImplemented {
	return ErrNotImplemented{message: message}
}

func (e ErrNotImplemented) Error() string { return "Not implemented: " + e.message }

// IsNotImplementedError returns true if the given error is am ErrNotImplemented.
func IsNotImplementedError(err error) bool {
	_, ok := err.(ErrNotImplemented)
	return ok
}

// ErrCannotCommunicate represents a failure in backend communication.
type ErrCannotCommunicate struct{ message string }

// NewErrCannotCommunicate returns a new ErrCannotCommunicate.
func NewErrCannotCommunicate(message string) ErrCannotCommunicate {
	return ErrCannotCommunicate{message: message}
}

func (e ErrCannotCommunicate) Error() string { return "Cannot communicate: " + e.message }

// IsCannotCommunicateError returns true if the given error is am ErrCannotCommunicate.
func IsCannotCommunicateError(err error) bool {
	_, ok := err.(ErrCannotCommunicate)
	return ok
}

// ErrLocked represents the error returned when the server api is locked..
type ErrLocked struct{ message string }

// NewErrLocked returns a new ErrCannotCommunicate.
func NewErrLocked(message string) ErrLocked {
	return ErrLocked{message: message}
}

func (e ErrLocked) Error() string { return "Cannot communicate: " + e.message }

// IsLockedError returns true if the given error is am ErrLocked.
func IsLockedError(err error) bool {
	_, ok := err.(ErrLocked)
	return ok
}

// ErrTransactionNotFound represents a failure to find a transaction.
type ErrTransactionNotFound struct{ message string }

// NewErrTransactionNotFound returns a new ErrTransactionNotFound.
func NewErrTransactionNotFound(message string) ErrTransactionNotFound {
	return ErrTransactionNotFound{message: message}
}

func (e ErrTransactionNotFound) Error() string { return "Transaction not found: " + e.message }

// IsTransactionNotFoundError returns true if the given error is am ErrTransactionNotFound.
func IsTransactionNotFoundError(err error) bool {
	_, ok := err.(ErrTransactionNotFound)
	return ok
}

// ErrConstraintViolation represents a failure to find a transaction.
type ErrConstraintViolation struct{ message string }

// NewErrConstraintViolation returns a new ErrConstraintViolation.
func NewErrConstraintViolation(message string) ErrConstraintViolation {
	return ErrConstraintViolation{message: message}
}

func (e ErrConstraintViolation) Error() string { return "Constraint violation: " + e.message }

// IsConstraintViolationError returns true if the given error is am ErrConstraintViolation.
func IsConstraintViolationError(err error) bool {
	_, ok := err.(ErrConstraintViolation)
	return ok
}

// ErrDisconnected represents an error due user disconnection.
type ErrDisconnected struct{ message string }

// NewErrDisconnected returns a new ErrDisconnected.
func NewErrDisconnected(message string) ErrDisconnected {
	return ErrDisconnected{message: message}
}

func (e ErrDisconnected) Error() string { return "Disconnected: " + e.message }

// IsDisconnectedError returns true if the given error is am ErrDisconnected.
func IsDisconnectedError(err error) bool {
	_, ok := err.(ErrDisconnected)
	return ok
}

// ErrTooManyRequests represents the error returned when the server api is locked.
type ErrTooManyRequests struct{ message string }

// NewErrTooManyRequests returns a new ErrTooManyRequests.
func NewErrTooManyRequests(message string) ErrTooManyRequests {
	return ErrTooManyRequests{message: message}
}

func (e ErrTooManyRequests) Error() string { return "Too many requests: " + e.message }

// IsTooManyRequestsError returns true if the given error is am ErrTooManyRequests.
func IsTooManyRequestsError(err error) bool {
	_, ok := err.(ErrTooManyRequests)
	return ok
}

// ErrTLS represents the error returned when there is a TLS error.
type ErrTLS struct{ message string }

// NewErrTLS returns a new ErrTLS.
func NewErrTLS(message string) ErrTLS {
	return ErrTLS{message: message}
}

func (e ErrTLS) Error() string { return "TLS error: " + e.message }

// IsTLSError returns true if the given error is am ErrTLS.
func IsTLSError(err error) bool {
	_, ok := err.(ErrTLS)
	return ok
}
