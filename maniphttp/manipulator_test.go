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
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/idempotency"
	"go.aporeto.io/manipulate/internal/tracing"
	"go.aporeto.io/manipulate/maniptest"
)

func TestHTTP_NewSHTTPm(t *testing.T) {

	Convey("When I create a new HTTP manipulator", t, func() {

		m := NewHTTPManipulator("http://url.com", "username", "password", "myns").(*httpManipulator)

		Convey("Then the property Username should be 'username'", func() {
			So(m.username, ShouldEqual, "username")
		})

		Convey("Then the property Password should be 'password'", func() {
			So(m.password, ShouldEqual, "password")
		})

		Convey("Then the property URL should be 'http://url.com'", func() {
			So(m.url, ShouldEqual, "http://url.com")
		})

		Convey("Then the property namespace should be 'myns'", func() {
			So(m.namespace, ShouldEqual, "myns")
		})

		Convey("Then the it should implement Manpilater interface", func() {

			var i interface{} = m
			var ok bool
			_, ok = i.(manipulate.Manipulator)
			So(ok, ShouldBeTrue)
		})
	})
}

func TestHTTP_makeAuthorizationHeaders(t *testing.T) {

	Convey("Given I create a new HTTP manipulator", t, func() {

		Convey("When I prepare the Authorization", func() {

			m := NewHTTPManipulator("http://url.com", "username", "password", "").(*httpManipulator)
			h := m.makeAuthorizationHeaders()

			Convey("Then the header should be correct", func() {
				So(h, ShouldEqual, "username password")
			})
		})
	})
}

func TestHTTP_prepareHeaders(t *testing.T) {

	Convey("Given I create an authenticated session", t, func() {

		m := NewHTTPManipulator("http://fake.com", "username", "password", "myns").(*httpManipulator)

		m.globalHeaders = http.Header{
			"Header-1": []string{"hey"},
			"Header-2": []string{"ho"},
		}

		Convey("Given I create a Request", func() {

			req, _ := http.NewRequest("GET", "http://fake.com", nil)

			Convey("When I prepareHeaders with a no context", func() {

				m.prepareHeaders(req, manipulate.NewContext(context.Background()))

				Convey("Then I should have a the X-Namespace set to 'myns'", func() {
					So(req.Header.Get("X-Namespace"), ShouldEqual, "myns")
				})

				Convey("Then I should not have a value for X-Count-Total", func() {
					So(req.Header.Get("X-Count-Total"), ShouldEqual, "")
				})

				Convey("Then I should get the global headers", func() {
					So(req.Header.Get("Header-1"), ShouldEqual, "hey")
					So(req.Header.Get("Header-2"), ShouldEqual, "ho")
				})
			})

			Convey("When I prepareHeaders with various options", func() {

				ctx := manipulate.NewContext(
					context.Background(),
					manipulate.ContextOptionTracking("tid", "type"),
					manipulate.ContextOptionReadConsistency(manipulate.ReadConsistencyStrong),
					manipulate.ContextOptionWriteConsistency(manipulate.WriteConsistencyStrong),
					manipulate.ContextOptionFields([]string{"a", "b"}),
				)

				ctx.(idempotency.Keyer).SetIdempotencyKey("coucou")

				m.prepareHeaders(req, ctx)

				Convey("Then I should have a value for X-External-Tracking-ID", func() {
					So(req.Header.Get("X-External-Tracking-ID"), ShouldEqual, "tid")
				})

				Convey("Then I should have a value for X-External-Tracking-Type", func() {
					So(req.Header.Get("X-External-Tracking-Type"), ShouldEqual, "type")
				})

				Convey("Then I should have a value for X-Read-Consistency", func() {
					So(req.Header.Get("X-Read-Consistency"), ShouldEqual, "strong")
				})

				Convey("Then I should have a value for X-Write-Consistency", func() {
					So(req.Header.Get("X-Write-Consistency"), ShouldEqual, "strong")
				})

				Convey("Then I should have a value for Idempotency-Key", func() {
					So(req.Header.Get("Idempotency-Key"), ShouldEqual, "coucou")
				})

				Convey("Then I should have a value for X-Fields", func() {
					So(req.Header["X-Fields"], ShouldResemble, []string{"a", "b"})
				})
			})
		})
	})
}

