package memdbvortex

import (
	"reflect"
	"testing"

	memdb "github.com/hashicorp/go-memdb"
	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate/manipvortex/config"
)

func datastoreIndexConfig() map[string]*config.MemDBIdentity {
	return map[string]*config.MemDBIdentity{
		testmodel.ListIdentity.Category: &config.MemDBIdentity{
			Identity: testmodel.ListIdentity,
			Indexes: []*config.IndexConfig{
				&config.IndexConfig{
					Name:      "id",
					Type:      config.String,
					Unique:    true,
					Attribute: "ID",
				},
				&config.IndexConfig{
					Name:      "Name",
					Type:      config.String,
					Unique:    false,
					Attribute: "Name",
				},
				&config.IndexConfig{
					Name:      "Slice",
					Type:      config.Slice,
					Unique:    false,
					Attribute: "Slice",
				},
				&config.IndexConfig{
					Name:      "Map",
					Type:      config.Map,
					Unique:    false,
					Attribute: "Map",
				},
				&config.IndexConfig{
					Name:      "Bool",
					Type:      config.Boolean,
					Unique:    false,
					Attribute: "Bool",
				},
				&config.IndexConfig{
					Name:      "StringBased",
					Type:      config.StringBased,
					Unique:    false,
					Attribute: "StringBased",
				},
			},
		},
	}
}

func Test_NewDatastore(t *testing.T) {
	t.Parallel()

	Convey("When I try to create a new datastore, I should ge the right structure", t, func() {
		d, err := NewDatastore(datastoreIndexConfig())
		So(err, ShouldBeNil)
		So(d, ShouldNotBeNil)
		So(d.schema, ShouldNotBeNil)
		So(d.schema.Tables, ShouldNotBeNil)
		So(d.db, ShouldNotBeNil)

		Convey("And the schema must be correct", func() {
			So(len(d.schema.Tables), ShouldEqual, 1)
			So(d.schema.Tables, ShouldContainKey, testmodel.ListIdentity.Category)
			So(len(d.schema.Tables[testmodel.ListIdentity.Category].Indexes), ShouldEqual, 6)
			So(d.schema.Tables[testmodel.ListIdentity.Category].Indexes["id"],
				ShouldResemble,
				&memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "ID"},
				},
			)
			So(d.schema.Tables[testmodel.ListIdentity.Category].Indexes["Name"],
				ShouldResemble,
				&memdb.IndexSchema{
					Name:    "Name",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Name"},
				},
			)
			So(d.schema.Tables[testmodel.ListIdentity.Category].Indexes["Slice"],
				ShouldResemble,
				&memdb.IndexSchema{
					Name:    "Slice",
					Unique:  false,
					Indexer: &memdb.StringSliceFieldIndex{Field: "Slice"},
				},
			)
			So(d.schema.Tables[testmodel.ListIdentity.Category].Indexes["Map"],
				ShouldResemble,
				&memdb.IndexSchema{
					Name:    "Map",
					Unique:  false,
					Indexer: &memdb.StringMapFieldIndex{Field: "Map"},
				},
			)
			So(d.schema.Tables[testmodel.ListIdentity.Category].Indexes["StringBased"],
				ShouldResemble,
				&memdb.IndexSchema{
					Name:    "StringBased",
					Unique:  false,
					Indexer: &StringBasedFieldIndex{Field: "StringBased"},
				},
			)
			boolIndex := d.schema.Tables[testmodel.ListIdentity.Category].Indexes["Bool"]
			So(boolIndex.Name, ShouldResemble, "Bool")
			So(boolIndex.Unique, ShouldBeFalse)
			So(reflect.TypeOf(boolIndex.Indexer), ShouldEqual, reflect.TypeOf(&memdb.ConditionalIndex{}))

		})
	})
}

func Test_DatastoreFlush(t *testing.T) {
	Convey("Given a valid data store", t, func() {
		d, err := NewDatastore(datastoreIndexConfig())
		So(err, ShouldBeNil)

		Convey("When I flush it, the db should be new", func() {
			oldDb := d.db

			err := d.Flush()
			So(err, ShouldBeNil)

			So(oldDb, ShouldNotResemble, d.db)
		})
	})
}

func Test_boolIndex(t *testing.T) {

	type testObject struct {
		somevalue      bool
		someothervalue string
	}

	Convey("When I call boolindex", t, func() {

		Convey("When I use a good data structure", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			val, err := boolIndex(t, "somevalue")
			So(err, ShouldBeNil)
			So(val, ShouldBeTrue)
		})

		Convey("When I use a good data structure with a bad field", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			_, err := boolIndex(t, "no value")
			So(err, ShouldNotBeNil)
		})

		Convey("When I use nil", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			_, err := boolIndex(t, "no value")
			So(err, ShouldNotBeNil)
		})

		Convey("When I use a bad type field", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			_, err := boolIndex(t, "somestring")
			So(err, ShouldNotBeNil)
		})
	})
}
