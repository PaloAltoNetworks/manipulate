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

		context := NewContext(context.Background())

		Convey("Then my context should be initiliazed", func() {
			So(context.Page, ShouldEqual, 0)
			So(context.PageSize, ShouldEqual, 0)
		})
	})
}

func TestMethodString(t *testing.T) {

	Convey("Given I create a new context and calle the method string", t, func() {

		context := NewContext(context.Background())
		context.Page = 1
		context.PageSize = 100
		context.Version = 12

		Convey("Then my context should be initiliazed", func() {
			So(context.String(), ShouldEqual, "<Context page:1 pagesize:100 filter:<nil> version:12>")
		})
	})
}
