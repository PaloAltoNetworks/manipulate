// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {

	apomock.RegisterFunc("gocql", "gocql.HostFilterFunc.Accept", (HostFilterFunc).AuxMockAccept)
	apomock.RegisterFunc("gocql", "gocql.AcceptAllFilter", AuxMockAcceptAllFilter)
	apomock.RegisterFunc("gocql", "gocql.DenyAllFilter", AuxMockDenyAllFilter)
	apomock.RegisterFunc("gocql", "gocql.DataCentreHostFilter", AuxMockDataCentreHostFilter)
	apomock.RegisterFunc("gocql", "gocql.WhiteListHostFilter", AuxMockWhiteListHostFilter)
}

const ()

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//
type HostFilter interface {
	Accept(host *HostInfo) bool
}

type HostFilterFunc func(host *HostInfo) bool

//
// Mock: (recvfn HostFilterFunc)Accept(arghost *HostInfo)(reta bool)
//

type MockArgsTypeHostFilterFuncAccept struct {
	ApomockCallNumber int
	Arghost           *HostInfo
}

var LastMockArgsHostFilterFuncAccept MockArgsTypeHostFilterFuncAccept

// (recvfn HostFilterFunc)AuxMockAccept(arghost *HostInfo)(reta bool) - Generated mock function
func (recvfn HostFilterFunc) AuxMockAccept(arghost *HostInfo) (reta bool) {
	LastMockArgsHostFilterFuncAccept = MockArgsTypeHostFilterFuncAccept{
		ApomockCallNumber: AuxMockGetRecorderAuxMockHostFilterFuncAccept(),
		Arghost:           arghost,
	}
	rargs, rerr := apomock.GetNext("gocql.HostFilterFunc.Accept")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.HostFilterFunc.Accept")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.HostFilterFunc.Accept")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockHostFilterFuncAccept  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockHostFilterFuncAccept int = 0

var condRecorderAuxMockHostFilterFuncAccept *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockHostFilterFuncAccept(i int) {
	condRecorderAuxMockHostFilterFuncAccept.L.Lock()
	for recorderAuxMockHostFilterFuncAccept < i {
		condRecorderAuxMockHostFilterFuncAccept.Wait()
	}
	condRecorderAuxMockHostFilterFuncAccept.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockHostFilterFuncAccept() {
	condRecorderAuxMockHostFilterFuncAccept.L.Lock()
	recorderAuxMockHostFilterFuncAccept++
	condRecorderAuxMockHostFilterFuncAccept.L.Unlock()
	condRecorderAuxMockHostFilterFuncAccept.Broadcast()
}
func AuxMockGetRecorderAuxMockHostFilterFuncAccept() (ret int) {
	condRecorderAuxMockHostFilterFuncAccept.L.Lock()
	ret = recorderAuxMockHostFilterFuncAccept
	condRecorderAuxMockHostFilterFuncAccept.L.Unlock()
	return
}

// (recvfn HostFilterFunc)Accept - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvfn HostFilterFunc) Accept(arghost *HostInfo) (reta bool) {
	FuncAuxMockHostFilterFuncAccept, ok := apomock.GetRegisteredFunc("gocql.HostFilterFunc.Accept")
	if ok {
		reta = FuncAuxMockHostFilterFuncAccept.(func(recvfn HostFilterFunc, arghost *HostInfo) (reta bool))(recvfn, arghost)
	} else {
		panic("FuncAuxMockHostFilterFuncAccept ")
	}
	AuxMockIncrementRecorderAuxMockHostFilterFuncAccept()
	return
}

//
// Mock: AcceptAllFilter()(reta HostFilter)
//

type MockArgsTypeAcceptAllFilter struct {
	ApomockCallNumber int
}

var LastMockArgsAcceptAllFilter MockArgsTypeAcceptAllFilter

// AuxMockAcceptAllFilter()(reta HostFilter) - Generated mock function
func AuxMockAcceptAllFilter() (reta HostFilter) {
	rargs, rerr := apomock.GetNext("gocql.AcceptAllFilter")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.AcceptAllFilter")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.AcceptAllFilter")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(HostFilter)
	}
	return
}

// RecorderAuxMockAcceptAllFilter  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockAcceptAllFilter int = 0

