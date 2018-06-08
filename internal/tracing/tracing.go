package tracing

import (
	"context"

	"go.aporeto.io/manipulate"
	"github.com/opentracing/opentracing-go"
)

// StartTrace starts a new trace from the root span if any.
func StartTrace(mctx *manipulate.Context, name string) opentracing.Span {

	if mctx == nil {
		sp, _ := opentracing.StartSpanFromContext(context.Background(), name)
		return sp
	}

	sp, _ := opentracing.StartSpanFromContext(mctx.Context(), name)

	sp.SetTag("manipulate.context.api_version", mctx.Version)
	sp.SetTag("manipulate.context.page", mctx.Page)
	sp.SetTag("manipulate.context.page_size", mctx.PageSize)
	sp.SetTag("manipulate.context.override_protection", mctx.OverrideProtection)
	sp.SetTag("manipulate.context.recursive", mctx.Recursive)

	if mctx.Namespace != "" {
		sp.SetTag("manipulate.context.namespace", mctx.Namespace)
	} else {
		sp.SetTag("manipulate.context.namespace", "manipulator-default")
	}

	if len(mctx.Parameters) >= 0 {
		sp.SetTag("manipulate.context.parameters", mctx.Parameters)
	}

	if mctx.Filter != nil {
		sp.SetTag("manipulate.context.filter", mctx.Filter.String())
	}

	return sp
}
