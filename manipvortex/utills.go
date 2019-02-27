package manipvortex

import "go.aporeto.io/elemental"

// isCommonIdentity will validate that all objects in the operation have the same identity.
// We do not support calls with different identities.
func isCommonIdentity(objects ...elemental.Identifiable) bool {
	if len(objects) == 0 {
		return false
	}

	first := objects[0].Identity()
	for _, obj := range objects {
		if !first.IsEqual(obj.Identity()) {
			return false
		}
	}

	return true
}
