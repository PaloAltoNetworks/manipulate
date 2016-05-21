// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"
import "net"

import "golang.org/x/net/context"

import "bufio"
import "crypto/tls"
import "errors"

import "github.com/gocql/gocql/apointernal/streams"

import "time"

func init() {
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructCallReq, apomockNewStructCallReq)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructInflightPrepare, apomockNewStructInflightPrepare)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructPreparedStatment, apomockNewStructPreparedStatment)

	apomock.RegisterFunc("gocql", "gocql.PasswordAuthenticator.Success", (PasswordAuthenticator).AuxMockSuccess)
	apomock.RegisterFunc("gocql", "gocql.Connect", AuxMockConnect)
	apomock.RegisterFunc("gocql", "gocql.Conn.discardFrame", (*Conn).AuxMockdiscardFrame)
	apomock.RegisterFunc("gocql", "gocql.Conn.executeQuery", (*Conn).AuxMockexecuteQuery)
	apomock.RegisterFunc("gocql", "gocql.Conn.Address", (*Conn).AuxMockAddress)
	apomock.RegisterFunc("gocql", "gocql.Conn.authenticateHandshake", (*Conn).AuxMockauthenticateHandshake)
	apomock.RegisterFunc("gocql", "gocql.Conn.closeWithError", (*Conn).AuxMockcloseWithError)
	apomock.RegisterFunc("gocql", "gocql.Conn.recv", (*Conn).AuxMockrecv)
	apomock.RegisterFunc("gocql", "gocql.Conn.setKeepalive", (*Conn).AuxMocksetKeepalive)
	apomock.RegisterFunc("gocql", "gocql.PasswordAuthenticator.Challenge", (PasswordAuthenticator).AuxMockChallenge)
	apomock.RegisterFunc("gocql", "gocql.Conn.Read", (*Conn).AuxMockRead)
	apomock.RegisterFunc("gocql", "gocql.Conn.Close", (*Conn).AuxMockClose)
	apomock.RegisterFunc("gocql", "gocql.Conn.releaseStream", (*Conn).AuxMockreleaseStream)
	apomock.RegisterFunc("gocql", "gocql.Conn.Pick", (*Conn).AuxMockPick)
	apomock.RegisterFunc("gocql", "gocql.Conn.awaitSchemaAgreement", (*Conn).AuxMockawaitSchemaAgreement)
	apomock.RegisterFunc("gocql", "gocql.JoinHostPort", AuxMockJoinHostPort)
	apomock.RegisterFunc("gocql", "gocql.Conn.Closed", (*Conn).AuxMockClosed)
	apomock.RegisterFunc("gocql", "gocql.Conn.UseKeyspace", (*Conn).AuxMockUseKeyspace)
	apomock.RegisterFunc("gocql", "gocql.approve", AuxMockapprove)
	apomock.RegisterFunc("gocql", "gocql.Conn.serve", (*Conn).AuxMockserve)
	apomock.RegisterFunc("gocql", "gocql.Conn.AvailableStreams", (*Conn).AuxMockAvailableStreams)
	apomock.RegisterFunc("gocql", "gocql.connErrorHandlerFn.HandleError", (connErrorHandlerFn).AuxMockHandleError)
	apomock.RegisterFunc("gocql", "gocql.Conn.handleTimeout", (*Conn).AuxMockhandleTimeout)
	apomock.RegisterFunc("gocql", "gocql.Conn.prepareStatement", (*Conn).AuxMockprepareStatement)
	apomock.RegisterFunc("gocql", "gocql.Conn.executeBatch", (*Conn).AuxMockexecuteBatch)
	apomock.RegisterFunc("gocql", "gocql.Conn.Write", (*Conn).AuxMockWrite)
	apomock.RegisterFunc("gocql", "gocql.Conn.startup", (*Conn).AuxMockstartup)
	apomock.RegisterFunc("gocql", "gocql.Conn.exec", (*Conn).AuxMockexec)
	apomock.RegisterFunc("gocql", "gocql.Conn.query", (*Conn).AuxMockquery)
}

const (
	ApomockStructCallReq          = 7
	ApomockStructInflightPrepare  = 8
	ApomockStructPreparedStatment = 9
)

var (
	approvedAuthenticators = [...]string{"org.apache.cassandra.auth.PasswordAuthenticator", "com.instaclustr.cassandra.auth.SharedSecretAuthenticator"}
)

var TimeoutLimit int64 = 10

var (
	streamPool = sync.Pool{New: func() interface{} {
		return &callReq{resp: make(chan error)}
	}}
)

var (
	ErrQueryArgLength    = errors.New("gocql: query argument length mismatch")
	ErrTimeoutNoResponse = errors.New("gocql: no response received from cassandra within timeout period")
	ErrTooManyTimeouts   = errors.New("gocql: too many query timeouts on the connection")
	ErrConnectionClosed  = errors.New("gocql: connection closed waiting for response")
	ErrNoStreams         = errors.New("gocql: no streams available on connection")
)

//
// Internal Types: in this package and their exportable versions
//
type connErrorHandlerFn func(conn *Conn, err error, closed bool)
type callReq struct {
	resp     chan error
	framer   *framer
	timeout  chan struct{}
	streamID int
	timer    *time.Timer
}
type inflightPrepare struct {
	wg               sync.WaitGroup
	err              error
	preparedStatment *preparedStatment
}
type preparedStatment struct {
	id       []byte
	request  preparedMetadata
	response resultMetadata
}

//
// External Types: in this package
//
type ConnErrorHandler interface {
	HandleError(conn *Conn, err error, closed bool)
}

type Conn struct {
	conn            net.Conn
	r               *bufio.Reader
	timeout         time.Duration
	cfg             *ConnConfig
	headerBuf       [maxFrameHeaderSize]byte
	streams         *streams.IDGenerator
	mu              sync.RWMutex
	calls           map[int]*callReq
	errorHandler    ConnErrorHandler
	compressor      Compressor
	auth            Authenticator
	addr            string
	version         uint8
	currentKeyspace string
	host            *HostInfo
	session         *Session
	closed          int32
	quit            chan struct{}
	timeouts        int64
}

type Authenticator interface {
	Challenge(req []byte) (resp []byte, auth Authenticator, err error)
	Success(data []byte) error
}

type SslOptions struct {
	tls.Config
	CertPath               string
	KeyPath                string
	CaPath                 string
	EnableHostVerification bool
}

type ConnConfig struct {
	ProtoVersion  int
	CQLVersion    string
	Timeout       time.Duration
	Compressor    Compressor
	Authenticator Authenticator
	Keepalive     time.Duration
	tlsConfig     *tls.Config
}

type PasswordAuthenticator struct {
	Username string
	Password string
}

func apomockNewStructCallReq() interface{}          { return &callReq{} }
func apomockNewStructInflightPrepare() interface{}  { return &inflightPrepare{} }
func apomockNewStructPreparedStatment() interface{} { return &preparedStatment{} }

//
// Mock: (recvp PasswordAuthenticator)Success(argdata []byte)(reta error)
//

type MockArgsTypePasswordAuthenticatorSuccess struct {
	ApomockCallNumber int
	Argdata           []byte
}

var LastMockArgsPasswordAuthenticatorSuccess MockArgsTypePasswordAuthenticatorSuccess

