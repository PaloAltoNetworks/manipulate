package customcodecs

import (
	"crypto/rand"
	"io"
	"os"
	"testing"

	"github.com/asdine/storm/codec/json"
	"github.com/stretchr/testify/require"
	testmodel "go.aporeto.io/elemental/test/model"
)

func writeRandomData(f *os.File, n int64) error {

	_, err := io.CopyN(f, rand.Reader, n)
	return err
}

func Test_FileSizeValidatorJSON(t *testing.T) {

	f, err := os.CreateTemp("", "")
	require.Nil(t, err)
	require.FileExists(t, f.Name())

	defer os.RemoveAll(f.Name()) // nolint: errcheck

	size := int64(65 * 1024)
	c := NewFileSizeValidator(f.Name(), size, json.Codec)
	require.NotNil(t, c)

	l := &testmodel.List{
		Name: "Centos",
	}

	d, err := c.Marshal(l)
	require.Nil(t, err)
	require.NotNil(t, d)

	l1 := &testmodel.List{}
	err = c.Unmarshal(d, l1)
	require.Nil(t, err)
	require.Equal(t, l, l1)
}

func Test_FileSizeValidator(t *testing.T) {

	f, err := os.CreateTemp("", "")
	require.Nil(t, err)
	require.FileExists(t, f.Name())

	defer os.RemoveAll(f.Name()) // nolint: errcheck

	size := int64(32 * 1024)
	c := NewFileSizeValidator(f.Name(), size, NewDummyCodec())
	require.NotNil(t, c)

	name := c.Name()
	require.Equal(t, "fileSizeValidator-dummyCodec", name)

	l := &testmodel.List{
		Name: "Centos",
	}

	err = writeRandomData(f, 200)
	require.Nil(t, err)

	d, err := c.Marshal(l)
	require.Nil(t, err)
	require.NotNil(t, d)

	err = writeRandomData(f, size+12)
	require.Nil(t, err)

	l = &testmodel.List{
		Name: "Centos",
	}

	_, err = c.Marshal(l)
	require.Equal(t, ErrExceedsSize, err)

	info, err := os.Stat(f.Name())
	require.Nil(t, err)
	require.NotNil(t, info)
	require.LessOrEqual(t, info.Size(), int64((32*1024)+12+200))

	c = NewFileSizeValidator("dummyPath", size, json.Codec)
	require.NotNil(t, c)

	l = &testmodel.List{
		Name: "Centos",
	}

	_, err = c.Marshal(l)
	require.NotNil(t, err)
}

func Test_mmapSize(t *testing.T) {

	sz, err := mmapSize(31 * 1024)
	require.Nil(t, err)
	require.Equal(t, int64(32*1024), sz)

	sz, err = mmapSize(33 * 1024)
	require.Nil(t, err)
	require.Equal(t, int64(2*32*1024), sz)

	sz, err = mmapSize(maxMapSize + 1)
	require.NotNil(t, err)
	require.Equal(t, int64(0), sz)

	sz, err = mmapSize(maxMmapStep + 1)
	require.Nil(t, err)
	require.Equal(t, int64(2*maxMmapStep), sz)
}
