// -----------------------------------------------------------------
// 2016 (c) Aporeto Inc.
// Auto-Generated Aporeto Mock
// DO NOT HAND EDIT !
// -----------------------------------------------------------------
package gocql

import "sync"

import "github.com/aporeto-inc/kennebec/apomock"

func init() {
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructSchemaDescriber, apomockNewStructSchemaDescriber)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructTypeParser, apomockNewStructTypeParser)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructTypeParserResult, apomockNewStructTypeParserResult)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructTypeParserClassNode, apomockNewStructTypeParserClassNode)
	apomock.RegisterInternalType("github.com/gocql/gocql", ApomockStructTypeParserParamNode, apomockNewStructTypeParserParamNode)

	apomock.RegisterFunc("gocql", "gocql.schemaDescriber.getSchema", (*schemaDescriber).AuxMockgetSchema)
	apomock.RegisterFunc("gocql", "gocql.compileMetadata", AuxMockcompileMetadata)
	apomock.RegisterFunc("gocql", "gocql.parseType", AuxMockparseType)
	apomock.RegisterFunc("gocql", "gocql.typeParser.parse", (*typeParser).AuxMockparse)
	apomock.RegisterFunc("gocql", "gocql.typeParserClassNode.asTypeInfo", (*typeParserClassNode).AuxMockasTypeInfo)
	apomock.RegisterFunc("gocql", "gocql.typeParser.parseClassNode", (*typeParser).AuxMockparseClassNode)
	apomock.RegisterFunc("gocql", "gocql.schemaDescriber.clearSchema", (*schemaDescriber).AuxMockclearSchema)
	apomock.RegisterFunc("gocql", "gocql.compileV1Metadata", AuxMockcompileV1Metadata)
	apomock.RegisterFunc("gocql", "gocql.compileV2Metadata", AuxMockcompileV2Metadata)
	apomock.RegisterFunc("gocql", "gocql.getTableMetadata", AuxMockgetTableMetadata)
	apomock.RegisterFunc("gocql", "gocql.getColumnMetadata", AuxMockgetColumnMetadata)
	apomock.RegisterFunc("gocql", "gocql.schemaDescriber.refreshSchema", (*schemaDescriber).AuxMockrefreshSchema)
	apomock.RegisterFunc("gocql", "gocql.componentColumnCountOfType", AuxMockcomponentColumnCountOfType)
	apomock.RegisterFunc("gocql", "gocql.newSchemaDescriber", AuxMocknewSchemaDescriber)
	apomock.RegisterFunc("gocql", "gocql.getKeyspaceMetadata", AuxMockgetKeyspaceMetadata)
	apomock.RegisterFunc("gocql", "gocql.typeParser.parseParamNodes", (*typeParser).AuxMockparseParamNodes)
	apomock.RegisterFunc("gocql", "gocql.typeParser.skipWhitespace", (*typeParser).AuxMockskipWhitespace)
	apomock.RegisterFunc("gocql", "gocql.isWhitespaceChar", AuxMockisWhitespaceChar)
	apomock.RegisterFunc("gocql", "gocql.typeParser.nextIdentifier", (*typeParser).AuxMocknextIdentifier)
	apomock.RegisterFunc("gocql", "gocql.isIdentifierChar", AuxMockisIdentifierChar)
}

const (
	ASC  ColumnOrder = false
	DESC             = true
)

const (
	PARTITION_KEY  = "partition_key"
	CLUSTERING_KEY = "clustering_key"
	REGULAR        = "regular"
	COMPACT_VALUE  = "compact_value"
)

const (
	DEFAULT_KEY_ALIAS    = "key"
	DEFAULT_COLUMN_ALIAS = "column"
	DEFAULT_VALUE_ALIAS  = "value"
)

const (
	REVERSED_TYPE   = "org.apache.cassandra.db.marshal.ReversedType"
	COMPOSITE_TYPE  = "org.apache.cassandra.db.marshal.CompositeType"
	COLLECTION_TYPE = "org.apache.cassandra.db.marshal.ColumnToCollectionType"
	LIST_TYPE       = "org.apache.cassandra.db.marshal.ListType"
	SET_TYPE        = "org.apache.cassandra.db.marshal.SetType"
	MAP_TYPE        = "org.apache.cassandra.db.marshal.MapType"
)

const (
	ApomockStructSchemaDescriber     = 54
	ApomockStructTypeParser          = 55
	ApomockStructTypeParserResult    = 56
	ApomockStructTypeParserClassNode = 57
	ApomockStructTypeParserParamNode = 58
)

//
// Internal Types: in this package and their exportable versions
//
type schemaDescriber struct {
	session *Session
	mu      sync.Mutex
	cache   map[string]*KeyspaceMetadata
}
type typeParser struct {
	input string
	index int
}
type typeParserResult struct {
	isComposite bool
	types       []TypeInfo
	reversed    []bool
	collections map[string]TypeInfo
}
type typeParserClassNode struct {
	name   string
	params []typeParserParamNode
	input  string
}
type typeParserParamNode struct {
	name  *string
	class typeParserClassNode
}

//
// External Types: in this package
//
type TableMetadata struct {
	Keyspace          string
	Name              string
	KeyValidator      string
	Comparator        string
	DefaultValidator  string
	KeyAliases        []string
	ColumnAliases     []string
	ValueAlias        string
	PartitionKey      []*ColumnMetadata
	ClusteringColumns []*ColumnMetadata
	Columns           map[string]*ColumnMetadata
	OrderedColumns    []string
}

type ColumnMetadata struct {
	Keyspace        string
	Table           string
	Name            string
	ComponentIndex  int
	Kind            string
	Validator       string
	Type            TypeInfo
	ClusteringOrder string
	Order           ColumnOrder
	Index           ColumnIndexMetadata
}

type ColumnOrder bool

type KeyspaceMetadata struct {
	Name            string
	DurableWrites   bool
	StrategyClass   string
	StrategyOptions map[string]interface{}
	Tables          map[string]*TableMetadata
}

type ColumnIndexMetadata struct {
	Name    string
	Type    string
	Options map[string]interface{}
}

func apomockNewStructSchemaDescriber() interface{}     { return &schemaDescriber{} }
func apomockNewStructTypeParser() interface{}          { return &typeParser{} }
func apomockNewStructTypeParserResult() interface{}    { return &typeParserResult{} }
func apomockNewStructTypeParserClassNode() interface{} { return &typeParserClassNode{} }
func apomockNewStructTypeParserParamNode() interface{} { return &typeParserParamNode{} }

//
// Mock: (recvs *schemaDescriber)getSchema(argkeyspaceName string)(reta *KeyspaceMetadata, retb error)
//

type MockArgsTypeschemaDescribergetSchema struct {
	ApomockCallNumber int
	ArgkeyspaceName   string
}

var LastMockArgsschemaDescribergetSchema MockArgsTypeschemaDescribergetSchema

// (recvs *schemaDescriber)AuxMockgetSchema(argkeyspaceName string)(reta *KeyspaceMetadata, retb error) - Generated mock function
func (recvs *schemaDescriber) AuxMockgetSchema(argkeyspaceName string) (reta *KeyspaceMetadata, retb error) {
	LastMockArgsschemaDescribergetSchema = MockArgsTypeschemaDescribergetSchema{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrschemaDescribergetSchema(),
		ArgkeyspaceName:   argkeyspaceName,
	}
	rargs, rerr := apomock.GetNext("gocql.schemaDescriber.getSchema")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.schemaDescriber.getSchema")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.schemaDescriber.getSchema")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*KeyspaceMetadata)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockPtrschemaDescribergetSchema  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrschemaDescribergetSchema int = 0

