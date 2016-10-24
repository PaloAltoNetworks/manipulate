package manipmemory

import (
	"testing"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMemManipulator_NewMemoryManipulator(t *testing.T) {

	Convey("Given I create a new MemoryManipulator with bad schema", t, func() {

		Convey("Then it should panic", func() {
			So(func() { NewMemoryManipulator(nil) }, ShouldPanic)
		})
	})
}

func TestMemManipulator_Create(t *testing.T) {

	Convey("Given I have a memory manipulator and a person", t, func() {

		m := NewMemoryManipulator(Schema)
		p := &Person{
			Name: "Antoine",
		}

		Convey("When I create person", func() {

			err := m.Create(nil, p)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then person ID should not be empty", func() {
				So(p.ID, ShouldNotBeEmpty)
			})

			Convey("When I retrieve the person in a second structure", func() {

				p2 := &Person{
					ID:   p.ID,
					Name: "not good",
				}

				err := m.Retrieve(nil, p2)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then p2 should be p", func() {
					So(p2, ShouldResemble, p)
				})
			})
		})

		Convey("When I create an object that is not part of the schema", func() {

			err := m.Create(nil, &NotPerson{})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestMemManipulator_Retrieve(t *testing.T) {

	Convey("Given I have a memory manipulator and a person", t, func() {

		m := NewMemoryManipulator(Schema)
		p1 := &Person{
			Name: "Antoine1",
		}

		m.Create(nil, p1)

		Convey("When I retrieve the person", func() {

			ps := &Person{
				ID: p1.ID,
			}

			err := m.Retrieve(nil, ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should  have retrieved p1 and p2", func() {
				So(ps, ShouldResemble, p1)
			})
		})

		Convey("When I retrieve a non existing person", func() {

			ps := &Person{
				ID: "not-good",
			}

			err := m.Retrieve(nil, ps)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I retrieve an object that is not part of the schema", func() {

			err := m.Retrieve(nil, &NotPerson{})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

	})
}

func TestMemManipulator_RetrieveMany(t *testing.T) {

	Convey("Given I have a memory manipulator and a person", t, func() {

		m := NewMemoryManipulator(Schema)
		p1 := &Person{
			Name: "Antoine1",
		}
		p2 := &Person{
			Name: "Antoine2",
		}

		m.Create(nil, p1, p2)

		Convey("When I retrieve the persons", func() {

			ps := []*Person{}

			err := m.RetrieveMany(nil, PersonIdentity, &ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should  have retrieved p1 and p2", func() {
				So(ps, ShouldContain, p1)
				So(ps, ShouldContain, p2)
			})
		})

		Convey("When I retrieve the persons with a filter that matches p1", func() {

			ps := []*Person{}

			ctx := manipulate.NewContextWithFilter(
				manipulate.NewFilterComposer().WithKey("Name").Equals("Antoine1").Done(),
			)

			err := m.RetrieveMany(ctx, PersonIdentity, &ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should only have retrieved p1", func() {
				So(ps, ShouldContain, p1)
				So(ps, ShouldNotContain, p2)
			})
		})

		Convey("When I retrieve the persons with a bad filter", func() {

			ps := []*Person{}

			ctx := manipulate.NewContextWithFilter(
				manipulate.NewFilterComposer().WithKey("Bad").Equals("Antoine1").Done(),
			)

			err := m.RetrieveMany(ctx, PersonIdentity, &ps)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.(elemental.Error).Code, ShouldEqual, manipulate.ErrCannotExecuteQuery)
			})
		})
	})
}

