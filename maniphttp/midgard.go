package maniphttp

import (
	"crypto/tls"
	"crypto/x509"
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
			log.Info("Refreshing Midgard token...")
			token, err := midgardClient.IssueFromCertificate(certificates)
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
