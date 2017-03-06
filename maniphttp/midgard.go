package maniphttp

import (
	"crypto/tls"
	"time"

	midgardclient "github.com/aporeto-inc/midgard-lib/client"
)

func renewMidgardToken(
	midgardClient *midgardclient.Client,
	manipulator *httpManipulator,
	certificates []tls.Certificate,
	refreshInterval time.Duration,
	stopCh chan bool,
) {
	nextRefresh := time.Now().Add(refreshInterval)

	for {
		select {
		case <-time.Tick(time.Minute):

			now := time.Now()
			if now.Before(nextRefresh) {
				continue
			}

			log.Info("Refreshing Midgard token...")
			token, err := midgardClient.IssueFromCertificate(certificates)
			if err != nil {
				log.WithError(err).Error("Unable to renew token.")
				break
			}

			// TODO: there is a race condition here. It cannot happen, but it should be fixed.
			manipulator.password = token
			nextRefresh = now.Add(refreshInterval)

		case <-stopCh:
			return
		}
	}
}
