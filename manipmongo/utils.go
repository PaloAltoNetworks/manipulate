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

package manipmongo

import (
	"strings"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"go.aporeto.io/manipulate"
)

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

	o := []string{} // nolint: prealloc

	for _, f := range order {

		if f == "" {
			continue
		}

		if f == "ID" || f == "id" {
			f = "_id"
		}

		o = append(o, strings.ToLower(invertSortKey(f, inverted)))
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

func makeFieldsSelector(fields []string) bson.M {

	if len(fields) == 0 {
		return nil
	}

	sels := bson.M{}
	for _, f := range fields {
		if f == "" {
			continue
		}
		if f == "ID" || f == "id" {
			f = "_id"
		}
		sels[strings.ToLower(f)] = 1
	}

	if len(sels) == 0 {
		return nil
	}

	return sels
}

func convertReadConsistency(c manipulate.ReadConsistency) mgo.Mode {
	switch c {
	case manipulate.ReadConsistencyEventual:
		return mgo.Eventual
	case manipulate.ReadConsistencyMonotonic:
		return mgo.Monotonic
	case manipulate.ReadConsistencyNearest:
		return mgo.Nearest
	case manipulate.ReadConsistencyStrong:
		return mgo.Strong
	default:
		return -1
	}
}

func convertWriteConsistency(c manipulate.WriteConsistency) *mgo.Safe {
	switch c {
	case manipulate.WriteConsistencyNone:
		return nil
	case manipulate.WriteConsistencyStrong:
		return &mgo.Safe{WMode: "majority"}
	case manipulate.WriteConsistencyStrongest:
		return &mgo.Safe{WMode: "majority", J: true}
	default:
		return &mgo.Safe{}
	}
}
