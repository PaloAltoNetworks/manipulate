// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

import "errors"

import "bytes"
import "math/big"

func init() {

	apomock.RegisterFunc("gocql", "gocql.unmarshalBool", AuxMockunmarshalBool)
	apomock.RegisterFunc("gocql", "gocql.marshalTimestamp", AuxMockmarshalTimestamp)
	apomock.RegisterFunc("gocql", "gocql.marshalTuple", AuxMockmarshalTuple)
	apomock.RegisterFunc("gocql", "gocql.unmarshalErrorf", AuxMockunmarshalErrorf)
	apomock.RegisterFunc("gocql", "gocql.unmarshalVarint", AuxMockunmarshalVarint)
	apomock.RegisterFunc("gocql", "gocql.unmarshalMap", AuxMockunmarshalMap)
	apomock.RegisterFunc("gocql", "gocql.unmarshalUUID", AuxMockunmarshalUUID)
	apomock.RegisterFunc("gocql", "gocql.unmarshalUDT", AuxMockunmarshalUDT)
	apomock.RegisterFunc("gocql", "gocql.Type.String", (Type).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.marshalMap", AuxMockmarshalMap)
	apomock.RegisterFunc("gocql", "gocql.unmarshalNullable", AuxMockunmarshalNullable)
	apomock.RegisterFunc("gocql", "gocql.unmarshalVarchar", AuxMockunmarshalVarchar)
	apomock.RegisterFunc("gocql", "gocql.marshalBool", AuxMockmarshalBool)
	apomock.RegisterFunc("gocql", "gocql.marshalDouble", AuxMockmarshalDouble)
	apomock.RegisterFunc("gocql", "gocql.writeCollectionSize", AuxMockwriteCollectionSize)
	apomock.RegisterFunc("gocql", "gocql.marshalVarchar", AuxMockmarshalVarchar)
	apomock.RegisterFunc("gocql", "gocql.bytesToInt64", AuxMockbytesToInt64)
	apomock.RegisterFunc("gocql", "gocql.unmarshalBigInt", AuxMockunmarshalBigInt)
	apomock.RegisterFunc("gocql", "gocql.CollectionType.String", (CollectionType).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.unmarshalTimestamp", AuxMockunmarshalTimestamp)
	apomock.RegisterFunc("gocql", "gocql.NativeType.Version", (NativeType).AuxMockVersion)
	apomock.RegisterFunc("gocql", "gocql.unmarshalIntlike", AuxMockunmarshalIntlike)
	apomock.RegisterFunc("gocql", "gocql.marshalUUID", AuxMockmarshalUUID)
	apomock.RegisterFunc("gocql", "gocql.marshalUDT", AuxMockmarshalUDT)
	apomock.RegisterFunc("gocql", "gocql.MarshalError.Error", (MarshalError).AuxMockError)
	apomock.RegisterFunc("gocql", "gocql.encBool", AuxMockencBool)
	apomock.RegisterFunc("gocql", "gocql.unmarshalDecimal", AuxMockunmarshalDecimal)
	apomock.RegisterFunc("gocql", "gocql.marshalList", AuxMockmarshalList)
	apomock.RegisterFunc("gocql", "gocql.marshalInet", AuxMockmarshalInet)
	apomock.RegisterFunc("gocql", "gocql.bytesToUint64", AuxMockbytesToUint64)
	apomock.RegisterFunc("gocql", "gocql.decBigInt", AuxMockdecBigInt)
	apomock.RegisterFunc("gocql", "gocql.unmarshalDouble", AuxMockunmarshalDouble)
	apomock.RegisterFunc("gocql", "gocql.NativeType.New", (NativeType).AuxMockNew)
	apomock.RegisterFunc("gocql", "gocql.NativeType.Custom", (NativeType).AuxMockCustom)
	apomock.RegisterFunc("gocql", "gocql.decBigInt2C", AuxMockdecBigInt2C)
	apomock.RegisterFunc("gocql", "gocql.readCollectionSize", AuxMockreadCollectionSize)
	apomock.RegisterFunc("gocql", "gocql.TupleTypeInfo.New", (TupleTypeInfo).AuxMockNew)
	apomock.RegisterFunc("gocql", "gocql.UDTTypeInfo.String", (UDTTypeInfo).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.NativeType.String", (NativeType).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.UDTTypeInfo.New", (UDTTypeInfo).AuxMockNew)
	apomock.RegisterFunc("gocql", "gocql.Unmarshal", AuxMockUnmarshal)
	apomock.RegisterFunc("gocql", "gocql.isNullableValue", AuxMockisNullableValue)
	apomock.RegisterFunc("gocql", "gocql.isNullData", AuxMockisNullData)
	apomock.RegisterFunc("gocql", "gocql.encBigInt", AuxMockencBigInt)
	apomock.RegisterFunc("gocql", "gocql.unmarshalTuple", AuxMockunmarshalTuple)
	apomock.RegisterFunc("gocql", "gocql.unmarshalList", AuxMockunmarshalList)
	apomock.RegisterFunc("gocql", "gocql.unmarshalTimeUUID", AuxMockunmarshalTimeUUID)
	apomock.RegisterFunc("gocql", "gocql.marshalInt", AuxMockmarshalInt)
	apomock.RegisterFunc("gocql", "gocql.encInt", AuxMockencInt)
	apomock.RegisterFunc("gocql", "gocql.decBool", AuxMockdecBool)
	apomock.RegisterFunc("gocql", "gocql.unmarshalFloat", AuxMockunmarshalFloat)
	apomock.RegisterFunc("gocql", "gocql.encBigInt2C", AuxMockencBigInt2C)
	apomock.RegisterFunc("gocql", "gocql.marshalBigInt", AuxMockmarshalBigInt)
	apomock.RegisterFunc("gocql", "gocql.unmarshalInet", AuxMockunmarshalInet)
	apomock.RegisterFunc("gocql", "gocql.decInt", AuxMockdecInt)
	apomock.RegisterFunc("gocql", "gocql.marshalVarint", AuxMockmarshalVarint)
	apomock.RegisterFunc("gocql", "gocql.marshalErrorf", AuxMockmarshalErrorf)
	apomock.RegisterFunc("gocql", "gocql.UnmarshalError.Error", (UnmarshalError).AuxMockError)
	apomock.RegisterFunc("gocql", "gocql.Marshal", AuxMockMarshal)
	apomock.RegisterFunc("gocql", "gocql.marshalDecimal", AuxMockmarshalDecimal)
	apomock.RegisterFunc("gocql", "gocql.CollectionType.New", (CollectionType).AuxMockNew)
	apomock.RegisterFunc("gocql", "gocql.unmarshalInt", AuxMockunmarshalInt)
	apomock.RegisterFunc("gocql", "gocql.marshalFloat", AuxMockmarshalFloat)
	apomock.RegisterFunc("gocql", "gocql.NativeType.Type", (NativeType).AuxMockType)
}

const (
	TypeCustom    Type = 0x0000
	TypeAscii     Type = 0x0001
	TypeBigInt    Type = 0x0002
	TypeBlob      Type = 0x0003
	TypeBoolean   Type = 0x0004
	TypeCounter   Type = 0x0005
	TypeDecimal   Type = 0x0006
	TypeDouble    Type = 0x0007
	TypeFloat     Type = 0x0008
	TypeInt       Type = 0x0009
	TypeText      Type = 0x000A
	TypeTimestamp Type = 0x000B
	TypeUUID      Type = 0x000C
	TypeVarchar   Type = 0x000D
	TypeVarint    Type = 0x000E
	TypeTimeUUID  Type = 0x000F
	TypeInet      Type = 0x0010
	TypeList      Type = 0x0020
	TypeMap       Type = 0x0021
	TypeSet       Type = 0x0022
	TypeUDT       Type = 0x0030
	TypeTuple     Type = 0x0031
)

const ()

var (
	bigOne = big.NewInt(1)
)

var (
	ErrorUDTUnavailable = errors.New("UDT are not available on protocols less than 3, please update config")
)

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//
type Marshaler interface {
	MarshalCQL(info TypeInfo) ([]byte, error)
}

type CollectionType struct {
	NativeType
	Key  TypeInfo
	Elem TypeInfo
}

type Type int

type UnmarshalError string

type Unmarshaler interface {
	UnmarshalCQL(info TypeInfo, data []byte) error
}

type UDTMarshaler interface {
	MarshalUDT(name string, info TypeInfo) ([]byte, error)
}

type UDTUnmarshaler interface {
	UnmarshalUDT(name string, info TypeInfo, data []byte) error
}

type TypeInfo interface {
	Type() Type
	Version() byte
	Custom() string
	New() interface{}
}

type NativeType struct {
	proto  byte
	typ    Type
	custom string
}

type TupleTypeInfo struct {
	NativeType
	Elems []TypeInfo
}

type UDTField struct {
	Name string
	Type TypeInfo
}

type UDTTypeInfo struct {
	NativeType
	KeySpace string
	Name     string
	Elements []UDTField
}

type MarshalError string

//
// Mock: unmarshalBool(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalBool struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalBool MockArgsTypeunmarshalBool

// AuxMockunmarshalBool(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalBool(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalBool = MockArgsTypeunmarshalBool{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalBool(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalBool")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalBool")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalBool")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalBool  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalBool int = 0

var condRecorderAuxMockunmarshalBool *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalBool(i int) {
	condRecorderAuxMockunmarshalBool.L.Lock()
	for recorderAuxMockunmarshalBool < i {
		condRecorderAuxMockunmarshalBool.Wait()
	}
	condRecorderAuxMockunmarshalBool.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalBool() {
	condRecorderAuxMockunmarshalBool.L.Lock()
	recorderAuxMockunmarshalBool++
	condRecorderAuxMockunmarshalBool.L.Unlock()
	condRecorderAuxMockunmarshalBool.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalBool() (ret int) {
	condRecorderAuxMockunmarshalBool.L.Lock()
	ret = recorderAuxMockunmarshalBool
	condRecorderAuxMockunmarshalBool.L.Unlock()
	return
}

// unmarshalBool - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalBool(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalBool, ok := apomock.GetRegisteredFunc("gocql.unmarshalBool")
	if ok {
		reta = FuncAuxMockunmarshalBool.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalBool ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalBool()
	return
}

//
// Mock: marshalTimestamp(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalTimestamp struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalTimestamp MockArgsTypemarshalTimestamp

// AuxMockmarshalTimestamp(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalTimestamp(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalTimestamp = MockArgsTypemarshalTimestamp{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalTimestamp(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalTimestamp")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalTimestamp")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalTimestamp")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalTimestamp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalTimestamp int = 0

var condRecorderAuxMockmarshalTimestamp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalTimestamp(i int) {
	condRecorderAuxMockmarshalTimestamp.L.Lock()
	for recorderAuxMockmarshalTimestamp < i {
		condRecorderAuxMockmarshalTimestamp.Wait()
	}
	condRecorderAuxMockmarshalTimestamp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalTimestamp() {
	condRecorderAuxMockmarshalTimestamp.L.Lock()
	recorderAuxMockmarshalTimestamp++
	condRecorderAuxMockmarshalTimestamp.L.Unlock()
	condRecorderAuxMockmarshalTimestamp.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalTimestamp() (ret int) {
	condRecorderAuxMockmarshalTimestamp.L.Lock()
	ret = recorderAuxMockmarshalTimestamp
	condRecorderAuxMockmarshalTimestamp.L.Unlock()
	return
}

// marshalTimestamp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalTimestamp(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalTimestamp, ok := apomock.GetRegisteredFunc("gocql.marshalTimestamp")
	if ok {
		reta, retb = FuncAuxMockmarshalTimestamp.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalTimestamp ")
	}
	AuxMockIncrementRecorderAuxMockmarshalTimestamp()
	return
}

//
// Mock: marshalTuple(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalTuple struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalTuple MockArgsTypemarshalTuple

// AuxMockmarshalTuple(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalTuple(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalTuple = MockArgsTypemarshalTuple{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalTuple(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalTuple")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalTuple")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalTuple")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalTuple  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalTuple int = 0

var condRecorderAuxMockmarshalTuple *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalTuple(i int) {
	condRecorderAuxMockmarshalTuple.L.Lock()
	for recorderAuxMockmarshalTuple < i {
		condRecorderAuxMockmarshalTuple.Wait()
	}
	condRecorderAuxMockmarshalTuple.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalTuple() {
	condRecorderAuxMockmarshalTuple.L.Lock()
	recorderAuxMockmarshalTuple++
	condRecorderAuxMockmarshalTuple.L.Unlock()
	condRecorderAuxMockmarshalTuple.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalTuple() (ret int) {
	condRecorderAuxMockmarshalTuple.L.Lock()
	ret = recorderAuxMockmarshalTuple
	condRecorderAuxMockmarshalTuple.L.Unlock()
	return
}

// marshalTuple - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalTuple(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalTuple, ok := apomock.GetRegisteredFunc("gocql.marshalTuple")
	if ok {
		reta, retb = FuncAuxMockmarshalTuple.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalTuple ")
	}
	AuxMockIncrementRecorderAuxMockmarshalTuple()
	return
}

//
// Mock: unmarshalErrorf(argformat string, args ...interface{})(reta UnmarshalError)
//

type MockArgsTypeunmarshalErrorf struct {
	ApomockCallNumber int
	Argformat         string
	Args              []interface{}
}

var LastMockArgsunmarshalErrorf MockArgsTypeunmarshalErrorf

// AuxMockunmarshalErrorf(argformat string, args ...interface{})(reta UnmarshalError) - Generated mock function
func AuxMockunmarshalErrorf(argformat string, args ...interface{}) (reta UnmarshalError) {
	LastMockArgsunmarshalErrorf = MockArgsTypeunmarshalErrorf{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalErrorf(),
		Argformat:         argformat,
		Args:              args,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalErrorf")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalErrorf")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalErrorf")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(UnmarshalError)
	}
	return
}

// RecorderAuxMockunmarshalErrorf  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalErrorf int = 0

var condRecorderAuxMockunmarshalErrorf *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalErrorf(i int) {
	condRecorderAuxMockunmarshalErrorf.L.Lock()
	for recorderAuxMockunmarshalErrorf < i {
		condRecorderAuxMockunmarshalErrorf.Wait()
	}
	condRecorderAuxMockunmarshalErrorf.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalErrorf() {
	condRecorderAuxMockunmarshalErrorf.L.Lock()
	recorderAuxMockunmarshalErrorf++
	condRecorderAuxMockunmarshalErrorf.L.Unlock()
	condRecorderAuxMockunmarshalErrorf.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalErrorf() (ret int) {
	condRecorderAuxMockunmarshalErrorf.L.Lock()
	ret = recorderAuxMockunmarshalErrorf
	condRecorderAuxMockunmarshalErrorf.L.Unlock()
	return
}

// unmarshalErrorf - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalErrorf(argformat string, args ...interface{}) (reta UnmarshalError) {
	FuncAuxMockunmarshalErrorf, ok := apomock.GetRegisteredFunc("gocql.unmarshalErrorf")
	if ok {
		reta = FuncAuxMockunmarshalErrorf.(func(argformat string, args ...interface{}) (reta UnmarshalError))(argformat, args...)
	} else {
		panic("FuncAuxMockunmarshalErrorf ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalErrorf()
	return
}

//
// Mock: unmarshalVarint(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalVarint struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalVarint MockArgsTypeunmarshalVarint

// AuxMockunmarshalVarint(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalVarint(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalVarint = MockArgsTypeunmarshalVarint{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalVarint(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalVarint")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalVarint")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalVarint")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalVarint  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalVarint int = 0

var condRecorderAuxMockunmarshalVarint *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalVarint(i int) {
	condRecorderAuxMockunmarshalVarint.L.Lock()
	for recorderAuxMockunmarshalVarint < i {
		condRecorderAuxMockunmarshalVarint.Wait()
	}
	condRecorderAuxMockunmarshalVarint.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalVarint() {
	condRecorderAuxMockunmarshalVarint.L.Lock()
	recorderAuxMockunmarshalVarint++
	condRecorderAuxMockunmarshalVarint.L.Unlock()
	condRecorderAuxMockunmarshalVarint.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalVarint() (ret int) {
	condRecorderAuxMockunmarshalVarint.L.Lock()
	ret = recorderAuxMockunmarshalVarint
	condRecorderAuxMockunmarshalVarint.L.Unlock()
	return
}

// unmarshalVarint - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalVarint(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalVarint, ok := apomock.GetRegisteredFunc("gocql.unmarshalVarint")
	if ok {
		reta = FuncAuxMockunmarshalVarint.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalVarint ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalVarint()
	return
}

//
// Mock: unmarshalMap(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalMap struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalMap MockArgsTypeunmarshalMap

// AuxMockunmarshalMap(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalMap(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalMap = MockArgsTypeunmarshalMap{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalMap(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalMap")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalMap")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalMap")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalMap  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalMap int = 0

var condRecorderAuxMockunmarshalMap *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalMap(i int) {
	condRecorderAuxMockunmarshalMap.L.Lock()
	for recorderAuxMockunmarshalMap < i {
		condRecorderAuxMockunmarshalMap.Wait()
	}
	condRecorderAuxMockunmarshalMap.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalMap() {
	condRecorderAuxMockunmarshalMap.L.Lock()
	recorderAuxMockunmarshalMap++
	condRecorderAuxMockunmarshalMap.L.Unlock()
	condRecorderAuxMockunmarshalMap.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalMap() (ret int) {
	condRecorderAuxMockunmarshalMap.L.Lock()
	ret = recorderAuxMockunmarshalMap
	condRecorderAuxMockunmarshalMap.L.Unlock()
	return
}

// unmarshalMap - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalMap(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalMap, ok := apomock.GetRegisteredFunc("gocql.unmarshalMap")
	if ok {
		reta = FuncAuxMockunmarshalMap.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalMap ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalMap()
	return
}

//
// Mock: unmarshalUUID(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalUUID struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalUUID MockArgsTypeunmarshalUUID

// AuxMockunmarshalUUID(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalUUID(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalUUID = MockArgsTypeunmarshalUUID{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalUUID(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalUUID")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalUUID")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalUUID")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalUUID  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalUUID int = 0

var condRecorderAuxMockunmarshalUUID *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalUUID(i int) {
	condRecorderAuxMockunmarshalUUID.L.Lock()
	for recorderAuxMockunmarshalUUID < i {
		condRecorderAuxMockunmarshalUUID.Wait()
	}
	condRecorderAuxMockunmarshalUUID.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalUUID() {
	condRecorderAuxMockunmarshalUUID.L.Lock()
	recorderAuxMockunmarshalUUID++
	condRecorderAuxMockunmarshalUUID.L.Unlock()
	condRecorderAuxMockunmarshalUUID.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalUUID() (ret int) {
	condRecorderAuxMockunmarshalUUID.L.Lock()
	ret = recorderAuxMockunmarshalUUID
	condRecorderAuxMockunmarshalUUID.L.Unlock()
	return
}

// unmarshalUUID - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalUUID(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalUUID, ok := apomock.GetRegisteredFunc("gocql.unmarshalUUID")
	if ok {
		reta = FuncAuxMockunmarshalUUID.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalUUID ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalUUID()
	return
}

//
// Mock: unmarshalUDT(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalUDT struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalUDT MockArgsTypeunmarshalUDT

// AuxMockunmarshalUDT(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalUDT(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalUDT = MockArgsTypeunmarshalUDT{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalUDT(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalUDT")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalUDT")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalUDT")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalUDT  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalUDT int = 0

var condRecorderAuxMockunmarshalUDT *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalUDT(i int) {
	condRecorderAuxMockunmarshalUDT.L.Lock()
	for recorderAuxMockunmarshalUDT < i {
		condRecorderAuxMockunmarshalUDT.Wait()
	}
	condRecorderAuxMockunmarshalUDT.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalUDT() {
	condRecorderAuxMockunmarshalUDT.L.Lock()
	recorderAuxMockunmarshalUDT++
	condRecorderAuxMockunmarshalUDT.L.Unlock()
	condRecorderAuxMockunmarshalUDT.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalUDT() (ret int) {
	condRecorderAuxMockunmarshalUDT.L.Lock()
	ret = recorderAuxMockunmarshalUDT
	condRecorderAuxMockunmarshalUDT.L.Unlock()
	return
}

// unmarshalUDT - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalUDT(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalUDT, ok := apomock.GetRegisteredFunc("gocql.unmarshalUDT")
	if ok {
		reta = FuncAuxMockunmarshalUDT.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalUDT ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalUDT()
	return
}

//
// Mock: (recvt Type)String()(reta string)
//

type MockArgsTypeTypeString struct {
	ApomockCallNumber int
}

var LastMockArgsTypeString MockArgsTypeTypeString

// (recvt Type)AuxMockString()(reta string) - Generated mock function
func (recvt Type) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.Type.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Type.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Type.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockTypeString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockTypeString int = 0

var condRecorderAuxMockTypeString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockTypeString(i int) {
	condRecorderAuxMockTypeString.L.Lock()
	for recorderAuxMockTypeString < i {
		condRecorderAuxMockTypeString.Wait()
	}
	condRecorderAuxMockTypeString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockTypeString() {
	condRecorderAuxMockTypeString.L.Lock()
	recorderAuxMockTypeString++
	condRecorderAuxMockTypeString.L.Unlock()
	condRecorderAuxMockTypeString.Broadcast()
}
func AuxMockGetRecorderAuxMockTypeString() (ret int) {
	condRecorderAuxMockTypeString.L.Lock()
	ret = recorderAuxMockTypeString
	condRecorderAuxMockTypeString.L.Unlock()
	return
}

// (recvt Type)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt Type) String() (reta string) {
	FuncAuxMockTypeString, ok := apomock.GetRegisteredFunc("gocql.Type.String")
	if ok {
		reta = FuncAuxMockTypeString.(func(recvt Type) (reta string))(recvt)
	} else {
		panic("FuncAuxMockTypeString ")
	}
	AuxMockIncrementRecorderAuxMockTypeString()
	return
}

//
// Mock: marshalMap(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalMap struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalMap MockArgsTypemarshalMap

// AuxMockmarshalMap(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalMap(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalMap = MockArgsTypemarshalMap{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalMap(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalMap")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalMap")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalMap")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalMap  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalMap int = 0

var condRecorderAuxMockmarshalMap *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalMap(i int) {
	condRecorderAuxMockmarshalMap.L.Lock()
	for recorderAuxMockmarshalMap < i {
		condRecorderAuxMockmarshalMap.Wait()
	}
	condRecorderAuxMockmarshalMap.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalMap() {
	condRecorderAuxMockmarshalMap.L.Lock()
	recorderAuxMockmarshalMap++
	condRecorderAuxMockmarshalMap.L.Unlock()
	condRecorderAuxMockmarshalMap.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalMap() (ret int) {
	condRecorderAuxMockmarshalMap.L.Lock()
	ret = recorderAuxMockmarshalMap
	condRecorderAuxMockmarshalMap.L.Unlock()
	return
}

// marshalMap - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalMap(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalMap, ok := apomock.GetRegisteredFunc("gocql.marshalMap")
	if ok {
		reta, retb = FuncAuxMockmarshalMap.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalMap ")
	}
	AuxMockIncrementRecorderAuxMockmarshalMap()
	return
}

//
// Mock: unmarshalNullable(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalNullable struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalNullable MockArgsTypeunmarshalNullable

// AuxMockunmarshalNullable(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalNullable(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalNullable = MockArgsTypeunmarshalNullable{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalNullable(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalNullable")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalNullable")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalNullable")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalNullable  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalNullable int = 0

var condRecorderAuxMockunmarshalNullable *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalNullable(i int) {
	condRecorderAuxMockunmarshalNullable.L.Lock()
	for recorderAuxMockunmarshalNullable < i {
		condRecorderAuxMockunmarshalNullable.Wait()
	}
	condRecorderAuxMockunmarshalNullable.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalNullable() {
	condRecorderAuxMockunmarshalNullable.L.Lock()
	recorderAuxMockunmarshalNullable++
	condRecorderAuxMockunmarshalNullable.L.Unlock()
	condRecorderAuxMockunmarshalNullable.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalNullable() (ret int) {
	condRecorderAuxMockunmarshalNullable.L.Lock()
	ret = recorderAuxMockunmarshalNullable
	condRecorderAuxMockunmarshalNullable.L.Unlock()
	return
}

// unmarshalNullable - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalNullable(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalNullable, ok := apomock.GetRegisteredFunc("gocql.unmarshalNullable")
	if ok {
		reta = FuncAuxMockunmarshalNullable.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalNullable ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalNullable()
	return
}

//
// Mock: unmarshalVarchar(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalVarchar struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalVarchar MockArgsTypeunmarshalVarchar

// AuxMockunmarshalVarchar(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalVarchar(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalVarchar = MockArgsTypeunmarshalVarchar{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalVarchar(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalVarchar")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalVarchar")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalVarchar")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalVarchar  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalVarchar int = 0

var condRecorderAuxMockunmarshalVarchar *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalVarchar(i int) {
	condRecorderAuxMockunmarshalVarchar.L.Lock()
	for recorderAuxMockunmarshalVarchar < i {
		condRecorderAuxMockunmarshalVarchar.Wait()
	}
	condRecorderAuxMockunmarshalVarchar.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalVarchar() {
	condRecorderAuxMockunmarshalVarchar.L.Lock()
	recorderAuxMockunmarshalVarchar++
	condRecorderAuxMockunmarshalVarchar.L.Unlock()
	condRecorderAuxMockunmarshalVarchar.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalVarchar() (ret int) {
	condRecorderAuxMockunmarshalVarchar.L.Lock()
	ret = recorderAuxMockunmarshalVarchar
	condRecorderAuxMockunmarshalVarchar.L.Unlock()
	return
}

// unmarshalVarchar - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalVarchar(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalVarchar, ok := apomock.GetRegisteredFunc("gocql.unmarshalVarchar")
	if ok {
		reta = FuncAuxMockunmarshalVarchar.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalVarchar ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalVarchar()
	return
}

//
// Mock: marshalBool(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalBool struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalBool MockArgsTypemarshalBool

// AuxMockmarshalBool(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalBool(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalBool = MockArgsTypemarshalBool{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalBool(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalBool")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalBool")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalBool")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalBool  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalBool int = 0

var condRecorderAuxMockmarshalBool *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalBool(i int) {
	condRecorderAuxMockmarshalBool.L.Lock()
	for recorderAuxMockmarshalBool < i {
		condRecorderAuxMockmarshalBool.Wait()
	}
	condRecorderAuxMockmarshalBool.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalBool() {
	condRecorderAuxMockmarshalBool.L.Lock()
	recorderAuxMockmarshalBool++
	condRecorderAuxMockmarshalBool.L.Unlock()
	condRecorderAuxMockmarshalBool.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalBool() (ret int) {
	condRecorderAuxMockmarshalBool.L.Lock()
	ret = recorderAuxMockmarshalBool
	condRecorderAuxMockmarshalBool.L.Unlock()
	return
}

// marshalBool - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalBool(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalBool, ok := apomock.GetRegisteredFunc("gocql.marshalBool")
	if ok {
		reta, retb = FuncAuxMockmarshalBool.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalBool ")
	}
	AuxMockIncrementRecorderAuxMockmarshalBool()
	return
}

//
// Mock: marshalDouble(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalDouble struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalDouble MockArgsTypemarshalDouble

// AuxMockmarshalDouble(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalDouble(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalDouble = MockArgsTypemarshalDouble{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalDouble(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalDouble")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalDouble")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalDouble")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalDouble  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalDouble int = 0

var condRecorderAuxMockmarshalDouble *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalDouble(i int) {
	condRecorderAuxMockmarshalDouble.L.Lock()
	for recorderAuxMockmarshalDouble < i {
		condRecorderAuxMockmarshalDouble.Wait()
	}
	condRecorderAuxMockmarshalDouble.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalDouble() {
	condRecorderAuxMockmarshalDouble.L.Lock()
	recorderAuxMockmarshalDouble++
	condRecorderAuxMockmarshalDouble.L.Unlock()
	condRecorderAuxMockmarshalDouble.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalDouble() (ret int) {
	condRecorderAuxMockmarshalDouble.L.Lock()
	ret = recorderAuxMockmarshalDouble
	condRecorderAuxMockmarshalDouble.L.Unlock()
	return
}

// marshalDouble - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalDouble(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalDouble, ok := apomock.GetRegisteredFunc("gocql.marshalDouble")
	if ok {
		reta, retb = FuncAuxMockmarshalDouble.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalDouble ")
	}
	AuxMockIncrementRecorderAuxMockmarshalDouble()
	return
}

//
// Mock: writeCollectionSize(arginfo CollectionType, argn int, argbuf *bytes.Buffer)(reta error)
//

type MockArgsTypewriteCollectionSize struct {
	ApomockCallNumber int
	Arginfo           CollectionType
	Argn              int
	Argbuf            *bytes.Buffer
}

var LastMockArgswriteCollectionSize MockArgsTypewriteCollectionSize

// AuxMockwriteCollectionSize(arginfo CollectionType, argn int, argbuf *bytes.Buffer)(reta error) - Generated mock function
func AuxMockwriteCollectionSize(arginfo CollectionType, argn int, argbuf *bytes.Buffer) (reta error) {
	LastMockArgswriteCollectionSize = MockArgsTypewriteCollectionSize{
		ApomockCallNumber: AuxMockGetRecorderAuxMockwriteCollectionSize(),
		Arginfo:           arginfo,
		Argn:              argn,
		Argbuf:            argbuf,
	}
	rargs, rerr := apomock.GetNext("gocql.writeCollectionSize")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeCollectionSize")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeCollectionSize")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockwriteCollectionSize  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockwriteCollectionSize int = 0

var condRecorderAuxMockwriteCollectionSize *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockwriteCollectionSize(i int) {
	condRecorderAuxMockwriteCollectionSize.L.Lock()
	for recorderAuxMockwriteCollectionSize < i {
		condRecorderAuxMockwriteCollectionSize.Wait()
	}
	condRecorderAuxMockwriteCollectionSize.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockwriteCollectionSize() {
	condRecorderAuxMockwriteCollectionSize.L.Lock()
	recorderAuxMockwriteCollectionSize++
	condRecorderAuxMockwriteCollectionSize.L.Unlock()
	condRecorderAuxMockwriteCollectionSize.Broadcast()
}
func AuxMockGetRecorderAuxMockwriteCollectionSize() (ret int) {
	condRecorderAuxMockwriteCollectionSize.L.Lock()
	ret = recorderAuxMockwriteCollectionSize
	condRecorderAuxMockwriteCollectionSize.L.Unlock()
	return
}

// writeCollectionSize - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func writeCollectionSize(arginfo CollectionType, argn int, argbuf *bytes.Buffer) (reta error) {
	FuncAuxMockwriteCollectionSize, ok := apomock.GetRegisteredFunc("gocql.writeCollectionSize")
	if ok {
		reta = FuncAuxMockwriteCollectionSize.(func(arginfo CollectionType, argn int, argbuf *bytes.Buffer) (reta error))(arginfo, argn, argbuf)
	} else {
		panic("FuncAuxMockwriteCollectionSize ")
	}
	AuxMockIncrementRecorderAuxMockwriteCollectionSize()
	return
}

//
// Mock: marshalVarchar(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalVarchar struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalVarchar MockArgsTypemarshalVarchar

// AuxMockmarshalVarchar(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalVarchar(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalVarchar = MockArgsTypemarshalVarchar{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalVarchar(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalVarchar")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalVarchar")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalVarchar")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalVarchar  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalVarchar int = 0

var condRecorderAuxMockmarshalVarchar *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalVarchar(i int) {
	condRecorderAuxMockmarshalVarchar.L.Lock()
	for recorderAuxMockmarshalVarchar < i {
		condRecorderAuxMockmarshalVarchar.Wait()
	}
	condRecorderAuxMockmarshalVarchar.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalVarchar() {
	condRecorderAuxMockmarshalVarchar.L.Lock()
	recorderAuxMockmarshalVarchar++
	condRecorderAuxMockmarshalVarchar.L.Unlock()
	condRecorderAuxMockmarshalVarchar.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalVarchar() (ret int) {
	condRecorderAuxMockmarshalVarchar.L.Lock()
	ret = recorderAuxMockmarshalVarchar
	condRecorderAuxMockmarshalVarchar.L.Unlock()
	return
}

// marshalVarchar - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalVarchar(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalVarchar, ok := apomock.GetRegisteredFunc("gocql.marshalVarchar")
	if ok {
		reta, retb = FuncAuxMockmarshalVarchar.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalVarchar ")
	}
	AuxMockIncrementRecorderAuxMockmarshalVarchar()
	return
}

//
// Mock: bytesToInt64(argdata []byte)(retret int64)
//

type MockArgsTypebytesToInt64 struct {
	ApomockCallNumber int
	Argdata           []byte
}

var LastMockArgsbytesToInt64 MockArgsTypebytesToInt64

// AuxMockbytesToInt64(argdata []byte)(retret int64) - Generated mock function
func AuxMockbytesToInt64(argdata []byte) (retret int64) {
	LastMockArgsbytesToInt64 = MockArgsTypebytesToInt64{
		ApomockCallNumber: AuxMockGetRecorderAuxMockbytesToInt64(),
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.bytesToInt64")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.bytesToInt64")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.bytesToInt64")
	}
	if rargs.GetArg(0) != nil {
		retret = rargs.GetArg(0).(int64)
	}
	return
}

// RecorderAuxMockbytesToInt64  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockbytesToInt64 int = 0

var condRecorderAuxMockbytesToInt64 *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockbytesToInt64(i int) {
	condRecorderAuxMockbytesToInt64.L.Lock()
	for recorderAuxMockbytesToInt64 < i {
		condRecorderAuxMockbytesToInt64.Wait()
	}
	condRecorderAuxMockbytesToInt64.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockbytesToInt64() {
	condRecorderAuxMockbytesToInt64.L.Lock()
	recorderAuxMockbytesToInt64++
	condRecorderAuxMockbytesToInt64.L.Unlock()
	condRecorderAuxMockbytesToInt64.Broadcast()
}
func AuxMockGetRecorderAuxMockbytesToInt64() (ret int) {
	condRecorderAuxMockbytesToInt64.L.Lock()
	ret = recorderAuxMockbytesToInt64
	condRecorderAuxMockbytesToInt64.L.Unlock()
	return
}

// bytesToInt64 - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func bytesToInt64(argdata []byte) (retret int64) {
	FuncAuxMockbytesToInt64, ok := apomock.GetRegisteredFunc("gocql.bytesToInt64")
	if ok {
		retret = FuncAuxMockbytesToInt64.(func(argdata []byte) (retret int64))(argdata)
	} else {
		panic("FuncAuxMockbytesToInt64 ")
	}
	AuxMockIncrementRecorderAuxMockbytesToInt64()
	return
}

//
// Mock: unmarshalBigInt(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalBigInt struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalBigInt MockArgsTypeunmarshalBigInt

// AuxMockunmarshalBigInt(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalBigInt(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalBigInt = MockArgsTypeunmarshalBigInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalBigInt(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalBigInt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalBigInt")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalBigInt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalBigInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalBigInt int = 0

var condRecorderAuxMockunmarshalBigInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalBigInt(i int) {
	condRecorderAuxMockunmarshalBigInt.L.Lock()
	for recorderAuxMockunmarshalBigInt < i {
		condRecorderAuxMockunmarshalBigInt.Wait()
	}
	condRecorderAuxMockunmarshalBigInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalBigInt() {
	condRecorderAuxMockunmarshalBigInt.L.Lock()
	recorderAuxMockunmarshalBigInt++
	condRecorderAuxMockunmarshalBigInt.L.Unlock()
	condRecorderAuxMockunmarshalBigInt.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalBigInt() (ret int) {
	condRecorderAuxMockunmarshalBigInt.L.Lock()
	ret = recorderAuxMockunmarshalBigInt
	condRecorderAuxMockunmarshalBigInt.L.Unlock()
	return
}

// unmarshalBigInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalBigInt(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalBigInt, ok := apomock.GetRegisteredFunc("gocql.unmarshalBigInt")
	if ok {
		reta = FuncAuxMockunmarshalBigInt.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalBigInt ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalBigInt()
	return
}

//
// Mock: (recvc CollectionType)String()(reta string)
//

type MockArgsTypeCollectionTypeString struct {
	ApomockCallNumber int
}

var LastMockArgsCollectionTypeString MockArgsTypeCollectionTypeString

// (recvc CollectionType)AuxMockString()(reta string) - Generated mock function
func (recvc CollectionType) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.CollectionType.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.CollectionType.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.CollectionType.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockCollectionTypeString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockCollectionTypeString int = 0

var condRecorderAuxMockCollectionTypeString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockCollectionTypeString(i int) {
	condRecorderAuxMockCollectionTypeString.L.Lock()
	for recorderAuxMockCollectionTypeString < i {
		condRecorderAuxMockCollectionTypeString.Wait()
	}
	condRecorderAuxMockCollectionTypeString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockCollectionTypeString() {
	condRecorderAuxMockCollectionTypeString.L.Lock()
	recorderAuxMockCollectionTypeString++
	condRecorderAuxMockCollectionTypeString.L.Unlock()
	condRecorderAuxMockCollectionTypeString.Broadcast()
}
func AuxMockGetRecorderAuxMockCollectionTypeString() (ret int) {
	condRecorderAuxMockCollectionTypeString.L.Lock()
	ret = recorderAuxMockCollectionTypeString
	condRecorderAuxMockCollectionTypeString.L.Unlock()
	return
}

// (recvc CollectionType)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc CollectionType) String() (reta string) {
	FuncAuxMockCollectionTypeString, ok := apomock.GetRegisteredFunc("gocql.CollectionType.String")
	if ok {
		reta = FuncAuxMockCollectionTypeString.(func(recvc CollectionType) (reta string))(recvc)
	} else {
		panic("FuncAuxMockCollectionTypeString ")
	}
	AuxMockIncrementRecorderAuxMockCollectionTypeString()
	return
}

//
// Mock: unmarshalTimestamp(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalTimestamp struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalTimestamp MockArgsTypeunmarshalTimestamp

// AuxMockunmarshalTimestamp(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalTimestamp(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalTimestamp = MockArgsTypeunmarshalTimestamp{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalTimestamp(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalTimestamp")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalTimestamp")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalTimestamp")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalTimestamp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalTimestamp int = 0

var condRecorderAuxMockunmarshalTimestamp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalTimestamp(i int) {
	condRecorderAuxMockunmarshalTimestamp.L.Lock()
	for recorderAuxMockunmarshalTimestamp < i {
		condRecorderAuxMockunmarshalTimestamp.Wait()
	}
	condRecorderAuxMockunmarshalTimestamp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalTimestamp() {
	condRecorderAuxMockunmarshalTimestamp.L.Lock()
	recorderAuxMockunmarshalTimestamp++
	condRecorderAuxMockunmarshalTimestamp.L.Unlock()
	condRecorderAuxMockunmarshalTimestamp.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalTimestamp() (ret int) {
	condRecorderAuxMockunmarshalTimestamp.L.Lock()
	ret = recorderAuxMockunmarshalTimestamp
	condRecorderAuxMockunmarshalTimestamp.L.Unlock()
	return
}

// unmarshalTimestamp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalTimestamp(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalTimestamp, ok := apomock.GetRegisteredFunc("gocql.unmarshalTimestamp")
	if ok {
		reta = FuncAuxMockunmarshalTimestamp.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalTimestamp ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalTimestamp()
	return
}

//
// Mock: (recvs NativeType)Version()(reta byte)
//

type MockArgsTypeNativeTypeVersion struct {
	ApomockCallNumber int
}

var LastMockArgsNativeTypeVersion MockArgsTypeNativeTypeVersion

// (recvs NativeType)AuxMockVersion()(reta byte) - Generated mock function
func (recvs NativeType) AuxMockVersion() (reta byte) {
	rargs, rerr := apomock.GetNext("gocql.NativeType.Version")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.NativeType.Version")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.NativeType.Version")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(byte)
	}
	return
}

// RecorderAuxMockNativeTypeVersion  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNativeTypeVersion int = 0

var condRecorderAuxMockNativeTypeVersion *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNativeTypeVersion(i int) {
	condRecorderAuxMockNativeTypeVersion.L.Lock()
	for recorderAuxMockNativeTypeVersion < i {
		condRecorderAuxMockNativeTypeVersion.Wait()
	}
	condRecorderAuxMockNativeTypeVersion.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNativeTypeVersion() {
	condRecorderAuxMockNativeTypeVersion.L.Lock()
	recorderAuxMockNativeTypeVersion++
	condRecorderAuxMockNativeTypeVersion.L.Unlock()
	condRecorderAuxMockNativeTypeVersion.Broadcast()
}
func AuxMockGetRecorderAuxMockNativeTypeVersion() (ret int) {
	condRecorderAuxMockNativeTypeVersion.L.Lock()
	ret = recorderAuxMockNativeTypeVersion
	condRecorderAuxMockNativeTypeVersion.L.Unlock()
	return
}

// (recvs NativeType)Version - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs NativeType) Version() (reta byte) {
	FuncAuxMockNativeTypeVersion, ok := apomock.GetRegisteredFunc("gocql.NativeType.Version")
	if ok {
		reta = FuncAuxMockNativeTypeVersion.(func(recvs NativeType) (reta byte))(recvs)
	} else {
		panic("FuncAuxMockNativeTypeVersion ")
	}
	AuxMockIncrementRecorderAuxMockNativeTypeVersion()
	return
}

//
// Mock: unmarshalIntlike(arginfo TypeInfo, argint64Val int64, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalIntlike struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argint64Val       int64
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalIntlike MockArgsTypeunmarshalIntlike

// AuxMockunmarshalIntlike(arginfo TypeInfo, argint64Val int64, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalIntlike(arginfo TypeInfo, argint64Val int64, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalIntlike = MockArgsTypeunmarshalIntlike{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalIntlike(),
		Arginfo:           arginfo,
		Argint64Val:       argint64Val,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalIntlike")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalIntlike")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalIntlike")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalIntlike  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalIntlike int = 0

var condRecorderAuxMockunmarshalIntlike *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalIntlike(i int) {
	condRecorderAuxMockunmarshalIntlike.L.Lock()
	for recorderAuxMockunmarshalIntlike < i {
		condRecorderAuxMockunmarshalIntlike.Wait()
	}
	condRecorderAuxMockunmarshalIntlike.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalIntlike() {
	condRecorderAuxMockunmarshalIntlike.L.Lock()
	recorderAuxMockunmarshalIntlike++
	condRecorderAuxMockunmarshalIntlike.L.Unlock()
	condRecorderAuxMockunmarshalIntlike.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalIntlike() (ret int) {
	condRecorderAuxMockunmarshalIntlike.L.Lock()
	ret = recorderAuxMockunmarshalIntlike
	condRecorderAuxMockunmarshalIntlike.L.Unlock()
	return
}

// unmarshalIntlike - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalIntlike(arginfo TypeInfo, argint64Val int64, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalIntlike, ok := apomock.GetRegisteredFunc("gocql.unmarshalIntlike")
	if ok {
		reta = FuncAuxMockunmarshalIntlike.(func(arginfo TypeInfo, argint64Val int64, argdata []byte, argvalue interface{}) (reta error))(arginfo, argint64Val, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalIntlike ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalIntlike()
	return
}

//
// Mock: marshalUUID(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalUUID struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalUUID MockArgsTypemarshalUUID

// AuxMockmarshalUUID(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalUUID(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalUUID = MockArgsTypemarshalUUID{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalUUID(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalUUID")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalUUID")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalUUID")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalUUID  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalUUID int = 0

var condRecorderAuxMockmarshalUUID *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalUUID(i int) {
	condRecorderAuxMockmarshalUUID.L.Lock()
	for recorderAuxMockmarshalUUID < i {
		condRecorderAuxMockmarshalUUID.Wait()
	}
	condRecorderAuxMockmarshalUUID.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalUUID() {
	condRecorderAuxMockmarshalUUID.L.Lock()
	recorderAuxMockmarshalUUID++
	condRecorderAuxMockmarshalUUID.L.Unlock()
	condRecorderAuxMockmarshalUUID.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalUUID() (ret int) {
	condRecorderAuxMockmarshalUUID.L.Lock()
	ret = recorderAuxMockmarshalUUID
	condRecorderAuxMockmarshalUUID.L.Unlock()
	return
}

// marshalUUID - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalUUID(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalUUID, ok := apomock.GetRegisteredFunc("gocql.marshalUUID")
	if ok {
		reta, retb = FuncAuxMockmarshalUUID.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalUUID ")
	}
	AuxMockIncrementRecorderAuxMockmarshalUUID()
	return
}

//
// Mock: marshalUDT(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalUDT struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalUDT MockArgsTypemarshalUDT

// AuxMockmarshalUDT(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalUDT(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalUDT = MockArgsTypemarshalUDT{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalUDT(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalUDT")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalUDT")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalUDT")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalUDT  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalUDT int = 0

var condRecorderAuxMockmarshalUDT *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalUDT(i int) {
	condRecorderAuxMockmarshalUDT.L.Lock()
	for recorderAuxMockmarshalUDT < i {
		condRecorderAuxMockmarshalUDT.Wait()
	}
	condRecorderAuxMockmarshalUDT.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalUDT() {
	condRecorderAuxMockmarshalUDT.L.Lock()
	recorderAuxMockmarshalUDT++
	condRecorderAuxMockmarshalUDT.L.Unlock()
	condRecorderAuxMockmarshalUDT.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalUDT() (ret int) {
	condRecorderAuxMockmarshalUDT.L.Lock()
	ret = recorderAuxMockmarshalUDT
	condRecorderAuxMockmarshalUDT.L.Unlock()
	return
}

// marshalUDT - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalUDT(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalUDT, ok := apomock.GetRegisteredFunc("gocql.marshalUDT")
	if ok {
		reta, retb = FuncAuxMockmarshalUDT.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalUDT ")
	}
	AuxMockIncrementRecorderAuxMockmarshalUDT()
	return
}

//
// Mock: (recvm MarshalError)Error()(reta string)
//

type MockArgsTypeMarshalErrorError struct {
	ApomockCallNumber int
}

var LastMockArgsMarshalErrorError MockArgsTypeMarshalErrorError

// (recvm MarshalError)AuxMockError()(reta string) - Generated mock function
func (recvm MarshalError) AuxMockError() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.MarshalError.Error")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.MarshalError.Error")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.MarshalError.Error")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockMarshalErrorError  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockMarshalErrorError int = 0

var condRecorderAuxMockMarshalErrorError *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockMarshalErrorError(i int) {
	condRecorderAuxMockMarshalErrorError.L.Lock()
	for recorderAuxMockMarshalErrorError < i {
		condRecorderAuxMockMarshalErrorError.Wait()
	}
	condRecorderAuxMockMarshalErrorError.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockMarshalErrorError() {
	condRecorderAuxMockMarshalErrorError.L.Lock()
	recorderAuxMockMarshalErrorError++
	condRecorderAuxMockMarshalErrorError.L.Unlock()
	condRecorderAuxMockMarshalErrorError.Broadcast()
}
func AuxMockGetRecorderAuxMockMarshalErrorError() (ret int) {
	condRecorderAuxMockMarshalErrorError.L.Lock()
	ret = recorderAuxMockMarshalErrorError
	condRecorderAuxMockMarshalErrorError.L.Unlock()
	return
}

// (recvm MarshalError)Error - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvm MarshalError) Error() (reta string) {
	FuncAuxMockMarshalErrorError, ok := apomock.GetRegisteredFunc("gocql.MarshalError.Error")
	if ok {
		reta = FuncAuxMockMarshalErrorError.(func(recvm MarshalError) (reta string))(recvm)
	} else {
		panic("FuncAuxMockMarshalErrorError ")
	}
	AuxMockIncrementRecorderAuxMockMarshalErrorError()
	return
}

//
// Mock: encBool(argv bool)(reta []byte)
//

type MockArgsTypeencBool struct {
	ApomockCallNumber int
	Argv              bool
}

var LastMockArgsencBool MockArgsTypeencBool

// AuxMockencBool(argv bool)(reta []byte) - Generated mock function
func AuxMockencBool(argv bool) (reta []byte) {
	LastMockArgsencBool = MockArgsTypeencBool{
		ApomockCallNumber: AuxMockGetRecorderAuxMockencBool(),
		Argv:              argv,
	}
	rargs, rerr := apomock.GetNext("gocql.encBool")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.encBool")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.encBool")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockencBool  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockencBool int = 0

var condRecorderAuxMockencBool *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockencBool(i int) {
	condRecorderAuxMockencBool.L.Lock()
	for recorderAuxMockencBool < i {
		condRecorderAuxMockencBool.Wait()
	}
	condRecorderAuxMockencBool.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockencBool() {
	condRecorderAuxMockencBool.L.Lock()
	recorderAuxMockencBool++
	condRecorderAuxMockencBool.L.Unlock()
	condRecorderAuxMockencBool.Broadcast()
}
func AuxMockGetRecorderAuxMockencBool() (ret int) {
	condRecorderAuxMockencBool.L.Lock()
	ret = recorderAuxMockencBool
	condRecorderAuxMockencBool.L.Unlock()
	return
}

// encBool - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func encBool(argv bool) (reta []byte) {
	FuncAuxMockencBool, ok := apomock.GetRegisteredFunc("gocql.encBool")
	if ok {
		reta = FuncAuxMockencBool.(func(argv bool) (reta []byte))(argv)
	} else {
		panic("FuncAuxMockencBool ")
	}
	AuxMockIncrementRecorderAuxMockencBool()
	return
}

//
// Mock: unmarshalDecimal(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalDecimal struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalDecimal MockArgsTypeunmarshalDecimal

// AuxMockunmarshalDecimal(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalDecimal(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalDecimal = MockArgsTypeunmarshalDecimal{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalDecimal(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalDecimal")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalDecimal")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalDecimal")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalDecimal  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalDecimal int = 0

var condRecorderAuxMockunmarshalDecimal *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalDecimal(i int) {
	condRecorderAuxMockunmarshalDecimal.L.Lock()
	for recorderAuxMockunmarshalDecimal < i {
		condRecorderAuxMockunmarshalDecimal.Wait()
	}
	condRecorderAuxMockunmarshalDecimal.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalDecimal() {
	condRecorderAuxMockunmarshalDecimal.L.Lock()
	recorderAuxMockunmarshalDecimal++
	condRecorderAuxMockunmarshalDecimal.L.Unlock()
	condRecorderAuxMockunmarshalDecimal.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalDecimal() (ret int) {
	condRecorderAuxMockunmarshalDecimal.L.Lock()
	ret = recorderAuxMockunmarshalDecimal
	condRecorderAuxMockunmarshalDecimal.L.Unlock()
	return
}

// unmarshalDecimal - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalDecimal(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalDecimal, ok := apomock.GetRegisteredFunc("gocql.unmarshalDecimal")
	if ok {
		reta = FuncAuxMockunmarshalDecimal.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalDecimal ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalDecimal()
	return
}

//
// Mock: marshalList(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalList struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalList MockArgsTypemarshalList

// AuxMockmarshalList(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalList(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalList = MockArgsTypemarshalList{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalList(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalList")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalList")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalList")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalList  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalList int = 0

var condRecorderAuxMockmarshalList *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalList(i int) {
	condRecorderAuxMockmarshalList.L.Lock()
	for recorderAuxMockmarshalList < i {
		condRecorderAuxMockmarshalList.Wait()
	}
	condRecorderAuxMockmarshalList.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalList() {
	condRecorderAuxMockmarshalList.L.Lock()
	recorderAuxMockmarshalList++
	condRecorderAuxMockmarshalList.L.Unlock()
	condRecorderAuxMockmarshalList.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalList() (ret int) {
	condRecorderAuxMockmarshalList.L.Lock()
	ret = recorderAuxMockmarshalList
	condRecorderAuxMockmarshalList.L.Unlock()
	return
}

// marshalList - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalList(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalList, ok := apomock.GetRegisteredFunc("gocql.marshalList")
	if ok {
		reta, retb = FuncAuxMockmarshalList.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalList ")
	}
	AuxMockIncrementRecorderAuxMockmarshalList()
	return
}

//
// Mock: marshalInet(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalInet struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalInet MockArgsTypemarshalInet

// AuxMockmarshalInet(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalInet(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalInet = MockArgsTypemarshalInet{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalInet(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalInet")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalInet")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalInet")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalInet  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalInet int = 0

var condRecorderAuxMockmarshalInet *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalInet(i int) {
	condRecorderAuxMockmarshalInet.L.Lock()
	for recorderAuxMockmarshalInet < i {
		condRecorderAuxMockmarshalInet.Wait()
	}
	condRecorderAuxMockmarshalInet.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalInet() {
	condRecorderAuxMockmarshalInet.L.Lock()
	recorderAuxMockmarshalInet++
	condRecorderAuxMockmarshalInet.L.Unlock()
	condRecorderAuxMockmarshalInet.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalInet() (ret int) {
	condRecorderAuxMockmarshalInet.L.Lock()
	ret = recorderAuxMockmarshalInet
	condRecorderAuxMockmarshalInet.L.Unlock()
	return
}

// marshalInet - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalInet(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalInet, ok := apomock.GetRegisteredFunc("gocql.marshalInet")
	if ok {
		reta, retb = FuncAuxMockmarshalInet.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalInet ")
	}
	AuxMockIncrementRecorderAuxMockmarshalInet()
	return
}

//
// Mock: bytesToUint64(argdata []byte)(retret uint64)
//

type MockArgsTypebytesToUint64 struct {
	ApomockCallNumber int
	Argdata           []byte
}

var LastMockArgsbytesToUint64 MockArgsTypebytesToUint64

// AuxMockbytesToUint64(argdata []byte)(retret uint64) - Generated mock function
func AuxMockbytesToUint64(argdata []byte) (retret uint64) {
	LastMockArgsbytesToUint64 = MockArgsTypebytesToUint64{
		ApomockCallNumber: AuxMockGetRecorderAuxMockbytesToUint64(),
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.bytesToUint64")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.bytesToUint64")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.bytesToUint64")
	}
	if rargs.GetArg(0) != nil {
		retret = rargs.GetArg(0).(uint64)
	}
	return
}

// RecorderAuxMockbytesToUint64  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockbytesToUint64 int = 0

var condRecorderAuxMockbytesToUint64 *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockbytesToUint64(i int) {
	condRecorderAuxMockbytesToUint64.L.Lock()
	for recorderAuxMockbytesToUint64 < i {
		condRecorderAuxMockbytesToUint64.Wait()
	}
	condRecorderAuxMockbytesToUint64.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockbytesToUint64() {
	condRecorderAuxMockbytesToUint64.L.Lock()
	recorderAuxMockbytesToUint64++
	condRecorderAuxMockbytesToUint64.L.Unlock()
	condRecorderAuxMockbytesToUint64.Broadcast()
}
func AuxMockGetRecorderAuxMockbytesToUint64() (ret int) {
	condRecorderAuxMockbytesToUint64.L.Lock()
	ret = recorderAuxMockbytesToUint64
	condRecorderAuxMockbytesToUint64.L.Unlock()
	return
}

// bytesToUint64 - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func bytesToUint64(argdata []byte) (retret uint64) {
	FuncAuxMockbytesToUint64, ok := apomock.GetRegisteredFunc("gocql.bytesToUint64")
	if ok {
		retret = FuncAuxMockbytesToUint64.(func(argdata []byte) (retret uint64))(argdata)
	} else {
		panic("FuncAuxMockbytesToUint64 ")
	}
	AuxMockIncrementRecorderAuxMockbytesToUint64()
	return
}

//
// Mock: decBigInt(argdata []byte)(reta int64)
//

type MockArgsTypedecBigInt struct {
	ApomockCallNumber int
	Argdata           []byte
}

var LastMockArgsdecBigInt MockArgsTypedecBigInt

// AuxMockdecBigInt(argdata []byte)(reta int64) - Generated mock function
func AuxMockdecBigInt(argdata []byte) (reta int64) {
	LastMockArgsdecBigInt = MockArgsTypedecBigInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockdecBigInt(),
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.decBigInt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.decBigInt")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.decBigInt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int64)
	}
	return
}

// RecorderAuxMockdecBigInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockdecBigInt int = 0

var condRecorderAuxMockdecBigInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockdecBigInt(i int) {
	condRecorderAuxMockdecBigInt.L.Lock()
	for recorderAuxMockdecBigInt < i {
		condRecorderAuxMockdecBigInt.Wait()
	}
	condRecorderAuxMockdecBigInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockdecBigInt() {
	condRecorderAuxMockdecBigInt.L.Lock()
	recorderAuxMockdecBigInt++
	condRecorderAuxMockdecBigInt.L.Unlock()
	condRecorderAuxMockdecBigInt.Broadcast()
}
func AuxMockGetRecorderAuxMockdecBigInt() (ret int) {
	condRecorderAuxMockdecBigInt.L.Lock()
	ret = recorderAuxMockdecBigInt
	condRecorderAuxMockdecBigInt.L.Unlock()
	return
}

// decBigInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func decBigInt(argdata []byte) (reta int64) {
	FuncAuxMockdecBigInt, ok := apomock.GetRegisteredFunc("gocql.decBigInt")
	if ok {
		reta = FuncAuxMockdecBigInt.(func(argdata []byte) (reta int64))(argdata)
	} else {
		panic("FuncAuxMockdecBigInt ")
	}
	AuxMockIncrementRecorderAuxMockdecBigInt()
	return
}

//
// Mock: unmarshalDouble(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalDouble struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalDouble MockArgsTypeunmarshalDouble

// AuxMockunmarshalDouble(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalDouble(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalDouble = MockArgsTypeunmarshalDouble{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalDouble(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalDouble")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalDouble")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalDouble")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalDouble  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalDouble int = 0

var condRecorderAuxMockunmarshalDouble *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalDouble(i int) {
	condRecorderAuxMockunmarshalDouble.L.Lock()
	for recorderAuxMockunmarshalDouble < i {
		condRecorderAuxMockunmarshalDouble.Wait()
	}
	condRecorderAuxMockunmarshalDouble.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalDouble() {
	condRecorderAuxMockunmarshalDouble.L.Lock()
	recorderAuxMockunmarshalDouble++
	condRecorderAuxMockunmarshalDouble.L.Unlock()
	condRecorderAuxMockunmarshalDouble.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalDouble() (ret int) {
	condRecorderAuxMockunmarshalDouble.L.Lock()
	ret = recorderAuxMockunmarshalDouble
	condRecorderAuxMockunmarshalDouble.L.Unlock()
	return
}

// unmarshalDouble - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalDouble(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalDouble, ok := apomock.GetRegisteredFunc("gocql.unmarshalDouble")
	if ok {
		reta = FuncAuxMockunmarshalDouble.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalDouble ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalDouble()
	return
}

//
// Mock: (recvt NativeType)New()(reta interface{})
//

type MockArgsTypeNativeTypeNew struct {
	ApomockCallNumber int
}

var LastMockArgsNativeTypeNew MockArgsTypeNativeTypeNew

// (recvt NativeType)AuxMockNew()(reta interface{}) - Generated mock function
func (recvt NativeType) AuxMockNew() (reta interface{}) {
	rargs, rerr := apomock.GetNext("gocql.NativeType.New")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.NativeType.New")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.NativeType.New")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(interface{})
	}
	return
}

// RecorderAuxMockNativeTypeNew  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNativeTypeNew int = 0

var condRecorderAuxMockNativeTypeNew *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNativeTypeNew(i int) {
	condRecorderAuxMockNativeTypeNew.L.Lock()
	for recorderAuxMockNativeTypeNew < i {
		condRecorderAuxMockNativeTypeNew.Wait()
	}
	condRecorderAuxMockNativeTypeNew.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNativeTypeNew() {
	condRecorderAuxMockNativeTypeNew.L.Lock()
	recorderAuxMockNativeTypeNew++
	condRecorderAuxMockNativeTypeNew.L.Unlock()
	condRecorderAuxMockNativeTypeNew.Broadcast()
}
func AuxMockGetRecorderAuxMockNativeTypeNew() (ret int) {
	condRecorderAuxMockNativeTypeNew.L.Lock()
	ret = recorderAuxMockNativeTypeNew
	condRecorderAuxMockNativeTypeNew.L.Unlock()
	return
}

// (recvt NativeType)New - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt NativeType) New() (reta interface{}) {
	FuncAuxMockNativeTypeNew, ok := apomock.GetRegisteredFunc("gocql.NativeType.New")
	if ok {
		reta = FuncAuxMockNativeTypeNew.(func(recvt NativeType) (reta interface{}))(recvt)
	} else {
		panic("FuncAuxMockNativeTypeNew ")
	}
	AuxMockIncrementRecorderAuxMockNativeTypeNew()
	return
}

//
// Mock: (recvs NativeType)Custom()(reta string)
//

type MockArgsTypeNativeTypeCustom struct {
	ApomockCallNumber int
}

var LastMockArgsNativeTypeCustom MockArgsTypeNativeTypeCustom

// (recvs NativeType)AuxMockCustom()(reta string) - Generated mock function
func (recvs NativeType) AuxMockCustom() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.NativeType.Custom")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.NativeType.Custom")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.NativeType.Custom")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockNativeTypeCustom  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNativeTypeCustom int = 0

var condRecorderAuxMockNativeTypeCustom *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNativeTypeCustom(i int) {
	condRecorderAuxMockNativeTypeCustom.L.Lock()
	for recorderAuxMockNativeTypeCustom < i {
		condRecorderAuxMockNativeTypeCustom.Wait()
	}
	condRecorderAuxMockNativeTypeCustom.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNativeTypeCustom() {
	condRecorderAuxMockNativeTypeCustom.L.Lock()
	recorderAuxMockNativeTypeCustom++
	condRecorderAuxMockNativeTypeCustom.L.Unlock()
	condRecorderAuxMockNativeTypeCustom.Broadcast()
}
func AuxMockGetRecorderAuxMockNativeTypeCustom() (ret int) {
	condRecorderAuxMockNativeTypeCustom.L.Lock()
	ret = recorderAuxMockNativeTypeCustom
	condRecorderAuxMockNativeTypeCustom.L.Unlock()
	return
}

// (recvs NativeType)Custom - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs NativeType) Custom() (reta string) {
	FuncAuxMockNativeTypeCustom, ok := apomock.GetRegisteredFunc("gocql.NativeType.Custom")
	if ok {
		reta = FuncAuxMockNativeTypeCustom.(func(recvs NativeType) (reta string))(recvs)
	} else {
		panic("FuncAuxMockNativeTypeCustom ")
	}
	AuxMockIncrementRecorderAuxMockNativeTypeCustom()
	return
}

//
// Mock: decBigInt2C(argdata []byte, argn *big.Int)(reta *big.Int)
//

type MockArgsTypedecBigInt2C struct {
	ApomockCallNumber int
	Argdata           []byte
	Argn              *big.Int
}

var LastMockArgsdecBigInt2C MockArgsTypedecBigInt2C

// AuxMockdecBigInt2C(argdata []byte, argn *big.Int)(reta *big.Int) - Generated mock function
func AuxMockdecBigInt2C(argdata []byte, argn *big.Int) (reta *big.Int) {
	LastMockArgsdecBigInt2C = MockArgsTypedecBigInt2C{
		ApomockCallNumber: AuxMockGetRecorderAuxMockdecBigInt2C(),
		Argdata:           argdata,
		Argn:              argn,
	}
	rargs, rerr := apomock.GetNext("gocql.decBigInt2C")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.decBigInt2C")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.decBigInt2C")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*big.Int)
	}
	return
}

// RecorderAuxMockdecBigInt2C  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockdecBigInt2C int = 0

var condRecorderAuxMockdecBigInt2C *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockdecBigInt2C(i int) {
	condRecorderAuxMockdecBigInt2C.L.Lock()
	for recorderAuxMockdecBigInt2C < i {
		condRecorderAuxMockdecBigInt2C.Wait()
	}
	condRecorderAuxMockdecBigInt2C.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockdecBigInt2C() {
	condRecorderAuxMockdecBigInt2C.L.Lock()
	recorderAuxMockdecBigInt2C++
	condRecorderAuxMockdecBigInt2C.L.Unlock()
	condRecorderAuxMockdecBigInt2C.Broadcast()
}
func AuxMockGetRecorderAuxMockdecBigInt2C() (ret int) {
	condRecorderAuxMockdecBigInt2C.L.Lock()
	ret = recorderAuxMockdecBigInt2C
	condRecorderAuxMockdecBigInt2C.L.Unlock()
	return
}

// decBigInt2C - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func decBigInt2C(argdata []byte, argn *big.Int) (reta *big.Int) {
	FuncAuxMockdecBigInt2C, ok := apomock.GetRegisteredFunc("gocql.decBigInt2C")
	if ok {
		reta = FuncAuxMockdecBigInt2C.(func(argdata []byte, argn *big.Int) (reta *big.Int))(argdata, argn)
	} else {
		panic("FuncAuxMockdecBigInt2C ")
	}
	AuxMockIncrementRecorderAuxMockdecBigInt2C()
	return
}

//
// Mock: readCollectionSize(arginfo CollectionType, argdata []byte)(retsize int, retread int)
//

type MockArgsTypereadCollectionSize struct {
	ApomockCallNumber int
	Arginfo           CollectionType
	Argdata           []byte
}

var LastMockArgsreadCollectionSize MockArgsTypereadCollectionSize

// AuxMockreadCollectionSize(arginfo CollectionType, argdata []byte)(retsize int, retread int) - Generated mock function
func AuxMockreadCollectionSize(arginfo CollectionType, argdata []byte) (retsize int, retread int) {
	LastMockArgsreadCollectionSize = MockArgsTypereadCollectionSize{
		ApomockCallNumber: AuxMockGetRecorderAuxMockreadCollectionSize(),
		Arginfo:           arginfo,
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.readCollectionSize")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.readCollectionSize")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.readCollectionSize")
	}
	if rargs.GetArg(0) != nil {
		retsize = rargs.GetArg(0).(int)
	}
	if rargs.GetArg(1) != nil {
		retread = rargs.GetArg(1).(int)
	}
	return
}

// RecorderAuxMockreadCollectionSize  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockreadCollectionSize int = 0

var condRecorderAuxMockreadCollectionSize *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockreadCollectionSize(i int) {
	condRecorderAuxMockreadCollectionSize.L.Lock()
	for recorderAuxMockreadCollectionSize < i {
		condRecorderAuxMockreadCollectionSize.Wait()
	}
	condRecorderAuxMockreadCollectionSize.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockreadCollectionSize() {
	condRecorderAuxMockreadCollectionSize.L.Lock()
	recorderAuxMockreadCollectionSize++
	condRecorderAuxMockreadCollectionSize.L.Unlock()
	condRecorderAuxMockreadCollectionSize.Broadcast()
}
func AuxMockGetRecorderAuxMockreadCollectionSize() (ret int) {
	condRecorderAuxMockreadCollectionSize.L.Lock()
	ret = recorderAuxMockreadCollectionSize
	condRecorderAuxMockreadCollectionSize.L.Unlock()
	return
}

// readCollectionSize - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func readCollectionSize(arginfo CollectionType, argdata []byte) (retsize int, retread int) {
	FuncAuxMockreadCollectionSize, ok := apomock.GetRegisteredFunc("gocql.readCollectionSize")
	if ok {
		retsize, retread = FuncAuxMockreadCollectionSize.(func(arginfo CollectionType, argdata []byte) (retsize int, retread int))(arginfo, argdata)
	} else {
		panic("FuncAuxMockreadCollectionSize ")
	}
	AuxMockIncrementRecorderAuxMockreadCollectionSize()
	return
}

//
// Mock: (recvt TupleTypeInfo)New()(reta interface{})
//

type MockArgsTypeTupleTypeInfoNew struct {
	ApomockCallNumber int
}

var LastMockArgsTupleTypeInfoNew MockArgsTypeTupleTypeInfoNew

// (recvt TupleTypeInfo)AuxMockNew()(reta interface{}) - Generated mock function
func (recvt TupleTypeInfo) AuxMockNew() (reta interface{}) {
	rargs, rerr := apomock.GetNext("gocql.TupleTypeInfo.New")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.TupleTypeInfo.New")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.TupleTypeInfo.New")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(interface{})
	}
	return
}

// RecorderAuxMockTupleTypeInfoNew  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockTupleTypeInfoNew int = 0

var condRecorderAuxMockTupleTypeInfoNew *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockTupleTypeInfoNew(i int) {
	condRecorderAuxMockTupleTypeInfoNew.L.Lock()
	for recorderAuxMockTupleTypeInfoNew < i {
		condRecorderAuxMockTupleTypeInfoNew.Wait()
	}
	condRecorderAuxMockTupleTypeInfoNew.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockTupleTypeInfoNew() {
	condRecorderAuxMockTupleTypeInfoNew.L.Lock()
	recorderAuxMockTupleTypeInfoNew++
	condRecorderAuxMockTupleTypeInfoNew.L.Unlock()
	condRecorderAuxMockTupleTypeInfoNew.Broadcast()
}
func AuxMockGetRecorderAuxMockTupleTypeInfoNew() (ret int) {
	condRecorderAuxMockTupleTypeInfoNew.L.Lock()
	ret = recorderAuxMockTupleTypeInfoNew
	condRecorderAuxMockTupleTypeInfoNew.L.Unlock()
	return
}

// (recvt TupleTypeInfo)New - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt TupleTypeInfo) New() (reta interface{}) {
	FuncAuxMockTupleTypeInfoNew, ok := apomock.GetRegisteredFunc("gocql.TupleTypeInfo.New")
	if ok {
		reta = FuncAuxMockTupleTypeInfoNew.(func(recvt TupleTypeInfo) (reta interface{}))(recvt)
	} else {
		panic("FuncAuxMockTupleTypeInfoNew ")
	}
	AuxMockIncrementRecorderAuxMockTupleTypeInfoNew()
	return
}

//
// Mock: (recvu UDTTypeInfo)String()(reta string)
//

type MockArgsTypeUDTTypeInfoString struct {
	ApomockCallNumber int
}

var LastMockArgsUDTTypeInfoString MockArgsTypeUDTTypeInfoString

// (recvu UDTTypeInfo)AuxMockString()(reta string) - Generated mock function
func (recvu UDTTypeInfo) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.UDTTypeInfo.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UDTTypeInfo.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UDTTypeInfo.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockUDTTypeInfoString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUDTTypeInfoString int = 0

var condRecorderAuxMockUDTTypeInfoString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUDTTypeInfoString(i int) {
	condRecorderAuxMockUDTTypeInfoString.L.Lock()
	for recorderAuxMockUDTTypeInfoString < i {
		condRecorderAuxMockUDTTypeInfoString.Wait()
	}
	condRecorderAuxMockUDTTypeInfoString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUDTTypeInfoString() {
	condRecorderAuxMockUDTTypeInfoString.L.Lock()
	recorderAuxMockUDTTypeInfoString++
	condRecorderAuxMockUDTTypeInfoString.L.Unlock()
	condRecorderAuxMockUDTTypeInfoString.Broadcast()
}
func AuxMockGetRecorderAuxMockUDTTypeInfoString() (ret int) {
	condRecorderAuxMockUDTTypeInfoString.L.Lock()
	ret = recorderAuxMockUDTTypeInfoString
	condRecorderAuxMockUDTTypeInfoString.L.Unlock()
	return
}

// (recvu UDTTypeInfo)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu UDTTypeInfo) String() (reta string) {
	FuncAuxMockUDTTypeInfoString, ok := apomock.GetRegisteredFunc("gocql.UDTTypeInfo.String")
	if ok {
		reta = FuncAuxMockUDTTypeInfoString.(func(recvu UDTTypeInfo) (reta string))(recvu)
	} else {
		panic("FuncAuxMockUDTTypeInfoString ")
	}
	AuxMockIncrementRecorderAuxMockUDTTypeInfoString()
	return
}

//
// Mock: (recvs NativeType)String()(reta string)
//

type MockArgsTypeNativeTypeString struct {
	ApomockCallNumber int
}

var LastMockArgsNativeTypeString MockArgsTypeNativeTypeString

// (recvs NativeType)AuxMockString()(reta string) - Generated mock function
func (recvs NativeType) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.NativeType.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.NativeType.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.NativeType.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockNativeTypeString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNativeTypeString int = 0

var condRecorderAuxMockNativeTypeString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNativeTypeString(i int) {
	condRecorderAuxMockNativeTypeString.L.Lock()
	for recorderAuxMockNativeTypeString < i {
		condRecorderAuxMockNativeTypeString.Wait()
	}
	condRecorderAuxMockNativeTypeString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNativeTypeString() {
	condRecorderAuxMockNativeTypeString.L.Lock()
	recorderAuxMockNativeTypeString++
	condRecorderAuxMockNativeTypeString.L.Unlock()
	condRecorderAuxMockNativeTypeString.Broadcast()
}
func AuxMockGetRecorderAuxMockNativeTypeString() (ret int) {
	condRecorderAuxMockNativeTypeString.L.Lock()
	ret = recorderAuxMockNativeTypeString
	condRecorderAuxMockNativeTypeString.L.Unlock()
	return
}

// (recvs NativeType)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs NativeType) String() (reta string) {
	FuncAuxMockNativeTypeString, ok := apomock.GetRegisteredFunc("gocql.NativeType.String")
	if ok {
		reta = FuncAuxMockNativeTypeString.(func(recvs NativeType) (reta string))(recvs)
	} else {
		panic("FuncAuxMockNativeTypeString ")
	}
	AuxMockIncrementRecorderAuxMockNativeTypeString()
	return
}

//
// Mock: (recvu UDTTypeInfo)New()(reta interface{})
//

type MockArgsTypeUDTTypeInfoNew struct {
	ApomockCallNumber int
}

var LastMockArgsUDTTypeInfoNew MockArgsTypeUDTTypeInfoNew

// (recvu UDTTypeInfo)AuxMockNew()(reta interface{}) - Generated mock function
func (recvu UDTTypeInfo) AuxMockNew() (reta interface{}) {
	rargs, rerr := apomock.GetNext("gocql.UDTTypeInfo.New")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UDTTypeInfo.New")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UDTTypeInfo.New")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(interface{})
	}
	return
}

// RecorderAuxMockUDTTypeInfoNew  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUDTTypeInfoNew int = 0

var condRecorderAuxMockUDTTypeInfoNew *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUDTTypeInfoNew(i int) {
	condRecorderAuxMockUDTTypeInfoNew.L.Lock()
	for recorderAuxMockUDTTypeInfoNew < i {
		condRecorderAuxMockUDTTypeInfoNew.Wait()
	}
	condRecorderAuxMockUDTTypeInfoNew.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUDTTypeInfoNew() {
	condRecorderAuxMockUDTTypeInfoNew.L.Lock()
	recorderAuxMockUDTTypeInfoNew++
	condRecorderAuxMockUDTTypeInfoNew.L.Unlock()
	condRecorderAuxMockUDTTypeInfoNew.Broadcast()
}
func AuxMockGetRecorderAuxMockUDTTypeInfoNew() (ret int) {
	condRecorderAuxMockUDTTypeInfoNew.L.Lock()
	ret = recorderAuxMockUDTTypeInfoNew
	condRecorderAuxMockUDTTypeInfoNew.L.Unlock()
	return
}

// (recvu UDTTypeInfo)New - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu UDTTypeInfo) New() (reta interface{}) {
	FuncAuxMockUDTTypeInfoNew, ok := apomock.GetRegisteredFunc("gocql.UDTTypeInfo.New")
	if ok {
		reta = FuncAuxMockUDTTypeInfoNew.(func(recvu UDTTypeInfo) (reta interface{}))(recvu)
	} else {
		panic("FuncAuxMockUDTTypeInfoNew ")
	}
	AuxMockIncrementRecorderAuxMockUDTTypeInfoNew()
	return
}

//
// Mock: Unmarshal(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeUnmarshal struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsUnmarshal MockArgsTypeUnmarshal

// AuxMockUnmarshal(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockUnmarshal(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsUnmarshal = MockArgsTypeUnmarshal{
		ApomockCallNumber: AuxMockGetRecorderAuxMockUnmarshal(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.Unmarshal")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Unmarshal")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Unmarshal")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockUnmarshal  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUnmarshal int = 0

var condRecorderAuxMockUnmarshal *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUnmarshal(i int) {
	condRecorderAuxMockUnmarshal.L.Lock()
	for recorderAuxMockUnmarshal < i {
		condRecorderAuxMockUnmarshal.Wait()
	}
	condRecorderAuxMockUnmarshal.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUnmarshal() {
	condRecorderAuxMockUnmarshal.L.Lock()
	recorderAuxMockUnmarshal++
	condRecorderAuxMockUnmarshal.L.Unlock()
	condRecorderAuxMockUnmarshal.Broadcast()
}
func AuxMockGetRecorderAuxMockUnmarshal() (ret int) {
	condRecorderAuxMockUnmarshal.L.Lock()
	ret = recorderAuxMockUnmarshal
	condRecorderAuxMockUnmarshal.L.Unlock()
	return
}

// Unmarshal - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func Unmarshal(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockUnmarshal, ok := apomock.GetRegisteredFunc("gocql.Unmarshal")
	if ok {
		reta = FuncAuxMockUnmarshal.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockUnmarshal ")
	}
	AuxMockIncrementRecorderAuxMockUnmarshal()
	return
}

//
// Mock: isNullableValue(argvalue interface{})(reta bool)
//

type MockArgsTypeisNullableValue struct {
	ApomockCallNumber int
	Argvalue          interface{}
}

var LastMockArgsisNullableValue MockArgsTypeisNullableValue

// AuxMockisNullableValue(argvalue interface{})(reta bool) - Generated mock function
func AuxMockisNullableValue(argvalue interface{}) (reta bool) {
	LastMockArgsisNullableValue = MockArgsTypeisNullableValue{
		ApomockCallNumber: AuxMockGetRecorderAuxMockisNullableValue(),
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.isNullableValue")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.isNullableValue")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.isNullableValue")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockisNullableValue  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockisNullableValue int = 0

var condRecorderAuxMockisNullableValue *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockisNullableValue(i int) {
	condRecorderAuxMockisNullableValue.L.Lock()
	for recorderAuxMockisNullableValue < i {
		condRecorderAuxMockisNullableValue.Wait()
	}
	condRecorderAuxMockisNullableValue.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockisNullableValue() {
	condRecorderAuxMockisNullableValue.L.Lock()
	recorderAuxMockisNullableValue++
	condRecorderAuxMockisNullableValue.L.Unlock()
	condRecorderAuxMockisNullableValue.Broadcast()
}
func AuxMockGetRecorderAuxMockisNullableValue() (ret int) {
	condRecorderAuxMockisNullableValue.L.Lock()
	ret = recorderAuxMockisNullableValue
	condRecorderAuxMockisNullableValue.L.Unlock()
	return
}

// isNullableValue - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func isNullableValue(argvalue interface{}) (reta bool) {
	FuncAuxMockisNullableValue, ok := apomock.GetRegisteredFunc("gocql.isNullableValue")
	if ok {
		reta = FuncAuxMockisNullableValue.(func(argvalue interface{}) (reta bool))(argvalue)
	} else {
		panic("FuncAuxMockisNullableValue ")
	}
	AuxMockIncrementRecorderAuxMockisNullableValue()
	return
}

//
// Mock: isNullData(arginfo TypeInfo, argdata []byte)(reta bool)
//

type MockArgsTypeisNullData struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
}

var LastMockArgsisNullData MockArgsTypeisNullData

// AuxMockisNullData(arginfo TypeInfo, argdata []byte)(reta bool) - Generated mock function
func AuxMockisNullData(arginfo TypeInfo, argdata []byte) (reta bool) {
	LastMockArgsisNullData = MockArgsTypeisNullData{
		ApomockCallNumber: AuxMockGetRecorderAuxMockisNullData(),
		Arginfo:           arginfo,
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.isNullData")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.isNullData")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.isNullData")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockisNullData  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockisNullData int = 0

var condRecorderAuxMockisNullData *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockisNullData(i int) {
	condRecorderAuxMockisNullData.L.Lock()
	for recorderAuxMockisNullData < i {
		condRecorderAuxMockisNullData.Wait()
	}
	condRecorderAuxMockisNullData.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockisNullData() {
	condRecorderAuxMockisNullData.L.Lock()
	recorderAuxMockisNullData++
	condRecorderAuxMockisNullData.L.Unlock()
	condRecorderAuxMockisNullData.Broadcast()
}
func AuxMockGetRecorderAuxMockisNullData() (ret int) {
	condRecorderAuxMockisNullData.L.Lock()
	ret = recorderAuxMockisNullData
	condRecorderAuxMockisNullData.L.Unlock()
	return
}

// isNullData - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func isNullData(arginfo TypeInfo, argdata []byte) (reta bool) {
	FuncAuxMockisNullData, ok := apomock.GetRegisteredFunc("gocql.isNullData")
	if ok {
		reta = FuncAuxMockisNullData.(func(arginfo TypeInfo, argdata []byte) (reta bool))(arginfo, argdata)
	} else {
		panic("FuncAuxMockisNullData ")
	}
	AuxMockIncrementRecorderAuxMockisNullData()
	return
}

//
// Mock: encBigInt(argx int64)(reta []byte)
//

type MockArgsTypeencBigInt struct {
	ApomockCallNumber int
	Argx              int64
}

var LastMockArgsencBigInt MockArgsTypeencBigInt

// AuxMockencBigInt(argx int64)(reta []byte) - Generated mock function
func AuxMockencBigInt(argx int64) (reta []byte) {
	LastMockArgsencBigInt = MockArgsTypeencBigInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockencBigInt(),
		Argx:              argx,
	}
	rargs, rerr := apomock.GetNext("gocql.encBigInt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.encBigInt")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.encBigInt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockencBigInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockencBigInt int = 0

var condRecorderAuxMockencBigInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockencBigInt(i int) {
	condRecorderAuxMockencBigInt.L.Lock()
	for recorderAuxMockencBigInt < i {
		condRecorderAuxMockencBigInt.Wait()
	}
	condRecorderAuxMockencBigInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockencBigInt() {
	condRecorderAuxMockencBigInt.L.Lock()
	recorderAuxMockencBigInt++
	condRecorderAuxMockencBigInt.L.Unlock()
	condRecorderAuxMockencBigInt.Broadcast()
}
func AuxMockGetRecorderAuxMockencBigInt() (ret int) {
	condRecorderAuxMockencBigInt.L.Lock()
	ret = recorderAuxMockencBigInt
	condRecorderAuxMockencBigInt.L.Unlock()
	return
}

// encBigInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func encBigInt(argx int64) (reta []byte) {
	FuncAuxMockencBigInt, ok := apomock.GetRegisteredFunc("gocql.encBigInt")
	if ok {
		reta = FuncAuxMockencBigInt.(func(argx int64) (reta []byte))(argx)
	} else {
		panic("FuncAuxMockencBigInt ")
	}
	AuxMockIncrementRecorderAuxMockencBigInt()
	return
}

//
// Mock: unmarshalTuple(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalTuple struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalTuple MockArgsTypeunmarshalTuple

// AuxMockunmarshalTuple(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalTuple(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalTuple = MockArgsTypeunmarshalTuple{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalTuple(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalTuple")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalTuple")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalTuple")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalTuple  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalTuple int = 0

var condRecorderAuxMockunmarshalTuple *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalTuple(i int) {
	condRecorderAuxMockunmarshalTuple.L.Lock()
	for recorderAuxMockunmarshalTuple < i {
		condRecorderAuxMockunmarshalTuple.Wait()
	}
	condRecorderAuxMockunmarshalTuple.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalTuple() {
	condRecorderAuxMockunmarshalTuple.L.Lock()
	recorderAuxMockunmarshalTuple++
	condRecorderAuxMockunmarshalTuple.L.Unlock()
	condRecorderAuxMockunmarshalTuple.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalTuple() (ret int) {
	condRecorderAuxMockunmarshalTuple.L.Lock()
	ret = recorderAuxMockunmarshalTuple
	condRecorderAuxMockunmarshalTuple.L.Unlock()
	return
}

// unmarshalTuple - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalTuple(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalTuple, ok := apomock.GetRegisteredFunc("gocql.unmarshalTuple")
	if ok {
		reta = FuncAuxMockunmarshalTuple.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalTuple ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalTuple()
	return
}

//
// Mock: unmarshalList(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalList struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalList MockArgsTypeunmarshalList

// AuxMockunmarshalList(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalList(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalList = MockArgsTypeunmarshalList{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalList(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalList")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalList")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalList")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalList  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalList int = 0

var condRecorderAuxMockunmarshalList *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalList(i int) {
	condRecorderAuxMockunmarshalList.L.Lock()
	for recorderAuxMockunmarshalList < i {
		condRecorderAuxMockunmarshalList.Wait()
	}
	condRecorderAuxMockunmarshalList.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalList() {
	condRecorderAuxMockunmarshalList.L.Lock()
	recorderAuxMockunmarshalList++
	condRecorderAuxMockunmarshalList.L.Unlock()
	condRecorderAuxMockunmarshalList.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalList() (ret int) {
	condRecorderAuxMockunmarshalList.L.Lock()
	ret = recorderAuxMockunmarshalList
	condRecorderAuxMockunmarshalList.L.Unlock()
	return
}

// unmarshalList - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalList(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalList, ok := apomock.GetRegisteredFunc("gocql.unmarshalList")
	if ok {
		reta = FuncAuxMockunmarshalList.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalList ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalList()
	return
}

//
// Mock: unmarshalTimeUUID(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalTimeUUID struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalTimeUUID MockArgsTypeunmarshalTimeUUID

// AuxMockunmarshalTimeUUID(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalTimeUUID(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalTimeUUID = MockArgsTypeunmarshalTimeUUID{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalTimeUUID(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalTimeUUID")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalTimeUUID")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalTimeUUID")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalTimeUUID  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalTimeUUID int = 0

var condRecorderAuxMockunmarshalTimeUUID *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalTimeUUID(i int) {
	condRecorderAuxMockunmarshalTimeUUID.L.Lock()
	for recorderAuxMockunmarshalTimeUUID < i {
		condRecorderAuxMockunmarshalTimeUUID.Wait()
	}
	condRecorderAuxMockunmarshalTimeUUID.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalTimeUUID() {
	condRecorderAuxMockunmarshalTimeUUID.L.Lock()
	recorderAuxMockunmarshalTimeUUID++
	condRecorderAuxMockunmarshalTimeUUID.L.Unlock()
	condRecorderAuxMockunmarshalTimeUUID.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalTimeUUID() (ret int) {
	condRecorderAuxMockunmarshalTimeUUID.L.Lock()
	ret = recorderAuxMockunmarshalTimeUUID
	condRecorderAuxMockunmarshalTimeUUID.L.Unlock()
	return
}

// unmarshalTimeUUID - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalTimeUUID(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalTimeUUID, ok := apomock.GetRegisteredFunc("gocql.unmarshalTimeUUID")
	if ok {
		reta = FuncAuxMockunmarshalTimeUUID.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalTimeUUID ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalTimeUUID()
	return
}

//
// Mock: marshalInt(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalInt struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalInt MockArgsTypemarshalInt

// AuxMockmarshalInt(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalInt(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalInt = MockArgsTypemarshalInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalInt(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalInt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalInt")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalInt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalInt int = 0

var condRecorderAuxMockmarshalInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalInt(i int) {
	condRecorderAuxMockmarshalInt.L.Lock()
	for recorderAuxMockmarshalInt < i {
		condRecorderAuxMockmarshalInt.Wait()
	}
	condRecorderAuxMockmarshalInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalInt() {
	condRecorderAuxMockmarshalInt.L.Lock()
	recorderAuxMockmarshalInt++
	condRecorderAuxMockmarshalInt.L.Unlock()
	condRecorderAuxMockmarshalInt.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalInt() (ret int) {
	condRecorderAuxMockmarshalInt.L.Lock()
	ret = recorderAuxMockmarshalInt
	condRecorderAuxMockmarshalInt.L.Unlock()
	return
}

// marshalInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalInt(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalInt, ok := apomock.GetRegisteredFunc("gocql.marshalInt")
	if ok {
		reta, retb = FuncAuxMockmarshalInt.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalInt ")
	}
	AuxMockIncrementRecorderAuxMockmarshalInt()
	return
}

//
// Mock: encInt(argx int32)(reta []byte)
//

type MockArgsTypeencInt struct {
	ApomockCallNumber int
	Argx              int32
}

var LastMockArgsencInt MockArgsTypeencInt

// AuxMockencInt(argx int32)(reta []byte) - Generated mock function
func AuxMockencInt(argx int32) (reta []byte) {
	LastMockArgsencInt = MockArgsTypeencInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockencInt(),
		Argx:              argx,
	}
	rargs, rerr := apomock.GetNext("gocql.encInt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.encInt")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.encInt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockencInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockencInt int = 0

var condRecorderAuxMockencInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockencInt(i int) {
	condRecorderAuxMockencInt.L.Lock()
	for recorderAuxMockencInt < i {
		condRecorderAuxMockencInt.Wait()
	}
	condRecorderAuxMockencInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockencInt() {
	condRecorderAuxMockencInt.L.Lock()
	recorderAuxMockencInt++
	condRecorderAuxMockencInt.L.Unlock()
	condRecorderAuxMockencInt.Broadcast()
}
func AuxMockGetRecorderAuxMockencInt() (ret int) {
	condRecorderAuxMockencInt.L.Lock()
	ret = recorderAuxMockencInt
	condRecorderAuxMockencInt.L.Unlock()
	return
}

// encInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func encInt(argx int32) (reta []byte) {
	FuncAuxMockencInt, ok := apomock.GetRegisteredFunc("gocql.encInt")
	if ok {
		reta = FuncAuxMockencInt.(func(argx int32) (reta []byte))(argx)
	} else {
		panic("FuncAuxMockencInt ")
	}
	AuxMockIncrementRecorderAuxMockencInt()
	return
}

//
// Mock: decBool(argv []byte)(reta bool)
//

type MockArgsTypedecBool struct {
	ApomockCallNumber int
	Argv              []byte
}

var LastMockArgsdecBool MockArgsTypedecBool

// AuxMockdecBool(argv []byte)(reta bool) - Generated mock function
func AuxMockdecBool(argv []byte) (reta bool) {
	LastMockArgsdecBool = MockArgsTypedecBool{
		ApomockCallNumber: AuxMockGetRecorderAuxMockdecBool(),
		Argv:              argv,
	}
	rargs, rerr := apomock.GetNext("gocql.decBool")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.decBool")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.decBool")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockdecBool  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockdecBool int = 0

var condRecorderAuxMockdecBool *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockdecBool(i int) {
	condRecorderAuxMockdecBool.L.Lock()
	for recorderAuxMockdecBool < i {
		condRecorderAuxMockdecBool.Wait()
	}
	condRecorderAuxMockdecBool.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockdecBool() {
	condRecorderAuxMockdecBool.L.Lock()
	recorderAuxMockdecBool++
	condRecorderAuxMockdecBool.L.Unlock()
	condRecorderAuxMockdecBool.Broadcast()
}
func AuxMockGetRecorderAuxMockdecBool() (ret int) {
	condRecorderAuxMockdecBool.L.Lock()
	ret = recorderAuxMockdecBool
	condRecorderAuxMockdecBool.L.Unlock()
	return
}

// decBool - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func decBool(argv []byte) (reta bool) {
	FuncAuxMockdecBool, ok := apomock.GetRegisteredFunc("gocql.decBool")
	if ok {
		reta = FuncAuxMockdecBool.(func(argv []byte) (reta bool))(argv)
	} else {
		panic("FuncAuxMockdecBool ")
	}
	AuxMockIncrementRecorderAuxMockdecBool()
	return
}

//
// Mock: unmarshalFloat(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalFloat struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalFloat MockArgsTypeunmarshalFloat

// AuxMockunmarshalFloat(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalFloat(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalFloat = MockArgsTypeunmarshalFloat{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalFloat(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalFloat")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalFloat")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalFloat")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalFloat  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalFloat int = 0

var condRecorderAuxMockunmarshalFloat *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalFloat(i int) {
	condRecorderAuxMockunmarshalFloat.L.Lock()
	for recorderAuxMockunmarshalFloat < i {
		condRecorderAuxMockunmarshalFloat.Wait()
	}
	condRecorderAuxMockunmarshalFloat.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalFloat() {
	condRecorderAuxMockunmarshalFloat.L.Lock()
	recorderAuxMockunmarshalFloat++
	condRecorderAuxMockunmarshalFloat.L.Unlock()
	condRecorderAuxMockunmarshalFloat.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalFloat() (ret int) {
	condRecorderAuxMockunmarshalFloat.L.Lock()
	ret = recorderAuxMockunmarshalFloat
	condRecorderAuxMockunmarshalFloat.L.Unlock()
	return
}

// unmarshalFloat - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalFloat(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalFloat, ok := apomock.GetRegisteredFunc("gocql.unmarshalFloat")
	if ok {
		reta = FuncAuxMockunmarshalFloat.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalFloat ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalFloat()
	return
}

//
// Mock: encBigInt2C(argn *big.Int)(reta []byte)
//

type MockArgsTypeencBigInt2C struct {
	ApomockCallNumber int
	Argn              *big.Int
}

var LastMockArgsencBigInt2C MockArgsTypeencBigInt2C

// AuxMockencBigInt2C(argn *big.Int)(reta []byte) - Generated mock function
func AuxMockencBigInt2C(argn *big.Int) (reta []byte) {
	LastMockArgsencBigInt2C = MockArgsTypeencBigInt2C{
		ApomockCallNumber: AuxMockGetRecorderAuxMockencBigInt2C(),
		Argn:              argn,
	}
	rargs, rerr := apomock.GetNext("gocql.encBigInt2C")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.encBigInt2C")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.encBigInt2C")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockencBigInt2C  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockencBigInt2C int = 0

var condRecorderAuxMockencBigInt2C *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockencBigInt2C(i int) {
	condRecorderAuxMockencBigInt2C.L.Lock()
	for recorderAuxMockencBigInt2C < i {
		condRecorderAuxMockencBigInt2C.Wait()
	}
	condRecorderAuxMockencBigInt2C.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockencBigInt2C() {
	condRecorderAuxMockencBigInt2C.L.Lock()
	recorderAuxMockencBigInt2C++
	condRecorderAuxMockencBigInt2C.L.Unlock()
	condRecorderAuxMockencBigInt2C.Broadcast()
}
func AuxMockGetRecorderAuxMockencBigInt2C() (ret int) {
	condRecorderAuxMockencBigInt2C.L.Lock()
	ret = recorderAuxMockencBigInt2C
	condRecorderAuxMockencBigInt2C.L.Unlock()
	return
}

// encBigInt2C - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func encBigInt2C(argn *big.Int) (reta []byte) {
	FuncAuxMockencBigInt2C, ok := apomock.GetRegisteredFunc("gocql.encBigInt2C")
	if ok {
		reta = FuncAuxMockencBigInt2C.(func(argn *big.Int) (reta []byte))(argn)
	} else {
		panic("FuncAuxMockencBigInt2C ")
	}
	AuxMockIncrementRecorderAuxMockencBigInt2C()
	return
}

//
// Mock: marshalBigInt(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalBigInt struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalBigInt MockArgsTypemarshalBigInt

// AuxMockmarshalBigInt(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalBigInt(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalBigInt = MockArgsTypemarshalBigInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalBigInt(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalBigInt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalBigInt")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalBigInt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalBigInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalBigInt int = 0

var condRecorderAuxMockmarshalBigInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalBigInt(i int) {
	condRecorderAuxMockmarshalBigInt.L.Lock()
	for recorderAuxMockmarshalBigInt < i {
		condRecorderAuxMockmarshalBigInt.Wait()
	}
	condRecorderAuxMockmarshalBigInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalBigInt() {
	condRecorderAuxMockmarshalBigInt.L.Lock()
	recorderAuxMockmarshalBigInt++
	condRecorderAuxMockmarshalBigInt.L.Unlock()
	condRecorderAuxMockmarshalBigInt.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalBigInt() (ret int) {
	condRecorderAuxMockmarshalBigInt.L.Lock()
	ret = recorderAuxMockmarshalBigInt
	condRecorderAuxMockmarshalBigInt.L.Unlock()
	return
}

// marshalBigInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalBigInt(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalBigInt, ok := apomock.GetRegisteredFunc("gocql.marshalBigInt")
	if ok {
		reta, retb = FuncAuxMockmarshalBigInt.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalBigInt ")
	}
	AuxMockIncrementRecorderAuxMockmarshalBigInt()
	return
}

//
// Mock: unmarshalInet(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalInet struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalInet MockArgsTypeunmarshalInet

// AuxMockunmarshalInet(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalInet(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalInet = MockArgsTypeunmarshalInet{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalInet(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalInet")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalInet")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalInet")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalInet  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalInet int = 0

var condRecorderAuxMockunmarshalInet *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalInet(i int) {
	condRecorderAuxMockunmarshalInet.L.Lock()
	for recorderAuxMockunmarshalInet < i {
		condRecorderAuxMockunmarshalInet.Wait()
	}
	condRecorderAuxMockunmarshalInet.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalInet() {
	condRecorderAuxMockunmarshalInet.L.Lock()
	recorderAuxMockunmarshalInet++
	condRecorderAuxMockunmarshalInet.L.Unlock()
	condRecorderAuxMockunmarshalInet.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalInet() (ret int) {
	condRecorderAuxMockunmarshalInet.L.Lock()
	ret = recorderAuxMockunmarshalInet
	condRecorderAuxMockunmarshalInet.L.Unlock()
	return
}

// unmarshalInet - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalInet(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalInet, ok := apomock.GetRegisteredFunc("gocql.unmarshalInet")
	if ok {
		reta = FuncAuxMockunmarshalInet.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalInet ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalInet()
	return
}

//
// Mock: decInt(argx []byte)(reta int32)
//

type MockArgsTypedecInt struct {
	ApomockCallNumber int
	Argx              []byte
}

var LastMockArgsdecInt MockArgsTypedecInt

// AuxMockdecInt(argx []byte)(reta int32) - Generated mock function
func AuxMockdecInt(argx []byte) (reta int32) {
	LastMockArgsdecInt = MockArgsTypedecInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockdecInt(),
		Argx:              argx,
	}
	rargs, rerr := apomock.GetNext("gocql.decInt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.decInt")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.decInt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int32)
	}
	return
}

// RecorderAuxMockdecInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockdecInt int = 0

var condRecorderAuxMockdecInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockdecInt(i int) {
	condRecorderAuxMockdecInt.L.Lock()
	for recorderAuxMockdecInt < i {
		condRecorderAuxMockdecInt.Wait()
	}
	condRecorderAuxMockdecInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockdecInt() {
	condRecorderAuxMockdecInt.L.Lock()
	recorderAuxMockdecInt++
	condRecorderAuxMockdecInt.L.Unlock()
	condRecorderAuxMockdecInt.Broadcast()
}
func AuxMockGetRecorderAuxMockdecInt() (ret int) {
	condRecorderAuxMockdecInt.L.Lock()
	ret = recorderAuxMockdecInt
	condRecorderAuxMockdecInt.L.Unlock()
	return
}

// decInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func decInt(argx []byte) (reta int32) {
	FuncAuxMockdecInt, ok := apomock.GetRegisteredFunc("gocql.decInt")
	if ok {
		reta = FuncAuxMockdecInt.(func(argx []byte) (reta int32))(argx)
	} else {
		panic("FuncAuxMockdecInt ")
	}
	AuxMockIncrementRecorderAuxMockdecInt()
	return
}

//
// Mock: marshalVarint(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalVarint struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalVarint MockArgsTypemarshalVarint

// AuxMockmarshalVarint(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalVarint(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalVarint = MockArgsTypemarshalVarint{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalVarint(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalVarint")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalVarint")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalVarint")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalVarint  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalVarint int = 0

var condRecorderAuxMockmarshalVarint *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalVarint(i int) {
	condRecorderAuxMockmarshalVarint.L.Lock()
	for recorderAuxMockmarshalVarint < i {
		condRecorderAuxMockmarshalVarint.Wait()
	}
	condRecorderAuxMockmarshalVarint.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalVarint() {
	condRecorderAuxMockmarshalVarint.L.Lock()
	recorderAuxMockmarshalVarint++
	condRecorderAuxMockmarshalVarint.L.Unlock()
	condRecorderAuxMockmarshalVarint.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalVarint() (ret int) {
	condRecorderAuxMockmarshalVarint.L.Lock()
	ret = recorderAuxMockmarshalVarint
	condRecorderAuxMockmarshalVarint.L.Unlock()
	return
}

// marshalVarint - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalVarint(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalVarint, ok := apomock.GetRegisteredFunc("gocql.marshalVarint")
	if ok {
		reta, retb = FuncAuxMockmarshalVarint.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalVarint ")
	}
	AuxMockIncrementRecorderAuxMockmarshalVarint()
	return
}

//
// Mock: marshalErrorf(argformat string, args ...interface{})(reta MarshalError)
//

type MockArgsTypemarshalErrorf struct {
	ApomockCallNumber int
	Argformat         string
	Args              []interface{}
}

var LastMockArgsmarshalErrorf MockArgsTypemarshalErrorf

// AuxMockmarshalErrorf(argformat string, args ...interface{})(reta MarshalError) - Generated mock function
func AuxMockmarshalErrorf(argformat string, args ...interface{}) (reta MarshalError) {
	LastMockArgsmarshalErrorf = MockArgsTypemarshalErrorf{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalErrorf(),
		Argformat:         argformat,
		Args:              args,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalErrorf")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalErrorf")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.marshalErrorf")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(MarshalError)
	}
	return
}

// RecorderAuxMockmarshalErrorf  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalErrorf int = 0

var condRecorderAuxMockmarshalErrorf *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalErrorf(i int) {
	condRecorderAuxMockmarshalErrorf.L.Lock()
	for recorderAuxMockmarshalErrorf < i {
		condRecorderAuxMockmarshalErrorf.Wait()
	}
	condRecorderAuxMockmarshalErrorf.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalErrorf() {
	condRecorderAuxMockmarshalErrorf.L.Lock()
	recorderAuxMockmarshalErrorf++
	condRecorderAuxMockmarshalErrorf.L.Unlock()
	condRecorderAuxMockmarshalErrorf.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalErrorf() (ret int) {
	condRecorderAuxMockmarshalErrorf.L.Lock()
	ret = recorderAuxMockmarshalErrorf
	condRecorderAuxMockmarshalErrorf.L.Unlock()
	return
}

// marshalErrorf - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalErrorf(argformat string, args ...interface{}) (reta MarshalError) {
	FuncAuxMockmarshalErrorf, ok := apomock.GetRegisteredFunc("gocql.marshalErrorf")
	if ok {
		reta = FuncAuxMockmarshalErrorf.(func(argformat string, args ...interface{}) (reta MarshalError))(argformat, args...)
	} else {
		panic("FuncAuxMockmarshalErrorf ")
	}
	AuxMockIncrementRecorderAuxMockmarshalErrorf()
	return
}

//
// Mock: (recvm UnmarshalError)Error()(reta string)
//

type MockArgsTypeUnmarshalErrorError struct {
	ApomockCallNumber int
}

var LastMockArgsUnmarshalErrorError MockArgsTypeUnmarshalErrorError

// (recvm UnmarshalError)AuxMockError()(reta string) - Generated mock function
func (recvm UnmarshalError) AuxMockError() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.UnmarshalError.Error")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UnmarshalError.Error")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UnmarshalError.Error")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockUnmarshalErrorError  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUnmarshalErrorError int = 0

var condRecorderAuxMockUnmarshalErrorError *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUnmarshalErrorError(i int) {
	condRecorderAuxMockUnmarshalErrorError.L.Lock()
	for recorderAuxMockUnmarshalErrorError < i {
		condRecorderAuxMockUnmarshalErrorError.Wait()
	}
	condRecorderAuxMockUnmarshalErrorError.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUnmarshalErrorError() {
	condRecorderAuxMockUnmarshalErrorError.L.Lock()
	recorderAuxMockUnmarshalErrorError++
	condRecorderAuxMockUnmarshalErrorError.L.Unlock()
	condRecorderAuxMockUnmarshalErrorError.Broadcast()
}
func AuxMockGetRecorderAuxMockUnmarshalErrorError() (ret int) {
	condRecorderAuxMockUnmarshalErrorError.L.Lock()
	ret = recorderAuxMockUnmarshalErrorError
	condRecorderAuxMockUnmarshalErrorError.L.Unlock()
	return
}

// (recvm UnmarshalError)Error - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvm UnmarshalError) Error() (reta string) {
	FuncAuxMockUnmarshalErrorError, ok := apomock.GetRegisteredFunc("gocql.UnmarshalError.Error")
	if ok {
		reta = FuncAuxMockUnmarshalErrorError.(func(recvm UnmarshalError) (reta string))(recvm)
	} else {
		panic("FuncAuxMockUnmarshalErrorError ")
	}
	AuxMockIncrementRecorderAuxMockUnmarshalErrorError()
	return
}

//
// Mock: Marshal(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypeMarshal struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsMarshal MockArgsTypeMarshal

// AuxMockMarshal(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockMarshal(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsMarshal = MockArgsTypeMarshal{
		ApomockCallNumber: AuxMockGetRecorderAuxMockMarshal(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.Marshal")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Marshal")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.Marshal")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockMarshal  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockMarshal int = 0

var condRecorderAuxMockMarshal *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockMarshal(i int) {
	condRecorderAuxMockMarshal.L.Lock()
	for recorderAuxMockMarshal < i {
		condRecorderAuxMockMarshal.Wait()
	}
	condRecorderAuxMockMarshal.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockMarshal() {
	condRecorderAuxMockMarshal.L.Lock()
	recorderAuxMockMarshal++
	condRecorderAuxMockMarshal.L.Unlock()
	condRecorderAuxMockMarshal.Broadcast()
}
func AuxMockGetRecorderAuxMockMarshal() (ret int) {
	condRecorderAuxMockMarshal.L.Lock()
	ret = recorderAuxMockMarshal
	condRecorderAuxMockMarshal.L.Unlock()
	return
}

// Marshal - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func Marshal(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockMarshal, ok := apomock.GetRegisteredFunc("gocql.Marshal")
	if ok {
		reta, retb = FuncAuxMockMarshal.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockMarshal ")
	}
	AuxMockIncrementRecorderAuxMockMarshal()
	return
}

//
// Mock: marshalDecimal(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalDecimal struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalDecimal MockArgsTypemarshalDecimal

// AuxMockmarshalDecimal(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalDecimal(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalDecimal = MockArgsTypemarshalDecimal{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalDecimal(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalDecimal")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalDecimal")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalDecimal")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalDecimal  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalDecimal int = 0

var condRecorderAuxMockmarshalDecimal *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalDecimal(i int) {
	condRecorderAuxMockmarshalDecimal.L.Lock()
	for recorderAuxMockmarshalDecimal < i {
		condRecorderAuxMockmarshalDecimal.Wait()
	}
	condRecorderAuxMockmarshalDecimal.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalDecimal() {
	condRecorderAuxMockmarshalDecimal.L.Lock()
	recorderAuxMockmarshalDecimal++
	condRecorderAuxMockmarshalDecimal.L.Unlock()
	condRecorderAuxMockmarshalDecimal.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalDecimal() (ret int) {
	condRecorderAuxMockmarshalDecimal.L.Lock()
	ret = recorderAuxMockmarshalDecimal
	condRecorderAuxMockmarshalDecimal.L.Unlock()
	return
}

// marshalDecimal - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalDecimal(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalDecimal, ok := apomock.GetRegisteredFunc("gocql.marshalDecimal")
	if ok {
		reta, retb = FuncAuxMockmarshalDecimal.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalDecimal ")
	}
	AuxMockIncrementRecorderAuxMockmarshalDecimal()
	return
}

//
// Mock: (recvt CollectionType)New()(reta interface{})
//

type MockArgsTypeCollectionTypeNew struct {
	ApomockCallNumber int
}

var LastMockArgsCollectionTypeNew MockArgsTypeCollectionTypeNew

// (recvt CollectionType)AuxMockNew()(reta interface{}) - Generated mock function
func (recvt CollectionType) AuxMockNew() (reta interface{}) {
	rargs, rerr := apomock.GetNext("gocql.CollectionType.New")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.CollectionType.New")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.CollectionType.New")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(interface{})
	}
	return
}

// RecorderAuxMockCollectionTypeNew  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockCollectionTypeNew int = 0

var condRecorderAuxMockCollectionTypeNew *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockCollectionTypeNew(i int) {
	condRecorderAuxMockCollectionTypeNew.L.Lock()
	for recorderAuxMockCollectionTypeNew < i {
		condRecorderAuxMockCollectionTypeNew.Wait()
	}
	condRecorderAuxMockCollectionTypeNew.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockCollectionTypeNew() {
	condRecorderAuxMockCollectionTypeNew.L.Lock()
	recorderAuxMockCollectionTypeNew++
	condRecorderAuxMockCollectionTypeNew.L.Unlock()
	condRecorderAuxMockCollectionTypeNew.Broadcast()
}
func AuxMockGetRecorderAuxMockCollectionTypeNew() (ret int) {
	condRecorderAuxMockCollectionTypeNew.L.Lock()
	ret = recorderAuxMockCollectionTypeNew
	condRecorderAuxMockCollectionTypeNew.L.Unlock()
	return
}

// (recvt CollectionType)New - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt CollectionType) New() (reta interface{}) {
	FuncAuxMockCollectionTypeNew, ok := apomock.GetRegisteredFunc("gocql.CollectionType.New")
	if ok {
		reta = FuncAuxMockCollectionTypeNew.(func(recvt CollectionType) (reta interface{}))(recvt)
	} else {
		panic("FuncAuxMockCollectionTypeNew ")
	}
	AuxMockIncrementRecorderAuxMockCollectionTypeNew()
	return
}

//
// Mock: unmarshalInt(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error)
//

type MockArgsTypeunmarshalInt struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argdata           []byte
	Argvalue          interface{}
}

var LastMockArgsunmarshalInt MockArgsTypeunmarshalInt

// AuxMockunmarshalInt(arginfo TypeInfo, argdata []byte, argvalue interface{})(reta error) - Generated mock function
func AuxMockunmarshalInt(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	LastMockArgsunmarshalInt = MockArgsTypeunmarshalInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockunmarshalInt(),
		Arginfo:           arginfo,
		Argdata:           argdata,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.unmarshalInt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.unmarshalInt")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.unmarshalInt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockunmarshalInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockunmarshalInt int = 0

var condRecorderAuxMockunmarshalInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockunmarshalInt(i int) {
	condRecorderAuxMockunmarshalInt.L.Lock()
	for recorderAuxMockunmarshalInt < i {
		condRecorderAuxMockunmarshalInt.Wait()
	}
	condRecorderAuxMockunmarshalInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockunmarshalInt() {
	condRecorderAuxMockunmarshalInt.L.Lock()
	recorderAuxMockunmarshalInt++
	condRecorderAuxMockunmarshalInt.L.Unlock()
	condRecorderAuxMockunmarshalInt.Broadcast()
}
func AuxMockGetRecorderAuxMockunmarshalInt() (ret int) {
	condRecorderAuxMockunmarshalInt.L.Lock()
	ret = recorderAuxMockunmarshalInt
	condRecorderAuxMockunmarshalInt.L.Unlock()
	return
}

// unmarshalInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func unmarshalInt(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error) {
	FuncAuxMockunmarshalInt, ok := apomock.GetRegisteredFunc("gocql.unmarshalInt")
	if ok {
		reta = FuncAuxMockunmarshalInt.(func(arginfo TypeInfo, argdata []byte, argvalue interface{}) (reta error))(arginfo, argdata, argvalue)
	} else {
		panic("FuncAuxMockunmarshalInt ")
	}
	AuxMockIncrementRecorderAuxMockunmarshalInt()
	return
}

//
// Mock: marshalFloat(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error)
//

type MockArgsTypemarshalFloat struct {
	ApomockCallNumber int
	Arginfo           TypeInfo
	Argvalue          interface{}
}

var LastMockArgsmarshalFloat MockArgsTypemarshalFloat

// AuxMockmarshalFloat(arginfo TypeInfo, argvalue interface{})(reta []byte, retb error) - Generated mock function
func AuxMockmarshalFloat(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	LastMockArgsmarshalFloat = MockArgsTypemarshalFloat{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmarshalFloat(),
		Arginfo:           arginfo,
		Argvalue:          argvalue,
	}
	rargs, rerr := apomock.GetNext("gocql.marshalFloat")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.marshalFloat")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.marshalFloat")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockmarshalFloat  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmarshalFloat int = 0

var condRecorderAuxMockmarshalFloat *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmarshalFloat(i int) {
	condRecorderAuxMockmarshalFloat.L.Lock()
	for recorderAuxMockmarshalFloat < i {
		condRecorderAuxMockmarshalFloat.Wait()
	}
	condRecorderAuxMockmarshalFloat.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmarshalFloat() {
	condRecorderAuxMockmarshalFloat.L.Lock()
	recorderAuxMockmarshalFloat++
	condRecorderAuxMockmarshalFloat.L.Unlock()
	condRecorderAuxMockmarshalFloat.Broadcast()
}
func AuxMockGetRecorderAuxMockmarshalFloat() (ret int) {
	condRecorderAuxMockmarshalFloat.L.Lock()
	ret = recorderAuxMockmarshalFloat
	condRecorderAuxMockmarshalFloat.L.Unlock()
	return
}

// marshalFloat - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func marshalFloat(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error) {
	FuncAuxMockmarshalFloat, ok := apomock.GetRegisteredFunc("gocql.marshalFloat")
	if ok {
		reta, retb = FuncAuxMockmarshalFloat.(func(arginfo TypeInfo, argvalue interface{}) (reta []byte, retb error))(arginfo, argvalue)
	} else {
		panic("FuncAuxMockmarshalFloat ")
	}
	AuxMockIncrementRecorderAuxMockmarshalFloat()
	return
}

//
// Mock: (recvs NativeType)Type()(reta Type)
//

type MockArgsTypeNativeTypeType struct {
	ApomockCallNumber int
}

var LastMockArgsNativeTypeType MockArgsTypeNativeTypeType

// (recvs NativeType)AuxMockType()(reta Type) - Generated mock function
func (recvs NativeType) AuxMockType() (reta Type) {
	rargs, rerr := apomock.GetNext("gocql.NativeType.Type")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.NativeType.Type")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.NativeType.Type")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(Type)
	}
	return
}

// RecorderAuxMockNativeTypeType  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNativeTypeType int = 0

var condRecorderAuxMockNativeTypeType *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNativeTypeType(i int) {
	condRecorderAuxMockNativeTypeType.L.Lock()
	for recorderAuxMockNativeTypeType < i {
		condRecorderAuxMockNativeTypeType.Wait()
	}
	condRecorderAuxMockNativeTypeType.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNativeTypeType() {
	condRecorderAuxMockNativeTypeType.L.Lock()
	recorderAuxMockNativeTypeType++
	condRecorderAuxMockNativeTypeType.L.Unlock()
	condRecorderAuxMockNativeTypeType.Broadcast()
}
func AuxMockGetRecorderAuxMockNativeTypeType() (ret int) {
	condRecorderAuxMockNativeTypeType.L.Lock()
	ret = recorderAuxMockNativeTypeType
	condRecorderAuxMockNativeTypeType.L.Unlock()
	return
}

// (recvs NativeType)Type - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs NativeType) Type() (reta Type) {
	FuncAuxMockNativeTypeType, ok := apomock.GetRegisteredFunc("gocql.NativeType.Type")
	if ok {
		reta = FuncAuxMockNativeTypeType.(func(recvs NativeType) (reta Type))(recvs)
	} else {
		panic("FuncAuxMockNativeTypeType ")
	}
	AuxMockIncrementRecorderAuxMockNativeTypeType()
	return
}
