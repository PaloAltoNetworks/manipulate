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
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"go.aporeto.io/manipulate"
)

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
