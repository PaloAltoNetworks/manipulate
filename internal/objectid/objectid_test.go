package objectid

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo/bson"
)

func TestParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  bson.ObjectId
		want1 bool
	}{
		{
			"valid",
			args{
				"5d66b8f7919e0c446f0b4597",
			},
			bson.ObjectIdHex("5d66b8f7919e0c446f0b4597"),
			true,
		},
		{
			"not exa",
			args{
				"ZZZ6b8f7919e0c446f0b4597",
			},
			bson.ObjectId(""),
			false,
		},
		{
			"too short",
			args{
				"5d66b8f7919e0c446f0b459",
			},
			bson.ObjectId(""),
			false,
		},
		{
			"empty",
			args{
				"",
			},
			bson.ObjectId(""),
			false,
		},
		{
			"weird stuff",
			args{
				"hello world how are you",
			},
			bson.ObjectId(""),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Parse(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
