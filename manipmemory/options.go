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

package manipmemory

// An Option represents a maniphttp.Manipulator option.
type Option func(*config)

type config struct {
	noCopy bool
}

func newConfig() *config {
	return &config{}
}

// OptionNoCopy tells the manipulator to store the data
// as is without copying it. This is faster, but unsafe
// as pointers are stored as is, allowing random
// modifications. If you use this option, you must
// make sure you are not modifying the object your store
// or retrieve.
func OptionNoCopy(noCopy bool) Option {
	return func(c *config) {
		c.noCopy = noCopy
	}
}
