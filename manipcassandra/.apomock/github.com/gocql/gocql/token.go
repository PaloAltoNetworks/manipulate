// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "fmt"

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

import "math/big"

func init() {
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructRandomPartitioner, apomockNewStructRandomPartitioner)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructMurmur3Partitioner, apomockNewStructMurmur3Partitioner)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructOrderedPartitioner, apomockNewStructOrderedPartitioner)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructTokenRing, apomockNewStructTokenRing)

	apomock.RegisterFunc("gocql", "gocql.murmur3Partitioner.Hash", (murmur3Partitioner).AuxMockHash)
	apomock.RegisterFunc("gocql", "gocql.orderedToken.Less", (orderedToken).AuxMockLess)
	apomock.RegisterFunc("gocql", "gocql.randomPartitioner.ParseString", (randomPartitioner).AuxMockParseString)
	apomock.RegisterFunc("gocql", "gocql.tokenRing.GetHostForToken", (*tokenRing).AuxMockGetHostForToken)
	apomock.RegisterFunc("gocql", "gocql.murmur3Token.Less", (murmur3Token).AuxMockLess)
	apomock.RegisterFunc("gocql", "gocql.orderedPartitioner.Hash", (orderedPartitioner).AuxMockHash)
	apomock.RegisterFunc("gocql", "gocql.orderedPartitioner.ParseString", (orderedPartitioner).AuxMockParseString)
	apomock.RegisterFunc("gocql", "gocql.orderedToken.String", (orderedToken).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.tokenRing.Len", (*tokenRing).AuxMockLen)
	apomock.RegisterFunc("gocql", "gocql.murmur3Partitioner.ParseString", (murmur3Partitioner).AuxMockParseString)
	apomock.RegisterFunc("gocql", "gocql.murmur3Token.String", (murmur3Token).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.randomPartitioner.Name", (randomPartitioner).AuxMockName)
	apomock.RegisterFunc("gocql", "gocql.randomToken.String", (*randomToken).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.randomToken.Less", (*randomToken).AuxMockLess)
	apomock.RegisterFunc("gocql", "gocql.tokenRing.Swap", (*tokenRing).AuxMockSwap)
	apomock.RegisterFunc("gocql", "gocql.tokenRing.GetHostForPartitionKey", (*tokenRing).AuxMockGetHostForPartitionKey)
	apomock.RegisterFunc("gocql", "gocql.murmur3Partitioner.Name", (murmur3Partitioner).AuxMockName)
	apomock.RegisterFunc("gocql", "gocql.orderedPartitioner.Name", (orderedPartitioner).AuxMockName)
	apomock.RegisterFunc("gocql", "gocql.randomPartitioner.Hash", (randomPartitioner).AuxMockHash)
	apomock.RegisterFunc("gocql", "gocql.newTokenRing", AuxMocknewTokenRing)
	apomock.RegisterFunc("gocql", "gocql.tokenRing.Less", (*tokenRing).AuxMockLess)
	apomock.RegisterFunc("gocql", "gocql.tokenRing.String", (*tokenRing).AuxMockString)
}

const (
	ApomockStructRandomPartitioner  = 49
	ApomockStructMurmur3Partitioner = 50
	ApomockStructOrderedPartitioner = 51
	ApomockStructTokenRing          = 52
)

//
// Internal Types: in this package and their exportable versions
//
type token interface {
	fmt.Stringer
	Less(token) bool
}
type murmur3Token int64
type randomPartitioner struct{}
type randomToken big.Int
type partitioner interface {
	Name() string
	Hash([]byte) token
	ParseString(string) token
}
type murmur3Partitioner struct{}
type orderedPartitioner struct{}
type orderedToken []byte
type tokenRing struct {
	partitioner partitioner
	tokens      []token
	hosts       []*HostInfo
}

//
// External Types: in this package
//

func apomockNewStructRandomPartitioner() interface{}  { return &randomPartitioner{} }
func apomockNewStructMurmur3Partitioner() interface{} { return &murmur3Partitioner{} }
func apomockNewStructOrderedPartitioner() interface{} { return &orderedPartitioner{} }
func apomockNewStructTokenRing() interface{}          { return &tokenRing{} }

//
// Mock: (recvp murmur3Partitioner)Hash(argpartitionKey []byte)(reta token)
//

type MockArgsTypemurmur3PartitionerHash struct {
	ApomockCallNumber int
	ArgpartitionKey   []byte
}

var LastMockArgsmurmur3PartitionerHash MockArgsTypemurmur3PartitionerHash

// (recvp murmur3Partitioner)AuxMockHash(argpartitionKey []byte)(reta token) - Generated mock function
func (recvp murmur3Partitioner) AuxMockHash(argpartitionKey []byte) (reta token) {
	LastMockArgsmurmur3PartitionerHash = MockArgsTypemurmur3PartitionerHash{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmurmur3PartitionerHash(),
		ArgpartitionKey:   argpartitionKey,
	}
	rargs, rerr := apomock.GetNext("gocql.murmur3Partitioner.Hash")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.murmur3Partitioner.Hash")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.murmur3Partitioner.Hash")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(token)
	}
	return
}

// RecorderAuxMockmurmur3PartitionerHash  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmurmur3PartitionerHash int = 0

var condRecorderAuxMockmurmur3PartitionerHash *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmurmur3PartitionerHash(i int) {
	condRecorderAuxMockmurmur3PartitionerHash.L.Lock()
	for recorderAuxMockmurmur3PartitionerHash < i {
		condRecorderAuxMockmurmur3PartitionerHash.Wait()
	}
	condRecorderAuxMockmurmur3PartitionerHash.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmurmur3PartitionerHash() {
	condRecorderAuxMockmurmur3PartitionerHash.L.Lock()
	recorderAuxMockmurmur3PartitionerHash++
	condRecorderAuxMockmurmur3PartitionerHash.L.Unlock()
	condRecorderAuxMockmurmur3PartitionerHash.Broadcast()
}
func AuxMockGetRecorderAuxMockmurmur3PartitionerHash() (ret int) {
	condRecorderAuxMockmurmur3PartitionerHash.L.Lock()
	ret = recorderAuxMockmurmur3PartitionerHash
	condRecorderAuxMockmurmur3PartitionerHash.L.Unlock()
	return
}

// (recvp murmur3Partitioner)Hash - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp murmur3Partitioner) Hash(argpartitionKey []byte) (reta token) {
	FuncAuxMockmurmur3PartitionerHash, ok := apomock.GetRegisteredFunc("gocql.murmur3Partitioner.Hash")
	if ok {
		reta = FuncAuxMockmurmur3PartitionerHash.(func(recvp murmur3Partitioner, argpartitionKey []byte) (reta token))(recvp, argpartitionKey)
	} else {
		panic("FuncAuxMockmurmur3PartitionerHash ")
	}
	AuxMockIncrementRecorderAuxMockmurmur3PartitionerHash()
	return
}

