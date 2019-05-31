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
	"fmt"
	"io"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

func Test_invertSortKey(t *testing.T) {
	type args struct {
		k      string
		revert bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"string",
			args{
				"hello",
				false,
			},
			"hello",
		},
		{
			"string revert",
			args{
				"hello",
				true,
			},
			"-hello",
		},
		{
			"already reverted string",
			args{
				"-hello",
				false,
			},
			"-hello",
		},
		{
			"already reverted string revert",
			args{
				"-hello",
				true,
			},
			"hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertSortKey(tt.args.k, tt.args.revert); got != tt.want {
				t.Errorf("invertSortKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleQueryError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name      string
		args      args
		errString string
	}{
		{
			"net error",
			args{
				&net.OpError{
					Op:  "coucou",
					Err: fmt.Errorf("network sucks"),
				},
			},
			"Cannot communicate: coucou: network sucks",
		},
		{
			"err not found",
			args{
				mgo.ErrNotFound,
			},
			"Object not found: cannot find the object for the given ID",
		},
		{
			"err dup",
			args{
				&mgo.LastError{Code: 11000},
			},
			"Constraint violation: duplicate key.",
		},
		{
			"isConnectionError says yes",
			args{
				fmt.Errorf("lost connection to server"),
			},
			"Cannot communicate: lost connection to server",
		},
		{
			"isConnectionError says no",
			args{
				fmt.Errorf("no"),
			},
			"Unable to execute query: no",
		},

		{
			"err 6",
			args{
				&mgo.LastError{Code: 6, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 7",
			args{
				&mgo.LastError{Code: 7, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 71",
			args{
				&mgo.LastError{Code: 71, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 74",
			args{
				&mgo.LastError{Code: 74, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 91",
			args{
				&mgo.LastError{Code: 91, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 109",
			args{
				&mgo.LastError{Code: 109, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 189",
			args{
				&mgo.LastError{Code: 189, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 202",
			args{
				&mgo.LastError{Code: 202, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 216",
			args{
				&mgo.LastError{Code: 216, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 10107",
			args{
				&mgo.LastError{Code: 10107, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 13436",
			args{
				&mgo.LastError{Code: 13436, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 13435",
			args{
				&mgo.LastError{Code: 13435, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 11600",
			args{
				&mgo.LastError{Code: 11600, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 11602",
			args{
				&mgo.LastError{Code: 11602, Err: "boom"},
			},
			"Cannot communicate: boom",
		},
		{
			"err 424242",
			args{
				&mgo.LastError{Code: 424242, Err: "boom"},
			},
			"Unable to execute query: boom",
		},

		{
			"err 11602 QueryError ",
			args{
				&mgo.QueryError{Code: 424242, Message: "boom"},
			},
			"Unable to execute query: boom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := handleQueryError(tt.args.err)
			if tt.errString != err.Error() {
				t.Errorf("handleQueryError() error = %v, wantErr %v", err, tt.errString)
			}
		})
	}
}

func Test_makeFieldsSelector(t *testing.T) {
	type args struct {
		fields []string
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
			},
			bson.M{
				"_id": 1,
			},
		},
		{
			"id",
			args{
				[]string{"ID"},
			},
			bson.M{
				"_id": 1,
			},
		},
		{
			"empty",
			args{
				[]string{},
			},
			nil,
		},
		{
			"nil",
			args{
				nil,
			},
			nil,
		},
		{
			"only empty",
			args{
				[]string{"", ""},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeFieldsSelector(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeFieldsSelector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_applyOrdering(t *testing.T) {
	type args struct {
		order    []string
		inverted bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"simple",
			args{
				[]string{"NAME", "toto", ""},
				false,
			},
			[]string{"name", "toto"},
		},
		{
			"simple inverted",
			args{
				[]string{"NAME", "", "toto"},
				true,
			},
			[]string{"-name", "-toto"},
		},
		{
			"ID",
			args{
				[]string{"ID"},
				false,
			},
			[]string{"_id"},
		},
		{
			"ID inverted",
			args{
				[]string{"ID"},
				true,
			},
			[]string{"-_id"},
		},
		{
			"id",
			args{
				[]string{"id"},
				false,
			},
			[]string{"_id"},
		},
		{
			"id inverted",
			args{
				[]string{"id"},
				true,
			},
			[]string{"-_id"},
		},
		{
			"only empty",
			args{
				[]string{"", ""},
				false,
			},
			[]string{},
		},
		{
			"only empty inverted",
			args{
				[]string{"", ""},
				true,
			},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyOrdering(tt.args.order, tt.args.inverted); !reflect.DeepEqual(got, tt.want) {
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
		want mgo.Mode
	}{
		{
			"eventual",
			args{manipulate.ReadConsistencyEventual},
			mgo.Eventual,
		},
		{
			"monotonic",
			args{manipulate.ReadConsistencyMonotonic},
			mgo.Monotonic,
		},
		{
			"nearest",
			args{manipulate.ReadConsistencyNearest},
			mgo.Nearest,
		},
		{
			"strong",
			args{manipulate.ReadConsistencyStrong},
			mgo.Strong,
		},
		{
			"default",
			args{manipulate.ReadConsistencyDefault},
			-1,
		},
		{
			"something else",
			args{manipulate.ReadConsistency("else")},
			-1,
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
	tests := []struct {
		name string
		args args
		want *mgo.Safe
	}{
		{
			"none",
			args{manipulate.WriteConsistencyNone},
			nil,
		},
		{
			"strong",
			args{manipulate.WriteConsistencyStrong},
			&mgo.Safe{WMode: "majority"},
		},
		{
			"strongest",
			args{manipulate.WriteConsistencyStrongest},
			&mgo.Safe{WMode: "majority", J: true},
		},
		{
			"default",
			args{manipulate.WriteConsistencyDefault},
			&mgo.Safe{},
		},
		{
			"something else",
			args{manipulate.WriteConsistency("else")},
			&mgo.Safe{},
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

func TestRunQueryFunc(t *testing.T) {

	testIdentity := elemental.MakeIdentity("test", "tests")

	Convey("Given I have query function that works", t, func() {

		var try int
		var lastErr error
		var imctx *manipulate.Context

		f := func() (interface{}, error) { return "hello", nil }
		rf := func(i manipulate.RetryInfo) error {
			try = i.Try()
			lastErr = i.Err()
			m := i.Context()
			if m != nil {
				imctx = &m
			}
			return nil
		}

		Convey("When I call runQueryFunc", func() {

			mctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionRetryFunc(rf),
			)

			out, err := runQueryFunc(
				mctx,
				elemental.OperationCreate,
				testIdentity,
				f,
				nil,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldResemble, "hello")
			})

			Convey("Then try should be correct", func() {
				So(try, ShouldEqual, 0)
			})

			Convey("Then lastErr should be correct", func() {
				So(lastErr, ShouldBeNil)
			})

			Convey("Then imctx should be correct", func() {
				So(imctx, ShouldBeNil)
			})
		})
	})

	Convey("Given I have query function that return an non comm error", t, func() {

		var try int
		var lastErr error
		var imctx *manipulate.Context

		rf := func(i manipulate.RetryInfo) error {
			try = i.Try()
			lastErr = i.Err()
			m := i.Context()
			if m != nil {
				imctx = &m
			}
			return nil
		}

		f := func() (interface{}, error) { return nil, fmt.Errorf("boom") }

		Convey("When I call runQueryFunc", func() {

			out, err := runQueryFunc(
				manipulate.NewContext(
					context.Background(),
					manipulate.ContextOptionRetryFunc(rf),
				),
				elemental.OperationCreate,
				testIdentity,
				f,
				nil,
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Unable to execute query: boom")
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldBeNil)
			})

			Convey("Then try should be correct", func() {
				So(try, ShouldEqual, 0)
			})

			Convey("Then lastErr should be correct", func() {
				So(lastErr, ShouldBeNil)
			})

			Convey("Then imctx should be correct", func() {
				So(imctx, ShouldBeNil)
			})
		})
	})

	Convey("Given I have query function that returns a net.Error and works at second try", t, func() {

		var try int
		var lastErr error
		var operation elemental.Operation
		var identity elemental.Identity
		var imctx *manipulate.Context

		rf := func(i manipulate.RetryInfo) error {
			try = i.Try()
			lastErr = i.Err()
			m := i.Context()
			if m != nil {
				imctx = &m
			}

			operation = i.(RetryInfo).Operation
			identity = i.(RetryInfo).Identity

			return nil
		}

		f := func() (interface{}, error) {
			if try == 2 {
				return "hello", nil
			}
			return nil, &net.OpError{Err: fmt.Errorf("hello")}
		}

		Convey("When I call runQueryFunc", func() {

			mctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionRetryFunc(rf),
			)

			out, err := runQueryFunc(
				mctx,
				elemental.OperationCreate,
				testIdentity,
				f,
				nil,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldResemble, "hello")
			})

			Convey("Then try should be correct", func() {
				So(try, ShouldEqual, 2)
			})

			Convey("Then lastErr should be correct", func() {
				So(lastErr, ShouldNotBeNil)
				So(lastErr.Error(), ShouldEqual, "Cannot communicate: : hello")
			})

			Convey("Then imctx should be correct", func() {
				So(imctx, ShouldNotBeNil)
				So(*imctx, ShouldEqual, mctx)
			})

			Convey("Then operation should be correct", func() {
				So(operation, ShouldEqual, elemental.OperationCreate)
			})

			Convey("Then identity should be correct", func() {
				So(identity.IsEqual(testIdentity), ShouldBeTrue)
			})
		})
	})

	Convey("Given I have query function that returns a net.Error and and a retry func that returns an error", t, func() {

		f := func() (interface{}, error) {
			return nil, &net.OpError{Err: fmt.Errorf("hello")}
		}

		rf := func(i manipulate.RetryInfo) error { return fmt.Errorf("non: %s", i.Err().Error()) }

		Convey("When I call runQueryFunc", func() {

			out, err := runQueryFunc(
				manipulate.NewContext(
					context.Background(),
					manipulate.ContextOptionRetryFunc(rf),
				),
				elemental.OperationCreate,
				testIdentity,
				f,
				nil,
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "non: Cannot communicate: : hello")
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldBeNil)
			})
		})
	})

	Convey("Given I have query function that returns a net.Error and a retry func and a default retry func", t, func() {

		f := func() (interface{}, error) {
			return nil, &net.OpError{Err: fmt.Errorf("hello")}
		}

		rf := func(i manipulate.RetryInfo) error { return fmt.Errorf("non: %s", i.Err().Error()) }
		df := func(i manipulate.RetryInfo) error { return fmt.Errorf("oui: %s", i.Err().Error()) }

		Convey("When I call runQueryFunc", func() {

			out, err := runQueryFunc(
				manipulate.NewContext(
					context.Background(),
					manipulate.ContextOptionRetryFunc(rf),
				),
				elemental.OperationCreate,
				testIdentity,
				f,
				df,
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "non: Cannot communicate: : hello")
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldBeNil)
			})
		})
	})

	Convey("Given I have query function that returns a net.Error and a default retry func", t, func() {

		f := func() (interface{}, error) {
			return nil, &net.OpError{Err: fmt.Errorf("hello")}
		}

		df := func(i manipulate.RetryInfo) error { return fmt.Errorf("oui: %s", i.Err().Error()) }

		Convey("When I call runQueryFunc", func() {

			out, err := runQueryFunc(
				manipulate.NewContext(
					context.Background(),
				),
				elemental.OperationCreate,
				testIdentity,
				f,
				df,
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "oui: Cannot communicate: : hello")
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldBeNil)
			})
		})
	})

	Convey("Given I have query function that returns a net.Error and never works", t, func() {

		f := func() (interface{}, error) {
			return nil, &net.OpError{Err: fmt.Errorf("hello")}
		}

		rf := func(i manipulate.RetryInfo) error { return nil }

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		Convey("When I call runQueryFunc", func() {

			out, err := runQueryFunc(
				manipulate.NewContext(
					ctx,
					manipulate.ContextOptionRetryFunc(rf),
				),
				elemental.OperationCreate,
				testIdentity,
				f,
				nil,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Unable to execute query: context deadline exceeded")
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldBeNil)
			})
		})
	})
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
				fmt.Errorf("hey!"),
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

func Test_getErrorCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"*mgo.QueryError",
			args{
				&mgo.QueryError{Code: 42},
			},
			42,
		},
		{
			"*mgo.LastError",
			args{
				&mgo.LastError{Code: 42},
			},
			42,
		},

		{
			"*mgo.BulkError",
			args{
				&mgo.BulkError{ /* private */ },
			},
			0, // Should be 42. but that is sadly untestable... or is it?
		},
		{
			"",
			args{
				fmt.Errorf("yo!"),
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getErrorCode(tt.args.err); got != tt.want {
				t.Errorf("getErrorCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
