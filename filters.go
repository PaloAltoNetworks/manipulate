package manipulate

import "fmt"

// Comparators represent various comparison operations.
const (
	EqualComparator FilterComparator = iota + 1
	GreaterComparator
	LesserComparator
	InComparator
	ContainComparator
)

// An FilterComparator is the type of a operator used by a filter.
type FilterComparator int

// FilterComparators are a list of FilterOperator.
type FilterComparators []FilterComparator

// NewFilterComparators returns a new FilterKeys
func NewFilterComparators(operators ...FilterComparator) FilterComparators {

	fos := FilterComparators{}

	if len(operators) == 0 {
		return fos
	}

	for _, o := range operators {
		fos = append(fos, o)
	}

	return fos
}

// Then adds a new filter key to te receiver and returns it.
func (f FilterComparators) Then(operators ...FilterComparator) FilterComparators {

	for _, o := range operators {
		f = append(f, o)
	}

	return f
}

// Operators represent various operators.
const (
	OrOperator FilterOperator = iota + 1
	AndOperator
)

// An FilterOperator is the type of a operator used by a filter.
type FilterOperator int

// FilterOperators are a list of FilterOperator.
type FilterOperators []FilterOperator

// NewFilterOperators returns a new FilterOperators.
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

// NewFilterKeys returns a new FilterKeys.
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

// FilterValueComposer adds values and operators.
type FilterValueComposer interface {
	Equals(...interface{}) FilterKeyComposer
	GreaterThan(...interface{}) FilterKeyComposer
	LesserThan(...interface{}) FilterKeyComposer
	In(...interface{}) FilterKeyComposer
	Contains(...interface{}) FilterKeyComposer
}

// FilterKeyComposer composes a filter.
type FilterKeyComposer interface {
	WithKey(...string) FilterValueComposer
	AndKey(...string) FilterValueComposer
	OrKey(...string) FilterValueComposer
	Done() *Filter
}

// Filter is a filter struct which can be used with Cassandra
type Filter struct {
	Keys        FilterKeys
	Values      FilterValues
	Comparators FilterComparators
	Operators   FilterOperators
}

// NewFilter returns a new filter.
func NewFilter() *Filter {

	return &Filter{
		Keys:        NewFilterKeys(),
		Values:      NewFilterValues(),
		Comparators: NewFilterComparators(),
		Operators:   NewFilterOperators(),
	}
}

// NewFilterComposer returns a FilterComposer.
func NewFilterComposer() FilterKeyComposer {

	return NewFilter()
}

// Equals adds a an equality comparator to the FilterComposer.
func (f *Filter) Equals(values ...interface{}) FilterKeyComposer {
	f.Values = f.Values.Then(values...)
	f.Comparators = f.Comparators.Then(EqualComparator)
	return f
}

// GreaterThan adds a greater than comparator to the FilterComposer.
func (f *Filter) GreaterThan(values ...interface{}) FilterKeyComposer {
	f.Values = f.Values.Then(values...)
	f.Comparators = f.Comparators.Then(LesserComparator)
	return f
}

// LesserThan adds a lesser than comparator to the FilterComposer.
func (f *Filter) LesserThan(values ...interface{}) FilterKeyComposer {
	f.Values = f.Values.Then(values...)
	f.Comparators = f.Comparators.Then(GreaterComparator)
	return f
}

// In adds a in comparator to the FilterComposer.
func (f *Filter) In(values ...interface{}) FilterKeyComposer {
	f.Values = f.Values.Then(values...)
	f.Comparators = f.Comparators.Then(InComparator)
	return f
}

// Contains adds a contains comparator to the FilterComposer.
func (f *Filter) Contains(values ...interface{}) FilterKeyComposer {
	f.Values = f.Values.Then(values...)
	f.Comparators = f.Comparators.Then(ContainComparator)
	return f
}

// AndKey adds a and operator to the FilterComposer.
func (f *Filter) AndKey(keys ...string) FilterValueComposer {
	f.Operators = f.Operators.Then(AndOperator)
	f.Keys = f.Keys.Then(keys...)
	return f
}

// OrKey adds a or operator to the FilterComposer.
func (f *Filter) OrKey(keys ...string) FilterValueComposer {
	f.Operators = f.Operators.Then(OrOperator)
	f.Keys = f.Keys.Then(keys...)
	return f
}

// WithKey adds a key to FilterComposer.
func (f *Filter) WithKey(keys ...string) FilterValueComposer {
	f.Keys = f.Keys.Then(keys...)
	return f
}

// Done terminates the filter composition and returns the *Filter.
func (f *Filter) Done() *Filter {
	return f
}

func (f *Filter) String() string {
	return fmt.Sprintf("%v", *f)
}
