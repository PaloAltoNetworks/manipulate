// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import (
	"errors"
	"reflect"
	"testing"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gocql"
	"github.com/aporeto-inc/kennebec/apomock"
	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

// UserIdentity represents the Identity of the object
var TagIdentity = elemental.Identity{
	Name:     "tag",
	Category: "tag",
}

type Tag struct {
	ID          string `cql:"id"`
	Description string `cql:"description"`
	Name        string `cql:"name"`
	Type        int    `cql:"type"`
}

func (t *Tag) Identifier() string {
	return t.ID
}

// Identity returns the Identity of the object.
func (t *Tag) Identity() elemental.Identity {

	return TagIdentity
}

// SetIdentifier sets the value of the object's unique identifier.
func (t *Tag) SetIdentifier(ID string) {
	t.ID = ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (t *Tag) Validate() elemental.Errors {
	return nil
}

func ParseUUID(input string) gocql.UUID {
	var u gocql.UUID
	j := 0
	for _, r := range input {
		switch {
		case r == '-' && j&1 == 0:
			continue
		case r >= '0' && r <= '9' && j < 32:
			u[j/2] |= byte(r-'0') << uint(4-j&1*4)
		case r >= 'a' && r <= 'f' && j < 32:
			u[j/2] |= byte(r-'a'+10) << uint(4-j&1*4)
		case r >= 'A' && r <= 'F' && j < 32:
			u[j/2] |= byte(r-'A'+10) << uint(4-j&1*4)
		default:
			return gocql.UUID{}
		}
		j++
	}

	return u
}

func StringUUID(u gocql.UUID) string {
	var offsets = [...]int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34}
	const hexString = "0123456789abcdef"
	r := make([]byte, 36)
	for i, b := range u {
		r[offsets[i]] = hexString[b>>4]
		r[offsets[i]+1] = hexString[b&0xF]
	}
	r[8] = '-'
	r[13] = '-'
	r[18] = '-'
	r[23] = '-'
	return string(r)
}

func TestCassandra_NewCassandraStore(t *testing.T) {

	Convey("When I create a new CassandraStore", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		Convey("Then the it should implement Manipulator interface", func() {

			var i interface{} = store
			var ok bool
			_, ok = i.(manipulate.Manipulator)
			So(ok, ShouldBeTrue)
		})
	})
}

func TestCassandra_Stop(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method stop", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var closeHasBeenCalled bool

		apomock.Override("gocql.Session.Close", func(session *gocql.Session) {
			closeHasBeenCalled = true
		})

		store.Stop()

		Convey("Then the native session should be close", func() {
			So(store.nativeSession, ShouldBeNil)
			So(closeHasBeenCalled, ShouldBeTrue)
		})
	})
}

func TestCassandra_Query(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method query", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		command := "SELECT * FROM table WHERE ID IN (?,?,?)"
		values := []interface{}{"1", "2", "3"}

		var expectedCommand string
		var expectedValues []interface{}
		expectedQuery := &gocql.Query{}

		apomock.Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		query := store.Query(command, values)

		Convey("Then we should get the good query", func() {
			So(expectedCommand, ShouldEqual, command)
			So(expectedValues, ShouldResemble, values)
			So(expectedQuery, ShouldEqual, query)
		})

	})
}

func TestCassandre_Batch(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method batch", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedBatchType gocql.BatchType
		expectedBatch := &gocql.Batch{}

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			expectedBatchType = t
			return expectedBatch
		})

		newBatch := store.Batch()

		Convey("Then we should get the good batch", func() {
			So(expectedBatch, ShouldEqual, newBatch)
			So(expectedBatchType, ShouldEqual, gocql.UnloggedBatch)
		})
	})
}

func TestCassandre_BatchWithAsynchroneBatch(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method batch", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}
		store.UseAsynchroneBatch = true

		var expectedBatchType gocql.BatchType
		expectedBatch := &gocql.Batch{}

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			expectedBatchType = t
			return expectedBatch
		})

		store.asynchroneBatch = store.nativeSession.NewBatch(gocql.UnloggedBatch)
		newBatch := store.Batch()

		Convey("Then we should get the good batch", func() {
			So(store.asynchroneBatch, ShouldEqual, newBatch)
			So(expectedBatchType, ShouldEqual, gocql.UnloggedBatch)
		})
	})
}

func TestCassandre_Commit(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method Commit", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}
		store.UseAsynchroneBatch = true

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		store.asynchroneBatch = store.nativeSession.NewBatch(gocql.UnloggedBatch)
		batch := store.Batch()

		var expectedBatch *gocql.Batch

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			expectedBatch = b
			return nil
		})

		err := store.Commit()

		Convey("Then we should get the good batch", func() {
			So(expectedBatch, ShouldEqual, batch)
			So(store.asynchroneBatch, ShouldNotEqual, batch)
			So(err, ShouldBeNil)
		})
	})
}

