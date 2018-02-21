package manipulate

import (
	"bytes"
	"fmt"
	"strings"
)

// An FilterComparator is the type of a operator used by a filter.
type FilterComparator int

// FilterComparators are a list of FilterOperator.
type FilterComparators []FilterComparator

// Comparators represent various comparison operations.
const (
	EqualComparator FilterComparator = iota + 1
	NotEqualComparator
	GreaterComparator
	LesserComparator
	InComparator
	NotInComparator
	ContainComparator
	NotContainComparator
	MatchComparator
	NotMatchComparator
)

func (f FilterComparators) add(comparators ...FilterComparator) FilterComparators {

	for _, o := range comparators {
		f = append(f, o)
	}

	return f
}

// An FilterOperator is the type of a operator used by a filter.
type FilterOperator int

// FilterOperators are a list of FilterOperator.
type FilterOperators []FilterOperator

// Operators represent various operators.
const (
	InitialOperator FilterOperator = iota + 1
	OrOperator
	AndOperator
	OrFilterOperator
	AndFilterOperator
)

// FilterKeys represents a list of FilterKey.
type FilterKeys []string

// FilterValue represents a filter value.
type FilterValue []interface{}

// FilterValues represents a list of FilterValue.
type FilterValues [][]interface{}

// Then adds a new value to the receiver and returns it.
func (f FilterValues) add(values ...interface{}) FilterValues {

	fv := FilterValue{}
	for _, v := range values {
		fv = append(fv, v)
	}

	return append(f, fv)
}

// FilterValueComposer adds values and operators.
type FilterValueComposer interface {
	Equals(...interface{}) FilterKeyComposer
	NotEquals(...interface{}) FilterKeyComposer
	GreaterThan(...interface{}) FilterKeyComposer
	LesserThan(...interface{}) FilterKeyComposer
	In(...interface{}) FilterKeyComposer
	Contains(...interface{}) FilterKeyComposer
	Matches(...interface{}) FilterKeyComposer
}

// FilterKeyComposer composes a filter.
type FilterKeyComposer interface {
	WithKey(string) FilterValueComposer
	AndKey(string) FilterValueComposer
	OrKey(string) FilterValueComposer

	And(...*Filter) FilterKeyComposer
	Or(...*Filter) FilterKeyComposer

	Done() *Filter
}

// Filter is a filter struct which can be used with Cassandra
type Filter struct {
	keys        FilterKeys
	values      FilterValues
	comparators FilterComparators
	operators   FilterOperators
	ands        [][]*Filter
	ors         [][]*Filter
}

// NewFilter returns a new filter.
func NewFilter() *Filter {

	return &Filter{
		keys:        FilterKeys{},
		values:      FilterValues{},
		comparators: FilterComparators{},
		operators:   FilterOperators{},
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
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(EqualComparator)
	return f
}

// NotEquals adds a an non equality comparator to the FilterComposer.
func (f *Filter) NotEquals(values ...interface{}) FilterKeyComposer {
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(NotEqualComparator)
	return f
}

// GreaterThan adds a greater than comparator to the FilterComposer.
func (f *Filter) GreaterThan(values ...interface{}) FilterKeyComposer {
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(GreaterComparator)
	return f
}

// LesserThan adds a lesser than comparator to the FilterComposer.
func (f *Filter) LesserThan(values ...interface{}) FilterKeyComposer {
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(LesserComparator)
	return f
}

// In adds a in comparator to the FilterComposer.
func (f *Filter) In(values ...interface{}) FilterKeyComposer {
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(InComparator)
	return f
}

// Contains adds a contains comparator to the FilterComposer.
func (f *Filter) Contains(values ...interface{}) FilterKeyComposer {
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(ContainComparator)
	return f
}

// Matches adds a match comparator to the FilterComposer.
func (f *Filter) Matches(values ...interface{}) FilterKeyComposer {
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(MatchComparator)
	return f
}

// AndKey adds a and operator to the FilterComposer.
func (f *Filter) AndKey(key string) FilterValueComposer {
	f.operators = append(f.operators, AndOperator)
	f.keys = append(f.keys, key)
	return f
}

// OrKey adds a or operator to the FilterComposer.
func (f *Filter) OrKey(key string) FilterValueComposer {
	f.operators = append(f.operators, OrOperator)
	f.keys = append(f.keys, key)
	return f
}

// WithKey adds a key to FilterComposer.
func (f *Filter) WithKey(key string) FilterValueComposer {
	f.operators = append(f.operators, InitialOperator)
	f.keys = append(f.keys, key)
	f.ands = append(f.ands, nil)
	f.ors = append(f.ors, nil)
	return f
}

// And adds a new sub filter to FilterComposer.
func (f *Filter) And(filters ...*Filter) FilterKeyComposer {
	f.operators = append(f.operators, AndFilterOperator)
	f.keys = append(f.keys, "")
	f.ands = append(f.ands, filters)
	f.ors = append(f.ors, nil)
	return f
}

// Or adds a new sub filter to FilterComposer.
func (f *Filter) Or(filters ...*Filter) FilterKeyComposer {
	f.operators = append(f.operators, OrFilterOperator)
	f.keys = append(f.keys, "")
	f.ands = append(f.ands, nil)
	f.ors = append(f.ors, filters)
	return f
}

// Done terminates the filter composition and returns the *Filter.
func (f *Filter) Done() *Filter {
	return f
}

func (f *Filter) String() string {

	var buffer bytes.Buffer

	for i, operator := range f.operators {
		if i > 0 {
			writeString(&buffer, translateOperator(operator))
			writeString(&buffer, " ")
		}

		switch operator {

		case InitialOperator, AndOperator, OrOperator:
			writeString(&buffer, fmt.Sprintf("%v", f.keys[i]))
			writeString(&buffer, " ")
			writeString(&buffer, translateComparator(f.comparators[i]))
			writeString(&buffer, " ")
			writeString(&buffer, fmt.Sprintf("%v", f.values[i]))

		case AndFilterOperator:
			var strs []string
			for _, andf := range f.ands[i] {
				strs = append(strs, fmt.Sprintf("(%s)", andf))
			}
			writeString(&buffer, fmt.Sprintf("(%s)", strings.Join(strs, " and ")))

		case OrFilterOperator:
			var strs []string
			for _, orf := range f.ors[i] {
				strs = append(strs, fmt.Sprintf("(%s)", orf))
			}
			writeString(&buffer, fmt.Sprintf("(%s)", strings.Join(strs, " or ")))
		}

		if i+1 < len(f.operators) {
			writeString(&buffer, " ")
		}
	}

	return buffer.String()
}

func translateComparator(comparator FilterComparator) string {

	switch comparator {
	case EqualComparator:
		return "="
	case NotEqualComparator:
		return "!="
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
	case AndOperator, AndFilterOperator:
		return "and"
	case OrOperator, OrFilterOperator:
		return "or"
	}

	return ""
}

func writeString(buffer *bytes.Buffer, str string) {

	if _, err := buffer.WriteString(str); err != nil {
		panic(err)
	}
}
