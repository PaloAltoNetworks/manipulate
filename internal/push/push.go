package push

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/wsc"
)

const (
	eventChSize  = 1024
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
	tokenRenewed            chan struct{}
	currentToken            string
	currentTokenLock        *sync.Mutex
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
		currentTokenLock:        &sync.Mutex{},
		unregisterTokenNotifier: unregisterTokenNotifier,
		registerTokenNotifier:   registerTokenNotifier,
		events:                  make(chan *elemental.Event, eventChSize),
		errors:                  make(chan error, errorChSize),
		status:                  make(chan manipulate.SubscriberStatus, statusChSize),
		filters:                 make(chan *elemental.PushFilter, filterChSize),
		config: wsc.Config{
			PongWait:     10 * time.Second,
			WriteWait:    10 * time.Second,
			PingPeriod:   5 * time.Second,
			ReadChanSize: 2048,
			TLSConfig:    tlsConfig,
		},
	}
}

func (s *subscription) UpdateFilter(filter *elemental.PushFilter) { s.filters <- filter }
func (s *subscription) Events() chan *elemental.Event             { return s.events }
func (s *subscription) Errors() chan error                        { return s.errors }
func (s *subscription) Status() chan manipulate.SubscriberStatus  { return s.status }

func (s *subscription) Start(ctx context.Context, filter *elemental.PushFilter) {

	if filter != nil {
		s.filters <- filter
	}

	s.registerTokenNotifier(s.id, s.setCurrentToken)

	go s.listen(ctx)
}

func (s *subscription) connect(ctx context.Context, initial bool) (err error) {

	var resp *http.Response
	var try int

	for {

		if s.conn, resp, err = wsc.Connect(ctx, makeURL(s.url, s.ns, s.getCurrentToken(), s.recursive), s.config); err == nil {

			if initial {
				s.publishStatus(manipulate.SubscriberStatusInitialConnection)
			} else {
				s.publishStatus(manipulate.SubscriberStatusReconnection)
			}

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

			case err = <-s.conn.Error():
				s.publishError(err)

			case err = <-s.conn.Done():

				if err != nil {
					s.publishError(err)
				}

				break processingLoop

			case <-s.tokenRenewed:

				s.publishStatus(manipulate.SubscriberStatusTokenRenewal)
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

	// notify the main loop if needed
	select {
	case s.tokenRenewed <- struct{}{}:
	default:
	}
}

func (s *subscription) getCurrentToken() string {

	s.currentTokenLock.Lock()
	t := s.currentToken
	s.currentTokenLock.Unlock()

	return t
}
