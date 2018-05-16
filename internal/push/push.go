package push

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aporeto-inc/addedeffect/wsc"
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/gorilla/websocket"
)

const (
	eventChSize  = 1024
	errorChSize  = 64
	statusChSize = 8
	filterChSize = 2
)

type subscription struct {
	events       chan *elemental.Event
	errors       chan error
	status       chan manipulate.SubscriberStatus
	filters      chan *elemental.PushFilter
	conn         wsc.Websocket
	tokenManager manipulate.TokenManager
	token        string
	config       wsc.Config
	url          string
	ns           string
	recursive    bool
}

// NewSubscriberWithTokenManager creates a new Subscription using a manioulate.TokenManager.
// The Subscribtion will get a new token before each reconnection.
func NewSubscriberWithTokenManager(url string, ns string, tokenManager manipulate.TokenManager, tlsConfig *tls.Config, recursive bool) manipulate.Subscriber {

	return newSubscriber(url, ns, "", tokenManager, tlsConfig, recursive)
}

// NewSubscriberWithToken creates a new Subscription using a one time token.
// If the token expires, the subscriber will continue to work until disconnection.
// It is strongly advised to use NewSubscriberWithTokenManager for long running operation.
func NewSubscriberWithToken(url string, ns string, token string, tlsConfig *tls.Config, recursive bool) manipulate.Subscriber {

	return newSubscriber(url, ns, token, nil, tlsConfig, recursive)
}

func newSubscriber(
	url string,
	ns string,
	token string,
	tokenManager manipulate.TokenManager,
	tlsConfig *tls.Config,
	recursive bool,
) manipulate.Subscriber {

	config := wsc.Config{
		PongWait:   10 * time.Second,
		WriteWait:  10 * time.Second,
		PingPeriod: 5 * time.Second,
		TLSConfig:  tlsConfig,
	}

	s := &subscription{
		url:          url,
		ns:           ns,
		recursive:    recursive,
		token:        token,
		tokenManager: tokenManager,
		events:       make(chan *elemental.Event, eventChSize),
		errors:       make(chan error, errorChSize),
		status:       make(chan manipulate.SubscriberStatus, statusChSize),
		filters:      make(chan *elemental.PushFilter, filterChSize),
		config:       config,
	}

	return s
}

func (s *subscription) UpdateFilter(filter *elemental.PushFilter) { s.filters <- filter }
func (s *subscription) Events() chan *elemental.Event             { return s.events }
func (s *subscription) Errors() chan error                        { return s.errors }
func (s *subscription) Status() chan manipulate.SubscriberStatus  { return s.status }

func (s *subscription) Start(ctx context.Context, filter *elemental.PushFilter) {

	if filter != nil {
		s.filters <- filter
	}

	go s.listen(ctx)
}

func (s *subscription) connect(ctx context.Context, initial bool) (err error) {

	var resp *http.Response

	for {

		endpoint, err := s.computeEndpoint(ctx)
		if err != nil {
			return err
		}

		if s.conn, resp, err = wsc.Connect(ctx, endpoint, s.config); err == nil {

			if initial {
				s.publishStatus(manipulate.SubscriberStatusInitialConnection)
			} else {
				s.publishStatus(manipulate.SubscriberStatusReconnection)
			}

			return nil
		}

		if !isCommError(resp) {
			return decodeErrors(resp.Body)
		}

		if initial {
			s.publishStatus(manipulate.SubscriberStatusInitialConnectionFailure)
		} else {
			s.publishStatus(manipulate.SubscriberStatusReconnectionFailure)
		}

		select {
		case <-time.After(3 * time.Second):
		case <-ctx.Done():
			return manipulate.NewErrDisconnected("Disconnected per signal")
		}
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

			case err = <-s.conn.Done():

				if err != nil {
					s.publishError(err)
				}

				break processingLoop

			case <-ctx.Done():

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

func (s *subscription) computeEndpoint(ctx context.Context) (url string, err error) {

	token := s.token

	if s.tokenManager != nil {
		token, err = s.tokenManager.Issue(ctx)
		if err != nil {
			return "", err
		}
	}

	return makeURL(s.url, "events", s.ns, token, s.recursive), nil
}
