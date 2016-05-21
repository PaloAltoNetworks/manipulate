// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"
import "github.com/hailocab/go-hostpool"

import "sync/atomic"

func init() {
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructCowHostList, apomockNewStructCowHostList)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructSelectedHostPoolHost, apomockNewStructSelectedHostPoolHost)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructRoundRobinHostPolicy, apomockNewStructRoundRobinHostPolicy)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructTokenAwareHostPolicy, apomockNewStructTokenAwareHostPolicy)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructHostPoolHostPolicy, apomockNewStructHostPoolHostPolicy)

	apomock.RegisterFunc("gocql", "gocql.cowHostList.update", (*cowHostList).AuxMockupdate)
	apomock.RegisterFunc("gocql", "gocql.tokenAwareHostPolicy.Pick", (*tokenAwareHostPolicy).AuxMockPick)
	apomock.RegisterFunc("gocql", "gocql.hostPoolHostPolicy.HostUp", (*hostPoolHostPolicy).AuxMockHostUp)
	apomock.RegisterFunc("gocql", "gocql.hostPoolHostPolicy.HostDown", (*hostPoolHostPolicy).AuxMockHostDown)
	apomock.RegisterFunc("gocql", "gocql.roundRobinHostPolicy.RemoveHost", (*roundRobinHostPolicy).AuxMockRemoveHost)
	apomock.RegisterFunc("gocql", "gocql.tokenAwareHostPolicy.RemoveHost", (*tokenAwareHostPolicy).AuxMockRemoveHost)
	apomock.RegisterFunc("gocql", "gocql.hostPoolHostPolicy.RemoveHost", (*hostPoolHostPolicy).AuxMockRemoveHost)
	apomock.RegisterFunc("gocql", "gocql.cowHostList.String", (*cowHostList).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.cowHostList.get", (*cowHostList).AuxMockget)
	apomock.RegisterFunc("gocql", "gocql.roundRobinHostPolicy.AddHost", (*roundRobinHostPolicy).AuxMockAddHost)
	apomock.RegisterFunc("gocql", "gocql.tokenAwareHostPolicy.AddHost", (*tokenAwareHostPolicy).AuxMockAddHost)
	apomock.RegisterFunc("gocql", "gocql.hostPoolHostPolicy.Pick", (*hostPoolHostPolicy).AuxMockPick)
	apomock.RegisterFunc("gocql", "gocql.selectedHostPoolHost.Info", (selectedHostPoolHost).AuxMockInfo)
	apomock.RegisterFunc("gocql", "gocql.cowHostList.add", (*cowHostList).AuxMockadd)
	apomock.RegisterFunc("gocql", "gocql.RoundRobinHostPolicy", AuxMockRoundRobinHostPolicy)
	apomock.RegisterFunc("gocql", "gocql.roundRobinHostPolicy.SetPartitioner", (*roundRobinHostPolicy).AuxMockSetPartitioner)
	apomock.RegisterFunc("gocql", "gocql.roundRobinHostPolicy.HostDown", (*roundRobinHostPolicy).AuxMockHostDown)
	apomock.RegisterFunc("gocql", "gocql.tokenAwareHostPolicy.HostDown", (*tokenAwareHostPolicy).AuxMockHostDown)
	apomock.RegisterFunc("gocql", "gocql.tokenAwareHostPolicy.resetTokenRing", (*tokenAwareHostPolicy).AuxMockresetTokenRing)
	apomock.RegisterFunc("gocql", "gocql.cowHostList.remove", (*cowHostList).AuxMockremove)
	apomock.RegisterFunc("gocql", "gocql.SimpleRetryPolicy.Attempt", (*SimpleRetryPolicy).AuxMockAttempt)
	apomock.RegisterFunc("gocql", "gocql.roundRobinHostPolicy.Pick", (*roundRobinHostPolicy).AuxMockPick)
	apomock.RegisterFunc("gocql", "gocql.roundRobinHostPolicy.HostUp", (*roundRobinHostPolicy).AuxMockHostUp)
	apomock.RegisterFunc("gocql", "gocql.tokenAwareHostPolicy.SetPartitioner", (*tokenAwareHostPolicy).AuxMockSetPartitioner)
	apomock.RegisterFunc("gocql", "gocql.hostPoolHostPolicy.SetHosts", (*hostPoolHostPolicy).AuxMockSetHosts)
	apomock.RegisterFunc("gocql", "gocql.TokenAwareHostPolicy", AuxMockTokenAwareHostPolicy)
	apomock.RegisterFunc("gocql", "gocql.HostPoolHostPolicy", AuxMockHostPoolHostPolicy)
	apomock.RegisterFunc("gocql", "gocql.hostPoolHostPolicy.SetPartitioner", (*hostPoolHostPolicy).AuxMockSetPartitioner)
	apomock.RegisterFunc("gocql", "gocql.selectedHostPoolHost.Mark", (selectedHostPoolHost).AuxMockMark)
	apomock.RegisterFunc("gocql", "gocql.selectedHost.Info", (*selectedHost).AuxMockInfo)
	apomock.RegisterFunc("gocql", "gocql.cowHostList.set", (*cowHostList).AuxMockset)
	apomock.RegisterFunc("gocql", "gocql.selectedHost.Mark", (*selectedHost).AuxMockMark)
	apomock.RegisterFunc("gocql", "gocql.tokenAwareHostPolicy.HostUp", (*tokenAwareHostPolicy).AuxMockHostUp)
	apomock.RegisterFunc("gocql", "gocql.hostPoolHostPolicy.AddHost", (*hostPoolHostPolicy).AuxMockAddHost)
}

const (
	ApomockStructCowHostList          = 49
	ApomockStructSelectedHostPoolHost = 50
	ApomockStructRoundRobinHostPolicy = 51
	ApomockStructTokenAwareHostPolicy = 52
	ApomockStructHostPoolHostPolicy   = 53
)

//
// Internal Types: in this package and their exportable versions
//
type cowHostList struct {
	list atomic.Value
	mu   sync.Mutex
}
type selectedHostPoolHost struct {
	policy *hostPoolHostPolicy
	info   *HostInfo
	hostR  hostpool.HostPoolResponse
}
type selectedHost HostInfo
type roundRobinHostPolicy struct {
	hosts cowHostList
	pos   uint32
	mu    sync.RWMutex
}
type tokenAwareHostPolicy struct {
	hosts       cowHostList
	mu          sync.RWMutex
	partitioner string
	tokenRing   *tokenRing
	fallback    HostSelectionPolicy
}
type hostPoolHostPolicy struct {
	hp      hostpool.HostPool
	mu      sync.RWMutex
	hostMap map[string]*HostInfo
}

//
// External Types: in this package
//
type NextHost func() SelectedHost

type RetryableQuery interface {
	Attempts() int
	GetConsistency() Consistency
}

type RetryPolicy interface {
	Attempt(RetryableQuery) bool
}

type SimpleRetryPolicy struct{ NumRetries int }

type HostStateNotifier interface {
	AddHost(host *HostInfo)
	RemoveHost(addr string)
	HostUp(host *HostInfo)
	HostDown(addr string)
}

type HostSelectionPolicy interface {
	HostStateNotifier
	SetPartitioner
	Pick(ExecutableQuery) NextHost
}

type SelectedHost interface {
	Info() *HostInfo
	Mark(error)
}

func apomockNewStructCowHostList() interface{}          { return &cowHostList{} }
func apomockNewStructSelectedHostPoolHost() interface{} { return &selectedHostPoolHost{} }
func apomockNewStructRoundRobinHostPolicy() interface{} { return &roundRobinHostPolicy{} }
func apomockNewStructTokenAwareHostPolicy() interface{} { return &tokenAwareHostPolicy{} }
func apomockNewStructHostPoolHostPolicy() interface{}   { return &hostPoolHostPolicy{} }

//
// Mock: (recvc *cowHostList)update(arghost *HostInfo)()
//

type MockArgsTypecowHostListupdate struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgscowHostListupdate MockArgsTypecowHostListupdate

// (recvc *cowHostList)AuxMockupdate(arghost *HostInfo)() - Generated mock function
func (recvc *cowHostList) AuxMockupdate(arghost *HostInfo) {
	LastMockArgscowHostListupdate = MockArgsTypecowHostListupdate{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcowHostListupdate(),
		Arghost:           arghost,
	}
	return
}

// RecorderAuxMockPtrcowHostListupdate  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcowHostListupdate int = 0

var condRecorderAuxMockPtrcowHostListupdate *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcowHostListupdate(i int) {
	condRecorderAuxMockPtrcowHostListupdate.L.Lock()
	for recorderAuxMockPtrcowHostListupdate < i {
		condRecorderAuxMockPtrcowHostListupdate.Wait()
	}
	condRecorderAuxMockPtrcowHostListupdate.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcowHostListupdate() {
	condRecorderAuxMockPtrcowHostListupdate.L.Lock()
	recorderAuxMockPtrcowHostListupdate++
	condRecorderAuxMockPtrcowHostListupdate.L.Unlock()
	condRecorderAuxMockPtrcowHostListupdate.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcowHostListupdate() (ret int) {
	condRecorderAuxMockPtrcowHostListupdate.L.Lock()
	ret = recorderAuxMockPtrcowHostListupdate
	condRecorderAuxMockPtrcowHostListupdate.L.Unlock()
	return
}

// (recvc *cowHostList)update - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *cowHostList) update(arghost *HostInfo) {
	FuncAuxMockPtrcowHostListupdate, ok := apomock.GetRegisteredFunc("gocql.cowHostList.update")
	if ok {
		FuncAuxMockPtrcowHostListupdate.(func(recvc *cowHostList, arghost *HostInfo))(recvc, arghost)
	} else {
		panic("FuncAuxMockPtrcowHostListupdate ")
	}
	AuxMockIncrementRecorderAuxMockPtrcowHostListupdate()
	return
}

//
// Mock: (recvt *tokenAwareHostPolicy)Pick(argqry ExecutableQuery)(reta NextHost)
//

type MockArgsTypetokenAwareHostPolicyPick struct {
	ApomockCallNumber int
	Argqry            ExecutableQuery
}

var LastMockArgstokenAwareHostPolicyPick MockArgsTypetokenAwareHostPolicyPick

// (recvt *tokenAwareHostPolicy)AuxMockPick(argqry ExecutableQuery)(reta NextHost) - Generated mock function
func (recvt *tokenAwareHostPolicy) AuxMockPick(argqry ExecutableQuery) (reta NextHost) {
	LastMockArgstokenAwareHostPolicyPick = MockArgsTypetokenAwareHostPolicyPick{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicyPick(),
		Argqry:            argqry,
	}
	rargs, rerr := apomock.GetNext("gocql.tokenAwareHostPolicy.Pick")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.tokenAwareHostPolicy.Pick")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.tokenAwareHostPolicy.Pick")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(NextHost)
	}
	return
}

// RecorderAuxMockPtrtokenAwareHostPolicyPick  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenAwareHostPolicyPick int = 0

