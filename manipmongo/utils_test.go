package manipmongo

import (
	"testing"

	mgo "gopkg.in/mgo.v2"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtils_collectionFromIdentity(t *testing.T) {

	Convey("Given I a mgo.Database and an identity", t, func() {

		Convey("When I use collectionFromIdentity", func() {

			db := &mgo.Database{}

			c := collectionFromIdentity(db, PersonIdentity)

			Convey("Then collection should not be nil", func() {
				So(c, ShouldNotBeNil)
			})

			Convey("Then collection fullName should be nil", func() {
				So(c.FullName, ShouldEqual, ".persons")
			})
		})
	})
}
