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

package maniphttp

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/internal/idempotency"
	"go.aporeto.io/manipulate/internal/snip"
	"go.aporeto.io/manipulate/internal/tracing"
)

type httpManipulator struct {
	username           string
	password           string
	url                string
	namespace          string
	renewLock          *sync.RWMutex
	renewNotifiers     map[string]func(string)
	renewNotifiersLock *sync.RWMutex
	disableAutoRetry   bool

	// optionnable
	ctx           context.Context
	client        *http.Client
	tlsConfig     *tls.Config
	tokenManager  manipulate.TokenManager
	globalHeaders http.Header
	transport     *http.Transport
	encoding      elemental.EncodingType
}

// New returns a maniphttp.Manipulator configured according to the given suite of Option.
func New(ctx context.Context, url string, options ...Option) (manipulate.Manipulator, error) {

	if url == "" {
		panic("empty url")
	}

	// initialize solid varialbles.
	m := &httpManipulator{
		renewLock:          &sync.RWMutex{},
		renewNotifiersLock: &sync.RWMutex{},
		renewNotifiers:     map[string]func(string){},
		ctx:                ctx,
		url:                url,
		encoding:           elemental.EncodingTypeJSON,
	}

	// Apply the options.
	for _, opt := range options {
		opt(m)
	}

	if m.client == nil {

		m.client = getDefaultClient()

		if m.transport == nil {

			m.transport, m.url = getDefaultTransport(url)

			if m.tlsConfig == nil {
				m.tlsConfig = getDefaultTLSConfig()
			}

			m.transport.TLSClientConfig = m.tlsConfig
		}

		m.client.Transport = m.transport
	}

	if m.tokenManager != nil {

		ictx, cancel := context.WithTimeout(m.ctx, 30*time.Second)
		defer cancel()

		token, err := m.tokenManager.Issue(ictx)
		if err != nil {
			return nil, err
		}

		m.username = "Bearer"
		m.password = token

		go func() {
			tokenCh := make(chan string)
			go m.tokenManager.Run(m.ctx, tokenCh)

			for {
				select {
				case t := <-tokenCh:
					m.setPassword(t)
				case <-m.ctx.Done():
					return
				}
			}
		}()
	}

	return m, nil
}

func (s *httpManipulator) RetrieveMany(mctx manipulate.Context, dest elemental.Identifiables) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.retrieve_many.%s", dest.Identity().Category))
	defer sp.Finish()

	url, err := s.getURLForChildrenIdentity(mctx.Parent(), dest.Identity(), dest.Version(), mctx.Version())
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return manipulate.NewErrCannotBuildQuery(err.Error())
	}

	response, err := s.send(mctx, http.MethodGet, url, nil, sp, 0)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if response.StatusCode == http.StatusNoContent || response.ContentLength == 0 {
		return nil
	}

	if response.StatusCode != http.StatusNoContent {
		defer response.Body.Close() // nolint: errcheck
		if err := decodeData(response, dest); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}
	}

	// backport all default values that are empty.
	for _, o := range dest.List() {
		if a, ok := o.(elemental.AttributeSpecifiable); ok {
			elemental.ResetDefaultForZeroValues(a)
		}
	}

	return nil
}

func (s *httpManipulator) Retrieve(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(s.ctx)
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.retrieve.object.%s", object.Identity().Name))
	defer sp.Finish()

	url, err := s.getPersonalURL(object, mctx.Version())
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return manipulate.NewErrCannotBuildQuery(err.Error())
	}

	response, err := s.send(mctx, http.MethodGet, url, nil, sp, 0)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		defer response.Body.Close() // nolint: errcheck
		if err := decodeData(response, object); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}
	}

	// backport all default values that are empty.
	if a, ok := object.(elemental.AttributeSpecifiable); ok {
		elemental.ResetDefaultForZeroValues(a)
	}

	return nil
}

func (s *httpManipulator) Create(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	kmctx, _ := mctx.(idempotency.Keyer)
	if kmctx != nil && kmctx.IdempotencyKey() == "" {
		kmctx.SetIdempotencyKey(uuid.Must(uuid.NewV4()).String())
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.create.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()))
	defer sp.Finish()

	url, err := s.getURLForChildrenIdentity(mctx.Parent(), object.Identity(), object.Version(), mctx.Version())
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return manipulate.NewErrCannotBuildQuery(err.Error())
	}

	data, err := elemental.Encode(s.encoding, object)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return manipulate.NewErrCannotMarshal(err.Error())
	}

	response, err := s.send(mctx, http.MethodPost, url, data, sp, 0)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		defer response.Body.Close() // nolint: errcheck
		if err := decodeData(response, object); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}
	}

	// backport all default values that are empty.
	if a, ok := object.(elemental.AttributeSpecifiable); ok {
		elemental.ResetDefaultForZeroValues(a)
	}

	if kmctx != nil {
		kmctx.SetIdempotencyKey("")
	}

	return nil
}

