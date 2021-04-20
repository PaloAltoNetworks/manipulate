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

// This package provides type mapping for backward compatilility
// as manipulate.Filter moved to elemental.

// Filter is an alias of elemental.Filter
type Filter = elemental.Filter

// FilterKeyComposer is an alias of elemental.FilterKeyComposer
type FilterKeyComposer = elemental.FilterKeyComposer

// FilterParser is an alias of elemental.FilterParser
type FilterParser = elemental.FilterParser

// NewFilter returns a new Filter using the aliased type.
//
// Deprecated: manipulate.NewFilter is deprecated and aliased to elemental.NewFilter.
func NewFilter() *Filter {
	fmt.Println("DEPRECATED: manipulate.NewFilter is deprecated and aliased to elemental.NewFilter")
	return elemental.NewFilter()
}

// NewFilterComposer returns a new FilterKeyComposer using the aliased type.
//
// Deprecated: manipulate.NewFilterComposer is deprecated and aliased to elemental.NewFilterComposer.
func NewFilterComposer() FilterKeyComposer {
	fmt.Println("DEPRECATED: manipulate.NewFilterComposer is deprecated and aliased to elemental.NewFilterComposer")
	return elemental.NewFilter()
}

// NewFilterFromString returns a new NewFilterFromString using the aliased type.
//
// Deprecated: manipulate.NewFilterFromString is deprecated and aliased to elemental.NewFilterFromString.
func NewFilterFromString(filter string) (*Filter, error) {
	fmt.Println("DEPRECATED: manipulate.NewFilterFromString is deprecated and aliased to elemental.NewFilterFromString")
	return elemental.NewFilterFromString(filter)
}

// NewFilterParser returns a new NewFilterParser using the aliased type.
//
// Deprecated: manipulate.NewFilterParser is deprecated and aliased to elemental.NewFilterParser.
func NewFilterParser(input string) *FilterParser {
	fmt.Println("DEPRECATED: manipulate.NewFilterParser is deprecated and aliased to elemental.NewFilterParser")
	return elemental.NewFilterParser(input)
}