var condRecorderAuxMockAcceptAllFilter *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockAcceptAllFilter(i int) {
	condRecorderAuxMockAcceptAllFilter.L.Lock()
	for recorderAuxMockAcceptAllFilter < i {
		condRecorderAuxMockAcceptAllFilter.Wait()
	}
	condRecorderAuxMockAcceptAllFilter.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockAcceptAllFilter() {
	condRecorderAuxMockAcceptAllFilter.L.Lock()
	recorderAuxMockAcceptAllFilter++
	condRecorderAuxMockAcceptAllFilter.L.Unlock()
	condRecorderAuxMockAcceptAllFilter.Broadcast()
}
func AuxMockGetRecorderAuxMockAcceptAllFilter() (ret int) {
	condRecorderAuxMockAcceptAllFilter.L.Lock()
	ret = recorderAuxMockAcceptAllFilter
	condRecorderAuxMockAcceptAllFilter.L.Unlock()
	return
}

// AcceptAllFilter - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func AcceptAllFilter() (reta HostFilter) {
	FuncAuxMockAcceptAllFilter, ok := apomock.GetRegisteredFunc("gocql.AcceptAllFilter")
	if ok {
		reta = FuncAuxMockAcceptAllFilter.(func() (reta HostFilter))()
	} else {
		panic("FuncAuxMockAcceptAllFilter ")
	}
	AuxMockIncrementRecorderAuxMockAcceptAllFilter()
	return
}

//
// Mock: DenyAllFilter()(reta HostFilter)
//

type MockArgsTypeDenyAllFilter struct {
	ApomockCallNumber int
}

var LastMockArgsDenyAllFilter MockArgsTypeDenyAllFilter

// AuxMockDenyAllFilter()(reta HostFilter) - Generated mock function
func AuxMockDenyAllFilter() (reta HostFilter) {
	rargs, rerr := apomock.GetNext("gocql.DenyAllFilter")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.DenyAllFilter")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.DenyAllFilter")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(HostFilter)
	}
	return
}

// RecorderAuxMockDenyAllFilter  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockDenyAllFilter int = 0

var condRecorderAuxMockDenyAllFilter *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockDenyAllFilter(i int) {
	condRecorderAuxMockDenyAllFilter.L.Lock()
	for recorderAuxMockDenyAllFilter < i {
		condRecorderAuxMockDenyAllFilter.Wait()
	}
	condRecorderAuxMockDenyAllFilter.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockDenyAllFilter() {
	condRecorderAuxMockDenyAllFilter.L.Lock()
	recorderAuxMockDenyAllFilter++
	condRecorderAuxMockDenyAllFilter.L.Unlock()
	condRecorderAuxMockDenyAllFilter.Broadcast()
}
func AuxMockGetRecorderAuxMockDenyAllFilter() (ret int) {
	condRecorderAuxMockDenyAllFilter.L.Lock()
	ret = recorderAuxMockDenyAllFilter
	condRecorderAuxMockDenyAllFilter.L.Unlock()
	return
}

// DenyAllFilter - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func DenyAllFilter() (reta HostFilter) {
	FuncAuxMockDenyAllFilter, ok := apomock.GetRegisteredFunc("gocql.DenyAllFilter")
	if ok {
		reta = FuncAuxMockDenyAllFilter.(func() (reta HostFilter))()
	} else {
		panic("FuncAuxMockDenyAllFilter ")
	}
	AuxMockIncrementRecorderAuxMockDenyAllFilter()
	return
}

//
// Mock: DataCentreHostFilter(argdataCentre string)(reta HostFilter)
//

type MockArgsTypeDataCentreHostFilter struct {
	ApomockCallNumber int
	ArgdataCentre     string
}

var LastMockArgsDataCentreHostFilter MockArgsTypeDataCentreHostFilter

// AuxMockDataCentreHostFilter(argdataCentre string)(reta HostFilter) - Generated mock function
func AuxMockDataCentreHostFilter(argdataCentre string) (reta HostFilter) {
	LastMockArgsDataCentreHostFilter = MockArgsTypeDataCentreHostFilter{
		ApomockCallNumber: AuxMockGetRecorderAuxMockDataCentreHostFilter(),
		ArgdataCentre:     argdataCentre,
	}
	rargs, rerr := apomock.GetNext("gocql.DataCentreHostFilter")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.DataCentreHostFilter")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.DataCentreHostFilter")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(HostFilter)
	}
	return
}

