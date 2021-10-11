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
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniptest"
)

func TestManiphttp_ExtractCredentials(t *testing.T) {

	Convey("Given I have an httpmanipulator with credentials", t, func() {

		m := &httpManipulator{
			renewLock: sync.RWMutex{},
			username:  "a",
			password:  "b",
		}

		Convey("When I call ExtractCredentials", func() {

			u, p := ExtractCredentials(m)

			Convey("Then I should get the creds", func() {
				So(u, ShouldEqual, "a")
				So(p, ShouldEqual, "b")
			})
		})
	})

	Convey("Given I have a non http manipulator with credentials", t, func() {

		m := maniptest.NewTestManipulator()

		Convey("When I call ExtractCredentials", func() {

			Convey("Then it should panic", func() {
				So(func() { ExtractCredentials(m) }, ShouldPanicWith, "You can only pass a HTTP Manipulator to ExtractCredentials")
			})
		})
	})
}

func TestManiphttp_ExtractEndpoint(t *testing.T) {

	Convey("Given I have an httpmanipulator with endpoint", t, func() {

		m := &httpManipulator{
			url: "https://toto.com",
		}

		Convey("When I call ExtractEndpoint", func() {

			u := ExtractEndpoint(m)

			Convey("Then I should get the creds", func() {
				So(u, ShouldEqual, "https://toto.com")
			})
		})
	})

	Convey("Given I have a non http manipulator", t, func() {

		m := maniptest.NewTestManipulator()

		Convey("When I call ExtractEndpoint", func() {

			Convey("Then it should panic", func() {
				So(func() { ExtractEndpoint(m) }, ShouldPanicWith, "You can only pass a HTTP Manipulator to ExtractEndpoint")
			})
		})
	})
}

func TestManiphttp_ExtractNamespace(t *testing.T) {

	Convey("Given I have an httpmanipulator with namespace", t, func() {

		m := &httpManipulator{
			namespace: "/toto",
		}

		Convey("When I call ExtractNamespace", func() {

			u := ExtractNamespace(m)

			Convey("Then I should get the creds", func() {
				So(u, ShouldEqual, "/toto")
			})
		})
	})

	Convey("Given I have a non http manipulator", t, func() {

		m := maniptest.NewTestManipulator()

		Convey("When I call ExtractEndpoint", func() {

			Convey("Then it should panic", func() {
				So(func() { ExtractNamespace(m) }, ShouldPanicWith, "You can only pass a HTTP Manipulator to ExtractNamespace")
			})
		})
	})
}

func TestManiphttp_ExtractTLSConfig(t *testing.T) {

	Convey("Given I have an httpmanipulator with namespace", t, func() {

		tlsc := &tls.Config{}
		m := &httpManipulator{
			tlsConfig: tlsc,
		}

		Convey("When I call ExtractTLSConfig", func() {

			u := ExtractTLSConfig(m)

			Convey("Then I should get the tls config", func() {
				So(u, ShouldResemble, tlsc)
				So(u, ShouldNotEqual, tlsc)

			})
		})
	})

	Convey("Given I have a non http manipulator", t, func() {

		m := maniptest.NewTestManipulator()

		Convey("When I call ExtractEndpoint", func() {

			Convey("Then it should panic", func() {
				So(func() { ExtractTLSConfig(m) }, ShouldPanicWith, "You can only pass a HTTP Manipulator to ExtractTLSConfig")
			})
		})
	})
}

func TestManiphttp_SetGlobalHeaders(t *testing.T) {

	Convey("Given I have a manipulator and some header", t, func() {

		m := &httpManipulator{
			renewLock: sync.RWMutex{},
		}

		h := http.Header{
			"Header-1": []string{"hey"},
			"Header-2": []string{"ho"},
		}

		Convey("When I call SetGlobalHeaders", func() {

			SetGlobalHeaders(m, h)

			Convey("Then hte global header should be set", func() {
				So(m.globalHeaders, ShouldResemble, h)
			})
		})
	})

	Convey("Given I have a non http manipulator", t, func() {

		m := maniptest.NewTestManipulator()

		Convey("When I call SetGlobalHeaders", func() {

			Convey("Then it should panic", func() {
				So(func() { SetGlobalHeaders(m, nil) }, ShouldPanicWith, "You can only pass a HTTP Manipulator to SetGlobalHeaders")
			})
		})
	})
}

