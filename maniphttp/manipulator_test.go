// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package maniphttp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/elemental/test/model"
	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHTTP_NewSHTTPStore(t *testing.T) {

	Convey("When I create a new HTTPStore", t, func() {

		store := NewHTTPManipulator("username", "password", "http://url.com", "myns").(*httpManipulator)

		Convey("Then the property Username should be 'username'", func() {
			So(store.username, ShouldEqual, "username")
		})

		Convey("Then the property Password should be 'password'", func() {
			So(store.password, ShouldEqual, "password")
		})

		Convey("Then the property URL should be 'http://url.com'", func() {
			So(store.url, ShouldEqual, "http://url.com")
		})

		Convey("Then the property namespace should be 'myns'", func() {
			So(store.namespace, ShouldEqual, "myns")
		})

		Convey("Then the it should implement Manpilater interface", func() {

			var i interface{} = store
			var ok bool
			_, ok = i.(manipulate.Manipulator)
			So(ok, ShouldBeTrue)
		})
	})
}

/*
	Privates
*/
func TestHTTP_makeAuthorizationHeaders(t *testing.T) {

	Convey("Given I create a new HTTPStore", t, func() {

		Convey("When I prepare the Authorization", func() {

			store := NewHTTPManipulator("username", "password", "http://url.com", "").(*httpManipulator)
			h := store.makeAuthorizationHeaders()

			Convey("Then the header should be correct", func() {
				So(h, ShouldEqual, "username password")
			})
		})
	})
}

func TestHTTP_prepareHeaders(t *testing.T) {

	Convey("Given I create an authenticated session", t, func() {

		store := NewHTTPManipulator("username", "password", "http://fake.com", "myns").(*httpManipulator)

		Convey("Given I create a Request", func() {

			req, _ := http.NewRequest("GET", "http://fake.com", nil)

			Convey("When I prepareHeaders with a no context", func() {

				store.prepareHeaders(req, manipulate.NewContext())

				Convey("Then I should have a the X-Namespace set to 'myns'", func() {
					So(req.Header.Get("X-Namespace"), ShouldEqual, "myns")
				})

				Convey("Then I should not have a value for X-Count-Total", func() {
					So(req.Header.Get("X-Count-Total"), ShouldEqual, "")
				})
			})
		})
	})
}

func TestHTTP_readHeaders(t *testing.T) {

	Convey("Given I create a new HTTPStore an a Context", t, func() {

		store := NewHTTPManipulator("username", "password", "http://fake.com", "").(*httpManipulator)
		ctx := manipulate.NewContext()
		req := &http.Response{Header: http.Header{}}

		Convey("When I readHeaders with a no context", func() {

			store.readHeaders(req, nil)

			Convey("Then nothing should happen", func() {
			})
		})

		Convey("When I readHeaders with a request that has information", func() {

			req.Header.Set("X-Count-Total", "456")

			store.readHeaders(req, ctx)

			Convey("Then Context.X-Count-Total should be 456", func() {
				So(ctx.CountTotal, ShouldEqual, 456)
			})
		})
	})
}

