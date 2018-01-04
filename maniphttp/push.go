package maniphttp

import (
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/push"
)

// NewSubscriber returns a new subscription.
func NewSubscriber(manipulator manipulate.Manipulator, filter *elemental.PushFilter, recursive bool, maxConnTry int) (manipulate.Subscriber, error) {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You must pass a HTTP manipulator to maniphttp.NewSubscriper")
	}

	return push.NewSubscriber(m.url, m.namespace, m.currentPassword(), m.tlsConfig, recursive, maxConnTry)
}
