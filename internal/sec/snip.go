package sec

import (
	"fmt"
	"strings"
)

// Snip snips the given token from the given error.
func Snip(err error, token string) error {
	return fmt.Errorf("%s",
		strings.Replace(
			err.Error(),
			token,
			"[snip]",
			-1),
	)
}
