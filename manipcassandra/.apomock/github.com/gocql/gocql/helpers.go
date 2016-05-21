// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "reflect"
import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {

	apomock.RegisterFunc("gocql", "gocql.goType", AuxMockgoType)
	apomock.RegisterFunc("gocql", "gocql.getCassandraType", AuxMockgetCassandraType)
	apomock.RegisterFunc("gocql", "gocql.RowData.rowMap", (*RowData).AuxMockrowMap)
	apomock.RegisterFunc("gocql", "gocql.copyBytes", AuxMockcopyBytes)
	apomock.RegisterFunc("gocql", "gocql.Iter.RowData", (*Iter).AuxMockRowData)
	apomock.RegisterFunc("gocql", "gocql.Iter.SliceMap", (*Iter).AuxMockSliceMap)
	apomock.RegisterFunc("gocql", "gocql.Iter.MapScan", (*Iter).AuxMockMapScan)
	apomock.RegisterFunc("gocql", "gocql.dereference", AuxMockdereference)
	apomock.RegisterFunc("gocql", "gocql.getApacheCassandraType", AuxMockgetApacheCassandraType)
	apomock.RegisterFunc("gocql", "gocql.typeCanBeNull", AuxMocktypeCanBeNull)
	apomock.RegisterFunc("gocql", "gocql.TupleColumnName", AuxMockTupleColumnName)
}

const (
	CustomUUIDTypeUUID   CustomUUIDType = 0x0000
	CustomUUIDTypeString CustomUUIDType = 0x0001
)

const ()

var (
	DefaultUUIDType = CustomUUIDTypeUUID
)

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//
type CustomUUIDType int

type RowData struct {
	Columns []string
	Values  []interface{}
}

//
// Mock: goType(argt TypeInfo)(reta reflect.Type)
//

type MockArgsTypegoType struct {
	ApomockCallNumber int
	Argt              TypeInfo
}

var LastMockArgsgoType MockArgsTypegoType

// AuxMockgoType(argt TypeInfo)(reta reflect.Type) - Generated mock function
func AuxMockgoType(argt TypeInfo) (reta reflect.Type) {
	LastMockArgsgoType = MockArgsTypegoType{
		ApomockCallNumber: AuxMockGetRecorderAuxMockgoType(),
		Argt:              argt,
	}
	rargs, rerr := apomock.GetNext("gocql.goType")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.goType")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.goType")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(reflect.Type)
	}
	return
}

// RecorderAuxMockgoType  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockgoType int = 0

var condRecorderAuxMockgoType *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockgoType(i int) {
	condRecorderAuxMockgoType.L.Lock()
	for recorderAuxMockgoType < i {
		condRecorderAuxMockgoType.Wait()
	}
	condRecorderAuxMockgoType.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockgoType() {
	condRecorderAuxMockgoType.L.Lock()
	recorderAuxMockgoType++
	condRecorderAuxMockgoType.L.Unlock()
	condRecorderAuxMockgoType.Broadcast()
}
func AuxMockGetRecorderAuxMockgoType() (ret int) {
	condRecorderAuxMockgoType.L.Lock()
	ret = recorderAuxMockgoType
	condRecorderAuxMockgoType.L.Unlock()
	return
}

// goType - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func goType(argt TypeInfo) (reta reflect.Type) {
	FuncAuxMockgoType, ok := apomock.GetRegisteredFunc("gocql.goType")
	if ok {
		reta = FuncAuxMockgoType.(func(argt TypeInfo) (reta reflect.Type))(argt)
	} else {
		panic("FuncAuxMockgoType ")
	}
	AuxMockIncrementRecorderAuxMockgoType()
	return
}

//
// Mock: getCassandraType(argname string)(reta Type)
//

type MockArgsTypegetCassandraType struct {
	ApomockCallNumber int
	Argname           string
}

var LastMockArgsgetCassandraType MockArgsTypegetCassandraType

// AuxMockgetCassandraType(argname string)(reta Type) - Generated mock function
func AuxMockgetCassandraType(argname string) (reta Type) {
	LastMockArgsgetCassandraType = MockArgsTypegetCassandraType{
		ApomockCallNumber: AuxMockGetRecorderAuxMockgetCassandraType(),
		Argname:           argname,
	}
	rargs, rerr := apomock.GetNext("gocql.getCassandraType")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.getCassandraType")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.getCassandraType")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(Type)
	}
	return
}

