package manipcli

// Format Types
const (
	formatTypeColumn = "column"
	formatTypeHash   = "hash"
	formatTypeCount  = "count"
	formatTypeArray  = "array"
)

// Output Flags
const (
	flagOutputTable    = "table"
	flagOutputJSON     = "json"
	flagOutputNone     = "none"
	flagOutputDefault  = "default"
	flagOutputYAML     = "yaml"
	flagOutputTemplate = "template"
)

// General Flags
const (
	flagTrackingID    = "tracking-id"
	flagOutput        = "output"
	flagAPI           = "api"
	flagAPISkipVerify = "api-skip-verify"
	flagNamespace     = "namespace"
	flagToken         = "token"
	flagCACertPath    = "api-cacert"
	flagEncoding      = "encoding"
	flagRecursive     = "recursive"
	flagParameters    = "param"
)

// Create/update FLags
const (
	flagInteractive = "interactive"
	flagInputValues = "input-values"
	flagInputData   = "input-data"
	flagEditor      = "editor"
	flagInputFile   = "input-file"
	flagInputURL    = "input-url"
	flagInputSet    = "input-set"
	flagPrint       = "print"
	flagRender      = "render"
)

// List Flags
const (
	flagPageSize = "page-size"
	flagPage     = "page"
	flagFilter   = "filter"
	flagOrder    = "order"
)

// Delete Flags
const (
	flagForce   = "force"
	flagConfirm = "confirm"
)
