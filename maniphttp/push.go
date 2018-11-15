package maniphttp

import (
	"fmt"
	"strings"

	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/push"
)

// NewSubscriber returns a new subscription.
func NewSubscriber(manipulator manipulate.Manipulator, recursive bool) manipulate.Subscriber {

	return NewSubscriberWithEndpoint(manipulator, "/events", recursive)
}

// NewSubscriberWithEndpoint returns a new subscription connecting to specific endpoint.
func NewSubscriberWithEndpoint(manipulator manipulate.Manipulator, endpoint string, recursive bool) manipulate.Subscriber {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You must pass a HTTP manipulator to maniphttp.NewSubscriber or maniphttp.NewSubscriberWithEndpoint")
	}

	return push.NewSubscriber(
		fmt.Sprintf("%s/%s", m.url, strings.TrimLeft(endpoint, "/")),
		m.namespace,
		m.currentPassword(),
		m.registerRenewNotifier,
		m.unregisterRenewNotifier,
		m.tlsConfig,
		recursive,
	)
}
