package maniphttp

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_addQueryParameters(t *testing.T) {

	Convey("Given I have a request and a context", t, func() {

		request, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
		ctx := manipulate.NewContext()

		Convey("When I call the method addQueryParameters with parameters", func() {

			ctx.Parameters = url.Values{
				"a":    []string{"1"},
				"b b":  []string{"2 2"},
				"c+&=": []string{"3;:/"},
			}

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The query string should be properly filled with escaped parameters", func() {
				So(request.URL.RawQuery, ShouldEqual, "a=1&b+b=2+2&c%2B%26%3D=3%3B%3A%2F")
			})
		})

		Convey("When I call the method addQueryParameters with a filter", func() {

			ctx.Filter = manipulate.NewFilterComposer().WithKey("name").Equals("toto").WithKey("description").Equals("hello").Done()

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the query string should be correct ", func() {
				So(request.URL.RawQuery, ShouldEqual, "tag=%24name%3Dtoto&tag=%24description%3Dhello")
			})
		})

		Convey("When I call the method addQueryParameters with a order", func() {

			ctx.Order = []string{"a", "b", "c"}

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the query string should be correct ", func() {
				So(request.URL.RawQuery, ShouldEqual, "order=a&order=b&order=c")
			})
		})

		Convey("When I call the method addQueryParameters with a overide protection", func() {

			ctx.OverrideProtection = true

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the query string should be correct ", func() {
				So(request.URL.RawQuery, ShouldEqual, "override=true")
			})
		})

		Convey("When I call the method addQueryParameters with a pagination", func() {

			ctx.Page = 12
			ctx.PageSize = 42

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the query string should be correct ", func() {
				So(request.URL.RawQuery, ShouldEqual, "page=12&pagesize=42")
			})
		})

		Convey("When I call the method addQueryParameters with a recursive", func() {

			ctx.Recursive = true

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the query string should be correct ", func() {
				So(request.URL.RawQuery, ShouldEqual, "recursive=true")
			})
		})
	})

	Convey("Given I create an new HTTP request with existing query parameters", t, func() {

		request, _ := http.NewRequest(http.MethodGet, "http://test.com?x=1&y=2", nil)

		Convey("When I call the method addQueryParameters with a context with no Parameters", func() {

			err := addQueryParameters(request, manipulate.NewContext())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The query string should not change", func() {
				So(request.URL.RawQuery, ShouldEqual, "x=1&y=2")
			})
		})
	})
}