func TestHTTP_standardURI(t *testing.T) {

	Convey("Given I create a new Session and an object", t, func() {

		list := testmodel.NewList()

		store := NewHTTPManipulator("username", "password", "http://url.com", "").(*httpManipulator)

		Convey("When I check personal URI of a standard object with an ID", func() {

			list.SetIdentifier("xxx")

			Convey("When I use no version", func() {

				url, err := store.getPersonalURL(list, 0)

				Convey("Then it should be http://url.com/v/1/lists/xxx", func() {
					So(url, ShouldEqual, "http://url.com/v/1/lists/xxx")
				})

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("When I use a version", func() {

				url, err := store.getPersonalURL(list, 12)

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

				url := store.getGeneralURL(list, 0)

				Convey("Then it should be http://url.com/v/1/lists", func() {
					So(url, ShouldEqual, "http://url.com/v/1/lists")
				})
			})

			Convey("When I use a version", func() {

				url := store.getGeneralURL(list, 12)

				Convey("Then it should be http://url.com/v/12/lists", func() {
					So(url, ShouldEqual, "http://url.com/v/12/lists")
				})
			})
		})

		Convey("When I check children URL for a root object", func() {

			list.SetIdentifier("xxx")

			Convey("When I use no version", func() {

				url, err := store.getURLForChildrenIdentity(nil, testmodel.ListIdentity, 0, 0)

				Convey("Then URL of the children with ListIdentity should be http://url.com/lists", func() {
					So(url, ShouldEqual, "http://url.com/lists")
				})

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("When I use a version", func() {

				url, err := store.getURLForChildrenIdentity(nil, testmodel.ListIdentity, 0, 12)

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

				url, err := store.getURLForChildrenIdentity(list, testmodel.TaskIdentity, 0, 0)

				Convey("Then URL of the children with FakeRootIdentity should be http://url.com/v/1/lists/xxx/tasks", func() {
					So(url, ShouldEqual, "http://url.com/v/1/lists/xxx/tasks")
				})

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("When I use a version", func() {

				url, err := store.getURLForChildrenIdentity(list, testmodel.TaskIdentity, 0, 12)

				Convey("Then URL of the children with FakeRootIdentity should be http://url.com/v/12/lists/xxx/tasks", func() {
					So(url, ShouldEqual, "http://url.com/v/12/lists/xxx/tasks")
				})

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})

		})

		Convey("When I check the general URL of a standard object without an ID", func() {

			url, err := store.getURLForChildrenIdentity(list, testmodel.TaskIdentity, 0, 0)

			Convey("Then it should be ''", func() {
				So(url, ShouldEqual, "")
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I check the children URL for a standard object without an ID", func() {

			url, err := store.getURLForChildrenIdentity(list, testmodel.TaskIdentity, 0, 0)

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

	Convey("Given I have a session and a and working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"ID":"xxx","name":"the list 1"}`)
		}))
		defer ts.Close()

		store := NewHTTPManipulator("username", "password", ts.URL, "")

		Convey("When I fetch an entity", func() {

			list := testmodel.NewList()
			list.ID = "xxx"
			errs := store.Retrieve(nil, list)

			Convey("Then err should be nil", func() {
				So(errs, ShouldBeNil)
			})

			Convey("Then Name should pedro", func() {
				So(list.Name, ShouldEqual, "the list 1")
			})
		})

		Convey("When I fetch an entity with no ID", func() {

			list := testmodel.NewList()
			err := store.Retrieve(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotBuildQuery{})
			})
		})
	})

	Convey("Given I have a session and a and the server will return an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(422)
			fmt.Fprint(w, `[{"code":422,"description":"nope.","subject":"elemental","title":"Read Only Error","data":null}]`)
		}))
		defer ts.Close()

		store := NewHTTPManipulator("username", "password", ts.URL, "")

		Convey("When I fetch an entity", func() {

			list := testmodel.NewList()
			list.ID = "x"
			err := store.Retrieve(nil, list)

			Convey("Then error should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.(elemental.Errors).Code(), ShouldEqual, 422)
				So(err.(elemental.Errors)[0].(elemental.Error).Description, ShouldEqual, "nope.")
			})
		})
	})

	Convey("Given I have a session and a and the server will return bad json", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `not good at all`)
		}))
		defer ts.Close()

		store := NewHTTPManipulator("username", "password", ts.URL, "")

		Convey("When I fetch an entity", func() {

			list := testmodel.NewList()
			list.ID = "x"
			err := store.Retrieve(nil, list)

			Convey("Then error should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotUnmarshal{})
			})
		})
	})
}

