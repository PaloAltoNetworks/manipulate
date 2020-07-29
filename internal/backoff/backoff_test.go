package backoff

import (
	"testing"
	"time"
)

func TestNext(t *testing.T) {
	type args struct {
		try      int
		deadline time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			"basic with no deadline, try 0",
			args{
				0,
				time.Time{},
			},
			0 * time.Second,
		},
		{
			"basic with no deadline, try 1",
			args{
				1,
				time.Time{},
			},
			1 * time.Second,
		},
		{
			"basic with no deadline, try 2",
			args{
				2,
				time.Time{},
			},
			5 * time.Second,
		},
		{
			"basic with no deadline, try 3",
			args{
				3,
				time.Time{},
			},
			10 * time.Second,
		},
		{
			"basic with no deadline, try 4",
			args{
				4,
				time.Time{},
			},
			20 * time.Second,
		},
		{
			"basic with no deadline, try 5",
			args{
				5,
				time.Time{},
			},
			30 * time.Second,
		},
		{
			"basic with no deadline, try 10",
			args{
				10,
				time.Time{},
			},
			60 * time.Second,
		},
		{
			"basic with no deadline, try 11",
			args{
				11,
				time.Time{},
			},
			60 * time.Second,
		},
		{
			"basic with no deadline, try 12",
			args{
				12,
				time.Time{},
			},
			60 * time.Second,
		},
		{
			"basic with no deadline, try 13",
			args{
				13,
				time.Time{},
			},
			60 * time.Second,
		},

		{
			"deadline in 1s with, try 0",
			args{
				0,
				time.Now().Add(1 * time.Second).Round(time.Second),
			},
			0 * time.Second,
		},
		{
			"deadline in 4s with, try 1",
			args{
				1,
				time.Now().Add(4 * time.Second).Round(time.Second),
			},
			1 * time.Second,
		},
		{
			"deadline in 1s with, try 2",
			args{
				2,
				time.Now().Add(1 * time.Second).Round(time.Second),
			},
			1 * time.Second,
		},
		{
			"deadline in 19s with, try 3",
			args{
				3,
				time.Now().Add(19 * time.Second).Round(time.Second),
			},
			10 * time.Second,
		},
		{
			"deadline in 60s with, try 4",
			args{
				4,
				time.Now().Add(60 * time.Second).Round(time.Second),
			},
			20 * time.Second,
		},
		{
			"deadline in 27s with, try 20",
			args{
				20,
				time.Now().Add(27 * time.Second).Round(time.Second),
			},
			27 * time.Second,
		},
		{
			"deadline in 2700s with, try 20",
			args{
				20,
				time.Now().Add(2700 * time.Second).Round(time.Second),
			},
			60 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Next(tt.args.try, tt.args.deadline); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
