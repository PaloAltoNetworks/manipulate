package memdbvortex

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniptest"
	"go.aporeto.io/manipulate/manipvortex/config"
)

func newObject(name string, tags []string) *testmodel.List {
	o := testmodel.NewList()
	o.Name = name
	o.Slice = tags

	return o
}

func newDatastore() (*MemdbDatastore, error) {

	config := map[string]*config.MemDBIdentity{
		testmodel.ListIdentity.Category: &config.MemDBIdentity{
			Identity: testmodel.ListIdentity,
			Indexes: []*config.IndexConfig{
				&config.IndexConfig{
					Name:      "id",
					Type:      config.String,
					Unique:    true,
					Attribute: "ID",
				},
				&config.IndexConfig{
					Name:      "Name",
					Type:      config.String,
					Unique:    false,
					Attribute: "Name",
				},
				&config.IndexConfig{
					Name:      "Slice",
					Type:      config.Slice,
					Unique:    false,
					Attribute: "Slice",
				},
			},
		},
	}

	d, err := NewDatastore(config)
	if err != nil {
		return nil, err
	}

	if err := d.Run(); err != nil {
		return nil, err
	}

	return d, nil
}

func newIdentityProcessor(mode config.CacheMode) map[string]*config.ProcessorConfiguration {

	return map[string]*config.ProcessorConfiguration{
		testmodel.ListIdentity.Name: &config.ProcessorConfiguration{
			Identity:         testmodel.ListIdentity,
			Mode:             mode,
			QueueingDuration: time.Minute,
			CommitOnEvent:    true,
		},
	}
}

func Test_NewMemDBVortex(t *testing.T) {
	t.Parallel()

	Convey("When I create a new memdb vortex, I sould have correct structures", t, func() {
		m := maniptest.NewTestManipulator()
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")
		So(v, ShouldNotBeNil)
		So(v, ShouldHaveSameTypeAs, &MemDBVortex{})

		mv := v.(*MemDBVortex)

		So(mv.m, ShouldResemble, m)
		So(mv.memory, ShouldBeNil)
		So(mv.datastore, ShouldNotBeNil)
		So(mv.processors, ShouldNotBeNil)
		So(mv.transactionQueue, ShouldNotBeNil)

	})
}

func Test_UnsupportedMethods(t *testing.T) {
	t.Parallel()
	Convey("Given a new memdb vortex with no backend", t, func() {

		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, nil, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("When I try to delete many objects, I should get an error", func() {
			err := v.DeleteMany(nil, testmodel.ListIdentity)
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Given a new memdb vortex with a backend", t, func() {
		m := maniptest.NewTestManipulator()
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		Convey("When I try to delete many objects, the transaction must be forwarded", func() {
			m.MockDeleteMany(t, func(mctx manipulate.Context, identity elemental.Identity) error {
				return nil
			})
			err := v.DeleteMany(nil, testmodel.ListIdentity)
			So(err, ShouldBeNil)
		})

	})
}

func Test_Count(t *testing.T) {

	t.Parallel()
	Convey("Given a new memdb vortex with no backend", t, func() {
		m := maniptest.NewTestManipulator()
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("When I try to count the objects, I should get an error", func() {
			n, err := v.Count(nil, testmodel.ListIdentity)
			So(err, ShouldBeNil)
			So(n, ShouldEqual, 0)
		})

		Convey("If the data store is not initialized it should return err", func() {
			n, err := v.Count(nil, testmodel.ListIdentity)
			So(err, ShouldBeNil)
			So(n, ShouldEqual, 0)
		})
	})

}

func Test_Run(t *testing.T) {
	t.Parallel()
	Convey("Given a new memdb vortex with no backend", t, func() {
		m := maniptest.NewTestManipulator()
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		Convey("When I try to run it, and it is already running, it should error", func() {
			ctx, cancel := context.WithCancel(context.Background())
			v.(*MemDBVortex).started = true
			err := v.Run(ctx)
			So(err, ShouldNotBeNil)
			cancel()
		})

		Convey("When I try to run it, with a bad datastore, it should error", func() {
			ctx, cancel := context.WithCancel(context.Background())
			v.(*MemDBVortex).datastore = nil
			err := v.Run(ctx)
			So(err, ShouldNotBeNil)
			cancel()
		})

		Convey("When I try to run it, it should succeed", func() {
			ctx, cancel := context.WithCancel(context.Background())
			err := v.Run(ctx)
			So(err, ShouldBeNil)
			So(v.(*MemDBVortex).started, ShouldBeTrue)
			cancel()
		})

	})

	Convey("Given a new memdb vortex with log enabled", t, func() {
		m := maniptest.NewTestManipulator()
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "./testlog")

		defer os.Remove("./testlog") // nolint errcheck

		Convey("When I try to run it, it should succeed and the log channel should be there", func() {
			ctx, cancel := context.WithCancel(context.Background())
			err := v.Run(ctx)
			So(err, ShouldBeNil)
			So(v.(*MemDBVortex).started, ShouldBeTrue)
			So(v.(*MemDBVortex).logChannel, ShouldNotBeNil)
			cancel()
		})
	})

	Convey("Given a new memdb vortex with log enabled and a bad file", t, func() {
		m := maniptest.NewTestManipulator()
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "./bad-directory/test")

		Convey("When I try to run it, it should fail", func() {
			ctx, cancel := context.WithCancel(context.Background())
			err := v.Run(ctx)
			So(err, ShouldNotBeNil)
			cancel()
		})
	})

	Convey("Given a new memdb vortex with a backend", t, func() {
		m := maniptest.NewTestManipulator()
		s := maniptest.NewTestSubscriber()
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, m, s, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		Convey("When I try to run it and the bakend fails in the resync, it should error", func() {
			m.MockRetrieveMany(t, func(mctx manipulate.Context, dest elemental.Identifiables) error {
				return manipulate.NewErrObjectNotFound("testing")
			})
			ctx, cancel := context.WithCancel(context.Background())
			err := v.Run(ctx)
			So(err, ShouldNotBeNil)
			So(v.(*MemDBVortex).started, ShouldBeFalse)
			cancel()
		})

		Convey("When the backend succeds it should run", func() {
			m.MockRetrieveMany(t, func(mctx manipulate.Context, dest elemental.Identifiables) error {
				return nil
			})
			ctx, cancel := context.WithCancel(context.Background())
			err := v.Run(ctx)
			So(err, ShouldBeNil)
			So(v.(*MemDBVortex).started, ShouldBeTrue)
			cancel()
		})

	})
}