// RecorderAuxMockDataCentreHostFilter  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockDataCentreHostFilter int = 0

var condRecorderAuxMockDataCentreHostFilter *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockDataCentreHostFilter(i int) {
	condRecorderAuxMockDataCentreHostFilter.L.Lock()
	for recorderAuxMockDataCentreHostFilter < i {
		condRecorderAuxMockDataCentreHostFilter.Wait()
	}
	condRecorderAuxMockDataCentreHostFilter.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockDataCentreHostFilter() {
	condRecorderAuxMockDataCentreHostFilter.L.Lock()
	recorderAuxMockDataCentreHostFilter++
	condRecorderAuxMockDataCentreHostFilter.L.Unlock()
	condRecorderAuxMockDataCentreHostFilter.Broadcast()
}
func AuxMockGetRecorderAuxMockDataCentreHostFilter() (ret int) {
	condRecorderAuxMockDataCentreHostFilter.L.Lock()
	ret = recorderAuxMockDataCentreHostFilter
	condRecorderAuxMockDataCentreHostFilter.L.Unlock()
	return
}

// DataCentreHostFilter - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func DataCentreHostFilter(argdataCentre string) (reta HostFilter) {
	FuncAuxMockDataCentreHostFilter, ok := apomock.GetRegisteredFunc("gocql.DataCentreHostFilter")
	if ok {
		reta = FuncAuxMockDataCentreHostFilter.(func(argdataCentre string) (reta HostFilter))(argdataCentre)
	} else {
		panic("FuncAuxMockDataCentreHostFilter ")
	}
	AuxMockIncrementRecorderAuxMockDataCentreHostFilter()
	return
}

//
// Mock: WhiteListHostFilter(hosts ...string)(reta HostFilter)
//

type MockArgsTypeWhiteListHostFilter struct {
	ApomockCallNumber int
	Hosts             []string
}

var LastMockArgsWhiteListHostFilter MockArgsTypeWhiteListHostFilter

// AuxMockWhiteListHostFilter(hosts ...string)(reta HostFilter) - Generated mock function
func AuxMockWhiteListHostFilter(hosts ...string) (reta HostFilter) {
	LastMockArgsWhiteListHostFilter = MockArgsTypeWhiteListHostFilter{
		ApomockCallNumber: AuxMockGetRecorderAuxMockWhiteListHostFilter(),
		Hosts:             hosts,
	}
	rargs, rerr := apomock.GetNext("gocql.WhiteListHostFilter")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.WhiteListHostFilter")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.WhiteListHostFilter")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(HostFilter)
	}
	return
}

// RecorderAuxMockWhiteListHostFilter  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockWhiteListHostFilter int = 0

var condRecorderAuxMockWhiteListHostFilter *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockWhiteListHostFilter(i int) {
	condRecorderAuxMockWhiteListHostFilter.L.Lock()
	for recorderAuxMockWhiteListHostFilter < i {
		condRecorderAuxMockWhiteListHostFilter.Wait()
	}
	condRecorderAuxMockWhiteListHostFilter.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockWhiteListHostFilter() {
	condRecorderAuxMockWhiteListHostFilter.L.Lock()
	recorderAuxMockWhiteListHostFilter++
	condRecorderAuxMockWhiteListHostFilter.L.Unlock()
	condRecorderAuxMockWhiteListHostFilter.Broadcast()
}
func AuxMockGetRecorderAuxMockWhiteListHostFilter() (ret int) {
	condRecorderAuxMockWhiteListHostFilter.L.Lock()
	ret = recorderAuxMockWhiteListHostFilter
	condRecorderAuxMockWhiteListHostFilter.L.Unlock()
	return
}

// WhiteListHostFilter - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func WhiteListHostFilter(hosts ...string) (reta HostFilter) {
	FuncAuxMockWhiteListHostFilter, ok := apomock.GetRegisteredFunc("gocql.WhiteListHostFilter")
	if ok {
		reta = FuncAuxMockWhiteListHostFilter.(func(hosts ...string) (reta HostFilter))(hosts...)
	} else {
		panic("FuncAuxMockWhiteListHostFilter ")
	}
	AuxMockIncrementRecorderAuxMockWhiteListHostFilter()
	return
}
