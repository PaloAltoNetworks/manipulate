package manipulate

import (
	"bytes"
	"fmt"
)

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
func NewFilterComparators(comparators ...FilterComparator) FilterComparators {

	fcs := FilterComparators{}

	if len(comparators) == 0 {
		return fcs
	}

	for _, o := range comparators {
		fcs = append(fcs, o)
	}

	return fcs
}

// Then adds a new filter key to te receiver and returns it.
func (f FilterComparators) Then(comparators ...FilterComparator) FilterComparators {

	for _, o := range comparators {
		f = append(f, o)
	}

	return f
}

// Operators represent various operators.
const (
	InitialOperator FilterOperator = iota + 1
	OrOperator
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
	keys        FilterKeys
	values      FilterValues
	comparators FilterComparators
	operators   FilterOperators
}

// NewFilter returns a new filter.
func NewFilter() *Filter {

	return &Filter{
		keys:        NewFilterKeys(),
		values:      NewFilterValues(),
		comparators: NewFilterComparators(),
		operators:   NewFilterOperators(),
	}
}

// NewFilterComposer returns a FilterComposer.
func NewFilterComposer() FilterKeyComposer {

	return NewFilter()
}

// Keys returns the current keys.
func (f *Filter) Keys() FilterKeys {
	return f.keys
}

// Values returns the current values.
func (f *Filter) Values() FilterValues {
	return f.values
}

// Operators returns the current operators.
func (f *Filter) Operators() FilterOperators {
	return f.operators
}

// Comparators returns the current comparators.
func (f *Filter) Comparators() FilterComparators {
	return f.comparators
}

// Equals adds a an equality comparator to the FilterComposer.
func (f *Filter) Equals(values ...interface{}) FilterKeyComposer {
	f.values = f.values.Then(values...)
	f.comparators = f.comparators.Then(EqualComparator)
	return f
}

// GreaterThan adds a greater than comparator to the FilterComposer.
func (f *Filter) GreaterThan(values ...interface{}) FilterKeyComposer {
	f.values = f.values.Then(values...)
	f.comparators = f.comparators.Then(GreaterComparator)
	return f
}

// LesserThan adds a lesser than comparator to the FilterComposer.
func (f *Filter) LesserThan(values ...interface{}) FilterKeyComposer {
	f.values = f.values.Then(values...)
	f.comparators = f.comparators.Then(LesserComparator)
	return f
}

// In adds a in comparator to the FilterComposer.
func (f *Filter) In(values ...interface{}) FilterKeyComposer {
	f.values = f.values.Then(values...)
	f.comparators = f.comparators.Then(InComparator)
	return f
}

// Contains adds a contains comparator to the FilterComposer.
func (f *Filter) Contains(values ...interface{}) FilterKeyComposer {
	f.values = f.values.Then(values...)
	f.comparators = f.comparators.Then(ContainComparator)
	return f
}

// AndKey adds a and operator to the FilterComposer.
func (f *Filter) AndKey(keys ...string) FilterValueComposer {
	f.operators = f.operators.Then(AndOperator)
	f.keys = f.keys.Then(keys...)
	return f
}

// OrKey adds a or operator to the FilterComposer.
func (f *Filter) OrKey(keys ...string) FilterValueComposer {
	f.operators = f.operators.Then(OrOperator)
	f.keys = f.keys.Then(keys...)
	return f
}

// WithKey adds a key to FilterComposer.
func (f *Filter) WithKey(keys ...string) FilterValueComposer {
	f.operators = f.operators.Then(InitialOperator)
	f.keys = f.keys.Then(keys...)
	return f
}

// Done terminates the filter composition and returns the *Filter.
func (f *Filter) Done() *Filter {
	return f
}

func (f *Filter) String() string {

	var buffer bytes.Buffer

	for i, operator := range f.operators {
		WriteString(&buffer, translateOperator(operator))
		if i > 0 {
			WriteString(&buffer, " ")
		}
		WriteString(&buffer, fmt.Sprintf("%v", f.keys[i]))
		WriteString(&buffer, " ")
		WriteString(&buffer, translateComparator(f.comparators[i]))
		WriteString(&buffer, " ")
		WriteString(&buffer, fmt.Sprintf("%v", f.values[i]))

		if i+1 < len(f.operators) {
			WriteString(&buffer, " ")
		}
	}

	return buffer.String()
}

func translateComparator(comparator FilterComparator) string {

	switch comparator {
	case EqualComparator:
		return "="
	case GreaterComparator:
		return ">="
	case LesserComparator:
		return "<="
	case InComparator:
		return "in"
	case ContainComparator:
		return "contains"
	}

	return ""
}

func translateOperator(operator FilterOperator) string {

	switch operator {
	case InitialOperator:
		return ""
	case AndOperator:
		return "and"
	case OrOperator:
		return "or"
	}

	return ""
}