//
// Mock: (recvo orderedToken)Less(argtoken token)(reta bool)
//

type MockArgsTypeorderedTokenLess struct {
	ApomockCallNumber int
	Argtoken          token
}

var LastMockArgsorderedTokenLess MockArgsTypeorderedTokenLess

// (recvo orderedToken)AuxMockLess(argtoken token)(reta bool) - Generated mock function
func (recvo orderedToken) AuxMockLess(argtoken token) (reta bool) {
	LastMockArgsorderedTokenLess = MockArgsTypeorderedTokenLess{
		ApomockCallNumber: AuxMockGetRecorderAuxMockorderedTokenLess(),
		Argtoken:          argtoken,
	}
	rargs, rerr := apomock.GetNext("gocql.orderedToken.Less")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.orderedToken.Less")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.orderedToken.Less")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockorderedTokenLess  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockorderedTokenLess int = 0

var condRecorderAuxMockorderedTokenLess *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockorderedTokenLess(i int) {
	condRecorderAuxMockorderedTokenLess.L.Lock()
	for recorderAuxMockorderedTokenLess < i {
		condRecorderAuxMockorderedTokenLess.Wait()
	}
	condRecorderAuxMockorderedTokenLess.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockorderedTokenLess() {
	condRecorderAuxMockorderedTokenLess.L.Lock()
	recorderAuxMockorderedTokenLess++
	condRecorderAuxMockorderedTokenLess.L.Unlock()
	condRecorderAuxMockorderedTokenLess.Broadcast()
}
func AuxMockGetRecorderAuxMockorderedTokenLess() (ret int) {
	condRecorderAuxMockorderedTokenLess.L.Lock()
	ret = recorderAuxMockorderedTokenLess
	condRecorderAuxMockorderedTokenLess.L.Unlock()
	return
}

// (recvo orderedToken)Less - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvo orderedToken) Less(argtoken token) (reta bool) {
	FuncAuxMockorderedTokenLess, ok := apomock.GetRegisteredFunc("gocql.orderedToken.Less")
	if ok {
		reta = FuncAuxMockorderedTokenLess.(func(recvo orderedToken, argtoken token) (reta bool))(recvo, argtoken)
	} else {
		panic("FuncAuxMockorderedTokenLess ")
	}
	AuxMockIncrementRecorderAuxMockorderedTokenLess()
	return
}

//
// Mock: (recvp randomPartitioner)ParseString(argstr string)(reta token)
//

type MockArgsTyperandomPartitionerParseString struct {
	ApomockCallNumber int
	Argstr            string
}

var LastMockArgsrandomPartitionerParseString MockArgsTyperandomPartitionerParseString

// (recvp randomPartitioner)AuxMockParseString(argstr string)(reta token) - Generated mock function
func (recvp randomPartitioner) AuxMockParseString(argstr string) (reta token) {
	LastMockArgsrandomPartitionerParseString = MockArgsTyperandomPartitionerParseString{
		ApomockCallNumber: AuxMockGetRecorderAuxMockrandomPartitionerParseString(),
		Argstr:            argstr,
	}
	rargs, rerr := apomock.GetNext("gocql.randomPartitioner.ParseString")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.randomPartitioner.ParseString")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.randomPartitioner.ParseString")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(token)
	}
	return
}

// RecorderAuxMockrandomPartitionerParseString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockrandomPartitionerParseString int = 0

var condRecorderAuxMockrandomPartitionerParseString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockrandomPartitionerParseString(i int) {
	condRecorderAuxMockrandomPartitionerParseString.L.Lock()
	for recorderAuxMockrandomPartitionerParseString < i {
		condRecorderAuxMockrandomPartitionerParseString.Wait()
	}
	condRecorderAuxMockrandomPartitionerParseString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockrandomPartitionerParseString() {
	condRecorderAuxMockrandomPartitionerParseString.L.Lock()
	recorderAuxMockrandomPartitionerParseString++
	condRecorderAuxMockrandomPartitionerParseString.L.Unlock()
	condRecorderAuxMockrandomPartitionerParseString.Broadcast()
}
func AuxMockGetRecorderAuxMockrandomPartitionerParseString() (ret int) {
	condRecorderAuxMockrandomPartitionerParseString.L.Lock()
	ret = recorderAuxMockrandomPartitionerParseString
	condRecorderAuxMockrandomPartitionerParseString.L.Unlock()
	return
}

// (recvp randomPartitioner)ParseString - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp randomPartitioner) ParseString(argstr string) (reta token) {
	FuncAuxMockrandomPartitionerParseString, ok := apomock.GetRegisteredFunc("gocql.randomPartitioner.ParseString")
	if ok {
		reta = FuncAuxMockrandomPartitionerParseString.(func(recvp randomPartitioner, argstr string) (reta token))(recvp, argstr)
	} else {
		panic("FuncAuxMockrandomPartitionerParseString ")
	}
	AuxMockIncrementRecorderAuxMockrandomPartitionerParseString()
	return
}

//
// Mock: (recvt *tokenRing)GetHostForToken(argtoken token)(reta *HostInfo)
//

type MockArgsTypetokenRingGetHostForToken struct {
	ApomockCallNumber int
	Argtoken          token
}

var LastMockArgstokenRingGetHostForToken MockArgsTypetokenRingGetHostForToken

// (recvt *tokenRing)AuxMockGetHostForToken(argtoken token)(reta *HostInfo) - Generated mock function
func (recvt *tokenRing) AuxMockGetHostForToken(argtoken token) (reta *HostInfo) {
	LastMockArgstokenRingGetHostForToken = MockArgsTypetokenRingGetHostForToken{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrtokenRingGetHostForToken(),
		Argtoken:          argtoken,
	}
	rargs, rerr := apomock.GetNext("gocql.tokenRing.GetHostForToken")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.tokenRing.GetHostForToken")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.tokenRing.GetHostForToken")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrtokenRingGetHostForToken  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenRingGetHostForToken int = 0

