package manipbolt

import (
	"context"
	"os"
	"testing"

	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/json"
	"github.com/stretchr/testify/require"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/manipbolt/customcodecs"
)

func prepareDB(t *testing.T) (manipulate.TransactionalManipulator, []*testmodel.List) {

	f, err := os.CreateTemp("", "")
	require.Nil(t, err)
	require.FileExists(t, f.Name())

	c, err := customcodecs.NewRandomJSONTagGenerator()
	require.Nil(t, err)
	require.NotNil(t, c)

	m, err := New(f.Name(), testmodel.Manager(), OptionCodec(c))
	require.Nil(t, err)
	require.NotNil(t, m)

	return m, create(t, m)
}

func create(t *testing.T, m manipulate.Manipulator) []*testmodel.List {

	p1 := &testmodel.List{
		Name:     "Centos7",
		ParentID: "xyz",
		Slice:    []string{"$name=centos7", "category=centos", "a=b", "c=d"},
	}

	p2 := &testmodel.List{
		Name:  "Centos8",
		Slice: []string{"$name=centos8", "category=centos", "x=y", "w=z"},
	}

	p3 := &testmodel.List{
		Name:  "Rhel7",
		Slice: []string{"$name=rhel7", "category=rhel", "a=b", "x=y"},
	}

	p4 := &testmodel.List{
		Name:  "rhel8",
		Slice: []string{"$name=rhel8", "category=rhel", "a=b", "g=h"},
	}

	err := m.Create(nil, p1)
	require.Nil(t, err)
	err = m.Create(nil, p2)
	require.Nil(t, err)
	err = m.Create(nil, p3)
	require.Nil(t, err)
	err = m.Create(nil, p4)
	require.Nil(t, err)

	return []*testmodel.List{p1, p2, p3, p4}
}

func TestBoltManip_New(t *testing.T) {

	f, err := os.CreateTemp("", "")
	require.Nil(t, err)
	require.FileExists(t, f.Name())

	defer os.RemoveAll(f.Name()) // nolint: errcheck

	m, err := New(f.Name(), testmodel.Manager())
	require.Nil(t, err)
	require.NotNil(t, m)

	m, err = New("", testmodel.Manager())
	require.NotNil(t, err)
	require.Nil(t, m)
}

func TestBoltManip_Create(t *testing.T) {

	f, err := os.CreateTemp("", "")
	require.Nil(t, err)
	require.FileExists(t, f.Name())

	defer os.RemoveAll(f.Name()) // nolint: errcheck

	m, err := New(f.Name(), testmodel.Manager())
	require.Nil(t, err)
	require.NotNil(t, m)

	p := &testmodel.List{
		Name:  "ubuntu",
		Slice: []string{"$names=ubuntu16"},
	}

	err = m.Create(nil, p)
	require.Nil(t, err)
	require.NotEmpty(t, p.ID)

	p1 := &testmodel.List{
		ID:   p.ID,
		Name: "not good",
	}

	err = m.Retrieve(nil, p1)
	require.Nil(t, err)
	require.Equal(t, p, p1)
}

func TestBoltManip_ClosedDB(t *testing.T) {

	f, err := os.CreateTemp("", "")
	require.Nil(t, err)
	require.FileExists(t, f.Name())

	defer os.RemoveAll(f.Name()) // nolint: errcheck

	m, err := New(f.Name(), testmodel.Manager())
	require.Nil(t, err)
	require.NotNil(t, m)

	ctxn, err := m.(*boltManipulator).getDB().Begin(true)
	require.Nil(t, err)
	require.NotNil(t, ctxn)

	err = ctxn.Rollback()
	require.Nil(t, err)

	atxn, err := m.(*boltManipulator).getDB().Begin(false)
	require.Nil(t, err)
	require.NotNil(t, atxn)

	err = atxn.Rollback()
	require.Nil(t, err)

	err = m.(*boltManipulator).getDB().Close()
	require.Nil(t, err)

	tid := manipulate.NewTransactionID()
	m.(*boltManipulator).registerTxn(tid, ctxn)

	err = m.Commit(tid)
	require.NotNil(t, err)

	txn := m.(*boltManipulator).registeredTxnWithID(tid)
	require.NotNil(t, txn)

	tid = manipulate.NewTransactionID()
	m.(*boltManipulator).registerTxn(tid, atxn)

	ok := m.Abort(tid)
	require.False(t, ok)

	txn = m.(*boltManipulator).registeredTxnWithID(tid)
	require.NotNil(t, txn)

	p := &testmodel.List{
		Name:  "ubuntu",
		Slice: []string{"$names=ubuntu16"},
	}

	err = m.Create(nil, p)
	require.NotNil(t, err)

	p = &testmodel.List{
		Name:  "ubuntu",
		Slice: []string{"$names=ubuntu16"},
	}

	err = m.Update(nil, p)
	require.NotNil(t, err)

	p1 := &testmodel.List{
		ID:   p.ID,
		Name: "not good",
	}

	err = m.Retrieve(nil, p1)
	require.NotNil(t, err)

	err = m.Delete(nil, p1)
	require.NotNil(t, err)

	err = m.DeleteMany(nil, testmodel.ListIdentity)
	require.NotNil(t, err)

	lists := &testmodel.ListsList{}

	err = m.RetrieveMany(nil, lists)
	require.NotNil(t, err)

	_, err = m.Count(nil, testmodel.ListIdentity)
	require.NotNil(t, err)
}

