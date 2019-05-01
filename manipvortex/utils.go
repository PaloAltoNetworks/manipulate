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

package manipvortex

import (
	"go.aporeto.io/manipulate"
)

func isStrongReadConsistency(mctx manipulate.Context, processor *Processor, defaultConsistency manipulate.ReadConsistency) bool {

	if mctx != nil && mctx.ReadConsistency() != manipulate.ReadConsistencyDefault {
		return mctx.ReadConsistency() == manipulate.ReadConsistencyStrong
	}

	if processor != nil && processor.ReadConsistency != manipulate.ReadConsistencyDefault {
		return processor.ReadConsistency == manipulate.ReadConsistencyStrong
	}

	return defaultConsistency == manipulate.ReadConsistencyStrong
}

func isStrongWriteConsistency(mctx manipulate.Context, processor *Processor, defaultConsistency manipulate.WriteConsistency) bool {

	if mctx != nil && mctx.WriteConsistency() != manipulate.WriteConsistencyDefault {
		return mctx.WriteConsistency() == manipulate.WriteConsistencyStrong || mctx.WriteConsistency() == manipulate.WriteConsistencyStrongest
	}

	if processor != nil && processor.WriteConsistency != manipulate.WriteConsistencyDefault {
		return processor.WriteConsistency == manipulate.WriteConsistencyStrong || processor.WriteConsistency == manipulate.WriteConsistencyStrongest
	}

	return defaultConsistency == manipulate.WriteConsistencyStrong || defaultConsistency == manipulate.WriteConsistencyStrongest
}
