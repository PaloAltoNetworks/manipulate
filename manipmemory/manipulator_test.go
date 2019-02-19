package manipmemory

import (
	"context"
	"crypto/rand"
	"reflect"
	"strconv"
	"testing"

	memdb "github.com/hashicorp/go-memdb"
	. "github.com/smartystreets/goconvey/convey"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate"
)

func datastoreIndexConfig() map[string]*IdentitySchema {

	return map[string]*IdentitySchema{
		testmodel.ListIdentity.Category: &IdentitySchema{
			Identity: testmodel.ListIdentity,
			Indexes: []*Index{
				&Index{
					Name:      "id",
					Type:      IndexTypeString,
					Unique:    true,
					Attribute: "ID",
				},
				&Index{
					Name:      "name",
					Type:      IndexTypeString,
					Unique:    false,
					Attribute: "Name",
				},
				&Index{
					Name:      "slice",
					Type:      IndexTypeSlice,
					Unique:    false,
					Attribute: "Slice",
				},
			},
		},
	}
}

func TestMemManipulator_New(t *testing.T) {

	Convey("Given I create a new MemoryManipulator with bad schema", t, func() {

		Convey("Then it should return err", func() {
			_, err := New(nil)
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Given I create a new MemoryManipulator with a valid schema ", t, func() {

		m, err := New(
			map[string]*IdentitySchema{
				testmodel.ListIdentity.Category: &IdentitySchema{
					Identity: testmodel.ListIdentity,
					Indexes: []*Index{
						&Index{
							Name:      "id",
							Type:      IndexTypeString,
							Unique:    true,
							Attribute: "ID",
						},
						&Index{
							Name:      "Name",
							Type:      IndexTypeString,
							Unique:    false,
							Attribute: "Name",
						},
						&Index{
							Name:      "Slice",
							Type:      IndexTypeSlice,
							Unique:    false,
							Attribute: "Slice",
						},
						&Index{
							Name:      "Map",
							Type:      IndexTypeMap,
							Unique:    false,
							Attribute: "Map",
						},
						&Index{
							Name:      "Bool",
							Type:      IndexTypeBoolean,
							Unique:    false,
							Attribute: "Bool",
						},
						&Index{
							Name:      "StringBased",
							Type:      IndexTypeStringBased,
							Unique:    false,
							Attribute: "StringBased",
						},
					},
				},
			},
		)

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		d := m.(*memdbManipulator)

		Convey("Then the schema should be populated", func() {
			So(d.schema, ShouldNotBeNil)
			So(d.schema.Tables, ShouldNotBeNil)
			So(d.db, ShouldNotBeNil)

		})

		Convey("Then the schema should be correct", func() {
			So(len(d.schema.Tables), ShouldEqual, 1)
			So(d.schema.Tables, ShouldContainKey, testmodel.ListIdentity.Category)
			So(len(d.schema.Tables[testmodel.ListIdentity.Category].Indexes), ShouldEqual, 6)
			So(d.schema.Tables[testmodel.ListIdentity.Category].Indexes["id"],
				ShouldResemble,
				&memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "ID"},
				},
			)
			So(d.schema.Tables[testmodel.ListIdentity.Category].Indexes["Name"],
				ShouldResemble,
				&memdb.IndexSchema{
					Name:    "Name",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Name"},
				},
			)
			So(d.schema.Tables[testmodel.ListIdentity.Category].Indexes["Slice"],
				ShouldResemble,
				&memdb.IndexSchema{
					Name:    "Slice",
					Unique:  false,
					Indexer: &memdb.StringSliceFieldIndex{Field: "Slice"},
				},
			)
			So(d.schema.Tables[testmodel.ListIdentity.Category].Indexes["Map"],
				ShouldResemble,
				&memdb.IndexSchema{
					Name:    "Map",
					Unique:  false,
					Indexer: &memdb.StringMapFieldIndex{Field: "Map"},
				},
			)
			So(d.schema.Tables[testmodel.ListIdentity.Category].Indexes["StringBased"],
				ShouldResemble,
				&memdb.IndexSchema{
					Name:    "StringBased",
					Unique:  false,
					Indexer: &stringBasedFieldIndex{Field: "StringBased"},
				},
			)
			boolIndex := d.schema.Tables[testmodel.ListIdentity.Category].Indexes["Bool"]
			So(boolIndex.Name, ShouldResemble, "Bool")
			So(boolIndex.Unique, ShouldBeFalse)
			So(reflect.TypeOf(boolIndex.Indexer), ShouldEqual, reflect.TypeOf(&memdb.ConditionalIndex{}))

		})
	})
}

