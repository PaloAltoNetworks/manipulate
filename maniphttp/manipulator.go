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

func (s *httpManipulator) Create(context *manipulate.Context, children ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	errs := []error{}

	for _, child := range children {

		url, berr := s.getURLForChildrenIdentity(context.Parent, child.Identity())
		if berr != nil {
			errs = append(errs, berr)
			continue
		}

		buffer := &bytes.Buffer{}
		if err := json.NewEncoder(buffer).Encode(&child); err != nil {
			errs = append(errs, elemental.NewError("Cannot Encode JSON", err.Error(), "manipulate", 0))
			continue
		}

		request, err := http.NewRequest("POST", url, buffer)
		if err != nil {
			errs = append(errs, elemental.NewError("Bad Request", err.Error(), "manipulate", http.StatusBadRequest))
			continue
		}

		response, berrs := s.send(request, context)

		if berrs != nil {
			errs = append(errs, berrs)
			continue
		}

		defer response.Body.Close()
		if err := json.NewDecoder(response.Body).Decode(&child); err != nil {
			errs = append(errs, elemental.NewError("Cannot Decode JSON", err.Error(), "manipulate", 0))
			continue
		}
	}

	if len(errs) > 0 {
		return elemental.NewErrors(errs...)
	}

	return nil
}

func (s *httpManipulator) Retrieve(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	errs := []error{}

	for _, object := range objects {

		url, berr := s.getPersonalURL(object)
		if berr != nil {
			errs = append(errs, berr)
			continue
		}

		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			errs = append(errs, elemental.NewError("Bad Request", err.Error(), "manipulate", http.StatusBadRequest))
			continue
		}

		response, berrs := s.send(request, context)

		if berrs != nil {
			errs = append(errs, berrs)
			continue
		}

		defer response.Body.Close()
		if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
			errs = append(errs, elemental.NewError("Cannot Decode JSON", err.Error(), "manipulate", 0))
			continue
		}
	}

	if len(errs) > 0 {
		return elemental.NewErrors(errs...)
	}

	return nil
}

func (s *httpManipulator) Update(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	errs := []error{}

	for _, object := range objects {

		url, berr := s.getPersonalURL(object)
		if berr != nil {
			errs = append(errs, berr)
			continue
		}

		buffer := &bytes.Buffer{}
		if err := json.NewEncoder(buffer).Encode(object); err != nil {
			errs = append(errs, elemental.NewError("Unable Encode", err.Error(), "manipulate", 0))
			continue
		}

		request, err := http.NewRequest("PUT", url, buffer)
		if err != nil {
			errs = append(errs, elemental.NewError("Bad Request", err.Error(), "manipulate", http.StatusBadRequest))
			continue
		}

		response, berrs := s.send(request, context)

		if berrs != nil {
			errs = append(errs, berrs)
			continue
		}

		defer response.Body.Close()
		if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
			errs = append(errs, elemental.NewError("Cannot Decode JSON", err.Error(), "manipulate", 0))
			continue
		}
	}

	if len(errs) > 0 {
		return elemental.NewErrors(errs...)
	}

	return nil
}

func (s *httpManipulator) Delete(context *manipulate.Context, objects ...manipulate.Manipulable) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	errs := []error{}

	for _, object := range objects {

		url, berr := s.getPersonalURL(object)
		if berr != nil {
			errs = append(errs, berr)
			continue
		}

		request, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			errs = append(errs, elemental.NewError("Bad Request", err.Error(), "manipulate", http.StatusBadRequest))
			continue
		}

		_, berrs := s.send(request, context)
		if berrs != nil {
			errs = append(errs, berrs)
			continue
		}
	}

	if len(errs) > 0 {
		return elemental.NewErrors(errs...)
	}

	return nil
}

func (s *httpManipulator) RetrieveMany(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	url, err := s.getURLForChildrenIdentity(context.Parent, identity)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return elemental.NewError("Bad Request", err.Error(), "manipulate", http.StatusBadRequest)
	}

	response, berrs := s.send(request, context)

	if berrs != nil {
		return berrs
	}

	if response.StatusCode == http.StatusNoContent || response.ContentLength == 0 {
		return nil
	}

	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&dest); err != nil {
		return elemental.NewError("Cannot Decode JSON", err.Error(), "manipulate", 0)
	}

	return nil
}

func (s *httpManipulator) Count(*manipulate.Context, elemental.Identity) (int, error) {
	return 0, nil
}

func (s *httpManipulator) Assign(context *manipulate.Context, assignation *elemental.Assignation) error {

	if context == nil {
		context = manipulate.NewContext()
	}

	url, berr := s.getURLForChildrenIdentity(context.Parent, assignation.MembersIdentity)
	if berr != nil {
		return berr
	}

	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(assignation); err != nil {
		return elemental.NewError("Unable Encode", err.Error(), "manipulate", 0)
	}

	request, err := http.NewRequest("PATCH", url, buffer)
	if err != nil {
		return elemental.NewError("Bad Request", err.Error(), "manipulate", http.StatusBadRequest)
	}

	_, berrs := s.send(request, nil)

	if berrs != nil {
		return berrs
	}

	return nil
}

func (s *httpManipulator) Increment(context *manipulate.Context, name string, counter string, inc int, filterKeys []string, filterValues []interface{}) error {
	return fmt.Errorf("Increment is not implemented in http store")
}

func (s *httpManipulator) makeAuthorizationHeaders() string {

	return s.username + " " + s.password
}

func (s *httpManipulator) prepareHeaders(request *http.Request, context *manipulate.Context) error {

	if s.namespace != "" {
		request.Header.Set("X-Namespace", s.namespace)
	}

	if s.username != "" && s.password != "" {
		request.Header.Set("Authorization", s.makeAuthorizationHeaders())
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if context == nil {
		return nil
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

	return nil
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
		return "", elemental.NewError("URL Computing Error", "Cannot GetPersonalURL of an object with no ID set", "manipulate", 2)
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
		return response, elemental.NewError("Error while sending the request", err.Error(), "manipulate", 0)
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
			return nil, elemental.NewError("No data", "gen", "manipulate", 1)
		}

		return response, errs
	}

	s.readHeaders(response, context)

	return response, nil
}
