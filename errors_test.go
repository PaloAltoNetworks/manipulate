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
	"errors"
	"fmt"
	"testing"

	// nolint:revive // Allow dot imports for readability in tests
	. "github.com/smartystreets/goconvey/convey"
)

func genericErrorTest(
	t *testing.T,
	errorPrefix string,
	makeFactory func(error) error,
	makeFactoryOld func(string) error,
	verifierFunc func(error) bool,
) {

	Convey("When I create an error", t, func() {

		oerr := fmt.Errorf("this is a an error")
		err := makeFactory(oerr)

		Convey("Then it should be correct", func() {
			So(err.Error(), ShouldEqual, errorPrefix+"this is a an error")
			So(errors.Is(err, oerr), ShouldBeTrue)
			So(verifierFunc(err), ShouldBeTrue)
		})

		olderr := makeFactoryOld("this is a an error")
		Convey("Then the old error should be correct", func() {
			So(olderr.Error(), ShouldEqual, errorPrefix+"this is a an error")
			So(verifierFunc(err), ShouldBeTrue)
		})
	})
}

func TestThing_Function(t *testing.T) {

	genericErrorTest(
		t,
		"Unable to unmarshal data: ",
		func(err error) error { return ErrCannotUnmarshal{Err: err} },
		func(msg string) error { return NewErrCannotUnmarshal(msg) },
		IsCannotUnmarshalError,
	)

	genericErrorTest(
		t,
		"Unable to marshal data: ",
		func(err error) error { return ErrCannotMarshal{Err: err} },
		func(msg string) error { return NewErrCannotMarshal(msg) },
		IsCannotMarshalError,
	)

	genericErrorTest(
		t,
		"Object not found: ",
		func(err error) error { return ErrObjectNotFound{Err: err} },
		func(msg string) error { return NewErrObjectNotFound(msg) },
		IsObjectNotFoundError,
	)

	genericErrorTest(
		t,
		"Multiple objects found: ",
		func(err error) error { return ErrMultipleObjectsFound{Err: err} },
		func(msg string) error { return NewErrMultipleObjectsFound(msg) },
		IsMultipleObjectsFoundError,
	)

	genericErrorTest(
		t,
		"Unable to build query: ",
		func(err error) error { return ErrCannotBuildQuery{Err: err} },
		func(msg string) error { return NewErrCannotBuildQuery(msg) },
		IsCannotBuildQueryError,
	)

	genericErrorTest(
		t,
		"Unable to execute query: ",
		func(err error) error { return ErrCannotExecuteQuery{Err: err} },
		func(msg string) error { return NewErrCannotExecuteQuery(msg) },
		IsCannotExecuteQueryError,
	)

	genericErrorTest(
		t,
		"Unable to commit transaction: ",
		func(err error) error { return ErrCannotCommit{Err: err} },
		func(msg string) error { return NewErrCannotCommit(msg) },
		IsCannotCommitError,
	)

	genericErrorTest(
		t,
		"Not implemented: ",
		func(err error) error { return ErrNotImplemented{Err: err} },
		func(msg string) error { return NewErrNotImplemented(msg) },
		IsNotImplementedError,
	)

	genericErrorTest(
		t,
		"Cannot communicate: ",
		func(err error) error { return ErrCannotCommunicate{Err: err} },
		func(msg string) error { return NewErrCannotCommunicate(msg) },
		IsCannotCommunicateError,
	)

	genericErrorTest(
		t,
		"Cannot communicate: ",
		func(err error) error { return ErrLocked{Err: err} },
		func(msg string) error { return NewErrLocked(msg) },
		IsLockedError,
	)

	genericErrorTest(
		t,
		"Transaction not found: ",
		func(err error) error { return ErrTransactionNotFound{Err: err} },
		func(msg string) error { return NewErrTransactionNotFound(msg) },
		IsTransactionNotFoundError,
	)

	genericErrorTest(
		t,
		"Constraint violation: ",
		func(err error) error { return ErrConstraintViolation{Err: err} },
		func(msg string) error { return NewErrConstraintViolation(msg) },
		IsConstraintViolationError,
	)

	genericErrorTest(
		t,
		"Disconnected: ",
		func(err error) error { return ErrDisconnected{Err: err} },
		func(msg string) error { return NewErrDisconnected(msg) },
		IsDisconnectedError,
	)

	genericErrorTest(
		t,
		"Too many requests: ",
		func(err error) error { return ErrTooManyRequests{Err: err} },
		func(msg string) error { return NewErrTooManyRequests(msg) },
		IsTooManyRequestsError,
	)

	genericErrorTest(
		t,
		"TLS error: ",
		func(err error) error { return ErrTLS{Err: err} },
		func(msg string) error { return NewErrTLS(msg) },
		IsTLSError,
	)
}