func TestCassandre_Commit_Error(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method Commit", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}
		store.UseAsynchroneBatch = true

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		store.asynchroneBatch = store.nativeSession.NewBatch(gocql.UnloggedBatch)
		batch := store.Batch()

		var expectedBatch *gocql.Batch
		expectedError := errors.New("error batch")

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			expectedBatch = b
			return expectedError
		})

		newError := store.Commit()
		expectedErrors := elemental.Errors{elemental.NewError("CassandraStore batch commit failed", "error batch", "", 500)}

		Convey("Then we should get the good batch", func() {
			So(expectedBatch, ShouldEqual, batch)
			So(store.asynchroneBatch, ShouldEqual, batch)
			So(expectedBatch, ShouldEqual, batch)
			So(expectedErrors, ShouldResemble, newError)
		})
	})
}

func TestCassandra_ExecuteBatch(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method executeBash", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		batch := &gocql.Batch{}

		var expectedBatch *gocql.Batch
		expectedError := errors.New("error batch")

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			expectedBatch = b
			return expectedError
		})

		newError := store.ExecuteBatch(batch)

		Convey("Then we should see that execute bash has been called", func() {
			So(expectedBatch, ShouldEqual, batch)
			So(expectedError, ShouldEqual, newError)
		})
	})
}

func TestCassandra_sliceMaps_SliceMapError(t *testing.T) {
	Convey("When I call the method sliceMaps with the iterator which has an error", t, func() {

		iter := &gocql.Iter{}

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return nil, errors.New("Error iter")
		})

		maps, err := sliceMaps(iter)
		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore SliceMap error", "Error iter", "", 500)}

		Convey("Then I should get an error in return", func() {
			So(maps, ShouldBeNil)
			So(err[0], ShouldResemble, expectedErrors[0])
		})
	})
}

func TestCassandra_sliceMaps_IterCloseError(t *testing.T) {
	Convey("When I call the method sliceMaps with the iterator which has an error", t, func() {

		iter := &gocql.Iter{}

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return nil, nil
		})

		apomock.Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return errors.New("Iter close error")
		})

		maps, err := sliceMaps(iter)
		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore Iter Close error", "Iter close error", "", 500)}

		Convey("Then I should get an error in return", func() {
			So(maps, ShouldBeNil)
			So(err[0], ShouldResemble, expectedErrors[0])
		})
	})
}

func TestCassandra_sliceMaps_EmptySliceMap(t *testing.T) {
	Convey("When I call the method sliceMaps with the iterator which has an error when closing", t, func() {

		iter := &gocql.Iter{}

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return make([]map[string]interface{}, 0), nil
		})

		apomock.Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		maps, err := sliceMaps(iter)

		Convey("Then I should get an error in return", func() {
			So(maps, ShouldResemble, make([]map[string]interface{}, 0))
			So(err, ShouldBeNil)
		})
	})
}

func TestCassandra_sliceMaps(t *testing.T) {
	Convey("When I call the method sliceMaps with the iterator which has an error because an empty map", t, func() {

		iter := &gocql.Iter{}

		var expectedValues []map[string]interface{}

		value := make(map[string]interface{})
		value["ID"] = "123"
		value["Environment"] = 4

		expectedValues = append(expectedValues, value)

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedValues, nil
		})

		apomock.Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		maps, err := sliceMaps(iter)

		Convey("Then I should get an error in return", func() {
			So(err, ShouldBeNil)
			So(expectedValues, ShouldResemble, maps)
		})
	})
}

func TestCassandra_Start(t *testing.T) {
	Convey("When I call the method start, the session should be init with the cassandre store", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {

			return session, nil
		})

		store.Start()

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(store.nativeSession, ShouldEqual, session)
			So(store.asynchroneBatch, ShouldResemble, store.nativeSession.NewBatch(gocql.UnloggedBatch))
			So(cluster.Keyspace, ShouldEqual, "keyspace")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
		})

	})
}

func TestCassandra_StartWithPanic(t *testing.T) {
	Convey("When I call the method start, the session should be init with the cassandre store", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		apomock.Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			return &gocql.ClusterConfig{}
		})

		apomock.Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {

			return nil, errors.New("should panic session")
		})

		panicFunc := func() {
			store.Start()
		}

		Convey("Then I should get a panic", func() {
			So(panicFunc, ShouldPanic)
		})

	})
}

func TestCassandra_unmarshalInterfaceWithEmptySliceMap(t *testing.T) {

	Convey("When I call the method unmarshalInterface and get an error from sliceMaps", t, func() {

		iter := &gocql.Iter{}

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return []map[string]interface{}{}, nil
		})

		errs := unmarshalInterface(iter, nil)

		Convey("Then I should get a panic", func() {
			So(errs, ShouldBeNil)
		})

	})
}

