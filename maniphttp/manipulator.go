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
	"io"
	"net"
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
	"go.aporeto.io/manipulate/internal/backoff"
	"go.aporeto.io/manipulate/internal/idempotency"
	"go.aporeto.io/manipulate/internal/snip"
	"go.aporeto.io/manipulate/internal/tracing"
)

const (
	defaultGlobalContextTimeout = 2 * time.Minute
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
	disableCompression bool
	defaultRetryFunc   manipulate.RetryFunc

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
			m.transport.DisableCompression = m.disableCompression

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
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.retrieve_many.%s", dest.Identity().Category))
	defer sp.Finish()

	url, err := s.getURLForChildrenIdentity(mctx.Parent(), dest.Identity(), dest.Version(), mctx.Version())
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return manipulate.NewErrCannotBuildQuery(err.Error())
	}

	response, err := s.send(mctx, http.MethodGet, url, nil, dest, sp)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if response.StatusCode == http.StatusNoContent || response.ContentLength == 0 {
		return nil
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

	response, err := s.send(mctx, http.MethodGet, url, nil, object, sp)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if response.StatusCode == http.StatusNoContent {
		return nil
	}

	// backport all default values that are empty.
	if a, ok := object.(elemental.AttributeSpecifiable); ok {
		elemental.ResetDefaultForZeroValues(a)
	}

	return nil
}