func TestManiphttp_DirectSend(t *testing.T) {

	Convey("Given I have a manipulator and a test server", t, func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			if string(body) != "hello" {
				panic("wrong body received.")
			}

			if !strings.HasSuffix(r.RequestURI, "/toto") {
				panic(fmt.Sprintf("wrong url: %s", r.RequestURI))
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `"bonjour"`)
		}))
		defer ts.Close()

		m, _ := New(context.Background(), ts.URL)

		Convey("When I call DirectSend", func() {

			resp, err := DirectSend(m, nil, "toto", http.MethodPost, []byte("hello"))

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the response should be correct", func() {
				So(resp, ShouldNotBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
			})
		})
	})

	Convey("Given I have a non http manipulator", t, func() {

		m := maniptest.NewTestManipulator()

		Convey("When I call DirectSend", func() {

			Convey("Then it should panic", func() {
				So(func() { _, _ = DirectSend(m, nil, "", http.MethodPost, nil) }, ShouldPanicWith, "You can only pass a HTTP Manipulator to DirectSend")
			})
		})
	})
}

func TestManiphttp_BatchCreate(t *testing.T) {

	Convey("Given I have a manipulator and a test server and I use JSON encoding", t, func() {

		l1 := testmodel.NewList()
		l1.Name = "a"
		l2 := testmodel.NewList()
		l2.Name = "b"
		l3 := testmodel.NewList()
		l3.Name = "c"

		expectedBody := bytes.NewBuffer(nil)
		enc, cl := elemental.MakeStreamEncoder(elemental.EncodingTypeJSON, expectedBody)
		defer cl()
		_ = enc(l1)
		_ = enc(l2)
		_ = enc(l3)

		errCh := make(chan error, 1)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				errCh <- err
			}

			expectedEncoding := "application/json+batch"
			if ct := r.Header.Get("Content-Type"); ct != expectedEncoding {
				errCh <- fmt.Errorf("wrong content type received. want: '%s', got: '%s'", expectedEncoding, ct)
			}

			if !bytes.Equal(body, expectedBody.Bytes()) {
				errCh <- fmt.Errorf("wrong body received. want:\n%s\n\ngot:\n%s", string(body), expectedBody.String())
			}

			w.WriteHeader(http.StatusOK)
			errCh <- nil
		}))
		defer ts.Close()

		m, _ := New(
			context.Background(),
			ts.URL,
			OptionEncoding(elemental.EncodingTypeJSON),
		)

		Convey("When I call BatchCreate", func() {

			resp, err := BatchCreate(m, nil, l1, l2, l3)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the response should be correct", func() {
				So(resp, ShouldNotBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(<-errCh, ShouldBeNil)
			})
		})

		Convey("When I call BatchCreate with an unmarshallable object", func() {

			resp, err := BatchCreate(m, nil, testmodel.NewUnmarshalableList())

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `unable to encode application/json: json encode error: error marshalling`)
			})

			Convey("Then the response should be correct", func() {
				So(resp, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and a test server and I use MSGPACK encoding", t, func() {

		l1 := testmodel.NewList()
		l1.Name = "a"
		l2 := testmodel.NewList()
		l2.Name = "b"
		l3 := testmodel.NewList()
		l3.Name = "c"

		expectedBody := bytes.NewBuffer(nil)
		enc, cl := elemental.MakeStreamEncoder(elemental.EncodingTypeMSGPACK, expectedBody)
		defer cl()
		_ = enc(l1)
		_ = enc(l2)
		_ = enc(l3)

		errCh := make(chan error, 10)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				errCh <- err
			}

			expectedEncoding := "application/msgpack+batch"
			if ct := r.Header.Get("Content-Type"); ct != expectedEncoding {
				errCh <- fmt.Errorf("wrong content type received. want: '%s', got: '%s'", expectedEncoding, ct)
			}

			if !bytes.Equal(body, expectedBody.Bytes()) {
				errCh <- fmt.Errorf("wrong body received. want:\n%s\n\ngot:\n%s", string(body), expectedBody.String())
				return
			}

			w.WriteHeader(http.StatusOK)
			errCh <- nil
		}))
		defer ts.Close()

		m, _ := New(
			context.Background(),
			ts.URL,
			OptionEncoding(elemental.EncodingTypeMSGPACK),
		)

		Convey("When I call BatchCreate", func() {

			resp, err := BatchCreate(m, nil, l1, l2, l3)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the response should be correct", func() {
				So(resp, ShouldNotBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(<-errCh, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a manipulator and a test server and I use custom encoding", t, func() {

		l1 := testmodel.NewList()
		l1.Name = "a"
		l2 := testmodel.NewList()
		l2.Name = "b"
		l3 := testmodel.NewList()
		l3.Name = "c"

		expectedBody := bytes.NewBuffer(nil)
		enc, cl := elemental.MakeStreamEncoder(elemental.EncodingTypeJSON, expectedBody)
		defer cl()
		_ = enc(l1)
		_ = enc(l2)
		_ = enc(l3)

		errCh := make(chan error, 10)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				errCh <- err
			}

			expectedEncoding := "custom+batch"
			if ct := r.Header.Get("Content-Type"); ct != expectedEncoding {
				errCh <- fmt.Errorf("wrong content type received. want: '%s', got: '%s'", expectedEncoding, ct)
			}

			if !bytes.Equal(body, expectedBody.Bytes()) {
				errCh <- fmt.Errorf("wrong body received. want:\n%s\n\ngot:\n%s", string(body), expectedBody.String())
				return
			}

			w.WriteHeader(http.StatusOK)
			errCh <- nil
		}))
		defer ts.Close()

		m, _ := New(
			context.Background(),
			ts.URL,
			OptionEncoding(elemental.EncodingTypeMSGPACK),
		)

		Convey("When I call BatchCreate", func() {

			mctx := manipulate.NewContext(
				context.Background(),
				ContextOptionOverrideContentType("custom"),
			)
			resp, err := BatchCreate(m, mctx, l1, l2, l3)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the response should be correct", func() {
				So(resp, ShouldNotBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(<-errCh, ShouldBeNil)
			})
		})
	})

	Convey("When I call BatchCreate with a non httpmanipulator", t, func() {

		m := maniptest.NewTestManipulator()
		l1 := testmodel.NewList()

		Convey("Then it should panic", func() {
			So(func() { _, _ = BatchCreate(m, nil, l1) }, ShouldPanicWith, "You can only pass a HTTP Manipulator to BatchCreate")
		})
	})

	Convey("When I call BatchCreate with no objects", t, func() {

		m, _ := New(
			context.Background(),
			"https://totot.com",
			OptionEncoding(elemental.EncodingTypeMSGPACK),
		)

		Convey("Then it should panic", func() {
			So(func() { _, _ = BatchCreate(m, nil) }, ShouldPanicWith, "You must pass at least one object to BatchCreate")
		})
	})

}