func TestCassandra_unmarshalInterfaceWithErrorFromSliceMap(t *testing.T) {

	Convey("When I call the method unmarshalInterface and get an error from sliceMaps", t, func() {

		iter := &gocql.Iter{}

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return nil, errors.New("Error iter")
		})

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore SliceMap error", "Error iter", "", 500)}

		errs := unmarshalInterface(iter, nil)

		Convey("Then I should get a panic", func() {
			So(errs[0], ShouldResemble, expectedErrors[0])
		})

	})
}

func TestCassandra_unmarshalInterfaceWithErrorFromUnmarshal(t *testing.T) {

	Convey("When I call the method unmarshalInterface and get an error from Unmarshal", t, func() {

		iter := &gocql.Iter{}

		var expectedValues []map[string]interface{}

		value := make(map[string]interface{})
		value["ID"] = "123"
		value["Environment"] = 4

		expectedValues = append(expectedValues, value)

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedValues, nil
		})

		apomock.Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		apomock.Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {
			return errors.New("Errors from unmarshal")
		})

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore Unmarshal error", "Errors from unmarshal", "%!s(<nil>)", 500)}

		errs := unmarshalInterface(iter, nil)

		Convey("Then I should get a panic", func() {
			So(errs[0], ShouldResemble, expectedErrors[0])
		})

	})
}

func TestCassandra_unmarshalInterface(t *testing.T) {

	Convey("When I call the method unmarshalInterface", t, func() {

		iter := &gocql.Iter{}

		var expectedValues []map[string]interface{}

		value := make(map[string]interface{})
		value["ID"] = "123"
		value["Environment"] = 4

		expectedValues = append(expectedValues, value)

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedValues, nil
		})

		apomock.Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		var expectedDatas interface{}
		var expectedV interface{}

		apomock.Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {
			expectedDatas = data
			expectedV = v

			return nil
		})

		var dest []interface{}

		errs := unmarshalInterface(iter, dest)

		Convey("Then I should get a panic", func() {
			So(errs, ShouldBeNil)
			So(expectedDatas, ShouldResemble, expectedValues)
			So(expectedV, ShouldEqual, dest)
		})
	})
}

func TestCassandra_unmarshalManipulableWithErrorFromSliceMap(t *testing.T) {

	Convey("When I call the method unmarshalManipulable and get an error from sliceMaps", t, func() {

		iter := &gocql.Iter{}

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return nil, errors.New("Error iter")
		})

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore SliceMap error", "Error iter", "", 500)}

		errs := unmarshalManipulable(iter, nil)

		Convey("Then I should get a panic", func() {
			So(errs[0], ShouldResemble, expectedErrors[0])
		})

	})
}

func TestCassandra_unmarshalManipulableWithErrorFromUnmarshal(t *testing.T) {

	Convey("When I call the method unmarshalManipulable and get an error from Unmarshal", t, func() {

		iter := &gocql.Iter{}

		var expectedValues []map[string]interface{}

		value := make(map[string]interface{})
		value["ID"] = "123"
		value["Environment"] = 4

		expectedValues = append(expectedValues, value)

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedValues, nil
		})

		apomock.Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		apomock.Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {
			return errors.New("Errors from unmarshal")
		})

		tag := &Tag{}
		tag.Description = "Tag description"

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore Unmarshal error", "Errors from unmarshal", "&{ Tag description  %!s(int=0)}", 500)}
		errs := unmarshalManipulable(iter, []manipulate.Manipulable{tag})

		Convey("Then I should get a panic", func() {
			So(errs[0], ShouldResemble, expectedErrors[0])
		})

	})
}

func TestCassandra_unmarshalManipulable(t *testing.T) {

	Convey("When I call the method unmarshalManipulable", t, func() {

		iter := &gocql.Iter{}

		var expectedValues []map[string]interface{}

		value := make(map[string]interface{})
		value["ID"] = "123"
		value["Environment"] = 4

		value2 := make(map[string]interface{})
		value2["ID"] = "456"
		value2["Environment"] = 5

		expectedValues = append(expectedValues, value, value2)

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedValues, nil
		})

		apomock.Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		var expectedDatas1 interface{}
		var expectedV1 interface{}
		var expectedDatas2 interface{}
		var expectedV2 interface{}

		apomock.Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {

			if expectedDatas1 == nil {
				expectedDatas1 = data
				expectedV1 = v
				return nil
			}

			expectedDatas2 = data
			expectedV2 = v

			return nil
		})

		tag1 := &Tag{}
		tag1.Description = "Tag description"

		tag2 := &Tag{}
		tag2.Description = "Tag 2 description"

		errs := unmarshalManipulable(iter, []manipulate.Manipulable{tag1, tag2})

		Convey("Then I should get a panic", func() {
			So(errs, ShouldBeNil)

			So(expectedDatas1, ShouldResemble, value)
			So(expectedV1, ShouldEqual, tag1)

			So(expectedDatas2, ShouldResemble, value2)
			So(expectedV2, ShouldEqual, tag2)
		})
	})
}

