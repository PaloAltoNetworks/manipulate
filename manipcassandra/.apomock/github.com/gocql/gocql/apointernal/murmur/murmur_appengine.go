// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
// +build appengine

package murmur

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {

	apomock.RegisterFunc("murmur", "murmur.Murmur3H1", AuxMockMurmur3H1)
}

const ()

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//

//
// Mock: Murmur3H1(argdata []byte)(reta uint64)
//

type MockArgsTypeMurmur3H1 struct {
	ApomockCallNumber int
	Argdata           []byte
}

var LastMockArgsMurmur3H1 MockArgsTypeMurmur3H1

// AuxMockMurmur3H1(argdata []byte)(reta uint64) - Generated mock function
func AuxMockMurmur3H1(argdata []byte) (reta uint64) {
	LastMockArgsMurmur3H1 = MockArgsTypeMurmur3H1{
		ApomockCallNumber: AuxMockGetRecorderAuxMockMurmur3H1(),
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("murmur.Murmur3H1")
	if rerr != nil {
		panic("Error getting next entry for method: murmur.Murmur3H1")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:murmur.Murmur3H1")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(uint64)
	}
	return
}

// RecorderAuxMockMurmur3H1  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockMurmur3H1 int = 0

var condRecorderAuxMockMurmur3H1 *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockMurmur3H1(i int) {
	condRecorderAuxMockMurmur3H1.L.Lock()
	for recorderAuxMockMurmur3H1 < i {
		condRecorderAuxMockMurmur3H1.Wait()
	}
	condRecorderAuxMockMurmur3H1.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockMurmur3H1() {
	condRecorderAuxMockMurmur3H1.L.Lock()
	recorderAuxMockMurmur3H1++
	condRecorderAuxMockMurmur3H1.L.Unlock()
	condRecorderAuxMockMurmur3H1.Broadcast()
}
func AuxMockGetRecorderAuxMockMurmur3H1() (ret int) {
	condRecorderAuxMockMurmur3H1.L.Lock()
	ret = recorderAuxMockMurmur3H1
	condRecorderAuxMockMurmur3H1.L.Unlock()
	return
}

// Murmur3H1 - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func Murmur3H1(argdata []byte) (reta uint64) {
	FuncAuxMockMurmur3H1, ok := apomock.GetRegisteredFunc("murmur.Murmur3H1")
	if ok {
		reta = FuncAuxMockMurmur3H1.(func(argdata []byte) (reta uint64))(argdata)
	} else {
		panic("FuncAuxMockMurmur3H1 ")
	}
	AuxMockIncrementRecorderAuxMockMurmur3H1()
	return
}
