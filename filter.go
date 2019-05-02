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

type Filter = elemental.Filter
type FilterKeyComposer = elemental.FilterKeyComposer
type FilterParser = elemental.FilterParser

func NewFilter() *Filter {
	fmt.Println("DEPRECATED: manipulate.NewFilter is deprecated and aliased to elemental.NewFilter")
	return elemental.NewFilter()
}

func NewFilterComposer() FilterKeyComposer {
	fmt.Println("DEPRECATED: manipulate.NewFilterComposer is deprecated and aliased to elemental.NewFilterComposer")
	return elemental.NewFilter()
}

func NewFilterFromString(filter string) (*Filter, error) {
	fmt.Println("DEPRECATED: manipulate.NewFilterFromString is deprecated and aliased to elemental.NewFilterFromString")
	return elemental.NewFilterFromString(filter)
}

func NewFilterParser(input string) *FilterParser {
	fmt.Println("DEPRECATED: manipulate.NewFilterParser is deprecated and aliased to elemental.NewFilterParser")
	return elemental.NewFilterParser(input)
}