func Test_Resync(t *testing.T) {

	t.Parallel()

	Convey("Given a new memdb vortex with no backend", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, nil, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		Convey("When I try to Re-sync it, nothing should happen", func() {
			ctx, cancel := context.WithCancel(context.Background())
			err := v.Run(ctx)
			So(err, ShouldBeNil)
			err = v.ReSync(ctx)
			So(err, ShouldBeNil)
			cancel()
		})
	})

	Convey("Given a new memdb vortex with a backend and objects", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)
		m := maniptest.NewTestManipulator()
		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		obj1 := newObject("obj1", []string{"a=b"})
		obj1.ID = "ID1"
		obj2 := newObject("obj2", []string{"x=y"})
		obj2.ID = "ID2"

		retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
			if mctx.Page() > 1 {
				return nil
			}
			objects := testmodel.ListsList{obj1, obj2}
			*dest.(*testmodel.ListsList) = objects
			return nil
		}

		m.MockRetrieveMany(t, retrieveRespose)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("When I try to Re-sync it with no data, the db should be empty", func() {
			retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
				return nil
			}
			m.MockRetrieveMany(t, retrieveRespose)
			err := v.ReSync(ctx)
			So(err, ShouldBeNil)

			objects := testmodel.ListsList{}
			err = v.RetrieveMany(nil, &objects)
			So(err, ShouldBeNil)
			So(len(objects), ShouldEqual, 0)
		})
	})

	Convey("Given a new memdb vortex with a backend ", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)
		m := maniptest.NewTestManipulator()
		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
			return nil
		}

		m.MockRetrieveMany(t, retrieveRespose)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("When I try to Re-sync it with and the schema is bad, it should fail", func() {
			retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
				return nil
			}
			m.MockRetrieveMany(t, retrieveRespose)
			v.(*MemDBVortex).datastore = nil
			err := v.ReSync(ctx)
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Given a new memdb vortex with a backend, where the backend fails", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)

		m := maniptest.NewTestManipulator()
		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
			return manipulate.NewErrObjectNotFound("error test")
		}

		m.MockRetrieveMany(t, retrieveRespose)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err = v.Run(ctx)
		So(err, ShouldNotBeNil)

	})
}

func Test_RetrieveMany(t *testing.T) {

	t.Parallel()

	Convey("Given a new memdb vortex with a backend and a hook", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)
		m := maniptest.NewTestManipulator()
		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		objConfig := v.(*MemDBVortex).processors[testmodel.ListIdentity.Name]
		objConfig.RetrieveManyHook = func(m manipulate.Manipulator, mctx manipulate.Context, dest elemental.Identifiables) (bool, error) {
			if mctx.Parent() != nil {
				return false, fmt.Errorf("no parent")
			}
			if mctx.Page() > 1 {
				return false, nil
			}
			return true, nil
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		obj1 := newObject("obj1", []string{"a=b"})
		obj1.ID = "ID1"
		obj2 := newObject("obj2", []string{"x=y"})
		obj2.ID = "ID2"

		retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
			if mctx.Page() > 1 {
				return nil
			}
			objects := testmodel.ListsList{obj1, obj2}
			*dest.(*testmodel.ListsList) = objects
			return nil
		}

		m.MockRetrieveMany(t, retrieveRespose)
		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("When I request a retrieve many with a parent, it should go the backend only", func() {
			mctx := manipulate.NewContext(ctx, manipulate.ContextOptionParent(&testmodel.List{}))
			objects := testmodel.ListsList{}
			m.MockRetrieveMany(t, func(mctx manipulate.Context, dest elemental.Identifiables) error {
				return nil
			})
			err := v.RetrieveMany(mctx, &objects)
			So(err, ShouldBeNil)
		})

		Convey("When I request a retrieve many with no parent for first page, it should retrieve the data", func() {
			mctx := manipulate.NewContext(ctx)
			objects := testmodel.ListsList{}
			err := v.RetrieveMany(mctx, &objects)
			So(err, ShouldBeNil)
			So(len(objects), ShouldEqual, 2)
		})

		Convey("When I request a retrieve many with no parent second page, it should retrieve no data", func() {
			mctx := manipulate.NewContext(ctx, manipulate.ContextOptionPage(2, 100))
			objects := testmodel.ListsList{}
			err := v.RetrieveMany(mctx, &objects)
			So(err, ShouldBeNil)
			So(len(objects), ShouldEqual, 0)
		})
	})
}