func TestHTTP_readHeaders(t *testing.T) {

	Convey("Given I create a new HTTP manipulator and a Context", t, func() {

		m := NewHTTPManipulator("http://fake.com", "username", "password", "").(*httpManipulator)
		ctx := manipulate.NewContext(context.Background())
		req := &http.Response{Header: http.Header{}}

		Convey("When I readHeaders with a no context", func() {

			m.readHeaders(req, nil)

			Convey("Then nothing should happen", func() {
			})
		})

		Convey("When I readHeaders with a request that has count information", func() {

			req.Header.Set("X-Count-Total", "456")

			m.readHeaders(req, ctx)

			Convey("Then Context.Count() should be 456", func() {
				So(ctx.Count(), ShouldEqual, 456)
			})
		})

		Convey("When I readHeaders with a request that has messages", func() {

			req.Header["X-Messages"] = []string{"hello", "bonjour"}

			m.readHeaders(req, ctx)

			Convey("Then Context.Messages() should be 456", func() {
				So(ctx.Messages(), ShouldResemble, []string{"hello", "bonjour"})
			})
		})
	})
}

func TestHTTP_standardURI(t *testing.T) {

	Convey("Given I create a new manipulator and an object", t, func() {

		list := testmodel.NewList()

		m := NewHTTPManipulator("http://url.com", "username", "password", "").(*httpManipulator)

		Convey("When I check personal URI of a standard object with an ID", func() {

			list.SetIdentifier("xxx")

			Convey("When I use no version", func() {

				url, err := m.getPersonalURL(list, 0)

				Convey("Then it should be http://url.com/v/1/lists/xxx", func() {
					So(url, ShouldEqual, "http://url.com/v/1/lists/xxx")
				})

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("When I use a version", func() {

				url, err := m.getPersonalURL(list, 12)

				Convey("Then it should be http://url.com/v/12/lists/xxx", func() {
					So(url, ShouldEqual, "http://url.com/v/12/lists/xxx")
				})

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("When I check general URI of a standard object with an ID", func() {

			list.SetIdentifier("xxx")

			Convey("When I use no version", func() {

				url := m.getGeneralURL(list, 0)

				Convey("Then it should be http://url.com/v/1/lists", func() {
					So(url, ShouldEqual, "http://url.com/v/1/lists")
				})
			})

			Convey("When I use a version", func() {

				url := m.getGeneralURL(list, 12)

				Convey("Then it should be http://url.com/v/12/lists", func() {
					So(url, ShouldEqual, "http://url.com/v/12/lists")
				})
			})
		})

		Convey("When I check children URL for a root object", func() {

			list.SetIdentifier("xxx")

			Convey("When I use no version", func() {

				url, err := m.getURLForChildrenIdentity(nil, testmodel.ListIdentity, 0, 0)

				Convey("Then URL of the children with ListIdentity should be http://url.com/lists", func() {
					So(url, ShouldEqual, "http://url.com/lists")
				})

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("When I use a version", func() {

				url, err := m.getURLForChildrenIdentity(nil, testmodel.ListIdentity, 0, 12)

				Convey("Then URL of the children with ListIdentity should be http://url.com/v/12/lists", func() {
					So(url, ShouldEqual, "http://url.com/v/12/lists")
				})

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("When I check children URL for a standard object with an ID", func() {

			list.SetIdentifier("xxx")

			Convey("When I use no version", func() {

				url, err := m.getURLForChildrenIdentity(list, testmodel.TaskIdentity, 0, 0)

				Convey("Then URL of the children with FakeRootIdentity should be http://url.com/v/1/lists/xxx/tasks", func() {
					So(url, ShouldEqual, "http://url.com/v/1/lists/xxx/tasks")
				})

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("When I use a version", func() {

				url, err := m.getURLForChildrenIdentity(list, testmodel.TaskIdentity, 0, 12)

				Convey("Then URL of the children with FakeRootIdentity should be http://url.com/v/12/lists/xxx/tasks", func() {
					So(url, ShouldEqual, "http://url.com/v/12/lists/xxx/tasks")
				})

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})

		})

		Convey("When I check the general URL of a standard object without an ID", func() {

			url, err := m.getURLForChildrenIdentity(list, testmodel.TaskIdentity, 0, 0)

			Convey("Then it should be ''", func() {
				So(url, ShouldEqual, "")
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I check the children URL for a standard object without an ID", func() {

			url, err := m.getURLForChildrenIdentity(list, testmodel.TaskIdentity, 0, 0)

			Convey("Then it should be ''", func() {
				So(url, ShouldEqual, "")
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_Retrieve(t *testing.T) {

	Convey("Given I have an http manipulator and a and working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"ID":"xxx","name":"the list 1"}`)
		}))
		defer ts.Close()

		m := NewHTTPManipulator(ts.URL, "username", "password", "")

		Convey("When I fetch an entity", func() {

			list := testmodel.NewList()
			list.ID = "xxx"
			errs := m.Retrieve(nil, list)

			Convey("Then err should be nil", func() {
				So(errs, ShouldBeNil)
			})

			Convey("Then Name should pedro", func() {
				So(list.Name, ShouldEqual, "the list 1")
			})
		})

		Convey("When I fetch an entity with no ID", func() {

			list := testmodel.NewList()
			err := m.Retrieve(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotBuildQuery{})
			})
		})
	})

	Convey("Given I have an http manipulator and a and the server will return an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(422)
			fmt.Fprint(w, `[{"code":422,"description":"nope.","subject":"elemental","title":"Read Only Error","data":null}]`)
		}))
		defer ts.Close()

		m := NewHTTPManipulator(ts.URL, "username", "password", "")

		Convey("When I fetch an entity", func() {

			list := testmodel.NewList()
			list.ID = "x"
			err := m.Retrieve(nil, list)

			Convey("Then error should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.(elemental.Errors).Code(), ShouldEqual, 422)
				So(err.(elemental.Errors)[0].Description, ShouldEqual, "nope.")
			})
		})
	})

	Convey("Given I have an http manipulator and a and the server will return bad json", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `not good at all`)
		}))
		defer ts.Close()

		m := NewHTTPManipulator(ts.URL, "username", "password", "")

		Convey("When I fetch an entity", func() {

			list := testmodel.NewList()
			list.ID = "x"
			err := m.Retrieve(nil, list)

			Convey("Then error should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotUnmarshal{})
			})
		})
	})
}

func TestHTTP_Update(t *testing.T) {

	Convey("Given I have an http manipulator and a and working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"ID": "zzz", "parentType": "pedro", "parentID": "yyy"}`)
		}))
		defer ts.Close()

		m := NewHTTPManipulator(ts.URL, "username", "password", "")

		Convey("When I save an entity", func() {

			list := testmodel.NewList()
			list.ID = "yyy"
			errs := m.Update(nil, list)

			Convey("Then err should be nil", func() {
				So(errs, ShouldBeNil)
			})

			Convey("Then ID should 'zzz'", func() {
				So(list.Identifier(), ShouldEqual, "zzz")
			})
		})

		Convey("When I save an entity with no ID", func() {

			list := testmodel.NewList()
			errs := m.Update(nil, list)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I save an unmarshalable entity", func() {

			list := testmodel.NewUnmarshalableList()
			list.ID = "yyy"
			errs := m.Update(nil, list)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have an http manipulator and the server returns an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "nope")
		}))
		defer ts.Close()

		m := NewHTTPManipulator(ts.URL, "username", "password", "")

		Convey("When I save an entity", func() {

			list := testmodel.NewList()
			list.ID = "ddd"
			err := m.Update(nil, list)

			Convey("Then error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have an http manipulator and the server returns a bad json", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `bad json`)
		}))
		defer ts.Close()

		m := NewHTTPManipulator(ts.URL, "username", "password", "")

		Convey("When I save an entity", func() {

			list := testmodel.NewList()
			list.ID = "x"
			err := m.Update(nil, list)

			Convey("Then the error should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotUnmarshal{})
			})
		})
	})

	Convey("Given I have an http manipulator and the server returns no data", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
		}))
		defer ts.Close()

		m := NewHTTPManipulator(ts.URL, "username", "password", "")

		Convey("When I save an entity", func() {

			list := testmodel.NewList()

			Convey("Then it not should panic", func() {
				So(func() { _ = m.Update(nil, list) }, ShouldNotPanic)
			})
		})
	})
}