var condRecorderAuxMockPtrtokenRingGetHostForToken *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenRingGetHostForToken(i int) {
	condRecorderAuxMockPtrtokenRingGetHostForToken.L.Lock()
	for recorderAuxMockPtrtokenRingGetHostForToken < i {
		condRecorderAuxMockPtrtokenRingGetHostForToken.Wait()
	}
	condRecorderAuxMockPtrtokenRingGetHostForToken.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenRingGetHostForToken() {
	condRecorderAuxMockPtrtokenRingGetHostForToken.L.Lock()
	recorderAuxMockPtrtokenRingGetHostForToken++
	condRecorderAuxMockPtrtokenRingGetHostForToken.L.Unlock()
	condRecorderAuxMockPtrtokenRingGetHostForToken.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenRingGetHostForToken() (ret int) {
	condRecorderAuxMockPtrtokenRingGetHostForToken.L.Lock()
	ret = recorderAuxMockPtrtokenRingGetHostForToken
	condRecorderAuxMockPtrtokenRingGetHostForToken.L.Unlock()
	return
}

// (recvt *tokenRing)GetHostForToken - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenRing) GetHostForToken(argtoken token) (reta *HostInfo) {
	FuncAuxMockPtrtokenRingGetHostForToken, ok := apomock.GetRegisteredFunc("gocql.tokenRing.GetHostForToken")
	if ok {
		reta = FuncAuxMockPtrtokenRingGetHostForToken.(func(recvt *tokenRing, argtoken token) (reta *HostInfo))(recvt, argtoken)
	} else {
		panic("FuncAuxMockPtrtokenRingGetHostForToken ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenRingGetHostForToken()
	return
}

//
// Mock: (recvm murmur3Token)Less(argtoken token)(reta bool)
//

type MockArgsTypemurmur3TokenLess struct {
	ApomockCallNumber int
	Argtoken          token
}

var LastMockArgsmurmur3TokenLess MockArgsTypemurmur3TokenLess

// (recvm murmur3Token)AuxMockLess(argtoken token)(reta bool) - Generated mock function
func (recvm murmur3Token) AuxMockLess(argtoken token) (reta bool) {
	LastMockArgsmurmur3TokenLess = MockArgsTypemurmur3TokenLess{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmurmur3TokenLess(),
		Argtoken:          argtoken,
	}
	rargs, rerr := apomock.GetNext("gocql.murmur3Token.Less")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.murmur3Token.Less")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.murmur3Token.Less")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockmurmur3TokenLess  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmurmur3TokenLess int = 0

var condRecorderAuxMockmurmur3TokenLess *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmurmur3TokenLess(i int) {
	condRecorderAuxMockmurmur3TokenLess.L.Lock()
	for recorderAuxMockmurmur3TokenLess < i {
		condRecorderAuxMockmurmur3TokenLess.Wait()
	}
	condRecorderAuxMockmurmur3TokenLess.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmurmur3TokenLess() {
	condRecorderAuxMockmurmur3TokenLess.L.Lock()
	recorderAuxMockmurmur3TokenLess++
	condRecorderAuxMockmurmur3TokenLess.L.Unlock()
	condRecorderAuxMockmurmur3TokenLess.Broadcast()
}
func AuxMockGetRecorderAuxMockmurmur3TokenLess() (ret int) {
	condRecorderAuxMockmurmur3TokenLess.L.Lock()
	ret = recorderAuxMockmurmur3TokenLess
	condRecorderAuxMockmurmur3TokenLess.L.Unlock()
	return
}

// (recvm murmur3Token)Less - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvm murmur3Token) Less(argtoken token) (reta bool) {
	FuncAuxMockmurmur3TokenLess, ok := apomock.GetRegisteredFunc("gocql.murmur3Token.Less")
	if ok {
		reta = FuncAuxMockmurmur3TokenLess.(func(recvm murmur3Token, argtoken token) (reta bool))(recvm, argtoken)
	} else {
		panic("FuncAuxMockmurmur3TokenLess ")
	}
	AuxMockIncrementRecorderAuxMockmurmur3TokenLess()
	return
}

//
// Mock: (recvp orderedPartitioner)Hash(argpartitionKey []byte)(reta token)
//

type MockArgsTypeorderedPartitionerHash struct {
	ApomockCallNumber int
	ArgpartitionKey   []byte
}

var LastMockArgsorderedPartitionerHash MockArgsTypeorderedPartitionerHash

// (recvp orderedPartitioner)AuxMockHash(argpartitionKey []byte)(reta token) - Generated mock function
func (recvp orderedPartitioner) AuxMockHash(argpartitionKey []byte) (reta token) {
	LastMockArgsorderedPartitionerHash = MockArgsTypeorderedPartitionerHash{
		ApomockCallNumber: AuxMockGetRecorderAuxMockorderedPartitionerHash(),
		ArgpartitionKey:   argpartitionKey,
	}
	rargs, rerr := apomock.GetNext("gocql.orderedPartitioner.Hash")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.orderedPartitioner.Hash")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.orderedPartitioner.Hash")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(token)
	}
	return
}

// RecorderAuxMockorderedPartitionerHash  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockorderedPartitionerHash int = 0

var condRecorderAuxMockorderedPartitionerHash *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockorderedPartitionerHash(i int) {
	condRecorderAuxMockorderedPartitionerHash.L.Lock()
	for recorderAuxMockorderedPartitionerHash < i {
		condRecorderAuxMockorderedPartitionerHash.Wait()
	}
	condRecorderAuxMockorderedPartitionerHash.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockorderedPartitionerHash() {
	condRecorderAuxMockorderedPartitionerHash.L.Lock()
	recorderAuxMockorderedPartitionerHash++
	condRecorderAuxMockorderedPartitionerHash.L.Unlock()
	condRecorderAuxMockorderedPartitionerHash.Broadcast()
}
func AuxMockGetRecorderAuxMockorderedPartitionerHash() (ret int) {
	condRecorderAuxMockorderedPartitionerHash.L.Lock()
	ret = recorderAuxMockorderedPartitionerHash
	condRecorderAuxMockorderedPartitionerHash.L.Unlock()
	return
}

// (recvp orderedPartitioner)Hash - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp orderedPartitioner) Hash(argpartitionKey []byte) (reta token) {
	FuncAuxMockorderedPartitionerHash, ok := apomock.GetRegisteredFunc("gocql.orderedPartitioner.Hash")
	if ok {
		reta = FuncAuxMockorderedPartitionerHash.(func(recvp orderedPartitioner, argpartitionKey []byte) (reta token))(recvp, argpartitionKey)
	} else {
		panic("FuncAuxMockorderedPartitionerHash ")
	}
	AuxMockIncrementRecorderAuxMockorderedPartitionerHash()
	return
}

//
// Mock: (recvp orderedPartitioner)ParseString(argstr string)(reta token)
//

type MockArgsTypeorderedPartitionerParseString struct {
	ApomockCallNumber int
	Argstr            string
}

var LastMockArgsorderedPartitionerParseString MockArgsTypeorderedPartitionerParseString

// (recvp orderedPartitioner)AuxMockParseString(argstr string)(reta token) - Generated mock function
func (recvp orderedPartitioner) AuxMockParseString(argstr string) (reta token) {
	LastMockArgsorderedPartitionerParseString = MockArgsTypeorderedPartitionerParseString{
		ApomockCallNumber: AuxMockGetRecorderAuxMockorderedPartitionerParseString(),
		Argstr:            argstr,
	}
	rargs, rerr := apomock.GetNext("gocql.orderedPartitioner.ParseString")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.orderedPartitioner.ParseString")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.orderedPartitioner.ParseString")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(token)
	}
	return
}

// RecorderAuxMockorderedPartitionerParseString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockorderedPartitionerParseString int = 0

