// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {

	apomock.RegisterFunc("gocql", "gocql.SnappyCompressor.Name", (SnappyCompressor).AuxMockName)
	apomock.RegisterFunc("gocql", "gocql.SnappyCompressor.Encode", (SnappyCompressor).AuxMockEncode)
	apomock.RegisterFunc("gocql", "gocql.SnappyCompressor.Decode", (SnappyCompressor).AuxMockDecode)
}

const ()

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//
type Compressor interface {
	Name() string
	Encode(data []byte) ([]byte, error)
	Decode(data []byte) ([]byte, error)
}

type SnappyCompressor struct{}

//
// Mock: (recvs SnappyCompressor)Name()(reta string)
//

type MockArgsTypeSnappyCompressorName struct {
	ApomockCallNumber int
}

var LastMockArgsSnappyCompressorName MockArgsTypeSnappyCompressorName

// (recvs SnappyCompressor)AuxMockName()(reta string) - Generated mock function
func (recvs SnappyCompressor) AuxMockName() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.SnappyCompressor.Name")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.SnappyCompressor.Name")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.SnappyCompressor.Name")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockSnappyCompressorName  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockSnappyCompressorName int = 0

var condRecorderAuxMockSnappyCompressorName *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockSnappyCompressorName(i int) {
	condRecorderAuxMockSnappyCompressorName.L.Lock()
	for recorderAuxMockSnappyCompressorName < i {
		condRecorderAuxMockSnappyCompressorName.Wait()
	}
	condRecorderAuxMockSnappyCompressorName.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockSnappyCompressorName() {
	condRecorderAuxMockSnappyCompressorName.L.Lock()
	recorderAuxMockSnappyCompressorName++
	condRecorderAuxMockSnappyCompressorName.L.Unlock()
	condRecorderAuxMockSnappyCompressorName.Broadcast()
}
func AuxMockGetRecorderAuxMockSnappyCompressorName() (ret int) {
	condRecorderAuxMockSnappyCompressorName.L.Lock()
	ret = recorderAuxMockSnappyCompressorName
	condRecorderAuxMockSnappyCompressorName.L.Unlock()
	return
}

// (recvs SnappyCompressor)Name - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs SnappyCompressor) Name() (reta string) {
	FuncAuxMockSnappyCompressorName, ok := apomock.GetRegisteredFunc("gocql.SnappyCompressor.Name")
	if ok {
		reta = FuncAuxMockSnappyCompressorName.(func(recvs SnappyCompressor) (reta string))(recvs)
	} else {
		panic("FuncAuxMockSnappyCompressorName ")
	}
	AuxMockIncrementRecorderAuxMockSnappyCompressorName()
	return
}

//
// Mock: (recvs SnappyCompressor)Encode(argdata []byte)(reta []byte, retb error)
//

type MockArgsTypeSnappyCompressorEncode struct {
	ApomockCallNumber int
	Argdata           []byte
}

var LastMockArgsSnappyCompressorEncode MockArgsTypeSnappyCompressorEncode

