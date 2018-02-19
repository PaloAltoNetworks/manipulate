package tracing

import (
	"context"

	"github.com/aporeto-inc/manipulate"
	"github.com/opentracing/opentracing-go"
)

// StartTrace starts a new trace from the root span if any.
func StartTrace(mctx *manipulate.Context, name string) opentracing.Span {

	if mctx == nil {
		sp, _ := opentracing.StartSpanFromContext(context.Background(), name)
		return sp
	}

	sp, _ := opentracing.StartSpanFromContext(mctx, name)

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

// // InjectInElementalRequest injects the span info into the given elemental.Request.
// func InjectInElementalRequest(span opentracing.Span, request *elemental.Request) error {

// 	if span == nil {
// 		return nil
// 	}

// 	tracer := span.Tracer()

// 	if tracer == nil {
// 		return nil
// 	}

// 	return tracer.Inject(span.Context(), opentracing.TextMap, request.TrackingData)
// }

// // InjectInHTTPRequest injects the span info into the given http.Request.
// func InjectInHTTPRequest(span opentracing.Span, request *http.Request) error {

// 	if span == nil {
// 		return nil
// 	}

// 	tracer := span.Tracer()

// 	if tracer == nil {
// 		return nil
// 	}

// 	return tracer.Inject(span.Context(), opentracing.TextMap, opentracing.HTTPHeadersCarrier(request.Header))
// }