func TestExtractTransport(t *testing.T) {

	Convey("Given I have an httpmanipulator with namespace", t, func() {

		transport := &http.Transport{}
		m := &httpManipulator{
			transport: transport,
		}

		Convey("When I call ExtractTransport", func() {

			t := ExtractTransport(m)

			Convey("Then I should get the correct transport", func() {
				So(t, ShouldEqual, transport)
			})
		})
	})

	Convey("Given I have a non http manipulator", t, func() {

		m := maniptest.NewTestManipulator()

		Convey("When I call ExtractTransport", func() {

			Convey("Then it should panic", func() {
				So(func() { ExtractTransport(m) }, ShouldPanicWith, "You can only pass a HTTP Manipulator to ExtractTransport")
			})
		})
	})
}

func TestExtractClient(t *testing.T) {

	Convey("Given I have an httpmanipulator with namespace", t, func() {

		client := &http.Client{}
		m := &httpManipulator{
			client: client,
		}

		Convey("When I call ExtractClient", func() {

			t := ExtractClient(m)

			Convey("Then I should get the correct client", func() {
				So(t, ShouldEqual, client)
			})
		})
	})

	Convey("Given I have a non http manipulator", t, func() {

		m := maniptest.NewTestManipulator()

		Convey("When I call ExtractClient", func() {

			Convey("Then it should panic", func() {
				So(func() { ExtractClient(m) }, ShouldPanicWith, "You can only pass a HTTP Manipulator to ExtractClient")
			})
		})
	})
}