func Test_Retrieve(t *testing.T) {

	t.Parallel()

	Convey("Given a new memdb vortex with data and backend", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)
		m := maniptest.NewTestManipulator()
		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		obj1 := newObject("obj1", []string{"a=b"})
		obj1.ID = "ID1"
		obj2 := newObject("obj2", []string{"x=y"})
		obj2.ID = "ID2"

		retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
			if mctx.Page() > 1 {
				return nil
			}
			objects := testmodel.ListsList{obj1, obj2}
			*dest.(*testmodel.ListsList) = objects
			return nil
		}
		m.MockRetrieveMany(t, retrieveRespose)

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("When I read a valid object, I should get no error and the object", func() {
			o := newObject("", []string{})
			o.ID = "ID1"

			err := v.Retrieve(nil, o)
			So(err, ShouldBeNil)
			So(o, ShouldResemble, obj1)
		})

		Convey("When I read an invalid object, with no consistency, it should error", func() {
			o := newObject("", []string{})
			o.ID = "bad-id"

			err := v.Retrieve(nil, o)
			So(err, ShouldNotBeNil)
		})

		Convey("When I read an invalid object, with consistency and the backend fails it should error", func() {
			o := newObject("", []string{})
			o.ID = "bad-id"

			mctx := manipulate.NewContext(ctx, manipulate.ContextOptionReadConsistency(manipulate.ReadConsistencyStrong))
			m.MockRetrieve(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
				return manipulate.NewErrCannotCommunicate("test")
			})
			err := v.Retrieve(mctx, o)
			So(err, ShouldNotBeNil)
			So(manipulate.IsCannotCommunicateError(err), ShouldBeTrue)
		})

		Convey("When I read an invalid object, with consistency and the backend succeeds it should succeed", func() {
			o := newObject("someobject", []string{"a=b"})
			o.ID = "bad-id"

			mctx := manipulate.NewContext(ctx, manipulate.ContextOptionReadConsistency(manipulate.ReadConsistencyStrong))
			m.MockRetrieve(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
				return nil
			})
			err := v.Retrieve(mctx, o)
			So(err, ShouldBeNil)

			Convey("... and the object must be now in the db", func() {
				o := newObject("", []string{})
				o.ID = "bad-id"
				err := v.Retrieve(nil, o)
				So(err, ShouldBeNil)
			})
		})

		Convey("When I read an invalid object, with consistency and the backend succeeds but cache fails", func() {
			o := newObject("", []string{""})
			o.ID = "bad-id"

			mctx := manipulate.NewContext(ctx, manipulate.ContextOptionReadConsistency(manipulate.ReadConsistencyStrong))
			m.MockRetrieve(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
				return nil
			})
			err := v.Retrieve(mctx, o)
			So(err, ShouldNotBeNil)
		})
	})

}

func Test_Create(t *testing.T) {

	Convey("Give a new memdb vortex with a backend", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)
		m := maniptest.NewTestManipulator()
		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		objConfig := v.(*MemDBVortex).processors[testmodel.ListIdentity.Name]
		objConfig.UpstreamHook = func(method elemental.Operation, mctx manipulate.Context, objects []elemental.Identifiable) (bool, error) {
			if mctx.Parent() != nil {
				return false, nil
			}
			return true, nil
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("When I create objects", func() {

			Convey("When the backend succeeds, the object must be stored in the DB", func() {
				m.MockCreate(t, func(mctx manipulate.Context, objects ...elemental.Identifiable) error {
					o := objects[0].(*testmodel.List)
					o.ID = "ID1"
					return nil
				})

				obj := newObject("obj1", []string{"label"})

				err := v.Create(nil, obj)
				So(err, ShouldBeNil)
				So(obj.ID, ShouldResemble, "ID1")

				newObject := newObject("", []string{})
				newObject.ID = "ID1"

				err = v.Retrieve(nil, newObject)
				So(err, ShouldBeNil)
				So(newObject, ShouldResemble, obj)
			})

			Convey("When the backend fails, the object must be not be stored in the DB", func() {
				m.MockCreate(t, func(mctx manipulate.Context, objects ...elemental.Identifiable) error {
					return manipulate.NewErrConstraintViolation("test")
				})

				obj := newObject("obj1", []string{"label"})
				obj.ID = "obj1"

				err := v.Create(nil, obj)
				So(err, ShouldNotBeNil)

				newObject := newObject("", []string{})
				newObject.ID = "ID1"

				err = v.Retrieve(nil, newObject)
				So(err, ShouldNotBeNil)
			})

			Convey("When the has a hook function that prevents execution, it should not commit to the DB", func() {

				m.MockCreate(t, func(mctx manipulate.Context, objects ...elemental.Identifiable) error {
					return manipulate.NewErrConstraintViolation("test")
				})

				obj := newObject("obj1", []string{"label"})
				obj.ID = "obj1"

				mctx := manipulate.NewContext(ctx, manipulate.ContextOptionParent(&testmodel.List{}))

				err := v.Create(mctx, obj)
				So(err, ShouldBeNil)

				newObject := newObject("", []string{})
				newObject.ID = "ID1"

				err = v.Retrieve(nil, newObject)
				So(err, ShouldNotBeNil)
			})

		})

	})
}

