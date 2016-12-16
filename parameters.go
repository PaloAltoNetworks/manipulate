package manipulate

// Parameters is a parameter struct which can be used with Cassandra (fields)
// or HTTP (KeyValues)
type Parameters struct {
	IfNotExists bool
	IfExists    bool
	UsingTTL    bool
	OrderByDesc string
	OrderByAsc  string
	KeyValues   map[string]string
}
