package manipmongo

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
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
