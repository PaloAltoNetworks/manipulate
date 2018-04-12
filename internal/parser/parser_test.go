package parser

import (
	"testing"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParser_Operators(t *testing.T) {

	Convey("Given the operator: '=='", t, func() {
		parser := NewFilterParser("key == value")
		expectedFilter := manipulate.NewFilterComposer().WithKey("key").Equals("value").Done()

		Convey("When I run a parse", func() {
			filter, err := parser.Parse()

			Convey("Then there should be no error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the filter should match our expectations", func() {
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator '!='", t, func() {
		parser := NewFilterParser("key != value")
		expectedFilter := manipulate.NewFilterComposer().WithKey("key").NotEquals("value").Done()

		Convey("When I run a parse", func() {
			filter, err := parser.Parse()

			Convey("Then there should be no error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the filter should match our expectations", func() {
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'contains'", t, func() {
		parser := NewFilterParser(`key contains ["value1", "value2"]`)
		expectedFilter := manipulate.NewFilterComposer().WithKey("key").Contains("value1", "value2").Done()

		Convey("When I run a parse", func() {
			filter, err := parser.Parse()

			Convey("Then there should be no error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the filter should match our expectations", func() {
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'not contains'", t, func() {
		parser := NewFilterParser(`key not contains ["value1", "value2"]`)
		expectedFilter := manipulate.NewFilterComposer().WithKey("key").NotContains("value1", "value2").Done()

		Convey("When I run a parse", func() {
			filter, err := parser.Parse()

			Convey("Then there should be no error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the filter should match our expectations", func() {
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'in'", t, func() {
		parser := NewFilterParser(`key in ["value1", "value2"]`)
		expectedFilter := manipulate.NewFilterComposer().WithKey("key").In("value1", "value2").Done()

		Convey("When I run a parse", func() {
			filter, err := parser.Parse()

			Convey("Then there should be no error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the filter should match our expectations", func() {
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'not in'", t, func() {
		parser := NewFilterParser(`key not in ["value1", "value2"]`)
		expectedFilter := manipulate.NewFilterComposer().WithKey("key").NotIn("value1", "value2").Done()

		Convey("When I run a parse", func() {
			filter, err := parser.Parse()

			Convey("Then there should be no error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the filter should match our expectations", func() {
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the operator 'matches'", t, func() {
		parser := NewFilterParser(`key matches ["value1", "value2"]`)
		expectedFilter := manipulate.NewFilterComposer().WithKey("key").Matches("value1", "value2").Done()

		Convey("When I run a parse", func() {
			filter, err := parser.Parse()

			Convey("Then there should be no error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the filter should match our expectations", func() {
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

}

func TestParser_Values(t *testing.T) {

	Convey("Given the value: 'hello world'", t, func() {
		parser := NewFilterParser("key == \"hello world\"")
		expectedFilter := manipulate.NewFilterComposer().WithKey("key").Equals("hello world").Done()

		Convey("When I run a parse", func() {
			filter, err := parser.Parse()

			Convey("Then there should be no error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the filter should match our expectations", func() {
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

	Convey("Given the value: 'hello\"word'", t, func() {
		parser := NewFilterParser("key == \"hello\\\"world\"")
		expectedFilter := manipulate.NewFilterComposer().WithKey("key").Equals("hello\"world").Done()

		Convey("When I run a parse", func() {
			filter, err := parser.Parse()

			Convey("Then there should be no error", func() {
				So(err, ShouldEqual, nil)
			})

			Convey("Then the filter should match our expectations", func() {
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})
	})

}

func Test_Parser(t *testing.T) {

	// Valid cases
	Convey("Given the filter: namespace == chris and test == true", t, func() {
		parser := NewFilterParser("namespace == chris and test == true")
		filter, err := parser.Parse()

		expectedFilter := manipulate.NewFilterComposer().And(
			manipulate.NewFilterComposer().WithKey("namespace").Equals("chris").Done(),
			manipulate.NewFilterComposer().WithKey("test").Equals(true).Done(),
		).Done()

		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	Convey("Given the filter: namespace == chris and test == true or value > 30", t, func() {
		parser := NewFilterParser("namespace == chris and test == true or value > 30")
		filter, err := parser.Parse()

		expectedFilter := manipulate.NewFilterComposer().Or(
			manipulate.NewFilterComposer().And(
				manipulate.NewFilterComposer().WithKey("namespace").Equals("chris").Done(),
				manipulate.NewFilterComposer().WithKey("test").Equals(true).Done(),
			).Done(),
			manipulate.NewFilterComposer().WithKey("value").GreaterThan(30).Done(),
		).Done()

		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	Convey("Given the filter: namespace == chris or test == true and value > 30", t, func() {
		parser := NewFilterParser("namespace == chris or test == true and value > 30")
		filter, err := parser.Parse()

		expectedFilter := manipulate.NewFilterComposer().And(
			manipulate.NewFilterComposer().Or(
				manipulate.NewFilterComposer().WithKey("namespace").Equals("chris").Done(),
				manipulate.NewFilterComposer().WithKey("test").Equals(true).Done(),
			).Done(),
			manipulate.NewFilterComposer().WithKey("value").GreaterThan(30).Done(),
		).Done()

		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	Convey(`Given the filter: "namespace"=="chris" and "test"== true`, t, func() {
		parser := NewFilterParser(`"namespace"=="chris" and "test"== true`)
		filter, err := parser.Parse()

		expectedFilter := manipulate.NewFilterComposer().And(
			manipulate.NewFilterComposer().WithKey("namespace").Equals("chris").Done(),
			manipulate.NewFilterComposer().WithKey("test").Equals(true).Done(),
		).Done()

		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	Convey(`Given the filter: "namespace" == "chris" and "test" == true`, t, func() {
		parser := NewFilterParser(`"namespace" == "chris" and "test" == true`)
		filter, err := parser.Parse()

		expectedFilter := manipulate.NewFilterComposer().And(
			manipulate.NewFilterComposer().WithKey("namespace").Equals("chris").Done(),
			manipulate.NewFilterComposer().WithKey("test").Equals(true).Done(),
		).Done()

		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	Convey(`Given the filter: "age" <= 32 or "age" > 50`, t, func() {
		parser := NewFilterParser(`"age" <= 32 or "age" > 50`)
		filter, err := parser.Parse()

		expectedFilter := manipulate.NewFilterComposer().Or(
			manipulate.NewFilterComposer().WithKey("age").LesserThan(32).Done(),
			manipulate.NewFilterComposer().WithKey("age").GreaterThan(50).Done(),
		).Done()

		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	Convey(`Given the filter: ("file" matches "*.txt" and "file" contains "search")`, t, func() {
		parser := NewFilterParser(`("file" matches "*.txt" and "file" contains "search")`)
		filter, err := parser.Parse()

		expectedFilter := manipulate.NewFilterComposer().And(
			manipulate.NewFilterComposer().WithKey("file").Matches("*.txt").Done(),
			manipulate.NewFilterComposer().WithKey("file").Contains("search").Done(),
		).Done()

		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	Convey(`Given the filter: "namespace" == "/chris"`, t, func() {
		parser := NewFilterParser(`"namespace" == "/chris"`)

		expectedFilter := manipulate.NewFilterComposer().WithKey("namespace").Equals("/chris").Done()

		filter, err := parser.Parse()
		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	Convey(`Given the filter: "namespace" == "/chris" and test == true and ("name" == toto or "name" == tata)`, t, func() {
		parser := NewFilterParser(`"namespace" == "/chris" and test == true ("name" == toto or "name" == tata)`)

		expectedFilter := manipulate.NewFilterComposer().And(
			manipulate.NewFilterComposer().WithKey("namespace").Equals("/chris").Done(),
			manipulate.NewFilterComposer().WithKey("test").Equals(true).Done(),
			manipulate.NewFilterComposer().Or(
				manipulate.NewFilterComposer().WithKey("name").Equals("toto").Done(),
				manipulate.NewFilterComposer().WithKey("name").Equals("tata").Done(),
			).Done(),
		).Done()

		filter, err := parser.Parse()
		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	Convey(`Given the filter: (name == toto or name == tata) and age == 31`, t, func() {
		parser := NewFilterParser("(name == toto or name == tata) and age == 31")

		expectedFilter := manipulate.NewFilterComposer().And(
			manipulate.NewFilterComposer().Or(
				manipulate.NewFilterComposer().WithKey("name").Equals("toto").Done(),
				manipulate.NewFilterComposer().WithKey("name").Equals("tata").Done(),
			).Done(),
			manipulate.NewFilterComposer().WithKey("age").Equals(31).Done(),
		).Done()

		filter, err := parser.Parse()
		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	Convey(`Given the filter: (name == toto and name == tata) or (age == 31 and age == 32)`, t, func() {
		parser := NewFilterParser("(name == toto and name == tata) or (age == 31 and age == 32)")

		expectedFilter := manipulate.NewFilterComposer().Or(
			manipulate.NewFilterComposer().And(
				manipulate.NewFilterComposer().WithKey("name").Equals("toto").Done(),
				manipulate.NewFilterComposer().WithKey("name").Equals("tata").Done(),
			).Done(),
			manipulate.NewFilterComposer().And(
				manipulate.NewFilterComposer().WithKey("age").Equals(31).Done(),
				manipulate.NewFilterComposer().WithKey("age").Equals(32).Done(),
			).Done(),
		).Done()

		filter, err := parser.Parse()
		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	Convey(`Given the filter: name == toto and value == 38.9000 and ((name == toto and name == tata) or (age == 31 and age == 32) or protected contains [true, false])`, t, func() {
		parser := NewFilterParser("name == toto and value == 38.9000 and ((name == toto and name == tata) or (age == 31 and age == 32) or protected contains [true, false])")

		expectedFilter := manipulate.NewFilterComposer().And(
			manipulate.NewFilterComposer().WithKey("name").Equals("toto").Done(),
			manipulate.NewFilterComposer().WithKey("value").Equals(38.9).Done(),
			manipulate.NewFilterComposer().Or(
				manipulate.NewFilterComposer().And(
					manipulate.NewFilterComposer().WithKey("name").Equals("toto").Done(),
					manipulate.NewFilterComposer().WithKey("name").Equals("tata").Done(),
				).Done(),
				manipulate.NewFilterComposer().And(
					manipulate.NewFilterComposer().WithKey("age").Equals(31).Done(),
					manipulate.NewFilterComposer().WithKey("age").Equals(32).Done(),
				).Done(),
				manipulate.NewFilterComposer().WithKey("protected").In(true, false).Done(),
			).Done(),
		).Done()

		filter, err := parser.Parse()
		So(err, ShouldEqual, nil)
		So(filter, ShouldNotEqual, nil)
		So(filter.String(), ShouldEqual, expectedFilter.String())
	})

	// Error cases
	Convey(`Given the filter: "namespace == chris`, t, func() {
		parser := NewFilterParser(`"namespace == chris`)
		_, err := parser.Parse()

		So(err, ShouldNotEqual, nil)
		So(err.Error(), ShouldEqual, `missing quote after the word namespace`)
	})

	Convey(`Given the filter: namespace" == chris`, t, func() {
		parser := NewFilterParser(`namespace" == chris`)
		_, err := parser.Parse()

		So(err, ShouldNotEqual, nil)
		So(err.Error(), ShouldEqual, `invalid operator. found "`)
	})

	Convey(`Given the filter: namespace == "chris`, t, func() {
		parser := NewFilterParser(`namespace == "chris`)
		_, err := parser.Parse()

		So(err, ShouldNotEqual, nil)
		So(err.Error(), ShouldEqual, `missing quote after the value chris`)
	})

	Convey(`Given the filter: namespace == chris"`, t, func() {
		parser := NewFilterParser(`namespace == chris"`)
		_, err := parser.Parse()

		So(err, ShouldNotEqual, nil)
		So(err.Error(), ShouldEqual, `missing quote before the value chris`)
	})

	Convey(`Given the filter: namespace and chris"`, t, func() {
		parser := NewFilterParser(`namespace and chris"`)
		_, err := parser.Parse()

		So(err, ShouldNotEqual, nil)
		So(err.Error(), ShouldEqual, `invalid operator. found and`)
	})

	Convey(`Given the filter: namespace == and"`, t, func() {
		parser := NewFilterParser(`namespace == and"`)
		_, err := parser.Parse()

		So(err, ShouldNotEqual, nil)
		So(err.Error(), ShouldEqual, `invalid value. found and`)
	})

	Convey(`Given the filter: namespace==toto`, t, func() {
		parser := NewFilterParser("namespace==toto")
		_, err := parser.Parse()

		So(err, ShouldNotEqual, nil)
		So(err.Error(), ShouldContainSubstring, `invalid operator`) // Note: Not sure about this case.
	})
}

func Test_isLetter(t *testing.T) {
	Convey("Given I have a FilterParser", t, func() {
		So(isLetter('<'), ShouldEqual, true)
		So(isLetter('b'), ShouldEqual, true)
		So(isLetter(4), ShouldEqual, false)
	})
}

func TestParser_AdvancedFilter(t *testing.T) {

	advancedFilter := `"namespace" == "coucou" and "number" == 32.900000 and (("name" == "toto" and "value" == 1) and ("color" in ["red", "green", "blue", 43] and "something" in ["stuff"] or (("size" matches [".*"]) or ("size" == "medium" and "fat" == false) or ("size" in [true, false]))))`

	Convey(`Given I have an advanced complex filter and a parser`, t, func() {
		parser := NewFilterParser(advancedFilter)

		expectedFilter := manipulate.NewFilterComposer().And(
			manipulate.NewFilterComposer().WithKey("namespace").Equals("coucou").Done(),
			manipulate.NewFilterComposer().WithKey("number").Equals(32.9).Done(),
			manipulate.NewFilterComposer().And(

				manipulate.NewFilterComposer().And(
					manipulate.NewFilterComposer().WithKey("name").Equals("toto").Done(),
					manipulate.NewFilterComposer().WithKey("value").Equals(1).Done(),
				).Done(),

				manipulate.NewFilterComposer().Or(
					manipulate.NewFilterComposer().And(
						manipulate.NewFilterComposer().WithKey("color").Contains("red", "green", "blue", 43).Done(),
						manipulate.NewFilterComposer().WithKey("something").Contains("stuff").Done(),
					).Done(),

					manipulate.NewFilterComposer().Or(
						manipulate.NewFilterComposer().WithKey("size").Matches(".*").Done(),
						manipulate.NewFilterComposer().And(
							manipulate.NewFilterComposer().WithKey("size").Equals("medium").Done(),
							manipulate.NewFilterComposer().WithKey("fat").Equals(false).Done(),
						).Done(),
						manipulate.NewFilterComposer().
							WithKey("size").In(true, false).
							Done(),
					).Done(),
				).Done(),
			).Done(),
		).Done()

		Convey("When I run parse", func() {
			filter, err := parser.Parse()

			Convey("Then err should be nil", func() {
				So(err, ShouldEqual, nil)
			})
			Convey("Then the filter should be correct", func() {
				So(filter, ShouldNotEqual, nil)
				So(filter.String(), ShouldEqual, expectedFilter.String())
			})
		})

		Convey("When I chain multiple parse", func() {
			filter, err := parser.Parse()
			So(err, ShouldEqual, nil)

			p := NewFilterParser(filter.String())
			f, err := p.Parse()

			Convey("Then err should be nil", func() {
				So(err, ShouldEqual, nil)
			})
			Convey("Then the filter should be equivalent to the advanced filter", func() {
				So(f, ShouldNotEqual, nil)
				So(f.String(), ShouldEqual, filter.String())
			})
		})

	})

}
