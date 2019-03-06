package manipvortex

import (
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

// isCommonIdentity will validate that all objects in the operation have the same identity.
// We do not support calls with different identities.
func isCommonIdentity(objects ...elemental.Identifiable) bool {

	identites := map[string]struct{}{}

	for _, obj := range objects {
		if obj == nil {
			continue
		}

		identites[obj.Identity().Name] = struct{}{}
	}

	return len(identites) <= 1
}

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
