package push

import (
	"crypto/tls"
	"errors"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/wsutils"
	"github.com/gorilla/websocket"
)

const (
	eventChSize  = 1024
	errorChSize  = 64
	statusChSize = 8
)

type subscription struct {
	events       chan *elemental.Event
	errors       chan error
	status       chan manipulate.SubscriberStatus
	conn         *websocket.Conn
	stoppedLock  *sync.Mutex
	stopped      bool
	endpoint     string
	tlsConfig    *tls.Config
	maxConnRetry int
	filter       *elemental.PushFilter
	filterLock   *sync.Mutex
}

// NewSubscriber creates a new Subscription.
func NewSubscriber(
	url string,
	ns string,
	token string,
	tlsConfig *tls.Config,
	filter *elemental.PushFilter,
	recursive bool,
	maxConnRetry int,
) (manipulate.Subscriber, error) {

	s := &subscription{
		endpoint:     wsutils.MakeURL(url, "events", ns, token, recursive),
		maxConnRetry: maxConnRetry,
		tlsConfig:    tlsConfig,
		filter:       filter,
		stoppedLock:  &sync.Mutex{},
		filterLock:   &sync.Mutex{},
		events:       make(chan *elemental.Event, eventChSize),
		errors:       make(chan error, errorChSize),
		status:       make(chan manipulate.SubscriberStatus, statusChSize),
	}

	if err := s.connect(true); err != nil {
		return nil, err
	}

	select {
	case s.status <- manipulate.SubscriberStatusInitialConnection:
	default:
	}

	go s.listen()

	return s, nil
}

// Unsubscribe stop the subscription. After this, the Subscription must not be reused.
func (s *subscription) Unsubscribe() error {

	s.stoppedLock.Lock()
	s.stopped = true
	s.stoppedLock.Unlock()

	if s.conn == nil {
		return nil
	}

	return s.conn.Close()
}

// UpdateFilter updates the desired filter.
func (s *subscription) UpdateFilter(filter *elemental.PushFilter) error {

	s.filterLock.Lock()
	s.filter = filter
	s.filterLock.Unlock()

	return errors.New("UpdateFilter is not functional yet")
	// return s.conn.WriteJSON(filter) // this causes concurrent writes
}

// Events returns the event channel.
func (s *subscription) Events() chan *elemental.Event {
	return s.events
}

// Errors returns the error channel.
func (s *subscription) Errors() chan error {
	return s.errors
}

// Status returns the status channel.
func (s *subscription) Status() chan manipulate.SubscriberStatus {
	return s.status
}

func (s *subscription) currentFilter() *elemental.PushFilter {

	s.filterLock.Lock()
	defer s.filterLock.Unlock()

	return s.filter
}

func (s *subscription) connect(initial bool) (err error) {

	try := 0

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		close(c)
	}()

	for {
		s.conn, err = wsutils.Dial(
			s.endpoint,
			s.tlsConfig,
		)

		if err == nil {
			return nil
		}

		if initial && !manipulate.IsCannotCommunicateError(err) || s.isStopped() {
			return err
		}

		try++
		if s.maxConnRetry != -1 && try >= s.maxConnRetry {
			return err
		}

		select {
		case <-time.After(3 * time.Second):
		case <-c:
			return manipulate.NewErrDisconnected("Disconnected per signal")
		}

	}
}

func (s *subscription) listen() {

	var isReconnection bool
	var isDisconnection bool

	defer func() {
		s.status <- manipulate.SubscriberStatusFinalDisconnection
	}()

	// Connection loop.
	for {

		// If the subscriber is stopped, we return.
		if s.isStopped() {
			return
		}

		// If this is a disconnection, we publish the status event.
		if isDisconnection {
			select {
			case s.status <- manipulate.SubscriberStatusDisconnection:
			default:
			}
		}

		// If we have been disconnected, we try to reconnect.
		if isReconnection {
			if err := s.connect(false); err != nil {
				select {
				case s.errors <- err:
				default:
				}
				return
			}
		}

		// If we have a filter we install it.
		if filter := s.currentFilter(); filter != nil {
			if err := s.conn.WriteJSON(filter); err != nil {
				select {
				case s.errors <- err:
				default:
				}
				return
			}
		}

		// If this is a reconnection we publish the reconnection event
		if isReconnection {
			select {
			case s.status <- manipulate.SubscriberStatusReconnection:
			default:
			}
		}

		isReconnection = true
		isDisconnection = true

		// Read loop.
		for {

			event := &elemental.Event{}

			if err := s.conn.ReadJSON(event); err != nil {

				// If there is an error, but the user called Unsubscribe we simply return.
				if s.isStopped() {
					return
				}

				// We publish the error
				select {
				case s.errors <- err:
				default:
				}

				// We break the loop to try to reconnect.
				break
			}

			// Otherwise we publish the event and  we continue the read loop.
			select {
			case s.events <- event:
			default:
			}
		}
	}
}

func (s *subscription) isStopped() bool {

	s.stoppedLock.Lock()
	stopped := s.stopped
	s.stoppedLock.Unlock()

	return stopped
}
