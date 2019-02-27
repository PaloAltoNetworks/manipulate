package manipvortex

import (
	"time"

	"go.aporeto.io/manipulate"
)

type config struct {
	upstreamManipulator  manipulate.Manipulator
	upstreamSubscriber   manipulate.Subscriber
	logfile              string
	enableLog            bool
	poolLimit            int
	transactionQueue     chan *Transaction
	readConsistency      manipulate.ReadConsistency
	writeConsistency     manipulate.WriteConsistency
	defaultQueueDuration time.Duration
	defaultPageSize      int
}

func newConfig() *config {
	return &config{
		transactionQueue:     make(chan *Transaction, 1000),
		readConsistency:      manipulate.ReadConsistencyEventual,
		writeConsistency:     manipulate.WriteConsistencyStrong,
		defaultQueueDuration: time.Second,
		defaultPageSize:      10000,
	}
}

// Option represents an option can can be passed to NewContext.
type Option func(*config)

// OptionDefaultConsistency sets the default read and write consistency.
func OptionDefaultConsistency(read manipulate.ReadConsistency, write manipulate.WriteConsistency) Option {
	return func(cfg *config) {
		if read != manipulate.ReadConsistencyDefault {
			cfg.readConsistency = read
		}
		if write != manipulate.WriteConsistencyDefault {
			cfg.writeConsistency = write
		}
	}
}

// OptionUpstreamManipulator sets the upstream manipulator.
func OptionUpstreamManipulator(manipulator manipulate.Manipulator) Option {
	return func(cfg *config) {
		cfg.upstreamManipulator = manipulator
	}
}

// OptionUpstreamSubscriber sets the upstream subscriber.
func OptionUpstreamSubscriber(s manipulate.Subscriber) Option {
	return func(cfg *config) {
		cfg.upstreamSubscriber = s
	}
}

// OptionTransactionLog sets the transaction log file.
func OptionTransactionLog(filename string) Option {
	return func(cfg *config) {
		cfg.logfile = filename
		cfg.enableLog = filename != ""
	}
}

// OptionTransactionQueueLength sets the queue length of the
// transaction queue.
func OptionTransactionQueueLength(n int) Option {
	return func(cfg *config) {
		cfg.transactionQueue = make(chan *Transaction, n)
	}
}

// OptionTransactionQueueDuration sets the default queue transaction
// duration. Once expired, the transaction is discarded.
func OptionTransactionQueueDuration(d time.Duration) Option {
	return func(cfg *config) {
		cfg.defaultQueueDuration = d
	}
}

// OptionDefaultPageSize is the page size during fetching.
func OptionDefaultPageSize(defaultPageSize int) Option {
	return func(cfg *config) {
		cfg.defaultPageSize = defaultPageSize
	}
}