func (s *httpManipulator) Update(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	kmctx, _ := mctx.(idempotency.Keyer)
	if kmctx != nil && kmctx.IdempotencyKey() == "" {
		kmctx.SetIdempotencyKey(uuid.Must(uuid.NewV4()).String())
	}

	method := http.MethodPut
	if _, ok := object.(elemental.SparseIdentifiable); ok {
		method = http.MethodPatch
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.update.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()))
	defer sp.Finish()

	url, err := s.getPersonalURL(object, mctx.Version())
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return manipulate.NewErrCannotBuildQuery(err.Error())
	}

	data, err := elemental.Encode(s.encoding, object)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return manipulate.NewErrCannotMarshal(err.Error())
	}

	response, err := s.send(mctx, method, url, data, sp, 0)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		defer response.Body.Close() // nolint: errcheck
		if err := decodeData(response, object); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}
	}

	// backport all default values that are empty.
	if a, ok := object.(elemental.AttributeSpecifiable); ok {
		elemental.ResetDefaultForZeroValues(a)
	}

	if kmctx != nil {
		kmctx.SetIdempotencyKey("")
	}

	return nil
}

func (s *httpManipulator) Delete(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.delete.object.%s", object.Identity().Name))
	sp.LogFields(log.String("object_id", object.Identifier()))
	defer sp.Finish()

	url, err := s.getPersonalURL(object, mctx.Version())
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return manipulate.NewErrCannotBuildQuery(err.Error())
	}

	response, err := s.send(mctx, http.MethodDelete, url, nil, sp, 0)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		defer response.Body.Close() // nolint: errcheck
		if err := decodeData(response, object); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}
	}

	// backport all default values that are empty.
	if a, ok := object.(elemental.AttributeSpecifiable); ok {
		elemental.ResetDefaultForZeroValues(a)
	}

	return nil
}

func (s *httpManipulator) DeleteMany(mctx manipulate.Context, identity elemental.Identity) error {
	return manipulate.NewErrNotImplemented("DeleteMany not implemented in maniphttp")
}

func (s *httpManipulator) Count(mctx manipulate.Context, identity elemental.Identity) (int, error) {

	if mctx == nil {
		mctx = manipulate.NewContext(context.Background())
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.count.%s", identity.Category))
	defer sp.Finish()

	url, err := s.getURLForChildrenIdentity(mctx.Parent(), identity, 0, mctx.Version())
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return 0, manipulate.NewErrCannotBuildQuery(err.Error())
	}

	if _, err = s.send(mctx, http.MethodHead, url, nil, sp, 0); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return 0, err
	}

	return mctx.Count(), nil
}

func (s *httpManipulator) makeAuthorizationHeaders(username, password string) string {

	return username + " " + password
}

func (s *httpManipulator) prepareHeaders(request *http.Request, mctx manipulate.Context) {

	ns := mctx.Namespace()
	if ns == "" {
		ns = s.namespace
	}

	for k, v := range s.globalHeaders {
		request.Header[k] = v
	}

	request.Header.Set("Content-Type", string(s.encoding))
	request.Header.Set("Accept", string(s.encoding))
	request.Header.Set("Accept-Encoding", "gzip")

	if ns != "" {
		request.Header.Set("X-Namespace", ns)
	}

	username, password := mctx.Credentials()
	if password == "" {
		password = s.currentPassword()
	}

	if username == "" {
		username = s.username
	}

	if username != "" && password != "" {
		request.Header.Set("Authorization", s.makeAuthorizationHeaders(username, password))
	}

	if v := mctx.ExternalTrackingID(); v != "" {
		request.Header.Set("X-External-Tracking-ID", v)
	}

	if v := mctx.ExternalTrackingType(); v != "" {
		request.Header.Set("X-External-Tracking-Type", v)
	}

	if v := mctx.ReadConsistency(); v != manipulate.ReadConsistencyDefault {
		request.Header.Set("X-Read-Consistency", string(v))
	}

	if v := mctx.WriteConsistency(); v != manipulate.WriteConsistencyDefault {
		request.Header.Set("X-Write-Consistency", string(v))
	}

	if k, ok := mctx.(idempotency.Keyer); ok && k.IdempotencyKey() != "" {
		request.Header.Set("Idempotency-Key", k.IdempotencyKey())
	}

	for _, field := range mctx.Fields() {
		request.Header.Add("X-Fields", field)
	}
}

func (s *httpManipulator) readHeaders(response *http.Response, mctx manipulate.Context) {

	if mctx == nil {
		return
	}

	t, _ := strconv.Atoi(response.Header.Get("X-Count-Total"))

	mctx.SetCount(t)
	mctx.SetMessages(response.Header["X-Messages"])
}

func (s *httpManipulator) computeVersion(modelVersion int, mctxVersion int) string {

	if mctxVersion > 0 {
		return fmt.Sprintf("v/%d/", mctxVersion)
	}

	if modelVersion > 0 {
		return fmt.Sprintf("v/%d/", modelVersion)
	}

	return ""
}

