package manipbolt

import (
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/asdine/storm/q"
)

type operation int

const (
	getOp operation = iota
	deleteOp
	countOp
)

var _ q.FieldMatcher = &containsOrEqualMatcher{}
var _ q.FieldMatcher = &regexpMatcher{}
var _ q.Matcher = &fieldMatcherCaseInsensitive{}

type containsOrEqualMatcher struct {
	field string
	value interface{}
}

func (c *containsOrEqualMatcher) MatchField(v interface{}) (bool, error) {

	ev := reflect.ValueOf(v)
	if ev.Kind() != reflect.Slice {
		eq, ok := q.Eq(c.field, c.value).(q.FieldMatcher)
		if !ok {
			return false, nil
		}

		return eq.MatchField(v)
	}

	for i := 0; i < ev.Len(); i++ {
		e := ev.Index(i).Interface()

		if e == c.value {
			return true, nil
		}
	}

	return false, nil
}

// containsOrEqual implements q.FieldMatcher interface.
// It combines both equal and contains matcher in one.
// Refer corresponding unit tests for details.
func containsOrEqual(field string, v interface{}) q.Matcher {
	return newFieldMatcherCaseInsensitive(field, &containsOrEqualMatcher{field, v})
}

var regexpCache = struct {
	sync.RWMutex
	m map[string]*regexp.Regexp
}{m: make(map[string]*regexp.Regexp)}

type regexpMatcher struct {
	r   *regexp.Regexp
	err error
}

func (r *regexpMatcher) MatchField(v interface{}) (bool, error) {

	if r.err != nil {
		return false, r.err
	}

	return r.r.MatchString(v.(string)), nil
}

// regexMatcher creates a regexp matcher. It checks if the
// given field matches the given regexp. Note that this only
// supports fields of type string. Field is case insensitive.
func regexMatcher(field string, re string) q.Matcher {

	regexpCache.RLock()
	if r, ok := regexpCache.m[re]; ok {
		regexpCache.RUnlock()
		return newFieldMatcherCaseInsensitive(field, &regexpMatcher{r, nil})
	}
	regexpCache.RUnlock()

	regexpCache.Lock()
	r, err := regexp.Compile(re)
	if err == nil {
		regexpCache.m[re] = r
	}
	regexpCache.Unlock()

	return newFieldMatcherCaseInsensitive(field, &regexpMatcher{r, err})
}

type fieldMatcherCaseInsensitive struct {
	q.FieldMatcher
	field string
}

func (r fieldMatcherCaseInsensitive) Match(i interface{}) (bool, error) {

	v := reflect.Indirect(reflect.ValueOf(i))

	field := v.FieldByNameFunc(func(n string) bool {
		return strings.EqualFold(n, r.field)
	})

	if !field.IsValid() {
		return false, q.ErrUnknownField
	}

	return r.MatchField(field.Interface())
}

// newFieldMatcherCaseInsensitive creates a case insensitive
// Matcher for a given field.
func newFieldMatcherCaseInsensitive(field string, fm q.FieldMatcher) q.Matcher {
	return fieldMatcherCaseInsensitive{fm, field}
}
