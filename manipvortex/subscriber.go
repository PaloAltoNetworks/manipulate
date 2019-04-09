package manipvortex

import (
	"context"
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

// vortexSubscriber is the memdb vortex subscriber implementation.
type vortexSubscriber struct {
	v                       *vortexManipulator
	subscriberErrorChannel  chan error
	subscriberEventChannel  chan *elemental.Event
	subscriberStatusChannel chan manipulate.SubscriberStatus
	filter                  *elemental.PushFilter

	sync.RWMutex
}

// NewSubscriber creates a new vortex subscriber.
func NewSubscriber(m manipulate.Manipulator, queueSize int) (manipulate.Subscriber, error) {

	v, ok := m.(*vortexManipulator)
	if !ok {
		return nil, fmt.Errorf("NewSubscriber only works with Vortex manipulator")
	}

	if !v.hasBackendSubscriber() {
		return nil, fmt.Errorf("vortex has no upstream subscriber: local subscriptions not supported")
	}

	s := &vortexSubscriber{
		v:                       v,
		subscriberErrorChannel:  make(chan error, queueSize),
		subscriberEventChannel:  make(chan *elemental.Event, queueSize),
		subscriberStatusChannel: make(chan manipulate.SubscriberStatus, queueSize),
	}

	v.registerSubscriber(s)

	return s, nil
}

// Start starts the subscriber.
func (s *vortexSubscriber) Start(ctx context.Context, e *elemental.PushFilter) {

	s.UpdateFilter(e)
}

// UpdateFilter updates the current filter.
func (s *vortexSubscriber) UpdateFilter(e *elemental.PushFilter) {

	if e == nil {
		return
	}

	s.Lock()
	s.filter = e
	s.Unlock()

	s.v.updateFilter()
}

// Events returns the events channel.
func (s *vortexSubscriber) Events() chan *elemental.Event { return s.subscriberEventChannel }

// Errors returns the errors channel.
func (s *vortexSubscriber) Errors() chan error { return s.subscriberErrorChannel }

// Status returns the status channel.
func (s *vortexSubscriber) Status() chan manipulate.SubscriberStatus {
	return s.subscriberStatusChannel
}
