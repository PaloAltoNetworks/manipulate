package memdbvortex

import (
	"context"
	"fmt"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

// Subscriber is the memdb vortex subscriber implementation.
type Subscriber struct {
	v                       *MemDBVortex
	subscriberErrorChannel  chan error
	subscriberEventChannel  chan *elemental.Event
	subscriberStatusChannel chan manipulate.SubscriberStatus
	filter                  *elemental.PushFilter
}

// NewSubscriber creates a new subscriber.
func NewSubscriber(m manipulate.Manipulator, queueSize int) (manipulate.Subscriber, error) {

	v, ok := m.(*MemDBVortex)
	if !ok {
		return nil, fmt.Errorf("NewSubscriber only works with MemDBVortex manipulator")
	}

	if !v.hasBackendSubscriber() {
		return nil, fmt.Errorf("Vortex has no upstream subscriber - local subscriptions not supported")
	}

	s := &Subscriber{
		v:                       v,
		subscriberErrorChannel:  make(chan error, queueSize),
		subscriberEventChannel:  make(chan *elemental.Event, queueSize),
		subscriberStatusChannel: make(chan manipulate.SubscriberStatus, queueSize),
	}

	v.registerSubscriber(s)

	return s, nil
}

// Start starts the subscriber.
func (s *Subscriber) Start(ctx context.Context, e *elemental.PushFilter) {
	s.filter = e
	s.v.updateFilter()
}

// UpdateFilter updates the current filter.
func (s *Subscriber) UpdateFilter(e *elemental.PushFilter) {
	s.filter = e
	s.v.updateFilter()
}

// Events returns the events channel.
func (s *Subscriber) Events() chan *elemental.Event { return s.subscriberEventChannel }

// Errors returns the errors channel.
func (s *Subscriber) Errors() chan error { return s.subscriberErrorChannel }

// Status returns the status channel.
func (s *Subscriber) Status() chan manipulate.SubscriberStatus {
	return s.subscriberStatusChannel
}
