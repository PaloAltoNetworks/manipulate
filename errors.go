package manipulate

// ErrCannotUnmarshal represents unmarshaling error.
type ErrCannotUnmarshal struct{ message string }

// NewErrCannotUnmarshal returns a new NewErrCannotUnmarshal.
func NewErrCannotUnmarshal(message string) ErrCannotUnmarshal {
	return ErrCannotUnmarshal{message: message}
}

func (e ErrCannotUnmarshal) Error() string { return "Unable to unmarshal data: " + e.message }

// ErrCannotMarshal represents marshaling error.
type ErrCannotMarshal struct{ message string }

// NewErrCannotMarshal returns a new NewErrCannotMarshal.
func NewErrCannotMarshal(message string) ErrCannotMarshal {
	return ErrCannotMarshal{message: message}
}

func (e ErrCannotMarshal) Error() string { return "Unable to marshal data: " + e.message }

// ErrObjectNotFound represents object not found error.
type ErrObjectNotFound struct{ message string }

// NewErrObjectNotFound returns a new NewErrObjectNotFound.
func NewErrObjectNotFound(message string) ErrObjectNotFound {
	return ErrObjectNotFound{message: message}
}

func (e ErrObjectNotFound) Error() string { return "Object not found: " + e.message }

// ErrMultipleObjectsFound represents too many object found error.
type ErrMultipleObjectsFound struct{ message string }

// NewErrMultipleObjectsFound returns a new NewErrMultipleObjectsFound.
func NewErrMultipleObjectsFound(message string) ErrMultipleObjectsFound {
	return ErrMultipleObjectsFound{message: message}
}

func (e ErrMultipleObjectsFound) Error() string { return "Multiple objects found: " + e.message }

// ErrCannotBuildQuery represents query building error.
type ErrCannotBuildQuery struct{ message string }

// NewErrCannotBuildQuery returns a new NewErrCannotBuildQuery.
func NewErrCannotBuildQuery(message string) ErrCannotBuildQuery {
	return ErrCannotBuildQuery{message: message}
}

func (e ErrCannotBuildQuery) Error() string { return "Unable to build query: " + e.message }

// ErrCannotExecuteQuery represents query execution error.
type ErrCannotExecuteQuery struct{ message string }

// NewErrCannotExecuteQuery returns a new NewErrCannotExecuteQuery.
func NewErrCannotExecuteQuery(message string) ErrCannotExecuteQuery {
	return ErrCannotExecuteQuery{message: message}
}

func (e ErrCannotExecuteQuery) Error() string { return "Unable to execute query: " + e.message }

// ErrCannotCommit represents commit execution error.
type ErrCannotCommit struct{ message string }

// NewErrCannotCommit returns a new NewErrCannotCommit.
func NewErrCannotCommit(message string) ErrCannotCommit {
	return ErrCannotCommit{message: message}
}

func (e ErrCannotCommit) Error() string { return "Unable to commit transaction: " + e.message }

// ErrNotImplemented represents a non implemented function.
type ErrNotImplemented struct{ message string }

// NewErrNotImplemented returns a new NewErrNotImplemented.
func NewErrNotImplemented(message string) ErrNotImplemented {
	return ErrNotImplemented{message: message}
}

func (e ErrNotImplemented) Error() string { return "Not implemented: " + e.message }

// ErrCannotCommunicate represents a failure in backend communication.
type ErrCannotCommunicate struct{ message string }

// NewErrCannotCommunicate returns a new NewErrCannotCommunicate.
func NewErrCannotCommunicate(message string) ErrCannotCommunicate {
	return ErrCannotCommunicate{message: message}
}

func (e ErrCannotCommunicate) Error() string { return "Cannot communicate: " + e.message }

// ErrTransactionNotFound represents a failure to find a transaction.
type ErrTransactionNotFound struct{ message string }

// NewErrTransactionNotFound returns a new NewErrTransactionNotFound.
func NewErrTransactionNotFound(message string) ErrTransactionNotFound {
	return ErrTransactionNotFound{message: message}
}

func (e ErrTransactionNotFound) Error() string { return "Transaction not found: " + e.message }
