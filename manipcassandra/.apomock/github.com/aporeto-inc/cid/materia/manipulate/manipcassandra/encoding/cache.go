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
	apomock.RegisterInternalType("github.com/aporeto-inc/cid/materia/manipulate/manipcassandra/encoding", ApomockStructCache, apomockNewStructCache)
	apomock.RegisterInternalType("github.com/aporeto-inc/cid/materia/manipulate/manipcassandra/encoding", ApomockStructField, apomockNewStructField)

	apomock.RegisterFunc("cassandra", "cassandra.cachedTypeFields", AuxMockcachedTypeFields)
	apomock.RegisterFunc("cassandra", "cassandra.byName.Less", (byName).AuxMockLess)
	apomock.RegisterFunc("cassandra", "cassandra.byIndex.Len", (byIndex).AuxMockLen)
	apomock.RegisterFunc("cassandra", "cassandra.typeFields", AuxMocktypeFields)
	apomock.RegisterFunc("cassandra", "cassandra.dominantField", AuxMockdominantField)
	apomock.RegisterFunc("cassandra", "cassandra.byName.Len", (byName).AuxMockLen)
	apomock.RegisterFunc("cassandra", "cassandra.byName.Swap", (byName).AuxMockSwap)
	apomock.RegisterFunc("cassandra", "cassandra.byIndex.Swap", (byIndex).AuxMockSwap)
	apomock.RegisterFunc("cassandra", "cassandra.byIndex.Less", (byIndex).AuxMockLess)
}

const (
	ApomockStructCache = 0
	ApomockStructField = 1
)

var fieldCache cache

//
// Internal Types: in this package and their exportable versions
//
type cache struct {
	sync.RWMutex
	m map[reflect.Type][]field
}
type field struct {
	name                  string
	tag                   bool
	index                 []int
	typ                   reflect.Type
	omitEmpty             bool
	quoted                bool
	autoTimestamp         bool
	autoTimestampOverride bool
	isPrimaryKey          bool
}
type byName []field
type byIndex []field

//
// External Types: in this package
//

func apomockNewStructCache() interface{} { return &cache{} }
func apomockNewStructField() interface{} { return &field{} }

//
// Mock: cachedTypeFields(argt reflect.Type)(reta []field)
//

type MockArgsTypecachedTypeFields struct {
	ApomockCallNumber int
	Argt              reflect.Type
}

var LastMockArgscachedTypeFields MockArgsTypecachedTypeFields

// AuxMockcachedTypeFields(argt reflect.Type)(reta []field) - Generated mock function
func AuxMockcachedTypeFields(argt reflect.Type) (reta []field) {
	LastMockArgscachedTypeFields = MockArgsTypecachedTypeFields{
		ApomockCallNumber: AuxMockGetRecorderAuxMockcachedTypeFields(),
		Argt:              argt,
	}
	rargs, rerr := apomock.GetNext("cassandra.cachedTypeFields")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.cachedTypeFields")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.cachedTypeFields")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]field)
	}
	return
}

// RecorderAuxMockcachedTypeFields  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockcachedTypeFields int = 0

var condRecorderAuxMockcachedTypeFields *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockcachedTypeFields(i int) {
	condRecorderAuxMockcachedTypeFields.L.Lock()
	for recorderAuxMockcachedTypeFields < i {
		condRecorderAuxMockcachedTypeFields.Wait()
	}
	condRecorderAuxMockcachedTypeFields.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockcachedTypeFields() {
	condRecorderAuxMockcachedTypeFields.L.Lock()
	recorderAuxMockcachedTypeFields++
	condRecorderAuxMockcachedTypeFields.L.Unlock()
	condRecorderAuxMockcachedTypeFields.Broadcast()
}
func AuxMockGetRecorderAuxMockcachedTypeFields() (ret int) {
	condRecorderAuxMockcachedTypeFields.L.Lock()
	ret = recorderAuxMockcachedTypeFields
	condRecorderAuxMockcachedTypeFields.L.Unlock()
	return
}

// cachedTypeFields - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func cachedTypeFields(argt reflect.Type) (reta []field) {
	FuncAuxMockcachedTypeFields, ok := apomock.GetRegisteredFunc("cassandra.cachedTypeFields")
	if ok {
		reta = FuncAuxMockcachedTypeFields.(func(argt reflect.Type) (reta []field))(argt)
	} else {
		panic("FuncAuxMockcachedTypeFields ")
	}
	AuxMockIncrementRecorderAuxMockcachedTypeFields()
	return
}

