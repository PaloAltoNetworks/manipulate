package maniphttp

import (
	"crypto/tls"
	"time"

	log "github.com/Sirupsen/logrus"
	midgardclient "github.com/aporeto-inc/midgard-lib/client"
)

func renewMidgardToken(
	midgardClient *midgardclient.Client,
	manipulator *httpManipulator,
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
			// TODO: there is a race condition here. It cannot happen, but it should be fixed.
			manipulator.password = token
		case <-stopCh:
			return
		}
	}
}