// RecorderAuxMockgetCassandraType  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockgetCassandraType int = 0

var condRecorderAuxMockgetCassandraType *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockgetCassandraType(i int) {
	condRecorderAuxMockgetCassandraType.L.Lock()
	for recorderAuxMockgetCassandraType < i {
		condRecorderAuxMockgetCassandraType.Wait()
	}
	condRecorderAuxMockgetCassandraType.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockgetCassandraType() {
	condRecorderAuxMockgetCassandraType.L.Lock()
	recorderAuxMockgetCassandraType++
	condRecorderAuxMockgetCassandraType.L.Unlock()
	condRecorderAuxMockgetCassandraType.Broadcast()
}
func AuxMockGetRecorderAuxMockgetCassandraType() (ret int) {
	condRecorderAuxMockgetCassandraType.L.Lock()
	ret = recorderAuxMockgetCassandraType
	condRecorderAuxMockgetCassandraType.L.Unlock()
	return
}

// getCassandraType - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func getCassandraType(argname string) (reta Type) {
	FuncAuxMockgetCassandraType, ok := apomock.GetRegisteredFunc("gocql.getCassandraType")
	if ok {
		reta = FuncAuxMockgetCassandraType.(func(argname string) (reta Type))(argname)
	} else {
		panic("FuncAuxMockgetCassandraType ")
	}
	AuxMockIncrementRecorderAuxMockgetCassandraType()
	return
}

//
// Mock: (recvr *RowData)rowMap(argm map[string]interface{})()
//

type MockArgsTypeRowDatarowMap struct {
	ApomockCallNumber int
	Argm              map[string]interface{}
}

var LastMockArgsRowDatarowMap MockArgsTypeRowDatarowMap

// (recvr *RowData)AuxMockrowMap(argm map[string]interface{})() - Generated mock function
func (recvr *RowData) AuxMockrowMap(argm map[string]interface{}) {
	LastMockArgsRowDatarowMap = MockArgsTypeRowDatarowMap{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrRowDatarowMap(),
		Argm:              argm,
	}
	return
}

// RecorderAuxMockPtrRowDatarowMap  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrRowDatarowMap int = 0

