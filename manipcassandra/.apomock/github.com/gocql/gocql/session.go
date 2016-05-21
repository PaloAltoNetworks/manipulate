// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

import "golang.org/x/net/context"

import "io"

import "github.com/gocql/gocql/apointernal/lru"
import "errors"
import "time"

func init() {
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructRoutingKeyInfo, apomockNewStructRoutingKeyInfo)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructTraceWriter, apomockNewStructTraceWriter)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructRoutingKeyInfoLRU, apomockNewStructRoutingKeyInfoLRU)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructNextIter, apomockNewStructNextIter)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructInflightCachedEntry, apomockNewStructInflightCachedEntry)

	apomock.RegisterFunc("gocql", "gocql.Session.routingKeyInfo", (*Session).AuxMockroutingKeyInfo)
	apomock.RegisterFunc("gocql", "gocql.Query.PageSize", (*Query).AuxMockPageSize)
	apomock.RegisterFunc("gocql", "gocql.Session.NewBatch", (*Session).AuxMockNewBatch)
	apomock.RegisterFunc("gocql", "gocql.routingKeyInfoLRU.Remove", (*routingKeyInfoLRU).AuxMockRemove)
	apomock.RegisterFunc("gocql", "gocql.Session.getConn", (*Session).AuxMockgetConn)
	apomock.RegisterFunc("gocql", "gocql.Iter.GetCustomPayload", (*Iter).AuxMockGetCustomPayload)
	apomock.RegisterFunc("gocql", "gocql.Batch.Attempts", (*Batch).AuxMockAttempts)
	apomock.RegisterFunc("gocql", "gocql.Batch.GetRoutingKey", (*Batch).AuxMockGetRoutingKey)
	apomock.RegisterFunc("gocql", "gocql.Query.RetryPolicy", (*Query).AuxMockRetryPolicy)
	apomock.RegisterFunc("gocql", "gocql.Session.Bind", (*Session).AuxMockBind)
	apomock.RegisterFunc("gocql", "gocql.Query.Attempts", (*Query).AuxMockAttempts)
	apomock.RegisterFunc("gocql", "gocql.Query.WithContext", (*Query).AuxMockWithContext)
	apomock.RegisterFunc("gocql", "gocql.Query.execute", (*Query).AuxMockexecute)
	apomock.RegisterFunc("gocql", "gocql.Session.reconnectDownedHosts", (*Session).AuxMockreconnectDownedHosts)
	apomock.RegisterFunc("gocql", "gocql.Query.GetRoutingKey", (*Query).AuxMockGetRoutingKey)
	apomock.RegisterFunc("gocql", "gocql.Query.SerialConsistency", (*Query).AuxMockSerialConsistency)
	apomock.RegisterFunc("gocql", "gocql.Batch.RetryPolicy", (*Batch).AuxMockRetryPolicy)
	apomock.RegisterFunc("gocql", "gocql.routingKeyInfoLRU.Max", (*routingKeyInfoLRU).AuxMockMax)
	apomock.RegisterFunc("gocql", "gocql.Query.Consistency", (*Query).AuxMockConsistency)
	apomock.RegisterFunc("gocql", "gocql.Query.Exec", (*Query).AuxMockExec)
	apomock.RegisterFunc("gocql", "gocql.nextIter.fetch", (*nextIter).AuxMockfetch)
	apomock.RegisterFunc("gocql", "gocql.Query.attempt", (*Query).AuxMockattempt)
	apomock.RegisterFunc("gocql", "gocql.Query.Prefetch", (*Query).AuxMockPrefetch)
	apomock.RegisterFunc("gocql", "gocql.Query.ScanCAS", (*Query).AuxMockScanCAS)
	apomock.RegisterFunc("gocql", "gocql.Batch.DefaultTimestamp", (*Batch).AuxMockDefaultTimestamp)
	apomock.RegisterFunc("gocql", "gocql.NewErrProtocol", AuxMockNewErrProtocol)
	apomock.RegisterFunc("gocql", "gocql.Session.MapExecuteBatchCAS", (*Session).AuxMockMapExecuteBatchCAS)
	apomock.RegisterFunc("gocql", "gocql.Query.GetConsistency", (*Query).AuxMockGetConsistency)
	apomock.RegisterFunc("gocql", "gocql.Query.NoSkipMetadata", (*Query).AuxMockNoSkipMetadata)
	apomock.RegisterFunc("gocql", "gocql.Query.Iter", (*Query).AuxMockIter)
	apomock.RegisterFunc("gocql", "gocql.Iter.Columns", (*Iter).AuxMockColumns)
	apomock.RegisterFunc("gocql", "gocql.Iter.WillSwitchPage", (*Iter).AuxMockWillSwitchPage)
	apomock.RegisterFunc("gocql", "gocql.Batch.GetConsistency", (*Batch).AuxMockGetConsistency)
	apomock.RegisterFunc("gocql", "gocql.Batch.Query", (*Batch).AuxMockQuery)
	apomock.RegisterFunc("gocql", "gocql.Session.executeBatch", (*Session).AuxMockexecuteBatch)
	apomock.RegisterFunc("gocql", "gocql.Session.Closed", (*Session).AuxMockClosed)
	apomock.RegisterFunc("gocql", "gocql.Session.KeyspaceMetadata", (*Session).AuxMockKeyspaceMetadata)
	apomock.RegisterFunc("gocql", "gocql.Query.DefaultTimestamp", (*Query).AuxMockDefaultTimestamp)
	apomock.RegisterFunc("gocql", "gocql.Query.Bind", (*Query).AuxMockBind)
	apomock.RegisterFunc("gocql", "gocql.Batch.Latency", (*Batch).AuxMockLatency)
	apomock.RegisterFunc("gocql", "gocql.Batch.Bind", (*Batch).AuxMockBind)
	apomock.RegisterFunc("gocql", "gocql.Session.SetPrefetch", (*Session).AuxMockSetPrefetch)
	apomock.RegisterFunc("gocql", "gocql.Session.ExecuteBatchCAS", (*Session).AuxMockExecuteBatchCAS)
	apomock.RegisterFunc("gocql", "gocql.Query.Trace", (*Query).AuxMockTrace)
	apomock.RegisterFunc("gocql", "gocql.isUseStatement", AuxMockisUseStatement)
	apomock.RegisterFunc("gocql", "gocql.Iter.PageState", (*Iter).AuxMockPageState)
	apomock.RegisterFunc("gocql", "gocql.ColumnInfo.String", (ColumnInfo).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.Error.Error", (Error).AuxMockError)
	apomock.RegisterFunc("gocql", "gocql.Session.SetPageSize", (*Session).AuxMockSetPageSize)
	apomock.RegisterFunc("gocql", "gocql.Query.WithTimestamp", (*Query).AuxMockWithTimestamp)
	apomock.RegisterFunc("gocql", "gocql.Query.retryPolicy", (*Query).AuxMockretryPolicy)
	apomock.RegisterFunc("gocql", "gocql.Query.MapScanCAS", (*Query).AuxMockMapScanCAS)
	apomock.RegisterFunc("gocql", "gocql.Iter.NumRows", (*Iter).AuxMockNumRows)
	apomock.RegisterFunc("gocql", "gocql.routingKeyInfo.String", (*routingKeyInfo).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.Session.Query", (*Session).AuxMockQuery)
	apomock.RegisterFunc("gocql", "gocql.Session.connect", (*Session).AuxMockconnect)
	apomock.RegisterFunc("gocql", "gocql.Query.Scan", (*Query).AuxMockScan)
	apomock.RegisterFunc("gocql", "gocql.Batch.retryPolicy", (*Batch).AuxMockretryPolicy)
	apomock.RegisterFunc("gocql", "gocql.traceWriter.Trace", (*traceWriter).AuxMockTrace)
	apomock.RegisterFunc("gocql", "gocql.Session.Close", (*Session).AuxMockClose)
	apomock.RegisterFunc("gocql", "gocql.Iter.Host", (*Iter).AuxMockHost)
	apomock.RegisterFunc("gocql", "gocql.Iter.readColumn", (*Iter).AuxMockreadColumn)
	apomock.RegisterFunc("gocql", "gocql.Query.String", (Query).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.Session.ExecuteBatch", (*Session).AuxMockExecuteBatch)
	apomock.RegisterFunc("gocql", "gocql.NewBatch", AuxMockNewBatch)
	apomock.RegisterFunc("gocql", "gocql.Batch.attempt", (*Batch).AuxMockattempt)
	apomock.RegisterFunc("gocql", "gocql.Session.SetTrace", (*Session).AuxMockSetTrace)
	apomock.RegisterFunc("gocql", "gocql.Session.SetConsistency", (*Session).AuxMockSetConsistency)
	apomock.RegisterFunc("gocql", "gocql.Batch.execute", (*Batch).AuxMockexecute)
	apomock.RegisterFunc("gocql", "gocql.Query.Latency", (*Query).AuxMockLatency)
	apomock.RegisterFunc("gocql", "gocql.Query.RoutingKey", (*Query).AuxMockRoutingKey)
	apomock.RegisterFunc("gocql", "gocql.Query.shouldPrepare", (*Query).AuxMockshouldPrepare)
	apomock.RegisterFunc("gocql", "gocql.Iter.Close", (*Iter).AuxMockClose)
	apomock.RegisterFunc("gocql", "gocql.Batch.Size", (*Batch).AuxMockSize)
	apomock.RegisterFunc("gocql", "gocql.addrsToHosts", AuxMockaddrsToHosts)
	apomock.RegisterFunc("gocql", "gocql.NewTraceWriter", AuxMockNewTraceWriter)
	apomock.RegisterFunc("gocql", "gocql.Batch.SerialConsistency", (*Batch).AuxMockSerialConsistency)
	apomock.RegisterFunc("gocql", "gocql.Query.PageState", (*Query).AuxMockPageState)
	apomock.RegisterFunc("gocql", "gocql.Iter.checkErrAndNotFound", (*Iter).AuxMockcheckErrAndNotFound)
	apomock.RegisterFunc("gocql", "gocql.NewSession", AuxMockNewSession)
	apomock.RegisterFunc("gocql", "gocql.Query.MapScan", (*Query).AuxMockMapScan)
	apomock.RegisterFunc("gocql", "gocql.Iter.Scan", (*Iter).AuxMockScan)
	apomock.RegisterFunc("gocql", "gocql.Batch.WithContext", (*Batch).AuxMockWithContext)
	apomock.RegisterFunc("gocql", "gocql.Session.executeQuery", (*Session).AuxMockexecuteQuery)
}

const (
	LoggedBatch   BatchType = 0
	UnloggedBatch BatchType = 1
	CounterBatch  BatchType = 2
)

const BatchSizeMaximum = 65535

const (
	ApomockStructRoutingKeyInfo      = 10
	ApomockStructTraceWriter         = 11
	ApomockStructRoutingKeyInfoLRU   = 12
	ApomockStructNextIter            = 13
	ApomockStructInflightCachedEntry = 14
)

var (
	ErrNotFound      = errors.New("not found")
	ErrUnavailable   = errors.New("unavailable")
	ErrUnsupported   = errors.New("feature not supported")
	ErrTooManyStmts  = errors.New("too many statements")
	ErrUseStmt       = errors.New("use statements aren't supported. Please see https://github.com/gocql/gocql for explaination.")
	ErrSessionClosed = errors.New("session has been closed")
	ErrNoConnections = errors.New("qocql: no hosts available in the pool")
	ErrNoKeyspace    = errors.New("no keyspace provided")
	ErrNoMetadata    = errors.New("no metadata available")
)

//
// Internal Types: in this package and their exportable versions
//
type routingKeyInfo struct {
	indexes []int
	types   []TypeInfo
}
type traceWriter struct {
	session *Session
	w       io.Writer
	mu      sync.Mutex
}
type routingKeyInfoLRU struct {
	lru *lru.Cache
	mu  sync.Mutex
}
type nextIter struct {
	qry  Query
	pos  int
	once sync.Once
	next *Iter
}
type inflightCachedEntry struct {
	wg    sync.WaitGroup
	err   error
	value interface{}
}

//
// External Types: in this package
//
type ErrProtocol struct{ error }

type QueryInfo struct {
	Id          []byte
	Args        []ColumnInfo
	Rval        []ColumnInfo
	PKeyColumns []int
}

type Tracer interface {
	Trace(traceId []byte)
}

type Session struct {
	cons                Consistency
	pageSize            int
	prefetch            float64
	routingKeyInfoCache routingKeyInfoLRU
	schemaDescriber     *schemaDescriber
	trace               Tracer
	hostSource          *ringDescriber
	stmtsLRU            *preparedLRU
	connCfg             *ConnConfig
	executor            *queryExecutor
	pool                *policyConnPool
	policy              HostSelectionPolicy
	ring                ring
	metadata            clusterMetadata
	mu                  sync.RWMutex
	control             *controlConn
	nodeEvents          *eventDeouncer
	schemaEvents        *eventDeouncer
	hosts               []HostInfo
	useSystemSchema     bool
	cfg                 ClusterConfig
	closeMu             sync.RWMutex
	isClosed            bool
}

type Query struct {
	stmt                  string
	values                []interface{}
	cons                  Consistency
	pageSize              int
	routingKey            []byte
	routingKeyBuffer      []byte
	pageState             []byte
	prefetch              float64
	trace                 Tracer
	session               *Session
	rt                    RetryPolicy
	binding               func(q *QueryInfo) ([]interface{}, error)
	attempts              int
	totalLatency          int64
	serialCons            SerialConsistency
	defaultTimestamp      bool
	defaultTimestampValue int64
	disableSkipMetadata   bool
	context               context.Context
	disableAutoPage       bool
}

type BatchType byte

type Iter struct {
	err     error
	pos     int
	meta    resultMetadata
	numRows int
	next    *nextIter
	host    *HostInfo
	framer  *framer
	closed  int32
}

type Batch struct {
	Type             BatchType
	Entries          []BatchEntry
	Cons             Consistency
	rt               RetryPolicy
	attempts         int
	totalLatency     int64
	serialCons       SerialConsistency
	defaultTimestamp bool
	context          context.Context
}

type BatchEntry struct {
	Stmt    string
	Args    []interface{}
	binding func(q *QueryInfo) ([]interface{}, error)
}

type ColumnInfo struct {
	Keyspace string
	Table    string
	Name     string
	TypeInfo TypeInfo
}

type Error struct {
	Code    int
	Message string
}

func apomockNewStructRoutingKeyInfo() interface{}      { return &routingKeyInfo{} }
func apomockNewStructTraceWriter() interface{}         { return &traceWriter{} }
func apomockNewStructRoutingKeyInfoLRU() interface{}   { return &routingKeyInfoLRU{} }
func apomockNewStructNextIter() interface{}            { return &nextIter{} }
func apomockNewStructInflightCachedEntry() interface{} { return &inflightCachedEntry{} }

//
// Mock: (recvs *Session)routingKeyInfo(argctx context.Context, argstmt string)(reta *routingKeyInfo, retb error)
//

type MockArgsTypeSessionroutingKeyInfo struct {
	ApomockCallNumber int
	Argctx            context.Context
	Argstmt           string
}

var LastMockArgsSessionroutingKeyInfo MockArgsTypeSessionroutingKeyInfo

// (recvs *Session)AuxMockroutingKeyInfo(argctx context.Context, argstmt string)(reta *routingKeyInfo, retb error) - Generated mock function
func (recvs *Session) AuxMockroutingKeyInfo(argctx context.Context, argstmt string) (reta *routingKeyInfo, retb error) {
	LastMockArgsSessionroutingKeyInfo = MockArgsTypeSessionroutingKeyInfo{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionroutingKeyInfo(),
		Argctx:            argctx,
		Argstmt:           argstmt,
	}
	rargs, rerr := apomock.GetNext("gocql.Session.routingKeyInfo")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.routingKeyInfo")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Session.routingKeyInfo")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*routingKeyInfo)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrSessionroutingKeyInfo  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionroutingKeyInfo int = 0

var condRecorderAuxMockPtrSessionroutingKeyInfo *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionroutingKeyInfo(i int) {
	condRecorderAuxMockPtrSessionroutingKeyInfo.L.Lock()
	for recorderAuxMockPtrSessionroutingKeyInfo < i {
		condRecorderAuxMockPtrSessionroutingKeyInfo.Wait()
	}
	condRecorderAuxMockPtrSessionroutingKeyInfo.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionroutingKeyInfo() {
	condRecorderAuxMockPtrSessionroutingKeyInfo.L.Lock()
	recorderAuxMockPtrSessionroutingKeyInfo++
	condRecorderAuxMockPtrSessionroutingKeyInfo.L.Unlock()
	condRecorderAuxMockPtrSessionroutingKeyInfo.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionroutingKeyInfo() (ret int) {
	condRecorderAuxMockPtrSessionroutingKeyInfo.L.Lock()
	ret = recorderAuxMockPtrSessionroutingKeyInfo
	condRecorderAuxMockPtrSessionroutingKeyInfo.L.Unlock()
	return
}

// (recvs *Session)routingKeyInfo - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) routingKeyInfo(argctx context.Context, argstmt string) (reta *routingKeyInfo, retb error) {
	FuncAuxMockPtrSessionroutingKeyInfo, ok := apomock.GetRegisteredFunc("gocql.Session.routingKeyInfo")
	if ok {
		reta, retb = FuncAuxMockPtrSessionroutingKeyInfo.(func(recvs *Session, argctx context.Context, argstmt string) (reta *routingKeyInfo, retb error))(recvs, argctx, argstmt)
	} else {
		panic("FuncAuxMockPtrSessionroutingKeyInfo ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionroutingKeyInfo()
	return
}

//
// Mock: (recvq *Query)PageSize(argn int)(reta *Query)
//

type MockArgsTypeQueryPageSize struct {
	ApomockCallNumber int
	Argn              int
}

var LastMockArgsQueryPageSize MockArgsTypeQueryPageSize

// (recvq *Query)AuxMockPageSize(argn int)(reta *Query) - Generated mock function
func (recvq *Query) AuxMockPageSize(argn int) (reta *Query) {
	LastMockArgsQueryPageSize = MockArgsTypeQueryPageSize{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryPageSize(),
		Argn:              argn,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.PageSize")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.PageSize")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.PageSize")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryPageSize  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryPageSize int = 0

var condRecorderAuxMockPtrQueryPageSize *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryPageSize(i int) {
	condRecorderAuxMockPtrQueryPageSize.L.Lock()
	for recorderAuxMockPtrQueryPageSize < i {
		condRecorderAuxMockPtrQueryPageSize.Wait()
	}
	condRecorderAuxMockPtrQueryPageSize.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryPageSize() {
	condRecorderAuxMockPtrQueryPageSize.L.Lock()
	recorderAuxMockPtrQueryPageSize++
	condRecorderAuxMockPtrQueryPageSize.L.Unlock()
	condRecorderAuxMockPtrQueryPageSize.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryPageSize() (ret int) {
	condRecorderAuxMockPtrQueryPageSize.L.Lock()
	ret = recorderAuxMockPtrQueryPageSize
	condRecorderAuxMockPtrQueryPageSize.L.Unlock()
	return
}

// (recvq *Query)PageSize - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) PageSize(argn int) (reta *Query) {
	FuncAuxMockPtrQueryPageSize, ok := apomock.GetRegisteredFunc("gocql.Query.PageSize")
	if ok {
		reta = FuncAuxMockPtrQueryPageSize.(func(recvq *Query, argn int) (reta *Query))(recvq, argn)
	} else {
		panic("FuncAuxMockPtrQueryPageSize ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryPageSize()
	return
}

//
// Mock: (recvs *Session)NewBatch(argtyp BatchType)(reta *Batch)
//

type MockArgsTypeSessionNewBatch struct {
	ApomockCallNumber int
	Argtyp            BatchType
}

var LastMockArgsSessionNewBatch MockArgsTypeSessionNewBatch

// (recvs *Session)AuxMockNewBatch(argtyp BatchType)(reta *Batch) - Generated mock function
func (recvs *Session) AuxMockNewBatch(argtyp BatchType) (reta *Batch) {
	LastMockArgsSessionNewBatch = MockArgsTypeSessionNewBatch{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionNewBatch(),
		Argtyp:            argtyp,
	}
	rargs, rerr := apomock.GetNext("gocql.Session.NewBatch")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.NewBatch")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Session.NewBatch")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Batch)
	}
	return
}

// RecorderAuxMockPtrSessionNewBatch  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionNewBatch int = 0

var condRecorderAuxMockPtrSessionNewBatch *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionNewBatch(i int) {
	condRecorderAuxMockPtrSessionNewBatch.L.Lock()
	for recorderAuxMockPtrSessionNewBatch < i {
		condRecorderAuxMockPtrSessionNewBatch.Wait()
	}
	condRecorderAuxMockPtrSessionNewBatch.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionNewBatch() {
	condRecorderAuxMockPtrSessionNewBatch.L.Lock()
	recorderAuxMockPtrSessionNewBatch++
	condRecorderAuxMockPtrSessionNewBatch.L.Unlock()
	condRecorderAuxMockPtrSessionNewBatch.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionNewBatch() (ret int) {
	condRecorderAuxMockPtrSessionNewBatch.L.Lock()
	ret = recorderAuxMockPtrSessionNewBatch
	condRecorderAuxMockPtrSessionNewBatch.L.Unlock()
	return
}

// (recvs *Session)NewBatch - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) NewBatch(argtyp BatchType) (reta *Batch) {
	FuncAuxMockPtrSessionNewBatch, ok := apomock.GetRegisteredFunc("gocql.Session.NewBatch")
	if ok {
		reta = FuncAuxMockPtrSessionNewBatch.(func(recvs *Session, argtyp BatchType) (reta *Batch))(recvs, argtyp)
	} else {
		panic("FuncAuxMockPtrSessionNewBatch ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionNewBatch()
	return
}

//
// Mock: (recvr *routingKeyInfoLRU)Remove(argkey string)()
//

type MockArgsTyperoutingKeyInfoLRURemove struct {
	ApomockCallNumber int
	Argkey            string
}

var LastMockArgsroutingKeyInfoLRURemove MockArgsTyperoutingKeyInfoLRURemove

// (recvr *routingKeyInfoLRU)AuxMockRemove(argkey string)() - Generated mock function
func (recvr *routingKeyInfoLRU) AuxMockRemove(argkey string) {
	LastMockArgsroutingKeyInfoLRURemove = MockArgsTyperoutingKeyInfoLRURemove{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrroutingKeyInfoLRURemove(),
		Argkey:            argkey,
	}
	return
}

// RecorderAuxMockPtrroutingKeyInfoLRURemove  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrroutingKeyInfoLRURemove int = 0

var condRecorderAuxMockPtrroutingKeyInfoLRURemove *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrroutingKeyInfoLRURemove(i int) {
	condRecorderAuxMockPtrroutingKeyInfoLRURemove.L.Lock()
	for recorderAuxMockPtrroutingKeyInfoLRURemove < i {
		condRecorderAuxMockPtrroutingKeyInfoLRURemove.Wait()
	}
	condRecorderAuxMockPtrroutingKeyInfoLRURemove.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrroutingKeyInfoLRURemove() {
	condRecorderAuxMockPtrroutingKeyInfoLRURemove.L.Lock()
	recorderAuxMockPtrroutingKeyInfoLRURemove++
	condRecorderAuxMockPtrroutingKeyInfoLRURemove.L.Unlock()
	condRecorderAuxMockPtrroutingKeyInfoLRURemove.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrroutingKeyInfoLRURemove() (ret int) {
	condRecorderAuxMockPtrroutingKeyInfoLRURemove.L.Lock()
	ret = recorderAuxMockPtrroutingKeyInfoLRURemove
	condRecorderAuxMockPtrroutingKeyInfoLRURemove.L.Unlock()
	return
}

// (recvr *routingKeyInfoLRU)Remove - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *routingKeyInfoLRU) Remove(argkey string) {
	FuncAuxMockPtrroutingKeyInfoLRURemove, ok := apomock.GetRegisteredFunc("gocql.routingKeyInfoLRU.Remove")
	if ok {
		FuncAuxMockPtrroutingKeyInfoLRURemove.(func(recvr *routingKeyInfoLRU, argkey string))(recvr, argkey)
	} else {
		panic("FuncAuxMockPtrroutingKeyInfoLRURemove ")
	}
	AuxMockIncrementRecorderAuxMockPtrroutingKeyInfoLRURemove()
	return
}

//
// Mock: (recvs *Session)getConn()(reta *Conn)
//

type MockArgsTypeSessiongetConn struct {
	ApomockCallNumber int
}

var LastMockArgsSessiongetConn MockArgsTypeSessiongetConn

// (recvs *Session)AuxMockgetConn()(reta *Conn) - Generated mock function
func (recvs *Session) AuxMockgetConn() (reta *Conn) {
	rargs, rerr := apomock.GetNext("gocql.Session.getConn")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.getConn")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Session.getConn")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Conn)
	}
	return
}

// RecorderAuxMockPtrSessiongetConn  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessiongetConn int = 0

var condRecorderAuxMockPtrSessiongetConn *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessiongetConn(i int) {
	condRecorderAuxMockPtrSessiongetConn.L.Lock()
	for recorderAuxMockPtrSessiongetConn < i {
		condRecorderAuxMockPtrSessiongetConn.Wait()
	}
	condRecorderAuxMockPtrSessiongetConn.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessiongetConn() {
	condRecorderAuxMockPtrSessiongetConn.L.Lock()
	recorderAuxMockPtrSessiongetConn++
	condRecorderAuxMockPtrSessiongetConn.L.Unlock()
	condRecorderAuxMockPtrSessiongetConn.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessiongetConn() (ret int) {
	condRecorderAuxMockPtrSessiongetConn.L.Lock()
	ret = recorderAuxMockPtrSessiongetConn
	condRecorderAuxMockPtrSessiongetConn.L.Unlock()
	return
}

// (recvs *Session)getConn - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) getConn() (reta *Conn) {
	FuncAuxMockPtrSessiongetConn, ok := apomock.GetRegisteredFunc("gocql.Session.getConn")
	if ok {
		reta = FuncAuxMockPtrSessiongetConn.(func(recvs *Session) (reta *Conn))(recvs)
	} else {
		panic("FuncAuxMockPtrSessiongetConn ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessiongetConn()
	return
}

//
// Mock: (recviter *Iter)GetCustomPayload()(reta map[string][]byte)
//

type MockArgsTypeIterGetCustomPayload struct {
	ApomockCallNumber int
}

var LastMockArgsIterGetCustomPayload MockArgsTypeIterGetCustomPayload

// (recviter *Iter)AuxMockGetCustomPayload()(reta map[string][]byte) - Generated mock function
func (recviter *Iter) AuxMockGetCustomPayload() (reta map[string][]byte) {
	rargs, rerr := apomock.GetNext("gocql.Iter.GetCustomPayload")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.GetCustomPayload")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Iter.GetCustomPayload")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(map[string][]byte)
	}
	return
}

// RecorderAuxMockPtrIterGetCustomPayload  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterGetCustomPayload int = 0

var condRecorderAuxMockPtrIterGetCustomPayload *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterGetCustomPayload(i int) {
	condRecorderAuxMockPtrIterGetCustomPayload.L.Lock()
	for recorderAuxMockPtrIterGetCustomPayload < i {
		condRecorderAuxMockPtrIterGetCustomPayload.Wait()
	}
	condRecorderAuxMockPtrIterGetCustomPayload.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterGetCustomPayload() {
	condRecorderAuxMockPtrIterGetCustomPayload.L.Lock()
	recorderAuxMockPtrIterGetCustomPayload++
	condRecorderAuxMockPtrIterGetCustomPayload.L.Unlock()
	condRecorderAuxMockPtrIterGetCustomPayload.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterGetCustomPayload() (ret int) {
	condRecorderAuxMockPtrIterGetCustomPayload.L.Lock()
	ret = recorderAuxMockPtrIterGetCustomPayload
	condRecorderAuxMockPtrIterGetCustomPayload.L.Unlock()
	return
}

// (recviter *Iter)GetCustomPayload - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) GetCustomPayload() (reta map[string][]byte) {
	FuncAuxMockPtrIterGetCustomPayload, ok := apomock.GetRegisteredFunc("gocql.Iter.GetCustomPayload")
	if ok {
		reta = FuncAuxMockPtrIterGetCustomPayload.(func(recviter *Iter) (reta map[string][]byte))(recviter)
	} else {
		panic("FuncAuxMockPtrIterGetCustomPayload ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterGetCustomPayload()
	return
}

//
// Mock: (recvb *Batch)Attempts()(reta int)
//

type MockArgsTypeBatchAttempts struct {
	ApomockCallNumber int
}

var LastMockArgsBatchAttempts MockArgsTypeBatchAttempts

// (recvb *Batch)AuxMockAttempts()(reta int) - Generated mock function
func (recvb *Batch) AuxMockAttempts() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.Batch.Attempts")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Batch.Attempts")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Batch.Attempts")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrBatchAttempts  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchAttempts int = 0

var condRecorderAuxMockPtrBatchAttempts *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchAttempts(i int) {
	condRecorderAuxMockPtrBatchAttempts.L.Lock()
	for recorderAuxMockPtrBatchAttempts < i {
		condRecorderAuxMockPtrBatchAttempts.Wait()
	}
	condRecorderAuxMockPtrBatchAttempts.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchAttempts() {
	condRecorderAuxMockPtrBatchAttempts.L.Lock()
	recorderAuxMockPtrBatchAttempts++
	condRecorderAuxMockPtrBatchAttempts.L.Unlock()
	condRecorderAuxMockPtrBatchAttempts.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchAttempts() (ret int) {
	condRecorderAuxMockPtrBatchAttempts.L.Lock()
	ret = recorderAuxMockPtrBatchAttempts
	condRecorderAuxMockPtrBatchAttempts.L.Unlock()
	return
}

// (recvb *Batch)Attempts - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) Attempts() (reta int) {
	FuncAuxMockPtrBatchAttempts, ok := apomock.GetRegisteredFunc("gocql.Batch.Attempts")
	if ok {
		reta = FuncAuxMockPtrBatchAttempts.(func(recvb *Batch) (reta int))(recvb)
	} else {
		panic("FuncAuxMockPtrBatchAttempts ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchAttempts()
	return
}

//
// Mock: (recvb *Batch)GetRoutingKey()(reta []byte, retb error)
//

type MockArgsTypeBatchGetRoutingKey struct {
	ApomockCallNumber int
}

var LastMockArgsBatchGetRoutingKey MockArgsTypeBatchGetRoutingKey

// (recvb *Batch)AuxMockGetRoutingKey()(reta []byte, retb error) - Generated mock function
func (recvb *Batch) AuxMockGetRoutingKey() (reta []byte, retb error) {
	rargs, rerr := apomock.GetNext("gocql.Batch.GetRoutingKey")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Batch.GetRoutingKey")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Batch.GetRoutingKey")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrBatchGetRoutingKey  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchGetRoutingKey int = 0

var condRecorderAuxMockPtrBatchGetRoutingKey *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchGetRoutingKey(i int) {
	condRecorderAuxMockPtrBatchGetRoutingKey.L.Lock()
	for recorderAuxMockPtrBatchGetRoutingKey < i {
		condRecorderAuxMockPtrBatchGetRoutingKey.Wait()
	}
	condRecorderAuxMockPtrBatchGetRoutingKey.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchGetRoutingKey() {
	condRecorderAuxMockPtrBatchGetRoutingKey.L.Lock()
	recorderAuxMockPtrBatchGetRoutingKey++
	condRecorderAuxMockPtrBatchGetRoutingKey.L.Unlock()
	condRecorderAuxMockPtrBatchGetRoutingKey.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchGetRoutingKey() (ret int) {
	condRecorderAuxMockPtrBatchGetRoutingKey.L.Lock()
	ret = recorderAuxMockPtrBatchGetRoutingKey
	condRecorderAuxMockPtrBatchGetRoutingKey.L.Unlock()
	return
}

// (recvb *Batch)GetRoutingKey - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) GetRoutingKey() (reta []byte, retb error) {
	FuncAuxMockPtrBatchGetRoutingKey, ok := apomock.GetRegisteredFunc("gocql.Batch.GetRoutingKey")
	if ok {
		reta, retb = FuncAuxMockPtrBatchGetRoutingKey.(func(recvb *Batch) (reta []byte, retb error))(recvb)
	} else {
		panic("FuncAuxMockPtrBatchGetRoutingKey ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchGetRoutingKey()
	return
}

//
// Mock: (recvq *Query)RetryPolicy(argr RetryPolicy)(reta *Query)
//

type MockArgsTypeQueryRetryPolicy struct {
	ApomockCallNumber int
	Argr              RetryPolicy
}

var LastMockArgsQueryRetryPolicy MockArgsTypeQueryRetryPolicy

// (recvq *Query)AuxMockRetryPolicy(argr RetryPolicy)(reta *Query) - Generated mock function
func (recvq *Query) AuxMockRetryPolicy(argr RetryPolicy) (reta *Query) {
	LastMockArgsQueryRetryPolicy = MockArgsTypeQueryRetryPolicy{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryRetryPolicy(),
		Argr:              argr,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.RetryPolicy")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.RetryPolicy")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.RetryPolicy")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryRetryPolicy  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryRetryPolicy int = 0

var condRecorderAuxMockPtrQueryRetryPolicy *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryRetryPolicy(i int) {
	condRecorderAuxMockPtrQueryRetryPolicy.L.Lock()
	for recorderAuxMockPtrQueryRetryPolicy < i {
		condRecorderAuxMockPtrQueryRetryPolicy.Wait()
	}
	condRecorderAuxMockPtrQueryRetryPolicy.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryRetryPolicy() {
	condRecorderAuxMockPtrQueryRetryPolicy.L.Lock()
	recorderAuxMockPtrQueryRetryPolicy++
	condRecorderAuxMockPtrQueryRetryPolicy.L.Unlock()
	condRecorderAuxMockPtrQueryRetryPolicy.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryRetryPolicy() (ret int) {
	condRecorderAuxMockPtrQueryRetryPolicy.L.Lock()
	ret = recorderAuxMockPtrQueryRetryPolicy
	condRecorderAuxMockPtrQueryRetryPolicy.L.Unlock()
	return
}

// (recvq *Query)RetryPolicy - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) RetryPolicy(argr RetryPolicy) (reta *Query) {
	FuncAuxMockPtrQueryRetryPolicy, ok := apomock.GetRegisteredFunc("gocql.Query.RetryPolicy")
	if ok {
		reta = FuncAuxMockPtrQueryRetryPolicy.(func(recvq *Query, argr RetryPolicy) (reta *Query))(recvq, argr)
	} else {
		panic("FuncAuxMockPtrQueryRetryPolicy ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryRetryPolicy()
	return
}

//
// Mock: (recvs *Session)Bind(argstmt string, argb func(*QueryInfo) ([]interface{}, error))(reta *Query)
//

type MockArgsTypeSessionBind struct {
	ApomockCallNumber int
	Argstmt           string
	Argb              func(*QueryInfo) ([]interface{}, error)
}

var LastMockArgsSessionBind MockArgsTypeSessionBind

// (recvs *Session)AuxMockBind(argstmt string, argb func(*QueryInfo) ([]interface{}, error))(reta *Query) - Generated mock function
func (recvs *Session) AuxMockBind(argstmt string, argb func(*QueryInfo) ([]interface{}, error)) (reta *Query) {
	LastMockArgsSessionBind = MockArgsTypeSessionBind{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionBind(),
		Argstmt:           argstmt,
		Argb:              argb,
	}
	rargs, rerr := apomock.GetNext("gocql.Session.Bind")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.Bind")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Session.Bind")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrSessionBind  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionBind int = 0

var condRecorderAuxMockPtrSessionBind *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionBind(i int) {
	condRecorderAuxMockPtrSessionBind.L.Lock()
	for recorderAuxMockPtrSessionBind < i {
		condRecorderAuxMockPtrSessionBind.Wait()
	}
	condRecorderAuxMockPtrSessionBind.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionBind() {
	condRecorderAuxMockPtrSessionBind.L.Lock()
	recorderAuxMockPtrSessionBind++
	condRecorderAuxMockPtrSessionBind.L.Unlock()
	condRecorderAuxMockPtrSessionBind.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionBind() (ret int) {
	condRecorderAuxMockPtrSessionBind.L.Lock()
	ret = recorderAuxMockPtrSessionBind
	condRecorderAuxMockPtrSessionBind.L.Unlock()
	return
}

// (recvs *Session)Bind - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) Bind(argstmt string, argb func(*QueryInfo) ([]interface{}, error)) (reta *Query) {
	FuncAuxMockPtrSessionBind, ok := apomock.GetRegisteredFunc("gocql.Session.Bind")
	if ok {
		reta = FuncAuxMockPtrSessionBind.(func(recvs *Session, argstmt string, argb func(*QueryInfo) ([]interface{}, error)) (reta *Query))(recvs, argstmt, argb)
	} else {
		panic("FuncAuxMockPtrSessionBind ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionBind()
	return
}

//
// Mock: (recvq *Query)Attempts()(reta int)
//

type MockArgsTypeQueryAttempts struct {
	ApomockCallNumber int
}

var LastMockArgsQueryAttempts MockArgsTypeQueryAttempts

// (recvq *Query)AuxMockAttempts()(reta int) - Generated mock function
func (recvq *Query) AuxMockAttempts() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.Query.Attempts")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.Attempts")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.Attempts")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrQueryAttempts  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryAttempts int = 0

var condRecorderAuxMockPtrQueryAttempts *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryAttempts(i int) {
	condRecorderAuxMockPtrQueryAttempts.L.Lock()
	for recorderAuxMockPtrQueryAttempts < i {
		condRecorderAuxMockPtrQueryAttempts.Wait()
	}
	condRecorderAuxMockPtrQueryAttempts.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryAttempts() {
	condRecorderAuxMockPtrQueryAttempts.L.Lock()
	recorderAuxMockPtrQueryAttempts++
	condRecorderAuxMockPtrQueryAttempts.L.Unlock()
	condRecorderAuxMockPtrQueryAttempts.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryAttempts() (ret int) {
	condRecorderAuxMockPtrQueryAttempts.L.Lock()
	ret = recorderAuxMockPtrQueryAttempts
	condRecorderAuxMockPtrQueryAttempts.L.Unlock()
	return
}

// (recvq *Query)Attempts - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) Attempts() (reta int) {
	FuncAuxMockPtrQueryAttempts, ok := apomock.GetRegisteredFunc("gocql.Query.Attempts")
	if ok {
		reta = FuncAuxMockPtrQueryAttempts.(func(recvq *Query) (reta int))(recvq)
	} else {
		panic("FuncAuxMockPtrQueryAttempts ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryAttempts()
	return
}

//
// Mock: (recvq *Query)WithContext(argctx context.Context)(reta *Query)
//

type MockArgsTypeQueryWithContext struct {
	ApomockCallNumber int
	Argctx            context.Context
}

var LastMockArgsQueryWithContext MockArgsTypeQueryWithContext

// (recvq *Query)AuxMockWithContext(argctx context.Context)(reta *Query) - Generated mock function
func (recvq *Query) AuxMockWithContext(argctx context.Context) (reta *Query) {
	LastMockArgsQueryWithContext = MockArgsTypeQueryWithContext{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryWithContext(),
		Argctx:            argctx,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.WithContext")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.WithContext")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.WithContext")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryWithContext  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryWithContext int = 0

var condRecorderAuxMockPtrQueryWithContext *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryWithContext(i int) {
	condRecorderAuxMockPtrQueryWithContext.L.Lock()
	for recorderAuxMockPtrQueryWithContext < i {
		condRecorderAuxMockPtrQueryWithContext.Wait()
	}
	condRecorderAuxMockPtrQueryWithContext.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryWithContext() {
	condRecorderAuxMockPtrQueryWithContext.L.Lock()
	recorderAuxMockPtrQueryWithContext++
	condRecorderAuxMockPtrQueryWithContext.L.Unlock()
	condRecorderAuxMockPtrQueryWithContext.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryWithContext() (ret int) {
	condRecorderAuxMockPtrQueryWithContext.L.Lock()
	ret = recorderAuxMockPtrQueryWithContext
	condRecorderAuxMockPtrQueryWithContext.L.Unlock()
	return
}

// (recvq *Query)WithContext - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) WithContext(argctx context.Context) (reta *Query) {
	FuncAuxMockPtrQueryWithContext, ok := apomock.GetRegisteredFunc("gocql.Query.WithContext")
	if ok {
		reta = FuncAuxMockPtrQueryWithContext.(func(recvq *Query, argctx context.Context) (reta *Query))(recvq, argctx)
	} else {
		panic("FuncAuxMockPtrQueryWithContext ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryWithContext()
	return
}

//
// Mock: (recvq *Query)execute(argconn *Conn)(reta *Iter)
//

type MockArgsTypeQueryexecute struct {
	ApomockCallNumber int
	Argconn           *Conn
}

var LastMockArgsQueryexecute MockArgsTypeQueryexecute

// (recvq *Query)AuxMockexecute(argconn *Conn)(reta *Iter) - Generated mock function
func (recvq *Query) AuxMockexecute(argconn *Conn) (reta *Iter) {
	LastMockArgsQueryexecute = MockArgsTypeQueryexecute{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryexecute(),
		Argconn:           argconn,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.execute")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.execute")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.execute")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Iter)
	}
	return
}

// RecorderAuxMockPtrQueryexecute  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryexecute int = 0

var condRecorderAuxMockPtrQueryexecute *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryexecute(i int) {
	condRecorderAuxMockPtrQueryexecute.L.Lock()
	for recorderAuxMockPtrQueryexecute < i {
		condRecorderAuxMockPtrQueryexecute.Wait()
	}
	condRecorderAuxMockPtrQueryexecute.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryexecute() {
	condRecorderAuxMockPtrQueryexecute.L.Lock()
	recorderAuxMockPtrQueryexecute++
	condRecorderAuxMockPtrQueryexecute.L.Unlock()
	condRecorderAuxMockPtrQueryexecute.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryexecute() (ret int) {
	condRecorderAuxMockPtrQueryexecute.L.Lock()
	ret = recorderAuxMockPtrQueryexecute
	condRecorderAuxMockPtrQueryexecute.L.Unlock()
	return
}

// (recvq *Query)execute - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) execute(argconn *Conn) (reta *Iter) {
	FuncAuxMockPtrQueryexecute, ok := apomock.GetRegisteredFunc("gocql.Query.execute")
	if ok {
		reta = FuncAuxMockPtrQueryexecute.(func(recvq *Query, argconn *Conn) (reta *Iter))(recvq, argconn)
	} else {
		panic("FuncAuxMockPtrQueryexecute ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryexecute()
	return
}

//
// Mock: (recvs *Session)reconnectDownedHosts(argintv time.Duration)()
//

type MockArgsTypeSessionreconnectDownedHosts struct {
	ApomockCallNumber int
	Argintv           time.Duration
}

var LastMockArgsSessionreconnectDownedHosts MockArgsTypeSessionreconnectDownedHosts

// (recvs *Session)AuxMockreconnectDownedHosts(argintv time.Duration)() - Generated mock function
func (recvs *Session) AuxMockreconnectDownedHosts(argintv time.Duration) {
	LastMockArgsSessionreconnectDownedHosts = MockArgsTypeSessionreconnectDownedHosts{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionreconnectDownedHosts(),
		Argintv:           argintv,
	}
	return
}

// RecorderAuxMockPtrSessionreconnectDownedHosts  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionreconnectDownedHosts int = 0

var condRecorderAuxMockPtrSessionreconnectDownedHosts *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionreconnectDownedHosts(i int) {
	condRecorderAuxMockPtrSessionreconnectDownedHosts.L.Lock()
	for recorderAuxMockPtrSessionreconnectDownedHosts < i {
		condRecorderAuxMockPtrSessionreconnectDownedHosts.Wait()
	}
	condRecorderAuxMockPtrSessionreconnectDownedHosts.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionreconnectDownedHosts() {
	condRecorderAuxMockPtrSessionreconnectDownedHosts.L.Lock()
	recorderAuxMockPtrSessionreconnectDownedHosts++
	condRecorderAuxMockPtrSessionreconnectDownedHosts.L.Unlock()
	condRecorderAuxMockPtrSessionreconnectDownedHosts.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionreconnectDownedHosts() (ret int) {
	condRecorderAuxMockPtrSessionreconnectDownedHosts.L.Lock()
	ret = recorderAuxMockPtrSessionreconnectDownedHosts
	condRecorderAuxMockPtrSessionreconnectDownedHosts.L.Unlock()
	return
}

// (recvs *Session)reconnectDownedHosts - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) reconnectDownedHosts(argintv time.Duration) {
	FuncAuxMockPtrSessionreconnectDownedHosts, ok := apomock.GetRegisteredFunc("gocql.Session.reconnectDownedHosts")
	if ok {
		FuncAuxMockPtrSessionreconnectDownedHosts.(func(recvs *Session, argintv time.Duration))(recvs, argintv)
	} else {
		panic("FuncAuxMockPtrSessionreconnectDownedHosts ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionreconnectDownedHosts()
	return
}

//
// Mock: (recvq *Query)GetRoutingKey()(reta []byte, retb error)
//

type MockArgsTypeQueryGetRoutingKey struct {
	ApomockCallNumber int
}

var LastMockArgsQueryGetRoutingKey MockArgsTypeQueryGetRoutingKey

// (recvq *Query)AuxMockGetRoutingKey()(reta []byte, retb error) - Generated mock function
func (recvq *Query) AuxMockGetRoutingKey() (reta []byte, retb error) {
	rargs, rerr := apomock.GetNext("gocql.Query.GetRoutingKey")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.GetRoutingKey")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Query.GetRoutingKey")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrQueryGetRoutingKey  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryGetRoutingKey int = 0

var condRecorderAuxMockPtrQueryGetRoutingKey *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryGetRoutingKey(i int) {
	condRecorderAuxMockPtrQueryGetRoutingKey.L.Lock()
	for recorderAuxMockPtrQueryGetRoutingKey < i {
		condRecorderAuxMockPtrQueryGetRoutingKey.Wait()
	}
	condRecorderAuxMockPtrQueryGetRoutingKey.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryGetRoutingKey() {
	condRecorderAuxMockPtrQueryGetRoutingKey.L.Lock()
	recorderAuxMockPtrQueryGetRoutingKey++
	condRecorderAuxMockPtrQueryGetRoutingKey.L.Unlock()
	condRecorderAuxMockPtrQueryGetRoutingKey.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryGetRoutingKey() (ret int) {
	condRecorderAuxMockPtrQueryGetRoutingKey.L.Lock()
	ret = recorderAuxMockPtrQueryGetRoutingKey
	condRecorderAuxMockPtrQueryGetRoutingKey.L.Unlock()
	return
}

// (recvq *Query)GetRoutingKey - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) GetRoutingKey() (reta []byte, retb error) {
	FuncAuxMockPtrQueryGetRoutingKey, ok := apomock.GetRegisteredFunc("gocql.Query.GetRoutingKey")
	if ok {
		reta, retb = FuncAuxMockPtrQueryGetRoutingKey.(func(recvq *Query) (reta []byte, retb error))(recvq)
	} else {
		panic("FuncAuxMockPtrQueryGetRoutingKey ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryGetRoutingKey()
	return
}

//
// Mock: (recvq *Query)SerialConsistency(argcons SerialConsistency)(reta *Query)
//

type MockArgsTypeQuerySerialConsistency struct {
	ApomockCallNumber int
	Argcons           SerialConsistency
}

var LastMockArgsQuerySerialConsistency MockArgsTypeQuerySerialConsistency

// (recvq *Query)AuxMockSerialConsistency(argcons SerialConsistency)(reta *Query) - Generated mock function
func (recvq *Query) AuxMockSerialConsistency(argcons SerialConsistency) (reta *Query) {
	LastMockArgsQuerySerialConsistency = MockArgsTypeQuerySerialConsistency{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQuerySerialConsistency(),
		Argcons:           argcons,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.SerialConsistency")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.SerialConsistency")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.SerialConsistency")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQuerySerialConsistency  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQuerySerialConsistency int = 0

var condRecorderAuxMockPtrQuerySerialConsistency *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQuerySerialConsistency(i int) {
	condRecorderAuxMockPtrQuerySerialConsistency.L.Lock()
	for recorderAuxMockPtrQuerySerialConsistency < i {
		condRecorderAuxMockPtrQuerySerialConsistency.Wait()
	}
	condRecorderAuxMockPtrQuerySerialConsistency.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQuerySerialConsistency() {
	condRecorderAuxMockPtrQuerySerialConsistency.L.Lock()
	recorderAuxMockPtrQuerySerialConsistency++
	condRecorderAuxMockPtrQuerySerialConsistency.L.Unlock()
	condRecorderAuxMockPtrQuerySerialConsistency.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQuerySerialConsistency() (ret int) {
	condRecorderAuxMockPtrQuerySerialConsistency.L.Lock()
	ret = recorderAuxMockPtrQuerySerialConsistency
	condRecorderAuxMockPtrQuerySerialConsistency.L.Unlock()
	return
}

// (recvq *Query)SerialConsistency - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) SerialConsistency(argcons SerialConsistency) (reta *Query) {
	FuncAuxMockPtrQuerySerialConsistency, ok := apomock.GetRegisteredFunc("gocql.Query.SerialConsistency")
	if ok {
		reta = FuncAuxMockPtrQuerySerialConsistency.(func(recvq *Query, argcons SerialConsistency) (reta *Query))(recvq, argcons)
	} else {
		panic("FuncAuxMockPtrQuerySerialConsistency ")
	}
	AuxMockIncrementRecorderAuxMockPtrQuerySerialConsistency()
	return
}

//
// Mock: (recvb *Batch)RetryPolicy(argr RetryPolicy)(reta *Batch)
//

type MockArgsTypeBatchRetryPolicy struct {
	ApomockCallNumber int
	Argr              RetryPolicy
}

var LastMockArgsBatchRetryPolicy MockArgsTypeBatchRetryPolicy

// (recvb *Batch)AuxMockRetryPolicy(argr RetryPolicy)(reta *Batch) - Generated mock function
func (recvb *Batch) AuxMockRetryPolicy(argr RetryPolicy) (reta *Batch) {
	LastMockArgsBatchRetryPolicy = MockArgsTypeBatchRetryPolicy{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrBatchRetryPolicy(),
		Argr:              argr,
	}
	rargs, rerr := apomock.GetNext("gocql.Batch.RetryPolicy")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Batch.RetryPolicy")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Batch.RetryPolicy")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Batch)
	}
	return
}

// RecorderAuxMockPtrBatchRetryPolicy  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchRetryPolicy int = 0

var condRecorderAuxMockPtrBatchRetryPolicy *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchRetryPolicy(i int) {
	condRecorderAuxMockPtrBatchRetryPolicy.L.Lock()
	for recorderAuxMockPtrBatchRetryPolicy < i {
		condRecorderAuxMockPtrBatchRetryPolicy.Wait()
	}
	condRecorderAuxMockPtrBatchRetryPolicy.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchRetryPolicy() {
	condRecorderAuxMockPtrBatchRetryPolicy.L.Lock()
	recorderAuxMockPtrBatchRetryPolicy++
	condRecorderAuxMockPtrBatchRetryPolicy.L.Unlock()
	condRecorderAuxMockPtrBatchRetryPolicy.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchRetryPolicy() (ret int) {
	condRecorderAuxMockPtrBatchRetryPolicy.L.Lock()
	ret = recorderAuxMockPtrBatchRetryPolicy
	condRecorderAuxMockPtrBatchRetryPolicy.L.Unlock()
	return
}

// (recvb *Batch)RetryPolicy - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) RetryPolicy(argr RetryPolicy) (reta *Batch) {
	FuncAuxMockPtrBatchRetryPolicy, ok := apomock.GetRegisteredFunc("gocql.Batch.RetryPolicy")
	if ok {
		reta = FuncAuxMockPtrBatchRetryPolicy.(func(recvb *Batch, argr RetryPolicy) (reta *Batch))(recvb, argr)
	} else {
		panic("FuncAuxMockPtrBatchRetryPolicy ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchRetryPolicy()
	return
}

//
// Mock: (recvr *routingKeyInfoLRU)Max(argmax int)()
//

type MockArgsTyperoutingKeyInfoLRUMax struct {
	ApomockCallNumber int
	Argmax            int
}

var LastMockArgsroutingKeyInfoLRUMax MockArgsTyperoutingKeyInfoLRUMax

// (recvr *routingKeyInfoLRU)AuxMockMax(argmax int)() - Generated mock function
func (recvr *routingKeyInfoLRU) AuxMockMax(argmax int) {
	LastMockArgsroutingKeyInfoLRUMax = MockArgsTyperoutingKeyInfoLRUMax{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrroutingKeyInfoLRUMax(),
		Argmax:            argmax,
	}
	return
}

// RecorderAuxMockPtrroutingKeyInfoLRUMax  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrroutingKeyInfoLRUMax int = 0

var condRecorderAuxMockPtrroutingKeyInfoLRUMax *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrroutingKeyInfoLRUMax(i int) {
	condRecorderAuxMockPtrroutingKeyInfoLRUMax.L.Lock()
	for recorderAuxMockPtrroutingKeyInfoLRUMax < i {
		condRecorderAuxMockPtrroutingKeyInfoLRUMax.Wait()
	}
	condRecorderAuxMockPtrroutingKeyInfoLRUMax.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrroutingKeyInfoLRUMax() {
	condRecorderAuxMockPtrroutingKeyInfoLRUMax.L.Lock()
	recorderAuxMockPtrroutingKeyInfoLRUMax++
	condRecorderAuxMockPtrroutingKeyInfoLRUMax.L.Unlock()
	condRecorderAuxMockPtrroutingKeyInfoLRUMax.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrroutingKeyInfoLRUMax() (ret int) {
	condRecorderAuxMockPtrroutingKeyInfoLRUMax.L.Lock()
	ret = recorderAuxMockPtrroutingKeyInfoLRUMax
	condRecorderAuxMockPtrroutingKeyInfoLRUMax.L.Unlock()
	return
}

// (recvr *routingKeyInfoLRU)Max - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *routingKeyInfoLRU) Max(argmax int) {
	FuncAuxMockPtrroutingKeyInfoLRUMax, ok := apomock.GetRegisteredFunc("gocql.routingKeyInfoLRU.Max")
	if ok {
		FuncAuxMockPtrroutingKeyInfoLRUMax.(func(recvr *routingKeyInfoLRU, argmax int))(recvr, argmax)
	} else {
		panic("FuncAuxMockPtrroutingKeyInfoLRUMax ")
	}
	AuxMockIncrementRecorderAuxMockPtrroutingKeyInfoLRUMax()
	return
}

//
// Mock: (recvq *Query)Consistency(argc Consistency)(reta *Query)
//

type MockArgsTypeQueryConsistency struct {
	ApomockCallNumber int
	Argc              Consistency
}

var LastMockArgsQueryConsistency MockArgsTypeQueryConsistency

// (recvq *Query)AuxMockConsistency(argc Consistency)(reta *Query) - Generated mock function
func (recvq *Query) AuxMockConsistency(argc Consistency) (reta *Query) {
	LastMockArgsQueryConsistency = MockArgsTypeQueryConsistency{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryConsistency(),
		Argc:              argc,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.Consistency")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.Consistency")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.Consistency")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryConsistency  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryConsistency int = 0

var condRecorderAuxMockPtrQueryConsistency *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryConsistency(i int) {
	condRecorderAuxMockPtrQueryConsistency.L.Lock()
	for recorderAuxMockPtrQueryConsistency < i {
		condRecorderAuxMockPtrQueryConsistency.Wait()
	}
	condRecorderAuxMockPtrQueryConsistency.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryConsistency() {
	condRecorderAuxMockPtrQueryConsistency.L.Lock()
	recorderAuxMockPtrQueryConsistency++
	condRecorderAuxMockPtrQueryConsistency.L.Unlock()
	condRecorderAuxMockPtrQueryConsistency.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryConsistency() (ret int) {
	condRecorderAuxMockPtrQueryConsistency.L.Lock()
	ret = recorderAuxMockPtrQueryConsistency
	condRecorderAuxMockPtrQueryConsistency.L.Unlock()
	return
}

// (recvq *Query)Consistency - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) Consistency(argc Consistency) (reta *Query) {
	FuncAuxMockPtrQueryConsistency, ok := apomock.GetRegisteredFunc("gocql.Query.Consistency")
	if ok {
		reta = FuncAuxMockPtrQueryConsistency.(func(recvq *Query, argc Consistency) (reta *Query))(recvq, argc)
	} else {
		panic("FuncAuxMockPtrQueryConsistency ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryConsistency()
	return
}

//
// Mock: (recvq *Query)Exec()(reta error)
//

type MockArgsTypeQueryExec struct {
	ApomockCallNumber int
}

var LastMockArgsQueryExec MockArgsTypeQueryExec

// (recvq *Query)AuxMockExec()(reta error) - Generated mock function
func (recvq *Query) AuxMockExec() (reta error) {
	rargs, rerr := apomock.GetNext("gocql.Query.Exec")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.Exec")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.Exec")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrQueryExec  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryExec int = 0

var condRecorderAuxMockPtrQueryExec *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryExec(i int) {
	condRecorderAuxMockPtrQueryExec.L.Lock()
	for recorderAuxMockPtrQueryExec < i {
		condRecorderAuxMockPtrQueryExec.Wait()
	}
	condRecorderAuxMockPtrQueryExec.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryExec() {
	condRecorderAuxMockPtrQueryExec.L.Lock()
	recorderAuxMockPtrQueryExec++
	condRecorderAuxMockPtrQueryExec.L.Unlock()
	condRecorderAuxMockPtrQueryExec.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryExec() (ret int) {
	condRecorderAuxMockPtrQueryExec.L.Lock()
	ret = recorderAuxMockPtrQueryExec
	condRecorderAuxMockPtrQueryExec.L.Unlock()
	return
}

// (recvq *Query)Exec - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) Exec() (reta error) {
	FuncAuxMockPtrQueryExec, ok := apomock.GetRegisteredFunc("gocql.Query.Exec")
	if ok {
		reta = FuncAuxMockPtrQueryExec.(func(recvq *Query) (reta error))(recvq)
	} else {
		panic("FuncAuxMockPtrQueryExec ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryExec()
	return
}

//
// Mock: (recvn *nextIter)fetch()(reta *Iter)
//

type MockArgsTypenextIterfetch struct {
	ApomockCallNumber int
}

var LastMockArgsnextIterfetch MockArgsTypenextIterfetch

// (recvn *nextIter)AuxMockfetch()(reta *Iter) - Generated mock function
func (recvn *nextIter) AuxMockfetch() (reta *Iter) {
	rargs, rerr := apomock.GetNext("gocql.nextIter.fetch")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.nextIter.fetch")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.nextIter.fetch")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Iter)
	}
	return
}

// RecorderAuxMockPtrnextIterfetch  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrnextIterfetch int = 0

var condRecorderAuxMockPtrnextIterfetch *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrnextIterfetch(i int) {
	condRecorderAuxMockPtrnextIterfetch.L.Lock()
	for recorderAuxMockPtrnextIterfetch < i {
		condRecorderAuxMockPtrnextIterfetch.Wait()
	}
	condRecorderAuxMockPtrnextIterfetch.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrnextIterfetch() {
	condRecorderAuxMockPtrnextIterfetch.L.Lock()
	recorderAuxMockPtrnextIterfetch++
	condRecorderAuxMockPtrnextIterfetch.L.Unlock()
	condRecorderAuxMockPtrnextIterfetch.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrnextIterfetch() (ret int) {
	condRecorderAuxMockPtrnextIterfetch.L.Lock()
	ret = recorderAuxMockPtrnextIterfetch
	condRecorderAuxMockPtrnextIterfetch.L.Unlock()
	return
}

// (recvn *nextIter)fetch - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvn *nextIter) fetch() (reta *Iter) {
	FuncAuxMockPtrnextIterfetch, ok := apomock.GetRegisteredFunc("gocql.nextIter.fetch")
	if ok {
		reta = FuncAuxMockPtrnextIterfetch.(func(recvn *nextIter) (reta *Iter))(recvn)
	} else {
		panic("FuncAuxMockPtrnextIterfetch ")
	}
	AuxMockIncrementRecorderAuxMockPtrnextIterfetch()
	return
}

//
// Mock: (recvq *Query)attempt(argd time.Duration)()
//

type MockArgsTypeQueryattempt struct {
	ApomockCallNumber int
	Argd              time.Duration
}

var LastMockArgsQueryattempt MockArgsTypeQueryattempt

// (recvq *Query)AuxMockattempt(argd time.Duration)() - Generated mock function
func (recvq *Query) AuxMockattempt(argd time.Duration) {
	LastMockArgsQueryattempt = MockArgsTypeQueryattempt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryattempt(),
		Argd:              argd,
	}
	return
}

// RecorderAuxMockPtrQueryattempt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryattempt int = 0

var condRecorderAuxMockPtrQueryattempt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryattempt(i int) {
	condRecorderAuxMockPtrQueryattempt.L.Lock()
	for recorderAuxMockPtrQueryattempt < i {
		condRecorderAuxMockPtrQueryattempt.Wait()
	}
	condRecorderAuxMockPtrQueryattempt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryattempt() {
	condRecorderAuxMockPtrQueryattempt.L.Lock()
	recorderAuxMockPtrQueryattempt++
	condRecorderAuxMockPtrQueryattempt.L.Unlock()
	condRecorderAuxMockPtrQueryattempt.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryattempt() (ret int) {
	condRecorderAuxMockPtrQueryattempt.L.Lock()
	ret = recorderAuxMockPtrQueryattempt
	condRecorderAuxMockPtrQueryattempt.L.Unlock()
	return
}

// (recvq *Query)attempt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) attempt(argd time.Duration) {
	FuncAuxMockPtrQueryattempt, ok := apomock.GetRegisteredFunc("gocql.Query.attempt")
	if ok {
		FuncAuxMockPtrQueryattempt.(func(recvq *Query, argd time.Duration))(recvq, argd)
	} else {
		panic("FuncAuxMockPtrQueryattempt ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryattempt()
	return
}

//
// Mock: (recvq *Query)Prefetch(argp float64)(reta *Query)
//

type MockArgsTypeQueryPrefetch struct {
	ApomockCallNumber int
	Argp              float64
}

var LastMockArgsQueryPrefetch MockArgsTypeQueryPrefetch

// (recvq *Query)AuxMockPrefetch(argp float64)(reta *Query) - Generated mock function
func (recvq *Query) AuxMockPrefetch(argp float64) (reta *Query) {
	LastMockArgsQueryPrefetch = MockArgsTypeQueryPrefetch{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryPrefetch(),
		Argp:              argp,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.Prefetch")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.Prefetch")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.Prefetch")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryPrefetch  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryPrefetch int = 0

var condRecorderAuxMockPtrQueryPrefetch *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryPrefetch(i int) {
	condRecorderAuxMockPtrQueryPrefetch.L.Lock()
	for recorderAuxMockPtrQueryPrefetch < i {
		condRecorderAuxMockPtrQueryPrefetch.Wait()
	}
	condRecorderAuxMockPtrQueryPrefetch.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryPrefetch() {
	condRecorderAuxMockPtrQueryPrefetch.L.Lock()
	recorderAuxMockPtrQueryPrefetch++
	condRecorderAuxMockPtrQueryPrefetch.L.Unlock()
	condRecorderAuxMockPtrQueryPrefetch.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryPrefetch() (ret int) {
	condRecorderAuxMockPtrQueryPrefetch.L.Lock()
	ret = recorderAuxMockPtrQueryPrefetch
	condRecorderAuxMockPtrQueryPrefetch.L.Unlock()
	return
}

// (recvq *Query)Prefetch - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) Prefetch(argp float64) (reta *Query) {
	FuncAuxMockPtrQueryPrefetch, ok := apomock.GetRegisteredFunc("gocql.Query.Prefetch")
	if ok {
		reta = FuncAuxMockPtrQueryPrefetch.(func(recvq *Query, argp float64) (reta *Query))(recvq, argp)
	} else {
		panic("FuncAuxMockPtrQueryPrefetch ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryPrefetch()
	return
}

//
// Mock: (recvq *Query)ScanCAS(dest ...interface{})(retapplied bool, reterr error)
//

type MockArgsTypeQueryScanCAS struct {
	ApomockCallNumber int
	Dest              []interface{}
}

var LastMockArgsQueryScanCAS MockArgsTypeQueryScanCAS

// (recvq *Query)AuxMockScanCAS(dest ...interface{})(retapplied bool, reterr error) - Generated mock function
func (recvq *Query) AuxMockScanCAS(dest ...interface{}) (retapplied bool, reterr error) {
	LastMockArgsQueryScanCAS = MockArgsTypeQueryScanCAS{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryScanCAS(),
		Dest:              dest,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.ScanCAS")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.ScanCAS")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Query.ScanCAS")
	}
	if rargs.GetArg(0) != nil {
		retapplied = rargs.GetArg(0).(bool)
	}
	if rargs.GetArg(1) != nil {
		reterr = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrQueryScanCAS  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryScanCAS int = 0

var condRecorderAuxMockPtrQueryScanCAS *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryScanCAS(i int) {
	condRecorderAuxMockPtrQueryScanCAS.L.Lock()
	for recorderAuxMockPtrQueryScanCAS < i {
		condRecorderAuxMockPtrQueryScanCAS.Wait()
	}
	condRecorderAuxMockPtrQueryScanCAS.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryScanCAS() {
	condRecorderAuxMockPtrQueryScanCAS.L.Lock()
	recorderAuxMockPtrQueryScanCAS++
	condRecorderAuxMockPtrQueryScanCAS.L.Unlock()
	condRecorderAuxMockPtrQueryScanCAS.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryScanCAS() (ret int) {
	condRecorderAuxMockPtrQueryScanCAS.L.Lock()
	ret = recorderAuxMockPtrQueryScanCAS
	condRecorderAuxMockPtrQueryScanCAS.L.Unlock()
	return
}

// (recvq *Query)ScanCAS - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) ScanCAS(dest ...interface{}) (retapplied bool, reterr error) {
	FuncAuxMockPtrQueryScanCAS, ok := apomock.GetRegisteredFunc("gocql.Query.ScanCAS")
	if ok {
		retapplied, reterr = FuncAuxMockPtrQueryScanCAS.(func(recvq *Query, dest ...interface{}) (retapplied bool, reterr error))(recvq, dest...)
	} else {
		panic("FuncAuxMockPtrQueryScanCAS ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryScanCAS()
	return
}

//
// Mock: (recvb *Batch)DefaultTimestamp(argenable bool)(reta *Batch)
//

type MockArgsTypeBatchDefaultTimestamp struct {
	ApomockCallNumber int
	Argenable         bool
}

var LastMockArgsBatchDefaultTimestamp MockArgsTypeBatchDefaultTimestamp

// (recvb *Batch)AuxMockDefaultTimestamp(argenable bool)(reta *Batch) - Generated mock function
func (recvb *Batch) AuxMockDefaultTimestamp(argenable bool) (reta *Batch) {
	LastMockArgsBatchDefaultTimestamp = MockArgsTypeBatchDefaultTimestamp{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrBatchDefaultTimestamp(),
		Argenable:         argenable,
	}
	rargs, rerr := apomock.GetNext("gocql.Batch.DefaultTimestamp")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Batch.DefaultTimestamp")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Batch.DefaultTimestamp")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Batch)
	}
	return
}

// RecorderAuxMockPtrBatchDefaultTimestamp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchDefaultTimestamp int = 0

var condRecorderAuxMockPtrBatchDefaultTimestamp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchDefaultTimestamp(i int) {
	condRecorderAuxMockPtrBatchDefaultTimestamp.L.Lock()
	for recorderAuxMockPtrBatchDefaultTimestamp < i {
		condRecorderAuxMockPtrBatchDefaultTimestamp.Wait()
	}
	condRecorderAuxMockPtrBatchDefaultTimestamp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchDefaultTimestamp() {
	condRecorderAuxMockPtrBatchDefaultTimestamp.L.Lock()
	recorderAuxMockPtrBatchDefaultTimestamp++
	condRecorderAuxMockPtrBatchDefaultTimestamp.L.Unlock()
	condRecorderAuxMockPtrBatchDefaultTimestamp.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchDefaultTimestamp() (ret int) {
	condRecorderAuxMockPtrBatchDefaultTimestamp.L.Lock()
	ret = recorderAuxMockPtrBatchDefaultTimestamp
	condRecorderAuxMockPtrBatchDefaultTimestamp.L.Unlock()
	return
}

// (recvb *Batch)DefaultTimestamp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) DefaultTimestamp(argenable bool) (reta *Batch) {
	FuncAuxMockPtrBatchDefaultTimestamp, ok := apomock.GetRegisteredFunc("gocql.Batch.DefaultTimestamp")
	if ok {
		reta = FuncAuxMockPtrBatchDefaultTimestamp.(func(recvb *Batch, argenable bool) (reta *Batch))(recvb, argenable)
	} else {
		panic("FuncAuxMockPtrBatchDefaultTimestamp ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchDefaultTimestamp()
	return
}

//
// Mock: NewErrProtocol(argformat string, args ...interface{})(reta error)
//

type MockArgsTypeNewErrProtocol struct {
	ApomockCallNumber int
	Argformat         string
	Args              []interface{}
}

var LastMockArgsNewErrProtocol MockArgsTypeNewErrProtocol

// AuxMockNewErrProtocol(argformat string, args ...interface{})(reta error) - Generated mock function
func AuxMockNewErrProtocol(argformat string, args ...interface{}) (reta error) {
	LastMockArgsNewErrProtocol = MockArgsTypeNewErrProtocol{
		ApomockCallNumber: AuxMockGetRecorderAuxMockNewErrProtocol(),
		Argformat:         argformat,
		Args:              args,
	}
	rargs, rerr := apomock.GetNext("gocql.NewErrProtocol")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.NewErrProtocol")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.NewErrProtocol")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockNewErrProtocol  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNewErrProtocol int = 0

var condRecorderAuxMockNewErrProtocol *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNewErrProtocol(i int) {
	condRecorderAuxMockNewErrProtocol.L.Lock()
	for recorderAuxMockNewErrProtocol < i {
		condRecorderAuxMockNewErrProtocol.Wait()
	}
	condRecorderAuxMockNewErrProtocol.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNewErrProtocol() {
	condRecorderAuxMockNewErrProtocol.L.Lock()
	recorderAuxMockNewErrProtocol++
	condRecorderAuxMockNewErrProtocol.L.Unlock()
	condRecorderAuxMockNewErrProtocol.Broadcast()
}
func AuxMockGetRecorderAuxMockNewErrProtocol() (ret int) {
	condRecorderAuxMockNewErrProtocol.L.Lock()
	ret = recorderAuxMockNewErrProtocol
	condRecorderAuxMockNewErrProtocol.L.Unlock()
	return
}

// NewErrProtocol - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func NewErrProtocol(argformat string, args ...interface{}) (reta error) {
	FuncAuxMockNewErrProtocol, ok := apomock.GetRegisteredFunc("gocql.NewErrProtocol")
	if ok {
		reta = FuncAuxMockNewErrProtocol.(func(argformat string, args ...interface{}) (reta error))(argformat, args...)
	} else {
		panic("FuncAuxMockNewErrProtocol ")
	}
	AuxMockIncrementRecorderAuxMockNewErrProtocol()
	return
}

//
// Mock: (recvs *Session)MapExecuteBatchCAS(argbatch *Batch, argdest map[string]interface{})(retapplied bool, retiter *Iter, reterr error)
//

type MockArgsTypeSessionMapExecuteBatchCAS struct {
	ApomockCallNumber int
	Argbatch          *Batch
	Argdest           map[string]interface{}
}

var LastMockArgsSessionMapExecuteBatchCAS MockArgsTypeSessionMapExecuteBatchCAS

// (recvs *Session)AuxMockMapExecuteBatchCAS(argbatch *Batch, argdest map[string]interface{})(retapplied bool, retiter *Iter, reterr error) - Generated mock function
func (recvs *Session) AuxMockMapExecuteBatchCAS(argbatch *Batch, argdest map[string]interface{}) (retapplied bool, retiter *Iter, reterr error) {
	LastMockArgsSessionMapExecuteBatchCAS = MockArgsTypeSessionMapExecuteBatchCAS{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionMapExecuteBatchCAS(),
		Argbatch:          argbatch,
		Argdest:           argdest,
	}
	rargs, rerr := apomock.GetNext("gocql.Session.MapExecuteBatchCAS")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.MapExecuteBatchCAS")
	} else if rargs.NumArgs() != 3 {
		panic("All return parameters not provided for method:gocql.Session.MapExecuteBatchCAS")
	}
	if rargs.GetArg(0) != nil {
		retapplied = rargs.GetArg(0).(bool)
	}
	if rargs.GetArg(1) != nil {
		retiter = rargs.GetArg(1).(*Iter)
	}
	if rargs.GetArg(2) != nil {
		reterr = rargs.GetArg(2).(error)
	}
	return
}

// RecorderAuxMockPtrSessionMapExecuteBatchCAS  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionMapExecuteBatchCAS int = 0

var condRecorderAuxMockPtrSessionMapExecuteBatchCAS *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionMapExecuteBatchCAS(i int) {
	condRecorderAuxMockPtrSessionMapExecuteBatchCAS.L.Lock()
	for recorderAuxMockPtrSessionMapExecuteBatchCAS < i {
		condRecorderAuxMockPtrSessionMapExecuteBatchCAS.Wait()
	}
	condRecorderAuxMockPtrSessionMapExecuteBatchCAS.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionMapExecuteBatchCAS() {
	condRecorderAuxMockPtrSessionMapExecuteBatchCAS.L.Lock()
	recorderAuxMockPtrSessionMapExecuteBatchCAS++
	condRecorderAuxMockPtrSessionMapExecuteBatchCAS.L.Unlock()
	condRecorderAuxMockPtrSessionMapExecuteBatchCAS.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionMapExecuteBatchCAS() (ret int) {
	condRecorderAuxMockPtrSessionMapExecuteBatchCAS.L.Lock()
	ret = recorderAuxMockPtrSessionMapExecuteBatchCAS
	condRecorderAuxMockPtrSessionMapExecuteBatchCAS.L.Unlock()
	return
}

// (recvs *Session)MapExecuteBatchCAS - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) MapExecuteBatchCAS(argbatch *Batch, argdest map[string]interface{}) (retapplied bool, retiter *Iter, reterr error) {
	FuncAuxMockPtrSessionMapExecuteBatchCAS, ok := apomock.GetRegisteredFunc("gocql.Session.MapExecuteBatchCAS")
	if ok {
		retapplied, retiter, reterr = FuncAuxMockPtrSessionMapExecuteBatchCAS.(func(recvs *Session, argbatch *Batch, argdest map[string]interface{}) (retapplied bool, retiter *Iter, reterr error))(recvs, argbatch, argdest)
	} else {
		panic("FuncAuxMockPtrSessionMapExecuteBatchCAS ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionMapExecuteBatchCAS()
	return
}

//
// Mock: (recvq *Query)GetConsistency()(reta Consistency)
//

type MockArgsTypeQueryGetConsistency struct {
	ApomockCallNumber int
}

var LastMockArgsQueryGetConsistency MockArgsTypeQueryGetConsistency

// (recvq *Query)AuxMockGetConsistency()(reta Consistency) - Generated mock function
func (recvq *Query) AuxMockGetConsistency() (reta Consistency) {
	rargs, rerr := apomock.GetNext("gocql.Query.GetConsistency")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.GetConsistency")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.GetConsistency")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(Consistency)
	}
	return
}

// RecorderAuxMockPtrQueryGetConsistency  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryGetConsistency int = 0

var condRecorderAuxMockPtrQueryGetConsistency *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryGetConsistency(i int) {
	condRecorderAuxMockPtrQueryGetConsistency.L.Lock()
	for recorderAuxMockPtrQueryGetConsistency < i {
		condRecorderAuxMockPtrQueryGetConsistency.Wait()
	}
	condRecorderAuxMockPtrQueryGetConsistency.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryGetConsistency() {
	condRecorderAuxMockPtrQueryGetConsistency.L.Lock()
	recorderAuxMockPtrQueryGetConsistency++
	condRecorderAuxMockPtrQueryGetConsistency.L.Unlock()
	condRecorderAuxMockPtrQueryGetConsistency.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryGetConsistency() (ret int) {
	condRecorderAuxMockPtrQueryGetConsistency.L.Lock()
	ret = recorderAuxMockPtrQueryGetConsistency
	condRecorderAuxMockPtrQueryGetConsistency.L.Unlock()
	return
}

// (recvq *Query)GetConsistency - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) GetConsistency() (reta Consistency) {
	FuncAuxMockPtrQueryGetConsistency, ok := apomock.GetRegisteredFunc("gocql.Query.GetConsistency")
	if ok {
		reta = FuncAuxMockPtrQueryGetConsistency.(func(recvq *Query) (reta Consistency))(recvq)
	} else {
		panic("FuncAuxMockPtrQueryGetConsistency ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryGetConsistency()
	return
}

//
// Mock: (recvq *Query)NoSkipMetadata()(reta *Query)
//

type MockArgsTypeQueryNoSkipMetadata struct {
	ApomockCallNumber int
}

var LastMockArgsQueryNoSkipMetadata MockArgsTypeQueryNoSkipMetadata

// (recvq *Query)AuxMockNoSkipMetadata()(reta *Query) - Generated mock function
func (recvq *Query) AuxMockNoSkipMetadata() (reta *Query) {
	rargs, rerr := apomock.GetNext("gocql.Query.NoSkipMetadata")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.NoSkipMetadata")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.NoSkipMetadata")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryNoSkipMetadata  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryNoSkipMetadata int = 0

var condRecorderAuxMockPtrQueryNoSkipMetadata *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryNoSkipMetadata(i int) {
	condRecorderAuxMockPtrQueryNoSkipMetadata.L.Lock()
	for recorderAuxMockPtrQueryNoSkipMetadata < i {
		condRecorderAuxMockPtrQueryNoSkipMetadata.Wait()
	}
	condRecorderAuxMockPtrQueryNoSkipMetadata.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryNoSkipMetadata() {
	condRecorderAuxMockPtrQueryNoSkipMetadata.L.Lock()
	recorderAuxMockPtrQueryNoSkipMetadata++
	condRecorderAuxMockPtrQueryNoSkipMetadata.L.Unlock()
	condRecorderAuxMockPtrQueryNoSkipMetadata.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryNoSkipMetadata() (ret int) {
	condRecorderAuxMockPtrQueryNoSkipMetadata.L.Lock()
	ret = recorderAuxMockPtrQueryNoSkipMetadata
	condRecorderAuxMockPtrQueryNoSkipMetadata.L.Unlock()
	return
}

// (recvq *Query)NoSkipMetadata - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) NoSkipMetadata() (reta *Query) {
	FuncAuxMockPtrQueryNoSkipMetadata, ok := apomock.GetRegisteredFunc("gocql.Query.NoSkipMetadata")
	if ok {
		reta = FuncAuxMockPtrQueryNoSkipMetadata.(func(recvq *Query) (reta *Query))(recvq)
	} else {
		panic("FuncAuxMockPtrQueryNoSkipMetadata ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryNoSkipMetadata()
	return
}

//
// Mock: (recvq *Query)Iter()(reta *Iter)
//

type MockArgsTypeQueryIter struct {
	ApomockCallNumber int
}

var LastMockArgsQueryIter MockArgsTypeQueryIter

// (recvq *Query)AuxMockIter()(reta *Iter) - Generated mock function
func (recvq *Query) AuxMockIter() (reta *Iter) {
	rargs, rerr := apomock.GetNext("gocql.Query.Iter")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.Iter")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.Iter")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Iter)
	}
	return
}

// RecorderAuxMockPtrQueryIter  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryIter int = 0

var condRecorderAuxMockPtrQueryIter *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryIter(i int) {
	condRecorderAuxMockPtrQueryIter.L.Lock()
	for recorderAuxMockPtrQueryIter < i {
		condRecorderAuxMockPtrQueryIter.Wait()
	}
	condRecorderAuxMockPtrQueryIter.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryIter() {
	condRecorderAuxMockPtrQueryIter.L.Lock()
	recorderAuxMockPtrQueryIter++
	condRecorderAuxMockPtrQueryIter.L.Unlock()
	condRecorderAuxMockPtrQueryIter.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryIter() (ret int) {
	condRecorderAuxMockPtrQueryIter.L.Lock()
	ret = recorderAuxMockPtrQueryIter
	condRecorderAuxMockPtrQueryIter.L.Unlock()
	return
}

// (recvq *Query)Iter - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) Iter() (reta *Iter) {
	FuncAuxMockPtrQueryIter, ok := apomock.GetRegisteredFunc("gocql.Query.Iter")
	if ok {
		reta = FuncAuxMockPtrQueryIter.(func(recvq *Query) (reta *Iter))(recvq)
	} else {
		panic("FuncAuxMockPtrQueryIter ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryIter()
	return
}

//
// Mock: (recviter *Iter)Columns()(reta []ColumnInfo)
//

type MockArgsTypeIterColumns struct {
	ApomockCallNumber int
}

var LastMockArgsIterColumns MockArgsTypeIterColumns

// (recviter *Iter)AuxMockColumns()(reta []ColumnInfo) - Generated mock function
func (recviter *Iter) AuxMockColumns() (reta []ColumnInfo) {
	rargs, rerr := apomock.GetNext("gocql.Iter.Columns")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.Columns")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Iter.Columns")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]ColumnInfo)
	}
	return
}

// RecorderAuxMockPtrIterColumns  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterColumns int = 0

var condRecorderAuxMockPtrIterColumns *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterColumns(i int) {
	condRecorderAuxMockPtrIterColumns.L.Lock()
	for recorderAuxMockPtrIterColumns < i {
		condRecorderAuxMockPtrIterColumns.Wait()
	}
	condRecorderAuxMockPtrIterColumns.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterColumns() {
	condRecorderAuxMockPtrIterColumns.L.Lock()
	recorderAuxMockPtrIterColumns++
	condRecorderAuxMockPtrIterColumns.L.Unlock()
	condRecorderAuxMockPtrIterColumns.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterColumns() (ret int) {
	condRecorderAuxMockPtrIterColumns.L.Lock()
	ret = recorderAuxMockPtrIterColumns
	condRecorderAuxMockPtrIterColumns.L.Unlock()
	return
}

// (recviter *Iter)Columns - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) Columns() (reta []ColumnInfo) {
	FuncAuxMockPtrIterColumns, ok := apomock.GetRegisteredFunc("gocql.Iter.Columns")
	if ok {
		reta = FuncAuxMockPtrIterColumns.(func(recviter *Iter) (reta []ColumnInfo))(recviter)
	} else {
		panic("FuncAuxMockPtrIterColumns ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterColumns()
	return
}

//
// Mock: (recviter *Iter)WillSwitchPage()(reta bool)
//

type MockArgsTypeIterWillSwitchPage struct {
	ApomockCallNumber int
}

var LastMockArgsIterWillSwitchPage MockArgsTypeIterWillSwitchPage

// (recviter *Iter)AuxMockWillSwitchPage()(reta bool) - Generated mock function
func (recviter *Iter) AuxMockWillSwitchPage() (reta bool) {
	rargs, rerr := apomock.GetNext("gocql.Iter.WillSwitchPage")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.WillSwitchPage")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Iter.WillSwitchPage")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrIterWillSwitchPage  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterWillSwitchPage int = 0

var condRecorderAuxMockPtrIterWillSwitchPage *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterWillSwitchPage(i int) {
	condRecorderAuxMockPtrIterWillSwitchPage.L.Lock()
	for recorderAuxMockPtrIterWillSwitchPage < i {
		condRecorderAuxMockPtrIterWillSwitchPage.Wait()
	}
	condRecorderAuxMockPtrIterWillSwitchPage.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterWillSwitchPage() {
	condRecorderAuxMockPtrIterWillSwitchPage.L.Lock()
	recorderAuxMockPtrIterWillSwitchPage++
	condRecorderAuxMockPtrIterWillSwitchPage.L.Unlock()
	condRecorderAuxMockPtrIterWillSwitchPage.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterWillSwitchPage() (ret int) {
	condRecorderAuxMockPtrIterWillSwitchPage.L.Lock()
	ret = recorderAuxMockPtrIterWillSwitchPage
	condRecorderAuxMockPtrIterWillSwitchPage.L.Unlock()
	return
}

// (recviter *Iter)WillSwitchPage - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) WillSwitchPage() (reta bool) {
	FuncAuxMockPtrIterWillSwitchPage, ok := apomock.GetRegisteredFunc("gocql.Iter.WillSwitchPage")
	if ok {
		reta = FuncAuxMockPtrIterWillSwitchPage.(func(recviter *Iter) (reta bool))(recviter)
	} else {
		panic("FuncAuxMockPtrIterWillSwitchPage ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterWillSwitchPage()
	return
}

//
// Mock: (recvb *Batch)GetConsistency()(reta Consistency)
//

type MockArgsTypeBatchGetConsistency struct {
	ApomockCallNumber int
}

var LastMockArgsBatchGetConsistency MockArgsTypeBatchGetConsistency

// (recvb *Batch)AuxMockGetConsistency()(reta Consistency) - Generated mock function
func (recvb *Batch) AuxMockGetConsistency() (reta Consistency) {
	rargs, rerr := apomock.GetNext("gocql.Batch.GetConsistency")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Batch.GetConsistency")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Batch.GetConsistency")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(Consistency)
	}
	return
}

// RecorderAuxMockPtrBatchGetConsistency  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchGetConsistency int = 0

var condRecorderAuxMockPtrBatchGetConsistency *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchGetConsistency(i int) {
	condRecorderAuxMockPtrBatchGetConsistency.L.Lock()
	for recorderAuxMockPtrBatchGetConsistency < i {
		condRecorderAuxMockPtrBatchGetConsistency.Wait()
	}
	condRecorderAuxMockPtrBatchGetConsistency.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchGetConsistency() {
	condRecorderAuxMockPtrBatchGetConsistency.L.Lock()
	recorderAuxMockPtrBatchGetConsistency++
	condRecorderAuxMockPtrBatchGetConsistency.L.Unlock()
	condRecorderAuxMockPtrBatchGetConsistency.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchGetConsistency() (ret int) {
	condRecorderAuxMockPtrBatchGetConsistency.L.Lock()
	ret = recorderAuxMockPtrBatchGetConsistency
	condRecorderAuxMockPtrBatchGetConsistency.L.Unlock()
	return
}

// (recvb *Batch)GetConsistency - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) GetConsistency() (reta Consistency) {
	FuncAuxMockPtrBatchGetConsistency, ok := apomock.GetRegisteredFunc("gocql.Batch.GetConsistency")
	if ok {
		reta = FuncAuxMockPtrBatchGetConsistency.(func(recvb *Batch) (reta Consistency))(recvb)
	} else {
		panic("FuncAuxMockPtrBatchGetConsistency ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchGetConsistency()
	return
}

//
// Mock: (recvb *Batch)Query(argstmt string, args ...interface{})()
//

type MockArgsTypeBatchQuery struct {
	ApomockCallNumber int
	Argstmt           string
	Args              []interface{}
}

var LastMockArgsBatchQuery MockArgsTypeBatchQuery

// (recvb *Batch)AuxMockQuery(argstmt string, args ...interface{})() - Generated mock function
func (recvb *Batch) AuxMockQuery(argstmt string, args ...interface{}) {
	LastMockArgsBatchQuery = MockArgsTypeBatchQuery{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrBatchQuery(),
		Argstmt:           argstmt,
		Args:              args,
	}
	return
}

// RecorderAuxMockPtrBatchQuery  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchQuery int = 0

var condRecorderAuxMockPtrBatchQuery *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchQuery(i int) {
	condRecorderAuxMockPtrBatchQuery.L.Lock()
	for recorderAuxMockPtrBatchQuery < i {
		condRecorderAuxMockPtrBatchQuery.Wait()
	}
	condRecorderAuxMockPtrBatchQuery.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchQuery() {
	condRecorderAuxMockPtrBatchQuery.L.Lock()
	recorderAuxMockPtrBatchQuery++
	condRecorderAuxMockPtrBatchQuery.L.Unlock()
	condRecorderAuxMockPtrBatchQuery.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchQuery() (ret int) {
	condRecorderAuxMockPtrBatchQuery.L.Lock()
	ret = recorderAuxMockPtrBatchQuery
	condRecorderAuxMockPtrBatchQuery.L.Unlock()
	return
}

// (recvb *Batch)Query - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) Query(argstmt string, args ...interface{}) {
	FuncAuxMockPtrBatchQuery, ok := apomock.GetRegisteredFunc("gocql.Batch.Query")
	if ok {
		FuncAuxMockPtrBatchQuery.(func(recvb *Batch, argstmt string, args ...interface{}))(recvb, argstmt, args...)
	} else {
		panic("FuncAuxMockPtrBatchQuery ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchQuery()
	return
}

//
// Mock: (recvs *Session)executeBatch(argbatch *Batch)(reta *Iter)
//

type MockArgsTypeSessionexecuteBatch struct {
	ApomockCallNumber int
	Argbatch          *Batch
}

var LastMockArgsSessionexecuteBatch MockArgsTypeSessionexecuteBatch

// (recvs *Session)AuxMockexecuteBatch(argbatch *Batch)(reta *Iter) - Generated mock function
func (recvs *Session) AuxMockexecuteBatch(argbatch *Batch) (reta *Iter) {
	LastMockArgsSessionexecuteBatch = MockArgsTypeSessionexecuteBatch{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionexecuteBatch(),
		Argbatch:          argbatch,
	}
	rargs, rerr := apomock.GetNext("gocql.Session.executeBatch")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.executeBatch")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Session.executeBatch")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Iter)
	}
	return
}

// RecorderAuxMockPtrSessionexecuteBatch  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionexecuteBatch int = 0

var condRecorderAuxMockPtrSessionexecuteBatch *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionexecuteBatch(i int) {
	condRecorderAuxMockPtrSessionexecuteBatch.L.Lock()
	for recorderAuxMockPtrSessionexecuteBatch < i {
		condRecorderAuxMockPtrSessionexecuteBatch.Wait()
	}
	condRecorderAuxMockPtrSessionexecuteBatch.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionexecuteBatch() {
	condRecorderAuxMockPtrSessionexecuteBatch.L.Lock()
	recorderAuxMockPtrSessionexecuteBatch++
	condRecorderAuxMockPtrSessionexecuteBatch.L.Unlock()
	condRecorderAuxMockPtrSessionexecuteBatch.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionexecuteBatch() (ret int) {
	condRecorderAuxMockPtrSessionexecuteBatch.L.Lock()
	ret = recorderAuxMockPtrSessionexecuteBatch
	condRecorderAuxMockPtrSessionexecuteBatch.L.Unlock()
	return
}

// (recvs *Session)executeBatch - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) executeBatch(argbatch *Batch) (reta *Iter) {
	FuncAuxMockPtrSessionexecuteBatch, ok := apomock.GetRegisteredFunc("gocql.Session.executeBatch")
	if ok {
		reta = FuncAuxMockPtrSessionexecuteBatch.(func(recvs *Session, argbatch *Batch) (reta *Iter))(recvs, argbatch)
	} else {
		panic("FuncAuxMockPtrSessionexecuteBatch ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionexecuteBatch()
	return
}

//
// Mock: (recvs *Session)Closed()(reta bool)
//

type MockArgsTypeSessionClosed struct {
	ApomockCallNumber int
}

var LastMockArgsSessionClosed MockArgsTypeSessionClosed

// (recvs *Session)AuxMockClosed()(reta bool) - Generated mock function
func (recvs *Session) AuxMockClosed() (reta bool) {
	rargs, rerr := apomock.GetNext("gocql.Session.Closed")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.Closed")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Session.Closed")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrSessionClosed  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionClosed int = 0

var condRecorderAuxMockPtrSessionClosed *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionClosed(i int) {
	condRecorderAuxMockPtrSessionClosed.L.Lock()
	for recorderAuxMockPtrSessionClosed < i {
		condRecorderAuxMockPtrSessionClosed.Wait()
	}
	condRecorderAuxMockPtrSessionClosed.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionClosed() {
	condRecorderAuxMockPtrSessionClosed.L.Lock()
	recorderAuxMockPtrSessionClosed++
	condRecorderAuxMockPtrSessionClosed.L.Unlock()
	condRecorderAuxMockPtrSessionClosed.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionClosed() (ret int) {
	condRecorderAuxMockPtrSessionClosed.L.Lock()
	ret = recorderAuxMockPtrSessionClosed
	condRecorderAuxMockPtrSessionClosed.L.Unlock()
	return
}

// (recvs *Session)Closed - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) Closed() (reta bool) {
	FuncAuxMockPtrSessionClosed, ok := apomock.GetRegisteredFunc("gocql.Session.Closed")
	if ok {
		reta = FuncAuxMockPtrSessionClosed.(func(recvs *Session) (reta bool))(recvs)
	} else {
		panic("FuncAuxMockPtrSessionClosed ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionClosed()
	return
}

//
// Mock: (recvs *Session)KeyspaceMetadata(argkeyspace string)(reta *KeyspaceMetadata, retb error)
//

type MockArgsTypeSessionKeyspaceMetadata struct {
	ApomockCallNumber int
	Argkeyspace       string
}

var LastMockArgsSessionKeyspaceMetadata MockArgsTypeSessionKeyspaceMetadata

// (recvs *Session)AuxMockKeyspaceMetadata(argkeyspace string)(reta *KeyspaceMetadata, retb error) - Generated mock function
func (recvs *Session) AuxMockKeyspaceMetadata(argkeyspace string) (reta *KeyspaceMetadata, retb error) {
	LastMockArgsSessionKeyspaceMetadata = MockArgsTypeSessionKeyspaceMetadata{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionKeyspaceMetadata(),
		Argkeyspace:       argkeyspace,
	}
	rargs, rerr := apomock.GetNext("gocql.Session.KeyspaceMetadata")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.KeyspaceMetadata")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Session.KeyspaceMetadata")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*KeyspaceMetadata)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrSessionKeyspaceMetadata  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionKeyspaceMetadata int = 0

var condRecorderAuxMockPtrSessionKeyspaceMetadata *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionKeyspaceMetadata(i int) {
	condRecorderAuxMockPtrSessionKeyspaceMetadata.L.Lock()
	for recorderAuxMockPtrSessionKeyspaceMetadata < i {
		condRecorderAuxMockPtrSessionKeyspaceMetadata.Wait()
	}
	condRecorderAuxMockPtrSessionKeyspaceMetadata.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionKeyspaceMetadata() {
	condRecorderAuxMockPtrSessionKeyspaceMetadata.L.Lock()
	recorderAuxMockPtrSessionKeyspaceMetadata++
	condRecorderAuxMockPtrSessionKeyspaceMetadata.L.Unlock()
	condRecorderAuxMockPtrSessionKeyspaceMetadata.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionKeyspaceMetadata() (ret int) {
	condRecorderAuxMockPtrSessionKeyspaceMetadata.L.Lock()
	ret = recorderAuxMockPtrSessionKeyspaceMetadata
	condRecorderAuxMockPtrSessionKeyspaceMetadata.L.Unlock()
	return
}

// (recvs *Session)KeyspaceMetadata - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) KeyspaceMetadata(argkeyspace string) (reta *KeyspaceMetadata, retb error) {
	FuncAuxMockPtrSessionKeyspaceMetadata, ok := apomock.GetRegisteredFunc("gocql.Session.KeyspaceMetadata")
	if ok {
		reta, retb = FuncAuxMockPtrSessionKeyspaceMetadata.(func(recvs *Session, argkeyspace string) (reta *KeyspaceMetadata, retb error))(recvs, argkeyspace)
	} else {
		panic("FuncAuxMockPtrSessionKeyspaceMetadata ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionKeyspaceMetadata()
	return
}

//
// Mock: (recvq *Query)DefaultTimestamp(argenable bool)(reta *Query)
//

type MockArgsTypeQueryDefaultTimestamp struct {
	ApomockCallNumber int
	Argenable         bool
}

var LastMockArgsQueryDefaultTimestamp MockArgsTypeQueryDefaultTimestamp

// (recvq *Query)AuxMockDefaultTimestamp(argenable bool)(reta *Query) - Generated mock function
func (recvq *Query) AuxMockDefaultTimestamp(argenable bool) (reta *Query) {
	LastMockArgsQueryDefaultTimestamp = MockArgsTypeQueryDefaultTimestamp{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryDefaultTimestamp(),
		Argenable:         argenable,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.DefaultTimestamp")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.DefaultTimestamp")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.DefaultTimestamp")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryDefaultTimestamp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryDefaultTimestamp int = 0

var condRecorderAuxMockPtrQueryDefaultTimestamp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryDefaultTimestamp(i int) {
	condRecorderAuxMockPtrQueryDefaultTimestamp.L.Lock()
	for recorderAuxMockPtrQueryDefaultTimestamp < i {
		condRecorderAuxMockPtrQueryDefaultTimestamp.Wait()
	}
	condRecorderAuxMockPtrQueryDefaultTimestamp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryDefaultTimestamp() {
	condRecorderAuxMockPtrQueryDefaultTimestamp.L.Lock()
	recorderAuxMockPtrQueryDefaultTimestamp++
	condRecorderAuxMockPtrQueryDefaultTimestamp.L.Unlock()
	condRecorderAuxMockPtrQueryDefaultTimestamp.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryDefaultTimestamp() (ret int) {
	condRecorderAuxMockPtrQueryDefaultTimestamp.L.Lock()
	ret = recorderAuxMockPtrQueryDefaultTimestamp
	condRecorderAuxMockPtrQueryDefaultTimestamp.L.Unlock()
	return
}

// (recvq *Query)DefaultTimestamp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) DefaultTimestamp(argenable bool) (reta *Query) {
	FuncAuxMockPtrQueryDefaultTimestamp, ok := apomock.GetRegisteredFunc("gocql.Query.DefaultTimestamp")
	if ok {
		reta = FuncAuxMockPtrQueryDefaultTimestamp.(func(recvq *Query, argenable bool) (reta *Query))(recvq, argenable)
	} else {
		panic("FuncAuxMockPtrQueryDefaultTimestamp ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryDefaultTimestamp()
	return
}

//
// Mock: (recvq *Query)Bind(v ...interface{})(reta *Query)
//

type MockArgsTypeQueryBind struct {
	ApomockCallNumber int
	V                 []interface{}
}

var LastMockArgsQueryBind MockArgsTypeQueryBind

// (recvq *Query)AuxMockBind(v ...interface{})(reta *Query) - Generated mock function
func (recvq *Query) AuxMockBind(v ...interface{}) (reta *Query) {
	LastMockArgsQueryBind = MockArgsTypeQueryBind{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryBind(),
		V:                 v,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.Bind")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.Bind")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.Bind")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryBind  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryBind int = 0

var condRecorderAuxMockPtrQueryBind *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryBind(i int) {
	condRecorderAuxMockPtrQueryBind.L.Lock()
	for recorderAuxMockPtrQueryBind < i {
		condRecorderAuxMockPtrQueryBind.Wait()
	}
	condRecorderAuxMockPtrQueryBind.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryBind() {
	condRecorderAuxMockPtrQueryBind.L.Lock()
	recorderAuxMockPtrQueryBind++
	condRecorderAuxMockPtrQueryBind.L.Unlock()
	condRecorderAuxMockPtrQueryBind.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryBind() (ret int) {
	condRecorderAuxMockPtrQueryBind.L.Lock()
	ret = recorderAuxMockPtrQueryBind
	condRecorderAuxMockPtrQueryBind.L.Unlock()
	return
}

// (recvq *Query)Bind - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) Bind(v ...interface{}) (reta *Query) {
	FuncAuxMockPtrQueryBind, ok := apomock.GetRegisteredFunc("gocql.Query.Bind")
	if ok {
		reta = FuncAuxMockPtrQueryBind.(func(recvq *Query, v ...interface{}) (reta *Query))(recvq, v...)
	} else {
		panic("FuncAuxMockPtrQueryBind ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryBind()
	return
}

//
// Mock: (recvb *Batch)Latency()(reta int64)
//

type MockArgsTypeBatchLatency struct {
	ApomockCallNumber int
}

var LastMockArgsBatchLatency MockArgsTypeBatchLatency

// (recvb *Batch)AuxMockLatency()(reta int64) - Generated mock function
func (recvb *Batch) AuxMockLatency() (reta int64) {
	rargs, rerr := apomock.GetNext("gocql.Batch.Latency")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Batch.Latency")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Batch.Latency")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int64)
	}
	return
}

// RecorderAuxMockPtrBatchLatency  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchLatency int = 0

var condRecorderAuxMockPtrBatchLatency *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchLatency(i int) {
	condRecorderAuxMockPtrBatchLatency.L.Lock()
	for recorderAuxMockPtrBatchLatency < i {
		condRecorderAuxMockPtrBatchLatency.Wait()
	}
	condRecorderAuxMockPtrBatchLatency.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchLatency() {
	condRecorderAuxMockPtrBatchLatency.L.Lock()
	recorderAuxMockPtrBatchLatency++
	condRecorderAuxMockPtrBatchLatency.L.Unlock()
	condRecorderAuxMockPtrBatchLatency.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchLatency() (ret int) {
	condRecorderAuxMockPtrBatchLatency.L.Lock()
	ret = recorderAuxMockPtrBatchLatency
	condRecorderAuxMockPtrBatchLatency.L.Unlock()
	return
}

// (recvb *Batch)Latency - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) Latency() (reta int64) {
	FuncAuxMockPtrBatchLatency, ok := apomock.GetRegisteredFunc("gocql.Batch.Latency")
	if ok {
		reta = FuncAuxMockPtrBatchLatency.(func(recvb *Batch) (reta int64))(recvb)
	} else {
		panic("FuncAuxMockPtrBatchLatency ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchLatency()
	return
}

//
// Mock: (recvb *Batch)Bind(argstmt string, argbind func(*QueryInfo) ([]interface{}, error))()
//

type MockArgsTypeBatchBind struct {
	ApomockCallNumber int
	Argstmt           string
	Argbind           func(*QueryInfo) ([]interface{}, error)
}

var LastMockArgsBatchBind MockArgsTypeBatchBind

// (recvb *Batch)AuxMockBind(argstmt string, argbind func(*QueryInfo) ([]interface{}, error))() - Generated mock function
func (recvb *Batch) AuxMockBind(argstmt string, argbind func(*QueryInfo) ([]interface{}, error)) {
	LastMockArgsBatchBind = MockArgsTypeBatchBind{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrBatchBind(),
		Argstmt:           argstmt,
		Argbind:           argbind,
	}
	return
}

// RecorderAuxMockPtrBatchBind  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchBind int = 0

var condRecorderAuxMockPtrBatchBind *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchBind(i int) {
	condRecorderAuxMockPtrBatchBind.L.Lock()
	for recorderAuxMockPtrBatchBind < i {
		condRecorderAuxMockPtrBatchBind.Wait()
	}
	condRecorderAuxMockPtrBatchBind.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchBind() {
	condRecorderAuxMockPtrBatchBind.L.Lock()
	recorderAuxMockPtrBatchBind++
	condRecorderAuxMockPtrBatchBind.L.Unlock()
	condRecorderAuxMockPtrBatchBind.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchBind() (ret int) {
	condRecorderAuxMockPtrBatchBind.L.Lock()
	ret = recorderAuxMockPtrBatchBind
	condRecorderAuxMockPtrBatchBind.L.Unlock()
	return
}

// (recvb *Batch)Bind - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) Bind(argstmt string, argbind func(*QueryInfo) ([]interface{}, error)) {
	FuncAuxMockPtrBatchBind, ok := apomock.GetRegisteredFunc("gocql.Batch.Bind")
	if ok {
		FuncAuxMockPtrBatchBind.(func(recvb *Batch, argstmt string, argbind func(*QueryInfo) ([]interface{}, error)))(recvb, argstmt, argbind)
	} else {
		panic("FuncAuxMockPtrBatchBind ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchBind()
	return
}

//
// Mock: (recvs *Session)SetPrefetch(argp float64)()
//

type MockArgsTypeSessionSetPrefetch struct {
	ApomockCallNumber int
	Argp              float64
}

var LastMockArgsSessionSetPrefetch MockArgsTypeSessionSetPrefetch

// (recvs *Session)AuxMockSetPrefetch(argp float64)() - Generated mock function
func (recvs *Session) AuxMockSetPrefetch(argp float64) {
	LastMockArgsSessionSetPrefetch = MockArgsTypeSessionSetPrefetch{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionSetPrefetch(),
		Argp:              argp,
	}
	return
}

// RecorderAuxMockPtrSessionSetPrefetch  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionSetPrefetch int = 0

var condRecorderAuxMockPtrSessionSetPrefetch *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionSetPrefetch(i int) {
	condRecorderAuxMockPtrSessionSetPrefetch.L.Lock()
	for recorderAuxMockPtrSessionSetPrefetch < i {
		condRecorderAuxMockPtrSessionSetPrefetch.Wait()
	}
	condRecorderAuxMockPtrSessionSetPrefetch.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionSetPrefetch() {
	condRecorderAuxMockPtrSessionSetPrefetch.L.Lock()
	recorderAuxMockPtrSessionSetPrefetch++
	condRecorderAuxMockPtrSessionSetPrefetch.L.Unlock()
	condRecorderAuxMockPtrSessionSetPrefetch.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionSetPrefetch() (ret int) {
	condRecorderAuxMockPtrSessionSetPrefetch.L.Lock()
	ret = recorderAuxMockPtrSessionSetPrefetch
	condRecorderAuxMockPtrSessionSetPrefetch.L.Unlock()
	return
}

// (recvs *Session)SetPrefetch - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) SetPrefetch(argp float64) {
	FuncAuxMockPtrSessionSetPrefetch, ok := apomock.GetRegisteredFunc("gocql.Session.SetPrefetch")
	if ok {
		FuncAuxMockPtrSessionSetPrefetch.(func(recvs *Session, argp float64))(recvs, argp)
	} else {
		panic("FuncAuxMockPtrSessionSetPrefetch ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionSetPrefetch()
	return
}

//
// Mock: (recvs *Session)ExecuteBatchCAS(argbatch *Batch, dest ...interface{})(retapplied bool, retiter *Iter, reterr error)
//

type MockArgsTypeSessionExecuteBatchCAS struct {
	ApomockCallNumber int
	Argbatch          *Batch
	Dest              []interface{}
}

var LastMockArgsSessionExecuteBatchCAS MockArgsTypeSessionExecuteBatchCAS

// (recvs *Session)AuxMockExecuteBatchCAS(argbatch *Batch, dest ...interface{})(retapplied bool, retiter *Iter, reterr error) - Generated mock function
func (recvs *Session) AuxMockExecuteBatchCAS(argbatch *Batch, dest ...interface{}) (retapplied bool, retiter *Iter, reterr error) {
	LastMockArgsSessionExecuteBatchCAS = MockArgsTypeSessionExecuteBatchCAS{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionExecuteBatchCAS(),
		Argbatch:          argbatch,
		Dest:              dest,
	}
	rargs, rerr := apomock.GetNext("gocql.Session.ExecuteBatchCAS")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.ExecuteBatchCAS")
	} else if rargs.NumArgs() != 3 {
		panic("All return parameters not provided for method:gocql.Session.ExecuteBatchCAS")
	}
	if rargs.GetArg(0) != nil {
		retapplied = rargs.GetArg(0).(bool)
	}
	if rargs.GetArg(1) != nil {
		retiter = rargs.GetArg(1).(*Iter)
	}
	if rargs.GetArg(2) != nil {
		reterr = rargs.GetArg(2).(error)
	}
	return
}

// RecorderAuxMockPtrSessionExecuteBatchCAS  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionExecuteBatchCAS int = 0

var condRecorderAuxMockPtrSessionExecuteBatchCAS *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionExecuteBatchCAS(i int) {
	condRecorderAuxMockPtrSessionExecuteBatchCAS.L.Lock()
	for recorderAuxMockPtrSessionExecuteBatchCAS < i {
		condRecorderAuxMockPtrSessionExecuteBatchCAS.Wait()
	}
	condRecorderAuxMockPtrSessionExecuteBatchCAS.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionExecuteBatchCAS() {
	condRecorderAuxMockPtrSessionExecuteBatchCAS.L.Lock()
	recorderAuxMockPtrSessionExecuteBatchCAS++
	condRecorderAuxMockPtrSessionExecuteBatchCAS.L.Unlock()
	condRecorderAuxMockPtrSessionExecuteBatchCAS.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionExecuteBatchCAS() (ret int) {
	condRecorderAuxMockPtrSessionExecuteBatchCAS.L.Lock()
	ret = recorderAuxMockPtrSessionExecuteBatchCAS
	condRecorderAuxMockPtrSessionExecuteBatchCAS.L.Unlock()
	return
}

// (recvs *Session)ExecuteBatchCAS - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) ExecuteBatchCAS(argbatch *Batch, dest ...interface{}) (retapplied bool, retiter *Iter, reterr error) {
	FuncAuxMockPtrSessionExecuteBatchCAS, ok := apomock.GetRegisteredFunc("gocql.Session.ExecuteBatchCAS")
	if ok {
		retapplied, retiter, reterr = FuncAuxMockPtrSessionExecuteBatchCAS.(func(recvs *Session, argbatch *Batch, dest ...interface{}) (retapplied bool, retiter *Iter, reterr error))(recvs, argbatch, dest...)
	} else {
		panic("FuncAuxMockPtrSessionExecuteBatchCAS ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionExecuteBatchCAS()
	return
}

//
// Mock: (recvq *Query)Trace(argtrace Tracer)(reta *Query)
//

type MockArgsTypeQueryTrace struct {
	ApomockCallNumber int
	Argtrace          Tracer
}

var LastMockArgsQueryTrace MockArgsTypeQueryTrace

// (recvq *Query)AuxMockTrace(argtrace Tracer)(reta *Query) - Generated mock function
func (recvq *Query) AuxMockTrace(argtrace Tracer) (reta *Query) {
	LastMockArgsQueryTrace = MockArgsTypeQueryTrace{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryTrace(),
		Argtrace:          argtrace,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.Trace")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.Trace")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.Trace")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryTrace  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryTrace int = 0

var condRecorderAuxMockPtrQueryTrace *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryTrace(i int) {
	condRecorderAuxMockPtrQueryTrace.L.Lock()
	for recorderAuxMockPtrQueryTrace < i {
		condRecorderAuxMockPtrQueryTrace.Wait()
	}
	condRecorderAuxMockPtrQueryTrace.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryTrace() {
	condRecorderAuxMockPtrQueryTrace.L.Lock()
	recorderAuxMockPtrQueryTrace++
	condRecorderAuxMockPtrQueryTrace.L.Unlock()
	condRecorderAuxMockPtrQueryTrace.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryTrace() (ret int) {
	condRecorderAuxMockPtrQueryTrace.L.Lock()
	ret = recorderAuxMockPtrQueryTrace
	condRecorderAuxMockPtrQueryTrace.L.Unlock()
	return
}

// (recvq *Query)Trace - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) Trace(argtrace Tracer) (reta *Query) {
	FuncAuxMockPtrQueryTrace, ok := apomock.GetRegisteredFunc("gocql.Query.Trace")
	if ok {
		reta = FuncAuxMockPtrQueryTrace.(func(recvq *Query, argtrace Tracer) (reta *Query))(recvq, argtrace)
	} else {
		panic("FuncAuxMockPtrQueryTrace ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryTrace()
	return
}

//
// Mock: isUseStatement(argstmt string)(reta bool)
//

type MockArgsTypeisUseStatement struct {
	ApomockCallNumber int
	Argstmt           string
}

var LastMockArgsisUseStatement MockArgsTypeisUseStatement

// AuxMockisUseStatement(argstmt string)(reta bool) - Generated mock function
func AuxMockisUseStatement(argstmt string) (reta bool) {
	LastMockArgsisUseStatement = MockArgsTypeisUseStatement{
		ApomockCallNumber: AuxMockGetRecorderAuxMockisUseStatement(),
		Argstmt:           argstmt,
	}
	rargs, rerr := apomock.GetNext("gocql.isUseStatement")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.isUseStatement")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.isUseStatement")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockisUseStatement  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockisUseStatement int = 0

var condRecorderAuxMockisUseStatement *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockisUseStatement(i int) {
	condRecorderAuxMockisUseStatement.L.Lock()
	for recorderAuxMockisUseStatement < i {
		condRecorderAuxMockisUseStatement.Wait()
	}
	condRecorderAuxMockisUseStatement.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockisUseStatement() {
	condRecorderAuxMockisUseStatement.L.Lock()
	recorderAuxMockisUseStatement++
	condRecorderAuxMockisUseStatement.L.Unlock()
	condRecorderAuxMockisUseStatement.Broadcast()
}
func AuxMockGetRecorderAuxMockisUseStatement() (ret int) {
	condRecorderAuxMockisUseStatement.L.Lock()
	ret = recorderAuxMockisUseStatement
	condRecorderAuxMockisUseStatement.L.Unlock()
	return
}

// isUseStatement - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func isUseStatement(argstmt string) (reta bool) {
	FuncAuxMockisUseStatement, ok := apomock.GetRegisteredFunc("gocql.isUseStatement")
	if ok {
		reta = FuncAuxMockisUseStatement.(func(argstmt string) (reta bool))(argstmt)
	} else {
		panic("FuncAuxMockisUseStatement ")
	}
	AuxMockIncrementRecorderAuxMockisUseStatement()
	return
}

//
// Mock: (recviter *Iter)PageState()(reta []byte)
//

type MockArgsTypeIterPageState struct {
	ApomockCallNumber int
}

var LastMockArgsIterPageState MockArgsTypeIterPageState

// (recviter *Iter)AuxMockPageState()(reta []byte) - Generated mock function
func (recviter *Iter) AuxMockPageState() (reta []byte) {
	rargs, rerr := apomock.GetNext("gocql.Iter.PageState")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.PageState")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Iter.PageState")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockPtrIterPageState  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterPageState int = 0

var condRecorderAuxMockPtrIterPageState *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterPageState(i int) {
	condRecorderAuxMockPtrIterPageState.L.Lock()
	for recorderAuxMockPtrIterPageState < i {
		condRecorderAuxMockPtrIterPageState.Wait()
	}
	condRecorderAuxMockPtrIterPageState.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterPageState() {
	condRecorderAuxMockPtrIterPageState.L.Lock()
	recorderAuxMockPtrIterPageState++
	condRecorderAuxMockPtrIterPageState.L.Unlock()
	condRecorderAuxMockPtrIterPageState.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterPageState() (ret int) {
	condRecorderAuxMockPtrIterPageState.L.Lock()
	ret = recorderAuxMockPtrIterPageState
	condRecorderAuxMockPtrIterPageState.L.Unlock()
	return
}

// (recviter *Iter)PageState - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) PageState() (reta []byte) {
	FuncAuxMockPtrIterPageState, ok := apomock.GetRegisteredFunc("gocql.Iter.PageState")
	if ok {
		reta = FuncAuxMockPtrIterPageState.(func(recviter *Iter) (reta []byte))(recviter)
	} else {
		panic("FuncAuxMockPtrIterPageState ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterPageState()
	return
}

//
// Mock: (recvc ColumnInfo)String()(reta string)
//

type MockArgsTypeColumnInfoString struct {
	ApomockCallNumber int
}

var LastMockArgsColumnInfoString MockArgsTypeColumnInfoString

// (recvc ColumnInfo)AuxMockString()(reta string) - Generated mock function
func (recvc ColumnInfo) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.ColumnInfo.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ColumnInfo.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.ColumnInfo.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockColumnInfoString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockColumnInfoString int = 0

var condRecorderAuxMockColumnInfoString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockColumnInfoString(i int) {
	condRecorderAuxMockColumnInfoString.L.Lock()
	for recorderAuxMockColumnInfoString < i {
		condRecorderAuxMockColumnInfoString.Wait()
	}
	condRecorderAuxMockColumnInfoString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockColumnInfoString() {
	condRecorderAuxMockColumnInfoString.L.Lock()
	recorderAuxMockColumnInfoString++
	condRecorderAuxMockColumnInfoString.L.Unlock()
	condRecorderAuxMockColumnInfoString.Broadcast()
}
func AuxMockGetRecorderAuxMockColumnInfoString() (ret int) {
	condRecorderAuxMockColumnInfoString.L.Lock()
	ret = recorderAuxMockColumnInfoString
	condRecorderAuxMockColumnInfoString.L.Unlock()
	return
}

// (recvc ColumnInfo)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc ColumnInfo) String() (reta string) {
	FuncAuxMockColumnInfoString, ok := apomock.GetRegisteredFunc("gocql.ColumnInfo.String")
	if ok {
		reta = FuncAuxMockColumnInfoString.(func(recvc ColumnInfo) (reta string))(recvc)
	} else {
		panic("FuncAuxMockColumnInfoString ")
	}
	AuxMockIncrementRecorderAuxMockColumnInfoString()
	return
}

//
// Mock: (recve Error)Error()(reta string)
//

type MockArgsTypeErrorError struct {
	ApomockCallNumber int
}

var LastMockArgsErrorError MockArgsTypeErrorError

// (recve Error)AuxMockError()(reta string) - Generated mock function
func (recve Error) AuxMockError() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.Error.Error")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Error.Error")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Error.Error")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockErrorError  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockErrorError int = 0

var condRecorderAuxMockErrorError *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockErrorError(i int) {
	condRecorderAuxMockErrorError.L.Lock()
	for recorderAuxMockErrorError < i {
		condRecorderAuxMockErrorError.Wait()
	}
	condRecorderAuxMockErrorError.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockErrorError() {
	condRecorderAuxMockErrorError.L.Lock()
	recorderAuxMockErrorError++
	condRecorderAuxMockErrorError.L.Unlock()
	condRecorderAuxMockErrorError.Broadcast()
}
func AuxMockGetRecorderAuxMockErrorError() (ret int) {
	condRecorderAuxMockErrorError.L.Lock()
	ret = recorderAuxMockErrorError
	condRecorderAuxMockErrorError.L.Unlock()
	return
}

// (recve Error)Error - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve Error) Error() (reta string) {
	FuncAuxMockErrorError, ok := apomock.GetRegisteredFunc("gocql.Error.Error")
	if ok {
		reta = FuncAuxMockErrorError.(func(recve Error) (reta string))(recve)
	} else {
		panic("FuncAuxMockErrorError ")
	}
	AuxMockIncrementRecorderAuxMockErrorError()
	return
}

//
// Mock: (recvs *Session)SetPageSize(argn int)()
//

type MockArgsTypeSessionSetPageSize struct {
	ApomockCallNumber int
	Argn              int
}

var LastMockArgsSessionSetPageSize MockArgsTypeSessionSetPageSize

// (recvs *Session)AuxMockSetPageSize(argn int)() - Generated mock function
func (recvs *Session) AuxMockSetPageSize(argn int) {
	LastMockArgsSessionSetPageSize = MockArgsTypeSessionSetPageSize{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionSetPageSize(),
		Argn:              argn,
	}
	return
}

// RecorderAuxMockPtrSessionSetPageSize  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionSetPageSize int = 0

var condRecorderAuxMockPtrSessionSetPageSize *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionSetPageSize(i int) {
	condRecorderAuxMockPtrSessionSetPageSize.L.Lock()
	for recorderAuxMockPtrSessionSetPageSize < i {
		condRecorderAuxMockPtrSessionSetPageSize.Wait()
	}
	condRecorderAuxMockPtrSessionSetPageSize.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionSetPageSize() {
	condRecorderAuxMockPtrSessionSetPageSize.L.Lock()
	recorderAuxMockPtrSessionSetPageSize++
	condRecorderAuxMockPtrSessionSetPageSize.L.Unlock()
	condRecorderAuxMockPtrSessionSetPageSize.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionSetPageSize() (ret int) {
	condRecorderAuxMockPtrSessionSetPageSize.L.Lock()
	ret = recorderAuxMockPtrSessionSetPageSize
	condRecorderAuxMockPtrSessionSetPageSize.L.Unlock()
	return
}

// (recvs *Session)SetPageSize - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) SetPageSize(argn int) {
	FuncAuxMockPtrSessionSetPageSize, ok := apomock.GetRegisteredFunc("gocql.Session.SetPageSize")
	if ok {
		FuncAuxMockPtrSessionSetPageSize.(func(recvs *Session, argn int))(recvs, argn)
	} else {
		panic("FuncAuxMockPtrSessionSetPageSize ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionSetPageSize()
	return
}

//
// Mock: (recvq *Query)WithTimestamp(argtimestamp int64)(reta *Query)
//

type MockArgsTypeQueryWithTimestamp struct {
	ApomockCallNumber int
	Argtimestamp      int64
}

var LastMockArgsQueryWithTimestamp MockArgsTypeQueryWithTimestamp

// (recvq *Query)AuxMockWithTimestamp(argtimestamp int64)(reta *Query) - Generated mock function
func (recvq *Query) AuxMockWithTimestamp(argtimestamp int64) (reta *Query) {
	LastMockArgsQueryWithTimestamp = MockArgsTypeQueryWithTimestamp{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryWithTimestamp(),
		Argtimestamp:      argtimestamp,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.WithTimestamp")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.WithTimestamp")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.WithTimestamp")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryWithTimestamp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryWithTimestamp int = 0

var condRecorderAuxMockPtrQueryWithTimestamp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryWithTimestamp(i int) {
	condRecorderAuxMockPtrQueryWithTimestamp.L.Lock()
	for recorderAuxMockPtrQueryWithTimestamp < i {
		condRecorderAuxMockPtrQueryWithTimestamp.Wait()
	}
	condRecorderAuxMockPtrQueryWithTimestamp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryWithTimestamp() {
	condRecorderAuxMockPtrQueryWithTimestamp.L.Lock()
	recorderAuxMockPtrQueryWithTimestamp++
	condRecorderAuxMockPtrQueryWithTimestamp.L.Unlock()
	condRecorderAuxMockPtrQueryWithTimestamp.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryWithTimestamp() (ret int) {
	condRecorderAuxMockPtrQueryWithTimestamp.L.Lock()
	ret = recorderAuxMockPtrQueryWithTimestamp
	condRecorderAuxMockPtrQueryWithTimestamp.L.Unlock()
	return
}

// (recvq *Query)WithTimestamp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) WithTimestamp(argtimestamp int64) (reta *Query) {
	FuncAuxMockPtrQueryWithTimestamp, ok := apomock.GetRegisteredFunc("gocql.Query.WithTimestamp")
	if ok {
		reta = FuncAuxMockPtrQueryWithTimestamp.(func(recvq *Query, argtimestamp int64) (reta *Query))(recvq, argtimestamp)
	} else {
		panic("FuncAuxMockPtrQueryWithTimestamp ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryWithTimestamp()
	return
}

//
// Mock: (recvq *Query)retryPolicy()(reta RetryPolicy)
//

type MockArgsTypeQueryretryPolicy struct {
	ApomockCallNumber int
}

var LastMockArgsQueryretryPolicy MockArgsTypeQueryretryPolicy

// (recvq *Query)AuxMockretryPolicy()(reta RetryPolicy) - Generated mock function
func (recvq *Query) AuxMockretryPolicy() (reta RetryPolicy) {
	rargs, rerr := apomock.GetNext("gocql.Query.retryPolicy")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.retryPolicy")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.retryPolicy")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(RetryPolicy)
	}
	return
}

// RecorderAuxMockPtrQueryretryPolicy  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryretryPolicy int = 0

var condRecorderAuxMockPtrQueryretryPolicy *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryretryPolicy(i int) {
	condRecorderAuxMockPtrQueryretryPolicy.L.Lock()
	for recorderAuxMockPtrQueryretryPolicy < i {
		condRecorderAuxMockPtrQueryretryPolicy.Wait()
	}
	condRecorderAuxMockPtrQueryretryPolicy.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryretryPolicy() {
	condRecorderAuxMockPtrQueryretryPolicy.L.Lock()
	recorderAuxMockPtrQueryretryPolicy++
	condRecorderAuxMockPtrQueryretryPolicy.L.Unlock()
	condRecorderAuxMockPtrQueryretryPolicy.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryretryPolicy() (ret int) {
	condRecorderAuxMockPtrQueryretryPolicy.L.Lock()
	ret = recorderAuxMockPtrQueryretryPolicy
	condRecorderAuxMockPtrQueryretryPolicy.L.Unlock()
	return
}

// (recvq *Query)retryPolicy - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) retryPolicy() (reta RetryPolicy) {
	FuncAuxMockPtrQueryretryPolicy, ok := apomock.GetRegisteredFunc("gocql.Query.retryPolicy")
	if ok {
		reta = FuncAuxMockPtrQueryretryPolicy.(func(recvq *Query) (reta RetryPolicy))(recvq)
	} else {
		panic("FuncAuxMockPtrQueryretryPolicy ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryretryPolicy()
	return
}

//
// Mock: (recvq *Query)MapScanCAS(argdest map[string]interface{})(retapplied bool, reterr error)
//

type MockArgsTypeQueryMapScanCAS struct {
	ApomockCallNumber int
	Argdest           map[string]interface{}
}

var LastMockArgsQueryMapScanCAS MockArgsTypeQueryMapScanCAS

// (recvq *Query)AuxMockMapScanCAS(argdest map[string]interface{})(retapplied bool, reterr error) - Generated mock function
func (recvq *Query) AuxMockMapScanCAS(argdest map[string]interface{}) (retapplied bool, reterr error) {
	LastMockArgsQueryMapScanCAS = MockArgsTypeQueryMapScanCAS{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryMapScanCAS(),
		Argdest:           argdest,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.MapScanCAS")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.MapScanCAS")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Query.MapScanCAS")
	}
	if rargs.GetArg(0) != nil {
		retapplied = rargs.GetArg(0).(bool)
	}
	if rargs.GetArg(1) != nil {
		reterr = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrQueryMapScanCAS  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryMapScanCAS int = 0

var condRecorderAuxMockPtrQueryMapScanCAS *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryMapScanCAS(i int) {
	condRecorderAuxMockPtrQueryMapScanCAS.L.Lock()
	for recorderAuxMockPtrQueryMapScanCAS < i {
		condRecorderAuxMockPtrQueryMapScanCAS.Wait()
	}
	condRecorderAuxMockPtrQueryMapScanCAS.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryMapScanCAS() {
	condRecorderAuxMockPtrQueryMapScanCAS.L.Lock()
	recorderAuxMockPtrQueryMapScanCAS++
	condRecorderAuxMockPtrQueryMapScanCAS.L.Unlock()
	condRecorderAuxMockPtrQueryMapScanCAS.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryMapScanCAS() (ret int) {
	condRecorderAuxMockPtrQueryMapScanCAS.L.Lock()
	ret = recorderAuxMockPtrQueryMapScanCAS
	condRecorderAuxMockPtrQueryMapScanCAS.L.Unlock()
	return
}

// (recvq *Query)MapScanCAS - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) MapScanCAS(argdest map[string]interface{}) (retapplied bool, reterr error) {
	FuncAuxMockPtrQueryMapScanCAS, ok := apomock.GetRegisteredFunc("gocql.Query.MapScanCAS")
	if ok {
		retapplied, reterr = FuncAuxMockPtrQueryMapScanCAS.(func(recvq *Query, argdest map[string]interface{}) (retapplied bool, reterr error))(recvq, argdest)
	} else {
		panic("FuncAuxMockPtrQueryMapScanCAS ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryMapScanCAS()
	return
}

//
// Mock: (recviter *Iter)NumRows()(reta int)
//

type MockArgsTypeIterNumRows struct {
	ApomockCallNumber int
}

var LastMockArgsIterNumRows MockArgsTypeIterNumRows

// (recviter *Iter)AuxMockNumRows()(reta int) - Generated mock function
func (recviter *Iter) AuxMockNumRows() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.Iter.NumRows")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.NumRows")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Iter.NumRows")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrIterNumRows  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterNumRows int = 0

var condRecorderAuxMockPtrIterNumRows *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterNumRows(i int) {
	condRecorderAuxMockPtrIterNumRows.L.Lock()
	for recorderAuxMockPtrIterNumRows < i {
		condRecorderAuxMockPtrIterNumRows.Wait()
	}
	condRecorderAuxMockPtrIterNumRows.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterNumRows() {
	condRecorderAuxMockPtrIterNumRows.L.Lock()
	recorderAuxMockPtrIterNumRows++
	condRecorderAuxMockPtrIterNumRows.L.Unlock()
	condRecorderAuxMockPtrIterNumRows.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterNumRows() (ret int) {
	condRecorderAuxMockPtrIterNumRows.L.Lock()
	ret = recorderAuxMockPtrIterNumRows
	condRecorderAuxMockPtrIterNumRows.L.Unlock()
	return
}

// (recviter *Iter)NumRows - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) NumRows() (reta int) {
	FuncAuxMockPtrIterNumRows, ok := apomock.GetRegisteredFunc("gocql.Iter.NumRows")
	if ok {
		reta = FuncAuxMockPtrIterNumRows.(func(recviter *Iter) (reta int))(recviter)
	} else {
		panic("FuncAuxMockPtrIterNumRows ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterNumRows()
	return
}

//
// Mock: (recvr *routingKeyInfo)String()(reta string)
//

type MockArgsTyperoutingKeyInfoString struct {
	ApomockCallNumber int
}

var LastMockArgsroutingKeyInfoString MockArgsTyperoutingKeyInfoString

// (recvr *routingKeyInfo)AuxMockString()(reta string) - Generated mock function
func (recvr *routingKeyInfo) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.routingKeyInfo.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.routingKeyInfo.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.routingKeyInfo.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrroutingKeyInfoString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrroutingKeyInfoString int = 0

var condRecorderAuxMockPtrroutingKeyInfoString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrroutingKeyInfoString(i int) {
	condRecorderAuxMockPtrroutingKeyInfoString.L.Lock()
	for recorderAuxMockPtrroutingKeyInfoString < i {
		condRecorderAuxMockPtrroutingKeyInfoString.Wait()
	}
	condRecorderAuxMockPtrroutingKeyInfoString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrroutingKeyInfoString() {
	condRecorderAuxMockPtrroutingKeyInfoString.L.Lock()
	recorderAuxMockPtrroutingKeyInfoString++
	condRecorderAuxMockPtrroutingKeyInfoString.L.Unlock()
	condRecorderAuxMockPtrroutingKeyInfoString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrroutingKeyInfoString() (ret int) {
	condRecorderAuxMockPtrroutingKeyInfoString.L.Lock()
	ret = recorderAuxMockPtrroutingKeyInfoString
	condRecorderAuxMockPtrroutingKeyInfoString.L.Unlock()
	return
}

// (recvr *routingKeyInfo)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *routingKeyInfo) String() (reta string) {
	FuncAuxMockPtrroutingKeyInfoString, ok := apomock.GetRegisteredFunc("gocql.routingKeyInfo.String")
	if ok {
		reta = FuncAuxMockPtrroutingKeyInfoString.(func(recvr *routingKeyInfo) (reta string))(recvr)
	} else {
		panic("FuncAuxMockPtrroutingKeyInfoString ")
	}
	AuxMockIncrementRecorderAuxMockPtrroutingKeyInfoString()
	return
}

//
// Mock: (recvs *Session)Query(argstmt string, values ...interface{})(reta *Query)
//

type MockArgsTypeSessionQuery struct {
	ApomockCallNumber int
	Argstmt           string
	Values            []interface{}
}

var LastMockArgsSessionQuery MockArgsTypeSessionQuery

// (recvs *Session)AuxMockQuery(argstmt string, values ...interface{})(reta *Query) - Generated mock function
func (recvs *Session) AuxMockQuery(argstmt string, values ...interface{}) (reta *Query) {
	LastMockArgsSessionQuery = MockArgsTypeSessionQuery{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionQuery(),
		Argstmt:           argstmt,
		Values:            values,
	}
	rargs, rerr := apomock.GetNext("gocql.Session.Query")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.Query")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Session.Query")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrSessionQuery  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionQuery int = 0

var condRecorderAuxMockPtrSessionQuery *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionQuery(i int) {
	condRecorderAuxMockPtrSessionQuery.L.Lock()
	for recorderAuxMockPtrSessionQuery < i {
		condRecorderAuxMockPtrSessionQuery.Wait()
	}
	condRecorderAuxMockPtrSessionQuery.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionQuery() {
	condRecorderAuxMockPtrSessionQuery.L.Lock()
	recorderAuxMockPtrSessionQuery++
	condRecorderAuxMockPtrSessionQuery.L.Unlock()
	condRecorderAuxMockPtrSessionQuery.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionQuery() (ret int) {
	condRecorderAuxMockPtrSessionQuery.L.Lock()
	ret = recorderAuxMockPtrSessionQuery
	condRecorderAuxMockPtrSessionQuery.L.Unlock()
	return
}

// (recvs *Session)Query - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) Query(argstmt string, values ...interface{}) (reta *Query) {
	FuncAuxMockPtrSessionQuery, ok := apomock.GetRegisteredFunc("gocql.Session.Query")
	if ok {
		reta = FuncAuxMockPtrSessionQuery.(func(recvs *Session, argstmt string, values ...interface{}) (reta *Query))(recvs, argstmt, values...)
	} else {
		panic("FuncAuxMockPtrSessionQuery ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionQuery()
	return
}

//
// Mock: (recvs *Session)connect(argaddr string, argerrorHandler ConnErrorHandler, arghost *HostInfo)(reta *Conn, retb error)
//

type MockArgsTypeSessionconnect struct {
	ApomockCallNumber int
	Argaddr           string
	ArgerrorHandler   ConnErrorHandler
	Arghost           *HostInfo
}

var LastMockArgsSessionconnect MockArgsTypeSessionconnect

// (recvs *Session)AuxMockconnect(argaddr string, argerrorHandler ConnErrorHandler, arghost *HostInfo)(reta *Conn, retb error) - Generated mock function
func (recvs *Session) AuxMockconnect(argaddr string, argerrorHandler ConnErrorHandler, arghost *HostInfo) (reta *Conn, retb error) {
	LastMockArgsSessionconnect = MockArgsTypeSessionconnect{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionconnect(),
		Argaddr:           argaddr,
		ArgerrorHandler:   argerrorHandler,
		Arghost:           arghost,
	}
	rargs, rerr := apomock.GetNext("gocql.Session.connect")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.connect")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Session.connect")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Conn)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrSessionconnect  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionconnect int = 0

var condRecorderAuxMockPtrSessionconnect *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionconnect(i int) {
	condRecorderAuxMockPtrSessionconnect.L.Lock()
	for recorderAuxMockPtrSessionconnect < i {
		condRecorderAuxMockPtrSessionconnect.Wait()
	}
	condRecorderAuxMockPtrSessionconnect.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionconnect() {
	condRecorderAuxMockPtrSessionconnect.L.Lock()
	recorderAuxMockPtrSessionconnect++
	condRecorderAuxMockPtrSessionconnect.L.Unlock()
	condRecorderAuxMockPtrSessionconnect.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionconnect() (ret int) {
	condRecorderAuxMockPtrSessionconnect.L.Lock()
	ret = recorderAuxMockPtrSessionconnect
	condRecorderAuxMockPtrSessionconnect.L.Unlock()
	return
}

// (recvs *Session)connect - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) connect(argaddr string, argerrorHandler ConnErrorHandler, arghost *HostInfo) (reta *Conn, retb error) {
	FuncAuxMockPtrSessionconnect, ok := apomock.GetRegisteredFunc("gocql.Session.connect")
	if ok {
		reta, retb = FuncAuxMockPtrSessionconnect.(func(recvs *Session, argaddr string, argerrorHandler ConnErrorHandler, arghost *HostInfo) (reta *Conn, retb error))(recvs, argaddr, argerrorHandler, arghost)
	} else {
		panic("FuncAuxMockPtrSessionconnect ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionconnect()
	return
}

//
// Mock: (recvq *Query)Scan(dest ...interface{})(reta error)
//

type MockArgsTypeQueryScan struct {
	ApomockCallNumber int
	Dest              []interface{}
}

var LastMockArgsQueryScan MockArgsTypeQueryScan

// (recvq *Query)AuxMockScan(dest ...interface{})(reta error) - Generated mock function
func (recvq *Query) AuxMockScan(dest ...interface{}) (reta error) {
	LastMockArgsQueryScan = MockArgsTypeQueryScan{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryScan(),
		Dest:              dest,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.Scan")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.Scan")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.Scan")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrQueryScan  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryScan int = 0

var condRecorderAuxMockPtrQueryScan *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryScan(i int) {
	condRecorderAuxMockPtrQueryScan.L.Lock()
	for recorderAuxMockPtrQueryScan < i {
		condRecorderAuxMockPtrQueryScan.Wait()
	}
	condRecorderAuxMockPtrQueryScan.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryScan() {
	condRecorderAuxMockPtrQueryScan.L.Lock()
	recorderAuxMockPtrQueryScan++
	condRecorderAuxMockPtrQueryScan.L.Unlock()
	condRecorderAuxMockPtrQueryScan.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryScan() (ret int) {
	condRecorderAuxMockPtrQueryScan.L.Lock()
	ret = recorderAuxMockPtrQueryScan
	condRecorderAuxMockPtrQueryScan.L.Unlock()
	return
}

// (recvq *Query)Scan - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) Scan(dest ...interface{}) (reta error) {
	FuncAuxMockPtrQueryScan, ok := apomock.GetRegisteredFunc("gocql.Query.Scan")
	if ok {
		reta = FuncAuxMockPtrQueryScan.(func(recvq *Query, dest ...interface{}) (reta error))(recvq, dest...)
	} else {
		panic("FuncAuxMockPtrQueryScan ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryScan()
	return
}

//
// Mock: (recvb *Batch)retryPolicy()(reta RetryPolicy)
//

type MockArgsTypeBatchretryPolicy struct {
	ApomockCallNumber int
}

var LastMockArgsBatchretryPolicy MockArgsTypeBatchretryPolicy

// (recvb *Batch)AuxMockretryPolicy()(reta RetryPolicy) - Generated mock function
func (recvb *Batch) AuxMockretryPolicy() (reta RetryPolicy) {
	rargs, rerr := apomock.GetNext("gocql.Batch.retryPolicy")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Batch.retryPolicy")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Batch.retryPolicy")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(RetryPolicy)
	}
	return
}

// RecorderAuxMockPtrBatchretryPolicy  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchretryPolicy int = 0

var condRecorderAuxMockPtrBatchretryPolicy *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchretryPolicy(i int) {
	condRecorderAuxMockPtrBatchretryPolicy.L.Lock()
	for recorderAuxMockPtrBatchretryPolicy < i {
		condRecorderAuxMockPtrBatchretryPolicy.Wait()
	}
	condRecorderAuxMockPtrBatchretryPolicy.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchretryPolicy() {
	condRecorderAuxMockPtrBatchretryPolicy.L.Lock()
	recorderAuxMockPtrBatchretryPolicy++
	condRecorderAuxMockPtrBatchretryPolicy.L.Unlock()
	condRecorderAuxMockPtrBatchretryPolicy.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchretryPolicy() (ret int) {
	condRecorderAuxMockPtrBatchretryPolicy.L.Lock()
	ret = recorderAuxMockPtrBatchretryPolicy
	condRecorderAuxMockPtrBatchretryPolicy.L.Unlock()
	return
}

// (recvb *Batch)retryPolicy - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) retryPolicy() (reta RetryPolicy) {
	FuncAuxMockPtrBatchretryPolicy, ok := apomock.GetRegisteredFunc("gocql.Batch.retryPolicy")
	if ok {
		reta = FuncAuxMockPtrBatchretryPolicy.(func(recvb *Batch) (reta RetryPolicy))(recvb)
	} else {
		panic("FuncAuxMockPtrBatchretryPolicy ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchretryPolicy()
	return
}

//
// Mock: (recvt *traceWriter)Trace(argtraceId []byte)()
//

type MockArgsTypetraceWriterTrace struct {
	ApomockCallNumber int
	ArgtraceId        []byte
}

var LastMockArgstraceWriterTrace MockArgsTypetraceWriterTrace

// (recvt *traceWriter)AuxMockTrace(argtraceId []byte)() - Generated mock function
func (recvt *traceWriter) AuxMockTrace(argtraceId []byte) {
	LastMockArgstraceWriterTrace = MockArgsTypetraceWriterTrace{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrtraceWriterTrace(),
		ArgtraceId:        argtraceId,
	}
	return
}

// RecorderAuxMockPtrtraceWriterTrace  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtraceWriterTrace int = 0

var condRecorderAuxMockPtrtraceWriterTrace *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtraceWriterTrace(i int) {
	condRecorderAuxMockPtrtraceWriterTrace.L.Lock()
	for recorderAuxMockPtrtraceWriterTrace < i {
		condRecorderAuxMockPtrtraceWriterTrace.Wait()
	}
	condRecorderAuxMockPtrtraceWriterTrace.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtraceWriterTrace() {
	condRecorderAuxMockPtrtraceWriterTrace.L.Lock()
	recorderAuxMockPtrtraceWriterTrace++
	condRecorderAuxMockPtrtraceWriterTrace.L.Unlock()
	condRecorderAuxMockPtrtraceWriterTrace.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtraceWriterTrace() (ret int) {
	condRecorderAuxMockPtrtraceWriterTrace.L.Lock()
	ret = recorderAuxMockPtrtraceWriterTrace
	condRecorderAuxMockPtrtraceWriterTrace.L.Unlock()
	return
}

// (recvt *traceWriter)Trace - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *traceWriter) Trace(argtraceId []byte) {
	FuncAuxMockPtrtraceWriterTrace, ok := apomock.GetRegisteredFunc("gocql.traceWriter.Trace")
	if ok {
		FuncAuxMockPtrtraceWriterTrace.(func(recvt *traceWriter, argtraceId []byte))(recvt, argtraceId)
	} else {
		panic("FuncAuxMockPtrtraceWriterTrace ")
	}
	AuxMockIncrementRecorderAuxMockPtrtraceWriterTrace()
	return
}

//
// Mock: (recvs *Session)Close()()
//

type MockArgsTypeSessionClose struct {
	ApomockCallNumber int
}

var LastMockArgsSessionClose MockArgsTypeSessionClose

// (recvs *Session)AuxMockClose()() - Generated mock function
func (recvs *Session) AuxMockClose() {
	return
}

// RecorderAuxMockPtrSessionClose  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionClose int = 0

var condRecorderAuxMockPtrSessionClose *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionClose(i int) {
	condRecorderAuxMockPtrSessionClose.L.Lock()
	for recorderAuxMockPtrSessionClose < i {
		condRecorderAuxMockPtrSessionClose.Wait()
	}
	condRecorderAuxMockPtrSessionClose.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionClose() {
	condRecorderAuxMockPtrSessionClose.L.Lock()
	recorderAuxMockPtrSessionClose++
	condRecorderAuxMockPtrSessionClose.L.Unlock()
	condRecorderAuxMockPtrSessionClose.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionClose() (ret int) {
	condRecorderAuxMockPtrSessionClose.L.Lock()
	ret = recorderAuxMockPtrSessionClose
	condRecorderAuxMockPtrSessionClose.L.Unlock()
	return
}

// (recvs *Session)Close - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) Close() {
	FuncAuxMockPtrSessionClose, ok := apomock.GetRegisteredFunc("gocql.Session.Close")
	if ok {
		FuncAuxMockPtrSessionClose.(func(recvs *Session))(recvs)
	} else {
		panic("FuncAuxMockPtrSessionClose ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionClose()
	return
}

//
// Mock: (recviter *Iter)Host()(reta *HostInfo)
//

type MockArgsTypeIterHost struct {
	ApomockCallNumber int
}

var LastMockArgsIterHost MockArgsTypeIterHost

// (recviter *Iter)AuxMockHost()(reta *HostInfo) - Generated mock function
func (recviter *Iter) AuxMockHost() (reta *HostInfo) {
	rargs, rerr := apomock.GetNext("gocql.Iter.Host")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.Host")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Iter.Host")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrIterHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterHost int = 0

var condRecorderAuxMockPtrIterHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterHost(i int) {
	condRecorderAuxMockPtrIterHost.L.Lock()
	for recorderAuxMockPtrIterHost < i {
		condRecorderAuxMockPtrIterHost.Wait()
	}
	condRecorderAuxMockPtrIterHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterHost() {
	condRecorderAuxMockPtrIterHost.L.Lock()
	recorderAuxMockPtrIterHost++
	condRecorderAuxMockPtrIterHost.L.Unlock()
	condRecorderAuxMockPtrIterHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterHost() (ret int) {
	condRecorderAuxMockPtrIterHost.L.Lock()
	ret = recorderAuxMockPtrIterHost
	condRecorderAuxMockPtrIterHost.L.Unlock()
	return
}

// (recviter *Iter)Host - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) Host() (reta *HostInfo) {
	FuncAuxMockPtrIterHost, ok := apomock.GetRegisteredFunc("gocql.Iter.Host")
	if ok {
		reta = FuncAuxMockPtrIterHost.(func(recviter *Iter) (reta *HostInfo))(recviter)
	} else {
		panic("FuncAuxMockPtrIterHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterHost()
	return
}

//
// Mock: (recviter *Iter)readColumn()(reta []byte, retb error)
//

type MockArgsTypeIterreadColumn struct {
	ApomockCallNumber int
}

var LastMockArgsIterreadColumn MockArgsTypeIterreadColumn

// (recviter *Iter)AuxMockreadColumn()(reta []byte, retb error) - Generated mock function
func (recviter *Iter) AuxMockreadColumn() (reta []byte, retb error) {
	rargs, rerr := apomock.GetNext("gocql.Iter.readColumn")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.readColumn")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Iter.readColumn")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrIterreadColumn  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterreadColumn int = 0

var condRecorderAuxMockPtrIterreadColumn *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterreadColumn(i int) {
	condRecorderAuxMockPtrIterreadColumn.L.Lock()
	for recorderAuxMockPtrIterreadColumn < i {
		condRecorderAuxMockPtrIterreadColumn.Wait()
	}
	condRecorderAuxMockPtrIterreadColumn.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterreadColumn() {
	condRecorderAuxMockPtrIterreadColumn.L.Lock()
	recorderAuxMockPtrIterreadColumn++
	condRecorderAuxMockPtrIterreadColumn.L.Unlock()
	condRecorderAuxMockPtrIterreadColumn.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterreadColumn() (ret int) {
	condRecorderAuxMockPtrIterreadColumn.L.Lock()
	ret = recorderAuxMockPtrIterreadColumn
	condRecorderAuxMockPtrIterreadColumn.L.Unlock()
	return
}

// (recviter *Iter)readColumn - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) readColumn() (reta []byte, retb error) {
	FuncAuxMockPtrIterreadColumn, ok := apomock.GetRegisteredFunc("gocql.Iter.readColumn")
	if ok {
		reta, retb = FuncAuxMockPtrIterreadColumn.(func(recviter *Iter) (reta []byte, retb error))(recviter)
	} else {
		panic("FuncAuxMockPtrIterreadColumn ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterreadColumn()
	return
}

//
// Mock: (recvq Query)String()(reta string)
//

type MockArgsTypeQueryString struct {
	ApomockCallNumber int
}

var LastMockArgsQueryString MockArgsTypeQueryString

// (recvq Query)AuxMockString()(reta string) - Generated mock function
func (recvq Query) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.Query.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockQueryString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockQueryString int = 0

var condRecorderAuxMockQueryString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockQueryString(i int) {
	condRecorderAuxMockQueryString.L.Lock()
	for recorderAuxMockQueryString < i {
		condRecorderAuxMockQueryString.Wait()
	}
	condRecorderAuxMockQueryString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockQueryString() {
	condRecorderAuxMockQueryString.L.Lock()
	recorderAuxMockQueryString++
	condRecorderAuxMockQueryString.L.Unlock()
	condRecorderAuxMockQueryString.Broadcast()
}
func AuxMockGetRecorderAuxMockQueryString() (ret int) {
	condRecorderAuxMockQueryString.L.Lock()
	ret = recorderAuxMockQueryString
	condRecorderAuxMockQueryString.L.Unlock()
	return
}

// (recvq Query)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq Query) String() (reta string) {
	FuncAuxMockQueryString, ok := apomock.GetRegisteredFunc("gocql.Query.String")
	if ok {
		reta = FuncAuxMockQueryString.(func(recvq Query) (reta string))(recvq)
	} else {
		panic("FuncAuxMockQueryString ")
	}
	AuxMockIncrementRecorderAuxMockQueryString()
	return
}

//
// Mock: (recvs *Session)ExecuteBatch(argbatch *Batch)(reta error)
//

type MockArgsTypeSessionExecuteBatch struct {
	ApomockCallNumber int
	Argbatch          *Batch
}

var LastMockArgsSessionExecuteBatch MockArgsTypeSessionExecuteBatch

// (recvs *Session)AuxMockExecuteBatch(argbatch *Batch)(reta error) - Generated mock function
func (recvs *Session) AuxMockExecuteBatch(argbatch *Batch) (reta error) {
	LastMockArgsSessionExecuteBatch = MockArgsTypeSessionExecuteBatch{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionExecuteBatch(),
		Argbatch:          argbatch,
	}
	rargs, rerr := apomock.GetNext("gocql.Session.ExecuteBatch")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.ExecuteBatch")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Session.ExecuteBatch")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrSessionExecuteBatch  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionExecuteBatch int = 0

var condRecorderAuxMockPtrSessionExecuteBatch *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionExecuteBatch(i int) {
	condRecorderAuxMockPtrSessionExecuteBatch.L.Lock()
	for recorderAuxMockPtrSessionExecuteBatch < i {
		condRecorderAuxMockPtrSessionExecuteBatch.Wait()
	}
	condRecorderAuxMockPtrSessionExecuteBatch.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionExecuteBatch() {
	condRecorderAuxMockPtrSessionExecuteBatch.L.Lock()
	recorderAuxMockPtrSessionExecuteBatch++
	condRecorderAuxMockPtrSessionExecuteBatch.L.Unlock()
	condRecorderAuxMockPtrSessionExecuteBatch.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionExecuteBatch() (ret int) {
	condRecorderAuxMockPtrSessionExecuteBatch.L.Lock()
	ret = recorderAuxMockPtrSessionExecuteBatch
	condRecorderAuxMockPtrSessionExecuteBatch.L.Unlock()
	return
}

// (recvs *Session)ExecuteBatch - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) ExecuteBatch(argbatch *Batch) (reta error) {
	FuncAuxMockPtrSessionExecuteBatch, ok := apomock.GetRegisteredFunc("gocql.Session.ExecuteBatch")
	if ok {
		reta = FuncAuxMockPtrSessionExecuteBatch.(func(recvs *Session, argbatch *Batch) (reta error))(recvs, argbatch)
	} else {
		panic("FuncAuxMockPtrSessionExecuteBatch ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionExecuteBatch()
	return
}

//
// Mock: NewBatch(argtyp BatchType)(reta *Batch)
//

type MockArgsTypeNewBatch struct {
	ApomockCallNumber int
	Argtyp            BatchType
}

var LastMockArgsNewBatch MockArgsTypeNewBatch

// AuxMockNewBatch(argtyp BatchType)(reta *Batch) - Generated mock function
func AuxMockNewBatch(argtyp BatchType) (reta *Batch) {
	LastMockArgsNewBatch = MockArgsTypeNewBatch{
		ApomockCallNumber: AuxMockGetRecorderAuxMockNewBatch(),
		Argtyp:            argtyp,
	}
	rargs, rerr := apomock.GetNext("gocql.NewBatch")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.NewBatch")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.NewBatch")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Batch)
	}
	return
}

// RecorderAuxMockNewBatch  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNewBatch int = 0

var condRecorderAuxMockNewBatch *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNewBatch(i int) {
	condRecorderAuxMockNewBatch.L.Lock()
	for recorderAuxMockNewBatch < i {
		condRecorderAuxMockNewBatch.Wait()
	}
	condRecorderAuxMockNewBatch.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNewBatch() {
	condRecorderAuxMockNewBatch.L.Lock()
	recorderAuxMockNewBatch++
	condRecorderAuxMockNewBatch.L.Unlock()
	condRecorderAuxMockNewBatch.Broadcast()
}
func AuxMockGetRecorderAuxMockNewBatch() (ret int) {
	condRecorderAuxMockNewBatch.L.Lock()
	ret = recorderAuxMockNewBatch
	condRecorderAuxMockNewBatch.L.Unlock()
	return
}

// NewBatch - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func NewBatch(argtyp BatchType) (reta *Batch) {
	FuncAuxMockNewBatch, ok := apomock.GetRegisteredFunc("gocql.NewBatch")
	if ok {
		reta = FuncAuxMockNewBatch.(func(argtyp BatchType) (reta *Batch))(argtyp)
	} else {
		panic("FuncAuxMockNewBatch ")
	}
	AuxMockIncrementRecorderAuxMockNewBatch()
	return
}

//
// Mock: (recvb *Batch)attempt(argd time.Duration)()
//

type MockArgsTypeBatchattempt struct {
	ApomockCallNumber int
	Argd              time.Duration
}

var LastMockArgsBatchattempt MockArgsTypeBatchattempt

// (recvb *Batch)AuxMockattempt(argd time.Duration)() - Generated mock function
func (recvb *Batch) AuxMockattempt(argd time.Duration) {
	LastMockArgsBatchattempt = MockArgsTypeBatchattempt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrBatchattempt(),
		Argd:              argd,
	}
	return
}

// RecorderAuxMockPtrBatchattempt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchattempt int = 0

var condRecorderAuxMockPtrBatchattempt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchattempt(i int) {
	condRecorderAuxMockPtrBatchattempt.L.Lock()
	for recorderAuxMockPtrBatchattempt < i {
		condRecorderAuxMockPtrBatchattempt.Wait()
	}
	condRecorderAuxMockPtrBatchattempt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchattempt() {
	condRecorderAuxMockPtrBatchattempt.L.Lock()
	recorderAuxMockPtrBatchattempt++
	condRecorderAuxMockPtrBatchattempt.L.Unlock()
	condRecorderAuxMockPtrBatchattempt.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchattempt() (ret int) {
	condRecorderAuxMockPtrBatchattempt.L.Lock()
	ret = recorderAuxMockPtrBatchattempt
	condRecorderAuxMockPtrBatchattempt.L.Unlock()
	return
}

// (recvb *Batch)attempt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) attempt(argd time.Duration) {
	FuncAuxMockPtrBatchattempt, ok := apomock.GetRegisteredFunc("gocql.Batch.attempt")
	if ok {
		FuncAuxMockPtrBatchattempt.(func(recvb *Batch, argd time.Duration))(recvb, argd)
	} else {
		panic("FuncAuxMockPtrBatchattempt ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchattempt()
	return
}

//
// Mock: (recvs *Session)SetTrace(argtrace Tracer)()
//

type MockArgsTypeSessionSetTrace struct {
	ApomockCallNumber int
	Argtrace          Tracer
}

var LastMockArgsSessionSetTrace MockArgsTypeSessionSetTrace

// (recvs *Session)AuxMockSetTrace(argtrace Tracer)() - Generated mock function
func (recvs *Session) AuxMockSetTrace(argtrace Tracer) {
	LastMockArgsSessionSetTrace = MockArgsTypeSessionSetTrace{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionSetTrace(),
		Argtrace:          argtrace,
	}
	return
}

// RecorderAuxMockPtrSessionSetTrace  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionSetTrace int = 0

var condRecorderAuxMockPtrSessionSetTrace *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionSetTrace(i int) {
	condRecorderAuxMockPtrSessionSetTrace.L.Lock()
	for recorderAuxMockPtrSessionSetTrace < i {
		condRecorderAuxMockPtrSessionSetTrace.Wait()
	}
	condRecorderAuxMockPtrSessionSetTrace.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionSetTrace() {
	condRecorderAuxMockPtrSessionSetTrace.L.Lock()
	recorderAuxMockPtrSessionSetTrace++
	condRecorderAuxMockPtrSessionSetTrace.L.Unlock()
	condRecorderAuxMockPtrSessionSetTrace.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionSetTrace() (ret int) {
	condRecorderAuxMockPtrSessionSetTrace.L.Lock()
	ret = recorderAuxMockPtrSessionSetTrace
	condRecorderAuxMockPtrSessionSetTrace.L.Unlock()
	return
}

// (recvs *Session)SetTrace - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) SetTrace(argtrace Tracer) {
	FuncAuxMockPtrSessionSetTrace, ok := apomock.GetRegisteredFunc("gocql.Session.SetTrace")
	if ok {
		FuncAuxMockPtrSessionSetTrace.(func(recvs *Session, argtrace Tracer))(recvs, argtrace)
	} else {
		panic("FuncAuxMockPtrSessionSetTrace ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionSetTrace()
	return
}

//
// Mock: (recvs *Session)SetConsistency(argcons Consistency)()
//

type MockArgsTypeSessionSetConsistency struct {
	ApomockCallNumber int
	Argcons           Consistency
}

var LastMockArgsSessionSetConsistency MockArgsTypeSessionSetConsistency

// (recvs *Session)AuxMockSetConsistency(argcons Consistency)() - Generated mock function
func (recvs *Session) AuxMockSetConsistency(argcons Consistency) {
	LastMockArgsSessionSetConsistency = MockArgsTypeSessionSetConsistency{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionSetConsistency(),
		Argcons:           argcons,
	}
	return
}

// RecorderAuxMockPtrSessionSetConsistency  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionSetConsistency int = 0

var condRecorderAuxMockPtrSessionSetConsistency *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionSetConsistency(i int) {
	condRecorderAuxMockPtrSessionSetConsistency.L.Lock()
	for recorderAuxMockPtrSessionSetConsistency < i {
		condRecorderAuxMockPtrSessionSetConsistency.Wait()
	}
	condRecorderAuxMockPtrSessionSetConsistency.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionSetConsistency() {
	condRecorderAuxMockPtrSessionSetConsistency.L.Lock()
	recorderAuxMockPtrSessionSetConsistency++
	condRecorderAuxMockPtrSessionSetConsistency.L.Unlock()
	condRecorderAuxMockPtrSessionSetConsistency.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionSetConsistency() (ret int) {
	condRecorderAuxMockPtrSessionSetConsistency.L.Lock()
	ret = recorderAuxMockPtrSessionSetConsistency
	condRecorderAuxMockPtrSessionSetConsistency.L.Unlock()
	return
}

// (recvs *Session)SetConsistency - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) SetConsistency(argcons Consistency) {
	FuncAuxMockPtrSessionSetConsistency, ok := apomock.GetRegisteredFunc("gocql.Session.SetConsistency")
	if ok {
		FuncAuxMockPtrSessionSetConsistency.(func(recvs *Session, argcons Consistency))(recvs, argcons)
	} else {
		panic("FuncAuxMockPtrSessionSetConsistency ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionSetConsistency()
	return
}

//
// Mock: (recvb *Batch)execute(argconn *Conn)(reta *Iter)
//

type MockArgsTypeBatchexecute struct {
	ApomockCallNumber int
	Argconn           *Conn
}

var LastMockArgsBatchexecute MockArgsTypeBatchexecute

// (recvb *Batch)AuxMockexecute(argconn *Conn)(reta *Iter) - Generated mock function
func (recvb *Batch) AuxMockexecute(argconn *Conn) (reta *Iter) {
	LastMockArgsBatchexecute = MockArgsTypeBatchexecute{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrBatchexecute(),
		Argconn:           argconn,
	}
	rargs, rerr := apomock.GetNext("gocql.Batch.execute")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Batch.execute")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Batch.execute")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Iter)
	}
	return
}

// RecorderAuxMockPtrBatchexecute  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchexecute int = 0

var condRecorderAuxMockPtrBatchexecute *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchexecute(i int) {
	condRecorderAuxMockPtrBatchexecute.L.Lock()
	for recorderAuxMockPtrBatchexecute < i {
		condRecorderAuxMockPtrBatchexecute.Wait()
	}
	condRecorderAuxMockPtrBatchexecute.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchexecute() {
	condRecorderAuxMockPtrBatchexecute.L.Lock()
	recorderAuxMockPtrBatchexecute++
	condRecorderAuxMockPtrBatchexecute.L.Unlock()
	condRecorderAuxMockPtrBatchexecute.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchexecute() (ret int) {
	condRecorderAuxMockPtrBatchexecute.L.Lock()
	ret = recorderAuxMockPtrBatchexecute
	condRecorderAuxMockPtrBatchexecute.L.Unlock()
	return
}

// (recvb *Batch)execute - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) execute(argconn *Conn) (reta *Iter) {
	FuncAuxMockPtrBatchexecute, ok := apomock.GetRegisteredFunc("gocql.Batch.execute")
	if ok {
		reta = FuncAuxMockPtrBatchexecute.(func(recvb *Batch, argconn *Conn) (reta *Iter))(recvb, argconn)
	} else {
		panic("FuncAuxMockPtrBatchexecute ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchexecute()
	return
}

//
// Mock: (recvq *Query)Latency()(reta int64)
//

type MockArgsTypeQueryLatency struct {
	ApomockCallNumber int
}

var LastMockArgsQueryLatency MockArgsTypeQueryLatency

// (recvq *Query)AuxMockLatency()(reta int64) - Generated mock function
func (recvq *Query) AuxMockLatency() (reta int64) {
	rargs, rerr := apomock.GetNext("gocql.Query.Latency")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.Latency")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.Latency")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int64)
	}
	return
}

// RecorderAuxMockPtrQueryLatency  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryLatency int = 0

var condRecorderAuxMockPtrQueryLatency *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryLatency(i int) {
	condRecorderAuxMockPtrQueryLatency.L.Lock()
	for recorderAuxMockPtrQueryLatency < i {
		condRecorderAuxMockPtrQueryLatency.Wait()
	}
	condRecorderAuxMockPtrQueryLatency.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryLatency() {
	condRecorderAuxMockPtrQueryLatency.L.Lock()
	recorderAuxMockPtrQueryLatency++
	condRecorderAuxMockPtrQueryLatency.L.Unlock()
	condRecorderAuxMockPtrQueryLatency.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryLatency() (ret int) {
	condRecorderAuxMockPtrQueryLatency.L.Lock()
	ret = recorderAuxMockPtrQueryLatency
	condRecorderAuxMockPtrQueryLatency.L.Unlock()
	return
}

// (recvq *Query)Latency - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) Latency() (reta int64) {
	FuncAuxMockPtrQueryLatency, ok := apomock.GetRegisteredFunc("gocql.Query.Latency")
	if ok {
		reta = FuncAuxMockPtrQueryLatency.(func(recvq *Query) (reta int64))(recvq)
	} else {
		panic("FuncAuxMockPtrQueryLatency ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryLatency()
	return
}

//
// Mock: (recvq *Query)RoutingKey(argroutingKey []byte)(reta *Query)
//

type MockArgsTypeQueryRoutingKey struct {
	ApomockCallNumber int
	ArgroutingKey     []byte
}

var LastMockArgsQueryRoutingKey MockArgsTypeQueryRoutingKey

// (recvq *Query)AuxMockRoutingKey(argroutingKey []byte)(reta *Query) - Generated mock function
func (recvq *Query) AuxMockRoutingKey(argroutingKey []byte) (reta *Query) {
	LastMockArgsQueryRoutingKey = MockArgsTypeQueryRoutingKey{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryRoutingKey(),
		ArgroutingKey:     argroutingKey,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.RoutingKey")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.RoutingKey")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.RoutingKey")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryRoutingKey  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryRoutingKey int = 0

var condRecorderAuxMockPtrQueryRoutingKey *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryRoutingKey(i int) {
	condRecorderAuxMockPtrQueryRoutingKey.L.Lock()
	for recorderAuxMockPtrQueryRoutingKey < i {
		condRecorderAuxMockPtrQueryRoutingKey.Wait()
	}
	condRecorderAuxMockPtrQueryRoutingKey.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryRoutingKey() {
	condRecorderAuxMockPtrQueryRoutingKey.L.Lock()
	recorderAuxMockPtrQueryRoutingKey++
	condRecorderAuxMockPtrQueryRoutingKey.L.Unlock()
	condRecorderAuxMockPtrQueryRoutingKey.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryRoutingKey() (ret int) {
	condRecorderAuxMockPtrQueryRoutingKey.L.Lock()
	ret = recorderAuxMockPtrQueryRoutingKey
	condRecorderAuxMockPtrQueryRoutingKey.L.Unlock()
	return
}

// (recvq *Query)RoutingKey - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) RoutingKey(argroutingKey []byte) (reta *Query) {
	FuncAuxMockPtrQueryRoutingKey, ok := apomock.GetRegisteredFunc("gocql.Query.RoutingKey")
	if ok {
		reta = FuncAuxMockPtrQueryRoutingKey.(func(recvq *Query, argroutingKey []byte) (reta *Query))(recvq, argroutingKey)
	} else {
		panic("FuncAuxMockPtrQueryRoutingKey ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryRoutingKey()
	return
}

//
// Mock: (recvq *Query)shouldPrepare()(reta bool)
//

type MockArgsTypeQueryshouldPrepare struct {
	ApomockCallNumber int
}

var LastMockArgsQueryshouldPrepare MockArgsTypeQueryshouldPrepare

// (recvq *Query)AuxMockshouldPrepare()(reta bool) - Generated mock function
func (recvq *Query) AuxMockshouldPrepare() (reta bool) {
	rargs, rerr := apomock.GetNext("gocql.Query.shouldPrepare")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.shouldPrepare")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.shouldPrepare")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrQueryshouldPrepare  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryshouldPrepare int = 0

var condRecorderAuxMockPtrQueryshouldPrepare *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryshouldPrepare(i int) {
	condRecorderAuxMockPtrQueryshouldPrepare.L.Lock()
	for recorderAuxMockPtrQueryshouldPrepare < i {
		condRecorderAuxMockPtrQueryshouldPrepare.Wait()
	}
	condRecorderAuxMockPtrQueryshouldPrepare.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryshouldPrepare() {
	condRecorderAuxMockPtrQueryshouldPrepare.L.Lock()
	recorderAuxMockPtrQueryshouldPrepare++
	condRecorderAuxMockPtrQueryshouldPrepare.L.Unlock()
	condRecorderAuxMockPtrQueryshouldPrepare.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryshouldPrepare() (ret int) {
	condRecorderAuxMockPtrQueryshouldPrepare.L.Lock()
	ret = recorderAuxMockPtrQueryshouldPrepare
	condRecorderAuxMockPtrQueryshouldPrepare.L.Unlock()
	return
}

// (recvq *Query)shouldPrepare - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) shouldPrepare() (reta bool) {
	FuncAuxMockPtrQueryshouldPrepare, ok := apomock.GetRegisteredFunc("gocql.Query.shouldPrepare")
	if ok {
		reta = FuncAuxMockPtrQueryshouldPrepare.(func(recvq *Query) (reta bool))(recvq)
	} else {
		panic("FuncAuxMockPtrQueryshouldPrepare ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryshouldPrepare()
	return
}

//
// Mock: (recviter *Iter)Close()(reta error)
//

type MockArgsTypeIterClose struct {
	ApomockCallNumber int
}

var LastMockArgsIterClose MockArgsTypeIterClose

// (recviter *Iter)AuxMockClose()(reta error) - Generated mock function
func (recviter *Iter) AuxMockClose() (reta error) {
	rargs, rerr := apomock.GetNext("gocql.Iter.Close")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.Close")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Iter.Close")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrIterClose  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterClose int = 0

var condRecorderAuxMockPtrIterClose *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterClose(i int) {
	condRecorderAuxMockPtrIterClose.L.Lock()
	for recorderAuxMockPtrIterClose < i {
		condRecorderAuxMockPtrIterClose.Wait()
	}
	condRecorderAuxMockPtrIterClose.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterClose() {
	condRecorderAuxMockPtrIterClose.L.Lock()
	recorderAuxMockPtrIterClose++
	condRecorderAuxMockPtrIterClose.L.Unlock()
	condRecorderAuxMockPtrIterClose.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterClose() (ret int) {
	condRecorderAuxMockPtrIterClose.L.Lock()
	ret = recorderAuxMockPtrIterClose
	condRecorderAuxMockPtrIterClose.L.Unlock()
	return
}

// (recviter *Iter)Close - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) Close() (reta error) {
	FuncAuxMockPtrIterClose, ok := apomock.GetRegisteredFunc("gocql.Iter.Close")
	if ok {
		reta = FuncAuxMockPtrIterClose.(func(recviter *Iter) (reta error))(recviter)
	} else {
		panic("FuncAuxMockPtrIterClose ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterClose()
	return
}

//
// Mock: (recvb *Batch)Size()(reta int)
//

type MockArgsTypeBatchSize struct {
	ApomockCallNumber int
}

var LastMockArgsBatchSize MockArgsTypeBatchSize

// (recvb *Batch)AuxMockSize()(reta int) - Generated mock function
func (recvb *Batch) AuxMockSize() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.Batch.Size")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Batch.Size")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Batch.Size")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrBatchSize  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchSize int = 0

var condRecorderAuxMockPtrBatchSize *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchSize(i int) {
	condRecorderAuxMockPtrBatchSize.L.Lock()
	for recorderAuxMockPtrBatchSize < i {
		condRecorderAuxMockPtrBatchSize.Wait()
	}
	condRecorderAuxMockPtrBatchSize.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchSize() {
	condRecorderAuxMockPtrBatchSize.L.Lock()
	recorderAuxMockPtrBatchSize++
	condRecorderAuxMockPtrBatchSize.L.Unlock()
	condRecorderAuxMockPtrBatchSize.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchSize() (ret int) {
	condRecorderAuxMockPtrBatchSize.L.Lock()
	ret = recorderAuxMockPtrBatchSize
	condRecorderAuxMockPtrBatchSize.L.Unlock()
	return
}

// (recvb *Batch)Size - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) Size() (reta int) {
	FuncAuxMockPtrBatchSize, ok := apomock.GetRegisteredFunc("gocql.Batch.Size")
	if ok {
		reta = FuncAuxMockPtrBatchSize.(func(recvb *Batch) (reta int))(recvb)
	} else {
		panic("FuncAuxMockPtrBatchSize ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchSize()
	return
}

//
// Mock: addrsToHosts(argaddrs []string, argdefaultPort int)(reta []*HostInfo, retb error)
//

type MockArgsTypeaddrsToHosts struct {
	ApomockCallNumber int
	Argaddrs          []string
	ArgdefaultPort    int
}

var LastMockArgsaddrsToHosts MockArgsTypeaddrsToHosts

// AuxMockaddrsToHosts(argaddrs []string, argdefaultPort int)(reta []*HostInfo, retb error) - Generated mock function
func AuxMockaddrsToHosts(argaddrs []string, argdefaultPort int) (reta []*HostInfo, retb error) {
	LastMockArgsaddrsToHosts = MockArgsTypeaddrsToHosts{
		ApomockCallNumber: AuxMockGetRecorderAuxMockaddrsToHosts(),
		Argaddrs:          argaddrs,
		ArgdefaultPort:    argdefaultPort,
	}
	rargs, rerr := apomock.GetNext("gocql.addrsToHosts")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.addrsToHosts")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.addrsToHosts")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]*HostInfo)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockaddrsToHosts  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockaddrsToHosts int = 0

var condRecorderAuxMockaddrsToHosts *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockaddrsToHosts(i int) {
	condRecorderAuxMockaddrsToHosts.L.Lock()
	for recorderAuxMockaddrsToHosts < i {
		condRecorderAuxMockaddrsToHosts.Wait()
	}
	condRecorderAuxMockaddrsToHosts.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockaddrsToHosts() {
	condRecorderAuxMockaddrsToHosts.L.Lock()
	recorderAuxMockaddrsToHosts++
	condRecorderAuxMockaddrsToHosts.L.Unlock()
	condRecorderAuxMockaddrsToHosts.Broadcast()
}
func AuxMockGetRecorderAuxMockaddrsToHosts() (ret int) {
	condRecorderAuxMockaddrsToHosts.L.Lock()
	ret = recorderAuxMockaddrsToHosts
	condRecorderAuxMockaddrsToHosts.L.Unlock()
	return
}

// addrsToHosts - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func addrsToHosts(argaddrs []string, argdefaultPort int) (reta []*HostInfo, retb error) {
	FuncAuxMockaddrsToHosts, ok := apomock.GetRegisteredFunc("gocql.addrsToHosts")
	if ok {
		reta, retb = FuncAuxMockaddrsToHosts.(func(argaddrs []string, argdefaultPort int) (reta []*HostInfo, retb error))(argaddrs, argdefaultPort)
	} else {
		panic("FuncAuxMockaddrsToHosts ")
	}
	AuxMockIncrementRecorderAuxMockaddrsToHosts()
	return
}

//
// Mock: NewTraceWriter(argsession *Session, argw io.Writer)(reta Tracer)
//

type MockArgsTypeNewTraceWriter struct {
	ApomockCallNumber int
	Argsession        *Session
	Argw              io.Writer
}

var LastMockArgsNewTraceWriter MockArgsTypeNewTraceWriter

// AuxMockNewTraceWriter(argsession *Session, argw io.Writer)(reta Tracer) - Generated mock function
func AuxMockNewTraceWriter(argsession *Session, argw io.Writer) (reta Tracer) {
	LastMockArgsNewTraceWriter = MockArgsTypeNewTraceWriter{
		ApomockCallNumber: AuxMockGetRecorderAuxMockNewTraceWriter(),
		Argsession:        argsession,
		Argw:              argw,
	}
	rargs, rerr := apomock.GetNext("gocql.NewTraceWriter")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.NewTraceWriter")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.NewTraceWriter")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(Tracer)
	}
	return
}

// RecorderAuxMockNewTraceWriter  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNewTraceWriter int = 0

var condRecorderAuxMockNewTraceWriter *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNewTraceWriter(i int) {
	condRecorderAuxMockNewTraceWriter.L.Lock()
	for recorderAuxMockNewTraceWriter < i {
		condRecorderAuxMockNewTraceWriter.Wait()
	}
	condRecorderAuxMockNewTraceWriter.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNewTraceWriter() {
	condRecorderAuxMockNewTraceWriter.L.Lock()
	recorderAuxMockNewTraceWriter++
	condRecorderAuxMockNewTraceWriter.L.Unlock()
	condRecorderAuxMockNewTraceWriter.Broadcast()
}
func AuxMockGetRecorderAuxMockNewTraceWriter() (ret int) {
	condRecorderAuxMockNewTraceWriter.L.Lock()
	ret = recorderAuxMockNewTraceWriter
	condRecorderAuxMockNewTraceWriter.L.Unlock()
	return
}

// NewTraceWriter - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func NewTraceWriter(argsession *Session, argw io.Writer) (reta Tracer) {
	FuncAuxMockNewTraceWriter, ok := apomock.GetRegisteredFunc("gocql.NewTraceWriter")
	if ok {
		reta = FuncAuxMockNewTraceWriter.(func(argsession *Session, argw io.Writer) (reta Tracer))(argsession, argw)
	} else {
		panic("FuncAuxMockNewTraceWriter ")
	}
	AuxMockIncrementRecorderAuxMockNewTraceWriter()
	return
}

//
// Mock: (recvb *Batch)SerialConsistency(argcons SerialConsistency)(reta *Batch)
//

type MockArgsTypeBatchSerialConsistency struct {
	ApomockCallNumber int
	Argcons           SerialConsistency
}

var LastMockArgsBatchSerialConsistency MockArgsTypeBatchSerialConsistency

// (recvb *Batch)AuxMockSerialConsistency(argcons SerialConsistency)(reta *Batch) - Generated mock function
func (recvb *Batch) AuxMockSerialConsistency(argcons SerialConsistency) (reta *Batch) {
	LastMockArgsBatchSerialConsistency = MockArgsTypeBatchSerialConsistency{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrBatchSerialConsistency(),
		Argcons:           argcons,
	}
	rargs, rerr := apomock.GetNext("gocql.Batch.SerialConsistency")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Batch.SerialConsistency")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Batch.SerialConsistency")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Batch)
	}
	return
}

// RecorderAuxMockPtrBatchSerialConsistency  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchSerialConsistency int = 0

var condRecorderAuxMockPtrBatchSerialConsistency *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchSerialConsistency(i int) {
	condRecorderAuxMockPtrBatchSerialConsistency.L.Lock()
	for recorderAuxMockPtrBatchSerialConsistency < i {
		condRecorderAuxMockPtrBatchSerialConsistency.Wait()
	}
	condRecorderAuxMockPtrBatchSerialConsistency.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchSerialConsistency() {
	condRecorderAuxMockPtrBatchSerialConsistency.L.Lock()
	recorderAuxMockPtrBatchSerialConsistency++
	condRecorderAuxMockPtrBatchSerialConsistency.L.Unlock()
	condRecorderAuxMockPtrBatchSerialConsistency.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchSerialConsistency() (ret int) {
	condRecorderAuxMockPtrBatchSerialConsistency.L.Lock()
	ret = recorderAuxMockPtrBatchSerialConsistency
	condRecorderAuxMockPtrBatchSerialConsistency.L.Unlock()
	return
}

// (recvb *Batch)SerialConsistency - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) SerialConsistency(argcons SerialConsistency) (reta *Batch) {
	FuncAuxMockPtrBatchSerialConsistency, ok := apomock.GetRegisteredFunc("gocql.Batch.SerialConsistency")
	if ok {
		reta = FuncAuxMockPtrBatchSerialConsistency.(func(recvb *Batch, argcons SerialConsistency) (reta *Batch))(recvb, argcons)
	} else {
		panic("FuncAuxMockPtrBatchSerialConsistency ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchSerialConsistency()
	return
}

//
// Mock: (recvq *Query)PageState(argstate []byte)(reta *Query)
//

type MockArgsTypeQueryPageState struct {
	ApomockCallNumber int
	Argstate          []byte
}

var LastMockArgsQueryPageState MockArgsTypeQueryPageState

// (recvq *Query)AuxMockPageState(argstate []byte)(reta *Query) - Generated mock function
func (recvq *Query) AuxMockPageState(argstate []byte) (reta *Query) {
	LastMockArgsQueryPageState = MockArgsTypeQueryPageState{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryPageState(),
		Argstate:          argstate,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.PageState")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.PageState")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.PageState")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Query)
	}
	return
}

// RecorderAuxMockPtrQueryPageState  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryPageState int = 0

var condRecorderAuxMockPtrQueryPageState *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryPageState(i int) {
	condRecorderAuxMockPtrQueryPageState.L.Lock()
	for recorderAuxMockPtrQueryPageState < i {
		condRecorderAuxMockPtrQueryPageState.Wait()
	}
	condRecorderAuxMockPtrQueryPageState.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryPageState() {
	condRecorderAuxMockPtrQueryPageState.L.Lock()
	recorderAuxMockPtrQueryPageState++
	condRecorderAuxMockPtrQueryPageState.L.Unlock()
	condRecorderAuxMockPtrQueryPageState.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryPageState() (ret int) {
	condRecorderAuxMockPtrQueryPageState.L.Lock()
	ret = recorderAuxMockPtrQueryPageState
	condRecorderAuxMockPtrQueryPageState.L.Unlock()
	return
}

// (recvq *Query)PageState - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) PageState(argstate []byte) (reta *Query) {
	FuncAuxMockPtrQueryPageState, ok := apomock.GetRegisteredFunc("gocql.Query.PageState")
	if ok {
		reta = FuncAuxMockPtrQueryPageState.(func(recvq *Query, argstate []byte) (reta *Query))(recvq, argstate)
	} else {
		panic("FuncAuxMockPtrQueryPageState ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryPageState()
	return
}

//
// Mock: (recviter *Iter)checkErrAndNotFound()(reta error)
//

type MockArgsTypeItercheckErrAndNotFound struct {
	ApomockCallNumber int
}

var LastMockArgsItercheckErrAndNotFound MockArgsTypeItercheckErrAndNotFound

// (recviter *Iter)AuxMockcheckErrAndNotFound()(reta error) - Generated mock function
func (recviter *Iter) AuxMockcheckErrAndNotFound() (reta error) {
	rargs, rerr := apomock.GetNext("gocql.Iter.checkErrAndNotFound")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.checkErrAndNotFound")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Iter.checkErrAndNotFound")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrItercheckErrAndNotFound  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrItercheckErrAndNotFound int = 0

var condRecorderAuxMockPtrItercheckErrAndNotFound *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrItercheckErrAndNotFound(i int) {
	condRecorderAuxMockPtrItercheckErrAndNotFound.L.Lock()
	for recorderAuxMockPtrItercheckErrAndNotFound < i {
		condRecorderAuxMockPtrItercheckErrAndNotFound.Wait()
	}
	condRecorderAuxMockPtrItercheckErrAndNotFound.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrItercheckErrAndNotFound() {
	condRecorderAuxMockPtrItercheckErrAndNotFound.L.Lock()
	recorderAuxMockPtrItercheckErrAndNotFound++
	condRecorderAuxMockPtrItercheckErrAndNotFound.L.Unlock()
	condRecorderAuxMockPtrItercheckErrAndNotFound.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrItercheckErrAndNotFound() (ret int) {
	condRecorderAuxMockPtrItercheckErrAndNotFound.L.Lock()
	ret = recorderAuxMockPtrItercheckErrAndNotFound
	condRecorderAuxMockPtrItercheckErrAndNotFound.L.Unlock()
	return
}

// (recviter *Iter)checkErrAndNotFound - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) checkErrAndNotFound() (reta error) {
	FuncAuxMockPtrItercheckErrAndNotFound, ok := apomock.GetRegisteredFunc("gocql.Iter.checkErrAndNotFound")
	if ok {
		reta = FuncAuxMockPtrItercheckErrAndNotFound.(func(recviter *Iter) (reta error))(recviter)
	} else {
		panic("FuncAuxMockPtrItercheckErrAndNotFound ")
	}
	AuxMockIncrementRecorderAuxMockPtrItercheckErrAndNotFound()
	return
}

//
// Mock: NewSession(argcfg ClusterConfig)(reta *Session, retb error)
//

type MockArgsTypeNewSession struct {
	ApomockCallNumber int
	Argcfg            ClusterConfig
}

var LastMockArgsNewSession MockArgsTypeNewSession

// AuxMockNewSession(argcfg ClusterConfig)(reta *Session, retb error) - Generated mock function
func AuxMockNewSession(argcfg ClusterConfig) (reta *Session, retb error) {
	LastMockArgsNewSession = MockArgsTypeNewSession{
		ApomockCallNumber: AuxMockGetRecorderAuxMockNewSession(),
		Argcfg:            argcfg,
	}
	rargs, rerr := apomock.GetNext("gocql.NewSession")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.NewSession")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.NewSession")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Session)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockNewSession  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNewSession int = 0

var condRecorderAuxMockNewSession *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNewSession(i int) {
	condRecorderAuxMockNewSession.L.Lock()
	for recorderAuxMockNewSession < i {
		condRecorderAuxMockNewSession.Wait()
	}
	condRecorderAuxMockNewSession.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNewSession() {
	condRecorderAuxMockNewSession.L.Lock()
	recorderAuxMockNewSession++
	condRecorderAuxMockNewSession.L.Unlock()
	condRecorderAuxMockNewSession.Broadcast()
}
func AuxMockGetRecorderAuxMockNewSession() (ret int) {
	condRecorderAuxMockNewSession.L.Lock()
	ret = recorderAuxMockNewSession
	condRecorderAuxMockNewSession.L.Unlock()
	return
}

// NewSession - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func NewSession(argcfg ClusterConfig) (reta *Session, retb error) {
	FuncAuxMockNewSession, ok := apomock.GetRegisteredFunc("gocql.NewSession")
	if ok {
		reta, retb = FuncAuxMockNewSession.(func(argcfg ClusterConfig) (reta *Session, retb error))(argcfg)
	} else {
		panic("FuncAuxMockNewSession ")
	}
	AuxMockIncrementRecorderAuxMockNewSession()
	return
}

//
// Mock: (recvq *Query)MapScan(argm map[string]interface{})(reta error)
//

type MockArgsTypeQueryMapScan struct {
	ApomockCallNumber int
	Argm              map[string]interface{}
}

var LastMockArgsQueryMapScan MockArgsTypeQueryMapScan

// (recvq *Query)AuxMockMapScan(argm map[string]interface{})(reta error) - Generated mock function
func (recvq *Query) AuxMockMapScan(argm map[string]interface{}) (reta error) {
	LastMockArgsQueryMapScan = MockArgsTypeQueryMapScan{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrQueryMapScan(),
		Argm:              argm,
	}
	rargs, rerr := apomock.GetNext("gocql.Query.MapScan")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Query.MapScan")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Query.MapScan")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrQueryMapScan  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrQueryMapScan int = 0

var condRecorderAuxMockPtrQueryMapScan *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrQueryMapScan(i int) {
	condRecorderAuxMockPtrQueryMapScan.L.Lock()
	for recorderAuxMockPtrQueryMapScan < i {
		condRecorderAuxMockPtrQueryMapScan.Wait()
	}
	condRecorderAuxMockPtrQueryMapScan.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrQueryMapScan() {
	condRecorderAuxMockPtrQueryMapScan.L.Lock()
	recorderAuxMockPtrQueryMapScan++
	condRecorderAuxMockPtrQueryMapScan.L.Unlock()
	condRecorderAuxMockPtrQueryMapScan.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrQueryMapScan() (ret int) {
	condRecorderAuxMockPtrQueryMapScan.L.Lock()
	ret = recorderAuxMockPtrQueryMapScan
	condRecorderAuxMockPtrQueryMapScan.L.Unlock()
	return
}

// (recvq *Query)MapScan - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq *Query) MapScan(argm map[string]interface{}) (reta error) {
	FuncAuxMockPtrQueryMapScan, ok := apomock.GetRegisteredFunc("gocql.Query.MapScan")
	if ok {
		reta = FuncAuxMockPtrQueryMapScan.(func(recvq *Query, argm map[string]interface{}) (reta error))(recvq, argm)
	} else {
		panic("FuncAuxMockPtrQueryMapScan ")
	}
	AuxMockIncrementRecorderAuxMockPtrQueryMapScan()
	return
}

//
// Mock: (recviter *Iter)Scan(dest ...interface{})(reta bool)
//

type MockArgsTypeIterScan struct {
	ApomockCallNumber int
	Dest              []interface{}
}

var LastMockArgsIterScan MockArgsTypeIterScan

// (recviter *Iter)AuxMockScan(dest ...interface{})(reta bool) - Generated mock function
func (recviter *Iter) AuxMockScan(dest ...interface{}) (reta bool) {
	LastMockArgsIterScan = MockArgsTypeIterScan{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrIterScan(),
		Dest:              dest,
	}
	rargs, rerr := apomock.GetNext("gocql.Iter.Scan")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.Scan")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Iter.Scan")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrIterScan  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterScan int = 0

var condRecorderAuxMockPtrIterScan *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterScan(i int) {
	condRecorderAuxMockPtrIterScan.L.Lock()
	for recorderAuxMockPtrIterScan < i {
		condRecorderAuxMockPtrIterScan.Wait()
	}
	condRecorderAuxMockPtrIterScan.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterScan() {
	condRecorderAuxMockPtrIterScan.L.Lock()
	recorderAuxMockPtrIterScan++
	condRecorderAuxMockPtrIterScan.L.Unlock()
	condRecorderAuxMockPtrIterScan.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterScan() (ret int) {
	condRecorderAuxMockPtrIterScan.L.Lock()
	ret = recorderAuxMockPtrIterScan
	condRecorderAuxMockPtrIterScan.L.Unlock()
	return
}

// (recviter *Iter)Scan - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) Scan(dest ...interface{}) (reta bool) {
	FuncAuxMockPtrIterScan, ok := apomock.GetRegisteredFunc("gocql.Iter.Scan")
	if ok {
		reta = FuncAuxMockPtrIterScan.(func(recviter *Iter, dest ...interface{}) (reta bool))(recviter, dest...)
	} else {
		panic("FuncAuxMockPtrIterScan ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterScan()
	return
}

//
// Mock: (recvb *Batch)WithContext(argctx context.Context)(reta *Batch)
//

type MockArgsTypeBatchWithContext struct {
	ApomockCallNumber int
	Argctx            context.Context
}

var LastMockArgsBatchWithContext MockArgsTypeBatchWithContext

// (recvb *Batch)AuxMockWithContext(argctx context.Context)(reta *Batch) - Generated mock function
func (recvb *Batch) AuxMockWithContext(argctx context.Context) (reta *Batch) {
	LastMockArgsBatchWithContext = MockArgsTypeBatchWithContext{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrBatchWithContext(),
		Argctx:            argctx,
	}
	rargs, rerr := apomock.GetNext("gocql.Batch.WithContext")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Batch.WithContext")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Batch.WithContext")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Batch)
	}
	return
}

// RecorderAuxMockPtrBatchWithContext  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrBatchWithContext int = 0

var condRecorderAuxMockPtrBatchWithContext *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrBatchWithContext(i int) {
	condRecorderAuxMockPtrBatchWithContext.L.Lock()
	for recorderAuxMockPtrBatchWithContext < i {
		condRecorderAuxMockPtrBatchWithContext.Wait()
	}
	condRecorderAuxMockPtrBatchWithContext.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrBatchWithContext() {
	condRecorderAuxMockPtrBatchWithContext.L.Lock()
	recorderAuxMockPtrBatchWithContext++
	condRecorderAuxMockPtrBatchWithContext.L.Unlock()
	condRecorderAuxMockPtrBatchWithContext.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrBatchWithContext() (ret int) {
	condRecorderAuxMockPtrBatchWithContext.L.Lock()
	ret = recorderAuxMockPtrBatchWithContext
	condRecorderAuxMockPtrBatchWithContext.L.Unlock()
	return
}

// (recvb *Batch)WithContext - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvb *Batch) WithContext(argctx context.Context) (reta *Batch) {
	FuncAuxMockPtrBatchWithContext, ok := apomock.GetRegisteredFunc("gocql.Batch.WithContext")
	if ok {
		reta = FuncAuxMockPtrBatchWithContext.(func(recvb *Batch, argctx context.Context) (reta *Batch))(recvb, argctx)
	} else {
		panic("FuncAuxMockPtrBatchWithContext ")
	}
	AuxMockIncrementRecorderAuxMockPtrBatchWithContext()
	return
}

//
// Mock: (recvs *Session)executeQuery(argqry *Query)(reta *Iter)
//

type MockArgsTypeSessionexecuteQuery struct {
	ApomockCallNumber int
	Argqry            *Query
}

var LastMockArgsSessionexecuteQuery MockArgsTypeSessionexecuteQuery

// (recvs *Session)AuxMockexecuteQuery(argqry *Query)(reta *Iter) - Generated mock function
func (recvs *Session) AuxMockexecuteQuery(argqry *Query) (reta *Iter) {
	LastMockArgsSessionexecuteQuery = MockArgsTypeSessionexecuteQuery{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionexecuteQuery(),
		Argqry:            argqry,
	}
	rargs, rerr := apomock.GetNext("gocql.Session.executeQuery")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Session.executeQuery")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Session.executeQuery")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Iter)
	}
	return
}

// RecorderAuxMockPtrSessionexecuteQuery  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionexecuteQuery int = 0

var condRecorderAuxMockPtrSessionexecuteQuery *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionexecuteQuery(i int) {
	condRecorderAuxMockPtrSessionexecuteQuery.L.Lock()
	for recorderAuxMockPtrSessionexecuteQuery < i {
		condRecorderAuxMockPtrSessionexecuteQuery.Wait()
	}
	condRecorderAuxMockPtrSessionexecuteQuery.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionexecuteQuery() {
	condRecorderAuxMockPtrSessionexecuteQuery.L.Lock()
	recorderAuxMockPtrSessionexecuteQuery++
	condRecorderAuxMockPtrSessionexecuteQuery.L.Unlock()
	condRecorderAuxMockPtrSessionexecuteQuery.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionexecuteQuery() (ret int) {
	condRecorderAuxMockPtrSessionexecuteQuery.L.Lock()
	ret = recorderAuxMockPtrSessionexecuteQuery
	condRecorderAuxMockPtrSessionexecuteQuery.L.Unlock()
	return
}

// (recvs *Session)executeQuery - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) executeQuery(argqry *Query) (reta *Iter) {
	FuncAuxMockPtrSessionexecuteQuery, ok := apomock.GetRegisteredFunc("gocql.Session.executeQuery")
	if ok {
		reta = FuncAuxMockPtrSessionexecuteQuery.(func(recvs *Session, argqry *Query) (reta *Iter))(recvs, argqry)
	} else {
		panic("FuncAuxMockPtrSessionexecuteQuery ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionexecuteQuery()
	return
}
