package manipcli

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.aporeto.io/elemental"
	testmodel "go.aporeto.io/elemental/test/model"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniphttp"
	"go.aporeto.io/manipulate/maniptest"
)

func Test_ManipulatorMakerFromFlags(t *testing.T) {

	Convey("Given a command with the flags binded", t, func() {
		cmd := &cobra.Command{
			Use:   "test",
			Short: "this is a test",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return nil
			},
		}

		cmd.Flags().AddFlagSet(ManipulatorFlagSet())
		err := viper.BindPFlags(cmd.Flags())
		So(err, ShouldEqual, nil)

		Convey("When I set all flags and call ManipulatorMakerFromFlags", func() {
			cmd.Flags().Set(flagAPI, "https://test.com") // nolint
			cmd.Flags().Set(flagNamespace, "/test")      // nolint
			cmd.Flags().Set(flagEncoding, "msgpack")     // nolint
			cmd.Flags().Set(flagToken, "token1234")      // nolint
			cmd.Flags().Set(flagAPISkipVerify, "true")   // nolint

			m, err := ManipulatorMakerFromFlags()()

			Convey("Then I should get an error", func() {
				So(err, ShouldEqual, nil)
				So(maniphttp.ExtractNamespace(m), ShouldEqual, "/test")
				So(maniphttp.ExtractEndpoint(m), ShouldEqual, "https://test.com")
				So(maniphttp.ExtractEncoding(m), ShouldEqual, "application/msgpack")

				user, pass := maniphttp.ExtractCredentials(m)
				So(user, ShouldEqual, "Bearer")
				So(pass, ShouldEqual, "token1234")

			})
		})

		Convey("When I pass an unsupported encoding", func() {
			cmd.Flags().Set(flagEncoding, "unsupported") // nolint

			_, err := ManipulatorMakerFromFlags()()

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "unsupported encoding")
			})
		})

		Convey("When I pass an invalid capath", func() {
			cmd.Flags().Set(flagEncoding, "msgpack") // nolint
			cmd.Flags().Set(flagCACertPath, "boom")  // nolint

			_, err := ManipulatorMakerFromFlags()()

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "unable to load root ca pool")
			})
		})
	})
}

func Test_readViperFlags(t *testing.T) {

	Convey("Given a command and a task model", t, func() {
		cmd := &cobra.Command{
			Use:   "another command",
			Short: "this is a test",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return nil
			},
		}

		task := testmodel.NewTask()
		err := setViperFlags(cmd, task, testmodel.Manager(), "")
		So(err, ShouldEqual, nil)
		err = viper.BindPFlags(cmd.Flags())
		So(err, ShouldEqual, nil)

		cmd.Flags().Set("name", "test")                             // nolint
		cmd.Flags().Set("description", "a description")             // nolint
		cmd.Flags().Set("status", string(testmodel.TaskStatusDONE)) // nolint

		Convey("When I call readViperFlags with force=false", func() {

			err = readViperFlags(task, testmodel.Manager(), "")

			Convey("Then I should get an error", func() {
				So(err, ShouldEqual, nil)
				So(task.Name, ShouldEqual, "test")
				So(task.Description, ShouldEqual, "a description")
				So(task.Status, ShouldEqual, testmodel.TaskStatusDONE)

			})
		})
	})

	Convey("Given a nil identifiable", t, func() {
		Convey("When I call readViperFlags with force=false", func() {
			err := readViperFlags(nil, testmodel.Manager(), "")
			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "provided identifiable is nil")
			})
		})
	})
}

