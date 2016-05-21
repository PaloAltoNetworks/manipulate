// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

import "time"

func init() {

	apomock.RegisterFunc("gocql", "gocql.UUID.Time", (UUID).AuxMockTime)
	apomock.RegisterFunc("gocql", "gocql.UUID.UnmarshalJSON", (*UUID).AuxMockUnmarshalJSON)
	apomock.RegisterFunc("gocql", "gocql.UUID.UnmarshalText", (*UUID).AuxMockUnmarshalText)
	apomock.RegisterFunc("gocql", "gocql.UUID.MarshalJSON", (UUID).AuxMockMarshalJSON)
	apomock.RegisterFunc("gocql", "gocql.UUID.MarshalText", (UUID).AuxMockMarshalText)
	apomock.RegisterFunc("gocql", "gocql.ParseUUID", AuxMockParseUUID)
	apomock.RegisterFunc("gocql", "gocql.UUIDFromTime", AuxMockUUIDFromTime)
	apomock.RegisterFunc("gocql", "gocql.UUID.Variant", (UUID).AuxMockVariant)
	apomock.RegisterFunc("gocql", "gocql.UUID.Version", (UUID).AuxMockVersion)
	apomock.RegisterFunc("gocql", "gocql.UUID.Node", (UUID).AuxMockNode)
	apomock.RegisterFunc("gocql", "gocql.UUID.Timestamp", (UUID).AuxMockTimestamp)
	apomock.RegisterFunc("gocql", "gocql.UUIDFromBytes", AuxMockUUIDFromBytes)
	apomock.RegisterFunc("gocql", "gocql.RandomUUID", AuxMockRandomUUID)
	apomock.RegisterFunc("gocql", "gocql.TimeUUID", AuxMockTimeUUID)
	apomock.RegisterFunc("gocql", "gocql.UUID.String", (UUID).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.UUID.Bytes", (UUID).AuxMockBytes)
}

const (
	VariantNCSCompat = 0
	VariantIETF      = 2
	VariantMicrosoft = 6
	VariantFuture    = 7
)

const ()

var hardwareAddr []byte

var clockSeq uint32

var timeBase = time.Date(1582, time.October, 15, 0, 0, 0, 0, time.UTC).Unix()

//
// Internal Types: in this package and their exportable versions
//

//
// External Types: in this package
//
type UUID [16]byte

//
// Mock: (recvu UUID)Time()(reta time.Time)
//

type MockArgsTypeUUIDTime struct {
	ApomockCallNumber int
}

var LastMockArgsUUIDTime MockArgsTypeUUIDTime

// (recvu UUID)AuxMockTime()(reta time.Time) - Generated mock function
func (recvu UUID) AuxMockTime() (reta time.Time) {
	rargs, rerr := apomock.GetNext("gocql.UUID.Time")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUID.Time")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UUID.Time")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(time.Time)
	}
	return
}

// RecorderAuxMockUUIDTime  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUUIDTime int = 0

var condRecorderAuxMockUUIDTime *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUUIDTime(i int) {
	condRecorderAuxMockUUIDTime.L.Lock()
	for recorderAuxMockUUIDTime < i {
		condRecorderAuxMockUUIDTime.Wait()
	}
	condRecorderAuxMockUUIDTime.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUUIDTime() {
	condRecorderAuxMockUUIDTime.L.Lock()
	recorderAuxMockUUIDTime++
	condRecorderAuxMockUUIDTime.L.Unlock()
	condRecorderAuxMockUUIDTime.Broadcast()
}
func AuxMockGetRecorderAuxMockUUIDTime() (ret int) {
	condRecorderAuxMockUUIDTime.L.Lock()
	ret = recorderAuxMockUUIDTime
	condRecorderAuxMockUUIDTime.L.Unlock()
	return
}

// (recvu UUID)Time - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu UUID) Time() (reta time.Time) {
	FuncAuxMockUUIDTime, ok := apomock.GetRegisteredFunc("gocql.UUID.Time")
	if ok {
		reta = FuncAuxMockUUIDTime.(func(recvu UUID) (reta time.Time))(recvu)
	} else {
		panic("FuncAuxMockUUIDTime ")
	}
	AuxMockIncrementRecorderAuxMockUUIDTime()
	return
}

//
// Mock: (recvu *UUID)UnmarshalJSON(argdata []byte)(reta error)
//

type MockArgsTypeUUIDUnmarshalJSON struct {
	ApomockCallNumber int
	Argdata           []byte
}

