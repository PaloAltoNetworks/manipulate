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

	apomock.RegisterFunc("cassandra", "cassandra.isValidTag", AuxMockisValidTag)
	apomock.RegisterFunc("cassandra", "cassandra.tagOptions.Contains", (tagOptions).AuxMockContains)
	apomock.RegisterFunc("cassandra", "cassandra.tagForField", AuxMocktagForField)
	apomock.RegisterFunc("cassandra", "cassandra.parseTag", AuxMockparseTag)
}

const TagName = "cql"

const ()

//
// Internal Types: in this package and their exportable versions
//
type tagOptions string

//
// External Types: in this package
//

//
// Mock: isValidTag(args string)(reta bool)
//

type MockArgsTypeisValidTag struct {
	ApomockCallNumber int
	Args              string
}

var LastMockArgsisValidTag MockArgsTypeisValidTag

// AuxMockisValidTag(args string)(reta bool) - Generated mock function
func AuxMockisValidTag(args string) (reta bool) {
	LastMockArgsisValidTag = MockArgsTypeisValidTag{
		ApomockCallNumber: AuxMockGetRecorderAuxMockisValidTag(),
		Args:              args,
	}
	rargs, rerr := apomock.GetNext("cassandra.isValidTag")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.isValidTag")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.isValidTag")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockisValidTag  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockisValidTag int = 0

var condRecorderAuxMockisValidTag *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockisValidTag(i int) {
	condRecorderAuxMockisValidTag.L.Lock()
	for recorderAuxMockisValidTag < i {
		condRecorderAuxMockisValidTag.Wait()
	}
	condRecorderAuxMockisValidTag.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockisValidTag() {
	condRecorderAuxMockisValidTag.L.Lock()
	recorderAuxMockisValidTag++
	condRecorderAuxMockisValidTag.L.Unlock()
	condRecorderAuxMockisValidTag.Broadcast()
}
func AuxMockGetRecorderAuxMockisValidTag() (ret int) {
	condRecorderAuxMockisValidTag.L.Lock()
	ret = recorderAuxMockisValidTag
	condRecorderAuxMockisValidTag.L.Unlock()
	return
}

// isValidTag - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func isValidTag(args string) (reta bool) {
	FuncAuxMockisValidTag, ok := apomock.GetRegisteredFunc("cassandra.isValidTag")
	if ok {
		reta = FuncAuxMockisValidTag.(func(args string) (reta bool))(args)
	} else {
		panic("FuncAuxMockisValidTag ")
	}
	AuxMockIncrementRecorderAuxMockisValidTag()
	return
}

//
// Mock: (recvo tagOptions)Contains(argoptionName string)(reta bool)
//

type MockArgsTypetagOptionsContains struct {
	ApomockCallNumber int
	ArgoptionName     string
}

var LastMockArgstagOptionsContains MockArgsTypetagOptionsContains

// (recvo tagOptions)AuxMockContains(argoptionName string)(reta bool) - Generated mock function
func (recvo tagOptions) AuxMockContains(argoptionName string) (reta bool) {
	LastMockArgstagOptionsContains = MockArgsTypetagOptionsContains{
		ApomockCallNumber: AuxMockGetRecorderAuxMocktagOptionsContains(),
		ArgoptionName:     argoptionName,
	}
	rargs, rerr := apomock.GetNext("cassandra.tagOptions.Contains")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.tagOptions.Contains")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.tagOptions.Contains")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMocktagOptionsContains  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocktagOptionsContains int = 0

var condRecorderAuxMocktagOptionsContains *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocktagOptionsContains(i int) {
	condRecorderAuxMocktagOptionsContains.L.Lock()
	for recorderAuxMocktagOptionsContains < i {
		condRecorderAuxMocktagOptionsContains.Wait()
	}
	condRecorderAuxMocktagOptionsContains.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocktagOptionsContains() {
	condRecorderAuxMocktagOptionsContains.L.Lock()
	recorderAuxMocktagOptionsContains++
	condRecorderAuxMocktagOptionsContains.L.Unlock()
	condRecorderAuxMocktagOptionsContains.Broadcast()
}
func AuxMockGetRecorderAuxMocktagOptionsContains() (ret int) {
	condRecorderAuxMocktagOptionsContains.L.Lock()
	ret = recorderAuxMocktagOptionsContains
	condRecorderAuxMocktagOptionsContains.L.Unlock()
	return
}

// (recvo tagOptions)Contains - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvo tagOptions) Contains(argoptionName string) (reta bool) {
	FuncAuxMocktagOptionsContains, ok := apomock.GetRegisteredFunc("cassandra.tagOptions.Contains")
	if ok {
		reta = FuncAuxMocktagOptionsContains.(func(recvo tagOptions, argoptionName string) (reta bool))(recvo, argoptionName)
	} else {
		panic("FuncAuxMocktagOptionsContains ")
	}
	AuxMockIncrementRecorderAuxMocktagOptionsContains()
	return
}

