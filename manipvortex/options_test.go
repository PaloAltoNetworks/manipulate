package manipvortex

import (
	"testing"
	"time"

	"go.aporeto.io/manipulate"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/manipulate/maniptest"
)

func Test_newOptions(t *testing.T) {
	Convey("Given I call newConfig", t, func() {

		cfg := newConfig()

		Convey("Then it should be correct", func() {
			So(cfg.enableLog, ShouldBeFalse)
			So(cfg.logfile, ShouldBeEmpty)
			So(cfg.readConsistency, ShouldEqual, manipulate.ReadConsistencyEventual)
			So(cfg.writeConsistency, ShouldEqual, manipulate.WriteConsistencyStrong)
			So(cfg.upstreamManipulator, ShouldBeNil)
			So(cfg.upstreamSubscriber, ShouldBeNil)
			So(cfg.defaultQueueDuration, ShouldEqual, time.Second)
			So(cfg.transactionQueue, ShouldHaveSameTypeAs, make(chan *Transaction, 1000))
			So(cfg.defaultPageSize, ShouldEqual, 10000)
		})
	})
}

func Test_Options(t *testing.T) {
	Convey("Given a memdb vortex memory", t, func() {

		cfg := newConfig()

		Convey("OptionUpstreamManipulator should work", func() {
			m := maniptest.NewTestManipulator()
			OptionUpstreamManipulator(m)(cfg)
			So(cfg.upstreamManipulator, ShouldResemble, m)
		})

		Convey("OptionUpstreamSubscriber should work", func() {
			s := maniptest.NewTestSubscriber()
			OptionUpstreamSubscriber(s)(cfg)
			So(cfg.upstreamSubscriber, ShouldResemble, s)
		})

		Convey("OptionTransactionLog should work", func() {
			OptionTransactionLog("somefile")(cfg)
			So(cfg.logfile, ShouldResemble, "somefile")
			So(cfg.enableLog, ShouldBeTrue)
		})

		Convey("OptionTransactionQueueLength it should work", func() {
			OptionTransactionQueueLength(13)(cfg)
			So(cfg.transactionQueue, ShouldNotBeNil)
			So(cap(cfg.transactionQueue), ShouldEqual, 13)
		})

		Convey("OptionDefaultConsistency should work", func() {
			OptionDefaultConsistency(manipulate.ReadConsistencyStrong, manipulate.WriteConsistencyNone)(cfg)
			So(cfg.readConsistency, ShouldEqual, manipulate.ReadConsistencyStrong)
			So(cfg.writeConsistency, ShouldEqual, manipulate.WriteConsistencyNone)
		})

		Convey("OptionDefaultConsistency with defaults should work", func() {
			OptionDefaultConsistency(manipulate.ReadConsistencyDefault, manipulate.WriteConsistencyDefault)(cfg)
			So(cfg.readConsistency, ShouldEqual, manipulate.ReadConsistencyEventual)
			So(cfg.writeConsistency, ShouldEqual, manipulate.WriteConsistencyStrong)
		})

		Convey("OptionTransactionQueueDuration with defaults should work", func() {
			OptionTransactionQueueDuration(time.Minute)(cfg)
			So(cfg.defaultQueueDuration, ShouldEqual, time.Minute)
		})

		Convey("OptionDefaultPageSize with defaults should work", func() {
			OptionDefaultPageSize(12)(cfg)
			So(cfg.defaultPageSize, ShouldEqual, 12)
		})

		Convey("OptionPrefetcher with defaults should work", func() {
			p := NewTestPrefetcher()
			OptionPrefetcher(p)(cfg)
			So(cfg.prefetcher, ShouldEqual, p)
		})
	})
}
