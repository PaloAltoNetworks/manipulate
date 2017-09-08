package manipwebsocket

import (
	"fmt"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/sec"
	"github.com/aporeto-inc/manipulate/internal/sharedcompiler"
)

func decodeErrors(response *elemental.Response) error {

	es := []elemental.Error{}

	if err := response.Decode(&es); err != nil {
		return manipulate.NewErrCannotUnmarshal(err.Error())
	}

	errs := elemental.NewErrors()
	for _, e := range es {
		errs = append(errs, e)
	}

	return errs
}

func handleCommunicationError(m *websocketManipulator, err error) error {

	if _, ok := err.(manipulate.ErrDisconnected); ok {
		return err
	}

	if !m.isConnected() {
		return manipulate.NewErrDisconnected("disconnected per user request")
	}

	return manipulate.NewErrCannotCommunicate(sec.Snip(err, m.currentPassword()).Error())
}

func populateRequestFromContext(request *elemental.Request, ctx *manipulate.Context, o interface{}) error {

	if ctx.Filter != nil {
		var err error
		request.Parameters, err = sharedcompiler.CompileFilter(ctx.Filter)
		if err != nil {
			return err
		}
	}

	if ctx.Parameters != nil {
		for k, v := range ctx.Parameters.KeyValues {
			request.Parameters[k] = v
		}
	}

	if ctx.Parent != nil {
		request.ParentIdentity = ctx.Parent.Identity()
		request.ParentID = ctx.Parent.Identifier()
	}

	if ctx.Namespace != "" {
		request.Namespace = ctx.Namespace
	}

	if ctx.Recursive {
		request.Recursive = true
	}

	if ctx.Version == 0 {
		if v, ok := o.(elemental.Versionable); ok {
			request.Version = int(v.Version())
		}
	}

	request.ExternalTrackingID = ctx.ExternalTrackingID
	request.ExternalTrackingType = ctx.ExternalTrackingType
	request.Page = ctx.Page
	request.PageSize = ctx.PageSize
	request.OverrideProtection = ctx.OverrideProtection
	request.Order = append([]string{}, ctx.Order...)

	return nil
}

// SendRequest sends the given request with the given manipulator
func SendRequest(manipulator manipulate.Manipulator, request *elemental.Request) (*elemental.Response, error) {

	m, ok := manipulator.(*websocketManipulator)
	if !ok {
		return nil, fmt.Errorf("You can only pass a Websocket Manipulator to SendRequest")
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
