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

package snip

import (
	"errors"
	"testing"

	// nolint:revive // Allow dot imports for readability in tests
	. "github.com/smartystreets/goconvey/convey"
)

func TestTokenUtils_Snip(t *testing.T) {

	Convey("Given have a token and and error containing the token", t, func() {

		token := "token"
		err := errors.New("your token is token")

		Convey("When I call Snip", func() {

			e := Snip(err, token)

			Convey("Then err should have the reference to token snipped", func() {
				So(e.Error(), ShouldEqual, "your [snip] is [snip]")
			})
		})
	})

	Convey("Given have a token and and error that doesn't contain the token", t, func() {

		token := "token"
		err := errors.New("your secret is secret")

		Convey("When I call Snip", func() {

			e := Snip(err, token)

			Convey("Then err should have the reference to token snipped", func() {
				So(e.Error(), ShouldEqual, "your secret is secret")
			})
		})
	})

	Convey("Given I have a token and a nil error", t, func() {

		token := "token"

		Convey("When I call Snip", func() {

			e := Snip(nil, token)

			Convey("Then err should be nil", func() {
				So(e, ShouldBeNil)
			})
		})
	})
}