func TestBoltManip_CreateWithCodec(t *testing.T) {

	f, err := os.CreateTemp("", "")
	require.Nil(t, err)
	require.FileExists(t, f.Name())

	defer os.RemoveAll(f.Name()) // nolint: errcheck

	codec := customcodecs.NewFileSizeValidator(f.Name(), 64*1024, json.Codec)

	m, err := New(f.Name(), testmodel.Manager(), OptionCodec(codec))
	require.Nil(t, err)
	require.NotNil(t, m)

	p := &testmodel.List{
		Name:  "ubuntu",
		Slice: []string{"$names=ubuntu16"},
	}

	err = m.Create(nil, p)
	require.Nil(t, err)
	require.NotEmpty(t, p.ID)

	p1 := &testmodel.List{
		ID:   p.ID,
		Name: "not good",
	}

	err = m.Retrieve(nil, p1)
	require.Nil(t, err)
	require.Equal(t, p, p1)

	p2 := &testmodel.List{
		Name:  "ubuntu",
		Slice: []string{"$names=ubuntu18"},
	}

	err = m.Create(nil, p2)
	require.NotNil(t, err)
	require.IsType(t, manipulate.ErrCannotExecuteQuery{}, err)
	require.Contains(t, err.Error(), customcodecs.ErrExceedsSize.Error())

	info, err := os.Stat(f.Name())
	require.Nil(t, err)
	require.LessOrEqual(t, info.Size(), int64(64*1024))
}

func TestBoltManip_Count(t *testing.T) {

	m, _ := prepareDB(t)
	defer os.RemoveAll(m.(*boltManipulator).getDB().Bolt.Path()) // nolint: errcheck

	n, err := m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	n, err = m.Count(nil, testmodel.UserIdentity)
	require.Nil(t, err)
	require.Equal(t, 0, n)

	mctx := manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(
			elemental.NewFilterComposer().WithKey("slice").Contains("category=rhel", "a=b").Done(),
		),
	)

	n, err = m.Count(mctx, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 3, n)

	filter := elemental.NewFilterComposer().And(
		elemental.NewFilterComposer().
			WithKey("slice").Contains("category=rhel").Done(),
		elemental.NewFilterComposer().
			WithKey("slice").Contains("a=b").Done(),
		elemental.NewFilterComposer().
			WithKey("slice").Contains("g=h").Done(),
	).Done()

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(filter),
	)

	n, err = m.Count(mctx, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 1, n)

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(
			elemental.NewFilterComposer().WithKey("Name").Exists().Done(),
		),
	)

	_, err = m.Count(mctx, testmodel.ListIdentity)
	require.NotNil(t, err)
}

