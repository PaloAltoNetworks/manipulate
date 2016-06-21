// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

// ManipCassandraUnmarshalErrorTitle represents the title...
const (
	ManipCassandraUnmarshalErrorTitle                          = `manipcassandra: unmarshalInterface - cassandra.Unmarshal`
	ManipCassandraUnmarshalNumberObjectsAndSliceMapsErrorTitle = `manipcassandra: unmarshalManipulable - number of objects different to number of maps`
	ManipCassandraIteratorCloseErrorTitle                      = `manipcassandra: sliceMaps - issue with iterator close`
	ManipCassandraIteratorSliceMapErrorTitle                   = `manipcassandra: sliceMaps - issue with iterator sliceMap`
	ManipCassandraExecuteBatchErrorTitle                       = `manipcassandra: executeBatch of the batch`
	ManipCassandraIteratorScanErrorTitle                       = `manipcassandra: count - issue with iterator scan`
	ManipCassandraFieldsAndValuesErrorTitle                    = `manipcassandra: fieldAndValues`
	ManipCassandraPrimaryFieldsAndValuesErrorTitle             = `manipcassandra: primaryFieldAndValues`
	ManipCassandraQueryErrorTitle                              = `manipcassandra: session querry issue`
)

// ManipCassandraUnmarshalErrorDescription represents the description...
const (
	ManipCassandraUnmarshalErrorDescription                          = `Error when calling the method cassandra.Unmarshal with the objects %s and sliceMaps %s. Go error %s`
	ManipCassandraUnmarshalNumberObjectsAndSliceMapsErrorDescription = `The number of given objects %s and number of maps %s is different`
	ManipCassandraIteratorCloseErrorDescription                      = `An issue occurs when closing the iterator. Go error %s`
	ManipCassandraIteratorSliceMapErrorDescription                   = `An issue occurs when calling sliceMap of the iterator. Go error %s`
	ManipCassandraExecuteBatchErrorDescription                       = `An issue occurs when calling executeBatch of the batch. Go error %s`
	ManipCassandraIteratorScanErrorDescription                       = `An issue occurs when scanning the iterator.`
	ManipCassandraFieldsAndValuesErrorDescription                    = `An issue occurs when using the method fieldAndValues on the object %s. Go error %s`
	ManipCassandraPrimaryFieldsAndValuesErrorDescription             = `An issue occurs when using the method primaryFieldAndValues on the object %s. Go error %s`
	ManipCassandraQueryErrorDescription                              = `An issue occurs when executiong a query with the object %s. Go error %s`
)

// ManipCassandraUnmarshalErrorCode represents the code...
const (
	ManipCassandraUnmarshalErrorCode                          = 5000
	ManipCassandraUnmarshalNumberObjectsAndSliceMapsErrorCode = 5001
	ManipCassandraIteratorCloseErrorCode                      = 5002
	ManipCassandraIteratorSliceMapErrorCode                   = 5003
	ManipCassandraExecuteBatchErrorCode                       = 5004
	ManipCassandraIteratorScanErrorCode                       = 5005
	ManipCassandraFieldsAndValuesErrorCode                    = 5006
	ManipCassandraPrimaryFieldsAndValuesErrorCode             = 5007
	ManipCassandraQueryErrorCode                              = 5008
)