var LastMockArgsUUIDUnmarshalJSON MockArgsTypeUUIDUnmarshalJSON

// (recvu *UUID)AuxMockUnmarshalJSON(argdata []byte)(reta error) - Generated mock function
func (recvu *UUID) AuxMockUnmarshalJSON(argdata []byte) (reta error) {
	LastMockArgsUUIDUnmarshalJSON = MockArgsTypeUUIDUnmarshalJSON{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrUUIDUnmarshalJSON(),
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.UUID.UnmarshalJSON")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUID.UnmarshalJSON")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UUID.UnmarshalJSON")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrUUIDUnmarshalJSON  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrUUIDUnmarshalJSON int = 0

var condRecorderAuxMockPtrUUIDUnmarshalJSON *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrUUIDUnmarshalJSON(i int) {
	condRecorderAuxMockPtrUUIDUnmarshalJSON.L.Lock()
	for recorderAuxMockPtrUUIDUnmarshalJSON < i {
		condRecorderAuxMockPtrUUIDUnmarshalJSON.Wait()
	}
	condRecorderAuxMockPtrUUIDUnmarshalJSON.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrUUIDUnmarshalJSON() {
	condRecorderAuxMockPtrUUIDUnmarshalJSON.L.Lock()
	recorderAuxMockPtrUUIDUnmarshalJSON++
	condRecorderAuxMockPtrUUIDUnmarshalJSON.L.Unlock()
	condRecorderAuxMockPtrUUIDUnmarshalJSON.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrUUIDUnmarshalJSON() (ret int) {
	condRecorderAuxMockPtrUUIDUnmarshalJSON.L.Lock()
	ret = recorderAuxMockPtrUUIDUnmarshalJSON
	condRecorderAuxMockPtrUUIDUnmarshalJSON.L.Unlock()
	return
}

// (recvu *UUID)UnmarshalJSON - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu *UUID) UnmarshalJSON(argdata []byte) (reta error) {
	FuncAuxMockPtrUUIDUnmarshalJSON, ok := apomock.GetRegisteredFunc("gocql.UUID.UnmarshalJSON")
	if ok {
		reta = FuncAuxMockPtrUUIDUnmarshalJSON.(func(recvu *UUID, argdata []byte) (reta error))(recvu, argdata)
	} else {
		panic("FuncAuxMockPtrUUIDUnmarshalJSON ")
	}
	AuxMockIncrementRecorderAuxMockPtrUUIDUnmarshalJSON()
	return
}

//
// Mock: (recvu *UUID)UnmarshalText(argtext []byte)(reterr error)
//

type MockArgsTypeUUIDUnmarshalText struct {
	ApomockCallNumber int
	Argtext           []byte
}

var LastMockArgsUUIDUnmarshalText MockArgsTypeUUIDUnmarshalText

// (recvu *UUID)AuxMockUnmarshalText(argtext []byte)(reterr error) - Generated mock function
func (recvu *UUID) AuxMockUnmarshalText(argtext []byte) (reterr error) {
	LastMockArgsUUIDUnmarshalText = MockArgsTypeUUIDUnmarshalText{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrUUIDUnmarshalText(),
		Argtext:           argtext,
	}
	rargs, rerr := apomock.GetNext("gocql.UUID.UnmarshalText")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUID.UnmarshalText")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UUID.UnmarshalText")
	}
	if rargs.GetArg(0) != nil {
		reterr = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrUUIDUnmarshalText  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrUUIDUnmarshalText int = 0

var condRecorderAuxMockPtrUUIDUnmarshalText *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrUUIDUnmarshalText(i int) {
	condRecorderAuxMockPtrUUIDUnmarshalText.L.Lock()
	for recorderAuxMockPtrUUIDUnmarshalText < i {
		condRecorderAuxMockPtrUUIDUnmarshalText.Wait()
	}
	condRecorderAuxMockPtrUUIDUnmarshalText.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrUUIDUnmarshalText() {
	condRecorderAuxMockPtrUUIDUnmarshalText.L.Lock()
	recorderAuxMockPtrUUIDUnmarshalText++
	condRecorderAuxMockPtrUUIDUnmarshalText.L.Unlock()
	condRecorderAuxMockPtrUUIDUnmarshalText.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrUUIDUnmarshalText() (ret int) {
	condRecorderAuxMockPtrUUIDUnmarshalText.L.Lock()
	ret = recorderAuxMockPtrUUIDUnmarshalText
	condRecorderAuxMockPtrUUIDUnmarshalText.L.Unlock()
	return
}

// (recvu *UUID)UnmarshalText - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu *UUID) UnmarshalText(argtext []byte) (reterr error) {
	FuncAuxMockPtrUUIDUnmarshalText, ok := apomock.GetRegisteredFunc("gocql.UUID.UnmarshalText")
	if ok {
		reterr = FuncAuxMockPtrUUIDUnmarshalText.(func(recvu *UUID, argtext []byte) (reterr error))(recvu, argtext)
	} else {
		panic("FuncAuxMockPtrUUIDUnmarshalText ")
	}
	AuxMockIncrementRecorderAuxMockPtrUUIDUnmarshalText()
	return
}

//
// Mock: (recvu UUID)MarshalJSON()(reta []byte, retb error)
//

type MockArgsTypeUUIDMarshalJSON struct {
	ApomockCallNumber int
}

var LastMockArgsUUIDMarshalJSON MockArgsTypeUUIDMarshalJSON

// (recvu UUID)AuxMockMarshalJSON()(reta []byte, retb error) - Generated mock function
func (recvu UUID) AuxMockMarshalJSON() (reta []byte, retb error) {
	rargs, rerr := apomock.GetNext("gocql.UUID.MarshalJSON")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUID.MarshalJSON")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.UUID.MarshalJSON")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockUUIDMarshalJSON  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUUIDMarshalJSON int = 0

var condRecorderAuxMockUUIDMarshalJSON *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUUIDMarshalJSON(i int) {
	condRecorderAuxMockUUIDMarshalJSON.L.Lock()
	for recorderAuxMockUUIDMarshalJSON < i {
		condRecorderAuxMockUUIDMarshalJSON.Wait()
	}
	condRecorderAuxMockUUIDMarshalJSON.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUUIDMarshalJSON() {
	condRecorderAuxMockUUIDMarshalJSON.L.Lock()
	recorderAuxMockUUIDMarshalJSON++
	condRecorderAuxMockUUIDMarshalJSON.L.Unlock()
	condRecorderAuxMockUUIDMarshalJSON.Broadcast()
}
func AuxMockGetRecorderAuxMockUUIDMarshalJSON() (ret int) {
	condRecorderAuxMockUUIDMarshalJSON.L.Lock()
	ret = recorderAuxMockUUIDMarshalJSON
	condRecorderAuxMockUUIDMarshalJSON.L.Unlock()
	return
}

// (recvu UUID)MarshalJSON - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu UUID) MarshalJSON() (reta []byte, retb error) {
	FuncAuxMockUUIDMarshalJSON, ok := apomock.GetRegisteredFunc("gocql.UUID.MarshalJSON")
	if ok {
		reta, retb = FuncAuxMockUUIDMarshalJSON.(func(recvu UUID) (reta []byte, retb error))(recvu)
	} else {
		panic("FuncAuxMockUUIDMarshalJSON ")
	}
	AuxMockIncrementRecorderAuxMockUUIDMarshalJSON()
	return
}

//
// Mock: (recvu UUID)MarshalText()(reta []byte, retb error)
//

type MockArgsTypeUUIDMarshalText struct {
	ApomockCallNumber int
}

var LastMockArgsUUIDMarshalText MockArgsTypeUUIDMarshalText

// (recvu UUID)AuxMockMarshalText()(reta []byte, retb error) - Generated mock function
func (recvu UUID) AuxMockMarshalText() (reta []byte, retb error) {
	rargs, rerr := apomock.GetNext("gocql.UUID.MarshalText")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUID.MarshalText")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.UUID.MarshalText")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockUUIDMarshalText  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUUIDMarshalText int = 0

var condRecorderAuxMockUUIDMarshalText *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUUIDMarshalText(i int) {
	condRecorderAuxMockUUIDMarshalText.L.Lock()
	for recorderAuxMockUUIDMarshalText < i {
		condRecorderAuxMockUUIDMarshalText.Wait()
	}
	condRecorderAuxMockUUIDMarshalText.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUUIDMarshalText() {
	condRecorderAuxMockUUIDMarshalText.L.Lock()
	recorderAuxMockUUIDMarshalText++
	condRecorderAuxMockUUIDMarshalText.L.Unlock()
	condRecorderAuxMockUUIDMarshalText.Broadcast()
}
func AuxMockGetRecorderAuxMockUUIDMarshalText() (ret int) {
	condRecorderAuxMockUUIDMarshalText.L.Lock()
	ret = recorderAuxMockUUIDMarshalText
	condRecorderAuxMockUUIDMarshalText.L.Unlock()
	return
}

// (recvu UUID)MarshalText - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu UUID) MarshalText() (reta []byte, retb error) {
	FuncAuxMockUUIDMarshalText, ok := apomock.GetRegisteredFunc("gocql.UUID.MarshalText")
	if ok {
		reta, retb = FuncAuxMockUUIDMarshalText.(func(recvu UUID) (reta []byte, retb error))(recvu)
	} else {
		panic("FuncAuxMockUUIDMarshalText ")
	}
	AuxMockIncrementRecorderAuxMockUUIDMarshalText()
	return
}

//
// Mock: ParseUUID(arginput string)(reta UUID, retb error)
//

type MockArgsTypeParseUUID struct {
	ApomockCallNumber int
	Arginput          string
}

var LastMockArgsParseUUID MockArgsTypeParseUUID

// AuxMockParseUUID(arginput string)(reta UUID, retb error) - Generated mock function
func AuxMockParseUUID(arginput string) (reta UUID, retb error) {
	LastMockArgsParseUUID = MockArgsTypeParseUUID{
		ApomockCallNumber: AuxMockGetRecorderAuxMockParseUUID(),
		Arginput:          arginput,
	}
	rargs, rerr := apomock.GetNext("gocql.ParseUUID")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ParseUUID")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.ParseUUID")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(UUID)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockParseUUID  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockParseUUID int = 0

var condRecorderAuxMockParseUUID *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockParseUUID(i int) {
	condRecorderAuxMockParseUUID.L.Lock()
	for recorderAuxMockParseUUID < i {
		condRecorderAuxMockParseUUID.Wait()
	}
	condRecorderAuxMockParseUUID.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockParseUUID() {
	condRecorderAuxMockParseUUID.L.Lock()
	recorderAuxMockParseUUID++
	condRecorderAuxMockParseUUID.L.Unlock()
	condRecorderAuxMockParseUUID.Broadcast()
}
func AuxMockGetRecorderAuxMockParseUUID() (ret int) {
	condRecorderAuxMockParseUUID.L.Lock()
	ret = recorderAuxMockParseUUID
	condRecorderAuxMockParseUUID.L.Unlock()
	return
}

// ParseUUID - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func ParseUUID(arginput string) (reta UUID, retb error) {
	FuncAuxMockParseUUID, ok := apomock.GetRegisteredFunc("gocql.ParseUUID")
	if ok {
		reta, retb = FuncAuxMockParseUUID.(func(arginput string) (reta UUID, retb error))(arginput)
	} else {
		panic("FuncAuxMockParseUUID ")
	}
	AuxMockIncrementRecorderAuxMockParseUUID()
	return
}

//
// Mock: UUIDFromTime(argaTime time.Time)(reta UUID)
//

type MockArgsTypeUUIDFromTime struct {
	ApomockCallNumber int
	ArgaTime          time.Time
}

var LastMockArgsUUIDFromTime MockArgsTypeUUIDFromTime

// AuxMockUUIDFromTime(argaTime time.Time)(reta UUID) - Generated mock function
func AuxMockUUIDFromTime(argaTime time.Time) (reta UUID) {
	LastMockArgsUUIDFromTime = MockArgsTypeUUIDFromTime{
		ApomockCallNumber: AuxMockGetRecorderAuxMockUUIDFromTime(),
		ArgaTime:          argaTime,
	}
	rargs, rerr := apomock.GetNext("gocql.UUIDFromTime")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUIDFromTime")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UUIDFromTime")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(UUID)
	}
	return
}

// RecorderAuxMockUUIDFromTime  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUUIDFromTime int = 0

var condRecorderAuxMockUUIDFromTime *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUUIDFromTime(i int) {
	condRecorderAuxMockUUIDFromTime.L.Lock()
	for recorderAuxMockUUIDFromTime < i {
		condRecorderAuxMockUUIDFromTime.Wait()
	}
	condRecorderAuxMockUUIDFromTime.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUUIDFromTime() {
	condRecorderAuxMockUUIDFromTime.L.Lock()
	recorderAuxMockUUIDFromTime++
	condRecorderAuxMockUUIDFromTime.L.Unlock()
	condRecorderAuxMockUUIDFromTime.Broadcast()
}
func AuxMockGetRecorderAuxMockUUIDFromTime() (ret int) {
	condRecorderAuxMockUUIDFromTime.L.Lock()
	ret = recorderAuxMockUUIDFromTime
	condRecorderAuxMockUUIDFromTime.L.Unlock()
	return
}

// UUIDFromTime - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func UUIDFromTime(argaTime time.Time) (reta UUID) {
	FuncAuxMockUUIDFromTime, ok := apomock.GetRegisteredFunc("gocql.UUIDFromTime")
	if ok {
		reta = FuncAuxMockUUIDFromTime.(func(argaTime time.Time) (reta UUID))(argaTime)
	} else {
		panic("FuncAuxMockUUIDFromTime ")
	}
	AuxMockIncrementRecorderAuxMockUUIDFromTime()
	return
}

//
// Mock: (recvu UUID)Variant()(reta int)
//

type MockArgsTypeUUIDVariant struct {
	ApomockCallNumber int
}

var LastMockArgsUUIDVariant MockArgsTypeUUIDVariant

// (recvu UUID)AuxMockVariant()(reta int) - Generated mock function
func (recvu UUID) AuxMockVariant() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.UUID.Variant")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUID.Variant")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UUID.Variant")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockUUIDVariant  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUUIDVariant int = 0