func TestBoltManip_Update(t *testing.T) {

	m, pus := prepareDB(t)
	defer os.RemoveAll(m.(*boltManipulator).getDB().Bolt.Path()) // nolint: errcheck

	n, err := m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	l5 := &testmodel.List{
		ID:    pus[2].ID,
		Name:  "SIBI",
		Slice: nil,
	}

	err = m.Update(nil, l5)
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	p1 := &testmodel.List{
		ID:   l5.ID,
		Name: "not good",
	}

	err = m.Retrieve(nil, p1)
	require.Nil(t, err)
	require.Equal(t, p1, l5)

	pu := &testmodel.List{}

	err = m.Update(nil, pu)
	require.IsType(t, manipulate.ErrCannotExecuteQuery{}, err)
	require.Contains(t, err.Error(), storm.ErrNotFound.Error())

	ns := &testmodel.User{
		ID: "abc",
	}

	err = m.Update(nil, ns)
	require.IsType(t, manipulate.ErrCannotExecuteQuery{}, err)
	require.Contains(t, err.Error(), storm.ErrNotFound.Error())
}

func TestBoltManip_DeleteMany(t *testing.T) {

	m, _ := prepareDB(t)
	defer os.RemoveAll(m.(*boltManipulator).getDB().Bolt.Path()) // nolint: errcheck

	n, err := m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	mctx := manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(
			elemental.NewFilterComposer().WithKey("Name").Matches("^Cent").Done(),
		),
	)

	err = m.DeleteMany(mctx, testmodel.ListIdentity)
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 2, n)

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(
			elemental.NewFilterComposer().Done(),
		),
	)

	err = m.DeleteMany(mctx, testmodel.ListIdentity)
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 2, n)

	err = m.DeleteMany(nil, testmodel.ListIdentity)
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 0, n)

	err = m.DeleteMany(nil, testmodel.UserIdentity)
	require.Nil(t, err)

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(
			elemental.NewFilterComposer().WithKey("Name").Exists().Done(),
		),
	)

	err = m.DeleteMany(mctx, testmodel.ListIdentity)
	require.NotNil(t, err)
}

func TestBoltManip_Retrieve(t *testing.T) {

	m, pus := prepareDB(t)
	defer os.RemoveAll(m.(*boltManipulator).getDB().Bolt.Path()) // nolint: errcheck

	p := &testmodel.List{
		ID: pus[0].ID,
	}

	err := m.Retrieve(nil, p)
	require.Nil(t, err)
	require.Equal(t, pus[0], p)

	p1 := &testmodel.List{
		ID: "BADID",
	}

	err = m.Retrieve(nil, p1)
	require.IsType(t, manipulate.ErrCannotExecuteQuery{}, err)
	require.Contains(t, err.Error(), storm.ErrNotFound.Error())

	p2 := &testmodel.List{}

	err = m.Retrieve(nil, p2)
	require.IsType(t, manipulate.ErrCannotExecuteQuery{}, err)
	require.Contains(t, err.Error(), storm.ErrNotFound.Error())

	ns := &testmodel.User{}

	err = m.Retrieve(nil, ns)
	require.IsType(t, manipulate.ErrCannotExecuteQuery{}, err)
	require.Contains(t, err.Error(), storm.ErrNotFound.Error())
}

func TestBoltManip_Delete(t *testing.T) {

	m, pus := prepareDB(t)
	defer os.RemoveAll(m.(*boltManipulator).getDB().Bolt.Path()) // nolint: errcheck

	n, err := m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	p1 := &testmodel.List{
		ID:   pus[0].ID,
		Name: "not good",
	}

	err = m.Retrieve(nil, p1)
	require.Nil(t, err)
	require.Equal(t, pus[0], p1)

	err = m.Delete(nil, p1)
	require.Nil(t, err)

	err = m.Retrieve(nil, p1)
	require.IsType(t, manipulate.ErrCannotExecuteQuery{}, err)
	require.Contains(t, err.Error(), storm.ErrNotFound.Error())

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 3, n)

	p3 := &testmodel.List{
		ID:   pus[1].ID,
		Name: "not good",
	}

	err = m.Retrieve(nil, p3)
	require.Nil(t, err)
	require.Equal(t, pus[1], p3)

	ns := &testmodel.User{
		ID: "abc",
	}

	err = m.Delete(nil, ns)
	require.IsType(t, manipulate.ErrCannotExecuteQuery{}, err)
	require.Contains(t, err.Error(), storm.ErrNotFound.Error())

	pu := &testmodel.List{}

	err = m.Delete(nil, pu)
	require.IsType(t, manipulate.ErrCannotExecuteQuery{}, err)
	require.Contains(t, err.Error(), storm.ErrNotFound.Error())
}

