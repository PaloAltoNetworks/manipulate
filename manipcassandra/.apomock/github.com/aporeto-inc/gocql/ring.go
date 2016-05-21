// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructRing, apomockNewStructRing)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructClusterMetadata, apomockNewStructClusterMetadata)

	apomock.RegisterFunc("gocql", "gocql.ring.addHostIfMissing", (*ring).AuxMockaddHostIfMissing)
	apomock.RegisterFunc("gocql", "gocql.ring.removeHost", (*ring).AuxMockremoveHost)
	apomock.RegisterFunc("gocql", "gocql.clusterMetadata.setPartitioner", (*clusterMetadata).AuxMocksetPartitioner)
	apomock.RegisterFunc("gocql", "gocql.ring.rrHost", (*ring).AuxMockrrHost)
	apomock.RegisterFunc("gocql", "gocql.ring.getHost", (*ring).AuxMockgetHost)
	apomock.RegisterFunc("gocql", "gocql.ring.allHosts", (*ring).AuxMockallHosts)
	apomock.RegisterFunc("gocql", "gocql.ring.addHost", (*ring).AuxMockaddHost)
}

const (
	ApomockStructRing            = 60
	ApomockStructClusterMetadata = 61
)

//
// Internal Types: in this package and their exportable versions
//
type ring struct {
	endpoints []string
	mu        sync.RWMutex
	hosts     map[string]*HostInfo
	hostList  []*HostInfo
	pos       uint32
}
type clusterMetadata struct {
	mu          sync.RWMutex
	partitioner string
}

//
// External Types: in this package
//

func apomockNewStructRing() interface{}            { return &ring{} }
func apomockNewStructClusterMetadata() interface{} { return &clusterMetadata{} }

//
// Mock: (recvr *ring)addHostIfMissing(arghost *HostInfo)(reta *HostInfo, retb bool)
//

type MockArgsTyperingaddHostIfMissing struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgsringaddHostIfMissing MockArgsTyperingaddHostIfMissing

// (recvr *ring)AuxMockaddHostIfMissing(arghost *HostInfo)(reta *HostInfo, retb bool) - Generated mock function
func (recvr *ring) AuxMockaddHostIfMissing(arghost *HostInfo) (reta *HostInfo, retb bool) {
	LastMockArgsringaddHostIfMissing = MockArgsTyperingaddHostIfMissing{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrringaddHostIfMissing(),
		Arghost:           arghost,
	}
	rargs, rerr := apomock.GetNext("gocql.ring.addHostIfMissing")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ring.addHostIfMissing")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.ring.addHostIfMissing")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(bool)
	}
	return
}

// RecorderAuxMockPtrringaddHostIfMissing  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrringaddHostIfMissing int = 0