var condRecorderAuxMockorderedPartitionerParseString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockorderedPartitionerParseString(i int) {
	condRecorderAuxMockorderedPartitionerParseString.L.Lock()
	for recorderAuxMockorderedPartitionerParseString < i {
		condRecorderAuxMockorderedPartitionerParseString.Wait()
	}
	condRecorderAuxMockorderedPartitionerParseString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockorderedPartitionerParseString() {
	condRecorderAuxMockorderedPartitionerParseString.L.Lock()
	recorderAuxMockorderedPartitionerParseString++
	condRecorderAuxMockorderedPartitionerParseString.L.Unlock()
	condRecorderAuxMockorderedPartitionerParseString.Broadcast()
}
func AuxMockGetRecorderAuxMockorderedPartitionerParseString() (ret int) {
	condRecorderAuxMockorderedPartitionerParseString.L.Lock()
	ret = recorderAuxMockorderedPartitionerParseString
	condRecorderAuxMockorderedPartitionerParseString.L.Unlock()
	return
}

// (recvp orderedPartitioner)ParseString - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp orderedPartitioner) ParseString(argstr string) (reta token) {
	FuncAuxMockorderedPartitionerParseString, ok := apomock.GetRegisteredFunc("gocql.orderedPartitioner.ParseString")
	if ok {
		reta = FuncAuxMockorderedPartitionerParseString.(func(recvp orderedPartitioner, argstr string) (reta token))(recvp, argstr)
	} else {
		panic("FuncAuxMockorderedPartitionerParseString ")
	}
	AuxMockIncrementRecorderAuxMockorderedPartitionerParseString()
	return
}

//
// Mock: (recvo orderedToken)String()(reta string)
//

type MockArgsTypeorderedTokenString struct {
	ApomockCallNumber int
}

var LastMockArgsorderedTokenString MockArgsTypeorderedTokenString

// (recvo orderedToken)AuxMockString()(reta string) - Generated mock function
func (recvo orderedToken) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.orderedToken.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.orderedToken.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.orderedToken.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockorderedTokenString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockorderedTokenString int = 0

var condRecorderAuxMockorderedTokenString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockorderedTokenString(i int) {
	condRecorderAuxMockorderedTokenString.L.Lock()
	for recorderAuxMockorderedTokenString < i {
		condRecorderAuxMockorderedTokenString.Wait()
	}
	condRecorderAuxMockorderedTokenString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockorderedTokenString() {
	condRecorderAuxMockorderedTokenString.L.Lock()
	recorderAuxMockorderedTokenString++
	condRecorderAuxMockorderedTokenString.L.Unlock()
	condRecorderAuxMockorderedTokenString.Broadcast()
}
func AuxMockGetRecorderAuxMockorderedTokenString() (ret int) {
	condRecorderAuxMockorderedTokenString.L.Lock()
	ret = recorderAuxMockorderedTokenString
	condRecorderAuxMockorderedTokenString.L.Unlock()
	return
}

// (recvo orderedToken)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvo orderedToken) String() (reta string) {
	FuncAuxMockorderedTokenString, ok := apomock.GetRegisteredFunc("gocql.orderedToken.String")
	if ok {
		reta = FuncAuxMockorderedTokenString.(func(recvo orderedToken) (reta string))(recvo)
	} else {
		panic("FuncAuxMockorderedTokenString ")
	}
	AuxMockIncrementRecorderAuxMockorderedTokenString()
	return
}

//
// Mock: (recvt *tokenRing)Len()(reta int)
//

type MockArgsTypetokenRingLen struct {
	ApomockCallNumber int
}

var LastMockArgstokenRingLen MockArgsTypetokenRingLen

// (recvt *tokenRing)AuxMockLen()(reta int) - Generated mock function
func (recvt *tokenRing) AuxMockLen() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.tokenRing.Len")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.tokenRing.Len")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.tokenRing.Len")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrtokenRingLen  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenRingLen int = 0

