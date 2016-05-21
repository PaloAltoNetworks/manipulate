// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

import "crypto/tls"

func init() {
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructPolicyConnPool, apomockNewStructPolicyConnPool)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructHostConnPool, apomockNewStructHostConnPool)

	apomock.RegisterFunc("gocql", "gocql.setupTLSConfig", AuxMocksetupTLSConfig)
	apomock.RegisterFunc("gocql", "gocql.policyConnPool.Size", (*policyConnPool).AuxMockSize)
	apomock.RegisterFunc("gocql", "gocql.policyConnPool.getPool", (*policyConnPool).AuxMockgetPool)
	apomock.RegisterFunc("gocql", "gocql.policyConnPool.addHost", (*policyConnPool).AuxMockaddHost)
	apomock.RegisterFunc("gocql", "gocql.policyConnPool.removeHost", (*policyConnPool).AuxMockremoveHost)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.String", (*hostConnPool).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.logConnectErr", (*hostConnPool).AuxMocklogConnectErr)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.connectMany", (*hostConnPool).AuxMockconnectMany)
	apomock.RegisterFunc("gocql", "gocql.newPolicyConnPool", AuxMocknewPolicyConnPool)
	apomock.RegisterFunc("gocql", "gocql.policyConnPool.SetHosts", (*policyConnPool).AuxMockSetHosts)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.Pick", (*hostConnPool).AuxMockPick)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.fill", (*hostConnPool).AuxMockfill)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.fillingStopped", (*hostConnPool).AuxMockfillingStopped)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.HandleError", (*hostConnPool).AuxMockHandleError)
	apomock.RegisterFunc("gocql", "gocql.connConfig", AuxMockconnConfig)
	apomock.RegisterFunc("gocql", "gocql.policyConnPool.Close", (*policyConnPool).AuxMockClose)
	apomock.RegisterFunc("gocql", "gocql.policyConnPool.hostDown", (*policyConnPool).AuxMockhostDown)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.connect", (*hostConnPool).AuxMockconnect)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.drainLocked", (*hostConnPool).AuxMockdrainLocked)
	apomock.RegisterFunc("gocql", "gocql.policyConnPool.hostUp", (*policyConnPool).AuxMockhostUp)
	apomock.RegisterFunc("gocql", "gocql.newHostConnPool", AuxMocknewHostConnPool)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.Size", (*hostConnPool).AuxMockSize)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.Close", (*hostConnPool).AuxMockClose)
	apomock.RegisterFunc("gocql", "gocql.hostConnPool.drain", (*hostConnPool).AuxMockdrain)
}

const (
	ApomockStructPolicyConnPool = 0
	ApomockStructHostConnPool   = 1
)

//
// Internal Types: in this package and their exportable versions
//
type policyConnPool struct {
	session       *Session
	port          int
	numConns      int
	keyspace      string
	mu            sync.RWMutex
	hostConnPools map[string]*hostConnPool
	endpoints     []string
}
type hostConnPool struct {
	session  *Session
	host     *HostInfo
	port     int
	addr     string
	size     int
	keyspace string
	mu       sync.RWMutex
	conns    []*Conn
	closed   bool
	filling  bool
	pos      uint32
}

//
// External Types: in this package
//
type SetHosts interface {
	SetHosts(hosts []*HostInfo)
}

type SetPartitioner interface {
	SetPartitioner(partitioner string)
}

func apomockNewStructPolicyConnPool() interface{} { return &policyConnPool{} }
func apomockNewStructHostConnPool() interface{}   { return &hostConnPool{} }

//
// Mock: setupTLSConfig(argsslOpts *SslOptions)(reta *tls.Config, retb error)
//

type MockArgsTypesetupTLSConfig struct {
	ApomockCallNumber int
	ArgsslOpts        *SslOptions
}

var LastMockArgssetupTLSConfig MockArgsTypesetupTLSConfig

// AuxMocksetupTLSConfig(argsslOpts *SslOptions)(reta *tls.Config, retb error) - Generated mock function
func AuxMocksetupTLSConfig(argsslOpts *SslOptions) (reta *tls.Config, retb error) {
	LastMockArgssetupTLSConfig = MockArgsTypesetupTLSConfig{
		ApomockCallNumber: AuxMockGetRecorderAuxMocksetupTLSConfig(),
		ArgsslOpts:        argsslOpts,
	}
	rargs, rerr := apomock.GetNext("gocql.setupTLSConfig")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.setupTLSConfig")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.setupTLSConfig")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*tls.Config)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMocksetupTLSConfig  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocksetupTLSConfig int = 0

var condRecorderAuxMocksetupTLSConfig *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocksetupTLSConfig(i int) {
	condRecorderAuxMocksetupTLSConfig.L.Lock()
	for recorderAuxMocksetupTLSConfig < i {
		condRecorderAuxMocksetupTLSConfig.Wait()
	}
	condRecorderAuxMocksetupTLSConfig.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocksetupTLSConfig() {
	condRecorderAuxMocksetupTLSConfig.L.Lock()
	recorderAuxMocksetupTLSConfig++
	condRecorderAuxMocksetupTLSConfig.L.Unlock()
	condRecorderAuxMocksetupTLSConfig.Broadcast()
}
func AuxMockGetRecorderAuxMocksetupTLSConfig() (ret int) {
	condRecorderAuxMocksetupTLSConfig.L.Lock()
	ret = recorderAuxMocksetupTLSConfig
	condRecorderAuxMocksetupTLSConfig.L.Unlock()
	return
}

// setupTLSConfig - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func setupTLSConfig(argsslOpts *SslOptions) (reta *tls.Config, retb error) {
	FuncAuxMocksetupTLSConfig, ok := apomock.GetRegisteredFunc("gocql.setupTLSConfig")
	if ok {
		reta, retb = FuncAuxMocksetupTLSConfig.(func(argsslOpts *SslOptions) (reta *tls.Config, retb error))(argsslOpts)
	} else {
		panic("FuncAuxMocksetupTLSConfig ")
	}
	AuxMockIncrementRecorderAuxMocksetupTLSConfig()
	return
}

//
// Mock: (recvp *policyConnPool)Size()(reta int)
//

type MockArgsTypepolicyConnPoolSize struct {
	ApomockCallNumber int
}

var LastMockArgspolicyConnPoolSize MockArgsTypepolicyConnPoolSize

// (recvp *policyConnPool)AuxMockSize()(reta int) - Generated mock function
func (recvp *policyConnPool) AuxMockSize() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.policyConnPool.Size")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.policyConnPool.Size")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.policyConnPool.Size")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrpolicyConnPoolSize  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpolicyConnPoolSize int = 0