func TestCassandra_Count(t *testing.T) {

	Convey("When I call the method Count", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedCommand string
		var expectedValues []interface{}
		expectedQuery := &gocql.Query{}

		apomock.Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		apomock.Override("gocql.Iter.Scan", func(iter *gocql.Iter, i ...interface{}) bool {
			reflect.ValueOf(i[0]).Elem().SetInt(100)
			return true
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		count, err := store.Count(context, TagIdentity)

		Convey("Then I should get no error", func() {
			So(err, ShouldBeNil)
			So(expectedCommand, ShouldEqual, "SELECT COUNT(*) FROM tag LIMIT 10")
			So(expectedValues, ShouldResemble, []interface{}{})
			So(count, ShouldEqual, 100)
		})
	})
}

func TestCassandra_CountErrorScan(t *testing.T) {

	Convey("When I call the method Count", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedCommand string
		var expectedValues []interface{}
		expectedQuery := &gocql.Query{}

		apomock.Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		apomock.Override("gocql.Iter.Scan", func(iter *gocql.Iter, i ...interface{}) bool {
			return false
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		count, err := store.Count(context, TagIdentity)
		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore Iter Close error", "Error when scanning the iterator of a count command", "tag", 500)}

		Convey("Then I should get no error", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
			So(expectedCommand, ShouldEqual, "SELECT COUNT(*) FROM tag LIMIT 10")
			So(expectedValues, ShouldResemble, []interface{}{})
			So(count, ShouldEqual, -1)
		})
	})
}

func TestCassandra_CountErrorCloseIter(t *testing.T) {

	Convey("When I call the method Count", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedCommand string
		var expectedValues []interface{}
		expectedQuery := &gocql.Query{}

		apomock.Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		apomock.Override("gocql.Iter.Scan", func(iter *gocql.Iter, i ...interface{}) bool {
			return true
		})

		apomock.Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return errors.New("Iter close error")
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		count, err := store.Count(context, TagIdentity)
		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore Iter Close error", "Error when closing the iterator of a count command", "tag", 500)}

		Convey("Then I should get no error", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
			So(expectedCommand, ShouldEqual, "SELECT COUNT(*) FROM tag LIMIT 10")
			So(expectedValues, ShouldResemble, []interface{}{})
			So(count, ShouldEqual, -1)
		})
	})
}

func TestCassandra_RetrieveChildren(t *testing.T) {

	Convey("When I call the method RetrieveChildren", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedCommand string
		var expectedValues []interface{}
		expectedQuery := &gocql.Query{}

		apomock.Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		apomock.Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		var expectedV interface{}

		apomock.Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {
			expectedV = v
			return errors.New("Errors from unmarshal")
		})

		var tags []*Tag
		context := &manipulate.Context{}
		context.PageSize = 10

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore Unmarshal error", "Errors from unmarshal", "&[]", 500)}
		err := store.RetrieveChildren(context, nil, TagIdentity, &tags)

		Convey("Then I should get no error", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
			So(expectedCommand, ShouldEqual, "SELECT * FROM tag LIMIT 10")
			So(expectedValues, ShouldResemble, []interface{}{})
			So(expectedV, ShouldEqual, &tags)
		})
	})
}

func TestCassandra_RetrieveWithError(t *testing.T) {

	Convey("When I call the method Retrieve", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedCommand string
		var expectedValues []interface{}
		expectedQuery := &gocql.Query{}

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{}, []interface{}{}, nil
		})

		apomock.Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		tag := &Tag{}
		tag.ID = "1234"

		context := &manipulate.Context{}
		context.PageSize = 10

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore Unmarshal error", "The number of the given objects and the number of results fetched is different", "[%!s(*manipcassandra.Tag=&{1234   0})]", 500)}
		err := store.Retrieve(context, tag)

		Convey("Then I should get an error", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
			So(expectedCommand, ShouldEqual, "SELECT * FROM tag LIMIT 10")
			So(expectedValues, ShouldResemble, []interface{}{})
		})
	})
}

