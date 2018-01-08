// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package maniphttp

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/aporeto-inc/addedeffect/tokensnip"
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/auth"
	"github.com/aporeto-inc/manipulate/internal/tracing"
	"github.com/opentracing/opentracing-go"

	midgardclient "github.com/aporeto-inc/midgard-lib/client"
)

type httpManipulator struct {
	username      string
	password      string
	url           string
	namespace     string
	renewLock     *sync.Mutex
	client        *http.Client
	tlsConfig     *tls.Config
	mclient       *midgardclient.Client
	tokenValidity time.Duration
}

// NewHTTPManipulator returns a Manipulator backed by an ReST API.
func NewHTTPManipulator(username, password, url, namespace string) manipulate.Manipulator {

	CAPool, err := x509.SystemCertPool()
	if err != nil {
		zap.L().Fatal("Unable to load system root cert pool", zap.Error(err))
	}

	return NewHTTPManipulatorWithRootCA(username, password, url, namespace, CAPool, true)
}

// NewHTTPManipulatorWithRootCA returns a Manipulator backed by an ReST API using the given CAPool as root CA.
func NewHTTPManipulatorWithRootCA(username, password, url, namespace string, rootCAPool *x509.CertPool, skipTLSVerify bool) manipulate.Manipulator {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: skipTLSVerify,
		RootCAs:            rootCAPool,
	}

	return &httpManipulator{
		username:  username,
		password:  password,
		renewLock: &sync.Mutex{},
		url:       url,
		client: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				IdleConnTimeout:     30 * time.Second,
				MaxIdleConnsPerHost: 100,
				TLSClientConfig:     tlsConfig,
			},
		},
		namespace: namespace,
		tlsConfig: tlsConfig,
	}
}

// NewHTTPManipulatorWithMTLS returns a Manipulator backed by an ReST API using the given CAPool as root CA and client certificate.
func NewHTTPManipulatorWithMTLS(url string, namespace string, rootCAPool *x509.CertPool, clientCertificates []tls.Certificate, skipTLSVerify bool) manipulate.Manipulator {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: skipTLSVerify,
		RootCAs:            rootCAPool,
		Certificates:       clientCertificates,
	}

	return &httpManipulator{
		renewLock: &sync.Mutex{},
		namespace: namespace,
		url:       url,
		client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				IdleConnTimeout:     30 * time.Second,
				MaxIdleConnsPerHost: 100,
				TLSClientConfig:     tlsConfig,
			},
		},
		tlsConfig: tlsConfig,
	}
}

// NewHTTPManipulatorWithMidgardCertAuthentication returns a http backed manipulate.Manipulator
// using a certificates to authenticate against a Midgard server.
func NewHTTPManipulatorWithMidgardCertAuthentication(
	url string,
	midgardurl string,
	rootCAPool *x509.CertPool,
	certificates []tls.Certificate,
	namespace string,
	validity time.Duration,
	skipInsecure bool,
) (manipulate.Manipulator, func(), error) {

	sp := opentracing.StartSpan("maniphttp.authenthication")
	defer sp.Finish()

	mclient := midgardclient.NewClientWithCAPool(midgardurl, rootCAPool, skipInsecure)
	token, err := auth.IssueInitialToken(mclient, certificates, validity, sp)
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return nil, nil, err
	}

	tlsConfig := &tls.Config{
		Certificates:       certificates,
		InsecureSkipVerify: skipInsecure,
		RootCAs:            rootCAPool,
	}

	m := &httpManipulator{
		username:  "Bearer",
		password:  token,
		renewLock: &sync.Mutex{},
		url:       url,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
		namespace:     namespace,
		tlsConfig:     tlsConfig,
		mclient:       mclient,
		tokenValidity: validity,
	}

	stopCh := make(chan bool)

	go auth.RenewMidgardToken(m, stopCh)

	return m, func() { stopCh <- true }, nil
}

func (s *httpManipulator) RetrieveMany(context *manipulate.Context, dest elemental.ContentIdentifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, fmt.Sprintf("maniphttp.retrieve_many.%s", dest.ContentIdentity().Category), context)

	url, err := s.getURLForChildrenIdentity(context.Parent, dest.ContentIdentity(), dest.Version(), context.Version)
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return manipulate.NewErrCannotBuildQuery(err.Error())
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}
	if err = addQueryParameters(request, context); err != nil {
		return err
	}

	if err = tracing.InjectInHTTPRequest(sp, request); err != nil {
		tracing.FinishTraceWithError(sp, err)
		return err
	}

	response, err := s.send(request, context)
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return err
	}

	if response.StatusCode == http.StatusNoContent || response.ContentLength == 0 {
		tracing.FinishTrace(sp)
		return nil
	}

	if response.StatusCode != http.StatusNoContent {
		defer response.Body.Close() // nolint: errcheck
		if err := json.NewDecoder(response.Body).Decode(dest); err != nil {
			tracing.FinishTraceWithError(sp, err)
			return manipulate.NewErrCannotUnmarshal(err.Error())
		}
	}

	tracing.FinishTrace(sp)
	return nil
}

