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
	"time"

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
			So(mctx.Context(), ShouldEqual, ctx)
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
			ContextOptionTransationID(tid),
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

		rfunc := func(int, error) error { return nil }

		mctx := &mcontext{
			page:                 1,
			pageSize:             2,
			parent:               testmodel.NewList(),
			countTotal:           3,
			filter:               elemental.NewFilterComposer().WithKey("k").Equals("v").Done(),
			parameters:           url.Values{"a": []string{"b"}},
			transactionID:        NewTransactionID(),
			namespace:            "/",
			recursive:            true,
			overrideProtection:   true,
			createFinalizer:      nil,
			version:              4,
			externalTrackingID:   "externalTrackingID",
			externalTrackingType: "externalTrackingType",
			order:                []string{"a", "b"},
			fields:               []string{"a", "b"},
			ctx:                  context.Background(),
			requestTimeout:       42 * time.Second,
			idempotencyKey:       "ikey",
			retryFunc:            rfunc,
		}

		Convey("When I Derive without option", func() {

			copy := mctx.Derive().(*mcontext)

			Convey("Then the copy should resemble to the original", func() {
				So(copy.Page(), ShouldEqual, 1)
				So(copy.PageSize(), ShouldEqual, 2)
				So(copy.Parent(), ShouldEqual, mctx.parent)
				So(copy.Count(), ShouldEqual, mctx.countTotal)
				So(copy.filter.String(), ShouldEqual, `k == "v"`)
				So(copy.Parameters(), ShouldEqual, mctx.parameters)
				So(copy.TransactionID(), ShouldEqual, mctx.transactionID)
				So(copy.Namespace(), ShouldEqual, mctx.namespace)
				So(copy.Recursive(), ShouldEqual, mctx.recursive)
				So(copy.Override(), ShouldEqual, mctx.overrideProtection)
				So(copy.Finalizer(), ShouldEqual, mctx.createFinalizer)
				So(copy.Version(), ShouldEqual, mctx.version)
				So(copy.ExternalTrackingID(), ShouldEqual, mctx.externalTrackingID)
				So(copy.ExternalTrackingType(), ShouldEqual, mctx.externalTrackingType)
				So(copy.Order(), ShouldResemble, mctx.order)
				So(copy.Fields(), ShouldResemble, mctx.fields)
				So(copy.ctx, ShouldEqual, mctx.ctx)
				So(copy.IdempotencyKey(), ShouldEqual, "")
				So(copy.RequestTimeout(), ShouldEqual, 42*time.Second)
				So(copy.RetryFunc(), ShouldEqual, rfunc)
			})
		})

		Convey("When I Derive with options", func() {

			copy := mctx.Derive(
				ContextOptionPage(11, 12),
				ContextOptionFilter(elemental.NewFilterComposer().WithKey("k").Equals("v2").Done()),
			).(*mcontext)

			Convey("Then the copy should resemble to the original but for the changes", func() {
				So(copy.Page(), ShouldEqual, 11)
				So(copy.PageSize(), ShouldEqual, 12)
				So(copy.Parent(), ShouldEqual, mctx.parent)
				So(copy.Count(), ShouldEqual, mctx.countTotal)
				So(copy.filter.String(), ShouldEqual, `k == "v2"`)
				So(copy.Parameters(), ShouldEqual, mctx.parameters)
				So(copy.TransactionID(), ShouldEqual, mctx.transactionID)
				So(copy.Namespace(), ShouldEqual, mctx.namespace)
				So(copy.Recursive(), ShouldEqual, mctx.recursive)
				So(copy.Override(), ShouldEqual, mctx.overrideProtection)
				So(copy.Finalizer(), ShouldEqual, mctx.createFinalizer)
				So(copy.Version(), ShouldEqual, mctx.version)
				So(copy.ExternalTrackingID(), ShouldEqual, mctx.externalTrackingID)
				So(copy.ExternalTrackingType(), ShouldEqual, mctx.externalTrackingType)
				So(copy.Order(), ShouldResemble, mctx.order)
				So(copy.Fields(), ShouldResemble, mctx.fields)
				So(copy.ctx, ShouldEqual, mctx.ctx)
				So(copy.IdempotencyKey(), ShouldEqual, "")
				So(copy.RequestTimeout(), ShouldEqual, mctx.requestTimeout)
				So(copy.RetryFunc(), ShouldEqual, rfunc)
			})
		})
	})
}
