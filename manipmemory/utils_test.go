package manipmemory

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_boolIndex(t *testing.T) {

	type testObject struct {
		somevalue      bool
		someothervalue string
	}

	Convey("When I call boolindex", t, func() {

		Convey("When I use a good data structure", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			val, err := boolIndex(t, "somevalue")
			So(err, ShouldBeNil)
			So(val, ShouldBeTrue)
		})

		Convey("When I use a good data structure with a bad field", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			_, err := boolIndex(t, "no value")
			So(err, ShouldNotBeNil)
		})

		Convey("When I use nil", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			_, err := boolIndex(t, "no value")
			So(err, ShouldNotBeNil)
		})

		Convey("When I use a bad type field", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			_, err := boolIndex(t, "somestring")
			So(err, ShouldNotBeNil)
		})
	})
}