func (s *httpManipulator) getGeneralURL(o elemental.Identifiable, mctxVersion int) string {

	v := s.computeVersion(o.Version(), mctxVersion)

	return s.url + "/" + v + o.Identity().Category
}

func (s *httpManipulator) getPersonalURL(o elemental.Identifiable, mctxVersion int) (string, error) {

	if o.Identifier() == "" {
		return "", fmt.Errorf("cannot GetPersonalURL of an object with no ID set")
	}

	return s.getGeneralURL(o, mctxVersion) + "/" + o.Identifier(), nil
}

func (s *httpManipulator) getURLForChildrenIdentity(
	parent elemental.Identifiable,
	childrenIdentity elemental.Identity,
	modelVersion int,
	mctxVersion int,
) (string, error) {

	if parent == nil {
		v := s.computeVersion(modelVersion, mctxVersion)
		return s.url + "/" + v + childrenIdentity.Category, nil
	}

	url, err := s.getPersonalURL(parent, mctxVersion)
	if err != nil {
		return "", err
	}

	return url + "/" + childrenIdentity.Category, nil
}

func (s *httpManipulator) send(mctx manipulate.Context, method string, requrl string, body []byte, sp opentracing.Span, try int) (*http.Response, error) {

	request, err := http.NewRequest(method, requrl, bytes.NewBuffer(body))
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return nil, manipulate.NewErrCannotExecuteQuery(err.Error())
	}
	if err = addQueryParameters(request, mctx); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return nil, err
	}

	if err = sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(request.Header)); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return nil, err
	}

	retryErrCom := func(resp *http.Response, err error) (*http.Response, error) {

		if s.disableAutoRetry {
			return resp, err
		}

		if try >= 3 {
			return resp, err
		}
		<-time.After(time.Duration(try) * time.Second)
		try++
		return s.send(mctx, method, requrl, body, sp, try)
	}

	s.prepareHeaders(request, mctx)

	request = request.WithContext(mctx.Context())

	response, err := s.client.Do(request)
	if err != nil {
		if uerr, ok := err.(*url.Error); ok {
			switch uerr.Err.(type) {
			case x509.UnknownAuthorityError, x509.CertificateInvalidError, x509.HostnameError:
				return response, manipulate.NewErrTLS(err.Error())
			}
		}

		return retryErrCom(response, manipulate.NewErrCannotCommunicate(snip.Snip(err, s.currentPassword()).Error()))
	}

	if response.StatusCode == http.StatusBadGateway {
		return retryErrCom(response, manipulate.NewErrCannotCommunicate("Bad gateway"))
	}

	if response.StatusCode == http.StatusServiceUnavailable {
		return retryErrCom(response, manipulate.NewErrCannotCommunicate("Service unavailable"))
	}

	if response.StatusCode == http.StatusGatewayTimeout {
		return retryErrCom(response, manipulate.NewErrCannotCommunicate("Gateway timeout"))
	}

	if response.StatusCode == http.StatusLocked {
		return retryErrCom(response, manipulate.NewErrLocked("The api has been locked down by the server."))
	}

	// If we get a forbidden or auth error, we try to renew the token and retry the request 3 times
	if (response.StatusCode == http.StatusForbidden || response.StatusCode == http.StatusUnauthorized) && s.tokenManager != nil && try < 3 {

		<-time.After(time.Duration(try) * time.Second)
		try++

		token, err := s.tokenManager.Issue(mctx.Context())
		if err == nil {
			s.setPassword(token)
			return s.send(mctx, method, requrl, body, sp, try)
		}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {

		es := elemental.Errors{}

		defer response.Body.Close() // nolint: errcheck
		if err := decodeData(response, &es); err != nil {
			return nil, err
		}

		errs := elemental.NewErrors()

		for _, e := range es {
			errs = append(errs, e)
		}

		if response.StatusCode == http.StatusRequestTimeout {
			return retryErrCom(response, manipulate.NewErrCannotCommunicate(errs.Error()))
		}

		if response.StatusCode == http.StatusTooManyRequests {
			return retryErrCom(response, manipulate.NewErrTooManyRequests(errs.Error()))
		}

		return response, errs
	}

	s.readHeaders(response, mctx)

	return response, nil
}

func (s *httpManipulator) registerRenewNotifier(id string, f func(string)) {

	s.renewNotifiersLock.Lock()
	s.renewNotifiers[id] = f
	s.renewNotifiersLock.Unlock()
}

func (s *httpManipulator) unregisterRenewNotifier(id string) {

	s.renewNotifiersLock.Lock()
	delete(s.renewNotifiers, id)
	s.renewNotifiersLock.Unlock()
}

func (s *httpManipulator) setPassword(password string) {

	s.renewLock.Lock()
	s.password = password
	s.renewLock.Unlock()

	s.renewNotifiersLock.RLock()
	for _, f := range s.renewNotifiers {
		if f != nil {
			f(password)
		}
	}
	s.renewNotifiersLock.RUnlock()
}

func (s *httpManipulator) currentPassword() string {
	s.renewLock.RLock()
	p := s.password
	s.renewLock.RUnlock()
	return p
}