var condRecorderAuxMockUUIDVariant *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUUIDVariant(i int) {
	condRecorderAuxMockUUIDVariant.L.Lock()
	for recorderAuxMockUUIDVariant < i {
		condRecorderAuxMockUUIDVariant.Wait()
	}
	condRecorderAuxMockUUIDVariant.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUUIDVariant() {
	condRecorderAuxMockUUIDVariant.L.Lock()
	recorderAuxMockUUIDVariant++
	condRecorderAuxMockUUIDVariant.L.Unlock()
	condRecorderAuxMockUUIDVariant.Broadcast()
}
func AuxMockGetRecorderAuxMockUUIDVariant() (ret int) {
	condRecorderAuxMockUUIDVariant.L.Lock()
	ret = recorderAuxMockUUIDVariant
	condRecorderAuxMockUUIDVariant.L.Unlock()
	return
}

// (recvu UUID)Variant - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu UUID) Variant() (reta int) {
	FuncAuxMockUUIDVariant, ok := apomock.GetRegisteredFunc("gocql.UUID.Variant")
	if ok {
		reta = FuncAuxMockUUIDVariant.(func(recvu UUID) (reta int))(recvu)
	} else {
		panic("FuncAuxMockUUIDVariant ")
	}
	AuxMockIncrementRecorderAuxMockUUIDVariant()
	return
}

