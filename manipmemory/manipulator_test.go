package manipmemory

import (
	"testing"

	"github.com/aporeto-inc/elemental/test/model"
	"github.com/aporeto-inc/manipulate"

	memdb "github.com/hashicorp/go-memdb"
	. "github.com/smartystreets/goconvey/convey"
)

var Schema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"lists": &memdb.TableSchema{
			Name: "lists",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "ID"},
				},
				"Name": &memdb.IndexSchema{
					Name:    "Name",
					Indexer: &memdb.StringFieldIndex{Field: "Name"},
				},
			},
		},
	},
}

func TestMemManipulator_NewMemoryManipulator(t *testing.T) {

	Convey("Given I create a new MemoryManipulator with bad schema", t, func() {

		Convey("Then it should panic", func() {
			So(func() { NewMemoryManipulator(nil) }, ShouldPanic)
		})
	})
}

func TestMemManipulator_Create(t *testing.T) {

	Convey("Given I have a memory manipulator and a list", t, func() {

		m := NewMemoryManipulator(Schema)
		p := &testmodel.List{
			Name: "Antoine",
		}

		Convey("When I create list", func() {

			err := m.Create(nil, p)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then list ID should not be empty", func() {
				So(p.ID, ShouldNotBeEmpty)
			})

			Convey("When I retrieve the list in a second structure", func() {

				l2 := &testmodel.List{
					ID:   p.ID,
					Name: "not good",
				}

				err := m.Retrieve(nil, l2)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then l2 should be p", func() {
					So(l2, ShouldResemble, p)
				})
			})
		})

		Convey("When I create an object that is not part of the schema", func() {

			err := m.Create(nil, &testmodel.Task{})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestMemManipulator_Retrieve(t *testing.T) {

	Convey("Given I have a memory manipulator and a list", t, func() {

		m := NewMemoryManipulator(Schema)
		l1 := &testmodel.List{
			Name: "Antoine1",
		}

		_ = m.Create(nil, l1)

		Convey("When I retrieve the list", func() {

			ps := &testmodel.List{
				ID: l1.ID,
			}

			err := m.Retrieve(nil, ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should  have retrieved l1 and l2", func() {
				So(ps, ShouldResemble, l1)
			})
		})

		Convey("When I retrieve a non existing list", func() {

			ps := &testmodel.List{
				ID: "not-good",
			}

			err := m.Retrieve(nil, ps)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I retrieve an object that is not part of the schema", func() {

			err := m.Retrieve(nil, &testmodel.Task{})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

	})
}

func TestMemManipulator_RetrieveMany(t *testing.T) {

	Convey("Given I have a memory manipulator and a list", t, func() {

		m := NewMemoryManipulator(Schema)
		l1 := &testmodel.List{
			Name: "Antoine1",
		}
		l2 := &testmodel.List{
			Name: "Antoine2",
		}

		_ = m.Create(nil, l1, l2)

		Convey("When I retrieve the lists", func() {

			ps := testmodel.ListsList{}

			err := m.RetrieveMany(nil, &ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should  have retrieved l1 and l2", func() {
				So(ps, ShouldContain, l1)
				So(ps, ShouldContain, l2)
			})
		})

		Convey("When I retrieve the lists with a filter that matches l1", func() {

			ps := testmodel.ListsList{}

			mctx := manipulate.NewContext()
			mctx.Filter = manipulate.NewFilterComposer().WithKey("Name").Equals("Antoine1").Done()

			err := m.RetrieveMany(mctx, &ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should only have retrieved l1", func() {
				So(ps, ShouldContain, l1)
				So(ps, ShouldNotContain, l2)
			})
		})

		Convey("When I retrieve the lists with a bad filter", func() {

			ps := testmodel.ListsList{}

			mctx := manipulate.NewContext()
			mctx.Filter = manipulate.NewFilterComposer().WithKey("Bad").Equals("Antoine1").Done()

			err := m.RetrieveMany(mctx, &ps)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotExecuteQuery{})
			})
		})
	})
}

