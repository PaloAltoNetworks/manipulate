// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructErrorFrame, apomockNewStructErrorFrame)

	apomock.RegisterFunc("gocql", "gocql.errorFrame.Message", (errorFrame).AuxMockMessage)
	apomock.RegisterFunc("gocql", "gocql.errorFrame.Error", (errorFrame).AuxMockError)
	apomock.RegisterFunc("gocql", "gocql.errorFrame.String", (errorFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.RequestErrUnavailable.String", (*RequestErrUnavailable).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.errorFrame.Code", (errorFrame).AuxMockCode)
}

const (
	errServer          = 0x0000
	errProtocol        = 0x000A
	errCredentials     = 0x0100
	errUnavailable     = 0x1000
	errOverloaded      = 0x1001
	errBootstrapping   = 0x1002
	errTruncate        = 0x1003
	errWriteTimeout    = 0x1100
	errReadTimeout     = 0x1200
	errReadFailure     = 0x1300
	errFunctionFailure = 0x1400
	errWriteFailure    = 0x1500
	errSyntax          = 0x2000
	errUnauthorized    = 0x2100
	errInvalid         = 0x2200
	errConfig          = 0x2300
	errAlreadyExists   = 0x2400
	errUnprepared      = 0x2500
)

const (
	ApomockStructErrorFrame = 46
)

//
// Internal Types: in this package and their exportable versions
//
type errorFrame struct {
	frameHeader
	code    int
	message string
}

//
// External Types: in this package
//
type RequestErrWriteFailure struct {
	errorFrame
	Consistency Consistency
	Received    int
	BlockFor    int
	NumFailures int
	WriteType   string
}

type RequestErrAlreadyExists struct {
	errorFrame
	Keyspace string
	Table    string
}

type RequestErrReadFailure struct {
	errorFrame
	Consistency Consistency
	Received    int
	BlockFor    int
	NumFailures int
	DataPresent bool
}

type RequestError interface {
	Code() int
	Message() string
	Error() string
}

type RequestErrUnavailable struct {
	errorFrame
	Consistency Consistency
	Required    int
	Alive       int
}

type RequestErrWriteTimeout struct {
	errorFrame
	Consistency Consistency
	Received    int
	BlockFor    int
	WriteType   string
}

type RequestErrReadTimeout struct {
	errorFrame
	Consistency Consistency
	Received    int
	BlockFor    int
	DataPresent byte
}

type RequestErrUnprepared struct {
	errorFrame
	StatementId []byte
}

type RequestErrFunctionFailure struct {
	errorFrame
	Keyspace string
	Function string
	ArgTypes []string
}

func apomockNewStructErrorFrame() interface{} { return &errorFrame{} }

//
// Mock: (recve errorFrame)Message()(reta string)
//

type MockArgsTypeerrorFrameMessage struct {
	ApomockCallNumber int
}

var LastMockArgserrorFrameMessage MockArgsTypeerrorFrameMessage

// (recve errorFrame)AuxMockMessage()(reta string) - Generated mock function
func (recve errorFrame) AuxMockMessage() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.errorFrame.Message")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.errorFrame.Message")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.errorFrame.Message")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockerrorFrameMessage  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockerrorFrameMessage int = 0

var condRecorderAuxMockerrorFrameMessage *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockerrorFrameMessage(i int) {
	condRecorderAuxMockerrorFrameMessage.L.Lock()
	for recorderAuxMockerrorFrameMessage < i {
		condRecorderAuxMockerrorFrameMessage.Wait()
	}
	condRecorderAuxMockerrorFrameMessage.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockerrorFrameMessage() {
	condRecorderAuxMockerrorFrameMessage.L.Lock()
	recorderAuxMockerrorFrameMessage++
	condRecorderAuxMockerrorFrameMessage.L.Unlock()
	condRecorderAuxMockerrorFrameMessage.Broadcast()
}
func AuxMockGetRecorderAuxMockerrorFrameMessage() (ret int) {
	condRecorderAuxMockerrorFrameMessage.L.Lock()
	ret = recorderAuxMockerrorFrameMessage
	condRecorderAuxMockerrorFrameMessage.L.Unlock()
	return
}

// (recve errorFrame)Message - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve errorFrame) Message() (reta string) {
	FuncAuxMockerrorFrameMessage, ok := apomock.GetRegisteredFunc("gocql.errorFrame.Message")
	if ok {
		reta = FuncAuxMockerrorFrameMessage.(func(recve errorFrame) (reta string))(recve)
	} else {
		panic("FuncAuxMockerrorFrameMessage ")
	}
	AuxMockIncrementRecorderAuxMockerrorFrameMessage()
	return
}