//
// Mock: (recvx byName)Less(argi int, argj int)(reta bool)
//

type MockArgsTypebyNameLess struct {
	ApomockCallNumber int
	Argi              int
	Argj              int
}

var LastMockArgsbyNameLess MockArgsTypebyNameLess

// (recvx byName)AuxMockLess(argi int, argj int)(reta bool) - Generated mock function
func (recvx byName) AuxMockLess(argi int, argj int) (reta bool) {
	LastMockArgsbyNameLess = MockArgsTypebyNameLess{
		ApomockCallNumber: AuxMockGetRecorderAuxMockbyNameLess(),
		Argi:              argi,
		Argj:              argj,
	}
	rargs, rerr := apomock.GetNext("cassandra.byName.Less")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.byName.Less")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.byName.Less")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockbyNameLess  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockbyNameLess int = 0

var condRecorderAuxMockbyNameLess *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockbyNameLess(i int) {
	condRecorderAuxMockbyNameLess.L.Lock()
	for recorderAuxMockbyNameLess < i {
		condRecorderAuxMockbyNameLess.Wait()
	}
	condRecorderAuxMockbyNameLess.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockbyNameLess() {
	condRecorderAuxMockbyNameLess.L.Lock()
	recorderAuxMockbyNameLess++
	condRecorderAuxMockbyNameLess.L.Unlock()
	condRecorderAuxMockbyNameLess.Broadcast()
}
func AuxMockGetRecorderAuxMockbyNameLess() (ret int) {
	condRecorderAuxMockbyNameLess.L.Lock()
	ret = recorderAuxMockbyNameLess
	condRecorderAuxMockbyNameLess.L.Unlock()
	return
}

// (recvx byName)Less - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvx byName) Less(argi int, argj int) (reta bool) {
	FuncAuxMockbyNameLess, ok := apomock.GetRegisteredFunc("cassandra.byName.Less")
	if ok {
		reta = FuncAuxMockbyNameLess.(func(recvx byName, argi int, argj int) (reta bool))(recvx, argi, argj)
	} else {
		panic("FuncAuxMockbyNameLess ")
	}
	AuxMockIncrementRecorderAuxMockbyNameLess()
	return
}

//
// Mock: (recvx byIndex)Len()(reta int)
//

type MockArgsTypebyIndexLen struct {
	ApomockCallNumber int
}

var LastMockArgsbyIndexLen MockArgsTypebyIndexLen

// (recvx byIndex)AuxMockLen()(reta int) - Generated mock function
func (recvx byIndex) AuxMockLen() (reta int) {
	rargs, rerr := apomock.GetNext("cassandra.byIndex.Len")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.byIndex.Len")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.byIndex.Len")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockbyIndexLen  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockbyIndexLen int = 0

var condRecorderAuxMockbyIndexLen *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockbyIndexLen(i int) {
	condRecorderAuxMockbyIndexLen.L.Lock()
	for recorderAuxMockbyIndexLen < i {
		condRecorderAuxMockbyIndexLen.Wait()
	}
	condRecorderAuxMockbyIndexLen.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockbyIndexLen() {
	condRecorderAuxMockbyIndexLen.L.Lock()
	recorderAuxMockbyIndexLen++
	condRecorderAuxMockbyIndexLen.L.Unlock()
	condRecorderAuxMockbyIndexLen.Broadcast()
}
func AuxMockGetRecorderAuxMockbyIndexLen() (ret int) {
	condRecorderAuxMockbyIndexLen.L.Lock()
	ret = recorderAuxMockbyIndexLen
	condRecorderAuxMockbyIndexLen.L.Unlock()
	return
}

// (recvx byIndex)Len - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvx byIndex) Len() (reta int) {
	FuncAuxMockbyIndexLen, ok := apomock.GetRegisteredFunc("cassandra.byIndex.Len")
	if ok {
		reta = FuncAuxMockbyIndexLen.(func(recvx byIndex) (reta int))(recvx)
	} else {
		panic("FuncAuxMockbyIndexLen ")
	}
	AuxMockIncrementRecorderAuxMockbyIndexLen()
	return
}

