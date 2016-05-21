// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"
import "net"

import "errors"
import "io"

func init() {
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructPreparedMetadata, apomockNewStructPreparedMetadata)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructTopologyChangeEventFrame, apomockNewStructTopologyChangeEventFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructWriteRegisterFrame, apomockNewStructWriteRegisterFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructWriteQueryFrame, apomockNewStructWriteQueryFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructResultMetadata, apomockNewStructResultMetadata)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructResultRowsFrame, apomockNewStructResultRowsFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructStatusChangeEventFrame, apomockNewStructStatusChangeEventFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructWriteAuthResponseFrame, apomockNewStructWriteAuthResponseFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructQueryValues, apomockNewStructQueryValues)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructWriteExecuteFrame, apomockNewStructWriteExecuteFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructSupportedFrame, apomockNewStructSupportedFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructQueryParams, apomockNewStructQueryParams)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructAuthenticateFrame, apomockNewStructAuthenticateFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructResultKeyspaceFrame, apomockNewStructResultKeyspaceFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructSchemaChangeFunction, apomockNewStructSchemaChangeFunction)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructBatchStatment, apomockNewStructBatchStatment)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructWriteStartupFrame, apomockNewStructWriteStartupFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructAuthChallengeFrame, apomockNewStructAuthChallengeFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructReadyFrame, apomockNewStructReadyFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructResultVoidFrame, apomockNewStructResultVoidFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructSchemaChangeTable, apomockNewStructSchemaChangeTable)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructFramer, apomockNewStructFramer)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructFrameHeader, apomockNewStructFrameHeader)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructWritePrepareFrame, apomockNewStructWritePrepareFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructResultPreparedFrame, apomockNewStructResultPreparedFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructSchemaChangeKeyspace, apomockNewStructSchemaChangeKeyspace)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructAuthSuccessFrame, apomockNewStructAuthSuccessFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructWriteBatchFrame, apomockNewStructWriteBatchFrame)
	apomock.RegisterInternalType("github.com/aporeto-inc/gocql", ApomockStructWriteOptionsFrame, apomockNewStructWriteOptionsFrame)

	apomock.RegisterFunc("gocql", "gocql.writeStartupFrame.writeFrame", (*writeStartupFrame).AuxMockwriteFrame)
	apomock.RegisterFunc("gocql", "gocql.writeBatchFrame.writeFrame", (*writeBatchFrame).AuxMockwriteFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.writeConsistency", (*framer).AuxMockwriteConsistency)
	apomock.RegisterFunc("gocql", "gocql.frameOp.String", (frameOp).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.writeStartupFrame", (*framer).AuxMockwriteStartupFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.parseResultPrepared", (*framer).AuxMockparseResultPrepared)
	apomock.RegisterFunc("gocql", "gocql.authenticateFrame.String", (*authenticateFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.writeExecuteFrame", (*framer).AuxMockwriteExecuteFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.writeRegisterFrame", (*framer).AuxMockwriteRegisterFrame)
	apomock.RegisterFunc("gocql", "gocql.authChallengeFrame.String", (*authChallengeFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.readBytesInternal", (*framer).AuxMockreadBytesInternal)
	apomock.RegisterFunc("gocql", "gocql.framer.readStringMap", (*framer).AuxMockreadStringMap)
	apomock.RegisterFunc("gocql", "gocql.writeInt", AuxMockwriteInt)
	apomock.RegisterFunc("gocql", "gocql.resultMetadata.String", (resultMetadata).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.parseResultSchemaChange", (*framer).AuxMockparseResultSchemaChange)
	apomock.RegisterFunc("gocql", "gocql.framer.readUUID", (*framer).AuxMockreadUUID)
	apomock.RegisterFunc("gocql", "gocql.protoVersion.request", (protoVersion).AuxMockrequest)
	apomock.RegisterFunc("gocql", "gocql.framer.parsePreparedMetadata", (*framer).AuxMockparsePreparedMetadata)
	apomock.RegisterFunc("gocql", "gocql.framer.writeAuthResponseFrame", (*framer).AuxMockwriteAuthResponseFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.writeBatchFrame", (*framer).AuxMockwriteBatchFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.readBytes", (*framer).AuxMockreadBytes)
	apomock.RegisterFunc("gocql", "gocql.framer.readInt", (*framer).AuxMockreadInt)
	apomock.RegisterFunc("gocql", "gocql.framer.readStringList", (*framer).AuxMockreadStringList)
	apomock.RegisterFunc("gocql", "gocql.framer.parseErrorFrame", (*framer).AuxMockparseErrorFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.readTypeInfo", (*framer).AuxMockreadTypeInfo)
	apomock.RegisterFunc("gocql", "gocql.framer.readCol", (*framer).AuxMockreadCol)
	apomock.RegisterFunc("gocql", "gocql.framer.writeQueryParams", (*framer).AuxMockwriteQueryParams)
	apomock.RegisterFunc("gocql", "gocql.framer.readConsistency", (*framer).AuxMockreadConsistency)
	apomock.RegisterFunc("gocql", "gocql.framer.readLong", (*framer).AuxMockreadLong)
	apomock.RegisterFunc("gocql", "gocql.framer.readLongString", (*framer).AuxMockreadLongString)
	apomock.RegisterFunc("gocql", "gocql.protoVersion.version", (protoVersion).AuxMockversion)
	apomock.RegisterFunc("gocql", "gocql.framer.parseResultRows", (*framer).AuxMockparseResultRows)
	apomock.RegisterFunc("gocql", "gocql.framer.parseAuthChallengeFrame", (*framer).AuxMockparseAuthChallengeFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.writeQueryFrame", (*framer).AuxMockwriteQueryFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.writeInt", (*framer).AuxMockwriteInt)
	apomock.RegisterFunc("gocql", "gocql.framer.writeShortBytes", (*framer).AuxMockwriteShortBytes)
	apomock.RegisterFunc("gocql", "gocql.frameHeader.Header", (frameHeader).AuxMockHeader)
	apomock.RegisterFunc("gocql", "gocql.writeQueryFrame.writeFrame", (*writeQueryFrame).AuxMockwriteFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.writeByte", (*framer).AuxMockwriteByte)
	apomock.RegisterFunc("gocql", "gocql.protoVersion.String", (protoVersion).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.writeHeader", (*framer).AuxMockwriteHeader)
	apomock.RegisterFunc("gocql", "gocql.framer.setLength", (*framer).AuxMocksetLength)
	apomock.RegisterFunc("gocql", "gocql.framer.parseAuthenticateFrame", (*framer).AuxMockparseAuthenticateFrame)
	apomock.RegisterFunc("gocql", "gocql.appendLong", AuxMockappendLong)
	apomock.RegisterFunc("gocql", "gocql.readHeader", AuxMockreadHeader)
	apomock.RegisterFunc("gocql", "gocql.framer.writePrepareFrame", (*framer).AuxMockwritePrepareFrame)
	apomock.RegisterFunc("gocql", "gocql.writeExecuteFrame.String", (*writeExecuteFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.appendBytes", AuxMockappendBytes)
	apomock.RegisterFunc("gocql", "gocql.readInt", AuxMockreadInt)
	apomock.RegisterFunc("gocql", "gocql.frameHeader.String", (frameHeader).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.readFrame", (*framer).AuxMockreadFrame)
	apomock.RegisterFunc("gocql", "gocql.resultVoidFrame.String", (*resultVoidFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.resultKeyspaceFrame.String", (*resultKeyspaceFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.topologyChangeEventFrame.String", (topologyChangeEventFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.writeStringMap", (*framer).AuxMockwriteStringMap)
	apomock.RegisterFunc("gocql", "gocql.framer.parseResultMetadata", (*framer).AuxMockparseResultMetadata)
	apomock.RegisterFunc("gocql", "gocql.schemaChangeTable.String", (schemaChangeTable).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.parseEventFrame", (*framer).AuxMockparseEventFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.writeOptionsFrame", (*framer).AuxMockwriteOptionsFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.writeInet", (*framer).AuxMockwriteInet)
	apomock.RegisterFunc("gocql", "gocql.framer.readTrace", (*framer).AuxMockreadTrace)
	apomock.RegisterFunc("gocql", "gocql.resultRowsFrame.String", (*resultRowsFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.writeAuthResponseFrame.writeFrame", (*writeAuthResponseFrame).AuxMockwriteFrame)
	apomock.RegisterFunc("gocql", "gocql.protoVersion.response", (protoVersion).AuxMockresponse)
	apomock.RegisterFunc("gocql", "gocql.ParseConsistency", AuxMockParseConsistency)
	apomock.RegisterFunc("gocql", "gocql.writePrepareFrame.writeFrame", (*writePrepareFrame).AuxMockwriteFrame)
	apomock.RegisterFunc("gocql", "gocql.schemaChangeKeyspace.String", (schemaChangeKeyspace).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.statusChangeEventFrame.String", (statusChangeEventFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.readString", (*framer).AuxMockreadString)
	apomock.RegisterFunc("gocql", "gocql.framer.readStringMultiMap", (*framer).AuxMockreadStringMultiMap)
	apomock.RegisterFunc("gocql", "gocql.writeQueryFrame.String", (*writeQueryFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.writeAuthResponseFrame.String", (*writeAuthResponseFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.readShortBytes", (*framer).AuxMockreadShortBytes)
	apomock.RegisterFunc("gocql", "gocql.SerialConsistency.String", (SerialConsistency).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.readShort", AuxMockreadShort)
	apomock.RegisterFunc("gocql", "gocql.framer.readByte", (*framer).AuxMockreadByte)
	apomock.RegisterFunc("gocql", "gocql.framer.parseResultFrame", (*framer).AuxMockparseResultFrame)
	apomock.RegisterFunc("gocql", "gocql.writeRegisterFrame.writeFrame", (*writeRegisterFrame).AuxMockwriteFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.writeShort", (*framer).AuxMockwriteShort)
	apomock.RegisterFunc("gocql", "gocql.framer.parseReadyFrame", (*framer).AuxMockparseReadyFrame)
	apomock.RegisterFunc("gocql", "gocql.writeShort", AuxMockwriteShort)
	apomock.RegisterFunc("gocql", "gocql.framer.trace", (*framer).AuxMocktrace)
	apomock.RegisterFunc("gocql", "gocql.authSuccessFrame.String", (*authSuccessFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.newFramer", AuxMocknewFramer)
	apomock.RegisterFunc("gocql", "gocql.writeExecuteFrame.writeFrame", (*writeExecuteFrame).AuxMockwriteFrame)
	apomock.RegisterFunc("gocql", "gocql.appendShort", AuxMockappendShort)
	apomock.RegisterFunc("gocql", "gocql.framer.writeLong", (*framer).AuxMockwriteLong)
	apomock.RegisterFunc("gocql", "gocql.Consistency.String", (Consistency).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.parseSupportedFrame", (*framer).AuxMockparseSupportedFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.parseFrame", (*framer).AuxMockparseFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.readInet", (*framer).AuxMockreadInet)
	apomock.RegisterFunc("gocql", "gocql.framer.writeStringList", (*framer).AuxMockwriteStringList)
	apomock.RegisterFunc("gocql", "gocql.framer.readBytesMap", (*framer).AuxMockreadBytesMap)
	apomock.RegisterFunc("gocql", "gocql.queryParams.String", (queryParams).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.writeString", (*framer).AuxMockwriteString)
	apomock.RegisterFunc("gocql", "gocql.framer.writeLongString", (*framer).AuxMockwriteLongString)
	apomock.RegisterFunc("gocql", "gocql.writeStartupFrame.String", (writeStartupFrame).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.preparedMetadata.String", (preparedMetadata).AuxMockString)
	apomock.RegisterFunc("gocql", "gocql.framer.parseAuthSuccessFrame", (*framer).AuxMockparseAuthSuccessFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.readShort", (*framer).AuxMockreadShort)
	apomock.RegisterFunc("gocql", "gocql.appendInt", AuxMockappendInt)
	apomock.RegisterFunc("gocql", "gocql.framer.writeUUID", (*framer).AuxMockwriteUUID)
	apomock.RegisterFunc("gocql", "gocql.framer.finishWrite", (*framer).AuxMockfinishWrite)
	apomock.RegisterFunc("gocql", "gocql.framer.writeBytes", (*framer).AuxMockwriteBytes)
	apomock.RegisterFunc("gocql", "gocql.frameWriterFunc.writeFrame", (frameWriterFunc).AuxMockwriteFrame)
	apomock.RegisterFunc("gocql", "gocql.framer.parseResultSetKeyspace", (*framer).AuxMockparseResultSetKeyspace)
	apomock.RegisterFunc("gocql", "gocql.writeOptionsFrame.writeFrame", (*writeOptionsFrame).AuxMockwriteFrame)
}

const (
	protoDirectionMask = 0x80
	protoVersionMask   = 0x7F
	protoVersion1      = 0x01
	protoVersion2      = 0x02
	protoVersion3      = 0x03
	protoVersion4      = 0x04
	maxFrameSize       = 256 * 1024 * 1024
)

const (
	opError         frameOp = 0x00
	opStartup       frameOp = 0x01
	opReady         frameOp = 0x02
	opAuthenticate  frameOp = 0x03
	opOptions       frameOp = 0x05
	opSupported     frameOp = 0x06
	opQuery         frameOp = 0x07
	opResult        frameOp = 0x08
	opPrepare       frameOp = 0x09
	opExecute       frameOp = 0x0A
	opRegister      frameOp = 0x0B
	opEvent         frameOp = 0x0C
	opBatch         frameOp = 0x0D
	opAuthChallenge frameOp = 0x0E
	opAuthResponse  frameOp = 0x0F
	opAuthSuccess   frameOp = 0x10
)

const (
	resultKindVoid                 = 1
	resultKindRows                 = 2
	resultKindKeyspace             = 3
	resultKindPrepared             = 4
	resultKindSchemaChanged        = 5
	flagGlobalTableSpec       int  = 0x01
	flagHasMorePages          int  = 0x02
	flagNoMetaData            int  = 0x04
	flagValues                byte = 0x01
	flagSkipMetaData          byte = 0x02
	flagPageSize              byte = 0x04
	flagWithPagingState       byte = 0x08
	flagWithSerialConsistency byte = 0x10
	flagDefaultTimestamp      byte = 0x20
	flagWithNameValues        byte = 0x40
	flagCompress              byte = 0x01
	flagTracing               byte = 0x02
	flagCustomPayload         byte = 0x04
	flagWarning               byte = 0x08
)

const (
	Any         Consistency = 0x00
	One         Consistency = 0x01
	Two         Consistency = 0x02
	Three       Consistency = 0x03
	Quorum      Consistency = 0x04
	All         Consistency = 0x05
	LocalQuorum Consistency = 0x06
	EachQuorum  Consistency = 0x07
	LocalOne    Consistency = 0x0A
)

const (
	Serial      SerialConsistency = 0x08
	LocalSerial SerialConsistency = 0x09
)

const (
	apacheCassandraTypePrefix = "org.apache.cassandra.db.marshal."
)

const maxFrameHeaderSize = 9

const defaultBufSize = 128

const (
	ApomockStructPreparedMetadata         = 9
	ApomockStructTopologyChangeEventFrame = 10
	ApomockStructWriteRegisterFrame       = 11
	ApomockStructWriteQueryFrame          = 12
	ApomockStructResultMetadata           = 13
	ApomockStructResultRowsFrame          = 14
	ApomockStructStatusChangeEventFrame   = 15
	ApomockStructWriteAuthResponseFrame   = 16
	ApomockStructQueryValues              = 17
	ApomockStructWriteExecuteFrame        = 18
	ApomockStructSupportedFrame           = 19
	ApomockStructQueryParams              = 20
	ApomockStructAuthenticateFrame        = 21
	ApomockStructResultKeyspaceFrame      = 22
	ApomockStructSchemaChangeFunction     = 23
	ApomockStructBatchStatment            = 24
	ApomockStructWriteStartupFrame        = 25
	ApomockStructAuthChallengeFrame       = 26
	ApomockStructReadyFrame               = 27
	ApomockStructResultVoidFrame          = 28
	ApomockStructSchemaChangeTable        = 29
	ApomockStructFramer                   = 30
	ApomockStructFrameHeader              = 31
	ApomockStructWritePrepareFrame        = 32
	ApomockStructResultPreparedFrame      = 33
	ApomockStructSchemaChangeKeyspace     = 34
	ApomockStructAuthSuccessFrame         = 35
	ApomockStructWriteBatchFrame          = 36
	ApomockStructWriteOptionsFrame        = 37
)

var (
	ErrFrameTooBig = errors.New("frame length is bigger than the maximum allowed")
)

var framerPool = sync.Pool{New: func() interface{} {
	return &framer{wbuf: make([]byte, defaultBufSize), readBuffer: make([]byte, defaultBufSize)}
}}

//
// Internal Types: in this package and their exportable versions
//
type preparedMetadata struct {
	resultMetadata
	pkeyColumns []int
}
type topologyChangeEventFrame struct {
	frameHeader
	change string
	host   net.IP
	port   int
}
type frameWriter interface {
	writeFrame(framer *framer, streamID int) error
}
type writeRegisterFrame struct{ events []string }
type writeQueryFrame struct {
	statement string
	params    queryParams
}
type resultMetadata struct {
	flags          int
	pagingState    []byte
	columns        []ColumnInfo
	colCount       int
	actualColCount int
}
type resultRowsFrame struct {
	frameHeader
	meta    resultMetadata
	numRows int
}
type statusChangeEventFrame struct {
	frameHeader
	change string
	host   net.IP
	port   int
}
type writeAuthResponseFrame struct{ data []byte }
type queryValues struct {
	value []byte
	name  string
}
type frameWriterFunc func(framer *framer, streamID int) error
type writeExecuteFrame struct {
	preparedID []byte
	params     queryParams
}
type supportedFrame struct {
	frameHeader
	supported map[string][]string
}
type queryParams struct {
	consistency           Consistency
	skipMeta              bool
	values                []queryValues
	pageSize              int
	pagingState           []byte
	serialConsistency     SerialConsistency
	defaultTimestamp      bool
	defaultTimestampValue int64
}
type authenticateFrame struct {
	frameHeader
	class string
}
type resultKeyspaceFrame struct {
	frameHeader
	keyspace string
}
type schemaChangeFunction struct {
	frameHeader
	change   string
	keyspace string
	name     string
	args     []string
}
type batchStatment struct {
	preparedID []byte
	statement  string
	values     []queryValues
}
type frame interface {
	Header() frameHeader
}
type writeStartupFrame struct{ opts map[string]string }
type authChallengeFrame struct {
	frameHeader
	data []byte
}
type readyFrame struct{ frameHeader }
type resultVoidFrame struct{ frameHeader }
type schemaChangeTable struct {
	frameHeader
	change   string
	keyspace string
	object   string
}
type framer struct {
	r          io.Reader
	w          io.Writer
	proto      byte
	flags      byte
	compres    Compressor
	headSize   int
	header     *frameHeader
	traceID    []byte
	readBuffer []byte
	rbuf       []byte
	wbuf       []byte
}
type frameOp byte
type frameHeader struct {
	version       protoVersion
	flags         byte
	stream        int
	op            frameOp
	length        int
	customPayload map[string][]byte
}
type writePrepareFrame struct{ statement string }
type resultPreparedFrame struct {
	frameHeader
	preparedID []byte
	reqMeta    preparedMetadata
	respMeta   resultMetadata
}
type schemaChangeKeyspace struct {
	frameHeader
	change   string
	keyspace string
}
type authSuccessFrame struct {
	frameHeader
	data []byte
}
type writeBatchFrame struct {
	typ               BatchType
	statements        []batchStatment
	consistency       Consistency
	serialConsistency SerialConsistency
	defaultTimestamp  bool
}
type protoVersion byte
type writeOptionsFrame struct{}

//
// External Types: in this package
//
type SerialConsistency uint16

type Consistency uint16

func apomockNewStructPreparedMetadata() interface{}         { return &preparedMetadata{} }
func apomockNewStructTopologyChangeEventFrame() interface{} { return &topologyChangeEventFrame{} }
func apomockNewStructWriteRegisterFrame() interface{}       { return &writeRegisterFrame{} }
func apomockNewStructWriteQueryFrame() interface{}          { return &writeQueryFrame{} }
func apomockNewStructResultMetadata() interface{}           { return &resultMetadata{} }
func apomockNewStructResultRowsFrame() interface{}          { return &resultRowsFrame{} }
func apomockNewStructStatusChangeEventFrame() interface{}   { return &statusChangeEventFrame{} }
func apomockNewStructWriteAuthResponseFrame() interface{}   { return &writeAuthResponseFrame{} }
func apomockNewStructQueryValues() interface{}              { return &queryValues{} }
func apomockNewStructWriteExecuteFrame() interface{}        { return &writeExecuteFrame{} }
func apomockNewStructSupportedFrame() interface{}           { return &supportedFrame{} }
func apomockNewStructQueryParams() interface{}              { return &queryParams{} }
func apomockNewStructAuthenticateFrame() interface{}        { return &authenticateFrame{} }
func apomockNewStructResultKeyspaceFrame() interface{}      { return &resultKeyspaceFrame{} }
func apomockNewStructSchemaChangeFunction() interface{}     { return &schemaChangeFunction{} }
func apomockNewStructBatchStatment() interface{}            { return &batchStatment{} }
func apomockNewStructWriteStartupFrame() interface{}        { return &writeStartupFrame{} }
func apomockNewStructAuthChallengeFrame() interface{}       { return &authChallengeFrame{} }
func apomockNewStructReadyFrame() interface{}               { return &readyFrame{} }
func apomockNewStructResultVoidFrame() interface{}          { return &resultVoidFrame{} }
func apomockNewStructSchemaChangeTable() interface{}        { return &schemaChangeTable{} }
func apomockNewStructFramer() interface{}                   { return &framer{} }
func apomockNewStructFrameHeader() interface{}              { return &frameHeader{} }
func apomockNewStructWritePrepareFrame() interface{}        { return &writePrepareFrame{} }
func apomockNewStructResultPreparedFrame() interface{}      { return &resultPreparedFrame{} }
func apomockNewStructSchemaChangeKeyspace() interface{}     { return &schemaChangeKeyspace{} }
func apomockNewStructAuthSuccessFrame() interface{}         { return &authSuccessFrame{} }
func apomockNewStructWriteBatchFrame() interface{}          { return &writeBatchFrame{} }
func apomockNewStructWriteOptionsFrame() interface{}        { return &writeOptionsFrame{} }

//
// Mock: (recvw *writeStartupFrame)writeFrame(argframer *framer, argstreamID int)(reta error)
//

type MockArgsTypewriteStartupFramewriteFrame struct {
	ApomockCallNumber int
	Argframer         *framer
	ArgstreamID       int
}

var LastMockArgswriteStartupFramewriteFrame MockArgsTypewriteStartupFramewriteFrame

// (recvw *writeStartupFrame)AuxMockwriteFrame(argframer *framer, argstreamID int)(reta error) - Generated mock function
func (recvw *writeStartupFrame) AuxMockwriteFrame(argframer *framer, argstreamID int) (reta error) {
	LastMockArgswriteStartupFramewriteFrame = MockArgsTypewriteStartupFramewriteFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrwriteStartupFramewriteFrame(),
		Argframer:         argframer,
		ArgstreamID:       argstreamID,
	}
	rargs, rerr := apomock.GetNext("gocql.writeStartupFrame.writeFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeStartupFrame.writeFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeStartupFrame.writeFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrwriteStartupFramewriteFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrwriteStartupFramewriteFrame int = 0

var condRecorderAuxMockPtrwriteStartupFramewriteFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrwriteStartupFramewriteFrame(i int) {
	condRecorderAuxMockPtrwriteStartupFramewriteFrame.L.Lock()
	for recorderAuxMockPtrwriteStartupFramewriteFrame < i {
		condRecorderAuxMockPtrwriteStartupFramewriteFrame.Wait()
	}
	condRecorderAuxMockPtrwriteStartupFramewriteFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrwriteStartupFramewriteFrame() {
	condRecorderAuxMockPtrwriteStartupFramewriteFrame.L.Lock()
	recorderAuxMockPtrwriteStartupFramewriteFrame++
	condRecorderAuxMockPtrwriteStartupFramewriteFrame.L.Unlock()
	condRecorderAuxMockPtrwriteStartupFramewriteFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrwriteStartupFramewriteFrame() (ret int) {
	condRecorderAuxMockPtrwriteStartupFramewriteFrame.L.Lock()
	ret = recorderAuxMockPtrwriteStartupFramewriteFrame
	condRecorderAuxMockPtrwriteStartupFramewriteFrame.L.Unlock()
	return
}

// (recvw *writeStartupFrame)writeFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvw *writeStartupFrame) writeFrame(argframer *framer, argstreamID int) (reta error) {
	FuncAuxMockPtrwriteStartupFramewriteFrame, ok := apomock.GetRegisteredFunc("gocql.writeStartupFrame.writeFrame")
	if ok {
		reta = FuncAuxMockPtrwriteStartupFramewriteFrame.(func(recvw *writeStartupFrame, argframer *framer, argstreamID int) (reta error))(recvw, argframer, argstreamID)
	} else {
		panic("FuncAuxMockPtrwriteStartupFramewriteFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrwriteStartupFramewriteFrame()
	return
}

//
// Mock: (recvw *writeBatchFrame)writeFrame(argframer *framer, argstreamID int)(reta error)
//

type MockArgsTypewriteBatchFramewriteFrame struct {
	ApomockCallNumber int
	Argframer         *framer
	ArgstreamID       int
}

var LastMockArgswriteBatchFramewriteFrame MockArgsTypewriteBatchFramewriteFrame

// (recvw *writeBatchFrame)AuxMockwriteFrame(argframer *framer, argstreamID int)(reta error) - Generated mock function
func (recvw *writeBatchFrame) AuxMockwriteFrame(argframer *framer, argstreamID int) (reta error) {
	LastMockArgswriteBatchFramewriteFrame = MockArgsTypewriteBatchFramewriteFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrwriteBatchFramewriteFrame(),
		Argframer:         argframer,
		ArgstreamID:       argstreamID,
	}
	rargs, rerr := apomock.GetNext("gocql.writeBatchFrame.writeFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeBatchFrame.writeFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeBatchFrame.writeFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrwriteBatchFramewriteFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrwriteBatchFramewriteFrame int = 0

var condRecorderAuxMockPtrwriteBatchFramewriteFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrwriteBatchFramewriteFrame(i int) {
	condRecorderAuxMockPtrwriteBatchFramewriteFrame.L.Lock()
	for recorderAuxMockPtrwriteBatchFramewriteFrame < i {
		condRecorderAuxMockPtrwriteBatchFramewriteFrame.Wait()
	}
	condRecorderAuxMockPtrwriteBatchFramewriteFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrwriteBatchFramewriteFrame() {
	condRecorderAuxMockPtrwriteBatchFramewriteFrame.L.Lock()
	recorderAuxMockPtrwriteBatchFramewriteFrame++
	condRecorderAuxMockPtrwriteBatchFramewriteFrame.L.Unlock()
	condRecorderAuxMockPtrwriteBatchFramewriteFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrwriteBatchFramewriteFrame() (ret int) {
	condRecorderAuxMockPtrwriteBatchFramewriteFrame.L.Lock()
	ret = recorderAuxMockPtrwriteBatchFramewriteFrame
	condRecorderAuxMockPtrwriteBatchFramewriteFrame.L.Unlock()
	return
}

// (recvw *writeBatchFrame)writeFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvw *writeBatchFrame) writeFrame(argframer *framer, argstreamID int) (reta error) {
	FuncAuxMockPtrwriteBatchFramewriteFrame, ok := apomock.GetRegisteredFunc("gocql.writeBatchFrame.writeFrame")
	if ok {
		reta = FuncAuxMockPtrwriteBatchFramewriteFrame.(func(recvw *writeBatchFrame, argframer *framer, argstreamID int) (reta error))(recvw, argframer, argstreamID)
	} else {
		panic("FuncAuxMockPtrwriteBatchFramewriteFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrwriteBatchFramewriteFrame()
	return
}

//
// Mock: (recvf *framer)writeConsistency(argcons Consistency)()
//

type MockArgsTypeframerwriteConsistency struct {
	ApomockCallNumber int
	Argcons           Consistency
}

var LastMockArgsframerwriteConsistency MockArgsTypeframerwriteConsistency

// (recvf *framer)AuxMockwriteConsistency(argcons Consistency)() - Generated mock function
func (recvf *framer) AuxMockwriteConsistency(argcons Consistency) {
	LastMockArgsframerwriteConsistency = MockArgsTypeframerwriteConsistency{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteConsistency(),
		Argcons:           argcons,
	}
	return
}

// RecorderAuxMockPtrframerwriteConsistency  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteConsistency int = 0

var condRecorderAuxMockPtrframerwriteConsistency *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteConsistency(i int) {
	condRecorderAuxMockPtrframerwriteConsistency.L.Lock()
	for recorderAuxMockPtrframerwriteConsistency < i {
		condRecorderAuxMockPtrframerwriteConsistency.Wait()
	}
	condRecorderAuxMockPtrframerwriteConsistency.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteConsistency() {
	condRecorderAuxMockPtrframerwriteConsistency.L.Lock()
	recorderAuxMockPtrframerwriteConsistency++
	condRecorderAuxMockPtrframerwriteConsistency.L.Unlock()
	condRecorderAuxMockPtrframerwriteConsistency.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteConsistency() (ret int) {
	condRecorderAuxMockPtrframerwriteConsistency.L.Lock()
	ret = recorderAuxMockPtrframerwriteConsistency
	condRecorderAuxMockPtrframerwriteConsistency.L.Unlock()
	return
}

// (recvf *framer)writeConsistency - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeConsistency(argcons Consistency) {
	FuncAuxMockPtrframerwriteConsistency, ok := apomock.GetRegisteredFunc("gocql.framer.writeConsistency")
	if ok {
		FuncAuxMockPtrframerwriteConsistency.(func(recvf *framer, argcons Consistency))(recvf, argcons)
	} else {
		panic("FuncAuxMockPtrframerwriteConsistency ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteConsistency()
	return
}

//
// Mock: (recvf frameOp)String()(reta string)
//

type MockArgsTypeframeOpString struct {
	ApomockCallNumber int
}

var LastMockArgsframeOpString MockArgsTypeframeOpString

// (recvf frameOp)AuxMockString()(reta string) - Generated mock function
func (recvf frameOp) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.frameOp.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.frameOp.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.frameOp.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockframeOpString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockframeOpString int = 0

var condRecorderAuxMockframeOpString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockframeOpString(i int) {
	condRecorderAuxMockframeOpString.L.Lock()
	for recorderAuxMockframeOpString < i {
		condRecorderAuxMockframeOpString.Wait()
	}
	condRecorderAuxMockframeOpString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockframeOpString() {
	condRecorderAuxMockframeOpString.L.Lock()
	recorderAuxMockframeOpString++
	condRecorderAuxMockframeOpString.L.Unlock()
	condRecorderAuxMockframeOpString.Broadcast()
}
func AuxMockGetRecorderAuxMockframeOpString() (ret int) {
	condRecorderAuxMockframeOpString.L.Lock()
	ret = recorderAuxMockframeOpString
	condRecorderAuxMockframeOpString.L.Unlock()
	return
}

// (recvf frameOp)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf frameOp) String() (reta string) {
	FuncAuxMockframeOpString, ok := apomock.GetRegisteredFunc("gocql.frameOp.String")
	if ok {
		reta = FuncAuxMockframeOpString.(func(recvf frameOp) (reta string))(recvf)
	} else {
		panic("FuncAuxMockframeOpString ")
	}
	AuxMockIncrementRecorderAuxMockframeOpString()
	return
}

//
// Mock: (recvf *framer)writeStartupFrame(argstreamID int, argoptions map[string]string)(reta error)
//

type MockArgsTypeframerwriteStartupFrame struct {
	ApomockCallNumber int
	ArgstreamID       int
	Argoptions        map[string]string
}

var LastMockArgsframerwriteStartupFrame MockArgsTypeframerwriteStartupFrame

// (recvf *framer)AuxMockwriteStartupFrame(argstreamID int, argoptions map[string]string)(reta error) - Generated mock function
func (recvf *framer) AuxMockwriteStartupFrame(argstreamID int, argoptions map[string]string) (reta error) {
	LastMockArgsframerwriteStartupFrame = MockArgsTypeframerwriteStartupFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteStartupFrame(),
		ArgstreamID:       argstreamID,
		Argoptions:        argoptions,
	}
	rargs, rerr := apomock.GetNext("gocql.framer.writeStartupFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.writeStartupFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.writeStartupFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrframerwriteStartupFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteStartupFrame int = 0

var condRecorderAuxMockPtrframerwriteStartupFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteStartupFrame(i int) {
	condRecorderAuxMockPtrframerwriteStartupFrame.L.Lock()
	for recorderAuxMockPtrframerwriteStartupFrame < i {
		condRecorderAuxMockPtrframerwriteStartupFrame.Wait()
	}
	condRecorderAuxMockPtrframerwriteStartupFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteStartupFrame() {
	condRecorderAuxMockPtrframerwriteStartupFrame.L.Lock()
	recorderAuxMockPtrframerwriteStartupFrame++
	condRecorderAuxMockPtrframerwriteStartupFrame.L.Unlock()
	condRecorderAuxMockPtrframerwriteStartupFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteStartupFrame() (ret int) {
	condRecorderAuxMockPtrframerwriteStartupFrame.L.Lock()
	ret = recorderAuxMockPtrframerwriteStartupFrame
	condRecorderAuxMockPtrframerwriteStartupFrame.L.Unlock()
	return
}

// (recvf *framer)writeStartupFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeStartupFrame(argstreamID int, argoptions map[string]string) (reta error) {
	FuncAuxMockPtrframerwriteStartupFrame, ok := apomock.GetRegisteredFunc("gocql.framer.writeStartupFrame")
	if ok {
		reta = FuncAuxMockPtrframerwriteStartupFrame.(func(recvf *framer, argstreamID int, argoptions map[string]string) (reta error))(recvf, argstreamID, argoptions)
	} else {
		panic("FuncAuxMockPtrframerwriteStartupFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteStartupFrame()
	return
}

//
// Mock: (recvf *framer)parseResultPrepared()(reta frame)
//

type MockArgsTypeframerparseResultPrepared struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseResultPrepared MockArgsTypeframerparseResultPrepared

// (recvf *framer)AuxMockparseResultPrepared()(reta frame) - Generated mock function
func (recvf *framer) AuxMockparseResultPrepared() (reta frame) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseResultPrepared")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseResultPrepared")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseResultPrepared")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	return
}

// RecorderAuxMockPtrframerparseResultPrepared  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseResultPrepared int = 0

var condRecorderAuxMockPtrframerparseResultPrepared *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseResultPrepared(i int) {
	condRecorderAuxMockPtrframerparseResultPrepared.L.Lock()
	for recorderAuxMockPtrframerparseResultPrepared < i {
		condRecorderAuxMockPtrframerparseResultPrepared.Wait()
	}
	condRecorderAuxMockPtrframerparseResultPrepared.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseResultPrepared() {
	condRecorderAuxMockPtrframerparseResultPrepared.L.Lock()
	recorderAuxMockPtrframerparseResultPrepared++
	condRecorderAuxMockPtrframerparseResultPrepared.L.Unlock()
	condRecorderAuxMockPtrframerparseResultPrepared.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseResultPrepared() (ret int) {
	condRecorderAuxMockPtrframerparseResultPrepared.L.Lock()
	ret = recorderAuxMockPtrframerparseResultPrepared
	condRecorderAuxMockPtrframerparseResultPrepared.L.Unlock()
	return
}

// (recvf *framer)parseResultPrepared - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseResultPrepared() (reta frame) {
	FuncAuxMockPtrframerparseResultPrepared, ok := apomock.GetRegisteredFunc("gocql.framer.parseResultPrepared")
	if ok {
		reta = FuncAuxMockPtrframerparseResultPrepared.(func(recvf *framer) (reta frame))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseResultPrepared ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseResultPrepared()
	return
}

//
// Mock: (recva *authenticateFrame)String()(reta string)
//

type MockArgsTypeauthenticateFrameString struct {
	ApomockCallNumber int
}

var LastMockArgsauthenticateFrameString MockArgsTypeauthenticateFrameString

// (recva *authenticateFrame)AuxMockString()(reta string) - Generated mock function
func (recva *authenticateFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.authenticateFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.authenticateFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.authenticateFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrauthenticateFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrauthenticateFrameString int = 0

var condRecorderAuxMockPtrauthenticateFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrauthenticateFrameString(i int) {
	condRecorderAuxMockPtrauthenticateFrameString.L.Lock()
	for recorderAuxMockPtrauthenticateFrameString < i {
		condRecorderAuxMockPtrauthenticateFrameString.Wait()
	}
	condRecorderAuxMockPtrauthenticateFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrauthenticateFrameString() {
	condRecorderAuxMockPtrauthenticateFrameString.L.Lock()
	recorderAuxMockPtrauthenticateFrameString++
	condRecorderAuxMockPtrauthenticateFrameString.L.Unlock()
	condRecorderAuxMockPtrauthenticateFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrauthenticateFrameString() (ret int) {
	condRecorderAuxMockPtrauthenticateFrameString.L.Lock()
	ret = recorderAuxMockPtrauthenticateFrameString
	condRecorderAuxMockPtrauthenticateFrameString.L.Unlock()
	return
}

// (recva *authenticateFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recva *authenticateFrame) String() (reta string) {
	FuncAuxMockPtrauthenticateFrameString, ok := apomock.GetRegisteredFunc("gocql.authenticateFrame.String")
	if ok {
		reta = FuncAuxMockPtrauthenticateFrameString.(func(recva *authenticateFrame) (reta string))(recva)
	} else {
		panic("FuncAuxMockPtrauthenticateFrameString ")
	}
	AuxMockIncrementRecorderAuxMockPtrauthenticateFrameString()
	return
}

//
// Mock: (recvf *framer)writeExecuteFrame(argstreamID int, argpreparedID []byte, argparams *queryParams)(reta error)
//

type MockArgsTypeframerwriteExecuteFrame struct {
	ApomockCallNumber int
	ArgstreamID       int
	ArgpreparedID     []byte
	Argparams         *queryParams
}

var LastMockArgsframerwriteExecuteFrame MockArgsTypeframerwriteExecuteFrame

// (recvf *framer)AuxMockwriteExecuteFrame(argstreamID int, argpreparedID []byte, argparams *queryParams)(reta error) - Generated mock function
func (recvf *framer) AuxMockwriteExecuteFrame(argstreamID int, argpreparedID []byte, argparams *queryParams) (reta error) {
	LastMockArgsframerwriteExecuteFrame = MockArgsTypeframerwriteExecuteFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteExecuteFrame(),
		ArgstreamID:       argstreamID,
		ArgpreparedID:     argpreparedID,
		Argparams:         argparams,
	}
	rargs, rerr := apomock.GetNext("gocql.framer.writeExecuteFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.writeExecuteFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.writeExecuteFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrframerwriteExecuteFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteExecuteFrame int = 0

var condRecorderAuxMockPtrframerwriteExecuteFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteExecuteFrame(i int) {
	condRecorderAuxMockPtrframerwriteExecuteFrame.L.Lock()
	for recorderAuxMockPtrframerwriteExecuteFrame < i {
		condRecorderAuxMockPtrframerwriteExecuteFrame.Wait()
	}
	condRecorderAuxMockPtrframerwriteExecuteFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteExecuteFrame() {
	condRecorderAuxMockPtrframerwriteExecuteFrame.L.Lock()
	recorderAuxMockPtrframerwriteExecuteFrame++
	condRecorderAuxMockPtrframerwriteExecuteFrame.L.Unlock()
	condRecorderAuxMockPtrframerwriteExecuteFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteExecuteFrame() (ret int) {
	condRecorderAuxMockPtrframerwriteExecuteFrame.L.Lock()
	ret = recorderAuxMockPtrframerwriteExecuteFrame
	condRecorderAuxMockPtrframerwriteExecuteFrame.L.Unlock()
	return
}

// (recvf *framer)writeExecuteFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeExecuteFrame(argstreamID int, argpreparedID []byte, argparams *queryParams) (reta error) {
	FuncAuxMockPtrframerwriteExecuteFrame, ok := apomock.GetRegisteredFunc("gocql.framer.writeExecuteFrame")
	if ok {
		reta = FuncAuxMockPtrframerwriteExecuteFrame.(func(recvf *framer, argstreamID int, argpreparedID []byte, argparams *queryParams) (reta error))(recvf, argstreamID, argpreparedID, argparams)
	} else {
		panic("FuncAuxMockPtrframerwriteExecuteFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteExecuteFrame()
	return
}

//
// Mock: (recvf *framer)writeRegisterFrame(argstreamID int, argw *writeRegisterFrame)(reta error)
//

type MockArgsTypeframerwriteRegisterFrame struct {
	ApomockCallNumber int
	ArgstreamID       int
	Argw              *writeRegisterFrame
}

var LastMockArgsframerwriteRegisterFrame MockArgsTypeframerwriteRegisterFrame

// (recvf *framer)AuxMockwriteRegisterFrame(argstreamID int, argw *writeRegisterFrame)(reta error) - Generated mock function
func (recvf *framer) AuxMockwriteRegisterFrame(argstreamID int, argw *writeRegisterFrame) (reta error) {
	LastMockArgsframerwriteRegisterFrame = MockArgsTypeframerwriteRegisterFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteRegisterFrame(),
		ArgstreamID:       argstreamID,
		Argw:              argw,
	}
	rargs, rerr := apomock.GetNext("gocql.framer.writeRegisterFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.writeRegisterFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.writeRegisterFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrframerwriteRegisterFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteRegisterFrame int = 0

var condRecorderAuxMockPtrframerwriteRegisterFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteRegisterFrame(i int) {
	condRecorderAuxMockPtrframerwriteRegisterFrame.L.Lock()
	for recorderAuxMockPtrframerwriteRegisterFrame < i {
		condRecorderAuxMockPtrframerwriteRegisterFrame.Wait()
	}
	condRecorderAuxMockPtrframerwriteRegisterFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteRegisterFrame() {
	condRecorderAuxMockPtrframerwriteRegisterFrame.L.Lock()
	recorderAuxMockPtrframerwriteRegisterFrame++
	condRecorderAuxMockPtrframerwriteRegisterFrame.L.Unlock()
	condRecorderAuxMockPtrframerwriteRegisterFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteRegisterFrame() (ret int) {
	condRecorderAuxMockPtrframerwriteRegisterFrame.L.Lock()
	ret = recorderAuxMockPtrframerwriteRegisterFrame
	condRecorderAuxMockPtrframerwriteRegisterFrame.L.Unlock()
	return
}

// (recvf *framer)writeRegisterFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeRegisterFrame(argstreamID int, argw *writeRegisterFrame) (reta error) {
	FuncAuxMockPtrframerwriteRegisterFrame, ok := apomock.GetRegisteredFunc("gocql.framer.writeRegisterFrame")
	if ok {
		reta = FuncAuxMockPtrframerwriteRegisterFrame.(func(recvf *framer, argstreamID int, argw *writeRegisterFrame) (reta error))(recvf, argstreamID, argw)
	} else {
		panic("FuncAuxMockPtrframerwriteRegisterFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteRegisterFrame()
	return
}

//
// Mock: (recva *authChallengeFrame)String()(reta string)
//

type MockArgsTypeauthChallengeFrameString struct {
	ApomockCallNumber int
}

var LastMockArgsauthChallengeFrameString MockArgsTypeauthChallengeFrameString

// (recva *authChallengeFrame)AuxMockString()(reta string) - Generated mock function
func (recva *authChallengeFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.authChallengeFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.authChallengeFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.authChallengeFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrauthChallengeFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrauthChallengeFrameString int = 0

var condRecorderAuxMockPtrauthChallengeFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrauthChallengeFrameString(i int) {
	condRecorderAuxMockPtrauthChallengeFrameString.L.Lock()
	for recorderAuxMockPtrauthChallengeFrameString < i {
		condRecorderAuxMockPtrauthChallengeFrameString.Wait()
	}
	condRecorderAuxMockPtrauthChallengeFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrauthChallengeFrameString() {
	condRecorderAuxMockPtrauthChallengeFrameString.L.Lock()
	recorderAuxMockPtrauthChallengeFrameString++
	condRecorderAuxMockPtrauthChallengeFrameString.L.Unlock()
	condRecorderAuxMockPtrauthChallengeFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrauthChallengeFrameString() (ret int) {
	condRecorderAuxMockPtrauthChallengeFrameString.L.Lock()
	ret = recorderAuxMockPtrauthChallengeFrameString
	condRecorderAuxMockPtrauthChallengeFrameString.L.Unlock()
	return
}

// (recva *authChallengeFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recva *authChallengeFrame) String() (reta string) {
	FuncAuxMockPtrauthChallengeFrameString, ok := apomock.GetRegisteredFunc("gocql.authChallengeFrame.String")
	if ok {
		reta = FuncAuxMockPtrauthChallengeFrameString.(func(recva *authChallengeFrame) (reta string))(recva)
	} else {
		panic("FuncAuxMockPtrauthChallengeFrameString ")
	}
	AuxMockIncrementRecorderAuxMockPtrauthChallengeFrameString()
	return
}

//
// Mock: (recvf *framer)readBytesInternal()(reta []byte, retb error)
//

type MockArgsTypeframerreadBytesInternal struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadBytesInternal MockArgsTypeframerreadBytesInternal

// (recvf *framer)AuxMockreadBytesInternal()(reta []byte, retb error) - Generated mock function
func (recvf *framer) AuxMockreadBytesInternal() (reta []byte, retb error) {
	rargs, rerr := apomock.GetNext("gocql.framer.readBytesInternal")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readBytesInternal")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.framer.readBytesInternal")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrframerreadBytesInternal  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadBytesInternal int = 0

var condRecorderAuxMockPtrframerreadBytesInternal *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadBytesInternal(i int) {
	condRecorderAuxMockPtrframerreadBytesInternal.L.Lock()
	for recorderAuxMockPtrframerreadBytesInternal < i {
		condRecorderAuxMockPtrframerreadBytesInternal.Wait()
	}
	condRecorderAuxMockPtrframerreadBytesInternal.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadBytesInternal() {
	condRecorderAuxMockPtrframerreadBytesInternal.L.Lock()
	recorderAuxMockPtrframerreadBytesInternal++
	condRecorderAuxMockPtrframerreadBytesInternal.L.Unlock()
	condRecorderAuxMockPtrframerreadBytesInternal.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadBytesInternal() (ret int) {
	condRecorderAuxMockPtrframerreadBytesInternal.L.Lock()
	ret = recorderAuxMockPtrframerreadBytesInternal
	condRecorderAuxMockPtrframerreadBytesInternal.L.Unlock()
	return
}

// (recvf *framer)readBytesInternal - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readBytesInternal() (reta []byte, retb error) {
	FuncAuxMockPtrframerreadBytesInternal, ok := apomock.GetRegisteredFunc("gocql.framer.readBytesInternal")
	if ok {
		reta, retb = FuncAuxMockPtrframerreadBytesInternal.(func(recvf *framer) (reta []byte, retb error))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadBytesInternal ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadBytesInternal()
	return
}

//
// Mock: (recvf *framer)readStringMap()(reta map[string]string)
//

type MockArgsTypeframerreadStringMap struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadStringMap MockArgsTypeframerreadStringMap

// (recvf *framer)AuxMockreadStringMap()(reta map[string]string) - Generated mock function
func (recvf *framer) AuxMockreadStringMap() (reta map[string]string) {
	rargs, rerr := apomock.GetNext("gocql.framer.readStringMap")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readStringMap")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readStringMap")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(map[string]string)
	}
	return
}

// RecorderAuxMockPtrframerreadStringMap  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadStringMap int = 0

var condRecorderAuxMockPtrframerreadStringMap *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadStringMap(i int) {
	condRecorderAuxMockPtrframerreadStringMap.L.Lock()
	for recorderAuxMockPtrframerreadStringMap < i {
		condRecorderAuxMockPtrframerreadStringMap.Wait()
	}
	condRecorderAuxMockPtrframerreadStringMap.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadStringMap() {
	condRecorderAuxMockPtrframerreadStringMap.L.Lock()
	recorderAuxMockPtrframerreadStringMap++
	condRecorderAuxMockPtrframerreadStringMap.L.Unlock()
	condRecorderAuxMockPtrframerreadStringMap.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadStringMap() (ret int) {
	condRecorderAuxMockPtrframerreadStringMap.L.Lock()
	ret = recorderAuxMockPtrframerreadStringMap
	condRecorderAuxMockPtrframerreadStringMap.L.Unlock()
	return
}

// (recvf *framer)readStringMap - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readStringMap() (reta map[string]string) {
	FuncAuxMockPtrframerreadStringMap, ok := apomock.GetRegisteredFunc("gocql.framer.readStringMap")
	if ok {
		reta = FuncAuxMockPtrframerreadStringMap.(func(recvf *framer) (reta map[string]string))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadStringMap ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadStringMap()
	return
}

//
// Mock: writeInt(argp []byte, argn int32)()
//

type MockArgsTypewriteInt struct {
	ApomockCallNumber int
	Argp              []byte
	Argn              int32
}

var LastMockArgswriteInt MockArgsTypewriteInt

// AuxMockwriteInt(argp []byte, argn int32)() - Generated mock function
func AuxMockwriteInt(argp []byte, argn int32) {
	LastMockArgswriteInt = MockArgsTypewriteInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockwriteInt(),
		Argp:              argp,
		Argn:              argn,
	}
	return
}

// RecorderAuxMockwriteInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockwriteInt int = 0

var condRecorderAuxMockwriteInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockwriteInt(i int) {
	condRecorderAuxMockwriteInt.L.Lock()
	for recorderAuxMockwriteInt < i {
		condRecorderAuxMockwriteInt.Wait()
	}
	condRecorderAuxMockwriteInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockwriteInt() {
	condRecorderAuxMockwriteInt.L.Lock()
	recorderAuxMockwriteInt++
	condRecorderAuxMockwriteInt.L.Unlock()
	condRecorderAuxMockwriteInt.Broadcast()
}
func AuxMockGetRecorderAuxMockwriteInt() (ret int) {
	condRecorderAuxMockwriteInt.L.Lock()
	ret = recorderAuxMockwriteInt
	condRecorderAuxMockwriteInt.L.Unlock()
	return
}

// writeInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func writeInt(argp []byte, argn int32) {
	FuncAuxMockwriteInt, ok := apomock.GetRegisteredFunc("gocql.writeInt")
	if ok {
		FuncAuxMockwriteInt.(func(argp []byte, argn int32))(argp, argn)
	} else {
		panic("FuncAuxMockwriteInt ")
	}
	AuxMockIncrementRecorderAuxMockwriteInt()
	return
}

//
// Mock: (recvr resultMetadata)String()(reta string)
//

type MockArgsTyperesultMetadataString struct {
	ApomockCallNumber int
}

var LastMockArgsresultMetadataString MockArgsTyperesultMetadataString

// (recvr resultMetadata)AuxMockString()(reta string) - Generated mock function
func (recvr resultMetadata) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.resultMetadata.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.resultMetadata.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.resultMetadata.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockresultMetadataString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockresultMetadataString int = 0

var condRecorderAuxMockresultMetadataString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockresultMetadataString(i int) {
	condRecorderAuxMockresultMetadataString.L.Lock()
	for recorderAuxMockresultMetadataString < i {
		condRecorderAuxMockresultMetadataString.Wait()
	}
	condRecorderAuxMockresultMetadataString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockresultMetadataString() {
	condRecorderAuxMockresultMetadataString.L.Lock()
	recorderAuxMockresultMetadataString++
	condRecorderAuxMockresultMetadataString.L.Unlock()
	condRecorderAuxMockresultMetadataString.Broadcast()
}
func AuxMockGetRecorderAuxMockresultMetadataString() (ret int) {
	condRecorderAuxMockresultMetadataString.L.Lock()
	ret = recorderAuxMockresultMetadataString
	condRecorderAuxMockresultMetadataString.L.Unlock()
	return
}

// (recvr resultMetadata)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr resultMetadata) String() (reta string) {
	FuncAuxMockresultMetadataString, ok := apomock.GetRegisteredFunc("gocql.resultMetadata.String")
	if ok {
		reta = FuncAuxMockresultMetadataString.(func(recvr resultMetadata) (reta string))(recvr)
	} else {
		panic("FuncAuxMockresultMetadataString ")
	}
	AuxMockIncrementRecorderAuxMockresultMetadataString()
	return
}

//
// Mock: (recvf *framer)parseResultSchemaChange()(reta frame)
//

type MockArgsTypeframerparseResultSchemaChange struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseResultSchemaChange MockArgsTypeframerparseResultSchemaChange

// (recvf *framer)AuxMockparseResultSchemaChange()(reta frame) - Generated mock function
func (recvf *framer) AuxMockparseResultSchemaChange() (reta frame) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseResultSchemaChange")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseResultSchemaChange")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseResultSchemaChange")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	return
}

// RecorderAuxMockPtrframerparseResultSchemaChange  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseResultSchemaChange int = 0

var condRecorderAuxMockPtrframerparseResultSchemaChange *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseResultSchemaChange(i int) {
	condRecorderAuxMockPtrframerparseResultSchemaChange.L.Lock()
	for recorderAuxMockPtrframerparseResultSchemaChange < i {
		condRecorderAuxMockPtrframerparseResultSchemaChange.Wait()
	}
	condRecorderAuxMockPtrframerparseResultSchemaChange.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseResultSchemaChange() {
	condRecorderAuxMockPtrframerparseResultSchemaChange.L.Lock()
	recorderAuxMockPtrframerparseResultSchemaChange++
	condRecorderAuxMockPtrframerparseResultSchemaChange.L.Unlock()
	condRecorderAuxMockPtrframerparseResultSchemaChange.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseResultSchemaChange() (ret int) {
	condRecorderAuxMockPtrframerparseResultSchemaChange.L.Lock()
	ret = recorderAuxMockPtrframerparseResultSchemaChange
	condRecorderAuxMockPtrframerparseResultSchemaChange.L.Unlock()
	return
}

// (recvf *framer)parseResultSchemaChange - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseResultSchemaChange() (reta frame) {
	FuncAuxMockPtrframerparseResultSchemaChange, ok := apomock.GetRegisteredFunc("gocql.framer.parseResultSchemaChange")
	if ok {
		reta = FuncAuxMockPtrframerparseResultSchemaChange.(func(recvf *framer) (reta frame))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseResultSchemaChange ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseResultSchemaChange()
	return
}

//
// Mock: (recvf *framer)readUUID()(reta *UUID)
//

type MockArgsTypeframerreadUUID struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadUUID MockArgsTypeframerreadUUID

// (recvf *framer)AuxMockreadUUID()(reta *UUID) - Generated mock function
func (recvf *framer) AuxMockreadUUID() (reta *UUID) {
	rargs, rerr := apomock.GetNext("gocql.framer.readUUID")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readUUID")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readUUID")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*UUID)
	}
	return
}

// RecorderAuxMockPtrframerreadUUID  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadUUID int = 0

var condRecorderAuxMockPtrframerreadUUID *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadUUID(i int) {
	condRecorderAuxMockPtrframerreadUUID.L.Lock()
	for recorderAuxMockPtrframerreadUUID < i {
		condRecorderAuxMockPtrframerreadUUID.Wait()
	}
	condRecorderAuxMockPtrframerreadUUID.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadUUID() {
	condRecorderAuxMockPtrframerreadUUID.L.Lock()
	recorderAuxMockPtrframerreadUUID++
	condRecorderAuxMockPtrframerreadUUID.L.Unlock()
	condRecorderAuxMockPtrframerreadUUID.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadUUID() (ret int) {
	condRecorderAuxMockPtrframerreadUUID.L.Lock()
	ret = recorderAuxMockPtrframerreadUUID
	condRecorderAuxMockPtrframerreadUUID.L.Unlock()
	return
}

// (recvf *framer)readUUID - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readUUID() (reta *UUID) {
	FuncAuxMockPtrframerreadUUID, ok := apomock.GetRegisteredFunc("gocql.framer.readUUID")
	if ok {
		reta = FuncAuxMockPtrframerreadUUID.(func(recvf *framer) (reta *UUID))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadUUID ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadUUID()
	return
}

//
// Mock: (recvp protoVersion)request()(reta bool)
//

type MockArgsTypeprotoVersionrequest struct {
	ApomockCallNumber int
}

var LastMockArgsprotoVersionrequest MockArgsTypeprotoVersionrequest

// (recvp protoVersion)AuxMockrequest()(reta bool) - Generated mock function
func (recvp protoVersion) AuxMockrequest() (reta bool) {
	rargs, rerr := apomock.GetNext("gocql.protoVersion.request")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.protoVersion.request")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.protoVersion.request")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockprotoVersionrequest  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockprotoVersionrequest int = 0

var condRecorderAuxMockprotoVersionrequest *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockprotoVersionrequest(i int) {
	condRecorderAuxMockprotoVersionrequest.L.Lock()
	for recorderAuxMockprotoVersionrequest < i {
		condRecorderAuxMockprotoVersionrequest.Wait()
	}
	condRecorderAuxMockprotoVersionrequest.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockprotoVersionrequest() {
	condRecorderAuxMockprotoVersionrequest.L.Lock()
	recorderAuxMockprotoVersionrequest++
	condRecorderAuxMockprotoVersionrequest.L.Unlock()
	condRecorderAuxMockprotoVersionrequest.Broadcast()
}
func AuxMockGetRecorderAuxMockprotoVersionrequest() (ret int) {
	condRecorderAuxMockprotoVersionrequest.L.Lock()
	ret = recorderAuxMockprotoVersionrequest
	condRecorderAuxMockprotoVersionrequest.L.Unlock()
	return
}

// (recvp protoVersion)request - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp protoVersion) request() (reta bool) {
	FuncAuxMockprotoVersionrequest, ok := apomock.GetRegisteredFunc("gocql.protoVersion.request")
	if ok {
		reta = FuncAuxMockprotoVersionrequest.(func(recvp protoVersion) (reta bool))(recvp)
	} else {
		panic("FuncAuxMockprotoVersionrequest ")
	}
	AuxMockIncrementRecorderAuxMockprotoVersionrequest()
	return
}

//
// Mock: (recvf *framer)parsePreparedMetadata()(reta preparedMetadata)
//

type MockArgsTypeframerparsePreparedMetadata struct {
	ApomockCallNumber int
}

var LastMockArgsframerparsePreparedMetadata MockArgsTypeframerparsePreparedMetadata

// (recvf *framer)AuxMockparsePreparedMetadata()(reta preparedMetadata) - Generated mock function
func (recvf *framer) AuxMockparsePreparedMetadata() (reta preparedMetadata) {
	rargs, rerr := apomock.GetNext("gocql.framer.parsePreparedMetadata")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parsePreparedMetadata")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parsePreparedMetadata")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(preparedMetadata)
	}
	return
}

// RecorderAuxMockPtrframerparsePreparedMetadata  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparsePreparedMetadata int = 0

var condRecorderAuxMockPtrframerparsePreparedMetadata *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparsePreparedMetadata(i int) {
	condRecorderAuxMockPtrframerparsePreparedMetadata.L.Lock()
	for recorderAuxMockPtrframerparsePreparedMetadata < i {
		condRecorderAuxMockPtrframerparsePreparedMetadata.Wait()
	}
	condRecorderAuxMockPtrframerparsePreparedMetadata.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparsePreparedMetadata() {
	condRecorderAuxMockPtrframerparsePreparedMetadata.L.Lock()
	recorderAuxMockPtrframerparsePreparedMetadata++
	condRecorderAuxMockPtrframerparsePreparedMetadata.L.Unlock()
	condRecorderAuxMockPtrframerparsePreparedMetadata.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparsePreparedMetadata() (ret int) {
	condRecorderAuxMockPtrframerparsePreparedMetadata.L.Lock()
	ret = recorderAuxMockPtrframerparsePreparedMetadata
	condRecorderAuxMockPtrframerparsePreparedMetadata.L.Unlock()
	return
}

// (recvf *framer)parsePreparedMetadata - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parsePreparedMetadata() (reta preparedMetadata) {
	FuncAuxMockPtrframerparsePreparedMetadata, ok := apomock.GetRegisteredFunc("gocql.framer.parsePreparedMetadata")
	if ok {
		reta = FuncAuxMockPtrframerparsePreparedMetadata.(func(recvf *framer) (reta preparedMetadata))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparsePreparedMetadata ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparsePreparedMetadata()
	return
}

//
// Mock: (recvf *framer)writeAuthResponseFrame(argstreamID int, argdata []byte)(reta error)
//

type MockArgsTypeframerwriteAuthResponseFrame struct {
	ApomockCallNumber int
	ArgstreamID       int
	Argdata           []byte
}

var LastMockArgsframerwriteAuthResponseFrame MockArgsTypeframerwriteAuthResponseFrame

// (recvf *framer)AuxMockwriteAuthResponseFrame(argstreamID int, argdata []byte)(reta error) - Generated mock function
func (recvf *framer) AuxMockwriteAuthResponseFrame(argstreamID int, argdata []byte) (reta error) {
	LastMockArgsframerwriteAuthResponseFrame = MockArgsTypeframerwriteAuthResponseFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteAuthResponseFrame(),
		ArgstreamID:       argstreamID,
		Argdata:           argdata,
	}
	rargs, rerr := apomock.GetNext("gocql.framer.writeAuthResponseFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.writeAuthResponseFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.writeAuthResponseFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrframerwriteAuthResponseFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteAuthResponseFrame int = 0

var condRecorderAuxMockPtrframerwriteAuthResponseFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteAuthResponseFrame(i int) {
	condRecorderAuxMockPtrframerwriteAuthResponseFrame.L.Lock()
	for recorderAuxMockPtrframerwriteAuthResponseFrame < i {
		condRecorderAuxMockPtrframerwriteAuthResponseFrame.Wait()
	}
	condRecorderAuxMockPtrframerwriteAuthResponseFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteAuthResponseFrame() {
	condRecorderAuxMockPtrframerwriteAuthResponseFrame.L.Lock()
	recorderAuxMockPtrframerwriteAuthResponseFrame++
	condRecorderAuxMockPtrframerwriteAuthResponseFrame.L.Unlock()
	condRecorderAuxMockPtrframerwriteAuthResponseFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteAuthResponseFrame() (ret int) {
	condRecorderAuxMockPtrframerwriteAuthResponseFrame.L.Lock()
	ret = recorderAuxMockPtrframerwriteAuthResponseFrame
	condRecorderAuxMockPtrframerwriteAuthResponseFrame.L.Unlock()
	return
}

// (recvf *framer)writeAuthResponseFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeAuthResponseFrame(argstreamID int, argdata []byte) (reta error) {
	FuncAuxMockPtrframerwriteAuthResponseFrame, ok := apomock.GetRegisteredFunc("gocql.framer.writeAuthResponseFrame")
	if ok {
		reta = FuncAuxMockPtrframerwriteAuthResponseFrame.(func(recvf *framer, argstreamID int, argdata []byte) (reta error))(recvf, argstreamID, argdata)
	} else {
		panic("FuncAuxMockPtrframerwriteAuthResponseFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteAuthResponseFrame()
	return
}

//
// Mock: (recvf *framer)writeBatchFrame(argstreamID int, argw *writeBatchFrame)(reta error)
//

type MockArgsTypeframerwriteBatchFrame struct {
	ApomockCallNumber int
	ArgstreamID       int
	Argw              *writeBatchFrame
}

var LastMockArgsframerwriteBatchFrame MockArgsTypeframerwriteBatchFrame

// (recvf *framer)AuxMockwriteBatchFrame(argstreamID int, argw *writeBatchFrame)(reta error) - Generated mock function
func (recvf *framer) AuxMockwriteBatchFrame(argstreamID int, argw *writeBatchFrame) (reta error) {
	LastMockArgsframerwriteBatchFrame = MockArgsTypeframerwriteBatchFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteBatchFrame(),
		ArgstreamID:       argstreamID,
		Argw:              argw,
	}
	rargs, rerr := apomock.GetNext("gocql.framer.writeBatchFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.writeBatchFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.writeBatchFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrframerwriteBatchFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteBatchFrame int = 0

var condRecorderAuxMockPtrframerwriteBatchFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteBatchFrame(i int) {
	condRecorderAuxMockPtrframerwriteBatchFrame.L.Lock()
	for recorderAuxMockPtrframerwriteBatchFrame < i {
		condRecorderAuxMockPtrframerwriteBatchFrame.Wait()
	}
	condRecorderAuxMockPtrframerwriteBatchFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteBatchFrame() {
	condRecorderAuxMockPtrframerwriteBatchFrame.L.Lock()
	recorderAuxMockPtrframerwriteBatchFrame++
	condRecorderAuxMockPtrframerwriteBatchFrame.L.Unlock()
	condRecorderAuxMockPtrframerwriteBatchFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteBatchFrame() (ret int) {
	condRecorderAuxMockPtrframerwriteBatchFrame.L.Lock()
	ret = recorderAuxMockPtrframerwriteBatchFrame
	condRecorderAuxMockPtrframerwriteBatchFrame.L.Unlock()
	return
}

// (recvf *framer)writeBatchFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeBatchFrame(argstreamID int, argw *writeBatchFrame) (reta error) {
	FuncAuxMockPtrframerwriteBatchFrame, ok := apomock.GetRegisteredFunc("gocql.framer.writeBatchFrame")
	if ok {
		reta = FuncAuxMockPtrframerwriteBatchFrame.(func(recvf *framer, argstreamID int, argw *writeBatchFrame) (reta error))(recvf, argstreamID, argw)
	} else {
		panic("FuncAuxMockPtrframerwriteBatchFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteBatchFrame()
	return
}

//
// Mock: (recvf *framer)readBytes()(reta []byte)
//

type MockArgsTypeframerreadBytes struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadBytes MockArgsTypeframerreadBytes

// (recvf *framer)AuxMockreadBytes()(reta []byte) - Generated mock function
func (recvf *framer) AuxMockreadBytes() (reta []byte) {
	rargs, rerr := apomock.GetNext("gocql.framer.readBytes")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readBytes")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readBytes")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockPtrframerreadBytes  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadBytes int = 0

var condRecorderAuxMockPtrframerreadBytes *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadBytes(i int) {
	condRecorderAuxMockPtrframerreadBytes.L.Lock()
	for recorderAuxMockPtrframerreadBytes < i {
		condRecorderAuxMockPtrframerreadBytes.Wait()
	}
	condRecorderAuxMockPtrframerreadBytes.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadBytes() {
	condRecorderAuxMockPtrframerreadBytes.L.Lock()
	recorderAuxMockPtrframerreadBytes++
	condRecorderAuxMockPtrframerreadBytes.L.Unlock()
	condRecorderAuxMockPtrframerreadBytes.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadBytes() (ret int) {
	condRecorderAuxMockPtrframerreadBytes.L.Lock()
	ret = recorderAuxMockPtrframerreadBytes
	condRecorderAuxMockPtrframerreadBytes.L.Unlock()
	return
}

// (recvf *framer)readBytes - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readBytes() (reta []byte) {
	FuncAuxMockPtrframerreadBytes, ok := apomock.GetRegisteredFunc("gocql.framer.readBytes")
	if ok {
		reta = FuncAuxMockPtrframerreadBytes.(func(recvf *framer) (reta []byte))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadBytes ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadBytes()
	return
}

//
// Mock: (recvf *framer)readInt()(retn int)
//

type MockArgsTypeframerreadInt struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadInt MockArgsTypeframerreadInt

// (recvf *framer)AuxMockreadInt()(retn int) - Generated mock function
func (recvf *framer) AuxMockreadInt() (retn int) {
	rargs, rerr := apomock.GetNext("gocql.framer.readInt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readInt")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readInt")
	}
	if rargs.GetArg(0) != nil {
		retn = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockPtrframerreadInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadInt int = 0

var condRecorderAuxMockPtrframerreadInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadInt(i int) {
	condRecorderAuxMockPtrframerreadInt.L.Lock()
	for recorderAuxMockPtrframerreadInt < i {
		condRecorderAuxMockPtrframerreadInt.Wait()
	}
	condRecorderAuxMockPtrframerreadInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadInt() {
	condRecorderAuxMockPtrframerreadInt.L.Lock()
	recorderAuxMockPtrframerreadInt++
	condRecorderAuxMockPtrframerreadInt.L.Unlock()
	condRecorderAuxMockPtrframerreadInt.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadInt() (ret int) {
	condRecorderAuxMockPtrframerreadInt.L.Lock()
	ret = recorderAuxMockPtrframerreadInt
	condRecorderAuxMockPtrframerreadInt.L.Unlock()
	return
}

// (recvf *framer)readInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readInt() (retn int) {
	FuncAuxMockPtrframerreadInt, ok := apomock.GetRegisteredFunc("gocql.framer.readInt")
	if ok {
		retn = FuncAuxMockPtrframerreadInt.(func(recvf *framer) (retn int))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadInt ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadInt()
	return
}

//
// Mock: (recvf *framer)readStringList()(reta []string)
//

type MockArgsTypeframerreadStringList struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadStringList MockArgsTypeframerreadStringList

// (recvf *framer)AuxMockreadStringList()(reta []string) - Generated mock function
func (recvf *framer) AuxMockreadStringList() (reta []string) {
	rargs, rerr := apomock.GetNext("gocql.framer.readStringList")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readStringList")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readStringList")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]string)
	}
	return
}

// RecorderAuxMockPtrframerreadStringList  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadStringList int = 0

var condRecorderAuxMockPtrframerreadStringList *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadStringList(i int) {
	condRecorderAuxMockPtrframerreadStringList.L.Lock()
	for recorderAuxMockPtrframerreadStringList < i {
		condRecorderAuxMockPtrframerreadStringList.Wait()
	}
	condRecorderAuxMockPtrframerreadStringList.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadStringList() {
	condRecorderAuxMockPtrframerreadStringList.L.Lock()
	recorderAuxMockPtrframerreadStringList++
	condRecorderAuxMockPtrframerreadStringList.L.Unlock()
	condRecorderAuxMockPtrframerreadStringList.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadStringList() (ret int) {
	condRecorderAuxMockPtrframerreadStringList.L.Lock()
	ret = recorderAuxMockPtrframerreadStringList
	condRecorderAuxMockPtrframerreadStringList.L.Unlock()
	return
}

// (recvf *framer)readStringList - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readStringList() (reta []string) {
	FuncAuxMockPtrframerreadStringList, ok := apomock.GetRegisteredFunc("gocql.framer.readStringList")
	if ok {
		reta = FuncAuxMockPtrframerreadStringList.(func(recvf *framer) (reta []string))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadStringList ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadStringList()
	return
}

//
// Mock: (recvf *framer)parseErrorFrame()(reta frame)
//

type MockArgsTypeframerparseErrorFrame struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseErrorFrame MockArgsTypeframerparseErrorFrame

// (recvf *framer)AuxMockparseErrorFrame()(reta frame) - Generated mock function
func (recvf *framer) AuxMockparseErrorFrame() (reta frame) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseErrorFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseErrorFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseErrorFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	return
}

// RecorderAuxMockPtrframerparseErrorFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseErrorFrame int = 0

var condRecorderAuxMockPtrframerparseErrorFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseErrorFrame(i int) {
	condRecorderAuxMockPtrframerparseErrorFrame.L.Lock()
	for recorderAuxMockPtrframerparseErrorFrame < i {
		condRecorderAuxMockPtrframerparseErrorFrame.Wait()
	}
	condRecorderAuxMockPtrframerparseErrorFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseErrorFrame() {
	condRecorderAuxMockPtrframerparseErrorFrame.L.Lock()
	recorderAuxMockPtrframerparseErrorFrame++
	condRecorderAuxMockPtrframerparseErrorFrame.L.Unlock()
	condRecorderAuxMockPtrframerparseErrorFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseErrorFrame() (ret int) {
	condRecorderAuxMockPtrframerparseErrorFrame.L.Lock()
	ret = recorderAuxMockPtrframerparseErrorFrame
	condRecorderAuxMockPtrframerparseErrorFrame.L.Unlock()
	return
}

// (recvf *framer)parseErrorFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseErrorFrame() (reta frame) {
	FuncAuxMockPtrframerparseErrorFrame, ok := apomock.GetRegisteredFunc("gocql.framer.parseErrorFrame")
	if ok {
		reta = FuncAuxMockPtrframerparseErrorFrame.(func(recvf *framer) (reta frame))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseErrorFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseErrorFrame()
	return
}

//
// Mock: (recvf *framer)readTypeInfo()(reta TypeInfo)
//

type MockArgsTypeframerreadTypeInfo struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadTypeInfo MockArgsTypeframerreadTypeInfo

// (recvf *framer)AuxMockreadTypeInfo()(reta TypeInfo) - Generated mock function
func (recvf *framer) AuxMockreadTypeInfo() (reta TypeInfo) {
	rargs, rerr := apomock.GetNext("gocql.framer.readTypeInfo")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readTypeInfo")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readTypeInfo")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(TypeInfo)
	}
	return
}

// RecorderAuxMockPtrframerreadTypeInfo  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadTypeInfo int = 0

var condRecorderAuxMockPtrframerreadTypeInfo *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadTypeInfo(i int) {
	condRecorderAuxMockPtrframerreadTypeInfo.L.Lock()
	for recorderAuxMockPtrframerreadTypeInfo < i {
		condRecorderAuxMockPtrframerreadTypeInfo.Wait()
	}
	condRecorderAuxMockPtrframerreadTypeInfo.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadTypeInfo() {
	condRecorderAuxMockPtrframerreadTypeInfo.L.Lock()
	recorderAuxMockPtrframerreadTypeInfo++
	condRecorderAuxMockPtrframerreadTypeInfo.L.Unlock()
	condRecorderAuxMockPtrframerreadTypeInfo.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadTypeInfo() (ret int) {
	condRecorderAuxMockPtrframerreadTypeInfo.L.Lock()
	ret = recorderAuxMockPtrframerreadTypeInfo
	condRecorderAuxMockPtrframerreadTypeInfo.L.Unlock()
	return
}

// (recvf *framer)readTypeInfo - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readTypeInfo() (reta TypeInfo) {
	FuncAuxMockPtrframerreadTypeInfo, ok := apomock.GetRegisteredFunc("gocql.framer.readTypeInfo")
	if ok {
		reta = FuncAuxMockPtrframerreadTypeInfo.(func(recvf *framer) (reta TypeInfo))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadTypeInfo ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadTypeInfo()
	return
}

//
// Mock: (recvf *framer)readCol(argcol *ColumnInfo, argmeta *resultMetadata, argglobalSpec bool, argkeyspace string, argtable string)()
//

type MockArgsTypeframerreadCol struct {
	ApomockCallNumber int
	Argcol            *ColumnInfo
	Argmeta           *resultMetadata
	ArgglobalSpec     bool
	Argkeyspace       string
	Argtable          string
}

var LastMockArgsframerreadCol MockArgsTypeframerreadCol

// (recvf *framer)AuxMockreadCol(argcol *ColumnInfo, argmeta *resultMetadata, argglobalSpec bool, argkeyspace string, argtable string)() - Generated mock function
func (recvf *framer) AuxMockreadCol(argcol *ColumnInfo, argmeta *resultMetadata, argglobalSpec bool, argkeyspace string, argtable string) {
	LastMockArgsframerreadCol = MockArgsTypeframerreadCol{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerreadCol(),
		Argcol:            argcol,
		Argmeta:           argmeta,
		ArgglobalSpec:     argglobalSpec,
		Argkeyspace:       argkeyspace,
		Argtable:          argtable,
	}
	return
}

// RecorderAuxMockPtrframerreadCol  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadCol int = 0

var condRecorderAuxMockPtrframerreadCol *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadCol(i int) {
	condRecorderAuxMockPtrframerreadCol.L.Lock()
	for recorderAuxMockPtrframerreadCol < i {
		condRecorderAuxMockPtrframerreadCol.Wait()
	}
	condRecorderAuxMockPtrframerreadCol.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadCol() {
	condRecorderAuxMockPtrframerreadCol.L.Lock()
	recorderAuxMockPtrframerreadCol++
	condRecorderAuxMockPtrframerreadCol.L.Unlock()
	condRecorderAuxMockPtrframerreadCol.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadCol() (ret int) {
	condRecorderAuxMockPtrframerreadCol.L.Lock()
	ret = recorderAuxMockPtrframerreadCol
	condRecorderAuxMockPtrframerreadCol.L.Unlock()
	return
}

// (recvf *framer)readCol - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readCol(argcol *ColumnInfo, argmeta *resultMetadata, argglobalSpec bool, argkeyspace string, argtable string) {
	FuncAuxMockPtrframerreadCol, ok := apomock.GetRegisteredFunc("gocql.framer.readCol")
	if ok {
		FuncAuxMockPtrframerreadCol.(func(recvf *framer, argcol *ColumnInfo, argmeta *resultMetadata, argglobalSpec bool, argkeyspace string, argtable string))(recvf, argcol, argmeta, argglobalSpec, argkeyspace, argtable)
	} else {
		panic("FuncAuxMockPtrframerreadCol ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadCol()
	return
}

//
// Mock: (recvf *framer)writeQueryParams(argopts *queryParams)()
//

type MockArgsTypeframerwriteQueryParams struct {
	ApomockCallNumber int
	Argopts           *queryParams
}

var LastMockArgsframerwriteQueryParams MockArgsTypeframerwriteQueryParams

// (recvf *framer)AuxMockwriteQueryParams(argopts *queryParams)() - Generated mock function
func (recvf *framer) AuxMockwriteQueryParams(argopts *queryParams) {
	LastMockArgsframerwriteQueryParams = MockArgsTypeframerwriteQueryParams{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteQueryParams(),
		Argopts:           argopts,
	}
	return
}

// RecorderAuxMockPtrframerwriteQueryParams  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteQueryParams int = 0

var condRecorderAuxMockPtrframerwriteQueryParams *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteQueryParams(i int) {
	condRecorderAuxMockPtrframerwriteQueryParams.L.Lock()
	for recorderAuxMockPtrframerwriteQueryParams < i {
		condRecorderAuxMockPtrframerwriteQueryParams.Wait()
	}
	condRecorderAuxMockPtrframerwriteQueryParams.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteQueryParams() {
	condRecorderAuxMockPtrframerwriteQueryParams.L.Lock()
	recorderAuxMockPtrframerwriteQueryParams++
	condRecorderAuxMockPtrframerwriteQueryParams.L.Unlock()
	condRecorderAuxMockPtrframerwriteQueryParams.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteQueryParams() (ret int) {
	condRecorderAuxMockPtrframerwriteQueryParams.L.Lock()
	ret = recorderAuxMockPtrframerwriteQueryParams
	condRecorderAuxMockPtrframerwriteQueryParams.L.Unlock()
	return
}

// (recvf *framer)writeQueryParams - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeQueryParams(argopts *queryParams) {
	FuncAuxMockPtrframerwriteQueryParams, ok := apomock.GetRegisteredFunc("gocql.framer.writeQueryParams")
	if ok {
		FuncAuxMockPtrframerwriteQueryParams.(func(recvf *framer, argopts *queryParams))(recvf, argopts)
	} else {
		panic("FuncAuxMockPtrframerwriteQueryParams ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteQueryParams()
	return
}

//
// Mock: (recvf *framer)readConsistency()(reta Consistency)
//

type MockArgsTypeframerreadConsistency struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadConsistency MockArgsTypeframerreadConsistency

// (recvf *framer)AuxMockreadConsistency()(reta Consistency) - Generated mock function
func (recvf *framer) AuxMockreadConsistency() (reta Consistency) {
	rargs, rerr := apomock.GetNext("gocql.framer.readConsistency")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readConsistency")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readConsistency")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(Consistency)
	}
	return
}

// RecorderAuxMockPtrframerreadConsistency  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadConsistency int = 0

var condRecorderAuxMockPtrframerreadConsistency *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadConsistency(i int) {
	condRecorderAuxMockPtrframerreadConsistency.L.Lock()
	for recorderAuxMockPtrframerreadConsistency < i {
		condRecorderAuxMockPtrframerreadConsistency.Wait()
	}
	condRecorderAuxMockPtrframerreadConsistency.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadConsistency() {
	condRecorderAuxMockPtrframerreadConsistency.L.Lock()
	recorderAuxMockPtrframerreadConsistency++
	condRecorderAuxMockPtrframerreadConsistency.L.Unlock()
	condRecorderAuxMockPtrframerreadConsistency.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadConsistency() (ret int) {
	condRecorderAuxMockPtrframerreadConsistency.L.Lock()
	ret = recorderAuxMockPtrframerreadConsistency
	condRecorderAuxMockPtrframerreadConsistency.L.Unlock()
	return
}

// (recvf *framer)readConsistency - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readConsistency() (reta Consistency) {
	FuncAuxMockPtrframerreadConsistency, ok := apomock.GetRegisteredFunc("gocql.framer.readConsistency")
	if ok {
		reta = FuncAuxMockPtrframerreadConsistency.(func(recvf *framer) (reta Consistency))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadConsistency ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadConsistency()
	return
}

//
// Mock: (recvf *framer)readLong()(retn int64)
//

type MockArgsTypeframerreadLong struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadLong MockArgsTypeframerreadLong

// (recvf *framer)AuxMockreadLong()(retn int64) - Generated mock function
func (recvf *framer) AuxMockreadLong() (retn int64) {
	rargs, rerr := apomock.GetNext("gocql.framer.readLong")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readLong")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readLong")
	}
	if rargs.GetArg(0) != nil {
		retn = rargs.GetArg(0).(int64)
	}
	return
}

// RecorderAuxMockPtrframerreadLong  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadLong int = 0

var condRecorderAuxMockPtrframerreadLong *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadLong(i int) {
	condRecorderAuxMockPtrframerreadLong.L.Lock()
	for recorderAuxMockPtrframerreadLong < i {
		condRecorderAuxMockPtrframerreadLong.Wait()
	}
	condRecorderAuxMockPtrframerreadLong.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadLong() {
	condRecorderAuxMockPtrframerreadLong.L.Lock()
	recorderAuxMockPtrframerreadLong++
	condRecorderAuxMockPtrframerreadLong.L.Unlock()
	condRecorderAuxMockPtrframerreadLong.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadLong() (ret int) {
	condRecorderAuxMockPtrframerreadLong.L.Lock()
	ret = recorderAuxMockPtrframerreadLong
	condRecorderAuxMockPtrframerreadLong.L.Unlock()
	return
}

// (recvf *framer)readLong - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readLong() (retn int64) {
	FuncAuxMockPtrframerreadLong, ok := apomock.GetRegisteredFunc("gocql.framer.readLong")
	if ok {
		retn = FuncAuxMockPtrframerreadLong.(func(recvf *framer) (retn int64))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadLong ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadLong()
	return
}

//
// Mock: (recvf *framer)readLongString()(rets string)
//

type MockArgsTypeframerreadLongString struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadLongString MockArgsTypeframerreadLongString

// (recvf *framer)AuxMockreadLongString()(rets string) - Generated mock function
func (recvf *framer) AuxMockreadLongString() (rets string) {
	rargs, rerr := apomock.GetNext("gocql.framer.readLongString")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readLongString")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readLongString")
	}
	if rargs.GetArg(0) != nil {
		rets = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrframerreadLongString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadLongString int = 0

var condRecorderAuxMockPtrframerreadLongString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadLongString(i int) {
	condRecorderAuxMockPtrframerreadLongString.L.Lock()
	for recorderAuxMockPtrframerreadLongString < i {
		condRecorderAuxMockPtrframerreadLongString.Wait()
	}
	condRecorderAuxMockPtrframerreadLongString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadLongString() {
	condRecorderAuxMockPtrframerreadLongString.L.Lock()
	recorderAuxMockPtrframerreadLongString++
	condRecorderAuxMockPtrframerreadLongString.L.Unlock()
	condRecorderAuxMockPtrframerreadLongString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadLongString() (ret int) {
	condRecorderAuxMockPtrframerreadLongString.L.Lock()
	ret = recorderAuxMockPtrframerreadLongString
	condRecorderAuxMockPtrframerreadLongString.L.Unlock()
	return
}

// (recvf *framer)readLongString - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readLongString() (rets string) {
	FuncAuxMockPtrframerreadLongString, ok := apomock.GetRegisteredFunc("gocql.framer.readLongString")
	if ok {
		rets = FuncAuxMockPtrframerreadLongString.(func(recvf *framer) (rets string))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadLongString ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadLongString()
	return
}

//
// Mock: (recvp protoVersion)version()(reta byte)
//

type MockArgsTypeprotoVersionversion struct {
	ApomockCallNumber int
}

var LastMockArgsprotoVersionversion MockArgsTypeprotoVersionversion

// (recvp protoVersion)AuxMockversion()(reta byte) - Generated mock function
func (recvp protoVersion) AuxMockversion() (reta byte) {
	rargs, rerr := apomock.GetNext("gocql.protoVersion.version")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.protoVersion.version")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.protoVersion.version")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(byte)
	}
	return
}

// RecorderAuxMockprotoVersionversion  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockprotoVersionversion int = 0

var condRecorderAuxMockprotoVersionversion *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockprotoVersionversion(i int) {
	condRecorderAuxMockprotoVersionversion.L.Lock()
	for recorderAuxMockprotoVersionversion < i {
		condRecorderAuxMockprotoVersionversion.Wait()
	}
	condRecorderAuxMockprotoVersionversion.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockprotoVersionversion() {
	condRecorderAuxMockprotoVersionversion.L.Lock()
	recorderAuxMockprotoVersionversion++
	condRecorderAuxMockprotoVersionversion.L.Unlock()
	condRecorderAuxMockprotoVersionversion.Broadcast()
}
func AuxMockGetRecorderAuxMockprotoVersionversion() (ret int) {
	condRecorderAuxMockprotoVersionversion.L.Lock()
	ret = recorderAuxMockprotoVersionversion
	condRecorderAuxMockprotoVersionversion.L.Unlock()
	return
}

// (recvp protoVersion)version - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp protoVersion) version() (reta byte) {
	FuncAuxMockprotoVersionversion, ok := apomock.GetRegisteredFunc("gocql.protoVersion.version")
	if ok {
		reta = FuncAuxMockprotoVersionversion.(func(recvp protoVersion) (reta byte))(recvp)
	} else {
		panic("FuncAuxMockprotoVersionversion ")
	}
	AuxMockIncrementRecorderAuxMockprotoVersionversion()
	return
}

//
// Mock: (recvf *framer)parseResultRows()(reta frame)
//

type MockArgsTypeframerparseResultRows struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseResultRows MockArgsTypeframerparseResultRows

// (recvf *framer)AuxMockparseResultRows()(reta frame) - Generated mock function
func (recvf *framer) AuxMockparseResultRows() (reta frame) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseResultRows")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseResultRows")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseResultRows")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	return
}

// RecorderAuxMockPtrframerparseResultRows  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseResultRows int = 0

var condRecorderAuxMockPtrframerparseResultRows *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseResultRows(i int) {
	condRecorderAuxMockPtrframerparseResultRows.L.Lock()
	for recorderAuxMockPtrframerparseResultRows < i {
		condRecorderAuxMockPtrframerparseResultRows.Wait()
	}
	condRecorderAuxMockPtrframerparseResultRows.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseResultRows() {
	condRecorderAuxMockPtrframerparseResultRows.L.Lock()
	recorderAuxMockPtrframerparseResultRows++
	condRecorderAuxMockPtrframerparseResultRows.L.Unlock()
	condRecorderAuxMockPtrframerparseResultRows.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseResultRows() (ret int) {
	condRecorderAuxMockPtrframerparseResultRows.L.Lock()
	ret = recorderAuxMockPtrframerparseResultRows
	condRecorderAuxMockPtrframerparseResultRows.L.Unlock()
	return
}

// (recvf *framer)parseResultRows - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseResultRows() (reta frame) {
	FuncAuxMockPtrframerparseResultRows, ok := apomock.GetRegisteredFunc("gocql.framer.parseResultRows")
	if ok {
		reta = FuncAuxMockPtrframerparseResultRows.(func(recvf *framer) (reta frame))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseResultRows ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseResultRows()
	return
}

//
// Mock: (recvf *framer)parseAuthChallengeFrame()(reta frame)
//

type MockArgsTypeframerparseAuthChallengeFrame struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseAuthChallengeFrame MockArgsTypeframerparseAuthChallengeFrame

// (recvf *framer)AuxMockparseAuthChallengeFrame()(reta frame) - Generated mock function
func (recvf *framer) AuxMockparseAuthChallengeFrame() (reta frame) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseAuthChallengeFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseAuthChallengeFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseAuthChallengeFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	return
}

// RecorderAuxMockPtrframerparseAuthChallengeFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseAuthChallengeFrame int = 0

var condRecorderAuxMockPtrframerparseAuthChallengeFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseAuthChallengeFrame(i int) {
	condRecorderAuxMockPtrframerparseAuthChallengeFrame.L.Lock()
	for recorderAuxMockPtrframerparseAuthChallengeFrame < i {
		condRecorderAuxMockPtrframerparseAuthChallengeFrame.Wait()
	}
	condRecorderAuxMockPtrframerparseAuthChallengeFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseAuthChallengeFrame() {
	condRecorderAuxMockPtrframerparseAuthChallengeFrame.L.Lock()
	recorderAuxMockPtrframerparseAuthChallengeFrame++
	condRecorderAuxMockPtrframerparseAuthChallengeFrame.L.Unlock()
	condRecorderAuxMockPtrframerparseAuthChallengeFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseAuthChallengeFrame() (ret int) {
	condRecorderAuxMockPtrframerparseAuthChallengeFrame.L.Lock()
	ret = recorderAuxMockPtrframerparseAuthChallengeFrame
	condRecorderAuxMockPtrframerparseAuthChallengeFrame.L.Unlock()
	return
}

// (recvf *framer)parseAuthChallengeFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseAuthChallengeFrame() (reta frame) {
	FuncAuxMockPtrframerparseAuthChallengeFrame, ok := apomock.GetRegisteredFunc("gocql.framer.parseAuthChallengeFrame")
	if ok {
		reta = FuncAuxMockPtrframerparseAuthChallengeFrame.(func(recvf *framer) (reta frame))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseAuthChallengeFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseAuthChallengeFrame()
	return
}

//
// Mock: (recvf *framer)writeQueryFrame(argstreamID int, argstatement string, argparams *queryParams)(reta error)
//

type MockArgsTypeframerwriteQueryFrame struct {
	ApomockCallNumber int
	ArgstreamID       int
	Argstatement      string
	Argparams         *queryParams
}

var LastMockArgsframerwriteQueryFrame MockArgsTypeframerwriteQueryFrame

// (recvf *framer)AuxMockwriteQueryFrame(argstreamID int, argstatement string, argparams *queryParams)(reta error) - Generated mock function
func (recvf *framer) AuxMockwriteQueryFrame(argstreamID int, argstatement string, argparams *queryParams) (reta error) {
	LastMockArgsframerwriteQueryFrame = MockArgsTypeframerwriteQueryFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteQueryFrame(),
		ArgstreamID:       argstreamID,
		Argstatement:      argstatement,
		Argparams:         argparams,
	}
	rargs, rerr := apomock.GetNext("gocql.framer.writeQueryFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.writeQueryFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.writeQueryFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrframerwriteQueryFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteQueryFrame int = 0

var condRecorderAuxMockPtrframerwriteQueryFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteQueryFrame(i int) {
	condRecorderAuxMockPtrframerwriteQueryFrame.L.Lock()
	for recorderAuxMockPtrframerwriteQueryFrame < i {
		condRecorderAuxMockPtrframerwriteQueryFrame.Wait()
	}
	condRecorderAuxMockPtrframerwriteQueryFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteQueryFrame() {
	condRecorderAuxMockPtrframerwriteQueryFrame.L.Lock()
	recorderAuxMockPtrframerwriteQueryFrame++
	condRecorderAuxMockPtrframerwriteQueryFrame.L.Unlock()
	condRecorderAuxMockPtrframerwriteQueryFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteQueryFrame() (ret int) {
	condRecorderAuxMockPtrframerwriteQueryFrame.L.Lock()
	ret = recorderAuxMockPtrframerwriteQueryFrame
	condRecorderAuxMockPtrframerwriteQueryFrame.L.Unlock()
	return
}

// (recvf *framer)writeQueryFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeQueryFrame(argstreamID int, argstatement string, argparams *queryParams) (reta error) {
	FuncAuxMockPtrframerwriteQueryFrame, ok := apomock.GetRegisteredFunc("gocql.framer.writeQueryFrame")
	if ok {
		reta = FuncAuxMockPtrframerwriteQueryFrame.(func(recvf *framer, argstreamID int, argstatement string, argparams *queryParams) (reta error))(recvf, argstreamID, argstatement, argparams)
	} else {
		panic("FuncAuxMockPtrframerwriteQueryFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteQueryFrame()
	return
}

//
// Mock: (recvf *framer)writeInt(argn int32)()
//

type MockArgsTypeframerwriteInt struct {
	ApomockCallNumber int
	Argn              int32
}

var LastMockArgsframerwriteInt MockArgsTypeframerwriteInt

// (recvf *framer)AuxMockwriteInt(argn int32)() - Generated mock function
func (recvf *framer) AuxMockwriteInt(argn int32) {
	LastMockArgsframerwriteInt = MockArgsTypeframerwriteInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteInt(),
		Argn:              argn,
	}
	return
}

// RecorderAuxMockPtrframerwriteInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteInt int = 0

var condRecorderAuxMockPtrframerwriteInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteInt(i int) {
	condRecorderAuxMockPtrframerwriteInt.L.Lock()
	for recorderAuxMockPtrframerwriteInt < i {
		condRecorderAuxMockPtrframerwriteInt.Wait()
	}
	condRecorderAuxMockPtrframerwriteInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteInt() {
	condRecorderAuxMockPtrframerwriteInt.L.Lock()
	recorderAuxMockPtrframerwriteInt++
	condRecorderAuxMockPtrframerwriteInt.L.Unlock()
	condRecorderAuxMockPtrframerwriteInt.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteInt() (ret int) {
	condRecorderAuxMockPtrframerwriteInt.L.Lock()
	ret = recorderAuxMockPtrframerwriteInt
	condRecorderAuxMockPtrframerwriteInt.L.Unlock()
	return
}

// (recvf *framer)writeInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeInt(argn int32) {
	FuncAuxMockPtrframerwriteInt, ok := apomock.GetRegisteredFunc("gocql.framer.writeInt")
	if ok {
		FuncAuxMockPtrframerwriteInt.(func(recvf *framer, argn int32))(recvf, argn)
	} else {
		panic("FuncAuxMockPtrframerwriteInt ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteInt()
	return
}

//
// Mock: (recvf *framer)writeShortBytes(argp []byte)()
//

type MockArgsTypeframerwriteShortBytes struct {
	ApomockCallNumber int
	Argp              []byte
}

var LastMockArgsframerwriteShortBytes MockArgsTypeframerwriteShortBytes

// (recvf *framer)AuxMockwriteShortBytes(argp []byte)() - Generated mock function
func (recvf *framer) AuxMockwriteShortBytes(argp []byte) {
	LastMockArgsframerwriteShortBytes = MockArgsTypeframerwriteShortBytes{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteShortBytes(),
		Argp:              argp,
	}
	return
}

// RecorderAuxMockPtrframerwriteShortBytes  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteShortBytes int = 0

var condRecorderAuxMockPtrframerwriteShortBytes *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteShortBytes(i int) {
	condRecorderAuxMockPtrframerwriteShortBytes.L.Lock()
	for recorderAuxMockPtrframerwriteShortBytes < i {
		condRecorderAuxMockPtrframerwriteShortBytes.Wait()
	}
	condRecorderAuxMockPtrframerwriteShortBytes.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteShortBytes() {
	condRecorderAuxMockPtrframerwriteShortBytes.L.Lock()
	recorderAuxMockPtrframerwriteShortBytes++
	condRecorderAuxMockPtrframerwriteShortBytes.L.Unlock()
	condRecorderAuxMockPtrframerwriteShortBytes.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteShortBytes() (ret int) {
	condRecorderAuxMockPtrframerwriteShortBytes.L.Lock()
	ret = recorderAuxMockPtrframerwriteShortBytes
	condRecorderAuxMockPtrframerwriteShortBytes.L.Unlock()
	return
}

// (recvf *framer)writeShortBytes - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeShortBytes(argp []byte) {
	FuncAuxMockPtrframerwriteShortBytes, ok := apomock.GetRegisteredFunc("gocql.framer.writeShortBytes")
	if ok {
		FuncAuxMockPtrframerwriteShortBytes.(func(recvf *framer, argp []byte))(recvf, argp)
	} else {
		panic("FuncAuxMockPtrframerwriteShortBytes ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteShortBytes()
	return
}

//
// Mock: (recvf frameHeader)Header()(reta frameHeader)
//

type MockArgsTypeframeHeaderHeader struct {
	ApomockCallNumber int
}

var LastMockArgsframeHeaderHeader MockArgsTypeframeHeaderHeader

// (recvf frameHeader)AuxMockHeader()(reta frameHeader) - Generated mock function
func (recvf frameHeader) AuxMockHeader() (reta frameHeader) {
	rargs, rerr := apomock.GetNext("gocql.frameHeader.Header")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.frameHeader.Header")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.frameHeader.Header")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frameHeader)
	}
	return
}

// RecorderAuxMockframeHeaderHeader  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockframeHeaderHeader int = 0

var condRecorderAuxMockframeHeaderHeader *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockframeHeaderHeader(i int) {
	condRecorderAuxMockframeHeaderHeader.L.Lock()
	for recorderAuxMockframeHeaderHeader < i {
		condRecorderAuxMockframeHeaderHeader.Wait()
	}
	condRecorderAuxMockframeHeaderHeader.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockframeHeaderHeader() {
	condRecorderAuxMockframeHeaderHeader.L.Lock()
	recorderAuxMockframeHeaderHeader++
	condRecorderAuxMockframeHeaderHeader.L.Unlock()
	condRecorderAuxMockframeHeaderHeader.Broadcast()
}
func AuxMockGetRecorderAuxMockframeHeaderHeader() (ret int) {
	condRecorderAuxMockframeHeaderHeader.L.Lock()
	ret = recorderAuxMockframeHeaderHeader
	condRecorderAuxMockframeHeaderHeader.L.Unlock()
	return
}

// (recvf frameHeader)Header - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf frameHeader) Header() (reta frameHeader) {
	FuncAuxMockframeHeaderHeader, ok := apomock.GetRegisteredFunc("gocql.frameHeader.Header")
	if ok {
		reta = FuncAuxMockframeHeaderHeader.(func(recvf frameHeader) (reta frameHeader))(recvf)
	} else {
		panic("FuncAuxMockframeHeaderHeader ")
	}
	AuxMockIncrementRecorderAuxMockframeHeaderHeader()
	return
}

//
// Mock: (recvw *writeQueryFrame)writeFrame(argframer *framer, argstreamID int)(reta error)
//

type MockArgsTypewriteQueryFramewriteFrame struct {
	ApomockCallNumber int
	Argframer         *framer
	ArgstreamID       int
}

var LastMockArgswriteQueryFramewriteFrame MockArgsTypewriteQueryFramewriteFrame

// (recvw *writeQueryFrame)AuxMockwriteFrame(argframer *framer, argstreamID int)(reta error) - Generated mock function
func (recvw *writeQueryFrame) AuxMockwriteFrame(argframer *framer, argstreamID int) (reta error) {
	LastMockArgswriteQueryFramewriteFrame = MockArgsTypewriteQueryFramewriteFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrwriteQueryFramewriteFrame(),
		Argframer:         argframer,
		ArgstreamID:       argstreamID,
	}
	rargs, rerr := apomock.GetNext("gocql.writeQueryFrame.writeFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeQueryFrame.writeFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeQueryFrame.writeFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrwriteQueryFramewriteFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrwriteQueryFramewriteFrame int = 0

var condRecorderAuxMockPtrwriteQueryFramewriteFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrwriteQueryFramewriteFrame(i int) {
	condRecorderAuxMockPtrwriteQueryFramewriteFrame.L.Lock()
	for recorderAuxMockPtrwriteQueryFramewriteFrame < i {
		condRecorderAuxMockPtrwriteQueryFramewriteFrame.Wait()
	}
	condRecorderAuxMockPtrwriteQueryFramewriteFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrwriteQueryFramewriteFrame() {
	condRecorderAuxMockPtrwriteQueryFramewriteFrame.L.Lock()
	recorderAuxMockPtrwriteQueryFramewriteFrame++
	condRecorderAuxMockPtrwriteQueryFramewriteFrame.L.Unlock()
	condRecorderAuxMockPtrwriteQueryFramewriteFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrwriteQueryFramewriteFrame() (ret int) {
	condRecorderAuxMockPtrwriteQueryFramewriteFrame.L.Lock()
	ret = recorderAuxMockPtrwriteQueryFramewriteFrame
	condRecorderAuxMockPtrwriteQueryFramewriteFrame.L.Unlock()
	return
}

// (recvw *writeQueryFrame)writeFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvw *writeQueryFrame) writeFrame(argframer *framer, argstreamID int) (reta error) {
	FuncAuxMockPtrwriteQueryFramewriteFrame, ok := apomock.GetRegisteredFunc("gocql.writeQueryFrame.writeFrame")
	if ok {
		reta = FuncAuxMockPtrwriteQueryFramewriteFrame.(func(recvw *writeQueryFrame, argframer *framer, argstreamID int) (reta error))(recvw, argframer, argstreamID)
	} else {
		panic("FuncAuxMockPtrwriteQueryFramewriteFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrwriteQueryFramewriteFrame()
	return
}

//
// Mock: (recvf *framer)writeByte(argb byte)()
//

type MockArgsTypeframerwriteByte struct {
	ApomockCallNumber int
	Argb              byte
}

var LastMockArgsframerwriteByte MockArgsTypeframerwriteByte

// (recvf *framer)AuxMockwriteByte(argb byte)() - Generated mock function
func (recvf *framer) AuxMockwriteByte(argb byte) {
	LastMockArgsframerwriteByte = MockArgsTypeframerwriteByte{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteByte(),
		Argb:              argb,
	}
	return
}

// RecorderAuxMockPtrframerwriteByte  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteByte int = 0

var condRecorderAuxMockPtrframerwriteByte *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteByte(i int) {
	condRecorderAuxMockPtrframerwriteByte.L.Lock()
	for recorderAuxMockPtrframerwriteByte < i {
		condRecorderAuxMockPtrframerwriteByte.Wait()
	}
	condRecorderAuxMockPtrframerwriteByte.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteByte() {
	condRecorderAuxMockPtrframerwriteByte.L.Lock()
	recorderAuxMockPtrframerwriteByte++
	condRecorderAuxMockPtrframerwriteByte.L.Unlock()
	condRecorderAuxMockPtrframerwriteByte.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteByte() (ret int) {
	condRecorderAuxMockPtrframerwriteByte.L.Lock()
	ret = recorderAuxMockPtrframerwriteByte
	condRecorderAuxMockPtrframerwriteByte.L.Unlock()
	return
}

// (recvf *framer)writeByte - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeByte(argb byte) {
	FuncAuxMockPtrframerwriteByte, ok := apomock.GetRegisteredFunc("gocql.framer.writeByte")
	if ok {
		FuncAuxMockPtrframerwriteByte.(func(recvf *framer, argb byte))(recvf, argb)
	} else {
		panic("FuncAuxMockPtrframerwriteByte ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteByte()
	return
}

//
// Mock: (recvp protoVersion)String()(reta string)
//

type MockArgsTypeprotoVersionString struct {
	ApomockCallNumber int
}

var LastMockArgsprotoVersionString MockArgsTypeprotoVersionString

// (recvp protoVersion)AuxMockString()(reta string) - Generated mock function
func (recvp protoVersion) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.protoVersion.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.protoVersion.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.protoVersion.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockprotoVersionString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockprotoVersionString int = 0

var condRecorderAuxMockprotoVersionString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockprotoVersionString(i int) {
	condRecorderAuxMockprotoVersionString.L.Lock()
	for recorderAuxMockprotoVersionString < i {
		condRecorderAuxMockprotoVersionString.Wait()
	}
	condRecorderAuxMockprotoVersionString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockprotoVersionString() {
	condRecorderAuxMockprotoVersionString.L.Lock()
	recorderAuxMockprotoVersionString++
	condRecorderAuxMockprotoVersionString.L.Unlock()
	condRecorderAuxMockprotoVersionString.Broadcast()
}
func AuxMockGetRecorderAuxMockprotoVersionString() (ret int) {
	condRecorderAuxMockprotoVersionString.L.Lock()
	ret = recorderAuxMockprotoVersionString
	condRecorderAuxMockprotoVersionString.L.Unlock()
	return
}

// (recvp protoVersion)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp protoVersion) String() (reta string) {
	FuncAuxMockprotoVersionString, ok := apomock.GetRegisteredFunc("gocql.protoVersion.String")
	if ok {
		reta = FuncAuxMockprotoVersionString.(func(recvp protoVersion) (reta string))(recvp)
	} else {
		panic("FuncAuxMockprotoVersionString ")
	}
	AuxMockIncrementRecorderAuxMockprotoVersionString()
	return
}

//
// Mock: (recvf *framer)writeHeader(argflags byte, argop frameOp, argstream int)()
//

type MockArgsTypeframerwriteHeader struct {
	ApomockCallNumber int
	Argflags          byte
	Argop             frameOp
	Argstream         int
}

var LastMockArgsframerwriteHeader MockArgsTypeframerwriteHeader

// (recvf *framer)AuxMockwriteHeader(argflags byte, argop frameOp, argstream int)() - Generated mock function
func (recvf *framer) AuxMockwriteHeader(argflags byte, argop frameOp, argstream int) {
	LastMockArgsframerwriteHeader = MockArgsTypeframerwriteHeader{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteHeader(),
		Argflags:          argflags,
		Argop:             argop,
		Argstream:         argstream,
	}
	return
}

// RecorderAuxMockPtrframerwriteHeader  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteHeader int = 0

var condRecorderAuxMockPtrframerwriteHeader *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteHeader(i int) {
	condRecorderAuxMockPtrframerwriteHeader.L.Lock()
	for recorderAuxMockPtrframerwriteHeader < i {
		condRecorderAuxMockPtrframerwriteHeader.Wait()
	}
	condRecorderAuxMockPtrframerwriteHeader.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteHeader() {
	condRecorderAuxMockPtrframerwriteHeader.L.Lock()
	recorderAuxMockPtrframerwriteHeader++
	condRecorderAuxMockPtrframerwriteHeader.L.Unlock()
	condRecorderAuxMockPtrframerwriteHeader.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteHeader() (ret int) {
	condRecorderAuxMockPtrframerwriteHeader.L.Lock()
	ret = recorderAuxMockPtrframerwriteHeader
	condRecorderAuxMockPtrframerwriteHeader.L.Unlock()
	return
}

// (recvf *framer)writeHeader - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeHeader(argflags byte, argop frameOp, argstream int) {
	FuncAuxMockPtrframerwriteHeader, ok := apomock.GetRegisteredFunc("gocql.framer.writeHeader")
	if ok {
		FuncAuxMockPtrframerwriteHeader.(func(recvf *framer, argflags byte, argop frameOp, argstream int))(recvf, argflags, argop, argstream)
	} else {
		panic("FuncAuxMockPtrframerwriteHeader ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteHeader()
	return
}

//
// Mock: (recvf *framer)setLength(arglength int)()
//

type MockArgsTypeframersetLength struct {
	ApomockCallNumber int
	Arglength         int
}

var LastMockArgsframersetLength MockArgsTypeframersetLength

// (recvf *framer)AuxMocksetLength(arglength int)() - Generated mock function
func (recvf *framer) AuxMocksetLength(arglength int) {
	LastMockArgsframersetLength = MockArgsTypeframersetLength{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframersetLength(),
		Arglength:         arglength,
	}
	return
}

// RecorderAuxMockPtrframersetLength  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframersetLength int = 0

var condRecorderAuxMockPtrframersetLength *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframersetLength(i int) {
	condRecorderAuxMockPtrframersetLength.L.Lock()
	for recorderAuxMockPtrframersetLength < i {
		condRecorderAuxMockPtrframersetLength.Wait()
	}
	condRecorderAuxMockPtrframersetLength.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframersetLength() {
	condRecorderAuxMockPtrframersetLength.L.Lock()
	recorderAuxMockPtrframersetLength++
	condRecorderAuxMockPtrframersetLength.L.Unlock()
	condRecorderAuxMockPtrframersetLength.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframersetLength() (ret int) {
	condRecorderAuxMockPtrframersetLength.L.Lock()
	ret = recorderAuxMockPtrframersetLength
	condRecorderAuxMockPtrframersetLength.L.Unlock()
	return
}

// (recvf *framer)setLength - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) setLength(arglength int) {
	FuncAuxMockPtrframersetLength, ok := apomock.GetRegisteredFunc("gocql.framer.setLength")
	if ok {
		FuncAuxMockPtrframersetLength.(func(recvf *framer, arglength int))(recvf, arglength)
	} else {
		panic("FuncAuxMockPtrframersetLength ")
	}
	AuxMockIncrementRecorderAuxMockPtrframersetLength()
	return
}

//
// Mock: (recvf *framer)parseAuthenticateFrame()(reta frame)
//

type MockArgsTypeframerparseAuthenticateFrame struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseAuthenticateFrame MockArgsTypeframerparseAuthenticateFrame

// (recvf *framer)AuxMockparseAuthenticateFrame()(reta frame) - Generated mock function
func (recvf *framer) AuxMockparseAuthenticateFrame() (reta frame) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseAuthenticateFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseAuthenticateFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseAuthenticateFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	return
}

// RecorderAuxMockPtrframerparseAuthenticateFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseAuthenticateFrame int = 0

var condRecorderAuxMockPtrframerparseAuthenticateFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseAuthenticateFrame(i int) {
	condRecorderAuxMockPtrframerparseAuthenticateFrame.L.Lock()
	for recorderAuxMockPtrframerparseAuthenticateFrame < i {
		condRecorderAuxMockPtrframerparseAuthenticateFrame.Wait()
	}
	condRecorderAuxMockPtrframerparseAuthenticateFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseAuthenticateFrame() {
	condRecorderAuxMockPtrframerparseAuthenticateFrame.L.Lock()
	recorderAuxMockPtrframerparseAuthenticateFrame++
	condRecorderAuxMockPtrframerparseAuthenticateFrame.L.Unlock()
	condRecorderAuxMockPtrframerparseAuthenticateFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseAuthenticateFrame() (ret int) {
	condRecorderAuxMockPtrframerparseAuthenticateFrame.L.Lock()
	ret = recorderAuxMockPtrframerparseAuthenticateFrame
	condRecorderAuxMockPtrframerparseAuthenticateFrame.L.Unlock()
	return
}

// (recvf *framer)parseAuthenticateFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseAuthenticateFrame() (reta frame) {
	FuncAuxMockPtrframerparseAuthenticateFrame, ok := apomock.GetRegisteredFunc("gocql.framer.parseAuthenticateFrame")
	if ok {
		reta = FuncAuxMockPtrframerparseAuthenticateFrame.(func(recvf *framer) (reta frame))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseAuthenticateFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseAuthenticateFrame()
	return
}

//
// Mock: appendLong(argp []byte, argn int64)(reta []byte)
//

type MockArgsTypeappendLong struct {
	ApomockCallNumber int
	Argp              []byte
	Argn              int64
}

var LastMockArgsappendLong MockArgsTypeappendLong

// AuxMockappendLong(argp []byte, argn int64)(reta []byte) - Generated mock function
func AuxMockappendLong(argp []byte, argn int64) (reta []byte) {
	LastMockArgsappendLong = MockArgsTypeappendLong{
		ApomockCallNumber: AuxMockGetRecorderAuxMockappendLong(),
		Argp:              argp,
		Argn:              argn,
	}
	rargs, rerr := apomock.GetNext("gocql.appendLong")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.appendLong")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.appendLong")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockappendLong  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockappendLong int = 0

var condRecorderAuxMockappendLong *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockappendLong(i int) {
	condRecorderAuxMockappendLong.L.Lock()
	for recorderAuxMockappendLong < i {
		condRecorderAuxMockappendLong.Wait()
	}
	condRecorderAuxMockappendLong.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockappendLong() {
	condRecorderAuxMockappendLong.L.Lock()
	recorderAuxMockappendLong++
	condRecorderAuxMockappendLong.L.Unlock()
	condRecorderAuxMockappendLong.Broadcast()
}
func AuxMockGetRecorderAuxMockappendLong() (ret int) {
	condRecorderAuxMockappendLong.L.Lock()
	ret = recorderAuxMockappendLong
	condRecorderAuxMockappendLong.L.Unlock()
	return
}

// appendLong - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func appendLong(argp []byte, argn int64) (reta []byte) {
	FuncAuxMockappendLong, ok := apomock.GetRegisteredFunc("gocql.appendLong")
	if ok {
		reta = FuncAuxMockappendLong.(func(argp []byte, argn int64) (reta []byte))(argp, argn)
	} else {
		panic("FuncAuxMockappendLong ")
	}
	AuxMockIncrementRecorderAuxMockappendLong()
	return
}

//
// Mock: readHeader(argr io.Reader, argp []byte)(rethead frameHeader, reterr error)
//

type MockArgsTypereadHeader struct {
	ApomockCallNumber int
	Argr              io.Reader
	Argp              []byte
}

var LastMockArgsreadHeader MockArgsTypereadHeader

// AuxMockreadHeader(argr io.Reader, argp []byte)(rethead frameHeader, reterr error) - Generated mock function
func AuxMockreadHeader(argr io.Reader, argp []byte) (rethead frameHeader, reterr error) {
	LastMockArgsreadHeader = MockArgsTypereadHeader{
		ApomockCallNumber: AuxMockGetRecorderAuxMockreadHeader(),
		Argr:              argr,
		Argp:              argp,
	}
	rargs, rerr := apomock.GetNext("gocql.readHeader")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.readHeader")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.readHeader")
	}
	if rargs.GetArg(0) != nil {
		rethead = rargs.GetArg(0).(frameHeader)
	}
	if rargs.GetArg(1) != nil {
		reterr = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockreadHeader  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockreadHeader int = 0

var condRecorderAuxMockreadHeader *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockreadHeader(i int) {
	condRecorderAuxMockreadHeader.L.Lock()
	for recorderAuxMockreadHeader < i {
		condRecorderAuxMockreadHeader.Wait()
	}
	condRecorderAuxMockreadHeader.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockreadHeader() {
	condRecorderAuxMockreadHeader.L.Lock()
	recorderAuxMockreadHeader++
	condRecorderAuxMockreadHeader.L.Unlock()
	condRecorderAuxMockreadHeader.Broadcast()
}
func AuxMockGetRecorderAuxMockreadHeader() (ret int) {
	condRecorderAuxMockreadHeader.L.Lock()
	ret = recorderAuxMockreadHeader
	condRecorderAuxMockreadHeader.L.Unlock()
	return
}

// readHeader - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func readHeader(argr io.Reader, argp []byte) (rethead frameHeader, reterr error) {
	FuncAuxMockreadHeader, ok := apomock.GetRegisteredFunc("gocql.readHeader")
	if ok {
		rethead, reterr = FuncAuxMockreadHeader.(func(argr io.Reader, argp []byte) (rethead frameHeader, reterr error))(argr, argp)
	} else {
		panic("FuncAuxMockreadHeader ")
	}
	AuxMockIncrementRecorderAuxMockreadHeader()
	return
}

//
// Mock: (recvf *framer)writePrepareFrame(argstream int, argstatement string)(reta error)
//

type MockArgsTypeframerwritePrepareFrame struct {
	ApomockCallNumber int
	Argstream         int
	Argstatement      string
}

var LastMockArgsframerwritePrepareFrame MockArgsTypeframerwritePrepareFrame

// (recvf *framer)AuxMockwritePrepareFrame(argstream int, argstatement string)(reta error) - Generated mock function
func (recvf *framer) AuxMockwritePrepareFrame(argstream int, argstatement string) (reta error) {
	LastMockArgsframerwritePrepareFrame = MockArgsTypeframerwritePrepareFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwritePrepareFrame(),
		Argstream:         argstream,
		Argstatement:      argstatement,
	}
	rargs, rerr := apomock.GetNext("gocql.framer.writePrepareFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.writePrepareFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.writePrepareFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrframerwritePrepareFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwritePrepareFrame int = 0

var condRecorderAuxMockPtrframerwritePrepareFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwritePrepareFrame(i int) {
	condRecorderAuxMockPtrframerwritePrepareFrame.L.Lock()
	for recorderAuxMockPtrframerwritePrepareFrame < i {
		condRecorderAuxMockPtrframerwritePrepareFrame.Wait()
	}
	condRecorderAuxMockPtrframerwritePrepareFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwritePrepareFrame() {
	condRecorderAuxMockPtrframerwritePrepareFrame.L.Lock()
	recorderAuxMockPtrframerwritePrepareFrame++
	condRecorderAuxMockPtrframerwritePrepareFrame.L.Unlock()
	condRecorderAuxMockPtrframerwritePrepareFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwritePrepareFrame() (ret int) {
	condRecorderAuxMockPtrframerwritePrepareFrame.L.Lock()
	ret = recorderAuxMockPtrframerwritePrepareFrame
	condRecorderAuxMockPtrframerwritePrepareFrame.L.Unlock()
	return
}

// (recvf *framer)writePrepareFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writePrepareFrame(argstream int, argstatement string) (reta error) {
	FuncAuxMockPtrframerwritePrepareFrame, ok := apomock.GetRegisteredFunc("gocql.framer.writePrepareFrame")
	if ok {
		reta = FuncAuxMockPtrframerwritePrepareFrame.(func(recvf *framer, argstream int, argstatement string) (reta error))(recvf, argstream, argstatement)
	} else {
		panic("FuncAuxMockPtrframerwritePrepareFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwritePrepareFrame()
	return
}

//
// Mock: (recve *writeExecuteFrame)String()(reta string)
//

type MockArgsTypewriteExecuteFrameString struct {
	ApomockCallNumber int
}

var LastMockArgswriteExecuteFrameString MockArgsTypewriteExecuteFrameString

// (recve *writeExecuteFrame)AuxMockString()(reta string) - Generated mock function
func (recve *writeExecuteFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.writeExecuteFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeExecuteFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeExecuteFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrwriteExecuteFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrwriteExecuteFrameString int = 0

var condRecorderAuxMockPtrwriteExecuteFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrwriteExecuteFrameString(i int) {
	condRecorderAuxMockPtrwriteExecuteFrameString.L.Lock()
	for recorderAuxMockPtrwriteExecuteFrameString < i {
		condRecorderAuxMockPtrwriteExecuteFrameString.Wait()
	}
	condRecorderAuxMockPtrwriteExecuteFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrwriteExecuteFrameString() {
	condRecorderAuxMockPtrwriteExecuteFrameString.L.Lock()
	recorderAuxMockPtrwriteExecuteFrameString++
	condRecorderAuxMockPtrwriteExecuteFrameString.L.Unlock()
	condRecorderAuxMockPtrwriteExecuteFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrwriteExecuteFrameString() (ret int) {
	condRecorderAuxMockPtrwriteExecuteFrameString.L.Lock()
	ret = recorderAuxMockPtrwriteExecuteFrameString
	condRecorderAuxMockPtrwriteExecuteFrameString.L.Unlock()
	return
}

// (recve *writeExecuteFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve *writeExecuteFrame) String() (reta string) {
	FuncAuxMockPtrwriteExecuteFrameString, ok := apomock.GetRegisteredFunc("gocql.writeExecuteFrame.String")
	if ok {
		reta = FuncAuxMockPtrwriteExecuteFrameString.(func(recve *writeExecuteFrame) (reta string))(recve)
	} else {
		panic("FuncAuxMockPtrwriteExecuteFrameString ")
	}
	AuxMockIncrementRecorderAuxMockPtrwriteExecuteFrameString()
	return
}

//
// Mock: appendBytes(argp []byte, argd []byte)(reta []byte)
//

type MockArgsTypeappendBytes struct {
	ApomockCallNumber int
	Argp              []byte
	Argd              []byte
}

var LastMockArgsappendBytes MockArgsTypeappendBytes

// AuxMockappendBytes(argp []byte, argd []byte)(reta []byte) - Generated mock function
func AuxMockappendBytes(argp []byte, argd []byte) (reta []byte) {
	LastMockArgsappendBytes = MockArgsTypeappendBytes{
		ApomockCallNumber: AuxMockGetRecorderAuxMockappendBytes(),
		Argp:              argp,
		Argd:              argd,
	}
	rargs, rerr := apomock.GetNext("gocql.appendBytes")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.appendBytes")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.appendBytes")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockappendBytes  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockappendBytes int = 0

var condRecorderAuxMockappendBytes *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockappendBytes(i int) {
	condRecorderAuxMockappendBytes.L.Lock()
	for recorderAuxMockappendBytes < i {
		condRecorderAuxMockappendBytes.Wait()
	}
	condRecorderAuxMockappendBytes.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockappendBytes() {
	condRecorderAuxMockappendBytes.L.Lock()
	recorderAuxMockappendBytes++
	condRecorderAuxMockappendBytes.L.Unlock()
	condRecorderAuxMockappendBytes.Broadcast()
}
func AuxMockGetRecorderAuxMockappendBytes() (ret int) {
	condRecorderAuxMockappendBytes.L.Lock()
	ret = recorderAuxMockappendBytes
	condRecorderAuxMockappendBytes.L.Unlock()
	return
}

// appendBytes - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func appendBytes(argp []byte, argd []byte) (reta []byte) {
	FuncAuxMockappendBytes, ok := apomock.GetRegisteredFunc("gocql.appendBytes")
	if ok {
		reta = FuncAuxMockappendBytes.(func(argp []byte, argd []byte) (reta []byte))(argp, argd)
	} else {
		panic("FuncAuxMockappendBytes ")
	}
	AuxMockIncrementRecorderAuxMockappendBytes()
	return
}

//
// Mock: readInt(argp []byte)(reta int32)
//

type MockArgsTypereadInt struct {
	ApomockCallNumber int
	Argp              []byte
}

var LastMockArgsreadInt MockArgsTypereadInt

// AuxMockreadInt(argp []byte)(reta int32) - Generated mock function
func AuxMockreadInt(argp []byte) (reta int32) {
	LastMockArgsreadInt = MockArgsTypereadInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockreadInt(),
		Argp:              argp,
	}
	rargs, rerr := apomock.GetNext("gocql.readInt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.readInt")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.readInt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int32)
	}
	return
}

// RecorderAuxMockreadInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockreadInt int = 0

var condRecorderAuxMockreadInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockreadInt(i int) {
	condRecorderAuxMockreadInt.L.Lock()
	for recorderAuxMockreadInt < i {
		condRecorderAuxMockreadInt.Wait()
	}
	condRecorderAuxMockreadInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockreadInt() {
	condRecorderAuxMockreadInt.L.Lock()
	recorderAuxMockreadInt++
	condRecorderAuxMockreadInt.L.Unlock()
	condRecorderAuxMockreadInt.Broadcast()
}
func AuxMockGetRecorderAuxMockreadInt() (ret int) {
	condRecorderAuxMockreadInt.L.Lock()
	ret = recorderAuxMockreadInt
	condRecorderAuxMockreadInt.L.Unlock()
	return
}

// readInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func readInt(argp []byte) (reta int32) {
	FuncAuxMockreadInt, ok := apomock.GetRegisteredFunc("gocql.readInt")
	if ok {
		reta = FuncAuxMockreadInt.(func(argp []byte) (reta int32))(argp)
	} else {
		panic("FuncAuxMockreadInt ")
	}
	AuxMockIncrementRecorderAuxMockreadInt()
	return
}

//
// Mock: (recvf frameHeader)String()(reta string)
//

type MockArgsTypeframeHeaderString struct {
	ApomockCallNumber int
}

var LastMockArgsframeHeaderString MockArgsTypeframeHeaderString

// (recvf frameHeader)AuxMockString()(reta string) - Generated mock function
func (recvf frameHeader) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.frameHeader.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.frameHeader.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.frameHeader.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockframeHeaderString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockframeHeaderString int = 0

var condRecorderAuxMockframeHeaderString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockframeHeaderString(i int) {
	condRecorderAuxMockframeHeaderString.L.Lock()
	for recorderAuxMockframeHeaderString < i {
		condRecorderAuxMockframeHeaderString.Wait()
	}
	condRecorderAuxMockframeHeaderString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockframeHeaderString() {
	condRecorderAuxMockframeHeaderString.L.Lock()
	recorderAuxMockframeHeaderString++
	condRecorderAuxMockframeHeaderString.L.Unlock()
	condRecorderAuxMockframeHeaderString.Broadcast()
}
func AuxMockGetRecorderAuxMockframeHeaderString() (ret int) {
	condRecorderAuxMockframeHeaderString.L.Lock()
	ret = recorderAuxMockframeHeaderString
	condRecorderAuxMockframeHeaderString.L.Unlock()
	return
}

// (recvf frameHeader)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf frameHeader) String() (reta string) {
	FuncAuxMockframeHeaderString, ok := apomock.GetRegisteredFunc("gocql.frameHeader.String")
	if ok {
		reta = FuncAuxMockframeHeaderString.(func(recvf frameHeader) (reta string))(recvf)
	} else {
		panic("FuncAuxMockframeHeaderString ")
	}
	AuxMockIncrementRecorderAuxMockframeHeaderString()
	return
}

//
// Mock: (recvf *framer)readFrame(arghead *frameHeader)(reta error)
//

type MockArgsTypeframerreadFrame struct {
	ApomockCallNumber int
	Arghead           *frameHeader
}

var LastMockArgsframerreadFrame MockArgsTypeframerreadFrame

// (recvf *framer)AuxMockreadFrame(arghead *frameHeader)(reta error) - Generated mock function
func (recvf *framer) AuxMockreadFrame(arghead *frameHeader) (reta error) {
	LastMockArgsframerreadFrame = MockArgsTypeframerreadFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerreadFrame(),
		Arghead:           arghead,
	}
	rargs, rerr := apomock.GetNext("gocql.framer.readFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrframerreadFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadFrame int = 0

var condRecorderAuxMockPtrframerreadFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadFrame(i int) {
	condRecorderAuxMockPtrframerreadFrame.L.Lock()
	for recorderAuxMockPtrframerreadFrame < i {
		condRecorderAuxMockPtrframerreadFrame.Wait()
	}
	condRecorderAuxMockPtrframerreadFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadFrame() {
	condRecorderAuxMockPtrframerreadFrame.L.Lock()
	recorderAuxMockPtrframerreadFrame++
	condRecorderAuxMockPtrframerreadFrame.L.Unlock()
	condRecorderAuxMockPtrframerreadFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadFrame() (ret int) {
	condRecorderAuxMockPtrframerreadFrame.L.Lock()
	ret = recorderAuxMockPtrframerreadFrame
	condRecorderAuxMockPtrframerreadFrame.L.Unlock()
	return
}

// (recvf *framer)readFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readFrame(arghead *frameHeader) (reta error) {
	FuncAuxMockPtrframerreadFrame, ok := apomock.GetRegisteredFunc("gocql.framer.readFrame")
	if ok {
		reta = FuncAuxMockPtrframerreadFrame.(func(recvf *framer, arghead *frameHeader) (reta error))(recvf, arghead)
	} else {
		panic("FuncAuxMockPtrframerreadFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadFrame()
	return
}

//
// Mock: (recvf *resultVoidFrame)String()(reta string)
//

type MockArgsTyperesultVoidFrameString struct {
	ApomockCallNumber int
}

var LastMockArgsresultVoidFrameString MockArgsTyperesultVoidFrameString

// (recvf *resultVoidFrame)AuxMockString()(reta string) - Generated mock function
func (recvf *resultVoidFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.resultVoidFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.resultVoidFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.resultVoidFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrresultVoidFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrresultVoidFrameString int = 0

var condRecorderAuxMockPtrresultVoidFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrresultVoidFrameString(i int) {
	condRecorderAuxMockPtrresultVoidFrameString.L.Lock()
	for recorderAuxMockPtrresultVoidFrameString < i {
		condRecorderAuxMockPtrresultVoidFrameString.Wait()
	}
	condRecorderAuxMockPtrresultVoidFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrresultVoidFrameString() {
	condRecorderAuxMockPtrresultVoidFrameString.L.Lock()
	recorderAuxMockPtrresultVoidFrameString++
	condRecorderAuxMockPtrresultVoidFrameString.L.Unlock()
	condRecorderAuxMockPtrresultVoidFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrresultVoidFrameString() (ret int) {
	condRecorderAuxMockPtrresultVoidFrameString.L.Lock()
	ret = recorderAuxMockPtrresultVoidFrameString
	condRecorderAuxMockPtrresultVoidFrameString.L.Unlock()
	return
}

// (recvf *resultVoidFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *resultVoidFrame) String() (reta string) {
	FuncAuxMockPtrresultVoidFrameString, ok := apomock.GetRegisteredFunc("gocql.resultVoidFrame.String")
	if ok {
		reta = FuncAuxMockPtrresultVoidFrameString.(func(recvf *resultVoidFrame) (reta string))(recvf)
	} else {
		panic("FuncAuxMockPtrresultVoidFrameString ")
	}
	AuxMockIncrementRecorderAuxMockPtrresultVoidFrameString()
	return
}

//
// Mock: (recvr *resultKeyspaceFrame)String()(reta string)
//

type MockArgsTyperesultKeyspaceFrameString struct {
	ApomockCallNumber int
}

var LastMockArgsresultKeyspaceFrameString MockArgsTyperesultKeyspaceFrameString

// (recvr *resultKeyspaceFrame)AuxMockString()(reta string) - Generated mock function
func (recvr *resultKeyspaceFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.resultKeyspaceFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.resultKeyspaceFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.resultKeyspaceFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrresultKeyspaceFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrresultKeyspaceFrameString int = 0

var condRecorderAuxMockPtrresultKeyspaceFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrresultKeyspaceFrameString(i int) {
	condRecorderAuxMockPtrresultKeyspaceFrameString.L.Lock()
	for recorderAuxMockPtrresultKeyspaceFrameString < i {
		condRecorderAuxMockPtrresultKeyspaceFrameString.Wait()
	}
	condRecorderAuxMockPtrresultKeyspaceFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrresultKeyspaceFrameString() {
	condRecorderAuxMockPtrresultKeyspaceFrameString.L.Lock()
	recorderAuxMockPtrresultKeyspaceFrameString++
	condRecorderAuxMockPtrresultKeyspaceFrameString.L.Unlock()
	condRecorderAuxMockPtrresultKeyspaceFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrresultKeyspaceFrameString() (ret int) {
	condRecorderAuxMockPtrresultKeyspaceFrameString.L.Lock()
	ret = recorderAuxMockPtrresultKeyspaceFrameString
	condRecorderAuxMockPtrresultKeyspaceFrameString.L.Unlock()
	return
}

// (recvr *resultKeyspaceFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr *resultKeyspaceFrame) String() (reta string) {
	FuncAuxMockPtrresultKeyspaceFrameString, ok := apomock.GetRegisteredFunc("gocql.resultKeyspaceFrame.String")
	if ok {
		reta = FuncAuxMockPtrresultKeyspaceFrameString.(func(recvr *resultKeyspaceFrame) (reta string))(recvr)
	} else {
		panic("FuncAuxMockPtrresultKeyspaceFrameString ")
	}
	AuxMockIncrementRecorderAuxMockPtrresultKeyspaceFrameString()
	return
}

//
// Mock: (recvt topologyChangeEventFrame)String()(reta string)
//

type MockArgsTypetopologyChangeEventFrameString struct {
	ApomockCallNumber int
}

var LastMockArgstopologyChangeEventFrameString MockArgsTypetopologyChangeEventFrameString

// (recvt topologyChangeEventFrame)AuxMockString()(reta string) - Generated mock function
func (recvt topologyChangeEventFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.topologyChangeEventFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.topologyChangeEventFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.topologyChangeEventFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMocktopologyChangeEventFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocktopologyChangeEventFrameString int = 0

var condRecorderAuxMocktopologyChangeEventFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocktopologyChangeEventFrameString(i int) {
	condRecorderAuxMocktopologyChangeEventFrameString.L.Lock()
	for recorderAuxMocktopologyChangeEventFrameString < i {
		condRecorderAuxMocktopologyChangeEventFrameString.Wait()
	}
	condRecorderAuxMocktopologyChangeEventFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocktopologyChangeEventFrameString() {
	condRecorderAuxMocktopologyChangeEventFrameString.L.Lock()
	recorderAuxMocktopologyChangeEventFrameString++
	condRecorderAuxMocktopologyChangeEventFrameString.L.Unlock()
	condRecorderAuxMocktopologyChangeEventFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMocktopologyChangeEventFrameString() (ret int) {
	condRecorderAuxMocktopologyChangeEventFrameString.L.Lock()
	ret = recorderAuxMocktopologyChangeEventFrameString
	condRecorderAuxMocktopologyChangeEventFrameString.L.Unlock()
	return
}

// (recvt topologyChangeEventFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt topologyChangeEventFrame) String() (reta string) {
	FuncAuxMocktopologyChangeEventFrameString, ok := apomock.GetRegisteredFunc("gocql.topologyChangeEventFrame.String")
	if ok {
		reta = FuncAuxMocktopologyChangeEventFrameString.(func(recvt topologyChangeEventFrame) (reta string))(recvt)
	} else {
		panic("FuncAuxMocktopologyChangeEventFrameString ")
	}
	AuxMockIncrementRecorderAuxMocktopologyChangeEventFrameString()
	return
}

//
// Mock: (recvf *framer)writeStringMap(argm map[string]string)()
//

type MockArgsTypeframerwriteStringMap struct {
	ApomockCallNumber int
	Argm              map[string]string
}

var LastMockArgsframerwriteStringMap MockArgsTypeframerwriteStringMap

// (recvf *framer)AuxMockwriteStringMap(argm map[string]string)() - Generated mock function
func (recvf *framer) AuxMockwriteStringMap(argm map[string]string) {
	LastMockArgsframerwriteStringMap = MockArgsTypeframerwriteStringMap{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteStringMap(),
		Argm:              argm,
	}
	return
}

// RecorderAuxMockPtrframerwriteStringMap  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteStringMap int = 0

var condRecorderAuxMockPtrframerwriteStringMap *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteStringMap(i int) {
	condRecorderAuxMockPtrframerwriteStringMap.L.Lock()
	for recorderAuxMockPtrframerwriteStringMap < i {
		condRecorderAuxMockPtrframerwriteStringMap.Wait()
	}
	condRecorderAuxMockPtrframerwriteStringMap.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteStringMap() {
	condRecorderAuxMockPtrframerwriteStringMap.L.Lock()
	recorderAuxMockPtrframerwriteStringMap++
	condRecorderAuxMockPtrframerwriteStringMap.L.Unlock()
	condRecorderAuxMockPtrframerwriteStringMap.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteStringMap() (ret int) {
	condRecorderAuxMockPtrframerwriteStringMap.L.Lock()
	ret = recorderAuxMockPtrframerwriteStringMap
	condRecorderAuxMockPtrframerwriteStringMap.L.Unlock()
	return
}

// (recvf *framer)writeStringMap - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeStringMap(argm map[string]string) {
	FuncAuxMockPtrframerwriteStringMap, ok := apomock.GetRegisteredFunc("gocql.framer.writeStringMap")
	if ok {
		FuncAuxMockPtrframerwriteStringMap.(func(recvf *framer, argm map[string]string))(recvf, argm)
	} else {
		panic("FuncAuxMockPtrframerwriteStringMap ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteStringMap()
	return
}

//
// Mock: (recvf *framer)parseResultMetadata()(reta resultMetadata)
//

type MockArgsTypeframerparseResultMetadata struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseResultMetadata MockArgsTypeframerparseResultMetadata

// (recvf *framer)AuxMockparseResultMetadata()(reta resultMetadata) - Generated mock function
func (recvf *framer) AuxMockparseResultMetadata() (reta resultMetadata) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseResultMetadata")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseResultMetadata")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseResultMetadata")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(resultMetadata)
	}
	return
}

// RecorderAuxMockPtrframerparseResultMetadata  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseResultMetadata int = 0

var condRecorderAuxMockPtrframerparseResultMetadata *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseResultMetadata(i int) {
	condRecorderAuxMockPtrframerparseResultMetadata.L.Lock()
	for recorderAuxMockPtrframerparseResultMetadata < i {
		condRecorderAuxMockPtrframerparseResultMetadata.Wait()
	}
	condRecorderAuxMockPtrframerparseResultMetadata.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseResultMetadata() {
	condRecorderAuxMockPtrframerparseResultMetadata.L.Lock()
	recorderAuxMockPtrframerparseResultMetadata++
	condRecorderAuxMockPtrframerparseResultMetadata.L.Unlock()
	condRecorderAuxMockPtrframerparseResultMetadata.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseResultMetadata() (ret int) {
	condRecorderAuxMockPtrframerparseResultMetadata.L.Lock()
	ret = recorderAuxMockPtrframerparseResultMetadata
	condRecorderAuxMockPtrframerparseResultMetadata.L.Unlock()
	return
}

// (recvf *framer)parseResultMetadata - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseResultMetadata() (reta resultMetadata) {
	FuncAuxMockPtrframerparseResultMetadata, ok := apomock.GetRegisteredFunc("gocql.framer.parseResultMetadata")
	if ok {
		reta = FuncAuxMockPtrframerparseResultMetadata.(func(recvf *framer) (reta resultMetadata))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseResultMetadata ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseResultMetadata()
	return
}

//
// Mock: (recvf schemaChangeTable)String()(reta string)
//

type MockArgsTypeschemaChangeTableString struct {
	ApomockCallNumber int
}

var LastMockArgsschemaChangeTableString MockArgsTypeschemaChangeTableString

// (recvf schemaChangeTable)AuxMockString()(reta string) - Generated mock function
func (recvf schemaChangeTable) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.schemaChangeTable.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.schemaChangeTable.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.schemaChangeTable.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockschemaChangeTableString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockschemaChangeTableString int = 0

var condRecorderAuxMockschemaChangeTableString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockschemaChangeTableString(i int) {
	condRecorderAuxMockschemaChangeTableString.L.Lock()
	for recorderAuxMockschemaChangeTableString < i {
		condRecorderAuxMockschemaChangeTableString.Wait()
	}
	condRecorderAuxMockschemaChangeTableString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockschemaChangeTableString() {
	condRecorderAuxMockschemaChangeTableString.L.Lock()
	recorderAuxMockschemaChangeTableString++
	condRecorderAuxMockschemaChangeTableString.L.Unlock()
	condRecorderAuxMockschemaChangeTableString.Broadcast()
}
func AuxMockGetRecorderAuxMockschemaChangeTableString() (ret int) {
	condRecorderAuxMockschemaChangeTableString.L.Lock()
	ret = recorderAuxMockschemaChangeTableString
	condRecorderAuxMockschemaChangeTableString.L.Unlock()
	return
}

// (recvf schemaChangeTable)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf schemaChangeTable) String() (reta string) {
	FuncAuxMockschemaChangeTableString, ok := apomock.GetRegisteredFunc("gocql.schemaChangeTable.String")
	if ok {
		reta = FuncAuxMockschemaChangeTableString.(func(recvf schemaChangeTable) (reta string))(recvf)
	} else {
		panic("FuncAuxMockschemaChangeTableString ")
	}
	AuxMockIncrementRecorderAuxMockschemaChangeTableString()
	return
}

//
// Mock: (recvf *framer)parseEventFrame()(reta frame)
//

type MockArgsTypeframerparseEventFrame struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseEventFrame MockArgsTypeframerparseEventFrame

// (recvf *framer)AuxMockparseEventFrame()(reta frame) - Generated mock function
func (recvf *framer) AuxMockparseEventFrame() (reta frame) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseEventFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseEventFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseEventFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	return
}

// RecorderAuxMockPtrframerparseEventFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseEventFrame int = 0

var condRecorderAuxMockPtrframerparseEventFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseEventFrame(i int) {
	condRecorderAuxMockPtrframerparseEventFrame.L.Lock()
	for recorderAuxMockPtrframerparseEventFrame < i {
		condRecorderAuxMockPtrframerparseEventFrame.Wait()
	}
	condRecorderAuxMockPtrframerparseEventFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseEventFrame() {
	condRecorderAuxMockPtrframerparseEventFrame.L.Lock()
	recorderAuxMockPtrframerparseEventFrame++
	condRecorderAuxMockPtrframerparseEventFrame.L.Unlock()
	condRecorderAuxMockPtrframerparseEventFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseEventFrame() (ret int) {
	condRecorderAuxMockPtrframerparseEventFrame.L.Lock()
	ret = recorderAuxMockPtrframerparseEventFrame
	condRecorderAuxMockPtrframerparseEventFrame.L.Unlock()
	return
}

// (recvf *framer)parseEventFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseEventFrame() (reta frame) {
	FuncAuxMockPtrframerparseEventFrame, ok := apomock.GetRegisteredFunc("gocql.framer.parseEventFrame")
	if ok {
		reta = FuncAuxMockPtrframerparseEventFrame.(func(recvf *framer) (reta frame))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseEventFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseEventFrame()
	return
}

//
// Mock: (recvf *framer)writeOptionsFrame(argstream int, arg_ *writeOptionsFrame)(reta error)
//

type MockArgsTypeframerwriteOptionsFrame struct {
	ApomockCallNumber int
	Argstream         int
	Arg_              *writeOptionsFrame
}

var LastMockArgsframerwriteOptionsFrame MockArgsTypeframerwriteOptionsFrame

// (recvf *framer)AuxMockwriteOptionsFrame(argstream int, arg_ *writeOptionsFrame)(reta error) - Generated mock function
func (recvf *framer) AuxMockwriteOptionsFrame(argstream int, arg_ *writeOptionsFrame) (reta error) {
	LastMockArgsframerwriteOptionsFrame = MockArgsTypeframerwriteOptionsFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteOptionsFrame(),
		Argstream:         argstream,
		Arg_:              arg_,
	}
	rargs, rerr := apomock.GetNext("gocql.framer.writeOptionsFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.writeOptionsFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.writeOptionsFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrframerwriteOptionsFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteOptionsFrame int = 0

var condRecorderAuxMockPtrframerwriteOptionsFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteOptionsFrame(i int) {
	condRecorderAuxMockPtrframerwriteOptionsFrame.L.Lock()
	for recorderAuxMockPtrframerwriteOptionsFrame < i {
		condRecorderAuxMockPtrframerwriteOptionsFrame.Wait()
	}
	condRecorderAuxMockPtrframerwriteOptionsFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteOptionsFrame() {
	condRecorderAuxMockPtrframerwriteOptionsFrame.L.Lock()
	recorderAuxMockPtrframerwriteOptionsFrame++
	condRecorderAuxMockPtrframerwriteOptionsFrame.L.Unlock()
	condRecorderAuxMockPtrframerwriteOptionsFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteOptionsFrame() (ret int) {
	condRecorderAuxMockPtrframerwriteOptionsFrame.L.Lock()
	ret = recorderAuxMockPtrframerwriteOptionsFrame
	condRecorderAuxMockPtrframerwriteOptionsFrame.L.Unlock()
	return
}

// (recvf *framer)writeOptionsFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeOptionsFrame(argstream int, arg_ *writeOptionsFrame) (reta error) {
	FuncAuxMockPtrframerwriteOptionsFrame, ok := apomock.GetRegisteredFunc("gocql.framer.writeOptionsFrame")
	if ok {
		reta = FuncAuxMockPtrframerwriteOptionsFrame.(func(recvf *framer, argstream int, arg_ *writeOptionsFrame) (reta error))(recvf, argstream, arg_)
	} else {
		panic("FuncAuxMockPtrframerwriteOptionsFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteOptionsFrame()
	return
}

//
// Mock: (recvf *framer)writeInet(argip net.IP, argport int)()
//

type MockArgsTypeframerwriteInet struct {
	ApomockCallNumber int
	Argip             net.IP
	Argport           int
}

var LastMockArgsframerwriteInet MockArgsTypeframerwriteInet

// (recvf *framer)AuxMockwriteInet(argip net.IP, argport int)() - Generated mock function
func (recvf *framer) AuxMockwriteInet(argip net.IP, argport int) {
	LastMockArgsframerwriteInet = MockArgsTypeframerwriteInet{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteInet(),
		Argip:             argip,
		Argport:           argport,
	}
	return
}

// RecorderAuxMockPtrframerwriteInet  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteInet int = 0

var condRecorderAuxMockPtrframerwriteInet *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteInet(i int) {
	condRecorderAuxMockPtrframerwriteInet.L.Lock()
	for recorderAuxMockPtrframerwriteInet < i {
		condRecorderAuxMockPtrframerwriteInet.Wait()
	}
	condRecorderAuxMockPtrframerwriteInet.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteInet() {
	condRecorderAuxMockPtrframerwriteInet.L.Lock()
	recorderAuxMockPtrframerwriteInet++
	condRecorderAuxMockPtrframerwriteInet.L.Unlock()
	condRecorderAuxMockPtrframerwriteInet.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteInet() (ret int) {
	condRecorderAuxMockPtrframerwriteInet.L.Lock()
	ret = recorderAuxMockPtrframerwriteInet
	condRecorderAuxMockPtrframerwriteInet.L.Unlock()
	return
}

// (recvf *framer)writeInet - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeInet(argip net.IP, argport int) {
	FuncAuxMockPtrframerwriteInet, ok := apomock.GetRegisteredFunc("gocql.framer.writeInet")
	if ok {
		FuncAuxMockPtrframerwriteInet.(func(recvf *framer, argip net.IP, argport int))(recvf, argip, argport)
	} else {
		panic("FuncAuxMockPtrframerwriteInet ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteInet()
	return
}

//
// Mock: (recvf *framer)readTrace()()
//

type MockArgsTypeframerreadTrace struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadTrace MockArgsTypeframerreadTrace

// (recvf *framer)AuxMockreadTrace()() - Generated mock function
func (recvf *framer) AuxMockreadTrace() {
	return
}

// RecorderAuxMockPtrframerreadTrace  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadTrace int = 0

var condRecorderAuxMockPtrframerreadTrace *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadTrace(i int) {
	condRecorderAuxMockPtrframerreadTrace.L.Lock()
	for recorderAuxMockPtrframerreadTrace < i {
		condRecorderAuxMockPtrframerreadTrace.Wait()
	}
	condRecorderAuxMockPtrframerreadTrace.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadTrace() {
	condRecorderAuxMockPtrframerreadTrace.L.Lock()
	recorderAuxMockPtrframerreadTrace++
	condRecorderAuxMockPtrframerreadTrace.L.Unlock()
	condRecorderAuxMockPtrframerreadTrace.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadTrace() (ret int) {
	condRecorderAuxMockPtrframerreadTrace.L.Lock()
	ret = recorderAuxMockPtrframerreadTrace
	condRecorderAuxMockPtrframerreadTrace.L.Unlock()
	return
}

// (recvf *framer)readTrace - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readTrace() {
	FuncAuxMockPtrframerreadTrace, ok := apomock.GetRegisteredFunc("gocql.framer.readTrace")
	if ok {
		FuncAuxMockPtrframerreadTrace.(func(recvf *framer))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadTrace ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadTrace()
	return
}

//
// Mock: (recvf *resultRowsFrame)String()(reta string)
//

type MockArgsTyperesultRowsFrameString struct {
	ApomockCallNumber int
}

var LastMockArgsresultRowsFrameString MockArgsTyperesultRowsFrameString

// (recvf *resultRowsFrame)AuxMockString()(reta string) - Generated mock function
func (recvf *resultRowsFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.resultRowsFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.resultRowsFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.resultRowsFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrresultRowsFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrresultRowsFrameString int = 0

var condRecorderAuxMockPtrresultRowsFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrresultRowsFrameString(i int) {
	condRecorderAuxMockPtrresultRowsFrameString.L.Lock()
	for recorderAuxMockPtrresultRowsFrameString < i {
		condRecorderAuxMockPtrresultRowsFrameString.Wait()
	}
	condRecorderAuxMockPtrresultRowsFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrresultRowsFrameString() {
	condRecorderAuxMockPtrresultRowsFrameString.L.Lock()
	recorderAuxMockPtrresultRowsFrameString++
	condRecorderAuxMockPtrresultRowsFrameString.L.Unlock()
	condRecorderAuxMockPtrresultRowsFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrresultRowsFrameString() (ret int) {
	condRecorderAuxMockPtrresultRowsFrameString.L.Lock()
	ret = recorderAuxMockPtrresultRowsFrameString
	condRecorderAuxMockPtrresultRowsFrameString.L.Unlock()
	return
}

// (recvf *resultRowsFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *resultRowsFrame) String() (reta string) {
	FuncAuxMockPtrresultRowsFrameString, ok := apomock.GetRegisteredFunc("gocql.resultRowsFrame.String")
	if ok {
		reta = FuncAuxMockPtrresultRowsFrameString.(func(recvf *resultRowsFrame) (reta string))(recvf)
	} else {
		panic("FuncAuxMockPtrresultRowsFrameString ")
	}
	AuxMockIncrementRecorderAuxMockPtrresultRowsFrameString()
	return
}

//
// Mock: (recva *writeAuthResponseFrame)writeFrame(argframer *framer, argstreamID int)(reta error)
//

type MockArgsTypewriteAuthResponseFramewriteFrame struct {
	ApomockCallNumber int
	Argframer         *framer
	ArgstreamID       int
}

var LastMockArgswriteAuthResponseFramewriteFrame MockArgsTypewriteAuthResponseFramewriteFrame

// (recva *writeAuthResponseFrame)AuxMockwriteFrame(argframer *framer, argstreamID int)(reta error) - Generated mock function
func (recva *writeAuthResponseFrame) AuxMockwriteFrame(argframer *framer, argstreamID int) (reta error) {
	LastMockArgswriteAuthResponseFramewriteFrame = MockArgsTypewriteAuthResponseFramewriteFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrwriteAuthResponseFramewriteFrame(),
		Argframer:         argframer,
		ArgstreamID:       argstreamID,
	}
	rargs, rerr := apomock.GetNext("gocql.writeAuthResponseFrame.writeFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeAuthResponseFrame.writeFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeAuthResponseFrame.writeFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrwriteAuthResponseFramewriteFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrwriteAuthResponseFramewriteFrame int = 0

var condRecorderAuxMockPtrwriteAuthResponseFramewriteFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrwriteAuthResponseFramewriteFrame(i int) {
	condRecorderAuxMockPtrwriteAuthResponseFramewriteFrame.L.Lock()
	for recorderAuxMockPtrwriteAuthResponseFramewriteFrame < i {
		condRecorderAuxMockPtrwriteAuthResponseFramewriteFrame.Wait()
	}
	condRecorderAuxMockPtrwriteAuthResponseFramewriteFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrwriteAuthResponseFramewriteFrame() {
	condRecorderAuxMockPtrwriteAuthResponseFramewriteFrame.L.Lock()
	recorderAuxMockPtrwriteAuthResponseFramewriteFrame++
	condRecorderAuxMockPtrwriteAuthResponseFramewriteFrame.L.Unlock()
	condRecorderAuxMockPtrwriteAuthResponseFramewriteFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrwriteAuthResponseFramewriteFrame() (ret int) {
	condRecorderAuxMockPtrwriteAuthResponseFramewriteFrame.L.Lock()
	ret = recorderAuxMockPtrwriteAuthResponseFramewriteFrame
	condRecorderAuxMockPtrwriteAuthResponseFramewriteFrame.L.Unlock()
	return
}

// (recva *writeAuthResponseFrame)writeFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recva *writeAuthResponseFrame) writeFrame(argframer *framer, argstreamID int) (reta error) {
	FuncAuxMockPtrwriteAuthResponseFramewriteFrame, ok := apomock.GetRegisteredFunc("gocql.writeAuthResponseFrame.writeFrame")
	if ok {
		reta = FuncAuxMockPtrwriteAuthResponseFramewriteFrame.(func(recva *writeAuthResponseFrame, argframer *framer, argstreamID int) (reta error))(recva, argframer, argstreamID)
	} else {
		panic("FuncAuxMockPtrwriteAuthResponseFramewriteFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrwriteAuthResponseFramewriteFrame()
	return
}

//
// Mock: (recvp protoVersion)response()(reta bool)
//

type MockArgsTypeprotoVersionresponse struct {
	ApomockCallNumber int
}

var LastMockArgsprotoVersionresponse MockArgsTypeprotoVersionresponse

// (recvp protoVersion)AuxMockresponse()(reta bool) - Generated mock function
func (recvp protoVersion) AuxMockresponse() (reta bool) {
	rargs, rerr := apomock.GetNext("gocql.protoVersion.response")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.protoVersion.response")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.protoVersion.response")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockprotoVersionresponse  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockprotoVersionresponse int = 0

var condRecorderAuxMockprotoVersionresponse *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockprotoVersionresponse(i int) {
	condRecorderAuxMockprotoVersionresponse.L.Lock()
	for recorderAuxMockprotoVersionresponse < i {
		condRecorderAuxMockprotoVersionresponse.Wait()
	}
	condRecorderAuxMockprotoVersionresponse.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockprotoVersionresponse() {
	condRecorderAuxMockprotoVersionresponse.L.Lock()
	recorderAuxMockprotoVersionresponse++
	condRecorderAuxMockprotoVersionresponse.L.Unlock()
	condRecorderAuxMockprotoVersionresponse.Broadcast()
}
func AuxMockGetRecorderAuxMockprotoVersionresponse() (ret int) {
	condRecorderAuxMockprotoVersionresponse.L.Lock()
	ret = recorderAuxMockprotoVersionresponse
	condRecorderAuxMockprotoVersionresponse.L.Unlock()
	return
}

// (recvp protoVersion)response - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvp protoVersion) response() (reta bool) {
	FuncAuxMockprotoVersionresponse, ok := apomock.GetRegisteredFunc("gocql.protoVersion.response")
	if ok {
		reta = FuncAuxMockprotoVersionresponse.(func(recvp protoVersion) (reta bool))(recvp)
	} else {
		panic("FuncAuxMockprotoVersionresponse ")
	}
	AuxMockIncrementRecorderAuxMockprotoVersionresponse()
	return
}

//
// Mock: ParseConsistency(args string)(reta Consistency)
//

type MockArgsTypeParseConsistency struct {
	ApomockCallNumber int
	Args              string
}

var LastMockArgsParseConsistency MockArgsTypeParseConsistency

// AuxMockParseConsistency(args string)(reta Consistency) - Generated mock function
func AuxMockParseConsistency(args string) (reta Consistency) {
	LastMockArgsParseConsistency = MockArgsTypeParseConsistency{
		ApomockCallNumber: AuxMockGetRecorderAuxMockParseConsistency(),
		Args:              args,
	}
	rargs, rerr := apomock.GetNext("gocql.ParseConsistency")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.ParseConsistency")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.ParseConsistency")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(Consistency)
	}
	return
}

// RecorderAuxMockParseConsistency  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockParseConsistency int = 0

var condRecorderAuxMockParseConsistency *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockParseConsistency(i int) {
	condRecorderAuxMockParseConsistency.L.Lock()
	for recorderAuxMockParseConsistency < i {
		condRecorderAuxMockParseConsistency.Wait()
	}
	condRecorderAuxMockParseConsistency.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockParseConsistency() {
	condRecorderAuxMockParseConsistency.L.Lock()
	recorderAuxMockParseConsistency++
	condRecorderAuxMockParseConsistency.L.Unlock()
	condRecorderAuxMockParseConsistency.Broadcast()
}
func AuxMockGetRecorderAuxMockParseConsistency() (ret int) {
	condRecorderAuxMockParseConsistency.L.Lock()
	ret = recorderAuxMockParseConsistency
	condRecorderAuxMockParseConsistency.L.Unlock()
	return
}

// ParseConsistency - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func ParseConsistency(args string) (reta Consistency) {
	FuncAuxMockParseConsistency, ok := apomock.GetRegisteredFunc("gocql.ParseConsistency")
	if ok {
		reta = FuncAuxMockParseConsistency.(func(args string) (reta Consistency))(args)
	} else {
		panic("FuncAuxMockParseConsistency ")
	}
	AuxMockIncrementRecorderAuxMockParseConsistency()
	return
}

//
// Mock: (recvw *writePrepareFrame)writeFrame(argframer *framer, argstreamID int)(reta error)
//

type MockArgsTypewritePrepareFramewriteFrame struct {
	ApomockCallNumber int
	Argframer         *framer
	ArgstreamID       int
}

var LastMockArgswritePrepareFramewriteFrame MockArgsTypewritePrepareFramewriteFrame

// (recvw *writePrepareFrame)AuxMockwriteFrame(argframer *framer, argstreamID int)(reta error) - Generated mock function
func (recvw *writePrepareFrame) AuxMockwriteFrame(argframer *framer, argstreamID int) (reta error) {
	LastMockArgswritePrepareFramewriteFrame = MockArgsTypewritePrepareFramewriteFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrwritePrepareFramewriteFrame(),
		Argframer:         argframer,
		ArgstreamID:       argstreamID,
	}
	rargs, rerr := apomock.GetNext("gocql.writePrepareFrame.writeFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writePrepareFrame.writeFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writePrepareFrame.writeFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrwritePrepareFramewriteFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrwritePrepareFramewriteFrame int = 0

var condRecorderAuxMockPtrwritePrepareFramewriteFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrwritePrepareFramewriteFrame(i int) {
	condRecorderAuxMockPtrwritePrepareFramewriteFrame.L.Lock()
	for recorderAuxMockPtrwritePrepareFramewriteFrame < i {
		condRecorderAuxMockPtrwritePrepareFramewriteFrame.Wait()
	}
	condRecorderAuxMockPtrwritePrepareFramewriteFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrwritePrepareFramewriteFrame() {
	condRecorderAuxMockPtrwritePrepareFramewriteFrame.L.Lock()
	recorderAuxMockPtrwritePrepareFramewriteFrame++
	condRecorderAuxMockPtrwritePrepareFramewriteFrame.L.Unlock()
	condRecorderAuxMockPtrwritePrepareFramewriteFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrwritePrepareFramewriteFrame() (ret int) {
	condRecorderAuxMockPtrwritePrepareFramewriteFrame.L.Lock()
	ret = recorderAuxMockPtrwritePrepareFramewriteFrame
	condRecorderAuxMockPtrwritePrepareFramewriteFrame.L.Unlock()
	return
}

// (recvw *writePrepareFrame)writeFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvw *writePrepareFrame) writeFrame(argframer *framer, argstreamID int) (reta error) {
	FuncAuxMockPtrwritePrepareFramewriteFrame, ok := apomock.GetRegisteredFunc("gocql.writePrepareFrame.writeFrame")
	if ok {
		reta = FuncAuxMockPtrwritePrepareFramewriteFrame.(func(recvw *writePrepareFrame, argframer *framer, argstreamID int) (reta error))(recvw, argframer, argstreamID)
	} else {
		panic("FuncAuxMockPtrwritePrepareFramewriteFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrwritePrepareFramewriteFrame()
	return
}

//
// Mock: (recvf schemaChangeKeyspace)String()(reta string)
//

type MockArgsTypeschemaChangeKeyspaceString struct {
	ApomockCallNumber int
}

var LastMockArgsschemaChangeKeyspaceString MockArgsTypeschemaChangeKeyspaceString

// (recvf schemaChangeKeyspace)AuxMockString()(reta string) - Generated mock function
func (recvf schemaChangeKeyspace) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.schemaChangeKeyspace.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.schemaChangeKeyspace.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.schemaChangeKeyspace.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockschemaChangeKeyspaceString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockschemaChangeKeyspaceString int = 0

var condRecorderAuxMockschemaChangeKeyspaceString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockschemaChangeKeyspaceString(i int) {
	condRecorderAuxMockschemaChangeKeyspaceString.L.Lock()
	for recorderAuxMockschemaChangeKeyspaceString < i {
		condRecorderAuxMockschemaChangeKeyspaceString.Wait()
	}
	condRecorderAuxMockschemaChangeKeyspaceString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockschemaChangeKeyspaceString() {
	condRecorderAuxMockschemaChangeKeyspaceString.L.Lock()
	recorderAuxMockschemaChangeKeyspaceString++
	condRecorderAuxMockschemaChangeKeyspaceString.L.Unlock()
	condRecorderAuxMockschemaChangeKeyspaceString.Broadcast()
}
func AuxMockGetRecorderAuxMockschemaChangeKeyspaceString() (ret int) {
	condRecorderAuxMockschemaChangeKeyspaceString.L.Lock()
	ret = recorderAuxMockschemaChangeKeyspaceString
	condRecorderAuxMockschemaChangeKeyspaceString.L.Unlock()
	return
}

// (recvf schemaChangeKeyspace)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf schemaChangeKeyspace) String() (reta string) {
	FuncAuxMockschemaChangeKeyspaceString, ok := apomock.GetRegisteredFunc("gocql.schemaChangeKeyspace.String")
	if ok {
		reta = FuncAuxMockschemaChangeKeyspaceString.(func(recvf schemaChangeKeyspace) (reta string))(recvf)
	} else {
		panic("FuncAuxMockschemaChangeKeyspaceString ")
	}
	AuxMockIncrementRecorderAuxMockschemaChangeKeyspaceString()
	return
}

//
// Mock: (recvt statusChangeEventFrame)String()(reta string)
//

type MockArgsTypestatusChangeEventFrameString struct {
	ApomockCallNumber int
}

var LastMockArgsstatusChangeEventFrameString MockArgsTypestatusChangeEventFrameString

// (recvt statusChangeEventFrame)AuxMockString()(reta string) - Generated mock function
func (recvt statusChangeEventFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.statusChangeEventFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.statusChangeEventFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.statusChangeEventFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockstatusChangeEventFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockstatusChangeEventFrameString int = 0

var condRecorderAuxMockstatusChangeEventFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockstatusChangeEventFrameString(i int) {
	condRecorderAuxMockstatusChangeEventFrameString.L.Lock()
	for recorderAuxMockstatusChangeEventFrameString < i {
		condRecorderAuxMockstatusChangeEventFrameString.Wait()
	}
	condRecorderAuxMockstatusChangeEventFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockstatusChangeEventFrameString() {
	condRecorderAuxMockstatusChangeEventFrameString.L.Lock()
	recorderAuxMockstatusChangeEventFrameString++
	condRecorderAuxMockstatusChangeEventFrameString.L.Unlock()
	condRecorderAuxMockstatusChangeEventFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockstatusChangeEventFrameString() (ret int) {
	condRecorderAuxMockstatusChangeEventFrameString.L.Lock()
	ret = recorderAuxMockstatusChangeEventFrameString
	condRecorderAuxMockstatusChangeEventFrameString.L.Unlock()
	return
}

// (recvt statusChangeEventFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt statusChangeEventFrame) String() (reta string) {
	FuncAuxMockstatusChangeEventFrameString, ok := apomock.GetRegisteredFunc("gocql.statusChangeEventFrame.String")
	if ok {
		reta = FuncAuxMockstatusChangeEventFrameString.(func(recvt statusChangeEventFrame) (reta string))(recvt)
	} else {
		panic("FuncAuxMockstatusChangeEventFrameString ")
	}
	AuxMockIncrementRecorderAuxMockstatusChangeEventFrameString()
	return
}

//
// Mock: (recvf *framer)readString()(rets string)
//

type MockArgsTypeframerreadString struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadString MockArgsTypeframerreadString

// (recvf *framer)AuxMockreadString()(rets string) - Generated mock function
func (recvf *framer) AuxMockreadString() (rets string) {
	rargs, rerr := apomock.GetNext("gocql.framer.readString")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readString")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readString")
	}
	if rargs.GetArg(0) != nil {
		rets = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrframerreadString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadString int = 0

var condRecorderAuxMockPtrframerreadString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadString(i int) {
	condRecorderAuxMockPtrframerreadString.L.Lock()
	for recorderAuxMockPtrframerreadString < i {
		condRecorderAuxMockPtrframerreadString.Wait()
	}
	condRecorderAuxMockPtrframerreadString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadString() {
	condRecorderAuxMockPtrframerreadString.L.Lock()
	recorderAuxMockPtrframerreadString++
	condRecorderAuxMockPtrframerreadString.L.Unlock()
	condRecorderAuxMockPtrframerreadString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadString() (ret int) {
	condRecorderAuxMockPtrframerreadString.L.Lock()
	ret = recorderAuxMockPtrframerreadString
	condRecorderAuxMockPtrframerreadString.L.Unlock()
	return
}

// (recvf *framer)readString - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readString() (rets string) {
	FuncAuxMockPtrframerreadString, ok := apomock.GetRegisteredFunc("gocql.framer.readString")
	if ok {
		rets = FuncAuxMockPtrframerreadString.(func(recvf *framer) (rets string))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadString ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadString()
	return
}

//
// Mock: (recvf *framer)readStringMultiMap()(reta map[string][]string)
//

type MockArgsTypeframerreadStringMultiMap struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadStringMultiMap MockArgsTypeframerreadStringMultiMap

// (recvf *framer)AuxMockreadStringMultiMap()(reta map[string][]string) - Generated mock function
func (recvf *framer) AuxMockreadStringMultiMap() (reta map[string][]string) {
	rargs, rerr := apomock.GetNext("gocql.framer.readStringMultiMap")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readStringMultiMap")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readStringMultiMap")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(map[string][]string)
	}
	return
}

// RecorderAuxMockPtrframerreadStringMultiMap  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadStringMultiMap int = 0

var condRecorderAuxMockPtrframerreadStringMultiMap *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadStringMultiMap(i int) {
	condRecorderAuxMockPtrframerreadStringMultiMap.L.Lock()
	for recorderAuxMockPtrframerreadStringMultiMap < i {
		condRecorderAuxMockPtrframerreadStringMultiMap.Wait()
	}
	condRecorderAuxMockPtrframerreadStringMultiMap.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadStringMultiMap() {
	condRecorderAuxMockPtrframerreadStringMultiMap.L.Lock()
	recorderAuxMockPtrframerreadStringMultiMap++
	condRecorderAuxMockPtrframerreadStringMultiMap.L.Unlock()
	condRecorderAuxMockPtrframerreadStringMultiMap.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadStringMultiMap() (ret int) {
	condRecorderAuxMockPtrframerreadStringMultiMap.L.Lock()
	ret = recorderAuxMockPtrframerreadStringMultiMap
	condRecorderAuxMockPtrframerreadStringMultiMap.L.Unlock()
	return
}

// (recvf *framer)readStringMultiMap - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readStringMultiMap() (reta map[string][]string) {
	FuncAuxMockPtrframerreadStringMultiMap, ok := apomock.GetRegisteredFunc("gocql.framer.readStringMultiMap")
	if ok {
		reta = FuncAuxMockPtrframerreadStringMultiMap.(func(recvf *framer) (reta map[string][]string))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadStringMultiMap ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadStringMultiMap()
	return
}

//
// Mock: (recvw *writeQueryFrame)String()(reta string)
//

type MockArgsTypewriteQueryFrameString struct {
	ApomockCallNumber int
}

var LastMockArgswriteQueryFrameString MockArgsTypewriteQueryFrameString

// (recvw *writeQueryFrame)AuxMockString()(reta string) - Generated mock function
func (recvw *writeQueryFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.writeQueryFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeQueryFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeQueryFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrwriteQueryFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrwriteQueryFrameString int = 0

var condRecorderAuxMockPtrwriteQueryFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrwriteQueryFrameString(i int) {
	condRecorderAuxMockPtrwriteQueryFrameString.L.Lock()
	for recorderAuxMockPtrwriteQueryFrameString < i {
		condRecorderAuxMockPtrwriteQueryFrameString.Wait()
	}
	condRecorderAuxMockPtrwriteQueryFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrwriteQueryFrameString() {
	condRecorderAuxMockPtrwriteQueryFrameString.L.Lock()
	recorderAuxMockPtrwriteQueryFrameString++
	condRecorderAuxMockPtrwriteQueryFrameString.L.Unlock()
	condRecorderAuxMockPtrwriteQueryFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrwriteQueryFrameString() (ret int) {
	condRecorderAuxMockPtrwriteQueryFrameString.L.Lock()
	ret = recorderAuxMockPtrwriteQueryFrameString
	condRecorderAuxMockPtrwriteQueryFrameString.L.Unlock()
	return
}

// (recvw *writeQueryFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvw *writeQueryFrame) String() (reta string) {
	FuncAuxMockPtrwriteQueryFrameString, ok := apomock.GetRegisteredFunc("gocql.writeQueryFrame.String")
	if ok {
		reta = FuncAuxMockPtrwriteQueryFrameString.(func(recvw *writeQueryFrame) (reta string))(recvw)
	} else {
		panic("FuncAuxMockPtrwriteQueryFrameString ")
	}
	AuxMockIncrementRecorderAuxMockPtrwriteQueryFrameString()
	return
}

//
// Mock: (recva *writeAuthResponseFrame)String()(reta string)
//

type MockArgsTypewriteAuthResponseFrameString struct {
	ApomockCallNumber int
}

var LastMockArgswriteAuthResponseFrameString MockArgsTypewriteAuthResponseFrameString

// (recva *writeAuthResponseFrame)AuxMockString()(reta string) - Generated mock function
func (recva *writeAuthResponseFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.writeAuthResponseFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeAuthResponseFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeAuthResponseFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrwriteAuthResponseFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrwriteAuthResponseFrameString int = 0

var condRecorderAuxMockPtrwriteAuthResponseFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrwriteAuthResponseFrameString(i int) {
	condRecorderAuxMockPtrwriteAuthResponseFrameString.L.Lock()
	for recorderAuxMockPtrwriteAuthResponseFrameString < i {
		condRecorderAuxMockPtrwriteAuthResponseFrameString.Wait()
	}
	condRecorderAuxMockPtrwriteAuthResponseFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrwriteAuthResponseFrameString() {
	condRecorderAuxMockPtrwriteAuthResponseFrameString.L.Lock()
	recorderAuxMockPtrwriteAuthResponseFrameString++
	condRecorderAuxMockPtrwriteAuthResponseFrameString.L.Unlock()
	condRecorderAuxMockPtrwriteAuthResponseFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrwriteAuthResponseFrameString() (ret int) {
	condRecorderAuxMockPtrwriteAuthResponseFrameString.L.Lock()
	ret = recorderAuxMockPtrwriteAuthResponseFrameString
	condRecorderAuxMockPtrwriteAuthResponseFrameString.L.Unlock()
	return
}

// (recva *writeAuthResponseFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recva *writeAuthResponseFrame) String() (reta string) {
	FuncAuxMockPtrwriteAuthResponseFrameString, ok := apomock.GetRegisteredFunc("gocql.writeAuthResponseFrame.String")
	if ok {
		reta = FuncAuxMockPtrwriteAuthResponseFrameString.(func(recva *writeAuthResponseFrame) (reta string))(recva)
	} else {
		panic("FuncAuxMockPtrwriteAuthResponseFrameString ")
	}
	AuxMockIncrementRecorderAuxMockPtrwriteAuthResponseFrameString()
	return
}

//
// Mock: (recvf *framer)readShortBytes()(reta []byte)
//

type MockArgsTypeframerreadShortBytes struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadShortBytes MockArgsTypeframerreadShortBytes

// (recvf *framer)AuxMockreadShortBytes()(reta []byte) - Generated mock function
func (recvf *framer) AuxMockreadShortBytes() (reta []byte) {
	rargs, rerr := apomock.GetNext("gocql.framer.readShortBytes")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readShortBytes")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readShortBytes")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockPtrframerreadShortBytes  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadShortBytes int = 0

var condRecorderAuxMockPtrframerreadShortBytes *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadShortBytes(i int) {
	condRecorderAuxMockPtrframerreadShortBytes.L.Lock()
	for recorderAuxMockPtrframerreadShortBytes < i {
		condRecorderAuxMockPtrframerreadShortBytes.Wait()
	}
	condRecorderAuxMockPtrframerreadShortBytes.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadShortBytes() {
	condRecorderAuxMockPtrframerreadShortBytes.L.Lock()
	recorderAuxMockPtrframerreadShortBytes++
	condRecorderAuxMockPtrframerreadShortBytes.L.Unlock()
	condRecorderAuxMockPtrframerreadShortBytes.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadShortBytes() (ret int) {
	condRecorderAuxMockPtrframerreadShortBytes.L.Lock()
	ret = recorderAuxMockPtrframerreadShortBytes
	condRecorderAuxMockPtrframerreadShortBytes.L.Unlock()
	return
}

// (recvf *framer)readShortBytes - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readShortBytes() (reta []byte) {
	FuncAuxMockPtrframerreadShortBytes, ok := apomock.GetRegisteredFunc("gocql.framer.readShortBytes")
	if ok {
		reta = FuncAuxMockPtrframerreadShortBytes.(func(recvf *framer) (reta []byte))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadShortBytes ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadShortBytes()
	return
}

//
// Mock: (recvs SerialConsistency)String()(reta string)
//

type MockArgsTypeSerialConsistencyString struct {
	ApomockCallNumber int
}

var LastMockArgsSerialConsistencyString MockArgsTypeSerialConsistencyString

// (recvs SerialConsistency)AuxMockString()(reta string) - Generated mock function
func (recvs SerialConsistency) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.SerialConsistency.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.SerialConsistency.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.SerialConsistency.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockSerialConsistencyString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockSerialConsistencyString int = 0

var condRecorderAuxMockSerialConsistencyString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockSerialConsistencyString(i int) {
	condRecorderAuxMockSerialConsistencyString.L.Lock()
	for recorderAuxMockSerialConsistencyString < i {
		condRecorderAuxMockSerialConsistencyString.Wait()
	}
	condRecorderAuxMockSerialConsistencyString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockSerialConsistencyString() {
	condRecorderAuxMockSerialConsistencyString.L.Lock()
	recorderAuxMockSerialConsistencyString++
	condRecorderAuxMockSerialConsistencyString.L.Unlock()
	condRecorderAuxMockSerialConsistencyString.Broadcast()
}
func AuxMockGetRecorderAuxMockSerialConsistencyString() (ret int) {
	condRecorderAuxMockSerialConsistencyString.L.Lock()
	ret = recorderAuxMockSerialConsistencyString
	condRecorderAuxMockSerialConsistencyString.L.Unlock()
	return
}

// (recvs SerialConsistency)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs SerialConsistency) String() (reta string) {
	FuncAuxMockSerialConsistencyString, ok := apomock.GetRegisteredFunc("gocql.SerialConsistency.String")
	if ok {
		reta = FuncAuxMockSerialConsistencyString.(func(recvs SerialConsistency) (reta string))(recvs)
	} else {
		panic("FuncAuxMockSerialConsistencyString ")
	}
	AuxMockIncrementRecorderAuxMockSerialConsistencyString()
	return
}

//
// Mock: readShort(argp []byte)(reta uint16)
//

type MockArgsTypereadShort struct {
	ApomockCallNumber int
	Argp              []byte
}

var LastMockArgsreadShort MockArgsTypereadShort

// AuxMockreadShort(argp []byte)(reta uint16) - Generated mock function
func AuxMockreadShort(argp []byte) (reta uint16) {
	LastMockArgsreadShort = MockArgsTypereadShort{
		ApomockCallNumber: AuxMockGetRecorderAuxMockreadShort(),
		Argp:              argp,
	}
	rargs, rerr := apomock.GetNext("gocql.readShort")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.readShort")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.readShort")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(uint16)
	}
	return
}

// RecorderAuxMockreadShort  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockreadShort int = 0

var condRecorderAuxMockreadShort *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockreadShort(i int) {
	condRecorderAuxMockreadShort.L.Lock()
	for recorderAuxMockreadShort < i {
		condRecorderAuxMockreadShort.Wait()
	}
	condRecorderAuxMockreadShort.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockreadShort() {
	condRecorderAuxMockreadShort.L.Lock()
	recorderAuxMockreadShort++
	condRecorderAuxMockreadShort.L.Unlock()
	condRecorderAuxMockreadShort.Broadcast()
}
func AuxMockGetRecorderAuxMockreadShort() (ret int) {
	condRecorderAuxMockreadShort.L.Lock()
	ret = recorderAuxMockreadShort
	condRecorderAuxMockreadShort.L.Unlock()
	return
}

// readShort - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func readShort(argp []byte) (reta uint16) {
	FuncAuxMockreadShort, ok := apomock.GetRegisteredFunc("gocql.readShort")
	if ok {
		reta = FuncAuxMockreadShort.(func(argp []byte) (reta uint16))(argp)
	} else {
		panic("FuncAuxMockreadShort ")
	}
	AuxMockIncrementRecorderAuxMockreadShort()
	return
}

//
// Mock: (recvf *framer)readByte()(reta byte)
//

type MockArgsTypeframerreadByte struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadByte MockArgsTypeframerreadByte

// (recvf *framer)AuxMockreadByte()(reta byte) - Generated mock function
func (recvf *framer) AuxMockreadByte() (reta byte) {
	rargs, rerr := apomock.GetNext("gocql.framer.readByte")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readByte")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readByte")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(byte)
	}
	return
}

// RecorderAuxMockPtrframerreadByte  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadByte int = 0

var condRecorderAuxMockPtrframerreadByte *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadByte(i int) {
	condRecorderAuxMockPtrframerreadByte.L.Lock()
	for recorderAuxMockPtrframerreadByte < i {
		condRecorderAuxMockPtrframerreadByte.Wait()
	}
	condRecorderAuxMockPtrframerreadByte.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadByte() {
	condRecorderAuxMockPtrframerreadByte.L.Lock()
	recorderAuxMockPtrframerreadByte++
	condRecorderAuxMockPtrframerreadByte.L.Unlock()
	condRecorderAuxMockPtrframerreadByte.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadByte() (ret int) {
	condRecorderAuxMockPtrframerreadByte.L.Lock()
	ret = recorderAuxMockPtrframerreadByte
	condRecorderAuxMockPtrframerreadByte.L.Unlock()
	return
}

// (recvf *framer)readByte - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readByte() (reta byte) {
	FuncAuxMockPtrframerreadByte, ok := apomock.GetRegisteredFunc("gocql.framer.readByte")
	if ok {
		reta = FuncAuxMockPtrframerreadByte.(func(recvf *framer) (reta byte))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadByte ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadByte()
	return
}

//
// Mock: (recvf *framer)parseResultFrame()(reta frame, retb error)
//

type MockArgsTypeframerparseResultFrame struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseResultFrame MockArgsTypeframerparseResultFrame

// (recvf *framer)AuxMockparseResultFrame()(reta frame, retb error) - Generated mock function
func (recvf *framer) AuxMockparseResultFrame() (reta frame, retb error) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseResultFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseResultFrame")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.framer.parseResultFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrframerparseResultFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseResultFrame int = 0

var condRecorderAuxMockPtrframerparseResultFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseResultFrame(i int) {
	condRecorderAuxMockPtrframerparseResultFrame.L.Lock()
	for recorderAuxMockPtrframerparseResultFrame < i {
		condRecorderAuxMockPtrframerparseResultFrame.Wait()
	}
	condRecorderAuxMockPtrframerparseResultFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseResultFrame() {
	condRecorderAuxMockPtrframerparseResultFrame.L.Lock()
	recorderAuxMockPtrframerparseResultFrame++
	condRecorderAuxMockPtrframerparseResultFrame.L.Unlock()
	condRecorderAuxMockPtrframerparseResultFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseResultFrame() (ret int) {
	condRecorderAuxMockPtrframerparseResultFrame.L.Lock()
	ret = recorderAuxMockPtrframerparseResultFrame
	condRecorderAuxMockPtrframerparseResultFrame.L.Unlock()
	return
}

// (recvf *framer)parseResultFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseResultFrame() (reta frame, retb error) {
	FuncAuxMockPtrframerparseResultFrame, ok := apomock.GetRegisteredFunc("gocql.framer.parseResultFrame")
	if ok {
		reta, retb = FuncAuxMockPtrframerparseResultFrame.(func(recvf *framer) (reta frame, retb error))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseResultFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseResultFrame()
	return
}

//
// Mock: (recvw *writeRegisterFrame)writeFrame(argframer *framer, argstreamID int)(reta error)
//

type MockArgsTypewriteRegisterFramewriteFrame struct {
	ApomockCallNumber int
	Argframer         *framer
	ArgstreamID       int
}

var LastMockArgswriteRegisterFramewriteFrame MockArgsTypewriteRegisterFramewriteFrame

// (recvw *writeRegisterFrame)AuxMockwriteFrame(argframer *framer, argstreamID int)(reta error) - Generated mock function
func (recvw *writeRegisterFrame) AuxMockwriteFrame(argframer *framer, argstreamID int) (reta error) {
	LastMockArgswriteRegisterFramewriteFrame = MockArgsTypewriteRegisterFramewriteFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrwriteRegisterFramewriteFrame(),
		Argframer:         argframer,
		ArgstreamID:       argstreamID,
	}
	rargs, rerr := apomock.GetNext("gocql.writeRegisterFrame.writeFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeRegisterFrame.writeFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeRegisterFrame.writeFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrwriteRegisterFramewriteFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrwriteRegisterFramewriteFrame int = 0

var condRecorderAuxMockPtrwriteRegisterFramewriteFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrwriteRegisterFramewriteFrame(i int) {
	condRecorderAuxMockPtrwriteRegisterFramewriteFrame.L.Lock()
	for recorderAuxMockPtrwriteRegisterFramewriteFrame < i {
		condRecorderAuxMockPtrwriteRegisterFramewriteFrame.Wait()
	}
	condRecorderAuxMockPtrwriteRegisterFramewriteFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrwriteRegisterFramewriteFrame() {
	condRecorderAuxMockPtrwriteRegisterFramewriteFrame.L.Lock()
	recorderAuxMockPtrwriteRegisterFramewriteFrame++
	condRecorderAuxMockPtrwriteRegisterFramewriteFrame.L.Unlock()
	condRecorderAuxMockPtrwriteRegisterFramewriteFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrwriteRegisterFramewriteFrame() (ret int) {
	condRecorderAuxMockPtrwriteRegisterFramewriteFrame.L.Lock()
	ret = recorderAuxMockPtrwriteRegisterFramewriteFrame
	condRecorderAuxMockPtrwriteRegisterFramewriteFrame.L.Unlock()
	return
}

// (recvw *writeRegisterFrame)writeFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvw *writeRegisterFrame) writeFrame(argframer *framer, argstreamID int) (reta error) {
	FuncAuxMockPtrwriteRegisterFramewriteFrame, ok := apomock.GetRegisteredFunc("gocql.writeRegisterFrame.writeFrame")
	if ok {
		reta = FuncAuxMockPtrwriteRegisterFramewriteFrame.(func(recvw *writeRegisterFrame, argframer *framer, argstreamID int) (reta error))(recvw, argframer, argstreamID)
	} else {
		panic("FuncAuxMockPtrwriteRegisterFramewriteFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrwriteRegisterFramewriteFrame()
	return
}

//
// Mock: (recvf *framer)writeShort(argn uint16)()
//

type MockArgsTypeframerwriteShort struct {
	ApomockCallNumber int
	Argn              uint16
}

var LastMockArgsframerwriteShort MockArgsTypeframerwriteShort

// (recvf *framer)AuxMockwriteShort(argn uint16)() - Generated mock function
func (recvf *framer) AuxMockwriteShort(argn uint16) {
	LastMockArgsframerwriteShort = MockArgsTypeframerwriteShort{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteShort(),
		Argn:              argn,
	}
	return
}

// RecorderAuxMockPtrframerwriteShort  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteShort int = 0

var condRecorderAuxMockPtrframerwriteShort *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteShort(i int) {
	condRecorderAuxMockPtrframerwriteShort.L.Lock()
	for recorderAuxMockPtrframerwriteShort < i {
		condRecorderAuxMockPtrframerwriteShort.Wait()
	}
	condRecorderAuxMockPtrframerwriteShort.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteShort() {
	condRecorderAuxMockPtrframerwriteShort.L.Lock()
	recorderAuxMockPtrframerwriteShort++
	condRecorderAuxMockPtrframerwriteShort.L.Unlock()
	condRecorderAuxMockPtrframerwriteShort.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteShort() (ret int) {
	condRecorderAuxMockPtrframerwriteShort.L.Lock()
	ret = recorderAuxMockPtrframerwriteShort
	condRecorderAuxMockPtrframerwriteShort.L.Unlock()
	return
}

// (recvf *framer)writeShort - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeShort(argn uint16) {
	FuncAuxMockPtrframerwriteShort, ok := apomock.GetRegisteredFunc("gocql.framer.writeShort")
	if ok {
		FuncAuxMockPtrframerwriteShort.(func(recvf *framer, argn uint16))(recvf, argn)
	} else {
		panic("FuncAuxMockPtrframerwriteShort ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteShort()
	return
}

//
// Mock: (recvf *framer)parseReadyFrame()(reta frame)
//

type MockArgsTypeframerparseReadyFrame struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseReadyFrame MockArgsTypeframerparseReadyFrame

// (recvf *framer)AuxMockparseReadyFrame()(reta frame) - Generated mock function
func (recvf *framer) AuxMockparseReadyFrame() (reta frame) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseReadyFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseReadyFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseReadyFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	return
}

// RecorderAuxMockPtrframerparseReadyFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseReadyFrame int = 0

var condRecorderAuxMockPtrframerparseReadyFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseReadyFrame(i int) {
	condRecorderAuxMockPtrframerparseReadyFrame.L.Lock()
	for recorderAuxMockPtrframerparseReadyFrame < i {
		condRecorderAuxMockPtrframerparseReadyFrame.Wait()
	}
	condRecorderAuxMockPtrframerparseReadyFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseReadyFrame() {
	condRecorderAuxMockPtrframerparseReadyFrame.L.Lock()
	recorderAuxMockPtrframerparseReadyFrame++
	condRecorderAuxMockPtrframerparseReadyFrame.L.Unlock()
	condRecorderAuxMockPtrframerparseReadyFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseReadyFrame() (ret int) {
	condRecorderAuxMockPtrframerparseReadyFrame.L.Lock()
	ret = recorderAuxMockPtrframerparseReadyFrame
	condRecorderAuxMockPtrframerparseReadyFrame.L.Unlock()
	return
}

// (recvf *framer)parseReadyFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseReadyFrame() (reta frame) {
	FuncAuxMockPtrframerparseReadyFrame, ok := apomock.GetRegisteredFunc("gocql.framer.parseReadyFrame")
	if ok {
		reta = FuncAuxMockPtrframerparseReadyFrame.(func(recvf *framer) (reta frame))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseReadyFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseReadyFrame()
	return
}

//
// Mock: writeShort(argp []byte, argn uint16)()
//

type MockArgsTypewriteShort struct {
	ApomockCallNumber int
	Argp              []byte
	Argn              uint16
}

var LastMockArgswriteShort MockArgsTypewriteShort

// AuxMockwriteShort(argp []byte, argn uint16)() - Generated mock function
func AuxMockwriteShort(argp []byte, argn uint16) {
	LastMockArgswriteShort = MockArgsTypewriteShort{
		ApomockCallNumber: AuxMockGetRecorderAuxMockwriteShort(),
		Argp:              argp,
		Argn:              argn,
	}
	return
}

// RecorderAuxMockwriteShort  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockwriteShort int = 0

var condRecorderAuxMockwriteShort *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockwriteShort(i int) {
	condRecorderAuxMockwriteShort.L.Lock()
	for recorderAuxMockwriteShort < i {
		condRecorderAuxMockwriteShort.Wait()
	}
	condRecorderAuxMockwriteShort.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockwriteShort() {
	condRecorderAuxMockwriteShort.L.Lock()
	recorderAuxMockwriteShort++
	condRecorderAuxMockwriteShort.L.Unlock()
	condRecorderAuxMockwriteShort.Broadcast()
}
func AuxMockGetRecorderAuxMockwriteShort() (ret int) {
	condRecorderAuxMockwriteShort.L.Lock()
	ret = recorderAuxMockwriteShort
	condRecorderAuxMockwriteShort.L.Unlock()
	return
}

// writeShort - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func writeShort(argp []byte, argn uint16) {
	FuncAuxMockwriteShort, ok := apomock.GetRegisteredFunc("gocql.writeShort")
	if ok {
		FuncAuxMockwriteShort.(func(argp []byte, argn uint16))(argp, argn)
	} else {
		panic("FuncAuxMockwriteShort ")
	}
	AuxMockIncrementRecorderAuxMockwriteShort()
	return
}

//
// Mock: (recvf *framer)trace()()
//

type MockArgsTypeframertrace struct {
	ApomockCallNumber int
}

var LastMockArgsframertrace MockArgsTypeframertrace

// (recvf *framer)AuxMocktrace()() - Generated mock function
func (recvf *framer) AuxMocktrace() {
	return
}

// RecorderAuxMockPtrframertrace  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframertrace int = 0

var condRecorderAuxMockPtrframertrace *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframertrace(i int) {
	condRecorderAuxMockPtrframertrace.L.Lock()
	for recorderAuxMockPtrframertrace < i {
		condRecorderAuxMockPtrframertrace.Wait()
	}
	condRecorderAuxMockPtrframertrace.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframertrace() {
	condRecorderAuxMockPtrframertrace.L.Lock()
	recorderAuxMockPtrframertrace++
	condRecorderAuxMockPtrframertrace.L.Unlock()
	condRecorderAuxMockPtrframertrace.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframertrace() (ret int) {
	condRecorderAuxMockPtrframertrace.L.Lock()
	ret = recorderAuxMockPtrframertrace
	condRecorderAuxMockPtrframertrace.L.Unlock()
	return
}

// (recvf *framer)trace - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) trace() {
	FuncAuxMockPtrframertrace, ok := apomock.GetRegisteredFunc("gocql.framer.trace")
	if ok {
		FuncAuxMockPtrframertrace.(func(recvf *framer))(recvf)
	} else {
		panic("FuncAuxMockPtrframertrace ")
	}
	AuxMockIncrementRecorderAuxMockPtrframertrace()
	return
}

//
// Mock: (recva *authSuccessFrame)String()(reta string)
//

type MockArgsTypeauthSuccessFrameString struct {
	ApomockCallNumber int
}

var LastMockArgsauthSuccessFrameString MockArgsTypeauthSuccessFrameString

// (recva *authSuccessFrame)AuxMockString()(reta string) - Generated mock function
func (recva *authSuccessFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.authSuccessFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.authSuccessFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.authSuccessFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockPtrauthSuccessFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrauthSuccessFrameString int = 0

var condRecorderAuxMockPtrauthSuccessFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrauthSuccessFrameString(i int) {
	condRecorderAuxMockPtrauthSuccessFrameString.L.Lock()
	for recorderAuxMockPtrauthSuccessFrameString < i {
		condRecorderAuxMockPtrauthSuccessFrameString.Wait()
	}
	condRecorderAuxMockPtrauthSuccessFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrauthSuccessFrameString() {
	condRecorderAuxMockPtrauthSuccessFrameString.L.Lock()
	recorderAuxMockPtrauthSuccessFrameString++
	condRecorderAuxMockPtrauthSuccessFrameString.L.Unlock()
	condRecorderAuxMockPtrauthSuccessFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrauthSuccessFrameString() (ret int) {
	condRecorderAuxMockPtrauthSuccessFrameString.L.Lock()
	ret = recorderAuxMockPtrauthSuccessFrameString
	condRecorderAuxMockPtrauthSuccessFrameString.L.Unlock()
	return
}

// (recva *authSuccessFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recva *authSuccessFrame) String() (reta string) {
	FuncAuxMockPtrauthSuccessFrameString, ok := apomock.GetRegisteredFunc("gocql.authSuccessFrame.String")
	if ok {
		reta = FuncAuxMockPtrauthSuccessFrameString.(func(recva *authSuccessFrame) (reta string))(recva)
	} else {
		panic("FuncAuxMockPtrauthSuccessFrameString ")
	}
	AuxMockIncrementRecorderAuxMockPtrauthSuccessFrameString()
	return
}

//
// Mock: newFramer(argr io.Reader, argw io.Writer, argcompressor Compressor, argversion byte)(reta *framer)
//

type MockArgsTypenewFramer struct {
	ApomockCallNumber int
	Argr              io.Reader
	Argw              io.Writer
	Argcompressor     Compressor
	Argversion        byte
}

var LastMockArgsnewFramer MockArgsTypenewFramer

// AuxMocknewFramer(argr io.Reader, argw io.Writer, argcompressor Compressor, argversion byte)(reta *framer) - Generated mock function
func AuxMocknewFramer(argr io.Reader, argw io.Writer, argcompressor Compressor, argversion byte) (reta *framer) {
	LastMockArgsnewFramer = MockArgsTypenewFramer{
		ApomockCallNumber: AuxMockGetRecorderAuxMocknewFramer(),
		Argr:              argr,
		Argw:              argw,
		Argcompressor:     argcompressor,
		Argversion:        argversion,
	}
	rargs, rerr := apomock.GetNext("gocql.newFramer")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.newFramer")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.newFramer")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*framer)
	}
	return
}

// RecorderAuxMocknewFramer  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocknewFramer int = 0

var condRecorderAuxMocknewFramer *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocknewFramer(i int) {
	condRecorderAuxMocknewFramer.L.Lock()
	for recorderAuxMocknewFramer < i {
		condRecorderAuxMocknewFramer.Wait()
	}
	condRecorderAuxMocknewFramer.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocknewFramer() {
	condRecorderAuxMocknewFramer.L.Lock()
	recorderAuxMocknewFramer++
	condRecorderAuxMocknewFramer.L.Unlock()
	condRecorderAuxMocknewFramer.Broadcast()
}
func AuxMockGetRecorderAuxMocknewFramer() (ret int) {
	condRecorderAuxMocknewFramer.L.Lock()
	ret = recorderAuxMocknewFramer
	condRecorderAuxMocknewFramer.L.Unlock()
	return
}

// newFramer - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func newFramer(argr io.Reader, argw io.Writer, argcompressor Compressor, argversion byte) (reta *framer) {
	FuncAuxMocknewFramer, ok := apomock.GetRegisteredFunc("gocql.newFramer")
	if ok {
		reta = FuncAuxMocknewFramer.(func(argr io.Reader, argw io.Writer, argcompressor Compressor, argversion byte) (reta *framer))(argr, argw, argcompressor, argversion)
	} else {
		panic("FuncAuxMocknewFramer ")
	}
	AuxMockIncrementRecorderAuxMocknewFramer()
	return
}

//
// Mock: (recve *writeExecuteFrame)writeFrame(argfr *framer, argstreamID int)(reta error)
//

type MockArgsTypewriteExecuteFramewriteFrame struct {
	ApomockCallNumber int
	Argfr             *framer
	ArgstreamID       int
}

var LastMockArgswriteExecuteFramewriteFrame MockArgsTypewriteExecuteFramewriteFrame

// (recve *writeExecuteFrame)AuxMockwriteFrame(argfr *framer, argstreamID int)(reta error) - Generated mock function
func (recve *writeExecuteFrame) AuxMockwriteFrame(argfr *framer, argstreamID int) (reta error) {
	LastMockArgswriteExecuteFramewriteFrame = MockArgsTypewriteExecuteFramewriteFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrwriteExecuteFramewriteFrame(),
		Argfr:             argfr,
		ArgstreamID:       argstreamID,
	}
	rargs, rerr := apomock.GetNext("gocql.writeExecuteFrame.writeFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeExecuteFrame.writeFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeExecuteFrame.writeFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrwriteExecuteFramewriteFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrwriteExecuteFramewriteFrame int = 0

var condRecorderAuxMockPtrwriteExecuteFramewriteFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrwriteExecuteFramewriteFrame(i int) {
	condRecorderAuxMockPtrwriteExecuteFramewriteFrame.L.Lock()
	for recorderAuxMockPtrwriteExecuteFramewriteFrame < i {
		condRecorderAuxMockPtrwriteExecuteFramewriteFrame.Wait()
	}
	condRecorderAuxMockPtrwriteExecuteFramewriteFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrwriteExecuteFramewriteFrame() {
	condRecorderAuxMockPtrwriteExecuteFramewriteFrame.L.Lock()
	recorderAuxMockPtrwriteExecuteFramewriteFrame++
	condRecorderAuxMockPtrwriteExecuteFramewriteFrame.L.Unlock()
	condRecorderAuxMockPtrwriteExecuteFramewriteFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrwriteExecuteFramewriteFrame() (ret int) {
	condRecorderAuxMockPtrwriteExecuteFramewriteFrame.L.Lock()
	ret = recorderAuxMockPtrwriteExecuteFramewriteFrame
	condRecorderAuxMockPtrwriteExecuteFramewriteFrame.L.Unlock()
	return
}

// (recve *writeExecuteFrame)writeFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recve *writeExecuteFrame) writeFrame(argfr *framer, argstreamID int) (reta error) {
	FuncAuxMockPtrwriteExecuteFramewriteFrame, ok := apomock.GetRegisteredFunc("gocql.writeExecuteFrame.writeFrame")
	if ok {
		reta = FuncAuxMockPtrwriteExecuteFramewriteFrame.(func(recve *writeExecuteFrame, argfr *framer, argstreamID int) (reta error))(recve, argfr, argstreamID)
	} else {
		panic("FuncAuxMockPtrwriteExecuteFramewriteFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrwriteExecuteFramewriteFrame()
	return
}

//
// Mock: appendShort(argp []byte, argn uint16)(reta []byte)
//

type MockArgsTypeappendShort struct {
	ApomockCallNumber int
	Argp              []byte
	Argn              uint16
}

var LastMockArgsappendShort MockArgsTypeappendShort

// AuxMockappendShort(argp []byte, argn uint16)(reta []byte) - Generated mock function
func AuxMockappendShort(argp []byte, argn uint16) (reta []byte) {
	LastMockArgsappendShort = MockArgsTypeappendShort{
		ApomockCallNumber: AuxMockGetRecorderAuxMockappendShort(),
		Argp:              argp,
		Argn:              argn,
	}
	rargs, rerr := apomock.GetNext("gocql.appendShort")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.appendShort")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.appendShort")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockappendShort  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockappendShort int = 0

var condRecorderAuxMockappendShort *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockappendShort(i int) {
	condRecorderAuxMockappendShort.L.Lock()
	for recorderAuxMockappendShort < i {
		condRecorderAuxMockappendShort.Wait()
	}
	condRecorderAuxMockappendShort.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockappendShort() {
	condRecorderAuxMockappendShort.L.Lock()
	recorderAuxMockappendShort++
	condRecorderAuxMockappendShort.L.Unlock()
	condRecorderAuxMockappendShort.Broadcast()
}
func AuxMockGetRecorderAuxMockappendShort() (ret int) {
	condRecorderAuxMockappendShort.L.Lock()
	ret = recorderAuxMockappendShort
	condRecorderAuxMockappendShort.L.Unlock()
	return
}

// appendShort - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func appendShort(argp []byte, argn uint16) (reta []byte) {
	FuncAuxMockappendShort, ok := apomock.GetRegisteredFunc("gocql.appendShort")
	if ok {
		reta = FuncAuxMockappendShort.(func(argp []byte, argn uint16) (reta []byte))(argp, argn)
	} else {
		panic("FuncAuxMockappendShort ")
	}
	AuxMockIncrementRecorderAuxMockappendShort()
	return
}

//
// Mock: (recvf *framer)writeLong(argn int64)()
//

type MockArgsTypeframerwriteLong struct {
	ApomockCallNumber int
	Argn              int64
}

var LastMockArgsframerwriteLong MockArgsTypeframerwriteLong

// (recvf *framer)AuxMockwriteLong(argn int64)() - Generated mock function
func (recvf *framer) AuxMockwriteLong(argn int64) {
	LastMockArgsframerwriteLong = MockArgsTypeframerwriteLong{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteLong(),
		Argn:              argn,
	}
	return
}

// RecorderAuxMockPtrframerwriteLong  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteLong int = 0

var condRecorderAuxMockPtrframerwriteLong *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteLong(i int) {
	condRecorderAuxMockPtrframerwriteLong.L.Lock()
	for recorderAuxMockPtrframerwriteLong < i {
		condRecorderAuxMockPtrframerwriteLong.Wait()
	}
	condRecorderAuxMockPtrframerwriteLong.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteLong() {
	condRecorderAuxMockPtrframerwriteLong.L.Lock()
	recorderAuxMockPtrframerwriteLong++
	condRecorderAuxMockPtrframerwriteLong.L.Unlock()
	condRecorderAuxMockPtrframerwriteLong.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteLong() (ret int) {
	condRecorderAuxMockPtrframerwriteLong.L.Lock()
	ret = recorderAuxMockPtrframerwriteLong
	condRecorderAuxMockPtrframerwriteLong.L.Unlock()
	return
}

// (recvf *framer)writeLong - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeLong(argn int64) {
	FuncAuxMockPtrframerwriteLong, ok := apomock.GetRegisteredFunc("gocql.framer.writeLong")
	if ok {
		FuncAuxMockPtrframerwriteLong.(func(recvf *framer, argn int64))(recvf, argn)
	} else {
		panic("FuncAuxMockPtrframerwriteLong ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteLong()
	return
}

//
// Mock: (recvc Consistency)String()(reta string)
//

type MockArgsTypeConsistencyString struct {
	ApomockCallNumber int
}

var LastMockArgsConsistencyString MockArgsTypeConsistencyString

// (recvc Consistency)AuxMockString()(reta string) - Generated mock function
func (recvc Consistency) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.Consistency.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.Consistency.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.Consistency.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockConsistencyString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockConsistencyString int = 0

var condRecorderAuxMockConsistencyString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockConsistencyString(i int) {
	condRecorderAuxMockConsistencyString.L.Lock()
	for recorderAuxMockConsistencyString < i {
		condRecorderAuxMockConsistencyString.Wait()
	}
	condRecorderAuxMockConsistencyString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockConsistencyString() {
	condRecorderAuxMockConsistencyString.L.Lock()
	recorderAuxMockConsistencyString++
	condRecorderAuxMockConsistencyString.L.Unlock()
	condRecorderAuxMockConsistencyString.Broadcast()
}
func AuxMockGetRecorderAuxMockConsistencyString() (ret int) {
	condRecorderAuxMockConsistencyString.L.Lock()
	ret = recorderAuxMockConsistencyString
	condRecorderAuxMockConsistencyString.L.Unlock()
	return
}

// (recvc Consistency)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvc Consistency) String() (reta string) {
	FuncAuxMockConsistencyString, ok := apomock.GetRegisteredFunc("gocql.Consistency.String")
	if ok {
		reta = FuncAuxMockConsistencyString.(func(recvc Consistency) (reta string))(recvc)
	} else {
		panic("FuncAuxMockConsistencyString ")
	}
	AuxMockIncrementRecorderAuxMockConsistencyString()
	return
}

//
// Mock: (recvf *framer)parseSupportedFrame()(reta frame)
//

type MockArgsTypeframerparseSupportedFrame struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseSupportedFrame MockArgsTypeframerparseSupportedFrame

// (recvf *framer)AuxMockparseSupportedFrame()(reta frame) - Generated mock function
func (recvf *framer) AuxMockparseSupportedFrame() (reta frame) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseSupportedFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseSupportedFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseSupportedFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	return
}

// RecorderAuxMockPtrframerparseSupportedFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseSupportedFrame int = 0

var condRecorderAuxMockPtrframerparseSupportedFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseSupportedFrame(i int) {
	condRecorderAuxMockPtrframerparseSupportedFrame.L.Lock()
	for recorderAuxMockPtrframerparseSupportedFrame < i {
		condRecorderAuxMockPtrframerparseSupportedFrame.Wait()
	}
	condRecorderAuxMockPtrframerparseSupportedFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseSupportedFrame() {
	condRecorderAuxMockPtrframerparseSupportedFrame.L.Lock()
	recorderAuxMockPtrframerparseSupportedFrame++
	condRecorderAuxMockPtrframerparseSupportedFrame.L.Unlock()
	condRecorderAuxMockPtrframerparseSupportedFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseSupportedFrame() (ret int) {
	condRecorderAuxMockPtrframerparseSupportedFrame.L.Lock()
	ret = recorderAuxMockPtrframerparseSupportedFrame
	condRecorderAuxMockPtrframerparseSupportedFrame.L.Unlock()
	return
}

// (recvf *framer)parseSupportedFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseSupportedFrame() (reta frame) {
	FuncAuxMockPtrframerparseSupportedFrame, ok := apomock.GetRegisteredFunc("gocql.framer.parseSupportedFrame")
	if ok {
		reta = FuncAuxMockPtrframerparseSupportedFrame.(func(recvf *framer) (reta frame))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseSupportedFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseSupportedFrame()
	return
}

//
// Mock: (recvf *framer)parseFrame()(retframe frame, reterr error)
//

type MockArgsTypeframerparseFrame struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseFrame MockArgsTypeframerparseFrame

// (recvf *framer)AuxMockparseFrame()(retframe frame, reterr error) - Generated mock function
func (recvf *framer) AuxMockparseFrame() (retframe frame, reterr error) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseFrame")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.framer.parseFrame")
	}
	if rargs.GetArg(0) != nil {
		retframe = rargs.GetArg(0).(frame)
	}
	if rargs.GetArg(1) != nil {
		reterr = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrframerparseFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseFrame int = 0

var condRecorderAuxMockPtrframerparseFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseFrame(i int) {
	condRecorderAuxMockPtrframerparseFrame.L.Lock()
	for recorderAuxMockPtrframerparseFrame < i {
		condRecorderAuxMockPtrframerparseFrame.Wait()
	}
	condRecorderAuxMockPtrframerparseFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseFrame() {
	condRecorderAuxMockPtrframerparseFrame.L.Lock()
	recorderAuxMockPtrframerparseFrame++
	condRecorderAuxMockPtrframerparseFrame.L.Unlock()
	condRecorderAuxMockPtrframerparseFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseFrame() (ret int) {
	condRecorderAuxMockPtrframerparseFrame.L.Lock()
	ret = recorderAuxMockPtrframerparseFrame
	condRecorderAuxMockPtrframerparseFrame.L.Unlock()
	return
}

// (recvf *framer)parseFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseFrame() (retframe frame, reterr error) {
	FuncAuxMockPtrframerparseFrame, ok := apomock.GetRegisteredFunc("gocql.framer.parseFrame")
	if ok {
		retframe, reterr = FuncAuxMockPtrframerparseFrame.(func(recvf *framer) (retframe frame, reterr error))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseFrame()
	return
}

//
// Mock: (recvf *framer)readInet()(reta net.IP, retb int)
//

type MockArgsTypeframerreadInet struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadInet MockArgsTypeframerreadInet

// (recvf *framer)AuxMockreadInet()(reta net.IP, retb int) - Generated mock function
func (recvf *framer) AuxMockreadInet() (reta net.IP, retb int) {
	rargs, rerr := apomock.GetNext("gocql.framer.readInet")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readInet")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.framer.readInet")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(net.IP)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(int)
	}
	return
}

// RecorderAuxMockPtrframerreadInet  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadInet int = 0

var condRecorderAuxMockPtrframerreadInet *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadInet(i int) {
	condRecorderAuxMockPtrframerreadInet.L.Lock()
	for recorderAuxMockPtrframerreadInet < i {
		condRecorderAuxMockPtrframerreadInet.Wait()
	}
	condRecorderAuxMockPtrframerreadInet.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadInet() {
	condRecorderAuxMockPtrframerreadInet.L.Lock()
	recorderAuxMockPtrframerreadInet++
	condRecorderAuxMockPtrframerreadInet.L.Unlock()
	condRecorderAuxMockPtrframerreadInet.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadInet() (ret int) {
	condRecorderAuxMockPtrframerreadInet.L.Lock()
	ret = recorderAuxMockPtrframerreadInet
	condRecorderAuxMockPtrframerreadInet.L.Unlock()
	return
}

// (recvf *framer)readInet - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readInet() (reta net.IP, retb int) {
	FuncAuxMockPtrframerreadInet, ok := apomock.GetRegisteredFunc("gocql.framer.readInet")
	if ok {
		reta, retb = FuncAuxMockPtrframerreadInet.(func(recvf *framer) (reta net.IP, retb int))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadInet ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadInet()
	return
}

//
// Mock: (recvf *framer)writeStringList(argl []string)()
//

type MockArgsTypeframerwriteStringList struct {
	ApomockCallNumber int
	Argl              []string
}

var LastMockArgsframerwriteStringList MockArgsTypeframerwriteStringList

// (recvf *framer)AuxMockwriteStringList(argl []string)() - Generated mock function
func (recvf *framer) AuxMockwriteStringList(argl []string) {
	LastMockArgsframerwriteStringList = MockArgsTypeframerwriteStringList{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteStringList(),
		Argl:              argl,
	}
	return
}

// RecorderAuxMockPtrframerwriteStringList  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteStringList int = 0

var condRecorderAuxMockPtrframerwriteStringList *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteStringList(i int) {
	condRecorderAuxMockPtrframerwriteStringList.L.Lock()
	for recorderAuxMockPtrframerwriteStringList < i {
		condRecorderAuxMockPtrframerwriteStringList.Wait()
	}
	condRecorderAuxMockPtrframerwriteStringList.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteStringList() {
	condRecorderAuxMockPtrframerwriteStringList.L.Lock()
	recorderAuxMockPtrframerwriteStringList++
	condRecorderAuxMockPtrframerwriteStringList.L.Unlock()
	condRecorderAuxMockPtrframerwriteStringList.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteStringList() (ret int) {
	condRecorderAuxMockPtrframerwriteStringList.L.Lock()
	ret = recorderAuxMockPtrframerwriteStringList
	condRecorderAuxMockPtrframerwriteStringList.L.Unlock()
	return
}

// (recvf *framer)writeStringList - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeStringList(argl []string) {
	FuncAuxMockPtrframerwriteStringList, ok := apomock.GetRegisteredFunc("gocql.framer.writeStringList")
	if ok {
		FuncAuxMockPtrframerwriteStringList.(func(recvf *framer, argl []string))(recvf, argl)
	} else {
		panic("FuncAuxMockPtrframerwriteStringList ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteStringList()
	return
}

//
// Mock: (recvf *framer)readBytesMap()(reta map[string][]byte)
//

type MockArgsTypeframerreadBytesMap struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadBytesMap MockArgsTypeframerreadBytesMap

// (recvf *framer)AuxMockreadBytesMap()(reta map[string][]byte) - Generated mock function
func (recvf *framer) AuxMockreadBytesMap() (reta map[string][]byte) {
	rargs, rerr := apomock.GetNext("gocql.framer.readBytesMap")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readBytesMap")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readBytesMap")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(map[string][]byte)
	}
	return
}

// RecorderAuxMockPtrframerreadBytesMap  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadBytesMap int = 0

var condRecorderAuxMockPtrframerreadBytesMap *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadBytesMap(i int) {
	condRecorderAuxMockPtrframerreadBytesMap.L.Lock()
	for recorderAuxMockPtrframerreadBytesMap < i {
		condRecorderAuxMockPtrframerreadBytesMap.Wait()
	}
	condRecorderAuxMockPtrframerreadBytesMap.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadBytesMap() {
	condRecorderAuxMockPtrframerreadBytesMap.L.Lock()
	recorderAuxMockPtrframerreadBytesMap++
	condRecorderAuxMockPtrframerreadBytesMap.L.Unlock()
	condRecorderAuxMockPtrframerreadBytesMap.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadBytesMap() (ret int) {
	condRecorderAuxMockPtrframerreadBytesMap.L.Lock()
	ret = recorderAuxMockPtrframerreadBytesMap
	condRecorderAuxMockPtrframerreadBytesMap.L.Unlock()
	return
}

// (recvf *framer)readBytesMap - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readBytesMap() (reta map[string][]byte) {
	FuncAuxMockPtrframerreadBytesMap, ok := apomock.GetRegisteredFunc("gocql.framer.readBytesMap")
	if ok {
		reta = FuncAuxMockPtrframerreadBytesMap.(func(recvf *framer) (reta map[string][]byte))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadBytesMap ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadBytesMap()
	return
}

//
// Mock: (recvq queryParams)String()(reta string)
//

type MockArgsTypequeryParamsString struct {
	ApomockCallNumber int
}

var LastMockArgsqueryParamsString MockArgsTypequeryParamsString

// (recvq queryParams)AuxMockString()(reta string) - Generated mock function
func (recvq queryParams) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.queryParams.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.queryParams.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.queryParams.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockqueryParamsString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockqueryParamsString int = 0

var condRecorderAuxMockqueryParamsString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockqueryParamsString(i int) {
	condRecorderAuxMockqueryParamsString.L.Lock()
	for recorderAuxMockqueryParamsString < i {
		condRecorderAuxMockqueryParamsString.Wait()
	}
	condRecorderAuxMockqueryParamsString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockqueryParamsString() {
	condRecorderAuxMockqueryParamsString.L.Lock()
	recorderAuxMockqueryParamsString++
	condRecorderAuxMockqueryParamsString.L.Unlock()
	condRecorderAuxMockqueryParamsString.Broadcast()
}
func AuxMockGetRecorderAuxMockqueryParamsString() (ret int) {
	condRecorderAuxMockqueryParamsString.L.Lock()
	ret = recorderAuxMockqueryParamsString
	condRecorderAuxMockqueryParamsString.L.Unlock()
	return
}

// (recvq queryParams)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvq queryParams) String() (reta string) {
	FuncAuxMockqueryParamsString, ok := apomock.GetRegisteredFunc("gocql.queryParams.String")
	if ok {
		reta = FuncAuxMockqueryParamsString.(func(recvq queryParams) (reta string))(recvq)
	} else {
		panic("FuncAuxMockqueryParamsString ")
	}
	AuxMockIncrementRecorderAuxMockqueryParamsString()
	return
}

//
// Mock: (recvf *framer)writeString(args string)()
//

type MockArgsTypeframerwriteString struct {
	ApomockCallNumber int
	Args              string
}

var LastMockArgsframerwriteString MockArgsTypeframerwriteString

// (recvf *framer)AuxMockwriteString(args string)() - Generated mock function
func (recvf *framer) AuxMockwriteString(args string) {
	LastMockArgsframerwriteString = MockArgsTypeframerwriteString{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteString(),
		Args:              args,
	}
	return
}

// RecorderAuxMockPtrframerwriteString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteString int = 0

var condRecorderAuxMockPtrframerwriteString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteString(i int) {
	condRecorderAuxMockPtrframerwriteString.L.Lock()
	for recorderAuxMockPtrframerwriteString < i {
		condRecorderAuxMockPtrframerwriteString.Wait()
	}
	condRecorderAuxMockPtrframerwriteString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteString() {
	condRecorderAuxMockPtrframerwriteString.L.Lock()
	recorderAuxMockPtrframerwriteString++
	condRecorderAuxMockPtrframerwriteString.L.Unlock()
	condRecorderAuxMockPtrframerwriteString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteString() (ret int) {
	condRecorderAuxMockPtrframerwriteString.L.Lock()
	ret = recorderAuxMockPtrframerwriteString
	condRecorderAuxMockPtrframerwriteString.L.Unlock()
	return
}

// (recvf *framer)writeString - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeString(args string) {
	FuncAuxMockPtrframerwriteString, ok := apomock.GetRegisteredFunc("gocql.framer.writeString")
	if ok {
		FuncAuxMockPtrframerwriteString.(func(recvf *framer, args string))(recvf, args)
	} else {
		panic("FuncAuxMockPtrframerwriteString ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteString()
	return
}

//
// Mock: (recvf *framer)writeLongString(args string)()
//

type MockArgsTypeframerwriteLongString struct {
	ApomockCallNumber int
	Args              string
}

var LastMockArgsframerwriteLongString MockArgsTypeframerwriteLongString

// (recvf *framer)AuxMockwriteLongString(args string)() - Generated mock function
func (recvf *framer) AuxMockwriteLongString(args string) {
	LastMockArgsframerwriteLongString = MockArgsTypeframerwriteLongString{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteLongString(),
		Args:              args,
	}
	return
}

// RecorderAuxMockPtrframerwriteLongString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteLongString int = 0

var condRecorderAuxMockPtrframerwriteLongString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteLongString(i int) {
	condRecorderAuxMockPtrframerwriteLongString.L.Lock()
	for recorderAuxMockPtrframerwriteLongString < i {
		condRecorderAuxMockPtrframerwriteLongString.Wait()
	}
	condRecorderAuxMockPtrframerwriteLongString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteLongString() {
	condRecorderAuxMockPtrframerwriteLongString.L.Lock()
	recorderAuxMockPtrframerwriteLongString++
	condRecorderAuxMockPtrframerwriteLongString.L.Unlock()
	condRecorderAuxMockPtrframerwriteLongString.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteLongString() (ret int) {
	condRecorderAuxMockPtrframerwriteLongString.L.Lock()
	ret = recorderAuxMockPtrframerwriteLongString
	condRecorderAuxMockPtrframerwriteLongString.L.Unlock()
	return
}

// (recvf *framer)writeLongString - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeLongString(args string) {
	FuncAuxMockPtrframerwriteLongString, ok := apomock.GetRegisteredFunc("gocql.framer.writeLongString")
	if ok {
		FuncAuxMockPtrframerwriteLongString.(func(recvf *framer, args string))(recvf, args)
	} else {
		panic("FuncAuxMockPtrframerwriteLongString ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteLongString()
	return
}

//
// Mock: (recvw writeStartupFrame)String()(reta string)
//

type MockArgsTypewriteStartupFrameString struct {
	ApomockCallNumber int
}

var LastMockArgswriteStartupFrameString MockArgsTypewriteStartupFrameString

// (recvw writeStartupFrame)AuxMockString()(reta string) - Generated mock function
func (recvw writeStartupFrame) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.writeStartupFrame.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeStartupFrame.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeStartupFrame.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockwriteStartupFrameString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockwriteStartupFrameString int = 0

var condRecorderAuxMockwriteStartupFrameString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockwriteStartupFrameString(i int) {
	condRecorderAuxMockwriteStartupFrameString.L.Lock()
	for recorderAuxMockwriteStartupFrameString < i {
		condRecorderAuxMockwriteStartupFrameString.Wait()
	}
	condRecorderAuxMockwriteStartupFrameString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockwriteStartupFrameString() {
	condRecorderAuxMockwriteStartupFrameString.L.Lock()
	recorderAuxMockwriteStartupFrameString++
	condRecorderAuxMockwriteStartupFrameString.L.Unlock()
	condRecorderAuxMockwriteStartupFrameString.Broadcast()
}
func AuxMockGetRecorderAuxMockwriteStartupFrameString() (ret int) {
	condRecorderAuxMockwriteStartupFrameString.L.Lock()
	ret = recorderAuxMockwriteStartupFrameString
	condRecorderAuxMockwriteStartupFrameString.L.Unlock()
	return
}

// (recvw writeStartupFrame)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvw writeStartupFrame) String() (reta string) {
	FuncAuxMockwriteStartupFrameString, ok := apomock.GetRegisteredFunc("gocql.writeStartupFrame.String")
	if ok {
		reta = FuncAuxMockwriteStartupFrameString.(func(recvw writeStartupFrame) (reta string))(recvw)
	} else {
		panic("FuncAuxMockwriteStartupFrameString ")
	}
	AuxMockIncrementRecorderAuxMockwriteStartupFrameString()
	return
}

//
// Mock: (recvr preparedMetadata)String()(reta string)
//

type MockArgsTypepreparedMetadataString struct {
	ApomockCallNumber int
}

var LastMockArgspreparedMetadataString MockArgsTypepreparedMetadataString

// (recvr preparedMetadata)AuxMockString()(reta string) - Generated mock function
func (recvr preparedMetadata) AuxMockString() (reta string) {
	rargs, rerr := apomock.GetNext("gocql.preparedMetadata.String")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.preparedMetadata.String")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.preparedMetadata.String")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(string)
	}
	return
}

// RecorderAuxMockpreparedMetadataString  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockpreparedMetadataString int = 0

var condRecorderAuxMockpreparedMetadataString *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockpreparedMetadataString(i int) {
	condRecorderAuxMockpreparedMetadataString.L.Lock()
	for recorderAuxMockpreparedMetadataString < i {
		condRecorderAuxMockpreparedMetadataString.Wait()
	}
	condRecorderAuxMockpreparedMetadataString.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockpreparedMetadataString() {
	condRecorderAuxMockpreparedMetadataString.L.Lock()
	recorderAuxMockpreparedMetadataString++
	condRecorderAuxMockpreparedMetadataString.L.Unlock()
	condRecorderAuxMockpreparedMetadataString.Broadcast()
}
func AuxMockGetRecorderAuxMockpreparedMetadataString() (ret int) {
	condRecorderAuxMockpreparedMetadataString.L.Lock()
	ret = recorderAuxMockpreparedMetadataString
	condRecorderAuxMockpreparedMetadataString.L.Unlock()
	return
}

// (recvr preparedMetadata)String - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvr preparedMetadata) String() (reta string) {
	FuncAuxMockpreparedMetadataString, ok := apomock.GetRegisteredFunc("gocql.preparedMetadata.String")
	if ok {
		reta = FuncAuxMockpreparedMetadataString.(func(recvr preparedMetadata) (reta string))(recvr)
	} else {
		panic("FuncAuxMockpreparedMetadataString ")
	}
	AuxMockIncrementRecorderAuxMockpreparedMetadataString()
	return
}

//
// Mock: (recvf *framer)parseAuthSuccessFrame()(reta frame)
//

type MockArgsTypeframerparseAuthSuccessFrame struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseAuthSuccessFrame MockArgsTypeframerparseAuthSuccessFrame

// (recvf *framer)AuxMockparseAuthSuccessFrame()(reta frame) - Generated mock function
func (recvf *framer) AuxMockparseAuthSuccessFrame() (reta frame) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseAuthSuccessFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseAuthSuccessFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseAuthSuccessFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	return
}

// RecorderAuxMockPtrframerparseAuthSuccessFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseAuthSuccessFrame int = 0

var condRecorderAuxMockPtrframerparseAuthSuccessFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseAuthSuccessFrame(i int) {
	condRecorderAuxMockPtrframerparseAuthSuccessFrame.L.Lock()
	for recorderAuxMockPtrframerparseAuthSuccessFrame < i {
		condRecorderAuxMockPtrframerparseAuthSuccessFrame.Wait()
	}
	condRecorderAuxMockPtrframerparseAuthSuccessFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseAuthSuccessFrame() {
	condRecorderAuxMockPtrframerparseAuthSuccessFrame.L.Lock()
	recorderAuxMockPtrframerparseAuthSuccessFrame++
	condRecorderAuxMockPtrframerparseAuthSuccessFrame.L.Unlock()
	condRecorderAuxMockPtrframerparseAuthSuccessFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseAuthSuccessFrame() (ret int) {
	condRecorderAuxMockPtrframerparseAuthSuccessFrame.L.Lock()
	ret = recorderAuxMockPtrframerparseAuthSuccessFrame
	condRecorderAuxMockPtrframerparseAuthSuccessFrame.L.Unlock()
	return
}

// (recvf *framer)parseAuthSuccessFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseAuthSuccessFrame() (reta frame) {
	FuncAuxMockPtrframerparseAuthSuccessFrame, ok := apomock.GetRegisteredFunc("gocql.framer.parseAuthSuccessFrame")
	if ok {
		reta = FuncAuxMockPtrframerparseAuthSuccessFrame.(func(recvf *framer) (reta frame))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseAuthSuccessFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseAuthSuccessFrame()
	return
}

//
// Mock: (recvf *framer)readShort()(retn uint16)
//

type MockArgsTypeframerreadShort struct {
	ApomockCallNumber int
}

var LastMockArgsframerreadShort MockArgsTypeframerreadShort

// (recvf *framer)AuxMockreadShort()(retn uint16) - Generated mock function
func (recvf *framer) AuxMockreadShort() (retn uint16) {
	rargs, rerr := apomock.GetNext("gocql.framer.readShort")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.readShort")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.readShort")
	}
	if rargs.GetArg(0) != nil {
		retn = rargs.GetArg(0).(uint16)
	}
	return
}

// RecorderAuxMockPtrframerreadShort  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerreadShort int = 0

var condRecorderAuxMockPtrframerreadShort *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerreadShort(i int) {
	condRecorderAuxMockPtrframerreadShort.L.Lock()
	for recorderAuxMockPtrframerreadShort < i {
		condRecorderAuxMockPtrframerreadShort.Wait()
	}
	condRecorderAuxMockPtrframerreadShort.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerreadShort() {
	condRecorderAuxMockPtrframerreadShort.L.Lock()
	recorderAuxMockPtrframerreadShort++
	condRecorderAuxMockPtrframerreadShort.L.Unlock()
	condRecorderAuxMockPtrframerreadShort.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerreadShort() (ret int) {
	condRecorderAuxMockPtrframerreadShort.L.Lock()
	ret = recorderAuxMockPtrframerreadShort
	condRecorderAuxMockPtrframerreadShort.L.Unlock()
	return
}

// (recvf *framer)readShort - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) readShort() (retn uint16) {
	FuncAuxMockPtrframerreadShort, ok := apomock.GetRegisteredFunc("gocql.framer.readShort")
	if ok {
		retn = FuncAuxMockPtrframerreadShort.(func(recvf *framer) (retn uint16))(recvf)
	} else {
		panic("FuncAuxMockPtrframerreadShort ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerreadShort()
	return
}

//
// Mock: appendInt(argp []byte, argn int32)(reta []byte)
//

type MockArgsTypeappendInt struct {
	ApomockCallNumber int
	Argp              []byte
	Argn              int32
}

var LastMockArgsappendInt MockArgsTypeappendInt

// AuxMockappendInt(argp []byte, argn int32)(reta []byte) - Generated mock function
func AuxMockappendInt(argp []byte, argn int32) (reta []byte) {
	LastMockArgsappendInt = MockArgsTypeappendInt{
		ApomockCallNumber: AuxMockGetRecorderAuxMockappendInt(),
		Argp:              argp,
		Argn:              argn,
	}
	rargs, rerr := apomock.GetNext("gocql.appendInt")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.appendInt")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.appendInt")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]byte)
	}
	return
}

// RecorderAuxMockappendInt  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockappendInt int = 0

var condRecorderAuxMockappendInt *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockappendInt(i int) {
	condRecorderAuxMockappendInt.L.Lock()
	for recorderAuxMockappendInt < i {
		condRecorderAuxMockappendInt.Wait()
	}
	condRecorderAuxMockappendInt.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockappendInt() {
	condRecorderAuxMockappendInt.L.Lock()
	recorderAuxMockappendInt++
	condRecorderAuxMockappendInt.L.Unlock()
	condRecorderAuxMockappendInt.Broadcast()
}
func AuxMockGetRecorderAuxMockappendInt() (ret int) {
	condRecorderAuxMockappendInt.L.Lock()
	ret = recorderAuxMockappendInt
	condRecorderAuxMockappendInt.L.Unlock()
	return
}

// appendInt - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func appendInt(argp []byte, argn int32) (reta []byte) {
	FuncAuxMockappendInt, ok := apomock.GetRegisteredFunc("gocql.appendInt")
	if ok {
		reta = FuncAuxMockappendInt.(func(argp []byte, argn int32) (reta []byte))(argp, argn)
	} else {
		panic("FuncAuxMockappendInt ")
	}
	AuxMockIncrementRecorderAuxMockappendInt()
	return
}

//
// Mock: (recvf *framer)writeUUID(argu *UUID)()
//

type MockArgsTypeframerwriteUUID struct {
	ApomockCallNumber int
	Argu              *UUID
}

var LastMockArgsframerwriteUUID MockArgsTypeframerwriteUUID

// (recvf *framer)AuxMockwriteUUID(argu *UUID)() - Generated mock function
func (recvf *framer) AuxMockwriteUUID(argu *UUID) {
	LastMockArgsframerwriteUUID = MockArgsTypeframerwriteUUID{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteUUID(),
		Argu:              argu,
	}
	return
}

// RecorderAuxMockPtrframerwriteUUID  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteUUID int = 0

var condRecorderAuxMockPtrframerwriteUUID *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteUUID(i int) {
	condRecorderAuxMockPtrframerwriteUUID.L.Lock()
	for recorderAuxMockPtrframerwriteUUID < i {
		condRecorderAuxMockPtrframerwriteUUID.Wait()
	}
	condRecorderAuxMockPtrframerwriteUUID.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteUUID() {
	condRecorderAuxMockPtrframerwriteUUID.L.Lock()
	recorderAuxMockPtrframerwriteUUID++
	condRecorderAuxMockPtrframerwriteUUID.L.Unlock()
	condRecorderAuxMockPtrframerwriteUUID.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteUUID() (ret int) {
	condRecorderAuxMockPtrframerwriteUUID.L.Lock()
	ret = recorderAuxMockPtrframerwriteUUID
	condRecorderAuxMockPtrframerwriteUUID.L.Unlock()
	return
}

// (recvf *framer)writeUUID - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeUUID(argu *UUID) {
	FuncAuxMockPtrframerwriteUUID, ok := apomock.GetRegisteredFunc("gocql.framer.writeUUID")
	if ok {
		FuncAuxMockPtrframerwriteUUID.(func(recvf *framer, argu *UUID))(recvf, argu)
	} else {
		panic("FuncAuxMockPtrframerwriteUUID ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteUUID()
	return
}

//
// Mock: (recvf *framer)finishWrite()(reta error)
//

type MockArgsTypeframerfinishWrite struct {
	ApomockCallNumber int
}

var LastMockArgsframerfinishWrite MockArgsTypeframerfinishWrite

// (recvf *framer)AuxMockfinishWrite()(reta error) - Generated mock function
func (recvf *framer) AuxMockfinishWrite() (reta error) {
	rargs, rerr := apomock.GetNext("gocql.framer.finishWrite")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.finishWrite")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.finishWrite")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrframerfinishWrite  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerfinishWrite int = 0

var condRecorderAuxMockPtrframerfinishWrite *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerfinishWrite(i int) {
	condRecorderAuxMockPtrframerfinishWrite.L.Lock()
	for recorderAuxMockPtrframerfinishWrite < i {
		condRecorderAuxMockPtrframerfinishWrite.Wait()
	}
	condRecorderAuxMockPtrframerfinishWrite.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerfinishWrite() {
	condRecorderAuxMockPtrframerfinishWrite.L.Lock()
	recorderAuxMockPtrframerfinishWrite++
	condRecorderAuxMockPtrframerfinishWrite.L.Unlock()
	condRecorderAuxMockPtrframerfinishWrite.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerfinishWrite() (ret int) {
	condRecorderAuxMockPtrframerfinishWrite.L.Lock()
	ret = recorderAuxMockPtrframerfinishWrite
	condRecorderAuxMockPtrframerfinishWrite.L.Unlock()
	return
}

// (recvf *framer)finishWrite - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) finishWrite() (reta error) {
	FuncAuxMockPtrframerfinishWrite, ok := apomock.GetRegisteredFunc("gocql.framer.finishWrite")
	if ok {
		reta = FuncAuxMockPtrframerfinishWrite.(func(recvf *framer) (reta error))(recvf)
	} else {
		panic("FuncAuxMockPtrframerfinishWrite ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerfinishWrite()
	return
}

//
// Mock: (recvf *framer)writeBytes(argp []byte)()
//

type MockArgsTypeframerwriteBytes struct {
	ApomockCallNumber int
	Argp              []byte
}

var LastMockArgsframerwriteBytes MockArgsTypeframerwriteBytes

// (recvf *framer)AuxMockwriteBytes(argp []byte)() - Generated mock function
func (recvf *framer) AuxMockwriteBytes(argp []byte) {
	LastMockArgsframerwriteBytes = MockArgsTypeframerwriteBytes{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrframerwriteBytes(),
		Argp:              argp,
	}
	return
}

// RecorderAuxMockPtrframerwriteBytes  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerwriteBytes int = 0

var condRecorderAuxMockPtrframerwriteBytes *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerwriteBytes(i int) {
	condRecorderAuxMockPtrframerwriteBytes.L.Lock()
	for recorderAuxMockPtrframerwriteBytes < i {
		condRecorderAuxMockPtrframerwriteBytes.Wait()
	}
	condRecorderAuxMockPtrframerwriteBytes.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerwriteBytes() {
	condRecorderAuxMockPtrframerwriteBytes.L.Lock()
	recorderAuxMockPtrframerwriteBytes++
	condRecorderAuxMockPtrframerwriteBytes.L.Unlock()
	condRecorderAuxMockPtrframerwriteBytes.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerwriteBytes() (ret int) {
	condRecorderAuxMockPtrframerwriteBytes.L.Lock()
	ret = recorderAuxMockPtrframerwriteBytes
	condRecorderAuxMockPtrframerwriteBytes.L.Unlock()
	return
}

// (recvf *framer)writeBytes - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) writeBytes(argp []byte) {
	FuncAuxMockPtrframerwriteBytes, ok := apomock.GetRegisteredFunc("gocql.framer.writeBytes")
	if ok {
		FuncAuxMockPtrframerwriteBytes.(func(recvf *framer, argp []byte))(recvf, argp)
	} else {
		panic("FuncAuxMockPtrframerwriteBytes ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerwriteBytes()
	return
}

//
// Mock: (recvf frameWriterFunc)writeFrame(argframer *framer, argstreamID int)(reta error)
//

type MockArgsTypeframeWriterFuncwriteFrame struct {
	ApomockCallNumber int
	Argframer         *framer
	ArgstreamID       int
}

var LastMockArgsframeWriterFuncwriteFrame MockArgsTypeframeWriterFuncwriteFrame

// (recvf frameWriterFunc)AuxMockwriteFrame(argframer *framer, argstreamID int)(reta error) - Generated mock function
func (recvf frameWriterFunc) AuxMockwriteFrame(argframer *framer, argstreamID int) (reta error) {
	LastMockArgsframeWriterFuncwriteFrame = MockArgsTypeframeWriterFuncwriteFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockframeWriterFuncwriteFrame(),
		Argframer:         argframer,
		ArgstreamID:       argstreamID,
	}
	rargs, rerr := apomock.GetNext("gocql.frameWriterFunc.writeFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.frameWriterFunc.writeFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.frameWriterFunc.writeFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockframeWriterFuncwriteFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockframeWriterFuncwriteFrame int = 0

var condRecorderAuxMockframeWriterFuncwriteFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockframeWriterFuncwriteFrame(i int) {
	condRecorderAuxMockframeWriterFuncwriteFrame.L.Lock()
	for recorderAuxMockframeWriterFuncwriteFrame < i {
		condRecorderAuxMockframeWriterFuncwriteFrame.Wait()
	}
	condRecorderAuxMockframeWriterFuncwriteFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockframeWriterFuncwriteFrame() {
	condRecorderAuxMockframeWriterFuncwriteFrame.L.Lock()
	recorderAuxMockframeWriterFuncwriteFrame++
	condRecorderAuxMockframeWriterFuncwriteFrame.L.Unlock()
	condRecorderAuxMockframeWriterFuncwriteFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockframeWriterFuncwriteFrame() (ret int) {
	condRecorderAuxMockframeWriterFuncwriteFrame.L.Lock()
	ret = recorderAuxMockframeWriterFuncwriteFrame
	condRecorderAuxMockframeWriterFuncwriteFrame.L.Unlock()
	return
}

// (recvf frameWriterFunc)writeFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf frameWriterFunc) writeFrame(argframer *framer, argstreamID int) (reta error) {
	FuncAuxMockframeWriterFuncwriteFrame, ok := apomock.GetRegisteredFunc("gocql.frameWriterFunc.writeFrame")
	if ok {
		reta = FuncAuxMockframeWriterFuncwriteFrame.(func(recvf frameWriterFunc, argframer *framer, argstreamID int) (reta error))(recvf, argframer, argstreamID)
	} else {
		panic("FuncAuxMockframeWriterFuncwriteFrame ")
	}
	AuxMockIncrementRecorderAuxMockframeWriterFuncwriteFrame()
	return
}

//
// Mock: (recvf *framer)parseResultSetKeyspace()(reta frame)
//

type MockArgsTypeframerparseResultSetKeyspace struct {
	ApomockCallNumber int
}

var LastMockArgsframerparseResultSetKeyspace MockArgsTypeframerparseResultSetKeyspace

// (recvf *framer)AuxMockparseResultSetKeyspace()(reta frame) - Generated mock function
func (recvf *framer) AuxMockparseResultSetKeyspace() (reta frame) {
	rargs, rerr := apomock.GetNext("gocql.framer.parseResultSetKeyspace")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.framer.parseResultSetKeyspace")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.framer.parseResultSetKeyspace")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(frame)
	}
	return
}

// RecorderAuxMockPtrframerparseResultSetKeyspace  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrframerparseResultSetKeyspace int = 0

var condRecorderAuxMockPtrframerparseResultSetKeyspace *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrframerparseResultSetKeyspace(i int) {
	condRecorderAuxMockPtrframerparseResultSetKeyspace.L.Lock()
	for recorderAuxMockPtrframerparseResultSetKeyspace < i {
		condRecorderAuxMockPtrframerparseResultSetKeyspace.Wait()
	}
	condRecorderAuxMockPtrframerparseResultSetKeyspace.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrframerparseResultSetKeyspace() {
	condRecorderAuxMockPtrframerparseResultSetKeyspace.L.Lock()
	recorderAuxMockPtrframerparseResultSetKeyspace++
	condRecorderAuxMockPtrframerparseResultSetKeyspace.L.Unlock()
	condRecorderAuxMockPtrframerparseResultSetKeyspace.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrframerparseResultSetKeyspace() (ret int) {
	condRecorderAuxMockPtrframerparseResultSetKeyspace.L.Lock()
	ret = recorderAuxMockPtrframerparseResultSetKeyspace
	condRecorderAuxMockPtrframerparseResultSetKeyspace.L.Unlock()
	return
}

// (recvf *framer)parseResultSetKeyspace - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvf *framer) parseResultSetKeyspace() (reta frame) {
	FuncAuxMockPtrframerparseResultSetKeyspace, ok := apomock.GetRegisteredFunc("gocql.framer.parseResultSetKeyspace")
	if ok {
		reta = FuncAuxMockPtrframerparseResultSetKeyspace.(func(recvf *framer) (reta frame))(recvf)
	} else {
		panic("FuncAuxMockPtrframerparseResultSetKeyspace ")
	}
	AuxMockIncrementRecorderAuxMockPtrframerparseResultSetKeyspace()
	return
}

//
// Mock: (recvw *writeOptionsFrame)writeFrame(argframer *framer, argstreamID int)(reta error)
//

type MockArgsTypewriteOptionsFramewriteFrame struct {
	ApomockCallNumber int
	Argframer         *framer
	ArgstreamID       int
}

var LastMockArgswriteOptionsFramewriteFrame MockArgsTypewriteOptionsFramewriteFrame

// (recvw *writeOptionsFrame)AuxMockwriteFrame(argframer *framer, argstreamID int)(reta error) - Generated mock function
func (recvw *writeOptionsFrame) AuxMockwriteFrame(argframer *framer, argstreamID int) (reta error) {
	LastMockArgswriteOptionsFramewriteFrame = MockArgsTypewriteOptionsFramewriteFrame{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrwriteOptionsFramewriteFrame(),
		Argframer:         argframer,
		ArgstreamID:       argstreamID,
	}
	rargs, rerr := apomock.GetNext("gocql.writeOptionsFrame.writeFrame")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.writeOptionsFrame.writeFrame")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.writeOptionsFrame.writeFrame")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrwriteOptionsFramewriteFrame  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrwriteOptionsFramewriteFrame int = 0

var condRecorderAuxMockPtrwriteOptionsFramewriteFrame *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrwriteOptionsFramewriteFrame(i int) {
	condRecorderAuxMockPtrwriteOptionsFramewriteFrame.L.Lock()
	for recorderAuxMockPtrwriteOptionsFramewriteFrame < i {
		condRecorderAuxMockPtrwriteOptionsFramewriteFrame.Wait()
	}
	condRecorderAuxMockPtrwriteOptionsFramewriteFrame.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrwriteOptionsFramewriteFrame() {
	condRecorderAuxMockPtrwriteOptionsFramewriteFrame.L.Lock()
	recorderAuxMockPtrwriteOptionsFramewriteFrame++
	condRecorderAuxMockPtrwriteOptionsFramewriteFrame.L.Unlock()
	condRecorderAuxMockPtrwriteOptionsFramewriteFrame.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrwriteOptionsFramewriteFrame() (ret int) {
	condRecorderAuxMockPtrwriteOptionsFramewriteFrame.L.Lock()
	ret = recorderAuxMockPtrwriteOptionsFramewriteFrame
	condRecorderAuxMockPtrwriteOptionsFramewriteFrame.L.Unlock()
	return
}

// (recvw *writeOptionsFrame)writeFrame - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvw *writeOptionsFrame) writeFrame(argframer *framer, argstreamID int) (reta error) {
	FuncAuxMockPtrwriteOptionsFramewriteFrame, ok := apomock.GetRegisteredFunc("gocql.writeOptionsFrame.writeFrame")
	if ok {
		reta = FuncAuxMockPtrwriteOptionsFramewriteFrame.(func(recvw *writeOptionsFrame, argframer *framer, argstreamID int) (reta error))(recvw, argframer, argstreamID)
	} else {
		panic("FuncAuxMockPtrwriteOptionsFramewriteFrame ")
	}
	AuxMockIncrementRecorderAuxMockPtrwriteOptionsFramewriteFrame()
	return
}