func TestBoltManip_RetrieveMany(t *testing.T) {

	m, apus := prepareDB(t)
	defer os.RemoveAll(m.(*boltManipulator).getDB().Bolt.Path()) // nolint: errcheck

	err := m.RetrieveMany(nil, nil)
	require.NotNil(t, err)

	pus := testmodel.ListsList{}

	err = m.RetrieveMany(nil, &pus)
	require.Nil(t, err)
	require.Len(t, pus, 4)

	mctx := manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(
			elemental.NewFilterComposer().WithKey("Name").Equals("Centos7").Done(),
		),
	)

	pus = testmodel.ListsList{}

	err = m.RetrieveMany(mctx, &pus)
	require.Nil(t, err)
	require.Len(t, pus, 1)
	require.Equal(t, apus[0], pus[0])

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(
			elemental.NewFilterComposer().WithKey("parentID").Equals("xyz").Done(),
		),
	)

	pus = testmodel.ListsList{}

	err = m.RetrieveMany(mctx, &pus)
	require.Nil(t, err)
	require.Len(t, pus, 1)
	require.Equal(t, apus[0], pus[0])

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(
			elemental.NewFilterComposer().Done(),
		),
	)

	pus = testmodel.ListsList{}

	err = m.RetrieveMany(mctx, &pus)
	require.Nil(t, err)
	require.Len(t, pus, 0)

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(
			elemental.NewFilterComposer().WithKey("Name").Matches("^Cen").Done(),
		),
	)

	pus = testmodel.ListsList{}

	err = m.RetrieveMany(mctx, &pus)
	require.Nil(t, err)
	require.Len(t, pus, 2)
	require.Contains(t, pus, apus[0])
	require.Contains(t, pus, apus[1])

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(
			elemental.NewFilterComposer().WithKey("slice").Contains("category=rhel", "a=b", "none").Done(),
		),
	)

	pus = testmodel.ListsList{}

	err = m.RetrieveMany(mctx, &pus)
	require.Nil(t, err)
	require.Len(t, pus, 3)
	require.Contains(t, pus, apus[0])
	require.Contains(t, pus, apus[2])
	require.Contains(t, pus, apus[3])

	filter := elemental.NewFilterComposer().And(
		elemental.NewFilterComposer().
			WithKey("slice").Equals("category=rhel").Done(),
		elemental.NewFilterComposer().
			WithKey("slice").Equals("a=b").Done(),
		elemental.NewFilterComposer().
			WithKey("slice").Equals("g=h").Done(),
	).Done()

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(filter),
	)

	pus = testmodel.ListsList{}

	err = m.RetrieveMany(mctx, &pus)
	require.Nil(t, err)
	require.Len(t, pus, 1)
	require.Contains(t, pus, apus[3])

	filter = elemental.NewFilterComposer().Or(
		elemental.NewFilterComposer().
			WithKey("slice").Contains("category=rhel").Done(),
		elemental.NewFilterComposer().
			WithKey("slice").Contains("a=b").Done(),
		elemental.NewFilterComposer().
			WithKey("slice").Contains("g=h").Done(),
	).Done()

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(filter),
	)

	pus = testmodel.ListsList{}

	err = m.RetrieveMany(mctx, &pus)
	require.Nil(t, err)
	require.Len(t, pus, 3)
	require.Contains(t, pus, apus[0])
	require.Contains(t, pus, apus[2])
	require.Contains(t, pus, apus[3])

	pus = testmodel.ListsList{}

	err = m.RetrieveMany(nil, &pus)
	require.Nil(t, err)
	require.Len(t, pus, 4)
	require.Contains(t, pus, apus[0])
	require.Contains(t, pus, apus[1])
	require.Contains(t, pus, apus[2])
	require.Contains(t, pus, apus[3])

	nss := testmodel.ListsList{}

	err = m.RetrieveMany(nil, &nss)
	require.Nil(t, err)

	filter = elemental.NewFilterComposer().Or(
		elemental.NewFilterComposer().
			WithKey("out").Contains("category=rhel").Done(),
		elemental.NewFilterComposer().
			WithKey("slice").Contains("a=b").Done(),
		elemental.NewFilterComposer().
			WithKey("slice").Contains("g=h").Done(),
	).Done()

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionFilter(filter),
	)

	pus = testmodel.ListsList{}

	err = m.RetrieveMany(mctx, &pus)
	require.Nil(t, err)
	require.Len(t, pus, 0)
}