func TestMemManipulator_Update(t *testing.T) {

	Convey("Given I have a memory manipulator and a person", t, func() {

		m := NewMemoryManipulator(Schema)
		p := &Person{
			Name: "Antoine",
		}

		Convey("When I create the person", func() {

			err := m.Create(nil, p)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I update the person", func() {

				p.Name = "New Antoine"

				err := m.Update(nil, p)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("When I retrieve the person", func() {

					p2 := &Person{
						ID: p.ID,
					}

					err := m.Retrieve(nil, p2)

					Convey("Then err should be nil", func() {
						So(err, ShouldBeNil)
					})

					Convey("Then p2 should contains the updated data", func() {
						So(p2.Name, ShouldEqual, "New Antoine")
					})
				})
			})

			Convey("When I update the a non existing person", func() {

				pp := &Person{
					ID: "not-good",
				}

				err := m.Update(nil, pp)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestMemManipulator_Delete(t *testing.T) {

	Convey("Given I have a memory manipulator and a person", t, func() {

		m := NewMemoryManipulator(Schema)
		p := &Person{
			Name: "Antoine",
		}

		Convey("When I create the person", func() {

			err := m.Create(nil, p)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I delete the person", func() {

				err := m.Delete(nil, p)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("When I retrieve the persons using a wrapper", func() {

					ps := []*Person{}

					var err error

					func(zob interface{}) {
						err = m.RetrieveMany(nil, PersonIdentity, zob)
					}(&ps)

					Convey("Then err should be nil", func() {
						So(err, ShouldBeNil)
					})

					Convey("Then I the result should be empty", func() {
						So(len(ps), ShouldEqual, 0)
					})
				})

				Convey("When I retrieve the persons ", func() {

					ps := []*Person{}
					err := m.RetrieveMany(nil, PersonIdentity, &ps)

					Convey("Then err should be nil", func() {
						So(err, ShouldBeNil)
					})

					Convey("Then I the result should be empty", func() {
						So(len(ps), ShouldEqual, 0)
					})
				})
			})

			Convey("When I delete the a non existing person", func() {

				pp := &Person{
					ID: "not-good",
				}

				err := m.Delete(nil, pp)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestMemManipulator_Count(t *testing.T) {

	Convey("Given I have a memory manipulator and a person", t, func() {

		m := NewMemoryManipulator(Schema)
		p1 := &Person{
			Name: "Antoine1",
		}

		p2 := &Person{
			Name: "Antoine2",
		}

		Convey("When I create the person", func() {

			err := m.Create(nil, p1, p2)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I count the persons", func() {

				n, err := m.Count(nil, PersonIdentity)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then n should equal 1", func() {
					So(n, ShouldEqual, 2)
				})
			})

			Convey("When I delete the person", func() {

				err := m.Delete(nil, p1)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("When I count the persons", func() {

					n, err := m.Count(nil, PersonIdentity)

					Convey("Then err should be nil", func() {
						So(err, ShouldBeNil)
					})

					Convey("Then n should equal 0", func() {
						So(n, ShouldEqual, 1)
					})
				})
			})

			Convey("When I count with a bad filter", func() {

				ctx := manipulate.NewContextWithFilter(
					manipulate.NewFilterComposer().WithKey("Bad").Equals("Antoine1").Done(),
				)

				c, err := m.Count(ctx, PersonIdentity)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
					So(err.(elemental.Error).Code, ShouldEqual, manipulate.ErrCannotExecuteQuery)
				})

				Convey("Then c should equal -1", func() {
					So(c, ShouldEqual, -1)
				})
			})
		})
	})
}

func TestMemManipulator_Assign(t *testing.T) {

	Convey("Given I have a memory manipulator", t, func() {

		m := NewMemoryManipulator(Schema)

		Convey("When I call Assign", func() {

			err := m.Assign(nil, nil)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.(elemental.Error).Code, ShouldEqual, manipulate.ErrNotImplemented)
			})
		})
	})
}

func TestMemManipulator_Increment(t *testing.T) {

	Convey("Given I have a memory manipulator", t, func() {

		m := NewMemoryManipulator(Schema)

		Convey("When I call Increment", func() {

			err := m.Increment(nil, PersonIdentity, "counter", 1)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.(elemental.Error).Code, ShouldEqual, manipulate.ErrNotImplemented)
			})
		})
	})
}

func TestMemManipulator_Commit(t *testing.T) {

	Convey("Given I have a memory manipulator and a transaction ID", t, func() {

		m := NewMemoryManipulator(Schema)
		tid := manipulate.NewTransactionID()

		Convey("When I call Commit with a non existing tid", func() {

			err := m.Commit(tid)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.(elemental.Error).Code, ShouldEqual, manipulate.ErrCannotCommit)
			})
		})

		Convey("When I call Commit with an existing tid", func() {

			m.(*memdbManipulator).registerTxn(tid, m.(*memdbManipulator).db.Txn(true))
			err := m.Commit(tid)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestMemManipulator_Abort(t *testing.T) {

	Convey("Given I have a memory manipulator and a transaction ID", t, func() {

		m := NewMemoryManipulator(Schema)
		tid := manipulate.NewTransactionID()

		Convey("When I call Abort with a non existing tid", func() {

			ok := m.Abort(tid)

			Convey("Then ok should not false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call Commit with an existing tid", func() {

			m.(*memdbManipulator).registerTxn(tid, m.(*memdbManipulator).db.Txn(true))
			ok := m.Abort(tid)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})
	})
}

func TestMemManipulator_txnForID(t *testing.T) {

	Convey("Given I have a memory manipulator and a transaction ID", t, func() {

		m := NewMemoryManipulator(Schema)
		tid := manipulate.NewTransactionID()

		Convey("When I call txnForID with an empty ID", func() {

			txn := m.(*memdbManipulator).txnForID("")

			Convey("Then txn should not be nil", func() {
				So(txn, ShouldNotBeNil)
			})
		})

		Convey("When I call txnForID with an existing ID", func() {

			btxn := m.(*memdbManipulator).db.Txn(true)
			m.(*memdbManipulator).registerTxn(tid, btxn)
			txn := m.(*memdbManipulator).txnForID(tid)

			Convey("Then txn should not be nil", func() {
				So(txn, ShouldEqual, btxn)
			})
		})

		Convey("When I call txnForID with an non existing ID", func() {

			txn := m.(*memdbManipulator).txnForID(tid)

			Convey("Then txn should not be nil", func() {
				So(txn, ShouldNotBeNil)
			})
		})
	})
}
