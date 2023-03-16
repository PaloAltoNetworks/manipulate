package manipbolt

import (
	"testing"

	"github.com/asdine/storm/codec/json"
	"github.com/stretchr/testify/require"
)

func TestOptions(t *testing.T) {

	c := newConfig()
	require.NotNil(t, c)
	require.Nil(t, c.codec)

	OptionCodec(json.Codec)(c)
	require.NotNil(t, c.codec)
}