var condRecorderAuxMockPtrRowDatarowMap *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrRowDatarowMap(i int) {
	condRecorderAuxMockPtrRowDatarowMap.L.Lock()
	for recorderAuxMockPtrRowDatarowMap < i {
		condRecorderAuxMockPtrRowDatarowMap.Wait()
	}
	condRecorderAuxMockPtrRowDatarowMap.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrRowDatarowMap() {
	condRecorderAuxMockPtrRowDatarowMap.L.Lock()
	recorderAuxMockPtrRowDatarowMap++
	condRecorderAuxMockPtrRowDatarowMap.L.Unlock()
	condRecorderAuxMockPtrRowDatarowMap.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrRowDatarowMap() (ret int) {
	condRecorderAuxMockPtrRowDatarowMap.L.Lock()
	ret = recorderAuxMockPtrRowDatarowMap
	condRecorderAuxMockPtrRowDatarowMap.L.Unlock()
	return
}

// (recvr *RowData)rowMap - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *RowData) rowMap(argm map[string]interface{}) {
	FuncAuxMockPtrRowDatarowMap, ok := apomock.GetRegisteredFunc("gocql.RowData.rowMap")
	if ok {
		FuncAuxMockPtrRowDatarowMap.(func(recvr *RowData, argm map[string]interface{}))(recvr, argm)
	} else {
		panic("FuncAuxMockPtrRowDatarowMap ")
	}
	AuxMockIncrementRecorderAuxMockPtrRowDatarowMap()
	return
}

//
// Mock: copyBytes(argp []byte)(reta []byte)
//

type MockArgsTypecopyBytes struct {
	ApomockCallNumber int
	Argp              []byte
}

var LastMockArgscopyBytes MockArgsTypecopyBytes

// AuxMockcopyBytes(argp []byte)(reta []byte) - Generated mock function
func AuxMockcopyBytes(argp []byte) (reta []byte) {
	LastMockArgscopyBytes = MockArgsTypecopyBytes{
		ApomockCallNumber: AuxMockGetRecorderAuxMockcopyBytes(),
		Argp:              argp,
	}
	rargs, rerr := apomock.GetNext("gocql.copyBytes")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.copyBytes")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.copyBytes")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockcopyBytes  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockcopyBytes int = 0

var condRecorderAuxMockcopyBytes *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockcopyBytes(i int) {
	condRecorderAuxMockcopyBytes.L.Lock()
	for recorderAuxMockcopyBytes < i {
		condRecorderAuxMockcopyBytes.Wait()
	}
	condRecorderAuxMockcopyBytes.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockcopyBytes() {
	condRecorderAuxMockcopyBytes.L.Lock()
	recorderAuxMockcopyBytes++
	condRecorderAuxMockcopyBytes.L.Unlock()
	condRecorderAuxMockcopyBytes.Broadcast()
}
func AuxMockGetRecorderAuxMockcopyBytes() (ret int) {
	condRecorderAuxMockcopyBytes.L.Lock()
	ret = recorderAuxMockcopyBytes
	condRecorderAuxMockcopyBytes.L.Unlock()
	return
}

// copyBytes - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func copyBytes(argp []byte) (reta []byte) {
	FuncAuxMockcopyBytes, ok := apomock.GetRegisteredFunc("gocql.copyBytes")
	if ok {
		reta = FuncAuxMockcopyBytes.(func(argp []byte) (reta []byte))(argp)
	} else {
		panic("FuncAuxMockcopyBytes ")
	}
	AuxMockIncrementRecorderAuxMockcopyBytes()
	return
}

//
// Mock: (recviter *Iter)RowData()(reta RowData, retb error)
//

type MockArgsTypeIterRowData struct {
	ApomockCallNumber int
}

var LastMockArgsIterRowData MockArgsTypeIterRowData

// (recviter *Iter)AuxMockRowData()(reta RowData, retb error) - Generated mock function
func (recviter *Iter) AuxMockRowData() (reta RowData, retb error) {
	rargs, rerr := apomock.GetNext("gocql.Iter.RowData")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.RowData")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Iter.RowData")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(RowData)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrIterRowData  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterRowData int = 0

var condRecorderAuxMockPtrIterRowData *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterRowData(i int) {
	condRecorderAuxMockPtrIterRowData.L.Lock()
	for recorderAuxMockPtrIterRowData < i {
		condRecorderAuxMockPtrIterRowData.Wait()
	}
	condRecorderAuxMockPtrIterRowData.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterRowData() {
	condRecorderAuxMockPtrIterRowData.L.Lock()
	recorderAuxMockPtrIterRowData++
	condRecorderAuxMockPtrIterRowData.L.Unlock()
	condRecorderAuxMockPtrIterRowData.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterRowData() (ret int) {
	condRecorderAuxMockPtrIterRowData.L.Lock()
	ret = recorderAuxMockPtrIterRowData
	condRecorderAuxMockPtrIterRowData.L.Unlock()
	return
}

// (recviter *Iter)RowData - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) RowData() (reta RowData, retb error) {
	FuncAuxMockPtrIterRowData, ok := apomock.GetRegisteredFunc("gocql.Iter.RowData")
	if ok {
		reta, retb = FuncAuxMockPtrIterRowData.(func(recviter *Iter) (reta RowData, retb error))(recviter)
	} else {
		panic("FuncAuxMockPtrIterRowData ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterRowData()
	return
}

//
// Mock: (recviter *Iter)SliceMap()(reta []map[string]interface{}, retb error)
//

type MockArgsTypeIterSliceMap struct {
	ApomockCallNumber int
}

var LastMockArgsIterSliceMap MockArgsTypeIterSliceMap

// (recviter *Iter)AuxMockSliceMap()(reta []map[string]interface{}, retb error) - Generated mock function
func (recviter *Iter) AuxMockSliceMap() (reta []map[string]interface{}, retb error) {
	rargs, rerr := apomock.GetNext("gocql.Iter.SliceMap")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.SliceMap")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Iter.SliceMap")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]map[string]interface{})
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrIterSliceMap  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterSliceMap int = 0

var condRecorderAuxMockPtrIterSliceMap *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterSliceMap(i int) {
	condRecorderAuxMockPtrIterSliceMap.L.Lock()
	for recorderAuxMockPtrIterSliceMap < i {
		condRecorderAuxMockPtrIterSliceMap.Wait()
	}
	condRecorderAuxMockPtrIterSliceMap.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterSliceMap() {
	condRecorderAuxMockPtrIterSliceMap.L.Lock()
	recorderAuxMockPtrIterSliceMap++
	condRecorderAuxMockPtrIterSliceMap.L.Unlock()
	condRecorderAuxMockPtrIterSliceMap.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterSliceMap() (ret int) {
	condRecorderAuxMockPtrIterSliceMap.L.Lock()
	ret = recorderAuxMockPtrIterSliceMap
	condRecorderAuxMockPtrIterSliceMap.L.Unlock()
	return
}

// (recviter *Iter)SliceMap - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) SliceMap() (reta []map[string]interface{}, retb error) {
	FuncAuxMockPtrIterSliceMap, ok := apomock.GetRegisteredFunc("gocql.Iter.SliceMap")
	if ok {
		reta, retb = FuncAuxMockPtrIterSliceMap.(func(recviter *Iter) (reta []map[string]interface{}, retb error))(recviter)
	} else {
		panic("FuncAuxMockPtrIterSliceMap ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterSliceMap()
	return
}

//
// Mock: (recviter *Iter)MapScan(argm map[string]interface{})(reta bool)
//

type MockArgsTypeIterMapScan struct {
	ApomockCallNumber int
	Argm              map[string]interface{}
}

var LastMockArgsIterMapScan MockArgsTypeIterMapScan

// (recviter *Iter)AuxMockMapScan(argm map[string]interface{})(reta bool) - Generated mock function
func (recviter *Iter) AuxMockMapScan(argm map[string]interface{}) (reta bool) {
	LastMockArgsIterMapScan = MockArgsTypeIterMapScan{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrIterMapScan(),
		Argm:              argm,
	}
	rargs, rerr := apomock.GetNext("gocql.Iter.MapScan")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Iter.MapScan")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Iter.MapScan")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrIterMapScan  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIterMapScan int = 0

var condRecorderAuxMockPtrIterMapScan *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIterMapScan(i int) {
	condRecorderAuxMockPtrIterMapScan.L.Lock()
	for recorderAuxMockPtrIterMapScan < i {
		condRecorderAuxMockPtrIterMapScan.Wait()
	}
	condRecorderAuxMockPtrIterMapScan.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIterMapScan() {
	condRecorderAuxMockPtrIterMapScan.L.Lock()
	recorderAuxMockPtrIterMapScan++
	condRecorderAuxMockPtrIterMapScan.L.Unlock()
	condRecorderAuxMockPtrIterMapScan.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIterMapScan() (ret int) {
	condRecorderAuxMockPtrIterMapScan.L.Lock()
	ret = recorderAuxMockPtrIterMapScan
	condRecorderAuxMockPtrIterMapScan.L.Unlock()
	return
}

// (recviter *Iter)MapScan - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recviter *Iter) MapScan(argm map[string]interface{}) (reta bool) {
	FuncAuxMockPtrIterMapScan, ok := apomock.GetRegisteredFunc("gocql.Iter.MapScan")
	if ok {
		reta = FuncAuxMockPtrIterMapScan.(func(recviter *Iter, argm map[string]interface{}) (reta bool))(recviter, argm)
	} else {
		panic("FuncAuxMockPtrIterMapScan ")
	}
	AuxMockIncrementRecorderAuxMockPtrIterMapScan()
	return
}

//
// Mock: dereference(argi interface{})(reta interface{})
//

type MockArgsTypedereference struct {
	ApomockCallNumber int
	Argi              interface{}
}

var LastMockArgsdereference MockArgsTypedereference

// AuxMockdereference(argi interface{})(reta interface{}) - Generated mock function
func AuxMockdereference(argi interface{}) (reta interface{}) {
	LastMockArgsdereference = MockArgsTypedereference{
		ApomockCallNumber: AuxMockGetRecorderAuxMockdereference(),
		Argi:              argi,
	}
	rargs, rerr := apomock.GetNext("gocql.dereference")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.dereference")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.dereference")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(interface{})
	}
	return
}

// RecorderAuxMockdereference  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockdereference int = 0

var condRecorderAuxMockdereference *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockdereference(i int) {
	condRecorderAuxMockdereference.L.Lock()
	for recorderAuxMockdereference < i {
		condRecorderAuxMockdereference.Wait()
	}
	condRecorderAuxMockdereference.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockdereference() {
	condRecorderAuxMockdereference.L.Lock()
	recorderAuxMockdereference++
	condRecorderAuxMockdereference.L.Unlock()
	condRecorderAuxMockdereference.Broadcast()
}
func AuxMockGetRecorderAuxMockdereference() (ret int) {
	condRecorderAuxMockdereference.L.Lock()
	ret = recorderAuxMockdereference
	condRecorderAuxMockdereference.L.Unlock()
	return
}

// dereference - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func dereference(argi interface{}) (reta interface{}) {
	FuncAuxMockdereference, ok := apomock.GetRegisteredFunc("gocql.dereference")
	if ok {
		reta = FuncAuxMockdereference.(func(argi interface{}) (reta interface{}))(argi)
	} else {
		panic("FuncAuxMockdereference ")
	}
	AuxMockIncrementRecorderAuxMockdereference()
	return
}

//
// Mock: getApacheCassandraType(argclass string)(reta Type)
//

type MockArgsTypegetApacheCassandraType struct {
	ApomockCallNumber int
	Argclass          string
}

var LastMockArgsgetApacheCassandraType MockArgsTypegetApacheCassandraType

// AuxMockgetApacheCassandraType(argclass string)(reta Type) - Generated mock function
func AuxMockgetApacheCassandraType(argclass string) (reta Type) {
	LastMockArgsgetApacheCassandraType = MockArgsTypegetApacheCassandraType{
		ApomockCallNumber: AuxMockGetRecorderAuxMockgetApacheCassandraType(),
		Argclass:          argclass,
	}
	rargs, rerr := apomock.GetNext("gocql.getApacheCassandraType")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.getApacheCassandraType")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.getApacheCassandraType")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(Type)
	}
	return
}

// RecorderAuxMockgetApacheCassandraType  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockgetApacheCassandraType int = 0

var condRecorderAuxMockgetApacheCassandraType *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockgetApacheCassandraType(i int) {
	condRecorderAuxMockgetApacheCassandraType.L.Lock()
	for recorderAuxMockgetApacheCassandraType < i {
		condRecorderAuxMockgetApacheCassandraType.Wait()
	}
	condRecorderAuxMockgetApacheCassandraType.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockgetApacheCassandraType() {
	condRecorderAuxMockgetApacheCassandraType.L.Lock()
	recorderAuxMockgetApacheCassandraType++
	condRecorderAuxMockgetApacheCassandraType.L.Unlock()
	condRecorderAuxMockgetApacheCassandraType.Broadcast()
}
func AuxMockGetRecorderAuxMockgetApacheCassandraType() (ret int) {
	condRecorderAuxMockgetApacheCassandraType.L.Lock()
	ret = recorderAuxMockgetApacheCassandraType
	condRecorderAuxMockgetApacheCassandraType.L.Unlock()
	return
}

// getApacheCassandraType - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func getApacheCassandraType(argclass string) (reta Type) {
	FuncAuxMockgetApacheCassandraType, ok := apomock.GetRegisteredFunc("gocql.getApacheCassandraType")
	if ok {
		reta = FuncAuxMockgetApacheCassandraType.(func(argclass string) (reta Type))(argclass)
	} else {
		panic("FuncAuxMockgetApacheCassandraType ")
	}
	AuxMockIncrementRecorderAuxMockgetApacheCassandraType()
	return
}

//
// Mock: typeCanBeNull(argtyp TypeInfo)(reta bool)
//

type MockArgsTypetypeCanBeNull struct {
	ApomockCallNumber int
	Argtyp            TypeInfo
}

var LastMockArgstypeCanBeNull MockArgsTypetypeCanBeNull

// AuxMocktypeCanBeNull(argtyp TypeInfo)(reta bool) - Generated mock function
func AuxMocktypeCanBeNull(argtyp TypeInfo) (reta bool) {
	LastMockArgstypeCanBeNull = MockArgsTypetypeCanBeNull{
		ApomockCallNumber: AuxMockGetRecorderAuxMocktypeCanBeNull(),
		Argtyp:            argtyp,
	}
	rargs, rerr := apomock.GetNext("gocql.typeCanBeNull")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.typeCanBeNull")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.typeCanBeNull")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMocktypeCanBeNull  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocktypeCanBeNull int = 0

var condRecorderAuxMocktypeCanBeNull *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocktypeCanBeNull(i int) {
	condRecorderAuxMocktypeCanBeNull.L.Lock()
	for recorderAuxMocktypeCanBeNull < i {
		condRecorderAuxMocktypeCanBeNull.Wait()
	}
	condRecorderAuxMocktypeCanBeNull.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocktypeCanBeNull() {
	condRecorderAuxMocktypeCanBeNull.L.Lock()
	recorderAuxMocktypeCanBeNull++
	condRecorderAuxMocktypeCanBeNull.L.Unlock()
	condRecorderAuxMocktypeCanBeNull.Broadcast()
}
func AuxMockGetRecorderAuxMocktypeCanBeNull() (ret int) {
	condRecorderAuxMocktypeCanBeNull.L.Lock()
	ret = recorderAuxMocktypeCanBeNull
	condRecorderAuxMocktypeCanBeNull.L.Unlock()
	return
}

// typeCanBeNull - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func typeCanBeNull(argtyp TypeInfo) (reta bool) {
	FuncAuxMocktypeCanBeNull, ok := apomock.GetRegisteredFunc("gocql.typeCanBeNull")
	if ok {
		reta = FuncAuxMocktypeCanBeNull.(func(argtyp TypeInfo) (reta bool))(argtyp)
	} else {
		panic("FuncAuxMocktypeCanBeNull ")
	}
	AuxMockIncrementRecorderAuxMocktypeCanBeNull()
	return
}

//
// Mock: TupleColumnName(argc string, argn int)(reta string)
//

type MockArgsTypeTupleColumnName struct {
	ApomockCallNumber int
	Argc              string
	Argn              int
}

var LastMockArgsTupleColumnName MockArgsTypeTupleColumnName

// AuxMockTupleColumnName(argc string, argn int)(reta string) - Generated mock function
func AuxMockTupleColumnName(argc string, argn int) (reta string) {
	LastMockArgsTupleColumnName = MockArgsTypeTupleColumnName{
		ApomockCallNumber: AuxMockGetRecorderAuxMockTupleColumnName(),
		Argc:              argc,
		Argn:              argn,
	}
	rargs, rerr := apomock.GetNext("gocql.TupleColumnName")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.TupleColumnName")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.TupleColumnName")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockTupleColumnName  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockTupleColumnName int = 0

var condRecorderAuxMockTupleColumnName *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockTupleColumnName(i int) {
	condRecorderAuxMockTupleColumnName.L.Lock()
	for recorderAuxMockTupleColumnName < i {
		condRecorderAuxMockTupleColumnName.Wait()
	}
	condRecorderAuxMockTupleColumnName.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockTupleColumnName() {
	condRecorderAuxMockTupleColumnName.L.Lock()
	recorderAuxMockTupleColumnName++
	condRecorderAuxMockTupleColumnName.L.Unlock()
	condRecorderAuxMockTupleColumnName.Broadcast()
}
func AuxMockGetRecorderAuxMockTupleColumnName() (ret int) {
	condRecorderAuxMockTupleColumnName.L.Lock()
	ret = recorderAuxMockTupleColumnName
	condRecorderAuxMockTupleColumnName.L.Unlock()
	return
}

// TupleColumnName - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func TupleColumnName(argc string, argn int) (reta string) {
	FuncAuxMockTupleColumnName, ok := apomock.GetRegisteredFunc("gocql.TupleColumnName")
	if ok {
		reta = FuncAuxMockTupleColumnName.(func(argc string, argn int) (reta string))(argc, argn)
	} else {
		panic("FuncAuxMockTupleColumnName ")
	}
	AuxMockIncrementRecorderAuxMockTupleColumnName()
	return
}
