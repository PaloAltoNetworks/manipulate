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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
)

func TestNewFilter(t *testing.T) {

	Convey("Calling NewFilter should work", t, func() {
		f := NewFilter()
		So(f, ShouldHaveSameTypeAs, elemental.NewFilter())
	})

	Convey("Calling NewFilterComposer should work", t, func() {
		f := NewFilterComposer()
		So(f, ShouldHaveSameTypeAs, elemental.NewFilterComposer())
	})

	Convey("Calling NewFilterFromString should work", t, func() {
		f, _ := NewFilterFromString("a == a")
		ef, _ := elemental.NewFilterFromString("a == a")
		So(f, ShouldHaveSameTypeAs, ef)
	})

	Convey("Calling NewFilterParser should work", t, func() {
		f := NewFilterParser("a == a")
		So(f, ShouldHaveSameTypeAs, elemental.NewFilterParser("a == a"))
	})
}