func TestHTTP_Update(t *testing.T) {

	Convey("Given I have a session and a and working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"ID": "zzz", "parentType": "pedro", "parentID": "yyy"}`)
		}))
		defer ts.Close()

		store := NewHTTPManipulator("username", "password", ts.URL, "")

		Convey("When I save an entity", func() {

			list := testmodel.NewList()
			list.ID = "yyy"
			errs := store.Update(nil, list)

			Convey("Then err should be nil", func() {
				So(errs, ShouldBeNil)
			})

			Convey("Then ID should 'zzz'", func() {
				So(list.Identifier(), ShouldEqual, "zzz")
			})
		})

		Convey("When I save an entity with no ID", func() {

			list := testmodel.NewList()
			errs := store.Update(nil, list)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I save an unmarshalable entity", func() {

			list := testmodel.NewUnmarshalableList()
			list.ID = "yyy"
			errs := store.Update(nil, list)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a session and the server returns an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, "nope", 500)
		}))
		defer ts.Close()

		store := NewHTTPManipulator("username", "password", ts.URL, "")

		Convey("When I save an entity", func() {

			list := testmodel.NewList()
			list.ID = "ddd"
			err := store.Update(nil, list)

			Convey("Then error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a session and the server returns a bad json", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `bad json`)
		}))
		defer ts.Close()

		store := NewHTTPManipulator("username", "password", ts.URL, "")

		Convey("When I save an entity", func() {

			list := testmodel.NewList()
			list.ID = "x"
			err := store.Update(nil, list)

			Convey("Then the error should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotUnmarshal{})
			})
		})
	})

	Convey("Given I have a session and the server returns no data", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
		}))
		defer ts.Close()

		store := NewHTTPManipulator("username", "password", ts.URL, "")

		Convey("When I save an entity", func() {

			list := testmodel.NewList()

			Convey("Then it not should panic", func() {
				So(func() { _ = store.Update(nil, list) }, ShouldNotPanic)
			})
		})
	})
}

