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

		apomock.T(t).Override("gocql.Session.Close", func(session *gocql.Session) {
			closeHasBeenCalled = true
		})

		store.Stop()

		Convey("Then the native session should be close", func() {
			So(store.nativeSession, ShouldBeNil)
			So(closeHasBeenCalled, ShouldBeTrue)
		})
	})
}

func TestCassandre_BatchForID(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method batch with no id", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedBatchType gocql.BatchType
		expectedBatch := &gocql.Batch{}

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			expectedBatchType = t
			return expectedBatch
		})

		newBatch := store.batchForID("")

		Convey("Then we should get the good batch", func() {
			So(expectedBatch, ShouldEqual, newBatch)
			So(expectedBatchType, ShouldEqual, gocql.UnloggedBatch)
		})
	})
}

func TestCassandre_BatchForIDWithAnID(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method BatchForID, NewBatch should be called once", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedBatchType gocql.BatchType
		expectedBatch := &gocql.Batch{}
		var numberOfCalls int

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			expectedBatchType = t
			numberOfCalls++
			return expectedBatch
		})

		newBatch := store.batchForID("123")

		Convey("Then we should get the good batch", func() {
			So(expectedBatchType, ShouldEqual, gocql.UnloggedBatch)
			So(numberOfCalls, ShouldEqual, 1)
			So(store.batchRegistry["123"], ShouldEqual, newBatch)
		})
	})
}

func TestCassandre_CommitTransaction(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method CommitTransaction", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		batch := store.batchForID("123")

		var expectedBatch *gocql.Batch

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			expectedBatch = b
			return nil
		})

		err := store.Commit("123")

		Convey("Then we should get the good batch", func() {
			So(expectedBatch, ShouldEqual, batch)
			So(err, ShouldBeNil)
			So(store.batchRegistry["123"], ShouldBeNil)
		})
	})
}

func TestCassandre_CommitTransaction_ErrorBadID(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method Commit", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		errs := store.Commit("123")

		Convey("Then we should get the good batch", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotCommit)
		})
	})
}

func TestCassandre_CommitTransaction_Error(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method Commit", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		batch := store.batchForID("123")

		var expectedBatch *gocql.Batch

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			expectedBatch = b
			return errors.New("error batch")
		})

		errs := store.Commit("123")

		Convey("Then we should get the good batch", func() {
			So(expectedBatch, ShouldEqual, batch)
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotExecuteBatch)
			So(store.batchRegistry["123"], ShouldBeNil)
		})
	})
}

func TestCassandra_ExecuteBatch(t *testing.T) {
	Convey("When I create a new CassandraStore and call the method executeBash", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		batch := &gocql.Batch{}

		var expectedBatch *gocql.Batch
		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			expectedBatch = b
			return errors.New("error batch")
		})

		err := store.commitBatch(batch)

		Convey("Then we should see that execute bash has been called", func() {
			So(expectedBatch, ShouldEqual, batch)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestCassandra_sliceMaps_SliceMapError(t *testing.T) {
	Convey("When I call the method sliceMaps with the iterator which has an error", t, func() {

		iter := &gocql.Iter{}

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return nil, errors.New("Error iter")
		})

		maps, err := sliceMaps(iter)

		Convey("Then I should get an error in return", func() {
			So(maps, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestCassandra_sliceMaps_IterCloseError(t *testing.T) {
	Convey("When I call the method sliceMaps with the iterator which has an error", t, func() {

		iter := &gocql.Iter{}

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return nil, nil
		})

		apomock.T(t).Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return errors.New("Iter close error")
		})

		maps, err := sliceMaps(iter)

		Convey("Then I should get an error in return", func() {
			So(maps, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestCassandra_sliceMaps_EmptySliceMap(t *testing.T) {
	Convey("When I call the method sliceMaps with the iterator which has an error when closing", t, func() {

		iter := &gocql.Iter{}

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return make([]map[string]interface{}, 0), nil
		})

		apomock.T(t).Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
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

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedValues, nil
		})

		apomock.T(t).Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
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

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {

			return session, nil
		})

		store.Start()

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(store.nativeSession, ShouldEqual, session)
			So(cluster.Keyspace, ShouldEqual, "keyspace")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
		})

	})
}

func TestCassandra_StartWithError(t *testing.T) {
	Convey("When I call the method start, the session should be init with the cassandre store", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			return &gocql.ClusterConfig{}
		})

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {

			return nil, errors.New("should panic session")
		})

		err := store.Start()

		Convey("Then error should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})
}

func TestCassandra_unmarshalInterfaceWithEmptySliceMap(t *testing.T) {

	Convey("When I call the method unmarshalInterface and get an error from sliceMaps", t, func() {

		iter := &gocql.Iter{}

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return []map[string]interface{}{}, nil
		})

		errs := unmarshalManipulables(iter, nil)

		Convey("Then I should get a panic", func() {
			So(errs, ShouldBeNil)
		})

	})
}

