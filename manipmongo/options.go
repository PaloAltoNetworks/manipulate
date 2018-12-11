package manipmongo

import (
	"crypto/tls"
	"time"

	"go.aporeto.io/manipulate"
)

// An Option represents a maniphttp.Manipulator option.
type Option func(*config)

type config struct {
	username         string
	password         string
	authsource       string
	tlsConfig        *tls.Config
	poolLimit        int
	connectTimeout   time.Duration
	socketTimeout    time.Duration
	readConsistency  manipulate.ReadConsistency
	writeConsistency manipulate.WriteConsistency
	sharder          Sharder
}

func newConfig() *config {
	return &config{
		poolLimit:        4096,
		connectTimeout:   10 * time.Second,
		socketTimeout:    60 * time.Second,
		readConsistency:  manipulate.ReadConsistencyDefault,
		writeConsistency: manipulate.WriteConsistencyDefault,
	}
}

// OptionCredentials sets the username and password to use for authentication.
func OptionCredentials(username, password, authsource string) Option {
	return func(c *config) {
		c.username = username
		c.password = password
		c.authsource = authsource
	}
}

// OptionTLS sets the tls configuration for the connection.
func OptionTLS(tlsConfig *tls.Config) Option {
	return func(c *config) {
		c.tlsConfig = tlsConfig
	}
}

// OptionConnectionPoolLimit sets maximum size of the connection pool.
func OptionConnectionPoolLimit(poolLimit int) Option {
	return func(c *config) {
		c.poolLimit = poolLimit
	}
}

// OptionConnectionTimeout sets the connection timeout.
func OptionConnectionTimeout(connectTimeout time.Duration) Option {
	return func(c *config) {
		c.connectTimeout = connectTimeout
	}
}

// OptionSocketTimeout sets the socket timeout.
func OptionSocketTimeout(socketTimeout time.Duration) Option {
	return func(c *config) {
		c.socketTimeout = socketTimeout
	}
}

// OptionDefaultReadConsistencyMode sets the default read consistency mode.
func OptionDefaultReadConsistencyMode(consistency manipulate.ReadConsistency) Option {
	return func(c *config) {
		c.readConsistency = consistency
	}
}

// OptionDefaultWriteConsistencyMode sets the default write consistency mode.
func OptionDefaultWriteConsistencyMode(consistency manipulate.WriteConsistency) Option {
	return func(c *config) {
		c.writeConsistency = consistency
	}
}

// OptionSharder sets the sharder.
func OptionSharder(sharder Sharder) Option {
	return func(c *config) {
		c.sharder = sharder
	}
}
