package objectid

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  func() primitive.ObjectID
		want1 bool
	}{
		{
			"valid",
			args{
				"5d66b8f7919e0c446f0b4597",
			},
			func() primitive.ObjectID {
				id, _ := primitive.ObjectIDFromHex("5d66b8f7919e0c446f0b4597")
				return id
			},
			true,
		},
		{
			"not exa",
			args{
				"ZZZ6b8f7919e0c446f0b4597",
			},
			func() primitive.ObjectID {
				return primitive.NilObjectID
			},
			false,
		},
		{
			"too short",
			args{
				"5d66b8f7919e0c446f0b459",
			},
			func() primitive.ObjectID {
				return primitive.NilObjectID
			},
			false,
		},
		{
			"empty",
			args{
				"",
			},
			func() primitive.ObjectID {
				return primitive.NilObjectID
			},
			false,
		},
		{
			"weird stuff",
			args{
				"hello world how are you",
			},
			func() primitive.ObjectID {
				return primitive.NilObjectID
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Parse(tt.args.s)
			if !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want())
			}
			if got1 != tt.want1 {
				t.Errorf("Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
