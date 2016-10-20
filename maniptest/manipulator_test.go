package maniptest

import (
	"fmt"
	"testing"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTestManipulator_MockRetrieveMany(t *testing.T) {

	Convey("Given I have TestManipulator", t, func() {

		m := NewTestManipulator()

		Convey("When I call RetrieveMany without mock", func() {

			err := m.RetrieveMany(nil, PersonIdentity, nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I mock it to return an error", func() {

				m.MockRetrieveMany(t, func(context *manipulate.Context, identity elemental.Identity, dest interface{}) error {
					return fmt.Errorf("wow such error")
				})

				err := m.RetrieveMany(nil, PersonIdentity, nil)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("When I don't mock it", func() {

				err := m.RetrieveMany(nil, PersonIdentity, nil)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestTestManipulator_MockRetrieve(t *testing.T) {

	Convey("Given I have TestManipulator", t, func() {

		m := NewTestManipulator()

		Convey("When I call Retrieve without mock", func() {

			err := m.Retrieve(nil, nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I mock it to return an error", func() {

				m.MockRetrieve(t, func(ctx *manipulate.Context, objects ...manipulate.Manipulable) error {
					return fmt.Errorf("wow such error")
				})

				err := m.Retrieve(nil, nil)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("When I don't mock it", func() {

				err := m.Retrieve(nil, nil)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestTestManipulator_MockCreate(t *testing.T) {

	Convey("Given I have TestManipulator", t, func() {

		m := NewTestManipulator()

		Convey("When I call Retrieve without mock", func() {

			err := m.Create(nil, nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I mock it to return an error", func() {

				m.MockCreate(t, func(ctx *manipulate.Context, objects ...manipulate.Manipulable) error {
					return fmt.Errorf("wow such error")
				})

				err := m.Create(nil, nil)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("When I don't mock it", func() {

				err := m.Create(nil, nil)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestTestManipulator_MockUpdate(t *testing.T) {

	Convey("Given I have TestManipulator", t, func() {

		m := NewTestManipulator()

		Convey("When I call Retrieve without mock", func() {

			err := m.Update(nil, nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I mock it to return an error", func() {

				m.MockUpdate(t, func(ctx *manipulate.Context, objects ...manipulate.Manipulable) error {
					return fmt.Errorf("wow such error")
				})

				err := m.Update(nil, nil)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("When I don't mock it", func() {

				err := m.Update(nil, nil)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestTestManipulator_MockDelete(t *testing.T) {

	Convey("Given I have TestManipulator", t, func() {

		m := NewTestManipulator()

		Convey("When I call Retrieve without mock", func() {

			err := m.Delete(nil, nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I mock it to return an error", func() {

				m.MockDelete(t, func(ctx *manipulate.Context, objects ...manipulate.Manipulable) error {
					return fmt.Errorf("wow such error")
				})

				err := m.Delete(nil, nil)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("When I don't mock it", func() {

				err := m.Delete(nil, nil)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestTestManipulator_MockCount(t *testing.T) {

	Convey("Given I have TestManipulator", t, func() {

		m := NewTestManipulator()

		Convey("When I call Retrieve without mock", func() {

			c, err := m.Count(nil, PersonIdentity)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then c should be 0", func() {
				So(c, ShouldEqual, 0)
			})

			Convey("When I mock it to return an error", func() {

				m.MockCount(t, func(ctx *manipulate.Context, identity elemental.Identity) (int, error) {
					return -1, fmt.Errorf("wow such error")
				})

				c, err := m.Count(nil, PersonIdentity)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})

				Convey("Then c should be 0", func() {
					So(c, ShouldEqual, -1)
				})
			})

			Convey("When I don't mock it", func() {

				c, err := m.Count(nil, PersonIdentity)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then c should be 0", func() {
					So(c, ShouldEqual, 0)
				})
			})
		})
	})
}

func TestTestManipulator_MockAssign(t *testing.T) {

	Convey("Given I have TestManipulator", t, func() {

		m := NewTestManipulator()

		Convey("When I call Retrieve without mock", func() {

			err := m.Assign(nil, nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I mock it to return an error", func() {

				m.MockAssign(t, func(ctx *manipulate.Context, assignation *elemental.Assignation) error {
					return fmt.Errorf("wow such error")
				})

				err := m.Assign(nil, nil)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("When I don't mock it", func() {

				err := m.Assign(nil, nil)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestTestManipulator_MockIncrement(t *testing.T) {

	Convey("Given I have TestManipulator", t, func() {

		m := NewTestManipulator()

		Convey("When I call Retrieve without mock", func() {

			err := m.Increment(nil, "name", "counter", 1, nil, nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I mock it to return an error", func() {

				m.MockIncrement(t, func(ctx *manipulate.Context, name string, counter string, inc int, keys []string, values []interface{}) error {
					return fmt.Errorf("wow such error")
				})

				err := m.Increment(nil, "name", "counter", 1, nil, nil)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("When I don't mock it", func() {

				err := m.Increment(nil, "name", "counter", 1, nil, nil)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestTestManipulator_MockCommit(t *testing.T) {

	Convey("Given I have TestManipulator", t, func() {

		m := NewTestManipulator()

		Convey("When I call Retrieve without mock", func() {

			err := m.Commit(manipulate.NewTransactionID())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I mock it to return an error", func() {

				m.MockCommit(t, func(tid manipulate.TransactionID) error {
					return fmt.Errorf("wow such error")
				})

				err := m.Commit(manipulate.NewTransactionID())

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("When I don't mock it", func() {

				err := m.Commit(manipulate.NewTransactionID())

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestTestManipulator_MockAbort(t *testing.T) {

	Convey("Given I have TestManipulator", t, func() {

		m := NewTestManipulator()

		Convey("When I call Retrieve without mock", func() {

			ok := m.Abort(manipulate.NewTransactionID())

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})

			Convey("When I mock it to return an error", func() {

				m.MockAbort(t, func(tid manipulate.TransactionID) bool {
					return false
				})

				ok = m.Abort(manipulate.NewTransactionID())

				Convey("Then err should not be nil", func() {
					So(ok, ShouldBeFalse)
				})
			})

			Convey("When I don't mock it", func() {

				ok = m.Abort(manipulate.NewTransactionID())

				Convey("Then err should be nil", func() {
					So(ok, ShouldBeTrue)
				})
			})
		})
	})
}
