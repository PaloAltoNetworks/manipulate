// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipulate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMethodNewContext(t *testing.T) {

	Convey("Given I create a new context", t, func() {

		context := NewContext()

		Convey("Then my context should be initiliazed", func() {
			So(context.PageCurrent, ShouldEqual, 1)
			So(context.PageSize, ShouldEqual, 100)
		})
	})
}

func TestMethodNewContextWithFilter(t *testing.T) {

	Convey("Given I create a new context with filter", t, func() {

		filter := NewFilter()
		context := NewContextWithFilter(filter)

		Convey("Then my context should be initiliazed", func() {
			So(context.Filter, ShouldEqual, filter)
		})
	})
}

func TestMethodNewContextWithTransactionID(t *testing.T) {

	Convey("Given I create a new context iwth transactionID", t, func() {

		tid := NewTransactionID()
		context := NewContextWithTransactionID(tid)

		Convey("Then my context should be initiliazed", func() {
			So(context.TransactionID, ShouldEqual, tid)
		})
	})
}

func TestMethodString(t *testing.T) {

	Convey("Given I create a new context and calle the method string", t, func() {

		context := NewContext()

		Convey("Then my context should be initiliazed", func() {
			So(context.String(), ShouldEqual, "<Context page: 1, pagesize: 100> <Filter : <nil>>")
		})
	})
}

func TestMethodContextForIndex(t *testing.T) {

	Convey("Given I create a new context and calle the method ContextForIndex", t, func() {

		context := NewContext()

		Convey("Then I should get the good context", func() {
			So(ContextForIndex(context, -1), ShouldEqual, context)
			So(ContextForIndex(nil, -1), ShouldResemble, NewContext())
			So(ContextForIndex([]*Context{context}, 0), ShouldEqual, context)
		})
	})
}
