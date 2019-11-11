// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	data            testmodel.ListsList
	err             error
	stopAtIteration int
	cursor          int
	iteration       int
}

func (m *testManipulator) RetrieveMany(mctx Context, dest elemental.Identifiables) error {

	if m.err != nil {
		return m.err
	}

	if m.stopAtIteration != 0 {
		m.iteration++
		if m.iteration == m.stopAtIteration+1 {
			return nil
		}
	}

	if m.cursor > len(m.data) {
		return nil
	}

	for i, d := range m.data {
		if d.ID == mctx.After() {
			m.cursor = i + 1
		}
	}

	end := m.cursor + mctx.Limit()
	setNext := true
	if end >= len(m.data) {
		end = len(m.data)
		setNext = false
	}

	if setNext {
		if end-1 < 0 {
			mctx.SetNext(m.data[0].ID)
		} else {
			mctx.SetNext(m.data[end-1].ID)
		}
	}

	// fmt.Println("cursor:", m.cursor, "size:", size, "after:", mctx.After())
	*dest.(*testmodel.ListsList) = append(*dest.(*testmodel.ListsList), m.data[m.cursor:end]...)

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

func (m *testManipulator) DeleteMany(mctx Context, identity elemental.Identity) error {
	return nil
}

func (m *testManipulator) Count(mctx Context, identity elemental.Identity) (int, error) {
	return 0, nil
}

func TestDoIterFunc(t *testing.T) {

	Convey("Given I call doIterFunc with no manipulator", t, func() {

		Convey("Then it should panic", func() {
			So(
				func() {
					_ = doIterFunc(nil, nil, nil, nil, nil, 0, false) // nolint
				},
				ShouldPanicWith,
				"manipulator must not be nil",
			)
		})
	})

	Convey("Given I call doIterFunc with no iterator", t, func() {

		Convey("Then it should panic", func() {
			So(
				func() {
					_ = doIterFunc(nil, &testManipulator{}, nil, nil, nil, 0, false) // nolint
				},
				ShouldPanicWith,
				"iteratorFunc must not be nil",
			)
		})
	})

	Convey("Given I call doIterFunc with no identifiablesTemplate", t, func() {

		Convey("Then it should panic", func() {
			So(
				func() {
					_ = doIterFunc(nil, &testManipulator{}, nil, nil, func(elemental.Identifiables) error { return nil }, 0, false) // nolint
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

		Convey("When I call doIterFunc on a round page", func() {

			m := &testManipulator{
				data: makeData(40),
			}

			err := doIterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				10,
				false,
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

		Convey("When I call doIterFunc on a non round page", func() {

			m := &testManipulator{
				data: makeData(45),
			}

			err := doIterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				11,
				false,
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

		Convey("When I call doIterFunc with the default block size", func() {

			m := &testManipulator{
				data: makeData(45),
			}

			err := doIterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				0,
				false,
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

		Convey("When I call doIterFunc but there are no objects", func() {

			m := &testManipulator{
				data: makeData(0),
			}

			err := doIterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				0,
				false,
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

		Convey("When I call doIterFunc but manipulate returns an error", func() {

			m := &testManipulator{
				data: makeData(45),
				err:  fmt.Errorf("boom"),
			}

			err := doIterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				0,
				false,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "unable to retrieve objects for iteration 1: boom")
			})

			Convey("Then the number of calls to iter func should be correct", func() {
				So(called, ShouldEqual, 0)
			})

			Convey("Then the total data count should be correct", func() {
				So(ndata, ShouldEqual, 0)
			})
		})

		Convey("When I call doIterFunc but iter returns an error", func() {

			m := &testManipulator{
				data: makeData(45),
			}

			iter := func(data elemental.Identifiables) error {
				return fmt.Errorf("paf")
			}

			err := doIterFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				iter,
				0,
				false,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "iter function returned an error on iteration 1: paf")
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
				So(err.Error(), ShouldEqual, "unable to retrieve objects for iteration 1: pif")
			})

			Convey("Then dest should be correct", func() {
				So(dest, ShouldBeNil)
			})
		})
	})
}

func TestIterUntilFunc(t *testing.T) {

	Convey("Given I have a manipulator and some objects in the db", t, func() {

		m := &testManipulator{
			data:            makeData(45),
			stopAtIteration: 3,
		}

		Convey("When I call Iter", func() {

			dest := testmodel.ListsList{}
			err := IterUntilFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				func(block elemental.Identifiables) error {
					dest = append(dest, *block.(*testmodel.ListsList)...)
					return nil
				},
				10,
			)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then dest should be correct", func() {
				So(len(dest.List()), ShouldEqual, len(m.data[:30]))
			})
		})
	})

	Convey("Given I have a manipulator and no object in the db", t, func() {

		m := &testManipulator{
			data: makeData(0),
		}

		Convey("When I call Iter", func() {

			dest := testmodel.ListsList{}
			err := IterUntilFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				func(block elemental.Identifiables) error {
					dest = append(dest, *block.(*testmodel.ListsList)...)
					return nil
				},
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

			dest := testmodel.ListsList{}
			err := IterUntilFunc(
				context.Background(),
				m,
				testmodel.ListsList{},
				nil,
				func(block elemental.Identifiables) error {
					dest = append(dest, *block.(*testmodel.ListsList)...)
					return nil
				},
				10,
			)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "unable to retrieve objects for iteration 1: pif")
			})

			Convey("Then dest should be correct", func() {
				So(len(dest.List()), ShouldEqual, 0)
			})
		})
	})
}