//
// Mock: (recve errorFrame)Error()(reta string)
//

type MockArgsTypeerrorFrameError struct {
	ApomockCallNumber int
}

var LastMockArgserrorFrameError MockArgsTypeerrorFrameError

// (recve errorFrame)AuxMockError()(reta string) - Generated mock function
func (recve errorFrame) AuxMockError() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.errorFrame.Error")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.errorFrame.Error")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.errorFrame.Error")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockerrorFrameError  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockerrorFrameError int = 0

var condRecorderAuxMockerrorFrameError *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockerrorFrameError(i int) {
	condRecorderAuxMockerrorFrameError.L.Lock()
	for recorderAuxMockerrorFrameError < i {
		condRecorderAuxMockerrorFrameError.Wait()
	}
	condRecorderAuxMockerrorFrameError.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockerrorFrameError() {
	condRecorderAuxMockerrorFrameError.L.Lock()
	recorderAuxMockerrorFrameError++
	condRecorderAuxMockerrorFrameError.L.Unlock()
	condRecorderAuxMockerrorFrameError.Broadcast()
}
func AuxMockGetRecorderAuxMockerrorFrameError() (ret int) {
	condRecorderAuxMockerrorFrameError.L.Lock()
	ret = recorderAuxMockerrorFrameError
	condRecorderAuxMockerrorFrameError.L.Unlock()
	return
}

// (recve errorFrame)Error - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve errorFrame) Error() (reta string) {
	FuncAuxMockerrorFrameError, ok := apomock.GetRegisteredFunc("gocql.errorFrame.Error")
	if ok {
		reta = FuncAuxMockerrorFrameError.(func(recve errorFrame) (reta string))(recve)
	} else {
		panic("FuncAuxMockerrorFrameError ")
	}
	AuxMockIncrementRecorderAuxMockerrorFrameError()
	return
}

//
// Mock: (recve errorFrame)String()(reta string)
//

type MockArgsTypeerrorFrameString struct {
	ApomockCallNumber int
}

var LastMockArgserrorFrameString MockArgsTypeerrorFrameString

