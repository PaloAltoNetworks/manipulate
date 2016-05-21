// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package streams

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {

	apomock.RegisterFunc("streams", "streams.IDGenerator.Clear", (*IDGenerator).AuxMockClear)
	apomock.RegisterFunc("streams", "streams.bitfmt", AuxMockbitfmt)
	apomock.RegisterFunc("streams", "streams.bucketOffset", AuxMockbucketOffset)
	apomock.RegisterFunc("streams", "streams.IDGenerator.isSet", (*IDGenerator).AuxMockisSet)
	apomock.RegisterFunc("streams", "streams.IDGenerator.String", (*IDGenerator).AuxMockString)
	apomock.RegisterFunc("streams", "streams.isSet", AuxMockisSet)
	apomock.RegisterFunc("streams", "streams.IDGenerator.Available", (*IDGenerator).AuxMockAvailable)
	apomock.RegisterFunc("streams", "streams.New", AuxMockNew)
	apomock.RegisterFunc("streams", "streams.streamFromBucket", AuxMockstreamFromBucket)
	apomock.RegisterFunc("streams", "streams.IDGenerator.GetStream", (*IDGenerator).AuxMockGetStream)
	apomock.RegisterFunc("streams", "streams.streamOffset", AuxMockstreamOffset)
}

const bucketBits = 64

const ()

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//
type IDGenerator struct {
	NumStreams   int
	inuseStreams int32
	numBuckets   uint32
	streams      []uint64
	offset       uint32
}

//
// Mock: (recvs *IDGenerator)Clear(argstream int)(retinuse bool)
//

type MockArgsTypeIDGeneratorClear struct {
	ApomockCallNumber int
	Argstream         int
}

var LastMockArgsIDGeneratorClear MockArgsTypeIDGeneratorClear