var condRecorderAuxMockPtrtokenAwareHostPolicyPick *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenAwareHostPolicyPick(i int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyPick.L.Lock()
	for recorderAuxMockPtrtokenAwareHostPolicyPick < i {
		condRecorderAuxMockPtrtokenAwareHostPolicyPick.Wait()
	}
	condRecorderAuxMockPtrtokenAwareHostPolicyPick.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyPick() {
	condRecorderAuxMockPtrtokenAwareHostPolicyPick.L.Lock()
	recorderAuxMockPtrtokenAwareHostPolicyPick++
	condRecorderAuxMockPtrtokenAwareHostPolicyPick.L.Unlock()
	condRecorderAuxMockPtrtokenAwareHostPolicyPick.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicyPick() (ret int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyPick.L.Lock()
	ret = recorderAuxMockPtrtokenAwareHostPolicyPick
	condRecorderAuxMockPtrtokenAwareHostPolicyPick.L.Unlock()
	return
}

// (recvt *tokenAwareHostPolicy)Pick - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenAwareHostPolicy) Pick(argqry ExecutableQuery) (reta NextHost) {
	FuncAuxMockPtrtokenAwareHostPolicyPick, ok := apomock.GetRegisteredFunc("gocql.tokenAwareHostPolicy.Pick")
	if ok {
		reta = FuncAuxMockPtrtokenAwareHostPolicyPick.(func(recvt *tokenAwareHostPolicy, argqry ExecutableQuery) (reta NextHost))(recvt, argqry)
	} else {
		panic("FuncAuxMockPtrtokenAwareHostPolicyPick ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyPick()
	return
}

//
// Mock: (recvr *hostPoolHostPolicy)HostUp(arghost *HostInfo)()
//

type MockArgsTypehostPoolHostPolicyHostUp struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgshostPoolHostPolicyHostUp MockArgsTypehostPoolHostPolicyHostUp

// (recvr *hostPoolHostPolicy)AuxMockHostUp(arghost *HostInfo)() - Generated mock function
func (recvr *hostPoolHostPolicy) AuxMockHostUp(arghost *HostInfo) {
	LastMockArgshostPoolHostPolicyHostUp = MockArgsTypehostPoolHostPolicyHostUp{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrhostPoolHostPolicyHostUp(),
		Arghost:           arghost,
	}
	return
}

// RecorderAuxMockPtrhostPoolHostPolicyHostUp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostPoolHostPolicyHostUp int = 0

var condRecorderAuxMockPtrhostPoolHostPolicyHostUp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostPoolHostPolicyHostUp(i int) {
	condRecorderAuxMockPtrhostPoolHostPolicyHostUp.L.Lock()
	for recorderAuxMockPtrhostPoolHostPolicyHostUp < i {
		condRecorderAuxMockPtrhostPoolHostPolicyHostUp.Wait()
	}
	condRecorderAuxMockPtrhostPoolHostPolicyHostUp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicyHostUp() {
	condRecorderAuxMockPtrhostPoolHostPolicyHostUp.L.Lock()
	recorderAuxMockPtrhostPoolHostPolicyHostUp++
	condRecorderAuxMockPtrhostPoolHostPolicyHostUp.L.Unlock()
	condRecorderAuxMockPtrhostPoolHostPolicyHostUp.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostPoolHostPolicyHostUp() (ret int) {
	condRecorderAuxMockPtrhostPoolHostPolicyHostUp.L.Lock()
	ret = recorderAuxMockPtrhostPoolHostPolicyHostUp
	condRecorderAuxMockPtrhostPoolHostPolicyHostUp.L.Unlock()
	return
}

// (recvr *hostPoolHostPolicy)HostUp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *hostPoolHostPolicy) HostUp(arghost *HostInfo) {
	FuncAuxMockPtrhostPoolHostPolicyHostUp, ok := apomock.GetRegisteredFunc("gocql.hostPoolHostPolicy.HostUp")
	if ok {
		FuncAuxMockPtrhostPoolHostPolicyHostUp.(func(recvr *hostPoolHostPolicy, arghost *HostInfo))(recvr, arghost)
	} else {
		panic("FuncAuxMockPtrhostPoolHostPolicyHostUp ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicyHostUp()
	return
}

//
// Mock: (recvr *hostPoolHostPolicy)HostDown(argaddr string)()
//

type MockArgsTypehostPoolHostPolicyHostDown struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgshostPoolHostPolicyHostDown MockArgsTypehostPoolHostPolicyHostDown

// (recvr *hostPoolHostPolicy)AuxMockHostDown(argaddr string)() - Generated mock function
func (recvr *hostPoolHostPolicy) AuxMockHostDown(argaddr string) {
	LastMockArgshostPoolHostPolicyHostDown = MockArgsTypehostPoolHostPolicyHostDown{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrhostPoolHostPolicyHostDown(),
		Argaddr:           argaddr,
	}
	return
}

// RecorderAuxMockPtrhostPoolHostPolicyHostDown  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostPoolHostPolicyHostDown int = 0

var condRecorderAuxMockPtrhostPoolHostPolicyHostDown *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostPoolHostPolicyHostDown(i int) {
	condRecorderAuxMockPtrhostPoolHostPolicyHostDown.L.Lock()
	for recorderAuxMockPtrhostPoolHostPolicyHostDown < i {
		condRecorderAuxMockPtrhostPoolHostPolicyHostDown.Wait()
	}
	condRecorderAuxMockPtrhostPoolHostPolicyHostDown.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicyHostDown() {
	condRecorderAuxMockPtrhostPoolHostPolicyHostDown.L.Lock()
	recorderAuxMockPtrhostPoolHostPolicyHostDown++
	condRecorderAuxMockPtrhostPoolHostPolicyHostDown.L.Unlock()
	condRecorderAuxMockPtrhostPoolHostPolicyHostDown.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostPoolHostPolicyHostDown() (ret int) {
	condRecorderAuxMockPtrhostPoolHostPolicyHostDown.L.Lock()
	ret = recorderAuxMockPtrhostPoolHostPolicyHostDown
	condRecorderAuxMockPtrhostPoolHostPolicyHostDown.L.Unlock()
	return
}

// (recvr *hostPoolHostPolicy)HostDown - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *hostPoolHostPolicy) HostDown(argaddr string) {
	FuncAuxMockPtrhostPoolHostPolicyHostDown, ok := apomock.GetRegisteredFunc("gocql.hostPoolHostPolicy.HostDown")
	if ok {
		FuncAuxMockPtrhostPoolHostPolicyHostDown.(func(recvr *hostPoolHostPolicy, argaddr string))(recvr, argaddr)
	} else {
		panic("FuncAuxMockPtrhostPoolHostPolicyHostDown ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicyHostDown()
	return
}

//
// Mock: (recvr *roundRobinHostPolicy)RemoveHost(argaddr string)()
//

type MockArgsTyperoundRobinHostPolicyRemoveHost struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgsroundRobinHostPolicyRemoveHost MockArgsTyperoundRobinHostPolicyRemoveHost

// (recvr *roundRobinHostPolicy)AuxMockRemoveHost(argaddr string)() - Generated mock function
func (recvr *roundRobinHostPolicy) AuxMockRemoveHost(argaddr string) {
	LastMockArgsroundRobinHostPolicyRemoveHost = MockArgsTyperoundRobinHostPolicyRemoveHost{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrroundRobinHostPolicyRemoveHost(),
		Argaddr:           argaddr,
	}
	return
}

// RecorderAuxMockPtrroundRobinHostPolicyRemoveHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrroundRobinHostPolicyRemoveHost int = 0

var condRecorderAuxMockPtrroundRobinHostPolicyRemoveHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrroundRobinHostPolicyRemoveHost(i int) {
	condRecorderAuxMockPtrroundRobinHostPolicyRemoveHost.L.Lock()
	for recorderAuxMockPtrroundRobinHostPolicyRemoveHost < i {
		condRecorderAuxMockPtrroundRobinHostPolicyRemoveHost.Wait()
	}
	condRecorderAuxMockPtrroundRobinHostPolicyRemoveHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicyRemoveHost() {
	condRecorderAuxMockPtrroundRobinHostPolicyRemoveHost.L.Lock()
	recorderAuxMockPtrroundRobinHostPolicyRemoveHost++
	condRecorderAuxMockPtrroundRobinHostPolicyRemoveHost.L.Unlock()
	condRecorderAuxMockPtrroundRobinHostPolicyRemoveHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrroundRobinHostPolicyRemoveHost() (ret int) {
	condRecorderAuxMockPtrroundRobinHostPolicyRemoveHost.L.Lock()
	ret = recorderAuxMockPtrroundRobinHostPolicyRemoveHost
	condRecorderAuxMockPtrroundRobinHostPolicyRemoveHost.L.Unlock()
	return
}

// (recvr *roundRobinHostPolicy)RemoveHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *roundRobinHostPolicy) RemoveHost(argaddr string) {
	FuncAuxMockPtrroundRobinHostPolicyRemoveHost, ok := apomock.GetRegisteredFunc("gocql.roundRobinHostPolicy.RemoveHost")
	if ok {
		FuncAuxMockPtrroundRobinHostPolicyRemoveHost.(func(recvr *roundRobinHostPolicy, argaddr string))(recvr, argaddr)
	} else {
		panic("FuncAuxMockPtrroundRobinHostPolicyRemoveHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicyRemoveHost()
	return
}

//
// Mock: (recvt *tokenAwareHostPolicy)RemoveHost(argaddr string)()
//

type MockArgsTypetokenAwareHostPolicyRemoveHost struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgstokenAwareHostPolicyRemoveHost MockArgsTypetokenAwareHostPolicyRemoveHost

// (recvt *tokenAwareHostPolicy)AuxMockRemoveHost(argaddr string)() - Generated mock function
func (recvt *tokenAwareHostPolicy) AuxMockRemoveHost(argaddr string) {
	LastMockArgstokenAwareHostPolicyRemoveHost = MockArgsTypetokenAwareHostPolicyRemoveHost{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost(),
		Argaddr:           argaddr,
	}
	return
}

// RecorderAuxMockPtrtokenAwareHostPolicyRemoveHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenAwareHostPolicyRemoveHost int = 0

var condRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost(i int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost.L.Lock()
	for recorderAuxMockPtrtokenAwareHostPolicyRemoveHost < i {
		condRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost.Wait()
	}
	condRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost() {
	condRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost.L.Lock()
	recorderAuxMockPtrtokenAwareHostPolicyRemoveHost++
	condRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost.L.Unlock()
	condRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost() (ret int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost.L.Lock()
	ret = recorderAuxMockPtrtokenAwareHostPolicyRemoveHost
	condRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost.L.Unlock()
	return
}

// (recvt *tokenAwareHostPolicy)RemoveHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenAwareHostPolicy) RemoveHost(argaddr string) {
	FuncAuxMockPtrtokenAwareHostPolicyRemoveHost, ok := apomock.GetRegisteredFunc("gocql.tokenAwareHostPolicy.RemoveHost")
	if ok {
		FuncAuxMockPtrtokenAwareHostPolicyRemoveHost.(func(recvt *tokenAwareHostPolicy, argaddr string))(recvt, argaddr)
	} else {
		panic("FuncAuxMockPtrtokenAwareHostPolicyRemoveHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyRemoveHost()
	return
}

//
// Mock: (recvr *hostPoolHostPolicy)RemoveHost(argaddr string)()
//

type MockArgsTypehostPoolHostPolicyRemoveHost struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgshostPoolHostPolicyRemoveHost MockArgsTypehostPoolHostPolicyRemoveHost

// (recvr *hostPoolHostPolicy)AuxMockRemoveHost(argaddr string)() - Generated mock function
func (recvr *hostPoolHostPolicy) AuxMockRemoveHost(argaddr string) {
	LastMockArgshostPoolHostPolicyRemoveHost = MockArgsTypehostPoolHostPolicyRemoveHost{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrhostPoolHostPolicyRemoveHost(),
		Argaddr:           argaddr,
	}
	return
}

// RecorderAuxMockPtrhostPoolHostPolicyRemoveHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostPoolHostPolicyRemoveHost int = 0

var condRecorderAuxMockPtrhostPoolHostPolicyRemoveHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostPoolHostPolicyRemoveHost(i int) {
	condRecorderAuxMockPtrhostPoolHostPolicyRemoveHost.L.Lock()
	for recorderAuxMockPtrhostPoolHostPolicyRemoveHost < i {
		condRecorderAuxMockPtrhostPoolHostPolicyRemoveHost.Wait()
	}
	condRecorderAuxMockPtrhostPoolHostPolicyRemoveHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicyRemoveHost() {
	condRecorderAuxMockPtrhostPoolHostPolicyRemoveHost.L.Lock()
	recorderAuxMockPtrhostPoolHostPolicyRemoveHost++
	condRecorderAuxMockPtrhostPoolHostPolicyRemoveHost.L.Unlock()
	condRecorderAuxMockPtrhostPoolHostPolicyRemoveHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostPoolHostPolicyRemoveHost() (ret int) {
	condRecorderAuxMockPtrhostPoolHostPolicyRemoveHost.L.Lock()
	ret = recorderAuxMockPtrhostPoolHostPolicyRemoveHost
	condRecorderAuxMockPtrhostPoolHostPolicyRemoveHost.L.Unlock()
	return
}

// (recvr *hostPoolHostPolicy)RemoveHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *hostPoolHostPolicy) RemoveHost(argaddr string) {
	FuncAuxMockPtrhostPoolHostPolicyRemoveHost, ok := apomock.GetRegisteredFunc("gocql.hostPoolHostPolicy.RemoveHost")
	if ok {
		FuncAuxMockPtrhostPoolHostPolicyRemoveHost.(func(recvr *hostPoolHostPolicy, argaddr string))(recvr, argaddr)
	} else {
		panic("FuncAuxMockPtrhostPoolHostPolicyRemoveHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicyRemoveHost()
	return
}

//
// Mock: (recvc *cowHostList)String()(reta string)
//

type MockArgsTypecowHostListString struct {
	ApomockCallNumber int
}

var LastMockArgscowHostListString MockArgsTypecowHostListString

// (recvc *cowHostList)AuxMockString()(reta string) - Generated mock function
func (recvc *cowHostList) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.cowHostList.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.cowHostList.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.cowHostList.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrcowHostListString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcowHostListString int = 0

var condRecorderAuxMockPtrcowHostListString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcowHostListString(i int) {
	condRecorderAuxMockPtrcowHostListString.L.Lock()
	for recorderAuxMockPtrcowHostListString < i {
		condRecorderAuxMockPtrcowHostListString.Wait()
	}
	condRecorderAuxMockPtrcowHostListString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcowHostListString() {
	condRecorderAuxMockPtrcowHostListString.L.Lock()
	recorderAuxMockPtrcowHostListString++
	condRecorderAuxMockPtrcowHostListString.L.Unlock()
	condRecorderAuxMockPtrcowHostListString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcowHostListString() (ret int) {
	condRecorderAuxMockPtrcowHostListString.L.Lock()
	ret = recorderAuxMockPtrcowHostListString
	condRecorderAuxMockPtrcowHostListString.L.Unlock()
	return
}

// (recvc *cowHostList)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *cowHostList) String() (reta string) {
	FuncAuxMockPtrcowHostListString, ok := apomock.GetRegisteredFunc("gocql.cowHostList.String")
	if ok {
		reta = FuncAuxMockPtrcowHostListString.(func(recvc *cowHostList) (reta string))(recvc)
	} else {
		panic("FuncAuxMockPtrcowHostListString ")
	}
	AuxMockIncrementRecorderAuxMockPtrcowHostListString()
	return
}

//
// Mock: (recvc *cowHostList)get()(reta []*HostInfo)
//

type MockArgsTypecowHostListget struct {
	ApomockCallNumber int
}

var LastMockArgscowHostListget MockArgsTypecowHostListget

// (recvc *cowHostList)AuxMockget()(reta []*HostInfo) - Generated mock function
func (recvc *cowHostList) AuxMockget() (reta []*HostInfo) {
	rargs, rerr := apomock.GetNext("gocql.cowHostList.get")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.cowHostList.get")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.cowHostList.get")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]*HostInfo)
	}
	return
}

// RecorderAuxMockPtrcowHostListget  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcowHostListget int = 0

var condRecorderAuxMockPtrcowHostListget *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcowHostListget(i int) {
	condRecorderAuxMockPtrcowHostListget.L.Lock()
	for recorderAuxMockPtrcowHostListget < i {
		condRecorderAuxMockPtrcowHostListget.Wait()
	}
	condRecorderAuxMockPtrcowHostListget.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcowHostListget() {
	condRecorderAuxMockPtrcowHostListget.L.Lock()
	recorderAuxMockPtrcowHostListget++
	condRecorderAuxMockPtrcowHostListget.L.Unlock()
	condRecorderAuxMockPtrcowHostListget.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcowHostListget() (ret int) {
	condRecorderAuxMockPtrcowHostListget.L.Lock()
	ret = recorderAuxMockPtrcowHostListget
	condRecorderAuxMockPtrcowHostListget.L.Unlock()
	return
}

// (recvc *cowHostList)get - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *cowHostList) get() (reta []*HostInfo) {
	FuncAuxMockPtrcowHostListget, ok := apomock.GetRegisteredFunc("gocql.cowHostList.get")
	if ok {
		reta = FuncAuxMockPtrcowHostListget.(func(recvc *cowHostList) (reta []*HostInfo))(recvc)
	} else {
		panic("FuncAuxMockPtrcowHostListget ")
	}
	AuxMockIncrementRecorderAuxMockPtrcowHostListget()
	return
}

//
// Mock: (recvr *roundRobinHostPolicy)AddHost(arghost *HostInfo)()
//

type MockArgsTyperoundRobinHostPolicyAddHost struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgsroundRobinHostPolicyAddHost MockArgsTyperoundRobinHostPolicyAddHost

// (recvr *roundRobinHostPolicy)AuxMockAddHost(arghost *HostInfo)() - Generated mock function
func (recvr *roundRobinHostPolicy) AuxMockAddHost(arghost *HostInfo) {
	LastMockArgsroundRobinHostPolicyAddHost = MockArgsTyperoundRobinHostPolicyAddHost{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrroundRobinHostPolicyAddHost(),
		Arghost:           arghost,
	}
	return
}

// RecorderAuxMockPtrroundRobinHostPolicyAddHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrroundRobinHostPolicyAddHost int = 0

var condRecorderAuxMockPtrroundRobinHostPolicyAddHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrroundRobinHostPolicyAddHost(i int) {
	condRecorderAuxMockPtrroundRobinHostPolicyAddHost.L.Lock()
	for recorderAuxMockPtrroundRobinHostPolicyAddHost < i {
		condRecorderAuxMockPtrroundRobinHostPolicyAddHost.Wait()
	}
	condRecorderAuxMockPtrroundRobinHostPolicyAddHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicyAddHost() {
	condRecorderAuxMockPtrroundRobinHostPolicyAddHost.L.Lock()
	recorderAuxMockPtrroundRobinHostPolicyAddHost++
	condRecorderAuxMockPtrroundRobinHostPolicyAddHost.L.Unlock()
	condRecorderAuxMockPtrroundRobinHostPolicyAddHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrroundRobinHostPolicyAddHost() (ret int) {
	condRecorderAuxMockPtrroundRobinHostPolicyAddHost.L.Lock()
	ret = recorderAuxMockPtrroundRobinHostPolicyAddHost
	condRecorderAuxMockPtrroundRobinHostPolicyAddHost.L.Unlock()
	return
}

// (recvr *roundRobinHostPolicy)AddHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *roundRobinHostPolicy) AddHost(arghost *HostInfo) {
	FuncAuxMockPtrroundRobinHostPolicyAddHost, ok := apomock.GetRegisteredFunc("gocql.roundRobinHostPolicy.AddHost")
	if ok {
		FuncAuxMockPtrroundRobinHostPolicyAddHost.(func(recvr *roundRobinHostPolicy, arghost *HostInfo))(recvr, arghost)
	} else {
		panic("FuncAuxMockPtrroundRobinHostPolicyAddHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicyAddHost()
	return
}

//
// Mock: (recvt *tokenAwareHostPolicy)AddHost(arghost *HostInfo)()
//

type MockArgsTypetokenAwareHostPolicyAddHost struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgstokenAwareHostPolicyAddHost MockArgsTypetokenAwareHostPolicyAddHost

// (recvt *tokenAwareHostPolicy)AuxMockAddHost(arghost *HostInfo)() - Generated mock function
func (recvt *tokenAwareHostPolicy) AuxMockAddHost(arghost *HostInfo) {
	LastMockArgstokenAwareHostPolicyAddHost = MockArgsTypetokenAwareHostPolicyAddHost{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicyAddHost(),
		Arghost:           arghost,
	}
	return
}

// RecorderAuxMockPtrtokenAwareHostPolicyAddHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenAwareHostPolicyAddHost int = 0

var condRecorderAuxMockPtrtokenAwareHostPolicyAddHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenAwareHostPolicyAddHost(i int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyAddHost.L.Lock()
	for recorderAuxMockPtrtokenAwareHostPolicyAddHost < i {
		condRecorderAuxMockPtrtokenAwareHostPolicyAddHost.Wait()
	}
	condRecorderAuxMockPtrtokenAwareHostPolicyAddHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyAddHost() {
	condRecorderAuxMockPtrtokenAwareHostPolicyAddHost.L.Lock()
	recorderAuxMockPtrtokenAwareHostPolicyAddHost++
	condRecorderAuxMockPtrtokenAwareHostPolicyAddHost.L.Unlock()
	condRecorderAuxMockPtrtokenAwareHostPolicyAddHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicyAddHost() (ret int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyAddHost.L.Lock()
	ret = recorderAuxMockPtrtokenAwareHostPolicyAddHost
	condRecorderAuxMockPtrtokenAwareHostPolicyAddHost.L.Unlock()
	return
}

// (recvt *tokenAwareHostPolicy)AddHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenAwareHostPolicy) AddHost(arghost *HostInfo) {
	FuncAuxMockPtrtokenAwareHostPolicyAddHost, ok := apomock.GetRegisteredFunc("gocql.tokenAwareHostPolicy.AddHost")
	if ok {
		FuncAuxMockPtrtokenAwareHostPolicyAddHost.(func(recvt *tokenAwareHostPolicy, arghost *HostInfo))(recvt, arghost)
	} else {
		panic("FuncAuxMockPtrtokenAwareHostPolicyAddHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyAddHost()
	return
}

//
// Mock: (recvr *hostPoolHostPolicy)Pick(argqry ExecutableQuery)(reta NextHost)
//

type MockArgsTypehostPoolHostPolicyPick struct {
	ApomockCallNumber int
	Argqry            ExecutableQuery
}

var LastMockArgshostPoolHostPolicyPick MockArgsTypehostPoolHostPolicyPick

// (recvr *hostPoolHostPolicy)AuxMockPick(argqry ExecutableQuery)(reta NextHost) - Generated mock function
func (recvr *hostPoolHostPolicy) AuxMockPick(argqry ExecutableQuery) (reta NextHost) {
	LastMockArgshostPoolHostPolicyPick = MockArgsTypehostPoolHostPolicyPick{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrhostPoolHostPolicyPick(),
		Argqry:            argqry,
	}
	rargs, rerr := apomock.GetNext("gocql.hostPoolHostPolicy.Pick")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.hostPoolHostPolicy.Pick")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.hostPoolHostPolicy.Pick")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(NextHost)
	}
	return
}

// RecorderAuxMockPtrhostPoolHostPolicyPick  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostPoolHostPolicyPick int = 0

var condRecorderAuxMockPtrhostPoolHostPolicyPick *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostPoolHostPolicyPick(i int) {
	condRecorderAuxMockPtrhostPoolHostPolicyPick.L.Lock()
	for recorderAuxMockPtrhostPoolHostPolicyPick < i {
		condRecorderAuxMockPtrhostPoolHostPolicyPick.Wait()
	}
	condRecorderAuxMockPtrhostPoolHostPolicyPick.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicyPick() {
	condRecorderAuxMockPtrhostPoolHostPolicyPick.L.Lock()
	recorderAuxMockPtrhostPoolHostPolicyPick++
	condRecorderAuxMockPtrhostPoolHostPolicyPick.L.Unlock()
	condRecorderAuxMockPtrhostPoolHostPolicyPick.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostPoolHostPolicyPick() (ret int) {
	condRecorderAuxMockPtrhostPoolHostPolicyPick.L.Lock()
	ret = recorderAuxMockPtrhostPoolHostPolicyPick
	condRecorderAuxMockPtrhostPoolHostPolicyPick.L.Unlock()
	return
}

// (recvr *hostPoolHostPolicy)Pick - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *hostPoolHostPolicy) Pick(argqry ExecutableQuery) (reta NextHost) {
	FuncAuxMockPtrhostPoolHostPolicyPick, ok := apomock.GetRegisteredFunc("gocql.hostPoolHostPolicy.Pick")
	if ok {
		reta = FuncAuxMockPtrhostPoolHostPolicyPick.(func(recvr *hostPoolHostPolicy, argqry ExecutableQuery) (reta NextHost))(recvr, argqry)
	} else {
		panic("FuncAuxMockPtrhostPoolHostPolicyPick ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicyPick()
	return
}

//
// Mock: (recvhost selectedHostPoolHost)Info()(reta *HostInfo)
//

type MockArgsTypeselectedHostPoolHostInfo struct {
	ApomockCallNumber int
}

var LastMockArgsselectedHostPoolHostInfo MockArgsTypeselectedHostPoolHostInfo

// (recvhost selectedHostPoolHost)AuxMockInfo()(reta *HostInfo) - Generated mock function
func (recvhost selectedHostPoolHost) AuxMockInfo() (reta *HostInfo) {
	rargs, rerr := apomock.GetNext("gocql.selectedHostPoolHost.Info")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.selectedHostPoolHost.Info")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.selectedHostPoolHost.Info")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockselectedHostPoolHostInfo  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockselectedHostPoolHostInfo int = 0

var condRecorderAuxMockselectedHostPoolHostInfo *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockselectedHostPoolHostInfo(i int) {
	condRecorderAuxMockselectedHostPoolHostInfo.L.Lock()
	for recorderAuxMockselectedHostPoolHostInfo < i {
		condRecorderAuxMockselectedHostPoolHostInfo.Wait()
	}
	condRecorderAuxMockselectedHostPoolHostInfo.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockselectedHostPoolHostInfo() {
	condRecorderAuxMockselectedHostPoolHostInfo.L.Lock()
	recorderAuxMockselectedHostPoolHostInfo++
	condRecorderAuxMockselectedHostPoolHostInfo.L.Unlock()
	condRecorderAuxMockselectedHostPoolHostInfo.Broadcast()
}
func AuxMockGetRecorderAuxMockselectedHostPoolHostInfo() (ret int) {
	condRecorderAuxMockselectedHostPoolHostInfo.L.Lock()
	ret = recorderAuxMockselectedHostPoolHostInfo
	condRecorderAuxMockselectedHostPoolHostInfo.L.Unlock()
	return
}

// (recvhost selectedHostPoolHost)Info - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvhost selectedHostPoolHost) Info() (reta *HostInfo) {
	FuncAuxMockselectedHostPoolHostInfo, ok := apomock.GetRegisteredFunc("gocql.selectedHostPoolHost.Info")
	if ok {
		reta = FuncAuxMockselectedHostPoolHostInfo.(func(recvhost selectedHostPoolHost) (reta *HostInfo))(recvhost)
	} else {
		panic("FuncAuxMockselectedHostPoolHostInfo ")
	}
	AuxMockIncrementRecorderAuxMockselectedHostPoolHostInfo()
	return
}

//
// Mock: (recvc *cowHostList)add(arghost *HostInfo)(reta bool)
//

type MockArgsTypecowHostListadd struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgscowHostListadd MockArgsTypecowHostListadd

// (recvc *cowHostList)AuxMockadd(arghost *HostInfo)(reta bool) - Generated mock function
func (recvc *cowHostList) AuxMockadd(arghost *HostInfo) (reta bool) {
	LastMockArgscowHostListadd = MockArgsTypecowHostListadd{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcowHostListadd(),
		Arghost:           arghost,
	}
	rargs, rerr := apomock.GetNext("gocql.cowHostList.add")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.cowHostList.add")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.cowHostList.add")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrcowHostListadd  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcowHostListadd int = 0

var condRecorderAuxMockPtrcowHostListadd *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcowHostListadd(i int) {
	condRecorderAuxMockPtrcowHostListadd.L.Lock()
	for recorderAuxMockPtrcowHostListadd < i {
		condRecorderAuxMockPtrcowHostListadd.Wait()
	}
	condRecorderAuxMockPtrcowHostListadd.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcowHostListadd() {
	condRecorderAuxMockPtrcowHostListadd.L.Lock()
	recorderAuxMockPtrcowHostListadd++
	condRecorderAuxMockPtrcowHostListadd.L.Unlock()
	condRecorderAuxMockPtrcowHostListadd.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcowHostListadd() (ret int) {
	condRecorderAuxMockPtrcowHostListadd.L.Lock()
	ret = recorderAuxMockPtrcowHostListadd
	condRecorderAuxMockPtrcowHostListadd.L.Unlock()
	return
}

// (recvc *cowHostList)add - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *cowHostList) add(arghost *HostInfo) (reta bool) {
	FuncAuxMockPtrcowHostListadd, ok := apomock.GetRegisteredFunc("gocql.cowHostList.add")
	if ok {
		reta = FuncAuxMockPtrcowHostListadd.(func(recvc *cowHostList, arghost *HostInfo) (reta bool))(recvc, arghost)
	} else {
		panic("FuncAuxMockPtrcowHostListadd ")
	}
	AuxMockIncrementRecorderAuxMockPtrcowHostListadd()
	return
}

//
// Mock: RoundRobinHostPolicy()(reta HostSelectionPolicy)
//

type MockArgsTypeRoundRobinHostPolicy struct {
	ApomockCallNumber int
}

var LastMockArgsRoundRobinHostPolicy MockArgsTypeRoundRobinHostPolicy

// AuxMockRoundRobinHostPolicy()(reta HostSelectionPolicy) - Generated mock function
func AuxMockRoundRobinHostPolicy() (reta HostSelectionPolicy) {
	rargs, rerr := apomock.GetNext("gocql.RoundRobinHostPolicy")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.RoundRobinHostPolicy")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.RoundRobinHostPolicy")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(HostSelectionPolicy)
	}
	return
}

// RecorderAuxMockRoundRobinHostPolicy  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockRoundRobinHostPolicy int = 0

var condRecorderAuxMockRoundRobinHostPolicy *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockRoundRobinHostPolicy(i int) {
	condRecorderAuxMockRoundRobinHostPolicy.L.Lock()
	for recorderAuxMockRoundRobinHostPolicy < i {
		condRecorderAuxMockRoundRobinHostPolicy.Wait()
	}
	condRecorderAuxMockRoundRobinHostPolicy.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockRoundRobinHostPolicy() {
	condRecorderAuxMockRoundRobinHostPolicy.L.Lock()
	recorderAuxMockRoundRobinHostPolicy++
	condRecorderAuxMockRoundRobinHostPolicy.L.Unlock()
	condRecorderAuxMockRoundRobinHostPolicy.Broadcast()
}
func AuxMockGetRecorderAuxMockRoundRobinHostPolicy() (ret int) {
	condRecorderAuxMockRoundRobinHostPolicy.L.Lock()
	ret = recorderAuxMockRoundRobinHostPolicy
	condRecorderAuxMockRoundRobinHostPolicy.L.Unlock()
	return
}

// RoundRobinHostPolicy - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func RoundRobinHostPolicy() (reta HostSelectionPolicy) {
	FuncAuxMockRoundRobinHostPolicy, ok := apomock.GetRegisteredFunc("gocql.RoundRobinHostPolicy")
	if ok {
		reta = FuncAuxMockRoundRobinHostPolicy.(func() (reta HostSelectionPolicy))()
	} else {
		panic("FuncAuxMockRoundRobinHostPolicy ")
	}
	AuxMockIncrementRecorderAuxMockRoundRobinHostPolicy()
	return
}

//
// Mock: (recvr *roundRobinHostPolicy)SetPartitioner(argpartitioner string)()
//

type MockArgsTyperoundRobinHostPolicySetPartitioner struct {
	ApomockCallNumber int
	Argpartitioner    string
}

var LastMockArgsroundRobinHostPolicySetPartitioner MockArgsTyperoundRobinHostPolicySetPartitioner

// (recvr *roundRobinHostPolicy)AuxMockSetPartitioner(argpartitioner string)() - Generated mock function
func (recvr *roundRobinHostPolicy) AuxMockSetPartitioner(argpartitioner string) {
	LastMockArgsroundRobinHostPolicySetPartitioner = MockArgsTyperoundRobinHostPolicySetPartitioner{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrroundRobinHostPolicySetPartitioner(),
		Argpartitioner:    argpartitioner,
	}
	return
}

// RecorderAuxMockPtrroundRobinHostPolicySetPartitioner  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrroundRobinHostPolicySetPartitioner int = 0

var condRecorderAuxMockPtrroundRobinHostPolicySetPartitioner *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrroundRobinHostPolicySetPartitioner(i int) {
	condRecorderAuxMockPtrroundRobinHostPolicySetPartitioner.L.Lock()
	for recorderAuxMockPtrroundRobinHostPolicySetPartitioner < i {
		condRecorderAuxMockPtrroundRobinHostPolicySetPartitioner.Wait()
	}
	condRecorderAuxMockPtrroundRobinHostPolicySetPartitioner.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicySetPartitioner() {
	condRecorderAuxMockPtrroundRobinHostPolicySetPartitioner.L.Lock()
	recorderAuxMockPtrroundRobinHostPolicySetPartitioner++
	condRecorderAuxMockPtrroundRobinHostPolicySetPartitioner.L.Unlock()
	condRecorderAuxMockPtrroundRobinHostPolicySetPartitioner.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrroundRobinHostPolicySetPartitioner() (ret int) {
	condRecorderAuxMockPtrroundRobinHostPolicySetPartitioner.L.Lock()
	ret = recorderAuxMockPtrroundRobinHostPolicySetPartitioner
	condRecorderAuxMockPtrroundRobinHostPolicySetPartitioner.L.Unlock()
	return
}

// (recvr *roundRobinHostPolicy)SetPartitioner - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *roundRobinHostPolicy) SetPartitioner(argpartitioner string) {
	FuncAuxMockPtrroundRobinHostPolicySetPartitioner, ok := apomock.GetRegisteredFunc("gocql.roundRobinHostPolicy.SetPartitioner")
	if ok {
		FuncAuxMockPtrroundRobinHostPolicySetPartitioner.(func(recvr *roundRobinHostPolicy, argpartitioner string))(recvr, argpartitioner)
	} else {
		panic("FuncAuxMockPtrroundRobinHostPolicySetPartitioner ")
	}
	AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicySetPartitioner()
	return
}

//
// Mock: (recvr *roundRobinHostPolicy)HostDown(argaddr string)()
//

type MockArgsTyperoundRobinHostPolicyHostDown struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgsroundRobinHostPolicyHostDown MockArgsTyperoundRobinHostPolicyHostDown

// (recvr *roundRobinHostPolicy)AuxMockHostDown(argaddr string)() - Generated mock function
func (recvr *roundRobinHostPolicy) AuxMockHostDown(argaddr string) {
	LastMockArgsroundRobinHostPolicyHostDown = MockArgsTyperoundRobinHostPolicyHostDown{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrroundRobinHostPolicyHostDown(),
		Argaddr:           argaddr,
	}
	return
}

// RecorderAuxMockPtrroundRobinHostPolicyHostDown  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrroundRobinHostPolicyHostDown int = 0

var condRecorderAuxMockPtrroundRobinHostPolicyHostDown *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrroundRobinHostPolicyHostDown(i int) {
	condRecorderAuxMockPtrroundRobinHostPolicyHostDown.L.Lock()
	for recorderAuxMockPtrroundRobinHostPolicyHostDown < i {
		condRecorderAuxMockPtrroundRobinHostPolicyHostDown.Wait()
	}
	condRecorderAuxMockPtrroundRobinHostPolicyHostDown.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicyHostDown() {
	condRecorderAuxMockPtrroundRobinHostPolicyHostDown.L.Lock()
	recorderAuxMockPtrroundRobinHostPolicyHostDown++
	condRecorderAuxMockPtrroundRobinHostPolicyHostDown.L.Unlock()
	condRecorderAuxMockPtrroundRobinHostPolicyHostDown.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrroundRobinHostPolicyHostDown() (ret int) {
	condRecorderAuxMockPtrroundRobinHostPolicyHostDown.L.Lock()
	ret = recorderAuxMockPtrroundRobinHostPolicyHostDown
	condRecorderAuxMockPtrroundRobinHostPolicyHostDown.L.Unlock()
	return
}

// (recvr *roundRobinHostPolicy)HostDown - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *roundRobinHostPolicy) HostDown(argaddr string) {
	FuncAuxMockPtrroundRobinHostPolicyHostDown, ok := apomock.GetRegisteredFunc("gocql.roundRobinHostPolicy.HostDown")
	if ok {
		FuncAuxMockPtrroundRobinHostPolicyHostDown.(func(recvr *roundRobinHostPolicy, argaddr string))(recvr, argaddr)
	} else {
		panic("FuncAuxMockPtrroundRobinHostPolicyHostDown ")
	}
	AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicyHostDown()
	return
}

//
// Mock: (recvt *tokenAwareHostPolicy)HostDown(argaddr string)()
//

type MockArgsTypetokenAwareHostPolicyHostDown struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgstokenAwareHostPolicyHostDown MockArgsTypetokenAwareHostPolicyHostDown

// (recvt *tokenAwareHostPolicy)AuxMockHostDown(argaddr string)() - Generated mock function
func (recvt *tokenAwareHostPolicy) AuxMockHostDown(argaddr string) {
	LastMockArgstokenAwareHostPolicyHostDown = MockArgsTypetokenAwareHostPolicyHostDown{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicyHostDown(),
		Argaddr:           argaddr,
	}
	return
}

// RecorderAuxMockPtrtokenAwareHostPolicyHostDown  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenAwareHostPolicyHostDown int = 0

var condRecorderAuxMockPtrtokenAwareHostPolicyHostDown *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenAwareHostPolicyHostDown(i int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyHostDown.L.Lock()
	for recorderAuxMockPtrtokenAwareHostPolicyHostDown < i {
		condRecorderAuxMockPtrtokenAwareHostPolicyHostDown.Wait()
	}
	condRecorderAuxMockPtrtokenAwareHostPolicyHostDown.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyHostDown() {
	condRecorderAuxMockPtrtokenAwareHostPolicyHostDown.L.Lock()
	recorderAuxMockPtrtokenAwareHostPolicyHostDown++
	condRecorderAuxMockPtrtokenAwareHostPolicyHostDown.L.Unlock()
	condRecorderAuxMockPtrtokenAwareHostPolicyHostDown.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicyHostDown() (ret int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyHostDown.L.Lock()
	ret = recorderAuxMockPtrtokenAwareHostPolicyHostDown
	condRecorderAuxMockPtrtokenAwareHostPolicyHostDown.L.Unlock()
	return
}

// (recvt *tokenAwareHostPolicy)HostDown - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenAwareHostPolicy) HostDown(argaddr string) {
	FuncAuxMockPtrtokenAwareHostPolicyHostDown, ok := apomock.GetRegisteredFunc("gocql.tokenAwareHostPolicy.HostDown")
	if ok {
		FuncAuxMockPtrtokenAwareHostPolicyHostDown.(func(recvt *tokenAwareHostPolicy, argaddr string))(recvt, argaddr)
	} else {
		panic("FuncAuxMockPtrtokenAwareHostPolicyHostDown ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyHostDown()
	return
}

//
// Mock: (recvt *tokenAwareHostPolicy)resetTokenRing()()
//

type MockArgsTypetokenAwareHostPolicyresetTokenRing struct {
	ApomockCallNumber int
}

var LastMockArgstokenAwareHostPolicyresetTokenRing MockArgsTypetokenAwareHostPolicyresetTokenRing

// (recvt *tokenAwareHostPolicy)AuxMockresetTokenRing()() - Generated mock function
func (recvt *tokenAwareHostPolicy) AuxMockresetTokenRing() {
	return
}

// RecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenAwareHostPolicyresetTokenRing int = 0

var condRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing(i int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing.L.Lock()
	for recorderAuxMockPtrtokenAwareHostPolicyresetTokenRing < i {
		condRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing.Wait()
	}
	condRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing() {
	condRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing.L.Lock()
	recorderAuxMockPtrtokenAwareHostPolicyresetTokenRing++
	condRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing.L.Unlock()
	condRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing() (ret int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing.L.Lock()
	ret = recorderAuxMockPtrtokenAwareHostPolicyresetTokenRing
	condRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing.L.Unlock()
	return
}

// (recvt *tokenAwareHostPolicy)resetTokenRing - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenAwareHostPolicy) resetTokenRing() {
	FuncAuxMockPtrtokenAwareHostPolicyresetTokenRing, ok := apomock.GetRegisteredFunc("gocql.tokenAwareHostPolicy.resetTokenRing")
	if ok {
		FuncAuxMockPtrtokenAwareHostPolicyresetTokenRing.(func(recvt *tokenAwareHostPolicy))(recvt)
	} else {
		panic("FuncAuxMockPtrtokenAwareHostPolicyresetTokenRing ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyresetTokenRing()
	return
}

//
// Mock: (recvc *cowHostList)remove(argaddr string)(reta bool)
//

type MockArgsTypecowHostListremove struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgscowHostListremove MockArgsTypecowHostListremove

// (recvc *cowHostList)AuxMockremove(argaddr string)(reta bool) - Generated mock function
func (recvc *cowHostList) AuxMockremove(argaddr string) (reta bool) {
	LastMockArgscowHostListremove = MockArgsTypecowHostListremove{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcowHostListremove(),
		Argaddr:           argaddr,
	}
	rargs, rerr := apomock.GetNext("gocql.cowHostList.remove")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.cowHostList.remove")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.cowHostList.remove")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrcowHostListremove  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcowHostListremove int = 0

var condRecorderAuxMockPtrcowHostListremove *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcowHostListremove(i int) {
	condRecorderAuxMockPtrcowHostListremove.L.Lock()
	for recorderAuxMockPtrcowHostListremove < i {
		condRecorderAuxMockPtrcowHostListremove.Wait()
	}
	condRecorderAuxMockPtrcowHostListremove.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcowHostListremove() {
	condRecorderAuxMockPtrcowHostListremove.L.Lock()
	recorderAuxMockPtrcowHostListremove++
	condRecorderAuxMockPtrcowHostListremove.L.Unlock()
	condRecorderAuxMockPtrcowHostListremove.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcowHostListremove() (ret int) {
	condRecorderAuxMockPtrcowHostListremove.L.Lock()
	ret = recorderAuxMockPtrcowHostListremove
	condRecorderAuxMockPtrcowHostListremove.L.Unlock()
	return
}

// (recvc *cowHostList)remove - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *cowHostList) remove(argaddr string) (reta bool) {
	FuncAuxMockPtrcowHostListremove, ok := apomock.GetRegisteredFunc("gocql.cowHostList.remove")
	if ok {
		reta = FuncAuxMockPtrcowHostListremove.(func(recvc *cowHostList, argaddr string) (reta bool))(recvc, argaddr)
	} else {
		panic("FuncAuxMockPtrcowHostListremove ")
	}
	AuxMockIncrementRecorderAuxMockPtrcowHostListremove()
	return
}

//
// Mock: (recvs *SimpleRetryPolicy)Attempt(argq RetryableQuery)(reta bool)
//

type MockArgsTypeSimpleRetryPolicyAttempt struct {
	ApomockCallNumber int
	Argq              RetryableQuery
}

var LastMockArgsSimpleRetryPolicyAttempt MockArgsTypeSimpleRetryPolicyAttempt

// (recvs *SimpleRetryPolicy)AuxMockAttempt(argq RetryableQuery)(reta bool) - Generated mock function
func (recvs *SimpleRetryPolicy) AuxMockAttempt(argq RetryableQuery) (reta bool) {
	LastMockArgsSimpleRetryPolicyAttempt = MockArgsTypeSimpleRetryPolicyAttempt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSimpleRetryPolicyAttempt(),
		Argq:              argq,
	}
	rargs, rerr := apomock.GetNext("gocql.SimpleRetryPolicy.Attempt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.SimpleRetryPolicy.Attempt")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.SimpleRetryPolicy.Attempt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrSimpleRetryPolicyAttempt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSimpleRetryPolicyAttempt int = 0

var condRecorderAuxMockPtrSimpleRetryPolicyAttempt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSimpleRetryPolicyAttempt(i int) {
	condRecorderAuxMockPtrSimpleRetryPolicyAttempt.L.Lock()
	for recorderAuxMockPtrSimpleRetryPolicyAttempt < i {
		condRecorderAuxMockPtrSimpleRetryPolicyAttempt.Wait()
	}
	condRecorderAuxMockPtrSimpleRetryPolicyAttempt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSimpleRetryPolicyAttempt() {
	condRecorderAuxMockPtrSimpleRetryPolicyAttempt.L.Lock()
	recorderAuxMockPtrSimpleRetryPolicyAttempt++
	condRecorderAuxMockPtrSimpleRetryPolicyAttempt.L.Unlock()
	condRecorderAuxMockPtrSimpleRetryPolicyAttempt.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSimpleRetryPolicyAttempt() (ret int) {
	condRecorderAuxMockPtrSimpleRetryPolicyAttempt.L.Lock()
	ret = recorderAuxMockPtrSimpleRetryPolicyAttempt
	condRecorderAuxMockPtrSimpleRetryPolicyAttempt.L.Unlock()
	return
}

// (recvs *SimpleRetryPolicy)Attempt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *SimpleRetryPolicy) Attempt(argq RetryableQuery) (reta bool) {
	FuncAuxMockPtrSimpleRetryPolicyAttempt, ok := apomock.GetRegisteredFunc("gocql.SimpleRetryPolicy.Attempt")
	if ok {
		reta = FuncAuxMockPtrSimpleRetryPolicyAttempt.(func(recvs *SimpleRetryPolicy, argq RetryableQuery) (reta bool))(recvs, argq)
	} else {
		panic("FuncAuxMockPtrSimpleRetryPolicyAttempt ")
	}
	AuxMockIncrementRecorderAuxMockPtrSimpleRetryPolicyAttempt()
	return
}

//
// Mock: (recvr *roundRobinHostPolicy)Pick(argqry ExecutableQuery)(reta NextHost)
//

type MockArgsTyperoundRobinHostPolicyPick struct {
	ApomockCallNumber int
	Argqry            ExecutableQuery
}

var LastMockArgsroundRobinHostPolicyPick MockArgsTyperoundRobinHostPolicyPick

// (recvr *roundRobinHostPolicy)AuxMockPick(argqry ExecutableQuery)(reta NextHost) - Generated mock function
func (recvr *roundRobinHostPolicy) AuxMockPick(argqry ExecutableQuery) (reta NextHost) {
	LastMockArgsroundRobinHostPolicyPick = MockArgsTyperoundRobinHostPolicyPick{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrroundRobinHostPolicyPick(),
		Argqry:            argqry,
	}
	rargs, rerr := apomock.GetNext("gocql.roundRobinHostPolicy.Pick")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.roundRobinHostPolicy.Pick")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.roundRobinHostPolicy.Pick")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(NextHost)
	}
	return
}

// RecorderAuxMockPtrroundRobinHostPolicyPick  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrroundRobinHostPolicyPick int = 0

var condRecorderAuxMockPtrroundRobinHostPolicyPick *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrroundRobinHostPolicyPick(i int) {
	condRecorderAuxMockPtrroundRobinHostPolicyPick.L.Lock()
	for recorderAuxMockPtrroundRobinHostPolicyPick < i {
		condRecorderAuxMockPtrroundRobinHostPolicyPick.Wait()
	}
	condRecorderAuxMockPtrroundRobinHostPolicyPick.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicyPick() {
	condRecorderAuxMockPtrroundRobinHostPolicyPick.L.Lock()
	recorderAuxMockPtrroundRobinHostPolicyPick++
	condRecorderAuxMockPtrroundRobinHostPolicyPick.L.Unlock()
	condRecorderAuxMockPtrroundRobinHostPolicyPick.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrroundRobinHostPolicyPick() (ret int) {
	condRecorderAuxMockPtrroundRobinHostPolicyPick.L.Lock()
	ret = recorderAuxMockPtrroundRobinHostPolicyPick
	condRecorderAuxMockPtrroundRobinHostPolicyPick.L.Unlock()
	return
}

// (recvr *roundRobinHostPolicy)Pick - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *roundRobinHostPolicy) Pick(argqry ExecutableQuery) (reta NextHost) {
	FuncAuxMockPtrroundRobinHostPolicyPick, ok := apomock.GetRegisteredFunc("gocql.roundRobinHostPolicy.Pick")
	if ok {
		reta = FuncAuxMockPtrroundRobinHostPolicyPick.(func(recvr *roundRobinHostPolicy, argqry ExecutableQuery) (reta NextHost))(recvr, argqry)
	} else {
		panic("FuncAuxMockPtrroundRobinHostPolicyPick ")
	}
	AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicyPick()
	return
}

//
// Mock: (recvr *roundRobinHostPolicy)HostUp(arghost *HostInfo)()
//

type MockArgsTyperoundRobinHostPolicyHostUp struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgsroundRobinHostPolicyHostUp MockArgsTyperoundRobinHostPolicyHostUp

// (recvr *roundRobinHostPolicy)AuxMockHostUp(arghost *HostInfo)() - Generated mock function
func (recvr *roundRobinHostPolicy) AuxMockHostUp(arghost *HostInfo) {
	LastMockArgsroundRobinHostPolicyHostUp = MockArgsTyperoundRobinHostPolicyHostUp{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrroundRobinHostPolicyHostUp(),
		Arghost:           arghost,
	}
	return
}

// RecorderAuxMockPtrroundRobinHostPolicyHostUp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrroundRobinHostPolicyHostUp int = 0

var condRecorderAuxMockPtrroundRobinHostPolicyHostUp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrroundRobinHostPolicyHostUp(i int) {
	condRecorderAuxMockPtrroundRobinHostPolicyHostUp.L.Lock()
	for recorderAuxMockPtrroundRobinHostPolicyHostUp < i {
		condRecorderAuxMockPtrroundRobinHostPolicyHostUp.Wait()
	}
	condRecorderAuxMockPtrroundRobinHostPolicyHostUp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicyHostUp() {
	condRecorderAuxMockPtrroundRobinHostPolicyHostUp.L.Lock()
	recorderAuxMockPtrroundRobinHostPolicyHostUp++
	condRecorderAuxMockPtrroundRobinHostPolicyHostUp.L.Unlock()
	condRecorderAuxMockPtrroundRobinHostPolicyHostUp.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrroundRobinHostPolicyHostUp() (ret int) {
	condRecorderAuxMockPtrroundRobinHostPolicyHostUp.L.Lock()
	ret = recorderAuxMockPtrroundRobinHostPolicyHostUp
	condRecorderAuxMockPtrroundRobinHostPolicyHostUp.L.Unlock()
	return
}

// (recvr *roundRobinHostPolicy)HostUp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *roundRobinHostPolicy) HostUp(arghost *HostInfo) {
	FuncAuxMockPtrroundRobinHostPolicyHostUp, ok := apomock.GetRegisteredFunc("gocql.roundRobinHostPolicy.HostUp")
	if ok {
		FuncAuxMockPtrroundRobinHostPolicyHostUp.(func(recvr *roundRobinHostPolicy, arghost *HostInfo))(recvr, arghost)
	} else {
		panic("FuncAuxMockPtrroundRobinHostPolicyHostUp ")
	}
	AuxMockIncrementRecorderAuxMockPtrroundRobinHostPolicyHostUp()
	return
}

//
// Mock: (recvt *tokenAwareHostPolicy)SetPartitioner(argpartitioner string)()
//

type MockArgsTypetokenAwareHostPolicySetPartitioner struct {
	ApomockCallNumber int
	Argpartitioner    string
}

var LastMockArgstokenAwareHostPolicySetPartitioner MockArgsTypetokenAwareHostPolicySetPartitioner

// (recvt *tokenAwareHostPolicy)AuxMockSetPartitioner(argpartitioner string)() - Generated mock function
func (recvt *tokenAwareHostPolicy) AuxMockSetPartitioner(argpartitioner string) {
	LastMockArgstokenAwareHostPolicySetPartitioner = MockArgsTypetokenAwareHostPolicySetPartitioner{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner(),
		Argpartitioner:    argpartitioner,
	}
	return
}

// RecorderAuxMockPtrtokenAwareHostPolicySetPartitioner  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenAwareHostPolicySetPartitioner int = 0

var condRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner(i int) {
	condRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner.L.Lock()
	for recorderAuxMockPtrtokenAwareHostPolicySetPartitioner < i {
		condRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner.Wait()
	}
	condRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner() {
	condRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner.L.Lock()
	recorderAuxMockPtrtokenAwareHostPolicySetPartitioner++
	condRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner.L.Unlock()
	condRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner() (ret int) {
	condRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner.L.Lock()
	ret = recorderAuxMockPtrtokenAwareHostPolicySetPartitioner
	condRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner.L.Unlock()
	return
}

// (recvt *tokenAwareHostPolicy)SetPartitioner - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenAwareHostPolicy) SetPartitioner(argpartitioner string) {
	FuncAuxMockPtrtokenAwareHostPolicySetPartitioner, ok := apomock.GetRegisteredFunc("gocql.tokenAwareHostPolicy.SetPartitioner")
	if ok {
		FuncAuxMockPtrtokenAwareHostPolicySetPartitioner.(func(recvt *tokenAwareHostPolicy, argpartitioner string))(recvt, argpartitioner)
	} else {
		panic("FuncAuxMockPtrtokenAwareHostPolicySetPartitioner ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicySetPartitioner()
	return
}

//
// Mock: (recvr *hostPoolHostPolicy)SetHosts(arghosts []*HostInfo)()
//

type MockArgsTypehostPoolHostPolicySetHosts struct {
	ApomockCallNumber int
	Arghosts          []*HostInfo
}

var LastMockArgshostPoolHostPolicySetHosts MockArgsTypehostPoolHostPolicySetHosts

// (recvr *hostPoolHostPolicy)AuxMockSetHosts(arghosts []*HostInfo)() - Generated mock function
func (recvr *hostPoolHostPolicy) AuxMockSetHosts(arghosts []*HostInfo) {
	LastMockArgshostPoolHostPolicySetHosts = MockArgsTypehostPoolHostPolicySetHosts{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrhostPoolHostPolicySetHosts(),
		Arghosts:          arghosts,
	}
	return
}

// RecorderAuxMockPtrhostPoolHostPolicySetHosts  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostPoolHostPolicySetHosts int = 0

var condRecorderAuxMockPtrhostPoolHostPolicySetHosts *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostPoolHostPolicySetHosts(i int) {
	condRecorderAuxMockPtrhostPoolHostPolicySetHosts.L.Lock()
	for recorderAuxMockPtrhostPoolHostPolicySetHosts < i {
		condRecorderAuxMockPtrhostPoolHostPolicySetHosts.Wait()
	}
	condRecorderAuxMockPtrhostPoolHostPolicySetHosts.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicySetHosts() {
	condRecorderAuxMockPtrhostPoolHostPolicySetHosts.L.Lock()
	recorderAuxMockPtrhostPoolHostPolicySetHosts++
	condRecorderAuxMockPtrhostPoolHostPolicySetHosts.L.Unlock()
	condRecorderAuxMockPtrhostPoolHostPolicySetHosts.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostPoolHostPolicySetHosts() (ret int) {
	condRecorderAuxMockPtrhostPoolHostPolicySetHosts.L.Lock()
	ret = recorderAuxMockPtrhostPoolHostPolicySetHosts
	condRecorderAuxMockPtrhostPoolHostPolicySetHosts.L.Unlock()
	return
}

// (recvr *hostPoolHostPolicy)SetHosts - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *hostPoolHostPolicy) SetHosts(arghosts []*HostInfo) {
	FuncAuxMockPtrhostPoolHostPolicySetHosts, ok := apomock.GetRegisteredFunc("gocql.hostPoolHostPolicy.SetHosts")
	if ok {
		FuncAuxMockPtrhostPoolHostPolicySetHosts.(func(recvr *hostPoolHostPolicy, arghosts []*HostInfo))(recvr, arghosts)
	} else {
		panic("FuncAuxMockPtrhostPoolHostPolicySetHosts ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicySetHosts()
	return
}

//
// Mock: TokenAwareHostPolicy(argfallback HostSelectionPolicy)(reta HostSelectionPolicy)
//

type MockArgsTypeTokenAwareHostPolicy struct {
	ApomockCallNumber int
	Argfallback       HostSelectionPolicy
}

var LastMockArgsTokenAwareHostPolicy MockArgsTypeTokenAwareHostPolicy

// AuxMockTokenAwareHostPolicy(argfallback HostSelectionPolicy)(reta HostSelectionPolicy) - Generated mock function
func AuxMockTokenAwareHostPolicy(argfallback HostSelectionPolicy) (reta HostSelectionPolicy) {
	LastMockArgsTokenAwareHostPolicy = MockArgsTypeTokenAwareHostPolicy{
		ApomockCallNumber: AuxMockGetRecorderAuxMockTokenAwareHostPolicy(),
		Argfallback:       argfallback,
	}
	rargs, rerr := apomock.GetNext("gocql.TokenAwareHostPolicy")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.TokenAwareHostPolicy")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.TokenAwareHostPolicy")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(HostSelectionPolicy)
	}
	return
}

// RecorderAuxMockTokenAwareHostPolicy  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockTokenAwareHostPolicy int = 0

var condRecorderAuxMockTokenAwareHostPolicy *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockTokenAwareHostPolicy(i int) {
	condRecorderAuxMockTokenAwareHostPolicy.L.Lock()
	for recorderAuxMockTokenAwareHostPolicy < i {
		condRecorderAuxMockTokenAwareHostPolicy.Wait()
	}
	condRecorderAuxMockTokenAwareHostPolicy.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockTokenAwareHostPolicy() {
	condRecorderAuxMockTokenAwareHostPolicy.L.Lock()
	recorderAuxMockTokenAwareHostPolicy++
	condRecorderAuxMockTokenAwareHostPolicy.L.Unlock()
	condRecorderAuxMockTokenAwareHostPolicy.Broadcast()
}
func AuxMockGetRecorderAuxMockTokenAwareHostPolicy() (ret int) {
	condRecorderAuxMockTokenAwareHostPolicy.L.Lock()
	ret = recorderAuxMockTokenAwareHostPolicy
	condRecorderAuxMockTokenAwareHostPolicy.L.Unlock()
	return
}

// TokenAwareHostPolicy - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func TokenAwareHostPolicy(argfallback HostSelectionPolicy) (reta HostSelectionPolicy) {
	FuncAuxMockTokenAwareHostPolicy, ok := apomock.GetRegisteredFunc("gocql.TokenAwareHostPolicy")
	if ok {
		reta = FuncAuxMockTokenAwareHostPolicy.(func(argfallback HostSelectionPolicy) (reta HostSelectionPolicy))(argfallback)
	} else {
		panic("FuncAuxMockTokenAwareHostPolicy ")
	}
	AuxMockIncrementRecorderAuxMockTokenAwareHostPolicy()
	return
}

//
// Mock: HostPoolHostPolicy(arghp hostpool.HostPool)(reta HostSelectionPolicy)
//

type MockArgsTypeHostPoolHostPolicy struct {
	ApomockCallNumber int
	Arghp             hostpool.HostPool
}

var LastMockArgsHostPoolHostPolicy MockArgsTypeHostPoolHostPolicy

// AuxMockHostPoolHostPolicy(arghp hostpool.HostPool)(reta HostSelectionPolicy) - Generated mock function
func AuxMockHostPoolHostPolicy(arghp hostpool.HostPool) (reta HostSelectionPolicy) {
	LastMockArgsHostPoolHostPolicy = MockArgsTypeHostPoolHostPolicy{
		ApomockCallNumber: AuxMockGetRecorderAuxMockHostPoolHostPolicy(),
		Arghp:             arghp,
	}
	rargs, rerr := apomock.GetNext("gocql.HostPoolHostPolicy")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostPoolHostPolicy")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostPoolHostPolicy")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(HostSelectionPolicy)
	}
	return
}

// RecorderAuxMockHostPoolHostPolicy  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockHostPoolHostPolicy int = 0

var condRecorderAuxMockHostPoolHostPolicy *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockHostPoolHostPolicy(i int) {
	condRecorderAuxMockHostPoolHostPolicy.L.Lock()
	for recorderAuxMockHostPoolHostPolicy < i {
		condRecorderAuxMockHostPoolHostPolicy.Wait()
	}
	condRecorderAuxMockHostPoolHostPolicy.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockHostPoolHostPolicy() {
	condRecorderAuxMockHostPoolHostPolicy.L.Lock()
	recorderAuxMockHostPoolHostPolicy++
	condRecorderAuxMockHostPoolHostPolicy.L.Unlock()
	condRecorderAuxMockHostPoolHostPolicy.Broadcast()
}
func AuxMockGetRecorderAuxMockHostPoolHostPolicy() (ret int) {
	condRecorderAuxMockHostPoolHostPolicy.L.Lock()
	ret = recorderAuxMockHostPoolHostPolicy
	condRecorderAuxMockHostPoolHostPolicy.L.Unlock()
	return
}

// HostPoolHostPolicy - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func HostPoolHostPolicy(arghp hostpool.HostPool) (reta HostSelectionPolicy) {
	FuncAuxMockHostPoolHostPolicy, ok := apomock.GetRegisteredFunc("gocql.HostPoolHostPolicy")
	if ok {
		reta = FuncAuxMockHostPoolHostPolicy.(func(arghp hostpool.HostPool) (reta HostSelectionPolicy))(arghp)
	} else {
		panic("FuncAuxMockHostPoolHostPolicy ")
	}
	AuxMockIncrementRecorderAuxMockHostPoolHostPolicy()
	return
}

//
// Mock: (recvr *hostPoolHostPolicy)SetPartitioner(argpartitioner string)()
//

type MockArgsTypehostPoolHostPolicySetPartitioner struct {
	ApomockCallNumber int
	Argpartitioner    string
}

var LastMockArgshostPoolHostPolicySetPartitioner MockArgsTypehostPoolHostPolicySetPartitioner

// (recvr *hostPoolHostPolicy)AuxMockSetPartitioner(argpartitioner string)() - Generated mock function
func (recvr *hostPoolHostPolicy) AuxMockSetPartitioner(argpartitioner string) {
	LastMockArgshostPoolHostPolicySetPartitioner = MockArgsTypehostPoolHostPolicySetPartitioner{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrhostPoolHostPolicySetPartitioner(),
		Argpartitioner:    argpartitioner,
	}
	return
}

// RecorderAuxMockPtrhostPoolHostPolicySetPartitioner  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostPoolHostPolicySetPartitioner int = 0

var condRecorderAuxMockPtrhostPoolHostPolicySetPartitioner *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostPoolHostPolicySetPartitioner(i int) {
	condRecorderAuxMockPtrhostPoolHostPolicySetPartitioner.L.Lock()
	for recorderAuxMockPtrhostPoolHostPolicySetPartitioner < i {
		condRecorderAuxMockPtrhostPoolHostPolicySetPartitioner.Wait()
	}
	condRecorderAuxMockPtrhostPoolHostPolicySetPartitioner.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicySetPartitioner() {
	condRecorderAuxMockPtrhostPoolHostPolicySetPartitioner.L.Lock()
	recorderAuxMockPtrhostPoolHostPolicySetPartitioner++
	condRecorderAuxMockPtrhostPoolHostPolicySetPartitioner.L.Unlock()
	condRecorderAuxMockPtrhostPoolHostPolicySetPartitioner.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostPoolHostPolicySetPartitioner() (ret int) {
	condRecorderAuxMockPtrhostPoolHostPolicySetPartitioner.L.Lock()
	ret = recorderAuxMockPtrhostPoolHostPolicySetPartitioner
	condRecorderAuxMockPtrhostPoolHostPolicySetPartitioner.L.Unlock()
	return
}

// (recvr *hostPoolHostPolicy)SetPartitioner - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *hostPoolHostPolicy) SetPartitioner(argpartitioner string) {
	FuncAuxMockPtrhostPoolHostPolicySetPartitioner, ok := apomock.GetRegisteredFunc("gocql.hostPoolHostPolicy.SetPartitioner")
	if ok {
		FuncAuxMockPtrhostPoolHostPolicySetPartitioner.(func(recvr *hostPoolHostPolicy, argpartitioner string))(recvr, argpartitioner)
	} else {
		panic("FuncAuxMockPtrhostPoolHostPolicySetPartitioner ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicySetPartitioner()
	return
}

//
// Mock: (recvhost selectedHostPoolHost)Mark(argerr error)()
//

type MockArgsTypeselectedHostPoolHostMark struct {
	ApomockCallNumber int
	Argerr            error
}

var LastMockArgsselectedHostPoolHostMark MockArgsTypeselectedHostPoolHostMark

// (recvhost selectedHostPoolHost)AuxMockMark(argerr error)() - Generated mock function
func (recvhost selectedHostPoolHost) AuxMockMark(argerr error) {
	LastMockArgsselectedHostPoolHostMark = MockArgsTypeselectedHostPoolHostMark{
		ApomockCallNumber: AuxMockGetRecorderAuxMockselectedHostPoolHostMark(),
		Argerr:            argerr,
	}
	return
}

// RecorderAuxMockselectedHostPoolHostMark  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockselectedHostPoolHostMark int = 0

var condRecorderAuxMockselectedHostPoolHostMark *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockselectedHostPoolHostMark(i int) {
	condRecorderAuxMockselectedHostPoolHostMark.L.Lock()
	for recorderAuxMockselectedHostPoolHostMark < i {
		condRecorderAuxMockselectedHostPoolHostMark.Wait()
	}
	condRecorderAuxMockselectedHostPoolHostMark.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockselectedHostPoolHostMark() {
	condRecorderAuxMockselectedHostPoolHostMark.L.Lock()
	recorderAuxMockselectedHostPoolHostMark++
	condRecorderAuxMockselectedHostPoolHostMark.L.Unlock()
	condRecorderAuxMockselectedHostPoolHostMark.Broadcast()
}
func AuxMockGetRecorderAuxMockselectedHostPoolHostMark() (ret int) {
	condRecorderAuxMockselectedHostPoolHostMark.L.Lock()
	ret = recorderAuxMockselectedHostPoolHostMark
	condRecorderAuxMockselectedHostPoolHostMark.L.Unlock()
	return
}

// (recvhost selectedHostPoolHost)Mark - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvhost selectedHostPoolHost) Mark(argerr error) {
	FuncAuxMockselectedHostPoolHostMark, ok := apomock.GetRegisteredFunc("gocql.selectedHostPoolHost.Mark")
	if ok {
		FuncAuxMockselectedHostPoolHostMark.(func(recvhost selectedHostPoolHost, argerr error))(recvhost, argerr)
	} else {
		panic("FuncAuxMockselectedHostPoolHostMark ")
	}
	AuxMockIncrementRecorderAuxMockselectedHostPoolHostMark()
	return
}

//
// Mock: (recvhost *selectedHost)Info()(reta *HostInfo)
//

type MockArgsTypeselectedHostInfo struct {
	ApomockCallNumber int
}

var LastMockArgsselectedHostInfo MockArgsTypeselectedHostInfo

// (recvhost *selectedHost)AuxMockInfo()(reta *HostInfo) - Generated mock function
func (recvhost *selectedHost) AuxMockInfo() (reta *HostInfo) {
	rargs, rerr := apomock.GetNext("gocql.selectedHost.Info")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.selectedHost.Info")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.selectedHost.Info")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrselectedHostInfo  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrselectedHostInfo int = 0

var condRecorderAuxMockPtrselectedHostInfo *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrselectedHostInfo(i int) {
	condRecorderAuxMockPtrselectedHostInfo.L.Lock()
	for recorderAuxMockPtrselectedHostInfo < i {
		condRecorderAuxMockPtrselectedHostInfo.Wait()
	}
	condRecorderAuxMockPtrselectedHostInfo.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrselectedHostInfo() {
	condRecorderAuxMockPtrselectedHostInfo.L.Lock()
	recorderAuxMockPtrselectedHostInfo++
	condRecorderAuxMockPtrselectedHostInfo.L.Unlock()
	condRecorderAuxMockPtrselectedHostInfo.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrselectedHostInfo() (ret int) {
	condRecorderAuxMockPtrselectedHostInfo.L.Lock()
	ret = recorderAuxMockPtrselectedHostInfo
	condRecorderAuxMockPtrselectedHostInfo.L.Unlock()
	return
}

// (recvhost *selectedHost)Info - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvhost *selectedHost) Info() (reta *HostInfo) {
	FuncAuxMockPtrselectedHostInfo, ok := apomock.GetRegisteredFunc("gocql.selectedHost.Info")
	if ok {
		reta = FuncAuxMockPtrselectedHostInfo.(func(recvhost *selectedHost) (reta *HostInfo))(recvhost)
	} else {
		panic("FuncAuxMockPtrselectedHostInfo ")
	}
	AuxMockIncrementRecorderAuxMockPtrselectedHostInfo()
	return
}

//
// Mock: (recvc *cowHostList)set(arglist []*HostInfo)()
//

type MockArgsTypecowHostListset struct {
	ApomockCallNumber int
	Arglist           []*HostInfo
}

var LastMockArgscowHostListset MockArgsTypecowHostListset

// (recvc *cowHostList)AuxMockset(arglist []*HostInfo)() - Generated mock function
func (recvc *cowHostList) AuxMockset(arglist []*HostInfo) {
	LastMockArgscowHostListset = MockArgsTypecowHostListset{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrcowHostListset(),
		Arglist:           arglist,
	}
	return
}

// RecorderAuxMockPtrcowHostListset  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrcowHostListset int = 0

var condRecorderAuxMockPtrcowHostListset *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrcowHostListset(i int) {
	condRecorderAuxMockPtrcowHostListset.L.Lock()
	for recorderAuxMockPtrcowHostListset < i {
		condRecorderAuxMockPtrcowHostListset.Wait()
	}
	condRecorderAuxMockPtrcowHostListset.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrcowHostListset() {
	condRecorderAuxMockPtrcowHostListset.L.Lock()
	recorderAuxMockPtrcowHostListset++
	condRecorderAuxMockPtrcowHostListset.L.Unlock()
	condRecorderAuxMockPtrcowHostListset.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrcowHostListset() (ret int) {
	condRecorderAuxMockPtrcowHostListset.L.Lock()
	ret = recorderAuxMockPtrcowHostListset
	condRecorderAuxMockPtrcowHostListset.L.Unlock()
	return
}

// (recvc *cowHostList)set - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *cowHostList) set(arglist []*HostInfo) {
	FuncAuxMockPtrcowHostListset, ok := apomock.GetRegisteredFunc("gocql.cowHostList.set")
	if ok {
		FuncAuxMockPtrcowHostListset.(func(recvc *cowHostList, arglist []*HostInfo))(recvc, arglist)
	} else {
		panic("FuncAuxMockPtrcowHostListset ")
	}
	AuxMockIncrementRecorderAuxMockPtrcowHostListset()
	return
}

//
// Mock: (recvhost *selectedHost)Mark(argerr error)()
//

type MockArgsTypeselectedHostMark struct {
	ApomockCallNumber int
	Argerr            error
}

var LastMockArgsselectedHostMark MockArgsTypeselectedHostMark

// (recvhost *selectedHost)AuxMockMark(argerr error)() - Generated mock function
func (recvhost *selectedHost) AuxMockMark(argerr error) {
	LastMockArgsselectedHostMark = MockArgsTypeselectedHostMark{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrselectedHostMark(),
		Argerr:            argerr,
	}
	return
}

// RecorderAuxMockPtrselectedHostMark  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrselectedHostMark int = 0

var condRecorderAuxMockPtrselectedHostMark *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrselectedHostMark(i int) {
	condRecorderAuxMockPtrselectedHostMark.L.Lock()
	for recorderAuxMockPtrselectedHostMark < i {
		condRecorderAuxMockPtrselectedHostMark.Wait()
	}
	condRecorderAuxMockPtrselectedHostMark.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrselectedHostMark() {
	condRecorderAuxMockPtrselectedHostMark.L.Lock()
	recorderAuxMockPtrselectedHostMark++
	condRecorderAuxMockPtrselectedHostMark.L.Unlock()
	condRecorderAuxMockPtrselectedHostMark.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrselectedHostMark() (ret int) {
	condRecorderAuxMockPtrselectedHostMark.L.Lock()
	ret = recorderAuxMockPtrselectedHostMark
	condRecorderAuxMockPtrselectedHostMark.L.Unlock()
	return
}

// (recvhost *selectedHost)Mark - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvhost *selectedHost) Mark(argerr error) {
	FuncAuxMockPtrselectedHostMark, ok := apomock.GetRegisteredFunc("gocql.selectedHost.Mark")
	if ok {
		FuncAuxMockPtrselectedHostMark.(func(recvhost *selectedHost, argerr error))(recvhost, argerr)
	} else {
		panic("FuncAuxMockPtrselectedHostMark ")
	}
	AuxMockIncrementRecorderAuxMockPtrselectedHostMark()
	return
}

//
// Mock: (recvt *tokenAwareHostPolicy)HostUp(arghost *HostInfo)()
//

type MockArgsTypetokenAwareHostPolicyHostUp struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgstokenAwareHostPolicyHostUp MockArgsTypetokenAwareHostPolicyHostUp

// (recvt *tokenAwareHostPolicy)AuxMockHostUp(arghost *HostInfo)() - Generated mock function
func (recvt *tokenAwareHostPolicy) AuxMockHostUp(arghost *HostInfo) {
	LastMockArgstokenAwareHostPolicyHostUp = MockArgsTypetokenAwareHostPolicyHostUp{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicyHostUp(),
		Arghost:           arghost,
	}
	return
}

// RecorderAuxMockPtrtokenAwareHostPolicyHostUp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenAwareHostPolicyHostUp int = 0

var condRecorderAuxMockPtrtokenAwareHostPolicyHostUp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenAwareHostPolicyHostUp(i int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyHostUp.L.Lock()
	for recorderAuxMockPtrtokenAwareHostPolicyHostUp < i {
		condRecorderAuxMockPtrtokenAwareHostPolicyHostUp.Wait()
	}
	condRecorderAuxMockPtrtokenAwareHostPolicyHostUp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyHostUp() {
	condRecorderAuxMockPtrtokenAwareHostPolicyHostUp.L.Lock()
	recorderAuxMockPtrtokenAwareHostPolicyHostUp++
	condRecorderAuxMockPtrtokenAwareHostPolicyHostUp.L.Unlock()
	condRecorderAuxMockPtrtokenAwareHostPolicyHostUp.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenAwareHostPolicyHostUp() (ret int) {
	condRecorderAuxMockPtrtokenAwareHostPolicyHostUp.L.Lock()
	ret = recorderAuxMockPtrtokenAwareHostPolicyHostUp
	condRecorderAuxMockPtrtokenAwareHostPolicyHostUp.L.Unlock()
	return
}

// (recvt *tokenAwareHostPolicy)HostUp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenAwareHostPolicy) HostUp(arghost *HostInfo) {
	FuncAuxMockPtrtokenAwareHostPolicyHostUp, ok := apomock.GetRegisteredFunc("gocql.tokenAwareHostPolicy.HostUp")
	if ok {
		FuncAuxMockPtrtokenAwareHostPolicyHostUp.(func(recvt *tokenAwareHostPolicy, arghost *HostInfo))(recvt, arghost)
	} else {
		panic("FuncAuxMockPtrtokenAwareHostPolicyHostUp ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenAwareHostPolicyHostUp()
	return
}

//
// Mock: (recvr *hostPoolHostPolicy)AddHost(arghost *HostInfo)()
//

type MockArgsTypehostPoolHostPolicyAddHost struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgshostPoolHostPolicyAddHost MockArgsTypehostPoolHostPolicyAddHost

// (recvr *hostPoolHostPolicy)AuxMockAddHost(arghost *HostInfo)() - Generated mock function
func (recvr *hostPoolHostPolicy) AuxMockAddHost(arghost *HostInfo) {
	LastMockArgshostPoolHostPolicyAddHost = MockArgsTypehostPoolHostPolicyAddHost{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrhostPoolHostPolicyAddHost(),
		Arghost:           arghost,
	}
	return
}

// RecorderAuxMockPtrhostPoolHostPolicyAddHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostPoolHostPolicyAddHost int = 0

var condRecorderAuxMockPtrhostPoolHostPolicyAddHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostPoolHostPolicyAddHost(i int) {
	condRecorderAuxMockPtrhostPoolHostPolicyAddHost.L.Lock()
	for recorderAuxMockPtrhostPoolHostPolicyAddHost < i {
		condRecorderAuxMockPtrhostPoolHostPolicyAddHost.Wait()
	}
	condRecorderAuxMockPtrhostPoolHostPolicyAddHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicyAddHost() {
	condRecorderAuxMockPtrhostPoolHostPolicyAddHost.L.Lock()
	recorderAuxMockPtrhostPoolHostPolicyAddHost++
	condRecorderAuxMockPtrhostPoolHostPolicyAddHost.L.Unlock()
	condRecorderAuxMockPtrhostPoolHostPolicyAddHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostPoolHostPolicyAddHost() (ret int) {
	condRecorderAuxMockPtrhostPoolHostPolicyAddHost.L.Lock()
	ret = recorderAuxMockPtrhostPoolHostPolicyAddHost
	condRecorderAuxMockPtrhostPoolHostPolicyAddHost.L.Unlock()
	return
}

// (recvr *hostPoolHostPolicy)AddHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *hostPoolHostPolicy) AddHost(arghost *HostInfo) {
	FuncAuxMockPtrhostPoolHostPolicyAddHost, ok := apomock.GetRegisteredFunc("gocql.hostPoolHostPolicy.AddHost")
	if ok {
		FuncAuxMockPtrhostPoolHostPolicyAddHost.(func(recvr *hostPoolHostPolicy, arghost *HostInfo))(recvr, arghost)
	} else {
		panic("FuncAuxMockPtrhostPoolHostPolicyAddHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostPoolHostPolicyAddHost()
	return
}
