package push

import (
	"crypto/tls"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/wsutils"
	"github.com/gorilla/websocket"
)

type subscription struct {
	events      chan *elemental.Event
	errors      chan error
	stopCh      chan struct{}
	conn        *websocket.Conn
	stoppedLock *sync.Mutex
	stopped     bool
}

// NewSubscriber creates a new Subscription.
// You should not use this directly, and use either maniphttp.NewSubscription or manipwebsocket.NewSubscription.
func NewSubscriber(url string, ns string, token string, tlsConfig *tls.Config, recursive bool) (manipulate.Subscriber, error) {

	u := wsutils.MakeURL(url, "events", ns, token, recursive)
	conn, err := wsutils.Dial(u, tlsConfig)
	if err != nil {
		return nil, err
	}

	s := &subscription{
		events:      make(chan *elemental.Event),
		errors:      make(chan error),
		conn:        conn,
		stoppedLock: &sync.Mutex{},
	}

	s.listen()

	return s, nil
}

// UpdateFilter updates the desired filter.
func (s *subscription) UpdateFilter(filter *elemental.PushFilter) error {
	return s.conn.WriteJSON(filter)
}

// Unsubscribe stop the subscription. After this, the Subscription must not be reused.
func (s *subscription) Unsubscribe() error {

	s.stoppedLock.Lock()
	s.stopped = true
	s.stoppedLock.Unlock()

	return s.conn.Close()
}

// Events returns the event channel.
func (s *subscription) Events() chan *elemental.Event {
	return s.events
}

// Errors returns the error channel.
func (s *subscription) Errors() chan error {
	return s.errors
}

func (s *subscription) listen() {

	go func() {
		for {
			event := &elemental.Event{}

			if err := s.conn.ReadJSON(event); err != nil {

				s.stoppedLock.Lock()
				stopped := s.stopped
				s.stoppedLock.Unlock()

				if stopped {
					return
				}

				s.errors <- err
				continue
			}

			s.events <- event
		}
	}()
}