func TestCassandra_unmarshalInterfaceWithErrorFromSliceMap(t *testing.T) {

	Convey("When I call the method unmarshalInterface and get an error from sliceMaps", t, func() {

		iter := &gocql.Iter{}

		apomock.T(t).Override("gocql.Iter.NumRows", func(iter *gocql.Iter) int {
			return 1
		})

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return nil, errors.New("Error iter")
		})

		apomock.T(t).Override("gocql.Iter.NumRows", func(iter *gocql.Iter) int {
			return 1
		})

		err := unmarshalManipulables(iter, nil)

		Convey("Then I should get a panic", func() {
			So(err, ShouldNotBeNil)
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

		apomock.T(t).Override("gocql.Iter.NumRows", func(iter *gocql.Iter) int {
			return 1
		})

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedValues, nil
		})

		apomock.T(t).Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		apomock.T(t).Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {
			return errors.New("Errors from unmarshal")
		})

		err := unmarshalManipulables(iter, nil)

		Convey("Then I should get a panic", func() {
			So(err, ShouldNotBeNil)
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

		apomock.T(t).Override("gocql.Iter.NumRows", func(iter *gocql.Iter) int {
			return 1
		})

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedValues, nil
		})

		apomock.T(t).Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		var expectedDatas interface{}
		var expectedV interface{}

		apomock.T(t).Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {
			expectedDatas = data
			expectedV = v

			return nil
		})

		var dest []interface{}

		errs := unmarshalManipulables(iter, dest)

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

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return nil, errors.New("Error iter")
		})

		err := unmarshalManipulable(iter, nil)

		Convey("Then I should get a panic", func() {
			So(err, ShouldNotBeNil)
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

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedValues, nil
		})

		apomock.T(t).Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		apomock.T(t).Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {
			return errors.New("Errors from unmarshal")
		})

		tag := &Tag{}
		tag.Description = "Tag description"

		err := unmarshalManipulable(iter, tag)

		Convey("Then I should get a panic", func() {
			So(err, ShouldNotBeNil)
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

		expectedValues = append(expectedValues, value)

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedValues, nil
		})

		apomock.T(t).Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		var expectedDatas1 interface{}
		var expectedV1 interface{}

		apomock.T(t).Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {
			expectedDatas1 = data
			expectedV1 = v
			return nil
		})

		tag1 := &Tag{}
		tag1.Description = "Tag description"

		errs := unmarshalManipulable(iter, tag1)

		Convey("Then I should get a panic", func() {
			So(errs, ShouldBeNil)

			So(expectedDatas1, ShouldResemble, value)
			So(expectedV1, ShouldEqual, tag1)
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

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.T(t).Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		apomock.T(t).Override("gocql.Iter.Scan", func(iter *gocql.Iter, i ...interface{}) bool {
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

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.T(t).Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		apomock.T(t).Override("gocql.Iter.Scan", func(iter *gocql.Iter, i ...interface{}) bool {
			return false
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		count, errs := store.Count(context, TagIdentity)

		Convey("Then I should get no error", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotScan)
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

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.T(t).Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		apomock.T(t).Override("gocql.Iter.Scan", func(iter *gocql.Iter, i ...interface{}) bool {
			return true
		})

		apomock.T(t).Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return errors.New("Iter close error")
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		count, errs := store.Count(context, TagIdentity)

		Convey("Then I should get no error", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotCloseIterator)
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

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.T(t).Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		apomock.T(t).Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
			return nil
		})

		var expectedV interface{}

		apomock.T(t).Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {
			expectedV = v
			return errors.New("Errors from unmarshal")
		})

		var expectedMaps []map[string]interface{}

		value := make(map[string]interface{})
		value["ID"] = "123"
		value["Environment"] = 4

		expectedMaps = append(expectedMaps, value)

		apomock.T(t).Override("gocql.Iter.NumRows", func(iter *gocql.Iter) int {
			return 1
		})

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedMaps, nil
		})

		var tags []*Tag
		context := &manipulate.Context{}
		context.PageSize = 10

		tags = append(tags, &Tag{})
		errs := store.RetrieveChildren(context, nil, TagIdentity, &tags)

		Convey("Then I should get no error", func() {
			So(len(errs), ShouldEqual, 1)
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

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{}, []interface{}{}, nil
		})

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.T(t).Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		var expectedMaps []map[string]interface{}

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return expectedMaps, nil
		})

		tag := &Tag{}
		tag.ID = "1234"

		context := &manipulate.Context{}
		context.PageSize = 10

		errs := store.Retrieve(context, tag)

		Convey("Then I should get an error", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrObjectNotFound)
			So(expectedCommand, ShouldEqual, "SELECT * FROM tag LIMIT 10")
			So(expectedValues, ShouldResemble, []interface{}{})
		})
	})
}

