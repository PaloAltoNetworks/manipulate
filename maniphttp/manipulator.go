// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package maniphttp

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/aporeto-inc/addedeffect/tokenutils"
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/tracing"
	"github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"

	opentracing "github.com/opentracing/opentracing-go"
)

type httpManipulator struct {
	username     string
	password     string
	url          string
	namespace    string
	renewLock    *sync.Mutex
	client       *http.Client
	tlsConfig    *tls.Config
	tokenManager manipulate.TokenManager
}

// NewHTTPManipulator returns a Manipulator backed by an ReST API.
func NewHTTPManipulator(url, username, password, namespace string) manipulate.Manipulator {

	CAPool, err := x509.SystemCertPool()
	if err != nil {
		zap.L().Fatal("Unable to load system root cert pool", zap.Error(err))
	}

	return NewHTTPManipulatorWithTLS(
		url,
		username,
		password,
		namespace,
		&tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            CAPool,
		},
	)
}

// NewHTTPManipulatorWithTLS returns a Manipulator backed by an ReST API using the tls config.
func NewHTTPManipulatorWithTLS(url, username, password, namespace string, tlsConfig *tls.Config) manipulate.Manipulator {

	m, err := NewHTTPManipulatorWithTokenManager(context.Background(), url, namespace, tlsConfig, nil)
	if err != nil {
		panic(err)
	}

	m.(*httpManipulator).password = password
	m.(*httpManipulator).username = username

	return m
}

// NewHTTPManipulatorWithTokenManager returns a http backed manipulate.Manipulatorusing the given manipulate.TokenManager to manage the the token.
func NewHTTPManipulatorWithTokenManager(ctx context.Context, url string, namespace string, tlsConfig *tls.Config, tokenManager manipulate.TokenManager) (manipulate.Manipulator, error) {

	var username, password string

	if tokenManager != nil {

		issueCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		token, err := tokenManager.Issue(issueCtx)
		if err != nil {
			return nil, err
		}

		password = token
		username = "Bearer"
	}

	m := &httpManipulator{
		username:     username,
		password:     password,
		namespace:    namespace,
		tokenManager: tokenManager,
		renewLock:    &sync.Mutex{},
		url:          url,
		tlsConfig:    tlsConfig,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
	}

	if tokenManager != nil {
		go func() {
			tokenCh := make(chan string)

			go tokenManager.Run(ctx, tokenCh)

			for {
				select {
				case t := <-tokenCh:
					m.setPassword(t)
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	return m, nil
}

func (s *httpManipulator) RetrieveMany(mctx *manipulate.Context, dest elemental.ContentIdentifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.retrieve_many.%s", dest.ContentIdentity().Category))
	defer sp.Finish()

	url, err := s.getURLForChildrenIdentity(mctx.Parent, dest.ContentIdentity(), dest.Version(), mctx.Version)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return manipulate.NewErrCannotBuildQuery(err.Error())
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}
	if err = addQueryParameters(request, mctx); err != nil {
		return err
	}

	if err = sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(request.Header)); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return err
	}

	response, err := s.send(mctx, request)
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
		if err := json.NewDecoder(response.Body).Decode(dest); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotUnmarshal(err.Error())
		}
	}

	return nil
}

func (s *httpManipulator) Retrieve(mctx *manipulate.Context, objects ...elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	for _, object := range objects {

		sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.retrieve.object.%s", object.Identity().Name))
		defer sp.Finish()

		url, err := s.getPersonalURL(object, mctx.Version)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		if err = addQueryParameters(request, mctx); err != nil {
			return err
		}

		if err = sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(request.Header)); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}

		response, err := s.send(mctx, request)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}

		if response.StatusCode != http.StatusNoContent {
			defer response.Body.Close() // nolint: errcheck
			if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
				sp.SetTag("error", true)
				sp.LogFields(log.Error(err))
				return manipulate.NewErrCannotUnmarshal(err.Error())
			}
		}
	}

	return nil
}

func (s *httpManipulator) Create(mctx *manipulate.Context, objects ...elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	for _, child := range objects {

		sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.create.object.%s", child.Identity().Name))
		sp.LogFields(log.String("object_id", child.Identifier()))
		defer sp.Finish()

		url, err := s.getURLForChildrenIdentity(mctx.Parent, child.Identity(), child.Version(), mctx.Version)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		buffer := &bytes.Buffer{}
		if err = json.NewEncoder(buffer).Encode(&child); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotMarshal(err.Error())
		}

		request, err := http.NewRequest(http.MethodPost, url, buffer)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		if err = addQueryParameters(request, mctx); err != nil {
			return err
		}

		if err = sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(request.Header)); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}

		response, err := s.send(mctx, request)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}

		if response.StatusCode != http.StatusNoContent {
			defer response.Body.Close() // nolint: errcheck
			if err := json.NewDecoder(response.Body).Decode(&child); err != nil {
				sp.SetTag("error", true)
				sp.LogFields(log.Error(err))
				return manipulate.NewErrCannotUnmarshal(err.Error())
			}
		}
	}

	return nil
}

