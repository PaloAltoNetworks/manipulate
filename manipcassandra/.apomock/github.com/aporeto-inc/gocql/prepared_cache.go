// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"
import "github.com/aporeto-inc/gocql/apointernal/lru"

func init() {
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructPreparedLRU, apomockNewStructPreparedLRU)

	apomock.RegisterFunc("gocql", "gocql.preparedLRU.keyFor", (*preparedLRU).AuxMockkeyFor)
	apomock.RegisterFunc("gocql", "gocql.preparedLRU.max", (*preparedLRU).AuxMockmax)
	apomock.RegisterFunc("gocql", "gocql.preparedLRU.clear", (*preparedLRU).AuxMockclear)
	apomock.RegisterFunc("gocql", "gocql.preparedLRU.add", (*preparedLRU).AuxMockadd)
	apomock.RegisterFunc("gocql", "gocql.preparedLRU.remove", (*preparedLRU).AuxMockremove)
	apomock.RegisterFunc("gocql", "gocql.preparedLRU.execIfMissing", (*preparedLRU).AuxMockexecIfMissing)
}

const defaultMaxPreparedStmts = 1000

const (
	ApomockStructPreparedLRU = 2
)

//
// Internal Types: in this package and their exportable versions
//
type preparedLRU struct {
	mu  sync.Mutex
	lru *lru.Cache
}

//
// External Types: in this package
//

func apomockNewStructPreparedLRU() interface{} { return &preparedLRU{} }

//
// Mock: (recvp *preparedLRU)keyFor(argaddr string, argkeyspace string, argstatement string)(reta string)
//

type MockArgsTypepreparedLRUkeyFor struct {
	ApomockCallNumber int
	Argaddr           string
	Argkeyspace       string
	Argstatement      string
}

var LastMockArgspreparedLRUkeyFor MockArgsTypepreparedLRUkeyFor

// (recvp *preparedLRU)AuxMockkeyFor(argaddr string, argkeyspace string, argstatement string)(reta string) - Generated mock function
func (recvp *preparedLRU) AuxMockkeyFor(argaddr string, argkeyspace string, argstatement string) (reta string) {
	LastMockArgspreparedLRUkeyFor = MockArgsTypepreparedLRUkeyFor{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrpreparedLRUkeyFor(),
		Argaddr:           argaddr,
		Argkeyspace:       argkeyspace,
		Argstatement:      argstatement,
	}
	rargs, rerr := apomock.GetNext("gocql.preparedLRU.keyFor")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.preparedLRU.keyFor")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.preparedLRU.keyFor")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrpreparedLRUkeyFor  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpreparedLRUkeyFor int = 0