func TestCassandra_UpdateCollection(t *testing.T) {

	Convey("When I call the method UpdateCollection", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"123"}, nil
		})

		tag := &Tag{}
		tag.ID = "1234"

		a := &AttributeUpdater{}
		a.Key = "NAME"
		a.Operation = elemental.OperationSubstractive
		a.Values = "coucou"

		var expectedCommand string
		var expectedValues []interface{}
		expectedQuery := &gocql.Query{}

		apomock.Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.Override("gocql.Query.Exec", func(session *gocql.Query) error {
			return nil
		})

		context := &manipulate.Context{}
		err := store.UpdateCollection(context, a, tag)

		Convey("Then I should get an error", func() {
			So(err, ShouldBeNil)
			So(tag, ShouldResemble, &Tag{ID: "1234"})
			So(expectedCommand, ShouldResemble, "UPDATE tag SET NAME = NAME - ? WHERE ID = ?")
			So(expectedValues, ShouldResemble, []interface{}{"coucou", "123"})
		})
	})
}

func TestCassandra_UpdateCollectionWithErrorQuery(t *testing.T) {

	Convey("When I call the method UpdateCollection", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"123"}, nil
		})

		tag := &Tag{}
		tag.ID = "1234"

		a := &AttributeUpdater{}
		a.Key = "NAME"
		a.Operation = elemental.OperationSubstractive
		a.Values = "coucou"

		var expectedCommand string
		var expectedValues []interface{}
		expectedQuery := &gocql.Query{}

		apomock.Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.Override("gocql.Query.Exec", func(session *gocql.Query) error {
			return errors.New("CassandraStore query error")
		})

		context := &manipulate.Context{}

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore UpdateCollection failed", "CassandraStore query error", "&{1234   %!s(int=0)}", 500)}
		err := store.UpdateCollection(context, a, tag)

		Convey("Then I should get an error", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
			So(tag, ShouldResemble, &Tag{ID: "1234"})
			So(expectedCommand, ShouldResemble, "UPDATE tag SET NAME = NAME - ? WHERE ID = ?")
			So(expectedValues, ShouldResemble, []interface{}{"coucou", "123"})
		})
	})
}

func TestCassandra_UpdateCollectionWithErrorPrimaryFields(t *testing.T) {

	Convey("When I call the method UpdateCollection", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{}, []interface{}{}, errors.New("CassandraStore PrimaryFieldsAndValues error")
		})

		tag := &Tag{}
		tag.ID = "1234"

		context := &manipulate.Context{}

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore UpdateCollection failed", "CassandraStore PrimaryFieldsAndValues error", "&{1234   %!s(int=0)}", 500)}
		err := store.UpdateCollection(context, nil, tag)

		Convey("Then I should get an error", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
			So(tag, ShouldResemble, &Tag{ID: "1234"})
		})
	})
}

func TestCassandra_RetrieveWithErrorPrimaryFields(t *testing.T) {

	Convey("When I call the method Retrieve", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{}, []interface{}{}, errors.New("CassandraStore PrimaryFieldsAndValues error")
		})

		tag := &Tag{}
		tag.ID = "1234"

		context := &manipulate.Context{}
		context.PageSize = 10

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore retrieve failed", "CassandraStore PrimaryFieldsAndValues error", "&{1234   %!s(int=0)}", 500)}
		err := store.Retrieve(context, tag)

		Convey("Then I should get an error", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
			So(tag, ShouldResemble, &Tag{ID: "1234"})
		})
	})
}

func TestCassandra_Retrieve(t *testing.T) {

	Convey("When I call the method Retrieve", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedCommand string
		var expectedValues []interface{}
		expectedQuery := &gocql.Query{}

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"678"}, nil
		})

		apomock.Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		var sliceMaps []map[string]interface{}

		value := make(map[string]interface{})
		value["ID"] = "123"
		value["Environment"] = 4

		sliceMaps = append(sliceMaps, value)

		apomock.Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return sliceMaps, nil
		})

		apomock.Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {
			return nil
		})

		apomock.Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		tag := &Tag{}
		tag.ID = "1234"

		context := &manipulate.Context{}
		context.PageSize = 10

		err := store.Retrieve(context, tag)

		Convey("Then I should get an error", func() {
			So(err, ShouldBeNil)
			So(expectedCommand, ShouldEqual, "SELECT * FROM tag WHERE ID = ? LIMIT 10")
			So(expectedValues, ShouldResemble, []interface{}{"678"})
		})
	})
}

func TestCassandra_Delete(t *testing.T) {

	Convey("When I call the method Delete", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedBatchType gocql.BatchType
		expectedBatch := &gocql.Batch{}

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"123"}, nil
		})

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			expectedBatchType = t
			return expectedBatch
		})

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

			if expectedValues1 == nil {
				expectedQuery1 = command
				expectedValues1 = values
				return
			}

			expectedQuery2 = command
			expectedValues2 = values

		})

		var expectedExecutedBatch *gocql.Batch

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {

			expectedExecutedBatch = b
			return nil

		})

		tag1 := &Tag{}
		tag1.ID = "123"

		tag2 := &Tag{}
		tag2.ID = "456"

		context := &manipulate.Context{}
		context.PageSize = 10

		err := store.Delete(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err, ShouldBeNil)
			So(expectedBatchType, ShouldEqual, gocql.UnloggedBatch)
			So(expectedQuery1, ShouldEqual, "DELETE FROM tag WHERE ID = ? LIMIT 10")
			So(expectedQuery2, ShouldEqual, "DELETE FROM tag WHERE ID = ? LIMIT 10")
			So(expectedValues1, ShouldResemble, []interface{}{"123"})
			So(expectedValues2, ShouldResemble, []interface{}{"123"})
			So(expectedBatch, ShouldEqual, expectedExecutedBatch)
		})
	})
}

