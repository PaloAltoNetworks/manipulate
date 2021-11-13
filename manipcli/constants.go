package manipcli

// Format Types
const (
	formatTypeColumn = "column"
	formatTypeHash = "hash"
	formatTypeCount = "count"
	formatTypeArray = "array"
)

// Output Flags
const (
	flagOutputTable = "table"
	flagOutputJSON = "json"
	flagOutputNone = "none"
	flagOutputDefault = "default"
	flagOutputYAML = "yaml"
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

// List Flags
const (
	flagPageSize = "page-size"
	flagPage = "page"
	flagFilter = "filter"
	flagOrder = "order"
)

// Delete Flags
const (
	flagForce = "force"
	flagConfirm = "confirm"
)