var condRecorderAuxMockPtrpolicyConnPoolSize *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpolicyConnPoolSize(i int) {
	condRecorderAuxMockPtrpolicyConnPoolSize.L.Lock()
	for recorderAuxMockPtrpolicyConnPoolSize < i {
		condRecorderAuxMockPtrpolicyConnPoolSize.Wait()
	}
	condRecorderAuxMockPtrpolicyConnPoolSize.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolSize() {
	condRecorderAuxMockPtrpolicyConnPoolSize.L.Lock()
	recorderAuxMockPtrpolicyConnPoolSize++
	condRecorderAuxMockPtrpolicyConnPoolSize.L.Unlock()
	condRecorderAuxMockPtrpolicyConnPoolSize.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpolicyConnPoolSize() (ret int) {
	condRecorderAuxMockPtrpolicyConnPoolSize.L.Lock()
	ret = recorderAuxMockPtrpolicyConnPoolSize
	condRecorderAuxMockPtrpolicyConnPoolSize.L.Unlock()
	return
}

// (recvp *policyConnPool)Size - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *policyConnPool) Size() (reta int) {
	FuncAuxMockPtrpolicyConnPoolSize, ok := apomock.GetRegisteredFunc("gocql.policyConnPool.Size")
	if ok {
		reta = FuncAuxMockPtrpolicyConnPoolSize.(func(recvp *policyConnPool) (reta int))(recvp)
	} else {
		panic("FuncAuxMockPtrpolicyConnPoolSize ")
	}
	AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolSize()
	return
}

//
// Mock: (recvp *policyConnPool)getPool(argaddr string)(retpool *hostConnPool, retok bool)
//

type MockArgsTypepolicyConnPoolgetPool struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgspolicyConnPoolgetPool MockArgsTypepolicyConnPoolgetPool

// (recvp *policyConnPool)AuxMockgetPool(argaddr string)(retpool *hostConnPool, retok bool) - Generated mock function
func (recvp *policyConnPool) AuxMockgetPool(argaddr string) (retpool *hostConnPool, retok bool) {
	LastMockArgspolicyConnPoolgetPool = MockArgsTypepolicyConnPoolgetPool{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrpolicyConnPoolgetPool(),
		Argaddr:           argaddr,
	}
	rargs, rerr := apomock.GetNext("gocql.policyConnPool.getPool")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.policyConnPool.getPool")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.policyConnPool.getPool")
	}
	if rargs.GetArg(0) != nil {
		retpool = rargs.GetArg(0).(*hostConnPool)
	}
	if rargs.GetArg(1) != nil {
		retok = rargs.GetArg(1).(bool)
	}
	return
}

// RecorderAuxMockPtrpolicyConnPoolgetPool  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpolicyConnPoolgetPool int = 0

var condRecorderAuxMockPtrpolicyConnPoolgetPool *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpolicyConnPoolgetPool(i int) {
	condRecorderAuxMockPtrpolicyConnPoolgetPool.L.Lock()
	for recorderAuxMockPtrpolicyConnPoolgetPool < i {
		condRecorderAuxMockPtrpolicyConnPoolgetPool.Wait()
	}
	condRecorderAuxMockPtrpolicyConnPoolgetPool.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolgetPool() {
	condRecorderAuxMockPtrpolicyConnPoolgetPool.L.Lock()
	recorderAuxMockPtrpolicyConnPoolgetPool++
	condRecorderAuxMockPtrpolicyConnPoolgetPool.L.Unlock()
	condRecorderAuxMockPtrpolicyConnPoolgetPool.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpolicyConnPoolgetPool() (ret int) {
	condRecorderAuxMockPtrpolicyConnPoolgetPool.L.Lock()
	ret = recorderAuxMockPtrpolicyConnPoolgetPool
	condRecorderAuxMockPtrpolicyConnPoolgetPool.L.Unlock()
	return
}

// (recvp *policyConnPool)getPool - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *policyConnPool) getPool(argaddr string) (retpool *hostConnPool, retok bool) {
	FuncAuxMockPtrpolicyConnPoolgetPool, ok := apomock.GetRegisteredFunc("gocql.policyConnPool.getPool")
	if ok {
		retpool, retok = FuncAuxMockPtrpolicyConnPoolgetPool.(func(recvp *policyConnPool, argaddr string) (retpool *hostConnPool, retok bool))(recvp, argaddr)
	} else {
		panic("FuncAuxMockPtrpolicyConnPoolgetPool ")
	}
	AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolgetPool()
	return
}

//
// Mock: (recvp *policyConnPool)addHost(arghost *HostInfo)()
//

type MockArgsTypepolicyConnPooladdHost struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgspolicyConnPooladdHost MockArgsTypepolicyConnPooladdHost

// (recvp *policyConnPool)AuxMockaddHost(arghost *HostInfo)() - Generated mock function
func (recvp *policyConnPool) AuxMockaddHost(arghost *HostInfo) {
	LastMockArgspolicyConnPooladdHost = MockArgsTypepolicyConnPooladdHost{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrpolicyConnPooladdHost(),
		Arghost:           arghost,
	}
	return
}

// RecorderAuxMockPtrpolicyConnPooladdHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpolicyConnPooladdHost int = 0

var condRecorderAuxMockPtrpolicyConnPooladdHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpolicyConnPooladdHost(i int) {
	condRecorderAuxMockPtrpolicyConnPooladdHost.L.Lock()
	for recorderAuxMockPtrpolicyConnPooladdHost < i {
		condRecorderAuxMockPtrpolicyConnPooladdHost.Wait()
	}
	condRecorderAuxMockPtrpolicyConnPooladdHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpolicyConnPooladdHost() {
	condRecorderAuxMockPtrpolicyConnPooladdHost.L.Lock()
	recorderAuxMockPtrpolicyConnPooladdHost++
	condRecorderAuxMockPtrpolicyConnPooladdHost.L.Unlock()
	condRecorderAuxMockPtrpolicyConnPooladdHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpolicyConnPooladdHost() (ret int) {
	condRecorderAuxMockPtrpolicyConnPooladdHost.L.Lock()
	ret = recorderAuxMockPtrpolicyConnPooladdHost
	condRecorderAuxMockPtrpolicyConnPooladdHost.L.Unlock()
	return
}

// (recvp *policyConnPool)addHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *policyConnPool) addHost(arghost *HostInfo) {
	FuncAuxMockPtrpolicyConnPooladdHost, ok := apomock.GetRegisteredFunc("gocql.policyConnPool.addHost")
	if ok {
		FuncAuxMockPtrpolicyConnPooladdHost.(func(recvp *policyConnPool, arghost *HostInfo))(recvp, arghost)
	} else {
		panic("FuncAuxMockPtrpolicyConnPooladdHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrpolicyConnPooladdHost()
	return
}

//
// Mock: (recvp *policyConnPool)removeHost(argaddr string)()
//

type MockArgsTypepolicyConnPoolremoveHost struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgspolicyConnPoolremoveHost MockArgsTypepolicyConnPoolremoveHost

// (recvp *policyConnPool)AuxMockremoveHost(argaddr string)() - Generated mock function
func (recvp *policyConnPool) AuxMockremoveHost(argaddr string) {
	LastMockArgspolicyConnPoolremoveHost = MockArgsTypepolicyConnPoolremoveHost{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrpolicyConnPoolremoveHost(),
		Argaddr:           argaddr,
	}
	return
}

// RecorderAuxMockPtrpolicyConnPoolremoveHost  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpolicyConnPoolremoveHost int = 0

var condRecorderAuxMockPtrpolicyConnPoolremoveHost *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpolicyConnPoolremoveHost(i int) {
	condRecorderAuxMockPtrpolicyConnPoolremoveHost.L.Lock()
	for recorderAuxMockPtrpolicyConnPoolremoveHost < i {
		condRecorderAuxMockPtrpolicyConnPoolremoveHost.Wait()
	}
	condRecorderAuxMockPtrpolicyConnPoolremoveHost.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolremoveHost() {
	condRecorderAuxMockPtrpolicyConnPoolremoveHost.L.Lock()
	recorderAuxMockPtrpolicyConnPoolremoveHost++
	condRecorderAuxMockPtrpolicyConnPoolremoveHost.L.Unlock()
	condRecorderAuxMockPtrpolicyConnPoolremoveHost.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpolicyConnPoolremoveHost() (ret int) {
	condRecorderAuxMockPtrpolicyConnPoolremoveHost.L.Lock()
	ret = recorderAuxMockPtrpolicyConnPoolremoveHost
	condRecorderAuxMockPtrpolicyConnPoolremoveHost.L.Unlock()
	return
}

// (recvp *policyConnPool)removeHost - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *policyConnPool) removeHost(argaddr string) {
	FuncAuxMockPtrpolicyConnPoolremoveHost, ok := apomock.GetRegisteredFunc("gocql.policyConnPool.removeHost")
	if ok {
		FuncAuxMockPtrpolicyConnPoolremoveHost.(func(recvp *policyConnPool, argaddr string))(recvp, argaddr)
	} else {
		panic("FuncAuxMockPtrpolicyConnPoolremoveHost ")
	}
	AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolremoveHost()
	return
}

//
// Mock: (recvh *hostConnPool)String()(reta string)
//

type MockArgsTypehostConnPoolString struct {
	ApomockCallNumber int
}

var LastMockArgshostConnPoolString MockArgsTypehostConnPoolString

// (recvh *hostConnPool)AuxMockString()(reta string) - Generated mock function
func (recvh *hostConnPool) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.hostConnPool.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.hostConnPool.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.hostConnPool.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrhostConnPoolString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPoolString int = 0

var condRecorderAuxMockPtrhostConnPoolString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPoolString(i int) {
	condRecorderAuxMockPtrhostConnPoolString.L.Lock()
	for recorderAuxMockPtrhostConnPoolString < i {
		condRecorderAuxMockPtrhostConnPoolString.Wait()
	}
	condRecorderAuxMockPtrhostConnPoolString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPoolString() {
	condRecorderAuxMockPtrhostConnPoolString.L.Lock()
	recorderAuxMockPtrhostConnPoolString++
	condRecorderAuxMockPtrhostConnPoolString.L.Unlock()
	condRecorderAuxMockPtrhostConnPoolString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPoolString() (ret int) {
	condRecorderAuxMockPtrhostConnPoolString.L.Lock()
	ret = recorderAuxMockPtrhostConnPoolString
	condRecorderAuxMockPtrhostConnPoolString.L.Unlock()
	return
}

// (recvh *hostConnPool)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvh *hostConnPool) String() (reta string) {
	FuncAuxMockPtrhostConnPoolString, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.String")
	if ok {
		reta = FuncAuxMockPtrhostConnPoolString.(func(recvh *hostConnPool) (reta string))(recvh)
	} else {
		panic("FuncAuxMockPtrhostConnPoolString ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPoolString()
	return
}

//
// Mock: (recvpool *hostConnPool)logConnectErr(argerr error)()
//

type MockArgsTypehostConnPoollogConnectErr struct {
	ApomockCallNumber int
	Argerr            error
}

var LastMockArgshostConnPoollogConnectErr MockArgsTypehostConnPoollogConnectErr

// (recvpool *hostConnPool)AuxMocklogConnectErr(argerr error)() - Generated mock function
func (recvpool *hostConnPool) AuxMocklogConnectErr(argerr error) {
	LastMockArgshostConnPoollogConnectErr = MockArgsTypehostConnPoollogConnectErr{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrhostConnPoollogConnectErr(),
		Argerr:            argerr,
	}
	return
}

// RecorderAuxMockPtrhostConnPoollogConnectErr  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPoollogConnectErr int = 0

var condRecorderAuxMockPtrhostConnPoollogConnectErr *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPoollogConnectErr(i int) {
	condRecorderAuxMockPtrhostConnPoollogConnectErr.L.Lock()
	for recorderAuxMockPtrhostConnPoollogConnectErr < i {
		condRecorderAuxMockPtrhostConnPoollogConnectErr.Wait()
	}
	condRecorderAuxMockPtrhostConnPoollogConnectErr.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPoollogConnectErr() {
	condRecorderAuxMockPtrhostConnPoollogConnectErr.L.Lock()
	recorderAuxMockPtrhostConnPoollogConnectErr++
	condRecorderAuxMockPtrhostConnPoollogConnectErr.L.Unlock()
	condRecorderAuxMockPtrhostConnPoollogConnectErr.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPoollogConnectErr() (ret int) {
	condRecorderAuxMockPtrhostConnPoollogConnectErr.L.Lock()
	ret = recorderAuxMockPtrhostConnPoollogConnectErr
	condRecorderAuxMockPtrhostConnPoollogConnectErr.L.Unlock()
	return
}

// (recvpool *hostConnPool)logConnectErr - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvpool *hostConnPool) logConnectErr(argerr error) {
	FuncAuxMockPtrhostConnPoollogConnectErr, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.logConnectErr")
	if ok {
		FuncAuxMockPtrhostConnPoollogConnectErr.(func(recvpool *hostConnPool, argerr error))(recvpool, argerr)
	} else {
		panic("FuncAuxMockPtrhostConnPoollogConnectErr ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPoollogConnectErr()
	return
}

//
// Mock: (recvpool *hostConnPool)connectMany(argcount int)(reta error)
//

type MockArgsTypehostConnPoolconnectMany struct {
	ApomockCallNumber int
	Argcount          int
}

var LastMockArgshostConnPoolconnectMany MockArgsTypehostConnPoolconnectMany

// (recvpool *hostConnPool)AuxMockconnectMany(argcount int)(reta error) - Generated mock function
func (recvpool *hostConnPool) AuxMockconnectMany(argcount int) (reta error) {
	LastMockArgshostConnPoolconnectMany = MockArgsTypehostConnPoolconnectMany{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrhostConnPoolconnectMany(),
		Argcount:          argcount,
	}
	rargs, rerr := apomock.GetNext("gocql.hostConnPool.connectMany")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.hostConnPool.connectMany")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.hostConnPool.connectMany")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrhostConnPoolconnectMany  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPoolconnectMany int = 0

var condRecorderAuxMockPtrhostConnPoolconnectMany *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPoolconnectMany(i int) {
	condRecorderAuxMockPtrhostConnPoolconnectMany.L.Lock()
	for recorderAuxMockPtrhostConnPoolconnectMany < i {
		condRecorderAuxMockPtrhostConnPoolconnectMany.Wait()
	}
	condRecorderAuxMockPtrhostConnPoolconnectMany.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPoolconnectMany() {
	condRecorderAuxMockPtrhostConnPoolconnectMany.L.Lock()
	recorderAuxMockPtrhostConnPoolconnectMany++
	condRecorderAuxMockPtrhostConnPoolconnectMany.L.Unlock()
	condRecorderAuxMockPtrhostConnPoolconnectMany.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPoolconnectMany() (ret int) {
	condRecorderAuxMockPtrhostConnPoolconnectMany.L.Lock()
	ret = recorderAuxMockPtrhostConnPoolconnectMany
	condRecorderAuxMockPtrhostConnPoolconnectMany.L.Unlock()
	return
}

// (recvpool *hostConnPool)connectMany - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvpool *hostConnPool) connectMany(argcount int) (reta error) {
	FuncAuxMockPtrhostConnPoolconnectMany, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.connectMany")
	if ok {
		reta = FuncAuxMockPtrhostConnPoolconnectMany.(func(recvpool *hostConnPool, argcount int) (reta error))(recvpool, argcount)
	} else {
		panic("FuncAuxMockPtrhostConnPoolconnectMany ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPoolconnectMany()
	return
}

//
// Mock: newPolicyConnPool(argsession *Session)(reta *policyConnPool)
//

type MockArgsTypenewPolicyConnPool struct {
	ApomockCallNumber int
	Argsession        *Session
}

var LastMockArgsnewPolicyConnPool MockArgsTypenewPolicyConnPool

// AuxMocknewPolicyConnPool(argsession *Session)(reta *policyConnPool) - Generated mock function
func AuxMocknewPolicyConnPool(argsession *Session) (reta *policyConnPool) {
	LastMockArgsnewPolicyConnPool = MockArgsTypenewPolicyConnPool{
		ApomockCallNumber: AuxMockGetRecorderAuxMocknewPolicyConnPool(),
		Argsession:        argsession,
	}
	rargs, rerr := apomock.GetNext("gocql.newPolicyConnPool")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.newPolicyConnPool")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.newPolicyConnPool")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*policyConnPool)
	}
	return
}

// RecorderAuxMocknewPolicyConnPool  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocknewPolicyConnPool int = 0

var condRecorderAuxMocknewPolicyConnPool *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocknewPolicyConnPool(i int) {
	condRecorderAuxMocknewPolicyConnPool.L.Lock()
	for recorderAuxMocknewPolicyConnPool < i {
		condRecorderAuxMocknewPolicyConnPool.Wait()
	}
	condRecorderAuxMocknewPolicyConnPool.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocknewPolicyConnPool() {
	condRecorderAuxMocknewPolicyConnPool.L.Lock()
	recorderAuxMocknewPolicyConnPool++
	condRecorderAuxMocknewPolicyConnPool.L.Unlock()
	condRecorderAuxMocknewPolicyConnPool.Broadcast()
}
func AuxMockGetRecorderAuxMocknewPolicyConnPool() (ret int) {
	condRecorderAuxMocknewPolicyConnPool.L.Lock()
	ret = recorderAuxMocknewPolicyConnPool
	condRecorderAuxMocknewPolicyConnPool.L.Unlock()
	return
}

// newPolicyConnPool - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func newPolicyConnPool(argsession *Session) (reta *policyConnPool) {
	FuncAuxMocknewPolicyConnPool, ok := apomock.GetRegisteredFunc("gocql.newPolicyConnPool")
	if ok {
		reta = FuncAuxMocknewPolicyConnPool.(func(argsession *Session) (reta *policyConnPool))(argsession)
	} else {
		panic("FuncAuxMocknewPolicyConnPool ")
	}
	AuxMockIncrementRecorderAuxMocknewPolicyConnPool()
	return
}

//
// Mock: (recvp *policyConnPool)SetHosts(arghosts []*HostInfo)()
//

type MockArgsTypepolicyConnPoolSetHosts struct {
	ApomockCallNumber int
	Arghosts          []*HostInfo
}

var LastMockArgspolicyConnPoolSetHosts MockArgsTypepolicyConnPoolSetHosts

// (recvp *policyConnPool)AuxMockSetHosts(arghosts []*HostInfo)() - Generated mock function
func (recvp *policyConnPool) AuxMockSetHosts(arghosts []*HostInfo) {
	LastMockArgspolicyConnPoolSetHosts = MockArgsTypepolicyConnPoolSetHosts{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrpolicyConnPoolSetHosts(),
		Arghosts:          arghosts,
	}
	return
}

// RecorderAuxMockPtrpolicyConnPoolSetHosts  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpolicyConnPoolSetHosts int = 0

var condRecorderAuxMockPtrpolicyConnPoolSetHosts *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpolicyConnPoolSetHosts(i int) {
	condRecorderAuxMockPtrpolicyConnPoolSetHosts.L.Lock()
	for recorderAuxMockPtrpolicyConnPoolSetHosts < i {
		condRecorderAuxMockPtrpolicyConnPoolSetHosts.Wait()
	}
	condRecorderAuxMockPtrpolicyConnPoolSetHosts.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolSetHosts() {
	condRecorderAuxMockPtrpolicyConnPoolSetHosts.L.Lock()
	recorderAuxMockPtrpolicyConnPoolSetHosts++
	condRecorderAuxMockPtrpolicyConnPoolSetHosts.L.Unlock()
	condRecorderAuxMockPtrpolicyConnPoolSetHosts.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpolicyConnPoolSetHosts() (ret int) {
	condRecorderAuxMockPtrpolicyConnPoolSetHosts.L.Lock()
	ret = recorderAuxMockPtrpolicyConnPoolSetHosts
	condRecorderAuxMockPtrpolicyConnPoolSetHosts.L.Unlock()
	return
}

// (recvp *policyConnPool)SetHosts - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *policyConnPool) SetHosts(arghosts []*HostInfo) {
	FuncAuxMockPtrpolicyConnPoolSetHosts, ok := apomock.GetRegisteredFunc("gocql.policyConnPool.SetHosts")
	if ok {
		FuncAuxMockPtrpolicyConnPoolSetHosts.(func(recvp *policyConnPool, arghosts []*HostInfo))(recvp, arghosts)
	} else {
		panic("FuncAuxMockPtrpolicyConnPoolSetHosts ")
	}
	AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolSetHosts()
	return
}

//
// Mock: (recvpool *hostConnPool)Pick()(reta *Conn)
//

type MockArgsTypehostConnPoolPick struct {
	ApomockCallNumber int
}

var LastMockArgshostConnPoolPick MockArgsTypehostConnPoolPick

// (recvpool *hostConnPool)AuxMockPick()(reta *Conn) - Generated mock function
func (recvpool *hostConnPool) AuxMockPick() (reta *Conn) {
	rargs, rerr := apomock.GetNext("gocql.hostConnPool.Pick")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.hostConnPool.Pick")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.hostConnPool.Pick")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Conn)
	}
	return
}

// RecorderAuxMockPtrhostConnPoolPick  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPoolPick int = 0

var condRecorderAuxMockPtrhostConnPoolPick *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPoolPick(i int) {
	condRecorderAuxMockPtrhostConnPoolPick.L.Lock()
	for recorderAuxMockPtrhostConnPoolPick < i {
		condRecorderAuxMockPtrhostConnPoolPick.Wait()
	}
	condRecorderAuxMockPtrhostConnPoolPick.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPoolPick() {
	condRecorderAuxMockPtrhostConnPoolPick.L.Lock()
	recorderAuxMockPtrhostConnPoolPick++
	condRecorderAuxMockPtrhostConnPoolPick.L.Unlock()
	condRecorderAuxMockPtrhostConnPoolPick.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPoolPick() (ret int) {
	condRecorderAuxMockPtrhostConnPoolPick.L.Lock()
	ret = recorderAuxMockPtrhostConnPoolPick
	condRecorderAuxMockPtrhostConnPoolPick.L.Unlock()
	return
}

// (recvpool *hostConnPool)Pick - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvpool *hostConnPool) Pick() (reta *Conn) {
	FuncAuxMockPtrhostConnPoolPick, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.Pick")
	if ok {
		reta = FuncAuxMockPtrhostConnPoolPick.(func(recvpool *hostConnPool) (reta *Conn))(recvpool)
	} else {
		panic("FuncAuxMockPtrhostConnPoolPick ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPoolPick()
	return
}

//
// Mock: (recvpool *hostConnPool)fill()()
//

type MockArgsTypehostConnPoolfill struct {
	ApomockCallNumber int
}

var LastMockArgshostConnPoolfill MockArgsTypehostConnPoolfill

// (recvpool *hostConnPool)AuxMockfill()() - Generated mock function
func (recvpool *hostConnPool) AuxMockfill() {
	return
}

// RecorderAuxMockPtrhostConnPoolfill  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPoolfill int = 0

var condRecorderAuxMockPtrhostConnPoolfill *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPoolfill(i int) {
	condRecorderAuxMockPtrhostConnPoolfill.L.Lock()
	for recorderAuxMockPtrhostConnPoolfill < i {
		condRecorderAuxMockPtrhostConnPoolfill.Wait()
	}
	condRecorderAuxMockPtrhostConnPoolfill.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPoolfill() {
	condRecorderAuxMockPtrhostConnPoolfill.L.Lock()
	recorderAuxMockPtrhostConnPoolfill++
	condRecorderAuxMockPtrhostConnPoolfill.L.Unlock()
	condRecorderAuxMockPtrhostConnPoolfill.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPoolfill() (ret int) {
	condRecorderAuxMockPtrhostConnPoolfill.L.Lock()
	ret = recorderAuxMockPtrhostConnPoolfill
	condRecorderAuxMockPtrhostConnPoolfill.L.Unlock()
	return
}

// (recvpool *hostConnPool)fill - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvpool *hostConnPool) fill() {
	FuncAuxMockPtrhostConnPoolfill, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.fill")
	if ok {
		FuncAuxMockPtrhostConnPoolfill.(func(recvpool *hostConnPool))(recvpool)
	} else {
		panic("FuncAuxMockPtrhostConnPoolfill ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPoolfill()
	return
}

//
// Mock: (recvpool *hostConnPool)fillingStopped(arghadError bool)()
//

type MockArgsTypehostConnPoolfillingStopped struct {
	ApomockCallNumber int
	ArghadError       bool
}

var LastMockArgshostConnPoolfillingStopped MockArgsTypehostConnPoolfillingStopped

// (recvpool *hostConnPool)AuxMockfillingStopped(arghadError bool)() - Generated mock function
func (recvpool *hostConnPool) AuxMockfillingStopped(arghadError bool) {
	LastMockArgshostConnPoolfillingStopped = MockArgsTypehostConnPoolfillingStopped{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrhostConnPoolfillingStopped(),
		ArghadError:       arghadError,
	}
	return
}

// RecorderAuxMockPtrhostConnPoolfillingStopped  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPoolfillingStopped int = 0

var condRecorderAuxMockPtrhostConnPoolfillingStopped *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPoolfillingStopped(i int) {
	condRecorderAuxMockPtrhostConnPoolfillingStopped.L.Lock()
	for recorderAuxMockPtrhostConnPoolfillingStopped < i {
		condRecorderAuxMockPtrhostConnPoolfillingStopped.Wait()
	}
	condRecorderAuxMockPtrhostConnPoolfillingStopped.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPoolfillingStopped() {
	condRecorderAuxMockPtrhostConnPoolfillingStopped.L.Lock()
	recorderAuxMockPtrhostConnPoolfillingStopped++
	condRecorderAuxMockPtrhostConnPoolfillingStopped.L.Unlock()
	condRecorderAuxMockPtrhostConnPoolfillingStopped.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPoolfillingStopped() (ret int) {
	condRecorderAuxMockPtrhostConnPoolfillingStopped.L.Lock()
	ret = recorderAuxMockPtrhostConnPoolfillingStopped
	condRecorderAuxMockPtrhostConnPoolfillingStopped.L.Unlock()
	return
}

// (recvpool *hostConnPool)fillingStopped - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvpool *hostConnPool) fillingStopped(arghadError bool) {
	FuncAuxMockPtrhostConnPoolfillingStopped, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.fillingStopped")
	if ok {
		FuncAuxMockPtrhostConnPoolfillingStopped.(func(recvpool *hostConnPool, arghadError bool))(recvpool, arghadError)
	} else {
		panic("FuncAuxMockPtrhostConnPoolfillingStopped ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPoolfillingStopped()
	return
}

//
// Mock: (recvpool *hostConnPool)HandleError(argconn *Conn, argerr error, argclosed bool)()
//

type MockArgsTypehostConnPoolHandleError struct {
	ApomockCallNumber int
	Argconn           *Conn
	Argerr            error
	Argclosed         bool
}

var LastMockArgshostConnPoolHandleError MockArgsTypehostConnPoolHandleError

// (recvpool *hostConnPool)AuxMockHandleError(argconn *Conn, argerr error, argclosed bool)() - Generated mock function
func (recvpool *hostConnPool) AuxMockHandleError(argconn *Conn, argerr error, argclosed bool) {
	LastMockArgshostConnPoolHandleError = MockArgsTypehostConnPoolHandleError{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrhostConnPoolHandleError(),
		Argconn:           argconn,
		Argerr:            argerr,
		Argclosed:         argclosed,
	}
	return
}

// RecorderAuxMockPtrhostConnPoolHandleError  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPoolHandleError int = 0

var condRecorderAuxMockPtrhostConnPoolHandleError *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPoolHandleError(i int) {
	condRecorderAuxMockPtrhostConnPoolHandleError.L.Lock()
	for recorderAuxMockPtrhostConnPoolHandleError < i {
		condRecorderAuxMockPtrhostConnPoolHandleError.Wait()
	}
	condRecorderAuxMockPtrhostConnPoolHandleError.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPoolHandleError() {
	condRecorderAuxMockPtrhostConnPoolHandleError.L.Lock()
	recorderAuxMockPtrhostConnPoolHandleError++
	condRecorderAuxMockPtrhostConnPoolHandleError.L.Unlock()
	condRecorderAuxMockPtrhostConnPoolHandleError.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPoolHandleError() (ret int) {
	condRecorderAuxMockPtrhostConnPoolHandleError.L.Lock()
	ret = recorderAuxMockPtrhostConnPoolHandleError
	condRecorderAuxMockPtrhostConnPoolHandleError.L.Unlock()
	return
}

// (recvpool *hostConnPool)HandleError - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvpool *hostConnPool) HandleError(argconn *Conn, argerr error, argclosed bool) {
	FuncAuxMockPtrhostConnPoolHandleError, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.HandleError")
	if ok {
		FuncAuxMockPtrhostConnPoolHandleError.(func(recvpool *hostConnPool, argconn *Conn, argerr error, argclosed bool))(recvpool, argconn, argerr, argclosed)
	} else {
		panic("FuncAuxMockPtrhostConnPoolHandleError ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPoolHandleError()
	return
}

//
// Mock: connConfig(argsession *Session)(reta *ConnConfig, retb error)
//

type MockArgsTypeconnConfig struct {
	ApomockCallNumber int
	Argsession        *Session
}

var LastMockArgsconnConfig MockArgsTypeconnConfig

// AuxMockconnConfig(argsession *Session)(reta *ConnConfig, retb error) - Generated mock function
func AuxMockconnConfig(argsession *Session) (reta *ConnConfig, retb error) {
	LastMockArgsconnConfig = MockArgsTypeconnConfig{
		ApomockCallNumber: AuxMockGetRecorderAuxMockconnConfig(),
		Argsession:        argsession,
	}
	rargs, rerr := apomock.GetNext("gocql.connConfig")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.connConfig")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.connConfig")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*ConnConfig)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockconnConfig  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockconnConfig int = 0

var condRecorderAuxMockconnConfig *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockconnConfig(i int) {
	condRecorderAuxMockconnConfig.L.Lock()
	for recorderAuxMockconnConfig < i {
		condRecorderAuxMockconnConfig.Wait()
	}
	condRecorderAuxMockconnConfig.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockconnConfig() {
	condRecorderAuxMockconnConfig.L.Lock()
	recorderAuxMockconnConfig++
	condRecorderAuxMockconnConfig.L.Unlock()
	condRecorderAuxMockconnConfig.Broadcast()
}
func AuxMockGetRecorderAuxMockconnConfig() (ret int) {
	condRecorderAuxMockconnConfig.L.Lock()
	ret = recorderAuxMockconnConfig
	condRecorderAuxMockconnConfig.L.Unlock()
	return
}

// connConfig - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func connConfig(argsession *Session) (reta *ConnConfig, retb error) {
	FuncAuxMockconnConfig, ok := apomock.GetRegisteredFunc("gocql.connConfig")
	if ok {
		reta, retb = FuncAuxMockconnConfig.(func(argsession *Session) (reta *ConnConfig, retb error))(argsession)
	} else {
		panic("FuncAuxMockconnConfig ")
	}
	AuxMockIncrementRecorderAuxMockconnConfig()
	return
}

//
// Mock: (recvp *policyConnPool)Close()()
//

type MockArgsTypepolicyConnPoolClose struct {
	ApomockCallNumber int
}

var LastMockArgspolicyConnPoolClose MockArgsTypepolicyConnPoolClose

// (recvp *policyConnPool)AuxMockClose()() - Generated mock function
func (recvp *policyConnPool) AuxMockClose() {
	return
}

// RecorderAuxMockPtrpolicyConnPoolClose  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpolicyConnPoolClose int = 0

var condRecorderAuxMockPtrpolicyConnPoolClose *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpolicyConnPoolClose(i int) {
	condRecorderAuxMockPtrpolicyConnPoolClose.L.Lock()
	for recorderAuxMockPtrpolicyConnPoolClose < i {
		condRecorderAuxMockPtrpolicyConnPoolClose.Wait()
	}
	condRecorderAuxMockPtrpolicyConnPoolClose.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolClose() {
	condRecorderAuxMockPtrpolicyConnPoolClose.L.Lock()
	recorderAuxMockPtrpolicyConnPoolClose++
	condRecorderAuxMockPtrpolicyConnPoolClose.L.Unlock()
	condRecorderAuxMockPtrpolicyConnPoolClose.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpolicyConnPoolClose() (ret int) {
	condRecorderAuxMockPtrpolicyConnPoolClose.L.Lock()
	ret = recorderAuxMockPtrpolicyConnPoolClose
	condRecorderAuxMockPtrpolicyConnPoolClose.L.Unlock()
	return
}

// (recvp *policyConnPool)Close - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *policyConnPool) Close() {
	FuncAuxMockPtrpolicyConnPoolClose, ok := apomock.GetRegisteredFunc("gocql.policyConnPool.Close")
	if ok {
		FuncAuxMockPtrpolicyConnPoolClose.(func(recvp *policyConnPool))(recvp)
	} else {
		panic("FuncAuxMockPtrpolicyConnPoolClose ")
	}
	AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolClose()
	return
}

//
// Mock: (recvp *policyConnPool)hostDown(argaddr string)()
//

type MockArgsTypepolicyConnPoolhostDown struct {
	ApomockCallNumber int
	Argaddr           string
}

var LastMockArgspolicyConnPoolhostDown MockArgsTypepolicyConnPoolhostDown

// (recvp *policyConnPool)AuxMockhostDown(argaddr string)() - Generated mock function
func (recvp *policyConnPool) AuxMockhostDown(argaddr string) {
	LastMockArgspolicyConnPoolhostDown = MockArgsTypepolicyConnPoolhostDown{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrpolicyConnPoolhostDown(),
		Argaddr:           argaddr,
	}
	return
}

// RecorderAuxMockPtrpolicyConnPoolhostDown  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpolicyConnPoolhostDown int = 0

var condRecorderAuxMockPtrpolicyConnPoolhostDown *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpolicyConnPoolhostDown(i int) {
	condRecorderAuxMockPtrpolicyConnPoolhostDown.L.Lock()
	for recorderAuxMockPtrpolicyConnPoolhostDown < i {
		condRecorderAuxMockPtrpolicyConnPoolhostDown.Wait()
	}
	condRecorderAuxMockPtrpolicyConnPoolhostDown.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolhostDown() {
	condRecorderAuxMockPtrpolicyConnPoolhostDown.L.Lock()
	recorderAuxMockPtrpolicyConnPoolhostDown++
	condRecorderAuxMockPtrpolicyConnPoolhostDown.L.Unlock()
	condRecorderAuxMockPtrpolicyConnPoolhostDown.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpolicyConnPoolhostDown() (ret int) {
	condRecorderAuxMockPtrpolicyConnPoolhostDown.L.Lock()
	ret = recorderAuxMockPtrpolicyConnPoolhostDown
	condRecorderAuxMockPtrpolicyConnPoolhostDown.L.Unlock()
	return
}

// (recvp *policyConnPool)hostDown - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *policyConnPool) hostDown(argaddr string) {
	FuncAuxMockPtrpolicyConnPoolhostDown, ok := apomock.GetRegisteredFunc("gocql.policyConnPool.hostDown")
	if ok {
		FuncAuxMockPtrpolicyConnPoolhostDown.(func(recvp *policyConnPool, argaddr string))(recvp, argaddr)
	} else {
		panic("FuncAuxMockPtrpolicyConnPoolhostDown ")
	}
	AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolhostDown()
	return
}

//
// Mock: (recvpool *hostConnPool)connect()(reterr error)
//

type MockArgsTypehostConnPoolconnect struct {
	ApomockCallNumber int
}

var LastMockArgshostConnPoolconnect MockArgsTypehostConnPoolconnect

// (recvpool *hostConnPool)AuxMockconnect()(reterr error) - Generated mock function
func (recvpool *hostConnPool) AuxMockconnect() (reterr error) {
	rargs, rerr := apomock.GetNext("gocql.hostConnPool.connect")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.hostConnPool.connect")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.hostConnPool.connect")
	}
	if rargs.GetArg(0) != nil {
		reterr = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrhostConnPoolconnect  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPoolconnect int = 0

var condRecorderAuxMockPtrhostConnPoolconnect *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPoolconnect(i int) {
	condRecorderAuxMockPtrhostConnPoolconnect.L.Lock()
	for recorderAuxMockPtrhostConnPoolconnect < i {
		condRecorderAuxMockPtrhostConnPoolconnect.Wait()
	}
	condRecorderAuxMockPtrhostConnPoolconnect.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPoolconnect() {
	condRecorderAuxMockPtrhostConnPoolconnect.L.Lock()
	recorderAuxMockPtrhostConnPoolconnect++
	condRecorderAuxMockPtrhostConnPoolconnect.L.Unlock()
	condRecorderAuxMockPtrhostConnPoolconnect.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPoolconnect() (ret int) {
	condRecorderAuxMockPtrhostConnPoolconnect.L.Lock()
	ret = recorderAuxMockPtrhostConnPoolconnect
	condRecorderAuxMockPtrhostConnPoolconnect.L.Unlock()
	return
}

// (recvpool *hostConnPool)connect - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvpool *hostConnPool) connect() (reterr error) {
	FuncAuxMockPtrhostConnPoolconnect, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.connect")
	if ok {
		reterr = FuncAuxMockPtrhostConnPoolconnect.(func(recvpool *hostConnPool) (reterr error))(recvpool)
	} else {
		panic("FuncAuxMockPtrhostConnPoolconnect ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPoolconnect()
	return
}

//
// Mock: (recvpool *hostConnPool)drainLocked()()
//

type MockArgsTypehostConnPooldrainLocked struct {
	ApomockCallNumber int
}

var LastMockArgshostConnPooldrainLocked MockArgsTypehostConnPooldrainLocked

// (recvpool *hostConnPool)AuxMockdrainLocked()() - Generated mock function
func (recvpool *hostConnPool) AuxMockdrainLocked() {
	return
}

// RecorderAuxMockPtrhostConnPooldrainLocked  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPooldrainLocked int = 0

var condRecorderAuxMockPtrhostConnPooldrainLocked *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPooldrainLocked(i int) {
	condRecorderAuxMockPtrhostConnPooldrainLocked.L.Lock()
	for recorderAuxMockPtrhostConnPooldrainLocked < i {
		condRecorderAuxMockPtrhostConnPooldrainLocked.Wait()
	}
	condRecorderAuxMockPtrhostConnPooldrainLocked.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPooldrainLocked() {
	condRecorderAuxMockPtrhostConnPooldrainLocked.L.Lock()
	recorderAuxMockPtrhostConnPooldrainLocked++
	condRecorderAuxMockPtrhostConnPooldrainLocked.L.Unlock()
	condRecorderAuxMockPtrhostConnPooldrainLocked.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPooldrainLocked() (ret int) {
	condRecorderAuxMockPtrhostConnPooldrainLocked.L.Lock()
	ret = recorderAuxMockPtrhostConnPooldrainLocked
	condRecorderAuxMockPtrhostConnPooldrainLocked.L.Unlock()
	return
}

// (recvpool *hostConnPool)drainLocked - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvpool *hostConnPool) drainLocked() {
	FuncAuxMockPtrhostConnPooldrainLocked, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.drainLocked")
	if ok {
		FuncAuxMockPtrhostConnPooldrainLocked.(func(recvpool *hostConnPool))(recvpool)
	} else {
		panic("FuncAuxMockPtrhostConnPooldrainLocked ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPooldrainLocked()
	return
}

//
// Mock: (recvp *policyConnPool)hostUp(arghost *HostInfo)()
//

type MockArgsTypepolicyConnPoolhostUp struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgspolicyConnPoolhostUp MockArgsTypepolicyConnPoolhostUp

// (recvp *policyConnPool)AuxMockhostUp(arghost *HostInfo)() - Generated mock function
func (recvp *policyConnPool) AuxMockhostUp(arghost *HostInfo) {
	LastMockArgspolicyConnPoolhostUp = MockArgsTypepolicyConnPoolhostUp{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrpolicyConnPoolhostUp(),
		Arghost:           arghost,
	}
	return
}

// RecorderAuxMockPtrpolicyConnPoolhostUp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrpolicyConnPoolhostUp int = 0

var condRecorderAuxMockPtrpolicyConnPoolhostUp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrpolicyConnPoolhostUp(i int) {
	condRecorderAuxMockPtrpolicyConnPoolhostUp.L.Lock()
	for recorderAuxMockPtrpolicyConnPoolhostUp < i {
		condRecorderAuxMockPtrpolicyConnPoolhostUp.Wait()
	}
	condRecorderAuxMockPtrpolicyConnPoolhostUp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolhostUp() {
	condRecorderAuxMockPtrpolicyConnPoolhostUp.L.Lock()
	recorderAuxMockPtrpolicyConnPoolhostUp++
	condRecorderAuxMockPtrpolicyConnPoolhostUp.L.Unlock()
	condRecorderAuxMockPtrpolicyConnPoolhostUp.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrpolicyConnPoolhostUp() (ret int) {
	condRecorderAuxMockPtrpolicyConnPoolhostUp.L.Lock()
	ret = recorderAuxMockPtrpolicyConnPoolhostUp
	condRecorderAuxMockPtrpolicyConnPoolhostUp.L.Unlock()
	return
}

// (recvp *policyConnPool)hostUp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp *policyConnPool) hostUp(arghost *HostInfo) {
	FuncAuxMockPtrpolicyConnPoolhostUp, ok := apomock.GetRegisteredFunc("gocql.policyConnPool.hostUp")
	if ok {
		FuncAuxMockPtrpolicyConnPoolhostUp.(func(recvp *policyConnPool, arghost *HostInfo))(recvp, arghost)
	} else {
		panic("FuncAuxMockPtrpolicyConnPoolhostUp ")
	}
	AuxMockIncrementRecorderAuxMockPtrpolicyConnPoolhostUp()
	return
}

//
// Mock: newHostConnPool(argsession *Session, arghost *HostInfo, argport int, argsize int, argkeyspace string)(reta *hostConnPool)
//

type MockArgsTypenewHostConnPool struct {
	ApomockCallNumber int
	Argsession        *Session
	Arghost           *HostInfo
	Argport           int
	Argsize           int
	Argkeyspace       string
}

var LastMockArgsnewHostConnPool MockArgsTypenewHostConnPool

// AuxMocknewHostConnPool(argsession *Session, arghost *HostInfo, argport int, argsize int, argkeyspace string)(reta *hostConnPool) - Generated mock function
func AuxMocknewHostConnPool(argsession *Session, arghost *HostInfo, argport int, argsize int, argkeyspace string) (reta *hostConnPool) {
	LastMockArgsnewHostConnPool = MockArgsTypenewHostConnPool{
		ApomockCallNumber: AuxMockGetRecorderAuxMocknewHostConnPool(),
		Argsession:        argsession,
		Arghost:           arghost,
		Argport:           argport,
		Argsize:           argsize,
		Argkeyspace:       argkeyspace,
	}
	rargs, rerr := apomock.GetNext("gocql.newHostConnPool")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.newHostConnPool")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.newHostConnPool")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*hostConnPool)
	}
	return
}

// RecorderAuxMocknewHostConnPool  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocknewHostConnPool int = 0

var condRecorderAuxMocknewHostConnPool *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocknewHostConnPool(i int) {
	condRecorderAuxMocknewHostConnPool.L.Lock()
	for recorderAuxMocknewHostConnPool < i {
		condRecorderAuxMocknewHostConnPool.Wait()
	}
	condRecorderAuxMocknewHostConnPool.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocknewHostConnPool() {
	condRecorderAuxMocknewHostConnPool.L.Lock()
	recorderAuxMocknewHostConnPool++
	condRecorderAuxMocknewHostConnPool.L.Unlock()
	condRecorderAuxMocknewHostConnPool.Broadcast()
}
func AuxMockGetRecorderAuxMocknewHostConnPool() (ret int) {
	condRecorderAuxMocknewHostConnPool.L.Lock()
	ret = recorderAuxMocknewHostConnPool
	condRecorderAuxMocknewHostConnPool.L.Unlock()
	return
}

// newHostConnPool - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func newHostConnPool(argsession *Session, arghost *HostInfo, argport int, argsize int, argkeyspace string) (reta *hostConnPool) {
	FuncAuxMocknewHostConnPool, ok := apomock.GetRegisteredFunc("gocql.newHostConnPool")
	if ok {
		reta = FuncAuxMocknewHostConnPool.(func(argsession *Session, arghost *HostInfo, argport int, argsize int, argkeyspace string) (reta *hostConnPool))(argsession, arghost, argport, argsize, argkeyspace)
	} else {
		panic("FuncAuxMocknewHostConnPool ")
	}
	AuxMockIncrementRecorderAuxMocknewHostConnPool()
	return
}

//
// Mock: (recvpool *hostConnPool)Size()(reta int)
//

type MockArgsTypehostConnPoolSize struct {
	ApomockCallNumber int
}

var LastMockArgshostConnPoolSize MockArgsTypehostConnPoolSize

// (recvpool *hostConnPool)AuxMockSize()(reta int) - Generated mock function
func (recvpool *hostConnPool) AuxMockSize() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.hostConnPool.Size")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.hostConnPool.Size")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.hostConnPool.Size")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrhostConnPoolSize  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPoolSize int = 0

var condRecorderAuxMockPtrhostConnPoolSize *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPoolSize(i int) {
	condRecorderAuxMockPtrhostConnPoolSize.L.Lock()
	for recorderAuxMockPtrhostConnPoolSize < i {
		condRecorderAuxMockPtrhostConnPoolSize.Wait()
	}
	condRecorderAuxMockPtrhostConnPoolSize.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPoolSize() {
	condRecorderAuxMockPtrhostConnPoolSize.L.Lock()
	recorderAuxMockPtrhostConnPoolSize++
	condRecorderAuxMockPtrhostConnPoolSize.L.Unlock()
	condRecorderAuxMockPtrhostConnPoolSize.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPoolSize() (ret int) {
	condRecorderAuxMockPtrhostConnPoolSize.L.Lock()
	ret = recorderAuxMockPtrhostConnPoolSize
	condRecorderAuxMockPtrhostConnPoolSize.L.Unlock()
	return
}

// (recvpool *hostConnPool)Size - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvpool *hostConnPool) Size() (reta int) {
	FuncAuxMockPtrhostConnPoolSize, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.Size")
	if ok {
		reta = FuncAuxMockPtrhostConnPoolSize.(func(recvpool *hostConnPool) (reta int))(recvpool)
	} else {
		panic("FuncAuxMockPtrhostConnPoolSize ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPoolSize()
	return
}

//
// Mock: (recvpool *hostConnPool)Close()()
//

type MockArgsTypehostConnPoolClose struct {
	ApomockCallNumber int
}

var LastMockArgshostConnPoolClose MockArgsTypehostConnPoolClose

// (recvpool *hostConnPool)AuxMockClose()() - Generated mock function
func (recvpool *hostConnPool) AuxMockClose() {
	return
}

// RecorderAuxMockPtrhostConnPoolClose  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPoolClose int = 0

var condRecorderAuxMockPtrhostConnPoolClose *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPoolClose(i int) {
	condRecorderAuxMockPtrhostConnPoolClose.L.Lock()
	for recorderAuxMockPtrhostConnPoolClose < i {
		condRecorderAuxMockPtrhostConnPoolClose.Wait()
	}
	condRecorderAuxMockPtrhostConnPoolClose.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPoolClose() {
	condRecorderAuxMockPtrhostConnPoolClose.L.Lock()
	recorderAuxMockPtrhostConnPoolClose++
	condRecorderAuxMockPtrhostConnPoolClose.L.Unlock()
	condRecorderAuxMockPtrhostConnPoolClose.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPoolClose() (ret int) {
	condRecorderAuxMockPtrhostConnPoolClose.L.Lock()
	ret = recorderAuxMockPtrhostConnPoolClose
	condRecorderAuxMockPtrhostConnPoolClose.L.Unlock()
	return
}

// (recvpool *hostConnPool)Close - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvpool *hostConnPool) Close() {
	FuncAuxMockPtrhostConnPoolClose, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.Close")
	if ok {
		FuncAuxMockPtrhostConnPoolClose.(func(recvpool *hostConnPool))(recvpool)
	} else {
		panic("FuncAuxMockPtrhostConnPoolClose ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPoolClose()
	return
}

//
// Mock: (recvpool *hostConnPool)drain()()
//

type MockArgsTypehostConnPooldrain struct {
	ApomockCallNumber int
}

var LastMockArgshostConnPooldrain MockArgsTypehostConnPooldrain

// (recvpool *hostConnPool)AuxMockdrain()() - Generated mock function
func (recvpool *hostConnPool) AuxMockdrain() {
	return
}

// RecorderAuxMockPtrhostConnPooldrain  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrhostConnPooldrain int = 0

var condRecorderAuxMockPtrhostConnPooldrain *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrhostConnPooldrain(i int) {
	condRecorderAuxMockPtrhostConnPooldrain.L.Lock()
	for recorderAuxMockPtrhostConnPooldrain < i {
		condRecorderAuxMockPtrhostConnPooldrain.Wait()
	}
	condRecorderAuxMockPtrhostConnPooldrain.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrhostConnPooldrain() {
	condRecorderAuxMockPtrhostConnPooldrain.L.Lock()
	recorderAuxMockPtrhostConnPooldrain++
	condRecorderAuxMockPtrhostConnPooldrain.L.Unlock()
	condRecorderAuxMockPtrhostConnPooldrain.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrhostConnPooldrain() (ret int) {
	condRecorderAuxMockPtrhostConnPooldrain.L.Lock()
	ret = recorderAuxMockPtrhostConnPooldrain
	condRecorderAuxMockPtrhostConnPooldrain.L.Unlock()
	return
}

// (recvpool *hostConnPool)drain - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvpool *hostConnPool) drain() {
	FuncAuxMockPtrhostConnPooldrain, ok := apomock.GetRegisteredFunc("gocql.hostConnPool.drain")
	if ok {
		FuncAuxMockPtrhostConnPooldrain.(func(recvpool *hostConnPool))(recvpool)
	} else {
		panic("FuncAuxMockPtrhostConnPooldrain ")
	}
	AuxMockIncrementRecorderAuxMockPtrhostConnPooldrain()
	return
}