func TestHTTP_Delete(t *testing.T) {

	Convey("Given I have a session and a and working server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `[{"ID": "yyy"}]`)
		}))
		defer ts.Close()

		store := NewHTTPManipulator("username", "password", ts.URL, "")

		Convey("When I delete an entity", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			_ = store.Delete(nil, list)

			Convey("Then ID should 'xxx'", func() {
				So(list.Identifier(), ShouldEqual, "xxx")
			})
		})

		Convey("When I delete an entity with no ID", func() {

			store := NewHTTPManipulator("username", "password", "http://fake.com", "")

			list := testmodel.NewList()
			err := store.Delete(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a session and the server returns an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, "nope", 500)
		}))
		defer ts.Close()

		store := NewHTTPManipulator("username", "password", ts.URL, "")

		Convey("When I delete an entity", func() {

			list := testmodel.NewList()
			list.ID = "xxx"

			err := store.Delete(nil, list)

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

			store := NewHTTPManipulator("username", "password", ts.URL, "")

			var l testmodel.TasksList
			ctx := manipulate.NewContext()
			ctx.Parent = list
			errs := store.RetrieveMany(ctx, &l)

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

			store := NewHTTPManipulator("username", "password", "http://fake.com", "")

			list2 := testmodel.NewList()
			var l testmodel.TasksList

			ctx := manipulate.NewContext()
			ctx.Parent = list2

			errs := store.RetrieveMany(ctx, &l)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I fetch its children while there is no data", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
			}))
			defer ts.Close()

			store := NewHTTPManipulator("username", "password", ts.URL, "")

			e := testmodel.NewTask()
			var l testmodel.TasksList

			ctx := manipulate.NewContext()
			ctx.Parent = e

			errs := store.RetrieveMany(ctx, &l)

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
			store := NewHTTPManipulator("username", "password", ts.URL, "")

			var l testmodel.TasksList

			ctx := manipulate.NewContext()
			ctx.Parent = list

			_ = store.RetrieveMany(ctx, &l)

			Convey("Then the length of the children list should be 0", func() {
				So(len(l), ShouldEqual, 0)
			})
		})

		Convey("When I fetch the children and I got a communication error", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, "woops", 500)
			}))
			defer ts.Close()

			store := NewHTTPManipulator("username", "password", ts.URL, "")

			ctx := manipulate.NewContext()
			ctx.Parent = list

			var l testmodel.TasksList
			errs := store.RetrieveMany(ctx, &l)

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

			store := NewHTTPManipulator("username", "password", ts.URL, "")

			ctx := manipulate.NewContext()
			ctx.Parent = list

			var l testmodel.TasksList
			errs := store.RetrieveMany(ctx, &l)

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

			store := NewHTTPManipulator("username", "password", ts.URL, "")
			task := testmodel.NewTask()
			errs := store.Create(nil, list, task)

			Convey("Then the error should not be nil", func() {
				So(errs, ShouldBeNil)
			})

			Convey("Then ID of the new children should be zzz", func() {
				So(task.Identifier(), ShouldEqual, "zzz")
			})
		})

		Convey("When I create a child for a parent that has no ID", func() {

			store := NewHTTPManipulator("username", "password", "url.com", "")
			list2 := testmodel.NewList()
			task := testmodel.NewTask()
			ctx := manipulate.NewContext()
			ctx.Parent = list2

			errs := store.Create(ctx, task)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I create a child that is nil", func() {

			store := NewHTTPManipulator("username", "password", "http://fake.com", "")
			task := testmodel.NewUnmarshalableList() // c'mon, that's fine..
			errs := store.Create(nil, list, task)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I create a child and I got a communication error", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "", 500)
			}))
			defer ts.Close()

			store := NewHTTPManipulator("username", "password", ts.URL, "")
			task := testmodel.NewTask()
			errs := store.Create(nil, list, task)

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

			store := NewHTTPManipulator("username", "password", ts.URL, "")
			task := testmodel.NewTask()
			errs := store.Create(nil, list, task)

			Convey("Then the error should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I create an unmarshalable entity", func() {

			store := NewHTTPManipulator("username", "password", "", "")
			list := testmodel.NewUnmarshalableList()
			list.ID = "yyy"
			err := store.Create(nil, list)

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

			store := NewHTTPManipulator("username", "password", ts.URL, "")

			ctx := manipulate.NewContext()
			ctx.Parent = list
			num, errs := store.Count(ctx, testmodel.TaskIdentity)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldBeNil)
			})

			Convey("Then count should be 10", func() {
				So(num, ShouldEqual, 10)
			})
		})

		Convey("When I fetch its children but the parent has no ID", func() {

			store := NewHTTPManipulator("username", "password", "http://fake.com", "")

			list2 := testmodel.NewList()

			ctx := manipulate.NewContext()
			ctx.Parent = list2

			_, errs := store.Count(ctx, testmodel.TaskIdentity)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I fetch the children and I got a communication error", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, "woops", 500)
			}))
			defer ts.Close()

			store := NewHTTPManipulator("username", "password", ts.URL, "")

			ctx := manipulate.NewContext()
			ctx.Parent = list

			_, errs := store.Count(ctx, testmodel.TaskIdentity)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_Increment(t *testing.T) {

	Convey("Given I have a store", t, func() {

		store := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call Count", func() {
			err := store.Increment(nil, testmodel.ListIdentity, "counter", 1)

			Convey("Then err should should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrNotImplemented{})
			})
		})
	})
}

func TestHTTP_send(t *testing.T) {

	Convey("Given I have a store with bad url", t, func() {

		store := NewHTTPManipulator("username", "password", "", "")

		Convey("When I call send ", func() {

			req, _ := http.NewRequest(http.MethodPost, "nop", nil)
			_, err := store.(*httpManipulator).send(req, manipulate.NewContext())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommunicate{})
			})
		})
	})
}