func TestCassandra_UpdateCollection(t *testing.T) {

	Convey("When I call the method UpdateCollection", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"123"}, nil
		})

		tag := &Tag{}
		tag.ID = "1234"

		a := &AttributeUpdater{}
		a.Key = "NAME"
		a.AssignationType = elemental.AssignationTypeSubstract
		a.Values = "coucou"

		var expectedCommand string
		var expectedValues []interface{}
		expectedQuery := &gocql.Query{}

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.T(t).Override("gocql.Query.Exec", func(session *gocql.Query) error {
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

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"123"}, nil
		})

		tag := &Tag{}
		tag.ID = "1234"

		a := &AttributeUpdater{}
		a.Key = "NAME"
		a.AssignationType = elemental.AssignationTypeSubstract
		a.Values = "coucou"

		var expectedCommand string
		var expectedValues []interface{}
		expectedQuery := &gocql.Query{}

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.T(t).Override("gocql.Query.Exec", func(session *gocql.Query) error {
			return errors.New("CassandraStore query error")
		})

		context := &manipulate.Context{}

		errs := store.UpdateCollection(context, a, tag)

		Convey("Then I should get an error", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotExecuteQuery)
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

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{}, []interface{}{}, errors.New("CassandraStore PrimaryFieldsAndValues error")
		})

		tag := &Tag{}
		tag.ID = "1234"

		context := &manipulate.Context{}

		errs := store.UpdateCollection(context, nil, tag)

		Convey("Then I should get an error", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotExractPrimaryFieldsAndValues)
			So(tag, ShouldResemble, &Tag{ID: "1234"})
		})
	})
}

func TestCassandra_RetrieveWithErrorPrimaryFields(t *testing.T) {

	Convey("When I call the method Retrieve", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{}, []interface{}{}, errors.New("CassandraStore PrimaryFieldsAndValues error")
		})

		tag := &Tag{}
		tag.ID = "1234"

		context := &manipulate.Context{}
		context.PageSize = 10

		errs := store.Retrieve(context, tag)

		Convey("Then I should get an error", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotExractPrimaryFieldsAndValues)
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

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"678"}, nil
		})

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, command string, values ...interface{}) *gocql.Query {
			expectedCommand = command
			expectedValues = values

			return expectedQuery
		})

		apomock.T(t).Override("gocql.Query.Iter", func(q *gocql.Query) *gocql.Iter {
			return &gocql.Iter{}
		})

		var sliceMaps []map[string]interface{}

		value := make(map[string]interface{})
		value["ID"] = "123"
		value["Environment"] = 4

		sliceMaps = append(sliceMaps, value)

		apomock.T(t).Override("gocql.Iter.SliceMap", func(iter *gocql.Iter) ([]map[string]interface{}, error) {
			return sliceMaps, nil
		})

		apomock.T(t).Override("cassandra.Unmarshal", func(data interface{}, v interface{}) error {
			return nil
		})

		apomock.T(t).Override("gocql.Iter.Close", func(iter *gocql.Iter) error {
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

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"123"}, nil
		})

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			expectedBatchType = t
			return expectedBatch
		})

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.T(t).Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

			if expectedValues1 == nil {
				expectedQuery1 = command
				expectedValues1 = values
				return
			}

			expectedQuery2 = command
			expectedValues2 = values

		})

		var expectedExecutedBatch *gocql.Batch

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {

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
			So(expectedQuery1, ShouldEqual, "DELETE FROM tag WHERE ID = ?")
			So(expectedQuery2, ShouldEqual, "DELETE FROM tag WHERE ID = ?")
			So(expectedValues1, ShouldResemble, []interface{}{"123"})
			So(expectedValues2, ShouldResemble, []interface{}{"123"})
			So(expectedBatch, ShouldEqual, expectedExecutedBatch)
		})
	})
}

func TestCassandra_Delete_WithTransactionID(t *testing.T) {

	Convey("When I call the method Delete", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"123"}, nil
		})

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.T(t).Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

			if expectedValues1 == nil {
				expectedQuery1 = command
				expectedValues1 = values
				return
			}

			expectedQuery2 = command
			expectedValues2 = values

		})

		var numberOfCallOfExecuteBatch int

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {

			numberOfCallOfExecuteBatch++
			return nil

		})

		tag1 := &Tag{}
		tag1.ID = "123"

		tag2 := &Tag{}
		tag2.ID = "456"

		context := &manipulate.Context{}
		context.PageSize = 10
		context.TransactionID = "123"

		err := store.Delete(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err, ShouldBeNil)
			So(expectedQuery1, ShouldEqual, "DELETE FROM tag WHERE ID = ?")
			So(expectedQuery2, ShouldEqual, "DELETE FROM tag WHERE ID = ?")
			So(expectedValues1, ShouldResemble, []interface{}{"123"})
			So(expectedValues2, ShouldResemble, []interface{}{"123"})
			So(numberOfCallOfExecuteBatch, ShouldEqual, 0)
		})
	})
}

