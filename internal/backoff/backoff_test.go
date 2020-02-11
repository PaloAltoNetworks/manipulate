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
			0,
		},
		{
			"basic with no deadline, try 1",
			args{
				1,
				time.Time{},
			},
			1 * time.Millisecond,
		},
		{
			"basic with no deadline, try 2",
			args{
				2,
				time.Time{},
			},
			3 * time.Millisecond,
		},
		{
			"basic with no deadline, try 3",
			args{
				3,
				time.Time{},
			},
			7 * time.Millisecond,
		},
		{
			"basic with no deadline, try 4",
			args{
				4,
				time.Time{},
			},
			15 * time.Millisecond,
		},
		{
			"basic with no deadline, try 5",
			args{
				5,
				time.Time{},
			},
			31 * time.Millisecond,
		},
		{
			"basic with no deadline, try 10",
			args{
				10,
				time.Time{},
			},
			1*time.Second + 23*time.Millisecond,
		},
		{
			"basic with no deadline, try 11",
			args{
				11,
				time.Time{},
			},
			2*time.Second + 47*time.Millisecond,
		},
		{
			"basic with no deadline, try 12",
			args{
				12,
				time.Time{},
			},
			4*time.Second + 95*time.Millisecond,
		},
		{
			"basic with no deadline, try 13",
			args{
				13,
				time.Time{},
			},
			8 * time.Second,
		},

		{
			"deadline in 1s with, try 0",
			args{
				0,
				time.Now().Add(1 * time.Second).Round(time.Second),
			},
			0,
		},
		{
			"deadline in 1s with, try 1",
			args{
				1,
				time.Now().Add(1 * time.Second).Round(time.Second),
			},
			1 * time.Millisecond,
		},
		{
			"deadline in 1s with, try 2",
			args{
				2,
				time.Now().Add(1 * time.Second).Round(time.Second),
			},
			3 * time.Millisecond,
		},
		{
			"deadline in 1s with, try 3",
			args{
				3,
				time.Now().Add(1 * time.Second).Round(time.Second),
			},
			7 * time.Millisecond,
		},
		{
			"deadline in 1s with, try 4",
			args{
				4,
				time.Now().Add(1 * time.Second).Round(time.Second),
			},
			15 * time.Millisecond,
		},
		{
			"deadline in 1s with, try 5",
			args{
				5,
				time.Now().Add(1 * time.Second).Round(time.Second),
			},
			31 * time.Millisecond,
		},
		{
			"deadline in 1s with, try 10",
			args{
				10,
				time.Now().Add(1 * time.Second).Round(time.Second),
			},
			1 * time.Second,
		},
		{
			"deadline in 1s with, try 100",
			args{
				100,
				time.Now().Add(1 * time.Second).Round(time.Second),
			},
			1 * time.Second,
		},
		{
			"almost expired deadline with, try 1",
			args{
				1,
				time.Now().Add(1000 * time.Nanosecond),
			},
			1 * time.Millisecond,
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