// (recvs *IDGenerator)AuxMockClear(argstream int)(retinuse bool) - Generated mock function
func (recvs *IDGenerator) AuxMockClear(argstream int) (retinuse bool) {
	LastMockArgsIDGeneratorClear = MockArgsTypeIDGeneratorClear{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrIDGeneratorClear(),
		Argstream:         argstream,
	}
	rargs, rerr := apomock.GetNext("streams.IDGenerator.Clear")
	if rerr != nil {
		panic("Error getting next entry for method: streams.IDGenerator.Clear")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:streams.IDGenerator.Clear")
	}
	if rargs.GetArg(0) != nil {
		retinuse = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrIDGeneratorClear  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIDGeneratorClear int = 0

var condRecorderAuxMockPtrIDGeneratorClear *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIDGeneratorClear(i int) {
	condRecorderAuxMockPtrIDGeneratorClear.L.Lock()
	for recorderAuxMockPtrIDGeneratorClear < i {
		condRecorderAuxMockPtrIDGeneratorClear.Wait()
	}
	condRecorderAuxMockPtrIDGeneratorClear.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIDGeneratorClear() {
	condRecorderAuxMockPtrIDGeneratorClear.L.Lock()
	recorderAuxMockPtrIDGeneratorClear++
	condRecorderAuxMockPtrIDGeneratorClear.L.Unlock()
	condRecorderAuxMockPtrIDGeneratorClear.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIDGeneratorClear() (ret int) {
	condRecorderAuxMockPtrIDGeneratorClear.L.Lock()
	ret = recorderAuxMockPtrIDGeneratorClear
	condRecorderAuxMockPtrIDGeneratorClear.L.Unlock()
	return
}

// (recvs *IDGenerator)Clear - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *IDGenerator) Clear(argstream int) (retinuse bool) {
	FuncAuxMockPtrIDGeneratorClear, ok := apomock.GetRegisteredFunc("streams.IDGenerator.Clear")
	if ok {
		retinuse = FuncAuxMockPtrIDGeneratorClear.(func(recvs *IDGenerator, argstream int) (retinuse bool))(recvs, argstream)
	} else {
		panic("FuncAuxMockPtrIDGeneratorClear ")
	}
	AuxMockIncrementRecorderAuxMockPtrIDGeneratorClear()
	return
}

//
// Mock: bitfmt(argb uint64)(reta string)
//

type MockArgsTypebitfmt struct {
	ApomockCallNumber int
	Argb              uint64
}

var LastMockArgsbitfmt MockArgsTypebitfmt

// AuxMockbitfmt(argb uint64)(reta string) - Generated mock function
func AuxMockbitfmt(argb uint64) (reta string) {
	LastMockArgsbitfmt = MockArgsTypebitfmt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockbitfmt(),
		Argb:              argb,
	}
	rargs, rerr := apomock.GetNext("streams.bitfmt")
	if rerr != nil {
		panic("Error getting next entry for method: streams.bitfmt")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:streams.bitfmt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockbitfmt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockbitfmt int = 0

var condRecorderAuxMockbitfmt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockbitfmt(i int) {
	condRecorderAuxMockbitfmt.L.Lock()
	for recorderAuxMockbitfmt < i {
		condRecorderAuxMockbitfmt.Wait()
	}
	condRecorderAuxMockbitfmt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockbitfmt() {
	condRecorderAuxMockbitfmt.L.Lock()
	recorderAuxMockbitfmt++
	condRecorderAuxMockbitfmt.L.Unlock()
	condRecorderAuxMockbitfmt.Broadcast()
}
func AuxMockGetRecorderAuxMockbitfmt() (ret int) {
	condRecorderAuxMockbitfmt.L.Lock()
	ret = recorderAuxMockbitfmt
	condRecorderAuxMockbitfmt.L.Unlock()
	return
}

// bitfmt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func bitfmt(argb uint64) (reta string) {
	FuncAuxMockbitfmt, ok := apomock.GetRegisteredFunc("streams.bitfmt")
	if ok {
		reta = FuncAuxMockbitfmt.(func(argb uint64) (reta string))(argb)
	} else {
		panic("FuncAuxMockbitfmt ")
	}
	AuxMockIncrementRecorderAuxMockbitfmt()
	return
}

//
// Mock: bucketOffset(argi int)(reta int)
//

type MockArgsTypebucketOffset struct {
	ApomockCallNumber int
	Argi              int
}

var LastMockArgsbucketOffset MockArgsTypebucketOffset

// AuxMockbucketOffset(argi int)(reta int) - Generated mock function
func AuxMockbucketOffset(argi int) (reta int) {
	LastMockArgsbucketOffset = MockArgsTypebucketOffset{
		ApomockCallNumber: AuxMockGetRecorderAuxMockbucketOffset(),
		Argi:              argi,
	}
	rargs, rerr := apomock.GetNext("streams.bucketOffset")
	if rerr != nil {
		panic("Error getting next entry for method: streams.bucketOffset")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:streams.bucketOffset")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockbucketOffset  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockbucketOffset int = 0

var condRecorderAuxMockbucketOffset *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockbucketOffset(i int) {
	condRecorderAuxMockbucketOffset.L.Lock()
	for recorderAuxMockbucketOffset < i {
		condRecorderAuxMockbucketOffset.Wait()
	}
	condRecorderAuxMockbucketOffset.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockbucketOffset() {
	condRecorderAuxMockbucketOffset.L.Lock()
	recorderAuxMockbucketOffset++
	condRecorderAuxMockbucketOffset.L.Unlock()
	condRecorderAuxMockbucketOffset.Broadcast()
}
func AuxMockGetRecorderAuxMockbucketOffset() (ret int) {
	condRecorderAuxMockbucketOffset.L.Lock()
	ret = recorderAuxMockbucketOffset
	condRecorderAuxMockbucketOffset.L.Unlock()
	return
}

// bucketOffset - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func bucketOffset(argi int) (reta int) {
	FuncAuxMockbucketOffset, ok := apomock.GetRegisteredFunc("streams.bucketOffset")
	if ok {
		reta = FuncAuxMockbucketOffset.(func(argi int) (reta int))(argi)
	} else {
		panic("FuncAuxMockbucketOffset ")
	}
	AuxMockIncrementRecorderAuxMockbucketOffset()
	return
}

//
// Mock: (recvs *IDGenerator)isSet(argstream int)(reta bool)
//

type MockArgsTypeIDGeneratorisSet struct {
	ApomockCallNumber int
	Argstream         int
}

var LastMockArgsIDGeneratorisSet MockArgsTypeIDGeneratorisSet

// (recvs *IDGenerator)AuxMockisSet(argstream int)(reta bool) - Generated mock function
func (recvs *IDGenerator) AuxMockisSet(argstream int) (reta bool) {
	LastMockArgsIDGeneratorisSet = MockArgsTypeIDGeneratorisSet{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrIDGeneratorisSet(),
		Argstream:         argstream,
	}
	rargs, rerr := apomock.GetNext("streams.IDGenerator.isSet")
	if rerr != nil {
		panic("Error getting next entry for method: streams.IDGenerator.isSet")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:streams.IDGenerator.isSet")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrIDGeneratorisSet  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIDGeneratorisSet int = 0

var condRecorderAuxMockPtrIDGeneratorisSet *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIDGeneratorisSet(i int) {
	condRecorderAuxMockPtrIDGeneratorisSet.L.Lock()
	for recorderAuxMockPtrIDGeneratorisSet < i {
		condRecorderAuxMockPtrIDGeneratorisSet.Wait()
	}
	condRecorderAuxMockPtrIDGeneratorisSet.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIDGeneratorisSet() {
	condRecorderAuxMockPtrIDGeneratorisSet.L.Lock()
	recorderAuxMockPtrIDGeneratorisSet++
	condRecorderAuxMockPtrIDGeneratorisSet.L.Unlock()
	condRecorderAuxMockPtrIDGeneratorisSet.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIDGeneratorisSet() (ret int) {
	condRecorderAuxMockPtrIDGeneratorisSet.L.Lock()
	ret = recorderAuxMockPtrIDGeneratorisSet
	condRecorderAuxMockPtrIDGeneratorisSet.L.Unlock()
	return
}

// (recvs *IDGenerator)isSet - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *IDGenerator) isSet(argstream int) (reta bool) {
	FuncAuxMockPtrIDGeneratorisSet, ok := apomock.GetRegisteredFunc("streams.IDGenerator.isSet")
	if ok {
		reta = FuncAuxMockPtrIDGeneratorisSet.(func(recvs *IDGenerator, argstream int) (reta bool))(recvs, argstream)
	} else {
		panic("FuncAuxMockPtrIDGeneratorisSet ")
	}
	AuxMockIncrementRecorderAuxMockPtrIDGeneratorisSet()
	return
}

//
// Mock: (recvs *IDGenerator)String()(reta string)
//

type MockArgsTypeIDGeneratorString struct {
	ApomockCallNumber int
}

var LastMockArgsIDGeneratorString MockArgsTypeIDGeneratorString

// (recvs *IDGenerator)AuxMockString()(reta string) - Generated mock function
func (recvs *IDGenerator) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("streams.IDGenerator.String")
	if rerr != nil {
		panic("Error getting next entry for method: streams.IDGenerator.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:streams.IDGenerator.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrIDGeneratorString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIDGeneratorString int = 0

var condRecorderAuxMockPtrIDGeneratorString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIDGeneratorString(i int) {
	condRecorderAuxMockPtrIDGeneratorString.L.Lock()
	for recorderAuxMockPtrIDGeneratorString < i {
		condRecorderAuxMockPtrIDGeneratorString.Wait()
	}
	condRecorderAuxMockPtrIDGeneratorString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIDGeneratorString() {
	condRecorderAuxMockPtrIDGeneratorString.L.Lock()
	recorderAuxMockPtrIDGeneratorString++
	condRecorderAuxMockPtrIDGeneratorString.L.Unlock()
	condRecorderAuxMockPtrIDGeneratorString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIDGeneratorString() (ret int) {
	condRecorderAuxMockPtrIDGeneratorString.L.Lock()
	ret = recorderAuxMockPtrIDGeneratorString
	condRecorderAuxMockPtrIDGeneratorString.L.Unlock()
	return
}

// (recvs *IDGenerator)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *IDGenerator) String() (reta string) {
	FuncAuxMockPtrIDGeneratorString, ok := apomock.GetRegisteredFunc("streams.IDGenerator.String")
	if ok {
		reta = FuncAuxMockPtrIDGeneratorString.(func(recvs *IDGenerator) (reta string))(recvs)
	} else {
		panic("FuncAuxMockPtrIDGeneratorString ")
	}
	AuxMockIncrementRecorderAuxMockPtrIDGeneratorString()
	return
}

//
// Mock: isSet(argbits uint64, argstream int)(reta bool)
//

type MockArgsTypeisSet struct {
	ApomockCallNumber int
	Argbits           uint64
	Argstream         int
}

var LastMockArgsisSet MockArgsTypeisSet

// AuxMockisSet(argbits uint64, argstream int)(reta bool) - Generated mock function
func AuxMockisSet(argbits uint64, argstream int) (reta bool) {
	LastMockArgsisSet = MockArgsTypeisSet{
		ApomockCallNumber: AuxMockGetRecorderAuxMockisSet(),
		Argbits:           argbits,
		Argstream:         argstream,
	}
	rargs, rerr := apomock.GetNext("streams.isSet")
	if rerr != nil {
		panic("Error getting next entry for method: streams.isSet")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:streams.isSet")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockisSet  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockisSet int = 0

var condRecorderAuxMockisSet *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockisSet(i int) {
	condRecorderAuxMockisSet.L.Lock()
	for recorderAuxMockisSet < i {
		condRecorderAuxMockisSet.Wait()
	}
	condRecorderAuxMockisSet.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockisSet() {
	condRecorderAuxMockisSet.L.Lock()
	recorderAuxMockisSet++
	condRecorderAuxMockisSet.L.Unlock()
	condRecorderAuxMockisSet.Broadcast()
}
func AuxMockGetRecorderAuxMockisSet() (ret int) {
	condRecorderAuxMockisSet.L.Lock()
	ret = recorderAuxMockisSet
	condRecorderAuxMockisSet.L.Unlock()
	return
}

// isSet - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func isSet(argbits uint64, argstream int) (reta bool) {
	FuncAuxMockisSet, ok := apomock.GetRegisteredFunc("streams.isSet")
	if ok {
		reta = FuncAuxMockisSet.(func(argbits uint64, argstream int) (reta bool))(argbits, argstream)
	} else {
		panic("FuncAuxMockisSet ")
	}
	AuxMockIncrementRecorderAuxMockisSet()
	return
}

//
// Mock: (recvs *IDGenerator)Available()(reta int)
//

type MockArgsTypeIDGeneratorAvailable struct {
	ApomockCallNumber int
}

var LastMockArgsIDGeneratorAvailable MockArgsTypeIDGeneratorAvailable

// (recvs *IDGenerator)AuxMockAvailable()(reta int) - Generated mock function
func (recvs *IDGenerator) AuxMockAvailable() (reta int) {
	rargs, rerr := apomock.GetNext("streams.IDGenerator.Available")
	if rerr != nil {
		panic("Error getting next entry for method: streams.IDGenerator.Available")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:streams.IDGenerator.Available")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrIDGeneratorAvailable  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIDGeneratorAvailable int = 0

var condRecorderAuxMockPtrIDGeneratorAvailable *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIDGeneratorAvailable(i int) {
	condRecorderAuxMockPtrIDGeneratorAvailable.L.Lock()
	for recorderAuxMockPtrIDGeneratorAvailable < i {
		condRecorderAuxMockPtrIDGeneratorAvailable.Wait()
	}
	condRecorderAuxMockPtrIDGeneratorAvailable.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIDGeneratorAvailable() {
	condRecorderAuxMockPtrIDGeneratorAvailable.L.Lock()
	recorderAuxMockPtrIDGeneratorAvailable++
	condRecorderAuxMockPtrIDGeneratorAvailable.L.Unlock()
	condRecorderAuxMockPtrIDGeneratorAvailable.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIDGeneratorAvailable() (ret int) {
	condRecorderAuxMockPtrIDGeneratorAvailable.L.Lock()
	ret = recorderAuxMockPtrIDGeneratorAvailable
	condRecorderAuxMockPtrIDGeneratorAvailable.L.Unlock()
	return
}

// (recvs *IDGenerator)Available - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *IDGenerator) Available() (reta int) {
	FuncAuxMockPtrIDGeneratorAvailable, ok := apomock.GetRegisteredFunc("streams.IDGenerator.Available")
	if ok {
		reta = FuncAuxMockPtrIDGeneratorAvailable.(func(recvs *IDGenerator) (reta int))(recvs)
	} else {
		panic("FuncAuxMockPtrIDGeneratorAvailable ")
	}
	AuxMockIncrementRecorderAuxMockPtrIDGeneratorAvailable()
	return
}

//
// Mock: New(argprotocol int)(reta *IDGenerator)
//

type MockArgsTypeNew struct {
	ApomockCallNumber int
	Argprotocol       int
}

var LastMockArgsNew MockArgsTypeNew

// AuxMockNew(argprotocol int)(reta *IDGenerator) - Generated mock function
func AuxMockNew(argprotocol int) (reta *IDGenerator) {
	LastMockArgsNew = MockArgsTypeNew{
		ApomockCallNumber: AuxMockGetRecorderAuxMockNew(),
		Argprotocol:       argprotocol,
	}
	rargs, rerr := apomock.GetNext("streams.New")
	if rerr != nil {
		panic("Error getting next entry for method: streams.New")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:streams.New")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*IDGenerator)
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
func New(argprotocol int) (reta *IDGenerator) {
	FuncAuxMockNew, ok := apomock.GetRegisteredFunc("streams.New")
	if ok {
		reta = FuncAuxMockNew.(func(argprotocol int) (reta *IDGenerator))(argprotocol)
	} else {
		panic("FuncAuxMockNew ")
	}
	AuxMockIncrementRecorderAuxMockNew()
	return
}

//
// Mock: streamFromBucket(argbucket int, argstreamInBucket int)(reta int)
//

type MockArgsTypestreamFromBucket struct {
	ApomockCallNumber int
	Argbucket         int
	ArgstreamInBucket int
}

var LastMockArgsstreamFromBucket MockArgsTypestreamFromBucket

// AuxMockstreamFromBucket(argbucket int, argstreamInBucket int)(reta int) - Generated mock function
func AuxMockstreamFromBucket(argbucket int, argstreamInBucket int) (reta int) {
	LastMockArgsstreamFromBucket = MockArgsTypestreamFromBucket{
		ApomockCallNumber: AuxMockGetRecorderAuxMockstreamFromBucket(),
		Argbucket:         argbucket,
		ArgstreamInBucket: argstreamInBucket,
	}
	rargs, rerr := apomock.GetNext("streams.streamFromBucket")
	if rerr != nil {
		panic("Error getting next entry for method: streams.streamFromBucket")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:streams.streamFromBucket")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockstreamFromBucket  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockstreamFromBucket int = 0

var condRecorderAuxMockstreamFromBucket *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockstreamFromBucket(i int) {
	condRecorderAuxMockstreamFromBucket.L.Lock()
	for recorderAuxMockstreamFromBucket < i {
		condRecorderAuxMockstreamFromBucket.Wait()
	}
	condRecorderAuxMockstreamFromBucket.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockstreamFromBucket() {
	condRecorderAuxMockstreamFromBucket.L.Lock()
	recorderAuxMockstreamFromBucket++
	condRecorderAuxMockstreamFromBucket.L.Unlock()
	condRecorderAuxMockstreamFromBucket.Broadcast()
}
func AuxMockGetRecorderAuxMockstreamFromBucket() (ret int) {
	condRecorderAuxMockstreamFromBucket.L.Lock()
	ret = recorderAuxMockstreamFromBucket
	condRecorderAuxMockstreamFromBucket.L.Unlock()
	return
}

// streamFromBucket - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func streamFromBucket(argbucket int, argstreamInBucket int) (reta int) {
	FuncAuxMockstreamFromBucket, ok := apomock.GetRegisteredFunc("streams.streamFromBucket")
	if ok {
		reta = FuncAuxMockstreamFromBucket.(func(argbucket int, argstreamInBucket int) (reta int))(argbucket, argstreamInBucket)
	} else {
		panic("FuncAuxMockstreamFromBucket ")
	}
	AuxMockIncrementRecorderAuxMockstreamFromBucket()
	return
}

//
// Mock: (recvs *IDGenerator)GetStream()(reta int, retb bool)
//

type MockArgsTypeIDGeneratorGetStream struct {
	ApomockCallNumber int
}

var LastMockArgsIDGeneratorGetStream MockArgsTypeIDGeneratorGetStream

// (recvs *IDGenerator)AuxMockGetStream()(reta int, retb bool) - Generated mock function
func (recvs *IDGenerator) AuxMockGetStream() (reta int, retb bool) {
	rargs, rerr := apomock.GetNext("streams.IDGenerator.GetStream")
	if rerr != nil {
		panic("Error getting next entry for method: streams.IDGenerator.GetStream")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:streams.IDGenerator.GetStream")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(bool)
	}
	return
}

// RecorderAuxMockPtrIDGeneratorGetStream  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrIDGeneratorGetStream int = 0

var condRecorderAuxMockPtrIDGeneratorGetStream *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrIDGeneratorGetStream(i int) {
	condRecorderAuxMockPtrIDGeneratorGetStream.L.Lock()
	for recorderAuxMockPtrIDGeneratorGetStream < i {
		condRecorderAuxMockPtrIDGeneratorGetStream.Wait()
	}
	condRecorderAuxMockPtrIDGeneratorGetStream.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrIDGeneratorGetStream() {
	condRecorderAuxMockPtrIDGeneratorGetStream.L.Lock()
	recorderAuxMockPtrIDGeneratorGetStream++
	condRecorderAuxMockPtrIDGeneratorGetStream.L.Unlock()
	condRecorderAuxMockPtrIDGeneratorGetStream.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrIDGeneratorGetStream() (ret int) {
	condRecorderAuxMockPtrIDGeneratorGetStream.L.Lock()
	ret = recorderAuxMockPtrIDGeneratorGetStream
	condRecorderAuxMockPtrIDGeneratorGetStream.L.Unlock()
	return
}

// (recvs *IDGenerator)GetStream - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *IDGenerator) GetStream() (reta int, retb bool) {
	FuncAuxMockPtrIDGeneratorGetStream, ok := apomock.GetRegisteredFunc("streams.IDGenerator.GetStream")
	if ok {
		reta, retb = FuncAuxMockPtrIDGeneratorGetStream.(func(recvs *IDGenerator) (reta int, retb bool))(recvs)
	} else {
		panic("FuncAuxMockPtrIDGeneratorGetStream ")
	}
	AuxMockIncrementRecorderAuxMockPtrIDGeneratorGetStream()
	return
}

//
// Mock: streamOffset(argstream int)(reta uint64)
//

type MockArgsTypestreamOffset struct {
	ApomockCallNumber int
	Argstream         int
}

var LastMockArgsstreamOffset MockArgsTypestreamOffset

// AuxMockstreamOffset(argstream int)(reta uint64) - Generated mock function
func AuxMockstreamOffset(argstream int) (reta uint64) {
	LastMockArgsstreamOffset = MockArgsTypestreamOffset{
		ApomockCallNumber: AuxMockGetRecorderAuxMockstreamOffset(),
		Argstream:         argstream,
	}
	rargs, rerr := apomock.GetNext("streams.streamOffset")
	if rerr != nil {
		panic("Error getting next entry for method: streams.streamOffset")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:streams.streamOffset")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(uint64)
	}
	return
}

// RecorderAuxMockstreamOffset  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockstreamOffset int = 0

var condRecorderAuxMockstreamOffset *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockstreamOffset(i int) {
	condRecorderAuxMockstreamOffset.L.Lock()
	for recorderAuxMockstreamOffset < i {
		condRecorderAuxMockstreamOffset.Wait()
	}
	condRecorderAuxMockstreamOffset.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockstreamOffset() {
	condRecorderAuxMockstreamOffset.L.Lock()
	recorderAuxMockstreamOffset++
	condRecorderAuxMockstreamOffset.L.Unlock()
	condRecorderAuxMockstreamOffset.Broadcast()
}
func AuxMockGetRecorderAuxMockstreamOffset() (ret int) {
	condRecorderAuxMockstreamOffset.L.Lock()
	ret = recorderAuxMockstreamOffset
	condRecorderAuxMockstreamOffset.L.Unlock()
	return
}

// streamOffset - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func streamOffset(argstream int) (reta uint64) {
	FuncAuxMockstreamOffset, ok := apomock.GetRegisteredFunc("streams.streamOffset")
	if ok {
		reta = FuncAuxMockstreamOffset.(func(argstream int) (reta uint64))(argstream)
	} else {
		panic("FuncAuxMockstreamOffset ")
	}
	AuxMockIncrementRecorderAuxMockstreamOffset()
	return
}
