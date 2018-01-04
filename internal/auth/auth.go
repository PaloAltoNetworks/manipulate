package auth

import (
	"crypto/tls"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"

	"github.com/aporeto-inc/manipulate"
	midgardclient "github.com/aporeto-inc/midgard-lib/client"
	opentracing "github.com/opentracing/opentracing-go"
)

// RenewMidgardToken renews the midgard token using the given parameters.
func RenewMidgardToken(retriever manipulate.TokenRetrieveManipulator, stop chan bool) {

	nextRefresh := time.Now().Add(retriever.Validity() / 2)

	for {

		select {
		case <-time.After(time.Minute):

			now := time.Now()
			if now.Before(nextRefresh) {
				break
			}

			if err := retriever.RetrieveToken(); err != nil {
				zap.L().Error("Unable to renew Midgard token", zap.Error(err))
				break
			}

			nextRefresh = now.Add(retriever.Validity() / 2)
			zap.L().Info("Midgard token renewed")

		case <-stop:
			return
		}
	}
}

// IssueInitialToken issues the initial token. It will retr
func IssueInitialToken(mclient *midgardclient.Client, certificates []tls.Certificate, validity time.Duration, sp opentracing.Span) (string, error) {

	var token string
	var err error

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for i := 0; i < 12; i++ {

		token, err = mclient.IssueFromCertificateWithValidity(certificates, validity, sp)
		if err == nil {
			zap.L().Debug("Initial Midgard token issued")
			return token, nil
		}

		zap.L().Warn("Could not access midgard to issue initial token. Retrying in 5s", zap.Error(err))

		select {
		case <-time.After(5 * time.Second):
		case <-c:
			return "", manipulate.NewErrDisconnected("Disconnected per signal")
		}

	}

	return "", err
}