func Test_Flush(t *testing.T) {

	Convey("Given a valid data store", t, func() {

		m, _ := New(datastoreIndexConfig())

		d := m.(*memdbManipulator)

		Convey("When I flush it", func() {

			oldDb := d.db
			err := d.Flush(context.Background())

			Convey("Then there should be no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("The DB should be all new", func() {
				So(oldDb, ShouldNotResemble, d.db)
			})
		})
	})
}

func TestMemManipulator_Create(t *testing.T) {

	Convey("Given I have a memory manipulator and a list", t, func() {

		m, err := New(datastoreIndexConfig())
		So(err, ShouldBeNil)
		p := &testmodel.List{
			Name:  "Antoine",
			Slice: []string{"$names=antoine"},
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

		m, err := New(datastoreIndexConfig())
		So(err, ShouldBeNil)
		l1 := &testmodel.List{
			Name:  "Antoine1",
			Slice: []string{"$name=Antoine1"},
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

		m, err := New(datastoreIndexConfig())
		So(err, ShouldBeNil)
		l1 := &testmodel.List{
			Name:  "Antoine1",
			Slice: []string{"$name=antoine1", "category=antoine", "a=b", "c=d"},
		}
		l2 := &testmodel.List{
			Name:  "Antoine2",
			Slice: []string{"$name=antoine2", "category=antoine", "x=y", "w=z"},
		}
		l3 := &testmodel.List{
			Name:  "Dimitri1",
			Slice: []string{"$name=dimitri1", "category=dimitri", "a=b", "x=y"},
		}
		l4 := &testmodel.List{
			Name:  "Dimitri2",
			Slice: []string{"$name=dimitri2", "category=dimitri", "a=b", "x=y"},
		}

		_ = m.Create(nil, l1, l2, l3, l4)

		Convey("When I retrieve the lists", func() {

			ps := testmodel.ListsList{}

			err := m.RetrieveMany(nil, &ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should  have retrieved all the items", func() {
				So(len(ps), ShouldEqual, 4)
				So(ps, ShouldContain, l1)
				So(ps, ShouldContain, l2)
				So(ps, ShouldContain, l3)
				So(ps, ShouldContain, l4)
			})
		})

		Convey("When I retrieve the lists with a filter that matches l1", func() {

			ps := testmodel.ListsList{}

			mctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionFilter(
					manipulate.NewFilterComposer().WithKey("Name").Equals("Antoine1").
						WithKey("Slice").Contains("a=b").Done(),
				),
			)

			err := m.RetrieveMany(mctx, &ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should only have retrieved l1", func() {
				So(len(ps), ShouldEqual, 1)
				So(ps, ShouldContain, l1)
				So(ps, ShouldNotContain, l2)
			})
		})

		Convey("When I retrieve the lists with an OR filter that matches l1 and l2", func() {

			ps := testmodel.ListsList{}

			filter := manipulate.NewFilterComposer().Or(
				manipulate.NewFilterComposer().
					WithKey("Name").Equals("Antoine1").Done(),
				manipulate.NewFilterComposer().
					WithKey("Name").Equals("Antoine2").Done(),
			).Done()

			mctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionFilter(filter),
			)

			err := m.RetrieveMany(mctx, &ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should have both items in the list", func() {
				So(len(ps), ShouldEqual, 2)
				So(ps, ShouldContain, l1)
				So(ps, ShouldContain, l2)
			})
		})

		Convey("When I retrieve the lists with an AND filter that matches l3 and l4", func() {

			ps := testmodel.ListsList{}

			filter := manipulate.NewFilterComposer().And(
				manipulate.NewFilterComposer().
					WithKey("Slice").Equals("category=dimitri").Done(),
				manipulate.NewFilterComposer().
					WithKey("Slice").Equals("a=b").Done(),
			).Done()

			mctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionFilter(filter),
			)

			err := m.RetrieveMany(mctx, &ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should have two items in the list", func() {
				So(len(ps), ShouldEqual, 2)
				So(ps, ShouldContain, l3)
				So(ps, ShouldContain, l4)
			})
		})

		Convey("When I retrieve the lists with the Contains comparator", func() {

			ps := testmodel.ListsList{}

			filter := manipulate.NewFilterComposer().
				WithKey("Slice").Contains("category=dimitri", "a=b").Done()

			mctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionFilter(filter),
			)

			err := m.RetrieveMany(mctx, &ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should have two items in the list", func() {
				So(len(ps), ShouldEqual, 3)
				So(ps, ShouldContain, l1)
				So(ps, ShouldContain, l3)
				So(ps, ShouldContain, l4)
			})
		})

		Convey("When I retrieve the lists with an OR of Contains, I should get four items", func() {

			ps := testmodel.ListsList{}

			filter := manipulate.NewFilterComposer().Or(
				manipulate.NewFilterComposer().
					WithKey("Slice").Contains("category=dimitri", "a=b").Done(),
				manipulate.NewFilterComposer().
					WithKey("Slice").Contains("category=antoine").Done(),
				manipulate.NewFilterComposer().
					WithKey("Slice").Contains("x=y").Done(),
			).Done()

			mctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionFilter(filter),
			)

			err := m.RetrieveMany(mctx, &ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should have two items in the list", func() {
				So(len(ps), ShouldEqual, 4)
				So(ps, ShouldContain, l1)
				So(ps, ShouldContain, l2)
				So(ps, ShouldContain, l3)
				So(ps, ShouldContain, l4)
			})
		})

		Convey("When I retrieve the lists with a bad filter", func() {

			ps := testmodel.ListsList{}

			mctx := manipulate.NewContext(
				context.Background(),
				manipulate.ContextOptionFilter(
					manipulate.NewFilterComposer().WithKey("Bad").Equals("Antoine1").Done(),
				),
			)

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

		m, err := New(datastoreIndexConfig())
		So(err, ShouldBeNil)
		p := &testmodel.List{
			Name:  "Antoine",
			Slice: []string{"$names=antoine"},
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

		m, err := New(datastoreIndexConfig())
		So(err, ShouldBeNil)
		p := &testmodel.List{
			Name:  "Antoine",
			Slice: []string{"$name=antoine"},
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

		m, err := New(datastoreIndexConfig())
		So(err, ShouldBeNil)
		l1 := &testmodel.List{
			Name:  "Antoine1",
			Slice: []string{"$names=antoine1"},
		}

		l2 := &testmodel.List{
			Name:  "Antoine2",
			Slice: []string{"$name=antoine2"},
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

		m, err := New(datastoreIndexConfig())
		So(err, ShouldBeNil)
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

		m, err := New(datastoreIndexConfig())
		So(err, ShouldBeNil)
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

		m, err := New(datastoreIndexConfig())
		So(err, ShouldBeNil)
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

func BenchmarkRetrieveMany(b *testing.B) {
	b.StopTimer()

	m, err := New(datastoreIndexConfig())
	So(err, ShouldBeNil)
	err = populateDB(m, 10000)
	So(err, ShouldBeNil)

	filters := []*manipulate.Filter{
		manipulate.NewFilterComposer().
			WithKey("Slice").Contains("label1=10", "label2=10").
			Done(),
		manipulate.NewFilterComposer().
			WithKey("Slice").Contains("label1=10", "label2=10", "label=common").
			Done(),
	}
	b.StartTimer()

	for f, filter := range filters {
		b.Run("Filter: "+strconv.Itoa(f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {

				list := testmodel.ListsList{}

				mctx := manipulate.NewContext(
					context.Background(),
					manipulate.ContextOptionFilter(filter),
				)
				if err := m.RetrieveMany(mctx, &list); err != nil {
					b.Errorf("Error in retrieve many: %s", err.Error())
					b.FailNow()
				}
				if len(list) != 1 {
					b.Errorf("Length of list is wrong: %d %d %d", len(list), i, b.N)
					b.FailNow()
				}
			}
		})
	}
}

func populateDB(m manipulate.TransactionalManipulator, num int) error {

	for i := 0; i < num; i++ {

		iString := strconv.Itoa(i)

		l := testmodel.NewList()

		l.Name = "name" + iString

		l.Slice = []string{"label1=" + iString, "label2=" + iString, "label=common"}

		for j := 0; j < 14; j++ {
			b := make([]byte, 32)
			if _, err := rand.Read(b); err != nil {
				return err
			}
			l.Slice = append(l.Slice, string(b))
		}

		if err := m.Create(nil, l); err != nil {
			return err
		}
	}

	return nil
}
