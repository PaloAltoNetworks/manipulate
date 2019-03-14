package maniphttp

import (
	"crypto/tls"
	"net/http"
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

type fakeManipulator struct{}

func (*fakeManipulator) RetrieveMany(manipulate.Context, elemental.Identifiables) error {
	return nil
}
func (*fakeManipulator) Retrieve(manipulate.Context, ...elemental.Identifiable) error { return nil }
func (*fakeManipulator) Create(manipulate.Context, ...elemental.Identifiable) error   { return nil }
func (*fakeManipulator) Update(manipulate.Context, ...elemental.Identifiable) error   { return nil }
func (*fakeManipulator) Delete(manipulate.Context, ...elemental.Identifiable) error   { return nil }
func (*fakeManipulator) DeleteMany(manipulate.Context, elemental.Identity) error      { return nil }
func (*fakeManipulator) Count(manipulate.Context, elemental.Identity) (int, error)    { return 0, nil }

func TestManiphttp_ExtractCredentials(t *testing.T) {

	Convey("Given I have an httpmanipulator with credentials", t, func() {

		m := &httpManipulator{
			renewLock: &sync.RWMutex{},
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

		m := &fakeManipulator{}

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

		m := &fakeManipulator{}

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

		m := &fakeManipulator{}

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

		m := &fakeManipulator{}

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
			renewLock: &sync.RWMutex{},
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

		m := &fakeManipulator{}

		Convey("When I call SetGlobalHeaders", func() {

			Convey("Then it should panic", func() {
				So(func() { SetGlobalHeaders(m, nil) }, ShouldPanicWith, "You can only pass a HTTP Manipulator to SetGlobalHeaders")
			})
		})
	})
}