var condRecorderAuxMockPtrschemaDescribergetSchema *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrschemaDescribergetSchema(i int) {
	condRecorderAuxMockPtrschemaDescribergetSchema.L.Lock()
	for recorderAuxMockPtrschemaDescribergetSchema < i {
		condRecorderAuxMockPtrschemaDescribergetSchema.Wait()
	}
	condRecorderAuxMockPtrschemaDescribergetSchema.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrschemaDescribergetSchema() {
	condRecorderAuxMockPtrschemaDescribergetSchema.L.Lock()
	recorderAuxMockPtrschemaDescribergetSchema++
	condRecorderAuxMockPtrschemaDescribergetSchema.L.Unlock()
	condRecorderAuxMockPtrschemaDescribergetSchema.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrschemaDescribergetSchema() (ret int) {
	condRecorderAuxMockPtrschemaDescribergetSchema.L.Lock()
	ret = recorderAuxMockPtrschemaDescribergetSchema
	condRecorderAuxMockPtrschemaDescribergetSchema.L.Unlock()
	return
}

// (recvs *schemaDescriber)getSchema - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *schemaDescriber) getSchema(argkeyspaceName string) (reta *KeyspaceMetadata, retb error) {
	FuncAuxMockPtrschemaDescribergetSchema, ok := apomock.GetRegisteredFunc("gocql.schemaDescriber.getSchema")
	if ok {
		reta, retb = FuncAuxMockPtrschemaDescribergetSchema.(func(recvs *schemaDescriber, argkeyspaceName string) (reta *KeyspaceMetadata, retb error))(recvs, argkeyspaceName)
	} else {
		panic("FuncAuxMockPtrschemaDescribergetSchema ")
	}
	AuxMockIncrementRecorderAuxMockPtrschemaDescribergetSchema()
	return
}

//
// Mock: compileMetadata(argprotoVersion int, argkeyspace *KeyspaceMetadata, argtables []TableMetadata, argcolumns []ColumnMetadata)()
//

type MockArgsTypecompileMetadata struct {
	ApomockCallNumber int
	ArgprotoVersion   int
	Argkeyspace       *KeyspaceMetadata
	Argtables         []TableMetadata
	Argcolumns        []ColumnMetadata
}

var LastMockArgscompileMetadata MockArgsTypecompileMetadata

// AuxMockcompileMetadata(argprotoVersion int, argkeyspace *KeyspaceMetadata, argtables []TableMetadata, argcolumns []ColumnMetadata)() - Generated mock function
func AuxMockcompileMetadata(argprotoVersion int, argkeyspace *KeyspaceMetadata, argtables []TableMetadata, argcolumns []ColumnMetadata) {
	LastMockArgscompileMetadata = MockArgsTypecompileMetadata{
		ApomockCallNumber: AuxMockGetRecorderAuxMockcompileMetadata(),
		ArgprotoVersion:   argprotoVersion,
		Argkeyspace:       argkeyspace,
		Argtables:         argtables,
		Argcolumns:        argcolumns,
	}
	return
}

// RecorderAuxMockcompileMetadata  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockcompileMetadata int = 0

var condRecorderAuxMockcompileMetadata *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockcompileMetadata(i int) {
	condRecorderAuxMockcompileMetadata.L.Lock()
	for recorderAuxMockcompileMetadata < i {
		condRecorderAuxMockcompileMetadata.Wait()
	}
	condRecorderAuxMockcompileMetadata.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockcompileMetadata() {
	condRecorderAuxMockcompileMetadata.L.Lock()
	recorderAuxMockcompileMetadata++
	condRecorderAuxMockcompileMetadata.L.Unlock()
	condRecorderAuxMockcompileMetadata.Broadcast()
}
func AuxMockGetRecorderAuxMockcompileMetadata() (ret int) {
	condRecorderAuxMockcompileMetadata.L.Lock()
	ret = recorderAuxMockcompileMetadata
	condRecorderAuxMockcompileMetadata.L.Unlock()
	return
}

// compileMetadata - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func compileMetadata(argprotoVersion int, argkeyspace *KeyspaceMetadata, argtables []TableMetadata, argcolumns []ColumnMetadata) {
	FuncAuxMockcompileMetadata, ok := apomock.GetRegisteredFunc("gocql.compileMetadata")
	if ok {
		FuncAuxMockcompileMetadata.(func(argprotoVersion int, argkeyspace *KeyspaceMetadata, argtables []TableMetadata, argcolumns []ColumnMetadata))(argprotoVersion, argkeyspace, argtables, argcolumns)
	} else {
		panic("FuncAuxMockcompileMetadata ")
	}
	AuxMockIncrementRecorderAuxMockcompileMetadata()
	return
}

//
// Mock: parseType(argdef string)(reta typeParserResult)
//

type MockArgsTypeparseType struct {
	ApomockCallNumber int
	Argdef            string
}

var LastMockArgsparseType MockArgsTypeparseType

// AuxMockparseType(argdef string)(reta typeParserResult) - Generated mock function
func AuxMockparseType(argdef string) (reta typeParserResult) {
	LastMockArgsparseType = MockArgsTypeparseType{
		ApomockCallNumber: AuxMockGetRecorderAuxMockparseType(),
		Argdef:            argdef,
	}
	rargs, rerr := apomock.GetNext("gocql.parseType")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.parseType")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.parseType")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(typeParserResult)
	}
	return
}

// RecorderAuxMockparseType  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockparseType int = 0

var condRecorderAuxMockparseType *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockparseType(i int) {
	condRecorderAuxMockparseType.L.Lock()
	for recorderAuxMockparseType < i {
		condRecorderAuxMockparseType.Wait()
	}
	condRecorderAuxMockparseType.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockparseType() {
	condRecorderAuxMockparseType.L.Lock()
	recorderAuxMockparseType++
	condRecorderAuxMockparseType.L.Unlock()
	condRecorderAuxMockparseType.Broadcast()
}
func AuxMockGetRecorderAuxMockparseType() (ret int) {
	condRecorderAuxMockparseType.L.Lock()
	ret = recorderAuxMockparseType
	condRecorderAuxMockparseType.L.Unlock()
	return
}

// parseType - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func parseType(argdef string) (reta typeParserResult) {
	FuncAuxMockparseType, ok := apomock.GetRegisteredFunc("gocql.parseType")
	if ok {
		reta = FuncAuxMockparseType.(func(argdef string) (reta typeParserResult))(argdef)
	} else {
		panic("FuncAuxMockparseType ")
	}
	AuxMockIncrementRecorderAuxMockparseType()
	return
}

//
// Mock: (recvt *typeParser)parse()(reta typeParserResult)
//

type MockArgsTypetypeParserparse struct {
	ApomockCallNumber int
}

var LastMockArgstypeParserparse MockArgsTypetypeParserparse

// (recvt *typeParser)AuxMockparse()(reta typeParserResult) - Generated mock function
func (recvt *typeParser) AuxMockparse() (reta typeParserResult) {
	rargs, rerr := apomock.GetNext("gocql.typeParser.parse")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.typeParser.parse")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.typeParser.parse")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(typeParserResult)
	}
	return
}