func TestMemManipulator_Update(t *testing.T) {

	Convey("Given I have a memory manipulator and a list", t, func() {

		m := NewMemoryManipulator(Schema)
		p := &testmodel.List{
			Name: "Antoine",
		}

		Convey("When I create the list", func() {

			err := m.Create(nil, p)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I update the list", func() {

				p.Name = "New Antoine"

				err := m.Update(nil, p)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("When I retrieve the list", func() {

					l2 := &testmodel.List{
						ID: p.ID,
					}

					err := m.Retrieve(nil, l2)

					Convey("Then err should be nil", func() {
						So(err, ShouldBeNil)
					})

					Convey("Then l2 should contains the updated data", func() {
						So(l2.Name, ShouldEqual, "New Antoine")
					})
				})
			})

			Convey("When I update the a non existing list", func() {

				pp := &testmodel.List{
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

	Convey("Given I have a memory manipulator and a list", t, func() {

		m := NewMemoryManipulator(Schema)
		p := &testmodel.List{
			Name: "Antoine",
		}

		Convey("When I create the list", func() {

			err := m.Create(nil, p)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I delete the list", func() {

				err := m.Delete(nil, p)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("When I retrieve the lists using a wrapper", func() {

					ps := testmodel.ListsList{}

					var err error

					func(zob *testmodel.ListsList) {
						err = m.RetrieveMany(nil, zob)
					}(&ps)

					Convey("Then err should be nil", func() {
						So(err, ShouldBeNil)
					})

					Convey("Then I the result should be empty", func() {
						So(len(ps), ShouldEqual, 0)
					})
				})

				Convey("When I retrieve the lists ", func() {

					ps := testmodel.ListsList{}
					err := m.RetrieveMany(nil, &ps)

					Convey("Then err should be nil", func() {
						So(err, ShouldBeNil)
					})

					Convey("Then I the result should be empty", func() {
						So(len(ps), ShouldEqual, 0)
					})
				})
			})

			Convey("When I delete the a non existing list", func() {

				pp := &testmodel.List{
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

	Convey("Given I have a memory manipulator and a list", t, func() {

		m := NewMemoryManipulator(Schema)
		l1 := &testmodel.List{
			Name: "Antoine1",
		}

		l2 := &testmodel.List{
			Name: "Antoine2",
		}

		Convey("When I create the list", func() {

			err := m.Create(nil, l1, l2)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			// Convey("When I count the lists", func() {
			//
			// 	n, err := m.Count(nil, PersonIdentity)
			//
			// 	Convey("Then err should be nil", func() {
			// 		So(err, ShouldBeNil)
			// 	})
			//
			// 	Convey("Then n should equal 1", func() {
			// 		So(n, ShouldEqual, 2)
			// 	})
			// })

			Convey("When I delete the list", func() {

				err := m.Delete(nil, l1)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				// Convey("When I count the lists", func() {
				//
				// 	n, err := m.Count(nil, PersonIdentity)
				//
				// 	Convey("Then err should be nil", func() {
				// 		So(err, ShouldBeNil)
				// 	})
				//
				// 	Convey("Then n should equal 0", func() {
				// 		So(n, ShouldEqual, 1)
				// 	})
				// })
			})

			// Convey("When I count with a bad filter", func() {
			//
			// 	ctx := manipulate.NewContextWithFilter(
			// 		manipulate.NewFilterComposer().WithKey("Bad").Equals("Antoine1").Done(),
			// 	)
			//
			// 	c, err := m.Count(ctx, PersonIdentity)
			//
			// 	Convey("Then err should not be nil", func() {
			// 		So(err, ShouldNotBeNil)
			// 		So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotExecuteQuery{})
			// 	})
			//
			// 	Convey("Then c should equal -1", func() {
			// 		So(c, ShouldEqual, -1)
			// 	})
			// })
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
				So(err, ShouldHaveSameTypeAs, manipulate.ErrCannotCommit{})
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
