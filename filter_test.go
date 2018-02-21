package manipulate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFilterComparator_NewFilterComparators(t *testing.T) {

	Convey("Given I create a empty FilterComparators with 2 comparators", t, func() {

		fc := FilterComparators{EqualComparator, InComparator}

		Convey("Then it should should not be empty", func() {
			So(len(fc), ShouldEqual, 2)
		})

		Convey("Then it should be correctly populated", func() {
			So(fc, ShouldResemble, FilterComparators{EqualComparator, InComparator})
		})

		Convey("When I use Then to add 2 other comparators", func() {

			fc = fc.add(ContainComparator, GreaterComparator)

			Convey("Then it should be correctly populated", func() {
				So(fc, ShouldResemble, FilterComparators{
					EqualComparator,
					InComparator,
					ContainComparator,
					GreaterComparator,
				})
			})
		})
	})
}

func TestFilter_NewFilter(t *testing.T) {

	Convey("Given I create a new filter", t, func() {

		f := NewFilter()

		Convey("Then the filter should implement the correct interfaces", func() {
			So(f, ShouldImplement, (*FilterKeyComposer)(nil))
			So(f, ShouldImplement, (*FilterValueComposer)(nil))
		})

		Convey("Then all values should be initialized", func() {
			So(f.keys, ShouldResemble, FilterKeys{})
			So(f.values, ShouldResemble, FilterValues{})
			So(f.operators, ShouldResemble, FilterOperators{})
			So(f.comparators, ShouldResemble, FilterComparators{})
			So(f.ands, ShouldResemble, [][]*Filter(nil))
			So(f.ors, ShouldResemble, [][]*Filter(nil))
		})
	})
}

func TestFilter_NewComposer(t *testing.T) {

	Convey("Given I create a new FilterComposer", t, func() {

		f := NewFilterComposer().Done()

		Convey("When I add the initial Equals statement", func() {

			f.WithKey("hello").Equals(1, 2)

			Convey("Then the filter should be correctly populated", func() {
				So(f.keys, ShouldResemble, FilterKeys{"hello"})
				So(f.values, ShouldResemble, FilterValues{FilterValue{1, 2}})
				So(f.operators, ShouldResemble, FilterOperators{InitialOperator})
				So(f.comparators, ShouldResemble, FilterComparators{EqualComparator})
			})

			Convey("When I add a new GreaterThan statement", func() {

				f.AndKey("gt").GreaterThan(12)

				Convey("Then the filter should be correctly populated", func() {

					So(f.keys, ShouldResemble, FilterKeys{
						"hello",
						"gt",
					})
					So(f.values, ShouldResemble, FilterValues{
						FilterValue{1, 2},
						FilterValue{12},
					})
					So(f.operators, ShouldResemble, FilterOperators{
						InitialOperator,
						AndOperator,
					})
					So(f.comparators, ShouldResemble, FilterComparators{
						EqualComparator,
						GreaterComparator,
					})

					Convey("When I add a new LesserThan statement", func() {

						f.OrKey("lt").LesserThan(13)

						Convey("Then the filter should be correctly populated", func() {
							So(f.keys, ShouldResemble, FilterKeys{
								"hello",
								"gt",
								"lt",
							})
							So(f.values, ShouldResemble, FilterValues{
								FilterValue{1, 2},
								FilterValue{12},
								FilterValue{13},
							})
							So(f.operators, ShouldResemble, FilterOperators{
								InitialOperator,
								AndOperator,
								OrOperator,
							})
							So(f.comparators, ShouldResemble, FilterComparators{
								EqualComparator,
								GreaterComparator,
								LesserComparator,
							})
						})

						Convey("Then the string representation should be correct", func() {
							So(f.String(), ShouldEqual, "hello = [1 2] and gt >= [12] or lt <= [13]")
						})
					})
				})
			})

			Convey("When I add a new In statement", func() {

				f.AndKey("in").In("a", "b", "c")

				Convey("Then the filter should be correctly populated", func() {
					So(f.keys, ShouldResemble, FilterKeys{
						"hello",
						"in",
					})
					So(f.values, ShouldResemble, FilterValues{
						FilterValue{1, 2},
						FilterValue{"a", "b", "c"},
					})
					So(f.operators, ShouldResemble, FilterOperators{
						InitialOperator,
						AndOperator,
					})
					So(f.comparators, ShouldResemble, FilterComparators{
						EqualComparator,
						InComparator,
					})

					Convey("When I add a new Contains statement", func() {

						f.AndKey("ctn").Contains(false)

						Convey("Then the filter should be correctly populated", func() {
							So(f.keys, ShouldResemble, FilterKeys{
								"hello",
								"in",
								"ctn",
							})
							So(f.values, ShouldResemble, FilterValues{
								FilterValue{1, 2},
								FilterValue{"a", "b", "c"},
								FilterValue{false},
							})
							So(f.operators, ShouldResemble, FilterOperators{
								InitialOperator,
								AndOperator,
								AndOperator,
							})
							So(f.comparators, ShouldResemble, FilterComparators{
								EqualComparator,
								InComparator,
								ContainComparator,
							})

							Convey("Then the string representation should be correct", func() {
								So(f.String(), ShouldEqual, "hello = [1 2] and in in [a b c] and ctn contains [false]")
							})
						})
					})
				})
			})

			Convey("When I add a new difference comparator", func() {
				f.AndKey("x").NotEquals(true)

				Convey("Then the filter should be correctly populated", func() {
					So(f.keys, ShouldResemble, FilterKeys{
						"hello",
						"x",
					})
					So(f.values, ShouldResemble, FilterValues{
						FilterValue{1, 2},
						FilterValue{true},
					})
					So(f.operators, ShouldResemble, FilterOperators{
						InitialOperator,
						AndOperator,
					})
					So(f.comparators, ShouldResemble, FilterComparators{
						EqualComparator,
						NotEqualComparator,
					})
					So(f.String(), ShouldEqual, "hello = [1 2] and x != [true]")
				})
			})
		})
	})
}

func TestFilter_SubFilters(t *testing.T) {

	Convey("Given I have a composed filters", t, func() {

		f := NewFilterComposer().
			WithKey("namespace").Equals("coucou").
			And(
				NewFilterComposer().
					WithKey("name").Equals("toto").
					AndKey("surname").Equals("titi").
					Done(),
				NewFilterComposer().
					WithKey("color").Equals("blue").
					Or(
						NewFilterComposer().
							WithKey("size").Equals("big").
							Done(),
						NewFilterComposer().
							WithKey("size").Equals("medium").
							Done(),
					).
					Done(),
			).
			Done()

		Convey("When I call string it should be correct ", func() {

			So(f.String(), ShouldEqual, "namespace = [coucou] and ((name = [toto] and surname = [titi]) and (color = [blue] or ((size = [big]) or (size = [medium]))))")
		})
	})
}
