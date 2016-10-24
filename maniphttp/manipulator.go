// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package maniphttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"

	log "github.com/Sirupsen/logrus"
)

type httpManipulator struct {
	username  string
	password  string
	url       string
	namespace string
	client    *http.Client
}

// NewHTTPManipulator returns a Manipulator backed by an ReST API.
func NewHTTPManipulator(username, password, url, namespace string, tlsConfig *TLSConfiguration) manipulate.Manipulator {

	if tlsConfig == nil {
		tlsConfig = NewTLSConfiguration("", "", "", false)
	}

	client, err := tlsConfig.makeHTTPClient()
	if err != nil {
		panic(fmt.Sprintf("Invalid TLSConfiguration: %s", err))
	}

	return &httpManipulator{
		username:  username,
		password:  password,
		url:       url,
		client:    client,
		namespace: namespace,
	}
}

func (s *httpManipulator) RetrieveMany(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	url, err := s.getURLForChildrenIdentity(context.Parent, identity)
	if err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotBuildQuery)
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
	}

	response, err := s.send(request, context)
	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusNoContent || response.ContentLength == 0 {
		return nil
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if err := json.NewDecoder(response.Body).Decode(&dest); err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotUnmarshal)
	}

	return nil
}

func (s *httpManipulator) Retrieve(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	for _, object := range objects {

		url, err := s.getPersonalURL(object)
		if err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotBuildQuery)
		}

		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
		}

		response, err := s.send(request, context)
		if err != nil {
			return err
		}

		defer func() {
			if err := response.Body.Close(); err != nil {
				panic(err)
			}
		}()

		if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotUnmarshal)
		}
	}

	return nil
}

func (s *httpManipulator) Create(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	for _, child := range objects {

		url, err := s.getURLForChildrenIdentity(context.Parent, child.Identity())
		if err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotBuildQuery)
		}

		buffer := &bytes.Buffer{}
		if err1 := json.NewEncoder(buffer).Encode(&child); err1 != nil {
			return manipulate.NewError(err1.Error(), manipulate.ErrCannotMarshal)
		}

		request, err := http.NewRequest(http.MethodPost, url, buffer)
		if err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
		}

		response, err := s.send(request, context)
		if err != nil {
			return err
		}

		defer func() {
			if err := response.Body.Close(); err != nil {
				panic(err)
			}
		}()

		if err := json.NewDecoder(response.Body).Decode(&child); err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotUnmarshal)
		}
	}

	return nil
}

func (s *httpManipulator) Update(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	for _, object := range objects {

		url, err := s.getPersonalURL(object)
		if err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotBuildQuery)
		}

		buffer := &bytes.Buffer{}
		if err1 := json.NewEncoder(buffer).Encode(object); err1 != nil {
			return manipulate.NewError(err1.Error(), manipulate.ErrCannotMarshal)
		}

		request, err := http.NewRequest(http.MethodPut, url, buffer)
		if err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
		}

		response, err := s.send(request, context)
		if err != nil {
			return err
		}

		defer func() {
			if err := response.Body.Close(); err != nil {
				panic(err)
			}
		}()

		if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotUnmarshal)
		}
	}

	return nil
}

func (s *httpManipulator) Delete(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	for _, object := range objects {

		url, err := s.getPersonalURL(object)
		if err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotBuildQuery)
		}

		request, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
		}

		_, err = s.send(request, context)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *httpManipulator) Count(*manipulate.Context, elemental.Identity) (int, error) {
	return -1, manipulate.NewError("Count is not implemented in HTTPStore", manipulate.ErrNotImplemented)
}

func (s *httpManipulator) Assign(context *manipulate.Context, assignation *elemental.Assignation) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	url, err := s.getURLForChildrenIdentity(context.Parent, assignation.MembersIdentity)
	if err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotBuildQuery)
	}

	buffer := &bytes.Buffer{}
	if err1 := json.NewEncoder(buffer).Encode(assignation); err1 != nil {
		return manipulate.NewError(err1.Error(), manipulate.ErrCannotMarshal)
	}

	request, err := http.NewRequest(http.MethodPatch, url, buffer)
	if err != nil {
		return manipulate.NewError(err.Error(), manipulate.ErrCannotExecuteQuery)
	}

	_, err = s.send(request, nil)

	return err
}

func (s *httpManipulator) Increment(context *manipulate.Context, identity elemental.Identity, counter string, inc int) error {

	return manipulate.NewError("Increment is not implemented in HTTPStore", manipulate.ErrNotImplemented)
}

func (s *httpManipulator) makeAuthorizationHeaders() string {

	return s.username + " " + s.password
}

func (s *httpManipulator) prepareHeaders(request *http.Request, context *manipulate.Context) {

	if s.namespace != "" {
		request.Header.Set("X-Namespace", s.namespace)
	}

	if s.username != "" && s.password != "" {
		request.Header.Set("Authorization", s.makeAuthorizationHeaders())
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if context == nil {
		return
	}

	if context.PageSize != -1 {
		request.Header.Set("X-Page-Size", strconv.Itoa(context.PageSize))
	}

	if context.PageCurrent != -1 {
		request.Header.Set("X-Page-Current", strconv.Itoa(context.PageCurrent))
	}

	if context.PageSize > 0 {
		request.Header.Set("X-Page-Size", strconv.Itoa(context.PageSize))
	}

	return
}

func (s *httpManipulator) readHeaders(response *http.Response, context *manipulate.Context) {

	if context == nil {
		return
	}

	context.PageCurrent, _ = strconv.Atoi(response.Header.Get("X-Page-Current"))
	context.PageSize, _ = strconv.Atoi(response.Header.Get("X-Page-Size"))
	context.PageFirst = response.Header.Get("X-Page-First")
	context.PagePrev = response.Header.Get("X-Page-Prev")
	context.PageNext = response.Header.Get("X-Page-Next")
	context.PageLast = response.Header.Get("X-Page-Last")

	context.CountLocal, _ = strconv.Atoi(response.Header.Get("X-Count-Local"))
	context.CountTotal, _ = strconv.Atoi(response.Header.Get("X-Count-Total"))
}

func (s *httpManipulator) getGeneralURL(o manipulate.Manipulable) string {

	return s.url + "/" + o.Identity().Category
}

func (s *httpManipulator) getPersonalURL(o manipulate.Manipulable) (string, error) {

	if o.Identifier() == "" {
		return "", fmt.Errorf("Cannot GetPersonalURL of an object with no ID set")
	}

	return s.getGeneralURL(o) + "/" + o.Identifier(), nil
}

func (s *httpManipulator) getURLForChildrenIdentity(parent manipulate.Manipulable, childrenIdentity elemental.Identity) (string, error) {

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
		log.WithFields(log.Fields{
			"package":  "maniphttp",
			"method":   request.Method,
			"url":      request.URL,
			"request":  request,
			"response": response,
			"error":    err,
		}).Debug("Unable to send the request.")
		return response, manipulate.NewError(err.Error(), manipulate.ErrCannotCommunicate)
	}

	log.WithFields(log.Fields{
		"package":  "maniphttp",
		"method":   request.Method,
		"url":      request.URL,
		"request":  request,
		"response": response,
	}).Debug("Request sent.")

	if response.StatusCode < 200 || response.StatusCode >= 300 {

		errs := elemental.Errors{}

		if err := json.NewDecoder(response.Body).Decode(&errs); err != nil {
			return nil, err
		}

		return response, errs
	}

	s.readHeaders(response, context)

	return response, nil
}