// RecorderAuxMockPtrtypeParserparse  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtypeParserparse int = 0

var condRecorderAuxMockPtrtypeParserparse *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtypeParserparse(i int) {
	condRecorderAuxMockPtrtypeParserparse.L.Lock()
	for recorderAuxMockPtrtypeParserparse < i {
		condRecorderAuxMockPtrtypeParserparse.Wait()
	}
	condRecorderAuxMockPtrtypeParserparse.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtypeParserparse() {
	condRecorderAuxMockPtrtypeParserparse.L.Lock()
	recorderAuxMockPtrtypeParserparse++
	condRecorderAuxMockPtrtypeParserparse.L.Unlock()
	condRecorderAuxMockPtrtypeParserparse.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtypeParserparse() (ret int) {
	condRecorderAuxMockPtrtypeParserparse.L.Lock()
	ret = recorderAuxMockPtrtypeParserparse
	condRecorderAuxMockPtrtypeParserparse.L.Unlock()
	return
}

// (recvt *typeParser)parse - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *typeParser) parse() (reta typeParserResult) {
	FuncAuxMockPtrtypeParserparse, ok := apomock.GetRegisteredFunc("gocql.typeParser.parse")
	if ok {
		reta = FuncAuxMockPtrtypeParserparse.(func(recvt *typeParser) (reta typeParserResult))(recvt)
	} else {
		panic("FuncAuxMockPtrtypeParserparse ")
	}
	AuxMockIncrementRecorderAuxMockPtrtypeParserparse()
	return
}

//
// Mock: (recvclass *typeParserClassNode)asTypeInfo()(reta TypeInfo)
//

type MockArgsTypetypeParserClassNodeasTypeInfo struct {
	ApomockCallNumber int
}

var LastMockArgstypeParserClassNodeasTypeInfo MockArgsTypetypeParserClassNodeasTypeInfo

// (recvclass *typeParserClassNode)AuxMockasTypeInfo()(reta TypeInfo) - Generated mock function
func (recvclass *typeParserClassNode) AuxMockasTypeInfo() (reta TypeInfo) {
	rargs, rerr := apomock.GetNext("gocql.typeParserClassNode.asTypeInfo")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.typeParserClassNode.asTypeInfo")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.typeParserClassNode.asTypeInfo")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(TypeInfo)
	}
	return
}

// RecorderAuxMockPtrtypeParserClassNodeasTypeInfo  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtypeParserClassNodeasTypeInfo int = 0

var condRecorderAuxMockPtrtypeParserClassNodeasTypeInfo *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtypeParserClassNodeasTypeInfo(i int) {
	condRecorderAuxMockPtrtypeParserClassNodeasTypeInfo.L.Lock()
	for recorderAuxMockPtrtypeParserClassNodeasTypeInfo < i {
		condRecorderAuxMockPtrtypeParserClassNodeasTypeInfo.Wait()
	}
	condRecorderAuxMockPtrtypeParserClassNodeasTypeInfo.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtypeParserClassNodeasTypeInfo() {
	condRecorderAuxMockPtrtypeParserClassNodeasTypeInfo.L.Lock()
	recorderAuxMockPtrtypeParserClassNodeasTypeInfo++
	condRecorderAuxMockPtrtypeParserClassNodeasTypeInfo.L.Unlock()
	condRecorderAuxMockPtrtypeParserClassNodeasTypeInfo.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtypeParserClassNodeasTypeInfo() (ret int) {
	condRecorderAuxMockPtrtypeParserClassNodeasTypeInfo.L.Lock()
	ret = recorderAuxMockPtrtypeParserClassNodeasTypeInfo
	condRecorderAuxMockPtrtypeParserClassNodeasTypeInfo.L.Unlock()
	return
}

// (recvclass *typeParserClassNode)asTypeInfo - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvclass *typeParserClassNode) asTypeInfo() (reta TypeInfo) {
	FuncAuxMockPtrtypeParserClassNodeasTypeInfo, ok := apomock.GetRegisteredFunc("gocql.typeParserClassNode.asTypeInfo")
	if ok {
		reta = FuncAuxMockPtrtypeParserClassNodeasTypeInfo.(func(recvclass *typeParserClassNode) (reta TypeInfo))(recvclass)
	} else {
		panic("FuncAuxMockPtrtypeParserClassNodeasTypeInfo ")
	}
	AuxMockIncrementRecorderAuxMockPtrtypeParserClassNodeasTypeInfo()
	return
}

//
// Mock: (recvt *typeParser)parseClassNode()(retnode *typeParserClassNode, retok bool)
//

type MockArgsTypetypeParserparseClassNode struct {
	ApomockCallNumber int
}

var LastMockArgstypeParserparseClassNode MockArgsTypetypeParserparseClassNode

// (recvt *typeParser)AuxMockparseClassNode()(retnode *typeParserClassNode, retok bool) - Generated mock function
func (recvt *typeParser) AuxMockparseClassNode() (retnode *typeParserClassNode, retok bool) {
	rargs, rerr := apomock.GetNext("gocql.typeParser.parseClassNode")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.typeParser.parseClassNode")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.typeParser.parseClassNode")
	}
	if rargs.GetArg(0) != nil {
		retnode = rargs.GetArg(0).(*typeParserClassNode)
	}
	if rargs.GetArg(1) != nil {
		retok = rargs.GetArg(1).(bool)
	}
	return
}

// RecorderAuxMockPtrtypeParserparseClassNode  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtypeParserparseClassNode int = 0

