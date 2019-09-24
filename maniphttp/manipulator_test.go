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
	"sync/atomic"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/idempotency"
	"go.aporeto.io/manipulate/internal/tracing"
	"go.aporeto.io/manipulate/maniptest"
	"golang.org/x/sync/errgroup"
)

func TestHTTP_New(t *testing.T) {

	Convey("When I create a simple manipulator", t, func() {

		mm, _ := New(
			context.Background(),
			"http://url.com",
			OptionCredentials("username", "password"),
			OptionNamespace("myns"),
		)
		m := mm.(*httpManipulator)

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

		Convey("Then the it should implement Manipulator interface", func() {

			var i interface{} = m
			var ok bool
			_, ok = i.(manipulate.Manipulator)
			So(ok, ShouldBeTrue)
		})
	})

	Convey("When I create a manipulator with empty url", t, func() {

		Convey("Then it should panic", func() {
			So(func() { New(context.Background(), "") }, ShouldPanicWith, "empty url") // nolint
		})
	})

	Convey("When I create a manipulator with a token manager that works", t, func() {

		tm := maniptest.NewTestTokenManager()
		tm.MockIssue(t, func(context.Context) (string, error) { return "old-token", nil })

		mm, err := New(context.Background(), "http://url.com", OptionTokenManager(tm))
		m := mm.(*httpManipulator)

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Then username should be correct", func() {
			So(m.username, ShouldEqual, "Bearer")
		})

		Convey("Then password should be old-token", func() {
			So(m.currentPassword(), ShouldEqual, "old-token")
		})

		tm.MockRun(t, func(ctx context.Context, tokenCh chan string) {
			tokenCh <- "new-token"
		})

		time.Sleep(2 * time.Second) // concourse is sometimes slow...

		Convey("Then password should be new-token", func() {
			So(m.currentPassword(), ShouldEqual, "new-token")
		})
	})

	Convey("When I create a new HTTP manipulator with a token manager that fails", t, func() {

		tm := maniptest.NewTestTokenManager()
		tm.MockIssue(t, func(context.Context) (string, error) { return "", fmt.Errorf("paf") })

		_, err := New(context.Background(), "http://url.com", OptionTokenManager(tm))

		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "paf")
		})
	})
}

