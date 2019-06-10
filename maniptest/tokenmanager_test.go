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

package maniptest

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTestSubscriber_MockIssue(t *testing.T) {

	Convey("Given I have TestTokenManager", t, func() {

		tm := NewTestTokenManager()

		Convey("When I call Issue without mock", func() {

			token, err := tm.Issue(context.Background())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then token should be correct", func() {
				So(token, ShouldBeEmpty)
			})
		})

		Convey("When I call Issue with a mock", func() {

			tm.MockIssue(t, func(ctx context.Context) (string, error) {
				return "token", fmt.Errorf("test")
			})

			token, err := tm.Issue(context.Background())

			Convey("Then err should be correct", func() {
				So(err.Error(), ShouldEqual, "test")
			})

			Convey("Then token should be correct", func() {
				So(token, ShouldEqual, "token")
			})
		})
	})
}

func TestTestSubscriber_MockRun(t *testing.T) {

	Convey("Given I have TestTokenManager", t, func() {

		tm := NewTestTokenManager()

		Convey("When I call Run without mock", func() {

			tm.Run(context.Background(), nil)

			Convey("Then nothing should happen", func() {
			})
		})

		Convey("When I call Issue with a mock", func() {

			tm.MockRun(t, func(ctx context.Context, tokenCh chan string) {
				go func() { tokenCh <- "coucou" }()
			})

			ch := make(chan string)
			tm.Run(context.Background(), ch)

			Convey("Then the value in the chan should be correct", func() {
				So(<-ch, ShouldEqual, "coucou")
			})
		})
	})
}
