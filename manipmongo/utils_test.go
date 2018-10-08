package manipmongo

import (
	"testing"

	"github.com/globalsign/mgo"
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
