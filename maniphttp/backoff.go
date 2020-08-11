package maniphttp

import (
	"math/rand"
	"time"
)

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
		time.Duration(1500+rand.Intn(1000)) * time.Millisecond,  // t in (1.5, 2.5)
		time.Duration(3000+rand.Intn(1000)) * time.Millisecond,  // t in (3,4)
		time.Duration(7000+rand.Intn(2000)) * time.Millisecond,  // t in (7,9)
		time.Duration(14000+rand.Intn(2000)) * time.Millisecond, // t in (14,16)
		time.Duration(30000+rand.Intn(2000)) * time.Millisecond, // t in (30,32)
		time.Duration(62000+rand.Intn(2000)) * time.Millisecond, // t in (62, 64)
	}

	testingBackoffCurve = []time.Duration{
		0,
		1 * time.Millisecond,
		10 * time.Millisecond,
	}
)
