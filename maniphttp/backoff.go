package maniphttp

import "time"

var (
	defaultBackoffCurve = []time.Duration{
		0,
		1 * time.Second,
		5 * time.Second,
		10 * time.Second,
		20 * time.Second,
		30 * time.Second,
		60 * time.Second,
	}

	strongBackoffCurve = []time.Duration{
		2 * time.Second,
		4 * time.Second,
		8 * time.Second,
		16 * time.Second,
		32 * time.Second,
		64 * time.Second,
	}

	testingBackoffCurve = []time.Duration{
		0,
		1 * time.Millisecond,
		10 * time.Millisecond,
	}
)
