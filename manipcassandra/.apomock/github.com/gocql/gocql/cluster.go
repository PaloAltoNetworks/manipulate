// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"
import "errors"
import "time"

func init() {

	apomock.RegisterFunc("gocql", "gocql.PoolConfig.buildPool", (PoolConfig).AuxMockbuildPool)
	apomock.RegisterFunc("gocql", "gocql.DiscoveryConfig.matchFilter", (DiscoveryConfig).AuxMockmatchFilter)
	apomock.RegisterFunc("gocql", "gocql.NewCluster", AuxMockNewCluster)
	apomock.RegisterFunc("gocql", "gocql.ClusterConfig.CreateSession", (*ClusterConfig).AuxMockCreateSession)
}

const ()

var (
	ErrNoHosts              = errors.New("no hosts provided")
	ErrNoConnectionsStarted = errors.New("no connections were made when creating the session")
	ErrHostQueryFailed      = errors.New("unable to populate Hosts")
)

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//
type PoolConfig struct{ HostSelectionPolicy HostSelectionPolicy }

type DiscoveryConfig struct {
	DcFilter   string
	RackFilter string
	Sleep      time.Duration
}

type ClusterConfig struct {
	Hosts                    []string
	CQLVersion               string
	ProtoVersion             int
	Timeout                  time.Duration
	Port                     int
	Keyspace                 string
	NumConns                 int
	Consistency              Consistency
	Compressor               Compressor
	Authenticator            Authenticator
	RetryPolicy              RetryPolicy
	SocketKeepalive          time.Duration
	MaxPreparedStmts         int
	MaxRoutingKeyInfo        int
	PageSize                 int
	SerialConsistency        SerialConsistency
	SslOpts                  *SslOptions
	DefaultTimestamp         bool
	PoolConfig               PoolConfig
	Discovery                DiscoveryConfig
	ReconnectInterval        time.Duration
	MaxWaitSchemaAgreement   time.Duration
	HostFilter               HostFilter
	IgnorePeerAddr           bool
	DisableInitialHostLookup bool
	Events                   struct {
		DisableNodeStatusEvents bool
		DisableTopologyEvents   bool
		DisableSchemaEvents     bool
	}
	disableControlConn bool
}

//
// Mock: (recvp PoolConfig)buildPool(argsession *Session)(reta *policyConnPool)
//

type MockArgsTypePoolConfigbuildPool struct {
	ApomockCallNumber int
	Argsession        *Session
}

var LastMockArgsPoolConfigbuildPool MockArgsTypePoolConfigbuildPool

// (recvp PoolConfig)AuxMockbuildPool(argsession *Session)(reta *policyConnPool) - Generated mock function
func (recvp PoolConfig) AuxMockbuildPool(argsession *Session) (reta *policyConnPool) {
	LastMockArgsPoolConfigbuildPool = MockArgsTypePoolConfigbuildPool{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPoolConfigbuildPool(),
		Argsession:        argsession,
	}
	rargs, rerr := apomock.GetNext("gocql.PoolConfig.buildPool")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.PoolConfig.buildPool")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.PoolConfig.buildPool")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*policyConnPool)
	}
	return
}

// RecorderAuxMockPoolConfigbuildPool  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPoolConfigbuildPool int = 0

var condRecorderAuxMockPoolConfigbuildPool *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPoolConfigbuildPool(i int) {
	condRecorderAuxMockPoolConfigbuildPool.L.Lock()
	for recorderAuxMockPoolConfigbuildPool < i {
		condRecorderAuxMockPoolConfigbuildPool.Wait()
	}
	condRecorderAuxMockPoolConfigbuildPool.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPoolConfigbuildPool() {
	condRecorderAuxMockPoolConfigbuildPool.L.Lock()
	recorderAuxMockPoolConfigbuildPool++
	condRecorderAuxMockPoolConfigbuildPool.L.Unlock()
	condRecorderAuxMockPoolConfigbuildPool.Broadcast()
}
func AuxMockGetRecorderAuxMockPoolConfigbuildPool() (ret int) {
	condRecorderAuxMockPoolConfigbuildPool.L.Lock()
	ret = recorderAuxMockPoolConfigbuildPool
	condRecorderAuxMockPoolConfigbuildPool.L.Unlock()
	return
}

// (recvp PoolConfig)buildPool - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp PoolConfig) buildPool(argsession *Session) (reta *policyConnPool) {
	FuncAuxMockPoolConfigbuildPool, ok := apomock.GetRegisteredFunc("gocql.PoolConfig.buildPool")
	if ok {
		reta = FuncAuxMockPoolConfigbuildPool.(func(recvp PoolConfig, argsession *Session) (reta *policyConnPool))(recvp, argsession)
	} else {
		panic("FuncAuxMockPoolConfigbuildPool ")
	}
	AuxMockIncrementRecorderAuxMockPoolConfigbuildPool()
	return
}

