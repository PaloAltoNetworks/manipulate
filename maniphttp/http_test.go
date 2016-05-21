// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package maniphttp

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHTTP_NewSHTTPStore(t *testing.T) {

	Convey("When I create a new HTTPStore", t, func() {

		store := NewHTTPStore("username", "password", "http://url.com", "")

		Convey("Then the property Username should be 'username'", func() {
			So(store.username, ShouldEqual, "username")
		})

		Convey("Then the property Password should be 'password'", func() {
			So(store.password, ShouldEqual, "password")
		})

		Convey("Then the property URL should be 'http://url.com'", func() {
			So(store.url, ShouldEqual, "http://url.com")
		})

		Convey("Then the it should implement Manpilater interface", func() {

			var i interface{} = store
			var ok bool
			_, ok = i.(manipulate.Manipulator)
			So(ok, ShouldBeTrue)
		})
	})
}

func TestHTTP_SetInsecureSkipVerify(t *testing.T) {

	Convey("Given I create a new HTTPStore", t, func() {

		store := NewHTTPStore("username", "password", "http://url.com", "")

		Convey("When I set the insecure check skip to true", func() {

			store.SetInsecureSkipVerify(true)

			Convey("Then err should be nil", func() {
				t := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
				So(store.client.Transport, ShouldResemble, t)
			})
		})
	})
}

/*
	Privates
*/
func TestHTTP_makeAuthorizationHeaders(t *testing.T) {

	Convey("Given I create a new HTTPStore", t, func() {

		Convey("When I prepare the Authorization", func() {

			store := NewHTTPStore("username", "password", "http://url.com", "")
			h := store.makeAuthorizationHeaders()

			Convey("Then the header should be 'TODO=", func() {
				So(h, ShouldEqual, "TODO")
			})
		})
	})
}

