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

package manipvortex

import (
	"testing"
	"time"

	"go.aporeto.io/manipulate"
	"golang.org/x/time/rate"

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
			So(cfg.upstreamReconciler, ShouldBeNil)
			So(cfg.downstreamReconciler, ShouldBeNil)
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

		Convey("OptionUpstreamReconciler with defaults should work", func() {
			r := NewTestReconciler()
			OptionUpstreamReconciler(r)(cfg)
			So(cfg.upstreamReconciler, ShouldEqual, r)
		})

		Convey("OptionDownstreamReconciler with defaults should work", func() {
			r := NewTestReconciler()
			OptionDownstreamReconciler(r)(cfg)
			So(cfg.downstreamReconciler, ShouldEqual, r)
		})

		Convey("OptRateLimiting should work", func() {
			OptionRateLimiting(2, 5)(cfg)
			So(cfg.rateLimiter, ShouldResemble, rate.NewLimiter(2, 5))
		})

		Convey("OptRateLimiting with defaults should work", func() {
			OptionRateLimiting(0, 0)(cfg)
			So(cfg.rateLimiter, ShouldResemble, rate.NewLimiter(3, 6))
		})
	})
}
