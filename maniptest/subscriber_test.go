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

package maniptest

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

func TestTestSubscriber_MockStart(t *testing.T) {

	Convey("Given I have TestSubscriber", t, func() {

		m := NewTestSubscriber()

		Convey("When I call Start with a panic mock", func() {
			m.MockStart(t, func(ctx context.Context, filter *elemental.PushConfig) {
				panic("test")
			})

			So(func() { m.Start(context.Background(), nil) }, ShouldPanic)
		})
	})
}

func TestTestSubscriber_MockUpdateFilter(t *testing.T) {
	Convey("Given I have TestSubscriber", t, func() {

		m := NewTestSubscriber()

		Convey("When I call UpdateFilter with a panic mock", func() {
			m.MockUpdateFilter(t, func(filter *elemental.PushConfig) {
				panic("test")
			})

			So(func() { m.UpdateFilter(nil) }, ShouldPanic)
		})
	})
}

func TestTestSubscriber_Events(t *testing.T) {

	Convey("Given a TestSubscriber", t, func() {

		m := NewTestSubscriber()

		Convey("When I call Events with no mock, it should return nil", func() {
			e := m.Events()
			So(e, ShouldBeNil)
		})

		Convey("When I call Events with an Event channel mock", func() {
			eventChannel := make(chan *elemental.Event)

			m.MockEvents(t, func() chan *elemental.Event {
				return eventChannel
			})

			e := m.Events()
			So(e, ShouldNotBeNil)
			So(e, ShouldResemble, eventChannel)
		})
	})
}

func TestTestSubscriber_Errors(t *testing.T) {

	Convey("Given a TestSubscriber", t, func() {

		m := NewTestSubscriber()

		Convey("When I call Errors with no mock, it should return nil", func() {
			e := m.Errors()
			So(e, ShouldBeNil)
		})

		Convey("When I call Events with an error channel mock", func() {
			errorChannel := make(chan error)

			m.MockErrors(t, func() chan error {
				return errorChannel
			})

			e := m.Errors()
			So(e, ShouldNotBeNil)
			So(e, ShouldResemble, errorChannel)
		})
	})
}

func TestTestSubscriber_Status(t *testing.T) {

	Convey("Given a TestSubscriber", t, func() {

		m := NewTestSubscriber()

		Convey("When I call Status with no mock, it should return nil", func() {
			e := m.Status()
			So(e, ShouldBeNil)
		})

		Convey("When I call Status with an status channel mock", func() {
			statusChannel := make(chan manipulate.SubscriberStatus)

			m.MockStatus(t, func() chan manipulate.SubscriberStatus {
				return statusChannel
			})

			e := m.Status()
			So(e, ShouldNotBeNil)
			So(e, ShouldResemble, statusChannel)
		})

	})
}
