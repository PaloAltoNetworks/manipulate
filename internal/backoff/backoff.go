package backoff

import (
	"math"
	"time"
)

const maxBackoff = 8000

// Next computes the next backoff time for a give try number
// with a hard given deadline.
func Next(try int, deadline time.Time) time.Duration {

	wait := time.Duration(math.Min(math.Pow(2.0, float64(try))-1, maxBackoff)) * time.Millisecond

	if deadline.IsZero() {
		return wait
	}

	now := time.Now().Round(time.Second)
	if now.Add(wait).After(deadline) && deadline.Sub(now) > 0 {
		return deadline.Sub(now)
	}

	return wait
}
