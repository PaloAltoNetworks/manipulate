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

package maniphttp

import (
	"go.aporeto.io/manipulate"
)

// A RetryInfo contains information about a retry,
type RetryInfo struct {
	URL    string
	Method string

	err  error
	try  int
	mctx manipulate.Context
}

// Try returns the try number.
func (i RetryInfo) Try() int {
	return i.try
}

// Err returns the error that caused the retry.
func (i RetryInfo) Err() error {
	return i.err
}

// Context returns the manipulate.Context used.
func (i RetryInfo) Context() manipulate.Context {
	return i.mctx
}