// (recve errorFrame)AuxMockString()(reta string) - Generated mock function
func (recve errorFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.errorFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.errorFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.errorFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockerrorFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockerrorFrameString int = 0

var condRecorderAuxMockerrorFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockerrorFrameString(i int) {
	condRecorderAuxMockerrorFrameString.L.Lock()
	for recorderAuxMockerrorFrameString < i {
		condRecorderAuxMockerrorFrameString.Wait()
	}
	condRecorderAuxMockerrorFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockerrorFrameString() {
	condRecorderAuxMockerrorFrameString.L.Lock()
	recorderAuxMockerrorFrameString++
	condRecorderAuxMockerrorFrameString.L.Unlock()
	condRecorderAuxMockerrorFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockerrorFrameString() (ret int) {
	condRecorderAuxMockerrorFrameString.L.Lock()
	ret = recorderAuxMockerrorFrameString
	condRecorderAuxMockerrorFrameString.L.Unlock()
	return
}

// (recve errorFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve errorFrame) String() (reta string) {
	FuncAuxMockerrorFrameString, ok := apomock.GetRegisteredFunc("gocql.errorFrame.String")
	if ok {
		reta = FuncAuxMockerrorFrameString.(func(recve errorFrame) (reta string))(recve)
	} else {
		panic("FuncAuxMockerrorFrameString ")
	}
	AuxMockIncrementRecorderAuxMockerrorFrameString()
	return
}

//
// Mock: (recve *RequestErrUnavailable)String()(reta string)
//

type MockArgsTypeRequestErrUnavailableString struct {
	ApomockCallNumber int
}

var LastMockArgsRequestErrUnavailableString MockArgsTypeRequestErrUnavailableString

// (recve *RequestErrUnavailable)AuxMockString()(reta string) - Generated mock function
func (recve *RequestErrUnavailable) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.RequestErrUnavailable.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.RequestErrUnavailable.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.RequestErrUnavailable.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrRequestErrUnavailableString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrRequestErrUnavailableString int = 0

var condRecorderAuxMockPtrRequestErrUnavailableString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrRequestErrUnavailableString(i int) {
	condRecorderAuxMockPtrRequestErrUnavailableString.L.Lock()
	for recorderAuxMockPtrRequestErrUnavailableString < i {
		condRecorderAuxMockPtrRequestErrUnavailableString.Wait()
	}
	condRecorderAuxMockPtrRequestErrUnavailableString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrRequestErrUnavailableString() {
	condRecorderAuxMockPtrRequestErrUnavailableString.L.Lock()
	recorderAuxMockPtrRequestErrUnavailableString++
	condRecorderAuxMockPtrRequestErrUnavailableString.L.Unlock()
	condRecorderAuxMockPtrRequestErrUnavailableString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrRequestErrUnavailableString() (ret int) {
	condRecorderAuxMockPtrRequestErrUnavailableString.L.Lock()
	ret = recorderAuxMockPtrRequestErrUnavailableString
	condRecorderAuxMockPtrRequestErrUnavailableString.L.Unlock()
	return
}

// (recve *RequestErrUnavailable)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve *RequestErrUnavailable) String() (reta string) {
	FuncAuxMockPtrRequestErrUnavailableString, ok := apomock.GetRegisteredFunc("gocql.RequestErrUnavailable.String")
	if ok {
		reta = FuncAuxMockPtrRequestErrUnavailableString.(func(recve *RequestErrUnavailable) (reta string))(recve)
	} else {
		panic("FuncAuxMockPtrRequestErrUnavailableString ")
	}
	AuxMockIncrementRecorderAuxMockPtrRequestErrUnavailableString()
	return
}

//
// Mock: (recve errorFrame)Code()(reta int)
//

type MockArgsTypeerrorFrameCode struct {
	ApomockCallNumber int
}

var LastMockArgserrorFrameCode MockArgsTypeerrorFrameCode

// (recve errorFrame)AuxMockCode()(reta int) - Generated mock function
func (recve errorFrame) AuxMockCode() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.errorFrame.Code")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.errorFrame.Code")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.errorFrame.Code")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockerrorFrameCode  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockerrorFrameCode int = 0

var condRecorderAuxMockerrorFrameCode *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockerrorFrameCode(i int) {
	condRecorderAuxMockerrorFrameCode.L.Lock()
	for recorderAuxMockerrorFrameCode < i {
		condRecorderAuxMockerrorFrameCode.Wait()
	}
	condRecorderAuxMockerrorFrameCode.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockerrorFrameCode() {
	condRecorderAuxMockerrorFrameCode.L.Lock()
	recorderAuxMockerrorFrameCode++
	condRecorderAuxMockerrorFrameCode.L.Unlock()
	condRecorderAuxMockerrorFrameCode.Broadcast()
}
func AuxMockGetRecorderAuxMockerrorFrameCode() (ret int) {
	condRecorderAuxMockerrorFrameCode.L.Lock()
	ret = recorderAuxMockerrorFrameCode
	condRecorderAuxMockerrorFrameCode.L.Unlock()
	return
}

// (recve errorFrame)Code - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve errorFrame) Code() (reta int) {
	FuncAuxMockerrorFrameCode, ok := apomock.GetRegisteredFunc("gocql.errorFrame.Code")
	if ok {
		reta = FuncAuxMockerrorFrameCode.(func(recve errorFrame) (reta int))(recve)
	} else {
		panic("FuncAuxMockerrorFrameCode ")
	}
	AuxMockIncrementRecorderAuxMockerrorFrameCode()
	return
}
