// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {

	apomock.RegisterFunc("gocql", "gocql.RoundRobin.Size", (*RoundRobin).AuxMockSize)
	apomock.RegisterFunc("gocql", "gocql.RoundRobin.Pick", (*RoundRobin).AuxMockPick)
	apomock.RegisterFunc("gocql", "gocql.RoundRobin.Close", (*RoundRobin).AuxMockClose)
	apomock.RegisterFunc("gocql", "gocql.NewRoundRobin", AuxMockNewRoundRobin)
	apomock.RegisterFunc("gocql", "gocql.RoundRobin.AddNode", (*RoundRobin).AuxMockAddNode)
	apomock.RegisterFunc("gocql", "gocql.RoundRobin.RemoveNode", (*RoundRobin).AuxMockRemoveNode)
}

const ()

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//
type Node interface {
	Pick(qry *Query) *Conn
	Close()
}

type RoundRobin struct {
	pool []Node
	pos  uint32
	mu   sync.RWMutex
}

//
// Mock: (recvr *RoundRobin)Size()(reta int)
//

type MockArgsTypeRoundRobinSize struct {
	ApomockCallNumber int
}

var LastMockArgsRoundRobinSize MockArgsTypeRoundRobinSize

// (recvr *RoundRobin)AuxMockSize()(reta int) - Generated mock function
func (recvr *RoundRobin) AuxMockSize() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.RoundRobin.Size")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.RoundRobin.Size")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.RoundRobin.Size")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrRoundRobinSize  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrRoundRobinSize int = 0

var condRecorderAuxMockPtrRoundRobinSize *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrRoundRobinSize(i int) {
	condRecorderAuxMockPtrRoundRobinSize.L.Lock()
	for recorderAuxMockPtrRoundRobinSize < i {
		condRecorderAuxMockPtrRoundRobinSize.Wait()
	}
	condRecorderAuxMockPtrRoundRobinSize.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrRoundRobinSize() {
	condRecorderAuxMockPtrRoundRobinSize.L.Lock()
	recorderAuxMockPtrRoundRobinSize++
	condRecorderAuxMockPtrRoundRobinSize.L.Unlock()
	condRecorderAuxMockPtrRoundRobinSize.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrRoundRobinSize() (ret int) {
	condRecorderAuxMockPtrRoundRobinSize.L.Lock()
	ret = recorderAuxMockPtrRoundRobinSize
	condRecorderAuxMockPtrRoundRobinSize.L.Unlock()
	return
}

// (recvr *RoundRobin)Size - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *RoundRobin) Size() (reta int) {
	FuncAuxMockPtrRoundRobinSize, ok := apomock.GetRegisteredFunc("gocql.RoundRobin.Size")
	if ok {
		reta = FuncAuxMockPtrRoundRobinSize.(func(recvr *RoundRobin) (reta int))(recvr)
	} else {
		panic("FuncAuxMockPtrRoundRobinSize ")
	}
	AuxMockIncrementRecorderAuxMockPtrRoundRobinSize()
	return
}

//
// Mock: (recvr *RoundRobin)Pick(argqry *Query)(reta *Conn)
//

type MockArgsTypeRoundRobinPick struct {
	ApomockCallNumber int
	Argqry            *Query
}

var LastMockArgsRoundRobinPick MockArgsTypeRoundRobinPick

// (recvr *RoundRobin)AuxMockPick(argqry *Query)(reta *Conn) - Generated mock function
func (recvr *RoundRobin) AuxMockPick(argqry *Query) (reta *Conn) {
	LastMockArgsRoundRobinPick = MockArgsTypeRoundRobinPick{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrRoundRobinPick(),
		Argqry:            argqry,
	}
	rargs, rerr := apomock.GetNext("gocql.RoundRobin.Pick")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.RoundRobin.Pick")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.RoundRobin.Pick")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Conn)
	}
	return
}

// RecorderAuxMockPtrRoundRobinPick  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrRoundRobinPick int = 0