//
// Mock: (recvu UUID)Version()(reta int)
//

type MockArgsTypeUUIDVersion struct {
	ApomockCallNumber int
}

var LastMockArgsUUIDVersion MockArgsTypeUUIDVersion

// (recvu UUID)AuxMockVersion()(reta int) - Generated mock function
func (recvu UUID) AuxMockVersion() (reta int) {
	rargs, rerr := apomock.GetNext("gocql.UUID.Version")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUID.Version")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UUID.Version")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockUUIDVersion  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUUIDVersion int = 0

var condRecorderAuxMockUUIDVersion *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUUIDVersion(i int) {
	condRecorderAuxMockUUIDVersion.L.Lock()
	for recorderAuxMockUUIDVersion < i {
		condRecorderAuxMockUUIDVersion.Wait()
	}
	condRecorderAuxMockUUIDVersion.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUUIDVersion() {
	condRecorderAuxMockUUIDVersion.L.Lock()
	recorderAuxMockUUIDVersion++
	condRecorderAuxMockUUIDVersion.L.Unlock()
	condRecorderAuxMockUUIDVersion.Broadcast()
}
func AuxMockGetRecorderAuxMockUUIDVersion() (ret int) {
	condRecorderAuxMockUUIDVersion.L.Lock()
	ret = recorderAuxMockUUIDVersion
	condRecorderAuxMockUUIDVersion.L.Unlock()
	return
}

// (recvu UUID)Version - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu UUID) Version() (reta int) {
	FuncAuxMockUUIDVersion, ok := apomock.GetRegisteredFunc("gocql.UUID.Version")
	if ok {
		reta = FuncAuxMockUUIDVersion.(func(recvu UUID) (reta int))(recvu)
	} else {
		panic("FuncAuxMockUUIDVersion ")
	}
	AuxMockIncrementRecorderAuxMockUUIDVersion()
	return
}

