package manipmemory

import (
	"go.aporeto.io/elemental"
)

// IndexType is the data type of the index.
type IndexType int

// Values of IndexType.
const (
	IndexTypeString IndexType = iota
	IndexTypeSlice
	IndexTypeMap
	IndexTypeBoolean
	IndexTypeStringBased
)

// Index configures the attributes that must be indexed.
type Index struct {

	// Name of the index. Must match an attribute of elemental.
	Name string

	// Type of the index.
	Type IndexType

	// If there is a unique requirement on the index. At least
	// one of the indexes must have this set.
	Unique bool

	// Attribute is the elemental attribute name.
	Attribute string
}

// IdentitySchema is the configuration of the indexes for the associated identity.
type IdentitySchema struct {
	// Identity of the object.
	Identity elemental.Identity

	// Indexes of the object
	Indexes []*Index
}