func TestHTTP_Delete(t *testing.T) {

	Convey("Given I have an http manipulator and a and working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"ID":"yyy"}`)) // nolint
		}))
		defer ts.Close()

		m := NewHTTPManipulator(ts.URL, "username", "password", "")

		Convey("When I delete an entity", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			err := m.Delete(nil, list)
			if err != nil {
				panic(err)
			}

			Convey("Then ID should 'yyy'", func() {
				So(list.Identifier(), ShouldEqual, "yyy")
			})
		})

		Convey("When I delete an entity with no ID", func() {

			m := NewHTTPManipulator("username", "password", "http://fake.com", "")

			list := testmodel.NewList()
			err := m.Delete(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have an http manipulator and the server returns an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "nope")
		}))
		defer ts.Close()

		m := NewHTTPManipulator(ts.URL, "username", "password", "")

		Convey("When I delete an entity", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			err := m.Delete(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_RetrieveMany(t *testing.T) {

	Convey("Given I have an existing object", t, func() {

		list := testmodel.NewList()
		list.ID = "xxx"

		Convey("When I fetch its children with success", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `[{"ID": "1", "name": "name1"}, {"ID": "2", "name": "name2"}]`)
			}))
			defer ts.Close()

			m := NewHTTPManipulator(ts.URL, "username", "password", "")

			var l testmodel.TasksList
			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list),
			)

			errs := m.RetrieveMany(ctx, &l)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldBeNil)
			})

			Convey("Then the length of the children list should be 2", func() {
				So(len(l), ShouldEqual, 2)
			})

			Convey("Then the first child ID should be 1 and Name name1", func() {
				So(l[0].Identifier(), ShouldEqual, "1")
				So(l[0].Name, ShouldEqual, "name1")
			})

			Convey("Then the second child ID should be 2 Name name1", func() {
				So(l[1].Identifier(), ShouldEqual, "2")
				So(l[1].Name, ShouldEqual, "name2")
			})

			Convey("Then the identity of the children should be FakeIdentity", func() {
				So(l[0].Identity(), ShouldResemble, testmodel.TaskIdentity)
				So(l[1].Identity(), ShouldResemble, testmodel.TaskIdentity)
			})
		})

		Convey("When I fetch its children but the parent has no ID", func() {

			m := NewHTTPManipulator("username", "password", "http://fake.com", "")

			list2 := testmodel.NewList()
			var l testmodel.TasksList

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list2),
			)

			errs := m.RetrieveMany(ctx, &l)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I fetch its children while there is no data", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
			}))
			defer ts.Close()

			m := NewHTTPManipulator(ts.URL, "username", "password", "")

			e := testmodel.NewTask()
			var l testmodel.TasksList

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(e),
			)

			errs := m.RetrieveMany(ctx, &l)

			Convey("Then the length of the children list should be 0", func() {
				So(l, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I fetch its children while there is none, but I still get an empty array", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `[]`)
			}))
			defer ts.Close()
			m := NewHTTPManipulator(ts.URL, "username", "password", "")

			var l testmodel.TasksList

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list),
			)

			_ = m.RetrieveMany(ctx, &l)

			Convey("Then the length of the children list should be 0", func() {
				So(len(l), ShouldEqual, 0)
			})
		})

		Convey("When I fetch the children and I got a communication error", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "woops")
			}))
			defer ts.Close()

			m := NewHTTPManipulator(ts.URL, "username", "password", "")

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list),
			)

			var l testmodel.TasksList
			errs := m.RetrieveMany(ctx, &l)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I fetch the children I got a bad json", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `[not good]`)
			}))
			defer ts.Close()

			m := NewHTTPManipulator(ts.URL, "username", "password", "")

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list),
			)

			var l testmodel.TasksList
			errs := m.RetrieveMany(ctx, &l)

			Convey("Then the error should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_Create(t *testing.T) {

	Convey("Given I have an existing object", t, func() {

		list := testmodel.NewList()
		list.ID = "xxx"

		Convey("When I create a child with success", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				fmt.Fprint(w, `{"ID": "zzz"}`)
			}))
			defer ts.Close()

			m := NewHTTPManipulator(ts.URL, "username", "password", "")
			errs := m.Create(nil, list)

			Convey("Then the error should not be nil", func() {
				So(errs, ShouldBeNil)
			})

			Convey("Then ID of the new children should be zzz", func() {
				So(list.Identifier(), ShouldEqual, "zzz")
			})
		})

		Convey("When I create a child for a parent that has no ID", func() {

			m := NewHTTPManipulator("username", "password", "url.com", "")
			list2 := testmodel.NewList()
			task := testmodel.NewTask()
			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list2),
			)

			errs := m.Create(ctx, task)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I create a child that is nil", func() {

			m := NewHTTPManipulator("username", "password", "http://fake.com", "")
			errs := m.Create(nil, list)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I create a child and I got a communication error", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "")
			}))
			defer ts.Close()

			m := NewHTTPManipulator(ts.URL, "username", "password", "")
			errs := m.Create(nil, list)

			Convey("Then error should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I create a child I got a bad json", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				fmt.Fprint(w, `[{"bad"}]`)
			}))
			defer ts.Close()

			m := NewHTTPManipulator(ts.URL, "username", "password", "")
			errs := m.Create(nil, list)

			Convey("Then the error should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I create an unmarshalable entity", func() {

			m := NewHTTPManipulator("username", "password", "", "")
			list := testmodel.NewUnmarshalableList()
			list.ID = "yyy"
			err := m.Create(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotMarshal{})
			})
		})
	})
}

