package maniphttp

import (
	"crypto/tls"

	"go.aporeto.io/manipulate"
)

// ExtractCredentials extracts the username and password from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractCredentials(manipulator manipulate.Manipulator) (string, string) {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractCredentials")
	}

	m.renewLock.Lock()
	u, p := m.username, m.password
	m.renewLock.Unlock()

	return u, p
}

// ExtractEndpoint extracts the endpoint url from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractEndpoint(manipulator manipulate.Manipulator) string {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractEndpoint")
	}

	return m.url
}

// ExtractTLSConfig extracts the tls config from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractTLSConfig(manipulator manipulate.Manipulator) *tls.Config {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractEndpoint")
	}

	return m.tlsConfig
}