//
// Mock: (recvd DiscoveryConfig)matchFilter(arghost *HostInfo)(reta bool)
//

type MockArgsTypeDiscoveryConfigmatchFilter struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgsDiscoveryConfigmatchFilter MockArgsTypeDiscoveryConfigmatchFilter

// (recvd DiscoveryConfig)AuxMockmatchFilter(arghost *HostInfo)(reta bool) - Generated mock function
func (recvd DiscoveryConfig) AuxMockmatchFilter(arghost *HostInfo) (reta bool) {
	LastMockArgsDiscoveryConfigmatchFilter = MockArgsTypeDiscoveryConfigmatchFilter{
		ApomockCallNumber: AuxMockGetRecorderAuxMockDiscoveryConfigmatchFilter(),
		Arghost:           arghost,
	}
	rargs, rerr := apomock.GetNext("gocql.DiscoveryConfig.matchFilter")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.DiscoveryConfig.matchFilter")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.DiscoveryConfig.matchFilter")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockDiscoveryConfigmatchFilter  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockDiscoveryConfigmatchFilter int = 0

var condRecorderAuxMockDiscoveryConfigmatchFilter *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockDiscoveryConfigmatchFilter(i int) {
	condRecorderAuxMockDiscoveryConfigmatchFilter.L.Lock()
	for recorderAuxMockDiscoveryConfigmatchFilter < i {
		condRecorderAuxMockDiscoveryConfigmatchFilter.Wait()
	}
	condRecorderAuxMockDiscoveryConfigmatchFilter.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockDiscoveryConfigmatchFilter() {
	condRecorderAuxMockDiscoveryConfigmatchFilter.L.Lock()
	recorderAuxMockDiscoveryConfigmatchFilter++
	condRecorderAuxMockDiscoveryConfigmatchFilter.L.Unlock()
	condRecorderAuxMockDiscoveryConfigmatchFilter.Broadcast()
}
func AuxMockGetRecorderAuxMockDiscoveryConfigmatchFilter() (ret int) {
	condRecorderAuxMockDiscoveryConfigmatchFilter.L.Lock()
	ret = recorderAuxMockDiscoveryConfigmatchFilter
	condRecorderAuxMockDiscoveryConfigmatchFilter.L.Unlock()
	return
}

// (recvd DiscoveryConfig)matchFilter - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvd DiscoveryConfig) matchFilter(arghost *HostInfo) (reta bool) {
	FuncAuxMockDiscoveryConfigmatchFilter, ok := apomock.GetRegisteredFunc("gocql.DiscoveryConfig.matchFilter")
	if ok {
		reta = FuncAuxMockDiscoveryConfigmatchFilter.(func(recvd DiscoveryConfig, arghost *HostInfo) (reta bool))(recvd, arghost)
	} else {
		panic("FuncAuxMockDiscoveryConfigmatchFilter ")
	}
	AuxMockIncrementRecorderAuxMockDiscoveryConfigmatchFilter()
	return
}

//
// Mock: NewCluster(hosts ...string)(reta *ClusterConfig)
//

type MockArgsTypeNewCluster struct {
	ApomockCallNumber int
	Hosts             []string
}

var LastMockArgsNewCluster MockArgsTypeNewCluster

