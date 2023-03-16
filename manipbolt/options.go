package manipbolt

import "github.com/asdine/storm/codec"

// An Option represents a manipbolt.Manipulator option.
type Option func(*config)

type config struct {
	codec codec.MarshalUnmarshaler
}

func newConfig() *config {
	return &config{}
}

// OptionCodec used to set a custom encoder and decoder.
// The default is storm codec which is JSON.
func OptionCodec(codec codec.MarshalUnmarshaler) Option {
	return func(c *config) {
		c.codec = codec
	}
}