func TestCassandra_DeleteError(t *testing.T) {

	Convey("When I call the method Delete", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return errors.New("Error batch")
		})

		tag1 := &Tag{}
		tag1.ID = "123"

		tag2 := &Tag{}
		tag2.ID = "456"

		context := &manipulate.Context{}
		context.PageSize = 10

		errs := store.Delete(context, tag1, tag2)

		Convey("Then I should get an error", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotExecuteBatch)
		})
	})
}

func TestCassandra_DeleteWithErrorPrimaryFields(t *testing.T) {

	Convey("When I call the method Retrieve", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{}, []interface{}{}, errors.New("CassandraStore PrimaryFieldsAndValues error")
		})

		tag := &Tag{}
		tag.ID = "1234"

		context := &manipulate.Context{}
		context.PageSize = 10

		errs := store.Delete(context, tag)

		Convey("Then I should get an error", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotExractPrimaryFieldsAndValues)
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

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			expectedBatchType = t
			return expectedBatch
		})

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.T(t).Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

			if expectedValues1 == nil {
				expectedQuery1 = command
				expectedValues1 = values
				return
			}

			expectedQuery2 = command
			expectedValues2 = values

		})

		var expectedExecutedBatch *gocql.Batch

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {

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

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"456"}, nil
		})

		apomock.T(t).Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {

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
			So(expectedQuery1, ShouldEqual, "UPDATE tag SET description = ? WHERE ID = ?")
			So(expectedQuery2, ShouldEqual, "UPDATE tag SET description = ? WHERE ID = ?")
			So(expectedValues1, ShouldResemble, []interface{}{"description 1", "456"})
			So(expectedValues2, ShouldResemble, []interface{}{"description 2", "456"})
			So(expectedBatch, ShouldEqual, expectedExecutedBatch)
			So(expectedValue1, ShouldEqual, tag1)
			So(expectedValue2, ShouldEqual, tag2)
		})
	})
}

func TestCassandra_Update_WithTransactionID(t *testing.T) {

	Convey("When I call the method Update", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.T(t).Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

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

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID"}, []interface{}{"456"}, nil
		})

		apomock.T(t).Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {

			if expectedValue1 == nil {
				expectedValue1 = val
				return []string{"ID", "description"}, []interface{}{"123", "description 1"}, nil
			}

			expectedValue2 = val
			return []string{"ID", "description"}, []interface{}{"456", "description 2"}, nil
		})

		var numberOfCallOfExecuteBatch int

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {

			numberOfCallOfExecuteBatch++
			return nil

		})

		context := &manipulate.Context{}
		context.PageSize = 10
		context.Attributes = []string{"description"}
		context.TransactionID = "123"

		err := store.Update(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err, ShouldBeNil)
			So(expectedQuery1, ShouldEqual, "UPDATE tag SET description = ? WHERE ID = ?")
			So(expectedQuery2, ShouldEqual, "UPDATE tag SET description = ? WHERE ID = ?")
			So(expectedValues1, ShouldResemble, []interface{}{"description 1", "456"})
			So(expectedValues2, ShouldResemble, []interface{}{"description 2", "456"})
			So(expectedValue1, ShouldEqual, tag1)
			So(expectedValue2, ShouldEqual, tag2)
			So(numberOfCallOfExecuteBatch, ShouldEqual, 0)
		})
	})
}

func TestCassandra_Update_ErrorFieldsAndValues(t *testing.T) {

	Convey("When I call the method Update", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return nil
		})

		tag1 := &Tag{}
		tag1.ID = "123"
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.ID = "456"
		tag2.Description = "description 2"

		apomock.T(t).Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return nil, nil, errors.New("Error from FieldsAndValues")
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		errs := store.Update(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotExtractFieldsAndValues)
		})
	})
}

func TestCassandra_Update_ErrorPrimaryFieldsAndValues(t *testing.T) {

	Convey("When I call the method Update", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
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

		errs := store.Update(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotExractPrimaryFieldsAndValues)
		})
	})
}

func TestCassandra_Update_ErrorExecuteBatch(t *testing.T) {

	Convey("When I call the method Update", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return nil
		})

		tag1 := &Tag{}
		tag1.ID = "123"
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.ID = "456"
		tag2.Description = "description 2"

		apomock.T(t).Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID", "description"}, []interface{}{"123", "description 1"}, nil
		})

		apomock.T(t).Override("cassandra.PrimaryFieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{}, []interface{}{}, nil
		})

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return errors.New("Error batch")
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		errs := store.Update(context, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotExecuteBatch)
		})
	})
}

