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
	"context"
	"testing"

	"go.aporeto.io/manipulate"

	// nolint:revive // Allow dot imports for readability in tests
	. "github.com/smartystreets/goconvey/convey"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate/maniptest"
)

func Test_NewSubscriber(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	Convey("Given a valid manipulator with a backend", t, func() {
		m := maniptest.NewTestManipulator()
		s := maniptest.NewTestSubscriber()
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v, err := New(
			ctx,
			d,
			newIdentityProcessor(manipulate.ReadConsistencyEventual, manipulate.WriteConsistencyStrong),
			testmodel.Manager(),
			OptionUpstreamManipulator(m),
			OptionUpstreamSubscriber(s),
		)
		So(err, ShouldBeNil)

		Convey("When I request a new subscriber, it should be valid", func() {
			s, err := NewSubscriber(v, 100)
			So(err, ShouldBeNil)
			So(s, ShouldNotBeNil)
		})
	})

	Convey("Given an invalid manipulator, the method should error", t, func() {
		_, err := NewSubscriber(nil, 100)
		So(err, ShouldNotBeNil)
	})

	Convey("Given a valid manipulator with no upstream subscriber, it should error", t, func() {
		m := maniptest.NewTestManipulator()
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v, err := New(
			ctx,
			d,
			newIdentityProcessor(manipulate.ReadConsistencyEventual, manipulate.WriteConsistencyStrong),
			testmodel.Manager(),
			OptionUpstreamManipulator(m),
		)
		So(err, ShouldBeNil)

		Convey("When I request a new subscriber, it should be valid", func() {
			_, err := NewSubscriber(v, 100)
			So(err, ShouldNotBeNil)
		})
	})
}

func Test_SubscriberMethods(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	Convey("Given a valid manipulator with a backend subscriber", t, func() {
		m := maniptest.NewTestManipulator()
		us := maniptest.NewTestSubscriber()
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v, err := New(
			ctx,
			d,
			newIdentityProcessor(manipulate.ReadConsistencyEventual, manipulate.WriteConsistencyStrong),
			testmodel.Manager(),
			OptionUpstreamManipulator(m),
			OptionUpstreamSubscriber(us),
		)
		So(err, ShouldBeNil)
		So(v, ShouldNotBeNil)
		s, err := NewSubscriber(v, 100)
		So(err, ShouldBeNil)

		Convey("When I retrieve the events channel, it should not be nil", func() {
			ch := s.Events()
			So(ch, ShouldNotBeNil)
			So(cap(ch), ShouldEqual, 100)
		})

		Convey("When I retrieve the errors channel, it should not be nil", func() {
			ch := s.Errors()
			So(ch, ShouldNotBeNil)
			So(cap(ch), ShouldEqual, 100)
		})

		Convey("When I retrieve the status channel, it should not be nil", func() {
			ch := s.Status()
			So(ch, ShouldNotBeNil)
			So(cap(ch), ShouldEqual, 100)
		})
	})
}