//
// Mock: typeFields(argt reflect.Type)(reta []field)
//

type MockArgsTypetypeFields struct {
	ApomockCallNumber int
	Argt              reflect.Type
}

var LastMockArgstypeFields MockArgsTypetypeFields

// AuxMocktypeFields(argt reflect.Type)(reta []field) - Generated mock function
func AuxMocktypeFields(argt reflect.Type) (reta []field) {
	LastMockArgstypeFields = MockArgsTypetypeFields{
		ApomockCallNumber: AuxMockGetRecorderAuxMocktypeFields(),
		Argt:              argt,
	}
	rargs, rerr := apomock.GetNext("cassandra.typeFields")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.typeFields")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.typeFields")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]field)
	}
	return
}

// RecorderAuxMocktypeFields  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocktypeFields int = 0

var condRecorderAuxMocktypeFields *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocktypeFields(i int) {
	condRecorderAuxMocktypeFields.L.Lock()
	for recorderAuxMocktypeFields < i {
		condRecorderAuxMocktypeFields.Wait()
	}
	condRecorderAuxMocktypeFields.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocktypeFields() {
	condRecorderAuxMocktypeFields.L.Lock()
	recorderAuxMocktypeFields++
	condRecorderAuxMocktypeFields.L.Unlock()
	condRecorderAuxMocktypeFields.Broadcast()
}
func AuxMockGetRecorderAuxMocktypeFields() (ret int) {
	condRecorderAuxMocktypeFields.L.Lock()
	ret = recorderAuxMocktypeFields
	condRecorderAuxMocktypeFields.L.Unlock()
	return
}

// typeFields - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func typeFields(argt reflect.Type) (reta []field) {
	FuncAuxMocktypeFields, ok := apomock.GetRegisteredFunc("cassandra.typeFields")
	if ok {
		reta = FuncAuxMocktypeFields.(func(argt reflect.Type) (reta []field))(argt)
	} else {
		panic("FuncAuxMocktypeFields ")
	}
	AuxMockIncrementRecorderAuxMocktypeFields()
	return
}

//
// Mock: dominantField(argfields []field)(reta field, retb bool)
//

type MockArgsTypedominantField struct {
	ApomockCallNumber int
	Argfields         []field
}

var LastMockArgsdominantField MockArgsTypedominantField

// AuxMockdominantField(argfields []field)(reta field, retb bool) - Generated mock function
func AuxMockdominantField(argfields []field) (reta field, retb bool) {
	LastMockArgsdominantField = MockArgsTypedominantField{
		ApomockCallNumber: AuxMockGetRecorderAuxMockdominantField(),
		Argfields:         argfields,
	}
	rargs, rerr := apomock.GetNext("cassandra.dominantField")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.dominantField")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:cassandra.dominantField")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(field)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(bool)
	}
	return
}

// RecorderAuxMockdominantField  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockdominantField int = 0

var condRecorderAuxMockdominantField *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockdominantField(i int) {
	condRecorderAuxMockdominantField.L.Lock()
	for recorderAuxMockdominantField < i {
		condRecorderAuxMockdominantField.Wait()
	}
	condRecorderAuxMockdominantField.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockdominantField() {
	condRecorderAuxMockdominantField.L.Lock()
	recorderAuxMockdominantField++
	condRecorderAuxMockdominantField.L.Unlock()
	condRecorderAuxMockdominantField.Broadcast()
}
func AuxMockGetRecorderAuxMockdominantField() (ret int) {
	condRecorderAuxMockdominantField.L.Lock()
	ret = recorderAuxMockdominantField
	condRecorderAuxMockdominantField.L.Unlock()
	return
}

// dominantField - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func dominantField(argfields []field) (reta field, retb bool) {
	FuncAuxMockdominantField, ok := apomock.GetRegisteredFunc("cassandra.dominantField")
	if ok {
		reta, retb = FuncAuxMockdominantField.(func(argfields []field) (reta field, retb bool))(argfields)
	} else {
		panic("FuncAuxMockdominantField ")
	}
	AuxMockIncrementRecorderAuxMockdominantField()
	return
}

//
// Mock: (recvx byName)Len()(reta int)
//