func TestCassandra_Create(t *testing.T) {

	Convey("When I call the method Create", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		var expectedBatchType gocql.BatchType
		expectedBatch := &gocql.Batch{}

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			expectedBatchType = t
			return expectedBatch
		})

		apomock.T(t).Override("gocql.TimeUUID", func() gocql.UUID {
			return ParseUUID("7c469413-12ed-11e6-ac73-f45c89941b79")
		})

		apomock.T(t).Override("gocql.UUID.String", func(uuid gocql.UUID) string {
			return StringUUID(uuid)
		})

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.T(t).Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

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

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {

			expectedExecutedBatch = b
			return nil

		})

		tag1 := &Tag{}
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.Description = "description 2"

		var expectedValue1 interface{}
		var expectedValue2 interface{}

		apomock.T(t).Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {

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
			So(expectedQuery1, ShouldEqual, "INSERT INTO tag (ID, description) VALUES (?, ?)")
			So(expectedQuery2, ShouldEqual, "INSERT INTO tag (ID, description) VALUES (?, ?)")
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

func TestCassandra_Create_WithTransacationID(t *testing.T) {

	Convey("When I call the method Create", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("gocql.TimeUUID", func() gocql.UUID {
			return ParseUUID("7c469413-12ed-11e6-ac73-f45c89941b79")
		})

		apomock.T(t).Override("gocql.UUID.String", func(uuid gocql.UUID) string {
			return StringUUID(uuid)
		})

		var expectedQuery1 string
		var expectedValues1 interface{}

		var expectedQuery2 string
		var expectedValues2 interface{}

		apomock.T(t).Override("gocql.Batch.Query", func(b *gocql.Batch, command string, values ...interface{}) {

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

		apomock.T(t).Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {

			if expectedValue1 == nil {
				expectedValue1 = val
				return []string{"ID", "description"}, []interface{}{"123", "description 1"}, nil
			}

			expectedValue2 = val
			return []string{"ID", "description"}, []interface{}{"456", "description 2"}, nil
		})

		var numberOfCallOfExecuteBatch int

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {

			numberOfCallOfExecuteBatch++
			return nil

		})

		context := &manipulate.Context{}
		context.PageSize = 10
		context.TransactionID = "123"

		err := store.Create(context, nil, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(err, ShouldBeNil)
			So(expectedQuery1, ShouldEqual, "INSERT INTO tag (ID, description) VALUES (?, ?)")
			So(expectedQuery2, ShouldEqual, "INSERT INTO tag (ID, description) VALUES (?, ?)")
			So(expectedValues1, ShouldResemble, []interface{}{"123", "description 1"})
			So(expectedValues2, ShouldResemble, []interface{}{"456", "description 2"})
			So(expectedValue1, ShouldEqual, tag1)
			So(expectedValue2, ShouldEqual, tag2)
			So(len(tag2.ID), ShouldEqual, 36)
			So(len(tag2.ID), ShouldEqual, 36)
			So(numberOfCallOfExecuteBatch, ShouldEqual, 0)
		})
	})
}

func TestCassandra_Create_ErrorFieldsAndValues(t *testing.T) {

	Convey("When I call the method Create", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return nil
		})

		apomock.T(t).Override("gocql.TimeUUID", func() gocql.UUID {
			return ParseUUID("7c469413-12ed-11e6-ac73-f45c89941b79")
		})

		apomock.T(t).Override("gocql.UUID.String", func(uuid gocql.UUID) string {
			return StringUUID(uuid)
		})

		tag1 := &Tag{}
		tag1.ID = "123"
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.ID = "456"
		tag2.Description = "description 2"

		apomock.T(t).Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return nil, nil, errors.New("Error from FieldsAndValues")
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		errs := store.Create(context, nil, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotExtractFieldsAndValues)
			So(len(tag2.ID), ShouldEqual, 0)
			So(len(tag2.ID), ShouldEqual, 0)
		})
	})
}

func TestCassandra_Create_ErrorExecuteBatch(t *testing.T) {

	Convey("When I call the method Create", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		store.nativeSession = &gocql.Session{}

		apomock.T(t).Override("gocql.Session.NewBatch", func(session *gocql.Session, t gocql.BatchType) *gocql.Batch {
			return &gocql.Batch{}
		})

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return nil
		})

		apomock.T(t).Override("gocql.TimeUUID", func() gocql.UUID {
			return ParseUUID("7c469413-12ed-11e6-ac73-f45c89941b79")
		})

		apomock.T(t).Override("gocql.UUID.String", func(uuid gocql.UUID) string {
			return StringUUID(uuid)
		})

		tag1 := &Tag{}
		tag1.ID = "123"
		tag1.Description = "description 1"

		tag2 := &Tag{}
		tag2.ID = "456"
		tag2.Description = "description 2"

		apomock.T(t).Override("cassandra.FieldsAndValues", func(val interface{}) ([]string, []interface{}, error) {
			return []string{"ID", "description"}, []interface{}{"123", "description 1"}, nil
		})

		apomock.T(t).Override("gocql.Session.ExecuteBatch", func(session *gocql.Session, b *gocql.Batch) error {
			return errors.New("Error batch")
		})

		context := &manipulate.Context{}
		context.PageSize = 10

		errs := store.Create(context, nil, tag1, tag2)

		Convey("Then everything should have been well called", func() {
			So(len(errs), ShouldEqual, 1)
			So(errs[0].Code, ShouldEqual, ErrCannotExecuteBatch)
			So(len(tag2.ID), ShouldEqual, 0)
			So(len(tag2.ID), ShouldEqual, 0)
		})
	})
}