func TestCassandra_Delete_AsynchroneBatch(t *testing.T) {

	Convey("When I call the method Delete", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}
		store.asynchroneBatch = &gocql.Batch{}
		store.UseAsynchroneBatch = true

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"123"}, nil
		})

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

			if expectedValues1 == nil {
				expectedQuery1 = command
				expectedValues1 = values
				return
			}

			expectedQuery2 = command
			expectedValues2 = values

		})

		tag1 := &Tag{}
		tag1.ID = "123"

		tag2 := &Tag{}
		tag2.ID = "456"

		context := &manipulate.Context{}
		context.PageSize = 10

		err := store.Delete(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err, ShouldBeNil)
			So(expectedQuery1, ShouldEqual, "DELETE FROM tag WHERE ID = ? LIMIT 10")
			So(expectedQuery2, ShouldEqual, "DELETE FROM tag WHERE ID = ? LIMIT 10")
			So(expectedValues1, ShouldResemble, []interface{}{"123"})
			So(expectedValues2, ShouldResemble, []interface{}{"123"})
		})
	})
}

func TestCassandra_DeleteError(t *testing.T) {

	Convey("When I call the method Delete", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return errors.New("Error batch")
		})

		tag1 := &Tag{}
		tag1.ID = "123"

		tag2 := &Tag{}
		tag2.ID = "456"

		context := &manipulate.Context{}
		context.PageSize = 10

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore delete failed", "Error batch", "[%!s(*manipcassandra.Tag=&{123   0}) %!s(*manipcassandra.Tag=&{456   0})]", 500)}
		err := store.Delete(context, tag1, tag2)

		Convey("Then I should get an error", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
		})
	})
}

func TestCassandra_DeleteWithErrorPrimaryFields(t *testing.T) {

	Convey("When I call the method Retrieve", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{}, []interface{}{}, errors.New("CassandraStore PrimaryFieldsAndValues error")
		})

		tag := &Tag{}
		tag.ID = "1234"

		context := &manipulate.Context{}
		context.PageSize = 10

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore delete failed", "CassandraStore PrimaryFieldsAndValues error", "&{1234   %!s(int=0)}", 500)}
		err := store.Delete(context, tag)

		Convey("Then I should get an error", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
			So(tag, ShouldResemble, &Tag{ID: "1234"})
		})
	})
}

func TestCassandra_Update(t *testing.T) {

	Convey("When I call the method Update", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedBatchType gocql.BatchType
		expectedBatch := &gocql.Batch{}

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			expectedBatchType = t
			return expectedBatch
		})

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

			if expectedValues1 == nil {
				expectedQuery1 = command
				expectedValues1 = values
				return
			}

			expectedQuery2 = command
			expectedValues2 = values

		})

		var expectedExecutedBatch *gocql.Batch

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {

			expectedExecutedBatch = b
			return nil

		})

		tag1 := &Tag{}
		tag1.ID = "123"
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.ID = "456"
		tag2.Description = "description 2"

		var expectedValue1 interface{}
		var expectedValue2 interface{}

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"456"}, nil
		})

		apomock.Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {

			if expectedValue1 == nil {
				expectedValue1 = val
				return []string{"ID", "description"}, []interface{}{"123", "description 1"}, nil
			}

			expectedValue2 = val
			return []string{"ID", "description"}, []interface{}{"456", "description 2"}, nil
		})

		context := &manipulate.Context{}
		context.PageSize = 10
		context.Attributes = []string{"description"}

		err := store.Update(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err, ShouldBeNil)
			So(expectedBatchType, ShouldEqual, gocql.UnloggedBatch)
			So(expectedQuery1, ShouldEqual, "UPDATE tag SET description = ? WHERE ID = ? LIMIT 10")
			So(expectedQuery2, ShouldEqual, "UPDATE tag SET description = ? WHERE ID = ? LIMIT 10")
			So(expectedValues1, ShouldResemble, []interface{}{"description 1", "456"})
			So(expectedValues2, ShouldResemble, []interface{}{"description 2", "456"})
			So(expectedBatch, ShouldEqual, expectedExecutedBatch)
			So(expectedValue1, ShouldEqual, tag1)
			So(expectedValue2, ShouldEqual, tag2)
		})
	})
}

