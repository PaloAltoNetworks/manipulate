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
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/objectid"
)

const (
	descendingOrderPrefix       = "-"
	errInvalidQueryInvalidRegex = "regular expression is invalid"
	errInvalidQueryBadRegex     = "$regex has to be a string"
)

func applyOrdering(order []string, spec elemental.AttributeSpecifiable) []string {

	o := []string{} // nolint: prealloc

	for _, f := range order {

		if f == "" {
			continue
		}

		if spec != nil {
			trimmed := strings.TrimPrefix(f, descendingOrderPrefix)
			if attrSpec := spec.SpecificationForAttribute(trimmed); attrSpec.BSONFieldName != "" {
				original := f
				f = attrSpec.BSONFieldName
				// if we stripped the "-" from the field name, we add it back to the BSON representation of the field name.
				if trimmed != original {
					f = fmt.Sprintf("%s%s", descendingOrderPrefix, f)
				}
			}
		} else {
			if f == "ID" || f == "id" {
				f = "_id"
			}

			if f == "-ID" || f == "-id" {
				f = "-_id"
			}
		}

		o = append(o, strings.ToLower(f))
	}

	return o
}

func prepareNextFilter(collection *mgo.Collection, orderingField string, next string) (bson.D, error) {

	var id any
	if oid, ok := objectid.Parse(next); ok {
		id = oid
	} else {
		id = next
	}

	if orderingField == "" {
		return bson.D{
			{
				Name: "_id",
				Value: bson.D{
					{
						Name:  "$gt",
						Value: id,
					},
				},
			},
		}, nil
	}

	comp := "$gt"
	if strings.HasPrefix(orderingField, "-") {
		orderingField = strings.TrimPrefix(orderingField, "-")
		comp = "$lt"
	}

	doc := bson.M{}
	if err := collection.FindId(id).Select(bson.M{orderingField: 1}).One(&doc); err != nil {
		return nil, HandleQueryError(err)
	}

	return bson.D{
		{
			Name: orderingField,
			Value: bson.D{
				{
					Name:  comp,
					Value: doc[orderingField],
				},
			},
		},
	}, nil
}