func TestCassandraCreateKeySpace(t *testing.T) {
	Convey("When I call the method createKeySpace, a new key sapce should be created", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, nil
		})

		var expectedQueryString string
		newQuery := &gocql.Query{}

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, query string, i ...interface{}) *gocql.Query {
			expectedQueryString = query
			return newQuery
		})

		var expectedQuery *gocql.Query
		apomock.T(t).Override("gocql.Query.Exec", func(query *gocql.Query) error {
			expectedQuery = query
			return nil
		})

		var sessionCloseCalled bool

		apomock.T(t).Override("gocql.Session.Close", func(session *gocql.Session) {
			sessionCloseCalled = true
		})

		err := store.CreateKeySpace(5)

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(expectedQueryString, ShouldEqual, "CREATE KEYSPACE keyspace WITH replication = {'class' : 'SimpleStrategy', 'replication_factor': 5}")
			So(expectedQuery, ShouldEqual, newQuery)
			So(sessionCloseCalled, ShouldBeTrue)
			So(err, ShouldBeNil)
		})
	})
}

func TestCassandraCreateKeySpaceWithErrorFromSession(t *testing.T) {
	Convey("When I call the method createKeySpace and got an error from the creation of the session", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, errors.New("should panic session")
		})

		err := store.CreateKeySpace(5)

		Convey("Then I should have get a session", func() {
			So(err, ShouldResemble, errors.New("should panic session"))
		})
	})
}

func TestCassandraCreateKeySpaceErrorQuery(t *testing.T) {
	Convey("When I call the method createKeySpace, with an error from the query", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, nil
		})

		var expectedQueryString string
		newQuery := &gocql.Query{}

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, query string, i ...interface{}) *gocql.Query {
			expectedQueryString = query
			return newQuery
		})

		var expectedQuery *gocql.Query
		apomock.T(t).Override("gocql.Query.Exec", func(query *gocql.Query) error {
			expectedQuery = query
			return errors.New("query error")
		})

		var sessionCloseCalled bool

		apomock.T(t).Override("gocql.Session.Close", func(session *gocql.Session) {
			sessionCloseCalled = true
		})

		err := store.CreateKeySpace(5)

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(expectedQueryString, ShouldEqual, "CREATE KEYSPACE keyspace WITH replication = {'class' : 'SimpleStrategy', 'replication_factor': 5}")
			So(expectedQuery, ShouldEqual, newQuery)
			So(sessionCloseCalled, ShouldBeTrue)
			So(err, ShouldResemble, errors.New("query error"))
		})
	})
}

func TestCassandraDropKeySpace(t *testing.T) {
	Convey("When I call the method DropKeySpace, a new key sapce should be created", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, nil
		})

		var expectedQueryString string
		newQuery := &gocql.Query{}

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, query string, i ...interface{}) *gocql.Query {
			expectedQueryString = query
			return newQuery
		})

		var expectedQuery *gocql.Query
		apomock.T(t).Override("gocql.Query.Exec", func(query *gocql.Query) error {
			expectedQuery = query
			return nil
		})

		var sessionCloseCalled bool

		apomock.T(t).Override("gocql.Session.Close", func(session *gocql.Session) {
			sessionCloseCalled = true
		})

		err := store.DropKeySpace()

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(expectedQueryString, ShouldEqual, "DROP KEYSPACE IF EXISTS keyspace")
			So(expectedQuery, ShouldEqual, newQuery)
			So(sessionCloseCalled, ShouldBeTrue)
			So(err, ShouldBeNil)
		})
	})
}

func TestCassandraDropKeySpaceWithErrorFromSession(t *testing.T) {
	Convey("When I call the method DropKeySpace and got an error from the creation of the session", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, errors.New("should panic session")
		})

		err := store.DropKeySpace()

		Convey("Then I should have get a session", func() {
			So(err, ShouldResemble, errors.New("should panic session"))
		})
	})
}

