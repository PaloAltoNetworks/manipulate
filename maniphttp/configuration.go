package maniphttp

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

// TLSConfiguration holds various TLS configuration details
type TLSConfiguration struct {
	CACertificatePath  string
	CertificatePath    string
	CertificateKeyPath string
	SkipInsecure       bool
}

// NewTLSConfiguration returns a new TLSConfiguration
func NewTLSConfiguration(cacert, cert, key string, skip bool) *TLSConfiguration {

	if cacert != "" && (cert == "" || key == "") {
		panic("You must set cacert, cert and key parameters")
	}

	if cert != "" && (cacert == "" || key == "") {
		panic("You must set cacert, cert and key parameters")
	}

	if key != "" && (cacert == "" || cert == "") {
		panic("You must set cacert, cert and key parameters")
	}

	return &TLSConfiguration{
		CACertificatePath:  cacert,
		CertificatePath:    cert,
		CertificateKeyPath: key,
		SkipInsecure:       skip,
	}
}

func (t *TLSConfiguration) makeHTTPClient() (*http.Client, error) {

	if t.CACertificatePath == "" {
		return &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: t.SkipInsecure,
				},
			},
		}, nil
	}

	cert, err := tls.LoadX509KeyPair(t.CertificatePath, t.CertificateKeyPath)
	if err != nil {
		return nil, err
	}

	CACert, err := ioutil.ReadFile(t.CACertificatePath)
	if err != nil {
		return nil, err
	}

	CACertPool := x509.NewCertPool()
	CACertPool.AppendCertsFromPEM(CACert)

	TLSConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      CACertPool,
	}

	TLSConfig.BuildNameToCertificate()

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: t.SkipInsecure,
			},
		},
	}, nil
}
