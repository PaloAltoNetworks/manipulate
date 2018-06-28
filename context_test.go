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

		context := NewContext()

		Convey("Then my context should be initiliazed", func() {
			So(context.Page(), ShouldEqual, 0)
			So(context.PageSize(), ShouldEqual, 0)
		})
	})
}

func TestMethodWithContext(t *testing.T) {

	Convey("Given I create a new context with a context", t, func() {

		ctx := context.Background()
		context := NewContext(
			ContextOptionContext(ctx),
			ContextOptionPage(1, 0),
		)

		Convey("Then my context should be initiliazed", func() {
			So(context.Page(), ShouldEqual, 1)
			So(context.Context(), ShouldEqual, ctx)
		})
	})

	Convey("Given I create a new context with a nil context", t, func() {

		Convey("Then it should panic", func() {
			So(func() { NewContext(ContextOptionContext(nil)) }, ShouldPanicWith, "nil context") // nolint
		})
	})

	Convey("Given I create a new context without context", t, func() {

		context := NewContext()

		Convey("Then it should panic", func() {
			So(context.Context(), ShouldNotBeNil)
		})
	})
}

func TestMethodNewContextWithFilter(t *testing.T) {

	Convey("Given I create a new context with filter", t, func() {

		filter := NewFilter()
		context := NewContext(
			ContextOptionFilter(filter),
		)

		Convey("Then my context should be initiliazed", func() {
			So(context.Filter(), ShouldEqual, filter)
		})
	})
}

func TestMethodNewContextWithTransactionID(t *testing.T) {

	Convey("Given I create a new context with transactionID", t, func() {

		tid := NewTransactionID()
		context := NewContext(
			ContextOptionTransationID(tid),
		)

		Convey("Then my context should be initiliazed", func() {
			So(context.TransactionID(), ShouldEqual, tid)
		})
	})
}

func TestMethodString(t *testing.T) {

	Convey("Given I create a new context and calle the method string", t, func() {

		context := NewContext(
			ContextOptionPage(1, 100),
			ContextOptionVersion(12),
		)

		Convey("Then my context should be initiliazed", func() {
			So(context.String(), ShouldEqual, "<Context page:1 pagesize:100 filter:<nil> version:12>")
		})
	})
}
