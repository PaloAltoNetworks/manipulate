// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package cassandra

import "reflect"
import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {

	apomock.RegisterFunc("cassandra", "cassandra.Marshal", AuxMockMarshal)
	apomock.RegisterFunc("cassandra", "cassandra.fieldByIndex", AuxMockfieldByIndex)
	apomock.RegisterFunc("cassandra", "cassandra.FieldsAndValues", AuxMockFieldsAndValues)
	apomock.RegisterFunc("cassandra", "cassandra.PrimaryFieldsAndValues", AuxMockPrimaryFieldsAndValues)
	apomock.RegisterFunc("cassandra", "cassandra.isEmptyValue", AuxMockisEmptyValue)
}

const ()

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//

//
// Mock: Marshal(argv interface{})(reta map[string]interface{}, retb error)
//

type MockArgsTypeMarshal struct {
	ApomockCallNumber int
	Argv              interface{}
}

var LastMockArgsMarshal MockArgsTypeMarshal

// AuxMockMarshal(argv interface{})(reta map[string]interface{}, retb error) - Generated mock function
func AuxMockMarshal(argv interface{}) (reta map[string]interface{}, retb error) {
	LastMockArgsMarshal = MockArgsTypeMarshal{
		ApomockCallNumber: AuxMockGetRecorderAuxMockMarshal(),
		Argv:              argv,
	}
	rargs, rerr := apomock.GetNext("cassandra.Marshal")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.Marshal")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:cassandra.Marshal")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(map[string]interface{})
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
func Marshal(argv interface{}) (reta map[string]interface{}, retb error) {
	FuncAuxMockMarshal, ok := apomock.GetRegisteredFunc("cassandra.Marshal")
	if ok {
		reta, retb = FuncAuxMockMarshal.(func(argv interface{}) (reta map[string]interface{}, retb error))(argv)
	} else {
		panic("FuncAuxMockMarshal ")
	}
	AuxMockIncrementRecorderAuxMockMarshal()
	return
}

//
// Mock: fieldByIndex(argv reflect.Value, argindex []int)(reta reflect.Value)
//

type MockArgsTypefieldByIndex struct {
	ApomockCallNumber int
	Argv              reflect.Value
	Argindex          []int
}

var LastMockArgsfieldByIndex MockArgsTypefieldByIndex

// AuxMockfieldByIndex(argv reflect.Value, argindex []int)(reta reflect.Value) - Generated mock function
func AuxMockfieldByIndex(argv reflect.Value, argindex []int) (reta reflect.Value) {
	LastMockArgsfieldByIndex = MockArgsTypefieldByIndex{
		ApomockCallNumber: AuxMockGetRecorderAuxMockfieldByIndex(),
		Argv:              argv,
		Argindex:          argindex,
	}
	rargs, rerr := apomock.GetNext("cassandra.fieldByIndex")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.fieldByIndex")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.fieldByIndex")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(reflect.Value)
	}
	return
}

// RecorderAuxMockfieldByIndex  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockfieldByIndex int = 0

var condRecorderAuxMockfieldByIndex *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockfieldByIndex(i int) {
	condRecorderAuxMockfieldByIndex.L.Lock()
	for recorderAuxMockfieldByIndex < i {
		condRecorderAuxMockfieldByIndex.Wait()
	}
	condRecorderAuxMockfieldByIndex.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockfieldByIndex() {
	condRecorderAuxMockfieldByIndex.L.Lock()
	recorderAuxMockfieldByIndex++
	condRecorderAuxMockfieldByIndex.L.Unlock()
	condRecorderAuxMockfieldByIndex.Broadcast()
}
func AuxMockGetRecorderAuxMockfieldByIndex() (ret int) {
	condRecorderAuxMockfieldByIndex.L.Lock()
	ret = recorderAuxMockfieldByIndex
	condRecorderAuxMockfieldByIndex.L.Unlock()
	return
}

// fieldByIndex - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func fieldByIndex(argv reflect.Value, argindex []int) (reta reflect.Value) {
	FuncAuxMockfieldByIndex, ok := apomock.GetRegisteredFunc("cassandra.fieldByIndex")
	if ok {
		reta = FuncAuxMockfieldByIndex.(func(argv reflect.Value, argindex []int) (reta reflect.Value))(argv, argindex)
	} else {
		panic("FuncAuxMockfieldByIndex ")
	}
	AuxMockIncrementRecorderAuxMockfieldByIndex()
	return
}

