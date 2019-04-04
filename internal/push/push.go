package push

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/wsc"
)

const (
	eventChSize  = 2048
	errorChSize  = 64
	statusChSize = 8
	filterChSize = 2
)

type subscription struct {
	id                      string
	config                  wsc.Config
	conn                    wsc.Websocket
	errors                  chan error
	events                  chan *elemental.Event
	ns                      string
	recursive               bool
	status                  chan manipulate.SubscriberStatus
	url                     string
	filters                 chan *elemental.PushFilter
	currentFilter           *elemental.PushFilter
	currentFilterLock       sync.RWMutex
	tokenRenewed            chan struct{}
	currentToken            string
	currentTokenLock        sync.RWMutex
	unregisterTokenNotifier func(string)
	registerTokenNotifier   func(string, func(string))
}

// NewSubscriber creates a new Subscription.
func NewSubscriber(
	url string,
	ns string,
	token string,
	registerTokenNotifier func(string, func(string)),
	unregisterTokenNotifier func(string),
	tlsConfig *tls.Config,
	recursive bool,
) manipulate.Subscriber {

	return &subscription{
		id:                      uuid.Must(uuid.NewV4()).String(),
		url:                     url,
		ns:                      ns,
		recursive:               recursive,
		currentToken:            token,
		tokenRenewed:            make(chan struct{}),
		currentTokenLock:        sync.RWMutex{},
		unregisterTokenNotifier: unregisterTokenNotifier,
		registerTokenNotifier:   registerTokenNotifier,
		events:                  make(chan *elemental.Event, eventChSize),
		errors:                  make(chan error, errorChSize),
		status:                  make(chan manipulate.SubscriberStatus, statusChSize),
		filters:                 make(chan *elemental.PushFilter, filterChSize),
		currentFilterLock:       sync.RWMutex{},
		config: wsc.Config{
			PongWait:     10 * time.Second,
			WriteWait:    10 * time.Second,
			PingPeriod:   5 * time.Second,
			ReadChanSize: 2048,
			TLSConfig:    tlsConfig,
		},
	}
}

func (s *subscription) Events() chan *elemental.Event            { return s.events }
func (s *subscription) Errors() chan error                       { return s.errors }
func (s *subscription) Status() chan manipulate.SubscriberStatus { return s.status }

func (s *subscription) Start(ctx context.Context, filter *elemental.PushFilter) {

	if filter != nil {
		s.setCurrentFilter(filter)
	}

	s.registerTokenNotifier(s.id, s.setCurrentToken)

	go s.listen(ctx)
}

func (s *subscription) UpdateFilter(filter *elemental.PushFilter) {

	s.setCurrentFilter(filter)

	select {
	case s.filters <- filter:
	default:
	}
}

func (s *subscription) connect(ctx context.Context, initial bool) (err error) {

	var resp *http.Response
	var try int

	for {

		if s.conn != nil {
			s.conn.Close(0)
		}

		if resp != nil {
			resp.Body.Close() // nolint
		}

		if s.conn, resp, err = wsc.Connect(ctx, makeURL(s.url, s.ns, s.getCurrentToken(), s.recursive), s.config); err == nil {

			if initial {
				s.publishStatus(manipulate.SubscriberStatusInitialConnection)
			} else {
				s.publishStatus(manipulate.SubscriberStatusReconnection)
			}

			resp.Body.Close() // nolint

			try = 0
			return nil
		}

		if initial {
			s.publishStatus(manipulate.SubscriberStatusInitialConnectionFailure)
		} else {
			s.publishStatus(manipulate.SubscriberStatusReconnectionFailure)
		}

		if resp == nil {
			s.errors <- err
		} else if resp.StatusCode != http.StatusSwitchingProtocols {
			s.errors <- decodeErrors(resp.Body)
		}

		select {
		case <-time.After(nextBackoff(try)):
		case <-ctx.Done():
			s.publishStatus(manipulate.SubscriberStatusFinalDisconnection)
		}
		try++
	}
}

func (s *subscription) listen(ctx context.Context) {

	var err error
	var isReconnection bool
	var filterData []byte

	for {

		if err = s.connect(ctx, !isReconnection); err != nil {
			s.publishError(err)
			return
		}
		// If we have a current filter, we send it right away
		if f := s.getCurrentFilter(); f != nil {
			select {
			case s.filters <- f:
			default:
			}
		}

		isReconnection = true

	processingLoop:
		for {

			select {

			case filter := <-s.filters:

				filterData, err = json.Marshal(filter)
				if err != nil {
					s.publishError(err)
					continue
				}

				s.conn.Write(filterData)

			case data := <-s.conn.Read():

				event := &elemental.Event{}
				if err = json.Unmarshal(data, event); err != nil {
					s.publishError(err)
					continue
				}

				s.publishEvent(event)

			case <-s.tokenRenewed:
				s.publishStatus(manipulate.SubscriberStatusTokenRenewal)

			case err = <-s.conn.Error():
				s.publishError(err)

			case err = <-s.conn.Done():

				if err != nil {
					s.publishError(err)
				}

				break processingLoop

			case <-ctx.Done():

				s.unregisterTokenNotifier(s.id)
				s.conn.Close(websocket.CloseGoingAway)
				s.publishStatus(manipulate.SubscriberStatusFinalDisconnection)
				return
			}
		}

		s.publishStatus(manipulate.SubscriberStatusDisconnection)
	}
}

func (s *subscription) publishError(err error) {
	select {
	case s.errors <- err:
	default:
	}
}

func (s *subscription) publishEvent(evt *elemental.Event) {
	select {
	case s.events <- evt:
	default:
		s.publishError(fmt.Errorf("unable to forward event: channel full"))
	}
}

func (s *subscription) publishStatus(st manipulate.SubscriberStatus) {
	select {
	case s.status <- st:
	default:
	}
}

func (s *subscription) setCurrentToken(t string) {

	s.currentTokenLock.Lock()
	s.currentToken = t
	s.currentTokenLock.Unlock()

	// We get the current filter
	filter := s.getCurrentFilter()
	if filter == nil {
		// if it's nil, we create an empty one.
		filter = &elemental.PushFilter{}
	}
	filter.Params = url.Values{"token": []string{t}}

	s.UpdateFilter(filter)

	// notify the main loop if needed
	select {
	case s.tokenRenewed <- struct{}{}:
	default:
	}
}

func (s *subscription) getCurrentToken() string {

	s.currentTokenLock.RLock()
	t := s.currentToken
	s.currentTokenLock.RUnlock()

	return t
}

func (s *subscription) setCurrentFilter(f *elemental.PushFilter) {

	s.currentFilterLock.Lock()
	s.currentFilter = f
	s.currentFilterLock.Unlock()
}

func (s *subscription) getCurrentFilter() *elemental.PushFilter {

	s.currentFilterLock.RLock()
	defer s.currentFilterLock.RUnlock()

	return s.currentFilter
}
