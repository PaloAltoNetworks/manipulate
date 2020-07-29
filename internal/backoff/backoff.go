package backoff

import (
	"time"
)

// NextWithCurve computes the next backoff time for a given try number,
// optional (non zero) hard deadline using the given backoffs curve.
func NextWithCurve(try int, deadline time.Time, curve []time.Duration) time.Duration {

	if len(curve) == 0 {
		return 0
	}

	var wait time.Duration
	if try >= len(curve) {
		wait = curve[len(curve)-1]
	} else {
		wait = curve[try]
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
