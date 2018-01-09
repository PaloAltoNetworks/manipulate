package manipwebsocket

import (
	"fmt"
	"sync"
	"testing"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/internal/testdata"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHelpers_decodeErrors(t *testing.T) {

	Convey("Given I have an elemental.Response with some errors", t, func() {

		resp := elemental.NewResponse()
		errs := elemental.NewErrors(
			elemental.NewError("title", "description", "subject", 500),
			elemental.NewError("title", "description", "subject", 500),
		)
		resp.Encode(errs) // nolint: errcheck

		Convey("When I call decodeError", func() {

			err := decodeErrors(resp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then err should resemble to the original err", func() {
				So(err, ShouldResemble, errs)
			})
		})
	})

	Convey("Given I have an elemental.Response with broken data", t, func() {

		resp := elemental.NewResponse()
		resp.Data = []byte("bad data")

		Convey("When I call decodeError", func() {

			err := decodeErrors(resp)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then err should be an NewErrCannotUnmarshal", func() {
				So(err, ShouldHaveSameTypeAs, manipulate.NewErrCannotUnmarshal("poof"))
			})
		})
	})
}

func TestHelpers_handleCommunicationError(t *testing.T) {

	Convey("Given I have connected manipulator and some classic error", t, func() {

		m := &websocketManipulator{
			connectedLock: &sync.Mutex{},
			renewLock:     &sync.Mutex{},
			connected:     true,
		}

		e := fmt.Errorf("woopsy")

		Convey("When I call handleCommunicationError", func() {

			err := handleCommunicationError(m, e)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then err should be a NewErrCannotCommunicate", func() {
				So(err, ShouldHaveSameTypeAs, manipulate.NewErrCannotCommunicate(""))
			})
		})
	})

	Convey("Given I have disconnected manipulator and some classic error", t, func() {

		m := &websocketManipulator{
			connectedLock: &sync.Mutex{},
			renewLock:     &sync.Mutex{},
		}

		e := fmt.Errorf("woopsy")

		Convey("When I call handleCommunicationError", func() {

			err := handleCommunicationError(m, e)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then err should be a NewErrDisconnected", func() {
				So(err, ShouldHaveSameTypeAs, manipulate.NewErrDisconnected(""))
			})
		})
	})

	Convey("Given I have connect manipulator and some ErrDisconnected error", t, func() {

		m := &websocketManipulator{
			connectedLock: &sync.Mutex{},
			renewLock:     &sync.Mutex{},
			connected:     true,
		}

		e := manipulate.NewErrDisconnected("boom")

		Convey("When I call handleCommunicationError", func() {

			err := handleCommunicationError(m, e)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then err should be a NewErrDisconnected", func() {
				So(err, ShouldResemble, e)
			})
		})
	})
}

func TestHelpers_populateRequestFromContext(t *testing.T) {

	Convey("Given I have a request and a context", t, func() {

		mctx := manipulate.NewContextWithFilter(
			manipulate.NewFilterComposer().
				WithKey("key").Equals(true).
				Done(),
		)
		mctx.Parameters.Add("key", "value")

		parent := testdata.NewList()
		parent.ID = "xxx"
		mctx.Parent = parent

		mctx.Namespace = "/ns"
		mctx.Recursive = true

		mctx.ExternalTrackingID = "tid"
		mctx.ExternalTrackingType = "ttype"
		mctx.Page = 2
		mctx.PageSize = 20
		mctx.OverrideProtection = true
		mctx.Order = []string{"a", "b"}

		req := elemental.NewRequest()
		o := testdata.NewList()

		Convey("When I call populateRequestFromContext", func() {

			err := populateRequestFromContext(req, mctx, o)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then request parameters should be correct", func() {
				So(req.Parameters.Get("key"), ShouldEqual, "value")
			})

			Convey("Then request parent should be correct", func() {
				So(req.ParentIdentity.Name, ShouldEqual, testdata.ListIdentity.Name)
				So(req.ParentID, ShouldEqual, "xxx")
			})

			Convey("Then request namespace should be correct", func() {
				So(req.Namespace, ShouldEqual, "/ns")
			})

			Convey("Then request recursive should be correct", func() {
				So(req.Recursive, ShouldBeTrue)
			})

			Convey("Then request version should be correct", func() {
				So(req.Version, ShouldEqual, 1)
			})

			Convey("Then request pagination should be correct", func() {
				So(req.Page, ShouldEqual, 2)
				So(req.PageSize, ShouldEqual, 20)
			})

			Convey("Then request tracking should be correct", func() {
				So(req.ExternalTrackingID, ShouldEqual, "tid")
				So(req.ExternalTrackingType, ShouldEqual, "ttype")
			})

			Convey("Then request protection should be correct", func() {
				So(req.OverrideProtection, ShouldBeTrue)
			})

			Convey("Then request order should be correct", func() {
				So(req.Order, ShouldResemble, []string{"a", "b"})
			})
		})
	})
}
