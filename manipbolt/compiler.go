package manipbolt

import (
	"fmt"

	"github.com/asdine/storm/q"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

func compileFilter(f *elemental.Filter) (q.Matcher, error) {

	if len(f.Operators()) == 0 {
		return q.And(), nil
	}

	matchers := []q.Matcher{}

	for i, operator := range f.Operators() {

		switch operator {

		case elemental.AndOperator:

			k := f.Keys()[i]
			values := f.Values()[i]
			items := []q.Matcher{}

			switch f.Comparators()[i] {

			case elemental.EqualComparator:

				items = append(items, containsOrEqual(k, values[0]))

			case elemental.MatchComparator:

				subs := []q.Matcher{}
				for _, value := range values {

					v, ok := value.(string)
					if !ok {
						return nil, manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("regex only supports string: %v", value)}
					}

					subs = append(subs, regexMatcher(k, v))
				}

				items = append(items, q.Or(subs...))

			case elemental.ContainComparator:

				subs := []q.Matcher{}
				for _, value := range values {
					subs = append(subs, containsOrEqual(k, value))
				}

				items = append(items, q.Or(subs...))

			default:
				return nil, manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("invalid comparator: %d", f.Comparators()[i])}
			}

			matchers = append(matchers, items...)

		case elemental.AndFilterOperator:

			subs := []q.Matcher{}
			for _, sub := range f.AndFilters()[i] {

				matcher, err := compileFilter(sub)
				if err != nil {
					return nil, err
				}

				subs = append(subs, matcher)
			}

			matchers = append(matchers, q.And(subs...))

		case elemental.OrFilterOperator:

			subs := []q.Matcher{}
			for _, sub := range f.OrFilters()[i] {

				matcher, err := compileFilter(sub)
				if err != nil {
					return nil, err
				}

				subs = append(subs, matcher)
			}

			matchers = append(matchers, q.Or(subs...))

		default:
			return nil, manipulate.ErrCannotBuildQuery{Err: fmt.Errorf("invalid operator: %d", operator)}
		}
	}

	return q.And(matchers...), nil
}
