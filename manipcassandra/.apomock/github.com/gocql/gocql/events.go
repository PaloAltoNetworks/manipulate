// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

import "net"
import "time"

func init() {
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructEventDeouncer, apomockNewStructEventDeouncer)

	apomock.RegisterFunc("gocql", "gocql.Session.handleSchemaEvent", (*Session).AuxMockhandleSchemaEvent)
	apomock.RegisterFunc("gocql", "gocql.Session.handleNodeEvent", (*Session).AuxMockhandleNodeEvent)
	apomock.RegisterFunc("gocql", "gocql.Session.handleNodeUp", (*Session).AuxMockhandleNodeUp)
	apomock.RegisterFunc("gocql", "gocql.Session.handleNodeDown", (*Session).AuxMockhandleNodeDown)
	apomock.RegisterFunc("gocql", "gocql.eventDeouncer.stop", (*eventDeouncer).AuxMockstop)
	apomock.RegisterFunc("gocql", "gocql.eventDeouncer.flusher", (*eventDeouncer).AuxMockflusher)
	apomock.RegisterFunc("gocql", "gocql.eventDeouncer.flush", (*eventDeouncer).AuxMockflush)
	apomock.RegisterFunc("gocql", "gocql.Session.handleEvent", (*Session).AuxMockhandleEvent)
	apomock.RegisterFunc("gocql", "gocql.newEventDeouncer", AuxMocknewEventDeouncer)
	apomock.RegisterFunc("gocql", "gocql.eventDeouncer.debounce", (*eventDeouncer).AuxMockdebounce)
	apomock.RegisterFunc("gocql", "gocql.Session.handleNewNode", (*Session).AuxMockhandleNewNode)
	apomock.RegisterFunc("gocql", "gocql.Session.handleRemovedNode", (*Session).AuxMockhandleRemovedNode)
}

const (
	eventBufferSize   = 1000
	eventDebounceTime = 1 * time.Second
)

const (
	ApomockStructEventDeouncer = 47
)

//
// Internal Types: in this package and their exportable versions
//
type eventDeouncer struct {
	name     string
	timer    *time.Timer
	mu       sync.Mutex
	events   []frame
	callback func([]frame)
	quit     chan struct{}
}

//
// External Types: in this package
//

func apomockNewStructEventDeouncer() interface{} { return &eventDeouncer{} }

//
// Mock: (recvs *Session)handleSchemaEvent(argframes []frame)()
//

type MockArgsTypeSessionhandleSchemaEvent struct {
	ApomockCallNumber int
	Argframes         []frame
}

var LastMockArgsSessionhandleSchemaEvent MockArgsTypeSessionhandleSchemaEvent

// (recvs *Session)AuxMockhandleSchemaEvent(argframes []frame)() - Generated mock function
func (recvs *Session) AuxMockhandleSchemaEvent(argframes []frame) {
	LastMockArgsSessionhandleSchemaEvent = MockArgsTypeSessionhandleSchemaEvent{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionhandleSchemaEvent(),
		Argframes:         argframes,
	}
	return
}

// RecorderAuxMockPtrSessionhandleSchemaEvent  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionhandleSchemaEvent int = 0

