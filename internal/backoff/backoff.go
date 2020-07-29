package backoff

import (
	"time"
)

// Mathematically advanced backoff curve :)
var backoffs = []time.Duration{
	0 * time.Second,
	1 * time.Second,
	5 * time.Second,
	10 * time.Second,
	20 * time.Second,
	30 * time.Second,
	60 * time.Second,
}

// Next computes the next backoff time for a give try number
// with a hard given deadline.
func Next(try int, deadline time.Time) time.Duration {

	var wait time.Duration
	if try > len(backoffs) {
		wait = backoffs[len(backoffs)-1]
	} else {
		wait = backoffs[try]
	}

	if deadline.IsZero() {
		return wait
	}

	now := time.Now().Round(time.Second)
	if now.Add(wait).After(deadline) && deadline.Sub(now) > 0 {
		return deadline.Sub(now)
	}

	return wait
}
