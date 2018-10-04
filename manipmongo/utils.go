package manipmongo

import (
	"strings"

	"go.aporeto.io/manipulate"

	"github.com/globalsign/mgo"
	"go.aporeto.io/elemental"
)

// collectionFromIdentity returns the mgo*.Collection associated to the given Identity from the
// given *mgo.Database.
func collectionFromIdentity(db *mgo.Database, identity elemental.Identity, prefix string) *mgo.Collection {

	// if prefix != "" {
	// 	return db.C(prefix + "-" + identity.Name)
	// }

	return db.C(identity.Name)
}

// invertSortKey eventually inverts the given sorting key.
func invertSortKey(k string, revert bool) string {

	if !revert {
		return k
	}

	if strings.HasPrefix(k, "-") {
		return k[1:]
	}

	return "-" + k
}

func applyOrdering(order []string, inverted bool) []string {

	o := make([]string, len(order))
	for i := 0; i < len(order); i++ {
		o[i] = strings.ToLower(invertSortKey(order[i], inverted))
	}

	return o
}

func handleQueryError(err error) error {

	if err == mgo.ErrNotFound {
		return manipulate.NewErrObjectNotFound("cannot find the object for the given ID")
	}

	if mgo.IsDup(err) {
		return manipulate.NewErrConstraintViolation("duplicate key.")
	}
	// see https://github.com/mongodb/mongo/blob/master/src/mongo/base/error_codes.err
	switch getErrorCode(err) {
	case 6, 7, 71, 74, 91, 109, 189, 202, 216, 262, 10107, 13436, 13435, 11600, 11602:
		// HostUnreachable
		// HostNotFound,
		// ReplicaSetNotFound,
		// NodeNotFound,
		// ConfigurationInProgress,
		// ShutdownInProgress
		// PrimarySteppedDown,
		// NetworkInterfaceExceededTimeLimit
		// ElectionInProgress
		// ExceededTimeLimit
		// NotMaster
		// NotMasterOrSecondary
		// NotMasterNoSlaveOk
		// InterruptedAtShutdown
		// InterruptedDueToStepDown
		return manipulate.NewErrCannotCommunicate(err.Error())
	default:
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}
}

func getErrorCode(err error) int {

	switch e := err.(type) {

	case *mgo.QueryError:
		return e.Code

	case *mgo.LastError:
		return e.Code

	case *mgo.BulkError:
		// we just get the first
		for _, c := range e.Cases() {
			return getErrorCode(c.Err)
		}
	}

	return 0
}
