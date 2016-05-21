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
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructRingDescriber, apomockNewStructRingDescriber)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructCassVersion, apomockNewStructCassVersion)

	apomock.RegisterFunc("gocql", "gocql.cassVersion.Set", (*cassVersion).AuxMockSet)
	apomock.RegisterFunc("gocql", "gocql.cassVersion.UnmarshalCQL", (*cassVersion).AuxMockUnmarshalCQL)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.setHostID", (*HostInfo).AuxMocksetHostID)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.setVersion", (*HostInfo).AuxMocksetVersion)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.IsUp", (*HostInfo).AuxMockIsUp)
	apomock.RegisterFunc("gocql", "gocql.nodeState.String", (nodeState).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.Peer", (*HostInfo).AuxMockPeer)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.Rack", (*HostInfo).AuxMockRack)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.update", (*HostInfo).AuxMockupdate)
	apomock.RegisterFunc("gocql", "gocql.checkSystemLocal", AuxMockcheckSystemLocal)
	apomock.RegisterFunc("gocql", "gocql.ringDescriber.GetHosts", (*ringDescriber).AuxMockGetHosts)
	apomock.RegisterFunc("gocql", "gocql.ringDescriber.refreshRing", (*ringDescriber).AuxMockrefreshRing)
	apomock.RegisterFunc("gocql", "gocql.cassVersion.nodeUpDelay", (cassVersion).AuxMocknodeUpDelay)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.Equal", (*HostInfo).AuxMockEqual)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.setRack", (*HostInfo).AuxMocksetRack)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.Port", (*HostInfo).AuxMockPort)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.setPort", (*HostInfo).AuxMocksetPort)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.setPeer", (*HostInfo).AuxMocksetPeer)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.DataCenter", (*HostInfo).AuxMockDataCenter)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.HostID", (*HostInfo).AuxMockHostID)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.Version", (*HostInfo).AuxMockVersion)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.String", (*HostInfo).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.ringDescriber.matchFilter", (*ringDescriber).AuxMockmatchFilter)
	apomock.RegisterFunc("gocql", "gocql.cassVersion.Before", (cassVersion).AuxMockBefore)
	apomock.RegisterFunc("gocql", "gocql.cassVersion.String", (cassVersion).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.setDataCenter", (*HostInfo).AuxMocksetDataCenter)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.setState", (*HostInfo).AuxMocksetState)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.Tokens", (*HostInfo).AuxMockTokens)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.setTokens", (*HostInfo).AuxMocksetTokens)
	apomock.RegisterFunc("gocql", "gocql.cassVersion.unmarshal", (*cassVersion).AuxMockunmarshal)
	apomock.RegisterFunc("gocql", "gocql.HostInfo.State", (*HostInfo).AuxMockState)
}

const (
	NodeUp nodeState = iota
	NodeDown
)

const (
	ApomockStructRingDescriber = 58
	ApomockStructCassVersion   = 59
)

//
// Internal Types: in this package and their exportable versions
//
type ringDescriber struct {
	dcFilter        string
	rackFilter      string
	session         *Session
	closeChan       chan bool
	localHasRpcAddr bool
	mu              sync.Mutex
	prevHosts       []*HostInfo
	prevPartitioner string
}
type nodeState int32
type cassVersion struct{ Major, Minor, Patch int }

//
// External Types: in this package
//
type HostInfo struct {
	mu         sync.RWMutex
	peer       string
	port       int
	dataCenter string
	rack       string
	hostId     string
	version    cassVersion
	state      nodeState
	tokens     []string
}

func apomockNewStructRingDescriber() interface{} { return &ringDescriber{} }
func apomockNewStructCassVersion() interface{}   { return &cassVersion{} }

//
// Mock: (recvc *cassVersion)Set(argv string)(reta error)
//

type MockArgsTypecassVersionSet struct {
	ApomockCallNumber int
	Argv              string
}

var LastMockArgscassVersionSet MockArgsTypecassVersionSet

// (recvc *cassVersion)AuxMockSet(argv string)(reta error) - Generated mock function
func (recvc *cassVersion) AuxMockSet(argv string) (reta error) {
	LastMockArgscassVersionSet = MockArgsTypecassVersionSet{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcassVersionSet(),
		Argv:              argv,
	}
	rargs, rerr := apomock.GetNext("gocql.cassVersion.Set")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.cassVersion.Set")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.cassVersion.Set")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrcassVersionSet  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcassVersionSet int = 0

var condRecorderAuxMockPtrcassVersionSet *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcassVersionSet(i int) {
	condRecorderAuxMockPtrcassVersionSet.L.Lock()
	for recorderAuxMockPtrcassVersionSet < i {
		condRecorderAuxMockPtrcassVersionSet.Wait()
	}
	condRecorderAuxMockPtrcassVersionSet.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcassVersionSet() {
	condRecorderAuxMockPtrcassVersionSet.L.Lock()
	recorderAuxMockPtrcassVersionSet++
	condRecorderAuxMockPtrcassVersionSet.L.Unlock()
	condRecorderAuxMockPtrcassVersionSet.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcassVersionSet() (ret int) {
	condRecorderAuxMockPtrcassVersionSet.L.Lock()
	ret = recorderAuxMockPtrcassVersionSet
	condRecorderAuxMockPtrcassVersionSet.L.Unlock()
	return
}

// (recvc *cassVersion)Set - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *cassVersion) Set(argv string) (reta error) {
	FuncAuxMockPtrcassVersionSet, ok := apomock.GetRegisteredFunc("gocql.cassVersion.Set")
	if ok {
		reta = FuncAuxMockPtrcassVersionSet.(func(recvc *cassVersion, argv string) (reta error))(recvc, argv)
	} else {
		panic("FuncAuxMockPtrcassVersionSet ")
	}
	AuxMockIncrementRecorderAuxMockPtrcassVersionSet()
	return
}

//
// Mock: (recvc *cassVersion)UnmarshalCQL(arginfo TypeInfo, argdata []byte)(reta error)
//

type MockArgsTypecassVersionUnmarshalCQL struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
}

var LastMockArgscassVersionUnmarshalCQL MockArgsTypecassVersionUnmarshalCQL

