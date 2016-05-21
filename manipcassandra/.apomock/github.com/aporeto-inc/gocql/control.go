// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

import "math/rand"

import "errors"
import "net"
import "sync/atomic"

func init() {
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructControlConn, apomockNewStructControlConn)

	apomock.RegisterFunc("gocql", "gocql.createControlConn", AuxMockcreateControlConn)
	apomock.RegisterFunc("gocql", "gocql.controlConn.heartBeat", (*controlConn).AuxMockheartBeat)
	apomock.RegisterFunc("gocql", "gocql.controlConn.reconnect", (*controlConn).AuxMockreconnect)
	apomock.RegisterFunc("gocql", "gocql.controlConn.writeFrame", (*controlConn).AuxMockwriteFrame)
	apomock.RegisterFunc("gocql", "gocql.controlConn.shuffleDial", (*controlConn).AuxMockshuffleDial)
	apomock.RegisterFunc("gocql", "gocql.controlConn.connect", (*controlConn).AuxMockconnect)
	apomock.RegisterFunc("gocql", "gocql.controlConn.HandleError", (*controlConn).AuxMockHandleError)
	apomock.RegisterFunc("gocql", "gocql.controlConn.withConn", (*controlConn).AuxMockwithConn)
	apomock.RegisterFunc("gocql", "gocql.controlConn.awaitSchemaAgreement", (*controlConn).AuxMockawaitSchemaAgreement)
	apomock.RegisterFunc("gocql", "gocql.hostInfo", AuxMockhostInfo)
	apomock.RegisterFunc("gocql", "gocql.controlConn.registerEvents", (*controlConn).AuxMockregisterEvents)
	apomock.RegisterFunc("gocql", "gocql.controlConn.query", (*controlConn).AuxMockquery)
	apomock.RegisterFunc("gocql", "gocql.controlConn.fetchHostInfo", (*controlConn).AuxMockfetchHostInfo)
	apomock.RegisterFunc("gocql", "gocql.controlConn.setupConn", (*controlConn).AuxMocksetupConn)
	apomock.RegisterFunc("gocql", "gocql.controlConn.addr", (*controlConn).AuxMockaddr)
	apomock.RegisterFunc("gocql", "gocql.controlConn.close", (*controlConn).AuxMockclose)
}

const (
	ApomockStructControlConn = 8
)

var (
	randr *rand.Rand
)

var errNoControl = errors.New("gocql: no control connection available")

//
// Internal Types: in this package and their exportable versions
//
type controlConn struct {
	session *Session
	conn    atomic.Value
	retry   RetryPolicy
	started int32
	quit    chan struct{}
}

//
// External Types: in this package
//

func apomockNewStructControlConn() interface{} { return &controlConn{} }

//
// Mock: createControlConn(argsession *Session)(reta *controlConn)
//

type MockArgsTypecreateControlConn struct {
	ApomockCallNumber int
	Argsession        *Session
}

var LastMockArgscreateControlConn MockArgsTypecreateControlConn

// AuxMockcreateControlConn(argsession *Session)(reta *controlConn) - Generated mock function
func AuxMockcreateControlConn(argsession *Session) (reta *controlConn) {
	LastMockArgscreateControlConn = MockArgsTypecreateControlConn{
		ApomockCallNumber: AuxMockGetRecorderAuxMockcreateControlConn(),
		Argsession:        argsession,
	}
	rargs, rerr := apomock.GetNext("gocql.createControlConn")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.createControlConn")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.createControlConn")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*controlConn)
	}
	return
}

// RecorderAuxMockcreateControlConn  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockcreateControlConn int = 0

var condRecorderAuxMockcreateControlConn *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockcreateControlConn(i int) {
	condRecorderAuxMockcreateControlConn.L.Lock()
	for recorderAuxMockcreateControlConn < i {
		condRecorderAuxMockcreateControlConn.Wait()
	}
	condRecorderAuxMockcreateControlConn.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockcreateControlConn() {
	condRecorderAuxMockcreateControlConn.L.Lock()
	recorderAuxMockcreateControlConn++
	condRecorderAuxMockcreateControlConn.L.Unlock()
	condRecorderAuxMockcreateControlConn.Broadcast()
}
func AuxMockGetRecorderAuxMockcreateControlConn() (ret int) {
	condRecorderAuxMockcreateControlConn.L.Lock()
	ret = recorderAuxMockcreateControlConn
	condRecorderAuxMockcreateControlConn.L.Unlock()
	return
}

// createControlConn - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func createControlConn(argsession *Session) (reta *controlConn) {
	FuncAuxMockcreateControlConn, ok := apomock.GetRegisteredFunc("gocql.createControlConn")
	if ok {
		reta = FuncAuxMockcreateControlConn.(func(argsession *Session) (reta *controlConn))(argsession)
	} else {
		panic("FuncAuxMockcreateControlConn ")
	}
	AuxMockIncrementRecorderAuxMockcreateControlConn()
	return
}

//
// Mock: (recvc *controlConn)heartBeat()()
//

type MockArgsTypecontrolConnheartBeat struct {
	ApomockCallNumber int
}