func (s *httpManipulator) Create(mctx manipulate.Context, object elemental.Identifiable) error {

	if mctx == nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
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

	response, err := s.send(mctx, http.MethodPost, url, data, object, sp)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if response.StatusCode == http.StatusNoContent {
		return nil
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
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
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

	response, err := s.send(mctx, method, url, data, object, sp)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if response.StatusCode == http.StatusNoContent {
		return nil
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
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
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

	response, err := s.send(mctx, http.MethodDelete, url, nil, object, sp)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	if response.StatusCode == http.StatusNoContent {
		return nil
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
		ctx, cancel := context.WithTimeout(context.Background(), defaultGlobalContextTimeout)
		defer cancel()
		mctx = manipulate.NewContext(ctx)
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.count.%s", identity.Category))
	defer sp.Finish()

	url, err := s.getURLForChildrenIdentity(mctx.Parent(), identity, 0, mctx.Version())
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return 0, manipulate.NewErrCannotBuildQuery(err.Error())
	}

	if _, err = s.send(mctx, http.MethodHead, url, nil, nil, sp); err != nil {
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

	if v := mctx.ClientIP(); v != "" {
		request.Header.Set("X-Forwarded-For", v)
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

func (s *httpManipulator) send(
	mctx manipulate.Context,
	method string,
	requrl string,
	body []byte,
	dest interface{},
	sp opentracing.Span,
) (*http.Response, error) {

	var try int         // try number. Starts at 0
	var lastError error // last error before retry.

	// We get the context deadline.
	deadline, ok := mctx.Context().Deadline()
	if !ok {
		deadline = time.Now().Add(time.Hour) // long, but not completely unlimited.
	}

	// We bucketize the deadline into multiple retries
	// a make it a minimum of 10sec
	subContextTimeout := time.Until(deadline) / 10
	if subContextTimeout < 10*time.Second {
		subContextTimeout = 10 * time.Second
	}

	// Helpers to deal with current request canceling
	var cancelReq context.CancelFunc
	cancelCurrentRequest := func() {
		if cancelReq != nil {
			cancelReq()
		}
	}
	defer cancelCurrentRequest()

	// Helpers to deal with closing the body of the current request
	var bodyCloser io.ReadCloser
	closeCurrentBody := func() {
		if bodyCloser != nil {
			bodyCloser.Close() // nolint
		}
	}
	defer closeCurrentBody()

	// Function that creates a new request to avoid reusing some buffers.
	// It also sets the current request cancel function.
	newRequest := func() (*http.Request, error) {

		req, err := http.NewRequest(method, requrl, bytes.NewBuffer(body))
		if err != nil {
			return nil, manipulate.NewErrCannotBuildQuery(err.Error())
		}

		if err = addQueryParameters(req, mctx); err != nil {
			return nil, err
		}

		if err = sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
			return nil, err
		}

		// We injects the header from mctx.
		s.prepareHeaders(req, mctx)

		ctx, cancel := context.WithTimeout(mctx.Context(), subContextTimeout)
		cancelReq = cancel

		return req.WithContext(ctx), nil
	}

	// Main retry loop
	for {

		// We spawn a new request
		request, err := newRequest()
		if err != nil {
			return nil, err
		}

		// We launch the request
		response, err := s.client.Do(request)

		if err != nil {

			// Per doc, client.Do always returns an *url.Error.
			uerr := err.(*url.Error)

			// We check for constant errors.
			switch uerr.Err {

			case context.DeadlineExceeded:
				if lastError == nil {
					lastError = manipulate.NewErrCannotCommunicate(snip.Snip(err, s.currentPassword()).Error())
				}
				goto RETRY

			case io.ErrUnexpectedEOF, io.EOF:
				lastError = manipulate.NewErrCannotCommunicate(snip.Snip(err, s.currentPassword()).Error())
				goto RETRY
			}

			// We check for error types.
			switch uerr.Err.(type) {

			case net.Error:
				lastError = manipulate.NewErrCannotCommunicate(snip.Snip(err, s.currentPassword()).Error())
				goto RETRY

			case x509.UnknownAuthorityError, x509.CertificateInvalidError, x509.HostnameError:
				return nil, manipulate.NewErrTLS(err.Error())

			default:
				return nil, manipulate.NewErrCannotExecuteQuery(err.Error())
			}
		}

		// We passed the basic error, we have a body.
		// We register it so next loop will be clean.
		bodyCloser = request.Body

		// We check for http status codes that triggers a retry
		switch response.StatusCode {

		case http.StatusBadGateway:
			lastError = manipulate.NewErrCannotCommunicate("Bad gateway")
			goto RETRY

		case http.StatusServiceUnavailable:
			lastError = manipulate.NewErrCannotCommunicate("Service unavailable")
			goto RETRY

		case http.StatusGatewayTimeout:
			lastError = manipulate.NewErrCannotCommunicate("Gateway timeout")
			goto RETRY

		case http.StatusLocked:
			lastError = manipulate.NewErrLocked("The api has been locked down by the server.")
			goto RETRY

		case http.StatusRequestTimeout:
			lastError = manipulate.NewErrCannotCommunicate("Request Timeout")
			goto RETRY

		case http.StatusTooManyRequests:
			lastError = manipulate.NewErrTooManyRequests("Too Many Requests")
			goto RETRY

		case http.StatusForbidden, http.StatusUnauthorized:

			// This is a special case where we try to renew our token
			// in case of 401 or 403 error.
			// This is retried twice.
			if s.tokenManager != nil {

				for i := 1; i < 3; i++ {
					time.Sleep(time.Duration(i) * time.Second)

					token, err := s.tokenManager.Issue(mctx.Context())
					if err == nil {
						s.setPassword(token)
						goto RETRY
					}
				}
			}

			// If we could not renew the token, we continue.
		}

		// We backport header info into mctx
		s.readHeaders(response, mctx)

		// If we have some other errors, we decode them.
		if response.StatusCode < 200 || response.StatusCode >= 300 {
			errs := elemental.NewErrors()
			if err := decodeData(response, &errs); err != nil {
				return nil, err
			}

			return nil, errs
		}

		//
		// From now on, this is a success.
		//

		// If we have content, we return the response.
		// The body will be drained by the defered call to closeCurrentBody().
		if response.StatusCode == http.StatusNoContent || response.ContentLength == 0 || dest == nil {
			return response, nil
		}

		// If we have a given dest to decode, we decode it now.
		if dest != nil {
			if err := decodeData(response, dest); err != nil {
				return nil, err
			}
		}

		// And we return the response
		return response, nil

	RETRY:
		//
		// From now on, this is a failure.
		//

		// We cancel any pending request context and the pending body.
		cancelCurrentRequest()
		closeCurrentBody()

		// If the manipulator has auto retry disabled we return the last error
		if s.disableAutoRetry {
			return nil, lastError
		}

		info := RetryInfo{
			URL:    requrl,
			Method: method,
			try:    try,
			mctx:   mctx,
			err:    lastError,
		}

		// We run the eventual retry funcs.
		if rf := mctx.RetryFunc(); rf != nil {
			if rerr := rf(info); rerr != nil {
				return nil, rerr
			}
		} else if s.defaultRetryFunc != nil {
			if rerr := s.defaultRetryFunc(info); rerr != nil {
				return nil, rerr
			}
		}

		// We check is the main context expired.
		// and if so, we return the last error
		if time.Until(deadline) <= 0 {
			return nil, lastError
		}

		// Then we sleep backoff and we restart the retry loop.
		time.Sleep(backoff.Next(try, deadline))
		try++
	}
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
