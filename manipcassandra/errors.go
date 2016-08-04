// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

const (
	// ManipCassandraDatabaseError represents the an internal db error.
	ManipCassandraDatabaseError = "Database Manipulation Error"
)

// ManipCassandraUnmarshalErrorCode represents the code...
const (
	ManipCassandraUnmarshalErrorCode                          = 5000
	ManipCassandraUnmarshalNumberObjectsAndSliceMapsErrorCode = 5001
	ManipCassandraIteratorCloseErrorCode                      = 5002
	ManipCassandraExecuteBatchErrorCode                       = 5004
	ManipCassandraIteratorScanErrorCode                       = 5005
	ManipCassandraFieldsAndValuesErrorCode                    = 5006
	ManipCassandraPrimaryFieldsAndValuesErrorCode             = 5007
	ManipCassandraQueryErrorCode                              = 5008
	ManipCassandraCommitTransactionErrorCode                  = 5009
)
