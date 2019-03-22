package maniphttp

import (
	"crypto/tls"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSubscriberOption(t *testing.T) {

	Convey("Given I have manipulator", t, func() {

		m := &httpManipulator{
			url:       "https://toto.com",
			namespace: "mns",
			tlsConfig: &tls.Config{},
		}

		Convey("When I newSubscribeConfig ", func() {

			cfg := newSubscribeConfig(m)

			Convey("Then cfg should be correct", func() {
				So(cfg.recursive, ShouldBeFalse)
				So(cfg.endpoint, ShouldEqual, "events")
				So(cfg.namespace, ShouldEqual, "mns")
				So(cfg.tlsConfig, ShouldEqual, m.tlsConfig)
			})

		})
	})
}

func TestOptions(t *testing.T) {

	m := &httpManipulator{
		namespace: "mns",
		tlsConfig: &tls.Config{},
	}

	Convey("SubscriberOptionRecursive should work", t, func() {
		cfg := newSubscribeConfig(m)
		SubscriberOptionRecursive(true)(&cfg)
		So(cfg.recursive, ShouldBeTrue)
	})

	Convey("SubscriberOptionNamespace should work", t, func() {
		cfg := newSubscribeConfig(m)
		SubscriberOptionNamespace("/toto")(&cfg)
		So(cfg.namespace, ShouldEqual, "/toto")
	})

	Convey("SubscriberOptionEndpoint should work", t, func() {
		cfg := newSubscribeConfig(m)
		SubscriberOptionEndpoint("/labas/")(&cfg)
		So(cfg.endpoint, ShouldEqual, "labas")
	})
}
