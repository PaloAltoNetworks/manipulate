package manipulate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFilterComparator_NewFilterComparators(t *testing.T) {

	Convey("Given I create a new empty FilterComparator", t, func() {

		fc := NewFilterComparators()

		Convey("Then it should should be empty", func() {
			So(len(fc), ShouldEqual, 0)
		})
	})

	Convey("Given I create a empty FilterComparators with 2 comparators", t, func() {

		fc := NewFilterComparators(EqualComparator, InComparator)

		Convey("Then it should should not be empty", func() {
			So(len(fc), ShouldEqual, 2)
		})

		Convey("Then it should be correctly populated", func() {
			So(fc, ShouldResemble, FilterComparators{EqualComparator, InComparator})
		})

		Convey("When I use Then to add 2 other comparators", func() {

			fc = fc.Then(ContainComparator, GreaterComparator)

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

func TestFilterOperator_NewFilterOperators(t *testing.T) {

	Convey("Given I create a new empty FilterOperators", t, func() {

		fo := NewFilterOperators()

		Convey("Then it should should be empty", func() {
			So(len(fo), ShouldEqual, 0)
		})
	})

	Convey("Given I create a empty FilterOperators with 2 operators", t, func() {

		fo := NewFilterOperators(AndOperator, OrOperator)

		Convey("Then it should should not be empty", func() {
			So(len(fo), ShouldEqual, 2)
		})

		Convey("Then it should be correctly populated", func() {
			So(fo, ShouldResemble, FilterOperators{AndOperator, OrOperator})
		})

		Convey("When I use Then to add 2 other operators", func() {

			fo = fo.Then(OrOperator, AndOperator)

			Convey("Then it should be correctly populated", func() {
				So(fo, ShouldResemble, FilterOperators{
					AndOperator,
					OrOperator,
					OrOperator,
					AndOperator,
				})
			})
		})
	})
}

func TestFilterKey_NewFilterKeys(t *testing.T) {

	Convey("Given I create a new empty FilterKeys", t, func() {

		fk := NewFilterKeys()

		Convey("Then it should should be empty", func() {
			So(len(fk), ShouldEqual, 0)
		})
	})

	Convey("Given I create a empty FilterKeys with 2 fields", t, func() {

		fk := NewFilterKeys("id", "name")

		Convey("Then it should should not be empty", func() {
			So(len(fk), ShouldEqual, 1)
		})

		Convey("Then it should be correctly populated", func() {
			So(fk, ShouldResemble, FilterKeys{FilterKey{"id", "name"}})
		})

		Convey("When I use Then to add 2 other fields", func() {

			fk = fk.Then("value", "type")

			Convey("Then it should be correctly populated", func() {
				So(fk, ShouldResemble, FilterKeys{
					FilterKey{"id", "name"},
					FilterKey{"value", "type"},
				})
			})
		})
	})
}

func TestFilterKey_NewFilterValues(t *testing.T) {

	Convey("Given I create a new empty FilterValues", t, func() {

		fk := NewFilterValues()

		Convey("Then it should should be empty", func() {
			So(len(fk), ShouldEqual, 0)
		})
	})

	Convey("Given I create a empty FilterValues with 2 fields", t, func() {

		fk := NewFilterValues(0, "name")

		Convey("Then it should should not be empty", func() {
			So(len(fk), ShouldEqual, 1)
		})

		Convey("Then it should be correctly populated", func() {
			So(fk, ShouldResemble, FilterValues{FilterValue{0, "name"}})
		})

		Convey("When I use Then to add 2 other fields", func() {

			fk = fk.Then("value", false)

			Convey("Then it should be correctly populated", func() {
				So(fk, ShouldResemble, FilterValues{
					FilterValue{0, "name"},
					FilterValue{"value", false},
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
			So(f.keys, ShouldResemble, NewFilterKeys())
			So(f.values, ShouldResemble, NewFilterValues())
			So(f.operators, ShouldResemble, NewFilterOperators())
			So(f.comparators, ShouldResemble, NewFilterComparators())
		})
	})
}

func TestFilter_NewComposer(t *testing.T) {

	Convey("Given I create a new FilterComposer", t, func() {

		f := NewFilterComposer().Done()

		Convey("When I add the initial Equals statement", func() {

			f.WithKey("hello", "world").Equals(1, 2)

			Convey("Then the filter should be correctly populated", func() {
				So(f.keys, ShouldResemble, FilterKeys{FilterKey{"hello", "world"}})
				So(f.values, ShouldResemble, FilterValues{FilterValue{1, 2}})
				So(f.operators, ShouldResemble, FilterOperators{InitialOperator})
				So(f.comparators, ShouldResemble, FilterComparators{EqualComparator})
			})

			Convey("When I add a new GreaterThan statement", func() {

				f.AndKey("gt").GreaterThan(12)

				Convey("Then the filter should be correctly populated", func() {

					So(f.keys, ShouldResemble, FilterKeys{
						FilterKey{"hello", "world"},
						FilterKey{"gt"},
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
								FilterKey{"hello", "world"},
								FilterKey{"gt"},
								FilterKey{"lt"},
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
							So(f.String(), ShouldEqual, "[hello world] = [1 2] and [gt] >= [12] or [lt] <= [13]")
						})
					})
				})
			})

			Convey("When I add a new In statement", func() {

				f.AndKey("in").In("a", "b", "c")

				Convey("Then the filter should be correctly populated", func() {
					So(f.keys, ShouldResemble, FilterKeys{
						FilterKey{"hello", "world"},
						FilterKey{"in"},
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

						f.OrKey("ctn").Contains(false)

						Convey("Then the filter should be correctly populated", func() {
							So(f.keys, ShouldResemble, FilterKeys{
								FilterKey{"hello", "world"},
								FilterKey{"in"},
								FilterKey{"ctn"},
							})
							So(f.values, ShouldResemble, FilterValues{
								FilterValue{1, 2},
								FilterValue{"a", "b", "c"},
								FilterValue{false},
							})
							So(f.operators, ShouldResemble, FilterOperators{
								InitialOperator,
								AndOperator,
								OrOperator,
							})
							So(f.comparators, ShouldResemble, FilterComparators{
								EqualComparator,
								InComparator,
								ContainComparator,
							})

							Convey("Then the string representation should be correct", func() {
								So(f.String(), ShouldEqual, "[hello world] = [1 2] and [in] in [a b c] or [ctn] contains [false]")
							})
						})
					})
				})
			})
		})
	})
}