//
// Mock: FieldsAndValues(argval interface{})(reta []string, retb []interface{}, retc error)
//

type MockArgsTypeFieldsAndValues struct {
	ApomockCallNumber int
	Argval            interface{}
}

var LastMockArgsFieldsAndValues MockArgsTypeFieldsAndValues

// AuxMockFieldsAndValues(argval interface{})(reta []string, retb []interface{}, retc error) - Generated mock function
func AuxMockFieldsAndValues(argval interface{}) (reta []string, retb []interface{}, retc error) {
	LastMockArgsFieldsAndValues = MockArgsTypeFieldsAndValues{
		ApomockCallNumber: AuxMockGetRecorderAuxMockFieldsAndValues(),
		Argval:            argval,
	}
	rargs, rerr := apomock.GetNext("cassandra.FieldsAndValues")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.FieldsAndValues")
	} else if rargs.NumArgs() != 3 {
		panic("All return parameters not provided for method:cassandra.FieldsAndValues")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]string)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).([]interface{})
	}
	if rargs.GetArg(2) != nil {
		retc = rargs.GetArg(2).(error)
	}
	return
}

// RecorderAuxMockFieldsAndValues  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockFieldsAndValues int = 0

var condRecorderAuxMockFieldsAndValues *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockFieldsAndValues(i int) {
	condRecorderAuxMockFieldsAndValues.L.Lock()
	for recorderAuxMockFieldsAndValues < i {
		condRecorderAuxMockFieldsAndValues.Wait()
	}
	condRecorderAuxMockFieldsAndValues.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockFieldsAndValues() {
	condRecorderAuxMockFieldsAndValues.L.Lock()
	recorderAuxMockFieldsAndValues++
	condRecorderAuxMockFieldsAndValues.L.Unlock()
	condRecorderAuxMockFieldsAndValues.Broadcast()
}
func AuxMockGetRecorderAuxMockFieldsAndValues() (ret int) {
	condRecorderAuxMockFieldsAndValues.L.Lock()
	ret = recorderAuxMockFieldsAndValues
	condRecorderAuxMockFieldsAndValues.L.Unlock()
	return
}

// FieldsAndValues - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func FieldsAndValues(argval interface{}) (reta []string, retb []interface{}, retc error) {
	FuncAuxMockFieldsAndValues, ok := apomock.GetRegisteredFunc("cassandra.FieldsAndValues")
	if ok {
		reta, retb, retc = FuncAuxMockFieldsAndValues.(func(argval interface{}) (reta []string, retb []interface{}, retc error))(argval)
	} else {
		panic("FuncAuxMockFieldsAndValues ")
	}
	AuxMockIncrementRecorderAuxMockFieldsAndValues()
	return
}

//
// Mock: PrimaryFieldsAndValues(argval interface{})(reta []string, retb []interface{}, retc error)
//

type MockArgsTypePrimaryFieldsAndValues struct {
	ApomockCallNumber int
	Argval            interface{}
}

var LastMockArgsPrimaryFieldsAndValues MockArgsTypePrimaryFieldsAndValues

// AuxMockPrimaryFieldsAndValues(argval interface{})(reta []string, retb []interface{}, retc error) - Generated mock function
func AuxMockPrimaryFieldsAndValues(argval interface{}) (reta []string, retb []interface{}, retc error) {
	LastMockArgsPrimaryFieldsAndValues = MockArgsTypePrimaryFieldsAndValues{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPrimaryFieldsAndValues(),
		Argval:            argval,
	}
	rargs, rerr := apomock.GetNext("cassandra.PrimaryFieldsAndValues")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.PrimaryFieldsAndValues")
	} else if rargs.NumArgs() != 3 {
		panic("All return parameters not provided for method:cassandra.PrimaryFieldsAndValues")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]string)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).([]interface{})
	}
	if rargs.GetArg(2) != nil {
		retc = rargs.GetArg(2).(error)
	}
	return
}