// (recvs SnappyCompressor)AuxMockEncode(argdata []byte)(reta []byte, retb error) - Generated mock function
func (recvs SnappyCompressor) AuxMockEncode(argdata []byte) (reta []byte, retb error) {
	LastMockArgsSnappyCompressorEncode = MockArgsTypeSnappyCompressorEncode{
		ApomockCallNumber: AuxMockGetRecorderAuxMockSnappyCompressorEncode(),
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.SnappyCompressor.Encode")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.SnappyCompressor.Encode")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.SnappyCompressor.Encode")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockSnappyCompressorEncode  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockSnappyCompressorEncode int = 0

var condRecorderAuxMockSnappyCompressorEncode *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockSnappyCompressorEncode(i int) {
	condRecorderAuxMockSnappyCompressorEncode.L.Lock()
	for recorderAuxMockSnappyCompressorEncode < i {
		condRecorderAuxMockSnappyCompressorEncode.Wait()
	}
	condRecorderAuxMockSnappyCompressorEncode.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockSnappyCompressorEncode() {
	condRecorderAuxMockSnappyCompressorEncode.L.Lock()
	recorderAuxMockSnappyCompressorEncode++
	condRecorderAuxMockSnappyCompressorEncode.L.Unlock()
	condRecorderAuxMockSnappyCompressorEncode.Broadcast()
}
func AuxMockGetRecorderAuxMockSnappyCompressorEncode() (ret int) {
	condRecorderAuxMockSnappyCompressorEncode.L.Lock()
	ret = recorderAuxMockSnappyCompressorEncode
	condRecorderAuxMockSnappyCompressorEncode.L.Unlock()
	return
}

// (recvs SnappyCompressor)Encode - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs SnappyCompressor) Encode(argdata []byte) (reta []byte, retb error) {
	FuncAuxMockSnappyCompressorEncode, ok := apomock.GetRegisteredFunc("gocql.SnappyCompressor.Encode")
	if ok {
		reta, retb = FuncAuxMockSnappyCompressorEncode.(func(recvs SnappyCompressor, argdata []byte) (reta []byte, retb error))(recvs, argdata)
	} else {
		panic("FuncAuxMockSnappyCompressorEncode ")
	}
	AuxMockIncrementRecorderAuxMockSnappyCompressorEncode()
	return
}

//
// Mock: (recvs SnappyCompressor)Decode(argdata []byte)(reta []byte, retb error)
//

type MockArgsTypeSnappyCompressorDecode struct {
	ApomockCallNumber int
	Argdata           []byte
}

var LastMockArgsSnappyCompressorDecode MockArgsTypeSnappyCompressorDecode

// (recvs SnappyCompressor)AuxMockDecode(argdata []byte)(reta []byte, retb error) - Generated mock function
func (recvs SnappyCompressor) AuxMockDecode(argdata []byte) (reta []byte, retb error) {
	LastMockArgsSnappyCompressorDecode = MockArgsTypeSnappyCompressorDecode{
		ApomockCallNumber: AuxMockGetRecorderAuxMockSnappyCompressorDecode(),
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.SnappyCompressor.Decode")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.SnappyCompressor.Decode")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.SnappyCompressor.Decode")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockSnappyCompressorDecode  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockSnappyCompressorDecode int = 0

var condRecorderAuxMockSnappyCompressorDecode *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockSnappyCompressorDecode(i int) {
	condRecorderAuxMockSnappyCompressorDecode.L.Lock()
	for recorderAuxMockSnappyCompressorDecode < i {
		condRecorderAuxMockSnappyCompressorDecode.Wait()
	}
	condRecorderAuxMockSnappyCompressorDecode.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockSnappyCompressorDecode() {
	condRecorderAuxMockSnappyCompressorDecode.L.Lock()
	recorderAuxMockSnappyCompressorDecode++
	condRecorderAuxMockSnappyCompressorDecode.L.Unlock()
	condRecorderAuxMockSnappyCompressorDecode.Broadcast()
}
func AuxMockGetRecorderAuxMockSnappyCompressorDecode() (ret int) {
	condRecorderAuxMockSnappyCompressorDecode.L.Lock()
	ret = recorderAuxMockSnappyCompressorDecode
	condRecorderAuxMockSnappyCompressorDecode.L.Unlock()
	return
}

// (recvs SnappyCompressor)Decode - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs SnappyCompressor) Decode(argdata []byte) (reta []byte, retb error) {
	FuncAuxMockSnappyCompressorDecode, ok := apomock.GetRegisteredFunc("gocql.SnappyCompressor.Decode")
	if ok {
		reta, retb = FuncAuxMockSnappyCompressorDecode.(func(recvs SnappyCompressor, argdata []byte) (reta []byte, retb error))(recvs, argdata)
	} else {
		panic("FuncAuxMockSnappyCompressorDecode ")
	}
	AuxMockIncrementRecorderAuxMockSnappyCompressorDecode()
	return
}
