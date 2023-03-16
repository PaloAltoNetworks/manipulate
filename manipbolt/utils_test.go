package manipbolt

import (
	"errors"
	"regexp"
	"testing"
)

func Test_containsOrEqualMatcher_MatchField(t *testing.T) {
	type fields struct {
		field string
		value interface{}
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "simple string equal value",
			fields: fields{
				"key",
				"abc",
			},
			args: args{
				"abc",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "simple string unequal value",
			fields: fields{
				"key",
				"abc",
			},
			args: args{
				"nope",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "simple num equal value",
			fields: fields{
				"key",
				23,
			},
			args: args{
				23,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "simple num unequal value",
			fields: fields{
				"key",
				45,
			},
			args: args{
				60,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "string list contains value",
			fields: fields{
				"key",
				"abc",
			},
			args: args{
				[]string{"doit", "abc", "yes"},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "string list not contains value",
			fields: fields{
				"key",
				"nope",
			},
			args: args{
				[]string{"doit", "abc", "yes"},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "num list contains value",
			fields: fields{
				"key",
				24,
			},
			args: args{
				[]int{56, 75, 24},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "num list not contains value",
			fields: fields{
				"key",
				35,
			},
			args: args{
				[]int{56, 75, 24},
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &containsOrEqualMatcher{
				field: tt.fields.field,
				value: tt.fields.value,
			}
			got, err := c.MatchField(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("containsOrEqualMatcher.MatchField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("containsOrEqualMatcher.MatchField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_regexpMatcher_MatchField(t *testing.T) {
	type fields struct {
		r   *regexp.Regexp
		err error
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				r:   nil,
				err: errors.New("failed"),
			},
			args: args{
				"abc",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "valid",
			fields: fields{
				r:   regexp.MustCompile("^ab"),
				err: nil,
			},
			args: args{
				"abcef",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "valid non-match",
			fields: fields{
				r:   regexp.MustCompile("^bca"),
				err: nil,
			},
			args: args{
				"abcef",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &regexpMatcher{
				r:   tt.fields.r,
				err: tt.fields.err,
			}
			got, err := r.MatchField(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("regexpMatcher.MatchField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("regexpMatcher.MatchField() = %v, want %v", got, tt.want)
			}
		})
	}
}
