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
	"fmt"
	"net/http"
	"strings"

	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/push"
)

type subscribeConfig struct {
	namespace           string
	credentialCookieKey string
	endpoint            string
	supportErrorEvents  bool
	recursive           bool
	tlsConfig           *tls.Config
}

func newSubscribeConfig(m *httpManipulator) subscribeConfig {
	return subscribeConfig{
		endpoint:  "events",
		namespace: m.namespace,
		tlsConfig: m.tlsConfig,
	}
}

// SubscriberOption represents option to NewSubscriber.
type SubscriberOption func(*subscribeConfig)

// SubscriberOptionRecursive makes the subscriber to listen
// to events in current namespace and all children.
func SubscriberOptionRecursive(recursive bool) SubscriberOption {
	return func(cfg *subscribeConfig) {
		cfg.recursive = recursive
	}
}

// SubscriberOptionNamespace sets the namespace from where the subscription
// should start.
// By default it is the same as the manipulator.
func SubscriberOptionNamespace(namespace string) SubscriberOption {
	return func(cfg *subscribeConfig) {
		cfg.namespace = namespace
	}
}

// SubscriberOptionEndpoint sets the endpint to connect to.
// By default it is /events.
func SubscriberOptionEndpoint(endpoint string) SubscriberOption {
	return func(cfg *subscribeConfig) {
		cfg.endpoint = strings.TrimRight(strings.TrimLeft(endpoint, "/"), "/")
	}
}

// SubscriberSendCredentialsAsCookie makes the subscriber send the
// crendentials as cookie using the provided key
func SubscriberSendCredentialsAsCookie(key string) SubscriberOption {
	return func(cfg *subscribeConfig) {
		cfg.credentialCookieKey = key
	}
}

// SubscriberOptionSupportErrorEvents will result in connecting to the socket server by declaring that you are capable of
// handling error events.
func SubscriberOptionSupportErrorEvents() SubscriberOption {
	return func(cfg *subscribeConfig) {
		cfg.supportErrorEvents = true
	}
}

// NewSubscriber returns a new subscription.
func NewSubscriber(manipulator manipulate.Manipulator, options ...SubscriberOption) manipulate.Subscriber {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You must pass a HTTP manipulator to maniphttp.NewSubscriber or maniphttp.NewSubscriberWithEndpoint")
	}

	cfg := newSubscribeConfig(m)
	for _, opt := range options {
		if opt == nil {
			panic("nil passed as subscriber option")
		}
		opt(&cfg)
	}

	return push.NewSubscriber(
		fmt.Sprintf("%s/%s", m.url, cfg.endpoint),
		cfg.namespace,
		m.currentPassword(),
		m.registerRenewNotifier,
		m.unregisterRenewNotifier,
		cfg.tlsConfig,
		http.Header{
			"Content-Type": []string{string(m.encoding)},
			"Accept":       []string{string(m.encoding)},
		},
		cfg.supportErrorEvents,
		cfg.recursive,
		cfg.credentialCookieKey,
	)
}

// NewSubscriberWithEndpoint returns a new subscription connecting to specific endpoint.
func NewSubscriberWithEndpoint(manipulator manipulate.Manipulator, endpoint string, recursive bool) manipulate.Subscriber {

	return NewSubscriber(manipulator, SubscriberOptionRecursive(recursive), SubscriberOptionEndpoint(endpoint))
}
