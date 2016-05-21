// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package lru

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"
import "container/list"

func init() {
	apomock.RegisterInternalType("github.com/gocql/gocql/internal/lru", ApomockStructEntry, apomockNewStructEntry)

	apomock.RegisterFunc("lru", "lru.Cache.Add", (*Cache).AuxMockAdd)
	apomock.RegisterFunc("lru", "lru.Cache.Get", (*Cache).AuxMockGet)
	apomock.RegisterFunc("lru", "lru.Cache.Remove", (*Cache).AuxMockRemove)
	apomock.RegisterFunc("lru", "lru.Cache.RemoveOldest", (*Cache).AuxMockRemoveOldest)
	apomock.RegisterFunc("lru", "lru.Cache.removeElement", (*Cache).AuxMockremoveElement)
	apomock.RegisterFunc("lru", "lru.Cache.Len", (*Cache).AuxMockLen)
	apomock.RegisterFunc("lru", "lru.New", AuxMockNew)
}

const (
	ApomockStructEntry = 0
)

//
// Internal Types: in this package and their exportable versions
//
type entry struct {
	key   string
	value interface{}
}

//
// External Types: in this package
//
type Cache struct {
	MaxEntries int
	OnEvicted  func(key string, value interface{})
	ll         *list.List
	cache      map[string]*list.Element
}

func apomockNewStructEntry() interface{} { return &entry{} }

//
// Mock: (recvc *Cache)Add(argkey string, argvalue interface{})()
//

type MockArgsTypeCacheAdd struct {
	ApomockCallNumber int
	Argkey            string
	Argvalue          interface{}
}

var LastMockArgsCacheAdd MockArgsTypeCacheAdd

// (recvc *Cache)AuxMockAdd(argkey string, argvalue interface{})() - Generated mock function
func (recvc *Cache) AuxMockAdd(argkey string, argvalue interface{}) {
	LastMockArgsCacheAdd = MockArgsTypeCacheAdd{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrCacheAdd(),
		Argkey:            argkey,
		Argvalue:          argvalue,
	}
	return
}

// RecorderAuxMockPtrCacheAdd  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrCacheAdd int = 0

