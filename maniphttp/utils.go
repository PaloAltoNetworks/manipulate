package maniphttp

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

// NewCertificateAndPool generate everything need for https session
func NewCertificateAndPool(certificatePath string, keyPath string, CAPath string) (*tls.Certificate, *x509.CertPool, error) {

	cert, err := tls.LoadX509KeyPair(certificatePath, keyPath)

	if err != nil {
		return nil, nil, err
	}

	caCert, err := ioutil.ReadFile(CAPath)

	if err != nil {
		return nil, nil, err
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	return &cert, pool, nil
}
