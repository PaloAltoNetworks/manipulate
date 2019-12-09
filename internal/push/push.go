// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package push

import (
	"context"
	"crypto/tls"
	"fmt"
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
	filters                 chan *elemental.PushConfig
	currentFilter           *elemental.PushConfig
	currentFilterLock       sync.RWMutex
	currentToken            string
	currentTokenLock        sync.RWMutex
	unregisterTokenNotifier func(string)
	registerTokenNotifier   func(string, func(string))
	readEncoding            elemental.EncodingType
	writeEncoding           elemental.EncodingType
}

// NewSubscriber creates a new Subscription.
func NewSubscriber(
	url string,
	ns string,
	token string,
	registerTokenNotifier func(string, func(string)),
	unregisterTokenNotifier func(string),
	tlsConfig *tls.Config,
	headers http.Header,
	recursive bool,
) manipulate.Subscriber {

	readEncoding, writeEncoding, err := elemental.EncodingFromHeaders(headers)
	if err != nil {
		panic(err)
	}

	return &subscription{
		id:                      uuid.Must(uuid.NewV4()).String(),
		url:                     url,
		ns:                      ns,
		recursive:               recursive,
		currentToken:            token,
		currentTokenLock:        sync.RWMutex{},
		unregisterTokenNotifier: unregisterTokenNotifier,
		registerTokenNotifier:   registerTokenNotifier,
		events:                  make(chan *elemental.Event, eventChSize),
		errors:                  make(chan error, errorChSize),
		status:                  make(chan manipulate.SubscriberStatus, statusChSize),
		filters:                 make(chan *elemental.PushConfig, filterChSize),
		currentFilterLock:       sync.RWMutex{},
		readEncoding:            readEncoding,
		writeEncoding:           writeEncoding,
		config: wsc.Config{
			PongWait:     10 * time.Second,
			WriteWait:    10 * time.Second,
			PingPeriod:   5 * time.Second,
			ReadChanSize: 2048,
			TLSConfig:    tlsConfig,
			Headers:      headers,
		},
	}
}

func (s *subscription) Events() chan *elemental.Event            { return s.events }
func (s *subscription) Errors() chan error                       { return s.errors }
func (s *subscription) Status() chan manipulate.SubscriberStatus { return s.status }

func (s *subscription) Start(ctx context.Context, filter *elemental.PushConfig) {

	if filter != nil {
		s.setCurrentFilter(filter)
	}

	s.registerTokenNotifier(s.id, s.setCurrentToken)

	go s.listen(ctx)
}

func (s *subscription) UpdateFilter(filter *elemental.PushConfig) {

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
			_ = resp.Body.Close() // nolint
		}

		if s.conn, resp, err = wsc.Connect(ctx, makeURL(s.url, s.ns, s.getCurrentToken(), s.recursive), s.config); err == nil {

			if initial {
				s.publishStatus(manipulate.SubscriberStatusInitialConnection)
			} else {
				s.publishStatus(manipulate.SubscriberStatusReconnection)
			}

			_ = resp.Body.Close() // nolint

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

				filterData, err = elemental.Encode(s.writeEncoding, filter)
				if err != nil {
					s.publishError(err)
					continue
				}

				s.conn.Write(filterData)

			case data := <-s.conn.Read():

				event := &elemental.Event{}
				if err = elemental.Decode(s.readEncoding, data, event); err != nil {
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
		filter = elemental.NewPushConfig()
	}
	filter.SetParameter("token", t)

	s.UpdateFilter(filter)
	s.publishStatus(manipulate.SubscriberStatusTokenRenewal)
}

func (s *subscription) getCurrentToken() string {

	s.currentTokenLock.RLock()
	t := s.currentToken
	s.currentTokenLock.RUnlock()

	return t
}

func (s *subscription) setCurrentFilter(f *elemental.PushConfig) {

	s.currentFilterLock.Lock()
	s.currentFilter = f
	s.currentFilterLock.Unlock()
}

func (s *subscription) getCurrentFilter() *elemental.PushConfig {

	s.currentFilterLock.RLock()
	defer s.currentFilterLock.RUnlock()

	return s.currentFilter
}