func TestHTTP_RetrieveMany(t *testing.T) {

	Convey("Given I have a manipulator and a working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `[{"ID": "1", "name": "name1"}, {"ID": "2", "name": "name2"}]`)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I retrieve the objects", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			mctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list),
			)

			var l testmodel.TasksList
			err := m.RetrieveMany(mctx, &l)

			Convey("Then err should not be nil", func() {
				So(err, ShouldBeNil)
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

		Convey("When I retrieve objects from a parent that has no ID", func() {

			list2 := testmodel.NewList()
			var l testmodel.TasksList

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list2),
			)

			err := m.RetrieveMany(ctx, &l)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I retrieve nil objects", func() {

			err := m.RetrieveMany(nil, nil)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and the server returns no data", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I retrieve the objects", func() {

			e := testmodel.NewTask()
			var l testmodel.TasksList

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(e),
			)

			err := m.RetrieveMany(ctx, &l)

			Convey("Then the length of the children list should be 0", func() {
				So(l, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and the server returns an empty array", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `[]`)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I retrieve the objects", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

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
	})

	Convey("Given I have a manipulator and the server returns an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "woops")
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		var l testmodel.TasksList
		err := m.RetrieveMany(nil, &l)

		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Given I have a manipulator and the server returns a bad json", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `[not good]`)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I retrieve the objects", func() {
			var l testmodel.TasksList
			err := m.RetrieveMany(nil, &l)

			Convey("Then the error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_Retrieve(t *testing.T) {

	Convey("Given I have a manipulator and a working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"ID":"xxx","name":"the list 1"}`)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I retrieve an object", func() {

			list := testmodel.NewList()
			list.ID = "xxx"
			err := m.Retrieve(nil, list)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then retrieve object should be correct", func() {
				So(list.Name, ShouldEqual, "the list 1")
			})
		})

		Convey("When I retrieve an object with no ID", func() {

			list := testmodel.NewList()
			err := m.Retrieve(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotBuildQuery{})
			})
		})

		Convey("When I retrieve a nil object", func() {

			err := m.Retrieve(nil, nil)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and the server returns no data", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I retrieve the objects", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			err := m.Retrieve(nil, list)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and a and the server will return an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(422)
			fmt.Fprint(w, `[{"code":422,"description":"nope.","subject":"elemental","title":"Read Only Error","data":null}]`)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I retrieve an object", func() {

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

	Convey("Given I have a manipulator and the server returns a bad json", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `not good at all`)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I retrieve an object", func() {

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

func TestHTTP_Create(t *testing.T) {

	Convey("Given I have a manipulator and a working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"ID": "zzz"}`)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I create the objects", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			err := m.Create(nil, list)

			Convey("Then the error should not be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then ID of the new children should be zzz", func() {
				So(list.Identifier(), ShouldEqual, "zzz")
			})
		})

		Convey("When I create an unmarshalable entity", func() {

			list := testmodel.NewUnmarshalableList()
			err := m.Create(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotMarshal{})
			})
		})

		Convey("When I create a nil object", func() {

			err := m.Create(nil, nil)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I create a child for a parent that has no ID", func() {

			list2 := testmodel.NewList()
			task := testmodel.NewTask()
			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list2),
			)

			err := m.Create(ctx, task)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I create a child that is nil", func() {

			err := m.Create(nil, nil)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

	})

	Convey("Given I have a manipulator and the server returns no data", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I create the object", func() {

			list := testmodel.NewList()

			err := m.Create(nil, list)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and a server that returns an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "")
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I call create", func() {

			list := testmodel.NewList()

			err := m.Create(nil, list)

			Convey("Then error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and a server that returns a bad json", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `[{"bad"}]`)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I call create", func() {

			list := testmodel.NewList()

			err := m.Create(nil, list)

			Convey("Then the error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_Update(t *testing.T) {

	Convey("Given I have a manipulator and a working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"ID": "zzz", "parentType": "pedro", "parentID": "yyy"}`)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I update an object", func() {

			list := testmodel.NewList()
			list.ID = "yyy"
			err := m.Update(nil, list)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the object should be correct", func() {
				So(list.Identifier(), ShouldEqual, "zzz")
			})
		})

		Convey("When I update an object with no ID", func() {

			list := testmodel.NewList()
			err := m.Update(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotBuildQuery{})
			})
		})

		Convey("When I update a nil object", func() {

			err := m.Update(nil, nil)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I update a sparse object", func() {

			list := testmodel.NewSparseList()
			id := "yyy"
			list.ID = &id

			err := m.Update(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I update an unmarshalable object", func() {

			list := testmodel.NewUnmarshalableList()
			list.ID = "yyy"
			err := m.Update(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and the server returns no data", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I update the object", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			err := m.Update(nil, list)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and the server returns an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "nope")
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I save an entity", func() {

			list := testmodel.NewList()
			list.ID = "ddd"
			err := m.Update(nil, list)

			Convey("Then error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and the server returns a bad json", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `bad json`)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I update an object", func() {

			list := testmodel.NewList()
			list.ID = "x"
			err := m.Update(nil, list)

			Convey("Then the error should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotUnmarshal{})
			})
		})
	})
}

func TestHTTP_Delete(t *testing.T) {

	Convey("Given I have a manipulator and a working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"ID":"yyy"}`)) // nolint
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I delete an object", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			err := m.Delete(nil, list)
			if err != nil {
				panic(err)
			}

			Convey("Then object should be correct", func() {
				So(list.Identifier(), ShouldEqual, "yyy")
			})
		})

		Convey("When I delete an entity with no ID", func() {

			mm, _ := New(context.Background(), "http://fake.com")
			m := mm.(*httpManipulator)

			list := testmodel.NewList()
			err := m.Delete(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I delete a nil object", func() {

			err := m.Delete(nil, nil)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotBuildQuery{})
			})
		})
	})

	Convey("Given I have a manipulator and the server returns no data", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I update the object", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			err := m.Delete(nil, list)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and the server returns an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "nope")
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I delete an object", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			err := m.Delete(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_DeleteMany(t *testing.T) {

	Convey("Given I have a manipulator", t, func() {

		mm, _ := New(context.Background(), "https://fake.com")
		m := mm.(*httpManipulator)

		err := m.DeleteMany(nil, testmodel.TaskIdentity)

		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})
}

