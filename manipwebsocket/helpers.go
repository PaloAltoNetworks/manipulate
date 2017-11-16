package manipwebsocket

import (
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
)

// SendRequest sends the given request with the given manipulator
func SendRequest(manipulator manipulate.Manipulator, request *elemental.Request) (*elemental.Response, error) {

	m, ok := manipulator.(*websocketManipulator)
	if !ok {
		panic("You can only pass a Websocket Manipulator to SendRequest")
	}

	// @TODO: make this configurable. I have no idea who is using this function.
	return m.send(request, manipulate.NewContext().Timeout)
}

// IsConnected checks the connection state of the manipulator.
func IsConnected(manipulator manipulate.Manipulator) bool {

	m, ok := manipulator.(*websocketManipulator)
	if !ok {
		return false
	}

	if !m.isConnected() {
		return false
	}

	m.wsLock.Lock()
	defer m.wsLock.Unlock()

	return m.ws != nil
}

// ExtractCredentials extracts the username and password from the given manipulator.
// Note: the given manipulator must be an WebSocket Manipulator or it will return an error.
func ExtractCredentials(manipulator manipulate.Manipulator) (string, string, error) {

	m, ok := manipulator.(*websocketManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractCredentials")
	}

	m.renewLock.Lock()
	u, p := m.username, m.password
	m.renewLock.Unlock()

	return u, p, nil
}
