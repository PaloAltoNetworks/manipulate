package maniptest

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTestSubscriber_MockIssue(t *testing.T) {

	Convey("Given I have TestTokenManager", t, func() {

		tm := NewTestTokenManager()

		Convey("When I call Issue without mock", func() {

			token, err := tm.Issue(context.Background())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then token should be correct", func() {
				So(token, ShouldBeEmpty)
			})
		})

		Convey("When I call Issue with a mock", func() {

			tm.MockIssue(t, func(ctx context.Context) (string, error) {
				return "token", fmt.Errorf("test")
			})

			token, err := tm.Issue(context.Background())

			Convey("Then err should be correct", func() {
				So(err.Error(), ShouldEqual, "test")
			})

			Convey("Then token should be correct", func() {
				So(token, ShouldEqual, "token")
			})
		})
	})
}

func TestTestSubscriber_MockRun(t *testing.T) {

	Convey("Given I have TestTokenManager", t, func() {

		tm := NewTestTokenManager()

		Convey("When I call Run without mock", func() {

			tm.Run(context.Background(), nil)

			Convey("Then nothing should happen", func() {
			})
		})

		Convey("When I call Issue with a mock", func() {

			tm.MockRun(t, func(ctx context.Context, tokenCh chan string) {
				go func() { tokenCh <- "coucou" }()
			})

			ch := make(chan string)
			tm.Run(context.Background(), ch)

			Convey("Then teh value in the chan should be correct", func() {
				So(<-ch, ShouldEqual, "coucou")
			})
		})
	})
}
