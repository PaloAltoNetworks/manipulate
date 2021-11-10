package cli

// Format Types
const (
	// FormatTypeColumn represents the value "column"
	FormatTypeColumn = "column"
	// FormatTypeHash represents the value "hash"
	FormatTypeHash = "hash"
	// FormatTypeCount represents the value "count"
	FormatTypeCount = "count"
	// FormatTypeArray represents the value "array"
	FormatTypeArray = "array"
)

// Output Flags
const (
	// FlagOutputTable represents the value "table"
	FlagOutputTable = "table"
	// FlagOutputJSON represents the value "json"
	FlagOutputJSON = "json"
	// FlagOutputNone represents the value "none"
	FlagOutputNone = "none"
	// FlagOutputDefault represents the value "default"
	FlagOutputDefault = "default"
	// FlagOutputYAML represents the value "json"
	FlagOutputYAML = "yaml"
	// FlagOutputTemplate represents the value "template"
	FlagOutputTemplate = "template"
)

// General Flags
const (
	// FlagTrackingID represents the value "tracking-id"
	FlagTrackingID = "tracking-id"
	// FlagOutput represents the value "output"
	FlagOutput = "output"
	// FlagAPI represents the key "api"
	FlagAPI = "api"
	// FlagAPISkipVerify represents the key "api-skip-verify"
	FlagAPISkipVerify = "api-skip-verify"
	// FlagNamespace represents the key "namespace"
	FlagNamespace = "namespace"
	// FlagToken represents the key "token"
	FlagToken = "token"
	// FlagCACertPath represents the key "api-cacert"
	FlagCACertPath = "api-cacert"
	// FlagAppCredentials represents the key "creds"
	FlagAppCredentials = "creds"
	// FlagEncoding represents the key "encoding"
	FlagEncoding = "encoding"
	// FlagRecursive represents the key "recursive"
	FlagRecursive = "recursive"
	// FlagParameters represents the value "param"
	FlagParameters = "param"
)

// List Flags
const (
	// FlagPageSize represents the value "page-size"
	FlagPageSize = "page-size"
	// FlagPageKey represents the value "page"
	FlagPage = "page"
	// FlagFilterKey represents the value "filter"
	FlagFilter = "filter"
	// FlagOrderKey represents the value "order"
	FlagOrder = "order"
)

// Delete Flags
const (
	// FlagForce represents the value "force"
	FlagForce = "force"
)
