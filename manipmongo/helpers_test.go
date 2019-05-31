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

package manipmongo

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

func TestRunQuery(t *testing.T) {

	testIdentity := elemental.MakeIdentity("test", "tests")

	Convey("Given I have query function that works", t, func() {

		var try int
		var lastErr error
		var imctx *manipulate.Context

		f := func() (interface{}, error) { return "hello", nil }
		rf := func(i manipulate.RetryInfo) error {
			try = i.Try()
			lastErr = i.Err()
			m := i.Context()
			if m != nil {
				imctx = &m
			}
			return nil
		}

		Convey("When I call RunQuery", func() {

			mctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionRetryFunc(rf),
			)

			out, err := RunQuery(
				mctx,
				f,
				RetryInfo{
					Operation:        elemental.OperationCreate, // we miss DeleteMany
					Identity:         testIdentity,
					defaultRetryFunc: nil,
				},
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldResemble, "hello")
			})

			Convey("Then try should be correct", func() {
				So(try, ShouldEqual, 0)
			})

			Convey("Then lastErr should be correct", func() {
				So(lastErr, ShouldBeNil)
			})

			Convey("Then imctx should be correct", func() {
				So(imctx, ShouldBeNil)
			})
		})
	})

	Convey("Given I have query function that return an non comm error", t, func() {

		var try int
		var lastErr error
		var imctx *manipulate.Context

		rf := func(i manipulate.RetryInfo) error {
			try = i.Try()
			lastErr = i.Err()
			m := i.Context()
			if m != nil {
				imctx = &m
			}
			return nil
		}

		f := func() (interface{}, error) { return nil, fmt.Errorf("boom") }

		Convey("When I call RunQuery", func() {

			out, err := RunQuery(
				manipulate.NewContext(
					context.Background(),
					manipulate.ContextOptionRetryFunc(rf),
				),
				f,
				RetryInfo{
					Operation:        elemental.OperationCreate, // we miss DeleteMany
					Identity:         testIdentity,
					defaultRetryFunc: nil,
				},
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Unable to execute query: boom")
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldBeNil)
			})

			Convey("Then try should be correct", func() {
				So(try, ShouldEqual, 0)
			})

			Convey("Then lastErr should be correct", func() {
				So(lastErr, ShouldBeNil)
			})

			Convey("Then imctx should be correct", func() {
				So(imctx, ShouldBeNil)
			})
		})
	})

	Convey("Given I have query function that returns a net.Error and works at second try", t, func() {

		var try int
		var lastErr error
		var operation elemental.Operation
		var identity elemental.Identity
		var imctx *manipulate.Context

		rf := func(i manipulate.RetryInfo) error {
			try = i.Try()
			lastErr = i.Err()
			m := i.Context()
			if m != nil {
				imctx = &m
			}

			operation = i.(RetryInfo).Operation
			identity = i.(RetryInfo).Identity

			return nil
		}

		f := func() (interface{}, error) {
			if try == 2 {
				return "hello", nil
			}
			return nil, &net.OpError{Err: fmt.Errorf("hello")}
		}

		Convey("When I call RunQuery", func() {

			mctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionRetryFunc(rf),
			)

			out, err := RunQuery(
				mctx,
				f,
				RetryInfo{
					Operation:        elemental.OperationCreate, // we miss DeleteMany
					Identity:         testIdentity,
					defaultRetryFunc: nil,
				},
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldResemble, "hello")
			})

			Convey("Then try should be correct", func() {
				So(try, ShouldEqual, 2)
			})

			Convey("Then lastErr should be correct", func() {
				So(lastErr, ShouldNotBeNil)
				So(lastErr.Error(), ShouldEqual, "Cannot communicate: : hello")
			})

			Convey("Then imctx should be correct", func() {
				So(imctx, ShouldNotBeNil)
				So(*imctx, ShouldEqual, mctx)
			})

			Convey("Then operation should be correct", func() {
				So(operation, ShouldEqual, elemental.OperationCreate)
			})

			Convey("Then identity should be correct", func() {
				So(identity.IsEqual(testIdentity), ShouldBeTrue)
			})
		})
	})

	Convey("Given I have query function that returns a net.Error and and a retry func that returns an error", t, func() {

		f := func() (interface{}, error) {
			return nil, &net.OpError{Err: fmt.Errorf("hello")}
		}

		rf := func(i manipulate.RetryInfo) error { return fmt.Errorf("non: %s", i.Err().Error()) }

		Convey("When I call RunQuery", func() {

			out, err := RunQuery(
				manipulate.NewContext(
					context.Background(),
					manipulate.ContextOptionRetryFunc(rf),
				),
				f,
				RetryInfo{
					Operation:        elemental.OperationCreate, // we miss DeleteMany
					Identity:         testIdentity,
					defaultRetryFunc: nil,
				},
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "non: Cannot communicate: : hello")
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldBeNil)
			})
		})
	})

	Convey("Given I have query function that returns a net.Error and a retry func and a default retry func", t, func() {

		f := func() (interface{}, error) {
			return nil, &net.OpError{Err: fmt.Errorf("hello")}
		}

		rf := func(i manipulate.RetryInfo) error { return fmt.Errorf("non: %s", i.Err().Error()) }
		df := func(i manipulate.RetryInfo) error { return fmt.Errorf("oui: %s", i.Err().Error()) }

		Convey("When I call RunQuery", func() {

			out, err := RunQuery(
				manipulate.NewContext(
					context.Background(),
					manipulate.ContextOptionRetryFunc(rf),
				),
				f,
				RetryInfo{
					Operation:        elemental.OperationCreate, // we miss DeleteMany
					Identity:         testIdentity,
					defaultRetryFunc: df,
				},
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "non: Cannot communicate: : hello")
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldBeNil)
			})
		})
	})

	Convey("Given I have query function that returns a net.Error and a default retry func", t, func() {

		f := func() (interface{}, error) {
			return nil, &net.OpError{Err: fmt.Errorf("hello")}
		}

		df := func(i manipulate.RetryInfo) error { return fmt.Errorf("oui: %s", i.Err().Error()) }

		Convey("When I call RunQuery", func() {

			out, err := RunQuery(
				manipulate.NewContext(
					context.Background(),
				),
				f,
				RetryInfo{
					Operation:        elemental.OperationCreate, // we miss DeleteMany
					Identity:         testIdentity,
					defaultRetryFunc: df,
				},
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "oui: Cannot communicate: : hello")
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldBeNil)
			})
		})
	})

	Convey("Given I have query function that returns a net.Error and never works", t, func() {

		f := func() (interface{}, error) {
			return nil, &net.OpError{Err: fmt.Errorf("hello")}
		}

		rf := func(i manipulate.RetryInfo) error { return nil }

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		Convey("When I call RunQuery", func() {

			out, err := RunQuery(
				manipulate.NewContext(
					ctx,
					manipulate.ContextOptionRetryFunc(rf),
				),
				f,
				RetryInfo{
					Operation:        elemental.OperationCreate, // we miss DeleteMany
					Identity:         testIdentity,
					defaultRetryFunc: nil,
				},
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Unable to execute query: context deadline exceeded")
			})

			Convey("Then out should be correct", func() {
				So(out, ShouldBeNil)
			})
		})
	})
}