//
// Mock: (recvu UUID)Node()(reta []byte)
//

type MockArgsTypeUUIDNode struct {
	ApomockCallNumber int
}

var LastMockArgsUUIDNode MockArgsTypeUUIDNode

// (recvu UUID)AuxMockNode()(reta []byte) - Generated mock function
func (recvu UUID) AuxMockNode() (reta []byte) {
	rargs, rerr := apomock.GetNext("gocql.UUID.Node")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUID.Node")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UUID.Node")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockUUIDNode  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUUIDNode int = 0

var condRecorderAuxMockUUIDNode *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUUIDNode(i int) {
	condRecorderAuxMockUUIDNode.L.Lock()
	for recorderAuxMockUUIDNode < i {
		condRecorderAuxMockUUIDNode.Wait()
	}
	condRecorderAuxMockUUIDNode.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUUIDNode() {
	condRecorderAuxMockUUIDNode.L.Lock()
	recorderAuxMockUUIDNode++
	condRecorderAuxMockUUIDNode.L.Unlock()
	condRecorderAuxMockUUIDNode.Broadcast()
}
func AuxMockGetRecorderAuxMockUUIDNode() (ret int) {
	condRecorderAuxMockUUIDNode.L.Lock()
	ret = recorderAuxMockUUIDNode
	condRecorderAuxMockUUIDNode.L.Unlock()
	return
}

// (recvu UUID)Node - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu UUID) Node() (reta []byte) {
	FuncAuxMockUUIDNode, ok := apomock.GetRegisteredFunc("gocql.UUID.Node")
	if ok {
		reta = FuncAuxMockUUIDNode.(func(recvu UUID) (reta []byte))(recvu)
	} else {
		panic("FuncAuxMockUUIDNode ")
	}
	AuxMockIncrementRecorderAuxMockUUIDNode()
	return
}

