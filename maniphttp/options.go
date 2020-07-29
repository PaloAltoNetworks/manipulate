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
	"time"

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

// OptionDefaultRetryFunc sets the default retry func to use
// if manipulate.Context does not have one.
func OptionDefaultRetryFunc(f manipulate.RetryFunc) Option {
	return func(m *httpManipulator) {
		m.defaultRetryFunc = f
	}
}

// OptionDisableCompression disables the gzip compression
// in http transport. This only has effect if you don't set
// a custom transport.
func OptionDisableCompression() Option {
	return func(m *httpManipulator) {
		m.disableCompression = true
	}
}

// OptionSendCredentialsAsCookie configures the manipulator to
// send the password as a cookie using the provided key.
func OptionSendCredentialsAsCookie(key string) Option {
	return func(m *httpManipulator) {
		m.tokenCookieKey = key
	}
}

// OptionSimulateFailures will inject random error during
// low level communication with the remote API.
//
// The key of the map is a float between 0 and 1 that will
// give the percentage of chance for simulating the failure,
// and error it should return.
//
// For instance, take the following map:
//      map[float64]error{
//          0.10: manipulate.NewErrCannotBuildQuery("oh no"),
//          0.25: manipulate.NewErrCannotCommunicate("service is gone"),
//      }
//
// It will return manipulate.NewErrCannotBuildQuery around 10% of the requests,
// manipulate.NewErrCannotCommunicate around 25% of the requests.
// This is obviously designed for simulating backend failures and should
// not be used in production, obviously.
func OptionSimulateFailures(failureSimulations map[float64]error) Option {
	return func(m *httpManipulator) {
		m.failureSimulations = failureSimulations
	}
}

// OptionTCPUserTimeout configures the manipulator to
// have a custom tcp user timeout.
func OptionTCPUserTimeout(t time.Duration) Option {
	return func(m *httpManipulator) {
		m.tcpUserTimeout = t
	}
}

// OptionBackoffCurve configures the backoff curve
// the manipulator will use when performing internal retry
// operations.
// Default curve is: 0s, 1s, 4s, 10s, 20s, 30s, 60s
func OptionBackoffCurve(curve []time.Duration) Option {
	return func(m *httpManipulator) {
		m.backoffCurve = curve
	}
}

// OptionStrongBackoffCurve configures the strong backoff curve
// the manipulator will use when performing internal retry
// operations that necessitate to wait more than usual like
// a 429 code.
// Default curve is: 10s, 20s, 30s
func OptionStrongBackoffCurve(curve []time.Duration) Option {
	return func(m *httpManipulator) {
		m.strongBackoffCurve = curve
	}
}

var (
	opaqueKeyOverrideHeaderContentType = "maniphttp.opaqueKeyOverrideHeaderContentType"
	opaqueKeyOverrideHeaderAccept      = "maniphttp.opaqueKeyOverrideHeaderAccept"
)

type opaquer interface {
	Opaque() map[string]interface{}
}

// ContextOptionOverrideContentType is an advanced feature that allows you
// to override with the actual Content-Type header value.
// This should not be used in 99% of the case.
func ContextOptionOverrideContentType(encoding string) manipulate.ContextOption {

	return func(c manipulate.Context) {
		c.(opaquer).Opaque()[opaqueKeyOverrideHeaderContentType] = encoding
	}
}

// ContextOptionOverrideAccept is an advanced feature that allows you
// to override with the actual Accept header value.
// This should not be used in 99% of the case.
func ContextOptionOverrideAccept(accept string) manipulate.ContextOption {

	return func(c manipulate.Context) {
		c.(opaquer).Opaque()[opaqueKeyOverrideHeaderAccept] = accept
	}
}