func (s *httpManipulator) Retrieve(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, "maniphttp.retrieve", context)
	defer tracing.FinishTrace(sp)

	for _, object := range objects {

		subSp := tracing.StartTrace(sp, fmt.Sprintf("maniphttp.retrieve.object.%s", object.Identity().Name), context)
		tracing.SetTag(subSp, "maniphttp.retrieve.object.id", object.Identifier())

		url, err := s.getPersonalURL(object, context.Version)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		if err = addQueryParameters(request, context); err != nil {
			return err
		}

		if err = tracing.InjectInHTTPRequest(subSp, request); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		response, err := s.send(request, context)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if response.StatusCode != http.StatusNoContent {
			defer response.Body.Close() // nolint: errcheck
			if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
				tracing.FinishTraceWithError(subSp, err)
				return manipulate.NewErrCannotUnmarshal(err.Error())
			}
		}

		tracing.FinishTrace(subSp)
	}

	return nil
}

func (s *httpManipulator) Create(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, "maniphttp.create", context)
	defer tracing.FinishTrace(sp)

	for _, child := range objects {

		subSp := tracing.StartTrace(sp, fmt.Sprintf("maniphttp.create.object.%s", child.Identity().Name), context)
		tracing.SetTag(subSp, "maniphttp.create.object.id", child.Identifier())

		url, err := s.getURLForChildrenIdentity(context.Parent, child.Identity(), child.Version(), context.Version)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		buffer := &bytes.Buffer{}
		if err = json.NewEncoder(buffer).Encode(&child); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotMarshal(err.Error())
		}

		request, err := http.NewRequest(http.MethodPost, url, buffer)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		if err = addQueryParameters(request, context); err != nil {
			return err
		}

		if err = tracing.InjectInHTTPRequest(subSp, request); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		response, err := s.send(request, context)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if response.StatusCode != http.StatusNoContent {
			defer response.Body.Close() // nolint: errcheck
			if err := json.NewDecoder(response.Body).Decode(&child); err != nil {
				tracing.FinishTraceWithError(subSp, err)
				return manipulate.NewErrCannotUnmarshal(err.Error())
			}
		}

		tracing.FinishTrace(subSp)
	}

	return nil
}

func (s *httpManipulator) Update(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, "maniphttp.update", context)
	defer tracing.FinishTrace(sp)

	for _, object := range objects {

		subSp := tracing.StartTrace(sp, fmt.Sprintf("maniphttp.update.object.%s", object.Identity().Name), context)
		tracing.SetTag(subSp, "maniphttp.update.object.id", object.Identifier())

		url, err := s.getPersonalURL(object, context.Version)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		buffer := &bytes.Buffer{}
		if err = json.NewEncoder(buffer).Encode(object); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotMarshal(err.Error())
		}

		request, err := http.NewRequest(http.MethodPut, url, buffer)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		if err = addQueryParameters(request, context); err != nil {
			return err
		}

		if err = tracing.InjectInHTTPRequest(subSp, request); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		response, err := s.send(request, context)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if response.StatusCode != http.StatusNoContent {
			defer response.Body.Close() // nolint: errcheck
			if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
				tracing.FinishTraceWithError(subSp, err)
				return manipulate.NewErrCannotUnmarshal(err.Error())
			}
		}

		tracing.FinishTrace(subSp)
	}

	return nil
}

func (s *httpManipulator) Delete(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, "maniphttp.delete", context)
	defer tracing.FinishTrace(sp)

	for _, object := range objects {

		subSp := tracing.StartTrace(sp, fmt.Sprintf("maniphttp.delete.object.%s", object.Identity().Name), context)
		tracing.SetTag(subSp, "maniphttp.delete.object.id", object.Identifier())

		url, err := s.getPersonalURL(object, context.Version)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		request, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		if err = addQueryParameters(request, context); err != nil {
			return err
		}

		if err = tracing.InjectInHTTPRequest(subSp, request); err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		response, err := s.send(request, context)
		if err != nil {
			tracing.FinishTraceWithError(subSp, err)
			return err
		}

		if response.StatusCode != http.StatusNoContent {
			defer response.Body.Close() // nolint: errcheck
			if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
				tracing.FinishTraceWithError(subSp, err)
				return manipulate.NewErrCannotUnmarshal(err.Error())
			}
		}

		tracing.FinishTrace(subSp)
	}

	return nil
}

func (s *httpManipulator) DeleteMany(context *manipulate.Context, identity elemental.Identity) error {
	return manipulate.NewErrNotImplemented("DeleteMany not implemented in maniphttp")
}

func (s *httpManipulator) Count(context *manipulate.Context, identity elemental.Identity) (int, error) {

	if context == nil {
		context = manipulate.NewContext()
	}

	sp := tracing.StartTrace(context.TrackingSpan, fmt.Sprintf("maniphttp.count.%s", identity.Category), context)

	url, err := s.getURLForChildrenIdentity(context.Parent, identity, 0, context.Version)
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return 0, manipulate.NewErrCannotBuildQuery(err.Error())
	}

	request, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return 0, manipulate.NewErrCannotExecuteQuery(err.Error())
	}
	if err = addQueryParameters(request, context); err != nil {
		return 0, err
	}

	if err = tracing.InjectInHTTPRequest(sp, request); err != nil {
		tracing.FinishTraceWithError(sp, err)
		return 0, err
	}

	_, err = s.send(request, context)
	if err != nil {
		tracing.FinishTraceWithError(sp, err)
		return 0, err
	}

	tracing.FinishTrace(sp)

	return context.CountTotal, nil
}

