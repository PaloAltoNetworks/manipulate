package manipmemory

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMemManipulator_Create(t *testing.T) {

	Convey("Given I have a memory manipulator and a person", t, func() {

		m := NewMemoryManipulator(Schema)
		p := &Person{
			Name: "Antoine",
		}

		Convey("When I create person", func() {

			err := m.Create(nil, nil, p)

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
	})
}

func TestMemManipulator_RetrieveChildren(t *testing.T) {

	Convey("Given I have a memory manipulator and a person", t, func() {

		m := NewMemoryManipulator(Schema)
		p1 := &Person{
			Name: "Antoine1",
		}
		p2 := &Person{
			Name: "Antoine2",
		}

		m.Create(nil, nil, p1, p2)

		Convey("When I retrieve the persons", func() {

			ps := []*Person{}

			err := m.RetrieveChildren(nil, nil, PersonIdentity, &ps)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then I should  have retrieved p1 and p2", func() {
				So(ps, ShouldContain, p1)
				So(ps, ShouldContain, p2)
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

			err := m.Create(nil, nil, p)

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

			err := m.Create(nil, nil, p)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("When I delete the person", func() {

				err := m.Delete(nil, p)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("When I retrieve the persons", func() {

					ps := []*Person{}

					err := m.RetrieveChildren(nil, nil, PersonIdentity, &ps)

					Convey("Then err should be nil", func() {
						So(err, ShouldBeNil)
					})

					Convey("Then I the result should be empty", func() {
						So(len(ps), ShouldEqual, 0)
					})
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

			err := m.Create(nil, nil, p1, p2)

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
		})
	})
}
