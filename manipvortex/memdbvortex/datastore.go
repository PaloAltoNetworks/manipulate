package memdbvortex

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	memdb "github.com/hashicorp/go-memdb"
	"go.aporeto.io/manipulate/manipvortex/config"
)

// MemdbDatastore is the datastore of the vortex. It must be initialized
// first and then provided to the vortext.
type MemdbDatastore struct {
	db     *memdb.MemDB
	schema *memdb.DBSchema
	sync.Mutex
}

// NewDatastore creates a new datastore backed by a memdb.
func NewDatastore(c map[string]*config.MemDBIdentity) (*MemdbDatastore, error) {

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{},
	}

	for table, cfg := range c {
		index, err := createSchema(cfg)
		if err != nil {
			return nil, err
		}
		schema.Tables[table] = index
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	return &MemdbDatastore{
		schema: schema,
		db:     db,
	}, nil
}

// Flush will flush the datastore essentially creating a new one.
func (d *MemdbDatastore) Flush() error {
	d.Lock()
	defer d.Unlock()

	db, err := memdb.NewMemDB(d.schema)
	if err != nil {
		return err
	}

	d.db = db

	return nil
}

// GetDB returns the actual memdb.
func (d *MemdbDatastore) GetDB() *memdb.MemDB {
	d.Lock()
	defer d.Unlock()

	return d.db
}

// createSchema creates the memdb schema from the configuration of the identities.
func createSchema(c *config.MemDBIdentity) (*memdb.TableSchema, error) {

	tableSchema := &memdb.TableSchema{
		Name:    c.Identity.Category,
		Indexes: map[string]*memdb.IndexSchema{},
	}

	for _, index := range c.Indexes {

		var indexConfig memdb.Indexer

		switch index.Type {

		case config.Slice:
			indexConfig = &memdb.StringSliceFieldIndex{Field: index.Attribute}

		case config.Map:
			indexConfig = &memdb.StringMapFieldIndex{Field: index.Attribute}

		case config.String:
			indexConfig = &memdb.StringFieldIndex{Field: index.Attribute}

		case config.Boolean:
			attr := index.Attribute
			indexConfig = &memdb.ConditionalIndex{Conditional: func(obj interface{}) (bool, error) {
				return boolIndex(obj, attr)
			}}

		case config.StringBased:
			indexConfig = &StringBasedFieldIndex{Field: index.Attribute}

		default: // if the caller is a bozo
			return nil, fmt.Errorf("invalid index type: %d", index.Type)
		}

		tableSchema.Indexes[index.Name] = &memdb.IndexSchema{
			Name:    index.Name,
			Unique:  index.Unique,
			Indexer: indexConfig,
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

// StringBasedFieldIndex is used to extract a field from an object
// using reflection and builds an index on that field. The Indexer
// takes objects that the underlying is string, even though the original
// type is not string. For example, if you declare a type as
//     type ABC string
// then you should use this indexer. It implements the memdb indexer
// interface.
type StringBasedFieldIndex struct {
	Field     string
	Lowercase bool
}

// FromObject implements the memdb indexer interface.
func (s *StringBasedFieldIndex) FromObject(obj interface{}) (bool, []byte, error) {
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
func (s *StringBasedFieldIndex) FromArgs(args ...interface{}) ([]byte, error) {
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