// (recvp PasswordAuthenticator)AuxMockSuccess(argdata []byte)(reta error) - Generated mock function
func (recvp PasswordAuthenticator) AuxMockSuccess(argdata []byte) (reta error) {
	LastMockArgsPasswordAuthenticatorSuccess = MockArgsTypePasswordAuthenticatorSuccess{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPasswordAuthenticatorSuccess(),
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.PasswordAuthenticator.Success")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.PasswordAuthenticator.Success")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.PasswordAuthenticator.Success")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPasswordAuthenticatorSuccess  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPasswordAuthenticatorSuccess int = 0

var condRecorderAuxMockPasswordAuthenticatorSuccess *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPasswordAuthenticatorSuccess(i int) {
	condRecorderAuxMockPasswordAuthenticatorSuccess.L.Lock()
	for recorderAuxMockPasswordAuthenticatorSuccess < i {
		condRecorderAuxMockPasswordAuthenticatorSuccess.Wait()
	}
	condRecorderAuxMockPasswordAuthenticatorSuccess.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPasswordAuthenticatorSuccess() {
	condRecorderAuxMockPasswordAuthenticatorSuccess.L.Lock()
	recorderAuxMockPasswordAuthenticatorSuccess++
	condRecorderAuxMockPasswordAuthenticatorSuccess.L.Unlock()
	condRecorderAuxMockPasswordAuthenticatorSuccess.Broadcast()
}
func AuxMockGetRecorderAuxMockPasswordAuthenticatorSuccess() (ret int) {
	condRecorderAuxMockPasswordAuthenticatorSuccess.L.Lock()
	ret = recorderAuxMockPasswordAuthenticatorSuccess
	condRecorderAuxMockPasswordAuthenticatorSuccess.L.Unlock()
	return
}

// (recvp PasswordAuthenticator)Success - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp PasswordAuthenticator) Success(argdata []byte) (reta error) {
	FuncAuxMockPasswordAuthenticatorSuccess, ok := apomock.GetRegisteredFunc("gocql.PasswordAuthenticator.Success")
	if ok {
		reta = FuncAuxMockPasswordAuthenticatorSuccess.(func(recvp PasswordAuthenticator, argdata []byte) (reta error))(recvp, argdata)
	} else {
		panic("FuncAuxMockPasswordAuthenticatorSuccess ")
	}
	AuxMockIncrementRecorderAuxMockPasswordAuthenticatorSuccess()
	return
}

//
// Mock: Connect(arghost *HostInfo, argaddr string, argcfg *ConnConfig, argerrorHandler ConnErrorHandler, argsession *Session)(reta *Conn, retb error)
//

type MockArgsTypeConnect struct {
	ApomockCallNumber int
	Arghost           *HostInfo
	Argaddr           string
	Argcfg            *ConnConfig
	ArgerrorHandler   ConnErrorHandler
	Argsession        *Session
}

var LastMockArgsConnect MockArgsTypeConnect

// AuxMockConnect(arghost *HostInfo, argaddr string, argcfg *ConnConfig, argerrorHandler ConnErrorHandler, argsession *Session)(reta *Conn, retb error) - Generated mock function
func AuxMockConnect(arghost *HostInfo, argaddr string, argcfg *ConnConfig, argerrorHandler ConnErrorHandler, argsession *Session) (reta *Conn, retb error) {
	LastMockArgsConnect = MockArgsTypeConnect{
		ApomockCallNumber: AuxMockGetRecorderAuxMockConnect(),
		Arghost:           arghost,
		Argaddr:           argaddr,
		Argcfg:            argcfg,
		ArgerrorHandler:   argerrorHandler,
		Argsession:        argsession,
	}
	rargs, rerr := apomock.GetNext("gocql.Connect")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Connect")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Connect")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Conn)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockConnect  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockConnect int = 0

var condRecorderAuxMockConnect *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockConnect(i int) {
	condRecorderAuxMockConnect.L.Lock()
	for recorderAuxMockConnect < i {
		condRecorderAuxMockConnect.Wait()
	}
	condRecorderAuxMockConnect.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockConnect() {
	condRecorderAuxMockConnect.L.Lock()
	recorderAuxMockConnect++
	condRecorderAuxMockConnect.L.Unlock()
	condRecorderAuxMockConnect.Broadcast()
}
func AuxMockGetRecorderAuxMockConnect() (ret int) {
	condRecorderAuxMockConnect.L.Lock()
	ret = recorderAuxMockConnect
	condRecorderAuxMockConnect.L.Unlock()
	return
}

// Connect - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func Connect(arghost *HostInfo, argaddr string, argcfg *ConnConfig, argerrorHandler ConnErrorHandler, argsession *Session) (reta *Conn, retb error) {
	FuncAuxMockConnect, ok := apomock.GetRegisteredFunc("gocql.Connect")
	if ok {
		reta, retb = FuncAuxMockConnect.(func(arghost *HostInfo, argaddr string, argcfg *ConnConfig, argerrorHandler ConnErrorHandler, argsession *Session) (reta *Conn, retb error))(arghost, argaddr, argcfg, argerrorHandler, argsession)
	} else {
		panic("FuncAuxMockConnect ")
	}
	AuxMockIncrementRecorderAuxMockConnect()
	return
}

//
// Mock: (recvc *Conn)discardFrame(arghead frameHeader)(reta error)
//

type MockArgsTypeConndiscardFrame struct {
	ApomockCallNumber int
	Arghead           frameHeader
}

var LastMockArgsConndiscardFrame MockArgsTypeConndiscardFrame