func TestCassandraDropKeySpaceErrorQuery(t *testing.T) {
	Convey("When I call the method DropKeySpace, with an error from the query", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, nil
		})

		var expectedQueryString string
		newQuery := &gocql.Query{}

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, query string, i ...interface{}) *gocql.Query {
			expectedQueryString = query
			return newQuery
		})

		var expectedQuery *gocql.Query
		apomock.T(t).Override("gocql.Query.Exec", func(query *gocql.Query) error {
			expectedQuery = query
			return errors.New("query error")
		})

		var sessionCloseCalled bool

		apomock.T(t).Override("gocql.Session.Close", func(session *gocql.Session) {
			sessionCloseCalled = true
		})

		err := store.DropKeySpace()

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(expectedQueryString, ShouldEqual, "DROP KEYSPACE IF EXISTS keyspace")
			So(expectedQuery, ShouldEqual, newQuery)
			So(sessionCloseCalled, ShouldBeTrue)
			So(err, ShouldResemble, errors.New("query error"))
		})
	})
}

func TestCassandraDoesKeyspaceExist(t *testing.T) {
	Convey("When I call the method DoesKeyspaceExist, with a keyspace which is exist", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, nil
		})

		var sessionCloseCalled bool

		apomock.T(t).Override("gocql.Session.Close", func(session *gocql.Session) {
			sessionCloseCalled = true
		})

		keyspaceMetadata := &gocql.KeyspaceMetadata{}
		keyspaceMetadata.Tables = make(map[string]*gocql.TableMetadata)
		keyspaceMetadata.Tables["1"] = &gocql.TableMetadata{}

		apomock.T(t).Override("gocql.Session.KeyspaceMetadata", func(session *gocql.Session, keyspace string) (*gocql.KeyspaceMetadata, error) {
			return keyspaceMetadata, nil
		})

		ok, err := store.DoesKeyspaceExist()

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(sessionCloseCalled, ShouldBeTrue)
			So(err, ShouldBeNil)
			So(ok, ShouldBeTrue)
		})
	})
}

func TestCassandraDoesKeyspaceExistWithNoKeyspace(t *testing.T) {
	Convey("When I call the method DoesKeyspaceExist, with a keyspace which is not exist", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, nil
		})

		var sessionCloseCalled bool

		apomock.T(t).Override("gocql.Session.Close", func(session *gocql.Session) {
			sessionCloseCalled = true
		})

		keyspaceMetadata := &gocql.KeyspaceMetadata{}

		apomock.T(t).Override("gocql.Session.KeyspaceMetadata", func(session *gocql.Session, keyspace string) (*gocql.KeyspaceMetadata, error) {
			return keyspaceMetadata, nil
		})

		ok, err := store.DoesKeyspaceExist()

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(sessionCloseCalled, ShouldBeTrue)
			So(err, ShouldBeNil)
			So(ok, ShouldBeFalse)
		})
	})
}

func TestCassandraDoesKeyspaceExistWithErrorFromKeySpace(t *testing.T) {
	Convey("When I call the method DoesKeyspaceExist, with a keyspace which an error", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, nil
		})

		var sessionCloseCalled bool

		apomock.T(t).Override("gocql.Session.Close", func(session *gocql.Session) {
			sessionCloseCalled = true
		})

		apomock.T(t).Override("gocql.Session.KeyspaceMetadata", func(session *gocql.Session, keyspace string) (*gocql.KeyspaceMetadata, error) {
			return nil, errors.New("error from KeyspaceMetadata")
		})

		ok, err := store.DoesKeyspaceExist()

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(sessionCloseCalled, ShouldBeTrue)
			So(err, ShouldResemble, errors.New("error from KeyspaceMetadata"))
			So(ok, ShouldBeFalse)
		})
	})
}

func TestCassandraDoesKeyspaceExistWithErrorFromSessio(t *testing.T) {
	Convey("When I call the method DoesKeyspaceExist, with an error from the session", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return nil, errors.New("should panic session")
		})

		ok, err := store.DoesKeyspaceExist()

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(err, ShouldResemble, errors.New("should panic session"))
			So(ok, ShouldBeFalse)
		})
	})
}

func TestCassandraExecuteScript(t *testing.T) {
	Convey("When I call the method ExecuteScript, with a good script", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, nil
		})

		var expectedQueryString1 string
		var expectedQueryString2 string
		newQuery := &gocql.Query{}

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, query string, i ...interface{}) *gocql.Query {

			if expectedQueryString1 == "" {
				expectedQueryString1 = query
			} else {
				expectedQueryString2 = query
			}

			return newQuery
		})

		var expectedQuery *gocql.Query
		apomock.T(t).Override("gocql.Query.Exec", func(query *gocql.Query) error {
			expectedQuery = query
			return nil
		})

		var sessionCloseCalled bool

		apomock.T(t).Override("gocql.Session.Close", func(session *gocql.Session) {
			sessionCloseCalled = true
		})

		err := store.ExecuteScript("Alexandre;\nAntoine")

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "keyspace")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(sessionCloseCalled, ShouldBeTrue)
			So(expectedQueryString1, ShouldResemble, "Alexandre")
			So(expectedQueryString2, ShouldResemble, "Antoine")
			So(expectedQuery, ShouldEqual, newQuery)
			So(err, ShouldBeNil)
		})
	})
}