//
// Mock: (recvu UUID)Timestamp()(reta int64)
//

type MockArgsTypeUUIDTimestamp struct {
	ApomockCallNumber int
}

var LastMockArgsUUIDTimestamp MockArgsTypeUUIDTimestamp

// (recvu UUID)AuxMockTimestamp()(reta int64) - Generated mock function
func (recvu UUID) AuxMockTimestamp() (reta int64) {
	rargs, rerr := apomock.GetNext("gocql.UUID.Timestamp")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUID.Timestamp")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UUID.Timestamp")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int64)
	}
	return
}

// RecorderAuxMockUUIDTimestamp  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUUIDTimestamp int = 0

var condRecorderAuxMockUUIDTimestamp *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUUIDTimestamp(i int) {
	condRecorderAuxMockUUIDTimestamp.L.Lock()
	for recorderAuxMockUUIDTimestamp < i {
		condRecorderAuxMockUUIDTimestamp.Wait()
	}
	condRecorderAuxMockUUIDTimestamp.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUUIDTimestamp() {
	condRecorderAuxMockUUIDTimestamp.L.Lock()
	recorderAuxMockUUIDTimestamp++
	condRecorderAuxMockUUIDTimestamp.L.Unlock()
	condRecorderAuxMockUUIDTimestamp.Broadcast()
}
func AuxMockGetRecorderAuxMockUUIDTimestamp() (ret int) {
	condRecorderAuxMockUUIDTimestamp.L.Lock()
	ret = recorderAuxMockUUIDTimestamp
	condRecorderAuxMockUUIDTimestamp.L.Unlock()
	return
}

// (recvu UUID)Timestamp - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu UUID) Timestamp() (reta int64) {
	FuncAuxMockUUIDTimestamp, ok := apomock.GetRegisteredFunc("gocql.UUID.Timestamp")
	if ok {
		reta = FuncAuxMockUUIDTimestamp.(func(recvu UUID) (reta int64))(recvu)
	} else {
		panic("FuncAuxMockUUIDTimestamp ")
	}
	AuxMockIncrementRecorderAuxMockUUIDTimestamp()
	return
}

//
// Mock: UUIDFromBytes(arginput []byte)(reta UUID, retb error)
//

type MockArgsTypeUUIDFromBytes struct {
	ApomockCallNumber int
	Arginput          []byte
}

var LastMockArgsUUIDFromBytes MockArgsTypeUUIDFromBytes

// AuxMockUUIDFromBytes(arginput []byte)(reta UUID, retb error) - Generated mock function
func AuxMockUUIDFromBytes(arginput []byte) (reta UUID, retb error) {
	LastMockArgsUUIDFromBytes = MockArgsTypeUUIDFromBytes{
		ApomockCallNumber: AuxMockGetRecorderAuxMockUUIDFromBytes(),
		Arginput:          arginput,
	}
	rargs, rerr := apomock.GetNext("gocql.UUIDFromBytes")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUIDFromBytes")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.UUIDFromBytes")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(UUID)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockUUIDFromBytes  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUUIDFromBytes int = 0

