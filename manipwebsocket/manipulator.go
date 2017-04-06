// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipwebsocket

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/tracing"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/websocket"

	midgard "github.com/aporeto-inc/midgard-lib/client"
)

type websocketManipulator struct {
	responsesChanRegistry     map[string]chan *elemental.Response
	responsesChanRegistryLock *sync.Mutex
	renewLock                 *sync.Mutex
	namespace                 string
	password                  string
	receiveAll                bool
	running                   bool
	runningLock               *sync.Mutex
	tlsConfig                 *tls.Config
	url                       string
	username                  string
	wsLock                    *sync.Mutex
	ws                        *websocket.Conn
}

// NewWebSocketManipulator returns a Manipulator backed by a websocket API.
func NewWebSocketManipulator(username, password, url, namespace string) (manipulate.EventManipulator, func(), error) {

	CAPool, err := x509.SystemCertPool()
	if err != nil {
		logrus.Error("Unable to load system root cert pool. tls fallback to unsecure.")
	}

	return NewWebSocketManipulatorWithRootCA(username, password, url, namespace, CAPool, true)
}

// NewWebSocketManipulatorWithRootCA returns a Manipulator backed by an ReST API using the given CAPool as root CA.
func NewWebSocketManipulatorWithRootCA(username, password, url, namespace string, rootCAPool *x509.CertPool, skipTLSVerify bool) (manipulate.EventManipulator, func(), error) {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: skipTLSVerify,
		RootCAs:            rootCAPool,
	}

	m := &websocketManipulator{
		username:                  username,
		password:                  password,
		url:                       url,
		namespace:                 namespace,
		tlsConfig:                 tlsConfig,
		responsesChanRegistry:     map[string]chan *elemental.Response{},
		responsesChanRegistryLock: &sync.Mutex{},
		renewLock:                 &sync.Mutex{},
		runningLock:               &sync.Mutex{},
		wsLock:                    &sync.Mutex{},
		running:                   true,
	}

	if err := m.connect(); err != nil {
		return nil, nil, err
	}

	go m.listen()

	return m, func() {
		m.wsLock.Lock()
		if m.ws != nil && m.ws.IsClientConn() {
			m.runningLock.Lock()
			m.running = false
			m.runningLock.Unlock()
			m.ws.Close() // nolint: errcheck
		}
		m.wsLock.Unlock()
	}, nil
}

// NewWebSocketManipulatorWithMidgardCertAuthentication returns a http backed manipulate.Manipulator
// using a certificates to authenticate against a Midgard server.
func NewWebSocketManipulatorWithMidgardCertAuthentication(url string, midgardurl string, rootCAPool *x509.CertPool, clientCAPool *x509.CertPool, certificates []tls.Certificate, namespace string, refreshInterval time.Duration, skipInsecure bool) (manipulate.EventManipulator, func(), error) {

	sp := opentracing.StartSpan("manipwebsocket.authenthication")
	defer sp.Finish()

	mclient := midgard.NewClientWithCAPool(midgardurl, rootCAPool, clientCAPool, skipInsecure)
	token, err := mclient.IssueFromCertificateWithValidity(certificates, 24*time.Hour, sp)
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return nil, nil, err
	}

	m, stop, err := NewWebSocketManipulatorWithRootCA("Bearer", token, url, namespace, rootCAPool, skipInsecure)
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return nil, nil, err
	}

	stopCh := make(chan bool)

	go m.(*websocketManipulator).renewMidgardToken(mclient, certificates, refreshInterval, stopCh)

	return m, func() { stop(); stopCh <- true }, err
}

func (s *websocketManipulator) RetrieveMany(context *manipulate.Context, dest elemental.ContentIdentifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, fmt.Sprintf("manipwebsocket.retrieve_many.%s", dest.ContentIdentity().Category), context)

	req := elemental.NewRequest()
	req.Namespace = s.namespace
	req.Operation = elemental.OperationRetrieveMany
	req.Identity = dest.ContentIdentity()
	req.Username = s.username
	req.Password = s.currentPassword()

	if err := populateRequestFromContext(req, context); err != nil {
		tracing.FinishTraceWithError(sp, err)
		return err
	}

	tracing.Inject(sp, req)

	resp, err := s.send(req)
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return err
	}

	if err := resp.Decode(dest); err != nil {
		tracing.FinishTraceWithError(sp, err)
		return manipulate.NewErrCannotUnmarshal(err.Error())
	}

	tracing.FinishTrace(sp)

	return nil
}

