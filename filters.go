package manipulate

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// An FilterComparator is the type of a operator used by a filter.
type FilterComparator int

// FilterComparators are a list of FilterOperator.
type FilterComparators []FilterComparator

// Comparators represent various comparison operations.
const (
	EqualComparator FilterComparator = iota
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
	AndOperator FilterOperator = iota
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

// SubFilter is the type of subfilter
type SubFilter []*Filter

// SubFilters is is a list SubFilter,
type SubFilters []SubFilter

// FilterValueComposer adds values and operators.
type FilterValueComposer interface {
	Equals(interface{}) FilterKeyComposer
	NotEquals(interface{}) FilterKeyComposer
	GreaterThan(interface{}) FilterKeyComposer
	LesserThan(interface{}) FilterKeyComposer
	In(...interface{}) FilterKeyComposer
	NotIn(...interface{}) FilterKeyComposer
	Contains(...interface{}) FilterKeyComposer
	NotContains(...interface{}) FilterKeyComposer
	Matches(...interface{}) FilterKeyComposer
}

// FilterKeyComposer composes a filter.
type FilterKeyComposer interface {
	WithKey(string) FilterValueComposer

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
	ands        SubFilters
	ors         SubFilters
}

// NewFilter returns a new filter.
func NewFilter() *Filter {

	return &Filter{
		keys:        FilterKeys{},
		values:      FilterValues{},
		comparators: FilterComparators{},
		operators:   FilterOperators{},
		ands:        SubFilters{},
		ors:         SubFilters{},
	}
}

// NewFilterComposer returns a FilterComposer.
func NewFilterComposer() FilterKeyComposer {

	return NewFilter()
}

// NewFilterFromString returns a new filter computed from the given string.
func NewFilterFromString(filter string) (*Filter, error) {

	f, err := NewFilterParser(filter).Parse()
	if err != nil {
		return nil, err
	}

	return f.Done(), nil
}

// Keys returns the current keys.
func (f *Filter) Keys() FilterKeys {
	return append(FilterKeys{}, f.keys...)
}

// Values returns the current values.
func (f *Filter) Values() FilterValues {
	return append(FilterValues{}, f.values...)
}

// Operators returns the current operators.
func (f *Filter) Operators() FilterOperators {
	return append(FilterOperators{}, f.operators...)
}

// Comparators returns the current comparators.
func (f *Filter) Comparators() FilterComparators {
	return append(FilterComparators{}, f.comparators...)
}

// OrFilters returns the current ors sub filters.
func (f *Filter) OrFilters() SubFilters {
	return append(SubFilters{}, f.ors...)
}

// AndFilters returns the current and sub filters.
func (f *Filter) AndFilters() SubFilters {
	return append(SubFilters{}, f.ands...)
}

// Equals adds a an equality comparator to the FilterComposer.
func (f *Filter) Equals(value interface{}) FilterKeyComposer {
	f.values = f.values.add(value)
	f.comparators = f.comparators.add(EqualComparator)
	return f
}

// NotEquals adds a an non equality comparator to the FilterComposer.
func (f *Filter) NotEquals(value interface{}) FilterKeyComposer {
	f.values = f.values.add(value)
	f.comparators = f.comparators.add(NotEqualComparator)
	return f
}

// GreaterThan adds a greater than comparator to the FilterComposer.
func (f *Filter) GreaterThan(value interface{}) FilterKeyComposer {
	f.values = f.values.add(value)
	f.comparators = f.comparators.add(GreaterComparator)
	return f
}

// LesserThan adds a lesser than comparator to the FilterComposer.
func (f *Filter) LesserThan(value interface{}) FilterKeyComposer {
	f.values = f.values.add(value)
	f.comparators = f.comparators.add(LesserComparator)
	return f
}

// In adds a in comparator to the FilterComposer.
func (f *Filter) In(values ...interface{}) FilterKeyComposer {
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(InComparator)
	return f
}

// NotIn adds a not in comparator to the FilterComposer.
func (f *Filter) NotIn(values ...interface{}) FilterKeyComposer {
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(NotInComparator)
	return f
}

// Contains adds a contains comparator to the FilterComposer.
func (f *Filter) Contains(values ...interface{}) FilterKeyComposer {
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(ContainComparator)
	return f
}

// NotContains adds a contains comparator to the FilterComposer.
func (f *Filter) NotContains(values ...interface{}) FilterKeyComposer {
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(NotContainComparator)
	return f
}

// Matches adds a match comparator to the FilterComposer.
func (f *Filter) Matches(values ...interface{}) FilterKeyComposer {
	f.values = f.values.add(values...)
	f.comparators = f.comparators.add(MatchComparator)
	return f
}

// WithKey adds a key to FilterComposer.
func (f *Filter) WithKey(key string) FilterValueComposer {
	f.operators = append(f.operators, AndOperator)
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

		case AndOperator:
			writeString(&buffer, fmt.Sprintf(`%s`, f.keys[i]))
			writeString(&buffer, " ")
			writeString(&buffer, translateComparator(f.comparators[i]))
			writeString(&buffer, " ")
			writeString(&buffer, translateValue(f.comparators[i], f.values[i]))

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
		return "=="
	case NotEqualComparator:
		return "!="
	case GreaterComparator:
		return ">="
	case LesserComparator:
		return "<="
	case InComparator:
		return "in"
	case NotInComparator:
		return "not in"
	case ContainComparator:
		return "contains"
	case NotContainComparator:
		return "not contains"
	case MatchComparator:
		return "matches"
	default:
		panic(fmt.Sprintf("Unknown comparator: %d", comparator))
	}
}

func translateOperator(operator FilterOperator) string {

	switch operator {
	case AndOperator, AndFilterOperator:
		return "and"
	case OrFilterOperator:
		return "or"
	default:
		panic(fmt.Sprintf("Unknown operator: %d", operator))
	}
}

func translateValue(comparator FilterComparator, value interface{}) string {

	v := reflect.ValueOf(value)
	if comparator != ContainComparator && comparator != InComparator && comparator != MatchComparator {
		if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
			v = reflect.ValueOf(v.Index(0).Interface())
		}
	}

	switch v.Kind() {

	case reflect.String:
		return fmt.Sprintf(`"%s"`, v.Interface())

	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Int8, reflect.Uint, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uint8:
		return fmt.Sprintf(`%d`, v.Interface())

	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf(`%f`, v.Interface())

	case reflect.Bool:
		return fmt.Sprintf(`%t`, v.Interface())

	case reflect.Slice, reflect.Array:
		var final []string
		for i := 0; i < v.Len(); i++ {

			final = append(final, translateValue(comparator, v.Index(i).Interface()))
		}
		return fmt.Sprintf(`[%s]`, strings.Join(final, ", "))

	default:
		return fmt.Sprintf(`%v`, v.Interface())
	}
}

func writeString(buffer *bytes.Buffer, str string) {

	if _, err := buffer.WriteString(str); err != nil {
		panic(err)
	}
}
