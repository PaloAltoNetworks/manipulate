package maniphttp

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_addQueryParameters(t *testing.T) {

	createTestData := func() (ctx *manipulate.Context, benchmark string) {
		ctx = manipulate.NewContext()
		ctx.Parameters = url.Values{
			"a":    []string{"1"},
			"b b":  []string{"2 2"},
			"c+&=": []string{"3;:/"},
		}
		ctx.Page = 1
		ctx.PageSize = 100
		ctx.Recursive = true
		benchmark = "a=1&b+b=2+2&c%2B%26%3D=3%3B%3A%2F&page=1&pagesize=100&recursive=true"
		return
	}

	Convey("Given I create an new HTTP request with no query parameters", t, func() {

		request, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)

		Convey("When I call the method addQueryParameters with parameters", func() {

			ctx, benchmark := createTestData()
			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The query string should be properly filled with escaped parameters", func() {
				So(request.URL.RawQuery, ShouldEqual, benchmark)
			})
		})
	})

	Convey("Given I create an new HTTP request with existing query parameters", t, func() {

		request, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
		request.URL.RawQuery = "x=1&y=2"

		Convey("When I call the method addQueryParameters with a context with no Parameters", func() {

			err := addQueryParameters(request, manipulate.NewContext())
			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The query string should not change", func() {
				So(request.URL.RawQuery, ShouldEqual, "x=1&y=2")
			})
		})

		Convey("When I call the method addQueryParameters with a context with no KeyValues in Parameters", func() {

			ctx := manipulate.NewContext()
			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The query string should not change", func() {
				So(request.URL.RawQuery, ShouldEqual, "x=1&y=2")
			})
		})

		Convey("When I call the method addQueryParameters with parameters", func() {

			ctx, benchmark := createTestData()
			err := addQueryParameters(request, ctx)
			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The query string should be properly appended with escaped parameters", func() {
				So(request.URL.RawQuery, ShouldEqual, benchmark+"&x=1&y=2")
			})
		})

		Convey("When I call the method addQueryParameters with a filter", func() {

			ctx, benchmark := createTestData()
			ctx.Filter = manipulate.NewFilterComposer().WithKey("name").Equals("toto").AndKey("description").Equals("hello").Done()

			err := addQueryParameters(request, ctx)
			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The query string should be properly appended with escaped parameters", func() {
				So(request.URL.RawQuery, ShouldEqual, benchmark+"&tag=%24name%3Dtoto&tag=%24description%3Dhello&x=1&y=2")
			})
		})
	})
}
