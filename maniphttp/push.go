package maniphttp

import (
	"fmt"
	"strings"

	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/push"
)

// NewSubscriber returns a new subscription.
func NewSubscriber(manipulator manipulate.Manipulator, recursive bool) manipulate.Subscriber {

	return NewSubscriberWithEndpoint(manipulator, "/highwind/events", recursive)
}

// NewSubscriberWithEndpoint returns a new subscription connecting to specific endpoint.
func NewSubscriberWithEndpoint(manipulator manipulate.Manipulator, endpoint string, recursive bool) manipulate.Subscriber {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You must pass a HTTP manipulator to maniphttp.NewSubscriber or maniphttp.NewSubscriberWithEndpoint")
	}

	return push.NewSubscriber(fmt.Sprintf("%s/%s", m.url, strings.TrimLeft(endpoint, "/")), m.namespace, m.currentPassword, m.tlsConfig, recursive)
}