var condRecorderAuxMockPtrtypeParserparseClassNode *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtypeParserparseClassNode(i int) {
	condRecorderAuxMockPtrtypeParserparseClassNode.L.Lock()
	for recorderAuxMockPtrtypeParserparseClassNode < i {
		condRecorderAuxMockPtrtypeParserparseClassNode.Wait()
	}
	condRecorderAuxMockPtrtypeParserparseClassNode.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtypeParserparseClassNode() {
	condRecorderAuxMockPtrtypeParserparseClassNode.L.Lock()
	recorderAuxMockPtrtypeParserparseClassNode++
	condRecorderAuxMockPtrtypeParserparseClassNode.L.Unlock()
	condRecorderAuxMockPtrtypeParserparseClassNode.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtypeParserparseClassNode() (ret int) {
	condRecorderAuxMockPtrtypeParserparseClassNode.L.Lock()
	ret = recorderAuxMockPtrtypeParserparseClassNode
	condRecorderAuxMockPtrtypeParserparseClassNode.L.Unlock()
	return
}

// (recvt *typeParser)parseClassNode - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *typeParser) parseClassNode() (retnode *typeParserClassNode, retok bool) {
	FuncAuxMockPtrtypeParserparseClassNode, ok := apomock.GetRegisteredFunc("gocql.typeParser.parseClassNode")
	if ok {
		retnode, retok = FuncAuxMockPtrtypeParserparseClassNode.(func(recvt *typeParser) (retnode *typeParserClassNode, retok bool))(recvt)
	} else {
		panic("FuncAuxMockPtrtypeParserparseClassNode ")
	}
	AuxMockIncrementRecorderAuxMockPtrtypeParserparseClassNode()
	return
}

//
// Mock: (recvs *schemaDescriber)clearSchema(argkeyspaceName string)()
//

type MockArgsTypeschemaDescriberclearSchema struct {
	ApomockCallNumber int
	ArgkeyspaceName   string
}

var LastMockArgsschemaDescriberclearSchema MockArgsTypeschemaDescriberclearSchema

// (recvs *schemaDescriber)AuxMockclearSchema(argkeyspaceName string)() - Generated mock function
func (recvs *schemaDescriber) AuxMockclearSchema(argkeyspaceName string) {
	LastMockArgsschemaDescriberclearSchema = MockArgsTypeschemaDescriberclearSchema{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrschemaDescriberclearSchema(),
		ArgkeyspaceName:   argkeyspaceName,
	}
	return
}

// RecorderAuxMockPtrschemaDescriberclearSchema  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrschemaDescriberclearSchema int = 0

var condRecorderAuxMockPtrschemaDescriberclearSchema *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrschemaDescriberclearSchema(i int) {
	condRecorderAuxMockPtrschemaDescriberclearSchema.L.Lock()
	for recorderAuxMockPtrschemaDescriberclearSchema < i {
		condRecorderAuxMockPtrschemaDescriberclearSchema.Wait()
	}
	condRecorderAuxMockPtrschemaDescriberclearSchema.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrschemaDescriberclearSchema() {
	condRecorderAuxMockPtrschemaDescriberclearSchema.L.Lock()
	recorderAuxMockPtrschemaDescriberclearSchema++
	condRecorderAuxMockPtrschemaDescriberclearSchema.L.Unlock()
	condRecorderAuxMockPtrschemaDescriberclearSchema.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrschemaDescriberclearSchema() (ret int) {
	condRecorderAuxMockPtrschemaDescriberclearSchema.L.Lock()
	ret = recorderAuxMockPtrschemaDescriberclearSchema
	condRecorderAuxMockPtrschemaDescriberclearSchema.L.Unlock()
	return
}

// (recvs *schemaDescriber)clearSchema - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *schemaDescriber) clearSchema(argkeyspaceName string) {
	FuncAuxMockPtrschemaDescriberclearSchema, ok := apomock.GetRegisteredFunc("gocql.schemaDescriber.clearSchema")
	if ok {
		FuncAuxMockPtrschemaDescriberclearSchema.(func(recvs *schemaDescriber, argkeyspaceName string))(recvs, argkeyspaceName)
	} else {
		panic("FuncAuxMockPtrschemaDescriberclearSchema ")
	}
	AuxMockIncrementRecorderAuxMockPtrschemaDescriberclearSchema()
	return
}

//
// Mock: compileV1Metadata(argtables []TableMetadata)()
//

type MockArgsTypecompileV1Metadata struct {
	ApomockCallNumber int
	Argtables         []TableMetadata
}

var LastMockArgscompileV1Metadata MockArgsTypecompileV1Metadata

// AuxMockcompileV1Metadata(argtables []TableMetadata)() - Generated mock function
func AuxMockcompileV1Metadata(argtables []TableMetadata) {
	LastMockArgscompileV1Metadata = MockArgsTypecompileV1Metadata{
		ApomockCallNumber: AuxMockGetRecorderAuxMockcompileV1Metadata(),
		Argtables:         argtables,
	}
	return
}

// RecorderAuxMockcompileV1Metadata  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockcompileV1Metadata int = 0

var condRecorderAuxMockcompileV1Metadata *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockcompileV1Metadata(i int) {
	condRecorderAuxMockcompileV1Metadata.L.Lock()
	for recorderAuxMockcompileV1Metadata < i {
		condRecorderAuxMockcompileV1Metadata.Wait()
	}
	condRecorderAuxMockcompileV1Metadata.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockcompileV1Metadata() {
	condRecorderAuxMockcompileV1Metadata.L.Lock()
	recorderAuxMockcompileV1Metadata++
	condRecorderAuxMockcompileV1Metadata.L.Unlock()
	condRecorderAuxMockcompileV1Metadata.Broadcast()
}
func AuxMockGetRecorderAuxMockcompileV1Metadata() (ret int) {
	condRecorderAuxMockcompileV1Metadata.L.Lock()
	ret = recorderAuxMockcompileV1Metadata
	condRecorderAuxMockcompileV1Metadata.L.Unlock()
	return
}

// compileV1Metadata - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func compileV1Metadata(argtables []TableMetadata) {
	FuncAuxMockcompileV1Metadata, ok := apomock.GetRegisteredFunc("gocql.compileV1Metadata")
	if ok {
		FuncAuxMockcompileV1Metadata.(func(argtables []TableMetadata))(argtables)
	} else {
		panic("FuncAuxMockcompileV1Metadata ")
	}
	AuxMockIncrementRecorderAuxMockcompileV1Metadata()
	return
}

//
// Mock: compileV2Metadata(argtables []TableMetadata)()
//

type MockArgsTypecompileV2Metadata struct {
	ApomockCallNumber int
	Argtables         []TableMetadata
}

var LastMockArgscompileV2Metadata MockArgsTypecompileV2Metadata

// AuxMockcompileV2Metadata(argtables []TableMetadata)() - Generated mock function
func AuxMockcompileV2Metadata(argtables []TableMetadata) {
	LastMockArgscompileV2Metadata = MockArgsTypecompileV2Metadata{
		ApomockCallNumber: AuxMockGetRecorderAuxMockcompileV2Metadata(),
		Argtables:         argtables,
	}
	return
}

// RecorderAuxMockcompileV2Metadata  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockcompileV2Metadata int = 0

var condRecorderAuxMockcompileV2Metadata *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockcompileV2Metadata(i int) {
	condRecorderAuxMockcompileV2Metadata.L.Lock()
	for recorderAuxMockcompileV2Metadata < i {
		condRecorderAuxMockcompileV2Metadata.Wait()
	}
	condRecorderAuxMockcompileV2Metadata.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockcompileV2Metadata() {
	condRecorderAuxMockcompileV2Metadata.L.Lock()
	recorderAuxMockcompileV2Metadata++
	condRecorderAuxMockcompileV2Metadata.L.Unlock()
	condRecorderAuxMockcompileV2Metadata.Broadcast()
}
func AuxMockGetRecorderAuxMockcompileV2Metadata() (ret int) {
	condRecorderAuxMockcompileV2Metadata.L.Lock()
	ret = recorderAuxMockcompileV2Metadata
	condRecorderAuxMockcompileV2Metadata.L.Unlock()
	return
}

// compileV2Metadata - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func compileV2Metadata(argtables []TableMetadata) {
	FuncAuxMockcompileV2Metadata, ok := apomock.GetRegisteredFunc("gocql.compileV2Metadata")
	if ok {
		FuncAuxMockcompileV2Metadata.(func(argtables []TableMetadata))(argtables)
	} else {
		panic("FuncAuxMockcompileV2Metadata ")
	}
	AuxMockIncrementRecorderAuxMockcompileV2Metadata()
	return
}

//
// Mock: getTableMetadata(argsession *Session, argkeyspaceName string)(reta []TableMetadata, retb error)
//

type MockArgsTypegetTableMetadata struct {
	ApomockCallNumber int
	Argsession        *Session
	ArgkeyspaceName   string
}

var LastMockArgsgetTableMetadata MockArgsTypegetTableMetadata

// AuxMockgetTableMetadata(argsession *Session, argkeyspaceName string)(reta []TableMetadata, retb error) - Generated mock function
func AuxMockgetTableMetadata(argsession *Session, argkeyspaceName string) (reta []TableMetadata, retb error) {
	LastMockArgsgetTableMetadata = MockArgsTypegetTableMetadata{
		ApomockCallNumber: AuxMockGetRecorderAuxMockgetTableMetadata(),
		Argsession:        argsession,
		ArgkeyspaceName:   argkeyspaceName,
	}
	rargs, rerr := apomock.GetNext("gocql.getTableMetadata")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.getTableMetadata")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.getTableMetadata")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]TableMetadata)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockgetTableMetadata  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockgetTableMetadata int = 0