// HandleQueryError handles the provided upstream error returned by Mongo by returning a corresponding manipulate error type.
func HandleQueryError(err error) error {

	if _, ok := err.(net.Error); ok {
		return manipulate.ErrCannotCommunicate{Err: err}
	}

	if err == mgo.ErrNotFound {
		return manipulate.ErrObjectNotFound{Err: fmt.Errorf("cannot find the object for the given ID")}
	}

	if mgo.IsDup(err) {
		return manipulate.ErrConstraintViolation{Err: fmt.Errorf("duplicate key")}
	}

	if isConnectionError(err) {
		return manipulate.ErrCannotCommunicate{Err: err}
	}

	if ok, err := invalidQuery(err); ok {
		return err
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
		return manipulate.ErrCannotCommunicate{Err: err}
	default:
		return manipulate.ErrCannotExecuteQuery{Err: err}
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

func invalidQuery(err error) (bool, error) {

	qErr, ok := queryError(err)
	if !ok {
		return false, nil
	}

	errCopyLower := strings.ToLower(qErr.Message)
	switch {
	case qErr.Code == 2 && strings.Contains(errCopyLower, errInvalidQueryBadRegex):
		return true, manipulate.ErrInvalidQuery{
			DueToFilter: true,
			Err:         qErr,
		}
	case qErr.Code == 51091 && strings.Contains(errCopyLower, errInvalidQueryInvalidRegex):
		return true, manipulate.ErrInvalidQuery{
			DueToFilter: true,
			Err:         qErr,
		}
	default:
		return false, nil
	}
}

func queryError(err error) (*mgo.QueryError, bool) {

	if err == nil {
		return nil, false
	}

	switch e := err.(type) {
	case *mgo.QueryError:
		return e, true
	case *mgo.BulkError:
		for _, c := range e.Cases() {
			return queryError(c.Err)
		}
	}

	return nil, false
}

func isConnectionError(err error) bool {

	if err == nil {
		return false
	}

	// Stolen from mongodb code. this is ugly.
	const (
		errLostConnection               = "lost connection to server"
		errNoReachableServers           = "no reachable servers"
		errReplTimeoutPrefix            = "waiting for replication timed out"
		errCouldNotContactPrimaryPrefix = "could not contact primary for replica set"
		errWriteResultsUnavailable      = "write results unavailable from"
		errCouldNotFindPrimaryPrefix    = `could not find host matching read preference { mode: "primary"`
		errUnableToTargetPrefix         = "unable to target"
		errNotMaster                    = "not master"
		errConnectionRefusedSuffix      = "connection refused"
	)

	lowerCaseError := strings.ToLower(err.Error())
	if lowerCaseError == errNoReachableServers ||
		err == io.EOF ||
		strings.Contains(lowerCaseError, errLostConnection) ||
		strings.Contains(lowerCaseError, errReplTimeoutPrefix) ||
		strings.Contains(lowerCaseError, errCouldNotContactPrimaryPrefix) ||
		strings.Contains(lowerCaseError, errWriteResultsUnavailable) ||
		strings.Contains(lowerCaseError, errCouldNotFindPrimaryPrefix) ||
		strings.Contains(lowerCaseError, errUnableToTargetPrefix) ||
		lowerCaseError == errNotMaster ||
		strings.HasSuffix(lowerCaseError, errConnectionRefusedSuffix) {
		return true
	}
	return false
}

func makeFieldsSelector(fields []string, spec elemental.AttributeSpecifiable) bson.M {

	if len(fields) == 0 {
		return nil
	}

	sels := bson.M{}
	for _, f := range fields {

		if f == "" {
			continue
		}

		f = strings.ToLower(strings.TrimPrefix(f, descendingOrderPrefix))
		if spec != nil {
			// if a spec has been provided, use it to look up the BSON field name if there is an entry for the attribute.
			// if no entry was found for the attribute in the provided spec default to whatever value was provided for
			// the attribute.
			if as := spec.SpecificationForAttribute(f); as.BSONFieldName != "" {
				f = as.BSONFieldName
			}
		} else {
			if f == "id" {
				f = "_id"
			}
		}

		sels[f] = 1
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
	case manipulate.ReadConsistencyWeakest:
		return mgo.SecondaryPreferred
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

func explainIfNeeded(
	query *mgo.Query,
	filter bson.D,
	identity elemental.Identity,
	operation elemental.Operation,
	explainMap map[elemental.Identity]map[elemental.Operation]struct{},
) func() error {

	if len(explainMap) == 0 {
		return nil
	}

	exp, ok := explainMap[identity]
	if !ok {
		return nil
	}

	if len(exp) == 0 {
		return func() error { return explain(query, operation, identity, filter) }
	}

	if _, ok = exp[operation]; ok {
		return func() error { return explain(query, operation, identity, filter) }
	}

	return nil
}

func explain(query *mgo.Query, operation elemental.Operation, identity elemental.Identity, filter bson.D) error {

	r := bson.M{}
	if err := query.Explain(&r); err != nil {
		return fmt.Errorf("unable to explain: %s", err)
	}

	f := "<none>"
	if filter != nil {
		fdata, err := json.MarshalIndent(filter, "", "  ")
		if err != nil {
			return fmt.Errorf("unable to marshal filter: %s", err)
		}
		f = string(fdata)
	}

	rdata, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal explanation: %s", err)
	}

	fmt.Println("")
	fmt.Println("--------------------------------")
	fmt.Printf("Operation:  %s\n", operation)
	fmt.Printf("Identity:   %s\n", identity.Name)
	fmt.Printf("Filter:     %s\n", f)
	fmt.Println("Explanation:")
	fmt.Println(string(rdata))
	fmt.Println("--------------------------------")
	fmt.Println("")

	return nil
}
