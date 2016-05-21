// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package cassandra

import "reflect"
import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {

	apomock.RegisterFunc("cassandra", "cassandra.Unmarshal", AuxMockUnmarshal)
	apomock.RegisterFunc("cassandra", "cassandra.unmarshal", AuxMockunmarshal)
}

const ()

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//

//
// Mock: Unmarshal(argdata interface{}, argv interface{})(reta error)
//

type MockArgsTypeUnmarshal struct {
	ApomockCallNumber int
	Argdata           interface{}
	Argv              interface{}
}

var LastMockArgsUnmarshal MockArgsTypeUnmarshal

// AuxMockUnmarshal(argdata interface{}, argv interface{})(reta error) - Generated mock function
func AuxMockUnmarshal(argdata interface{}, argv interface{}) (reta error) {
	LastMockArgsUnmarshal = MockArgsTypeUnmarshal{
		ApomockCallNumber: AuxMockGetRecorderAuxMockUnmarshal(),
		Argdata:           argdata,
		Argv:              argv,
	}
	rargs, rerr := apomock.GetNext("cassandra.Unmarshal")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.Unmarshal")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.Unmarshal")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockUnmarshal  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUnmarshal int = 0

var condRecorderAuxMockUnmarshal *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUnmarshal(i int) {
	condRecorderAuxMockUnmarshal.L.Lock()
	for recorderAuxMockUnmarshal < i {
		condRecorderAuxMockUnmarshal.Wait()
	}
	condRecorderAuxMockUnmarshal.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUnmarshal() {
	condRecorderAuxMockUnmarshal.L.Lock()
	recorderAuxMockUnmarshal++
	condRecorderAuxMockUnmarshal.L.Unlock()
	condRecorderAuxMockUnmarshal.Broadcast()
}
func AuxMockGetRecorderAuxMockUnmarshal() (ret int) {
	condRecorderAuxMockUnmarshal.L.Lock()
	ret = recorderAuxMockUnmarshal
	condRecorderAuxMockUnmarshal.L.Unlock()
	return
}

// Unmarshal - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func Unmarshal(argdata interface{}, argv interface{}) (reta error) {
	FuncAuxMockUnmarshal, ok := apomock.GetRegisteredFunc("cassandra.Unmarshal")
	if ok {
		reta = FuncAuxMockUnmarshal.(func(argdata interface{}, argv interface{}) (reta error))(argdata, argv)
	} else {
		panic("FuncAuxMockUnmarshal ")
	}
	AuxMockIncrementRecorderAuxMockUnmarshal()
	return
}

//
// Mock: unmarshal(argval reflect.Value, argdata map[string]interface{}, argfieldsMap map[string]field)()
//

type MockArgsTypeunmarshal struct {
	ApomockCallNumber int
	Argval            reflect.Value
	Argdata           map[string]interface{}
	ArgfieldsMap      map[string]field
}

var LastMockArgsunmarshal MockArgsTypeunmarshal

// AuxMockunmarshal(argval reflect.Value, argdata map[string]interface{}, argfieldsMap map[string]field)() - Generated mock function
func AuxMockunmarshal(argval reflect.Value, argdata map[string]interface{}, argfieldsMap map[string]field) {
	LastMockArgsunmarshal = MockArgsTypeunmarshal{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshal(),
		Argval:            argval,
		Argdata:           argdata,
		ArgfieldsMap:      argfieldsMap,
	}
	return
}

// RecorderAuxMockunmarshal  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshal int = 0

var condRecorderAuxMockunmarshal *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshal(i int) {
	condRecorderAuxMockunmarshal.L.Lock()
	for recorderAuxMockunmarshal < i {
		condRecorderAuxMockunmarshal.Wait()
	}
	condRecorderAuxMockunmarshal.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshal() {
	condRecorderAuxMockunmarshal.L.Lock()
	recorderAuxMockunmarshal++
	condRecorderAuxMockunmarshal.L.Unlock()
	condRecorderAuxMockunmarshal.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshal() (ret int) {
	condRecorderAuxMockunmarshal.L.Lock()
	ret = recorderAuxMockunmarshal
	condRecorderAuxMockunmarshal.L.Unlock()
	return
}

// unmarshal - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshal(argval reflect.Value, argdata map[string]interface{}, argfieldsMap map[string]field) {
	FuncAuxMockunmarshal, ok := apomock.GetRegisteredFunc("cassandra.unmarshal")
	if ok {
		FuncAuxMockunmarshal.(func(argval reflect.Value, argdata map[string]interface{}, argfieldsMap map[string]field))(argval, argdata, argfieldsMap)
	} else {
		panic("FuncAuxMockunmarshal ")
	}
	AuxMockIncrementRecorderAuxMockunmarshal()
	return
}