var condRecorderAuxMockPtrSessionhandleSchemaEvent *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionhandleSchemaEvent(i int) {
	condRecorderAuxMockPtrSessionhandleSchemaEvent.L.Lock()
	for recorderAuxMockPtrSessionhandleSchemaEvent < i {
		condRecorderAuxMockPtrSessionhandleSchemaEvent.Wait()
	}
	condRecorderAuxMockPtrSessionhandleSchemaEvent.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionhandleSchemaEvent() {
	condRecorderAuxMockPtrSessionhandleSchemaEvent.L.Lock()
	recorderAuxMockPtrSessionhandleSchemaEvent++
	condRecorderAuxMockPtrSessionhandleSchemaEvent.L.Unlock()
	condRecorderAuxMockPtrSessionhandleSchemaEvent.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionhandleSchemaEvent() (ret int) {
	condRecorderAuxMockPtrSessionhandleSchemaEvent.L.Lock()
	ret = recorderAuxMockPtrSessionhandleSchemaEvent
	condRecorderAuxMockPtrSessionhandleSchemaEvent.L.Unlock()
	return
}

// (recvs *Session)handleSchemaEvent - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) handleSchemaEvent(argframes []frame) {
	FuncAuxMockPtrSessionhandleSchemaEvent, ok := apomock.GetRegisteredFunc("gocql.Session.handleSchemaEvent")
	if ok {
		FuncAuxMockPtrSessionhandleSchemaEvent.(func(recvs *Session, argframes []frame))(recvs, argframes)
	} else {
		panic("FuncAuxMockPtrSessionhandleSchemaEvent ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionhandleSchemaEvent()
	return
}

//
// Mock: (recvs *Session)handleNodeEvent(argframes []frame)()
//

type MockArgsTypeSessionhandleNodeEvent struct {
	ApomockCallNumber int
	Argframes         []frame
}

var LastMockArgsSessionhandleNodeEvent MockArgsTypeSessionhandleNodeEvent

// (recvs *Session)AuxMockhandleNodeEvent(argframes []frame)() - Generated mock function
func (recvs *Session) AuxMockhandleNodeEvent(argframes []frame) {
	LastMockArgsSessionhandleNodeEvent = MockArgsTypeSessionhandleNodeEvent{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionhandleNodeEvent(),
		Argframes:         argframes,
	}
	return
}

// RecorderAuxMockPtrSessionhandleNodeEvent  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionhandleNodeEvent int = 0

var condRecorderAuxMockPtrSessionhandleNodeEvent *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionhandleNodeEvent(i int) {
	condRecorderAuxMockPtrSessionhandleNodeEvent.L.Lock()
	for recorderAuxMockPtrSessionhandleNodeEvent < i {
		condRecorderAuxMockPtrSessionhandleNodeEvent.Wait()
	}
	condRecorderAuxMockPtrSessionhandleNodeEvent.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionhandleNodeEvent() {
	condRecorderAuxMockPtrSessionhandleNodeEvent.L.Lock()
	recorderAuxMockPtrSessionhandleNodeEvent++
	condRecorderAuxMockPtrSessionhandleNodeEvent.L.Unlock()
	condRecorderAuxMockPtrSessionhandleNodeEvent.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionhandleNodeEvent() (ret int) {
	condRecorderAuxMockPtrSessionhandleNodeEvent.L.Lock()
	ret = recorderAuxMockPtrSessionhandleNodeEvent
	condRecorderAuxMockPtrSessionhandleNodeEvent.L.Unlock()
	return
}

// (recvs *Session)handleNodeEvent - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) handleNodeEvent(argframes []frame) {
	FuncAuxMockPtrSessionhandleNodeEvent, ok := apomock.GetRegisteredFunc("gocql.Session.handleNodeEvent")
	if ok {
		FuncAuxMockPtrSessionhandleNodeEvent.(func(recvs *Session, argframes []frame))(recvs, argframes)
	} else {
		panic("FuncAuxMockPtrSessionhandleNodeEvent ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionhandleNodeEvent()
	return
}

//
// Mock: (recvs *Session)handleNodeUp(argip net.IP, argport int, argwaitForBinary bool)()
//

type MockArgsTypeSessionhandleNodeUp struct {
	ApomockCallNumber int
	Argip             net.IP
	Argport           int
	ArgwaitForBinary  bool
}

var LastMockArgsSessionhandleNodeUp MockArgsTypeSessionhandleNodeUp

// (recvs *Session)AuxMockhandleNodeUp(argip net.IP, argport int, argwaitForBinary bool)() - Generated mock function
func (recvs *Session) AuxMockhandleNodeUp(argip net.IP, argport int, argwaitForBinary bool) {
	LastMockArgsSessionhandleNodeUp = MockArgsTypeSessionhandleNodeUp{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionhandleNodeUp(),
		Argip:             argip,
		Argport:           argport,
		ArgwaitForBinary:  argwaitForBinary,
	}
	return
}

// RecorderAuxMockPtrSessionhandleNodeUp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionhandleNodeUp int = 0

var condRecorderAuxMockPtrSessionhandleNodeUp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionhandleNodeUp(i int) {
	condRecorderAuxMockPtrSessionhandleNodeUp.L.Lock()
	for recorderAuxMockPtrSessionhandleNodeUp < i {
		condRecorderAuxMockPtrSessionhandleNodeUp.Wait()
	}
	condRecorderAuxMockPtrSessionhandleNodeUp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionhandleNodeUp() {
	condRecorderAuxMockPtrSessionhandleNodeUp.L.Lock()
	recorderAuxMockPtrSessionhandleNodeUp++
	condRecorderAuxMockPtrSessionhandleNodeUp.L.Unlock()
	condRecorderAuxMockPtrSessionhandleNodeUp.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionhandleNodeUp() (ret int) {
	condRecorderAuxMockPtrSessionhandleNodeUp.L.Lock()
	ret = recorderAuxMockPtrSessionhandleNodeUp
	condRecorderAuxMockPtrSessionhandleNodeUp.L.Unlock()
	return
}

// (recvs *Session)handleNodeUp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) handleNodeUp(argip net.IP, argport int, argwaitForBinary bool) {
	FuncAuxMockPtrSessionhandleNodeUp, ok := apomock.GetRegisteredFunc("gocql.Session.handleNodeUp")
	if ok {
		FuncAuxMockPtrSessionhandleNodeUp.(func(recvs *Session, argip net.IP, argport int, argwaitForBinary bool))(recvs, argip, argport, argwaitForBinary)
	} else {
		panic("FuncAuxMockPtrSessionhandleNodeUp ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionhandleNodeUp()
	return
}

//
// Mock: (recvs *Session)handleNodeDown(argip net.IP, argport int)()
//

type MockArgsTypeSessionhandleNodeDown struct {
	ApomockCallNumber int
	Argip             net.IP
	Argport           int
}

var LastMockArgsSessionhandleNodeDown MockArgsTypeSessionhandleNodeDown

// (recvs *Session)AuxMockhandleNodeDown(argip net.IP, argport int)() - Generated mock function
func (recvs *Session) AuxMockhandleNodeDown(argip net.IP, argport int) {
	LastMockArgsSessionhandleNodeDown = MockArgsTypeSessionhandleNodeDown{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionhandleNodeDown(),
		Argip:             argip,
		Argport:           argport,
	}
	return
}

// RecorderAuxMockPtrSessionhandleNodeDown  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionhandleNodeDown int = 0

var condRecorderAuxMockPtrSessionhandleNodeDown *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionhandleNodeDown(i int) {
	condRecorderAuxMockPtrSessionhandleNodeDown.L.Lock()
	for recorderAuxMockPtrSessionhandleNodeDown < i {
		condRecorderAuxMockPtrSessionhandleNodeDown.Wait()
	}
	condRecorderAuxMockPtrSessionhandleNodeDown.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionhandleNodeDown() {
	condRecorderAuxMockPtrSessionhandleNodeDown.L.Lock()
	recorderAuxMockPtrSessionhandleNodeDown++
	condRecorderAuxMockPtrSessionhandleNodeDown.L.Unlock()
	condRecorderAuxMockPtrSessionhandleNodeDown.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionhandleNodeDown() (ret int) {
	condRecorderAuxMockPtrSessionhandleNodeDown.L.Lock()
	ret = recorderAuxMockPtrSessionhandleNodeDown
	condRecorderAuxMockPtrSessionhandleNodeDown.L.Unlock()
	return
}

// (recvs *Session)handleNodeDown - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) handleNodeDown(argip net.IP, argport int) {
	FuncAuxMockPtrSessionhandleNodeDown, ok := apomock.GetRegisteredFunc("gocql.Session.handleNodeDown")
	if ok {
		FuncAuxMockPtrSessionhandleNodeDown.(func(recvs *Session, argip net.IP, argport int))(recvs, argip, argport)
	} else {
		panic("FuncAuxMockPtrSessionhandleNodeDown ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionhandleNodeDown()
	return
}

//
// Mock: (recve *eventDeouncer)stop()()
//

type MockArgsTypeeventDeouncerstop struct {
	ApomockCallNumber int
}

var LastMockArgseventDeouncerstop MockArgsTypeeventDeouncerstop

// (recve *eventDeouncer)AuxMockstop()() - Generated mock function
func (recve *eventDeouncer) AuxMockstop() {
	return
}

// RecorderAuxMockPtreventDeouncerstop  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtreventDeouncerstop int = 0

var condRecorderAuxMockPtreventDeouncerstop *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtreventDeouncerstop(i int) {
	condRecorderAuxMockPtreventDeouncerstop.L.Lock()
	for recorderAuxMockPtreventDeouncerstop < i {
		condRecorderAuxMockPtreventDeouncerstop.Wait()
	}
	condRecorderAuxMockPtreventDeouncerstop.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtreventDeouncerstop() {
	condRecorderAuxMockPtreventDeouncerstop.L.Lock()
	recorderAuxMockPtreventDeouncerstop++
	condRecorderAuxMockPtreventDeouncerstop.L.Unlock()
	condRecorderAuxMockPtreventDeouncerstop.Broadcast()
}
func AuxMockGetRecorderAuxMockPtreventDeouncerstop() (ret int) {
	condRecorderAuxMockPtreventDeouncerstop.L.Lock()
	ret = recorderAuxMockPtreventDeouncerstop
	condRecorderAuxMockPtreventDeouncerstop.L.Unlock()
	return
}

// (recve *eventDeouncer)stop - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve *eventDeouncer) stop() {
	FuncAuxMockPtreventDeouncerstop, ok := apomock.GetRegisteredFunc("gocql.eventDeouncer.stop")
	if ok {
		FuncAuxMockPtreventDeouncerstop.(func(recve *eventDeouncer))(recve)
	} else {
		panic("FuncAuxMockPtreventDeouncerstop ")
	}
	AuxMockIncrementRecorderAuxMockPtreventDeouncerstop()
	return
}

//
// Mock: (recve *eventDeouncer)flusher()()
//

type MockArgsTypeeventDeouncerflusher struct {
	ApomockCallNumber int
}

var LastMockArgseventDeouncerflusher MockArgsTypeeventDeouncerflusher

// (recve *eventDeouncer)AuxMockflusher()() - Generated mock function
func (recve *eventDeouncer) AuxMockflusher() {
	return
}

// RecorderAuxMockPtreventDeouncerflusher  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtreventDeouncerflusher int = 0

var condRecorderAuxMockPtreventDeouncerflusher *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtreventDeouncerflusher(i int) {
	condRecorderAuxMockPtreventDeouncerflusher.L.Lock()
	for recorderAuxMockPtreventDeouncerflusher < i {
		condRecorderAuxMockPtreventDeouncerflusher.Wait()
	}
	condRecorderAuxMockPtreventDeouncerflusher.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtreventDeouncerflusher() {
	condRecorderAuxMockPtreventDeouncerflusher.L.Lock()
	recorderAuxMockPtreventDeouncerflusher++
	condRecorderAuxMockPtreventDeouncerflusher.L.Unlock()
	condRecorderAuxMockPtreventDeouncerflusher.Broadcast()
}
func AuxMockGetRecorderAuxMockPtreventDeouncerflusher() (ret int) {
	condRecorderAuxMockPtreventDeouncerflusher.L.Lock()
	ret = recorderAuxMockPtreventDeouncerflusher
	condRecorderAuxMockPtreventDeouncerflusher.L.Unlock()
	return
}

// (recve *eventDeouncer)flusher - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve *eventDeouncer) flusher() {
	FuncAuxMockPtreventDeouncerflusher, ok := apomock.GetRegisteredFunc("gocql.eventDeouncer.flusher")
	if ok {
		FuncAuxMockPtreventDeouncerflusher.(func(recve *eventDeouncer))(recve)
	} else {
		panic("FuncAuxMockPtreventDeouncerflusher ")
	}
	AuxMockIncrementRecorderAuxMockPtreventDeouncerflusher()
	return
}

//
// Mock: (recve *eventDeouncer)flush()()
//

type MockArgsTypeeventDeouncerflush struct {
	ApomockCallNumber int
}

var LastMockArgseventDeouncerflush MockArgsTypeeventDeouncerflush

// (recve *eventDeouncer)AuxMockflush()() - Generated mock function
func (recve *eventDeouncer) AuxMockflush() {
	return
}

// RecorderAuxMockPtreventDeouncerflush  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtreventDeouncerflush int = 0

var condRecorderAuxMockPtreventDeouncerflush *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtreventDeouncerflush(i int) {
	condRecorderAuxMockPtreventDeouncerflush.L.Lock()
	for recorderAuxMockPtreventDeouncerflush < i {
		condRecorderAuxMockPtreventDeouncerflush.Wait()
	}
	condRecorderAuxMockPtreventDeouncerflush.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtreventDeouncerflush() {
	condRecorderAuxMockPtreventDeouncerflush.L.Lock()
	recorderAuxMockPtreventDeouncerflush++
	condRecorderAuxMockPtreventDeouncerflush.L.Unlock()
	condRecorderAuxMockPtreventDeouncerflush.Broadcast()
}
func AuxMockGetRecorderAuxMockPtreventDeouncerflush() (ret int) {
	condRecorderAuxMockPtreventDeouncerflush.L.Lock()
	ret = recorderAuxMockPtreventDeouncerflush
	condRecorderAuxMockPtreventDeouncerflush.L.Unlock()
	return
}

// (recve *eventDeouncer)flush - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve *eventDeouncer) flush() {
	FuncAuxMockPtreventDeouncerflush, ok := apomock.GetRegisteredFunc("gocql.eventDeouncer.flush")
	if ok {
		FuncAuxMockPtreventDeouncerflush.(func(recve *eventDeouncer))(recve)
	} else {
		panic("FuncAuxMockPtreventDeouncerflush ")
	}
	AuxMockIncrementRecorderAuxMockPtreventDeouncerflush()
	return
}

//
// Mock: (recvs *Session)handleEvent(argframer *framer)()
//

type MockArgsTypeSessionhandleEvent struct {
	ApomockCallNumber int
	Argframer         *framer
}

var LastMockArgsSessionhandleEvent MockArgsTypeSessionhandleEvent

// (recvs *Session)AuxMockhandleEvent(argframer *framer)() - Generated mock function
func (recvs *Session) AuxMockhandleEvent(argframer *framer) {
	LastMockArgsSessionhandleEvent = MockArgsTypeSessionhandleEvent{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionhandleEvent(),
		Argframer:         argframer,
	}
	return
}

// RecorderAuxMockPtrSessionhandleEvent  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionhandleEvent int = 0

var condRecorderAuxMockPtrSessionhandleEvent *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionhandleEvent(i int) {
	condRecorderAuxMockPtrSessionhandleEvent.L.Lock()
	for recorderAuxMockPtrSessionhandleEvent < i {
		condRecorderAuxMockPtrSessionhandleEvent.Wait()
	}
	condRecorderAuxMockPtrSessionhandleEvent.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionhandleEvent() {
	condRecorderAuxMockPtrSessionhandleEvent.L.Lock()
	recorderAuxMockPtrSessionhandleEvent++
	condRecorderAuxMockPtrSessionhandleEvent.L.Unlock()
	condRecorderAuxMockPtrSessionhandleEvent.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionhandleEvent() (ret int) {
	condRecorderAuxMockPtrSessionhandleEvent.L.Lock()
	ret = recorderAuxMockPtrSessionhandleEvent
	condRecorderAuxMockPtrSessionhandleEvent.L.Unlock()
	return
}

// (recvs *Session)handleEvent - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) handleEvent(argframer *framer) {
	FuncAuxMockPtrSessionhandleEvent, ok := apomock.GetRegisteredFunc("gocql.Session.handleEvent")
	if ok {
		FuncAuxMockPtrSessionhandleEvent.(func(recvs *Session, argframer *framer))(recvs, argframer)
	} else {
		panic("FuncAuxMockPtrSessionhandleEvent ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionhandleEvent()
	return
}

//
// Mock: newEventDeouncer(argname string, argeventHandler func([]frame))(reta *eventDeouncer)
//

type MockArgsTypenewEventDeouncer struct {
	ApomockCallNumber int
	Argname           string
	ArgeventHandler   func([]frame)
}

var LastMockArgsnewEventDeouncer MockArgsTypenewEventDeouncer

// AuxMocknewEventDeouncer(argname string, argeventHandler func([]frame))(reta *eventDeouncer) - Generated mock function
func AuxMocknewEventDeouncer(argname string, argeventHandler func([]frame)) (reta *eventDeouncer) {
	LastMockArgsnewEventDeouncer = MockArgsTypenewEventDeouncer{
		ApomockCallNumber: AuxMockGetRecorderAuxMocknewEventDeouncer(),
		Argname:           argname,
		ArgeventHandler:   argeventHandler,
	}
	rargs, rerr := apomock.GetNext("gocql.newEventDeouncer")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.newEventDeouncer")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.newEventDeouncer")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*eventDeouncer)
	}
	return
}

// RecorderAuxMocknewEventDeouncer  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocknewEventDeouncer int = 0

var condRecorderAuxMocknewEventDeouncer *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocknewEventDeouncer(i int) {
	condRecorderAuxMocknewEventDeouncer.L.Lock()
	for recorderAuxMocknewEventDeouncer < i {
		condRecorderAuxMocknewEventDeouncer.Wait()
	}
	condRecorderAuxMocknewEventDeouncer.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocknewEventDeouncer() {
	condRecorderAuxMocknewEventDeouncer.L.Lock()
	recorderAuxMocknewEventDeouncer++
	condRecorderAuxMocknewEventDeouncer.L.Unlock()
	condRecorderAuxMocknewEventDeouncer.Broadcast()
}
func AuxMockGetRecorderAuxMocknewEventDeouncer() (ret int) {
	condRecorderAuxMocknewEventDeouncer.L.Lock()
	ret = recorderAuxMocknewEventDeouncer
	condRecorderAuxMocknewEventDeouncer.L.Unlock()
	return
}

// newEventDeouncer - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func newEventDeouncer(argname string, argeventHandler func([]frame)) (reta *eventDeouncer) {
	FuncAuxMocknewEventDeouncer, ok := apomock.GetRegisteredFunc("gocql.newEventDeouncer")
	if ok {
		reta = FuncAuxMocknewEventDeouncer.(func(argname string, argeventHandler func([]frame)) (reta *eventDeouncer))(argname, argeventHandler)
	} else {
		panic("FuncAuxMocknewEventDeouncer ")
	}
	AuxMockIncrementRecorderAuxMocknewEventDeouncer()
	return
}

//
// Mock: (recve *eventDeouncer)debounce(argframe frame)()
//

type MockArgsTypeeventDeouncerdebounce struct {
	ApomockCallNumber int
	Argframe          frame
}

var LastMockArgseventDeouncerdebounce MockArgsTypeeventDeouncerdebounce

// (recve *eventDeouncer)AuxMockdebounce(argframe frame)() - Generated mock function
func (recve *eventDeouncer) AuxMockdebounce(argframe frame) {
	LastMockArgseventDeouncerdebounce = MockArgsTypeeventDeouncerdebounce{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtreventDeouncerdebounce(),
		Argframe:          argframe,
	}
	return
}

// RecorderAuxMockPtreventDeouncerdebounce  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtreventDeouncerdebounce int = 0

var condRecorderAuxMockPtreventDeouncerdebounce *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtreventDeouncerdebounce(i int) {
	condRecorderAuxMockPtreventDeouncerdebounce.L.Lock()
	for recorderAuxMockPtreventDeouncerdebounce < i {
		condRecorderAuxMockPtreventDeouncerdebounce.Wait()
	}
	condRecorderAuxMockPtreventDeouncerdebounce.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtreventDeouncerdebounce() {
	condRecorderAuxMockPtreventDeouncerdebounce.L.Lock()
	recorderAuxMockPtreventDeouncerdebounce++
	condRecorderAuxMockPtreventDeouncerdebounce.L.Unlock()
	condRecorderAuxMockPtreventDeouncerdebounce.Broadcast()
}
func AuxMockGetRecorderAuxMockPtreventDeouncerdebounce() (ret int) {
	condRecorderAuxMockPtreventDeouncerdebounce.L.Lock()
	ret = recorderAuxMockPtreventDeouncerdebounce
	condRecorderAuxMockPtreventDeouncerdebounce.L.Unlock()
	return
}

// (recve *eventDeouncer)debounce - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve *eventDeouncer) debounce(argframe frame) {
	FuncAuxMockPtreventDeouncerdebounce, ok := apomock.GetRegisteredFunc("gocql.eventDeouncer.debounce")
	if ok {
		FuncAuxMockPtreventDeouncerdebounce.(func(recve *eventDeouncer, argframe frame))(recve, argframe)
	} else {
		panic("FuncAuxMockPtreventDeouncerdebounce ")
	}
	AuxMockIncrementRecorderAuxMockPtreventDeouncerdebounce()
	return
}

//
// Mock: (recvs *Session)handleNewNode(arghost net.IP, argport int, argwaitForBinary bool)()
//

type MockArgsTypeSessionhandleNewNode struct {
	ApomockCallNumber int
	Arghost           net.IP
	Argport           int
	ArgwaitForBinary  bool
}

var LastMockArgsSessionhandleNewNode MockArgsTypeSessionhandleNewNode

// (recvs *Session)AuxMockhandleNewNode(arghost net.IP, argport int, argwaitForBinary bool)() - Generated mock function
func (recvs *Session) AuxMockhandleNewNode(arghost net.IP, argport int, argwaitForBinary bool) {
	LastMockArgsSessionhandleNewNode = MockArgsTypeSessionhandleNewNode{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionhandleNewNode(),
		Arghost:           arghost,
		Argport:           argport,
		ArgwaitForBinary:  argwaitForBinary,
	}
	return
}

// RecorderAuxMockPtrSessionhandleNewNode  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionhandleNewNode int = 0

var condRecorderAuxMockPtrSessionhandleNewNode *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionhandleNewNode(i int) {
	condRecorderAuxMockPtrSessionhandleNewNode.L.Lock()
	for recorderAuxMockPtrSessionhandleNewNode < i {
		condRecorderAuxMockPtrSessionhandleNewNode.Wait()
	}
	condRecorderAuxMockPtrSessionhandleNewNode.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionhandleNewNode() {
	condRecorderAuxMockPtrSessionhandleNewNode.L.Lock()
	recorderAuxMockPtrSessionhandleNewNode++
	condRecorderAuxMockPtrSessionhandleNewNode.L.Unlock()
	condRecorderAuxMockPtrSessionhandleNewNode.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionhandleNewNode() (ret int) {
	condRecorderAuxMockPtrSessionhandleNewNode.L.Lock()
	ret = recorderAuxMockPtrSessionhandleNewNode
	condRecorderAuxMockPtrSessionhandleNewNode.L.Unlock()
	return
}

// (recvs *Session)handleNewNode - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) handleNewNode(arghost net.IP, argport int, argwaitForBinary bool) {
	FuncAuxMockPtrSessionhandleNewNode, ok := apomock.GetRegisteredFunc("gocql.Session.handleNewNode")
	if ok {
		FuncAuxMockPtrSessionhandleNewNode.(func(recvs *Session, arghost net.IP, argport int, argwaitForBinary bool))(recvs, arghost, argport, argwaitForBinary)
	} else {
		panic("FuncAuxMockPtrSessionhandleNewNode ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionhandleNewNode()
	return
}

//
// Mock: (recvs *Session)handleRemovedNode(argip net.IP, argport int)()
//

type MockArgsTypeSessionhandleRemovedNode struct {
	ApomockCallNumber int
	Argip             net.IP
	Argport           int
}

var LastMockArgsSessionhandleRemovedNode MockArgsTypeSessionhandleRemovedNode

// (recvs *Session)AuxMockhandleRemovedNode(argip net.IP, argport int)() - Generated mock function
func (recvs *Session) AuxMockhandleRemovedNode(argip net.IP, argport int) {
	LastMockArgsSessionhandleRemovedNode = MockArgsTypeSessionhandleRemovedNode{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrSessionhandleRemovedNode(),
		Argip:             argip,
		Argport:           argport,
	}
	return
}

// RecorderAuxMockPtrSessionhandleRemovedNode  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrSessionhandleRemovedNode int = 0

var condRecorderAuxMockPtrSessionhandleRemovedNode *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrSessionhandleRemovedNode(i int) {
	condRecorderAuxMockPtrSessionhandleRemovedNode.L.Lock()
	for recorderAuxMockPtrSessionhandleRemovedNode < i {
		condRecorderAuxMockPtrSessionhandleRemovedNode.Wait()
	}
	condRecorderAuxMockPtrSessionhandleRemovedNode.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrSessionhandleRemovedNode() {
	condRecorderAuxMockPtrSessionhandleRemovedNode.L.Lock()
	recorderAuxMockPtrSessionhandleRemovedNode++
	condRecorderAuxMockPtrSessionhandleRemovedNode.L.Unlock()
	condRecorderAuxMockPtrSessionhandleRemovedNode.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrSessionhandleRemovedNode() (ret int) {
	condRecorderAuxMockPtrSessionhandleRemovedNode.L.Lock()
	ret = recorderAuxMockPtrSessionhandleRemovedNode
	condRecorderAuxMockPtrSessionhandleRemovedNode.L.Unlock()
	return
}

// (recvs *Session)handleRemovedNode - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *Session) handleRemovedNode(argip net.IP, argport int) {
	FuncAuxMockPtrSessionhandleRemovedNode, ok := apomock.GetRegisteredFunc("gocql.Session.handleRemovedNode")
	if ok {
		FuncAuxMockPtrSessionhandleRemovedNode.(func(recvs *Session, argip net.IP, argport int))(recvs, argip, argport)
	} else {
		panic("FuncAuxMockPtrSessionhandleRemovedNode ")
	}
	AuxMockIncrementRecorderAuxMockPtrSessionhandleRemovedNode()
	return
}