func (s *websocketManipulator) Retrieve(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, "manipwebsocket.retrieve", context)
	defer tracing.FinishTrace(sp)

	for _, object := range objects {

		subSp := tracing.StartTrace(sp, fmt.Sprintf("manipwebsocket.retrieve.object.%s", object.Identity().Name), context)
		tracing.SetTag(subSp, "manipwebsocket.retrieve.object.id", object.Identifier())

		req := elemental.NewRequest()
		req.Namespace = s.namespace
		req.Operation = elemental.OperationRetrieve
		req.Identity = object.Identity()
		req.Username = s.username
		req.Password = s.currentPassword()
		req.ObjectID = object.Identifier()

		if err := populateRequestFromContext(req, context); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if err := req.Encode(object); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotMarshal(err.Error())
		}

		tracing.Inject(subSp, req)

		resp, err := s.send(req)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if err := resp.Decode(&object); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotUnmarshal(err.Error())
		}

		tracing.FinishTrace(subSp)
	}

	return nil
}

func (s *websocketManipulator) Create(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, "manipwebsocket.create", context)
	defer tracing.FinishTrace(sp)

	for _, object := range objects {

		subSp := tracing.StartTrace(sp, fmt.Sprintf("manipwebsocket.create.object.%s", object.Identity().Name), context)
		tracing.SetTag(subSp, "manipwebsocket.create.object.id", object.Identifier())

		req := elemental.NewRequest()
		req.Namespace = s.namespace
		req.Operation = elemental.OperationCreate
		req.Identity = object.Identity()
		req.Username = s.username
		req.Password = s.currentPassword()

		if err := populateRequestFromContext(req, context); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if err := req.Encode(object); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotMarshal(err.Error())
		}

		tracing.Inject(subSp, req)

		resp, err := s.send(req)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if err := resp.Decode(&object); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotUnmarshal(err.Error())
		}

		tracing.FinishTrace(subSp)
	}

	return nil
}

func (s *websocketManipulator) Update(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, "manipwebsocket.update", context)
	defer tracing.FinishTrace(sp)

	for _, object := range objects {

		subSp := tracing.StartTrace(sp, fmt.Sprintf("manipwebsocket.update.object.%s", object.Identity().Name), context)
		tracing.SetTag(subSp, "manipwebsocket.update.object.id", object.Identifier())

		req := elemental.NewRequest()
		req.Namespace = s.namespace
		req.Operation = elemental.OperationUpdate
		req.Identity = object.Identity()
		req.Username = s.username
		req.Password = s.currentPassword()
		req.ObjectID = object.Identifier()

		if err := populateRequestFromContext(req, context); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if err := req.Encode(object); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotMarshal(err.Error())
		}

		tracing.Inject(subSp, req)

		resp, err := s.send(req)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if err := resp.Decode(&object); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotUnmarshal(err.Error())
		}

		tracing.FinishTrace(subSp)
	}

	return nil
}

func (s *websocketManipulator) Delete(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, "manipwebsocket.delete", context)
	defer tracing.FinishTrace(sp)

	for _, object := range objects {

		subSp := tracing.StartTrace(sp, fmt.Sprintf("manipwebsocket.delete.object.%s", object.Identity().Name), context)
		tracing.SetTag(subSp, "manipwebsocket.delete.object.id", object.Identifier())

		req := elemental.NewRequest()
		req.Namespace = s.namespace
		req.Operation = elemental.OperationDelete
		req.Identity = object.Identity()
		req.Username = s.username
		req.Password = s.currentPassword()
		req.ObjectID = object.Identifier()

		if err := populateRequestFromContext(req, context); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if err := req.Encode(object); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotMarshal(err.Error())
		}

		tracing.Inject(subSp, req)

		resp, err := s.send(req)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if err := resp.Decode(&object); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotUnmarshal(err.Error())
		}

		tracing.FinishTrace(subSp)
	}

	return nil
}

func (s *websocketManipulator) DeleteMany(context *manipulate.Context, identity elemental.Identity) error {

	return manipulate.NewErrNotImplemented("DeleteMany not implemented in manipwebsocket")
}

