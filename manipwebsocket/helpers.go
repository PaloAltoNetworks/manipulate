package manipwebsocket

import (
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
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

func populateRequestFromContext(request *elemental.Request, ctx *manipulate.Context) {

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
}