func TestHTTP_prepareHeaders(t *testing.T) {

	Convey("Given I create an authenticated session", t, func() {

		store := NewHTTPStore("username", "password", "http://fake.com", "")

		Convey("Given I create a Request", func() {

			req, _ := http.NewRequest("GET", "http://fake.com", nil)

			Convey("When I prepareHeaders with a no context", func() {

				store.prepareHeaders(req, nil)

				Convey("Then I should not have a value for X-Page-Current", func() {
					So(req.Header.Get("X-Page-Current"), ShouldEqual, "")
				})

				Convey("Then I should not have a value for X-Page-Size", func() {
					So(req.Header.Get("X-Page-Size"), ShouldEqual, "")
				})

				Convey("Then I should not have a value for X-Page-First", func() {
					So(req.Header.Get("X-Page-First"), ShouldEqual, "")
				})

				Convey("Then I should not have a value for X-Page-Prev", func() {
					So(req.Header.Get("X-Page-Prev"), ShouldEqual, "")
				})

				Convey("Then I should not have a value for X-Page-Next", func() {
					So(req.Header.Get("X-Page-Next"), ShouldEqual, "")
				})

				Convey("Then I should not have a value for X-Page-Last", func() {
					So(req.Header.Get("X-Page-Last"), ShouldEqual, "")
				})

				Convey("Then I should not have a value for X-Count-Local", func() {
					So(req.Header.Get("X-Count-Local"), ShouldEqual, "")
				})

				Convey("Then I should not have a value for X-Count-Total", func() {
					So(req.Header.Get("X-Count-Total"), ShouldEqual, "")
				})
			})
		})

		Convey("Given I create a Request and a Context", func() {

			req, _ := http.NewRequest("GET", "http://fake.com", nil)
			ctx := manipulate.NewContext()

			Convey("When I prepareHeaders witha fetching info that has a all fields", func() {
				ctx.PageCurrent = 2
				ctx.PageSize = 42

				store.prepareHeaders(req, ctx)

				Convey("Then I should have a the X-Page-Current set to 2", func() {
					So(req.Header.Get("X-Page-Current"), ShouldEqual, "2")
				})

				Convey("Then I should have a the X-Page-Size set to 42", func() {
					So(req.Header.Get("X-Page-Size"), ShouldEqual, "42")
				})

				Convey("Then I should not have a value for X-Page-First", func() {
					So(req.Header.Get("X-Page-First"), ShouldEqual, "")
				})

				Convey("Then I should not have a value for X-Page-Prev", func() {
					So(req.Header.Get("X-Page-Prev"), ShouldEqual, "")
				})

				Convey("Then I should not have a value for X-Page-Next", func() {
					So(req.Header.Get("X-Page-Next"), ShouldEqual, "")
				})

				Convey("Then I should not have a value for X-Page-Last", func() {
					So(req.Header.Get("X-Page-Last"), ShouldEqual, "")
				})

				Convey("Then I should not have a value for X-Count-Local", func() {
					So(req.Header.Get("X-Count-Local"), ShouldEqual, "")
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

		store := NewHTTPStore("username", "password", "http://fake.com", "")
		ctx := manipulate.NewContext()
		req := &http.Response{Header: http.Header{}}

		Convey("When I readHeaders with a no context", func() {

			store.readHeaders(req, nil)

			Convey("Then nothing should happen", func() {
			})
		})

		Convey("When I readHeaders with a request that has information", func() {

			req.Header.Set("X-Page-Current", "3")
			req.Header.Set("X-Page-Size", "42")
			req.Header.Set("X-Page-First", "http://fake.com/?page=1")
			req.Header.Set("X-Page-Prev", "http://fake.com/?page=2")
			req.Header.Set("X-Page-Next", "http://fake.com/?page=4")
			req.Header.Set("X-Page-Last", "http://fake.com/?page=10")
			req.Header.Set("X-Count-Local", "123")
			req.Header.Set("X-Count-Total", "456")

			store.readHeaders(req, ctx)

			Convey("Then Context.PageCurrent should be 3", func() {
				So(ctx.PageCurrent, ShouldEqual, 3)
			})

			Convey("Then Context.PageSize should be 42", func() {
				So(ctx.PageSize, ShouldEqual, 42)
			})

			Convey("Then Context.PageFirst should be correct", func() {
				So(ctx.PageFirst, ShouldEqual, "http://fake.com/?page=1")
			})

			Convey("Then Context.PagePrev should be correct", func() {
				So(ctx.PagePrev, ShouldEqual, "http://fake.com/?page=2")
			})

			Convey("Then Context.PageNext should be correct", func() {
				So(ctx.PageNext, ShouldEqual, "http://fake.com/?page=4")
			})

			Convey("Then Context.PageLast should be correct", func() {
				So(ctx.PageLast, ShouldEqual, "http://fake.com/?page=10")
			})

			Convey("Then Context.X-Count-Local should be 123", func() {
				So(ctx.CountLocal, ShouldEqual, 123)
			})

			Convey("Then Context.X-Count-Total should be 456", func() {
				So(ctx.CountTotal, ShouldEqual, 456)
			})
		})
	})
}

func TestHTTP_standardURI(t *testing.T) {

	Convey("Given I create a new Session and an object", t, func() {

		list := NewList()

		store := NewHTTPStore("username", "password", "http://url.com", "")

		Convey("When I check personal URI of a standard object with an ID", func() {

			list.SetIdentifier("xxx")
			url, err := store.getPersonalURL(list)

			Convey("Then it should be http://url.com/lists/xxx", func() {
				So(url, ShouldEqual, "http://url.com/lists/xxx")
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I check general URI of a standard object with an ID", func() {

			list.SetIdentifier("xxx")
			url := store.getGeneralURL(list)

			Convey("Then it should be http://url.com/lists", func() {
				So(url, ShouldEqual, "http://url.com/lists")
			})
		})

		Convey("When I check children URL for a standard object with an ID", func() {

			list.SetIdentifier("xxx")
			url, err := store.getURLForChildrenIdentity(list, TaskIdentity)

			Convey("Then URL of the children with FakeRootIdentity should be http://url.com/lists/xxx/tasks", func() {
				So(url, ShouldEqual, "http://url.com/lists/xxx/tasks")
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I check the general URL of a standard object without an ID", func() {

			url, err := store.getURLForChildrenIdentity(list, TaskIdentity)

			Convey("Then it should be ''", func() {
				So(url, ShouldEqual, "")
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I check the children URL for a standard object without an ID", func() {

			url, err := store.getURLForChildrenIdentity(list, TaskIdentity)

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

		store := NewHTTPStore("username", "password", ts.URL, "")

		Convey("When I fetch an entity", func() {

			list := NewList()
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

			list := NewList()
			errs := store.Retrieve(nil, list)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a session and a and the server will return an error", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, "bad comm", 500)
		}))
		defer ts.Close()

		store := NewHTTPStore("username", "password", ts.URL, "")

		Convey("When I fetch an entity", func() {

			list := NewList()
			errs := store.Retrieve(nil, list)

			Convey("Then error should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a session and a and the server will return bad json", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `not good at all`)
		}))
		defer ts.Close()

		store := NewHTTPStore("username", "password", ts.URL, "")

		Convey("When I fetch an entity", func() {

			list := NewList()
			errs := store.Retrieve(nil, list)

			Convey("Then error should not be nil", func() {
				So(errs, ShouldNotBeNil)
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

		store := NewHTTPStore("username", "password", ts.URL, "")

		Convey("When I save an entity", func() {

			list := NewList()
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

			list := NewList()
			errs := store.Update(nil, list)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I save an unmarshalable entity", func() {

			list := NewUnmarshalableList()
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

		store := NewHTTPStore("username", "password", ts.URL, "")

		Convey("When I save an entity", func() {

			list := NewList()
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

		store := NewHTTPStore("username", "password", ts.URL, "")

		Convey("When I save an entity", func() {

			list := NewList()
			err := store.Update(nil, list)

			Convey("Then the error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a session and the server returns no data", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
		}))
		defer ts.Close()

		store := NewHTTPStore("username", "password", ts.URL, "")

		Convey("When I save an entity", func() {

			list := NewList()

			Convey("Then it not should panic", func() {
				So(func() { store.Update(nil, list) }, ShouldNotPanic)
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

		store := NewHTTPStore("username", "password", ts.URL, "")

		Convey("When I delete an entity", func() {

			list := NewList()
			list.ID = "xxx"

			store.Delete(nil, list)

			Convey("Then ID should 'xxx'", func() {
				So(list.Identifier(), ShouldEqual, "xxx")
			})
		})

		Convey("When I delete an entity with no ID", func() {

			store := NewHTTPStore("username", "password", "http://fake.com", "")

			list := NewList()
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

		store := NewHTTPStore("username", "password", ts.URL, "")

		Convey("When I delete an entity", func() {

			list := NewList()
			list.ID = "xxx"

			err := store.Delete(nil, list)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_RetrieveChildren(t *testing.T) {

	Convey("Given I have an existing object", t, func() {

		list := NewList()
		list.ID = "xxx"

		Convey("When I fetch its children with success", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `[{"ID": "1", "name": "name1"}, {"ID": "2", "name": "name2"}]`)
			}))
			defer ts.Close()

			store := NewHTTPStore("username", "password", ts.URL, "")

			var l TasksList
			errs := store.RetrieveChildren(nil, list, TaskIdentity, &l)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldBeNil)
			})

			Convey("Then the lenght of the children list should be 2", func() {
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
				So(l[0].Identity(), ShouldResemble, TaskIdentity)
				So(l[1].Identity(), ShouldResemble, TaskIdentity)
			})
		})

		Convey("When I fetch its children but the parent has no ID", func() {

			store := NewHTTPStore("username", "password", "http://fake.com", "")

			list2 := NewList()
			var l TasksList
			errs := store.RetrieveChildren(nil, list2, TaskIdentity, &l)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I fetch its children while there is no data", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
			}))
			defer ts.Close()

			store := NewHTTPStore("username", "password", ts.URL, "")

			e := NewTask()
			var l TasksList
			errs := store.RetrieveChildren(nil, e, TaskIdentity, &l)

			Convey("Then the lenght of the children list should be 0", func() {
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
			store := NewHTTPStore("username", "password", ts.URL, "")

			var l TasksList
			store.RetrieveChildren(nil, list, TaskIdentity, &l)

			Convey("Then the lenght of the children list should be 0", func() {
				So(len(l), ShouldEqual, 0)
			})
		})

		Convey("When I fetch the children and I got a communication error", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, "woops", 500)
			}))
			defer ts.Close()

			store := NewHTTPStore("username", "password", ts.URL, "")

			var l TasksList
			errs := store.RetrieveChildren(nil, list, TaskIdentity, &l)

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

			store := NewHTTPStore("username", "password", ts.URL, "")

			var l TasksList
			errs := store.RetrieveChildren(nil, list, TaskIdentity, &l)

			Convey("Then the error should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_Create(t *testing.T) {

	Convey("Given I have an existing object", t, func() {

		list := NewList()
		list.ID = "xxx"

		Convey("When I create a child with success", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				fmt.Fprint(w, `{"ID": "zzz"}`)
			}))
			defer ts.Close()

			store := NewHTTPStore("username", "password", ts.URL, "")
			task := NewTask()
			errs := store.Create(nil, list, task)

			Convey("Then the error should not be nil", func() {
				So(errs, ShouldBeNil)
			})

			Convey("Then ID of the new children should be zzz", func() {
				So(task.Identifier(), ShouldEqual, "zzz")
			})
		})

		Convey("When I create a child for a parent that has no ID", func() {

			store := NewHTTPStore("username", "password", "url.com", "")
			list2 := NewList()
			task := NewTask()
			errs := store.Create(nil, list2, task)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I create a child that is nil", func() {

			store := NewHTTPStore("username", "password", "http://fake.com", "")
			task := NewUnmarshalableList() // c'mon, that's fine..
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

			store := NewHTTPStore("username", "password", ts.URL, "")
			task := NewTask()
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

			store := NewHTTPStore("username", "password", ts.URL, "")
			task := NewTask()
			errs := store.Create(nil, list, task)

			Convey("Then the error should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})
}

func TestHTTP_Assign(t *testing.T) {

	Convey("Given I have two existing objects", t, func() {

		l := NewList()
		l.ID = "a"

		Convey("When I assign them with success", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			defer ts.Close()

			session := NewHTTPStore("username", "password", ts.URL, "")

			t1 := NewTask()
			t1.ID = "xxx"
			t2 := NewTask()
			t2.ID = "yyy"
			assignation := elemental.NewAssignation(elemental.OperationAdditive, TaskIdentity, t1, t2)
			errs := session.Assign(nil, l, assignation)

			Convey("Then err should be nil", func() {
				So(errs, ShouldBeNil)
			})
		})

		Convey("When I assign objects to a parent that has no ID", func() {

			session := NewHTTPStore("username", "password", "http://fake.com", "")

			l1 := NewList()
			t2 := NewTask()
			t2.ID = "yyy"
			assignation := elemental.NewAssignation(elemental.OperationAdditive, TaskIdentity, t2)

			errs := session.Assign(nil, l1, assignation)

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})

		Convey("When I assign them I got an communication error", func() {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "woops", 500)
			}))
			defer ts.Close()

			session := NewHTTPStore("username", "password", ts.URL, "")

			t1 := NewTask()
			t1.ID = "xxx"
			t2 := NewTask()
			t2.ID = "yyy"
			assignation := elemental.NewAssignation(elemental.OperationAdditive, TaskIdentity, t1, t2)
			errs := session.Assign(nil, l, assignation)

			Convey("Then errs should not be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})
}

//
// /*
// 	Events
// */
// func TestHTTP_NextEvent(t *testing.T) {
//
// 	Convey("When I use NextEvent and I receive a valid push notification", t, func() {
//
// 		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.Header().Set("Content-Type", "application/json")
// 			fmt.Fprint(w, `{"uuid": "y", "events": [{"type": "CREATE", "entityType": "thing", "updateMechanism": "DEFAULT", "entities": []}]}`)
// 		}))
// 		defer ts.Close()
//
// 		r := NewFakeRootObject()
//
// 		session := NewHTTPStore("username", "password", "organization", ts.URL, r, "")
//
// 		lID := "x"
// 		var notif *Notification
// 		c := make(NotificationsChannel)
// 		go session.NextEvent(c, lID)
//
// 		select {
// 		case notif = <-c:
// 		case <-time.After(10 * time.Millisecond):
// 		}
//
// 		Convey("Then notification should not be nil", func() {
// 			So(notif, ShouldNotBeNil)
// 		})
//
// 		Convey("Then last Event ID should be y", func() {
// 			So(notif.UUID, ShouldEqual, "y")
// 		})
// 	})
