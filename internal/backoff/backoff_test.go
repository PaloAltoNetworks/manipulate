package backoff

import (
	"testing"
	"time"
)

func TestNext(t *testing.T) {

	type args struct {
		try      int
		deadline time.Time
		curve    []time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			"empty curve no deadlie",
			args{
				0,
				time.Time{},
				nil,
			},
			0 * time.Second,
		},

		{
			"try 0 no deadline",
			args{
				0,
				time.Time{},
				[]time.Duration{
					1 * time.Second,
					2 * time.Second,
					3 * time.Second,
				},
			},
			1 * time.Second,
		},
		{
			"try 1 no deadline",
			args{
				1,
				time.Time{},
				[]time.Duration{
					1 * time.Second,
					2 * time.Second,
					3 * time.Second,
				},
			},
			2 * time.Second,
		},
		{
			"try 2 no deadline",
			args{
				2,
				time.Time{},
				[]time.Duration{
					1 * time.Second,
					2 * time.Second,
					3 * time.Second,
				},
			},
			3 * time.Second,
		},
		{
			"try 2+ no deadline",
			args{
				3,
				time.Time{},
				[]time.Duration{
					1 * time.Second,
					2 * time.Second,
					3 * time.Second,
				},
			},
			3 * time.Second,
		},

		{
			"try before deadline",
			args{
				3,
				time.Now().Add(4 * time.Second),
				[]time.Duration{
					1 * time.Second,
					2 * time.Second,
					3 * time.Second,
				},
			},
			3 * time.Second,
		},
		{
			"try after deadline",
			args{
				3,
				time.Now().Add(5 * time.Second).Round(time.Second),
				[]time.Duration{
					10 * time.Second,
				},
			},
			5 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextWithCurve(tt.args.try, tt.args.deadline, tt.args.curve); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
