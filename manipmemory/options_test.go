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

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_newConfig(t *testing.T) {

	Convey("Given call newConfig", t, func() {

		c := newConfig()

		Convey("Then I should get the default config", func() {
			So(c.noCopy, ShouldBeFalse)
		})
	})
}

func Test_Options(t *testing.T) {

	Convey("Calling OptionCredentials should work", t, func() {
		c := newConfig()
		OptionNoCopy(true)(c)
		So(c.noCopy, ShouldBeTrue)
	})
}