var condRecorderAuxMockPtrpreparedLRUkeyFor *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpreparedLRUkeyFor(i int) {
	condRecorderAuxMockPtrpreparedLRUkeyFor.L.Lock()
	for recorderAuxMockPtrpreparedLRUkeyFor < i {
		condRecorderAuxMockPtrpreparedLRUkeyFor.Wait()
	}
	condRecorderAuxMockPtrpreparedLRUkeyFor.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpreparedLRUkeyFor() {
	condRecorderAuxMockPtrpreparedLRUkeyFor.L.Lock()
	recorderAuxMockPtrpreparedLRUkeyFor++
	condRecorderAuxMockPtrpreparedLRUkeyFor.L.Unlock()
	condRecorderAuxMockPtrpreparedLRUkeyFor.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpreparedLRUkeyFor() (ret int) {
	condRecorderAuxMockPtrpreparedLRUkeyFor.L.Lock()
	ret = recorderAuxMockPtrpreparedLRUkeyFor
	condRecorderAuxMockPtrpreparedLRUkeyFor.L.Unlock()
	return
}

// (recvp *preparedLRU)keyFor - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *preparedLRU) keyFor(argaddr string, argkeyspace string, argstatement string) (reta string) {
	FuncAuxMockPtrpreparedLRUkeyFor, ok := apomock.GetRegisteredFunc("gocql.preparedLRU.keyFor")
	if ok {
		reta = FuncAuxMockPtrpreparedLRUkeyFor.(func(recvp *preparedLRU, argaddr string, argkeyspace string, argstatement string) (reta string))(recvp, argaddr, argkeyspace, argstatement)
	} else {
		panic("FuncAuxMockPtrpreparedLRUkeyFor ")
	}
	AuxMockIncrementRecorderAuxMockPtrpreparedLRUkeyFor()
	return
}

//
// Mock: (recvp *preparedLRU)max(argmax int)()
//

type MockArgsTypepreparedLRUmax struct {
	ApomockCallNumber int
	Argmax            int
}

var LastMockArgspreparedLRUmax MockArgsTypepreparedLRUmax

// (recvp *preparedLRU)AuxMockmax(argmax int)() - Generated mock function
func (recvp *preparedLRU) AuxMockmax(argmax int) {
	LastMockArgspreparedLRUmax = MockArgsTypepreparedLRUmax{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrpreparedLRUmax(),
		Argmax:            argmax,
	}
	return
}

// RecorderAuxMockPtrpreparedLRUmax  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpreparedLRUmax int = 0

var condRecorderAuxMockPtrpreparedLRUmax *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpreparedLRUmax(i int) {
	condRecorderAuxMockPtrpreparedLRUmax.L.Lock()
	for recorderAuxMockPtrpreparedLRUmax < i {
		condRecorderAuxMockPtrpreparedLRUmax.Wait()
	}
	condRecorderAuxMockPtrpreparedLRUmax.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpreparedLRUmax() {
	condRecorderAuxMockPtrpreparedLRUmax.L.Lock()
	recorderAuxMockPtrpreparedLRUmax++
	condRecorderAuxMockPtrpreparedLRUmax.L.Unlock()
	condRecorderAuxMockPtrpreparedLRUmax.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpreparedLRUmax() (ret int) {
	condRecorderAuxMockPtrpreparedLRUmax.L.Lock()
	ret = recorderAuxMockPtrpreparedLRUmax
	condRecorderAuxMockPtrpreparedLRUmax.L.Unlock()
	return
}

// (recvp *preparedLRU)max - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *preparedLRU) max(argmax int) {
	FuncAuxMockPtrpreparedLRUmax, ok := apomock.GetRegisteredFunc("gocql.preparedLRU.max")
	if ok {
		FuncAuxMockPtrpreparedLRUmax.(func(recvp *preparedLRU, argmax int))(recvp, argmax)
	} else {
		panic("FuncAuxMockPtrpreparedLRUmax ")
	}
	AuxMockIncrementRecorderAuxMockPtrpreparedLRUmax()
	return
}

//
// Mock: (recvp *preparedLRU)clear()()
//

type MockArgsTypepreparedLRUclear struct {
	ApomockCallNumber int
}

var LastMockArgspreparedLRUclear MockArgsTypepreparedLRUclear

// (recvp *preparedLRU)AuxMockclear()() - Generated mock function
func (recvp *preparedLRU) AuxMockclear() {
	return
}

// RecorderAuxMockPtrpreparedLRUclear  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpreparedLRUclear int = 0

var condRecorderAuxMockPtrpreparedLRUclear *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpreparedLRUclear(i int) {
	condRecorderAuxMockPtrpreparedLRUclear.L.Lock()
	for recorderAuxMockPtrpreparedLRUclear < i {
		condRecorderAuxMockPtrpreparedLRUclear.Wait()
	}
	condRecorderAuxMockPtrpreparedLRUclear.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpreparedLRUclear() {
	condRecorderAuxMockPtrpreparedLRUclear.L.Lock()
	recorderAuxMockPtrpreparedLRUclear++
	condRecorderAuxMockPtrpreparedLRUclear.L.Unlock()
	condRecorderAuxMockPtrpreparedLRUclear.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpreparedLRUclear() (ret int) {
	condRecorderAuxMockPtrpreparedLRUclear.L.Lock()
	ret = recorderAuxMockPtrpreparedLRUclear
	condRecorderAuxMockPtrpreparedLRUclear.L.Unlock()
	return
}

// (recvp *preparedLRU)clear - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *preparedLRU) clear() {
	FuncAuxMockPtrpreparedLRUclear, ok := apomock.GetRegisteredFunc("gocql.preparedLRU.clear")
	if ok {
		FuncAuxMockPtrpreparedLRUclear.(func(recvp *preparedLRU))(recvp)
	} else {
		panic("FuncAuxMockPtrpreparedLRUclear ")
	}
	AuxMockIncrementRecorderAuxMockPtrpreparedLRUclear()
	return
}

//
// Mock: (recvp *preparedLRU)add(argkey string, argval *inflightPrepare)()
//

type MockArgsTypepreparedLRUadd struct {
	ApomockCallNumber int
	Argkey            string
	Argval            *inflightPrepare
}

var LastMockArgspreparedLRUadd MockArgsTypepreparedLRUadd

// (recvp *preparedLRU)AuxMockadd(argkey string, argval *inflightPrepare)() - Generated mock function
func (recvp *preparedLRU) AuxMockadd(argkey string, argval *inflightPrepare) {
	LastMockArgspreparedLRUadd = MockArgsTypepreparedLRUadd{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrpreparedLRUadd(),
		Argkey:            argkey,
		Argval:            argval,
	}
	return
}

// RecorderAuxMockPtrpreparedLRUadd  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpreparedLRUadd int = 0

var condRecorderAuxMockPtrpreparedLRUadd *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpreparedLRUadd(i int) {
	condRecorderAuxMockPtrpreparedLRUadd.L.Lock()
	for recorderAuxMockPtrpreparedLRUadd < i {
		condRecorderAuxMockPtrpreparedLRUadd.Wait()
	}
	condRecorderAuxMockPtrpreparedLRUadd.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpreparedLRUadd() {
	condRecorderAuxMockPtrpreparedLRUadd.L.Lock()
	recorderAuxMockPtrpreparedLRUadd++
	condRecorderAuxMockPtrpreparedLRUadd.L.Unlock()
	condRecorderAuxMockPtrpreparedLRUadd.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpreparedLRUadd() (ret int) {
	condRecorderAuxMockPtrpreparedLRUadd.L.Lock()
	ret = recorderAuxMockPtrpreparedLRUadd
	condRecorderAuxMockPtrpreparedLRUadd.L.Unlock()
	return
}

// (recvp *preparedLRU)add - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *preparedLRU) add(argkey string, argval *inflightPrepare) {
	FuncAuxMockPtrpreparedLRUadd, ok := apomock.GetRegisteredFunc("gocql.preparedLRU.add")
	if ok {
		FuncAuxMockPtrpreparedLRUadd.(func(recvp *preparedLRU, argkey string, argval *inflightPrepare))(recvp, argkey, argval)
	} else {
		panic("FuncAuxMockPtrpreparedLRUadd ")
	}
	AuxMockIncrementRecorderAuxMockPtrpreparedLRUadd()
	return
}

//
// Mock: (recvp *preparedLRU)remove(argkey string)(reta bool)
//

type MockArgsTypepreparedLRUremove struct {
	ApomockCallNumber int
	Argkey            string
}

var LastMockArgspreparedLRUremove MockArgsTypepreparedLRUremove

// (recvp *preparedLRU)AuxMockremove(argkey string)(reta bool) - Generated mock function
func (recvp *preparedLRU) AuxMockremove(argkey string) (reta bool) {
	LastMockArgspreparedLRUremove = MockArgsTypepreparedLRUremove{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrpreparedLRUremove(),
		Argkey:            argkey,
	}
	rargs, rerr := apomock.GetNext("gocql.preparedLRU.remove")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.preparedLRU.remove")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.preparedLRU.remove")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrpreparedLRUremove  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpreparedLRUremove int = 0

var condRecorderAuxMockPtrpreparedLRUremove *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpreparedLRUremove(i int) {
	condRecorderAuxMockPtrpreparedLRUremove.L.Lock()
	for recorderAuxMockPtrpreparedLRUremove < i {
		condRecorderAuxMockPtrpreparedLRUremove.Wait()
	}
	condRecorderAuxMockPtrpreparedLRUremove.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpreparedLRUremove() {
	condRecorderAuxMockPtrpreparedLRUremove.L.Lock()
	recorderAuxMockPtrpreparedLRUremove++
	condRecorderAuxMockPtrpreparedLRUremove.L.Unlock()
	condRecorderAuxMockPtrpreparedLRUremove.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpreparedLRUremove() (ret int) {
	condRecorderAuxMockPtrpreparedLRUremove.L.Lock()
	ret = recorderAuxMockPtrpreparedLRUremove
	condRecorderAuxMockPtrpreparedLRUremove.L.Unlock()
	return
}

// (recvp *preparedLRU)remove - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *preparedLRU) remove(argkey string) (reta bool) {
	FuncAuxMockPtrpreparedLRUremove, ok := apomock.GetRegisteredFunc("gocql.preparedLRU.remove")
	if ok {
		reta = FuncAuxMockPtrpreparedLRUremove.(func(recvp *preparedLRU, argkey string) (reta bool))(recvp, argkey)
	} else {
		panic("FuncAuxMockPtrpreparedLRUremove ")
	}
	AuxMockIncrementRecorderAuxMockPtrpreparedLRUremove()
	return
}

//
// Mock: (recvp *preparedLRU)execIfMissing(argkey string, argfn func(*lru.Cache) *inflightPrepare)(reta *inflightPrepare, retb bool)
//

type MockArgsTypepreparedLRUexecIfMissing struct {
	ApomockCallNumber int
	Argkey            string
	Argfn             func(*lru.Cache) *inflightPrepare
}

var LastMockArgspreparedLRUexecIfMissing MockArgsTypepreparedLRUexecIfMissing

// (recvp *preparedLRU)AuxMockexecIfMissing(argkey string, argfn func(*lru.Cache) *inflightPrepare)(reta *inflightPrepare, retb bool) - Generated mock function
func (recvp *preparedLRU) AuxMockexecIfMissing(argkey string, argfn func(*lru.Cache) *inflightPrepare) (reta *inflightPrepare, retb bool) {
	LastMockArgspreparedLRUexecIfMissing = MockArgsTypepreparedLRUexecIfMissing{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrpreparedLRUexecIfMissing(),
		Argkey:            argkey,
		Argfn:             argfn,
	}
	rargs, rerr := apomock.GetNext("gocql.preparedLRU.execIfMissing")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.preparedLRU.execIfMissing")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.preparedLRU.execIfMissing")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*inflightPrepare)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(bool)
	}
	return
}

// RecorderAuxMockPtrpreparedLRUexecIfMissing  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpreparedLRUexecIfMissing int = 0

var condRecorderAuxMockPtrpreparedLRUexecIfMissing *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpreparedLRUexecIfMissing(i int) {
	condRecorderAuxMockPtrpreparedLRUexecIfMissing.L.Lock()
	for recorderAuxMockPtrpreparedLRUexecIfMissing < i {
		condRecorderAuxMockPtrpreparedLRUexecIfMissing.Wait()
	}
	condRecorderAuxMockPtrpreparedLRUexecIfMissing.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpreparedLRUexecIfMissing() {
	condRecorderAuxMockPtrpreparedLRUexecIfMissing.L.Lock()
	recorderAuxMockPtrpreparedLRUexecIfMissing++
	condRecorderAuxMockPtrpreparedLRUexecIfMissing.L.Unlock()
	condRecorderAuxMockPtrpreparedLRUexecIfMissing.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpreparedLRUexecIfMissing() (ret int) {
	condRecorderAuxMockPtrpreparedLRUexecIfMissing.L.Lock()
	ret = recorderAuxMockPtrpreparedLRUexecIfMissing
	condRecorderAuxMockPtrpreparedLRUexecIfMissing.L.Unlock()
	return
}

// (recvp *preparedLRU)execIfMissing - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *preparedLRU) execIfMissing(argkey string, argfn func(*lru.Cache) *inflightPrepare) (reta *inflightPrepare, retb bool) {
	FuncAuxMockPtrpreparedLRUexecIfMissing, ok := apomock.GetRegisteredFunc("gocql.preparedLRU.execIfMissing")
	if ok {
		reta, retb = FuncAuxMockPtrpreparedLRUexecIfMissing.(func(recvp *preparedLRU, argkey string, argfn func(*lru.Cache) *inflightPrepare) (reta *inflightPrepare, retb bool))(recvp, argkey, argfn)
	} else {
		panic("FuncAuxMockPtrpreparedLRUexecIfMissing ")
	}
	AuxMockIncrementRecorderAuxMockPtrpreparedLRUexecIfMissing()
	return
}