var condRecorderAuxMockPtrRoundRobinPick *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrRoundRobinPick(i int) {
	condRecorderAuxMockPtrRoundRobinPick.L.Lock()
	for recorderAuxMockPtrRoundRobinPick < i {
		condRecorderAuxMockPtrRoundRobinPick.Wait()
	}
	condRecorderAuxMockPtrRoundRobinPick.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrRoundRobinPick() {
	condRecorderAuxMockPtrRoundRobinPick.L.Lock()
	recorderAuxMockPtrRoundRobinPick++
	condRecorderAuxMockPtrRoundRobinPick.L.Unlock()
	condRecorderAuxMockPtrRoundRobinPick.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrRoundRobinPick() (ret int) {
	condRecorderAuxMockPtrRoundRobinPick.L.Lock()
	ret = recorderAuxMockPtrRoundRobinPick
	condRecorderAuxMockPtrRoundRobinPick.L.Unlock()
	return
}

// (recvr *RoundRobin)Pick - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *RoundRobin) Pick(argqry *Query) (reta *Conn) {
	FuncAuxMockPtrRoundRobinPick, ok := apomock.GetRegisteredFunc("gocql.RoundRobin.Pick")
	if ok {
		reta = FuncAuxMockPtrRoundRobinPick.(func(recvr *RoundRobin, argqry *Query) (reta *Conn))(recvr, argqry)
	} else {
		panic("FuncAuxMockPtrRoundRobinPick ")
	}
	AuxMockIncrementRecorderAuxMockPtrRoundRobinPick()
	return
}

//
// Mock: (recvr *RoundRobin)Close()()
//

type MockArgsTypeRoundRobinClose struct {
	ApomockCallNumber int
}

var LastMockArgsRoundRobinClose MockArgsTypeRoundRobinClose

// (recvr *RoundRobin)AuxMockClose()() - Generated mock function
func (recvr *RoundRobin) AuxMockClose() {
	return
}

// RecorderAuxMockPtrRoundRobinClose  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrRoundRobinClose int = 0

var condRecorderAuxMockPtrRoundRobinClose *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrRoundRobinClose(i int) {
	condRecorderAuxMockPtrRoundRobinClose.L.Lock()
	for recorderAuxMockPtrRoundRobinClose < i {
		condRecorderAuxMockPtrRoundRobinClose.Wait()
	}
	condRecorderAuxMockPtrRoundRobinClose.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrRoundRobinClose() {
	condRecorderAuxMockPtrRoundRobinClose.L.Lock()
	recorderAuxMockPtrRoundRobinClose++
	condRecorderAuxMockPtrRoundRobinClose.L.Unlock()
	condRecorderAuxMockPtrRoundRobinClose.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrRoundRobinClose() (ret int) {
	condRecorderAuxMockPtrRoundRobinClose.L.Lock()
	ret = recorderAuxMockPtrRoundRobinClose
	condRecorderAuxMockPtrRoundRobinClose.L.Unlock()
	return
}

// (recvr *RoundRobin)Close - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *RoundRobin) Close() {
	FuncAuxMockPtrRoundRobinClose, ok := apomock.GetRegisteredFunc("gocql.RoundRobin.Close")
	if ok {
		FuncAuxMockPtrRoundRobinClose.(func(recvr *RoundRobin))(recvr)
	} else {
		panic("FuncAuxMockPtrRoundRobinClose ")
	}
	AuxMockIncrementRecorderAuxMockPtrRoundRobinClose()
	return
}

//
// Mock: NewRoundRobin()(reta *RoundRobin)
//

type MockArgsTypeNewRoundRobin struct {
	ApomockCallNumber int
}

var LastMockArgsNewRoundRobin MockArgsTypeNewRoundRobin

// AuxMockNewRoundRobin()(reta *RoundRobin) - Generated mock function
func AuxMockNewRoundRobin() (reta *RoundRobin) {
	rargs, rerr := apomock.GetNext("gocql.NewRoundRobin")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.NewRoundRobin")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.NewRoundRobin")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*RoundRobin)
	}
	return
}

// RecorderAuxMockNewRoundRobin  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNewRoundRobin int = 0

var condRecorderAuxMockNewRoundRobin *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNewRoundRobin(i int) {
	condRecorderAuxMockNewRoundRobin.L.Lock()
	for recorderAuxMockNewRoundRobin < i {
		condRecorderAuxMockNewRoundRobin.Wait()
	}
	condRecorderAuxMockNewRoundRobin.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNewRoundRobin() {
	condRecorderAuxMockNewRoundRobin.L.Lock()
	recorderAuxMockNewRoundRobin++
	condRecorderAuxMockNewRoundRobin.L.Unlock()
	condRecorderAuxMockNewRoundRobin.Broadcast()
}
func AuxMockGetRecorderAuxMockNewRoundRobin() (ret int) {
	condRecorderAuxMockNewRoundRobin.L.Lock()
	ret = recorderAuxMockNewRoundRobin
	condRecorderAuxMockNewRoundRobin.L.Unlock()
	return
}

// NewRoundRobin - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func NewRoundRobin() (reta *RoundRobin) {
	FuncAuxMockNewRoundRobin, ok := apomock.GetRegisteredFunc("gocql.NewRoundRobin")
	if ok {
		reta = FuncAuxMockNewRoundRobin.(func() (reta *RoundRobin))()
	} else {
		panic("FuncAuxMockNewRoundRobin ")
	}
	AuxMockIncrementRecorderAuxMockNewRoundRobin()
	return
}

//
// Mock: (recvr *RoundRobin)AddNode(argnode Node)()
//

