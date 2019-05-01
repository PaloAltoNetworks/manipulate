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

func TestTransaction_NewTransactionID(t *testing.T) {

	Convey("Given I create a NewTransactionID", t, func() {

		tid := NewTransactionID()

		Convey("Then it should have the correct type", func() {
			So(tid, ShouldHaveSameTypeAs, TransactionID("ttt"))
		})

		Convey("Then it's len you be correct", func() {
			So(len(tid), ShouldEqual, 36)
		})
	})
}
