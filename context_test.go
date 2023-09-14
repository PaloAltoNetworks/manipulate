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
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
)

func TestMethodNewContext(t *testing.T) {

	Convey("Given I create a new context", t, func() {

		mctx := NewContext(context.Background())

		Convey("Then my context should be initiliazed", func() {
			So(mctx.Page(), ShouldEqual, 0)
			So(mctx.PageSize(), ShouldEqual, 0)
			So(mctx.RetryRatio(), ShouldEqual, 4)
			So(mctx.WriteConsistency(), ShouldEqual, WriteConsistencyDefault)
			So(mctx.ReadConsistency(), ShouldEqual, ReadConsistencyDefault)
		})
	})
}

func TestMethodWithContext(t *testing.T) {

	Convey("Given I create a new context with a context", t, func() {

		ctx := context.Background()
		mctx := NewContext(
			ctx,
			ContextOptionPage(1, 0),
		)

		Convey("Then my context should be initiliazed", func() {
			So(mctx.Page(), ShouldEqual, 1)
			So(mctx.Context(), ShouldResemble, ctx)
		})
	})

	Convey("Given I create a new context with a nil context", t, func() {

		Convey("Then it should panic", func() {
			So(func() { NewContext(nil) }, ShouldPanicWith, "nil context") // nolint
		})
	})

	Convey("Given I create a new context without context", t, func() {

		mctx := NewContext(context.Background())

		Convey("Then it should panic", func() {
			So(mctx.Context(), ShouldNotBeNil)
		})
	})
}

func TestMethodNewContextWithFilter(t *testing.T) {

	Convey("Given I create a new context with filter", t, func() {

		filter := elemental.NewFilter()
		mctx := NewContext(
			context.Background(),
			ContextOptionFilter(filter),
		)

		Convey("Then my context should be initiliazed", func() {
			So(mctx.Filter(), ShouldEqual, filter)
		})
	})
}

func TestMethodNewContextWithTransactionID(t *testing.T) {

	Convey("Given I create a new context with transactionID", t, func() {

		tid := NewTransactionID()
		mctx := NewContext(
			context.Background(),
			ContextOptionTransactionID(tid),
		)

		Convey("Then my context should be initiliazed", func() {
			So(mctx.TransactionID(), ShouldEqual, tid)
		})
	})
}

func TestMethodString(t *testing.T) {

	Convey("Given I create a new context and calle the method string", t, func() {

		mctx := NewContext(
			context.Background(),
			ContextOptionPage(1, 100),
			ContextOptionVersion(12),
		)

		Convey("Then my context should be initiliazed", func() {
			So(mctx.String(), ShouldEqual, "<Context page:1 pagesize:100 filter:<nil> version:12>")
		})
	})
}

