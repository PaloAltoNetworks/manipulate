package maniphttp

import (
	"crypto/tls"
	"net/http"

	"go.aporeto.io/manipulate"
)

// An Option represents a maniphttp.Manipulator option.
type Option func(*httpManipulator)

// OptCredentials sets the username and password to use for authentication.
func OptCredentials(username, password string) Option {
	return func(m *httpManipulator) {
		m.username = username
		m.password = password
	}
}

// OptToken sets JWT token. If you use for authentication.
//
// If you also use OptCredentials or OptTokenManager, the last one will take precendence.
func OptToken(token string) Option {
	return func(m *httpManipulator) {
		m.username = "Bearer"
		m.password = token
	}
}

// OptTokenManager sets manipulate.TokenManager to handle token auto renewal.
//
// If you also use OptCredentials or OptToken, the last one will take precendence.
func OptTokenManager(tokenManager manipulate.TokenManager) Option {
	return func(m *httpManipulator) {
		m.tokenManager = tokenManager
	}
}

// OptNamespace sets the namespace.
func OptNamespace(ns string) Option {
	return func(m *httpManipulator) {
		m.namespace = ns
	}
}

// OptAdditonalHeaders sets the additional http.Header that will be sent.
func OptAdditonalHeaders(headers http.Header) Option {
	return func(m *httpManipulator) {
		m.globalHeaders = headers
	}
}

// OptHTTPClient sets internal full *http.Client.
//
// If you use this option you are responsible for configuring the *http.Transport
// and transport's *tls.Config). OptHTTPTransport or OptTLSConfig
// will have no effect if you use this option.
func OptHTTPClient(client *http.Client) Option {
	return func(m *httpManipulator) {
		m.client = client
	}
}

// OptHTTPTransport sets internal *http.Transport.
//
// If you use this option you are responsible for configuring the *tls.Config.
// OptTLSConfig will have no effect if you use this option.
func OptHTTPTransport(transport *http.Transport) Option {
	return func(m *httpManipulator) {
		m.transport = transport
	}
}

// OptTLSConfig sets the tls.Config to use for the manipulator.
func OptTLSConfig(tlsConfig *tls.Config) Option {
	return func(m *httpManipulator) {
		m.tlsConfig = tlsConfig
	}
}
