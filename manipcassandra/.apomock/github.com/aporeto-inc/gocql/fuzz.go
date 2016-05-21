// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
// +build gofuzz

package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {

	apomock.RegisterFunc("gocql", "gocql.Fuzz", AuxMockFuzz)
}

const ()

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//

//
// Mock: Fuzz(argdata []byte)(reta int)
//

type MockArgsTypeFuzz struct {
	ApomockCallNumber int
	Argdata           []byte
}

var LastMockArgsFuzz MockArgsTypeFuzz

// AuxMockFuzz(argdata []byte)(reta int) - Generated mock function
func AuxMockFuzz(argdata []byte) (reta int) {
	LastMockArgsFuzz = MockArgsTypeFuzz{
		ApomockCallNumber: AuxMockGetRecorderAuxMockFuzz(),
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.Fuzz")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Fuzz")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Fuzz")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockFuzz  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockFuzz int = 0

var condRecorderAuxMockFuzz *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockFuzz(i int) {
	condRecorderAuxMockFuzz.L.Lock()
	for recorderAuxMockFuzz < i {
		condRecorderAuxMockFuzz.Wait()
	}
	condRecorderAuxMockFuzz.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockFuzz() {
	condRecorderAuxMockFuzz.L.Lock()
	recorderAuxMockFuzz++
	condRecorderAuxMockFuzz.L.Unlock()
	condRecorderAuxMockFuzz.Broadcast()
}
func AuxMockGetRecorderAuxMockFuzz() (ret int) {
	condRecorderAuxMockFuzz.L.Lock()
	ret = recorderAuxMockFuzz
	condRecorderAuxMockFuzz.L.Unlock()
	return
}

// Fuzz - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func Fuzz(argdata []byte) (reta int) {
	FuncAuxMockFuzz, ok := apomock.GetRegisteredFunc("gocql.Fuzz")
	if ok {
		reta = FuncAuxMockFuzz.(func(argdata []byte) (reta int))(argdata)
	} else {
		panic("FuncAuxMockFuzz ")
	}
	AuxMockIncrementRecorderAuxMockFuzz()
	return
}
