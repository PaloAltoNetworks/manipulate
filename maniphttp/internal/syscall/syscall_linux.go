// +build linux

package syscall

// package to set to the low-level/OS settings

import (
	"os"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

const (
	defaultTCPNetworkTimeout = 30 * time.Second
)

// SetTCPUserTimeout sets the TCP timeout for a socket connection
func SetTCPUserTimeout(t time.Duration) func(string, string, syscall.RawConn) error {

	// return if the tcpUserTimeout is not set.
	if t == 0 {
		return nil
	}
	tcpTimeout := defaultTCPNetworkTimeout

	if t > 0 && t != defaultTCPNetworkTimeout {
		tcpTimeout = t
	}

	return func(network, address string, c syscall.RawConn) error {
		var sysErr error
		var err = c.Control(func(fd uintptr) {
			sysErr = syscall.SetsockoptInt(int(fd), syscall.SOL_TCP, unix.TCP_USER_TIMEOUT,
				int(tcpTimeout.Milliseconds()))
		})
		if sysErr != nil {
			return os.NewSyscallError("setsockopt", sysErr)
		}
		return err
	}
}
