package backoff

import (
	"math"
	"time"
)

const maxBackoff = 8000

func Next(try int, deadline time.Time) time.Duration {

	wait := time.Duration(math.Min(math.Pow(2.0, float64(try))-1, maxBackoff)) * time.Millisecond
	now := time.Now()
	remaining := deadline.Sub(now)

	if wait > remaining {
		return remaining
	}

	return wait
}