var condRecorderAuxMockPtrtokenRingLen *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenRingLen(i int) {
	condRecorderAuxMockPtrtokenRingLen.L.Lock()
	for recorderAuxMockPtrtokenRingLen < i {
		condRecorderAuxMockPtrtokenRingLen.Wait()
	}
	condRecorderAuxMockPtrtokenRingLen.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenRingLen() {
	condRecorderAuxMockPtrtokenRingLen.L.Lock()
	recorderAuxMockPtrtokenRingLen++
	condRecorderAuxMockPtrtokenRingLen.L.Unlock()
	condRecorderAuxMockPtrtokenRingLen.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenRingLen() (ret int) {
	condRecorderAuxMockPtrtokenRingLen.L.Lock()
	ret = recorderAuxMockPtrtokenRingLen
	condRecorderAuxMockPtrtokenRingLen.L.Unlock()
	return
}

// (recvt *tokenRing)Len - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenRing) Len() (reta int) {
	FuncAuxMockPtrtokenRingLen, ok := apomock.GetRegisteredFunc("gocql.tokenRing.Len")
	if ok {
		reta = FuncAuxMockPtrtokenRingLen.(func(recvt *tokenRing) (reta int))(recvt)
	} else {
		panic("FuncAuxMockPtrtokenRingLen ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenRingLen()
	return
}

//
// Mock: (recvp murmur3Partitioner)ParseString(argstr string)(reta token)
//

type MockArgsTypemurmur3PartitionerParseString struct {
	ApomockCallNumber int
	Argstr            string
}

var LastMockArgsmurmur3PartitionerParseString MockArgsTypemurmur3PartitionerParseString

// (recvp murmur3Partitioner)AuxMockParseString(argstr string)(reta token) - Generated mock function
func (recvp murmur3Partitioner) AuxMockParseString(argstr string) (reta token) {
	LastMockArgsmurmur3PartitionerParseString = MockArgsTypemurmur3PartitionerParseString{
		ApomockCallNumber: AuxMockGetRecorderAuxMockmurmur3PartitionerParseString(),
		Argstr:            argstr,
	}
	rargs, rerr := apomock.GetNext("gocql.murmur3Partitioner.ParseString")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.murmur3Partitioner.ParseString")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.murmur3Partitioner.ParseString")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(token)
	}
	return
}

// RecorderAuxMockmurmur3PartitionerParseString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmurmur3PartitionerParseString int = 0

var condRecorderAuxMockmurmur3PartitionerParseString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmurmur3PartitionerParseString(i int) {
	condRecorderAuxMockmurmur3PartitionerParseString.L.Lock()
	for recorderAuxMockmurmur3PartitionerParseString < i {
		condRecorderAuxMockmurmur3PartitionerParseString.Wait()
	}
	condRecorderAuxMockmurmur3PartitionerParseString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmurmur3PartitionerParseString() {
	condRecorderAuxMockmurmur3PartitionerParseString.L.Lock()
	recorderAuxMockmurmur3PartitionerParseString++
	condRecorderAuxMockmurmur3PartitionerParseString.L.Unlock()
	condRecorderAuxMockmurmur3PartitionerParseString.Broadcast()
}
func AuxMockGetRecorderAuxMockmurmur3PartitionerParseString() (ret int) {
	condRecorderAuxMockmurmur3PartitionerParseString.L.Lock()
	ret = recorderAuxMockmurmur3PartitionerParseString
	condRecorderAuxMockmurmur3PartitionerParseString.L.Unlock()
	return
}

// (recvp murmur3Partitioner)ParseString - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp murmur3Partitioner) ParseString(argstr string) (reta token) {
	FuncAuxMockmurmur3PartitionerParseString, ok := apomock.GetRegisteredFunc("gocql.murmur3Partitioner.ParseString")
	if ok {
		reta = FuncAuxMockmurmur3PartitionerParseString.(func(recvp murmur3Partitioner, argstr string) (reta token))(recvp, argstr)
	} else {
		panic("FuncAuxMockmurmur3PartitionerParseString ")
	}
	AuxMockIncrementRecorderAuxMockmurmur3PartitionerParseString()
	return
}

//
// Mock: (recvm murmur3Token)String()(reta string)
//

type MockArgsTypemurmur3TokenString struct {
	ApomockCallNumber int
}

var LastMockArgsmurmur3TokenString MockArgsTypemurmur3TokenString

// (recvm murmur3Token)AuxMockString()(reta string) - Generated mock function
func (recvm murmur3Token) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.murmur3Token.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.murmur3Token.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.murmur3Token.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockmurmur3TokenString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmurmur3TokenString int = 0

var condRecorderAuxMockmurmur3TokenString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmurmur3TokenString(i int) {
	condRecorderAuxMockmurmur3TokenString.L.Lock()
	for recorderAuxMockmurmur3TokenString < i {
		condRecorderAuxMockmurmur3TokenString.Wait()
	}
	condRecorderAuxMockmurmur3TokenString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmurmur3TokenString() {
	condRecorderAuxMockmurmur3TokenString.L.Lock()
	recorderAuxMockmurmur3TokenString++
	condRecorderAuxMockmurmur3TokenString.L.Unlock()
	condRecorderAuxMockmurmur3TokenString.Broadcast()
}
func AuxMockGetRecorderAuxMockmurmur3TokenString() (ret int) {
	condRecorderAuxMockmurmur3TokenString.L.Lock()
	ret = recorderAuxMockmurmur3TokenString
	condRecorderAuxMockmurmur3TokenString.L.Unlock()
	return
}

// (recvm murmur3Token)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvm murmur3Token) String() (reta string) {
	FuncAuxMockmurmur3TokenString, ok := apomock.GetRegisteredFunc("gocql.murmur3Token.String")
	if ok {
		reta = FuncAuxMockmurmur3TokenString.(func(recvm murmur3Token) (reta string))(recvm)
	} else {
		panic("FuncAuxMockmurmur3TokenString ")
	}
	AuxMockIncrementRecorderAuxMockmurmur3TokenString()
	return
}

//
// Mock: (recvr randomPartitioner)Name()(reta string)
//

type MockArgsTyperandomPartitionerName struct {
	ApomockCallNumber int
}

var LastMockArgsrandomPartitionerName MockArgsTyperandomPartitionerName

// (recvr randomPartitioner)AuxMockName()(reta string) - Generated mock function
func (recvr randomPartitioner) AuxMockName() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.randomPartitioner.Name")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.randomPartitioner.Name")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.randomPartitioner.Name")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockrandomPartitionerName  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockrandomPartitionerName int = 0

var condRecorderAuxMockrandomPartitionerName *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockrandomPartitionerName(i int) {
	condRecorderAuxMockrandomPartitionerName.L.Lock()
	for recorderAuxMockrandomPartitionerName < i {
		condRecorderAuxMockrandomPartitionerName.Wait()
	}
	condRecorderAuxMockrandomPartitionerName.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockrandomPartitionerName() {
	condRecorderAuxMockrandomPartitionerName.L.Lock()
	recorderAuxMockrandomPartitionerName++
	condRecorderAuxMockrandomPartitionerName.L.Unlock()
	condRecorderAuxMockrandomPartitionerName.Broadcast()
}
func AuxMockGetRecorderAuxMockrandomPartitionerName() (ret int) {
	condRecorderAuxMockrandomPartitionerName.L.Lock()
	ret = recorderAuxMockrandomPartitionerName
	condRecorderAuxMockrandomPartitionerName.L.Unlock()
	return
}

// (recvr randomPartitioner)Name - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr randomPartitioner) Name() (reta string) {
	FuncAuxMockrandomPartitionerName, ok := apomock.GetRegisteredFunc("gocql.randomPartitioner.Name")
	if ok {
		reta = FuncAuxMockrandomPartitionerName.(func(recvr randomPartitioner) (reta string))(recvr)
	} else {
		panic("FuncAuxMockrandomPartitionerName ")
	}
	AuxMockIncrementRecorderAuxMockrandomPartitionerName()
	return
}

//
// Mock: (recvr *randomToken)String()(reta string)
//

type MockArgsTyperandomTokenString struct {
	ApomockCallNumber int
}

var LastMockArgsrandomTokenString MockArgsTyperandomTokenString

// (recvr *randomToken)AuxMockString()(reta string) - Generated mock function
func (recvr *randomToken) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.randomToken.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.randomToken.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.randomToken.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrrandomTokenString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrrandomTokenString int = 0