func Test_retrieveObjectByIDOrByName(t *testing.T) {

	Convey("Given a fake manipulator that works fine", t, func() {

		retrieveManyOutput := testmodel.TasksList{}
		expectedID := "617aec75a829de0001da2032"
		expectedName := "mytask"

		m := maniptest.NewTestManipulator()

		m.MockRetrieve(t, func(mctx manipulate.Context, object elemental.Identifiable) error {
			object.SetIdentifier(expectedID)
			object.(*testmodel.Task).Name = expectedName
			return nil
		})

		m.MockRetrieveMany(t, func(mctx manipulate.Context, dest elemental.Identifiables) error {

			if mctx.Filter().String() == fmt.Sprintf(`name == "%s"`, expectedName) {
				*dest.(*testmodel.TasksList) = retrieveManyOutput
			}
			return nil
		})

		Convey("When I call retrieveObjectByIDOrByName with an ID", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

			defer cancel()
			mctx := manipulate.NewContext(ctx)

			identifiable, err := retrieveObjectByIDOrByName(
				mctx,
				m,
				testmodel.TaskIdentity,
				expectedID,
				testmodel.Manager(),
			)

			Convey("Then I should get an error", func() {
				So(err, ShouldEqual, nil)
				task := identifiable.(*testmodel.Task)
				So(task.ID, ShouldEqual, expectedID)
				So(task.Name, ShouldEqual, expectedName)
			})
		})

		Convey("When I call retrieveObjectByIDOrByName with a valid name", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			mctx := manipulate.NewContext(ctx)

			retrieveManyOutput = testmodel.TasksList{
				{
					ID:   expectedID,
					Name: expectedName,
				},
			}

			identifiable, err := retrieveObjectByIDOrByName(
				mctx,
				m,
				testmodel.TaskIdentity,
				expectedName,
				testmodel.Manager(),
			)

			Convey("Then I should get an error", func() {
				So(err, ShouldEqual, nil)
				task := identifiable.(*testmodel.Task)
				So(task.ID, ShouldEqual, expectedID)
				So(task.Name, ShouldEqual, expectedName)
			})
		})

		Convey("When I call retrieveObjectByIDOrByName with an unknown name", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			mctx := manipulate.NewContext(ctx)

			_, err := retrieveObjectByIDOrByName(
				mctx,
				m,
				testmodel.TaskIdentity,
				"unknown-name",
				testmodel.Manager(),
			)

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "no task found with id or name")
			})
		})
	})

	Convey("Given a fake that returns error", t, func() {

		expectedID := "617aec75a829de0001da2032"
		expectedName := "mytask"

		m := maniptest.NewTestManipulator()
		m.MockRetrieve(t, func(mctx manipulate.Context, object elemental.Identifiable) error {
			return fmt.Errorf("unable to retrieve")
		})

		m.MockRetrieveMany(t, func(mctx manipulate.Context, dest elemental.Identifiables) error {
			return fmt.Errorf("unable to retrieve many")
		})

		Convey("When I call retrieveObjectByIDOrByName with an ID", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			mctx := manipulate.NewContext(ctx)

			_, err := retrieveObjectByIDOrByName(
				mctx,
				m,
				testmodel.TaskIdentity,
				expectedID,
				testmodel.Manager(),
			)

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "unable to retrieve")
			})
		})

		Convey("When I call retrieveObjectByIDOrByName with a name", func() {

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			mctx := manipulate.NewContext(ctx)

			_, err := retrieveObjectByIDOrByName(
				mctx,
				m,
				testmodel.TaskIdentity,
				expectedName,
				testmodel.Manager(),
			)

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "unable to retrieve many")
			})
		})
	})
}

func Test_setViperFlags(t *testing.T) {

	Convey("Given a command", t, func() {

		cmd := &cobra.Command{
			Use:   "test",
			Short: "this is a test",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return nil
			},
		}

		Convey("When I call setViperFlags with an empty identifiable", func() {
			err := setViperFlags(cmd, nil, testmodel.Manager(), "")

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "provided identifiable is nil")
			})
		})

		Convey("When I call setViperFlags with a valid identifiable", func() {
			task := testmodel.NewTask()
			err := setViperFlags(cmd, task, testmodel.Manager(), "")

			Convey("Then I should get an error", func() {
				So(err, ShouldEqual, nil)

				flags := cmd.Flags()

				// primary key
				So(flags.Lookup("ID"), ShouldEqual, nil)

				// regular field
				So(flags.Lookup("description"), ShouldNotEqual, nil)

				// autogenerated
				So(flags.Lookup("parentID"), ShouldEqual, nil)

				// readonly
				So(flags.Lookup("parentType"), ShouldEqual, nil)

				// enum
				So(flags.Lookup("status"), ShouldNotEqual, nil)
			})
		})
	})
}