type MockArgsTypebyNameLen struct {
	ApomockCallNumber int
}

var LastMockArgsbyNameLen MockArgsTypebyNameLen

// (recvx byName)AuxMockLen()(reta int) - Generated mock function
func (recvx byName) AuxMockLen() (reta int) {
	rargs, rerr := apomock.GetNext("cassandra.byName.Len")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.byName.Len")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.byName.Len")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockbyNameLen  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockbyNameLen int = 0

var condRecorderAuxMockbyNameLen *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockbyNameLen(i int) {
	condRecorderAuxMockbyNameLen.L.Lock()
	for recorderAuxMockbyNameLen < i {
		condRecorderAuxMockbyNameLen.Wait()
	}
	condRecorderAuxMockbyNameLen.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockbyNameLen() {
	condRecorderAuxMockbyNameLen.L.Lock()
	recorderAuxMockbyNameLen++
	condRecorderAuxMockbyNameLen.L.Unlock()
	condRecorderAuxMockbyNameLen.Broadcast()
}
func AuxMockGetRecorderAuxMockbyNameLen() (ret int) {
	condRecorderAuxMockbyNameLen.L.Lock()
	ret = recorderAuxMockbyNameLen
	condRecorderAuxMockbyNameLen.L.Unlock()
	return
}

// (recvx byName)Len - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvx byName) Len() (reta int) {
	FuncAuxMockbyNameLen, ok := apomock.GetRegisteredFunc("cassandra.byName.Len")
	if ok {
		reta = FuncAuxMockbyNameLen.(func(recvx byName) (reta int))(recvx)
	} else {
		panic("FuncAuxMockbyNameLen ")
	}
	AuxMockIncrementRecorderAuxMockbyNameLen()
	return
}

//
// Mock: (recvx byName)Swap(argi int, argj int)()
//

type MockArgsTypebyNameSwap struct {
	ApomockCallNumber int
	Argi              int
	Argj              int
}

var LastMockArgsbyNameSwap MockArgsTypebyNameSwap

// (recvx byName)AuxMockSwap(argi int, argj int)() - Generated mock function
func (recvx byName) AuxMockSwap(argi int, argj int) {
	LastMockArgsbyNameSwap = MockArgsTypebyNameSwap{
		ApomockCallNumber: AuxMockGetRecorderAuxMockbyNameSwap(),
		Argi:              argi,
		Argj:              argj,
	}
	return
}

// RecorderAuxMockbyNameSwap  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockbyNameSwap int = 0

var condRecorderAuxMockbyNameSwap *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockbyNameSwap(i int) {
	condRecorderAuxMockbyNameSwap.L.Lock()
	for recorderAuxMockbyNameSwap < i {
		condRecorderAuxMockbyNameSwap.Wait()
	}
	condRecorderAuxMockbyNameSwap.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockbyNameSwap() {
	condRecorderAuxMockbyNameSwap.L.Lock()
	recorderAuxMockbyNameSwap++
	condRecorderAuxMockbyNameSwap.L.Unlock()
	condRecorderAuxMockbyNameSwap.Broadcast()
}
func AuxMockGetRecorderAuxMockbyNameSwap() (ret int) {
	condRecorderAuxMockbyNameSwap.L.Lock()
	ret = recorderAuxMockbyNameSwap
	condRecorderAuxMockbyNameSwap.L.Unlock()
	return
}

// (recvx byName)Swap - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvx byName) Swap(argi int, argj int) {
	FuncAuxMockbyNameSwap, ok := apomock.GetRegisteredFunc("cassandra.byName.Swap")
	if ok {
		FuncAuxMockbyNameSwap.(func(recvx byName, argi int, argj int))(recvx, argi, argj)
	} else {
		panic("FuncAuxMockbyNameSwap ")
	}
	AuxMockIncrementRecorderAuxMockbyNameSwap()
	return
}

//
// Mock: (recvx byIndex)Swap(argi int, argj int)()
//

type MockArgsTypebyIndexSwap struct {
	ApomockCallNumber int
	Argi              int
	Argj              int
}

var LastMockArgsbyIndexSwap MockArgsTypebyIndexSwap