func TestCassandra_Update_AsynchroneBatch(t *testing.T) {

	Convey("When I call the method Update", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}
		store.asynchroneBatch = &gocql.Batch{}
		store.UseAsynchroneBatch = true

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

			if expectedValues1 == nil {
				expectedQuery1 = command
				expectedValues1 = values
				return
			}

			expectedQuery2 = command
			expectedValues2 = values

		})

		tag1 := &Tag{}
		tag1.ID = "123"
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.ID = "456"
		tag2.Description = "description 2"

		var expectedValue1 interface{}
		var expectedValue2 interface{}

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"456"}, nil
		})

		apomock.Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {

			if expectedValue1 == nil {
				expectedValue1 = val
				return []string{"ID", "description"}, []interface{}{"123", "description 1"}, nil
			}

			expectedValue2 = val
			return []string{"ID", "description"}, []interface{}{"456", "description 2"}, nil
		})

		context := &manipulate.Context{}
		context.PageSize = 10
		context.Attributes = []string{"description"}

		err := store.Update(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err, ShouldBeNil)
			So(expectedQuery1, ShouldEqual, "UPDATE tag SET description = ? WHERE ID = ? LIMIT 10")
			So(expectedQuery2, ShouldEqual, "UPDATE tag SET description = ? WHERE ID = ? LIMIT 10")
			So(expectedValues1, ShouldResemble, []interface{}{"description 1", "456"})
			So(expectedValues2, ShouldResemble, []interface{}{"description 2", "456"})
			So(expectedValue1, ShouldEqual, tag1)
			So(expectedValue2, ShouldEqual, tag2)
		})
	})
}

func TestCassandra_Update_ErrorFieldsAndValues(t *testing.T) {

	Convey("When I call the method Update", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return nil
		})

		tag1 := &Tag{}
		tag1.ID = "123"
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.ID = "456"
		tag2.Description = "description 2"

		apomock.Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return nil, nil, errors.New("Error from FieldsAndValues")
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore update failed", "Error from FieldsAndValues", "&{123 description 1  %!s(int=0)}", 500)}
		err := store.Update(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
		})
	})
}

func TestCassandra_Update_ErrorPrimaryFieldsAndValues(t *testing.T) {

	Convey("When I call the method Update", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{}, []interface{}{}, errors.New("CassandraStore PrimaryFieldsAndValues error")
		})

		tag1 := &Tag{}
		tag1.ID = "123"
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.ID = "456"
		tag2.Description = "description 2"

		context := &manipulate.Context{}
		context.PageSize = 10

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore update failed", "CassandraStore PrimaryFieldsAndValues error", "&{123 description 1  %!s(int=0)}", 500)}
		err := store.Update(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
		})
	})
}

func TestCassandra_Update_ErrorExecuteBatch(t *testing.T) {

	Convey("When I call the method Update", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return nil
		})

		tag1 := &Tag{}
		tag1.ID = "123"
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.ID = "456"
		tag2.Description = "description 2"

		apomock.Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID", "description"}, []interface{}{"123", "description 1"}, nil
		})

		apomock.Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{}, []interface{}{}, nil
		})

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return errors.New("Error batch")
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore update failed", "Error batch", "[%!s(*manipcassandra.Tag=&{123 description 1  0}) %!s(*manipcassandra.Tag=&{456 description 2  0})]", 500)}
		err := store.Update(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
		})
	})
}

func TestCassandra_Create(t *testing.T) {

	Convey("When I call the method Create", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedBatchType gocql.BatchType
		expectedBatch := &gocql.Batch{}

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			expectedBatchType = t
			return expectedBatch
		})

		apomock.Override("gocql.TimeUUID", func() gocql.UUID {
			return ParseUUID("7c469413-12ed-11e6-ac73-f45c89941b79")
		})

		apomock.Override("gocql.UUID.String", func(uuid gocql.UUID) string {
			return StringUUID(uuid)
		})

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

			t.Log("couocu")

			if expectedValues1 == nil {
				expectedQuery1 = command
				expectedValues1 = values
				return
			}

			expectedQuery2 = command
			expectedValues2 = values

		})

		var expectedExecutedBatch *gocql.Batch

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {

			expectedExecutedBatch = b
			return nil

		})

		tag1 := &Tag{}
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.Description = "description 2"

		var expectedValue1 interface{}
		var expectedValue2 interface{}

		apomock.Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {

			if expectedValue1 == nil {
				expectedValue1 = val
				return []string{"ID", "description"}, []interface{}{"123", "description 1"}, nil
			}

			expectedValue2 = val
			return []string{"ID", "description"}, []interface{}{"456", "description 2"}, nil
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		err := store.Create(context, nil, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err, ShouldBeNil)
			So(expectedBatchType, ShouldEqual, gocql.UnloggedBatch)
			So(expectedQuery1, ShouldEqual, "INSERT INTO tag (ID, description) VALUES (?, ?) LIMIT 10")
			So(expectedQuery2, ShouldEqual, "INSERT INTO tag (ID, description) VALUES (?, ?) LIMIT 10")
			So(expectedValues1, ShouldResemble, []interface{}{"123", "description 1"})
			So(expectedValues2, ShouldResemble, []interface{}{"456", "description 2"})
			So(expectedBatch, ShouldEqual, expectedExecutedBatch)
			So(expectedValue1, ShouldEqual, tag1)
			So(expectedValue2, ShouldEqual, tag2)
			So(len(tag2.ID), ShouldEqual, 36)
			So(len(tag2.ID), ShouldEqual, 36)
		})
	})
}

