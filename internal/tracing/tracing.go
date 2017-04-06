package tracing

import (
	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// StartTrace starts a new trace from the root span if any.
func StartTrace(rootSpan opentracing.Span, name string, mctx *manipulate.Context) opentracing.Span {

	if rootSpan == nil {
		return nil
	}

	sp := opentracing.StartSpan(name, opentracing.ChildOf(rootSpan.Context()))

	if mctx != nil {
		sp.SetTag("manipulate.context.api_version", mctx.Version)
		sp.SetTag("manipulate.context.page", mctx.Page)
		sp.SetTag("manipulate.context.page_size", mctx.PageSize)
		sp.SetTag("manipulate.context.override_protection", mctx.OverrideProtection)
		sp.SetTag("manipulate.context.recursive", mctx.Recursive)
		sp.SetTag("manipulate.context.namespace", mctx.Namespace)

		if len(mctx.Parameters.KeyValues) >= 0 {
			sp.SetTag("manipulate.context.parameters", mctx.Parameters.KeyValues)
		}

		if mctx.Filter != nil {
			sp.SetTag("manipulate.context.filter", mctx.Filter.String())
		}
	}

	return sp
}

// SetTag sets a tag to the given span
func SetTag(span opentracing.Span, key string, value interface{}) {

	if span == nil {
		return
	}

	span.SetTag(key, value)
}

// FinishTrace finish the given span if any.
func FinishTrace(span opentracing.Span) {

	if span == nil {
		return
	}

	span.Finish()
}

// FinishTraceWithError finish the given span if any as an error.
func FinishTraceWithError(span opentracing.Span, err error) {

	if span == nil {
		return
	}

	ext.Error.Set(span, true)
	span.LogEventWithPayload("Error", err.Error())
	span.Finish()
}

// Inject injects the span info into the given elemental.Request
func Inject(span opentracing.Span, request *elemental.Request) error {

	if span == nil {
		return nil
	}

	tracer := span.Tracer()

	if tracer == nil {
		return nil
	}

	return tracer.Inject(span.Context(), opentracing.TextMap, request.TrackingData)
}