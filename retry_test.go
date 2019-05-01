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
	"errors"
	"testing"
	"time"

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

	Convey("Given I have a context and a manipulate function that returns an error", t, func() {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		m := func() error {
			return NewErrCannotCommit("nope")
		}

		Convey("When I call Retry", func() {

			err := Retry(ctx, m, nil)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(IsCannotCommitError(err), ShouldBeTrue)
			})
		})
	})

	Convey("Given I have a context and a manipulate function that returns an communication error and retry func that stops the retrying", t, func() {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var tryN int
		r := func(t int, e error) error {
			tryN = t
			if t == 2 {
				return errors.New("stop that")
			}
			return nil
		}

		m := func() error {
			return NewErrCannotCommunicate("where are you?")
		}

		Convey("When I call Retry", func() {

			err := Retry(ctx, m, r)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(IsCannotCommitError(err), ShouldBeFalse)
				So(err.Error(), ShouldEqual, "stop that")
			})

			Convey("Then it should have retryied once", func() {
				So(tryN, ShouldEqual, 2)
			})
		})
	})

	Convey("Given I have a context and a manipulate function that returns 2 communication error before being ok", t, func() {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var tryN int
		r := func(t int, e error) error {
			tryN = t
			return nil
		}

		m := func() error {
			if tryN < 2 {
				return NewErrCannotCommunicate("where are you?")
			}

			return nil
		}

		Convey("When I call Retry", func() {

			err := Retry(ctx, m, r)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then it should have retryied once", func() {
				So(tryN, ShouldEqual, 2)
			})
		})
	})

	Convey("Given I have a context and a manipulate function that returns communication errors until context is canceled", t, func() {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		m := func() error {
			return NewErrCannotCommunicate("where are you?")
		}

		Convey("When I call Retry", func() {

			go func() {
				<-time.After(300 * time.Millisecond)
				cancel()
			}()

			err := Retry(ctx, m, nil)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Disconnected: interupted by context: context canceled. original error: Cannot communicate: where are you?")
			})

		})
	})
}