func TestCassandraExecuteScriptWithError(t *testing.T) {
	Convey("When I call the method ExecuteScript, with a error", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, nil
		})

		var expectedQueryString1 string
		var expectedQueryString2 string
		newQuery := &gocql.Query{}

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, query string, i ...interface{}) *gocql.Query {

			if expectedQueryString1 == "" {
				expectedQueryString1 = query
			} else {
				expectedQueryString2 = query
			}

			return newQuery
		})

		var expectedQuery *gocql.Query
		apomock.T(t).Override("gocql.Query.Exec", func(query *gocql.Query) error {
			expectedQuery = query
			return errors.New("should error exec")
		})

		var sessionCloseCalled bool

		apomock.T(t).Override("gocql.Session.Close", func(session *gocql.Session) {
			sessionCloseCalled = true
		})

		err := store.ExecuteScript("Alexandre;\nAntoine")

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "keyspace")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(sessionCloseCalled, ShouldBeTrue)
			So(expectedQueryString1, ShouldResemble, "Alexandre")
			So(expectedQueryString2, ShouldResemble, "")
			So(expectedQuery, ShouldEqual, newQuery)
			So(err, ShouldResemble, errors.New("should error exec"))
		})
	})
}

func TestCassandraExecuteScriptWithErrorSession(t *testing.T) {
	Convey("When I call the method ExecuteScript, with a error session", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, errors.New("should error panic")
		})

		err := store.ExecuteScript("Alexandre;\nAntoine")

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "keyspace")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(err, ShouldResemble, errors.New("should error panic"))
		})
	})
}

func TestCassandraExecuteScriptWithEmptyLine(t *testing.T) {
	Convey("When I call the method ExecuteScript, with a good script", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)

		var expectedServers []string
		cluster := &gocql.ClusterConfig{}

		apomock.T(t).Override("gocql.NewCluster", func(servers ...string) *gocql.ClusterConfig {
			expectedServers = servers
			return cluster
		})

		session := &gocql.Session{}

		apomock.T(t).Override("gocql.ClusterConfig.CreateSession", func(c *gocql.ClusterConfig) (*gocql.Session, error) {
			return session, nil
		})

		var expectedQueryString1 string
		var expectedQueryString2 string
		newQuery := &gocql.Query{}

		apomock.T(t).Override("gocql.Session.Query", func(session *gocql.Session, query string, i ...interface{}) *gocql.Query {

			if expectedQueryString1 == "" {
				expectedQueryString1 = query
			} else {
				expectedQueryString2 = query
			}

			return newQuery
		})

		var expectedQuery *gocql.Query
		apomock.T(t).Override("gocql.Query.Exec", func(query *gocql.Query) error {
			expectedQuery = query
			return nil
		})

		var sessionCloseCalled bool

		apomock.T(t).Override("gocql.Session.Close", func(session *gocql.Session) {
			sessionCloseCalled = true
		})

		err := store.ExecuteScript("Alexandre;\n;\nAntoine")

		Convey("Then I should have a native session", func() {
			So(store.Servers, ShouldResemble, expectedServers)
			So(cluster.Keyspace, ShouldEqual, "keyspace")
			So(cluster.Consistency, ShouldEqual, gocql.Quorum)
			So(cluster.ProtoVersion, ShouldEqual, 1)
			So(sessionCloseCalled, ShouldBeTrue)
			So(expectedQueryString1, ShouldResemble, "Alexandre")
			So(expectedQueryString2, ShouldResemble, "Antoine")
			So(expectedQuery, ShouldEqual, newQuery)
			So(err, ShouldBeNil)
		})
	})
}

func TestCassandraAbort(t *testing.T) {

	Convey("Given I have a store with a transaction", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		tid := manipulate.TransactionID("tid")

		store.batchRegistry[tid] = &gocql.Batch{}

		Convey("When I use Abort", func() {

			ok := store.Abort(tid)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("Then the transactionID should have been removed", func() {
				So(store.batchRegistry, ShouldNotContainKey, tid)
			})
		})
	})

	Convey("Given I have a store with no transaction", t, func() {

		store := NewCassandraStore([]string{"1.2.3.4", "1.2.3.5"}, "keyspace", 1)
		tid := manipulate.TransactionID("tid")

		Convey("When I use Abort", func() {

			ok := store.Abort(tid)

			Convey("Then ok should not be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}