var condRecorderAuxMockUUIDFromBytes *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUUIDFromBytes(i int) {
	condRecorderAuxMockUUIDFromBytes.L.Lock()
	for recorderAuxMockUUIDFromBytes < i {
		condRecorderAuxMockUUIDFromBytes.Wait()
	}
	condRecorderAuxMockUUIDFromBytes.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUUIDFromBytes() {
	condRecorderAuxMockUUIDFromBytes.L.Lock()
	recorderAuxMockUUIDFromBytes++
	condRecorderAuxMockUUIDFromBytes.L.Unlock()
	condRecorderAuxMockUUIDFromBytes.Broadcast()
}
func AuxMockGetRecorderAuxMockUUIDFromBytes() (ret int) {
	condRecorderAuxMockUUIDFromBytes.L.Lock()
	ret = recorderAuxMockUUIDFromBytes
	condRecorderAuxMockUUIDFromBytes.L.Unlock()
	return
}

// UUIDFromBytes - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func UUIDFromBytes(arginput []byte) (reta UUID, retb error) {
	FuncAuxMockUUIDFromBytes, ok := apomock.GetRegisteredFunc("gocql.UUIDFromBytes")
	if ok {
		reta, retb = FuncAuxMockUUIDFromBytes.(func(arginput []byte) (reta UUID, retb error))(arginput)
	} else {
		panic("FuncAuxMockUUIDFromBytes ")
	}
	AuxMockIncrementRecorderAuxMockUUIDFromBytes()
	return
}

//
// Mock: RandomUUID()(reta UUID, retb error)
//

type MockArgsTypeRandomUUID struct {
	ApomockCallNumber int
}

var LastMockArgsRandomUUID MockArgsTypeRandomUUID

// AuxMockRandomUUID()(reta UUID, retb error) - Generated mock function
func AuxMockRandomUUID() (reta UUID, retb error) {
	rargs, rerr := apomock.GetNext("gocql.RandomUUID")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.RandomUUID")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.RandomUUID")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(UUID)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockRandomUUID  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockRandomUUID int = 0

var condRecorderAuxMockRandomUUID *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockRandomUUID(i int) {
	condRecorderAuxMockRandomUUID.L.Lock()
	for recorderAuxMockRandomUUID < i {
		condRecorderAuxMockRandomUUID.Wait()
	}
	condRecorderAuxMockRandomUUID.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockRandomUUID() {
	condRecorderAuxMockRandomUUID.L.Lock()
	recorderAuxMockRandomUUID++
	condRecorderAuxMockRandomUUID.L.Unlock()
	condRecorderAuxMockRandomUUID.Broadcast()
}
func AuxMockGetRecorderAuxMockRandomUUID() (ret int) {
	condRecorderAuxMockRandomUUID.L.Lock()
	ret = recorderAuxMockRandomUUID
	condRecorderAuxMockRandomUUID.L.Unlock()
	return
}

// RandomUUID - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func RandomUUID() (reta UUID, retb error) {
	FuncAuxMockRandomUUID, ok := apomock.GetRegisteredFunc("gocql.RandomUUID")
	if ok {
		reta, retb = FuncAuxMockRandomUUID.(func() (reta UUID, retb error))()
	} else {
		panic("FuncAuxMockRandomUUID ")
	}
	AuxMockIncrementRecorderAuxMockRandomUUID()
	return
}

//
// Mock: TimeUUID()(reta UUID)
//

type MockArgsTypeTimeUUID struct {
	ApomockCallNumber int
}

var LastMockArgsTimeUUID MockArgsTypeTimeUUID

// AuxMockTimeUUID()(reta UUID) - Generated mock function
func AuxMockTimeUUID() (reta UUID) {
	rargs, rerr := apomock.GetNext("gocql.TimeUUID")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.TimeUUID")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.TimeUUID")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(UUID)
	}
	return
}

// RecorderAuxMockTimeUUID  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockTimeUUID int = 0

