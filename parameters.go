package manipulate

import "net/url"

// Parameters is a parameter struct which can be used with Cassandra (fields)
// or HTTP (KeyValues)
type Parameters struct {
	KeyValues url.Values
}

// NewParameters returns a new Parameter.
func NewParameters() *Parameters {
	return &Parameters{
		KeyValues: url.Values{},
	}
}
