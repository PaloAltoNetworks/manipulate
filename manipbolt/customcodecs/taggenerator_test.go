package customcodecs

import (
	"testing"

	"github.com/stretchr/testify/require"
	testmodel "go.aporeto.io/elemental/test/model"
)

func Test_TagGenerator(t *testing.T) {

	tr, err := NewRandomJSONTagGenerator()
	require.Nil(t, err)

	l := &testmodel.List{
		Name:      "Centos",
		Unexposed: "abc",
	}

	d, err := tr.Marshal(l)
	require.Nil(t, err)
	require.NotNil(t, d)

	l1 := &testmodel.List{}
	err = tr.Unmarshal(d, l1)
	require.Nil(t, err)
	require.Equal(t, l, l1)
	require.Equal(t, "abc", l1.Unexposed)

	i := 45

	d, err = tr.Marshal(&i)
	require.Nil(t, err)
	require.NotNil(t, d)

	var i1 int
	err = tr.Unmarshal(d, &i1)
	require.Nil(t, err)
	require.Equal(t, i, i1)

	name := tr.Name()
	require.Equal(t, "randomJSONTagGenerator", name)

	type userJSON struct {
		Username string          `json:"uname"`
		Password string          `json:"-"`
		Age      int             `json:"age"`
		Company  string          `json:"company"`
		List     *testmodel.List `json:"list"`
		User     *userJSON
	}

	nt := &userJSON{}

	nt.Username = "ABC"
	nt.Password = "XYZ123"
	nt.Age = 23
	nt.Company = "RETO"
	nt.List = l1

	d, err = tr.Marshal(nt)
	require.Nil(t, err)
	require.NotNil(t, d)

	ot := userJSON{}

	err = tr.Unmarshal(d, ot)
	require.NotNil(t, err)

	at := &userJSON{}

	err = tr.Unmarshal(d, at)
	require.Nil(t, err)
	require.Equal(t, nt.List, at.List)
}