// (recvc *cassVersion)AuxMockUnmarshalCQL(arginfo TypeInfo, argdata []byte)(reta error) - Generated mock function
func (recvc *cassVersion) AuxMockUnmarshalCQL(arginfo TypeInfo, argdata []byte) (reta error) {
	LastMockArgscassVersionUnmarshalCQL = MockArgsTypecassVersionUnmarshalCQL{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcassVersionUnmarshalCQL(),
		Arginfo:           arginfo,
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.cassVersion.UnmarshalCQL")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.cassVersion.UnmarshalCQL")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.cassVersion.UnmarshalCQL")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrcassVersionUnmarshalCQL  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcassVersionUnmarshalCQL int = 0

var condRecorderAuxMockPtrcassVersionUnmarshalCQL *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcassVersionUnmarshalCQL(i int) {
	condRecorderAuxMockPtrcassVersionUnmarshalCQL.L.Lock()
	for recorderAuxMockPtrcassVersionUnmarshalCQL < i {
		condRecorderAuxMockPtrcassVersionUnmarshalCQL.Wait()
	}
	condRecorderAuxMockPtrcassVersionUnmarshalCQL.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcassVersionUnmarshalCQL() {
	condRecorderAuxMockPtrcassVersionUnmarshalCQL.L.Lock()
	recorderAuxMockPtrcassVersionUnmarshalCQL++
	condRecorderAuxMockPtrcassVersionUnmarshalCQL.L.Unlock()
	condRecorderAuxMockPtrcassVersionUnmarshalCQL.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcassVersionUnmarshalCQL() (ret int) {
	condRecorderAuxMockPtrcassVersionUnmarshalCQL.L.Lock()
	ret = recorderAuxMockPtrcassVersionUnmarshalCQL
	condRecorderAuxMockPtrcassVersionUnmarshalCQL.L.Unlock()
	return
}

// (recvc *cassVersion)UnmarshalCQL - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *cassVersion) UnmarshalCQL(arginfo TypeInfo, argdata []byte) (reta error) {
	FuncAuxMockPtrcassVersionUnmarshalCQL, ok := apomock.GetRegisteredFunc("gocql.cassVersion.UnmarshalCQL")
	if ok {
		reta = FuncAuxMockPtrcassVersionUnmarshalCQL.(func(recvc *cassVersion, arginfo TypeInfo, argdata []byte) (reta error))(recvc, arginfo, argdata)
	} else {
		panic("FuncAuxMockPtrcassVersionUnmarshalCQL ")
	}
	AuxMockIncrementRecorderAuxMockPtrcassVersionUnmarshalCQL()
	return
}

//
// Mock: (recvh *HostInfo)setHostID(arghostID string)(reta *HostInfo)
//

type MockArgsTypeHostInfosetHostID struct {
	ApomockCallNumber int
	ArghostID         string
}

var LastMockArgsHostInfosetHostID MockArgsTypeHostInfosetHostID

// (recvh *HostInfo)AuxMocksetHostID(arghostID string)(reta *HostInfo) - Generated mock function
func (recvh *HostInfo) AuxMocksetHostID(arghostID string) (reta *HostInfo) {
	LastMockArgsHostInfosetHostID = MockArgsTypeHostInfosetHostID{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrHostInfosetHostID(),
		ArghostID:         arghostID,
	}
	rargs, rerr := apomock.GetNext("gocql.HostInfo.setHostID")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.setHostID")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.setHostID")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrHostInfosetHostID  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfosetHostID int = 0

var condRecorderAuxMockPtrHostInfosetHostID *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfosetHostID(i int) {
	condRecorderAuxMockPtrHostInfosetHostID.L.Lock()
	for recorderAuxMockPtrHostInfosetHostID < i {
		condRecorderAuxMockPtrHostInfosetHostID.Wait()
	}
	condRecorderAuxMockPtrHostInfosetHostID.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfosetHostID() {
	condRecorderAuxMockPtrHostInfosetHostID.L.Lock()
	recorderAuxMockPtrHostInfosetHostID++
	condRecorderAuxMockPtrHostInfosetHostID.L.Unlock()
	condRecorderAuxMockPtrHostInfosetHostID.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfosetHostID() (ret int) {
	condRecorderAuxMockPtrHostInfosetHostID.L.Lock()
	ret = recorderAuxMockPtrHostInfosetHostID
	condRecorderAuxMockPtrHostInfosetHostID.L.Unlock()
	return
}

// (recvh *HostInfo)setHostID - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) setHostID(arghostID string) (reta *HostInfo) {
	FuncAuxMockPtrHostInfosetHostID, ok := apomock.GetRegisteredFunc("gocql.HostInfo.setHostID")
	if ok {
		reta = FuncAuxMockPtrHostInfosetHostID.(func(recvh *HostInfo, arghostID string) (reta *HostInfo))(recvh, arghostID)
	} else {
		panic("FuncAuxMockPtrHostInfosetHostID ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfosetHostID()
	return
}

//
// Mock: (recvh *HostInfo)setVersion(argmajor int, argminor int, argpatch int)(reta *HostInfo)
//

type MockArgsTypeHostInfosetVersion struct {
	ApomockCallNumber int
	Argmajor          int
	Argminor          int
	Argpatch          int
}

var LastMockArgsHostInfosetVersion MockArgsTypeHostInfosetVersion

// (recvh *HostInfo)AuxMocksetVersion(argmajor int, argminor int, argpatch int)(reta *HostInfo) - Generated mock function
func (recvh *HostInfo) AuxMocksetVersion(argmajor int, argminor int, argpatch int) (reta *HostInfo) {
	LastMockArgsHostInfosetVersion = MockArgsTypeHostInfosetVersion{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrHostInfosetVersion(),
		Argmajor:          argmajor,
		Argminor:          argminor,
		Argpatch:          argpatch,
	}
	rargs, rerr := apomock.GetNext("gocql.HostInfo.setVersion")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.setVersion")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.setVersion")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrHostInfosetVersion  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfosetVersion int = 0

var condRecorderAuxMockPtrHostInfosetVersion *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfosetVersion(i int) {
	condRecorderAuxMockPtrHostInfosetVersion.L.Lock()
	for recorderAuxMockPtrHostInfosetVersion < i {
		condRecorderAuxMockPtrHostInfosetVersion.Wait()
	}
	condRecorderAuxMockPtrHostInfosetVersion.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfosetVersion() {
	condRecorderAuxMockPtrHostInfosetVersion.L.Lock()
	recorderAuxMockPtrHostInfosetVersion++
	condRecorderAuxMockPtrHostInfosetVersion.L.Unlock()
	condRecorderAuxMockPtrHostInfosetVersion.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfosetVersion() (ret int) {
	condRecorderAuxMockPtrHostInfosetVersion.L.Lock()
	ret = recorderAuxMockPtrHostInfosetVersion
	condRecorderAuxMockPtrHostInfosetVersion.L.Unlock()
	return
}

// (recvh *HostInfo)setVersion - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) setVersion(argmajor int, argminor int, argpatch int) (reta *HostInfo) {
	FuncAuxMockPtrHostInfosetVersion, ok := apomock.GetRegisteredFunc("gocql.HostInfo.setVersion")
	if ok {
		reta = FuncAuxMockPtrHostInfosetVersion.(func(recvh *HostInfo, argmajor int, argminor int, argpatch int) (reta *HostInfo))(recvh, argmajor, argminor, argpatch)
	} else {
		panic("FuncAuxMockPtrHostInfosetVersion ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfosetVersion()
	return
}

//
// Mock: (recvh *HostInfo)IsUp()(reta bool)
//

type MockArgsTypeHostInfoIsUp struct {
	ApomockCallNumber int
}

var LastMockArgsHostInfoIsUp MockArgsTypeHostInfoIsUp

// (recvh *HostInfo)AuxMockIsUp()(reta bool) - Generated mock function
func (recvh *HostInfo) AuxMockIsUp() (reta bool) {
	rargs, rerr := apomock.GetNext("gocql.HostInfo.IsUp")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.IsUp")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.IsUp")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrHostInfoIsUp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoIsUp int = 0

var condRecorderAuxMockPtrHostInfoIsUp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoIsUp(i int) {
	condRecorderAuxMockPtrHostInfoIsUp.L.Lock()
	for recorderAuxMockPtrHostInfoIsUp < i {
		condRecorderAuxMockPtrHostInfoIsUp.Wait()
	}
	condRecorderAuxMockPtrHostInfoIsUp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoIsUp() {
	condRecorderAuxMockPtrHostInfoIsUp.L.Lock()
	recorderAuxMockPtrHostInfoIsUp++
	condRecorderAuxMockPtrHostInfoIsUp.L.Unlock()
	condRecorderAuxMockPtrHostInfoIsUp.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoIsUp() (ret int) {
	condRecorderAuxMockPtrHostInfoIsUp.L.Lock()
	ret = recorderAuxMockPtrHostInfoIsUp
	condRecorderAuxMockPtrHostInfoIsUp.L.Unlock()
	return
}

// (recvh *HostInfo)IsUp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) IsUp() (reta bool) {
	FuncAuxMockPtrHostInfoIsUp, ok := apomock.GetRegisteredFunc("gocql.HostInfo.IsUp")
	if ok {
		reta = FuncAuxMockPtrHostInfoIsUp.(func(recvh *HostInfo) (reta bool))(recvh)
	} else {
		panic("FuncAuxMockPtrHostInfoIsUp ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoIsUp()
	return
}

//
// Mock: (recvn nodeState)String()(reta string)
//

type MockArgsTypenodeStateString struct {
	ApomockCallNumber int
}

var LastMockArgsnodeStateString MockArgsTypenodeStateString

// (recvn nodeState)AuxMockString()(reta string) - Generated mock function
func (recvn nodeState) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.nodeState.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.nodeState.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.nodeState.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMocknodeStateString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocknodeStateString int = 0

var condRecorderAuxMocknodeStateString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocknodeStateString(i int) {
	condRecorderAuxMocknodeStateString.L.Lock()
	for recorderAuxMocknodeStateString < i {
		condRecorderAuxMocknodeStateString.Wait()
	}
	condRecorderAuxMocknodeStateString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocknodeStateString() {
	condRecorderAuxMocknodeStateString.L.Lock()
	recorderAuxMocknodeStateString++
	condRecorderAuxMocknodeStateString.L.Unlock()
	condRecorderAuxMocknodeStateString.Broadcast()
}
func AuxMockGetRecorderAuxMocknodeStateString() (ret int) {
	condRecorderAuxMocknodeStateString.L.Lock()
	ret = recorderAuxMocknodeStateString
	condRecorderAuxMocknodeStateString.L.Unlock()
	return
}

// (recvn nodeState)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvn nodeState) String() (reta string) {
	FuncAuxMocknodeStateString, ok := apomock.GetRegisteredFunc("gocql.nodeState.String")
	if ok {
		reta = FuncAuxMocknodeStateString.(func(recvn nodeState) (reta string))(recvn)
	} else {
		panic("FuncAuxMocknodeStateString ")
	}
	AuxMockIncrementRecorderAuxMocknodeStateString()
	return
}

//
// Mock: (recvh *HostInfo)Peer()(reta string)
//

type MockArgsTypeHostInfoPeer struct {
	ApomockCallNumber int
}

var LastMockArgsHostInfoPeer MockArgsTypeHostInfoPeer

// (recvh *HostInfo)AuxMockPeer()(reta string) - Generated mock function
func (recvh *HostInfo) AuxMockPeer() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.HostInfo.Peer")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.Peer")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.Peer")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrHostInfoPeer  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoPeer int = 0

var condRecorderAuxMockPtrHostInfoPeer *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoPeer(i int) {
	condRecorderAuxMockPtrHostInfoPeer.L.Lock()
	for recorderAuxMockPtrHostInfoPeer < i {
		condRecorderAuxMockPtrHostInfoPeer.Wait()
	}
	condRecorderAuxMockPtrHostInfoPeer.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoPeer() {
	condRecorderAuxMockPtrHostInfoPeer.L.Lock()
	recorderAuxMockPtrHostInfoPeer++
	condRecorderAuxMockPtrHostInfoPeer.L.Unlock()
	condRecorderAuxMockPtrHostInfoPeer.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoPeer() (ret int) {
	condRecorderAuxMockPtrHostInfoPeer.L.Lock()
	ret = recorderAuxMockPtrHostInfoPeer
	condRecorderAuxMockPtrHostInfoPeer.L.Unlock()
	return
}

// (recvh *HostInfo)Peer - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) Peer() (reta string) {
	FuncAuxMockPtrHostInfoPeer, ok := apomock.GetRegisteredFunc("gocql.HostInfo.Peer")
	if ok {
		reta = FuncAuxMockPtrHostInfoPeer.(func(recvh *HostInfo) (reta string))(recvh)
	} else {
		panic("FuncAuxMockPtrHostInfoPeer ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoPeer()
	return
}

//
// Mock: (recvh *HostInfo)Rack()(reta string)
//

type MockArgsTypeHostInfoRack struct {
	ApomockCallNumber int
}

var LastMockArgsHostInfoRack MockArgsTypeHostInfoRack

// (recvh *HostInfo)AuxMockRack()(reta string) - Generated mock function
func (recvh *HostInfo) AuxMockRack() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.HostInfo.Rack")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.Rack")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.Rack")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrHostInfoRack  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoRack int = 0

var condRecorderAuxMockPtrHostInfoRack *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoRack(i int) {
	condRecorderAuxMockPtrHostInfoRack.L.Lock()
	for recorderAuxMockPtrHostInfoRack < i {
		condRecorderAuxMockPtrHostInfoRack.Wait()
	}
	condRecorderAuxMockPtrHostInfoRack.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoRack() {
	condRecorderAuxMockPtrHostInfoRack.L.Lock()
	recorderAuxMockPtrHostInfoRack++
	condRecorderAuxMockPtrHostInfoRack.L.Unlock()
	condRecorderAuxMockPtrHostInfoRack.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoRack() (ret int) {
	condRecorderAuxMockPtrHostInfoRack.L.Lock()
	ret = recorderAuxMockPtrHostInfoRack
	condRecorderAuxMockPtrHostInfoRack.L.Unlock()
	return
}

// (recvh *HostInfo)Rack - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) Rack() (reta string) {
	FuncAuxMockPtrHostInfoRack, ok := apomock.GetRegisteredFunc("gocql.HostInfo.Rack")
	if ok {
		reta = FuncAuxMockPtrHostInfoRack.(func(recvh *HostInfo) (reta string))(recvh)
	} else {
		panic("FuncAuxMockPtrHostInfoRack ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoRack()
	return
}

//
// Mock: (recvh *HostInfo)update(argfrom *HostInfo)()
//

type MockArgsTypeHostInfoupdate struct {
	ApomockCallNumber int
	Argfrom           *HostInfo
}

var LastMockArgsHostInfoupdate MockArgsTypeHostInfoupdate

// (recvh *HostInfo)AuxMockupdate(argfrom *HostInfo)() - Generated mock function
func (recvh *HostInfo) AuxMockupdate(argfrom *HostInfo) {
	LastMockArgsHostInfoupdate = MockArgsTypeHostInfoupdate{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrHostInfoupdate(),
		Argfrom:           argfrom,
	}
	return
}

// RecorderAuxMockPtrHostInfoupdate  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoupdate int = 0

var condRecorderAuxMockPtrHostInfoupdate *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoupdate(i int) {
	condRecorderAuxMockPtrHostInfoupdate.L.Lock()
	for recorderAuxMockPtrHostInfoupdate < i {
		condRecorderAuxMockPtrHostInfoupdate.Wait()
	}
	condRecorderAuxMockPtrHostInfoupdate.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoupdate() {
	condRecorderAuxMockPtrHostInfoupdate.L.Lock()
	recorderAuxMockPtrHostInfoupdate++
	condRecorderAuxMockPtrHostInfoupdate.L.Unlock()
	condRecorderAuxMockPtrHostInfoupdate.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoupdate() (ret int) {
	condRecorderAuxMockPtrHostInfoupdate.L.Lock()
	ret = recorderAuxMockPtrHostInfoupdate
	condRecorderAuxMockPtrHostInfoupdate.L.Unlock()
	return
}

// (recvh *HostInfo)update - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) update(argfrom *HostInfo) {
	FuncAuxMockPtrHostInfoupdate, ok := apomock.GetRegisteredFunc("gocql.HostInfo.update")
	if ok {
		FuncAuxMockPtrHostInfoupdate.(func(recvh *HostInfo, argfrom *HostInfo))(recvh, argfrom)
	} else {
		panic("FuncAuxMockPtrHostInfoupdate ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoupdate()
	return
}

//
// Mock: checkSystemLocal(argcontrol *controlConn)(reta bool, retb error)
//

type MockArgsTypecheckSystemLocal struct {
	ApomockCallNumber int
	Argcontrol        *controlConn
}

var LastMockArgscheckSystemLocal MockArgsTypecheckSystemLocal

// AuxMockcheckSystemLocal(argcontrol *controlConn)(reta bool, retb error) - Generated mock function
func AuxMockcheckSystemLocal(argcontrol *controlConn) (reta bool, retb error) {
	LastMockArgscheckSystemLocal = MockArgsTypecheckSystemLocal{
		ApomockCallNumber: AuxMockGetRecorderAuxMockcheckSystemLocal(),
		Argcontrol:        argcontrol,
	}
	rargs, rerr := apomock.GetNext("gocql.checkSystemLocal")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.checkSystemLocal")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.checkSystemLocal")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockcheckSystemLocal  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockcheckSystemLocal int = 0

var condRecorderAuxMockcheckSystemLocal *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockcheckSystemLocal(i int) {
	condRecorderAuxMockcheckSystemLocal.L.Lock()
	for recorderAuxMockcheckSystemLocal < i {
		condRecorderAuxMockcheckSystemLocal.Wait()
	}
	condRecorderAuxMockcheckSystemLocal.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockcheckSystemLocal() {
	condRecorderAuxMockcheckSystemLocal.L.Lock()
	recorderAuxMockcheckSystemLocal++
	condRecorderAuxMockcheckSystemLocal.L.Unlock()
	condRecorderAuxMockcheckSystemLocal.Broadcast()
}
func AuxMockGetRecorderAuxMockcheckSystemLocal() (ret int) {
	condRecorderAuxMockcheckSystemLocal.L.Lock()
	ret = recorderAuxMockcheckSystemLocal
	condRecorderAuxMockcheckSystemLocal.L.Unlock()
	return
}

// checkSystemLocal - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func checkSystemLocal(argcontrol *controlConn) (reta bool, retb error) {
	FuncAuxMockcheckSystemLocal, ok := apomock.GetRegisteredFunc("gocql.checkSystemLocal")
	if ok {
		reta, retb = FuncAuxMockcheckSystemLocal.(func(argcontrol *controlConn) (reta bool, retb error))(argcontrol)
	} else {
		panic("FuncAuxMockcheckSystemLocal ")
	}
	AuxMockIncrementRecorderAuxMockcheckSystemLocal()
	return
}

//
// Mock: (recvr *ringDescriber)GetHosts()(rethosts []*HostInfo, retpartitioner string, reterr error)
//

type MockArgsTyperingDescriberGetHosts struct {
	ApomockCallNumber int
}

var LastMockArgsringDescriberGetHosts MockArgsTyperingDescriberGetHosts

// (recvr *ringDescriber)AuxMockGetHosts()(rethosts []*HostInfo, retpartitioner string, reterr error) - Generated mock function
func (recvr *ringDescriber) AuxMockGetHosts() (rethosts []*HostInfo, retpartitioner string, reterr error) {
	rargs, rerr := apomock.GetNext("gocql.ringDescriber.GetHosts")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ringDescriber.GetHosts")
	} else if rargs.NumArgs() != 3 {
		panic("All return parameters not provided for method:gocql.ringDescriber.GetHosts")
	}
	if rargs.GetArg(0) != nil {
		rethosts = rargs.GetArg(0).([]*HostInfo)
	}
	if rargs.GetArg(1) != nil {
		retpartitioner = rargs.GetArg(1).(string)
	}
	if rargs.GetArg(2) != nil {
		reterr = rargs.GetArg(2).(error)
	}
	return
}

// RecorderAuxMockPtrringDescriberGetHosts  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrringDescriberGetHosts int = 0

var condRecorderAuxMockPtrringDescriberGetHosts *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrringDescriberGetHosts(i int) {
	condRecorderAuxMockPtrringDescriberGetHosts.L.Lock()
	for recorderAuxMockPtrringDescriberGetHosts < i {
		condRecorderAuxMockPtrringDescriberGetHosts.Wait()
	}
	condRecorderAuxMockPtrringDescriberGetHosts.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrringDescriberGetHosts() {
	condRecorderAuxMockPtrringDescriberGetHosts.L.Lock()
	recorderAuxMockPtrringDescriberGetHosts++
	condRecorderAuxMockPtrringDescriberGetHosts.L.Unlock()
	condRecorderAuxMockPtrringDescriberGetHosts.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrringDescriberGetHosts() (ret int) {
	condRecorderAuxMockPtrringDescriberGetHosts.L.Lock()
	ret = recorderAuxMockPtrringDescriberGetHosts
	condRecorderAuxMockPtrringDescriberGetHosts.L.Unlock()
	return
}

// (recvr *ringDescriber)GetHosts - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *ringDescriber) GetHosts() (rethosts []*HostInfo, retpartitioner string, reterr error) {
	FuncAuxMockPtrringDescriberGetHosts, ok := apomock.GetRegisteredFunc("gocql.ringDescriber.GetHosts")
	if ok {
		rethosts, retpartitioner, reterr = FuncAuxMockPtrringDescriberGetHosts.(func(recvr *ringDescriber) (rethosts []*HostInfo, retpartitioner string, reterr error))(recvr)
	} else {
		panic("FuncAuxMockPtrringDescriberGetHosts ")
	}
	AuxMockIncrementRecorderAuxMockPtrringDescriberGetHosts()
	return
}

//
// Mock: (recvr *ringDescriber)refreshRing()(reta error)
//

type MockArgsTyperingDescriberrefreshRing struct {
	ApomockCallNumber int
}

var LastMockArgsringDescriberrefreshRing MockArgsTyperingDescriberrefreshRing

// (recvr *ringDescriber)AuxMockrefreshRing()(reta error) - Generated mock function
func (recvr *ringDescriber) AuxMockrefreshRing() (reta error) {
	rargs, rerr := apomock.GetNext("gocql.ringDescriber.refreshRing")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ringDescriber.refreshRing")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.ringDescriber.refreshRing")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrringDescriberrefreshRing  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrringDescriberrefreshRing int = 0

var condRecorderAuxMockPtrringDescriberrefreshRing *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrringDescriberrefreshRing(i int) {
	condRecorderAuxMockPtrringDescriberrefreshRing.L.Lock()
	for recorderAuxMockPtrringDescriberrefreshRing < i {
		condRecorderAuxMockPtrringDescriberrefreshRing.Wait()
	}
	condRecorderAuxMockPtrringDescriberrefreshRing.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrringDescriberrefreshRing() {
	condRecorderAuxMockPtrringDescriberrefreshRing.L.Lock()
	recorderAuxMockPtrringDescriberrefreshRing++
	condRecorderAuxMockPtrringDescriberrefreshRing.L.Unlock()
	condRecorderAuxMockPtrringDescriberrefreshRing.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrringDescriberrefreshRing() (ret int) {
	condRecorderAuxMockPtrringDescriberrefreshRing.L.Lock()
	ret = recorderAuxMockPtrringDescriberrefreshRing
	condRecorderAuxMockPtrringDescriberrefreshRing.L.Unlock()
	return
}

// (recvr *ringDescriber)refreshRing - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *ringDescriber) refreshRing() (reta error) {
	FuncAuxMockPtrringDescriberrefreshRing, ok := apomock.GetRegisteredFunc("gocql.ringDescriber.refreshRing")
	if ok {
		reta = FuncAuxMockPtrringDescriberrefreshRing.(func(recvr *ringDescriber) (reta error))(recvr)
	} else {
		panic("FuncAuxMockPtrringDescriberrefreshRing ")
	}
	AuxMockIncrementRecorderAuxMockPtrringDescriberrefreshRing()
	return
}

//
// Mock: (recvc cassVersion)nodeUpDelay()(reta time.Duration)
//

type MockArgsTypecassVersionnodeUpDelay struct {
	ApomockCallNumber int
}

var LastMockArgscassVersionnodeUpDelay MockArgsTypecassVersionnodeUpDelay

// (recvc cassVersion)AuxMocknodeUpDelay()(reta time.Duration) - Generated mock function
func (recvc cassVersion) AuxMocknodeUpDelay() (reta time.Duration) {
	rargs, rerr := apomock.GetNext("gocql.cassVersion.nodeUpDelay")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.cassVersion.nodeUpDelay")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.cassVersion.nodeUpDelay")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(time.Duration)
	}
	return
}

// RecorderAuxMockcassVersionnodeUpDelay  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockcassVersionnodeUpDelay int = 0

var condRecorderAuxMockcassVersionnodeUpDelay *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockcassVersionnodeUpDelay(i int) {
	condRecorderAuxMockcassVersionnodeUpDelay.L.Lock()
	for recorderAuxMockcassVersionnodeUpDelay < i {
		condRecorderAuxMockcassVersionnodeUpDelay.Wait()
	}
	condRecorderAuxMockcassVersionnodeUpDelay.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockcassVersionnodeUpDelay() {
	condRecorderAuxMockcassVersionnodeUpDelay.L.Lock()
	recorderAuxMockcassVersionnodeUpDelay++
	condRecorderAuxMockcassVersionnodeUpDelay.L.Unlock()
	condRecorderAuxMockcassVersionnodeUpDelay.Broadcast()
}
func AuxMockGetRecorderAuxMockcassVersionnodeUpDelay() (ret int) {
	condRecorderAuxMockcassVersionnodeUpDelay.L.Lock()
	ret = recorderAuxMockcassVersionnodeUpDelay
	condRecorderAuxMockcassVersionnodeUpDelay.L.Unlock()
	return
}

// (recvc cassVersion)nodeUpDelay - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc cassVersion) nodeUpDelay() (reta time.Duration) {
	FuncAuxMockcassVersionnodeUpDelay, ok := apomock.GetRegisteredFunc("gocql.cassVersion.nodeUpDelay")
	if ok {
		reta = FuncAuxMockcassVersionnodeUpDelay.(func(recvc cassVersion) (reta time.Duration))(recvc)
	} else {
		panic("FuncAuxMockcassVersionnodeUpDelay ")
	}
	AuxMockIncrementRecorderAuxMockcassVersionnodeUpDelay()
	return
}

//
// Mock: (recvh *HostInfo)Equal(arghost *HostInfo)(reta bool)
//

type MockArgsTypeHostInfoEqual struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgsHostInfoEqual MockArgsTypeHostInfoEqual

// (recvh *HostInfo)AuxMockEqual(arghost *HostInfo)(reta bool) - Generated mock function
func (recvh *HostInfo) AuxMockEqual(arghost *HostInfo) (reta bool) {
	LastMockArgsHostInfoEqual = MockArgsTypeHostInfoEqual{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrHostInfoEqual(),
		Arghost:           arghost,
	}
	rargs, rerr := apomock.GetNext("gocql.HostInfo.Equal")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.Equal")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.Equal")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrHostInfoEqual  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoEqual int = 0

var condRecorderAuxMockPtrHostInfoEqual *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoEqual(i int) {
	condRecorderAuxMockPtrHostInfoEqual.L.Lock()
	for recorderAuxMockPtrHostInfoEqual < i {
		condRecorderAuxMockPtrHostInfoEqual.Wait()
	}
	condRecorderAuxMockPtrHostInfoEqual.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoEqual() {
	condRecorderAuxMockPtrHostInfoEqual.L.Lock()
	recorderAuxMockPtrHostInfoEqual++
	condRecorderAuxMockPtrHostInfoEqual.L.Unlock()
	condRecorderAuxMockPtrHostInfoEqual.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoEqual() (ret int) {
	condRecorderAuxMockPtrHostInfoEqual.L.Lock()
	ret = recorderAuxMockPtrHostInfoEqual
	condRecorderAuxMockPtrHostInfoEqual.L.Unlock()
	return
}

// (recvh *HostInfo)Equal - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) Equal(arghost *HostInfo) (reta bool) {
	FuncAuxMockPtrHostInfoEqual, ok := apomock.GetRegisteredFunc("gocql.HostInfo.Equal")
	if ok {
		reta = FuncAuxMockPtrHostInfoEqual.(func(recvh *HostInfo, arghost *HostInfo) (reta bool))(recvh, arghost)
	} else {
		panic("FuncAuxMockPtrHostInfoEqual ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoEqual()
	return
}

//
// Mock: (recvh *HostInfo)setRack(argrack string)(reta *HostInfo)
//

type MockArgsTypeHostInfosetRack struct {
	ApomockCallNumber int
	Argrack           string
}

var LastMockArgsHostInfosetRack MockArgsTypeHostInfosetRack

// (recvh *HostInfo)AuxMocksetRack(argrack string)(reta *HostInfo) - Generated mock function
func (recvh *HostInfo) AuxMocksetRack(argrack string) (reta *HostInfo) {
	LastMockArgsHostInfosetRack = MockArgsTypeHostInfosetRack{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrHostInfosetRack(),
		Argrack:           argrack,
	}
	rargs, rerr := apomock.GetNext("gocql.HostInfo.setRack")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.setRack")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.setRack")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrHostInfosetRack  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfosetRack int = 0

var condRecorderAuxMockPtrHostInfosetRack *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfosetRack(i int) {
	condRecorderAuxMockPtrHostInfosetRack.L.Lock()
	for recorderAuxMockPtrHostInfosetRack < i {
		condRecorderAuxMockPtrHostInfosetRack.Wait()
	}
	condRecorderAuxMockPtrHostInfosetRack.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfosetRack() {
	condRecorderAuxMockPtrHostInfosetRack.L.Lock()
	recorderAuxMockPtrHostInfosetRack++
	condRecorderAuxMockPtrHostInfosetRack.L.Unlock()
	condRecorderAuxMockPtrHostInfosetRack.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfosetRack() (ret int) {
	condRecorderAuxMockPtrHostInfosetRack.L.Lock()
	ret = recorderAuxMockPtrHostInfosetRack
	condRecorderAuxMockPtrHostInfosetRack.L.Unlock()
	return
}

// (recvh *HostInfo)setRack - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) setRack(argrack string) (reta *HostInfo) {
	FuncAuxMockPtrHostInfosetRack, ok := apomock.GetRegisteredFunc("gocql.HostInfo.setRack")
	if ok {
		reta = FuncAuxMockPtrHostInfosetRack.(func(recvh *HostInfo, argrack string) (reta *HostInfo))(recvh, argrack)
	} else {
		panic("FuncAuxMockPtrHostInfosetRack ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfosetRack()
	return
}

//
// Mock: (recvh *HostInfo)Port()(reta int)
//

type MockArgsTypeHostInfoPort struct {
	ApomockCallNumber int
}

var LastMockArgsHostInfoPort MockArgsTypeHostInfoPort

// (recvh *HostInfo)AuxMockPort()(reta int) - Generated mock function
func (recvh *HostInfo) AuxMockPort() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.HostInfo.Port")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.Port")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.Port")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrHostInfoPort  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoPort int = 0

var condRecorderAuxMockPtrHostInfoPort *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoPort(i int) {
	condRecorderAuxMockPtrHostInfoPort.L.Lock()
	for recorderAuxMockPtrHostInfoPort < i {
		condRecorderAuxMockPtrHostInfoPort.Wait()
	}
	condRecorderAuxMockPtrHostInfoPort.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoPort() {
	condRecorderAuxMockPtrHostInfoPort.L.Lock()
	recorderAuxMockPtrHostInfoPort++
	condRecorderAuxMockPtrHostInfoPort.L.Unlock()
	condRecorderAuxMockPtrHostInfoPort.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoPort() (ret int) {
	condRecorderAuxMockPtrHostInfoPort.L.Lock()
	ret = recorderAuxMockPtrHostInfoPort
	condRecorderAuxMockPtrHostInfoPort.L.Unlock()
	return
}

// (recvh *HostInfo)Port - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) Port() (reta int) {
	FuncAuxMockPtrHostInfoPort, ok := apomock.GetRegisteredFunc("gocql.HostInfo.Port")
	if ok {
		reta = FuncAuxMockPtrHostInfoPort.(func(recvh *HostInfo) (reta int))(recvh)
	} else {
		panic("FuncAuxMockPtrHostInfoPort ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoPort()
	return
}

//
// Mock: (recvh *HostInfo)setPort(argport int)(reta *HostInfo)
//

type MockArgsTypeHostInfosetPort struct {
	ApomockCallNumber int
	Argport           int
}

var LastMockArgsHostInfosetPort MockArgsTypeHostInfosetPort

// (recvh *HostInfo)AuxMocksetPort(argport int)(reta *HostInfo) - Generated mock function
func (recvh *HostInfo) AuxMocksetPort(argport int) (reta *HostInfo) {
	LastMockArgsHostInfosetPort = MockArgsTypeHostInfosetPort{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrHostInfosetPort(),
		Argport:           argport,
	}
	rargs, rerr := apomock.GetNext("gocql.HostInfo.setPort")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.setPort")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.setPort")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrHostInfosetPort  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfosetPort int = 0

var condRecorderAuxMockPtrHostInfosetPort *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfosetPort(i int) {
	condRecorderAuxMockPtrHostInfosetPort.L.Lock()
	for recorderAuxMockPtrHostInfosetPort < i {
		condRecorderAuxMockPtrHostInfosetPort.Wait()
	}
	condRecorderAuxMockPtrHostInfosetPort.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfosetPort() {
	condRecorderAuxMockPtrHostInfosetPort.L.Lock()
	recorderAuxMockPtrHostInfosetPort++
	condRecorderAuxMockPtrHostInfosetPort.L.Unlock()
	condRecorderAuxMockPtrHostInfosetPort.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfosetPort() (ret int) {
	condRecorderAuxMockPtrHostInfosetPort.L.Lock()
	ret = recorderAuxMockPtrHostInfosetPort
	condRecorderAuxMockPtrHostInfosetPort.L.Unlock()
	return
}

// (recvh *HostInfo)setPort - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) setPort(argport int) (reta *HostInfo) {
	FuncAuxMockPtrHostInfosetPort, ok := apomock.GetRegisteredFunc("gocql.HostInfo.setPort")
	if ok {
		reta = FuncAuxMockPtrHostInfosetPort.(func(recvh *HostInfo, argport int) (reta *HostInfo))(recvh, argport)
	} else {
		panic("FuncAuxMockPtrHostInfosetPort ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfosetPort()
	return
}

//
// Mock: (recvh *HostInfo)setPeer(argpeer string)(reta *HostInfo)
//

type MockArgsTypeHostInfosetPeer struct {
	ApomockCallNumber int
	Argpeer           string
}

var LastMockArgsHostInfosetPeer MockArgsTypeHostInfosetPeer

// (recvh *HostInfo)AuxMocksetPeer(argpeer string)(reta *HostInfo) - Generated mock function
func (recvh *HostInfo) AuxMocksetPeer(argpeer string) (reta *HostInfo) {
	LastMockArgsHostInfosetPeer = MockArgsTypeHostInfosetPeer{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrHostInfosetPeer(),
		Argpeer:           argpeer,
	}
	rargs, rerr := apomock.GetNext("gocql.HostInfo.setPeer")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.setPeer")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.setPeer")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrHostInfosetPeer  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfosetPeer int = 0

var condRecorderAuxMockPtrHostInfosetPeer *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfosetPeer(i int) {
	condRecorderAuxMockPtrHostInfosetPeer.L.Lock()
	for recorderAuxMockPtrHostInfosetPeer < i {
		condRecorderAuxMockPtrHostInfosetPeer.Wait()
	}
	condRecorderAuxMockPtrHostInfosetPeer.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfosetPeer() {
	condRecorderAuxMockPtrHostInfosetPeer.L.Lock()
	recorderAuxMockPtrHostInfosetPeer++
	condRecorderAuxMockPtrHostInfosetPeer.L.Unlock()
	condRecorderAuxMockPtrHostInfosetPeer.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfosetPeer() (ret int) {
	condRecorderAuxMockPtrHostInfosetPeer.L.Lock()
	ret = recorderAuxMockPtrHostInfosetPeer
	condRecorderAuxMockPtrHostInfosetPeer.L.Unlock()
	return
}

// (recvh *HostInfo)setPeer - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) setPeer(argpeer string) (reta *HostInfo) {
	FuncAuxMockPtrHostInfosetPeer, ok := apomock.GetRegisteredFunc("gocql.HostInfo.setPeer")
	if ok {
		reta = FuncAuxMockPtrHostInfosetPeer.(func(recvh *HostInfo, argpeer string) (reta *HostInfo))(recvh, argpeer)
	} else {
		panic("FuncAuxMockPtrHostInfosetPeer ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfosetPeer()
	return
}

//
// Mock: (recvh *HostInfo)DataCenter()(reta string)
//

type MockArgsTypeHostInfoDataCenter struct {
	ApomockCallNumber int
}

var LastMockArgsHostInfoDataCenter MockArgsTypeHostInfoDataCenter

// (recvh *HostInfo)AuxMockDataCenter()(reta string) - Generated mock function
func (recvh *HostInfo) AuxMockDataCenter() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.HostInfo.DataCenter")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.DataCenter")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.DataCenter")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrHostInfoDataCenter  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoDataCenter int = 0

var condRecorderAuxMockPtrHostInfoDataCenter *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoDataCenter(i int) {
	condRecorderAuxMockPtrHostInfoDataCenter.L.Lock()
	for recorderAuxMockPtrHostInfoDataCenter < i {
		condRecorderAuxMockPtrHostInfoDataCenter.Wait()
	}
	condRecorderAuxMockPtrHostInfoDataCenter.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoDataCenter() {
	condRecorderAuxMockPtrHostInfoDataCenter.L.Lock()
	recorderAuxMockPtrHostInfoDataCenter++
	condRecorderAuxMockPtrHostInfoDataCenter.L.Unlock()
	condRecorderAuxMockPtrHostInfoDataCenter.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoDataCenter() (ret int) {
	condRecorderAuxMockPtrHostInfoDataCenter.L.Lock()
	ret = recorderAuxMockPtrHostInfoDataCenter
	condRecorderAuxMockPtrHostInfoDataCenter.L.Unlock()
	return
}

// (recvh *HostInfo)DataCenter - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) DataCenter() (reta string) {
	FuncAuxMockPtrHostInfoDataCenter, ok := apomock.GetRegisteredFunc("gocql.HostInfo.DataCenter")
	if ok {
		reta = FuncAuxMockPtrHostInfoDataCenter.(func(recvh *HostInfo) (reta string))(recvh)
	} else {
		panic("FuncAuxMockPtrHostInfoDataCenter ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoDataCenter()
	return
}

//
// Mock: (recvh *HostInfo)HostID()(reta string)
//

type MockArgsTypeHostInfoHostID struct {
	ApomockCallNumber int
}

var LastMockArgsHostInfoHostID MockArgsTypeHostInfoHostID

// (recvh *HostInfo)AuxMockHostID()(reta string) - Generated mock function
func (recvh *HostInfo) AuxMockHostID() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.HostInfo.HostID")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.HostID")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.HostID")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrHostInfoHostID  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoHostID int = 0

var condRecorderAuxMockPtrHostInfoHostID *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoHostID(i int) {
	condRecorderAuxMockPtrHostInfoHostID.L.Lock()
	for recorderAuxMockPtrHostInfoHostID < i {
		condRecorderAuxMockPtrHostInfoHostID.Wait()
	}
	condRecorderAuxMockPtrHostInfoHostID.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoHostID() {
	condRecorderAuxMockPtrHostInfoHostID.L.Lock()
	recorderAuxMockPtrHostInfoHostID++
	condRecorderAuxMockPtrHostInfoHostID.L.Unlock()
	condRecorderAuxMockPtrHostInfoHostID.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoHostID() (ret int) {
	condRecorderAuxMockPtrHostInfoHostID.L.Lock()
	ret = recorderAuxMockPtrHostInfoHostID
	condRecorderAuxMockPtrHostInfoHostID.L.Unlock()
	return
}

// (recvh *HostInfo)HostID - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) HostID() (reta string) {
	FuncAuxMockPtrHostInfoHostID, ok := apomock.GetRegisteredFunc("gocql.HostInfo.HostID")
	if ok {
		reta = FuncAuxMockPtrHostInfoHostID.(func(recvh *HostInfo) (reta string))(recvh)
	} else {
		panic("FuncAuxMockPtrHostInfoHostID ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoHostID()
	return
}

//
// Mock: (recvh *HostInfo)Version()(reta cassVersion)
//

type MockArgsTypeHostInfoVersion struct {
	ApomockCallNumber int
}

var LastMockArgsHostInfoVersion MockArgsTypeHostInfoVersion

// (recvh *HostInfo)AuxMockVersion()(reta cassVersion) - Generated mock function
func (recvh *HostInfo) AuxMockVersion() (reta cassVersion) {
	rargs, rerr := apomock.GetNext("gocql.HostInfo.Version")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.Version")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.Version")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(cassVersion)
	}
	return
}

// RecorderAuxMockPtrHostInfoVersion  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoVersion int = 0

var condRecorderAuxMockPtrHostInfoVersion *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoVersion(i int) {
	condRecorderAuxMockPtrHostInfoVersion.L.Lock()
	for recorderAuxMockPtrHostInfoVersion < i {
		condRecorderAuxMockPtrHostInfoVersion.Wait()
	}
	condRecorderAuxMockPtrHostInfoVersion.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoVersion() {
	condRecorderAuxMockPtrHostInfoVersion.L.Lock()
	recorderAuxMockPtrHostInfoVersion++
	condRecorderAuxMockPtrHostInfoVersion.L.Unlock()
	condRecorderAuxMockPtrHostInfoVersion.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoVersion() (ret int) {
	condRecorderAuxMockPtrHostInfoVersion.L.Lock()
	ret = recorderAuxMockPtrHostInfoVersion
	condRecorderAuxMockPtrHostInfoVersion.L.Unlock()
	return
}

// (recvh *HostInfo)Version - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) Version() (reta cassVersion) {
	FuncAuxMockPtrHostInfoVersion, ok := apomock.GetRegisteredFunc("gocql.HostInfo.Version")
	if ok {
		reta = FuncAuxMockPtrHostInfoVersion.(func(recvh *HostInfo) (reta cassVersion))(recvh)
	} else {
		panic("FuncAuxMockPtrHostInfoVersion ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoVersion()
	return
}

//
// Mock: (recvh *HostInfo)String()(reta string)
//

type MockArgsTypeHostInfoString struct {
	ApomockCallNumber int
}

var LastMockArgsHostInfoString MockArgsTypeHostInfoString

// (recvh *HostInfo)AuxMockString()(reta string) - Generated mock function
func (recvh *HostInfo) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.HostInfo.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrHostInfoString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoString int = 0

var condRecorderAuxMockPtrHostInfoString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoString(i int) {
	condRecorderAuxMockPtrHostInfoString.L.Lock()
	for recorderAuxMockPtrHostInfoString < i {
		condRecorderAuxMockPtrHostInfoString.Wait()
	}
	condRecorderAuxMockPtrHostInfoString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoString() {
	condRecorderAuxMockPtrHostInfoString.L.Lock()
	recorderAuxMockPtrHostInfoString++
	condRecorderAuxMockPtrHostInfoString.L.Unlock()
	condRecorderAuxMockPtrHostInfoString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoString() (ret int) {
	condRecorderAuxMockPtrHostInfoString.L.Lock()
	ret = recorderAuxMockPtrHostInfoString
	condRecorderAuxMockPtrHostInfoString.L.Unlock()
	return
}

// (recvh *HostInfo)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) String() (reta string) {
	FuncAuxMockPtrHostInfoString, ok := apomock.GetRegisteredFunc("gocql.HostInfo.String")
	if ok {
		reta = FuncAuxMockPtrHostInfoString.(func(recvh *HostInfo) (reta string))(recvh)
	} else {
		panic("FuncAuxMockPtrHostInfoString ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoString()
	return
}

//
// Mock: (recvr *ringDescriber)matchFilter(arghost *HostInfo)(reta bool)
//

type MockArgsTyperingDescribermatchFilter struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgsringDescribermatchFilter MockArgsTyperingDescribermatchFilter

// (recvr *ringDescriber)AuxMockmatchFilter(arghost *HostInfo)(reta bool) - Generated mock function
func (recvr *ringDescriber) AuxMockmatchFilter(arghost *HostInfo) (reta bool) {
	LastMockArgsringDescribermatchFilter = MockArgsTyperingDescribermatchFilter{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrringDescribermatchFilter(),
		Arghost:           arghost,
	}
	rargs, rerr := apomock.GetNext("gocql.ringDescriber.matchFilter")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ringDescriber.matchFilter")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.ringDescriber.matchFilter")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrringDescribermatchFilter  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrringDescribermatchFilter int = 0

var condRecorderAuxMockPtrringDescribermatchFilter *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrringDescribermatchFilter(i int) {
	condRecorderAuxMockPtrringDescribermatchFilter.L.Lock()
	for recorderAuxMockPtrringDescribermatchFilter < i {
		condRecorderAuxMockPtrringDescribermatchFilter.Wait()
	}
	condRecorderAuxMockPtrringDescribermatchFilter.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrringDescribermatchFilter() {
	condRecorderAuxMockPtrringDescribermatchFilter.L.Lock()
	recorderAuxMockPtrringDescribermatchFilter++
	condRecorderAuxMockPtrringDescribermatchFilter.L.Unlock()
	condRecorderAuxMockPtrringDescribermatchFilter.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrringDescribermatchFilter() (ret int) {
	condRecorderAuxMockPtrringDescribermatchFilter.L.Lock()
	ret = recorderAuxMockPtrringDescribermatchFilter
	condRecorderAuxMockPtrringDescribermatchFilter.L.Unlock()
	return
}

// (recvr *ringDescriber)matchFilter - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *ringDescriber) matchFilter(arghost *HostInfo) (reta bool) {
	FuncAuxMockPtrringDescribermatchFilter, ok := apomock.GetRegisteredFunc("gocql.ringDescriber.matchFilter")
	if ok {
		reta = FuncAuxMockPtrringDescribermatchFilter.(func(recvr *ringDescriber, arghost *HostInfo) (reta bool))(recvr, arghost)
	} else {
		panic("FuncAuxMockPtrringDescribermatchFilter ")
	}
	AuxMockIncrementRecorderAuxMockPtrringDescribermatchFilter()
	return
}

//
// Mock: (recvc cassVersion)Before(argmajor int, argminor int, argpatch int)(reta bool)
//

type MockArgsTypecassVersionBefore struct {
	ApomockCallNumber int
	Argmajor          int
	Argminor          int
	Argpatch          int
}

var LastMockArgscassVersionBefore MockArgsTypecassVersionBefore

// (recvc cassVersion)AuxMockBefore(argmajor int, argminor int, argpatch int)(reta bool) - Generated mock function
func (recvc cassVersion) AuxMockBefore(argmajor int, argminor int, argpatch int) (reta bool) {
	LastMockArgscassVersionBefore = MockArgsTypecassVersionBefore{
		ApomockCallNumber: AuxMockGetRecorderAuxMockcassVersionBefore(),
		Argmajor:          argmajor,
		Argminor:          argminor,
		Argpatch:          argpatch,
	}
	rargs, rerr := apomock.GetNext("gocql.cassVersion.Before")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.cassVersion.Before")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.cassVersion.Before")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockcassVersionBefore  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockcassVersionBefore int = 0

var condRecorderAuxMockcassVersionBefore *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockcassVersionBefore(i int) {
	condRecorderAuxMockcassVersionBefore.L.Lock()
	for recorderAuxMockcassVersionBefore < i {
		condRecorderAuxMockcassVersionBefore.Wait()
	}
	condRecorderAuxMockcassVersionBefore.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockcassVersionBefore() {
	condRecorderAuxMockcassVersionBefore.L.Lock()
	recorderAuxMockcassVersionBefore++
	condRecorderAuxMockcassVersionBefore.L.Unlock()
	condRecorderAuxMockcassVersionBefore.Broadcast()
}
func AuxMockGetRecorderAuxMockcassVersionBefore() (ret int) {
	condRecorderAuxMockcassVersionBefore.L.Lock()
	ret = recorderAuxMockcassVersionBefore
	condRecorderAuxMockcassVersionBefore.L.Unlock()
	return
}

// (recvc cassVersion)Before - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc cassVersion) Before(argmajor int, argminor int, argpatch int) (reta bool) {
	FuncAuxMockcassVersionBefore, ok := apomock.GetRegisteredFunc("gocql.cassVersion.Before")
	if ok {
		reta = FuncAuxMockcassVersionBefore.(func(recvc cassVersion, argmajor int, argminor int, argpatch int) (reta bool))(recvc, argmajor, argminor, argpatch)
	} else {
		panic("FuncAuxMockcassVersionBefore ")
	}
	AuxMockIncrementRecorderAuxMockcassVersionBefore()
	return
}

//
// Mock: (recvc cassVersion)String()(reta string)
//

type MockArgsTypecassVersionString struct {
	ApomockCallNumber int
}

var LastMockArgscassVersionString MockArgsTypecassVersionString

// (recvc cassVersion)AuxMockString()(reta string) - Generated mock function
func (recvc cassVersion) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.cassVersion.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.cassVersion.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.cassVersion.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockcassVersionString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockcassVersionString int = 0

var condRecorderAuxMockcassVersionString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockcassVersionString(i int) {
	condRecorderAuxMockcassVersionString.L.Lock()
	for recorderAuxMockcassVersionString < i {
		condRecorderAuxMockcassVersionString.Wait()
	}
	condRecorderAuxMockcassVersionString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockcassVersionString() {
	condRecorderAuxMockcassVersionString.L.Lock()
	recorderAuxMockcassVersionString++
	condRecorderAuxMockcassVersionString.L.Unlock()
	condRecorderAuxMockcassVersionString.Broadcast()
}
func AuxMockGetRecorderAuxMockcassVersionString() (ret int) {
	condRecorderAuxMockcassVersionString.L.Lock()
	ret = recorderAuxMockcassVersionString
	condRecorderAuxMockcassVersionString.L.Unlock()
	return
}

// (recvc cassVersion)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc cassVersion) String() (reta string) {
	FuncAuxMockcassVersionString, ok := apomock.GetRegisteredFunc("gocql.cassVersion.String")
	if ok {
		reta = FuncAuxMockcassVersionString.(func(recvc cassVersion) (reta string))(recvc)
	} else {
		panic("FuncAuxMockcassVersionString ")
	}
	AuxMockIncrementRecorderAuxMockcassVersionString()
	return
}

//
// Mock: (recvh *HostInfo)setDataCenter(argdataCenter string)(reta *HostInfo)
//

type MockArgsTypeHostInfosetDataCenter struct {
	ApomockCallNumber int
	ArgdataCenter     string
}

var LastMockArgsHostInfosetDataCenter MockArgsTypeHostInfosetDataCenter

// (recvh *HostInfo)AuxMocksetDataCenter(argdataCenter string)(reta *HostInfo) - Generated mock function
func (recvh *HostInfo) AuxMocksetDataCenter(argdataCenter string) (reta *HostInfo) {
	LastMockArgsHostInfosetDataCenter = MockArgsTypeHostInfosetDataCenter{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrHostInfosetDataCenter(),
		ArgdataCenter:     argdataCenter,
	}
	rargs, rerr := apomock.GetNext("gocql.HostInfo.setDataCenter")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.setDataCenter")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.setDataCenter")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrHostInfosetDataCenter  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfosetDataCenter int = 0

var condRecorderAuxMockPtrHostInfosetDataCenter *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfosetDataCenter(i int) {
	condRecorderAuxMockPtrHostInfosetDataCenter.L.Lock()
	for recorderAuxMockPtrHostInfosetDataCenter < i {
		condRecorderAuxMockPtrHostInfosetDataCenter.Wait()
	}
	condRecorderAuxMockPtrHostInfosetDataCenter.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfosetDataCenter() {
	condRecorderAuxMockPtrHostInfosetDataCenter.L.Lock()
	recorderAuxMockPtrHostInfosetDataCenter++
	condRecorderAuxMockPtrHostInfosetDataCenter.L.Unlock()
	condRecorderAuxMockPtrHostInfosetDataCenter.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfosetDataCenter() (ret int) {
	condRecorderAuxMockPtrHostInfosetDataCenter.L.Lock()
	ret = recorderAuxMockPtrHostInfosetDataCenter
	condRecorderAuxMockPtrHostInfosetDataCenter.L.Unlock()
	return
}

// (recvh *HostInfo)setDataCenter - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) setDataCenter(argdataCenter string) (reta *HostInfo) {
	FuncAuxMockPtrHostInfosetDataCenter, ok := apomock.GetRegisteredFunc("gocql.HostInfo.setDataCenter")
	if ok {
		reta = FuncAuxMockPtrHostInfosetDataCenter.(func(recvh *HostInfo, argdataCenter string) (reta *HostInfo))(recvh, argdataCenter)
	} else {
		panic("FuncAuxMockPtrHostInfosetDataCenter ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfosetDataCenter()
	return
}

//
// Mock: (recvh *HostInfo)setState(argstate nodeState)(reta *HostInfo)
//

type MockArgsTypeHostInfosetState struct {
	ApomockCallNumber int
	Argstate          nodeState
}

var LastMockArgsHostInfosetState MockArgsTypeHostInfosetState

// (recvh *HostInfo)AuxMocksetState(argstate nodeState)(reta *HostInfo) - Generated mock function
func (recvh *HostInfo) AuxMocksetState(argstate nodeState) (reta *HostInfo) {
	LastMockArgsHostInfosetState = MockArgsTypeHostInfosetState{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrHostInfosetState(),
		Argstate:          argstate,
	}
	rargs, rerr := apomock.GetNext("gocql.HostInfo.setState")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.setState")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.setState")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrHostInfosetState  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfosetState int = 0

var condRecorderAuxMockPtrHostInfosetState *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfosetState(i int) {
	condRecorderAuxMockPtrHostInfosetState.L.Lock()
	for recorderAuxMockPtrHostInfosetState < i {
		condRecorderAuxMockPtrHostInfosetState.Wait()
	}
	condRecorderAuxMockPtrHostInfosetState.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfosetState() {
	condRecorderAuxMockPtrHostInfosetState.L.Lock()
	recorderAuxMockPtrHostInfosetState++
	condRecorderAuxMockPtrHostInfosetState.L.Unlock()
	condRecorderAuxMockPtrHostInfosetState.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfosetState() (ret int) {
	condRecorderAuxMockPtrHostInfosetState.L.Lock()
	ret = recorderAuxMockPtrHostInfosetState
	condRecorderAuxMockPtrHostInfosetState.L.Unlock()
	return
}

// (recvh *HostInfo)setState - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) setState(argstate nodeState) (reta *HostInfo) {
	FuncAuxMockPtrHostInfosetState, ok := apomock.GetRegisteredFunc("gocql.HostInfo.setState")
	if ok {
		reta = FuncAuxMockPtrHostInfosetState.(func(recvh *HostInfo, argstate nodeState) (reta *HostInfo))(recvh, argstate)
	} else {
		panic("FuncAuxMockPtrHostInfosetState ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfosetState()
	return
}

//
// Mock: (recvh *HostInfo)Tokens()(reta []string)
//

type MockArgsTypeHostInfoTokens struct {
	ApomockCallNumber int
}

var LastMockArgsHostInfoTokens MockArgsTypeHostInfoTokens

// (recvh *HostInfo)AuxMockTokens()(reta []string) - Generated mock function
func (recvh *HostInfo) AuxMockTokens() (reta []string) {
	rargs, rerr := apomock.GetNext("gocql.HostInfo.Tokens")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.Tokens")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.Tokens")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]string)
	}
	return
}

// RecorderAuxMockPtrHostInfoTokens  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoTokens int = 0

var condRecorderAuxMockPtrHostInfoTokens *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoTokens(i int) {
	condRecorderAuxMockPtrHostInfoTokens.L.Lock()
	for recorderAuxMockPtrHostInfoTokens < i {
		condRecorderAuxMockPtrHostInfoTokens.Wait()
	}
	condRecorderAuxMockPtrHostInfoTokens.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoTokens() {
	condRecorderAuxMockPtrHostInfoTokens.L.Lock()
	recorderAuxMockPtrHostInfoTokens++
	condRecorderAuxMockPtrHostInfoTokens.L.Unlock()
	condRecorderAuxMockPtrHostInfoTokens.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoTokens() (ret int) {
	condRecorderAuxMockPtrHostInfoTokens.L.Lock()
	ret = recorderAuxMockPtrHostInfoTokens
	condRecorderAuxMockPtrHostInfoTokens.L.Unlock()
	return
}

// (recvh *HostInfo)Tokens - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) Tokens() (reta []string) {
	FuncAuxMockPtrHostInfoTokens, ok := apomock.GetRegisteredFunc("gocql.HostInfo.Tokens")
	if ok {
		reta = FuncAuxMockPtrHostInfoTokens.(func(recvh *HostInfo) (reta []string))(recvh)
	} else {
		panic("FuncAuxMockPtrHostInfoTokens ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoTokens()
	return
}

//
// Mock: (recvh *HostInfo)setTokens(argtokens []string)(reta *HostInfo)
//

type MockArgsTypeHostInfosetTokens struct {
	ApomockCallNumber int
	Argtokens         []string
}

var LastMockArgsHostInfosetTokens MockArgsTypeHostInfosetTokens

// (recvh *HostInfo)AuxMocksetTokens(argtokens []string)(reta *HostInfo) - Generated mock function
func (recvh *HostInfo) AuxMocksetTokens(argtokens []string) (reta *HostInfo) {
	LastMockArgsHostInfosetTokens = MockArgsTypeHostInfosetTokens{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrHostInfosetTokens(),
		Argtokens:         argtokens,
	}
	rargs, rerr := apomock.GetNext("gocql.HostInfo.setTokens")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.setTokens")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.setTokens")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrHostInfosetTokens  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfosetTokens int = 0

var condRecorderAuxMockPtrHostInfosetTokens *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfosetTokens(i int) {
	condRecorderAuxMockPtrHostInfosetTokens.L.Lock()
	for recorderAuxMockPtrHostInfosetTokens < i {
		condRecorderAuxMockPtrHostInfosetTokens.Wait()
	}
	condRecorderAuxMockPtrHostInfosetTokens.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfosetTokens() {
	condRecorderAuxMockPtrHostInfosetTokens.L.Lock()
	recorderAuxMockPtrHostInfosetTokens++
	condRecorderAuxMockPtrHostInfosetTokens.L.Unlock()
	condRecorderAuxMockPtrHostInfosetTokens.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfosetTokens() (ret int) {
	condRecorderAuxMockPtrHostInfosetTokens.L.Lock()
	ret = recorderAuxMockPtrHostInfosetTokens
	condRecorderAuxMockPtrHostInfosetTokens.L.Unlock()
	return
}

// (recvh *HostInfo)setTokens - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) setTokens(argtokens []string) (reta *HostInfo) {
	FuncAuxMockPtrHostInfosetTokens, ok := apomock.GetRegisteredFunc("gocql.HostInfo.setTokens")
	if ok {
		reta = FuncAuxMockPtrHostInfosetTokens.(func(recvh *HostInfo, argtokens []string) (reta *HostInfo))(recvh, argtokens)
	} else {
		panic("FuncAuxMockPtrHostInfosetTokens ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfosetTokens()
	return
}

//
// Mock: (recvc *cassVersion)unmarshal(argdata []byte)(reta error)
//

type MockArgsTypecassVersionunmarshal struct {
	ApomockCallNumber int
	Argdata           []byte
}

var LastMockArgscassVersionunmarshal MockArgsTypecassVersionunmarshal

// (recvc *cassVersion)AuxMockunmarshal(argdata []byte)(reta error) - Generated mock function
func (recvc *cassVersion) AuxMockunmarshal(argdata []byte) (reta error) {
	LastMockArgscassVersionunmarshal = MockArgsTypecassVersionunmarshal{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcassVersionunmarshal(),
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.cassVersion.unmarshal")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.cassVersion.unmarshal")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.cassVersion.unmarshal")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrcassVersionunmarshal  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcassVersionunmarshal int = 0

var condRecorderAuxMockPtrcassVersionunmarshal *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcassVersionunmarshal(i int) {
	condRecorderAuxMockPtrcassVersionunmarshal.L.Lock()
	for recorderAuxMockPtrcassVersionunmarshal < i {
		condRecorderAuxMockPtrcassVersionunmarshal.Wait()
	}
	condRecorderAuxMockPtrcassVersionunmarshal.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcassVersionunmarshal() {
	condRecorderAuxMockPtrcassVersionunmarshal.L.Lock()
	recorderAuxMockPtrcassVersionunmarshal++
	condRecorderAuxMockPtrcassVersionunmarshal.L.Unlock()
	condRecorderAuxMockPtrcassVersionunmarshal.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcassVersionunmarshal() (ret int) {
	condRecorderAuxMockPtrcassVersionunmarshal.L.Lock()
	ret = recorderAuxMockPtrcassVersionunmarshal
	condRecorderAuxMockPtrcassVersionunmarshal.L.Unlock()
	return
}

// (recvc *cassVersion)unmarshal - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *cassVersion) unmarshal(argdata []byte) (reta error) {
	FuncAuxMockPtrcassVersionunmarshal, ok := apomock.GetRegisteredFunc("gocql.cassVersion.unmarshal")
	if ok {
		reta = FuncAuxMockPtrcassVersionunmarshal.(func(recvc *cassVersion, argdata []byte) (reta error))(recvc, argdata)
	} else {
		panic("FuncAuxMockPtrcassVersionunmarshal ")
	}
	AuxMockIncrementRecorderAuxMockPtrcassVersionunmarshal()
	return
}

//
// Mock: (recvh *HostInfo)State()(reta nodeState)
//

type MockArgsTypeHostInfoState struct {
	ApomockCallNumber int
}

var LastMockArgsHostInfoState MockArgsTypeHostInfoState

// (recvh *HostInfo)AuxMockState()(reta nodeState) - Generated mock function
func (recvh *HostInfo) AuxMockState() (reta nodeState) {
	rargs, rerr := apomock.GetNext("gocql.HostInfo.State")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostInfo.State")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostInfo.State")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(nodeState)
	}
	return
}

// RecorderAuxMockPtrHostInfoState  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrHostInfoState int = 0

var condRecorderAuxMockPtrHostInfoState *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrHostInfoState(i int) {
	condRecorderAuxMockPtrHostInfoState.L.Lock()
	for recorderAuxMockPtrHostInfoState < i {
		condRecorderAuxMockPtrHostInfoState.Wait()
	}
	condRecorderAuxMockPtrHostInfoState.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrHostInfoState() {
	condRecorderAuxMockPtrHostInfoState.L.Lock()
	recorderAuxMockPtrHostInfoState++
	condRecorderAuxMockPtrHostInfoState.L.Unlock()
	condRecorderAuxMockPtrHostInfoState.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrHostInfoState() (ret int) {
	condRecorderAuxMockPtrHostInfoState.L.Lock()
	ret = recorderAuxMockPtrHostInfoState
	condRecorderAuxMockPtrHostInfoState.L.Unlock()
	return
}

// (recvh *HostInfo)State - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *HostInfo) State() (reta nodeState) {
	FuncAuxMockPtrHostInfoState, ok := apomock.GetRegisteredFunc("gocql.HostInfo.State")
	if ok {
		reta = FuncAuxMockPtrHostInfoState.(func(recvh *HostInfo) (reta nodeState))(recvh)
	} else {
		panic("FuncAuxMockPtrHostInfoState ")
	}
	AuxMockIncrementRecorderAuxMockPtrHostInfoState()
	return
}
