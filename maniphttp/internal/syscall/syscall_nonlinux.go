// +build !linux

package syscall

import (
	"syscall"
	"time"
)

// package to set to the low-level/OS settings

// SetTCPUserTimeout sets the TCP timeout for a socket connection
func SetTCPUserTimeout(d time.Duration) func(string, string, syscall.RawConn) error {
	return nil
}
