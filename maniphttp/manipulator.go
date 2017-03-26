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
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"

	midgardclient "github.com/aporeto-inc/midgard-lib/client"
)

// Logger contains the main logger.
var Logger = logrus.New()

var log = Logger.WithField("package", "github.com/aporeto-inc/manipulate/maniphttp")

type httpManipulator struct {
	username  string
	password  string
	url       string
	namespace string
	client    *http.Client
	tlsConfig *tls.Config
}

// NewHTTPManipulator returns a Manipulator backed by an ReST API.
func NewHTTPManipulator(username, password, url, namespace string) manipulate.Manipulator {

	CAPool, err := x509.SystemCertPool()
	if err != nil {
		log.Error("Unable to load system root cert pool. tls fallback to unsecure.")
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
		username: username,
		password: password,
		url:      url,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
		namespace: namespace,
		tlsConfig: tlsConfig,
	}
}

// NewHTTPManipulatorWithMidgardCertAuthentication returns a http backed manipulate.Manipulator
// using a certificates to authenticate against a Midgard server.
func NewHTTPManipulatorWithMidgardCertAuthentication(
	url string,
	midgardurl string,
	rootCAPool *x509.CertPool,
	clientCAPool *x509.CertPool,
	certificates []tls.Certificate,
	namespace string,
	refreshInterval time.Duration,
	skipInsecure bool,
) (manipulate.Manipulator, func(), error) {

	mclient := midgardclient.NewClientWithCAPool(midgardurl, rootCAPool, clientCAPool, skipInsecure)
	token, err := mclient.IssueFromCertificate(certificates)
	if err != nil {
		return nil, nil, err
	}

	tlsConfig := &tls.Config{
		Certificates:       certificates,
		InsecureSkipVerify: skipInsecure,
		RootCAs:            rootCAPool,
		ClientCAs:          clientCAPool,
	}

	m := &httpManipulator{
		username: "Bearer",
		password: token,
		url:      url,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
		namespace: namespace,
		tlsConfig: tlsConfig,
	}

	stopCh := make(chan bool)

	go renewMidgardToken(mclient, m, certificates, refreshInterval, stopCh)

	return m, func() { stopCh <- true }, nil
}

func (s *httpManipulator) RetrieveMany(context *manipulate.Context, dest elemental.ContentIdentifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	url, err := s.getURLForChildrenIdentity(context.Parent, dest.ContentIdentity())
	if err != nil {
		return manipulate.NewErrCannotBuildQuery(err.Error())
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}
	addQueryParameters(request, context)

	response, err := s.send(request, context)
	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusNoContent || response.ContentLength == 0 {
		return nil
	}

	defer response.Body.Close() // nolint: errcheck
	if err := json.NewDecoder(response.Body).Decode(dest); err != nil {
		return manipulate.NewErrCannotUnmarshal(err.Error())
	}

	return nil
}

func (s *httpManipulator) Retrieve(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	for _, object := range objects {

		url, err := s.getPersonalURL(object)
		if err != nil {
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		addQueryParameters(request, context)

		response, err := s.send(request, context)
		if err != nil {
			return err
		}

		defer response.Body.Close() // nolint: errcheck
		if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
			return manipulate.NewErrCannotUnmarshal(err.Error())
		}
	}

	return nil
}