func Test_Update(t *testing.T) {

	t.Parallel()

	Convey("Give a new memdb vortex with a backend with an object", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)
		m := maniptest.NewTestManipulator()
		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		objConfig := v.(*MemDBVortex).processors[testmodel.ListIdentity.Name]
		objConfig.UpstreamHook = func(method elemental.Operation, mctx manipulate.Context, objects []elemental.Identifiable) (bool, error) {
			if mctx.Parent() != nil {
				return false, nil
			}
			return true, nil
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		obj := newObject("obj1", []string{"label"})
		obj.ID = "obj1"

		m.MockCreate(t, func(mctx manipulate.Context, objects ...elemental.Identifiable) error {
			o := objects[0].(*testmodel.List)
			o.ID = "ID1"
			return nil
		})

		err = v.Create(nil, obj)
		So(err, ShouldBeNil)

		Convey("When I update the object", func() {

			Convey("When the backend succeeds, the object must be stored in the DB", func() {
				m.MockUpdate(t, func(mctx manipulate.Context, objects ...elemental.Identifiable) error {
					return nil
				})

				updatedObject := newObject("", []string{"a=b"})
				updatedObject.Name = "test"
				updatedObject.ID = "ID1"
				err = v.Update(nil, updatedObject)
				So(err, ShouldBeNil)

				readobject := newObject("", []string{})
				readobject.ID = "ID1"

				err = v.Retrieve(nil, readobject)
				So(err, ShouldBeNil)
				So(updatedObject.Name, ShouldResemble, readobject.Name)
			})

			Convey("When the backend fails, the object must not be updated", func() {
				m.MockUpdate(t, func(mctx manipulate.Context, objects ...elemental.Identifiable) error {
					return manipulate.NewErrCannotBuildQuery("test")
				})

				updatedObject := newObject("", []string{"a=b"})
				updatedObject.Name = "test"
				updatedObject.ID = "ID1"
				err = v.Update(nil, updatedObject)
				So(err, ShouldNotBeNil)

				readobject := newObject("", []string{})
				readobject.ID = "ID1"

				err = v.Retrieve(nil, readobject)
				So(err, ShouldBeNil)
				So(updatedObject.Name, ShouldNotResemble, readobject.Name)
			})

			Convey("When the vortex has a hook function that prevents updates, the object must not be updated", func() {
				m.MockUpdate(t, func(mctx manipulate.Context, objects ...elemental.Identifiable) error {
					return manipulate.NewErrCannotBuildQuery("test")
				})

				updatedObject := newObject("", []string{"a=b"})
				updatedObject.Name = "test"
				updatedObject.ID = "ID1"

				mctx := manipulate.NewContext(ctx, manipulate.ContextOptionParent(&testmodel.List{}))
				err = v.Update(mctx, updatedObject)
				So(err, ShouldBeNil)

				readobject := newObject("", []string{})
				readobject.ID = "ID1"

				err = v.Retrieve(nil, readobject)
				So(err, ShouldBeNil)
				So(updatedObject.Name, ShouldNotResemble, readobject.Name)
			})
		})

	})
}

func Test_Delete(t *testing.T) {

	t.Parallel()

	Convey("Give a new memdb vortex with a backend with an object", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)
		m := maniptest.NewTestManipulator()
		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		objConfig := v.(*MemDBVortex).processors[testmodel.ListIdentity.Name]
		objConfig.UpstreamHook = func(method elemental.Operation, mctx manipulate.Context, objects []elemental.Identifiable) (bool, error) {
			if mctx.Parent() != nil {
				return false, nil
			}
			return true, nil
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		obj := newObject("obj1", []string{"label"})
		obj.ID = "obj1"

		m.MockCreate(t, func(mctx manipulate.Context, objects ...elemental.Identifiable) error {
			o := objects[0].(*testmodel.List)
			o.ID = "ID1"
			return nil
		})

		err = v.Create(nil, obj)
		So(err, ShouldBeNil)

		Convey("When I delete the object", func() {

			Convey("When the backend succeeds, the object must be deleted", func() {
				m.MockDelete(t, func(mctx manipulate.Context, objects ...elemental.Identifiable) error {
					return nil
				})

				updatedObject := newObject("", []string{})
				updatedObject.ID = "ID1"
				err = v.Delete(nil, updatedObject)
				So(err, ShouldBeNil)

				readobject := newObject("", []string{})
				readobject.ID = "ID1"

				err = v.Retrieve(nil, readobject)
				So(err, ShouldNotBeNil)
			})

			Convey("When the backend fail, the object must not be deleted", func() {
				m.MockDelete(t, func(mctx manipulate.Context, objects ...elemental.Identifiable) error {
					return manipulate.NewErrCannotBuildQuery("test")
				})

				updatedObject := newObject("", []string{"a=b"})
				updatedObject.ID = "ID1"
				err = v.Delete(nil, updatedObject)
				So(err, ShouldNotBeNil)

				readobject := newObject("", []string{})
				readobject.ID = "ID1"

				err = v.Retrieve(nil, readobject)
				So(err, ShouldBeNil)
			})

			Convey("When the vortex has a hook function that prevents deletes, the object must not be deleted", func() {
				m.MockDelete(t, func(mctx manipulate.Context, objects ...elemental.Identifiable) error {
					return manipulate.NewErrCannotBuildQuery("test")
				})

				updatedObject := newObject("", []string{"a=b"})
				updatedObject.ID = "ID1"

				mctx := manipulate.NewContext(ctx, manipulate.ContextOptionParent(&testmodel.List{}))
				err = v.Delete(mctx, updatedObject)
				So(err, ShouldBeNil)

				readobject := newObject("", []string{})
				readobject.ID = "ID1"

				err = v.Retrieve(nil, readobject)
				So(err, ShouldBeNil)
			})
		})

	})
}