// AuxMockNewCluster(hosts ...string)(reta *ClusterConfig) - Generated mock function
func AuxMockNewCluster(hosts ...string) (reta *ClusterConfig) {
	LastMockArgsNewCluster = MockArgsTypeNewCluster{
		ApomockCallNumber: AuxMockGetRecorderAuxMockNewCluster(),
		Hosts:             hosts,
	}
	rargs, rerr := apomock.GetNext("gocql.NewCluster")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.NewCluster")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.NewCluster")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*ClusterConfig)
	}
	return
}

// RecorderAuxMockNewCluster  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNewCluster int = 0

var condRecorderAuxMockNewCluster *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNewCluster(i int) {
	condRecorderAuxMockNewCluster.L.Lock()
	for recorderAuxMockNewCluster < i {
		condRecorderAuxMockNewCluster.Wait()
	}
	condRecorderAuxMockNewCluster.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNewCluster() {
	condRecorderAuxMockNewCluster.L.Lock()
	recorderAuxMockNewCluster++
	condRecorderAuxMockNewCluster.L.Unlock()
	condRecorderAuxMockNewCluster.Broadcast()
}
func AuxMockGetRecorderAuxMockNewCluster() (ret int) {
	condRecorderAuxMockNewCluster.L.Lock()
	ret = recorderAuxMockNewCluster
	condRecorderAuxMockNewCluster.L.Unlock()
	return
}

// NewCluster - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func NewCluster(hosts ...string) (reta *ClusterConfig) {
	FuncAuxMockNewCluster, ok := apomock.GetRegisteredFunc("gocql.NewCluster")
	if ok {
		reta = FuncAuxMockNewCluster.(func(hosts ...string) (reta *ClusterConfig))(hosts...)
	} else {
		panic("FuncAuxMockNewCluster ")
	}
	AuxMockIncrementRecorderAuxMockNewCluster()
	return
}

//
// Mock: (recvcfg *ClusterConfig)CreateSession()(reta *Session, retb error)
//

type MockArgsTypeClusterConfigCreateSession struct {
	ApomockCallNumber int
}

var LastMockArgsClusterConfigCreateSession MockArgsTypeClusterConfigCreateSession

// (recvcfg *ClusterConfig)AuxMockCreateSession()(reta *Session, retb error) - Generated mock function
func (recvcfg *ClusterConfig) AuxMockCreateSession() (reta *Session, retb error) {
	rargs, rerr := apomock.GetNext("gocql.ClusterConfig.CreateSession")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ClusterConfig.CreateSession")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.ClusterConfig.CreateSession")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Session)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrClusterConfigCreateSession  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrClusterConfigCreateSession int = 0

var condRecorderAuxMockPtrClusterConfigCreateSession *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrClusterConfigCreateSession(i int) {
	condRecorderAuxMockPtrClusterConfigCreateSession.L.Lock()
	for recorderAuxMockPtrClusterConfigCreateSession < i {
		condRecorderAuxMockPtrClusterConfigCreateSession.Wait()
	}
	condRecorderAuxMockPtrClusterConfigCreateSession.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrClusterConfigCreateSession() {
	condRecorderAuxMockPtrClusterConfigCreateSession.L.Lock()
	recorderAuxMockPtrClusterConfigCreateSession++
	condRecorderAuxMockPtrClusterConfigCreateSession.L.Unlock()
	condRecorderAuxMockPtrClusterConfigCreateSession.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrClusterConfigCreateSession() (ret int) {
	condRecorderAuxMockPtrClusterConfigCreateSession.L.Lock()
	ret = recorderAuxMockPtrClusterConfigCreateSession
	condRecorderAuxMockPtrClusterConfigCreateSession.L.Unlock()
	return
}

// (recvcfg *ClusterConfig)CreateSession - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvcfg *ClusterConfig) CreateSession() (reta *Session, retb error) {
	FuncAuxMockPtrClusterConfigCreateSession, ok := apomock.GetRegisteredFunc("gocql.ClusterConfig.CreateSession")
	if ok {
		reta, retb = FuncAuxMockPtrClusterConfigCreateSession.(func(recvcfg *ClusterConfig) (reta *Session, retb error))(recvcfg)
	} else {
		panic("FuncAuxMockPtrClusterConfigCreateSession ")
	}
	AuxMockIncrementRecorderAuxMockPtrClusterConfigCreateSession()
	return
}
