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
	"fmt"
	"reflect"
	"strings"

	memdb "github.com/hashicorp/go-memdb"
	"go.aporeto.io/elemental"
)

// stringBasedFieldIndex is used to extract a field from an object
// using reflection and builds an index on that field. The Indexer
// takes objects that the underlying is string, even though the original
// type is not string. For example, if you declare a type as
//
//	type ABC string
//
// then you should use this indexer. It implements the memdb indexer
// interface.
type stringBasedFieldIndex struct {
	Field     string
	Lowercase bool
}

// FromObject implements the memdb indexer interface.
func (s *stringBasedFieldIndex) FromObject(obj interface{}) (bool, []byte, error) {
	v := reflect.ValueOf(obj)
	v = reflect.Indirect(v) // Dereference the pointer if any

	fv := v.FieldByName(s.Field)
	if !fv.IsValid() {
		return false, nil,
			fmt.Errorf("field '%s' for %#v is invalid", s.Field, obj)
	}

	val := fv.String()
	if val == "" {
		return false, nil, nil
	}

	if s.Lowercase {
		val = strings.ToLower(val)
	}

	// Add the null character as a terminator
	val += "\x00"
	return true, []byte(val), nil
}

// FromArgs implements the memdb indexer interface.
func (s *stringBasedFieldIndex) FromArgs(args ...interface{}) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("must provide only a single argument")
	}

	t := reflect.TypeOf(args[0])
	if t.Kind() != reflect.String {
		return nil, fmt.Errorf("argument must be a string: %#v", args[0])
	}
	arg := reflect.ValueOf(args[0]).String()

	if s.Lowercase {
		arg = strings.ToLower(arg)
	}

	// Add the null character as a terminator
	arg += "\x00"
	return []byte(arg), nil
}

// createSchema creates the memdb schema from the configuration of the identities.
func createSchema(c *IdentitySchema) (*memdb.TableSchema, error) {

	tableSchema := &memdb.TableSchema{
		Name:    c.Identity.Category,
		Indexes: map[string]*memdb.IndexSchema{},
	}

	for _, index := range c.Indexes {

		var indexConfig memdb.Indexer

		switch index.Type {

		case IndexTypeSlice:
			indexConfig = &memdb.StringSliceFieldIndex{Field: index.Attribute}

		case IndexTypeMap:
			indexConfig = &memdb.StringMapFieldIndex{Field: index.Attribute}

		case IndexTypeString:
			indexConfig = &memdb.StringFieldIndex{Field: index.Attribute}

		case IndexTypeBoolean:
			attr := index.Attribute
			indexConfig = &memdb.ConditionalIndex{Conditional: func(obj interface{}) (bool, error) {
				return boolIndex(obj, attr)
			}}

		case IndexTypeStringBased:
			indexConfig = &stringBasedFieldIndex{Field: index.Attribute}

		default: // if the caller is a bozo
			return nil, fmt.Errorf("invalid index type: %d", index.Type)
		}

		tableSchema.Indexes[index.Name] = &memdb.IndexSchema{
			Name:         index.Name,
			Unique:       index.Unique,
			Indexer:      indexConfig,
			AllowMissing: true,
		}
	}

	return tableSchema, nil
}

// boolIndex is a conditional indexer for booleans.
func boolIndex(obj interface{}, field string) (bool, error) {

	v := reflect.ValueOf(obj)
	v = reflect.Indirect(v) // Dereference the pointer if any

	fv := v.FieldByName(field)
	if !fv.IsValid() {
		return false, fmt.Errorf("field '%s' for %#v is invalid", field, obj)
	}

	return fv.Bool(), nil
}

func mergeIn(target, source *map[string]elemental.Identifiable) {
	for k, v := range *source {
		(*target)[k] = v
	}
}

func intersection(target, source *map[string]elemental.Identifiable, queryStart bool) {

	combined := map[string]elemental.Identifiable{}

	for k, v := range *source {
		if _, ok := (*target)[k]; ok || queryStart {
			combined[k] = v
		}
	}

	*target = combined
}