// (recvc *Conn)AuxMockdiscardFrame(arghead frameHeader)(reta error) - Generated mock function
func (recvc *Conn) AuxMockdiscardFrame(arghead frameHeader) (reta error) {
	LastMockArgsConndiscardFrame = MockArgsTypeConndiscardFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConndiscardFrame(),
		Arghead:           arghead,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.discardFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.discardFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.discardFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrConndiscardFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConndiscardFrame int = 0

var condRecorderAuxMockPtrConndiscardFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConndiscardFrame(i int) {
	condRecorderAuxMockPtrConndiscardFrame.L.Lock()
	for recorderAuxMockPtrConndiscardFrame < i {
		condRecorderAuxMockPtrConndiscardFrame.Wait()
	}
	condRecorderAuxMockPtrConndiscardFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConndiscardFrame() {
	condRecorderAuxMockPtrConndiscardFrame.L.Lock()
	recorderAuxMockPtrConndiscardFrame++
	condRecorderAuxMockPtrConndiscardFrame.L.Unlock()
	condRecorderAuxMockPtrConndiscardFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConndiscardFrame() (ret int) {
	condRecorderAuxMockPtrConndiscardFrame.L.Lock()
	ret = recorderAuxMockPtrConndiscardFrame
	condRecorderAuxMockPtrConndiscardFrame.L.Unlock()
	return
}

// (recvc *Conn)discardFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) discardFrame(arghead frameHeader) (reta error) {
	FuncAuxMockPtrConndiscardFrame, ok := apomock.GetRegisteredFunc("gocql.Conn.discardFrame")
	if ok {
		reta = FuncAuxMockPtrConndiscardFrame.(func(recvc *Conn, arghead frameHeader) (reta error))(recvc, arghead)
	} else {
		panic("FuncAuxMockPtrConndiscardFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrConndiscardFrame()
	return
}

//
// Mock: (recvc *Conn)executeQuery(argqry *Query)(reta *Iter)
//

type MockArgsTypeConnexecuteQuery struct {
	ApomockCallNumber int
	Argqry            *Query
}

var LastMockArgsConnexecuteQuery MockArgsTypeConnexecuteQuery

// (recvc *Conn)AuxMockexecuteQuery(argqry *Query)(reta *Iter) - Generated mock function
func (recvc *Conn) AuxMockexecuteQuery(argqry *Query) (reta *Iter) {
	LastMockArgsConnexecuteQuery = MockArgsTypeConnexecuteQuery{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnexecuteQuery(),
		Argqry:            argqry,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.executeQuery")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.executeQuery")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.executeQuery")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Iter)
	}
	return
}

// RecorderAuxMockPtrConnexecuteQuery  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnexecuteQuery int = 0

var condRecorderAuxMockPtrConnexecuteQuery *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnexecuteQuery(i int) {
	condRecorderAuxMockPtrConnexecuteQuery.L.Lock()
	for recorderAuxMockPtrConnexecuteQuery < i {
		condRecorderAuxMockPtrConnexecuteQuery.Wait()
	}
	condRecorderAuxMockPtrConnexecuteQuery.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnexecuteQuery() {
	condRecorderAuxMockPtrConnexecuteQuery.L.Lock()
	recorderAuxMockPtrConnexecuteQuery++
	condRecorderAuxMockPtrConnexecuteQuery.L.Unlock()
	condRecorderAuxMockPtrConnexecuteQuery.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnexecuteQuery() (ret int) {
	condRecorderAuxMockPtrConnexecuteQuery.L.Lock()
	ret = recorderAuxMockPtrConnexecuteQuery
	condRecorderAuxMockPtrConnexecuteQuery.L.Unlock()
	return
}

// (recvc *Conn)executeQuery - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) executeQuery(argqry *Query) (reta *Iter) {
	FuncAuxMockPtrConnexecuteQuery, ok := apomock.GetRegisteredFunc("gocql.Conn.executeQuery")
	if ok {
		reta = FuncAuxMockPtrConnexecuteQuery.(func(recvc *Conn, argqry *Query) (reta *Iter))(recvc, argqry)
	} else {
		panic("FuncAuxMockPtrConnexecuteQuery ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnexecuteQuery()
	return
}

//
// Mock: (recvc *Conn)Address()(reta string)
//

type MockArgsTypeConnAddress struct {
	ApomockCallNumber int
}

var LastMockArgsConnAddress MockArgsTypeConnAddress

// (recvc *Conn)AuxMockAddress()(reta string) - Generated mock function
func (recvc *Conn) AuxMockAddress() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.Conn.Address")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.Address")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.Address")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrConnAddress  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnAddress int = 0

var condRecorderAuxMockPtrConnAddress *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnAddress(i int) {
	condRecorderAuxMockPtrConnAddress.L.Lock()
	for recorderAuxMockPtrConnAddress < i {
		condRecorderAuxMockPtrConnAddress.Wait()
	}
	condRecorderAuxMockPtrConnAddress.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnAddress() {
	condRecorderAuxMockPtrConnAddress.L.Lock()
	recorderAuxMockPtrConnAddress++
	condRecorderAuxMockPtrConnAddress.L.Unlock()
	condRecorderAuxMockPtrConnAddress.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnAddress() (ret int) {
	condRecorderAuxMockPtrConnAddress.L.Lock()
	ret = recorderAuxMockPtrConnAddress
	condRecorderAuxMockPtrConnAddress.L.Unlock()
	return
}

// (recvc *Conn)Address - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) Address() (reta string) {
	FuncAuxMockPtrConnAddress, ok := apomock.GetRegisteredFunc("gocql.Conn.Address")
	if ok {
		reta = FuncAuxMockPtrConnAddress.(func(recvc *Conn) (reta string))(recvc)
	} else {
		panic("FuncAuxMockPtrConnAddress ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnAddress()
	return
}

//
// Mock: (recvc *Conn)authenticateHandshake(argctx context.Context, argauthFrame *authenticateFrame, argframeTicker chan struct{})(reta error)
//

type MockArgsTypeConnauthenticateHandshake struct {
	ApomockCallNumber int
	Argctx            context.Context
	ArgauthFrame      *authenticateFrame
	ArgframeTicker    chan struct{}
}

var LastMockArgsConnauthenticateHandshake MockArgsTypeConnauthenticateHandshake

// (recvc *Conn)AuxMockauthenticateHandshake(argctx context.Context, argauthFrame *authenticateFrame, argframeTicker chan struct{})(reta error) - Generated mock function
func (recvc *Conn) AuxMockauthenticateHandshake(argctx context.Context, argauthFrame *authenticateFrame, argframeTicker chan struct{}) (reta error) {
	LastMockArgsConnauthenticateHandshake = MockArgsTypeConnauthenticateHandshake{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnauthenticateHandshake(),
		Argctx:            argctx,
		ArgauthFrame:      argauthFrame,
		ArgframeTicker:    argframeTicker,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.authenticateHandshake")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.authenticateHandshake")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.authenticateHandshake")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrConnauthenticateHandshake  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnauthenticateHandshake int = 0

var condRecorderAuxMockPtrConnauthenticateHandshake *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnauthenticateHandshake(i int) {
	condRecorderAuxMockPtrConnauthenticateHandshake.L.Lock()
	for recorderAuxMockPtrConnauthenticateHandshake < i {
		condRecorderAuxMockPtrConnauthenticateHandshake.Wait()
	}
	condRecorderAuxMockPtrConnauthenticateHandshake.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnauthenticateHandshake() {
	condRecorderAuxMockPtrConnauthenticateHandshake.L.Lock()
	recorderAuxMockPtrConnauthenticateHandshake++
	condRecorderAuxMockPtrConnauthenticateHandshake.L.Unlock()
	condRecorderAuxMockPtrConnauthenticateHandshake.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnauthenticateHandshake() (ret int) {
	condRecorderAuxMockPtrConnauthenticateHandshake.L.Lock()
	ret = recorderAuxMockPtrConnauthenticateHandshake
	condRecorderAuxMockPtrConnauthenticateHandshake.L.Unlock()
	return
}

// (recvc *Conn)authenticateHandshake - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) authenticateHandshake(argctx context.Context, argauthFrame *authenticateFrame, argframeTicker chan struct{}) (reta error) {
	FuncAuxMockPtrConnauthenticateHandshake, ok := apomock.GetRegisteredFunc("gocql.Conn.authenticateHandshake")
	if ok {
		reta = FuncAuxMockPtrConnauthenticateHandshake.(func(recvc *Conn, argctx context.Context, argauthFrame *authenticateFrame, argframeTicker chan struct{}) (reta error))(recvc, argctx, argauthFrame, argframeTicker)
	} else {
		panic("FuncAuxMockPtrConnauthenticateHandshake ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnauthenticateHandshake()
	return
}

//
// Mock: (recvc *Conn)closeWithError(argerr error)()
//

type MockArgsTypeConncloseWithError struct {
	ApomockCallNumber int
	Argerr            error
}

var LastMockArgsConncloseWithError MockArgsTypeConncloseWithError

// (recvc *Conn)AuxMockcloseWithError(argerr error)() - Generated mock function
func (recvc *Conn) AuxMockcloseWithError(argerr error) {
	LastMockArgsConncloseWithError = MockArgsTypeConncloseWithError{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConncloseWithError(),
		Argerr:            argerr,
	}
	return
}

// RecorderAuxMockPtrConncloseWithError  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConncloseWithError int = 0

var condRecorderAuxMockPtrConncloseWithError *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConncloseWithError(i int) {
	condRecorderAuxMockPtrConncloseWithError.L.Lock()
	for recorderAuxMockPtrConncloseWithError < i {
		condRecorderAuxMockPtrConncloseWithError.Wait()
	}
	condRecorderAuxMockPtrConncloseWithError.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConncloseWithError() {
	condRecorderAuxMockPtrConncloseWithError.L.Lock()
	recorderAuxMockPtrConncloseWithError++
	condRecorderAuxMockPtrConncloseWithError.L.Unlock()
	condRecorderAuxMockPtrConncloseWithError.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConncloseWithError() (ret int) {
	condRecorderAuxMockPtrConncloseWithError.L.Lock()
	ret = recorderAuxMockPtrConncloseWithError
	condRecorderAuxMockPtrConncloseWithError.L.Unlock()
	return
}

// (recvc *Conn)closeWithError - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) closeWithError(argerr error) {
	FuncAuxMockPtrConncloseWithError, ok := apomock.GetRegisteredFunc("gocql.Conn.closeWithError")
	if ok {
		FuncAuxMockPtrConncloseWithError.(func(recvc *Conn, argerr error))(recvc, argerr)
	} else {
		panic("FuncAuxMockPtrConncloseWithError ")
	}
	AuxMockIncrementRecorderAuxMockPtrConncloseWithError()
	return
}

//
// Mock: (recvc *Conn)recv()(reta error)
//

type MockArgsTypeConnrecv struct {
	ApomockCallNumber int
}

var LastMockArgsConnrecv MockArgsTypeConnrecv

// (recvc *Conn)AuxMockrecv()(reta error) - Generated mock function
func (recvc *Conn) AuxMockrecv() (reta error) {
	rargs, rerr := apomock.GetNext("gocql.Conn.recv")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.recv")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.recv")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrConnrecv  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnrecv int = 0

var condRecorderAuxMockPtrConnrecv *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnrecv(i int) {
	condRecorderAuxMockPtrConnrecv.L.Lock()
	for recorderAuxMockPtrConnrecv < i {
		condRecorderAuxMockPtrConnrecv.Wait()
	}
	condRecorderAuxMockPtrConnrecv.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnrecv() {
	condRecorderAuxMockPtrConnrecv.L.Lock()
	recorderAuxMockPtrConnrecv++
	condRecorderAuxMockPtrConnrecv.L.Unlock()
	condRecorderAuxMockPtrConnrecv.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnrecv() (ret int) {
	condRecorderAuxMockPtrConnrecv.L.Lock()
	ret = recorderAuxMockPtrConnrecv
	condRecorderAuxMockPtrConnrecv.L.Unlock()
	return
}

// (recvc *Conn)recv - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) recv() (reta error) {
	FuncAuxMockPtrConnrecv, ok := apomock.GetRegisteredFunc("gocql.Conn.recv")
	if ok {
		reta = FuncAuxMockPtrConnrecv.(func(recvc *Conn) (reta error))(recvc)
	} else {
		panic("FuncAuxMockPtrConnrecv ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnrecv()
	return
}

//
// Mock: (recvc *Conn)setKeepalive(argd time.Duration)(reta error)
//

type MockArgsTypeConnsetKeepalive struct {
	ApomockCallNumber int
	Argd              time.Duration
}

var LastMockArgsConnsetKeepalive MockArgsTypeConnsetKeepalive

// (recvc *Conn)AuxMocksetKeepalive(argd time.Duration)(reta error) - Generated mock function
func (recvc *Conn) AuxMocksetKeepalive(argd time.Duration) (reta error) {
	LastMockArgsConnsetKeepalive = MockArgsTypeConnsetKeepalive{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnsetKeepalive(),
		Argd:              argd,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.setKeepalive")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.setKeepalive")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.setKeepalive")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrConnsetKeepalive  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnsetKeepalive int = 0

var condRecorderAuxMockPtrConnsetKeepalive *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnsetKeepalive(i int) {
	condRecorderAuxMockPtrConnsetKeepalive.L.Lock()
	for recorderAuxMockPtrConnsetKeepalive < i {
		condRecorderAuxMockPtrConnsetKeepalive.Wait()
	}
	condRecorderAuxMockPtrConnsetKeepalive.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnsetKeepalive() {
	condRecorderAuxMockPtrConnsetKeepalive.L.Lock()
	recorderAuxMockPtrConnsetKeepalive++
	condRecorderAuxMockPtrConnsetKeepalive.L.Unlock()
	condRecorderAuxMockPtrConnsetKeepalive.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnsetKeepalive() (ret int) {
	condRecorderAuxMockPtrConnsetKeepalive.L.Lock()
	ret = recorderAuxMockPtrConnsetKeepalive
	condRecorderAuxMockPtrConnsetKeepalive.L.Unlock()
	return
}

// (recvc *Conn)setKeepalive - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) setKeepalive(argd time.Duration) (reta error) {
	FuncAuxMockPtrConnsetKeepalive, ok := apomock.GetRegisteredFunc("gocql.Conn.setKeepalive")
	if ok {
		reta = FuncAuxMockPtrConnsetKeepalive.(func(recvc *Conn, argd time.Duration) (reta error))(recvc, argd)
	} else {
		panic("FuncAuxMockPtrConnsetKeepalive ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnsetKeepalive()
	return
}

//
// Mock: (recvp PasswordAuthenticator)Challenge(argreq []byte)(reta []byte, retb Authenticator, retc error)
//

type MockArgsTypePasswordAuthenticatorChallenge struct {
	ApomockCallNumber int
	Argreq            []byte
}

var LastMockArgsPasswordAuthenticatorChallenge MockArgsTypePasswordAuthenticatorChallenge

// (recvp PasswordAuthenticator)AuxMockChallenge(argreq []byte)(reta []byte, retb Authenticator, retc error) - Generated mock function
func (recvp PasswordAuthenticator) AuxMockChallenge(argreq []byte) (reta []byte, retb Authenticator, retc error) {
	LastMockArgsPasswordAuthenticatorChallenge = MockArgsTypePasswordAuthenticatorChallenge{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPasswordAuthenticatorChallenge(),
		Argreq:            argreq,
	}
	rargs, rerr := apomock.GetNext("gocql.PasswordAuthenticator.Challenge")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.PasswordAuthenticator.Challenge")
	} else if rargs.NumArgs() != 3 {
		panic("All return parameters not provided for method:gocql.PasswordAuthenticator.Challenge")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(Authenticator)
	}
	if rargs.GetArg(2) != nil {
		retc = rargs.GetArg(2).(error)
	}
	return
}

// RecorderAuxMockPasswordAuthenticatorChallenge  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPasswordAuthenticatorChallenge int = 0

var condRecorderAuxMockPasswordAuthenticatorChallenge *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPasswordAuthenticatorChallenge(i int) {
	condRecorderAuxMockPasswordAuthenticatorChallenge.L.Lock()
	for recorderAuxMockPasswordAuthenticatorChallenge < i {
		condRecorderAuxMockPasswordAuthenticatorChallenge.Wait()
	}
	condRecorderAuxMockPasswordAuthenticatorChallenge.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPasswordAuthenticatorChallenge() {
	condRecorderAuxMockPasswordAuthenticatorChallenge.L.Lock()
	recorderAuxMockPasswordAuthenticatorChallenge++
	condRecorderAuxMockPasswordAuthenticatorChallenge.L.Unlock()
	condRecorderAuxMockPasswordAuthenticatorChallenge.Broadcast()
}
func AuxMockGetRecorderAuxMockPasswordAuthenticatorChallenge() (ret int) {
	condRecorderAuxMockPasswordAuthenticatorChallenge.L.Lock()
	ret = recorderAuxMockPasswordAuthenticatorChallenge
	condRecorderAuxMockPasswordAuthenticatorChallenge.L.Unlock()
	return
}

// (recvp PasswordAuthenticator)Challenge - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp PasswordAuthenticator) Challenge(argreq []byte) (reta []byte, retb Authenticator, retc error) {
	FuncAuxMockPasswordAuthenticatorChallenge, ok := apomock.GetRegisteredFunc("gocql.PasswordAuthenticator.Challenge")
	if ok {
		reta, retb, retc = FuncAuxMockPasswordAuthenticatorChallenge.(func(recvp PasswordAuthenticator, argreq []byte) (reta []byte, retb Authenticator, retc error))(recvp, argreq)
	} else {
		panic("FuncAuxMockPasswordAuthenticatorChallenge ")
	}
	AuxMockIncrementRecorderAuxMockPasswordAuthenticatorChallenge()
	return
}

//
// Mock: (recvc *Conn)Read(argp []byte)(retn int, reterr error)
//

type MockArgsTypeConnRead struct {
	ApomockCallNumber int
	Argp              []byte
}

var LastMockArgsConnRead MockArgsTypeConnRead

// (recvc *Conn)AuxMockRead(argp []byte)(retn int, reterr error) - Generated mock function
func (recvc *Conn) AuxMockRead(argp []byte) (retn int, reterr error) {
	LastMockArgsConnRead = MockArgsTypeConnRead{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnRead(),
		Argp:              argp,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.Read")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.Read")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Conn.Read")
	}
	if rargs.GetArg(0) != nil {
		retn = rargs.GetArg(0).(int)
	}
	if rargs.GetArg(1) != nil {
		reterr = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrConnRead  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnRead int = 0

var condRecorderAuxMockPtrConnRead *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnRead(i int) {
	condRecorderAuxMockPtrConnRead.L.Lock()
	for recorderAuxMockPtrConnRead < i {
		condRecorderAuxMockPtrConnRead.Wait()
	}
	condRecorderAuxMockPtrConnRead.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnRead() {
	condRecorderAuxMockPtrConnRead.L.Lock()
	recorderAuxMockPtrConnRead++
	condRecorderAuxMockPtrConnRead.L.Unlock()
	condRecorderAuxMockPtrConnRead.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnRead() (ret int) {
	condRecorderAuxMockPtrConnRead.L.Lock()
	ret = recorderAuxMockPtrConnRead
	condRecorderAuxMockPtrConnRead.L.Unlock()
	return
}

// (recvc *Conn)Read - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) Read(argp []byte) (retn int, reterr error) {
	FuncAuxMockPtrConnRead, ok := apomock.GetRegisteredFunc("gocql.Conn.Read")
	if ok {
		retn, reterr = FuncAuxMockPtrConnRead.(func(recvc *Conn, argp []byte) (retn int, reterr error))(recvc, argp)
	} else {
		panic("FuncAuxMockPtrConnRead ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnRead()
	return
}

//
// Mock: (recvc *Conn)Close()()
//

type MockArgsTypeConnClose struct {
	ApomockCallNumber int
}

var LastMockArgsConnClose MockArgsTypeConnClose

// (recvc *Conn)AuxMockClose()() - Generated mock function
func (recvc *Conn) AuxMockClose() {
	return
}

// RecorderAuxMockPtrConnClose  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnClose int = 0

var condRecorderAuxMockPtrConnClose *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnClose(i int) {
	condRecorderAuxMockPtrConnClose.L.Lock()
	for recorderAuxMockPtrConnClose < i {
		condRecorderAuxMockPtrConnClose.Wait()
	}
	condRecorderAuxMockPtrConnClose.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnClose() {
	condRecorderAuxMockPtrConnClose.L.Lock()
	recorderAuxMockPtrConnClose++
	condRecorderAuxMockPtrConnClose.L.Unlock()
	condRecorderAuxMockPtrConnClose.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnClose() (ret int) {
	condRecorderAuxMockPtrConnClose.L.Lock()
	ret = recorderAuxMockPtrConnClose
	condRecorderAuxMockPtrConnClose.L.Unlock()
	return
}

// (recvc *Conn)Close - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) Close() {
	FuncAuxMockPtrConnClose, ok := apomock.GetRegisteredFunc("gocql.Conn.Close")
	if ok {
		FuncAuxMockPtrConnClose.(func(recvc *Conn))(recvc)
	} else {
		panic("FuncAuxMockPtrConnClose ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnClose()
	return
}

//
// Mock: (recvc *Conn)releaseStream(argstream int)()
//

type MockArgsTypeConnreleaseStream struct {
	ApomockCallNumber int
	Argstream         int
}

var LastMockArgsConnreleaseStream MockArgsTypeConnreleaseStream

// (recvc *Conn)AuxMockreleaseStream(argstream int)() - Generated mock function
func (recvc *Conn) AuxMockreleaseStream(argstream int) {
	LastMockArgsConnreleaseStream = MockArgsTypeConnreleaseStream{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnreleaseStream(),
		Argstream:         argstream,
	}
	return
}

// RecorderAuxMockPtrConnreleaseStream  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnreleaseStream int = 0

var condRecorderAuxMockPtrConnreleaseStream *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnreleaseStream(i int) {
	condRecorderAuxMockPtrConnreleaseStream.L.Lock()
	for recorderAuxMockPtrConnreleaseStream < i {
		condRecorderAuxMockPtrConnreleaseStream.Wait()
	}
	condRecorderAuxMockPtrConnreleaseStream.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnreleaseStream() {
	condRecorderAuxMockPtrConnreleaseStream.L.Lock()
	recorderAuxMockPtrConnreleaseStream++
	condRecorderAuxMockPtrConnreleaseStream.L.Unlock()
	condRecorderAuxMockPtrConnreleaseStream.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnreleaseStream() (ret int) {
	condRecorderAuxMockPtrConnreleaseStream.L.Lock()
	ret = recorderAuxMockPtrConnreleaseStream
	condRecorderAuxMockPtrConnreleaseStream.L.Unlock()
	return
}

// (recvc *Conn)releaseStream - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) releaseStream(argstream int) {
	FuncAuxMockPtrConnreleaseStream, ok := apomock.GetRegisteredFunc("gocql.Conn.releaseStream")
	if ok {
		FuncAuxMockPtrConnreleaseStream.(func(recvc *Conn, argstream int))(recvc, argstream)
	} else {
		panic("FuncAuxMockPtrConnreleaseStream ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnreleaseStream()
	return
}

//
// Mock: (recvc *Conn)Pick(argqry *Query)(reta *Conn)
//

type MockArgsTypeConnPick struct {
	ApomockCallNumber int
	Argqry            *Query
}

var LastMockArgsConnPick MockArgsTypeConnPick

// (recvc *Conn)AuxMockPick(argqry *Query)(reta *Conn) - Generated mock function
func (recvc *Conn) AuxMockPick(argqry *Query) (reta *Conn) {
	LastMockArgsConnPick = MockArgsTypeConnPick{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnPick(),
		Argqry:            argqry,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.Pick")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.Pick")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.Pick")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Conn)
	}
	return
}

// RecorderAuxMockPtrConnPick  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnPick int = 0

var condRecorderAuxMockPtrConnPick *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnPick(i int) {
	condRecorderAuxMockPtrConnPick.L.Lock()
	for recorderAuxMockPtrConnPick < i {
		condRecorderAuxMockPtrConnPick.Wait()
	}
	condRecorderAuxMockPtrConnPick.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnPick() {
	condRecorderAuxMockPtrConnPick.L.Lock()
	recorderAuxMockPtrConnPick++
	condRecorderAuxMockPtrConnPick.L.Unlock()
	condRecorderAuxMockPtrConnPick.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnPick() (ret int) {
	condRecorderAuxMockPtrConnPick.L.Lock()
	ret = recorderAuxMockPtrConnPick
	condRecorderAuxMockPtrConnPick.L.Unlock()
	return
}

// (recvc *Conn)Pick - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) Pick(argqry *Query) (reta *Conn) {
	FuncAuxMockPtrConnPick, ok := apomock.GetRegisteredFunc("gocql.Conn.Pick")
	if ok {
		reta = FuncAuxMockPtrConnPick.(func(recvc *Conn, argqry *Query) (reta *Conn))(recvc, argqry)
	} else {
		panic("FuncAuxMockPtrConnPick ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnPick()
	return
}

//
// Mock: (recvc *Conn)awaitSchemaAgreement()(reterr error)
//

type MockArgsTypeConnawaitSchemaAgreement struct {
	ApomockCallNumber int
}

var LastMockArgsConnawaitSchemaAgreement MockArgsTypeConnawaitSchemaAgreement

// (recvc *Conn)AuxMockawaitSchemaAgreement()(reterr error) - Generated mock function
func (recvc *Conn) AuxMockawaitSchemaAgreement() (reterr error) {
	rargs, rerr := apomock.GetNext("gocql.Conn.awaitSchemaAgreement")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.awaitSchemaAgreement")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.awaitSchemaAgreement")
	}
	if rargs.GetArg(0) != nil {
		reterr = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrConnawaitSchemaAgreement  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnawaitSchemaAgreement int = 0

var condRecorderAuxMockPtrConnawaitSchemaAgreement *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnawaitSchemaAgreement(i int) {
	condRecorderAuxMockPtrConnawaitSchemaAgreement.L.Lock()
	for recorderAuxMockPtrConnawaitSchemaAgreement < i {
		condRecorderAuxMockPtrConnawaitSchemaAgreement.Wait()
	}
	condRecorderAuxMockPtrConnawaitSchemaAgreement.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnawaitSchemaAgreement() {
	condRecorderAuxMockPtrConnawaitSchemaAgreement.L.Lock()
	recorderAuxMockPtrConnawaitSchemaAgreement++
	condRecorderAuxMockPtrConnawaitSchemaAgreement.L.Unlock()
	condRecorderAuxMockPtrConnawaitSchemaAgreement.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnawaitSchemaAgreement() (ret int) {
	condRecorderAuxMockPtrConnawaitSchemaAgreement.L.Lock()
	ret = recorderAuxMockPtrConnawaitSchemaAgreement
	condRecorderAuxMockPtrConnawaitSchemaAgreement.L.Unlock()
	return
}

// (recvc *Conn)awaitSchemaAgreement - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) awaitSchemaAgreement() (reterr error) {
	FuncAuxMockPtrConnawaitSchemaAgreement, ok := apomock.GetRegisteredFunc("gocql.Conn.awaitSchemaAgreement")
	if ok {
		reterr = FuncAuxMockPtrConnawaitSchemaAgreement.(func(recvc *Conn) (reterr error))(recvc)
	} else {
		panic("FuncAuxMockPtrConnawaitSchemaAgreement ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnawaitSchemaAgreement()
	return
}

//
// Mock: JoinHostPort(argaddr string, argport int)(reta string)
//

type MockArgsTypeJoinHostPort struct {
	ApomockCallNumber int
	Argaddr           string
	Argport           int
}

var LastMockArgsJoinHostPort MockArgsTypeJoinHostPort

// AuxMockJoinHostPort(argaddr string, argport int)(reta string) - Generated mock function
func AuxMockJoinHostPort(argaddr string, argport int) (reta string) {
	LastMockArgsJoinHostPort = MockArgsTypeJoinHostPort{
		ApomockCallNumber: AuxMockGetRecorderAuxMockJoinHostPort(),
		Argaddr:           argaddr,
		Argport:           argport,
	}
	rargs, rerr := apomock.GetNext("gocql.JoinHostPort")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.JoinHostPort")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.JoinHostPort")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockJoinHostPort  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockJoinHostPort int = 0

var condRecorderAuxMockJoinHostPort *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockJoinHostPort(i int) {
	condRecorderAuxMockJoinHostPort.L.Lock()
	for recorderAuxMockJoinHostPort < i {
		condRecorderAuxMockJoinHostPort.Wait()
	}
	condRecorderAuxMockJoinHostPort.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockJoinHostPort() {
	condRecorderAuxMockJoinHostPort.L.Lock()
	recorderAuxMockJoinHostPort++
	condRecorderAuxMockJoinHostPort.L.Unlock()
	condRecorderAuxMockJoinHostPort.Broadcast()
}
func AuxMockGetRecorderAuxMockJoinHostPort() (ret int) {
	condRecorderAuxMockJoinHostPort.L.Lock()
	ret = recorderAuxMockJoinHostPort
	condRecorderAuxMockJoinHostPort.L.Unlock()
	return
}

// JoinHostPort - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func JoinHostPort(argaddr string, argport int) (reta string) {
	FuncAuxMockJoinHostPort, ok := apomock.GetRegisteredFunc("gocql.JoinHostPort")
	if ok {
		reta = FuncAuxMockJoinHostPort.(func(argaddr string, argport int) (reta string))(argaddr, argport)
	} else {
		panic("FuncAuxMockJoinHostPort ")
	}
	AuxMockIncrementRecorderAuxMockJoinHostPort()
	return
}

//
// Mock: (recvc *Conn)Closed()(reta bool)
//

type MockArgsTypeConnClosed struct {
	ApomockCallNumber int
}

var LastMockArgsConnClosed MockArgsTypeConnClosed

// (recvc *Conn)AuxMockClosed()(reta bool) - Generated mock function
func (recvc *Conn) AuxMockClosed() (reta bool) {
	rargs, rerr := apomock.GetNext("gocql.Conn.Closed")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.Closed")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.Closed")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrConnClosed  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnClosed int = 0

var condRecorderAuxMockPtrConnClosed *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnClosed(i int) {
	condRecorderAuxMockPtrConnClosed.L.Lock()
	for recorderAuxMockPtrConnClosed < i {
		condRecorderAuxMockPtrConnClosed.Wait()
	}
	condRecorderAuxMockPtrConnClosed.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnClosed() {
	condRecorderAuxMockPtrConnClosed.L.Lock()
	recorderAuxMockPtrConnClosed++
	condRecorderAuxMockPtrConnClosed.L.Unlock()
	condRecorderAuxMockPtrConnClosed.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnClosed() (ret int) {
	condRecorderAuxMockPtrConnClosed.L.Lock()
	ret = recorderAuxMockPtrConnClosed
	condRecorderAuxMockPtrConnClosed.L.Unlock()
	return
}

// (recvc *Conn)Closed - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) Closed() (reta bool) {
	FuncAuxMockPtrConnClosed, ok := apomock.GetRegisteredFunc("gocql.Conn.Closed")
	if ok {
		reta = FuncAuxMockPtrConnClosed.(func(recvc *Conn) (reta bool))(recvc)
	} else {
		panic("FuncAuxMockPtrConnClosed ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnClosed()
	return
}

//
// Mock: (recvc *Conn)UseKeyspace(argkeyspace string)(reta error)
//

type MockArgsTypeConnUseKeyspace struct {
	ApomockCallNumber int
	Argkeyspace       string
}

var LastMockArgsConnUseKeyspace MockArgsTypeConnUseKeyspace

// (recvc *Conn)AuxMockUseKeyspace(argkeyspace string)(reta error) - Generated mock function
func (recvc *Conn) AuxMockUseKeyspace(argkeyspace string) (reta error) {
	LastMockArgsConnUseKeyspace = MockArgsTypeConnUseKeyspace{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnUseKeyspace(),
		Argkeyspace:       argkeyspace,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.UseKeyspace")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.UseKeyspace")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.UseKeyspace")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrConnUseKeyspace  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnUseKeyspace int = 0

var condRecorderAuxMockPtrConnUseKeyspace *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnUseKeyspace(i int) {
	condRecorderAuxMockPtrConnUseKeyspace.L.Lock()
	for recorderAuxMockPtrConnUseKeyspace < i {
		condRecorderAuxMockPtrConnUseKeyspace.Wait()
	}
	condRecorderAuxMockPtrConnUseKeyspace.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnUseKeyspace() {
	condRecorderAuxMockPtrConnUseKeyspace.L.Lock()
	recorderAuxMockPtrConnUseKeyspace++
	condRecorderAuxMockPtrConnUseKeyspace.L.Unlock()
	condRecorderAuxMockPtrConnUseKeyspace.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnUseKeyspace() (ret int) {
	condRecorderAuxMockPtrConnUseKeyspace.L.Lock()
	ret = recorderAuxMockPtrConnUseKeyspace
	condRecorderAuxMockPtrConnUseKeyspace.L.Unlock()
	return
}

// (recvc *Conn)UseKeyspace - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) UseKeyspace(argkeyspace string) (reta error) {
	FuncAuxMockPtrConnUseKeyspace, ok := apomock.GetRegisteredFunc("gocql.Conn.UseKeyspace")
	if ok {
		reta = FuncAuxMockPtrConnUseKeyspace.(func(recvc *Conn, argkeyspace string) (reta error))(recvc, argkeyspace)
	} else {
		panic("FuncAuxMockPtrConnUseKeyspace ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnUseKeyspace()
	return
}

//
// Mock: approve(argauthenticator string)(reta bool)
//

type MockArgsTypeapprove struct {
	ApomockCallNumber int
	Argauthenticator  string
}

var LastMockArgsapprove MockArgsTypeapprove

// AuxMockapprove(argauthenticator string)(reta bool) - Generated mock function
func AuxMockapprove(argauthenticator string) (reta bool) {
	LastMockArgsapprove = MockArgsTypeapprove{
		ApomockCallNumber: AuxMockGetRecorderAuxMockapprove(),
		Argauthenticator:  argauthenticator,
	}
	rargs, rerr := apomock.GetNext("gocql.approve")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.approve")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.approve")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockapprove  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockapprove int = 0

var condRecorderAuxMockapprove *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockapprove(i int) {
	condRecorderAuxMockapprove.L.Lock()
	for recorderAuxMockapprove < i {
		condRecorderAuxMockapprove.Wait()
	}
	condRecorderAuxMockapprove.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockapprove() {
	condRecorderAuxMockapprove.L.Lock()
	recorderAuxMockapprove++
	condRecorderAuxMockapprove.L.Unlock()
	condRecorderAuxMockapprove.Broadcast()
}
func AuxMockGetRecorderAuxMockapprove() (ret int) {
	condRecorderAuxMockapprove.L.Lock()
	ret = recorderAuxMockapprove
	condRecorderAuxMockapprove.L.Unlock()
	return
}

// approve - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func approve(argauthenticator string) (reta bool) {
	FuncAuxMockapprove, ok := apomock.GetRegisteredFunc("gocql.approve")
	if ok {
		reta = FuncAuxMockapprove.(func(argauthenticator string) (reta bool))(argauthenticator)
	} else {
		panic("FuncAuxMockapprove ")
	}
	AuxMockIncrementRecorderAuxMockapprove()
	return
}

//
// Mock: (recvc *Conn)serve()()
//

type MockArgsTypeConnserve struct {
	ApomockCallNumber int
}

var LastMockArgsConnserve MockArgsTypeConnserve

// (recvc *Conn)AuxMockserve()() - Generated mock function
func (recvc *Conn) AuxMockserve() {
	return
}

// RecorderAuxMockPtrConnserve  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnserve int = 0

var condRecorderAuxMockPtrConnserve *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnserve(i int) {
	condRecorderAuxMockPtrConnserve.L.Lock()
	for recorderAuxMockPtrConnserve < i {
		condRecorderAuxMockPtrConnserve.Wait()
	}
	condRecorderAuxMockPtrConnserve.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnserve() {
	condRecorderAuxMockPtrConnserve.L.Lock()
	recorderAuxMockPtrConnserve++
	condRecorderAuxMockPtrConnserve.L.Unlock()
	condRecorderAuxMockPtrConnserve.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnserve() (ret int) {
	condRecorderAuxMockPtrConnserve.L.Lock()
	ret = recorderAuxMockPtrConnserve
	condRecorderAuxMockPtrConnserve.L.Unlock()
	return
}

// (recvc *Conn)serve - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) serve() {
	FuncAuxMockPtrConnserve, ok := apomock.GetRegisteredFunc("gocql.Conn.serve")
	if ok {
		FuncAuxMockPtrConnserve.(func(recvc *Conn))(recvc)
	} else {
		panic("FuncAuxMockPtrConnserve ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnserve()
	return
}

//
// Mock: (recvc *Conn)AvailableStreams()(reta int)
//

type MockArgsTypeConnAvailableStreams struct {
	ApomockCallNumber int
}

var LastMockArgsConnAvailableStreams MockArgsTypeConnAvailableStreams

// (recvc *Conn)AuxMockAvailableStreams()(reta int) - Generated mock function
func (recvc *Conn) AuxMockAvailableStreams() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.Conn.AvailableStreams")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.AvailableStreams")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.AvailableStreams")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrConnAvailableStreams  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnAvailableStreams int = 0

var condRecorderAuxMockPtrConnAvailableStreams *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnAvailableStreams(i int) {
	condRecorderAuxMockPtrConnAvailableStreams.L.Lock()
	for recorderAuxMockPtrConnAvailableStreams < i {
		condRecorderAuxMockPtrConnAvailableStreams.Wait()
	}
	condRecorderAuxMockPtrConnAvailableStreams.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnAvailableStreams() {
	condRecorderAuxMockPtrConnAvailableStreams.L.Lock()
	recorderAuxMockPtrConnAvailableStreams++
	condRecorderAuxMockPtrConnAvailableStreams.L.Unlock()
	condRecorderAuxMockPtrConnAvailableStreams.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnAvailableStreams() (ret int) {
	condRecorderAuxMockPtrConnAvailableStreams.L.Lock()
	ret = recorderAuxMockPtrConnAvailableStreams
	condRecorderAuxMockPtrConnAvailableStreams.L.Unlock()
	return
}

// (recvc *Conn)AvailableStreams - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) AvailableStreams() (reta int) {
	FuncAuxMockPtrConnAvailableStreams, ok := apomock.GetRegisteredFunc("gocql.Conn.AvailableStreams")
	if ok {
		reta = FuncAuxMockPtrConnAvailableStreams.(func(recvc *Conn) (reta int))(recvc)
	} else {
		panic("FuncAuxMockPtrConnAvailableStreams ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnAvailableStreams()
	return
}

//
// Mock: (recvfn connErrorHandlerFn)HandleError(argconn *Conn, argerr error, argclosed bool)()
//

type MockArgsTypeconnErrorHandlerFnHandleError struct {
	ApomockCallNumber int
	Argconn           *Conn
	Argerr            error
	Argclosed         bool
}

var LastMockArgsconnErrorHandlerFnHandleError MockArgsTypeconnErrorHandlerFnHandleError

// (recvfn connErrorHandlerFn)AuxMockHandleError(argconn *Conn, argerr error, argclosed bool)() - Generated mock function
func (recvfn connErrorHandlerFn) AuxMockHandleError(argconn *Conn, argerr error, argclosed bool) {
	LastMockArgsconnErrorHandlerFnHandleError = MockArgsTypeconnErrorHandlerFnHandleError{
		ApomockCallNumber: AuxMockGetRecorderAuxMockconnErrorHandlerFnHandleError(),
		Argconn:           argconn,
		Argerr:            argerr,
		Argclosed:         argclosed,
	}
	return
}

// RecorderAuxMockconnErrorHandlerFnHandleError  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockconnErrorHandlerFnHandleError int = 0

var condRecorderAuxMockconnErrorHandlerFnHandleError *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockconnErrorHandlerFnHandleError(i int) {
	condRecorderAuxMockconnErrorHandlerFnHandleError.L.Lock()
	for recorderAuxMockconnErrorHandlerFnHandleError < i {
		condRecorderAuxMockconnErrorHandlerFnHandleError.Wait()
	}
	condRecorderAuxMockconnErrorHandlerFnHandleError.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockconnErrorHandlerFnHandleError() {
	condRecorderAuxMockconnErrorHandlerFnHandleError.L.Lock()
	recorderAuxMockconnErrorHandlerFnHandleError++
	condRecorderAuxMockconnErrorHandlerFnHandleError.L.Unlock()
	condRecorderAuxMockconnErrorHandlerFnHandleError.Broadcast()
}
func AuxMockGetRecorderAuxMockconnErrorHandlerFnHandleError() (ret int) {
	condRecorderAuxMockconnErrorHandlerFnHandleError.L.Lock()
	ret = recorderAuxMockconnErrorHandlerFnHandleError
	condRecorderAuxMockconnErrorHandlerFnHandleError.L.Unlock()
	return
}

// (recvfn connErrorHandlerFn)HandleError - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvfn connErrorHandlerFn) HandleError(argconn *Conn, argerr error, argclosed bool) {
	FuncAuxMockconnErrorHandlerFnHandleError, ok := apomock.GetRegisteredFunc("gocql.connErrorHandlerFn.HandleError")
	if ok {
		FuncAuxMockconnErrorHandlerFnHandleError.(func(recvfn connErrorHandlerFn, argconn *Conn, argerr error, argclosed bool))(recvfn, argconn, argerr, argclosed)
	} else {
		panic("FuncAuxMockconnErrorHandlerFnHandleError ")
	}
	AuxMockIncrementRecorderAuxMockconnErrorHandlerFnHandleError()
	return
}

//
// Mock: (recvc *Conn)handleTimeout()()
//

type MockArgsTypeConnhandleTimeout struct {
	ApomockCallNumber int
}

var LastMockArgsConnhandleTimeout MockArgsTypeConnhandleTimeout

// (recvc *Conn)AuxMockhandleTimeout()() - Generated mock function
func (recvc *Conn) AuxMockhandleTimeout() {
	return
}

// RecorderAuxMockPtrConnhandleTimeout  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnhandleTimeout int = 0

var condRecorderAuxMockPtrConnhandleTimeout *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnhandleTimeout(i int) {
	condRecorderAuxMockPtrConnhandleTimeout.L.Lock()
	for recorderAuxMockPtrConnhandleTimeout < i {
		condRecorderAuxMockPtrConnhandleTimeout.Wait()
	}
	condRecorderAuxMockPtrConnhandleTimeout.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnhandleTimeout() {
	condRecorderAuxMockPtrConnhandleTimeout.L.Lock()
	recorderAuxMockPtrConnhandleTimeout++
	condRecorderAuxMockPtrConnhandleTimeout.L.Unlock()
	condRecorderAuxMockPtrConnhandleTimeout.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnhandleTimeout() (ret int) {
	condRecorderAuxMockPtrConnhandleTimeout.L.Lock()
	ret = recorderAuxMockPtrConnhandleTimeout
	condRecorderAuxMockPtrConnhandleTimeout.L.Unlock()
	return
}

// (recvc *Conn)handleTimeout - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) handleTimeout() {
	FuncAuxMockPtrConnhandleTimeout, ok := apomock.GetRegisteredFunc("gocql.Conn.handleTimeout")
	if ok {
		FuncAuxMockPtrConnhandleTimeout.(func(recvc *Conn))(recvc)
	} else {
		panic("FuncAuxMockPtrConnhandleTimeout ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnhandleTimeout()
	return
}

//
// Mock: (recvc *Conn)prepareStatement(argctx context.Context, argstmt string, argtracer Tracer)(reta *preparedStatment, retb error)
//

type MockArgsTypeConnprepareStatement struct {
	ApomockCallNumber int
	Argctx            context.Context
	Argstmt           string
	Argtracer         Tracer
}

var LastMockArgsConnprepareStatement MockArgsTypeConnprepareStatement

// (recvc *Conn)AuxMockprepareStatement(argctx context.Context, argstmt string, argtracer Tracer)(reta *preparedStatment, retb error) - Generated mock function
func (recvc *Conn) AuxMockprepareStatement(argctx context.Context, argstmt string, argtracer Tracer) (reta *preparedStatment, retb error) {
	LastMockArgsConnprepareStatement = MockArgsTypeConnprepareStatement{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnprepareStatement(),
		Argctx:            argctx,
		Argstmt:           argstmt,
		Argtracer:         argtracer,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.prepareStatement")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.prepareStatement")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Conn.prepareStatement")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*preparedStatment)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrConnprepareStatement  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnprepareStatement int = 0

var condRecorderAuxMockPtrConnprepareStatement *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnprepareStatement(i int) {
	condRecorderAuxMockPtrConnprepareStatement.L.Lock()
	for recorderAuxMockPtrConnprepareStatement < i {
		condRecorderAuxMockPtrConnprepareStatement.Wait()
	}
	condRecorderAuxMockPtrConnprepareStatement.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnprepareStatement() {
	condRecorderAuxMockPtrConnprepareStatement.L.Lock()
	recorderAuxMockPtrConnprepareStatement++
	condRecorderAuxMockPtrConnprepareStatement.L.Unlock()
	condRecorderAuxMockPtrConnprepareStatement.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnprepareStatement() (ret int) {
	condRecorderAuxMockPtrConnprepareStatement.L.Lock()
	ret = recorderAuxMockPtrConnprepareStatement
	condRecorderAuxMockPtrConnprepareStatement.L.Unlock()
	return
}

// (recvc *Conn)prepareStatement - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) prepareStatement(argctx context.Context, argstmt string, argtracer Tracer) (reta *preparedStatment, retb error) {
	FuncAuxMockPtrConnprepareStatement, ok := apomock.GetRegisteredFunc("gocql.Conn.prepareStatement")
	if ok {
		reta, retb = FuncAuxMockPtrConnprepareStatement.(func(recvc *Conn, argctx context.Context, argstmt string, argtracer Tracer) (reta *preparedStatment, retb error))(recvc, argctx, argstmt, argtracer)
	} else {
		panic("FuncAuxMockPtrConnprepareStatement ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnprepareStatement()
	return
}

//
// Mock: (recvc *Conn)executeBatch(argbatch *Batch)(reta *Iter)
//

type MockArgsTypeConnexecuteBatch struct {
	ApomockCallNumber int
	Argbatch          *Batch
}

var LastMockArgsConnexecuteBatch MockArgsTypeConnexecuteBatch

// (recvc *Conn)AuxMockexecuteBatch(argbatch *Batch)(reta *Iter) - Generated mock function
func (recvc *Conn) AuxMockexecuteBatch(argbatch *Batch) (reta *Iter) {
	LastMockArgsConnexecuteBatch = MockArgsTypeConnexecuteBatch{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnexecuteBatch(),
		Argbatch:          argbatch,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.executeBatch")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.executeBatch")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.executeBatch")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Iter)
	}
	return
}

// RecorderAuxMockPtrConnexecuteBatch  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnexecuteBatch int = 0

var condRecorderAuxMockPtrConnexecuteBatch *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnexecuteBatch(i int) {
	condRecorderAuxMockPtrConnexecuteBatch.L.Lock()
	for recorderAuxMockPtrConnexecuteBatch < i {
		condRecorderAuxMockPtrConnexecuteBatch.Wait()
	}
	condRecorderAuxMockPtrConnexecuteBatch.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnexecuteBatch() {
	condRecorderAuxMockPtrConnexecuteBatch.L.Lock()
	recorderAuxMockPtrConnexecuteBatch++
	condRecorderAuxMockPtrConnexecuteBatch.L.Unlock()
	condRecorderAuxMockPtrConnexecuteBatch.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnexecuteBatch() (ret int) {
	condRecorderAuxMockPtrConnexecuteBatch.L.Lock()
	ret = recorderAuxMockPtrConnexecuteBatch
	condRecorderAuxMockPtrConnexecuteBatch.L.Unlock()
	return
}

// (recvc *Conn)executeBatch - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) executeBatch(argbatch *Batch) (reta *Iter) {
	FuncAuxMockPtrConnexecuteBatch, ok := apomock.GetRegisteredFunc("gocql.Conn.executeBatch")
	if ok {
		reta = FuncAuxMockPtrConnexecuteBatch.(func(recvc *Conn, argbatch *Batch) (reta *Iter))(recvc, argbatch)
	} else {
		panic("FuncAuxMockPtrConnexecuteBatch ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnexecuteBatch()
	return
}

//
// Mock: (recvc *Conn)Write(argp []byte)(reta int, retb error)
//

type MockArgsTypeConnWrite struct {
	ApomockCallNumber int
	Argp              []byte
}

var LastMockArgsConnWrite MockArgsTypeConnWrite

// (recvc *Conn)AuxMockWrite(argp []byte)(reta int, retb error) - Generated mock function
func (recvc *Conn) AuxMockWrite(argp []byte) (reta int, retb error) {
	LastMockArgsConnWrite = MockArgsTypeConnWrite{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnWrite(),
		Argp:              argp,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.Write")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.Write")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Conn.Write")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrConnWrite  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnWrite int = 0

var condRecorderAuxMockPtrConnWrite *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnWrite(i int) {
	condRecorderAuxMockPtrConnWrite.L.Lock()
	for recorderAuxMockPtrConnWrite < i {
		condRecorderAuxMockPtrConnWrite.Wait()
	}
	condRecorderAuxMockPtrConnWrite.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnWrite() {
	condRecorderAuxMockPtrConnWrite.L.Lock()
	recorderAuxMockPtrConnWrite++
	condRecorderAuxMockPtrConnWrite.L.Unlock()
	condRecorderAuxMockPtrConnWrite.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnWrite() (ret int) {
	condRecorderAuxMockPtrConnWrite.L.Lock()
	ret = recorderAuxMockPtrConnWrite
	condRecorderAuxMockPtrConnWrite.L.Unlock()
	return
}

// (recvc *Conn)Write - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) Write(argp []byte) (reta int, retb error) {
	FuncAuxMockPtrConnWrite, ok := apomock.GetRegisteredFunc("gocql.Conn.Write")
	if ok {
		reta, retb = FuncAuxMockPtrConnWrite.(func(recvc *Conn, argp []byte) (reta int, retb error))(recvc, argp)
	} else {
		panic("FuncAuxMockPtrConnWrite ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnWrite()
	return
}

//
// Mock: (recvc *Conn)startup(argctx context.Context, argframeTicker chan struct{})(reta error)
//

type MockArgsTypeConnstartup struct {
	ApomockCallNumber int
	Argctx            context.Context
	ArgframeTicker    chan struct{}
}

var LastMockArgsConnstartup MockArgsTypeConnstartup

// (recvc *Conn)AuxMockstartup(argctx context.Context, argframeTicker chan struct{})(reta error) - Generated mock function
func (recvc *Conn) AuxMockstartup(argctx context.Context, argframeTicker chan struct{}) (reta error) {
	LastMockArgsConnstartup = MockArgsTypeConnstartup{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnstartup(),
		Argctx:            argctx,
		ArgframeTicker:    argframeTicker,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.startup")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.startup")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.startup")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrConnstartup  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnstartup int = 0

var condRecorderAuxMockPtrConnstartup *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnstartup(i int) {
	condRecorderAuxMockPtrConnstartup.L.Lock()
	for recorderAuxMockPtrConnstartup < i {
		condRecorderAuxMockPtrConnstartup.Wait()
	}
	condRecorderAuxMockPtrConnstartup.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnstartup() {
	condRecorderAuxMockPtrConnstartup.L.Lock()
	recorderAuxMockPtrConnstartup++
	condRecorderAuxMockPtrConnstartup.L.Unlock()
	condRecorderAuxMockPtrConnstartup.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnstartup() (ret int) {
	condRecorderAuxMockPtrConnstartup.L.Lock()
	ret = recorderAuxMockPtrConnstartup
	condRecorderAuxMockPtrConnstartup.L.Unlock()
	return
}

// (recvc *Conn)startup - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) startup(argctx context.Context, argframeTicker chan struct{}) (reta error) {
	FuncAuxMockPtrConnstartup, ok := apomock.GetRegisteredFunc("gocql.Conn.startup")
	if ok {
		reta = FuncAuxMockPtrConnstartup.(func(recvc *Conn, argctx context.Context, argframeTicker chan struct{}) (reta error))(recvc, argctx, argframeTicker)
	} else {
		panic("FuncAuxMockPtrConnstartup ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnstartup()
	return
}

//
// Mock: (recvc *Conn)exec(argctx context.Context, argreq frameWriter, argtracer Tracer)(reta *framer, retb error)
//

type MockArgsTypeConnexec struct {
	ApomockCallNumber int
	Argctx            context.Context
	Argreq            frameWriter
	Argtracer         Tracer
}

var LastMockArgsConnexec MockArgsTypeConnexec

// (recvc *Conn)AuxMockexec(argctx context.Context, argreq frameWriter, argtracer Tracer)(reta *framer, retb error) - Generated mock function
func (recvc *Conn) AuxMockexec(argctx context.Context, argreq frameWriter, argtracer Tracer) (reta *framer, retb error) {
	LastMockArgsConnexec = MockArgsTypeConnexec{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnexec(),
		Argctx:            argctx,
		Argreq:            argreq,
		Argtracer:         argtracer,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.exec")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.exec")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Conn.exec")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*framer)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrConnexec  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnexec int = 0

var condRecorderAuxMockPtrConnexec *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnexec(i int) {
	condRecorderAuxMockPtrConnexec.L.Lock()
	for recorderAuxMockPtrConnexec < i {
		condRecorderAuxMockPtrConnexec.Wait()
	}
	condRecorderAuxMockPtrConnexec.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnexec() {
	condRecorderAuxMockPtrConnexec.L.Lock()
	recorderAuxMockPtrConnexec++
	condRecorderAuxMockPtrConnexec.L.Unlock()
	condRecorderAuxMockPtrConnexec.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnexec() (ret int) {
	condRecorderAuxMockPtrConnexec.L.Lock()
	ret = recorderAuxMockPtrConnexec
	condRecorderAuxMockPtrConnexec.L.Unlock()
	return
}

// (recvc *Conn)exec - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) exec(argctx context.Context, argreq frameWriter, argtracer Tracer) (reta *framer, retb error) {
	FuncAuxMockPtrConnexec, ok := apomock.GetRegisteredFunc("gocql.Conn.exec")
	if ok {
		reta, retb = FuncAuxMockPtrConnexec.(func(recvc *Conn, argctx context.Context, argreq frameWriter, argtracer Tracer) (reta *framer, retb error))(recvc, argctx, argreq, argtracer)
	} else {
		panic("FuncAuxMockPtrConnexec ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnexec()
	return
}

//
// Mock: (recvc *Conn)query(argstatement string, values ...interface{})(retiter *Iter)
//

type MockArgsTypeConnquery struct {
	ApomockCallNumber int
	Argstatement      string
	Values            []interface{}
}

var LastMockArgsConnquery MockArgsTypeConnquery

// (recvc *Conn)AuxMockquery(argstatement string, values ...interface{})(retiter *Iter) - Generated mock function
func (recvc *Conn) AuxMockquery(argstatement string, values ...interface{}) (retiter *Iter) {
	LastMockArgsConnquery = MockArgsTypeConnquery{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrConnquery(),
		Argstatement:      argstatement,
		Values:            values,
	}
	rargs, rerr := apomock.GetNext("gocql.Conn.query")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Conn.query")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Conn.query")
	}
	if rargs.GetArg(0) != nil {
		retiter = rargs.GetArg(0).(*Iter)
	}
	return
}

// RecorderAuxMockPtrConnquery  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrConnquery int = 0

var condRecorderAuxMockPtrConnquery *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrConnquery(i int) {
	condRecorderAuxMockPtrConnquery.L.Lock()
	for recorderAuxMockPtrConnquery < i {
		condRecorderAuxMockPtrConnquery.Wait()
	}
	condRecorderAuxMockPtrConnquery.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrConnquery() {
	condRecorderAuxMockPtrConnquery.L.Lock()
	recorderAuxMockPtrConnquery++
	condRecorderAuxMockPtrConnquery.L.Unlock()
	condRecorderAuxMockPtrConnquery.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrConnquery() (ret int) {
	condRecorderAuxMockPtrConnquery.L.Lock()
	ret = recorderAuxMockPtrConnquery
	condRecorderAuxMockPtrConnquery.L.Unlock()
	return
}

// (recvc *Conn)query - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Conn) query(argstatement string, values ...interface{}) (retiter *Iter) {
	FuncAuxMockPtrConnquery, ok := apomock.GetRegisteredFunc("gocql.Conn.query")
	if ok {
		retiter = FuncAuxMockPtrConnquery.(func(recvc *Conn, argstatement string, values ...interface{}) (retiter *Iter))(recvc, argstatement, values...)
	} else {
		panic("FuncAuxMockPtrConnquery ")
	}
	AuxMockIncrementRecorderAuxMockPtrConnquery()
	return
}
