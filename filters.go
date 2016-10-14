package manipulate

// Operators represent various comparison operations.
const (
	EqualOperator FilterOperator = iota + 1
	EqualOrSuperiorOperator
	EqualOrInferiorOperator
	InOperator
	ContainsOperator
)

// An FilterOperator is the type of a operator used by a filter.
type FilterOperator int

// FilterOperators are a list of FilterOperator.
type FilterOperators []FilterOperator

// NewFilterOperators returns a new FilterKeys
func NewFilterOperators(operators ...FilterOperator) FilterOperators {

	fos := FilterOperators{}

	if len(operators) == 0 {
		return fos
	}

	for _, o := range operators {
		fos = append(fos, o)
	}

	return fos
}

// Then adds a new filter key to te receiver and returns it.
func (f FilterOperators) Then(operators ...FilterOperator) FilterOperators {

	for _, o := range operators {
		f = append(f, o)
	}

	return f
}

// FilterKey represents a filter key.
type FilterKey []string

// FilterKeys represents a list of FilterKey.
type FilterKeys []FilterKey

// NewFilterKeys returns a new FilterKeys
func NewFilterKeys(keys ...string) FilterKeys {

	if len(keys) == 0 {
		return FilterKeys{}
	}

	fk := FilterKey{}

	for _, k := range keys {
		fk = append(fk, k)
	}

	return FilterKeys{fk}
}

// Then adds a new filter key to te receiver and returns it.
func (f FilterKeys) Then(keys ...string) FilterKeys {

	fk := FilterKey{}

	for _, k := range keys {
		fk = append(fk, k)
	}

	return append(f, fk)
}

// FilterValue represents a filter value.
type FilterValue []interface{}

// FilterValues represents a list of FilterValue.
type FilterValues [][]interface{}

// NewFilterValues returns a new FilterValues.
func NewFilterValues(values ...interface{}) FilterValues {

	if len(values) == 0 {
		return FilterValues{}
	}

	fv := FilterValue{}

	for _, v := range values {
		fv = append(fv, v)
	}

	return FilterValues{fv}
}

// Then adds a new value to the reciever and returns it.
func (f FilterValues) Then(values ...interface{}) FilterValues {

	fv := FilterValue{}

	for _, v := range values {
		fv = append(fv, v)
	}

	return append(f, fv)
}

// Filter is a filter struct which can be used with Cassandra
type Filter struct {
	Keys      FilterKeys
	Values    FilterValues
	Operators FilterOperators
}

// NewFilter returns a new filter with the given keys, values and operators.
func NewFilter(keys FilterKeys, values FilterValues, operators FilterOperators) *Filter {

	return &Filter{
		Keys:      keys,
		Values:    values,
		Operators: operators,
	}
}

// NewSimpleFilter returns a new Filter.
func NewSimpleFilter(key string, value interface{}, operator FilterOperator) *Filter {

	return NewFilter(NewFilterKeys(key), NewFilterValues(value), NewFilterOperators(operator))
}
