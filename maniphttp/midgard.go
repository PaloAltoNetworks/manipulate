package maniphttp

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	midgardclient "github.com/aporeto-inc/midgard/client"
)

func renewMidgardToken(
	midgardClient *midgardclient.Client,
	manipulator *httpManipulator,
	CAPool *x509.CertPool,
	certificates []tls.Certificate,
	refreshInterval time.Duration,
	stopCh chan bool,
) {
	for {
		select {
		case <-time.Tick(refreshInterval):
			fmt.Println("Refreshing token...")
			token, err := midgardClient.IssueFromCertificate(certificates, CAPool)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Error("Unable to renew token.")
			}
			manipulator.password = token
		case <-stopCh:
			return
		}
	}
}