var condRecorderAuxMockPtrrandomTokenString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrrandomTokenString(i int) {
	condRecorderAuxMockPtrrandomTokenString.L.Lock()
	for recorderAuxMockPtrrandomTokenString < i {
		condRecorderAuxMockPtrrandomTokenString.Wait()
	}
	condRecorderAuxMockPtrrandomTokenString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrrandomTokenString() {
	condRecorderAuxMockPtrrandomTokenString.L.Lock()
	recorderAuxMockPtrrandomTokenString++
	condRecorderAuxMockPtrrandomTokenString.L.Unlock()
	condRecorderAuxMockPtrrandomTokenString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrrandomTokenString() (ret int) {
	condRecorderAuxMockPtrrandomTokenString.L.Lock()
	ret = recorderAuxMockPtrrandomTokenString
	condRecorderAuxMockPtrrandomTokenString.L.Unlock()
	return
}

// (recvr *randomToken)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *randomToken) String() (reta string) {
	FuncAuxMockPtrrandomTokenString, ok := apomock.GetRegisteredFunc("gocql.randomToken.String")
	if ok {
		reta = FuncAuxMockPtrrandomTokenString.(func(recvr *randomToken) (reta string))(recvr)
	} else {
		panic("FuncAuxMockPtrrandomTokenString ")
	}
	AuxMockIncrementRecorderAuxMockPtrrandomTokenString()
	return
}

//
// Mock: (recvr *randomToken)Less(argtoken token)(reta bool)
//

type MockArgsTyperandomTokenLess struct {
	ApomockCallNumber int
	Argtoken          token
}

var LastMockArgsrandomTokenLess MockArgsTyperandomTokenLess

// (recvr *randomToken)AuxMockLess(argtoken token)(reta bool) - Generated mock function
func (recvr *randomToken) AuxMockLess(argtoken token) (reta bool) {
	LastMockArgsrandomTokenLess = MockArgsTyperandomTokenLess{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrrandomTokenLess(),
		Argtoken:          argtoken,
	}
	rargs, rerr := apomock.GetNext("gocql.randomToken.Less")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.randomToken.Less")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.randomToken.Less")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrrandomTokenLess  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrrandomTokenLess int = 0

var condRecorderAuxMockPtrrandomTokenLess *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrrandomTokenLess(i int) {
	condRecorderAuxMockPtrrandomTokenLess.L.Lock()
	for recorderAuxMockPtrrandomTokenLess < i {
		condRecorderAuxMockPtrrandomTokenLess.Wait()
	}
	condRecorderAuxMockPtrrandomTokenLess.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrrandomTokenLess() {
	condRecorderAuxMockPtrrandomTokenLess.L.Lock()
	recorderAuxMockPtrrandomTokenLess++
	condRecorderAuxMockPtrrandomTokenLess.L.Unlock()
	condRecorderAuxMockPtrrandomTokenLess.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrrandomTokenLess() (ret int) {
	condRecorderAuxMockPtrrandomTokenLess.L.Lock()
	ret = recorderAuxMockPtrrandomTokenLess
	condRecorderAuxMockPtrrandomTokenLess.L.Unlock()
	return
}

// (recvr *randomToken)Less - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *randomToken) Less(argtoken token) (reta bool) {
	FuncAuxMockPtrrandomTokenLess, ok := apomock.GetRegisteredFunc("gocql.randomToken.Less")
	if ok {
		reta = FuncAuxMockPtrrandomTokenLess.(func(recvr *randomToken, argtoken token) (reta bool))(recvr, argtoken)
	} else {
		panic("FuncAuxMockPtrrandomTokenLess ")
	}
	AuxMockIncrementRecorderAuxMockPtrrandomTokenLess()
	return
}

//
// Mock: (recvt *tokenRing)Swap(argi int, argj int)()
//

type MockArgsTypetokenRingSwap struct {
	ApomockCallNumber int
	Argi              int
	Argj              int
}

var LastMockArgstokenRingSwap MockArgsTypetokenRingSwap

// (recvt *tokenRing)AuxMockSwap(argi int, argj int)() - Generated mock function
func (recvt *tokenRing) AuxMockSwap(argi int, argj int) {
	LastMockArgstokenRingSwap = MockArgsTypetokenRingSwap{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrtokenRingSwap(),
		Argi:              argi,
		Argj:              argj,
	}
	return
}

// RecorderAuxMockPtrtokenRingSwap  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenRingSwap int = 0

var condRecorderAuxMockPtrtokenRingSwap *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenRingSwap(i int) {
	condRecorderAuxMockPtrtokenRingSwap.L.Lock()
	for recorderAuxMockPtrtokenRingSwap < i {
		condRecorderAuxMockPtrtokenRingSwap.Wait()
	}
	condRecorderAuxMockPtrtokenRingSwap.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenRingSwap() {
	condRecorderAuxMockPtrtokenRingSwap.L.Lock()
	recorderAuxMockPtrtokenRingSwap++
	condRecorderAuxMockPtrtokenRingSwap.L.Unlock()
	condRecorderAuxMockPtrtokenRingSwap.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenRingSwap() (ret int) {
	condRecorderAuxMockPtrtokenRingSwap.L.Lock()
	ret = recorderAuxMockPtrtokenRingSwap
	condRecorderAuxMockPtrtokenRingSwap.L.Unlock()
	return
}

// (recvt *tokenRing)Swap - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenRing) Swap(argi int, argj int) {
	FuncAuxMockPtrtokenRingSwap, ok := apomock.GetRegisteredFunc("gocql.tokenRing.Swap")
	if ok {
		FuncAuxMockPtrtokenRingSwap.(func(recvt *tokenRing, argi int, argj int))(recvt, argi, argj)
	} else {
		panic("FuncAuxMockPtrtokenRingSwap ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenRingSwap()
	return
}

//
// Mock: (recvt *tokenRing)GetHostForPartitionKey(argpartitionKey []byte)(reta *HostInfo)
//

type MockArgsTypetokenRingGetHostForPartitionKey struct {
	ApomockCallNumber int
	ArgpartitionKey   []byte
}

var LastMockArgstokenRingGetHostForPartitionKey MockArgsTypetokenRingGetHostForPartitionKey

// (recvt *tokenRing)AuxMockGetHostForPartitionKey(argpartitionKey []byte)(reta *HostInfo) - Generated mock function
func (recvt *tokenRing) AuxMockGetHostForPartitionKey(argpartitionKey []byte) (reta *HostInfo) {
	LastMockArgstokenRingGetHostForPartitionKey = MockArgsTypetokenRingGetHostForPartitionKey{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrtokenRingGetHostForPartitionKey(),
		ArgpartitionKey:   argpartitionKey,
	}
	rargs, rerr := apomock.GetNext("gocql.tokenRing.GetHostForPartitionKey")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.tokenRing.GetHostForPartitionKey")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.tokenRing.GetHostForPartitionKey")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*HostInfo)
	}
	return
}

// RecorderAuxMockPtrtokenRingGetHostForPartitionKey  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenRingGetHostForPartitionKey int = 0

var condRecorderAuxMockPtrtokenRingGetHostForPartitionKey *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenRingGetHostForPartitionKey(i int) {
	condRecorderAuxMockPtrtokenRingGetHostForPartitionKey.L.Lock()
	for recorderAuxMockPtrtokenRingGetHostForPartitionKey < i {
		condRecorderAuxMockPtrtokenRingGetHostForPartitionKey.Wait()
	}
	condRecorderAuxMockPtrtokenRingGetHostForPartitionKey.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenRingGetHostForPartitionKey() {
	condRecorderAuxMockPtrtokenRingGetHostForPartitionKey.L.Lock()
	recorderAuxMockPtrtokenRingGetHostForPartitionKey++
	condRecorderAuxMockPtrtokenRingGetHostForPartitionKey.L.Unlock()
	condRecorderAuxMockPtrtokenRingGetHostForPartitionKey.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenRingGetHostForPartitionKey() (ret int) {
	condRecorderAuxMockPtrtokenRingGetHostForPartitionKey.L.Lock()
	ret = recorderAuxMockPtrtokenRingGetHostForPartitionKey
	condRecorderAuxMockPtrtokenRingGetHostForPartitionKey.L.Unlock()
	return
}

// (recvt *tokenRing)GetHostForPartitionKey - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenRing) GetHostForPartitionKey(argpartitionKey []byte) (reta *HostInfo) {
	FuncAuxMockPtrtokenRingGetHostForPartitionKey, ok := apomock.GetRegisteredFunc("gocql.tokenRing.GetHostForPartitionKey")
	if ok {
		reta = FuncAuxMockPtrtokenRingGetHostForPartitionKey.(func(recvt *tokenRing, argpartitionKey []byte) (reta *HostInfo))(recvt, argpartitionKey)
	} else {
		panic("FuncAuxMockPtrtokenRingGetHostForPartitionKey ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenRingGetHostForPartitionKey()
	return
}

//
// Mock: (recvp murmur3Partitioner)Name()(reta string)
//

type MockArgsTypemurmur3PartitionerName struct {
	ApomockCallNumber int
}

var LastMockArgsmurmur3PartitionerName MockArgsTypemurmur3PartitionerName

// (recvp murmur3Partitioner)AuxMockName()(reta string) - Generated mock function
func (recvp murmur3Partitioner) AuxMockName() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.murmur3Partitioner.Name")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.murmur3Partitioner.Name")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.murmur3Partitioner.Name")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockmurmur3PartitionerName  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockmurmur3PartitionerName int = 0

var condRecorderAuxMockmurmur3PartitionerName *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockmurmur3PartitionerName(i int) {
	condRecorderAuxMockmurmur3PartitionerName.L.Lock()
	for recorderAuxMockmurmur3PartitionerName < i {
		condRecorderAuxMockmurmur3PartitionerName.Wait()
	}
	condRecorderAuxMockmurmur3PartitionerName.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockmurmur3PartitionerName() {
	condRecorderAuxMockmurmur3PartitionerName.L.Lock()
	recorderAuxMockmurmur3PartitionerName++
	condRecorderAuxMockmurmur3PartitionerName.L.Unlock()
	condRecorderAuxMockmurmur3PartitionerName.Broadcast()
}
func AuxMockGetRecorderAuxMockmurmur3PartitionerName() (ret int) {
	condRecorderAuxMockmurmur3PartitionerName.L.Lock()
	ret = recorderAuxMockmurmur3PartitionerName
	condRecorderAuxMockmurmur3PartitionerName.L.Unlock()
	return
}

// (recvp murmur3Partitioner)Name - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp murmur3Partitioner) Name() (reta string) {
	FuncAuxMockmurmur3PartitionerName, ok := apomock.GetRegisteredFunc("gocql.murmur3Partitioner.Name")
	if ok {
		reta = FuncAuxMockmurmur3PartitionerName.(func(recvp murmur3Partitioner) (reta string))(recvp)
	} else {
		panic("FuncAuxMockmurmur3PartitionerName ")
	}
	AuxMockIncrementRecorderAuxMockmurmur3PartitionerName()
	return
}

//
// Mock: (recvp orderedPartitioner)Name()(reta string)
//

type MockArgsTypeorderedPartitionerName struct {
	ApomockCallNumber int
}

var LastMockArgsorderedPartitionerName MockArgsTypeorderedPartitionerName

// (recvp orderedPartitioner)AuxMockName()(reta string) - Generated mock function
func (recvp orderedPartitioner) AuxMockName() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.orderedPartitioner.Name")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.orderedPartitioner.Name")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.orderedPartitioner.Name")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockorderedPartitionerName  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockorderedPartitionerName int = 0

var condRecorderAuxMockorderedPartitionerName *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockorderedPartitionerName(i int) {
	condRecorderAuxMockorderedPartitionerName.L.Lock()
	for recorderAuxMockorderedPartitionerName < i {
		condRecorderAuxMockorderedPartitionerName.Wait()
	}
	condRecorderAuxMockorderedPartitionerName.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockorderedPartitionerName() {
	condRecorderAuxMockorderedPartitionerName.L.Lock()
	recorderAuxMockorderedPartitionerName++
	condRecorderAuxMockorderedPartitionerName.L.Unlock()
	condRecorderAuxMockorderedPartitionerName.Broadcast()
}
func AuxMockGetRecorderAuxMockorderedPartitionerName() (ret int) {
	condRecorderAuxMockorderedPartitionerName.L.Lock()
	ret = recorderAuxMockorderedPartitionerName
	condRecorderAuxMockorderedPartitionerName.L.Unlock()
	return
}

// (recvp orderedPartitioner)Name - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp orderedPartitioner) Name() (reta string) {
	FuncAuxMockorderedPartitionerName, ok := apomock.GetRegisteredFunc("gocql.orderedPartitioner.Name")
	if ok {
		reta = FuncAuxMockorderedPartitionerName.(func(recvp orderedPartitioner) (reta string))(recvp)
	} else {
		panic("FuncAuxMockorderedPartitionerName ")
	}
	AuxMockIncrementRecorderAuxMockorderedPartitionerName()
	return
}

//
// Mock: (recvp randomPartitioner)Hash(argpartitionKey []byte)(reta token)
//

type MockArgsTyperandomPartitionerHash struct {
	ApomockCallNumber int
	ArgpartitionKey   []byte
}

var LastMockArgsrandomPartitionerHash MockArgsTyperandomPartitionerHash

// (recvp randomPartitioner)AuxMockHash(argpartitionKey []byte)(reta token) - Generated mock function
func (recvp randomPartitioner) AuxMockHash(argpartitionKey []byte) (reta token) {
	LastMockArgsrandomPartitionerHash = MockArgsTyperandomPartitionerHash{
		ApomockCallNumber: AuxMockGetRecorderAuxMockrandomPartitionerHash(),
		ArgpartitionKey:   argpartitionKey,
	}
	rargs, rerr := apomock.GetNext("gocql.randomPartitioner.Hash")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.randomPartitioner.Hash")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.randomPartitioner.Hash")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(token)
	}
	return
}

// RecorderAuxMockrandomPartitionerHash  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockrandomPartitionerHash int = 0

var condRecorderAuxMockrandomPartitionerHash *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockrandomPartitionerHash(i int) {
	condRecorderAuxMockrandomPartitionerHash.L.Lock()
	for recorderAuxMockrandomPartitionerHash < i {
		condRecorderAuxMockrandomPartitionerHash.Wait()
	}
	condRecorderAuxMockrandomPartitionerHash.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockrandomPartitionerHash() {
	condRecorderAuxMockrandomPartitionerHash.L.Lock()
	recorderAuxMockrandomPartitionerHash++
	condRecorderAuxMockrandomPartitionerHash.L.Unlock()
	condRecorderAuxMockrandomPartitionerHash.Broadcast()
}
func AuxMockGetRecorderAuxMockrandomPartitionerHash() (ret int) {
	condRecorderAuxMockrandomPartitionerHash.L.Lock()
	ret = recorderAuxMockrandomPartitionerHash
	condRecorderAuxMockrandomPartitionerHash.L.Unlock()
	return
}

// (recvp randomPartitioner)Hash - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp randomPartitioner) Hash(argpartitionKey []byte) (reta token) {
	FuncAuxMockrandomPartitionerHash, ok := apomock.GetRegisteredFunc("gocql.randomPartitioner.Hash")
	if ok {
		reta = FuncAuxMockrandomPartitionerHash.(func(recvp randomPartitioner, argpartitionKey []byte) (reta token))(recvp, argpartitionKey)
	} else {
		panic("FuncAuxMockrandomPartitionerHash ")
	}
	AuxMockIncrementRecorderAuxMockrandomPartitionerHash()
	return
}

//
// Mock: newTokenRing(argpartitioner string, arghosts []*HostInfo)(reta *tokenRing, retb error)
//

type MockArgsTypenewTokenRing struct {
	ApomockCallNumber int
	Argpartitioner    string
	Arghosts          []*HostInfo
}

var LastMockArgsnewTokenRing MockArgsTypenewTokenRing

// AuxMocknewTokenRing(argpartitioner string, arghosts []*HostInfo)(reta *tokenRing, retb error) - Generated mock function
func AuxMocknewTokenRing(argpartitioner string, arghosts []*HostInfo) (reta *tokenRing, retb error) {
	LastMockArgsnewTokenRing = MockArgsTypenewTokenRing{
		ApomockCallNumber: AuxMockGetRecorderAuxMocknewTokenRing(),
		Argpartitioner:    argpartitioner,
		Arghosts:          arghosts,
	}
	rargs, rerr := apomock.GetNext("gocql.newTokenRing")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.newTokenRing")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.newTokenRing")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*tokenRing)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMocknewTokenRing  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocknewTokenRing int = 0