var condRecorderAuxMockTimeUUID *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockTimeUUID(i int) {
	condRecorderAuxMockTimeUUID.L.Lock()
	for recorderAuxMockTimeUUID < i {
		condRecorderAuxMockTimeUUID.Wait()
	}
	condRecorderAuxMockTimeUUID.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockTimeUUID() {
	condRecorderAuxMockTimeUUID.L.Lock()
	recorderAuxMockTimeUUID++
	condRecorderAuxMockTimeUUID.L.Unlock()
	condRecorderAuxMockTimeUUID.Broadcast()
}
func AuxMockGetRecorderAuxMockTimeUUID() (ret int) {
	condRecorderAuxMockTimeUUID.L.Lock()
	ret = recorderAuxMockTimeUUID
	condRecorderAuxMockTimeUUID.L.Unlock()
	return
}

// TimeUUID - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func TimeUUID() (reta UUID) {
	FuncAuxMockTimeUUID, ok := apomock.GetRegisteredFunc("gocql.TimeUUID")
	if ok {
		reta = FuncAuxMockTimeUUID.(func() (reta UUID))()
	} else {
		panic("FuncAuxMockTimeUUID ")
	}
	AuxMockIncrementRecorderAuxMockTimeUUID()
	return
}

//
// Mock: (recvu UUID)String()(reta string)
//

type MockArgsTypeUUIDString struct {
	ApomockCallNumber int
}

var LastMockArgsUUIDString MockArgsTypeUUIDString

// (recvu UUID)AuxMockString()(reta string) - Generated mock function
func (recvu UUID) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.UUID.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUID.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UUID.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockUUIDString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUUIDString int = 0

var condRecorderAuxMockUUIDString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUUIDString(i int) {
	condRecorderAuxMockUUIDString.L.Lock()
	for recorderAuxMockUUIDString < i {
		condRecorderAuxMockUUIDString.Wait()
	}
	condRecorderAuxMockUUIDString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUUIDString() {
	condRecorderAuxMockUUIDString.L.Lock()
	recorderAuxMockUUIDString++
	condRecorderAuxMockUUIDString.L.Unlock()
	condRecorderAuxMockUUIDString.Broadcast()
}
func AuxMockGetRecorderAuxMockUUIDString() (ret int) {
	condRecorderAuxMockUUIDString.L.Lock()
	ret = recorderAuxMockUUIDString
	condRecorderAuxMockUUIDString.L.Unlock()
	return
}

// (recvu UUID)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu UUID) String() (reta string) {
	FuncAuxMockUUIDString, ok := apomock.GetRegisteredFunc("gocql.UUID.String")
	if ok {
		reta = FuncAuxMockUUIDString.(func(recvu UUID) (reta string))(recvu)
	} else {
		panic("FuncAuxMockUUIDString ")
	}
	AuxMockIncrementRecorderAuxMockUUIDString()
	return
}

//
// Mock: (recvu UUID)Bytes()(reta []byte)
//

type MockArgsTypeUUIDBytes struct {
	ApomockCallNumber int
}

var LastMockArgsUUIDBytes MockArgsTypeUUIDBytes

// (recvu UUID)AuxMockBytes()(reta []byte) - Generated mock function
func (recvu UUID) AuxMockBytes() (reta []byte) {
	rargs, rerr := apomock.GetNext("gocql.UUID.Bytes")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.UUID.Bytes")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.UUID.Bytes")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockUUIDBytes  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockUUIDBytes int = 0

var condRecorderAuxMockUUIDBytes *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockUUIDBytes(i int) {
	condRecorderAuxMockUUIDBytes.L.Lock()
	for recorderAuxMockUUIDBytes < i {
		condRecorderAuxMockUUIDBytes.Wait()
	}
	condRecorderAuxMockUUIDBytes.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockUUIDBytes() {
	condRecorderAuxMockUUIDBytes.L.Lock()
	recorderAuxMockUUIDBytes++
	condRecorderAuxMockUUIDBytes.L.Unlock()
	condRecorderAuxMockUUIDBytes.Broadcast()
}
func AuxMockGetRecorderAuxMockUUIDBytes() (ret int) {
	condRecorderAuxMockUUIDBytes.L.Lock()
	ret = recorderAuxMockUUIDBytes
	condRecorderAuxMockUUIDBytes.L.Unlock()
	return
}

// (recvu UUID)Bytes - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvu UUID) Bytes() (reta []byte) {
	FuncAuxMockUUIDBytes, ok := apomock.GetRegisteredFunc("gocql.UUID.Bytes")
	if ok {
		reta = FuncAuxMockUUIDBytes.(func(recvu UUID) (reta []byte))(recvu)
	} else {
		panic("FuncAuxMockUUIDBytes ")
	}
	AuxMockIncrementRecorderAuxMockUUIDBytes()
	return
}
