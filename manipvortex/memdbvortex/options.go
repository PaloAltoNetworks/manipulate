package memdbvortex

import (
	"go.aporeto.io/manipulate"
)

// Option represents an option can can be passed to NewContext.
type Option func(*MemDBVortex)

// OptionBackendManipulator sets the backend manipulator.
func OptionBackendManipulator(manipulator manipulate.Manipulator) Option {
	return func(m *MemDBVortex) {
		m.m = manipulator
	}
}

// OptionBackendSubscriber sets the backend subscriber.
func OptionBackendSubscriber(s manipulate.Subscriber) Option {
	return func(m *MemDBVortex) {
		m.s = s
	}
}

// OptionTransactionLog sets the transaction log file.
func OptionTransactionLog(filename string) Option {
	return func(m *MemDBVortex) {
		m.logfile = filename
		m.enableLog = filename != ""
	}
}

// OptionTransactionQueueLength sets the queue length of the
// transaction queue.
func OptionTransactionQueueLength(n int) Option {
	return func(m *MemDBVortex) {
		m.transactionQueue = make(chan *Transaction, n)
	}
}
