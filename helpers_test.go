package manipulate

import (
	"net/http"
	"testing"

	"github.com/aporeto-inc/elemental"
	. "github.com/smartystreets/goconvey/convey"
)

// UserIdentity represents the Identity of the object
var TagIdentity = elemental.Identity{
	Name:     "tag",
	Category: "tag",
}

type Tag struct {
	ID string `cql:"id"`
}

func (t *Tag) Identifier() string {
	return t.ID
}

// Identity returns the Identity of the object.
func (t *Tag) Identity() elemental.Identity {

	return TagIdentity
}

// SetIdentifier sets the value of the object's unique identifier.
func (t *Tag) SetIdentifier(ID string) {
	t.ID = ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (t *Tag) Validate() error {
	return nil
}

func TestMethodConvertPointerArrayToManipulables(t *testing.T) {

	Convey("Given I create a two objects manipulator", t, func() {

		tag := &Tag{}
		tag2 := &Tag{}

		var listTags []*Tag
		listTags = append(listTags, tag, tag2)

		Convey("Then I call the method ConvertArrayToManipulables", func() {

			m := ConvertArrayToManipulables(listTags)
			So(m, ShouldHaveSameTypeAs, []Manipulable{})
			So(m[0], ShouldEqual, tag)
			So(m[1], ShouldEqual, tag2)
		})
	})
}

func TestAddQueryParameters(t *testing.T) {

	Convey("Given I create an new HTTP request with no query parameters", t, func() {

		request, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)

		Convey("When I call the method AddQueryParameters with nil", func() {

			AddQueryParameters(request, nil)

			Convey("The query string should be empty", func() {
				So(request.URL.RawQuery, ShouldEqual, "")
			})
		})

		Convey("When I call the method AddQueryParameters with parameters", func() {

			AddQueryParameters(request, map[string]string{
				"a":    "1",
				"b b":  "2 2",
				"c+&=": "3;:/",
			})
			Convey("The query string should be properly filled with escaped parameters", func() {
				So(request.URL.RawQuery, ShouldEqual, "a=1&b+b=2+2&c%2B%26%3D=3%3B%3A%2F")
			})
		})
	})

	Convey("Given I create an new HTTP request with existing query parameters", t, func() {

		request, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
		request.URL.RawQuery = "x=1&y=2"

		Convey("When I call the method AddQueryParameters with nil", func() {

			AddQueryParameters(request, nil)

			Convey("The query string should not change", func() {
				So(request.URL.RawQuery, ShouldEqual, "x=1&y=2")
			})
		})

		Convey("When I call the method AddQueryParameters with parameters", func() {

			AddQueryParameters(request, map[string]string{
				"a":    "1",
				"b b":  "2 2",
				"c+&=": "3;:/",
			})
			Convey("The query string should be properly appended with escaped parameters", func() {
				So(request.URL.RawQuery, ShouldEqual, "a=1&b+b=2+2&c%2B%26%3D=3%3B%3A%2F&x=1&y=2")
			})
		})
	})

	Convey("When I call AddQueryParameters with nil for request", t, func() {

		Convey("It should not panic", func() {
			So(func() {
				AddQueryParameters(nil, nil)
			}, ShouldNotPanic)
		})
	})
}
