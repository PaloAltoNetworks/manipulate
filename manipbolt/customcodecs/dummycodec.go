package customcodecs

import "github.com/asdine/storm/codec"

var _ codec.MarshalUnmarshaler = &dummyCodec{}

type dummyCodec struct {
}

// NewDummyCodec returns a dummyCodec handle.
func NewDummyCodec() codec.MarshalUnmarshaler {

	return &dummyCodec{}
}

// Marshal returns empty bytes and nil error.
func (d *dummyCodec) Marshal(v interface{}) ([]byte, error) {
	return []byte{}, nil
}

// Unmarshal returns nil error.
func (d *dummyCodec) Unmarshal(b []byte, v interface{}) error {
	return nil
}

// Name of this codec.
func (d *dummyCodec) Name() string {
	return "dummyCodec"
}