var LastMockArgscontrolConnheartBeat MockArgsTypecontrolConnheartBeat

// (recvc *controlConn)AuxMockheartBeat()() - Generated mock function
func (recvc *controlConn) AuxMockheartBeat() {
	return
}

// RecorderAuxMockPtrcontrolConnheartBeat  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnheartBeat int = 0

var condRecorderAuxMockPtrcontrolConnheartBeat *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnheartBeat(i int) {
	condRecorderAuxMockPtrcontrolConnheartBeat.L.Lock()
	for recorderAuxMockPtrcontrolConnheartBeat < i {
		condRecorderAuxMockPtrcontrolConnheartBeat.Wait()
	}
	condRecorderAuxMockPtrcontrolConnheartBeat.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnheartBeat() {
	condRecorderAuxMockPtrcontrolConnheartBeat.L.Lock()
	recorderAuxMockPtrcontrolConnheartBeat++
	condRecorderAuxMockPtrcontrolConnheartBeat.L.Unlock()
	condRecorderAuxMockPtrcontrolConnheartBeat.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnheartBeat() (ret int) {
	condRecorderAuxMockPtrcontrolConnheartBeat.L.Lock()
	ret = recorderAuxMockPtrcontrolConnheartBeat
	condRecorderAuxMockPtrcontrolConnheartBeat.L.Unlock()
	return
}

// (recvc *controlConn)heartBeat - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) heartBeat() {
	FuncAuxMockPtrcontrolConnheartBeat, ok := apomock.GetRegisteredFunc("gocql.controlConn.heartBeat")
	if ok {
		FuncAuxMockPtrcontrolConnheartBeat.(func(recvc *controlConn))(recvc)
	} else {
		panic("FuncAuxMockPtrcontrolConnheartBeat ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnheartBeat()
	return
}

//
// Mock: (recvc *controlConn)reconnect(argrefreshring bool)()
//

type MockArgsTypecontrolConnreconnect struct {
	ApomockCallNumber int
	Argrefreshring    bool
}

var LastMockArgscontrolConnreconnect MockArgsTypecontrolConnreconnect

// (recvc *controlConn)AuxMockreconnect(argrefreshring bool)() - Generated mock function
func (recvc *controlConn) AuxMockreconnect(argrefreshring bool) {
	LastMockArgscontrolConnreconnect = MockArgsTypecontrolConnreconnect{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcontrolConnreconnect(),
		Argrefreshring:    argrefreshring,
	}
	return
}

// RecorderAuxMockPtrcontrolConnreconnect  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnreconnect int = 0

var condRecorderAuxMockPtrcontrolConnreconnect *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnreconnect(i int) {
	condRecorderAuxMockPtrcontrolConnreconnect.L.Lock()
	for recorderAuxMockPtrcontrolConnreconnect < i {
		condRecorderAuxMockPtrcontrolConnreconnect.Wait()
	}
	condRecorderAuxMockPtrcontrolConnreconnect.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnreconnect() {
	condRecorderAuxMockPtrcontrolConnreconnect.L.Lock()
	recorderAuxMockPtrcontrolConnreconnect++
	condRecorderAuxMockPtrcontrolConnreconnect.L.Unlock()
	condRecorderAuxMockPtrcontrolConnreconnect.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnreconnect() (ret int) {
	condRecorderAuxMockPtrcontrolConnreconnect.L.Lock()
	ret = recorderAuxMockPtrcontrolConnreconnect
	condRecorderAuxMockPtrcontrolConnreconnect.L.Unlock()
	return
}

// (recvc *controlConn)reconnect - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) reconnect(argrefreshring bool) {
	FuncAuxMockPtrcontrolConnreconnect, ok := apomock.GetRegisteredFunc("gocql.controlConn.reconnect")
	if ok {
		FuncAuxMockPtrcontrolConnreconnect.(func(recvc *controlConn, argrefreshring bool))(recvc, argrefreshring)
	} else {
		panic("FuncAuxMockPtrcontrolConnreconnect ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnreconnect()
	return
}

//
// Mock: (recvc *controlConn)writeFrame(argw frameWriter)(reta frame, retb error)
//

type MockArgsTypecontrolConnwriteFrame struct {
	ApomockCallNumber int
	Argw              frameWriter
}

var LastMockArgscontrolConnwriteFrame MockArgsTypecontrolConnwriteFrame

// (recvc *controlConn)AuxMockwriteFrame(argw frameWriter)(reta frame, retb error) - Generated mock function
func (recvc *controlConn) AuxMockwriteFrame(argw frameWriter) (reta frame, retb error) {
	LastMockArgscontrolConnwriteFrame = MockArgsTypecontrolConnwriteFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcontrolConnwriteFrame(),
		Argw:              argw,
	}
	rargs, rerr := apomock.GetNext("gocql.controlConn.writeFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.controlConn.writeFrame")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.controlConn.writeFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrcontrolConnwriteFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnwriteFrame int = 0

var condRecorderAuxMockPtrcontrolConnwriteFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnwriteFrame(i int) {
	condRecorderAuxMockPtrcontrolConnwriteFrame.L.Lock()
	for recorderAuxMockPtrcontrolConnwriteFrame < i {
		condRecorderAuxMockPtrcontrolConnwriteFrame.Wait()
	}
	condRecorderAuxMockPtrcontrolConnwriteFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnwriteFrame() {
	condRecorderAuxMockPtrcontrolConnwriteFrame.L.Lock()
	recorderAuxMockPtrcontrolConnwriteFrame++
	condRecorderAuxMockPtrcontrolConnwriteFrame.L.Unlock()
	condRecorderAuxMockPtrcontrolConnwriteFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnwriteFrame() (ret int) {
	condRecorderAuxMockPtrcontrolConnwriteFrame.L.Lock()
	ret = recorderAuxMockPtrcontrolConnwriteFrame
	condRecorderAuxMockPtrcontrolConnwriteFrame.L.Unlock()
	return
}

// (recvc *controlConn)writeFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) writeFrame(argw frameWriter) (reta frame, retb error) {
	FuncAuxMockPtrcontrolConnwriteFrame, ok := apomock.GetRegisteredFunc("gocql.controlConn.writeFrame")
	if ok {
		reta, retb = FuncAuxMockPtrcontrolConnwriteFrame.(func(recvc *controlConn, argw frameWriter) (reta frame, retb error))(recvc, argw)
	} else {
		panic("FuncAuxMockPtrcontrolConnwriteFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnwriteFrame()
	return
}

//
// Mock: (recvc *controlConn)shuffleDial(argendpoints []string)(retconn *Conn, reterr error)
//

type MockArgsTypecontrolConnshuffleDial struct {
	ApomockCallNumber int
	Argendpoints      []string
}

var LastMockArgscontrolConnshuffleDial MockArgsTypecontrolConnshuffleDial

// (recvc *controlConn)AuxMockshuffleDial(argendpoints []string)(retconn *Conn, reterr error) - Generated mock function
func (recvc *controlConn) AuxMockshuffleDial(argendpoints []string) (retconn *Conn, reterr error) {
	LastMockArgscontrolConnshuffleDial = MockArgsTypecontrolConnshuffleDial{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcontrolConnshuffleDial(),
		Argendpoints:      argendpoints,
	}
	rargs, rerr := apomock.GetNext("gocql.controlConn.shuffleDial")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.controlConn.shuffleDial")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.controlConn.shuffleDial")
	}
	if rargs.GetArg(0) != nil {
		retconn = rargs.GetArg(0).(*Conn)
	}
	if rargs.GetArg(1) != nil {
		reterr = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrcontrolConnshuffleDial  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnshuffleDial int = 0

var condRecorderAuxMockPtrcontrolConnshuffleDial *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnshuffleDial(i int) {
	condRecorderAuxMockPtrcontrolConnshuffleDial.L.Lock()
	for recorderAuxMockPtrcontrolConnshuffleDial < i {
		condRecorderAuxMockPtrcontrolConnshuffleDial.Wait()
	}
	condRecorderAuxMockPtrcontrolConnshuffleDial.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnshuffleDial() {
	condRecorderAuxMockPtrcontrolConnshuffleDial.L.Lock()
	recorderAuxMockPtrcontrolConnshuffleDial++
	condRecorderAuxMockPtrcontrolConnshuffleDial.L.Unlock()
	condRecorderAuxMockPtrcontrolConnshuffleDial.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnshuffleDial() (ret int) {
	condRecorderAuxMockPtrcontrolConnshuffleDial.L.Lock()
	ret = recorderAuxMockPtrcontrolConnshuffleDial
	condRecorderAuxMockPtrcontrolConnshuffleDial.L.Unlock()
	return
}

// (recvc *controlConn)shuffleDial - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) shuffleDial(argendpoints []string) (retconn *Conn, reterr error) {
	FuncAuxMockPtrcontrolConnshuffleDial, ok := apomock.GetRegisteredFunc("gocql.controlConn.shuffleDial")
	if ok {
		retconn, reterr = FuncAuxMockPtrcontrolConnshuffleDial.(func(recvc *controlConn, argendpoints []string) (retconn *Conn, reterr error))(recvc, argendpoints)
	} else {
		panic("FuncAuxMockPtrcontrolConnshuffleDial ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnshuffleDial()
	return
}

//
// Mock: (recvc *controlConn)connect(argendpoints []string)(reta error)
//

type MockArgsTypecontrolConnconnect struct {
	ApomockCallNumber int
	Argendpoints      []string
}

var LastMockArgscontrolConnconnect MockArgsTypecontrolConnconnect

// (recvc *controlConn)AuxMockconnect(argendpoints []string)(reta error) - Generated mock function
func (recvc *controlConn) AuxMockconnect(argendpoints []string) (reta error) {
	LastMockArgscontrolConnconnect = MockArgsTypecontrolConnconnect{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcontrolConnconnect(),
		Argendpoints:      argendpoints,
	}
	rargs, rerr := apomock.GetNext("gocql.controlConn.connect")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.controlConn.connect")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.controlConn.connect")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrcontrolConnconnect  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnconnect int = 0

var condRecorderAuxMockPtrcontrolConnconnect *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnconnect(i int) {
	condRecorderAuxMockPtrcontrolConnconnect.L.Lock()
	for recorderAuxMockPtrcontrolConnconnect < i {
		condRecorderAuxMockPtrcontrolConnconnect.Wait()
	}
	condRecorderAuxMockPtrcontrolConnconnect.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnconnect() {
	condRecorderAuxMockPtrcontrolConnconnect.L.Lock()
	recorderAuxMockPtrcontrolConnconnect++
	condRecorderAuxMockPtrcontrolConnconnect.L.Unlock()
	condRecorderAuxMockPtrcontrolConnconnect.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnconnect() (ret int) {
	condRecorderAuxMockPtrcontrolConnconnect.L.Lock()
	ret = recorderAuxMockPtrcontrolConnconnect
	condRecorderAuxMockPtrcontrolConnconnect.L.Unlock()
	return
}

// (recvc *controlConn)connect - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) connect(argendpoints []string) (reta error) {
	FuncAuxMockPtrcontrolConnconnect, ok := apomock.GetRegisteredFunc("gocql.controlConn.connect")
	if ok {
		reta = FuncAuxMockPtrcontrolConnconnect.(func(recvc *controlConn, argendpoints []string) (reta error))(recvc, argendpoints)
	} else {
		panic("FuncAuxMockPtrcontrolConnconnect ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnconnect()
	return
}

//
// Mock: (recvc *controlConn)HandleError(argconn *Conn, argerr error, argclosed bool)()
//

type MockArgsTypecontrolConnHandleError struct {
	ApomockCallNumber int
	Argconn           *Conn
	Argerr            error
	Argclosed         bool
}

var LastMockArgscontrolConnHandleError MockArgsTypecontrolConnHandleError

// (recvc *controlConn)AuxMockHandleError(argconn *Conn, argerr error, argclosed bool)() - Generated mock function
func (recvc *controlConn) AuxMockHandleError(argconn *Conn, argerr error, argclosed bool) {
	LastMockArgscontrolConnHandleError = MockArgsTypecontrolConnHandleError{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcontrolConnHandleError(),
		Argconn:           argconn,
		Argerr:            argerr,
		Argclosed:         argclosed,
	}
	return
}

// RecorderAuxMockPtrcontrolConnHandleError  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnHandleError int = 0

var condRecorderAuxMockPtrcontrolConnHandleError *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnHandleError(i int) {
	condRecorderAuxMockPtrcontrolConnHandleError.L.Lock()
	for recorderAuxMockPtrcontrolConnHandleError < i {
		condRecorderAuxMockPtrcontrolConnHandleError.Wait()
	}
	condRecorderAuxMockPtrcontrolConnHandleError.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnHandleError() {
	condRecorderAuxMockPtrcontrolConnHandleError.L.Lock()
	recorderAuxMockPtrcontrolConnHandleError++
	condRecorderAuxMockPtrcontrolConnHandleError.L.Unlock()
	condRecorderAuxMockPtrcontrolConnHandleError.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnHandleError() (ret int) {
	condRecorderAuxMockPtrcontrolConnHandleError.L.Lock()
	ret = recorderAuxMockPtrcontrolConnHandleError
	condRecorderAuxMockPtrcontrolConnHandleError.L.Unlock()
	return
}

// (recvc *controlConn)HandleError - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) HandleError(argconn *Conn, argerr error, argclosed bool) {
	FuncAuxMockPtrcontrolConnHandleError, ok := apomock.GetRegisteredFunc("gocql.controlConn.HandleError")
	if ok {
		FuncAuxMockPtrcontrolConnHandleError.(func(recvc *controlConn, argconn *Conn, argerr error, argclosed bool))(recvc, argconn, argerr, argclosed)
	} else {
		panic("FuncAuxMockPtrcontrolConnHandleError ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnHandleError()
	return
}

//
// Mock: (recvc *controlConn)withConn(argfn func(*Conn) *Iter)(reta *Iter)
//

type MockArgsTypecontrolConnwithConn struct {
	ApomockCallNumber int
	Argfn             func(*Conn) *Iter
}

var LastMockArgscontrolConnwithConn MockArgsTypecontrolConnwithConn

// (recvc *controlConn)AuxMockwithConn(argfn func(*Conn) *Iter)(reta *Iter) - Generated mock function
func (recvc *controlConn) AuxMockwithConn(argfn func(*Conn) *Iter) (reta *Iter) {
	LastMockArgscontrolConnwithConn = MockArgsTypecontrolConnwithConn{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcontrolConnwithConn(),
		Argfn:             argfn,
	}
	rargs, rerr := apomock.GetNext("gocql.controlConn.withConn")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.controlConn.withConn")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.controlConn.withConn")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Iter)
	}
	return
}

// RecorderAuxMockPtrcontrolConnwithConn  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnwithConn int = 0

var condRecorderAuxMockPtrcontrolConnwithConn *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnwithConn(i int) {
	condRecorderAuxMockPtrcontrolConnwithConn.L.Lock()
	for recorderAuxMockPtrcontrolConnwithConn < i {
		condRecorderAuxMockPtrcontrolConnwithConn.Wait()
	}
	condRecorderAuxMockPtrcontrolConnwithConn.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnwithConn() {
	condRecorderAuxMockPtrcontrolConnwithConn.L.Lock()
	recorderAuxMockPtrcontrolConnwithConn++
	condRecorderAuxMockPtrcontrolConnwithConn.L.Unlock()
	condRecorderAuxMockPtrcontrolConnwithConn.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnwithConn() (ret int) {
	condRecorderAuxMockPtrcontrolConnwithConn.L.Lock()
	ret = recorderAuxMockPtrcontrolConnwithConn
	condRecorderAuxMockPtrcontrolConnwithConn.L.Unlock()
	return
}

// (recvc *controlConn)withConn - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) withConn(argfn func(*Conn) *Iter) (reta *Iter) {
	FuncAuxMockPtrcontrolConnwithConn, ok := apomock.GetRegisteredFunc("gocql.controlConn.withConn")
	if ok {
		reta = FuncAuxMockPtrcontrolConnwithConn.(func(recvc *controlConn, argfn func(*Conn) *Iter) (reta *Iter))(recvc, argfn)
	} else {
		panic("FuncAuxMockPtrcontrolConnwithConn ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnwithConn()
	return
}

//
// Mock: (recvc *controlConn)awaitSchemaAgreement()(reta error)
//

type MockArgsTypecontrolConnawaitSchemaAgreement struct {
	ApomockCallNumber int
}

var LastMockArgscontrolConnawaitSchemaAgreement MockArgsTypecontrolConnawaitSchemaAgreement

// (recvc *controlConn)AuxMockawaitSchemaAgreement()(reta error) - Generated mock function
func (recvc *controlConn) AuxMockawaitSchemaAgreement() (reta error) {
	rargs, rerr := apomock.GetNext("gocql.controlConn.awaitSchemaAgreement")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.controlConn.awaitSchemaAgreement")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.controlConn.awaitSchemaAgreement")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrcontrolConnawaitSchemaAgreement  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnawaitSchemaAgreement int = 0

var condRecorderAuxMockPtrcontrolConnawaitSchemaAgreement *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnawaitSchemaAgreement(i int) {
	condRecorderAuxMockPtrcontrolConnawaitSchemaAgreement.L.Lock()
	for recorderAuxMockPtrcontrolConnawaitSchemaAgreement < i {
		condRecorderAuxMockPtrcontrolConnawaitSchemaAgreement.Wait()
	}
	condRecorderAuxMockPtrcontrolConnawaitSchemaAgreement.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnawaitSchemaAgreement() {
	condRecorderAuxMockPtrcontrolConnawaitSchemaAgreement.L.Lock()
	recorderAuxMockPtrcontrolConnawaitSchemaAgreement++
	condRecorderAuxMockPtrcontrolConnawaitSchemaAgreement.L.Unlock()
	condRecorderAuxMockPtrcontrolConnawaitSchemaAgreement.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnawaitSchemaAgreement() (ret int) {
	condRecorderAuxMockPtrcontrolConnawaitSchemaAgreement.L.Lock()
	ret = recorderAuxMockPtrcontrolConnawaitSchemaAgreement
	condRecorderAuxMockPtrcontrolConnawaitSchemaAgreement.L.Unlock()
	return
}

// (recvc *controlConn)awaitSchemaAgreement - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) awaitSchemaAgreement() (reta error) {
	FuncAuxMockPtrcontrolConnawaitSchemaAgreement, ok := apomock.GetRegisteredFunc("gocql.controlConn.awaitSchemaAgreement")
	if ok {
		reta = FuncAuxMockPtrcontrolConnawaitSchemaAgreement.(func(recvc *controlConn) (reta error))(recvc)
	} else {
		panic("FuncAuxMockPtrcontrolConnawaitSchemaAgreement ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnawaitSchemaAgreement()
	return
}

//
// Mock: hostInfo(argaddr string, argdefaultPort int)(reta *HostInfo, retb error)
//

type MockArgsTypehostInfo struct {
	ApomockCallNumber int
	Argaddr           string
	ArgdefaultPort    int
}

var LastMockArgshostInfo MockArgsTypehostInfo

// AuxMockhostInfo(argaddr string, argdefaultPort int)(reta *HostInfo, retb error) - Generated mock function
func AuxMockhostInfo(argaddr string, argdefaultPort int) (reta *HostInfo, retb error) {
	LastMockArgshostInfo = MockArgsTypehostInfo{
		ApomockCallNumber: AuxMockGetRecorderAuxMockhostInfo(),
		Argaddr:           argaddr,
		ArgdefaultPort:    argdefaultPort,
	}
	rargs, rerr := apomock.GetNext("gocql.hostInfo")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.hostInfo")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.hostInfo")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockhostInfo  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockhostInfo int = 0

var condRecorderAuxMockhostInfo *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockhostInfo(i int) {
	condRecorderAuxMockhostInfo.L.Lock()
	for recorderAuxMockhostInfo < i {
		condRecorderAuxMockhostInfo.Wait()
	}
	condRecorderAuxMockhostInfo.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockhostInfo() {
	condRecorderAuxMockhostInfo.L.Lock()
	recorderAuxMockhostInfo++
	condRecorderAuxMockhostInfo.L.Unlock()
	condRecorderAuxMockhostInfo.Broadcast()
}
func AuxMockGetRecorderAuxMockhostInfo() (ret int) {
	condRecorderAuxMockhostInfo.L.Lock()
	ret = recorderAuxMockhostInfo
	condRecorderAuxMockhostInfo.L.Unlock()
	return
}

// hostInfo - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func hostInfo(argaddr string, argdefaultPort int) (reta *HostInfo, retb error) {
	FuncAuxMockhostInfo, ok := apomock.GetRegisteredFunc("gocql.hostInfo")
	if ok {
		reta, retb = FuncAuxMockhostInfo.(func(argaddr string, argdefaultPort int) (reta *HostInfo, retb error))(argaddr, argdefaultPort)
	} else {
		panic("FuncAuxMockhostInfo ")
	}
	AuxMockIncrementRecorderAuxMockhostInfo()
	return
}

//
// Mock: (recvc *controlConn)registerEvents(argconn *Conn)(reta error)
//

type MockArgsTypecontrolConnregisterEvents struct {
	ApomockCallNumber int
	Argconn           *Conn
}

var LastMockArgscontrolConnregisterEvents MockArgsTypecontrolConnregisterEvents

// (recvc *controlConn)AuxMockregisterEvents(argconn *Conn)(reta error) - Generated mock function
func (recvc *controlConn) AuxMockregisterEvents(argconn *Conn) (reta error) {
	LastMockArgscontrolConnregisterEvents = MockArgsTypecontrolConnregisterEvents{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcontrolConnregisterEvents(),
		Argconn:           argconn,
	}
	rargs, rerr := apomock.GetNext("gocql.controlConn.registerEvents")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.controlConn.registerEvents")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.controlConn.registerEvents")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrcontrolConnregisterEvents  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnregisterEvents int = 0

var condRecorderAuxMockPtrcontrolConnregisterEvents *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnregisterEvents(i int) {
	condRecorderAuxMockPtrcontrolConnregisterEvents.L.Lock()
	for recorderAuxMockPtrcontrolConnregisterEvents < i {
		condRecorderAuxMockPtrcontrolConnregisterEvents.Wait()
	}
	condRecorderAuxMockPtrcontrolConnregisterEvents.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnregisterEvents() {
	condRecorderAuxMockPtrcontrolConnregisterEvents.L.Lock()
	recorderAuxMockPtrcontrolConnregisterEvents++
	condRecorderAuxMockPtrcontrolConnregisterEvents.L.Unlock()
	condRecorderAuxMockPtrcontrolConnregisterEvents.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnregisterEvents() (ret int) {
	condRecorderAuxMockPtrcontrolConnregisterEvents.L.Lock()
	ret = recorderAuxMockPtrcontrolConnregisterEvents
	condRecorderAuxMockPtrcontrolConnregisterEvents.L.Unlock()
	return
}

// (recvc *controlConn)registerEvents - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) registerEvents(argconn *Conn) (reta error) {
	FuncAuxMockPtrcontrolConnregisterEvents, ok := apomock.GetRegisteredFunc("gocql.controlConn.registerEvents")
	if ok {
		reta = FuncAuxMockPtrcontrolConnregisterEvents.(func(recvc *controlConn, argconn *Conn) (reta error))(recvc, argconn)
	} else {
		panic("FuncAuxMockPtrcontrolConnregisterEvents ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnregisterEvents()
	return
}

//
// Mock: (recvc *controlConn)query(argstatement string, values ...interface{})(retiter *Iter)
//

type MockArgsTypecontrolConnquery struct {
	ApomockCallNumber int
	Argstatement      string
	Values            []interface{}
}

var LastMockArgscontrolConnquery MockArgsTypecontrolConnquery

// (recvc *controlConn)AuxMockquery(argstatement string, values ...interface{})(retiter *Iter) - Generated mock function
func (recvc *controlConn) AuxMockquery(argstatement string, values ...interface{}) (retiter *Iter) {
	LastMockArgscontrolConnquery = MockArgsTypecontrolConnquery{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcontrolConnquery(),
		Argstatement:      argstatement,
		Values:            values,
	}
	rargs, rerr := apomock.GetNext("gocql.controlConn.query")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.controlConn.query")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.controlConn.query")
	}
	if rargs.GetArg(0) != nil {
		retiter = rargs.GetArg(0).(*Iter)
	}
	return
}

// RecorderAuxMockPtrcontrolConnquery  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnquery int = 0

var condRecorderAuxMockPtrcontrolConnquery *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnquery(i int) {
	condRecorderAuxMockPtrcontrolConnquery.L.Lock()
	for recorderAuxMockPtrcontrolConnquery < i {
		condRecorderAuxMockPtrcontrolConnquery.Wait()
	}
	condRecorderAuxMockPtrcontrolConnquery.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnquery() {
	condRecorderAuxMockPtrcontrolConnquery.L.Lock()
	recorderAuxMockPtrcontrolConnquery++
	condRecorderAuxMockPtrcontrolConnquery.L.Unlock()
	condRecorderAuxMockPtrcontrolConnquery.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnquery() (ret int) {
	condRecorderAuxMockPtrcontrolConnquery.L.Lock()
	ret = recorderAuxMockPtrcontrolConnquery
	condRecorderAuxMockPtrcontrolConnquery.L.Unlock()
	return
}

// (recvc *controlConn)query - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) query(argstatement string, values ...interface{}) (retiter *Iter) {
	FuncAuxMockPtrcontrolConnquery, ok := apomock.GetRegisteredFunc("gocql.controlConn.query")
	if ok {
		retiter = FuncAuxMockPtrcontrolConnquery.(func(recvc *controlConn, argstatement string, values ...interface{}) (retiter *Iter))(recvc, argstatement, values...)
	} else {
		panic("FuncAuxMockPtrcontrolConnquery ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnquery()
	return
}

//
// Mock: (recvc *controlConn)fetchHostInfo(argaddr net.IP, argport int)(reta *HostInfo, retb error)
//

type MockArgsTypecontrolConnfetchHostInfo struct {
	ApomockCallNumber int
	Argaddr           net.IP
	Argport           int
}

var LastMockArgscontrolConnfetchHostInfo MockArgsTypecontrolConnfetchHostInfo

// (recvc *controlConn)AuxMockfetchHostInfo(argaddr net.IP, argport int)(reta *HostInfo, retb error) - Generated mock function
func (recvc *controlConn) AuxMockfetchHostInfo(argaddr net.IP, argport int) (reta *HostInfo, retb error) {
	LastMockArgscontrolConnfetchHostInfo = MockArgsTypecontrolConnfetchHostInfo{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcontrolConnfetchHostInfo(),
		Argaddr:           argaddr,
		Argport:           argport,
	}
	rargs, rerr := apomock.GetNext("gocql.controlConn.fetchHostInfo")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.controlConn.fetchHostInfo")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.controlConn.fetchHostInfo")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrcontrolConnfetchHostInfo  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnfetchHostInfo int = 0

var condRecorderAuxMockPtrcontrolConnfetchHostInfo *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnfetchHostInfo(i int) {
	condRecorderAuxMockPtrcontrolConnfetchHostInfo.L.Lock()
	for recorderAuxMockPtrcontrolConnfetchHostInfo < i {
		condRecorderAuxMockPtrcontrolConnfetchHostInfo.Wait()
	}
	condRecorderAuxMockPtrcontrolConnfetchHostInfo.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnfetchHostInfo() {
	condRecorderAuxMockPtrcontrolConnfetchHostInfo.L.Lock()
	recorderAuxMockPtrcontrolConnfetchHostInfo++
	condRecorderAuxMockPtrcontrolConnfetchHostInfo.L.Unlock()
	condRecorderAuxMockPtrcontrolConnfetchHostInfo.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnfetchHostInfo() (ret int) {
	condRecorderAuxMockPtrcontrolConnfetchHostInfo.L.Lock()
	ret = recorderAuxMockPtrcontrolConnfetchHostInfo
	condRecorderAuxMockPtrcontrolConnfetchHostInfo.L.Unlock()
	return
}

// (recvc *controlConn)fetchHostInfo - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) fetchHostInfo(argaddr net.IP, argport int) (reta *HostInfo, retb error) {
	FuncAuxMockPtrcontrolConnfetchHostInfo, ok := apomock.GetRegisteredFunc("gocql.controlConn.fetchHostInfo")
	if ok {
		reta, retb = FuncAuxMockPtrcontrolConnfetchHostInfo.(func(recvc *controlConn, argaddr net.IP, argport int) (reta *HostInfo, retb error))(recvc, argaddr, argport)
	} else {
		panic("FuncAuxMockPtrcontrolConnfetchHostInfo ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnfetchHostInfo()
	return
}

//
// Mock: (recvc *controlConn)setupConn(argconn *Conn)(reta error)
//

type MockArgsTypecontrolConnsetupConn struct {
	ApomockCallNumber int
	Argconn           *Conn
}

var LastMockArgscontrolConnsetupConn MockArgsTypecontrolConnsetupConn

// (recvc *controlConn)AuxMocksetupConn(argconn *Conn)(reta error) - Generated mock function
func (recvc *controlConn) AuxMocksetupConn(argconn *Conn) (reta error) {
	LastMockArgscontrolConnsetupConn = MockArgsTypecontrolConnsetupConn{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcontrolConnsetupConn(),
		Argconn:           argconn,
	}
	rargs, rerr := apomock.GetNext("gocql.controlConn.setupConn")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.controlConn.setupConn")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.controlConn.setupConn")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrcontrolConnsetupConn  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnsetupConn int = 0

var condRecorderAuxMockPtrcontrolConnsetupConn *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnsetupConn(i int) {
	condRecorderAuxMockPtrcontrolConnsetupConn.L.Lock()
	for recorderAuxMockPtrcontrolConnsetupConn < i {
		condRecorderAuxMockPtrcontrolConnsetupConn.Wait()
	}
	condRecorderAuxMockPtrcontrolConnsetupConn.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnsetupConn() {
	condRecorderAuxMockPtrcontrolConnsetupConn.L.Lock()
	recorderAuxMockPtrcontrolConnsetupConn++
	condRecorderAuxMockPtrcontrolConnsetupConn.L.Unlock()
	condRecorderAuxMockPtrcontrolConnsetupConn.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnsetupConn() (ret int) {
	condRecorderAuxMockPtrcontrolConnsetupConn.L.Lock()
	ret = recorderAuxMockPtrcontrolConnsetupConn
	condRecorderAuxMockPtrcontrolConnsetupConn.L.Unlock()
	return
}

// (recvc *controlConn)setupConn - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) setupConn(argconn *Conn) (reta error) {
	FuncAuxMockPtrcontrolConnsetupConn, ok := apomock.GetRegisteredFunc("gocql.controlConn.setupConn")
	if ok {
		reta = FuncAuxMockPtrcontrolConnsetupConn.(func(recvc *controlConn, argconn *Conn) (reta error))(recvc, argconn)
	} else {
		panic("FuncAuxMockPtrcontrolConnsetupConn ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnsetupConn()
	return
}

//
// Mock: (recvc *controlConn)addr()(reta string)
//

type MockArgsTypecontrolConnaddr struct {
	ApomockCallNumber int
}

var LastMockArgscontrolConnaddr MockArgsTypecontrolConnaddr

// (recvc *controlConn)AuxMockaddr()(reta string) - Generated mock function
func (recvc *controlConn) AuxMockaddr() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.controlConn.addr")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.controlConn.addr")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.controlConn.addr")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrcontrolConnaddr  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnaddr int = 0

var condRecorderAuxMockPtrcontrolConnaddr *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnaddr(i int) {
	condRecorderAuxMockPtrcontrolConnaddr.L.Lock()
	for recorderAuxMockPtrcontrolConnaddr < i {
		condRecorderAuxMockPtrcontrolConnaddr.Wait()
	}
	condRecorderAuxMockPtrcontrolConnaddr.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnaddr() {
	condRecorderAuxMockPtrcontrolConnaddr.L.Lock()
	recorderAuxMockPtrcontrolConnaddr++
	condRecorderAuxMockPtrcontrolConnaddr.L.Unlock()
	condRecorderAuxMockPtrcontrolConnaddr.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnaddr() (ret int) {
	condRecorderAuxMockPtrcontrolConnaddr.L.Lock()
	ret = recorderAuxMockPtrcontrolConnaddr
	condRecorderAuxMockPtrcontrolConnaddr.L.Unlock()
	return
}

// (recvc *controlConn)addr - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) addr() (reta string) {
	FuncAuxMockPtrcontrolConnaddr, ok := apomock.GetRegisteredFunc("gocql.controlConn.addr")
	if ok {
		reta = FuncAuxMockPtrcontrolConnaddr.(func(recvc *controlConn) (reta string))(recvc)
	} else {
		panic("FuncAuxMockPtrcontrolConnaddr ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnaddr()
	return
}

//
// Mock: (recvc *controlConn)close()()
//

type MockArgsTypecontrolConnclose struct {
	ApomockCallNumber int
}

var LastMockArgscontrolConnclose MockArgsTypecontrolConnclose

// (recvc *controlConn)AuxMockclose()() - Generated mock function
func (recvc *controlConn) AuxMockclose() {
	return
}

// RecorderAuxMockPtrcontrolConnclose  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcontrolConnclose int = 0

var condRecorderAuxMockPtrcontrolConnclose *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcontrolConnclose(i int) {
	condRecorderAuxMockPtrcontrolConnclose.L.Lock()
	for recorderAuxMockPtrcontrolConnclose < i {
		condRecorderAuxMockPtrcontrolConnclose.Wait()
	}
	condRecorderAuxMockPtrcontrolConnclose.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcontrolConnclose() {
	condRecorderAuxMockPtrcontrolConnclose.L.Lock()
	recorderAuxMockPtrcontrolConnclose++
	condRecorderAuxMockPtrcontrolConnclose.L.Unlock()
	condRecorderAuxMockPtrcontrolConnclose.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcontrolConnclose() (ret int) {
	condRecorderAuxMockPtrcontrolConnclose.L.Lock()
	ret = recorderAuxMockPtrcontrolConnclose
	condRecorderAuxMockPtrcontrolConnclose.L.Unlock()
	return
}

// (recvc *controlConn)close - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *controlConn) close() {
	FuncAuxMockPtrcontrolConnclose, ok := apomock.GetRegisteredFunc("gocql.controlConn.close")
	if ok {
		FuncAuxMockPtrcontrolConnclose.(func(recvc *controlConn))(recvc)
	} else {
		panic("FuncAuxMockPtrcontrolConnclose ")
	}
	AuxMockIncrementRecorderAuxMockPtrcontrolConnclose()
	return
}
