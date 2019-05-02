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

package compiler

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
)

func TestFilter_CompileFilter(t *testing.T) {

	Convey("Given I create a new Filter", t, func() {

		f := elemental.NewFilterComposer().
			WithKey("name").Equals("thename").
			WithKey("ID").Equals("xxx").
			WithKey("associatedTags").Contains("yy=zz").
			Done()

		Convey("When I call CompileFilter on it", func() {

			v, err := CompileFilter(f)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the filter should be correct", func() {
				So(v.Get("q"), ShouldEqual, `name == "thename" and ID == "xxx" and associatedTags contains ["yy=zz"]`)
			})
		})
	})
}
