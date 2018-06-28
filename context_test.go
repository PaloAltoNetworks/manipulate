// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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
