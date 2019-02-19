package manipvortex

import (
	"go.aporeto.io/manipulate"
)

// Option represents an option can can be passed to NewContext.
type Option func(*vortexManipulator)

// OptionUpstreamManipulator sets the upstream manipulator.
func OptionUpstreamManipulator(manipulator manipulate.Manipulator) Option {
	return func(m *vortexManipulator) {
		m.upstreamManipulator = manipulator
	}
}

// OptionUpstreamSubscriber sets the upstream subscriber.
func OptionUpstreamSubscriber(s manipulate.Subscriber) Option {
	return func(m *vortexManipulator) {
		m.upstreamSubscriber = s
	}
}

// OptionTransactionLog sets the transaction log file.
func OptionTransactionLog(filename string) Option {
	return func(m *vortexManipulator) {
		m.logfile = filename
		m.enableLog = filename != ""
	}
}

// OptionTransactionQueueLength sets the queue length of the
// transaction queue.
func OptionTransactionQueueLength(n int) Option {
	return func(m *vortexManipulator) {
		m.transactionQueue = make(chan *Transaction, n)
	}
}
