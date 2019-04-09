package manipulate

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
)

func makeData(size int) testmodel.ListsList {
	out := make(testmodel.ListsList, size)
	for i := 0; i < size; i++ {
		out[i] = &testmodel.List{
			ID:   strconv.Itoa(i),
			Name: fmt.Sprintf("list #%d", i),
		}
	}
	return out
}

// A testManipulator is an empty TransactionalManipulator that can be easily mocked.
type testManipulator struct {
	data testmodel.ListsList
	err  error
}

func (m *testManipulator) RetrieveMany(mctx Context, dest elemental.Identifiables) error {

	if m.err != nil {
		return m.err
	}

	start := (mctx.Page() - 1) * mctx.PageSize()
	end := start + mctx.PageSize()

	if start > len(m.data) {
		return nil
	}

	if end > len(m.data) {
		end = len(m.data)
	}
	*dest.(*testmodel.ListsList) = append(*dest.(*testmodel.ListsList), m.data[start:end]...)

	return nil
}

func (m *testManipulator) Retrieve(mctx Context, object elemental.Identifiable) error {
	return nil
}

func (m *testManipulator) Create(mctx Context, object elemental.Identifiable) error {
	return nil
}

func (m *testManipulator) Update(mctx Context, object elemental.Identifiable) error {
	return nil
}

func (m *testManipulator) Delete(mctx Context, object elemental.Identifiable) error {
	return nil
}

// DeleteMany is part of the implementation of the Manipulator interface.
func (m *testManipulator) DeleteMany(mctx Context, identity elemental.Identity) error {
	return nil
}

func (m *testManipulator) Count(mctx Context, identity elemental.Identity) (int, error) {
	return 0, nil
}

func TestIterFunc(t *testing.T) {

	Convey("Given I call IterFunc with no manipulator", t, func() {

		Convey("Then it should panic", func() {
			So(
				func() {
					_ = IterFunc(nil, nil, nil, nil, nil, 0) // nolint
				},
				ShouldPanicWith,
				"manipulator must not be nil",
			)
		})
	})

	Convey("Given I call IterFunc with no iterator", t, func() {

		Convey("Then it should panic", func() {
			So(
				func() {
					_ = IterFunc(nil, &testManipulator{}, nil, nil, nil, 0) // nolint
				},
				ShouldPanicWith,
				"iteratorFunc must not be nil",
			)
		})
	})

	Convey("Given I call IterFunc with no identifiablesTemplate", t, func() {

		Convey("Then it should panic", func() {
			So(
				func() {
					_ = IterFunc(nil, &testManipulator{}, nil, nil, func(elemental.Identifiables) error { return nil }, 0) // nolint
				},
				ShouldPanicWith,
				"identifiablesTemplate must not be nil",
			)
		})
	})

	Convey("Given I have an iter function", t, func() {

		var called int
		var ndata int
		iter := func(data elemental.Identifiables) error {
			called++
			ndata += len(data.List())
			return nil
		}

		Convey("When I call IterFunc on a round page", func() {

			m := &testManipulator{
				data: makeData(40),
			}

			err := IterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				10,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the number of calls to iter func should be correct", func() {
				So(called, ShouldEqual, 4)
			})

			Convey("Then the total data count should be correct", func() {
				So(ndata, ShouldEqual, 40)
			})
		})

		Convey("When I call IterFunc on a non round page", func() {

			m := &testManipulator{
				data: makeData(45),
			}

			err := IterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				11,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the number of calls to iter func should be correct", func() {
				So(called, ShouldEqual, 5)
			})

			Convey("Then the total data count should be correct", func() {
				So(ndata, ShouldEqual, 45)
			})
		})

		Convey("When I call IterFunc with the default block size", func() {

			m := &testManipulator{
				data: makeData(45),
			}

			err := IterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				0,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the number of calls to iter func should be correct", func() {
				So(called, ShouldEqual, 1)
			})

			Convey("Then the total data count should be correct", func() {
				So(ndata, ShouldEqual, 45)
			})
		})

		Convey("When I call IterFunc but there are no objects", func() {

			m := &testManipulator{
				data: makeData(0),
			}

			err := IterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				0,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the number of calls to iter func should be correct", func() {
				So(called, ShouldEqual, 0)
			})

			Convey("Then the total data count should be correct", func() {
				So(ndata, ShouldEqual, 0)
			})
		})

		Convey("When I call IterFunc but manipulate returns an error", func() {

			m := &testManipulator{
				data: makeData(45),
				err:  fmt.Errorf("boom"),
			}

			err := IterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				0,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "unable to retrieve objects for page 1: boom")
			})

			Convey("Then the number of calls to iter func should be correct", func() {
				So(called, ShouldEqual, 0)
			})

			Convey("Then the total data count should be correct", func() {
				So(ndata, ShouldEqual, 0)
			})
		})

		Convey("When I call IterFunc but iter returns an error", func() {

			m := &testManipulator{
				data: makeData(45),
			}

			iter := func(data elemental.Identifiables) error {
				return fmt.Errorf("paf")
			}

			err := IterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				0,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "iter function returned an error on page 1: paf")
			})

			Convey("Then the number of calls to iter func should be correct", func() {
				So(called, ShouldEqual, 0)
			})

			Convey("Then the total data count should be correct", func() {
				So(ndata, ShouldEqual, 0)
			})
		})
	})
}

func TestIter(t *testing.T) {

	Convey("Given I have a manipulator and some objects in the db", t, func() {

		m := &testManipulator{
			data: makeData(45),
		}

		Convey("When I call Iter", func() {

			dest, err := Iter(
				context.Background(),
				m,
				nil,
				testmodel.ListsList{},
				10,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then dest should be correct", func() {
				So(len(dest.List()), ShouldEqual, len(m.data))
				So(dest, ShouldResemble, m.data)
			})
		})
	})

	Convey("Given I have a manipulator and no object in the db", t, func() {

		m := &testManipulator{
			data: makeData(0),
		}

		Convey("When I call Iter", func() {

			dest, err := Iter(
				context.Background(),
				m,
				nil,
				testmodel.ListsList{},
				10,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then dest should be correct", func() {
				So(len(dest.List()), ShouldEqual, 0)
			})
		})
	})

	Convey("Given I have a manipulator but it returns an error", t, func() {

		m := &testManipulator{
			data: makeData(43),
			err:  fmt.Errorf("pif"),
		}

		Convey("When I call Iter", func() {

			dest, err := Iter(
				context.Background(),
				m,
				nil,
				testmodel.ListsList{},
				10,
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "unable to retrieve objects for page 1: pif")
			})

			Convey("Then dest should be correct", func() {
				So(dest, ShouldBeNil)
			})
		})
	})
}