func TestContext_Derive(t *testing.T) {

	Convey("Given I have a context", t, func() {

		rfunc := func(RetryInfo) error { return nil }

		mctx := &mcontext{
			page:                 1,
			pageSize:             2,
			after:                "42",
			parent:               testmodel.NewList(),
			filter:               elemental.NewFilterComposer().WithKey("k").Equals("v").Done(),
			parameters:           url.Values{"a": []string{"b"}},
			transactionID:        NewTransactionID(),
			namespace:            "/",
			recursive:            true,
			overrideProtection:   true,
			propagated:           true,
			createFinalizer:      nil,
			version:              4,
			externalTrackingID:   "externalTrackingID",
			externalTrackingType: "externalTrackingType",
			order:                []string{"a", "b"},
			fields:               []string{"a", "b"},
			ctx:                  context.Background(),
			retryFunc:            rfunc,
			writeConsistency:     WriteConsistencyStrong,
			readConsistency:      ReadConsistencyMonotonic,
			clientIP:             "1.1.1.1",
			retryRatio:           12,
			opaque:               map[string]any{"a": "b"},
		}

		mctx.SetCount(3)
		mctx.SetMessages([]string{"hello"})
		mctx.SetIdempotencyKey("ikey")
		mctx.SetCredentials("user", "password")

		u, p := mctx.Credentials()
		So(u, ShouldEqual, "user")
		So(p, ShouldEqual, "password")

		Convey("When I Derive without option", func() {

			copied := mctx.Derive().(*mcontext)

			Convey("Then the copy should resemble to the original", func() {

				So(copied.Count(), ShouldEqual, 0)
				So(copied.IdempotencyKey(), ShouldEqual, "")
				So(copied.Messages(), ShouldBeNil)

				So(copied.ClientIP(), ShouldEqual, mctx.clientIP)
				So(copied.ExternalTrackingID(), ShouldEqual, mctx.externalTrackingID)
				So(copied.ExternalTrackingType(), ShouldEqual, mctx.externalTrackingType)
				So(copied.Fields(), ShouldResemble, mctx.fields)
				So(copied.Fields(), ShouldNotEqual, mctx.fields)
				So(copied.Filter().String(), ShouldEqual, `k == "v"`)
				So(copied.Finalizer(), ShouldEqual, mctx.createFinalizer)
				So(copied.Namespace(), ShouldEqual, mctx.namespace)
				So(copied.Order(), ShouldResemble, mctx.order)
				So(copied.Order(), ShouldNotEqual, mctx.order)
				So(copied.Override(), ShouldEqual, mctx.overrideProtection)
				So(copied.Page(), ShouldEqual, mctx.page)
				So(copied.PageSize(), ShouldEqual, mctx.pageSize)
				So(copied.After(), ShouldEqual, mctx.after)
				So(copied.Parameters(), ShouldResemble, mctx.parameters)
				So(copied.Parameters(), ShouldNotEqual, mctx.parameters)
				So(copied.Parent(), ShouldEqual, mctx.parent)
				So(copied.password, ShouldEqual, mctx.password)
				So(copied.ReadConsistency(), ShouldEqual, mctx.readConsistency)
				So(copied.Recursive(), ShouldEqual, mctx.recursive)
				So(copied.Propagated(), ShouldEqual, mctx.propagated)
				So(copied.RetryFunc(), ShouldEqual, rfunc)
				So(copied.String(), ShouldEqual, mctx.String())
				So(copied.TransactionID(), ShouldEqual, mctx.transactionID)
				So(copied.username, ShouldEqual, mctx.username)
				So(copied.Version(), ShouldEqual, mctx.version)
				So(copied.WriteConsistency(), ShouldEqual, mctx.writeConsistency)
				So(copied.Context(), ShouldResemble, mctx.ctx)
				So(copied.RetryRatio(), ShouldEqual, mctx.retryRatio)
				So(copied.Opaque(), ShouldResemble, mctx.opaque)
				So(copied.Opaque(), ShouldNotEqual, mctx.opaque)
			})
		})

		Convey("When I Derive with options", func() {

			copied := mctx.Derive(
				ContextOptionPage(11, 12),
				ContextOptionFilter(elemental.NewFilterComposer().WithKey("k").Equals("v2").Done()),
			).(*mcontext)

			Convey("Then the copy should resemble to the original but for the changes", func() {
				So(copied.Page(), ShouldEqual, 11)
				So(copied.PageSize(), ShouldEqual, 12)
				So(copied.Filter().String(), ShouldEqual, `k == "v2"`)
				So(copied.String(), ShouldNotEqual, mctx.String())

				So(copied.Count(), ShouldEqual, 0)
				So(copied.IdempotencyKey(), ShouldEqual, "")
				So(copied.Messages(), ShouldBeNil)

				So(copied.ClientIP(), ShouldEqual, mctx.clientIP)
				So(copied.ExternalTrackingID(), ShouldEqual, mctx.externalTrackingID)
				So(copied.ExternalTrackingType(), ShouldEqual, mctx.externalTrackingType)
				So(copied.Fields(), ShouldResemble, mctx.fields)
				So(copied.Fields(), ShouldNotEqual, mctx.fields)
				So(copied.Finalizer(), ShouldEqual, mctx.createFinalizer)
				So(copied.Namespace(), ShouldEqual, mctx.namespace)
				So(copied.Order(), ShouldResemble, mctx.order)
				So(copied.Order(), ShouldNotEqual, mctx.order)
				So(copied.Override(), ShouldEqual, mctx.overrideProtection)
				So(copied.Parameters(), ShouldResemble, mctx.parameters)
				So(copied.Parameters(), ShouldNotEqual, mctx.parameters)
				So(copied.Parent(), ShouldEqual, mctx.parent)
				So(copied.password, ShouldEqual, mctx.password)
				So(copied.ReadConsistency(), ShouldEqual, mctx.readConsistency)
				So(copied.Recursive(), ShouldEqual, mctx.recursive)
				So(copied.RetryFunc(), ShouldEqual, rfunc)
				So(copied.TransactionID(), ShouldEqual, mctx.transactionID)
				So(copied.username, ShouldEqual, mctx.username)
				So(copied.Version(), ShouldEqual, mctx.version)
				So(copied.WriteConsistency(), ShouldEqual, mctx.writeConsistency)
				So(copied.Context(), ShouldResemble, mctx.ctx)
				So(copied.RetryRatio(), ShouldEqual, mctx.retryRatio)
				So(copied.Opaque(), ShouldResemble, mctx.opaque)
				So(copied.Opaque(), ShouldNotEqual, mctx.opaque)
			})
		})
	})
}
