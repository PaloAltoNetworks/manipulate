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

	"go.uber.org/zap"

	"github.com/aporeto-inc/addedeffect/tokensnip"
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/tracing"
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

	for _, object := range objects {

		sp := tracing.StartTrace(context.TrackingSpan, fmt.Sprintf("maniphttp.retrieve.object.%s", object.Identity().Name), context)
		tracing.SetTag(sp, "maniphttp.retrieve.object.id", object.Identifier())

		url, err := s.getPersonalURL(object, context.Version)
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

		if response.StatusCode != http.StatusNoContent {
			defer response.Body.Close() // nolint: errcheck
			if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
				tracing.FinishTraceWithError(sp, err)
				return manipulate.NewErrCannotUnmarshal(err.Error())
			}
		}

		tracing.FinishTrace(sp)
	}

	return nil
}

func (s *httpManipulator) Create(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	for _, child := range objects {

		subSp := tracing.StartTrace(context.TrackingSpan, fmt.Sprintf("maniphttp.create.object.%s", child.Identity().Name), context)
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

	for _, object := range objects {

		sp := tracing.StartTrace(context.TrackingSpan, fmt.Sprintf("maniphttp.update.object.%s", object.Identity().Name), context)
		tracing.SetTag(sp, "maniphttp.update.object.id", object.Identifier())

		url, err := s.getPersonalURL(object, context.Version)
		if err != nil {
			tracing.FinishTraceWithError(sp, err)
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		buffer := &bytes.Buffer{}
		if err = json.NewEncoder(buffer).Encode(object); err != nil {
			tracing.FinishTraceWithError(sp, err)
			return manipulate.NewErrCannotMarshal(err.Error())
		}

		request, err := http.NewRequest(http.MethodPut, url, buffer)
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

		if response.StatusCode != http.StatusNoContent {
			defer response.Body.Close() // nolint: errcheck
			if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
				tracing.FinishTraceWithError(sp, err)
				return manipulate.NewErrCannotUnmarshal(err.Error())
			}
		}

		tracing.FinishTrace(sp)
	}

	return nil
}

func (s *httpManipulator) Delete(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	for _, object := range objects {

		sp := tracing.StartTrace(context.TrackingSpan, fmt.Sprintf("maniphttp.delete.object.%s", object.Identity().Name), context)
		tracing.SetTag(sp, "maniphttp.delete.object.id", object.Identifier())

		url, err := s.getPersonalURL(object, context.Version)
		if err != nil {
			tracing.FinishTraceWithError(sp, err)
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		request, err := http.NewRequest(http.MethodDelete, url, nil)
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

		if response.StatusCode != http.StatusNoContent {
			defer response.Body.Close() // nolint: errcheck
			if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
				tracing.FinishTraceWithError(sp, err)
				return manipulate.NewErrCannotUnmarshal(err.Error())
			}
		}

		tracing.FinishTrace(sp)
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