//
// Mock: tagForField(argsf reflect.StructField)(reta string)
//

type MockArgsTypetagForField struct {
	ApomockCallNumber int
	Argsf             reflect.StructField
}

var LastMockArgstagForField MockArgsTypetagForField

// AuxMocktagForField(argsf reflect.StructField)(reta string) - Generated mock function
func AuxMocktagForField(argsf reflect.StructField) (reta string) {
	LastMockArgstagForField = MockArgsTypetagForField{
		ApomockCallNumber: AuxMockGetRecorderAuxMocktagForField(),
		Argsf:             argsf,
	}
	rargs, rerr := apomock.GetNext("cassandra.tagForField")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.tagForField")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.tagForField")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMocktagForField  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocktagForField int = 0

var condRecorderAuxMocktagForField *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocktagForField(i int) {
	condRecorderAuxMocktagForField.L.Lock()
	for recorderAuxMocktagForField < i {
		condRecorderAuxMocktagForField.Wait()
	}
	condRecorderAuxMocktagForField.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocktagForField() {
	condRecorderAuxMocktagForField.L.Lock()
	recorderAuxMocktagForField++
	condRecorderAuxMocktagForField.L.Unlock()
	condRecorderAuxMocktagForField.Broadcast()
}
func AuxMockGetRecorderAuxMocktagForField() (ret int) {
	condRecorderAuxMocktagForField.L.Lock()
	ret = recorderAuxMocktagForField
	condRecorderAuxMocktagForField.L.Unlock()
	return
}

// tagForField - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func tagForField(argsf reflect.StructField) (reta string) {
	FuncAuxMocktagForField, ok := apomock.GetRegisteredFunc("cassandra.tagForField")
	if ok {
		reta = FuncAuxMocktagForField.(func(argsf reflect.StructField) (reta string))(argsf)
	} else {
		panic("FuncAuxMocktagForField ")
	}
	AuxMockIncrementRecorderAuxMocktagForField()
	return
}

//
// Mock: parseTag(argtag string)(reta string, retb tagOptions)
//

type MockArgsTypeparseTag struct {
	ApomockCallNumber int
	Argtag            string
}

var LastMockArgsparseTag MockArgsTypeparseTag

// AuxMockparseTag(argtag string)(reta string, retb tagOptions) - Generated mock function
func AuxMockparseTag(argtag string) (reta string, retb tagOptions) {
	LastMockArgsparseTag = MockArgsTypeparseTag{
		ApomockCallNumber: AuxMockGetRecorderAuxMockparseTag(),
		Argtag:            argtag,
	}
	rargs, rerr := apomock.GetNext("cassandra.parseTag")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.parseTag")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:cassandra.parseTag")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(tagOptions)
	}
	return
}

// RecorderAuxMockparseTag  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockparseTag int = 0

var condRecorderAuxMockparseTag *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockparseTag(i int) {
	condRecorderAuxMockparseTag.L.Lock()
	for recorderAuxMockparseTag < i {
		condRecorderAuxMockparseTag.Wait()
	}
	condRecorderAuxMockparseTag.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockparseTag() {
	condRecorderAuxMockparseTag.L.Lock()
	recorderAuxMockparseTag++
	condRecorderAuxMockparseTag.L.Unlock()
	condRecorderAuxMockparseTag.Broadcast()
}
func AuxMockGetRecorderAuxMockparseTag() (ret int) {
	condRecorderAuxMockparseTag.L.Lock()
	ret = recorderAuxMockparseTag
	condRecorderAuxMockparseTag.L.Unlock()
	return
}

// parseTag - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func parseTag(argtag string) (reta string, retb tagOptions) {
	FuncAuxMockparseTag, ok := apomock.GetRegisteredFunc("cassandra.parseTag")
	if ok {
		reta, retb = FuncAuxMockparseTag.(func(argtag string) (reta string, retb tagOptions))(argtag)
	} else {
		panic("FuncAuxMockparseTag ")
	}
	AuxMockIncrementRecorderAuxMockparseTag()
	return
}