func (s *websocketManipulator) Count(context *manipulate.Context, identity elemental.Identity) (int, error) {

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, fmt.Sprintf("manipwebsocket.count.%s", identity.Category), context)

	req := elemental.NewRequest()
	req.Namespace = s.namespace
	req.Operation = elemental.OperationInfo
	req.Identity = identity
	req.Username = s.username
	req.Password = s.currentPassword()

	if err := populateRequestFromContext(req, context); err != nil {
		tracing.FinishTraceWithError(sp, err)
		return 0, err
	}

	tracing.Inject(sp, req)

	resp, err := s.send(req)
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return 0, err
	}

	tracing.FinishTrace(sp)

	return resp.Total, nil
}

func (s *websocketManipulator) Assign(context *manipulate.Context, assignation *elemental.Assignation) error {

	return manipulate.NewErrNotImplemented("Assign not implemented in websocket manipulator")
}

func (s *websocketManipulator) Increment(context *manipulate.Context, identity elemental.Identity, counter string, inc int) error {

	return manipulate.NewErrNotImplemented("Increment not implemented in websocket manipulator")
}

func (s *websocketManipulator) Subscribe(
	filter *elemental.PushFilter,
	allNamespaces bool,
	handler manipulate.EventHandler,
	recoHandler manipulate.RecoveryHandler,
) (manipulate.EventUnsubscriber, manipulate.EventFilterUpdater, error) {

	var ws *websocket.Conn
	var stopped bool

	lock := &sync.Mutex{}
	readyChan := make(chan bool)

	needsPublishDisconnectFunc := true

	disconnectFunc := func() error {
		lock.Lock()
		defer lock.Unlock()
		stopped = true
		return ws.Close()
	}

	eventFilterSetterFunc := func(filter *elemental.PushFilter) error {
		lock.Lock()
		defer lock.Unlock()
		return websocket.JSON.Send(ws, filter)
	}

	go func() {

		var needsReconnectionHandlerCall bool

		for {
			url := strings.Replace(s.url, "http://", "ws://", 1)
			url = strings.Replace(url, "https://", "wss://", 1)
			url = url + "/events?token=" + s.currentPassword()
			if s.namespace != "" {
				url += "&namespace=" + s.namespace
			}

			if allNamespaces {
				url = url + "&mode=all"
			}

			config, err := websocket.NewConfig(url, url)
			if err != nil {
				panic(err)
			}
			config.TlsConfig = s.tlsConfig

			lock.Lock()
			if stopped {
				lock.Unlock()
				return
			}

			ws, err = websocket.DialConfig(config)
			lock.Unlock()
			if err != nil {
				logrus.Warn("Could not connect to websocket. Retrying in 5s")
				<-time.After(5 * time.Second)
				continue
			}

			if filter != nil {
				if err := websocket.JSON.Send(ws, filter); err != nil {
					handler(nil, err)
					break
				}
			}

			if needsPublishDisconnectFunc {
				readyChan <- true
			}
			needsPublishDisconnectFunc = false

			if needsReconnectionHandlerCall && recoHandler != nil {
				needsReconnectionHandlerCall = false
				recoHandler()
			}

			for {
				event := &elemental.Event{}
				err := websocket.JSON.Receive(ws, event)

				lock.Lock()
				if stopped {
					lock.Unlock()
					break
				}
				lock.Unlock()

				if err != nil {
					handler(nil, err)
					needsReconnectionHandlerCall = true
					break
				}

				handler(event, nil)
			}
		}
	}()

	<-readyChan

	return disconnectFunc, eventFilterSetterFunc, nil
}

func (s *websocketManipulator) connect() error {

	s.unregisterAllResponseChannels()

	destURL := strings.Replace(s.url, "http://", "ws://", 1)
	destURL = strings.Replace(destURL, "https://", "wss://", 1)
	destURL = destURL + "/wsapi?token=" + s.currentPassword()
	if s.namespace != "" {
		destURL += "&namespace=" + url.QueryEscape(s.namespace)
	}

	if s.receiveAll {
		destURL = destURL + "&mode=all"
	}

	config, err := websocket.NewConfig(destURL, destURL)
	if err != nil {
		return err
	}

	config.TlsConfig = s.tlsConfig

	s.wsLock.Lock()
	s.ws, err = websocket.DialConfig(config)
	s.wsLock.Unlock()
	if err != nil {
		return manipulate.NewErrCannotCommunicate(err.Error())
	}

	response := elemental.NewResponse()
	if err := websocket.JSON.Receive(s.ws, &response); err != nil {
		return manipulate.NewErrCannotCommunicate(err.Error())
	}

	if response.StatusCode != http.StatusOK {
		return decodeErrors(response)
	}

	return nil
}