var condRecorderAuxMockgetTableMetadata *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockgetTableMetadata(i int) {
	condRecorderAuxMockgetTableMetadata.L.Lock()
	for recorderAuxMockgetTableMetadata < i {
		condRecorderAuxMockgetTableMetadata.Wait()
	}
	condRecorderAuxMockgetTableMetadata.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockgetTableMetadata() {
	condRecorderAuxMockgetTableMetadata.L.Lock()
	recorderAuxMockgetTableMetadata++
	condRecorderAuxMockgetTableMetadata.L.Unlock()
	condRecorderAuxMockgetTableMetadata.Broadcast()
}
func AuxMockGetRecorderAuxMockgetTableMetadata() (ret int) {
	condRecorderAuxMockgetTableMetadata.L.Lock()
	ret = recorderAuxMockgetTableMetadata
	condRecorderAuxMockgetTableMetadata.L.Unlock()
	return
}

// getTableMetadata - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func getTableMetadata(argsession *Session, argkeyspaceName string) (reta []TableMetadata, retb error) {
	FuncAuxMockgetTableMetadata, ok := apomock.GetRegisteredFunc("gocql.getTableMetadata")
	if ok {
		reta, retb = FuncAuxMockgetTableMetadata.(func(argsession *Session, argkeyspaceName string) (reta []TableMetadata, retb error))(argsession, argkeyspaceName)
	} else {
		panic("FuncAuxMockgetTableMetadata ")
	}
	AuxMockIncrementRecorderAuxMockgetTableMetadata()
	return
}

//
// Mock: getColumnMetadata(argsession *Session, argkeyspaceName string)(reta []ColumnMetadata, retb error)
//

type MockArgsTypegetColumnMetadata struct {
	ApomockCallNumber int
	Argsession        *Session
	ArgkeyspaceName   string
}

var LastMockArgsgetColumnMetadata MockArgsTypegetColumnMetadata

// AuxMockgetColumnMetadata(argsession *Session, argkeyspaceName string)(reta []ColumnMetadata, retb error) - Generated mock function
func AuxMockgetColumnMetadata(argsession *Session, argkeyspaceName string) (reta []ColumnMetadata, retb error) {
	LastMockArgsgetColumnMetadata = MockArgsTypegetColumnMetadata{
		ApomockCallNumber: AuxMockGetRecorderAuxMockgetColumnMetadata(),
		Argsession:        argsession,
		ArgkeyspaceName:   argkeyspaceName,
	}
	rargs, rerr := apomock.GetNext("gocql.getColumnMetadata")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.getColumnMetadata")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.getColumnMetadata")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).([]ColumnMetadata)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockgetColumnMetadata  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockgetColumnMetadata int = 0

var condRecorderAuxMockgetColumnMetadata *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockgetColumnMetadata(i int) {
	condRecorderAuxMockgetColumnMetadata.L.Lock()
	for recorderAuxMockgetColumnMetadata < i {
		condRecorderAuxMockgetColumnMetadata.Wait()
	}
	condRecorderAuxMockgetColumnMetadata.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockgetColumnMetadata() {
	condRecorderAuxMockgetColumnMetadata.L.Lock()
	recorderAuxMockgetColumnMetadata++
	condRecorderAuxMockgetColumnMetadata.L.Unlock()
	condRecorderAuxMockgetColumnMetadata.Broadcast()
}
func AuxMockGetRecorderAuxMockgetColumnMetadata() (ret int) {
	condRecorderAuxMockgetColumnMetadata.L.Lock()
	ret = recorderAuxMockgetColumnMetadata
	condRecorderAuxMockgetColumnMetadata.L.Unlock()
	return
}

// getColumnMetadata - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func getColumnMetadata(argsession *Session, argkeyspaceName string) (reta []ColumnMetadata, retb error) {
	FuncAuxMockgetColumnMetadata, ok := apomock.GetRegisteredFunc("gocql.getColumnMetadata")
	if ok {
		reta, retb = FuncAuxMockgetColumnMetadata.(func(argsession *Session, argkeyspaceName string) (reta []ColumnMetadata, retb error))(argsession, argkeyspaceName)
	} else {
		panic("FuncAuxMockgetColumnMetadata ")
	}
	AuxMockIncrementRecorderAuxMockgetColumnMetadata()
	return
}

//
// Mock: (recvs *schemaDescriber)refreshSchema(argkeyspaceName string)(reta error)
//

type MockArgsTypeschemaDescriberrefreshSchema struct {
	ApomockCallNumber int
	ArgkeyspaceName   string
}

var LastMockArgsschemaDescriberrefreshSchema MockArgsTypeschemaDescriberrefreshSchema

// (recvs *schemaDescriber)AuxMockrefreshSchema(argkeyspaceName string)(reta error) - Generated mock function
func (recvs *schemaDescriber) AuxMockrefreshSchema(argkeyspaceName string) (reta error) {
	LastMockArgsschemaDescriberrefreshSchema = MockArgsTypeschemaDescriberrefreshSchema{
		ApomockCallNumber: AuxMockGetRecorderAuxMockPtrschemaDescriberrefreshSchema(),
		ArgkeyspaceName:   argkeyspaceName,
	}
	rargs, rerr := apomock.GetNext("gocql.schemaDescriber.refreshSchema")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.schemaDescriber.refreshSchema")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.schemaDescriber.refreshSchema")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(error)
	}
	return
}

// RecorderAuxMockPtrschemaDescriberrefreshSchema  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrschemaDescriberrefreshSchema int = 0

var condRecorderAuxMockPtrschemaDescriberrefreshSchema *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrschemaDescriberrefreshSchema(i int) {
	condRecorderAuxMockPtrschemaDescriberrefreshSchema.L.Lock()
	for recorderAuxMockPtrschemaDescriberrefreshSchema < i {
		condRecorderAuxMockPtrschemaDescriberrefreshSchema.Wait()
	}
	condRecorderAuxMockPtrschemaDescriberrefreshSchema.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrschemaDescriberrefreshSchema() {
	condRecorderAuxMockPtrschemaDescriberrefreshSchema.L.Lock()
	recorderAuxMockPtrschemaDescriberrefreshSchema++
	condRecorderAuxMockPtrschemaDescriberrefreshSchema.L.Unlock()
	condRecorderAuxMockPtrschemaDescriberrefreshSchema.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrschemaDescriberrefreshSchema() (ret int) {
	condRecorderAuxMockPtrschemaDescriberrefreshSchema.L.Lock()
	ret = recorderAuxMockPtrschemaDescriberrefreshSchema
	condRecorderAuxMockPtrschemaDescriberrefreshSchema.L.Unlock()
	return
}

// (recvs *schemaDescriber)refreshSchema - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvs *schemaDescriber) refreshSchema(argkeyspaceName string) (reta error) {
	FuncAuxMockPtrschemaDescriberrefreshSchema, ok := apomock.GetRegisteredFunc("gocql.schemaDescriber.refreshSchema")
	if ok {
		reta = FuncAuxMockPtrschemaDescriberrefreshSchema.(func(recvs *schemaDescriber, argkeyspaceName string) (reta error))(recvs, argkeyspaceName)
	} else {
		panic("FuncAuxMockPtrschemaDescriberrefreshSchema ")
	}
	AuxMockIncrementRecorderAuxMockPtrschemaDescriberrefreshSchema()
	return
}

//
// Mock: componentColumnCountOfType(argcolumns map[string]*ColumnMetadata, argkind string)(reta int)
//

type MockArgsTypecomponentColumnCountOfType struct {
	ApomockCallNumber int
	Argcolumns        map[string]*ColumnMetadata
	Argkind           string
}

var LastMockArgscomponentColumnCountOfType MockArgsTypecomponentColumnCountOfType

// AuxMockcomponentColumnCountOfType(argcolumns map[string]*ColumnMetadata, argkind string)(reta int) - Generated mock function
func AuxMockcomponentColumnCountOfType(argcolumns map[string]*ColumnMetadata, argkind string) (reta int) {
	LastMockArgscomponentColumnCountOfType = MockArgsTypecomponentColumnCountOfType{
		ApomockCallNumber: AuxMockGetRecorderAuxMockcomponentColumnCountOfType(),
		Argcolumns:        argcolumns,
		Argkind:           argkind,
	}
	rargs, rerr := apomock.GetNext("gocql.componentColumnCountOfType")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.componentColumnCountOfType")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.componentColumnCountOfType")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(int)
	}
	return
}

// RecorderAuxMockcomponentColumnCountOfType  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockcomponentColumnCountOfType int = 0

var condRecorderAuxMockcomponentColumnCountOfType *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockcomponentColumnCountOfType(i int) {
	condRecorderAuxMockcomponentColumnCountOfType.L.Lock()
	for recorderAuxMockcomponentColumnCountOfType < i {
		condRecorderAuxMockcomponentColumnCountOfType.Wait()
	}
	condRecorderAuxMockcomponentColumnCountOfType.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockcomponentColumnCountOfType() {
	condRecorderAuxMockcomponentColumnCountOfType.L.Lock()
	recorderAuxMockcomponentColumnCountOfType++
	condRecorderAuxMockcomponentColumnCountOfType.L.Unlock()
	condRecorderAuxMockcomponentColumnCountOfType.Broadcast()
}
func AuxMockGetRecorderAuxMockcomponentColumnCountOfType() (ret int) {
	condRecorderAuxMockcomponentColumnCountOfType.L.Lock()
	ret = recorderAuxMockcomponentColumnCountOfType
	condRecorderAuxMockcomponentColumnCountOfType.L.Unlock()
	return
}

// componentColumnCountOfType - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func componentColumnCountOfType(argcolumns map[string]*ColumnMetadata, argkind string) (reta int) {
	FuncAuxMockcomponentColumnCountOfType, ok := apomock.GetRegisteredFunc("gocql.componentColumnCountOfType")
	if ok {
		reta = FuncAuxMockcomponentColumnCountOfType.(func(argcolumns map[string]*ColumnMetadata, argkind string) (reta int))(argcolumns, argkind)
	} else {
		panic("FuncAuxMockcomponentColumnCountOfType ")
	}
	AuxMockIncrementRecorderAuxMockcomponentColumnCountOfType()
	return
}

//
// Mock: newSchemaDescriber(argsession *Session)(reta *schemaDescriber)
//

type MockArgsTypenewSchemaDescriber struct {
	ApomockCallNumber int
	Argsession        *Session
}

var LastMockArgsnewSchemaDescriber MockArgsTypenewSchemaDescriber

// AuxMocknewSchemaDescriber(argsession *Session)(reta *schemaDescriber) - Generated mock function
func AuxMocknewSchemaDescriber(argsession *Session) (reta *schemaDescriber) {
	LastMockArgsnewSchemaDescriber = MockArgsTypenewSchemaDescriber{
		ApomockCallNumber: AuxMockGetRecorderAuxMocknewSchemaDescriber(),
		Argsession:        argsession,
	}
	rargs, rerr := apomock.GetNext("gocql.newSchemaDescriber")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.newSchemaDescriber")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.newSchemaDescriber")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*schemaDescriber)
	}
	return
}

// RecorderAuxMocknewSchemaDescriber  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMocknewSchemaDescriber int = 0

var condRecorderAuxMocknewSchemaDescriber *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMocknewSchemaDescriber(i int) {
	condRecorderAuxMocknewSchemaDescriber.L.Lock()
	for recorderAuxMocknewSchemaDescriber < i {
		condRecorderAuxMocknewSchemaDescriber.Wait()
	}
	condRecorderAuxMocknewSchemaDescriber.L.Unlock()
}

func AuxMockIncrementRecorderAuxMocknewSchemaDescriber() {
	condRecorderAuxMocknewSchemaDescriber.L.Lock()
	recorderAuxMocknewSchemaDescriber++
	condRecorderAuxMocknewSchemaDescriber.L.Unlock()
	condRecorderAuxMocknewSchemaDescriber.Broadcast()
}
func AuxMockGetRecorderAuxMocknewSchemaDescriber() (ret int) {
	condRecorderAuxMocknewSchemaDescriber.L.Lock()
	ret = recorderAuxMocknewSchemaDescriber
	condRecorderAuxMocknewSchemaDescriber.L.Unlock()
	return
}

// newSchemaDescriber - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func newSchemaDescriber(argsession *Session) (reta *schemaDescriber) {
	FuncAuxMocknewSchemaDescriber, ok := apomock.GetRegisteredFunc("gocql.newSchemaDescriber")
	if ok {
		reta = FuncAuxMocknewSchemaDescriber.(func(argsession *Session) (reta *schemaDescriber))(argsession)
	} else {
		panic("FuncAuxMocknewSchemaDescriber ")
	}
	AuxMockIncrementRecorderAuxMocknewSchemaDescriber()
	return
}

//
// Mock: getKeyspaceMetadata(argsession *Session, argkeyspaceName string)(reta *KeyspaceMetadata, retb error)
//

type MockArgsTypegetKeyspaceMetadata struct {
	ApomockCallNumber int
	Argsession        *Session
	ArgkeyspaceName   string
}

var LastMockArgsgetKeyspaceMetadata MockArgsTypegetKeyspaceMetadata

// AuxMockgetKeyspaceMetadata(argsession *Session, argkeyspaceName string)(reta *KeyspaceMetadata, retb error) - Generated mock function
func AuxMockgetKeyspaceMetadata(argsession *Session, argkeyspaceName string) (reta *KeyspaceMetadata, retb error) {
	LastMockArgsgetKeyspaceMetadata = MockArgsTypegetKeyspaceMetadata{
		ApomockCallNumber: AuxMockGetRecorderAuxMockgetKeyspaceMetadata(),
		Argsession:        argsession,
		ArgkeyspaceName:   argkeyspaceName,
	}
	rargs, rerr := apomock.GetNext("gocql.getKeyspaceMetadata")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.getKeyspaceMetadata")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.getKeyspaceMetadata")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(*KeyspaceMetadata)
	}
	if rargs.GetArg(1) != nil {
		retb = rargs.GetArg(1).(error)
	}
	return
}

// RecorderAuxMockgetKeyspaceMetadata  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockgetKeyspaceMetadata int = 0

var condRecorderAuxMockgetKeyspaceMetadata *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockgetKeyspaceMetadata(i int) {
	condRecorderAuxMockgetKeyspaceMetadata.L.Lock()
	for recorderAuxMockgetKeyspaceMetadata < i {
		condRecorderAuxMockgetKeyspaceMetadata.Wait()
	}
	condRecorderAuxMockgetKeyspaceMetadata.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockgetKeyspaceMetadata() {
	condRecorderAuxMockgetKeyspaceMetadata.L.Lock()
	recorderAuxMockgetKeyspaceMetadata++
	condRecorderAuxMockgetKeyspaceMetadata.L.Unlock()
	condRecorderAuxMockgetKeyspaceMetadata.Broadcast()
}
func AuxMockGetRecorderAuxMockgetKeyspaceMetadata() (ret int) {
	condRecorderAuxMockgetKeyspaceMetadata.L.Lock()
	ret = recorderAuxMockgetKeyspaceMetadata
	condRecorderAuxMockgetKeyspaceMetadata.L.Unlock()
	return
}

// getKeyspaceMetadata - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func getKeyspaceMetadata(argsession *Session, argkeyspaceName string) (reta *KeyspaceMetadata, retb error) {
	FuncAuxMockgetKeyspaceMetadata, ok := apomock.GetRegisteredFunc("gocql.getKeyspaceMetadata")
	if ok {
		reta, retb = FuncAuxMockgetKeyspaceMetadata.(func(argsession *Session, argkeyspaceName string) (reta *KeyspaceMetadata, retb error))(argsession, argkeyspaceName)
	} else {
		panic("FuncAuxMockgetKeyspaceMetadata ")
	}
	AuxMockIncrementRecorderAuxMockgetKeyspaceMetadata()
	return
}

//
// Mock: (recvt *typeParser)parseParamNodes()(retparams []typeParserParamNode, retok bool)
//

type MockArgsTypetypeParserparseParamNodes struct {
	ApomockCallNumber int
}

var LastMockArgstypeParserparseParamNodes MockArgsTypetypeParserparseParamNodes

// (recvt *typeParser)AuxMockparseParamNodes()(retparams []typeParserParamNode, retok bool) - Generated mock function
func (recvt *typeParser) AuxMockparseParamNodes() (retparams []typeParserParamNode, retok bool) {
	rargs, rerr := apomock.GetNext("gocql.typeParser.parseParamNodes")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.typeParser.parseParamNodes")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.typeParser.parseParamNodes")
	}
	if rargs.GetArg(0) != nil {
		retparams = rargs.GetArg(0).([]typeParserParamNode)
	}
	if rargs.GetArg(1) != nil {
		retok = rargs.GetArg(1).(bool)
	}
	return
}

// RecorderAuxMockPtrtypeParserparseParamNodes  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtypeParserparseParamNodes int = 0

