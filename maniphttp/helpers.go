package maniphttp

import (
	"github.com/aporeto-inc/manipulate"
)

// ExtractCredentials extracts the username and password from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractCredentials(manipulator manipulate.Manipulator) (string, string) {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractCredentials")
	}

	m.renewLock.Lock()
	u, p := m.username, m.password
	m.renewLock.Unlock()

	return u, p
}
