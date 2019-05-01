// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tracing

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
	"go.aporeto.io/manipulate"
)

// StartTrace starts a new trace from the root span if any.
func StartTrace(mctx manipulate.Context, name string) opentracing.Span {

	if mctx == nil {
		sp, _ := opentracing.StartSpanFromContext(context.Background(), name)
		return sp
	}

	sp, _ := opentracing.StartSpanFromContext(mctx.Context(), name)

	sp.SetTag("manipulate.context.api_version", mctx.Version())
	sp.SetTag("manipulate.context.page", mctx.Page())
	sp.SetTag("manipulate.context.page_size", mctx.PageSize())
	sp.SetTag("manipulate.context.override_protection", mctx.Override())
	sp.SetTag("manipulate.context.recursive", mctx.Recursive())

	if mctx.Namespace() != "" {
		sp.SetTag("manipulate.context.namespace", mctx.Namespace())
	} else {
		sp.SetTag("manipulate.context.namespace", "manipulator-default")
	}

	if len(mctx.Parameters()) >= 0 {
		sp.SetTag("manipulate.context.parameters", mctx.Parameters())
	}

	if mctx.Filter() != nil {
		sp.SetTag("manipulate.context.filter", mctx.Filter().String())
	}

	return sp
}
