package customcodecs

import (
	"errors"
	"fmt"
	"os"

	"github.com/asdine/storm/codec"
)

var _ codec.MarshalUnmarshaler = &fileSizeValidator{}

// ErrExceedsSize represents the error when the file size exceeds the given size.
var ErrExceedsSize = errors.New("exceeds max size")

type fileSizeValidator struct {
	path  string
	size  int64
	codec codec.MarshalUnmarshaler
}

// NewFileSizeValidator wraps the given 'codec' with fileSizeValidator and returns the handle.
// Implements codec.MarshalUnmarshaler interface.
func NewFileSizeValidator(path string, size int64, codec codec.MarshalUnmarshaler) codec.MarshalUnmarshaler {

	return &fileSizeValidator{
		path:  path,
		size:  size,
		codec: codec,
	}
}

// Marshal validates if the current file size plus the new object size after encoding
// doesn't exceed the max size. If it exceeds, it returns ErrExceedsSize error.
// If the file is non existent, it returns *os.PathError error. Otherwise, returns
// the encoded data. Note that each call to Marshal will also call os.Stat to get
// the current file size.
func (f *fileSizeValidator) Marshal(v interface{}) ([]byte, error) {

	info, err := os.Stat(f.path)
	if err != nil {
		return nil, err
	}

	data, err := f.codec.Marshal(v)
	if err != nil {
		return nil, err
	}

	sz, err := mmapSize(int64(len(data)) + info.Size())
	if err != nil {
		return nil, err
	}

	if sz > f.size {
		return nil, ErrExceedsSize
	}

	return data, nil
}

// Unmarshal calls the underlying codec Unmarshal.
func (f *fileSizeValidator) Unmarshal(b []byte, v interface{}) error {
	return f.codec.Unmarshal(b, v)
}

// Name of this codec appended with underlying codec name.
func (f *fileSizeValidator) Name() string {
	return fmt.Sprintf("fileSizeValidator-%s", f.codec.Name())
}

/*
The below function including the constants are extracted from
bbolt package https://github.com/etcd-io/bbolt/blob/master/db.go#L391
We use these to estimate the size that bbolt internally grows the
database and use the value to validate with our max size and abort
if it exceeds. How bbolt allocates size is explained in detail here
https://github.com/boltdb/bolt/issues/308#issuecomment-74811638
*/

// maxMapSize represents the largest mmap size supported by Bolt.
const maxMapSize = 0xFFFFFFFFFFFF // 256TB

// The largest step that can be taken when remapping the mmap.
const maxMmapStep = 1 << 30 // 1GB

func mmapSize(sz int64) (int64, error) {
	// Double the size from 32KB until 1GB.
	for i := uint(15); i <= 30; i++ {
		if sz <= 1<<i {
			return 1 << i, nil
		}
	}

	// Verify the requested size is not above the maximum allowed.
	if sz > maxMapSize {
		return 0, fmt.Errorf("mmap too large")
	}

	// If larger than 1GB then grow by 1GB at a time.
	if remainder := sz % int64(maxMmapStep); remainder > 0 {
		sz += int64(maxMmapStep) - remainder
	}

	// Ensure that the mmap size is a multiple of the page size.
	// This should always be true since we're incrementing in MBs.
	pageSize := int64(os.Getpagesize())
	if (sz % pageSize) != 0 {
		sz = ((sz / pageSize) + 1) * pageSize
	}

	// If we've exceeded the max size then only grow up to the max size.
	if sz > maxMapSize {
		sz = maxMapSize
	}

	return sz, nil
}
