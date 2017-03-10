package manipwebsocket

import (
	"fmt"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/manipwebsocket/compiler"
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

func populateRequestFromContext(request *elemental.Request, ctx *manipulate.Context) error {

	if ctx.Filter != nil {
		var err error
		request.Parameters, err = compiler.CompileFilter(ctx.Filter)
		if err != nil {
			return err
		}
	}

	if ctx.Parameters != nil {
		for k, v := range ctx.Parameters.KeyValues {
			request.Parameters.Add(k, v)
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

	request.Page = ctx.Page
	request.PageSize = ctx.PageSize
	request.OverrideProtection = ctx.OverrideProtection
	request.Version = ctx.Version

	return nil
}

// SendRequest sends the given request with the given manipulator
func SendRequest(manipulator manipulate.Manipulator, request *elemental.Request) (*elemental.Response, error) {

	m, ok := manipulator.(*websocketManipulator)

	if !ok {
		return nil, fmt.Errorf("You can only pass a Websocket Manipulator to GetWebsocketConn")
	}

	return m.send(request)
}