func TestBoltManip_Flush(t *testing.T) {

	m, _ := prepareDB(t)
	defer os.RemoveAll(m.(*boltManipulator).getDB().Bolt.Path()) // nolint: errcheck

	n, err := m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	err = m.(manipulate.FlushableManipulator).Flush(context.Background())
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 0, n)

	err = m.(manipulate.FlushableManipulator).Flush(context.Background())
	require.Nil(t, err)

	_ = create(t, m)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)
}

func TestBoltManip_Commit(t *testing.T) {

	m, _ := prepareDB(t)
	defer os.RemoveAll(m.(*boltManipulator).getDB().Bolt.Path()) // nolint: errcheck

	n, err := m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	p := &testmodel.List{
		Name:  "SIBI",
		Slice: []string{"a=c", "app=centos"},
	}

	tid := manipulate.NewTransactionID()
	mctx := manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionTransactionID(tid),
	)

	err = m.Create(mctx, p)
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	p.Name = "CENTOS"

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionTransactionID(tid),
	)

	err = m.Create(mctx, p)
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	err = m.Commit(tid)
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 5, n)

	p1 := &testmodel.List{
		ID: p.ID,
	}

	err = m.Retrieve(nil, p1)
	require.Nil(t, err)
	require.Equal(t, p, p1)

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionTransactionID(tid),
	)

	p1.Name = "nginx"

	err = m.Update(mctx, p1)
	require.Nil(t, err)

	p2 := &testmodel.List{
		ID: p.ID,
	}

	err = m.Retrieve(nil, p2)
	require.Nil(t, err)
	require.Equal(t, "CENTOS", p2.Name)

	err = m.Commit(tid)
	require.Nil(t, err)

	p3 := &testmodel.List{
		ID: p.ID,
	}

	err = m.Retrieve(nil, p3)
	require.Nil(t, err)
	require.Equal(t, "nginx", p3.Name)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 5, n)

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionTransactionID(tid),
	)

	err = m.Delete(mctx, p3)
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 5, n)

	err = m.Commit(tid)
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	mctx = manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionTransactionID(tid),
	)

	err = m.DeleteMany(mctx, testmodel.ListIdentity)
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	err = m.Commit(tid)
	require.Nil(t, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 0, n)
}

func TestBoltManip_Abort(t *testing.T) {

	m, _ := prepareDB(t)
	defer os.RemoveAll(m.(*boltManipulator).getDB().Bolt.Path()) // nolint: errcheck

	tid := manipulate.NewTransactionID()

	ok := m.Abort(tid)
	require.False(t, ok)

	n, err := m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	p := &testmodel.List{
		Name:  "SIBI",
		Slice: []string{"a=c", "app=centos"},
	}

	mctx := manipulate.NewContext(
		context.Background(),
		manipulate.ContextOptionTransactionID(tid),
	)

	err = m.Create(mctx, p)
	require.Nil(t, err)

	txn := m.(*boltManipulator).registeredTxnWithID(tid)
	require.NotNil(t, txn)

	ok = m.Abort(tid)
	require.True(t, ok)

	txn = m.(*boltManipulator).registeredTxnWithID(tid)
	require.Nil(t, txn)

	err = m.Commit(tid)
	require.IsType(t, manipulate.ErrTransactionNotFound{}, err)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)
}

func TestBoltManip_NewOnExistingDB(t *testing.T) {

	m, _ := prepareDB(t)
	path := m.(*boltManipulator).getDB().Bolt.Path()

	defer os.RemoveAll(path) // nolint: errcheck

	n, err := m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)

	// Only one instance per db can be created. We close the
	// existing DB, since we already have an instance open.
	err = m.(*boltManipulator).getDB().Close()
	require.Nil(t, err)

	c, err := customcodecs.NewRandomJSONTagGenerator()
	require.Nil(t, err)
	require.NotNil(t, c)

	m, err = New(path, testmodel.Manager(), OptionCodec(c))
	require.Nil(t, err)
	require.NotNil(t, m)

	n, err = m.Count(nil, testmodel.ListIdentity)
	require.Nil(t, err)
	require.Equal(t, 4, n)
}
