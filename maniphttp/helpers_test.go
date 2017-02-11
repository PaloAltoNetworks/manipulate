package maniphttp

import (
	"net/http"
	"testing"

	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_addQueryParameters(t *testing.T) {

	createTestData := func() (ctx *manipulate.Context, benchmark string) {
		ctx = manipulate.NewContext()
		ctx.Parameters = &manipulate.Parameters{}
		ctx.Parameters.KeyValues = map[string]string{
			"a":    "1",
			"b b":  "2 2",
			"c+&=": "3;:/",
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
			addQueryParameters(request, ctx)
			Convey("The query string should be properly filled with escaped parameters", func() {
				So(request.URL.RawQuery, ShouldEqual, benchmark)
			})
		})
	})

	Convey("Given I create an new HTTP request with existing query parameters", t, func() {

		request, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
		request.URL.RawQuery = "x=1&y=2"

		Convey("When I call the method addQueryParameters with a context with no Parameters", func() {

			addQueryParameters(request, manipulate.NewContext())

			Convey("The query string should not change", func() {
				So(request.URL.RawQuery, ShouldEqual, "x=1&y=2")
			})
		})

		Convey("When I call the method addQueryParameters with a context with no KeyValues in Parameters", func() {

			ctx := manipulate.NewContext()
			ctx.Parameters = &manipulate.Parameters{}
			addQueryParameters(request, ctx)

			Convey("The query string should not change", func() {
				So(request.URL.RawQuery, ShouldEqual, "x=1&y=2")
			})
		})

		Convey("When I call the method addQueryParameters with parameters", func() {

			ctx, benchmark := createTestData()
			addQueryParameters(request, ctx)
			Convey("The query string should be properly appended with escaped parameters", func() {
				So(request.URL.RawQuery, ShouldEqual, benchmark+"&x=1&y=2")
			})
		})
	})
}
