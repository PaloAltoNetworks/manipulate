package maniphttp

import (
	"context"
	"crypto/tls"

	"go.aporeto.io/manipulate"
)

// NewHTTPManipulator returns a Manipulator backed by an ReST API.
func NewHTTPManipulator(url, username, password, namespace string) manipulate.Manipulator {

	// fmt.Println("Deprecated: maniphttp.NewHTTPManipulator. Please switch to maniphttp.New")
	m, err := New(context.Background(), url, OptCredentials(username, password), OptNamespace(namespace))
	if err != nil {
		panic(err)
	}

	return m
}

// NewHTTPManipulatorWithTLS returns a Manipulator backed by an ReST API using the tls config.
func NewHTTPManipulatorWithTLS(url, username, password, namespace string, tlsConfig *tls.Config) manipulate.Manipulator {

	// fmt.Println("Deprecated: maniphttp.NewHTTPManipulatorWithTLS. Please switch to maniphttp.New")
	m, err := New(context.Background(), url, OptCredentials(username, password), OptTLSConfig(tlsConfig))
	if err != nil {
		panic(err)
	}

	return m
}

// NewHTTPManipulatorWithTokenManager returns a http backed manipulate.Manipulatorusing the given manipulate.TokenManager to manage the the token.
func NewHTTPManipulatorWithTokenManager(ctx context.Context, url string, namespace string, tlsConfig *tls.Config, tokenManager manipulate.TokenManager) (manipulate.Manipulator, error) {

	// fmt.Println("Deprecated: maniphttp.NewHTTPManipulatorWithTokenManager. Please switch to maniphttp.New")
	return New(ctx, url, OptNamespace(namespace), OptTLSConfig(tlsConfig), OptTokenManager(tokenManager))
}