func (s *httpManipulator) Create(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	for _, child := range objects {

		url, err := s.getURLForChildrenIdentity(context.Parent, child.Identity())
		if err != nil {
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		buffer := &bytes.Buffer{}
		if err = json.NewEncoder(buffer).Encode(&child); err != nil {
			return manipulate.NewErrCannotMarshal(err.Error())
		}

		request, err := http.NewRequest(http.MethodPost, url, buffer)
		if err != nil {
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		addQueryParameters(request, context)

		response, err := s.send(request, context)
		if err != nil {
			return err
		}

		defer response.Body.Close() // nolint: errcheck
		if err := json.NewDecoder(response.Body).Decode(&child); err != nil {
			return manipulate.NewErrCannotUnmarshal(err.Error())
		}
	}

	return nil
}

func (s *httpManipulator) Update(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	for _, object := range objects {

		url, err := s.getPersonalURL(object)
		if err != nil {
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		buffer := &bytes.Buffer{}
		if err = json.NewEncoder(buffer).Encode(object); err != nil {
			return manipulate.NewErrCannotMarshal(err.Error())
		}

		request, err := http.NewRequest(http.MethodPut, url, buffer)
		if err != nil {
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		addQueryParameters(request, context)

		response, err := s.send(request, context)
		if err != nil {
			return err
		}

		defer response.Body.Close() // nolint: errcheck
		if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
			return manipulate.NewErrCannotUnmarshal(err.Error())
		}
	}

	return nil
}

func (s *httpManipulator) Delete(context *manipulate.Context, objects ...elemental.Identifiable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	for _, object := range objects {

		url, err := s.getPersonalURL(object)
		if err != nil {
			return manipulate.NewErrCannotBuildQuery(err.Error())
		}

		request, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return manipulate.NewErrCannotExecuteQuery(err.Error())
		}
		addQueryParameters(request, context)

		_, err = s.send(request, context)
		if err != nil {
			return err
		}
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

	url, err := s.getURLForChildrenIdentity(context.Parent, identity)
	if err != nil {
		return 0, manipulate.NewErrCannotBuildQuery(err.Error())
	}

	request, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return 0, manipulate.NewErrCannotExecuteQuery(err.Error())
	}
	addQueryParameters(request, context)

	_, err = s.send(request, context)
	if err != nil {
		return 0, err
	}

	return context.CountTotal, nil
}

func (s *httpManipulator) Assign(context *manipulate.Context, assignation *elemental.Assignation) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	url, err := s.getURLForChildrenIdentity(context.Parent, assignation.MembersIdentity)
	if err != nil {
		return manipulate.NewErrCannotBuildQuery(err.Error())
	}

	buffer := &bytes.Buffer{}
	if err = json.NewEncoder(buffer).Encode(assignation); err != nil {
		return manipulate.NewErrCannotMarshal(err.Error())
	}

	request, err := http.NewRequest(http.MethodPatch, url, buffer)
	if err != nil {
		return manipulate.NewErrCannotExecuteQuery(err.Error())
	}
	addQueryParameters(request, context)

	_, err = s.send(request, context)

	return err
}

func (s *httpManipulator) Increment(context *manipulate.Context, identity elemental.Identity, counter string, inc int) error {

	return manipulate.NewErrNotImplemented("Increment method not implemented in http manipulator")
}

func (s *httpManipulator) makeAuthorizationHeaders() string {

	return s.username + " " + s.password
}

func (s *httpManipulator) prepareHeaders(request *http.Request, context *manipulate.Context) {

	if context.Namespace == "" {
		context.Namespace = s.namespace
	}

	if context.Namespace != "" {
		request.Header.Set("X-Namespace", context.Namespace)
	}

	if s.username != "" && s.password != "" {
		request.Header.Set("Authorization", s.makeAuthorizationHeaders())
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	return
}

func (s *httpManipulator) readHeaders(response *http.Response, context *manipulate.Context) {

	if context == nil {
		return
	}

	context.CountTotal, _ = strconv.Atoi(response.Header.Get("X-Count-Total"))
}

func (s *httpManipulator) getGeneralURL(o elemental.Identifiable) string {

	return s.url + "/" + o.Identity().Category
}

func (s *httpManipulator) getPersonalURL(o elemental.Identifiable) (string, error) {

	if o.Identifier() == "" {
		return "", fmt.Errorf("Cannot GetPersonalURL of an object with no ID set")
	}

	return s.getGeneralURL(o) + "/" + o.Identifier(), nil
}

func (s *httpManipulator) getURLForChildrenIdentity(parent elemental.Identifiable, childrenIdentity elemental.Identity) (string, error) {

	if parent == nil {
		return s.url + "/" + childrenIdentity.Category, nil
	}

	url, err := s.getPersonalURL(parent)
	if err != nil {
		return "", err
	}

	return url + "/" + childrenIdentity.Category, nil
}

func (s *httpManipulator) send(request *http.Request, context *manipulate.Context) (*http.Response, error) {

	s.prepareHeaders(request, context)

	response, err := s.client.Do(request)

	if err != nil {
		log.WithFields(logrus.Fields{
			"method":   request.Method,
			"url":      request.URL,
			"request":  request,
			"response": response,
			"error":    err.Error(),
		}).Debug("Unable to send the request.")
		return response, manipulate.NewErrCannotCommunicate(err.Error())
	}

	log.WithFields(logrus.Fields{
		"method":   request.Method,
		"url":      request.URL,
		"request":  request,
		"response": response,
	}).Debug("Request sent.")

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
