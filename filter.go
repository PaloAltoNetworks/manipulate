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

package manipulate

import (
	"fmt"

	"go.aporeto.io/elemental"
)

// NewFiltersFromQueryParameters returns the filters matching any `q` parameters.
func NewFiltersFromQueryParameters(parameters elemental.Parameters) ([]*elemental.Filter, error) {

	filters := []*elemental.Filter{}

	for _, query := range parameters.Get("q").StringValues() {

		f, err := elemental.NewFilterFromString(query)
		if err != nil {
			return nil, fmt.Errorf("unable to parse filter in query parameter: %w", err)
		}

		filters = append(filters, f.Done())
	}

	return filters, nil
}

// NewNamespaceFilter returns a manipulate filter used to create the namespace filter.
func NewNamespaceFilter(namespace string, recursive bool) *elemental.Filter {

	return NewNamespaceFilterWithCustomProperty("namespace", namespace, recursive)
}

// NewNamespaceFilterWithCustomProperty allows to create a namespace filter based on a property that
// is different from `namespace`.
func NewNamespaceFilterWithCustomProperty(propertyName string, namespace string, recursive bool) *elemental.Filter {

	if namespace == "" {
		namespace = "/"
	}

	if !recursive {
		return elemental.NewFilterComposer().WithKey(propertyName).Equals(namespace).Done()
	}

	if namespace == "/" {
		return elemental.NewFilterComposer().WithKey(propertyName).Matches("^/").Done()
	}

	return elemental.NewFilterComposer().Or(
		elemental.NewFilterComposer().
			WithKey(propertyName).Equals(namespace).
			Done(),
		elemental.NewFilterComposer().
			WithKey(propertyName).Matches("^"+namespace+"/").
			Done(),
	).Done()
}
