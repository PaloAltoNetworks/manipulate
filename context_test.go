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
			So(ContextForIndex(nil, -1), ShouldEqual, nil)
			So(ContextForIndex([]*Context{context}, 0), ShouldEqual, context)
		})
	})
}