func TestHTTP_Count(t *testing.T) {

	Convey("Given I have an existing object", t, func() {

		list := testmodel.NewList()
		list.ID = "xxx"

		Convey("When I fetch the count of its children with success", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("X-Count-Total", "10")
			}))
			defer ts.Close()

			m := NewHTTPManipulator(ts.URL, "username", "password", "")

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list),
			)

			num, errs := m.Count(ctx, testmodel.TaskIdentity)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldBeNil)
			})

			Convey("Then count should be 10", func() {
				So(num, ShouldEqual, 10)
			})
		})

		Convey("When I fetch its children but the parent has no ID", func() {

			m := NewHTTPManipulator("username", "password", "http://fake.com", "")

			list2 := testmodel.NewList()

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list2),
			)

			_, errs := m.Count(ctx, testmodel.TaskIdentity)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I fetch the children and I got a communication error", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "woops")
			}))
			defer ts.Close()

			m := NewHTTPManipulator(ts.URL, "username", "password", "")

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list),
			)

			_, errs := m.Count(ctx, testmodel.TaskIdentity)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_send(t *testing.T) {

	sp := tracing.StartTrace(nil, "test")
	defer sp.Finish()

	Convey("Given I have a m with bad url", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send", func() {

			_, err := m.(*httpManipulator).send(manipulate.NewContext(context.Background()), http.MethodPost, "nop", nil, sp, 0)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, `Cannot communicate: Post nop: unsupported protocol scheme ""`)
			})
		})
	})

	Convey("Given I have a m with with a timeout", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 0)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, "https://google.com", nil, sp, 0)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: Post https://google.com: context deadline exceeded")
			})
		})
	})

	Convey("Given I have a m with with a 408", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusRequestTimeout)
				fmt.Fprint(w, `[{"code": 408, "title": "nope", "description": "boom"}]`)
			}))
			defer ts.Close()

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, sp, 0)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: error 408 (): nope: boom")
			})
		})
	})

	Convey("Given I have a m with with a 502", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadGateway)
				fmt.Fprint(w, `[{"code": 502, "title": "nope", "description": "boom"}]`)
			}))
			defer ts.Close()

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, sp, 0)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: Bad gateway")
			})
		})
	})

	Convey("Given I have a m with with a 503", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprint(w, `[{"code": 503, "title": "nope", "description": "boom"}]`)
			}))
			defer ts.Close()

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, sp, 0)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: Service unavailable")
			})
		})
	})

	Convey("Given I have a m with with a 504", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusGatewayTimeout)
				fmt.Fprint(w, `[{"code": 504, "title": "nope", "description": "boom"}]`)
			}))
			defer ts.Close()

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, sp, 0)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: Gateway timeout")
			})
		})
	})

	Convey("Given I have a m with with a 429", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				fmt.Fprint(w, `[{"code": 429, "title": "nope", "description": "boom"}]`)
			}))
			defer ts.Close()

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, sp, 0)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrTooManyRequests{})
				So(err.Error(), ShouldEqual, "Too many requests: error 429 (): nope: boom")
			})
		})
	})

	Convey("Given I have a m with with a 403 and no token manager", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprint(w, `[{"code": 403, "title": "nope", "description": "boom"}]`)
			}))
			defer ts.Close()

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, sp, 0)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "error 403 (): nope: boom")
			})
		})
	})

	Convey("Given I have a m with with a 403 and with a token manager", t, func() {

		var call int
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			call++
			if call == 1 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprint(w, `[{"code": 403, "title": "nope", "description": "boom"}]`)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNoContent)
			}
		}))
		defer ts.Close()

		m := NewHTTPManipulator("username", "password", "", "")
		m.(*httpManipulator).tokenManager = maniptest.NewTestTokenManager()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, sp, 0)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a m with with a 423", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusLocked)
				fmt.Fprint(w, `[{"code": 423, "title": "nope", "description": "boom"}]`)
			}))
			defer ts.Close()

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, sp, 0)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrLocked{})
				So(err.Error(), ShouldEqual, "Cannot communicate: The api has been locked down by the server.")
			})
		})
	})

	Convey("Given I have a m with with a unmarshalable error", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, `[{"code": 423, "]`)
			}))
			defer ts.Close()

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, sp, 0)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotUnmarshal{})
				So(err.Error(), ShouldEqual, `Unable to unmarshal data: unable to decode application/json: EOF. original data:
[{"code": 423, "]`)
			})
		})
	})

	Convey("Given I have a m with with a 500", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, `[{"code": 500, "title": "nope", "description": "boom"}]`)
			}))
			defer ts.Close()

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, sp, 0)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, elemental.Errors{})
				So(err.Error(), ShouldEqual, "error 500 (): nope: boom")
			})
		})
	})
}