func Test_parametersToURLValues(t *testing.T) {
	type args struct {
		params []string
	}
	tests := []struct {
		name    string
		args    args
		want    url.Values
		wantErr bool
	}{
		{
			"empty params",
			args{
				params: []string{},
			},
			url.Values{},
			false,
		},
		{
			"single param",
			args{
				params: []string{"a=b"},
			},
			url.Values{
				"a": {"b"},
			},
			false,
		},
		{
			"multiple params",
			args{
				params: []string{"a=b", "b=c"},
			},
			url.Values{
				"a": {"b"},
				"b": {"c"},
			},
			false,
		},
		{
			"invalid params",
			args{
				params: []string{"a", "b"},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parametersToURLValues(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("parametersToURLValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parametersToURLValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateOutputParameters(t *testing.T) {
	type args struct {
		output string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"empty output",
			args{
				output: "",
			},
			true,
		},
		{
			"invalid output",
			args{
				output: "unknown",
			},
			true,
		},
		{
			"valid output",
			args{
				output: flagOutputTable,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateOutputParameters(tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("validateOutputParameters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_shouldManageSpecification(t *testing.T) {
	type args struct {
		spec elemental.AttributeSpecification
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"spec is not exposed",
			args{
				spec: elemental.AttributeSpecification{
					Exposed: false,
				},
			},
			false,
		},
		{
			"spec is Primary Key",
			args{
				spec: elemental.AttributeSpecification{
					Exposed:    true,
					PrimaryKey: true,
				},
			},
			false,
		},
		{
			"spec is Autogenerated",
			args{
				spec: elemental.AttributeSpecification{
					Exposed:       true,
					PrimaryKey:    false,
					Autogenerated: true,
				},
			},
			false,
		},
		{
			"spec is Readonly",
			args{
				spec: elemental.AttributeSpecification{
					Exposed:       true,
					PrimaryKey:    false,
					Autogenerated: false,
					ReadOnly:      true,
				},
			},
			false,
		},
		{
			"valid spec",
			args{
				spec: elemental.AttributeSpecification{
					Exposed:       true,
					PrimaryKey:    false,
					Autogenerated: false,
					ReadOnly:      false,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if resp := shouldManageAttribute(tt.args.spec); resp != tt.want {
				t.Errorf("validateOutputParameters() bool = %v, want %v", resp, tt.want)
			}
		})
	}
}

func Test_nameToFlag(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"empty case",
			args{
				name: "",
			},
			"",
		},
		{
			"simple name",
			args{
				name: "name",
			},
			"name",
		},
		{
			"camelCase name",
			args{
				name: "nameOfThing",
			},
			"name-of-thing",
		},
		{
			"camelCase name with number",
			args{
				name: "name1OfThing",
			},
			"name-1-of-thing",
		},
		{
			"something with a number in between",
			args{
				name: "customerB2B",
			},
			"customer-b2b",
		},
		{
			"something with a number in between advanced",
			args{
				name: "customerB2BAndCo",
			},
			"customer-b2b-and-co",
		},
		{
			"something with a number in the beginning",
			args{
				name: "B2BAndCo",
			},
			"b2b-and-co",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nameToFlag(tt.args.name); got != tt.want {
				t.Errorf("nameToFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_renderTemplate(t *testing.T) {

	Convey("Given a valid template", t, func() {

		content := "hello {{ .name }}"
		values := map[string]string{
			"name": "world",
		}

		Convey("When I call renderTemplate", func() {

			data, err := renderTemplate(content, values)

			Convey("Then I should get no error", func() {
				So(err, ShouldEqual, nil)
				So(string(data), ShouldEqual, "hello world")
			})
		})
	})

	Convey("Given an invalid template", t, func() {

		content := "hello {{ .name "
		values := map[string]string{
			"name": "world",
		}

		Convey("When I call renderTemplate", func() {

			_, err := renderTemplate(content, values)

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
			})
		})
	})

}

func Test_generateFileData(t *testing.T) {

	Convey("Given a nil identifiable", t, func() {

		Convey("When I call generateFileData", func() {

			_, err := generateFileData(nil, "message", true, true, true, outputFormat{
				formatType: formatTypeHash,
				output:     flagOutputYAML,
			})

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "identifiable is nil")
			})
		})
	})

	Convey("Given a valid identifiable", t, func() {

		Convey("When I call generateFileData", func() {

			str, err := generateFileData(testmodel.NewTask(), "message", true, true, true, outputFormat{
				formatType: formatTypeHash,
				output:     flagOutputYAML,
			})

			Convey("Then I should get no error", func() {
				So(err, ShouldEqual, nil)
				So(str, ShouldEqual, `# message

description: ""
name: ""
status: TODO

# Here is a copy of the full original object you are editing:
#
# ID: ""
# description: ""
# name: ""
# parentID: ""
# parentType: ""`)
			})
		})
	})

}

func Test_splitParentInfo(t *testing.T) {
	type args struct {
		parent string
	}
	tests := []struct {
		name     string
		args     args
		wantName string
		wantID   string
		wantErr  bool
	}{
		{
			"valid format",
			args{
				parent: "resource/12345",
			},
			"resource",
			"12345",
			false,
		},
		{
			"empty parent",
			args{
				parent: "",
			},
			"",
			"",
			true,
		},
		{
			"no / separator",
			args{
				parent: "just-a-name",
			},
			"",
			"",
			true,
		},
		{
			"multiple / separator",
			args{
				parent: "just/a/name",
			},
			"",
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := splitParentInfo(tt.args.parent)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitParentInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantName {
				t.Errorf("splitParentInfo() got = %v, want %v", got, tt.wantName)
			}
			if got1 != tt.wantID {
				t.Errorf("splitParentInfo() got1 = %v, want %v", got1, tt.wantID)
			}
		})
	}
}

func TestReadData(t *testing.T) {

	expectedData := `name: chris`

	// Temporary template
	templateFile, err := os.CreateTemp("", "test-data-template-")
	if err != nil {
		panic(err.Error())
	}
	defer os.RemoveAll(templateFile.Name()) // nolint

	_, err = templateFile.Write([]byte("hello {{ .Values.name }}"))
	if err != nil {
		panic(err.Error())
	}

	// Temporary value file
	valuesFile, err := os.CreateTemp("", "test-data-values-")
	if err != nil {
		panic(err.Error())
	}
	defer os.RemoveAll(valuesFile.Name()) // nolint

	_, err = valuesFile.Write([]byte("name: jeanmichel"))
	if err != nil {
		panic(err.Error())
	}

	// Fake http server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expectedData)) // nolint: errcheck
	}))
	defer ts.Close()

	type args struct {
		apiurl     string
		namespace  string
		file       string
		url        string
		valuesFile string
		values     []string
		printOnly  bool
		mandatory  bool
	}
	tests := []struct {
		name     string
		args     args
		wantData []byte
		wantErr  bool
	}{
		{
			"read from url and return data",
			args{
				apiurl:     "https://api.test.com",
				namespace:  "/test",
				file:       "",
				url:        ts.URL,
				valuesFile: "",
				values:     []string{},
				printOnly:  false,
				mandatory:  false,
			},
			[]byte(expectedData),
			false,
		},
		{
			"read from url and return print-only data",
			args{
				apiurl:     "https://api.test.com",
				namespace:  "/test",
				file:       "",
				url:        ts.URL,
				valuesFile: "",
				values:     []string{},
				printOnly:  true,
				mandatory:  false,
			},
			[]byte(expectedData),
			false,
		},
		{
			"read from file without values",
			args{
				apiurl:     "https://api.test.com",
				namespace:  "/test",
				file:       templateFile.Name(),
				url:        "",
				valuesFile: "",
				values:     []string{},
				printOnly:  false,
				mandatory:  false,
			},
			[]byte("hello <no value>"),
			false,
		},
		{
			"read from file without mandatory file",
			args{
				apiurl:     "https://api.test.com",
				namespace:  "/test",
				file:       "",
				url:        "",
				valuesFile: "",
				values:     []string{},
				printOnly:  false,
				mandatory:  true,
			},
			nil,
			true,
		},
		{
			"read from file with value files",
			args{
				apiurl:     "https://api.test.com",
				namespace:  "/test",
				file:       templateFile.Name(),
				url:        "",
				valuesFile: valuesFile.Name(),
				values:     nil,
				printOnly:  false,
				mandatory:  false,
			},
			[]byte("hello jeanmichel"),
			false,
		},
		{
			"read from file with values override",
			args{
				apiurl:     "https://api.test.com",
				namespace:  "/test",
				file:       templateFile.Name(),
				url:        "",
				valuesFile: valuesFile.Name(),
				values:     []string{"name=chris"},
				printOnly:  false,
				mandatory:  false,
			},
			[]byte("hello chris"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := ReadData(tt.args.apiurl, tt.args.namespace, tt.args.file, tt.args.url, tt.args.valuesFile, tt.args.values, tt.args.printOnly, tt.args.mandatory)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("ReadData() = %v, want %v", string(gotData), string(tt.wantData))
			}
		})
	}
}
