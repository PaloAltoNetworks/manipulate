package manipulate

import "github.com/aporeto-inc/elemental"

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

var errorTitles = map[int]string{
	ErrCannotUnmarshal:                    "Unable to unmarshal data.",
	ErrObjectNotFound:                     "Object not found.",
	ErrCannotSlice:                        "Unable to slice objects.",
	ErrCannotCloseIterator:                "Unable to close query iterator.",
	ErrCannotExecuteBatch:                 "Unable to execute batch.",
	ErrCannotScan:                         "Unable to scan query.",
	ErrCannotExecuteQuery:                 "Unable to execute query.",
	ErrCannotExtractFieldsAndValues:       "Unable to extract fields or values.",
	ErrCannotExractPrimaryFieldsAndValues: "Unable to extract primary keys or values.",
	ErrCannotCommit:                       "Unable to commit transaction.",
}

// NewError returns a new manipulation error.
func NewError(err string, code int) error {
	return elemental.NewError(
		errorTitles[code],
		err,
		"manipulate",
		code,
	)
}