func Test_WithNoBackend(t *testing.T) {

	t.Parallel()

	Convey("Given a new memdb vortext with no backend", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, nil, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("If I try to resync with no backend, I should get no error", func() {
			err := v.ReSync(ctx)
			So(err, ShouldBeNil)
		})

		obj1 := newObject("obj1", []string{"a=b", "c=de", "common"})
		obj2 := newObject("obj2", []string{"x=y", "w=z", "common"})

		Convey("When I create an objects I should get no errors", func() {
			err := v.Create(nil, obj1, obj2)
			So(err, ShouldBeNil)

			Convey("When I retrieve the objects I created with retrieve many", func() {
				objects := testmodel.ListsList{}
				err := v.RetrieveMany(nil, &objects)
				So(err, ShouldBeNil)
				So(len(objects), ShouldEqual, 2)
				So(objects, ShouldContain, obj1)
				So(objects, ShouldContain, obj2)
			})

			Convey("When I retrieve a specific object with retrieve", func() {
				newObject := newObject("", []string{})
				newObject.ID = obj1.ID

				err := v.Retrieve(nil, newObject)
				So(err, ShouldBeNil)
				So(newObject, ShouldResemble, obj1)
			})

			Convey("When I update one of the objects", func() {
				obj1.Name = "newobject1"

				err := v.Update(nil, obj1)
				So(err, ShouldBeNil)

				Convey("I should read an updated object", func() {
					newObject := newObject("", []string{})
					newObject.ID = obj1.ID

					err := v.Retrieve(nil, newObject)
					So(err, ShouldBeNil)
					So(newObject, ShouldResemble, obj1)
					So(newObject.Name, ShouldResemble, "newobject1")
				})
			})

			Convey("When I delete one of the objects, it should be deleted", func() {
				err := v.Delete(nil, obj1)
				So(err, ShouldBeNil)

				Convey("The DB must only have one object", func() {
					objects := testmodel.ListsList{}
					err := v.RetrieveMany(nil, &objects)
					So(err, ShouldBeNil)
					So(len(objects), ShouldEqual, 1)
					So(objects, ShouldContain, obj2)
				})
			})

			Convey("When I flush the cache, it should have no objects", func() {
				err := v.Flush(context.Background())
				So(err, ShouldBeNil)
				objects := testmodel.ListsList{}
				err = v.RetrieveMany(nil, &objects)
				So(err, ShouldBeNil)
				So(len(objects), ShouldEqual, 0)

			})
		})
	})
}