var condRecorderAuxMockPtrringaddHostIfMissing *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrringaddHostIfMissing(i int) {
	condRecorderAuxMockPtrringaddHostIfMissing.L.Lock()
	for recorderAuxMockPtrringaddHostIfMissing < i {
		condRecorderAuxMockPtrringaddHostIfMissing.Wait()
	}
	condRecorderAuxMockPtrringaddHostIfMissing.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrringaddHostIfMissing() {
	condRecorderAuxMockPtrringaddHostIfMissing.L.Lock()
	recorderAuxMockPtrringaddHostIfMissing++
	condRecorderAuxMockPtrringaddHostIfMissing.L.Unlock()
	condRecorderAuxMockPtrringaddHostIfMissing.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrringaddHostIfMissing() (ret int) {
	condRecorderAuxMockPtrringaddHostIfMissing.L.Lock()
	ret = recorderAuxMockPtrringaddHostIfMissing
	condRecorderAuxMockPtrringaddHostIfMissing.L.Unlock()
	return
}

// (recvr *ring)addHostIfMissing - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *ring) addHostIfMissing(arghost *HostInfo) (reta *HostInfo, retb bool) {
	FuncAuxMockPtrringaddHostIfMissing, ok := apomock.GetRegisteredFunc("gocql.ring.addHostIfMissing")
	if ok {
		reta, retb = FuncAuxMockPtrringaddHostIfMissing.(func(recvr *ring, arghost *HostInfo) (reta *HostInfo, retb bool))(recvr, arghost)
	} else {
		panic("FuncAuxMockPtrringaddHostIfMissing ")
	}
	AuxMockIncrementRecorderAuxMockPtrringaddHostIfMissing()
	return
}

//
// Mock: (recvr *ring)removeHost(argaddr string)(reta bool)
//

type MockArgsTyperingremoveHost struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgsringremoveHost MockArgsTyperingremoveHost

// (recvr *ring)AuxMockremoveHost(argaddr string)(reta bool) - Generated mock function
func (recvr *ring) AuxMockremoveHost(argaddr string) (reta bool) {
	LastMockArgsringremoveHost = MockArgsTyperingremoveHost{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrringremoveHost(),
		Argaddr:           argaddr,
	}
	rargs, rerr := apomock.GetNext("gocql.ring.removeHost")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ring.removeHost")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.ring.removeHost")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrringremoveHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrringremoveHost int = 0

var condRecorderAuxMockPtrringremoveHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrringremoveHost(i int) {
	condRecorderAuxMockPtrringremoveHost.L.Lock()
	for recorderAuxMockPtrringremoveHost < i {
		condRecorderAuxMockPtrringremoveHost.Wait()
	}
	condRecorderAuxMockPtrringremoveHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrringremoveHost() {
	condRecorderAuxMockPtrringremoveHost.L.Lock()
	recorderAuxMockPtrringremoveHost++
	condRecorderAuxMockPtrringremoveHost.L.Unlock()
	condRecorderAuxMockPtrringremoveHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrringremoveHost() (ret int) {
	condRecorderAuxMockPtrringremoveHost.L.Lock()
	ret = recorderAuxMockPtrringremoveHost
	condRecorderAuxMockPtrringremoveHost.L.Unlock()
	return
}

// (recvr *ring)removeHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *ring) removeHost(argaddr string) (reta bool) {
	FuncAuxMockPtrringremoveHost, ok := apomock.GetRegisteredFunc("gocql.ring.removeHost")
	if ok {
		reta = FuncAuxMockPtrringremoveHost.(func(recvr *ring, argaddr string) (reta bool))(recvr, argaddr)
	} else {
		panic("FuncAuxMockPtrringremoveHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrringremoveHost()
	return
}

//
// Mock: (recvc *clusterMetadata)setPartitioner(argpartitioner string)()
//

type MockArgsTypeclusterMetadatasetPartitioner struct {
	ApomockCallNumber int
	Argpartitioner    string
}

var LastMockArgsclusterMetadatasetPartitioner MockArgsTypeclusterMetadatasetPartitioner

// (recvc *clusterMetadata)AuxMocksetPartitioner(argpartitioner string)() - Generated mock function
func (recvc *clusterMetadata) AuxMocksetPartitioner(argpartitioner string) {
	LastMockArgsclusterMetadatasetPartitioner = MockArgsTypeclusterMetadatasetPartitioner{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrclusterMetadatasetPartitioner(),
		Argpartitioner:    argpartitioner,
	}
	return
}

// RecorderAuxMockPtrclusterMetadatasetPartitioner  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrclusterMetadatasetPartitioner int = 0

var condRecorderAuxMockPtrclusterMetadatasetPartitioner *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrclusterMetadatasetPartitioner(i int) {
	condRecorderAuxMockPtrclusterMetadatasetPartitioner.L.Lock()
	for recorderAuxMockPtrclusterMetadatasetPartitioner < i {
		condRecorderAuxMockPtrclusterMetadatasetPartitioner.Wait()
	}
	condRecorderAuxMockPtrclusterMetadatasetPartitioner.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrclusterMetadatasetPartitioner() {
	condRecorderAuxMockPtrclusterMetadatasetPartitioner.L.Lock()
	recorderAuxMockPtrclusterMetadatasetPartitioner++
	condRecorderAuxMockPtrclusterMetadatasetPartitioner.L.Unlock()
	condRecorderAuxMockPtrclusterMetadatasetPartitioner.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrclusterMetadatasetPartitioner() (ret int) {
	condRecorderAuxMockPtrclusterMetadatasetPartitioner.L.Lock()
	ret = recorderAuxMockPtrclusterMetadatasetPartitioner
	condRecorderAuxMockPtrclusterMetadatasetPartitioner.L.Unlock()
	return
}

// (recvc *clusterMetadata)setPartitioner - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *clusterMetadata) setPartitioner(argpartitioner string) {
	FuncAuxMockPtrclusterMetadatasetPartitioner, ok := apomock.GetRegisteredFunc("gocql.clusterMetadata.setPartitioner")
	if ok {
		FuncAuxMockPtrclusterMetadatasetPartitioner.(func(recvc *clusterMetadata, argpartitioner string))(recvc, argpartitioner)
	} else {
		panic("FuncAuxMockPtrclusterMetadatasetPartitioner ")
	}
	AuxMockIncrementRecorderAuxMockPtrclusterMetadatasetPartitioner()
	return
}

//
// Mock: (recvr *ring)rrHost()(reta *HostInfo)
//

type MockArgsTyperingrrHost struct {
	ApomockCallNumber int
}

var LastMockArgsringrrHost MockArgsTyperingrrHost

// (recvr *ring)AuxMockrrHost()(reta *HostInfo) - Generated mock function
func (recvr *ring) AuxMockrrHost() (reta *HostInfo) {
	rargs, rerr := apomock.GetNext("gocql.ring.rrHost")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ring.rrHost")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.ring.rrHost")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrringrrHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrringrrHost int = 0

var condRecorderAuxMockPtrringrrHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrringrrHost(i int) {
	condRecorderAuxMockPtrringrrHost.L.Lock()
	for recorderAuxMockPtrringrrHost < i {
		condRecorderAuxMockPtrringrrHost.Wait()
	}
	condRecorderAuxMockPtrringrrHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrringrrHost() {
	condRecorderAuxMockPtrringrrHost.L.Lock()
	recorderAuxMockPtrringrrHost++
	condRecorderAuxMockPtrringrrHost.L.Unlock()
	condRecorderAuxMockPtrringrrHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrringrrHost() (ret int) {
	condRecorderAuxMockPtrringrrHost.L.Lock()
	ret = recorderAuxMockPtrringrrHost
	condRecorderAuxMockPtrringrrHost.L.Unlock()
	return
}

// (recvr *ring)rrHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *ring) rrHost() (reta *HostInfo) {
	FuncAuxMockPtrringrrHost, ok := apomock.GetRegisteredFunc("gocql.ring.rrHost")
	if ok {
		reta = FuncAuxMockPtrringrrHost.(func(recvr *ring) (reta *HostInfo))(recvr)
	} else {
		panic("FuncAuxMockPtrringrrHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrringrrHost()
	return
}

//
// Mock: (recvr *ring)getHost(argaddr string)(reta *HostInfo)
//

type MockArgsTyperinggetHost struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgsringgetHost MockArgsTyperinggetHost

// (recvr *ring)AuxMockgetHost(argaddr string)(reta *HostInfo) - Generated mock function
func (recvr *ring) AuxMockgetHost(argaddr string) (reta *HostInfo) {
	LastMockArgsringgetHost = MockArgsTyperinggetHost{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrringgetHost(),
		Argaddr:           argaddr,
	}
	rargs, rerr := apomock.GetNext("gocql.ring.getHost")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ring.getHost")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.ring.getHost")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrringgetHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrringgetHost int = 0

var condRecorderAuxMockPtrringgetHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrringgetHost(i int) {
	condRecorderAuxMockPtrringgetHost.L.Lock()
	for recorderAuxMockPtrringgetHost < i {
		condRecorderAuxMockPtrringgetHost.Wait()
	}
	condRecorderAuxMockPtrringgetHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrringgetHost() {
	condRecorderAuxMockPtrringgetHost.L.Lock()
	recorderAuxMockPtrringgetHost++
	condRecorderAuxMockPtrringgetHost.L.Unlock()
	condRecorderAuxMockPtrringgetHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrringgetHost() (ret int) {
	condRecorderAuxMockPtrringgetHost.L.Lock()
	ret = recorderAuxMockPtrringgetHost
	condRecorderAuxMockPtrringgetHost.L.Unlock()
	return
}

// (recvr *ring)getHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *ring) getHost(argaddr string) (reta *HostInfo) {
	FuncAuxMockPtrringgetHost, ok := apomock.GetRegisteredFunc("gocql.ring.getHost")
	if ok {
		reta = FuncAuxMockPtrringgetHost.(func(recvr *ring, argaddr string) (reta *HostInfo))(recvr, argaddr)
	} else {
		panic("FuncAuxMockPtrringgetHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrringgetHost()
	return
}

//
// Mock: (recvr *ring)allHosts()(reta []*HostInfo)
//

type MockArgsTyperingallHosts struct {
	ApomockCallNumber int
}

var LastMockArgsringallHosts MockArgsTyperingallHosts

// (recvr *ring)AuxMockallHosts()(reta []*HostInfo) - Generated mock function
func (recvr *ring) AuxMockallHosts() (reta []*HostInfo) {
	rargs, rerr := apomock.GetNext("gocql.ring.allHosts")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ring.allHosts")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.ring.allHosts")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]*HostInfo)
	}
	return
}

// RecorderAuxMockPtrringallHosts  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrringallHosts int = 0

var condRecorderAuxMockPtrringallHosts *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrringallHosts(i int) {
	condRecorderAuxMockPtrringallHosts.L.Lock()
	for recorderAuxMockPtrringallHosts < i {
		condRecorderAuxMockPtrringallHosts.Wait()
	}
	condRecorderAuxMockPtrringallHosts.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrringallHosts() {
	condRecorderAuxMockPtrringallHosts.L.Lock()
	recorderAuxMockPtrringallHosts++
	condRecorderAuxMockPtrringallHosts.L.Unlock()
	condRecorderAuxMockPtrringallHosts.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrringallHosts() (ret int) {
	condRecorderAuxMockPtrringallHosts.L.Lock()
	ret = recorderAuxMockPtrringallHosts
	condRecorderAuxMockPtrringallHosts.L.Unlock()
	return
}

// (recvr *ring)allHosts - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *ring) allHosts() (reta []*HostInfo) {
	FuncAuxMockPtrringallHosts, ok := apomock.GetRegisteredFunc("gocql.ring.allHosts")
	if ok {
		reta = FuncAuxMockPtrringallHosts.(func(recvr *ring) (reta []*HostInfo))(recvr)
	} else {
		panic("FuncAuxMockPtrringallHosts ")
	}
	AuxMockIncrementRecorderAuxMockPtrringallHosts()
	return
}

//
// Mock: (recvr *ring)addHost(arghost *HostInfo)(reta bool)
//

type MockArgsTyperingaddHost struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgsringaddHost MockArgsTyperingaddHost

// (recvr *ring)AuxMockaddHost(arghost *HostInfo)(reta bool) - Generated mock function
func (recvr *ring) AuxMockaddHost(arghost *HostInfo) (reta bool) {
	LastMockArgsringaddHost = MockArgsTyperingaddHost{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrringaddHost(),
		Arghost:           arghost,
	}
	rargs, rerr := apomock.GetNext("gocql.ring.addHost")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ring.addHost")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.ring.addHost")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrringaddHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrringaddHost int = 0

var condRecorderAuxMockPtrringaddHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrringaddHost(i int) {
	condRecorderAuxMockPtrringaddHost.L.Lock()
	for recorderAuxMockPtrringaddHost < i {
		condRecorderAuxMockPtrringaddHost.Wait()
	}
	condRecorderAuxMockPtrringaddHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrringaddHost() {
	condRecorderAuxMockPtrringaddHost.L.Lock()
	recorderAuxMockPtrringaddHost++
	condRecorderAuxMockPtrringaddHost.L.Unlock()
	condRecorderAuxMockPtrringaddHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrringaddHost() (ret int) {
	condRecorderAuxMockPtrringaddHost.L.Lock()
	ret = recorderAuxMockPtrringaddHost
	condRecorderAuxMockPtrringaddHost.L.Unlock()
	return
}

// (recvr *ring)addHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *ring) addHost(arghost *HostInfo) (reta bool) {
	FuncAuxMockPtrringaddHost, ok := apomock.GetRegisteredFunc("gocql.ring.addHost")
	if ok {
		reta = FuncAuxMockPtrringaddHost.(func(recvr *ring, arghost *HostInfo) (reta bool))(recvr, arghost)
	} else {
		panic("FuncAuxMockPtrringaddHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrringaddHost()
	return
}
