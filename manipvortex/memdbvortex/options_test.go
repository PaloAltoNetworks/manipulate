package memdbvortex

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/manipulate/maniptest"
)

func Test_Options(t *testing.T) {
	Convey("Given a memdb vortext memory", t, func() {
		v := &MemDBVortex{}

		Convey("When I set the manipulator it should work", func() {
			m := maniptest.NewTestManipulator()
			OptionBackendManipulator(m)(v)
			So(v.m, ShouldResemble, m)
		})

		Convey("When I set the subscriber it should work", func() {
			s := maniptest.NewTestSubscriber()
			OptionBackendSubscriber(s)(v)
			So(v.s, ShouldResemble, s)
		})

		Convey("When I set the transaction log it should work", func() {
			OptionTransactionLog("somefile")(v)
			So(v.logfile, ShouldResemble, "somefile")
			So(v.enableLog, ShouldBeTrue)
		})

		Convey("When I set the transaction queue length it should work", func() {
			OptionTransactionQueueLength(13)(v)
			So(v.transactionQueue, ShouldNotBeNil)
			So(cap(v.transactionQueue), ShouldEqual, 13)
		})
	})
}