var condRecorderAuxMocknewTokenRing *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocknewTokenRing(i int) {
	condRecorderAuxMocknewTokenRing.L.Lock()
	for recorderAuxMocknewTokenRing < i {
		condRecorderAuxMocknewTokenRing.Wait()
	}
	condRecorderAuxMocknewTokenRing.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocknewTokenRing() {
	condRecorderAuxMocknewTokenRing.L.Lock()
	recorderAuxMocknewTokenRing++
	condRecorderAuxMocknewTokenRing.L.Unlock()
	condRecorderAuxMocknewTokenRing.Broadcast()
}
func AuxMockGetRecorderAuxMocknewTokenRing() (ret int) {
	condRecorderAuxMocknewTokenRing.L.Lock()
	ret = recorderAuxMocknewTokenRing
	condRecorderAuxMocknewTokenRing.L.Unlock()
	return
}

// newTokenRing - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func newTokenRing(argpartitioner string, arghosts []*HostInfo) (reta *tokenRing, retb error) {
	FuncAuxMocknewTokenRing, ok := apomock.GetRegisteredFunc("gocql.newTokenRing")
	if ok {
		reta, retb = FuncAuxMocknewTokenRing.(func(argpartitioner string, arghosts []*HostInfo) (reta *tokenRing, retb error))(argpartitioner, arghosts)
	} else {
		panic("FuncAuxMocknewTokenRing ")
	}
	AuxMockIncrementRecorderAuxMocknewTokenRing()
	return
}

