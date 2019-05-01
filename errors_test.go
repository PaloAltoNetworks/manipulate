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
)

func genericErrorTest(
	t *testing.T,
	errorPrefix string,
	makeFactory func(string) error,
	verifierFunc func(error) bool,
) {

	Convey("When I create an error", t, func() {

		err := makeFactory("this is a an error")

		Convey("Then it should be correct", func() {
			So(err.Error(), ShouldEqual, errorPrefix+"this is a an error")
			So(verifierFunc(err), ShouldBeTrue)
		})
	})
}

func TestThing_Function(t *testing.T) {

	genericErrorTest(
		t,
		"Unable to unmarshal data: ",
		func(text string) error { return NewErrCannotUnmarshal(text) },
		IsCannotUnmarshalError,
	)

	genericErrorTest(
		t,
		"Unable to marshal data: ",
		func(text string) error { return NewErrCannotMarshal(text) },
		IsCannotMarshalError,
	)

	genericErrorTest(
		t,
		"Object not found: ",
		func(text string) error { return NewErrObjectNotFound(text) },
		IsObjectNotFoundError,
	)

	genericErrorTest(
		t,
		"Multiple objects found: ",
		func(text string) error { return NewErrMultipleObjectsFound(text) },
		IsMultipleObjectsFoundError,
	)

	genericErrorTest(
		t,
		"Unable to build query: ",
		func(text string) error { return NewErrCannotBuildQuery(text) },
		IsCannotBuildQueryError,
	)

	genericErrorTest(
		t,
		"Unable to execute query: ",
		func(text string) error { return NewErrCannotExecuteQuery(text) },
		IsCannotExecuteQueryError,
	)

	genericErrorTest(
		t,
		"Unable to commit transaction: ",
		func(text string) error { return NewErrCannotCommit(text) },
		IsCannotCommitError,
	)

	genericErrorTest(
		t,
		"Not implemented: ",
		func(text string) error { return NewErrNotImplemented(text) },
		IsNotImplementedError,
	)

	genericErrorTest(
		t,
		"Cannot communicate: ",
		func(text string) error { return NewErrCannotCommunicate(text) },
		IsCannotCommunicateError,
	)

	genericErrorTest(
		t,
		"Cannot communicate: ",
		func(text string) error { return NewErrLocked(text) },
		IsLockedError,
	)

	genericErrorTest(
		t,
		"Transaction not found: ",
		func(text string) error { return NewErrTransactionNotFound(text) },
		IsTransactionNotFoundError,
	)

	genericErrorTest(
		t,
		"Constraint violation: ",
		func(text string) error { return NewErrConstraintViolation(text) },
		IsConstraintViolationError,
	)

	genericErrorTest(
		t,
		"Disconnected: ",
		func(text string) error { return NewErrDisconnected(text) },
		IsDisconnectedError,
	)

	genericErrorTest(
		t,
		"Too many requests: ",
		func(text string) error { return NewErrTooManyRequests(text) },
		IsTooManyRequestsError,
	)

	genericErrorTest(
		t,
		"TLS error: ",
		func(text string) error { return NewErrTLS(text) },
		IsTLSError,
	)
}