var condRecorderAuxMockPtrtypeParserparseParamNodes *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtypeParserparseParamNodes(i int) {
	condRecorderAuxMockPtrtypeParserparseParamNodes.L.Lock()
	for recorderAuxMockPtrtypeParserparseParamNodes < i {
		condRecorderAuxMockPtrtypeParserparseParamNodes.Wait()
	}
	condRecorderAuxMockPtrtypeParserparseParamNodes.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtypeParserparseParamNodes() {
	condRecorderAuxMockPtrtypeParserparseParamNodes.L.Lock()
	recorderAuxMockPtrtypeParserparseParamNodes++
	condRecorderAuxMockPtrtypeParserparseParamNodes.L.Unlock()
	condRecorderAuxMockPtrtypeParserparseParamNodes.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtypeParserparseParamNodes() (ret int) {
	condRecorderAuxMockPtrtypeParserparseParamNodes.L.Lock()
	ret = recorderAuxMockPtrtypeParserparseParamNodes
	condRecorderAuxMockPtrtypeParserparseParamNodes.L.Unlock()
	return
}

// (recvt *typeParser)parseParamNodes - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *typeParser) parseParamNodes() (retparams []typeParserParamNode, retok bool) {
	FuncAuxMockPtrtypeParserparseParamNodes, ok := apomock.GetRegisteredFunc("gocql.typeParser.parseParamNodes")
	if ok {
		retparams, retok = FuncAuxMockPtrtypeParserparseParamNodes.(func(recvt *typeParser) (retparams []typeParserParamNode, retok bool))(recvt)
	} else {
		panic("FuncAuxMockPtrtypeParserparseParamNodes ")
	}
	AuxMockIncrementRecorderAuxMockPtrtypeParserparseParamNodes()
	return
}

//
// Mock: (recvt *typeParser)skipWhitespace()()
//

type MockArgsTypetypeParserskipWhitespace struct {
	ApomockCallNumber int
}

var LastMockArgstypeParserskipWhitespace MockArgsTypetypeParserskipWhitespace

// (recvt *typeParser)AuxMockskipWhitespace()() - Generated mock function
func (recvt *typeParser) AuxMockskipWhitespace() {
	return
}

// RecorderAuxMockPtrtypeParserskipWhitespace  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtypeParserskipWhitespace int = 0

var condRecorderAuxMockPtrtypeParserskipWhitespace *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtypeParserskipWhitespace(i int) {
	condRecorderAuxMockPtrtypeParserskipWhitespace.L.Lock()
	for recorderAuxMockPtrtypeParserskipWhitespace < i {
		condRecorderAuxMockPtrtypeParserskipWhitespace.Wait()
	}
	condRecorderAuxMockPtrtypeParserskipWhitespace.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtypeParserskipWhitespace() {
	condRecorderAuxMockPtrtypeParserskipWhitespace.L.Lock()
	recorderAuxMockPtrtypeParserskipWhitespace++
	condRecorderAuxMockPtrtypeParserskipWhitespace.L.Unlock()
	condRecorderAuxMockPtrtypeParserskipWhitespace.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtypeParserskipWhitespace() (ret int) {
	condRecorderAuxMockPtrtypeParserskipWhitespace.L.Lock()
	ret = recorderAuxMockPtrtypeParserskipWhitespace
	condRecorderAuxMockPtrtypeParserskipWhitespace.L.Unlock()
	return
}

// (recvt *typeParser)skipWhitespace - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *typeParser) skipWhitespace() {
	FuncAuxMockPtrtypeParserskipWhitespace, ok := apomock.GetRegisteredFunc("gocql.typeParser.skipWhitespace")
	if ok {
		FuncAuxMockPtrtypeParserskipWhitespace.(func(recvt *typeParser))(recvt)
	} else {
		panic("FuncAuxMockPtrtypeParserskipWhitespace ")
	}
	AuxMockIncrementRecorderAuxMockPtrtypeParserskipWhitespace()
	return
}

//
// Mock: isWhitespaceChar(argc byte)(reta bool)
//

type MockArgsTypeisWhitespaceChar struct {
	ApomockCallNumber int
	Argc              byte
}

var LastMockArgsisWhitespaceChar MockArgsTypeisWhitespaceChar

// AuxMockisWhitespaceChar(argc byte)(reta bool) - Generated mock function
func AuxMockisWhitespaceChar(argc byte) (reta bool) {
	LastMockArgsisWhitespaceChar = MockArgsTypeisWhitespaceChar{
		ApomockCallNumber: AuxMockGetRecorderAuxMockisWhitespaceChar(),
		Argc:              argc,
	}
	rargs, rerr := apomock.GetNext("gocql.isWhitespaceChar")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.isWhitespaceChar")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.isWhitespaceChar")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockisWhitespaceChar  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockisWhitespaceChar int = 0

var condRecorderAuxMockisWhitespaceChar *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockisWhitespaceChar(i int) {
	condRecorderAuxMockisWhitespaceChar.L.Lock()
	for recorderAuxMockisWhitespaceChar < i {
		condRecorderAuxMockisWhitespaceChar.Wait()
	}
	condRecorderAuxMockisWhitespaceChar.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockisWhitespaceChar() {
	condRecorderAuxMockisWhitespaceChar.L.Lock()
	recorderAuxMockisWhitespaceChar++
	condRecorderAuxMockisWhitespaceChar.L.Unlock()
	condRecorderAuxMockisWhitespaceChar.Broadcast()
}
func AuxMockGetRecorderAuxMockisWhitespaceChar() (ret int) {
	condRecorderAuxMockisWhitespaceChar.L.Lock()
	ret = recorderAuxMockisWhitespaceChar
	condRecorderAuxMockisWhitespaceChar.L.Unlock()
	return
}

// isWhitespaceChar - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func isWhitespaceChar(argc byte) (reta bool) {
	FuncAuxMockisWhitespaceChar, ok := apomock.GetRegisteredFunc("gocql.isWhitespaceChar")
	if ok {
		reta = FuncAuxMockisWhitespaceChar.(func(argc byte) (reta bool))(argc)
	} else {
		panic("FuncAuxMockisWhitespaceChar ")
	}
	AuxMockIncrementRecorderAuxMockisWhitespaceChar()
	return
}

//
// Mock: (recvt *typeParser)nextIdentifier()(retid string, retfound bool)
//

type MockArgsTypetypeParsernextIdentifier struct {
	ApomockCallNumber int
}

var LastMockArgstypeParsernextIdentifier MockArgsTypetypeParsernextIdentifier

// (recvt *typeParser)AuxMocknextIdentifier()(retid string, retfound bool) - Generated mock function
func (recvt *typeParser) AuxMocknextIdentifier() (retid string, retfound bool) {
	rargs, rerr := apomock.GetNext("gocql.typeParser.nextIdentifier")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.typeParser.nextIdentifier")
	} else if rargs.NumArgs() != 2 {
		panic("All return parameters not provided for method:gocql.typeParser.nextIdentifier")
	}
	if rargs.GetArg(0) != nil {
		retid = rargs.GetArg(0).(string)
	}
	if rargs.GetArg(1) != nil {
		retfound = rargs.GetArg(1).(bool)
	}
	return
}

// RecorderAuxMockPtrtypeParsernextIdentifier  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockPtrtypeParsernextIdentifier int = 0

var condRecorderAuxMockPtrtypeParsernextIdentifier *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockPtrtypeParsernextIdentifier(i int) {
	condRecorderAuxMockPtrtypeParsernextIdentifier.L.Lock()
	for recorderAuxMockPtrtypeParsernextIdentifier < i {
		condRecorderAuxMockPtrtypeParsernextIdentifier.Wait()
	}
	condRecorderAuxMockPtrtypeParsernextIdentifier.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockPtrtypeParsernextIdentifier() {
	condRecorderAuxMockPtrtypeParsernextIdentifier.L.Lock()
	recorderAuxMockPtrtypeParsernextIdentifier++
	condRecorderAuxMockPtrtypeParsernextIdentifier.L.Unlock()
	condRecorderAuxMockPtrtypeParsernextIdentifier.Broadcast()
}
func AuxMockGetRecorderAuxMockPtrtypeParsernextIdentifier() (ret int) {
	condRecorderAuxMockPtrtypeParsernextIdentifier.L.Lock()
	ret = recorderAuxMockPtrtypeParsernextIdentifier
	condRecorderAuxMockPtrtypeParsernextIdentifier.L.Unlock()
	return
}

// (recvt *typeParser)nextIdentifier - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func (recvt *typeParser) nextIdentifier() (retid string, retfound bool) {
	FuncAuxMockPtrtypeParsernextIdentifier, ok := apomock.GetRegisteredFunc("gocql.typeParser.nextIdentifier")
	if ok {
		retid, retfound = FuncAuxMockPtrtypeParsernextIdentifier.(func(recvt *typeParser) (retid string, retfound bool))(recvt)
	} else {
		panic("FuncAuxMockPtrtypeParsernextIdentifier ")
	}
	AuxMockIncrementRecorderAuxMockPtrtypeParsernextIdentifier()
	return
}

//
// Mock: isIdentifierChar(argc byte)(reta bool)
//

type MockArgsTypeisIdentifierChar struct {
	ApomockCallNumber int
	Argc              byte
}

var LastMockArgsisIdentifierChar MockArgsTypeisIdentifierChar

// AuxMockisIdentifierChar(argc byte)(reta bool) - Generated mock function
func AuxMockisIdentifierChar(argc byte) (reta bool) {
	LastMockArgsisIdentifierChar = MockArgsTypeisIdentifierChar{
		ApomockCallNumber: AuxMockGetRecorderAuxMockisIdentifierChar(),
		Argc:              argc,
	}
	rargs, rerr := apomock.GetNext("gocql.isIdentifierChar")
	if rerr != nil {
		panic("Error getting next entry for method: gocql.isIdentifierChar")
	} else if rargs.NumArgs() != 1 {
		panic("All return parameters not provided for method:gocql.isIdentifierChar")
	}
	if rargs.GetArg(0) != nil {
		reta = rargs.GetArg(0).(bool)
	}
	return
}

// RecorderAuxMockisIdentifierChar  - Stats Recorder to keep a count on how many times the function has been called
var recorderAuxMockisIdentifierChar int = 0

var condRecorderAuxMockisIdentifierChar *sync.Cond = &sync.Cond{L: &sync.Mutex{}}

func AuxMockWaitForRecorderAuxMockisIdentifierChar(i int) {
	condRecorderAuxMockisIdentifierChar.L.Lock()
	for recorderAuxMockisIdentifierChar < i {
		condRecorderAuxMockisIdentifierChar.Wait()
	}
	condRecorderAuxMockisIdentifierChar.L.Unlock()
}

func AuxMockIncrementRecorderAuxMockisIdentifierChar() {
	condRecorderAuxMockisIdentifierChar.L.Lock()
	recorderAuxMockisIdentifierChar++
	condRecorderAuxMockisIdentifierChar.L.Unlock()
	condRecorderAuxMockisIdentifierChar.Broadcast()
}
func AuxMockGetRecorderAuxMockisIdentifierChar() (ret int) {
	condRecorderAuxMockisIdentifierChar.L.Lock()
	ret = recorderAuxMockisIdentifierChar
	condRecorderAuxMockisIdentifierChar.L.Unlock()
	return
}

// isIdentifierChar - Mocked version of actual API. Does Recordkeeping and calls the user provided mock or default mock
func isIdentifierChar(argc byte) (reta bool) {
	FuncAuxMockisIdentifierChar, ok := apomock.GetRegisteredFunc("gocql.isIdentifierChar")
	if ok {
		reta = FuncAuxMockisIdentifierChar.(func(argc byte) (reta bool))(argc)
	} else {
		panic("FuncAuxMockisIdentifierChar ")
	}
	AuxMockIncrementRecorderAuxMockisIdentifierChar()
	return
}