// (recvx byIndex)AuxMockSwap(argi int, argj int)() - Generated mock function
func (recvx byIndex) AuxMockSwap(argi int, argj int) {
	LastMockArgsbyIndexSwap = MockArgsTypebyIndexSwap{
		ApomockCallNumber: AuxMockGetRecorderAuxMockbyIndexSwap(),
		Argi:              argi,
		Argj:              argj,
	}
	return
}

// RecorderAuxMockbyIndexSwap  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockbyIndexSwap int = 0

var condRecorderAuxMockbyIndexSwap *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockbyIndexSwap(i int) {
	condRecorderAuxMockbyIndexSwap.L.Lock()
	for recorderAuxMockbyIndexSwap < i {
		condRecorderAuxMockbyIndexSwap.Wait()
	}
	condRecorderAuxMockbyIndexSwap.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockbyIndexSwap() {
	condRecorderAuxMockbyIndexSwap.L.Lock()
	recorderAuxMockbyIndexSwap++
	condRecorderAuxMockbyIndexSwap.L.Unlock()
	condRecorderAuxMockbyIndexSwap.Broadcast()
}
func AuxMockGetRecorderAuxMockbyIndexSwap() (ret int) {
	condRecorderAuxMockbyIndexSwap.L.Lock()
	ret = recorderAuxMockbyIndexSwap
	condRecorderAuxMockbyIndexSwap.L.Unlock()
	return
}

// (recvx byIndex)Swap - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvx byIndex) Swap(argi int, argj int) {
	FuncAuxMockbyIndexSwap, ok := apomock.GetRegisteredFunc("cassandra.byIndex.Swap")
	if ok {
		FuncAuxMockbyIndexSwap.(func(recvx byIndex, argi int, argj int))(recvx, argi, argj)
	} else {
		panic("FuncAuxMockbyIndexSwap ")
	}
	AuxMockIncrementRecorderAuxMockbyIndexSwap()
	return
}

//
// Mock: (recvx byIndex)Less(argi int, argj int)(reta bool)
//

type MockArgsTypebyIndexLess struct {
	ApomockCallNumber int
	Argi              int
	Argj              int
}

var LastMockArgsbyIndexLess MockArgsTypebyIndexLess

// (recvx byIndex)AuxMockLess(argi int, argj int)(reta bool) - Generated mock function
func (recvx byIndex) AuxMockLess(argi int, argj int) (reta bool) {
	LastMockArgsbyIndexLess = MockArgsTypebyIndexLess{
		ApomockCallNumber: AuxMockGetRecorderAuxMockbyIndexLess(),
		Argi:              argi,
		Argj:              argj,
	}
	rargs, rerr := apomock.GetNext("cassandra.byIndex.Less")
	if rerr != nil {
		panic("Error getting next entry for method: cassandra.byIndex.Less")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:cassandra.byIndex.Less")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockbyIndexLess  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockbyIndexLess int = 0

var condRecorderAuxMockbyIndexLess *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockbyIndexLess(i int) {
	condRecorderAuxMockbyIndexLess.L.Lock()
	for recorderAuxMockbyIndexLess < i {
		condRecorderAuxMockbyIndexLess.Wait()
	}
	condRecorderAuxMockbyIndexLess.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockbyIndexLess() {
	condRecorderAuxMockbyIndexLess.L.Lock()
	recorderAuxMockbyIndexLess++
	condRecorderAuxMockbyIndexLess.L.Unlock()
	condRecorderAuxMockbyIndexLess.Broadcast()
}
func AuxMockGetRecorderAuxMockbyIndexLess() (ret int) {
	condRecorderAuxMockbyIndexLess.L.Lock()
	ret = recorderAuxMockbyIndexLess
	condRecorderAuxMockbyIndexLess.L.Unlock()
	return
}

// (recvx byIndex)Less - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvx byIndex) Less(argi int, argj int) (reta bool) {
	FuncAuxMockbyIndexLess, ok := apomock.GetRegisteredFunc("cassandra.byIndex.Less")
	if ok {
		reta = FuncAuxMockbyIndexLess.(func(recvx byIndex, argi int, argj int) (reta bool))(recvx, argi, argj)
	} else {
		panic("FuncAuxMockbyIndexLess ")
	}
	AuxMockIncrementRecorderAuxMockbyIndexLess()
	return
}