type MockArgsTypeRoundRobinAddNode struct {
	ApomockCallNumber int
	Argnode           Node
}

var LastMockArgsRoundRobinAddNode MockArgsTypeRoundRobinAddNode

// (recvr *RoundRobin)AuxMockAddNode(argnode Node)() - Generated mock function
func (recvr *RoundRobin) AuxMockAddNode(argnode Node) {
	LastMockArgsRoundRobinAddNode = MockArgsTypeRoundRobinAddNode{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrRoundRobinAddNode(),
		Argnode:           argnode,
	}
	return
}

// RecorderAuxMockPtrRoundRobinAddNode  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrRoundRobinAddNode int = 0

var condRecorderAuxMockPtrRoundRobinAddNode *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrRoundRobinAddNode(i int) {
	condRecorderAuxMockPtrRoundRobinAddNode.L.Lock()
	for recorderAuxMockPtrRoundRobinAddNode < i {
		condRecorderAuxMockPtrRoundRobinAddNode.Wait()
	}
	condRecorderAuxMockPtrRoundRobinAddNode.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrRoundRobinAddNode() {
	condRecorderAuxMockPtrRoundRobinAddNode.L.Lock()
	recorderAuxMockPtrRoundRobinAddNode++
	condRecorderAuxMockPtrRoundRobinAddNode.L.Unlock()
	condRecorderAuxMockPtrRoundRobinAddNode.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrRoundRobinAddNode() (ret int) {
	condRecorderAuxMockPtrRoundRobinAddNode.L.Lock()
	ret = recorderAuxMockPtrRoundRobinAddNode
	condRecorderAuxMockPtrRoundRobinAddNode.L.Unlock()
	return
}

// (recvr *RoundRobin)AddNode - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *RoundRobin) AddNode(argnode Node) {
	FuncAuxMockPtrRoundRobinAddNode, ok := apomock.GetRegisteredFunc("gocql.RoundRobin.AddNode")
	if ok {
		FuncAuxMockPtrRoundRobinAddNode.(func(recvr *RoundRobin, argnode Node))(recvr, argnode)
	} else {
		panic("FuncAuxMockPtrRoundRobinAddNode ")
	}
	AuxMockIncrementRecorderAuxMockPtrRoundRobinAddNode()
	return
}

//
// Mock: (recvr *RoundRobin)RemoveNode(argnode Node)()
//

type MockArgsTypeRoundRobinRemoveNode struct {
	ApomockCallNumber int
	Argnode           Node
}

var LastMockArgsRoundRobinRemoveNode MockArgsTypeRoundRobinRemoveNode

// (recvr *RoundRobin)AuxMockRemoveNode(argnode Node)() - Generated mock function
func (recvr *RoundRobin) AuxMockRemoveNode(argnode Node) {
	LastMockArgsRoundRobinRemoveNode = MockArgsTypeRoundRobinRemoveNode{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrRoundRobinRemoveNode(),
		Argnode:           argnode,
	}
	return
}

// RecorderAuxMockPtrRoundRobinRemoveNode  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrRoundRobinRemoveNode int = 0

var condRecorderAuxMockPtrRoundRobinRemoveNode *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrRoundRobinRemoveNode(i int) {
	condRecorderAuxMockPtrRoundRobinRemoveNode.L.Lock()
	for recorderAuxMockPtrRoundRobinRemoveNode < i {
		condRecorderAuxMockPtrRoundRobinRemoveNode.Wait()
	}
	condRecorderAuxMockPtrRoundRobinRemoveNode.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrRoundRobinRemoveNode() {
	condRecorderAuxMockPtrRoundRobinRemoveNode.L.Lock()
	recorderAuxMockPtrRoundRobinRemoveNode++
	condRecorderAuxMockPtrRoundRobinRemoveNode.L.Unlock()
	condRecorderAuxMockPtrRoundRobinRemoveNode.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrRoundRobinRemoveNode() (ret int) {
	condRecorderAuxMockPtrRoundRobinRemoveNode.L.Lock()
	ret = recorderAuxMockPtrRoundRobinRemoveNode
	condRecorderAuxMockPtrRoundRobinRemoveNode.L.Unlock()
	return
}

// (recvr *RoundRobin)RemoveNode - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *RoundRobin) RemoveNode(argnode Node) {
	FuncAuxMockPtrRoundRobinRemoveNode, ok := apomock.GetRegisteredFunc("gocql.RoundRobin.RemoveNode")
	if ok {
		FuncAuxMockPtrRoundRobinRemoveNode.(func(recvr *RoundRobin, argnode Node))(recvr, argnode)
	} else {
		panic("FuncAuxMockPtrRoundRobinRemoveNode ")
	}
	AuxMockIncrementRecorderAuxMockPtrRoundRobinRemoveNode()
	return
}