// RecorderAuxMockPrimaryFieldsAndValues  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPrimaryFieldsAndValues int = 0

var condRecorderAuxMockPrimaryFieldsAndValues *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPrimaryFieldsAndValues(i int) {
	condRecorderAuxMockPrimaryFieldsAndValues.L.Lock()
	for recorderAuxMockPrimaryFieldsAndValues < i {
		condRecorderAuxMockPrimaryFieldsAndValues.Wait()
	}
	condRecorderAuxMockPrimaryFieldsAndValues.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPrimaryFieldsAndValues() {
	condRecorderAuxMockPrimaryFieldsAndValues.L.Lock()
	recorderAuxMockPrimaryFieldsAndValues++
	condRecorderAuxMockPrimaryFieldsAndValues.L.Unlock()
	condRecorderAuxMockPrimaryFieldsAndValues.Broadcast()
}
func AuxMockGetRecorderAuxMockPrimaryFieldsAndValues() (ret int) {
	condRecorderAuxMockPrimaryFieldsAndValues.L.Lock()
	ret = recorderAuxMockPrimaryFieldsAndValues
	condRecorderAuxMockPrimaryFieldsAndValues.L.Unlock()
	return
}

// PrimaryFieldsAndValues - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func PrimaryFieldsAndValues(argval interface{}) (reta []string, retb []interface{}, retc error) {
	FuncAuxMockPrimaryFieldsAndValues, ok := apomock.GetRegisteredFunc("cassandra.PrimaryFieldsAndValues")
	if ok {
		reta, retb, retc = FuncAuxMockPrimaryFieldsAndValues.(func(argval interface{}) (reta []string, retb []interface{}, retc error))(argval)
	} else {
		panic("FuncAuxMockPrimaryFieldsAndValues ")
	}
	AuxMockIncrementRecorderAuxMockPrimaryFieldsAndValues()
	return
}

//
// Mock: isEmptyValue(argv reflect.Value)(reta bool)
//

type MockArgsTypeisEmptyValue struct {
	ApomockCallNumber int
	Argv              reflect.Value
}

var LastMockArgsisEmptyValue MockArgsTypeisEmptyValue

// AuxMockisEmptyValue(argv reflect.Value)(reta bool) - Generated mock function
func AuxMockisEmptyValue(argv reflect.Value) (reta bool) {
	LastMockArgsisEmptyValue = MockArgsTypeisEmptyValue{
		ApomockCallNumber: AuxMockGetRecorderAuxMockisEmptyValue(),
		Argv:              argv,
	}
	rargs, rerr := apomock.GetNext("cassandra.isEmptyValue")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.isEmptyValue")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.isEmptyValue")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockisEmptyValue  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockisEmptyValue int = 0

var condRecorderAuxMockisEmptyValue *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockisEmptyValue(i int) {
	condRecorderAuxMockisEmptyValue.L.Lock()
	for recorderAuxMockisEmptyValue < i {
		condRecorderAuxMockisEmptyValue.Wait()
	}
	condRecorderAuxMockisEmptyValue.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockisEmptyValue() {
	condRecorderAuxMockisEmptyValue.L.Lock()
	recorderAuxMockisEmptyValue++
	condRecorderAuxMockisEmptyValue.L.Unlock()
	condRecorderAuxMockisEmptyValue.Broadcast()
}
func AuxMockGetRecorderAuxMockisEmptyValue() (ret int) {
	condRecorderAuxMockisEmptyValue.L.Lock()
	ret = recorderAuxMockisEmptyValue
	condRecorderAuxMockisEmptyValue.L.Unlock()
	return
}

// isEmptyValue - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func isEmptyValue(argv reflect.Value) (reta bool) {
	FuncAuxMockisEmptyValue, ok := apomock.GetRegisteredFunc("cassandra.isEmptyValue")
	if ok {
		reta = FuncAuxMockisEmptyValue.(func(argv reflect.Value) (reta bool))(argv)
	} else {
		panic("FuncAuxMockisEmptyValue ")
	}
	AuxMockIncrementRecorderAuxMockisEmptyValue()
	return
}
