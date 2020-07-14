// +build !linux

package syscall

import (
	"syscall"
	"time"
)

// package to set to the low-level/OS settings

// MakeDialerControlFunc creates a custom control for the dailer
func MakeDialerControlFunc(d time.Duration) func(string, string, syscall.RawConn) error {
	return nil
}