//
// Mock: (recvt *tokenRing)Less(argi int, argj int)(reta bool)
//

type MockArgsTypetokenRingLess struct {
	ApomockCallNumber int
	Argi              int
	Argj              int
}

var LastMockArgstokenRingLess MockArgsTypetokenRingLess

// (recvt *tokenRing)AuxMockLess(argi int, argj int)(reta bool) - Generated mock function
func (recvt *tokenRing) AuxMockLess(argi int, argj int) (reta bool) {
	LastMockArgstokenRingLess = MockArgsTypetokenRingLess{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrtokenRingLess(),
		Argi:              argi,
		Argj:              argj,
	}
	rargs, rerr := apomock.GetNext("gocql.tokenRing.Less")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.tokenRing.Less")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.tokenRing.Less")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockPtrtokenRingLess  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenRingLess int = 0

var condRecorderAuxMockPtrtokenRingLess *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenRingLess(i int) {
	condRecorderAuxMockPtrtokenRingLess.L.Lock()
	for recorderAuxMockPtrtokenRingLess < i {
		condRecorderAuxMockPtrtokenRingLess.Wait()
	}
	condRecorderAuxMockPtrtokenRingLess.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenRingLess() {
	condRecorderAuxMockPtrtokenRingLess.L.Lock()
	recorderAuxMockPtrtokenRingLess++
	condRecorderAuxMockPtrtokenRingLess.L.Unlock()
	condRecorderAuxMockPtrtokenRingLess.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenRingLess() (ret int) {
	condRecorderAuxMockPtrtokenRingLess.L.Lock()
	ret = recorderAuxMockPtrtokenRingLess
	condRecorderAuxMockPtrtokenRingLess.L.Unlock()
	return
}

// (recvt *tokenRing)Less - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenRing) Less(argi int, argj int) (reta bool) {
	FuncAuxMockPtrtokenRingLess, ok := apomock.GetRegisteredFunc("gocql.tokenRing.Less")
	if ok {
		reta = FuncAuxMockPtrtokenRingLess.(func(recvt *tokenRing, argi int, argj int) (reta bool))(recvt, argi, argj)
	} else {
		panic("FuncAuxMockPtrtokenRingLess ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenRingLess()
	return
}

//
// Mock: (recvt *tokenRing)String()(reta string)
//

type MockArgsTypetokenRingString struct {
	ApomockCallNumber int
}

var LastMockArgstokenRingString MockArgsTypetokenRingString

// (recvt *tokenRing)AuxMockString()(reta string) - Generated mock function
func (recvt *tokenRing) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.tokenRing.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.tokenRing.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.tokenRing.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrtokenRingString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtokenRingString int = 0

var condRecorderAuxMockPtrtokenRingString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtokenRingString(i int) {
	condRecorderAuxMockPtrtokenRingString.L.Lock()
	for recorderAuxMockPtrtokenRingString < i {
		condRecorderAuxMockPtrtokenRingString.Wait()
	}
	condRecorderAuxMockPtrtokenRingString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtokenRingString() {
	condRecorderAuxMockPtrtokenRingString.L.Lock()
	recorderAuxMockPtrtokenRingString++
	condRecorderAuxMockPtrtokenRingString.L.Unlock()
	condRecorderAuxMockPtrtokenRingString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtokenRingString() (ret int) {
	condRecorderAuxMockPtrtokenRingString.L.Lock()
	ret = recorderAuxMockPtrtokenRingString
	condRecorderAuxMockPtrtokenRingString.L.Unlock()
	return
}

// (recvt *tokenRing)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *tokenRing) String() (reta string) {
	FuncAuxMockPtrtokenRingString, ok := apomock.GetRegisteredFunc("gocql.tokenRing.String")
	if ok {
		reta = FuncAuxMockPtrtokenRingString.(func(recvt *tokenRing) (reta string))(recvt)
	} else {
		panic("FuncAuxMockPtrtokenRingString ")
	}
	AuxMockIncrementRecorderAuxMockPtrtokenRingString()
	return
}
