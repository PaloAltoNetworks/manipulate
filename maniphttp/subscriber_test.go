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

package maniphttp

import (
	"crypto/tls"
	"testing"

	// nolint:revive // Allow dot imports for readability in tests
	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/manipulate/maniptest"
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
				So(cfg.supportErrorEvents, ShouldBeFalse)
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

	Convey("SubscriberSendCredentialsAsCookie should work", t, func() {
		cfg := newSubscribeConfig(m)
		SubscriberSendCredentialsAsCookie("creds")(&cfg)
		So(cfg.credentialCookieKey, ShouldEqual, "creds")
	})

	Convey("SubscriberOptionSupportErrorEvents should work", t, func() {
		cfg := newSubscribeConfig(m)
		SubscriberOptionSupportErrorEvents()(&cfg)
		So(cfg.supportErrorEvents, ShouldBeTrue)
	})
}

func TestNewSubscriber(t *testing.T) {

	m := &httpManipulator{
		url:       "https://toto.com",
		namespace: "mns",
		tlsConfig: &tls.Config{},
	}

	Convey("Creating a new subscriber should work", t, func() {

		out := NewSubscriber(m, SubscriberOptionEndpoint("/"))

		So(out, ShouldNotBeNil)
	})

	Convey("Creating a new subscriber with a nil option should panic", t, func() {

		So(func() { NewSubscriber(m, nil) }, ShouldPanicWith, "nil passed as subscriber option")
	})

	Convey("Creating a new subscriber with a non http manipulator should panic", t, func() {

		So(func() { NewSubscriber(maniptest.NewTestManipulator(), nil) }, ShouldPanicWith, "You must pass a HTTP manipulator to maniphttp.NewSubscriber or maniphttp.NewSubscriberWithEndpoint")
	})

}
