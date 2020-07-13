// +build linux

package syscall

// package to set to the low-level/OS settings

import (
	"os"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

// MakeDialerControlFunc creates a custom control for the dailer
func MakeDialerControlFunc(t time.Duration) func(string, string, syscall.RawConn) error {
	// return if the tcpUserTimeout is not set.
	if t == 0 {
		return nil
	}

	return func(network, address string, c syscall.RawConn) error {
		var sysErr error
		err := c.Control(func(fd uintptr) {
			sysErr = syscall.SetsockoptInt(int(fd), syscall.SOL_TCP, unix.TCP_USER_TIMEOUT,
				int(t.Milliseconds()))
		})
		if sysErr != nil {
			return os.NewSyscallError("setsockopt", sysErr)
		}
		return err
	}
}
