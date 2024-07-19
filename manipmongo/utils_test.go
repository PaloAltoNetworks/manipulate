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
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/golang/mock/gomock"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/manipmongo/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func Test_makeFieldsSelector(t *testing.T) {
	type args struct {
		fields    []string
		setupSpec func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable
	}
	tests := []struct {
		name string
		args args
		want bson.M
	}{
		{
			"simple",
			args{
				[]string{"MyField1", "myfield2", ""},
				nil,
			},
			bson.M{
				"myfield1": 1,
				"myfield2": 1,
			},
		},
		{
			"ID",
			args{
				[]string{"ID"},
				nil,
			},
			bson.M{
				"_id": 1,
			},
		},
		{
			"id",
			args{
				[]string{"ID"},
				nil,
			},
			bson.M{
				"_id": 1,
			},
		},
		{
			"inverted",
			args{
				[]string{"-something"},
				nil,
			},
			bson.M{
				"something": 1,
			},
		},
		{
			"empty",
			args{
				[]string{},
				nil,
			},
			nil,
		},
		{
			"nil",
			args{
				nil,
				nil,
			},
			nil,
		},
		{
			"only empty",
			args{
				[]string{"", ""},
				nil,
			},
			nil,
		},
		{
			"translate fields from provided spec - entry found",
			args{
				fields: []string{"FieldA"},
				setupSpec: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {

					spec := internal.NewMockAttributeSpecifiable(ctrl)
					spec.
						EXPECT().
						SpecificationForAttribute("fielda").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "a",
							},
						)

					return spec
				},
			},
			bson.M{
				"a": 1,
			},
		},
		{
			"translate fields from provided spec - no entry found - should default to whatever was provided",
			args{
				fields: []string{"FieldA"},
				setupSpec: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {

					spec := internal.NewMockAttributeSpecifiable(ctrl)
					spec.
						EXPECT().
						SpecificationForAttribute("fielda").
						Return(
							elemental.AttributeSpecification{

								// notice how no entry was found for 'fielda' therefore the value in the filter will be used.

								BSONFieldName: "",
							},
						)

					return spec
				},
			},
			bson.M{
				"fielda": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var spec elemental.AttributeSpecifiable
			if tt.args.setupSpec != nil {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				spec = tt.args.setupSpec(t, ctrl)
			}

			if got := makeFieldsSelector(tt.args.fields, spec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeFieldsSelector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_applyOrdering(t *testing.T) {
	type args struct {
		order     []string
		setupSpec func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple",
			args: args{
				order:     []string{"NAME", "toto", ""},
				setupSpec: nil,
			},
			want: []string{"name", "toto"},
		},

		{
			name: "ID",
			args: args{
				order:     []string{"ID"},
				setupSpec: nil,
			},
			want: []string{"_id"},
		},
		{
			name: "-ID",
			args: args{
				order:     []string{"-ID"},
				setupSpec: nil,
			},
			want: []string{"-_id"},
		},

		{
			name: "id",
			args: args{
				order:     []string{"id"},
				setupSpec: nil,
			},
			want: []string{"_id"},
		},
		{
			name: "-id",
			args: args{
				order:     []string{"-id"},
				setupSpec: nil,
			},
			want: []string{"-_id"},
		},

		{
			name: "_id",
			args: args{
				order:     []string{"_id"},
				setupSpec: nil,
			},
			want: []string{"_id"},
		},

		{
			name: "only empty",
			args: args{
				order:     []string{"", ""},
				setupSpec: nil,
			},
			want: []string{},
		},

		{
			name: "only empty",
			args: args{
				order:     []string{"", ""},
				setupSpec: nil,
			},
			want: []string{},
		},

		{
			name: "translate order keys from spec",
			args: args{
				order: []string{
					"FieldA",
					"FieldB",
				},
				setupSpec: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {

					spec := internal.NewMockAttributeSpecifiable(ctrl)
					spec.
						EXPECT().
						SpecificationForAttribute("FieldA").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "a",
							},
						)
					spec.
						EXPECT().
						SpecificationForAttribute("FieldB").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "b",
							},
						)

					return spec
				},
			},
			want: []string{"a", "b"},
		},

		{
			name: "translate order keys from spec w/ order prefix - one field",
			args: args{
				order: []string{
					"-FieldA",
					"FieldB",
				},
				setupSpec: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {

					spec := internal.NewMockAttributeSpecifiable(ctrl)
					spec.
						EXPECT().
						SpecificationForAttribute("FieldA").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "a",
							},
						)
					spec.
						EXPECT().
						SpecificationForAttribute("FieldB").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "b",
							},
						)

					return spec
				},
			},
			want: []string{"-a", "b"},
		},

		{
			name: "translate order keys from spec - ID field - upper case",
			args: args{
				order: []string{
					"ID",
				},
				setupSpec: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {

					spec := internal.NewMockAttributeSpecifiable(ctrl)
					spec.
						EXPECT().
						SpecificationForAttribute("ID").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "_id",
							},
						)
					return spec
				},
			},
			want: []string{"_id"},
		},

		{
			name: "translate order keys from spec - ID field - lower case",
			args: args{
				order: []string{
					"id",
				},
				setupSpec: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {

					spec := internal.NewMockAttributeSpecifiable(ctrl)
					spec.
						EXPECT().
						SpecificationForAttribute("id").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "_id",
							},
						)
					return spec
				},
			},
			want: []string{"_id"},
		},

		{
			name: "translate order keys from spec - ID field - upper case - desc",
			args: args{
				order: []string{
					"-ID",
				},
				setupSpec: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {

					spec := internal.NewMockAttributeSpecifiable(ctrl)
					spec.
						EXPECT().
						SpecificationForAttribute("ID").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "_id",
							},
						)
					return spec
				},
			},
			want: []string{"-_id"},
		},

		{
			name: "translate order keys from spec - ID field - lower case - desc",
			args: args{
				order: []string{
					"-id",
				},
				setupSpec: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {

					spec := internal.NewMockAttributeSpecifiable(ctrl)
					spec.
						EXPECT().
						SpecificationForAttribute("id").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "_id",
							},
						)
					return spec
				},
			},
			want: []string{"-_id"},
		},

		{
			name: "translate order keys from spec w/ order prefix - both fields",
			args: args{
				order: []string{
					"-FieldA",
					"-FieldB",
				},
				setupSpec: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {

					spec := internal.NewMockAttributeSpecifiable(ctrl)
					spec.
						EXPECT().
						SpecificationForAttribute("FieldA").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "a",
							},
						)
					spec.
						EXPECT().
						SpecificationForAttribute("FieldB").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "b",
							},
						)

					return spec
				},
			},
			want: []string{"-a", "-b"},
		},

		{
			name: "translate order keys from spec - default to provided value if nothing found in spec",
			args: args{
				order: []string{
					"YOLO",
				},
				setupSpec: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {

					spec := internal.NewMockAttributeSpecifiable(ctrl)
					spec.
						EXPECT().
						SpecificationForAttribute("YOLO").
						Return(
							elemental.AttributeSpecification{
								BSONFieldName: "",
							},
						)

					return spec
				},
			},
			want: []string{"yolo"},
		},

		{
			name: "only empty",
			args: args{
				order:     []string{"", ""},
				setupSpec: nil,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var spec elemental.AttributeSpecifiable
			if tt.args.setupSpec != nil {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				spec = tt.args.setupSpec(t, ctrl)
			}

			if got := applyOrdering(tt.args.order, spec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOrdering() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertReadConsistency(t *testing.T) {
	type args struct {
		c manipulate.ReadConsistency
	}
	tests := []struct {
		name string
		args args
		want *readconcern.ReadConcern
	}{
		{
			"eventual",
			args{manipulate.ReadConsistencyEventual},
			readconcern.Available(),
		},
		{
			"monotonic",
			args{manipulate.ReadConsistencyMonotonic},
			readconcern.Majority(),
		},
		{
			"nearest",
			args{manipulate.ReadConsistencyNearest},
			readconcern.Local(),
		},
		{
			"strong",
			args{manipulate.ReadConsistencyStrong},
			readconcern.Majority(),
		},
		{
			"weakest",
			args{manipulate.ReadConsistencyWeakest},
			readconcern.Available(),
		},
		{
			"default",
			args{manipulate.ReadConsistencyDefault},
			&readconcern.ReadConcern{},
		},
		{
			"something else",
			args{manipulate.ReadConsistency("else")},
			&readconcern.ReadConcern{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertReadConsistency(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertConsistency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertWriteConsistency(t *testing.T) {
	type args struct {
		c manipulate.WriteConsistency
	}
	trueVal := true
	tests := []struct {
		name string
		args args
		want *writeconcern.WriteConcern
	}{
		{
			"none",
			args{manipulate.WriteConsistencyNone},
			writeconcern.Unacknowledged(),
		},
		{
			"strong",
			args{manipulate.WriteConsistencyStrong},
			writeconcern.Majority(),
		},
		{
			"strongest",
			args{manipulate.WriteConsistencyStrongest},
			&writeconcern.WriteConcern{
				W:       "majority",
				Journal: &trueVal,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertWriteConsistency(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertWriteConsistency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isConnectionError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"nil",
			args{
				nil,
			},
			false,
		},
		{
			"lost connection to server",
			args{
				fmt.Errorf("lost connection to server"),
			},
			true,
		},
		{
			"no reachable servers",
			args{
				fmt.Errorf("no reachable servers"),
			},
			true,
		},
		{
			"waiting for replication timed out",
			args{
				fmt.Errorf("waiting for replication timed out"),
			},
			true,
		},
		{
			"could not contact primary for replica set",
			args{
				fmt.Errorf("could not contact primary for replica set"),
			},
			true,
		},
		{
			"write results unavailable from",
			args{
				fmt.Errorf("write results unavailable from"),
			},
			true,
		},
		{
			`could not find host matching read preference { mode: "primary"`,
			args{
				fmt.Errorf(`could not find host matching read preference { mode: "primary"`),
			},
			true,
		},
		{
			"unable to target",
			args{
				fmt.Errorf("unable to target"),
			},
			true,
		},
		{
			"Connection refused",
			args{
				fmt.Errorf("blah: connection refused"),
			},
			true,
		},
		{
			"EOF",
			args{
				io.EOF,
			},
			true,
		},
		{
			"nope",
			args{
				fmt.Errorf("hey"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isConnectionError(tt.args.err); got != tt.want {
				t.Errorf("isConnectionError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_invalidQuery(t *testing.T) {

	type args struct {
		err error
	}

	testCases := map[string]struct {
		input   args
		wantOk  bool
		wantErr error
	}{
		"code 2 - bad regex": {
			input: args{
				err: &mgo.QueryError{
					Code:    2,
					Message: errInvalidQueryBadRegex,
				},
			},
			wantOk: true,
			wantErr: manipulate.ErrInvalidQuery{
				DueToFilter: true,
				Err: &mgo.QueryError{
					Code:    2,
					Message: errInvalidQueryBadRegex,
				},
			},
		},
		"code 51091 - invalid regex": {
			input: args{
				err: &mgo.QueryError{
					Code:    51091,
					Message: errInvalidQueryInvalidRegex,
				},
			},
			wantOk: true,
			wantErr: manipulate.ErrInvalidQuery{
				DueToFilter: true,
				Err: &mgo.QueryError{
					Code:    51091,
					Message: errInvalidQueryInvalidRegex,
				},
			},
		},
		"nil": {
			input: args{
				err: nil,
			},
			wantOk:  false,
			wantErr: nil,
		},
		"not an invalid query error": {
			input: args{
				err: errors.New("some other error"),
			},
			wantOk:  false,
			wantErr: nil,
		},
	}

	for scenario, tc := range testCases {
		t.Run(scenario, func(t *testing.T) {
			ok, err := invalidQuery(tc.input.err)

			if ok != tc.wantOk {
				t.Errorf("wanted '%t', got '%t'", tc.wantOk, ok)
			}

			if ok && err == nil {
				t.Error("no error was returned when one was expected")
			}

			if !reflect.DeepEqual(err, tc.wantErr) {
				t.Log("Error types did not match")
				t.Errorf("\n"+
					"EXPECTED:\n"+
					"%+v\n"+
					"ACTUAL:\n"+
					"%+v",
					tc.wantErr,
					err,
				)
			}
		})
	}
}

func Test_explainIfNeeded(t *testing.T) {

	identity := elemental.MakeIdentity("thing", "things")

	type args struct {
		collection *mongo.Collection
		filter     bson.D
		identity   elemental.Identity
		operation  elemental.Operation
		explainMap map[elemental.Identity]map[elemental.Operation]struct{}
	}
	tests := []struct {
		name     string
		args     args
		wantFunc bool
	}{
		{
			"empty",
			args{
				nil,
				nil,
				identity,
				elemental.OperationCreate,
				map[elemental.Identity]map[elemental.Operation]struct{}{},
			},
			false,
		},
		{
			"nil",
			args{
				nil,
				nil,
				identity,
				elemental.OperationCreate,
				nil,
			},
			false,
		},
		{
			"matching exactly",
			args{
				nil,
				nil,
				identity,
				elemental.OperationCreate,
				map[elemental.Identity]map[elemental.Operation]struct{}{
					identity: {elemental.OperationCreate: {}},
				},
			},
			true,
		},
		{
			"matching with no operation",
			args{
				nil,
				nil,
				identity,
				elemental.OperationCreate,
				map[elemental.Identity]map[elemental.Operation]struct{}{
					identity: {},
				},
			},
			true,
		},
		{
			"matching with nil operation",
			args{
				nil,
				nil,
				identity,
				elemental.OperationCreate,
				map[elemental.Identity]map[elemental.Operation]struct{}{
					identity: nil,
				},
			},
			true,
		},
		{
			"not matching",
			args{
				nil,
				nil,
				identity,
				elemental.OperationCreate,
				map[elemental.Identity]map[elemental.Operation]struct{}{
					elemental.MakeIdentity("hello", "hellos"): {},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := explainIfNeeded(tt.args.collection, tt.args.filter, tt.args.identity, tt.args.operation, tt.args.explainMap); (got != nil) != tt.wantFunc {
				t.Errorf("explainIfNeeded() = %v, want %v", (got != nil), tt.wantFunc)
			}
		})
	}
}

// func TestSetMaxTime(t *testing.T) {

// 	Convey("Calling setMaxTime with a context with no deadline should work", t, func() {
// 		q := &mgo.Query{}
// 		q, err := setMaxTime(context.Background(), q)
// 		So(err, ShouldBeNil)
// 		qr := (&mgo.Query{}).SetMaxTime(defaultGlobalContextTimeout)
// 		So(q, ShouldResemble, qr)
// 	})

// 	Convey("Calling setMaxTime with a context with valid deadline should work", t, func() {
// 		q := &mgo.Query{}
// 		deadline := time.Now().Add(3 * time.Second)
// 		ctx, cancel := context.WithDeadline(context.Background(), deadline)
// 		defer cancel()

// 		q, err := setMaxTime(ctx, q)
// 		So(err, ShouldBeNil)
// 		qr := (&mgo.Query{}).SetMaxTime(time.Until(deadline))
// 		So(q, ShouldResemble, qr)
// 	})

// 	Convey("Calling setMaxTime with a context with expired deadline should not work", t, func() {
// 		q := &mgo.Query{}
// 		deadline := time.Now().Add(-3 * time.Second)
// 		ctx, cancel := context.WithDeadline(context.Background(), deadline)
// 		defer cancel()

// 		q, err := setMaxTime(ctx, q)
// 		So(err, ShouldNotBeNil)
// 		So(err.Error(), ShouldEqual, "Unable to build query: context deadline exceeded")
// 		So(q, ShouldBeNil)
// 	})

// 	Convey("Calling setMaxTime with a canceled context should not work", t, func() {
// 		q := &mgo.Query{}
// 		deadline := time.Now().Add(3 * time.Second)
// 		ctx, cancel := context.WithDeadline(context.Background(), deadline)
// 		cancel()

// 		q, err := setMaxTime(ctx, q)
// 		So(err, ShouldNotBeNil)
// 		So(err.Error(), ShouldEqual, "Unable to build query: context canceled")
// 		So(q, ShouldBeNil)
// 	})
// }
