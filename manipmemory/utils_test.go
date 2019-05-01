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

func Test_boolIndex(t *testing.T) {

	type testObject struct {
		somevalue      bool
		someothervalue string
	}

	Convey("When I call boolindex", t, func() {

		Convey("When I use a good data structure", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			val, err := boolIndex(t, "somevalue")
			So(err, ShouldBeNil)
			So(val, ShouldBeTrue)
		})

		Convey("When I use a good data structure with a bad field", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			_, err := boolIndex(t, "no value")
			So(err, ShouldNotBeNil)
		})

		Convey("When I use nil", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			_, err := boolIndex(t, "no value")
			So(err, ShouldNotBeNil)
		})

		Convey("When I use a bad type field", func() {
			t := &testObject{
				somevalue:      true,
				someothervalue: "somestring",
			}

			_, err := boolIndex(t, "somestring")
			So(err, ShouldNotBeNil)
		})
	})
}