func (s *httpManipulator) Update(mctx *manipulate.Context, objects ...elemental.Identifiable) error {

	if len(objects) == 0 {
		return nil
	}

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	for _, object := range objects {

		sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.update.object.%s", object.Identity().Name))
		sp.LogFields(log.String("object_id", object.Identifier()))
		defer sp.Finish()

		url, err := s.getPersonalURL(object, mctx.Version)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		buffer := &bytes.Buffer{}
		if err = json.NewEncoder(buffer).Encode(object); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotMarshal(err.Error())
		}

		request, err := http.NewRequest(http.MethodPut, url, buffer)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		if err = addQueryParameters(request, mctx); err != nil {
			return err
		}

		if err = sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(request.Header)); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}

		response, err := s.send(mctx, request)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}

		if response.StatusCode != http.StatusNoContent {
			defer response.Body.Close() // nolint: errcheck
			if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
				sp.SetTag("error", true)
				sp.LogFields(log.Error(err))
				return manipulate.NewErrCannotUnmarshal(err.Error())
			}
		}
	}

	return nil
}

func (s *httpManipulator) Delete(mctx *manipulate.Context, objects ...elemental.Identifiable) error {

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	for _, object := range objects {

		sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.delete.object.%s", object.Identity().Name))
		sp.LogFields(log.String("object_id", object.Identifier()))
		defer sp.Finish()

		url, err := s.getPersonalURL(object, mctx.Version)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		request, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		if err = addQueryParameters(request, mctx); err != nil {
			return err
		}

		if err = sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(request.Header)); err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}

		response, err := s.send(mctx, request)
		if err != nil {
			sp.SetTag("error", true)
			sp.LogFields(log.Error(err))
			return err
		}

		if response.StatusCode != http.StatusNoContent {
			defer response.Body.Close() // nolint: errcheck
			if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
				sp.SetTag("error", true)
				sp.LogFields(log.Error(err))
				return manipulate.NewErrCannotUnmarshal(err.Error())
			}
		}
	}

	return nil
}

func (s *httpManipulator) DeleteMany(mctx *manipulate.Context, identity elemental.Identity) error {
	return manipulate.NewErrNotImplemented("DeleteMany not implemented in maniphttp")
}

func (s *httpManipulator) Count(mctx *manipulate.Context, identity elemental.Identity) (int, error) {

	if mctx == nil {
		mctx = manipulate.NewContext()
	}

	sp := tracing.StartTrace(mctx, fmt.Sprintf("maniphttp.count.%s", identity.Category))
	defer sp.Finish()

	url, err := s.getURLForChildrenIdentity(mctx.Parent, identity, 0, mctx.Version)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return 0, manipulate.NewErrCannotBuildQuery(err.Error())
	}

	request, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return 0, manipulate.NewErrCannotExecuteQuery(err.Error())
	}

	if err = addQueryParameters(request, mctx); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return 0, err
	}

	if err = sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(request.Header)); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return 0, err
	}

	if _, err = s.send(mctx, request); err != nil {
		sp.SetTag("error", true)
		sp.LogFields(log.Error(err))
		return 0, err
	}

	return mctx.CountTotal, nil
}

func (s *httpManipulator) makeAuthorizationHeaders() string {

	return s.username + " " + s.currentPassword()
}

func (s *httpManipulator) prepareHeaders(request *http.Request, mctx *manipulate.Context) {

	if mctx.Namespace == "" {
		mctx.Namespace = s.namespace
	}

	if mctx.Namespace != "" {
		request.Header.Set("X-Namespace", mctx.Namespace)
	}

	if s.username != "" && s.currentPassword() != "" {
		request.Header.Set("Authorization", s.makeAuthorizationHeaders())
	}

	if mctx.ExternalTrackingID != "" {
		request.Header.Set("X-External-Tracking-ID", mctx.ExternalTrackingID)
	}

	if mctx.ExternalTrackingType != "" {
		request.Header.Set("X-External-Tracking-Type", mctx.ExternalTrackingType)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
}

func (s *httpManipulator) readHeaders(response *http.Response, mctx *manipulate.Context) {

	if mctx == nil {
		return
	}

	mctx.CountTotal, _ = strconv.Atoi(response.Header.Get("X-Count-Total"))
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
		return "", fmt.Errorf("Cannot GetPersonalURL of an object with no ID set")
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

func (s *httpManipulator) send(mctx *manipulate.Context, request *http.Request) (*http.Response, error) {

	s.prepareHeaders(request, mctx)

	response, err := s.client.Do(request)
	if err != nil {
		return response, manipulate.NewErrCannotCommunicate(tokenutils.Snip(err, s.currentPassword()).Error())
	}

	if response.StatusCode == http.StatusBadGateway {
		return response, manipulate.NewErrCannotCommunicate("Service unavailable")
	}

	if response.StatusCode == http.StatusLocked {
		return response, manipulate.NewErrLocked("The api has been locked down by the server.")
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

	s.readHeaders(response, mctx)

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
