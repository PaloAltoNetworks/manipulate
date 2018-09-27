package maniphttp

import (
	"crypto/tls"
	"net/http"

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
