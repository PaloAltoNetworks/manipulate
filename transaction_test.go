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
