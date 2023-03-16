package customcodecs

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/asdine/storm/codec"
	jsoniter "github.com/json-iterator/go"
)

var _ codec.MarshalUnmarshaler = &tagGenerator{}

type tagGenerator struct {
	api jsoniter.API
}

// NewRandomJSONTagGenerator generates a random JSON struct tag to find
// in struct. Its useful in cases when you don't want to consider
// struct tags while encoding/decoding. This can be solved in a
// lot of different ways. This is an attempt at generating a random
// key with 16 length and the json-iter will use this key to apply
// struct tag operations which of course will not exist.
func NewRandomJSONTagGenerator() (codec.MarshalUnmarshaler, error) {

	tagKey, err := generateRandomString(16)
	if err != nil {
		return nil, err
	}

	api := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 tagKey,
	}.Froze()

	return &tagGenerator{
		api: api,
	}, nil
}

// Marshal calls the underlying Marshal.
func (t *tagGenerator) Marshal(v interface{}) ([]byte, error) {
	return t.api.Marshal(v)
}

// Unmarshal calls the underlying Unmarshal.
func (t *tagGenerator) Unmarshal(b []byte, v interface{}) (err error) {
	return t.api.Unmarshal(b, v)
}

// Name of this codec.
func (t *tagGenerator) Name() string {
	return "randomJSONTagGenerator"
}

func generateRandomString(n int) (string, error) {

	b := make([]byte, n)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	s := hex.EncodeToString(b)

	return s[:n], nil
}