func (s *httpManipulator) Increment(context *manipulate.Context, identity elemental.Identity, counter string, inc int) error {

	return manipulate.NewErrNotImplemented("Increment method not implemented in http manipulator")
}

func (s *httpManipulator) makeAuthorizationHeaders() string {

	return s.username + " " + s.currentPassword()
}

func (s *httpManipulator) prepareHeaders(request *http.Request, context *manipulate.Context) {

	if context.Namespace == "" {
		context.Namespace = s.namespace
	}

	if context.Namespace != "" {
		request.Header.Set("X-Namespace", context.Namespace)
	}

	if s.username != "" && s.currentPassword() != "" {
		request.Header.Set("Authorization", s.makeAuthorizationHeaders())
	}

	if context.ExternalTrackingID != "" {
		request.Header.Set("X-External-Tracking-ID", context.ExternalTrackingID)
	}

	if context.ExternalTrackingType != "" {
		request.Header.Set("X-External-Tracking-Type", context.ExternalTrackingType)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
}

func (s *httpManipulator) readHeaders(response *http.Response, context *manipulate.Context) {

	if context == nil {
		return
	}

	context.CountTotal, _ = strconv.Atoi(response.Header.Get("X-Count-Total"))
}

func (s *httpManipulator) computeVersion(modelVersion int, contextVersion int) string {

	if contextVersion > 0 {
		return fmt.Sprintf("v/%d/", contextVersion)
	}

	if modelVersion > 0 {
		return fmt.Sprintf("v/%d/", modelVersion)
	}

	return ""
}

func (s *httpManipulator) getGeneralURL(o elemental.Identifiable, contextVersion int) string {

	v := s.computeVersion(o.Version(), contextVersion)

	return s.url + "/" + v + o.Identity().Category
}

func (s *httpManipulator) getPersonalURL(o elemental.Identifiable, contextVersion int) (string, error) {

	if o.Identifier() == "" {
		return "", fmt.Errorf("Cannot GetPersonalURL of an object with no ID set")
	}

	return s.getGeneralURL(o, contextVersion) + "/" + o.Identifier(), nil
}

func (s *httpManipulator) getURLForChildrenIdentity(
	parent elemental.Identifiable,
	childrenIdentity elemental.Identity,
	modelVersion int,
	contextVersion int,
) (string, error) {

	if parent == nil {
		v := s.computeVersion(modelVersion, contextVersion)
		return s.url + "/" + v + childrenIdentity.Category, nil
	}

	url, err := s.getPersonalURL(parent, contextVersion)
	if err != nil {
		return "", err
	}

	return url + "/" + childrenIdentity.Category, nil
}

func (s *httpManipulator) send(request *http.Request, context *manipulate.Context) (*http.Response, error) {

	s.prepareHeaders(request, context)

	response, err := s.client.Do(request)
	if err != nil {
		return response, manipulate.NewErrCannotCommunicate(tokensnip.Snip(err, s.currentPassword()).Error())
	}

	if response.StatusCode == http.StatusLocked {
		return response, manipulate.NewErrLocked("The api has been locked down by the server.")
	}

	if response.StatusCode == http.StatusUnauthorized && s.mclient != nil {
		zap.L().Warn("Authentication error. Trying to renew the token...")
		for i := 0; i < 3; i++ {
			if err := s.RetrieveToken(); err == nil {
				return nil, manipulate.NewErrCannotCommunicate("Temporary authentication error. token renewed, please retry.")
			}
			<-time.After(10 * time.Second)
		}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {

		es := []elemental.Error{}

		defer response.Body.Close() // nolint: errcheck
		if err := json.NewDecoder(response.Body).Decode(&es); err != nil {
			return nil, manipulate.NewErrCannotUnmarshal(err.Error())
		}

		errs := elemental.NewErrors()

		for _, e := range es {
			errs = append(errs, e)
		}

		return response, errs
	}

	s.readHeaders(response, context)

	return response, nil
}

func (s *httpManipulator) setPassword(password string) {
	s.renewLock.Lock()
	s.password = password
	s.renewLock.Unlock()
}

func (s *httpManipulator) currentPassword() string {
	s.renewLock.Lock()
	p := s.password
	s.renewLock.Unlock()
	return p
}

func (s *httpManipulator) RetrieveToken() error {

	if s.mclient == nil {
		return fmt.Errorf("this manipulator is not configured to renew tokens")
	}

	token, err := s.mclient.IssueFromCertificateWithValidity(s.tlsConfig.Certificates, s.tokenValidity, nil)
	if err != nil {
		return err
	}

	s.setPassword(token)

	return nil
}

func (s *httpManipulator) Validity() time.Duration {

	return s.tokenValidity
}