func TestCassandra_Create_AsynchroneBatch(t *testing.T) {

	Convey("When I call the method Create", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}
		store.asynchroneBatch = &gocql.Batch{}
		store.UseAsynchroneBatch = true

		apomock.Override("gocql.TimeUUID", func() gocql.UUID {
			return ParseUUID("7c469413-12ed-11e6-ac73-f45c89941b79")
		})

		apomock.Override("gocql.UUID.String", func(uuid gocql.UUID) string {
			return StringUUID(uuid)
		})

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

			t.Log("couocu")

			if expectedValues1 == nil {
				expectedQuery1 = command
				expectedValues1 = values
				return
			}

			expectedQuery2 = command
			expectedValues2 = values

		})

		tag1 := &Tag{}
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.Description = "description 2"

		var expectedValue1 interface{}
		var expectedValue2 interface{}

		apomock.Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {

			if expectedValue1 == nil {
				expectedValue1 = val
				return []string{"ID", "description"}, []interface{}{"123", "description 1"}, nil
			}

			expectedValue2 = val
			return []string{"ID", "description"}, []interface{}{"456", "description 2"}, nil
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		err := store.Create(context, nil, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err, ShouldBeNil)
			So(expectedQuery1, ShouldEqual, "INSERT INTO tag (ID, description) VALUES (?, ?) LIMIT 10")
			So(expectedQuery2, ShouldEqual, "INSERT INTO tag (ID, description) VALUES (?, ?) LIMIT 10")
			So(expectedValues1, ShouldResemble, []interface{}{"123", "description 1"})
			So(expectedValues2, ShouldResemble, []interface{}{"456", "description 2"})
			So(expectedValue1, ShouldEqual, tag1)
			So(expectedValue2, ShouldEqual, tag2)
			So(len(tag2.ID), ShouldEqual, 36)
			So(len(tag2.ID), ShouldEqual, 36)
		})
	})
}

func TestCassandra_Create_ErrorFieldsAndValues(t *testing.T) {

	Convey("When I call the method Create", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return nil
		})

		apomock.Override("gocql.TimeUUID", func() gocql.UUID {
			return ParseUUID("7c469413-12ed-11e6-ac73-f45c89941b79")
		})

		apomock.Override("gocql.UUID.String", func(uuid gocql.UUID) string {
			return StringUUID(uuid)
		})

		tag1 := &Tag{}
		tag1.ID = "123"
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.ID = "456"
		tag2.Description = "description 2"

		apomock.Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return nil, nil, errors.New("Error from FieldsAndValues")
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore create failed", "Error from FieldsAndValues", "&{ description 1  %!s(int=0)}", 500)}
		err := store.Create(context, nil, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
			So(len(tag2.ID), ShouldEqual, 0)
			So(len(tag2.ID), ShouldEqual, 0)
		})
	})
}

func TestCassandra_Create_ErrorExecuteBatch(t *testing.T) {

	Convey("When I call the method Create", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return nil
		})

		apomock.Override("gocql.TimeUUID", func() gocql.UUID {
			return ParseUUID("7c469413-12ed-11e6-ac73-f45c89941b79")
		})

		apomock.Override("gocql.UUID.String", func(uuid gocql.UUID) string {
			return StringUUID(uuid)
		})

		tag1 := &Tag{}
		tag1.ID = "123"
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.ID = "456"
		tag2.Description = "description 2"

		apomock.Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID", "description"}, []interface{}{"123", "description 1"}, nil
		})

		apomock.Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return errors.New("Error batch")
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		expectedErrors := []*elemental.Error{elemental.NewError("CassandraStore create failed", "Error batch", "[%!s(*manipcassandra.Tag=&{ description 1  0}) %!s(*manipcassandra.Tag=&{ description 2  0})]", 500)}
		err := store.Create(context, nil, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err[0], ShouldResemble, expectedErrors[0])
			So(len(tag2.ID), ShouldEqual, 0)
			So(len(tag2.ID), ShouldEqual, 0)
		})
	})
}
