package maniphttp

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"go.aporeto.io/manipulate"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_addQueryParameters(t *testing.T) {

	Convey("Given I have a request and a context", t, func() {

		request, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)

		Convey("When I call the method addQueryParameters with parameters", func() {

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParameters(
					url.Values{
						"a":    []string{"1"},
						"b b":  []string{"2 2"},
						"c+&=": []string{"3;:/"},
					},
				),
			)

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The query string should be properly filled with escaped parameters", func() {
				So(request.URL.RawQuery, ShouldEqual, "a=1&b+b=2+2&c%2B%26%3D=3%3B%3A%2F")
			})
		})

		Convey("When I call the method addQueryParameters with a filter", func() {

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionFilter(
					manipulate.NewFilterComposer().WithKey("name").Equals("toto").WithKey("description").Equals("hello").Done(),
				),
			)

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the query string should be correct ", func() {
				So(request.URL.RawQuery, ShouldEqual, "q=name+%3D%3D+%22toto%22+and+description+%3D%3D+%22hello%22")
			})
		})

		Convey("When I call the method addQueryParameters with a order", func() {

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionOrder("a", "b", "c"),
			)

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the query string should be correct ", func() {
				So(request.URL.RawQuery, ShouldEqual, "order=a&order=b&order=c")
			})
		})

		Convey("When I call the method addQueryParameters with a overide protection", func() {

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionOverride(true),
			)

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the query string should be correct ", func() {
				So(request.URL.RawQuery, ShouldEqual, "override=true")
			})
		})

		Convey("When I call the method addQueryParameters with a pagination", func() {

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionPage(12, 42),
			)

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the query string should be correct ", func() {
				So(request.URL.RawQuery, ShouldEqual, "page=12&pagesize=42")
			})
		})

		Convey("When I call the method addQueryParameters with a recursive", func() {

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionRecursive(true),
			)

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

			err := addQueryParameters(request, manipulate.NewContext(context.Background()))

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("The query string should not change", func() {
				So(request.URL.RawQuery, ShouldEqual, "x=1&y=2")
			})
		})
	})
}

type fakeReader struct{}

func (r *fakeReader) Read(p []byte) (n int, err error) { return 0, errors.New("boom") }

func Test_decodeData(t *testing.T) {

	Convey("Given I have valid json data in a reader", t, func() {

		r := &http.Response{
			Body: ioutil.NopCloser(bytes.NewBuffer([]byte(`{"name":"thename","age": 2}`))),
		}

		Convey("When I call decodeData", func() {

			dest := map[string]interface{}{}
			err := decodeData(r, &dest)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the dest should be correct", func() {
				So(len(dest), ShouldEqual, 2)
				So(dest["name"].(string), ShouldEqual, "thename")
				So(dest["age"].(float64), ShouldEqual, 2)
			})
		})
	})

	Convey("Given I have valid json gzipped data in a reader", t, func() {

		buf := bytes.NewBuffer(nil)
		zw := gzip.NewWriter(buf)
		_, err := zw.Write([]byte(`{"name":"thename","age": 2}`))
		if err != nil {
			panic(err)
		}
		zw.Close() // nolint

		r := &http.Response{
			Body: ioutil.NopCloser(buf),
			Header: http.Header{
				"Content-Encoding": []string{"gzip"},
			},
		}

		Convey("When I call decodeData", func() {

			dest := map[string]interface{}{}
			err := decodeData(r, &dest)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the dest should be correct", func() {
				So(len(dest), ShouldEqual, 2)
				So(dest["name"].(string), ShouldEqual, "thename")
				So(dest["age"].(float64), ShouldEqual, 2)
			})
		})
	})

	Convey("Given I have invalid valid json data in a reader", t, func() {

		r := &http.Response{
			Body: ioutil.NopCloser(bytes.NewBuffer([]byte(`<html><body>not json</body></html>`))),
		}

		Convey("When I call decodeData", func() {

			dest := map[string]interface{}{}
			err := decodeData(r, &dest)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Unable to unmarshal data: invalid character '<' looking for beginning of value. original data:\n<html><body>not json</body></html>")
			})

			Convey("Then the dest should be empty", func() {
				So(len(dest), ShouldEqual, 0)
			})
		})
	})

	Convey("Given I have a nil reader", t, func() {

		Convey("When I call decodeData", func() {

			r := &http.Response{
				Body: nil,
			}

			dest := map[string]interface{}{}
			err := decodeData(r, &dest)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `Unable to unmarshal data: nil reader`)
			})

			Convey("Then the dest should be empty", func() {
				So(len(dest), ShouldEqual, 0)
			})
		})
	})

	Convey("Given I have a reader that returns an errpr", t, func() {

		r := &http.Response{
			Body: ioutil.NopCloser(&fakeReader{}),
		}

		Convey("When I call decodeData", func() {

			dest := map[string]interface{}{}
			err := decodeData(r, &dest)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `Unable to unmarshal data: unable to read data: boom`)
			})

			Convey("Then the dest should be empty", func() {
				So(len(dest), ShouldEqual, 0)
			})
		})
	})
}