func (s *websocketManipulator) listen() {

	for {

		response := elemental.NewResponse()
		err := websocket.JSON.Receive(s.ws, &response)

		if err == nil {
			if ch := s.responseChannelForID(response.Request.RequestID); ch != nil {
				ch <- response
			}
			continue
		}

		s.runningLock.Lock()
		if !s.running {
			s.runningLock.Unlock()
			return
		}
		s.runningLock.Unlock()

		logrus.WithField("error", err).Warn("Websocket connection died. Reconnecting...")
		for {

			if err := s.connect(); err != nil {
				logrus.Warn("Websocket not available. Retrying in 5s...")
				<-time.After(5 * time.Second)
				continue
			}

			logrus.Info("Websocket connection restored.")
			break
		}
	}
}

func (s *websocketManipulator) send(request *elemental.Request) (*elemental.Response, error) {

	s.wsLock.Lock()
	if s.ws == nil {
		s.wsLock.Unlock()
		return nil, manipulate.NewErrCannotCommunicate("Websocket not initialized")
	}

	err := websocket.JSON.Send(s.ws, request)
	s.wsLock.Unlock()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"method":  request.Operation,
			"url":     s.url,
			"request": request.String(),
			"data":    string(request.Data),
			"error":   err.Error(),
		}).Debug("Unable to send the request.")
		return nil, manipulate.NewErrCannotCommunicate(err.Error())
	}

	ch := s.registerResponseChannel(request.RequestID)
	defer s.unregisterResponseChannel(request.RequestID)

	select {
	case response := <-ch:

		if response.StatusCode < 200 || response.StatusCode > 300 {
			return nil, decodeErrors(response)
		}

		return response, nil

	case <-time.After(30 * time.Second):
		return nil, manipulate.NewErrCannotCommunicate("Request timeout")
	}
}

func (s *websocketManipulator) registerResponseChannel(rid string) chan *elemental.Response {

	ch := make(chan *elemental.Response)
	s.responsesChanRegistryLock.Lock()
	s.responsesChanRegistry[rid] = ch
	s.responsesChanRegistryLock.Unlock()

	return ch
}

func (s *websocketManipulator) unregisterResponseChannel(rid string) {

	s.responsesChanRegistryLock.Lock()
	delete(s.responsesChanRegistry, rid)
	s.responsesChanRegistryLock.Unlock()
}

func (s *websocketManipulator) unregisterAllResponseChannels() {

	s.responsesChanRegistryLock.Lock()
	s.responsesChanRegistry = map[string]chan *elemental.Response{}
	s.responsesChanRegistryLock.Unlock()
}

func (s *websocketManipulator) responseChannelForID(rid string) chan *elemental.Response {

	s.responsesChanRegistryLock.Lock()
	defer s.responsesChanRegistryLock.Unlock()

	return s.responsesChanRegistry[rid]
}

func (s *websocketManipulator) setPassword(password string) {
	s.renewLock.Lock()
	s.password = password
	s.renewLock.Unlock()
}

func (s *websocketManipulator) currentPassword() string {
	s.renewLock.Lock()
	defer s.renewLock.Unlock()
	return s.password
}

func (s *websocketManipulator) renewMidgardToken(
	mclient *midgard.Client,
	certificates []tls.Certificate,
	interval time.Duration,
	stop chan bool,
) {

	for {
		nextRefresh := time.Now().Add(interval)

		select {
		case <-time.Tick(time.Minute):

			now := time.Now()
			if now.Before(nextRefresh) {
				continue
			}

			token, err := mclient.IssueFromCertificate(certificates)
			if err != nil {
				logrus.WithError(err).Error("Unable to renew token.")
				break
			}

			s.setPassword(token)

			logrus.Info("Midgard token refreshed")

		case <-stop:
			return
		}
	}
}
