// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maniphttp

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
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
					elemental.NewFilterComposer().WithKey("name").Equals("toto").WithKey("description").Equals("hello").Done(),
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

		Convey("When I call the method addQueryParameters with the propagate option", func() {

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionPropagated(true),
			)

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the query string should be correct ", func() {
				So(request.URL.RawQuery, ShouldEqual, "propagated=true")
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

		Convey("When I call the method addQueryParameters with lazy pagination", func() {

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionAfter("12", 42),
			)

			err := addQueryParameters(request, ctx)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the query string should be correct ", func() {
				So(request.URL.RawQuery, ShouldEqual, "after=12&limit=42")
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
			Body: io.NopCloser(bytes.NewBuffer([]byte(`{"name":"thename","age": 2}`))),
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
				So(dest["age"].(uint64), ShouldEqual, 2)
			})
		})
	})

	Convey("Given I have invalid valid json data in a reader", t, func() {

		r := &http.Response{
			Body: io.NopCloser(bytes.NewBuffer([]byte(`<html><body>not json</body></html>`))),
		}

		Convey("When I call decodeData", func() {

			dest := map[string]interface{}{}
			err := decodeData(r, &dest)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Unable to unmarshal data: unable to decode application/json: json decode error [pos 1]: read map - expect char '{' but got char '<'. original data:\n<html><body>not json</body></html>")
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
			Body: io.NopCloser(&fakeReader{}),
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
