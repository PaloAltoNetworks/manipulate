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
	"context"
	"testing"

	// nolint:revive // Allow dot imports for readability in tests
	. "github.com/smartystreets/goconvey/convey"
)

func TestManipulate_Retry(t *testing.T) {

	Convey("Given I have a context and a manipulate function that returns no error", t, func() {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		m := func() error {
			return nil
		}

		Convey("When I call Retry", func() {

			err := Retry(ctx, m, nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
