package manipmongo

import (
	"crypto/tls"
	"testing"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
)

type fakeSharder struct{}

func (*fakeSharder) Shard(elemental.Identifiable)            {}
func (*fakeSharder) FilterOne(elemental.Identifiable) bson.M { return nil }
func (*fakeSharder) FilterMany(elemental.Identity) bson.M    { return nil }

func TestManipMongo_newConfig(t *testing.T) {

	Convey("Given call newConfig", t, func() {

		c := newConfig()

		Convey("Then I should get the default config", func() {
			So(c.username, ShouldEqual, "")
			So(c.password, ShouldEqual, "")
			So(c.tlsConfig, ShouldEqual, nil)
			So(c.poolLimit, ShouldEqual, 4096)
			So(c.connectTimeout, ShouldEqual, 10*time.Second)
			So(c.socketTimeout, ShouldEqual, 60*time.Second)
			So(c.mode, ShouldEqual, mgo.Strong)
		})
	})
}

func TestManipMongo_Options(t *testing.T) {

	Convey("Calling OptionCredentials should work", t, func() {
		c := newConfig()
		OptionCredentials("user", "password", "authdb")(c)
		So(c.username, ShouldEqual, "user")
		So(c.password, ShouldEqual, "password")
		So(c.authsource, ShouldEqual, "authdb")
	})

	Convey("Calling OptionTLS should work", t, func() {
		c := newConfig()
		t := &tls.Config{}
		OptionTLS(t)(c)
		So(c.tlsConfig, ShouldEqual, t)
	})

	Convey("Calling OptionConnectionPoolLimit should work", t, func() {
		c := newConfig()
		OptionConnectionPoolLimit(12)(c)
		So(c.poolLimit, ShouldEqual, 12)
	})

	Convey("Calling OptionConnectionTimeout should work", t, func() {
		c := newConfig()
		OptionConnectionTimeout(12 * time.Second)(c)
		So(c.connectTimeout, ShouldEqual, 12*time.Second)
	})

	Convey("Calling OptionSocketTimeout should work", t, func() {
		c := newConfig()
		OptionSocketTimeout(12 * time.Second)(c)
		So(c.socketTimeout, ShouldEqual, 12*time.Second)
	})

	Convey("Calling OptionDefaultConsistencyMode should work", t, func() {
		c := newConfig()
		OptionDefaultConsistencyMode(mgo.Nearest)(c)
		So(c.mode, ShouldEqual, mgo.Nearest)
	})

	Convey("Calling OptionSharder should work", t, func() {
		c := newConfig()
		s := &fakeSharder{}
		OptionSharder(s)(c)
		So(c.sharder, ShouldEqual, s)
	})
}