func TestHTTP_setPassword(t *testing.T) {

	Convey("Given I have an http manipulator", t, func() {

		m := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call setPassword", func() {

			m.(*httpManipulator).setPassword("secret")

			Convey("Then it should set the password", func() {
				So(m.(*httpManipulator).currentPassword(), ShouldEqual, "secret")
			})
		})
	})
}

func TestHTTP_renewNotifiers(t *testing.T) {

	Convey("Given I have an http manipulator", t, func() {

		m := NewHTTPManipulator("username", "password", "", "").(*httpManipulator)

		var called1, called2 string
		notifier1 := func(p string) { called1 = p }
		notifier2 := func(p string) { called2 = p }

		Convey("When I register the notifiers", func() {

			m.registerRenewNotifier("1", notifier1)
			m.registerRenewNotifier("2", notifier2)

			Convey("When I call setPassword", func() {

				m.setPassword("changed")

				Convey("Then both notified should have been called", func() {
					So(called1, ShouldEqual, "changed")
					So(called2, ShouldEqual, "changed")
				})

				Convey("Then when I unregister notifier2", func() {

					m.unregisterRenewNotifier("2")

					Convey("When I call setPassword again", func() {

						m.setPassword("changed1")

						Convey("Then both notified should have been called", func() {
							So(called1, ShouldEqual, "changed1")
							So(called2, ShouldEqual, "changed")
						})
					})
				})
			})
		})
	})
}