func Test_WriteThroughBackend(t *testing.T) {

	t.Parallel()

	Convey("Given a memdb vortext with a write-through mode", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)

		m := maniptest.NewTestManipulator()
		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		obj1 := newObject("obj1", []string{"a=b"})
		obj1.ID = "ID1"
		obj2 := newObject("obj2", []string{"x=y"})
		obj2.ID = "ID2"

		retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
			if mctx.Page() > 1 {
				return nil
			}
			objects := testmodel.ListsList{obj1, obj2}
			*dest.(*testmodel.ListsList) = objects
			return nil
		}
		m.MockRetrieveMany(t, retrieveRespose)

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("When I retrieve the objects that are loaded after the sync, I should get the right objects", func() {
			objects := testmodel.ListsList{}
			err := v.RetrieveMany(nil, &objects)
			So(err, ShouldBeNil)
			So(len(objects), ShouldEqual, 2)
			So(objects, ShouldContain, obj1)
			So(objects, ShouldContain, obj2)
		})

		Convey("When I create a new object", func() {

			Convey("When the backend fails, the creation should fail", func() {
				obj3 := newObject("obj3", []string{"w=z"})

				m.MockCreate(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
					return fmt.Errorf("backend failed")
				})

				err := v.Create(nil, obj3)
				So(err, ShouldNotBeNil)

				objects := testmodel.ListsList{}
				err = v.RetrieveMany(nil, &objects)
				So(err, ShouldBeNil)
				So(len(objects), ShouldEqual, 2)
			})

			obj3 := newObject("obj3", []string{"w=z"})

			m.MockCreate(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
				o := objects[0]
				o.SetIdentifier("ID3")
				return nil
			})

			err := v.Create(nil, obj3)
			So(err, ShouldBeNil)
			So(obj3.Identifier(), ShouldResemble, "ID3")

			Convey("When I retrieve the object from the local db", func() {

				newObject := &testmodel.List{
					ID: "ID3",
				}
				err := v.Retrieve(nil, newObject)
				So(err, ShouldBeNil)
				So(newObject, ShouldResemble, obj3)
			})

			Convey("When I update the object, it must be updated in both DBs", func() {

				Convey("When the backend fails, the update should fail", func() {
					updatedObj := testmodel.NewList()
					updatedObj.Name = "newname"
					updatedObj.ID = obj3.ID

					m.MockUpdate(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
						return fmt.Errorf("backend is not there")
					})

					err := v.Update(nil, updatedObj)
					So(err, ShouldNotBeNil)

					newObject := &testmodel.List{
						ID: "ID3",
					}
					err = v.Retrieve(nil, newObject)
					So(err, ShouldBeNil)
					So(newObject.Name, ShouldNotResemble, "newname")
				})

				Convey("When the backend succeeds the object must be updated", func() {

					updatedObj := testmodel.NewList()
					updatedObj.Name = "newname"
					updatedObj.ID = obj3.ID
					updatedObj.Slice = obj3.Slice

					m.MockUpdate(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
						return nil
					})

					err := v.Update(nil, updatedObj)
					So(err, ShouldBeNil)

					newObject := &testmodel.List{
						ID: "ID3",
					}
					err = v.Retrieve(nil, newObject)
					So(err, ShouldBeNil)
					So(newObject, ShouldResemble, updatedObj)
				})
			})

			Convey("When I delete the object, it must be deleted in both DBs", func() {

				Convey("When the backend fails, the object should not be deleted", func() {
					newObject := &testmodel.List{
						ID: "ID3",
					}

					m.MockDelete(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
						return manipulate.NewErrCannotCommit("failed")
					})
					err = v.Delete(nil, newObject)
					So(err, ShouldNotBeNil)

					objects := testmodel.ListsList{}
					err := v.RetrieveMany(nil, &objects)
					So(err, ShouldBeNil)
					So(len(objects), ShouldEqual, 3)
					So(objects, ShouldContain, obj3)
				})

				Convey("When the backend succeeds, the object should be deleted", func() {

					newObject := &testmodel.List{
						ID: "ID3",
					}
					m.MockDelete(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
						return nil
					})
					err = v.Delete(nil, newObject)
					So(err, ShouldBeNil)

					objects := testmodel.ListsList{}
					err := v.RetrieveMany(nil, &objects)
					So(err, ShouldBeNil)
					So(len(objects), ShouldEqual, 2)
					So(objects, ShouldNotContain, obj3)
				})
			})
		})
	})
}

func Test_Monitor(t *testing.T) {

	t.Parallel()

	Convey("Given a valid memdb vortex with a subscriber", t, func() {
		m := maniptest.NewTestManipulator()
		s := maniptest.NewTestSubscriber()

		s.MockStart(t, func(ctx context.Context, filter *elemental.PushFilter) {})

		eventChannel := make(chan *elemental.Event)
		s.MockEvents(t, func() chan *elemental.Event { return eventChannel })

		errorsChannel := make(chan error)
		s.MockErrors(t, func() chan error { return errorsChannel })

		statusChannel := make(chan manipulate.SubscriberStatus)
		s.MockStatus(t, func() chan manipulate.SubscriberStatus {
			return statusChannel
		})

		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, m, s, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("When I push a create event, the object must be written in the DB", func() {
			obj := newObject("push1", []string{"test=push"})
			obj.ID = "push"

			event := elemental.NewEvent(elemental.EventCreate, obj)
			eventChannel <- event

			// Necessary sleep to allow event to be processed.
			time.Sleep(100 * time.Millisecond)

			objects := testmodel.ListsList{}
			err := v.RetrieveMany(nil, &objects)
			So(err, ShouldBeNil)
			So(len(objects), ShouldEqual, 1)
			So(objects[0].Name, ShouldResemble, "push1")
			So(objects[0].ID, ShouldResemble, "push")
			So(objects[0].Slice, ShouldResemble, []string{"test=push"})

			Convey("When I push an update event with a new name, the object must be updated", func() {
				obj := newObject("updatedpush", []string{"test=push"})
				obj.ID = "push"

				event := elemental.NewEvent(elemental.EventUpdate, obj)
				eventChannel <- event

				// Necessary sleep to allow event to be processed.
				time.Sleep(100 * time.Millisecond)

				objects := testmodel.ListsList{}
				err := v.RetrieveMany(nil, &objects)
				So(err, ShouldBeNil)
				So(len(objects), ShouldEqual, 1)
				So(objects[0].Name, ShouldResemble, "updatedpush")
				So(objects[0].ID, ShouldResemble, "push")
				So(objects[0].Slice, ShouldResemble, []string{"test=push"})
			})

			Convey("When I push a delete, the object must be deleted", func() {
				obj := newObject("updatedpush", []string{"test=push"})
				obj.ID = "push"

				event := elemental.NewEvent(elemental.EventDelete, obj)
				eventChannel <- event

				// Necessary sleep to allow event to be processed.
				time.Sleep(100 * time.Millisecond)

				objects := testmodel.ListsList{}
				err := v.RetrieveMany(nil, &objects)
				So(err, ShouldBeNil)
				So(len(objects), ShouldEqual, 0)
			})
		})

		Convey("When I push a reconnection status, the db must be resyced", func() {

			obj1 := newObject("obj1", []string{"a=b"})
			obj1.ID = "ID1"
			obj2 := newObject("obj2", []string{"x=y"})
			obj2.ID = "ID2"

			retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
				if mctx.Page() > 1 {
					return nil
				}
				objects := testmodel.ListsList{obj1, obj2}
				*dest.(*testmodel.ListsList) = objects
				return nil
			}
			m.MockRetrieveMany(t, retrieveRespose)

			statusChannel <- manipulate.SubscriberStatusReconnection

			time.Sleep(100 * time.Millisecond)

			objects := testmodel.ListsList{}
			err := v.RetrieveMany(nil, &objects)
			So(err, ShouldBeNil)
			So(len(objects), ShouldEqual, 2)
		})

		Convey("When I push a set of errors, followed by a reconnection, the db must resync", func() {

			obj1 := newObject("obj1", []string{"a=b"})
			obj1.ID = "ID1"
			obj2 := newObject("obj2", []string{"x=y"})
			obj2.ID = "ID2"

			retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
				if mctx.Page() > 1 {
					return nil
				}
				objects := testmodel.ListsList{obj1, obj2}
				*dest.(*testmodel.ListsList) = objects
				return nil
			}
			m.MockRetrieveMany(t, retrieveRespose)

			statusChannel <- manipulate.SubscriberStatusDisconnection
			statusChannel <- manipulate.SubscriberStatusInitialConnection
			errorsChannel <- fmt.Errorf("Some error")
			statusChannel <- manipulate.SubscriberStatusReconnection

			time.Sleep(100 * time.Millisecond)

			objects := testmodel.ListsList{}
			err := v.RetrieveMany(nil, &objects)
			So(err, ShouldBeNil)
			So(len(objects), ShouldEqual, 2)
		})

		Convey("When I push a final disconnection, the monitor should die", func() {

			retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
				t.Errorf("Must not be called")
				return nil
			}
			m.MockRetrieveMany(t, retrieveRespose)

			statusChannel <- manipulate.SubscriberStatusFinalDisconnection

			time.Sleep(100 * time.Millisecond)

			objects := testmodel.ListsList{}
			err := v.RetrieveMany(nil, &objects)
			So(err, ShouldBeNil)
			So(len(objects), ShouldEqual, 0)
		})
	})

}

