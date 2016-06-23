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
)

// HTTPStore represents a user session.
type HTTPStore struct {
	certificate string
	username    string
	password    string
	url         string
	namespace   string
	client      *http.Client
}

// NewHTTPStore returns a new *HTTPStore
func NewHTTPStore(username, password, url, namespace string, tlsConfig *TLSConfiguration) *HTTPStore {

	if tlsConfig == nil {
		tlsConfig = NewTLSConfiguration("", "", false)
	}

	client, err := tlsConfig.makeHTTPClient()
	if err != nil {
		panic(fmt.Sprintf("Invalid TLSConfiguration: %s", err))
	}

	return &HTTPStore{
		username:  username,
		password:  password,
		url:       url,
		client:    client,
		namespace: namespace,
	}
}

func (s *HTTPStore) makeAuthorizationHeaders() string {

	return "TODO"
}

func (s *HTTPStore) prepareHeaders(request *http.Request, context *manipulate.Context) *elemental.Error {

	if s.namespace != "" {
		request.Header.Set("X-Namespace", s.namespace)
	}

	request.Header.Set("Authorization", s.makeAuthorizationHeaders())
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

func (s *HTTPStore) readHeaders(response *http.Response, context *manipulate.Context) {

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

func (s *HTTPStore) getGeneralURL(o manipulate.Manipulable) string {

	return s.url + "/" + o.Identity().Category
}

func (s *HTTPStore) getPersonalURL(o manipulate.Manipulable) (string, *elemental.Error) {

	if o.Identifier() == "" {
		return "", elemental.NewError("URL Computing Error", "Cannot GetPersonalURL of an object with no ID set", "manipulate", 2)
	}

	return s.getGeneralURL(o) + "/" + o.Identifier(), nil
}

func (s *HTTPStore) getURLForChildrenIdentity(parent manipulate.Manipulable, childrenIdentity elemental.Identity) (string, *elemental.Error) {

	if parent == nil {
		return s.url + "/" + childrenIdentity.Category, nil
	}

	url, err := s.getPersonalURL(parent)
	if err != nil {
		return "", err
	}

	return url + "/" + childrenIdentity.Category, nil
}

func (s *HTTPStore) send(request *http.Request, context *manipulate.Context) (*http.Response, elemental.Errors) {

	s.prepareHeaders(request, context)

	response, err := s.client.Do(request)

	if err != nil {
		return response, elemental.NewErrors(elemental.NewError("Error while sending the request", err.Error(), "manipulate", 0))
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {

		var errs elemental.Errors

		if err := json.NewDecoder(response.Body).Decode(&errs); err != nil {
			errs = elemental.NewErrors(elemental.NewError("No data", "gen", "manipulate", 1))
		}

		return response, errs
	}

	s.readHeaders(response, context)
	return response, nil
}

// Create creates a new child Identifiable under the given parent Identifiable in the server.
func (s *HTTPStore) Create(contexts manipulate.Contexts, parent manipulate.Manipulable, children ...manipulate.Manipulable) elemental.Errors {

	errs := elemental.Errors{}

	// very stupid implementation for now
	for index, child := range children {

		url, berr := s.getURLForChildrenIdentity(parent, child.Identity())
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

		response, berrs := s.send(request, manipulate.ContextForIndex(contexts, index))

		if berrs != nil {
			errs = append(errs, berrs...)
			continue
		}

		defer response.Body.Close()
		if err := json.NewDecoder(response.Body).Decode(&child); err != nil {
			errs = append(errs, elemental.NewError("Cannot Decode JSON", err.Error(), "manipulate", 0))
			continue
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

// Retrieve fetchs the given Identifiable from the server.
func (s *HTTPStore) Retrieve(contexts manipulate.Contexts, objects ...manipulate.Manipulable) elemental.Errors {

	errs := elemental.Errors{}

	// very stupid implementation for now
	for index, object := range objects {

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

		response, berrs := s.send(request, manipulate.ContextForIndex(contexts, index))

		if berrs != nil {
			errs = append(errs, berrs...)
			continue
		}

		defer response.Body.Close()
		if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
			errs = append(errs, elemental.NewError("Cannot Decode JSON", err.Error(), "manipulate", 0))
			continue
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

// Update saves the given Identifiable into the server.
func (s *HTTPStore) Update(contexts manipulate.Contexts, objects ...manipulate.Manipulable) elemental.Errors {

	errs := elemental.Errors{}

	// very stupid implementation for now
	for index, object := range objects {

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

		response, berrs := s.send(request, manipulate.ContextForIndex(contexts, index))

		if berrs != nil {
			errs = append(errs, berrs...)
			continue
		}

		defer response.Body.Close()
		if err := json.NewDecoder(response.Body).Decode(&object); err != nil {
			errs = append(errs, elemental.NewError("Cannot Decode JSON", err.Error(), "manipulate", 0))
			continue
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

// Delete deletes the given Identifiable from the server.
func (s *HTTPStore) Delete(contexts manipulate.Contexts, objects ...manipulate.Manipulable) elemental.Errors {

	errs := elemental.Errors{}

	// very stupid implementation for now
	for index, object := range objects {

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

		_, berrs := s.send(request, manipulate.ContextForIndex(contexts, index))
		if berrs != nil {
			errs = append(errs, berrs...)
			continue
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

// RetrieveChildren fetches the children with of given parent identified by the given Identity.
func (s *HTTPStore) RetrieveChildren(contexts manipulate.Contexts, parent manipulate.Manipulable, identity elemental.Identity, dest interface{}) elemental.Errors {

	url, berr := s.getURLForChildrenIdentity(parent, identity)
	if berr != nil {
		return elemental.NewErrors(berr)
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		elemental.NewErrors(elemental.NewError("Bad Request", err.Error(), "manipulate", http.StatusBadRequest))
	}

	response, berrs := s.send(request, manipulate.ContextForIndex(contexts, 0))

	if berrs != nil {
		return berrs
	}

	if response.StatusCode == http.StatusNoContent || response.ContentLength == 0 {
		return nil
	}

	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&dest); err != nil {
		return elemental.NewErrors(elemental.NewError("Cannot Decode JSON", err.Error(), "manipulate", 0))
	}

	return nil
}

// Count count the number of element with the given Identity.
// Not implemented yet
func (s *HTTPStore) Count(manipulate.Contexts, elemental.Identity) (int, elemental.Errors) {
	return 0, nil
}

// Assign assigns the list of given child Identifiables to the given Identifiable parent in the server.
func (s *HTTPStore) Assign(contexts manipulate.Contexts, parent manipulate.Manipulable, assignation *elemental.Assignation) elemental.Errors {

	url, berr := s.getURLForChildrenIdentity(parent, assignation.MembersIdentity)
	if berr != nil {
		return elemental.NewErrors(berr)
	}

	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(assignation); err != nil {
		return elemental.NewErrors(elemental.NewError("Unable Encode", err.Error(), "manipulate", 0))
	}

	request, err := http.NewRequest("PATCH", url, buffer)
	if err != nil {
		return elemental.NewErrors(elemental.NewError("Bad Request", err.Error(), "manipulate", http.StatusBadRequest))
	}

	_, berrs := s.send(request, nil)

	if berrs != nil {
		return berrs
	}

	return nil
}

//
// // NextEvent will return the next notification from the backend as it occurs and will
// // send it to the correct channel.
// func (s *HTTPSeal) NextEvent(channel NotificationsChannel, lastEventID string) *Error {
//
// 	currentURL := s.URL + "/events"
// 	if lastEventID != "" {
// 		currentURL += "?uuid=" + lastEventID
// 	}
//
// 	request, err := http.NewRequest("GET", currentURL, nil)
// 	if err != nil {
// 		return NewError(http.StatusBadRequest, err.Error())
// 	}
//
// 	response, berr := s.send(request, nil)
// 	if berr != nil {
// 		return berr
// 	}
//
// 	notification := NewNotification()
// 	if err := json.NewDecoder(response.Body).Decode(notification); err != nil {
// 		return NewError(ErrorCodeJSONCannotDecode, err.Error())
// 	}
//
// 	if len(notification.Events) > 0 {
// 		channel <- notification
// 	}
//
// 	return nil
// }