var condRecorderAuxMockPtrCacheAdd *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrCacheAdd(i int) {
	condRecorderAuxMockPtrCacheAdd.L.Lock()
	for recorderAuxMockPtrCacheAdd < i {
		condRecorderAuxMockPtrCacheAdd.Wait()
	}
	condRecorderAuxMockPtrCacheAdd.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrCacheAdd() {
	condRecorderAuxMockPtrCacheAdd.L.Lock()
	recorderAuxMockPtrCacheAdd++
	condRecorderAuxMockPtrCacheAdd.L.Unlock()
	condRecorderAuxMockPtrCacheAdd.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrCacheAdd() (ret int) {
	condRecorderAuxMockPtrCacheAdd.L.Lock()
	ret = recorderAuxMockPtrCacheAdd
	condRecorderAuxMockPtrCacheAdd.L.Unlock()
	return
}

// (recvc *Cache)Add - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Cache) Add(argkey string, argvalue interface{}) {
	FuncAuxMockPtrCacheAdd, ok := apomock.GetRegisteredFunc("lru.Cache.Add")
	if ok {
		FuncAuxMockPtrCacheAdd.(func(recvc *Cache, argkey string, argvalue interface{}))(recvc, argkey, argvalue)
	} else {
		panic("FuncAuxMockPtrCacheAdd ")
	}
	AuxMockIncrementRecorderAuxMockPtrCacheAdd()
	return
}

//
// Mock: (recvc *Cache)Get(argkey string)(retvalue interface{}, retok bool)
//

type MockArgsTypeCacheGet struct {
	ApomockCallNumber int
	Argkey            string
}

var LastMockArgsCacheGet MockArgsTypeCacheGet

// (recvc *Cache)AuxMockGet(argkey string)(retvalue interface{}, retok bool) - Generated mock function
func (recvc *Cache) AuxMockGet(argkey string) (retvalue interface{}, retok bool) {
	LastMockArgsCacheGet = MockArgsTypeCacheGet{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrCacheGet(),
		Argkey:            argkey,
	}
	rargs, rerr := apomock.GetNext("lru.Cache.Get")
	if rerr != nil {
		panic("Error getting next entry for method: lru.Cache.Get")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:lru.Cache.Get")
	}
	if rargs.GetArg(0) != nil {
		retvalue = rargs.GetArg(0).(interface{})
	}
	if rargs.GetArg(1) != nil {
		retok = rargs.GetArg(1).(bool)
	}
	return
}

// RecorderAuxMockPtrCacheGet  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrCacheGet int = 0

var condRecorderAuxMockPtrCacheGet *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrCacheGet(i int) {
	condRecorderAuxMockPtrCacheGet.L.Lock()
	for recorderAuxMockPtrCacheGet < i {
		condRecorderAuxMockPtrCacheGet.Wait()
	}
	condRecorderAuxMockPtrCacheGet.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrCacheGet() {
	condRecorderAuxMockPtrCacheGet.L.Lock()
	recorderAuxMockPtrCacheGet++
	condRecorderAuxMockPtrCacheGet.L.Unlock()
	condRecorderAuxMockPtrCacheGet.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrCacheGet() (ret int) {
	condRecorderAuxMockPtrCacheGet.L.Lock()
	ret = recorderAuxMockPtrCacheGet
	condRecorderAuxMockPtrCacheGet.L.Unlock()
	return
}

// (recvc *Cache)Get - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Cache) Get(argkey string) (retvalue interface{}, retok bool) {
	FuncAuxMockPtrCacheGet, ok := apomock.GetRegisteredFunc("lru.Cache.Get")
	if ok {
		retvalue, retok = FuncAuxMockPtrCacheGet.(func(recvc *Cache, argkey string) (retvalue interface{}, retok bool))(recvc, argkey)
	} else {
		panic("FuncAuxMockPtrCacheGet ")
	}
	AuxMockIncrementRecorderAuxMockPtrCacheGet()
	return
}

//
// Mock: (recvc *Cache)Remove(argkey string)(reta bool)
//

type MockArgsTypeCacheRemove struct {
	ApomockCallNumber int
	Argkey            string
}

var LastMockArgsCacheRemove MockArgsTypeCacheRemove

// (recvc *Cache)AuxMockRemove(argkey string)(reta bool) - Generated mock function
func (recvc *Cache) AuxMockRemove(argkey string) (reta bool) {
	LastMockArgsCacheRemove = MockArgsTypeCacheRemove{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrCacheRemove(),
		Argkey:            argkey,
	}
	rargs, rerr := apomock.GetNext("lru.Cache.Remove")
	if rerr != nil {
		panic("Error getting next entry for method: lru.Cache.Remove")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:lru.Cache.Remove")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrCacheRemove  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrCacheRemove int = 0

var condRecorderAuxMockPtrCacheRemove *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrCacheRemove(i int) {
	condRecorderAuxMockPtrCacheRemove.L.Lock()
	for recorderAuxMockPtrCacheRemove < i {
		condRecorderAuxMockPtrCacheRemove.Wait()
	}
	condRecorderAuxMockPtrCacheRemove.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrCacheRemove() {
	condRecorderAuxMockPtrCacheRemove.L.Lock()
	recorderAuxMockPtrCacheRemove++
	condRecorderAuxMockPtrCacheRemove.L.Unlock()
	condRecorderAuxMockPtrCacheRemove.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrCacheRemove() (ret int) {
	condRecorderAuxMockPtrCacheRemove.L.Lock()
	ret = recorderAuxMockPtrCacheRemove
	condRecorderAuxMockPtrCacheRemove.L.Unlock()
	return
}

// (recvc *Cache)Remove - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Cache) Remove(argkey string) (reta bool) {
	FuncAuxMockPtrCacheRemove, ok := apomock.GetRegisteredFunc("lru.Cache.Remove")
	if ok {
		reta = FuncAuxMockPtrCacheRemove.(func(recvc *Cache, argkey string) (reta bool))(recvc, argkey)
	} else {
		panic("FuncAuxMockPtrCacheRemove ")
	}
	AuxMockIncrementRecorderAuxMockPtrCacheRemove()
	return
}

//
// Mock: (recvc *Cache)RemoveOldest()()
//

type MockArgsTypeCacheRemoveOldest struct {
	ApomockCallNumber int
}

var LastMockArgsCacheRemoveOldest MockArgsTypeCacheRemoveOldest

// (recvc *Cache)AuxMockRemoveOldest()() - Generated mock function
func (recvc *Cache) AuxMockRemoveOldest() {
	return
}

// RecorderAuxMockPtrCacheRemoveOldest  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrCacheRemoveOldest int = 0

var condRecorderAuxMockPtrCacheRemoveOldest *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrCacheRemoveOldest(i int) {
	condRecorderAuxMockPtrCacheRemoveOldest.L.Lock()
	for recorderAuxMockPtrCacheRemoveOldest < i {
		condRecorderAuxMockPtrCacheRemoveOldest.Wait()
	}
	condRecorderAuxMockPtrCacheRemoveOldest.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrCacheRemoveOldest() {
	condRecorderAuxMockPtrCacheRemoveOldest.L.Lock()
	recorderAuxMockPtrCacheRemoveOldest++
	condRecorderAuxMockPtrCacheRemoveOldest.L.Unlock()
	condRecorderAuxMockPtrCacheRemoveOldest.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrCacheRemoveOldest() (ret int) {
	condRecorderAuxMockPtrCacheRemoveOldest.L.Lock()
	ret = recorderAuxMockPtrCacheRemoveOldest
	condRecorderAuxMockPtrCacheRemoveOldest.L.Unlock()
	return
}

// (recvc *Cache)RemoveOldest - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Cache) RemoveOldest() {
	FuncAuxMockPtrCacheRemoveOldest, ok := apomock.GetRegisteredFunc("lru.Cache.RemoveOldest")
	if ok {
		FuncAuxMockPtrCacheRemoveOldest.(func(recvc *Cache))(recvc)
	} else {
		panic("FuncAuxMockPtrCacheRemoveOldest ")
	}
	AuxMockIncrementRecorderAuxMockPtrCacheRemoveOldest()
	return
}

//
// Mock: (recvc *Cache)removeElement(arge *list.Element)()
//

type MockArgsTypeCacheremoveElement struct {
	ApomockCallNumber int
	Arge              *list.Element
}

var LastMockArgsCacheremoveElement MockArgsTypeCacheremoveElement

// (recvc *Cache)AuxMockremoveElement(arge *list.Element)() - Generated mock function
func (recvc *Cache) AuxMockremoveElement(arge *list.Element) {
	LastMockArgsCacheremoveElement = MockArgsTypeCacheremoveElement{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrCacheremoveElement(),
		Arge:              arge,
	}
	return
}

// RecorderAuxMockPtrCacheremoveElement  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrCacheremoveElement int = 0

var condRecorderAuxMockPtrCacheremoveElement *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrCacheremoveElement(i int) {
	condRecorderAuxMockPtrCacheremoveElement.L.Lock()
	for recorderAuxMockPtrCacheremoveElement < i {
		condRecorderAuxMockPtrCacheremoveElement.Wait()
	}
	condRecorderAuxMockPtrCacheremoveElement.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrCacheremoveElement() {
	condRecorderAuxMockPtrCacheremoveElement.L.Lock()
	recorderAuxMockPtrCacheremoveElement++
	condRecorderAuxMockPtrCacheremoveElement.L.Unlock()
	condRecorderAuxMockPtrCacheremoveElement.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrCacheremoveElement() (ret int) {
	condRecorderAuxMockPtrCacheremoveElement.L.Lock()
	ret = recorderAuxMockPtrCacheremoveElement
	condRecorderAuxMockPtrCacheremoveElement.L.Unlock()
	return
}

// (recvc *Cache)removeElement - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Cache) removeElement(arge *list.Element) {
	FuncAuxMockPtrCacheremoveElement, ok := apomock.GetRegisteredFunc("lru.Cache.removeElement")
	if ok {
		FuncAuxMockPtrCacheremoveElement.(func(recvc *Cache, arge *list.Element))(recvc, arge)
	} else {
		panic("FuncAuxMockPtrCacheremoveElement ")
	}
	AuxMockIncrementRecorderAuxMockPtrCacheremoveElement()
	return
}

//
// Mock: (recvc *Cache)Len()(reta int)
//

type MockArgsTypeCacheLen struct {
	ApomockCallNumber int
}

var LastMockArgsCacheLen MockArgsTypeCacheLen

// (recvc *Cache)AuxMockLen()(reta int) - Generated mock function
func (recvc *Cache) AuxMockLen() (reta int) {
	rargs, rerr := apomock.GetNext("lru.Cache.Len")
	if rerr != nil {
		panic("Error getting next entry for method: lru.Cache.Len")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:lru.Cache.Len")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrCacheLen  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrCacheLen int = 0

var condRecorderAuxMockPtrCacheLen *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrCacheLen(i int) {
	condRecorderAuxMockPtrCacheLen.L.Lock()
	for recorderAuxMockPtrCacheLen < i {
		condRecorderAuxMockPtrCacheLen.Wait()
	}
	condRecorderAuxMockPtrCacheLen.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrCacheLen() {
	condRecorderAuxMockPtrCacheLen.L.Lock()
	recorderAuxMockPtrCacheLen++
	condRecorderAuxMockPtrCacheLen.L.Unlock()
	condRecorderAuxMockPtrCacheLen.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrCacheLen() (ret int) {
	condRecorderAuxMockPtrCacheLen.L.Lock()
	ret = recorderAuxMockPtrCacheLen
	condRecorderAuxMockPtrCacheLen.L.Unlock()
	return
}

// (recvc *Cache)Len - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc *Cache) Len() (reta int) {
	FuncAuxMockPtrCacheLen, ok := apomock.GetRegisteredFunc("lru.Cache.Len")
	if ok {
		reta = FuncAuxMockPtrCacheLen.(func(recvc *Cache) (reta int))(recvc)
	} else {
		panic("FuncAuxMockPtrCacheLen ")
	}
	AuxMockIncrementRecorderAuxMockPtrCacheLen()
	return
}

//
// Mock: New(argmaxEntries int)(reta *Cache)
//

type MockArgsTypeNew struct {
	ApomockCallNumber int
	ArgmaxEntries     int
}

var LastMockArgsNew MockArgsTypeNew

// AuxMockNew(argmaxEntries int)(reta *Cache) - Generated mock function
func AuxMockNew(argmaxEntries int) (reta *Cache) {
	LastMockArgsNew = MockArgsTypeNew{
		ApomockCallNumber: AuxMockGetRecorderAuxMockNew(),
		ArgmaxEntries:     argmaxEntries,
	}
	rargs, rerr := apomock.GetNext("lru.New")
	if rerr != nil {
		panic("Error getting next entry for method: lru.New")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:lru.New")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*Cache)
	}
	return
}

// RecorderAuxMockNew  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockNew int = 0

var condRecorderAuxMockNew *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockNew(i int) {
	condRecorderAuxMockNew.L.Lock()
	for recorderAuxMockNew < i {
		condRecorderAuxMockNew.Wait()
	}
	condRecorderAuxMockNew.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockNew() {
	condRecorderAuxMockNew.L.Lock()
	recorderAuxMockNew++
	condRecorderAuxMockNew.L.Unlock()
	condRecorderAuxMockNew.Broadcast()
}
func AuxMockGetRecorderAuxMockNew() (ret int) {
	condRecorderAuxMockNew.L.Lock()
	ret = recorderAuxMockNew
	condRecorderAuxMockNew.L.Unlock()
	return
}

// New - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func New(argmaxEntries int) (reta *Cache) {
	FuncAuxMockNew, ok := apomock.GetRegisteredFunc("lru.New")
	if ok {
		reta = FuncAuxMockNew.(func(argmaxEntries int) (reta *Cache))(argmaxEntries)
	} else {
		panic("FuncAuxMockNew ")
	}
	AuxMockIncrementRecorderAuxMockNew()
	return
}