func Test_WriteBackBackend(t *testing.T) {

	Convey("Given a memdb vortext with a write-back mode", t, func() {
		d, err := newDatastore()
		So(err, ShouldBeNil)

		m := maniptest.NewTestManipulator()
		v := NewMemDBVortex(d, m, nil, newIdentityProcessor(config.WriteBack), testmodel.Manager(), "")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		obj1 := newObject("obj1", []string{"a=b"})
		obj1.ID = "ID1"
		obj2 := newObject("obj2", []string{"x=y"})
		obj2.ID = "ID2"

		retrieveRespose := func(mctx manipulate.Context, dest elemental.Identifiables) error {
			if mctx.Page() > 1 {
				return nil
			}
			objects := testmodel.ListsList{obj1, obj2}
			*dest.(*testmodel.ListsList) = objects
			return nil
		}
		m.MockRetrieveMany(t, retrieveRespose)

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("When I retrieve the objects that are loaded after the sync, I should get the right objects", func() {
			objects := testmodel.ListsList{}
			err := v.RetrieveMany(nil, &objects)
			So(err, ShouldBeNil)
			So(len(objects), ShouldEqual, 2)
			So(objects, ShouldContain, obj1)
			So(objects, ShouldContain, obj2)
		})

		Convey("When I create a new object", func() {

			Convey("When the backend fails for more than the timeout, object must not be created", func() {
				obj3 := newObject("obj3", []string{"w=z"})
				obj3.ID = "ID3"

				m.MockCreate(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
					return manipulate.NewErrCannotCommunicate("testing")
				})

				err := v.Create(nil, obj3)
				Convey("I should not get error during create", func() {
					So(err, ShouldBeNil)
				})

				time.Sleep(1200 * time.Millisecond)
				objects := testmodel.ListsList{}
				err = v.(*MemDBVortex).memory.RetrieveMany(nil, &objects)

				Convey("After waiting for timeout, the object must not be there", func() {
					So(err, ShouldBeNil)
					So(len(objects), ShouldEqual, 2)
					So(len(v.(*MemDBVortex).transactionQueue), ShouldEqual, 0)
				})

			})

			Convey("When the backend event succeeds, the object must be created", func() {
				obj3 := newObject("obj3", []string{"w=z"})
				obj3.ID = "ID3"

				m.MockCreate(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
					return nil
				})

				err := v.Create(nil, obj3)
				Convey("I should no get error during create", func() {
					So(err, ShouldBeNil)
				})

				time.Sleep(1200 * time.Millisecond)
				objects := testmodel.ListsList{}
				err = v.RetrieveMany(nil, &objects)

				Convey("After waiting for timeout, the object must be in the database", func() {
					So(err, ShouldBeNil)
					So(len(objects), ShouldEqual, 3)
					So(len(v.(*MemDBVortex).transactionQueue), ShouldEqual, 0)
				})

				Convey("When I retrieve the object from the local db, it should be there", func() {

					newObject := &testmodel.List{
						ID: "ID3",
					}
					err := v.Retrieve(nil, newObject)
					So(err, ShouldBeNil)
					So(newObject.Name, ShouldResemble, "obj3")
				})
			})
		})

		Convey("When I update an object", func() {

			Convey("When the backend fails, the object must be not be updated", func() {
				obj3 := newObject("obj3", []string{"w=z"})
				obj3.ID = "ID3"

				updatedObject := newObject("obj3updated", []string{"w=z"})
				updatedObject.ID = "ID3"

				m.MockCreate(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
					return nil
				})

				m.MockUpdate(t, func(ctx manipulate.Context, object ...elemental.Identifiable) error {
					return manipulate.NewErrCannotCommunicate("test")
				})

				err := v.Create(nil, obj3)
				Convey("I should not get error during create", func() {
					So(err, ShouldBeNil)
				})

				err = v.Update(nil, updatedObject)
				Convey("I should not get error during update", func() {
					So(err, ShouldBeNil)
				})

				time.Sleep(1200 * time.Millisecond)

				objects := testmodel.ListsList{}
				err = v.RetrieveMany(nil, &objects)

				Convey("The original object must be there and not updated", func() {
					So(err, ShouldBeNil)
					So(len(objects), ShouldEqual, 3)
					So(objects, ShouldContain, obj3)
					So(len(v.(*MemDBVortex).transactionQueue), ShouldEqual, 0)
				})
			})

			Convey("When the backend succeeds, the object must be updated", func() {
				obj3 := newObject("obj3", []string{"w=z"})
				obj3.ID = "ID3"

				updatedObject := newObject("obj3updated", []string{"w=z"})
				updatedObject.ID = "ID3"

				m.MockCreate(t, func(ctx manipulate.Context, objects ...elemental.Identifiable) error {
					return nil
				})

				m.MockUpdate(t, func(ctx manipulate.Context, object ...elemental.Identifiable) error {
					return nil
				})

				err := v.Create(nil, obj3)
				Convey("I should not get error during create", func() {
					So(err, ShouldBeNil)
				})

				err = v.Update(nil, updatedObject)
				Convey("I should not get error during update", func() {
					So(err, ShouldBeNil)
				})

				time.Sleep(1200 * time.Millisecond)

				objects := testmodel.ListsList{}
				err = v.RetrieveMany(nil, &objects)

				Convey("The original object must be there and not updated", func() {
					So(err, ShouldBeNil)
					So(len(objects), ShouldEqual, 3)
					So(objects, ShouldContain, updatedObject)
					So(len(v.(*MemDBVortex).transactionQueue), ShouldEqual, 0)
				})
			})
		})
	})
}

