package manipmongo

import (
	"testing"

	"github.com/aporeto-inc/elemental/test/model"
	. "github.com/smartystreets/goconvey/convey"
	mgo "gopkg.in/mgo.v2"
)

func TestUtils_collectionFromIdentity(t *testing.T) {

	Convey("Given I a mgo.Database and an identity", t, func() {

		Convey("When I use collectionFromIdentity", func() {

			db := &mgo.Database{}

			c := collectionFromIdentity(db, testmodel.ListIdentity)

			Convey("Then collection should not be nil", func() {
				So(c, ShouldNotBeNil)
			})

			Convey("Then collection fullName should be nil", func() {
				So(c.FullName, ShouldEqual, ".list")
			})
		})
	})
}
