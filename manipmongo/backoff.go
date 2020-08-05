package manipmongo

import "time"

var (
	defaultBackoffCurve = []time.Duration{
		0,
		50 * time.Millisecond,
		100 * time.Millisecond,
		300 * time.Millisecond,
		1 * time.Second,
		5 * time.Second,
		10 * time.Second,
	}
)