func Test_SubscriberMethods(t *testing.T) {
	Convey("Given a memory DB vortex, with a subsriber and a manipulatr", t, func() {
		m := maniptest.NewTestManipulator()
		s := maniptest.NewTestSubscriber()

		s.MockStart(t, func(ctx context.Context, filter *elemental.PushFilter) {})

		eventChannel := make(chan *elemental.Event)
		s.MockEvents(t, func() chan *elemental.Event { return eventChannel })

		errorsChannel := make(chan error)
		s.MockErrors(t, func() chan error { return errorsChannel })

		statusChannel := make(chan manipulate.SubscriberStatus)
		s.MockStatus(t, func() chan manipulate.SubscriberStatus {
			return statusChannel
		})

		d, err := newDatastore()
		So(err, ShouldBeNil)

		v := NewMemDBVortex(d, m, s, newIdentityProcessor(config.WriteThrough), testmodel.Manager(), "")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err = v.Run(ctx)
		So(err, ShouldBeNil)

		Convey("The subscriber method should return itself", func() {
			ds := v.(*MemDBVortex).Subscriber()
			So(ds, ShouldEqual, v)
		})

		Convey("Events should return the subscriber events channel", func() {
			sv := v.Events()
			So(sv, ShouldNotBeNil)
			Convey("When I push an event to the main channel, it must be propagated", func() {
				obj := newObject("push1", []string{"test=push"})
				obj.ID = "push"

				event := elemental.NewEvent(elemental.EventCreate, obj)
				eventChannel <- event
				time.Sleep(1 * time.Millisecond)

				So(len(sv), ShouldEqual, 1)
				channelEvent := <-sv
				So(channelEvent, ShouldResemble, event)
			})
		})

		Convey("Errors should return the subscriber errors channel", func() {
			se := v.Errors()
			So(se, ShouldNotBeNil)

			Convey("When I push an error to the main channel, it must be propagated", func() {

				err := errors.New("eventerror")
				errorsChannel <- err
				time.Sleep(1 * time.Millisecond)

				So(len(se), ShouldEqual, 1)
				channelError := <-se
				So(channelError, ShouldResemble, err)
			})
		})

		Convey("Status should return the subscriber Status channel", func() {
			se := v.Status()
			So(se, ShouldNotBeNil)

			Convey("When a I push a status update it must be propagated", func() {
				statusChannel <- manipulate.SubscriberStatusDisconnection
				time.Sleep(1 * time.Millisecond)
				So(len(se), ShouldEqual, 1)
				channelStatus := <-se
				So(channelStatus, ShouldEqual, manipulate.SubscriberStatusDisconnection)
			})
		})
	})
}
