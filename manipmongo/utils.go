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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/objectid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
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

func prepareNextFilter(collection *mongo.Collection, orderingField string, next string) (bson.D, error) {

	var id any
	if oid, ok := objectid.Parse(next); ok {
		id = oid
	} else {
		id = next
	}

	if orderingField == "" {
		return bson.D{
			{
				Key: "_id",
				Value: bson.D{
					{
						Key:   "$gt",
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

	opts := options.FindOne().SetProjection(bson.M{"orderingField": 1})
	if err := collection.FindOne(context.Background(), bson.M{"_id": id}, opts).Decode(&doc); err != nil {
		log.Fatal(err)
	}

	return bson.D{
		{
			Key: orderingField,
			Value: bson.D{
				{
					Key:   comp,
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

	if err == mongo.ErrNoDocuments {
		return manipulate.ErrObjectNotFound{Err: fmt.Errorf("cannot find the object for the given ID")}
	}

	if mongo.IsDuplicateKeyError(err) {
		return manipulate.ErrConstraintViolation{Err: fmt.Errorf("duplicate key")}
	}

	if isConnectionError(err) {
		return manipulate.ErrCannotCommunicate{Err: err}
	}

	if ok, err := invalidQuery(err); ok {
		return err
	}

	if mongo.IsNetworkError(err) {
		return manipulate.ErrCannotCommunicate{Err: err}
	}

	return manipulate.ErrCannotExecuteQuery{Err: err}
}

func invalidQuery(err error) (bool, error) {

	ok, qErr := queryError(err)
	if !ok {
		return false, nil
	}

	return true, manipulate.ErrInvalidQuery{
		DueToFilter: true,
		Err:         qErr,
	}
}

func queryError(err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	switch e := err.(type) {
	case mongo.CommandError:
		return true, e
	case mongo.BulkWriteException:
		for _, writeErr := range e.WriteErrors {
			return queryError(writeErr)
		}
	}
	return false, err
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

func convertReadConsistency(c manipulate.ReadConsistency) *readconcern.ReadConcern {
	switch c {
	case manipulate.ReadConsistencyEventual:
		return readconcern.Available()
	case manipulate.ReadConsistencyMonotonic:
		return readconcern.Majority()
	case manipulate.ReadConsistencyNearest:
		return readconcern.Local()
	case manipulate.ReadConsistencyStrong:
		return readconcern.Majority()
	case manipulate.ReadConsistencyWeakest:
		return readconcern.Available()
	default:
		return &readconcern.ReadConcern{}
	}
}

func convertWriteConsistency(c manipulate.WriteConsistency) *writeconcern.WriteConcern {
	switch c {
	case manipulate.WriteConsistencyNone:
		return writeconcern.Unacknowledged()
	case manipulate.WriteConsistencyStrong:
		return writeconcern.Majority()
	case manipulate.WriteConsistencyStrongest:
		{
			journal := true
			return &writeconcern.WriteConcern{
				W:       "majority",
				Journal: &journal,
			}
		}
	default:
		return &writeconcern.WriteConcern{}
	}
}

func explainIfNeeded(
	collection *mongo.Collection,
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
		return func() error { return explain(collection, operation, identity, filter) }
	}

	if _, ok = exp[operation]; ok {
		return func() error { return explain(collection, operation, identity, filter) }
	}

	return nil
}

func explain(collection *mongo.Collection, operation elemental.Operation, identity elemental.Identity, filter bson.D) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	explainCmd := bson.D{
		{Key: "explain", Value: bson.D{
			{Key: "find", Value: collection.Name()},
			{Key: "filter", Value: filter},
		}},
	}

	var result bson.M
	if err := collection.Database().RunCommand(ctx, explainCmd).Decode(&result); err != nil {
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

	rdata, err := json.MarshalIndent(result, "", "  ")
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

func setMaxTime(ctx context.Context, q interface{}) (interface{}, error) {
	d, ok := ctx.Deadline()
	var mx time.Duration
	if !ok {
		mx = defaultGlobalContextTimeout
	} else {
		mx = time.Until(d)
	}

	if err := ctx.Err(); err != nil {
		return nil, manipulate.ErrCannotBuildQuery{Err: err}
	}

	switch opts := q.(type) {
	case *options.FindOptions:
		return opts.SetMaxTime(mx), nil
	case *options.FindOneOptions:
		return opts.SetMaxTime(mx), nil
	case *options.CountOptions:
		return opts.SetMaxTime(mx), nil
	default:
		return nil, manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("unsupported options type")}
	}
}
