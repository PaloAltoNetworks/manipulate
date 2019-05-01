// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
