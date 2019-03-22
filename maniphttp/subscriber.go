package maniphttp

import (
	"crypto/tls"
	"fmt"
	"strings"

	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/push"
)

type subscribeConfig struct {
	namespace string
	endpoint  string
	recursive bool
	tlsConfig *tls.Config
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
		cfg.recursive,
	)
}

// NewSubscriberWithEndpoint returns a new subscription connecting to specific endpoint.
func NewSubscriberWithEndpoint(manipulator manipulate.Manipulator, endpoint string, recursive bool) manipulate.Subscriber {

	return NewSubscriber(manipulator, SubscriberOptionRecursive(recursive), SubscriberOptionEndpoint(endpoint))
}
