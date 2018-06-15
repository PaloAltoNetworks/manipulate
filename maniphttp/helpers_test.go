package maniphttp

import (
	"sync"
	"testing"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"

	. "github.com/smartystreets/goconvey/convey"
)

type fakeManipulator struct{}

func (*fakeManipulator) RetrieveMany(*manipulate.Context, elemental.Identifiables) error {
	return nil
}
func (*fakeManipulator) Retrieve(*manipulate.Context, ...elemental.Identifiable) error { return nil }
func (*fakeManipulator) Create(*manipulate.Context, ...elemental.Identifiable) error   { return nil }
func (*fakeManipulator) Update(*manipulate.Context, ...elemental.Identifiable) error   { return nil }
func (*fakeManipulator) Delete(*manipulate.Context, ...elemental.Identifiable) error   { return nil }
func (*fakeManipulator) DeleteMany(*manipulate.Context, elemental.Identity) error      { return nil }
func (*fakeManipulator) Count(*manipulate.Context, elemental.Identity) (int, error)    { return 0, nil }

func TestManiphttp_ExtractCredentials(t *testing.T) {

	Convey("Given I have an httpmanipulator with credentials", t, func() {

		m := &httpManipulator{
			renewLock: &sync.Mutex{},
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