func TestHTTP_Count(t *testing.T) {

	Convey("Given I have a manipulator and a working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Count-Total", "10")
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I call count", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list),
			)

			num, err := m.Count(ctx, testmodel.TaskIdentity)

			Convey("Then err should not be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then count should be 10", func() {
				So(num, ShouldEqual, 10)
			})
		})

		Convey("When I count children but the parent has no ID", func() {

			list := testmodel.NewList()

			ctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionParent(list),
			)

			_, err := m.Count(ctx, testmodel.TaskIdentity)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and a server that returns an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "woops")
		}))
		defer ts.Close()

		mm, _ := New(context.Background(), ts.URL)
		m := mm.(*httpManipulator)

		Convey("When I call count", func() {
			_, err := m.Count(nil, testmodel.TaskIdentity)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_send(t *testing.T) {

	sp := tracing.StartTrace(nil, "test")
	defer sp.Finish()

	Convey("Given I have a m with bad url", t, func() {

		m, _ := New(context.Background(), "toto.com")

		Convey("When I call send", func() {

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(context.Background()), http.MethodPost, "nop", nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotExecuteQuery{})
				So(err.Error(), ShouldEqual, `Unable to execute query: Post nop: unsupported protocol scheme ""`)

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have an already canceled context", t, func() {

		m, _ := New(context.Background(), "toto.com")

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 0)
			cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, "https://google.com", nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: Post https://google.com: context deadline exceeded")

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning net.Error", t, func() {

		m, _ := New(context.Background(), "toto.com")

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, "https://NANANAN", nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})

				// On linux, the message is slightly different than on macOS
				So(err.Error(), ShouldStartWith, "Cannot communicate: Post https://NANANAN: dial tcp: lookup NANANAN")

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning EOF error", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("eof")
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, fmt.Sprintf("Cannot communicate: Post %s: EOF", ts.URL))

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning TLS error", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrTLS{})
				So(err.Error(), ShouldEqual, fmt.Sprintf("TLS error: Post %s: x509: certificate signed by unknown authority", ts.URL))

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server and a retry func", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusRequestTimeout)
			fmt.Fprint(w, `[{"code": 408, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			var t int
			var rerr error
			resp, err := m.(*httpManipulator).send(
				manipulate.NewContext(
					ctx,
					manipulate.ContextOptionRetryFunc(func(i manipulate.RetryInfo) error {
						t = i.Try()
						rerr = i.Err()
						return nil
					}),
				),
				http.MethodPost,
				ts.URL,
				nil,
				nil,
				sp,
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: Request Timeout")
				So(t, ShouldBeGreaterThan, 1)
				So(rerr.Error(), ShouldEqual, err.Error())

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server and a retry func that returns a error at try 3", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusRequestTimeout)
			fmt.Fprint(w, `[{"code": 408, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			var t int
			resp, err := m.(*httpManipulator).send(
				manipulate.NewContext(
					ctx,
					manipulate.ContextOptionRetryFunc(func(i manipulate.RetryInfo) error {
						t = i.Try()
						if t == 3 {
							return fmt.Errorf("bam")
						}
						return nil
					}),
				),
				http.MethodPost,
				ts.URL,
				nil,
				nil,
				sp,
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "bam")
				So(t, ShouldEqual, 3)

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server and a default retry func", t, func() {

		var t int
		var rerr error

		m, _ := New(
			context.Background(),
			"toto.com",
			OptionDefaultRetryFunc(func(i manipulate.RetryInfo) error {
				t = i.Try()
				rerr = i.Err()
				return nil
			}))

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusRequestTimeout)
			fmt.Fprint(w, `[{"code": 408, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(
				manipulate.NewContext(ctx),
				http.MethodPost,
				ts.URL,
				nil,
				nil,
				sp,
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: Request Timeout")
				So(t, ShouldBeGreaterThan, 1)
				So(rerr.Error(), ShouldEqual, err.Error())

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server and a default retry func that returns a error at try 3", t, func() {

		var t int

		m, _ := New(
			context.Background(),
			"toto.com",
			OptionDefaultRetryFunc(func(i manipulate.RetryInfo) error {
				t = i.Try()
				if t == 3 {
					return fmt.Errorf("bam")
				}
				return nil
			}))

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusRequestTimeout)
			fmt.Fprint(w, `[{"code": 408, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(
				manipulate.NewContext(ctx),
				http.MethodPost,
				ts.URL,
				nil,
				nil,
				sp,
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "bam")
				So(t, ShouldEqual, 3)

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server never returning", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(3 * time.Second)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(
				manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, fmt.Sprintf("Cannot communicate: Post %s: context deadline exceeded", ts.URL))

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning 408", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusRequestTimeout)
			fmt.Fprint(w, `[{"code": 408, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: Request Timeout")

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning 502", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprint(w, `[{"code": 502, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: Bad gateway")

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning 503", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, `[{"code": 503, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: Service unavailable")

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning 504", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusGatewayTimeout)
			fmt.Fprint(w, `[{"code": 504, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
				So(err.Error(), ShouldEqual, "Cannot communicate: Gateway timeout")

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning 429", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, `[{"code": 429, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrTooManyRequests{})
				So(err.Error(), ShouldEqual, "Too many requests: Too Many Requests")

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning 403 and no token manager", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, `[{"code": 403, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "error 403 (): nope: boom")

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning 403 and with a token manager that renews the token", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Authorization") != "Bearer ok-token" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprint(w, `[{"code": 403, "title": "nope", "description": "boom"}]`)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNoContent)
			}
		}))
		defer ts.Close()

		tm := maniptest.NewTestTokenManager()
		var tmCalled int64
		tm.MockIssue(t, func(context.Context) (string, error) {
			atomic.AddInt64(&tmCalled, 1)
			time.Sleep(300 * time.Millisecond)
			return "ok-token", nil
		})

		m, _ := New(context.Background(), "toto.com")
		m.(*httpManipulator).tokenManager = tm
		m.(*httpManipulator).username = "Bearer"
		m.(*httpManipulator).password = "token"
		m.(*httpManipulator).atomicRenewTokenFunc = elemental.AtomicJob(m.(*httpManipulator).renewToken)

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			var eg errgroup.Group

			eg.Go(func() error {
				_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)
				return err
			})
			eg.Go(func() error {
				_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)
				return err
			})
			eg.Go(func() error {
				_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)
				return err
			})
			eg.Go(func() error {
				_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)
				return err
			})

			err := eg.Wait()

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the token manager should have been called only once", func() {
				So(tmCalled, ShouldEqual, 1)
			})
		})
	})

	Convey("Given I have a server returning 403 with a token manager but a timeout during retry", t, func() {

		var call int
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			call++
			if call == 1 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprint(w, `[{"code": 403, "title": "nope", "description": "boom"}]`)
			} else {
				time.Sleep(5 * time.Second)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNoContent)
			}
		}))
		defer ts.Close()

		m, _ := New(context.Background(), "toto.com")
		m.(*httpManipulator).tokenManager = maniptest.NewTestTokenManager()
		m.(*httpManipulator).atomicRenewTokenFunc = elemental.AtomicJob(m.(*httpManipulator).renewToken)

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "error 403 (): nope: boom")
				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning 403 for all my requests, I should return an error and tokenmanager should only be invoked once", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, `[{"code": 403, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		tm := maniptest.NewTestTokenManager()
		var tmCalled int64
		tm.MockIssue(t, func(context.Context) (string, error) {
			atomic.AddInt64(&tmCalled, 1)
			time.Sleep(300 * time.Millisecond)
			return "ok-token", nil
		})

		m, _ := New(context.Background(), "toto.com")
		m.(*httpManipulator).tokenManager = tm
		m.(*httpManipulator).username = "Bearer"
		m.(*httpManipulator).password = "ok-token"
		m.(*httpManipulator).atomicRenewTokenFunc = elemental.AtomicJob(m.(*httpManipulator).renewToken)

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			_, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should be not nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then the token manager should have been called only once", func() {
				So(tmCalled, ShouldEqual, 1)
			})
		})
	})

	Convey("Given I have a server returning 423", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusLocked)
			fmt.Fprint(w, `[{"code": 423, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrLocked{})
				So(err.Error(), ShouldEqual, "Cannot communicate: The api has been locked down by the server.")

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning an unmarshalable body", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `[{"code": 423, "]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotUnmarshal{})
				So(err.Error(), ShouldEqual, "Unable to unmarshal data: unable to decode application/json: EOF. original data:\n[{\"code\": 423, \"]")

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning 500", t, func() {

		m, _ := New(context.Background(), "toto.com")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, `[{"code": 500, "title": "nope", "description": "boom"}]`)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, elemental.Errors{})
				So(err.Error(), ShouldEqual, "error 500 (): nope: boom")

				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning 201 but a simulation error with 100% chance", t, func() {

		m, _ := New(context.Background(), "toto.com", OptionSimulateFailures(
			map[float64]error{
				1.0: fmt.Errorf("simulated error"),
			},
		))

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("this should not be called")
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "simulated error")
				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a server returning 201 but a simulation error with 0% chance", t, func() {

		m, _ := New(context.Background(), "toto.com", OptionSimulateFailures(
			map[float64]error{
				0.0: fmt.Errorf("simulated error"),
			},
		))

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		}))
		defer ts.Close()

		Convey("When I call send", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			resp, err := m.(*httpManipulator).send(manipulate.NewContext(ctx), http.MethodPost, ts.URL, nil, nil, sp)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_makeAuthorizationHeaders(t *testing.T) {

	Convey("Given I create a new HTTP manipulator", t, func() {

		Convey("When I prepare the Authorization", func() {

			mm, _ := New(
				context.Background(),
				"http://url.com",
				OptionCredentials("username", "password"),
				OptionNamespace("myns"),
			)
			m := mm.(*httpManipulator)

			h := m.makeAuthorizationHeaders("username", "password")

			Convey("Then the header should be correct", func() {
				So(h, ShouldEqual, "username password")
			})
		})
	})
}

func TestHTTP_prepareHeaders(t *testing.T) {

	Convey("Given I create an authenticated session", t, func() {

		mm, _ := New(
			context.Background(),
			"http://fake.com",
			OptionCredentials("username", "password"),
			OptionNamespace("myns"),
			OptionAdditonalHeaders(http.Header{
				"Header-1": []string{"hey"},
				"Header-2": []string{"ho"},
			}),
		)
		m := mm.(*httpManipulator)

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
					manipulate.ContextOptionCredentials("username", "password"),
					manipulate.ContextOptionClientIP("10.1.1.1"),
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

				Convey("Then I should have a value for the Authorization", func() {
					So(req.Header.Get("Authorization"), ShouldResemble, "username password")
				})

				Convey("Then I should have a value for the client IP in X-Forwarded-For", func() {
					So(req.Header["X-Forwarded-For"], ShouldResemble, []string{"10.1.1.1"})
				})
			})
		})
	})
}

func TestHTTP_readHeaders(t *testing.T) {

	Convey("Given I create a new HTTP manipulator and a Context", t, func() {

		mm, _ := New(
			context.Background(),
			"http://fake.com",
			OptionCredentials("username", "password"),
		)
		m := mm.(*httpManipulator)

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

func TestHTTP_getPersonalURL(t *testing.T) {

	Convey("Given I create a new manipulator and an object", t, func() {

		list := testmodel.NewList()

		mm, _ := New(
			context.Background(),
			"http://url.com",
			OptionCredentials("username", "password"),
		)
		m := mm.(*httpManipulator)

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
func TestHTTP_setPassword(t *testing.T) {

	Convey("Given I have a manipulator", t, func() {

		m, _ := New(context.Background(), "toto.com")

		Convey("When I call setPassword", func() {

			m.(*httpManipulator).setPassword("secret")

			Convey("Then it should set the password", func() {
				So(m.(*httpManipulator).currentPassword(), ShouldEqual, "secret")
			})
		})
	})
}

func TestHTTP_renewNotifiers(t *testing.T) {

	Convey("Given I have a manipulator", t, func() {

		mm, _ := New(context.Background(), "toto.com")
		m := mm.(*httpManipulator)

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
