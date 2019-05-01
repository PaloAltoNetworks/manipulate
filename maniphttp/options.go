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
	"crypto/tls"
	"net/http"

	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

// An Option represents a maniphttp.Manipulator option.
type Option func(*httpManipulator)

// OptionCredentials sets the username and password to use for authentication.
func OptionCredentials(username, password string) Option {
	return func(m *httpManipulator) {
		m.username = username
		m.password = password
	}
}

// OptionToken sets JWT token. If you use for authentication.
//
// If you also use OptionCredentials or OptionTokenManager, the last one will take precedence.
func OptionToken(token string) Option {
	return func(m *httpManipulator) {
		m.username = "Bearer"
		m.password = token
	}
}

// OptionTokenManager sets manipulate.TokenManager to handle token auto renewal.
//
// If you also use OptionCredentials or OptionToken, the last one will take precedence.
func OptionTokenManager(tokenManager manipulate.TokenManager) Option {
	return func(m *httpManipulator) {
		m.tokenManager = tokenManager
	}
}

// OptionNamespace sets the namespace.
func OptionNamespace(ns string) Option {
	return func(m *httpManipulator) {
		m.namespace = ns
	}
}

// OptionAdditonalHeaders sets the additional http.Header that will be sent.
func OptionAdditonalHeaders(headers http.Header) Option {
	return func(m *httpManipulator) {
		m.globalHeaders = headers
	}
}

// OptionHTTPClient sets internal full *http.Client.
//
// If you use this option you are responsible for configuring the *http.Transport
// and transport's *tls.Config). OptionHTTPTransport or OptionTLSConfig
// will have no effect if you use this option.
func OptionHTTPClient(client *http.Client) Option {
	return func(m *httpManipulator) {
		m.client = client
	}
}

// OptionHTTPTransport sets internal *http.Transport.
//
// If you use this option you are responsible for configuring the *tls.Config.
// OptionTLSConfig will have no effect if you use this option.
func OptionHTTPTransport(transport *http.Transport) Option {
	return func(m *httpManipulator) {
		m.transport = transport
	}
}

// OptionTLSConfig sets the tls.Config to use for the manipulator.
func OptionTLSConfig(tlsConfig *tls.Config) Option {
	return func(m *httpManipulator) {
		m.tlsConfig = tlsConfig
	}
}

// OptionDisableBuiltInRetry disables the auto retry mechanism
// built in maniphttp Manipulator.
// By default, the manipulator will silently retry on communication
// error 3 times after 1s, 2s, and 3s.
func OptionDisableBuiltInRetry() Option {
	return func(m *httpManipulator) {
		m.disableAutoRetry = true
	}
}

// OptionEncoding sets the encoding/decoding type to use.
func OptionEncoding(enc elemental.EncodingType) Option {
	return func(m *httpManipulator) {
		m.encoding = enc
	}
}
