// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"
import "time"

func init() {
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructQueryExecutor, apomockNewStructQueryExecutor)

	apomock.RegisterFunc("gocql", "gocql.queryExecutor.executeQuery", (*queryExecutor).AuxMockexecuteQuery)
}

const (
	ApomockStructQueryExecutor = 38
)

//
// Internal Types: in this package and their exportable versions
//
type queryExecutor struct {
	pool   *policyConnPool
	policy HostSelectionPolicy
}

//
// External Types: in this package
//
type ExecutableQuery interface {
	execute(conn *Conn) *Iter
	attempt(time.Duration)
	retryPolicy() RetryPolicy
	GetRoutingKey() ([]byte, error)
	RetryableQuery
}

func apomockNewStructQueryExecutor() interface{} { return &queryExecutor{} }

//
// Mock: (recvq *queryExecutor)executeQuery(argqry ExecutableQuery)(reta *Iter, retb error)
//

type MockArgsTypequeryExecutorexecuteQuery struct {
	ApomockCallNumber int
	Argqry            ExecutableQuery
}

var LastMockArgsqueryExecutorexecuteQuery MockArgsTypequeryExecutorexecuteQuery

// (recvq *queryExecutor)AuxMockexecuteQuery(argqry ExecutableQuery)(reta *Iter, retb error) - Generated mock function
func (recvq *queryExecutor) AuxMockexecuteQuery(argqry ExecutableQuery) (reta *Iter, retb error) {
	LastMockArgsqueryExecutorexecuteQuery = MockArgsTypequeryExecutorexecuteQuery{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrqueryExecutorexecuteQuery(),
		Argqry:            argqry,
	}
	rargs, rerr := apomock.GetNext("gocql.queryExecutor.executeQuery")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.queryExecutor.executeQuery")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.queryExecutor.executeQuery")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Iter)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrqueryExecutorexecuteQuery  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrqueryExecutorexecuteQuery int = 0

var condRecorderAuxMockPtrqueryExecutorexecuteQuery *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrqueryExecutorexecuteQuery(i int) {
	condRecorderAuxMockPtrqueryExecutorexecuteQuery.L.Lock()
	for recorderAuxMockPtrqueryExecutorexecuteQuery < i {
		condRecorderAuxMockPtrqueryExecutorexecuteQuery.Wait()
	}
	condRecorderAuxMockPtrqueryExecutorexecuteQuery.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrqueryExecutorexecuteQuery() {
	condRecorderAuxMockPtrqueryExecutorexecuteQuery.L.Lock()
	recorderAuxMockPtrqueryExecutorexecuteQuery++
	condRecorderAuxMockPtrqueryExecutorexecuteQuery.L.Unlock()
	condRecorderAuxMockPtrqueryExecutorexecuteQuery.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrqueryExecutorexecuteQuery() (ret int) {
	condRecorderAuxMockPtrqueryExecutorexecuteQuery.L.Lock()
	ret = recorderAuxMockPtrqueryExecutorexecuteQuery
	condRecorderAuxMockPtrqueryExecutorexecuteQuery.L.Unlock()
	return
}

// (recvq *queryExecutor)executeQuery - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *queryExecutor) executeQuery(argqry ExecutableQuery) (reta *Iter, retb error) {
	FuncAuxMockPtrqueryExecutorexecuteQuery, ok := apomock.GetRegisteredFunc("gocql.queryExecutor.executeQuery")
	if ok {
		reta, retb = FuncAuxMockPtrqueryExecutorexecuteQuery.(func(recvq *queryExecutor, argqry ExecutableQuery) (reta *Iter, retb error))(recvq, argqry)
	} else {
		panic("FuncAuxMockPtrqueryExecutorexecuteQuery ")
	}
	AuxMockIncrementRecorderAuxMockPtrqueryExecutorexecuteQuery()
	return
}
