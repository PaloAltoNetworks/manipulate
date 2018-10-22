// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import (
	"context"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental/test/model"
)

func TestMethodNewContext(t *testing.T) {

	Convey("Given I create a new context", t, func() {

		mctx := NewContext(context.Background())

		Convey("Then my context should be initiliazed", func() {
			So(mctx.Page(), ShouldEqual, 0)
			So(mctx.PageSize(), ShouldEqual, 0)
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

		filter := NewFilter()
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

		mctx := &mcontext{
			page:                 1,
			pageSize:             2,
			parent:               testmodel.NewList(),
			countTotal:           3,
			filter:               NewFilterComposer().WithKey("k").Equals("v").Done(),
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
		}

		Convey("When I Derive without option", func() {

			copy := mctx.Derive()

			Convey("Then the copy should resemble to the original", func() {
				So(copy, ShouldResemble, mctx)
			})
		})

		Convey("When I Derive without witah options", func() {

			copy := mctx.Derive(
				ContextOptionPage(11, 12),
				ContextOptionFilter(NewFilterComposer().WithKey("k").Equals("v2").Done()),
			).(*mcontext)

			Convey("Then the copy should resemble to the original but for the changes", func() {
				So(copy.page, ShouldEqual, 11)
				So(copy.pageSize, ShouldEqual, 12)
				So(copy.parent, ShouldEqual, mctx.parent)
				So(copy.countTotal, ShouldEqual, mctx.countTotal)
				So(copy.filter.String(), ShouldEqual, `k == "v2"`)
				So(copy.parameters, ShouldEqual, mctx.parameters)
				So(copy.transactionID, ShouldEqual, mctx.transactionID)
				So(copy.namespace, ShouldEqual, mctx.namespace)
				So(copy.recursive, ShouldEqual, mctx.recursive)
				So(copy.overrideProtection, ShouldEqual, mctx.overrideProtection)
				So(copy.createFinalizer, ShouldEqual, mctx.createFinalizer)
				So(copy.version, ShouldEqual, mctx.version)
				So(copy.externalTrackingID, ShouldEqual, mctx.externalTrackingID)
				So(copy.externalTrackingType, ShouldEqual, mctx.externalTrackingType)
				So(copy.order, ShouldResemble, mctx.order)
				So(copy.fields, ShouldResemble, mctx.fields)
				So(copy.ctx, ShouldEqual, mctx.ctx)
			})
		})
	})
}
