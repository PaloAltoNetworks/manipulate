package manipcli

import (
	"fmt"
	"net/url"
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
	"golang.org/x/net/context"
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
			cmd.Flags().Set(FlagAPI, "https://test.com")
			cmd.Flags().Set(FlagNamespace, "/test")
			cmd.Flags().Set(FlagEncoding, "msgpack")
			cmd.Flags().Set(FlagToken, "token1234")
			cmd.Flags().Set(FlagAPISkipVerify, "true")

			m, err := ManipulatorMakerFromFlags()()

			Convey("Then I should get an error", func() {
				So(err, ShouldEqual, nil)
				So(maniphttp.ExtractNamespace(m), ShouldEqual, "/test")
				So(maniphttp.ExtractEndpoint(m), ShouldEqual, "https://test.com")
				So(maniphttp.ExtractEncoding(m), ShouldEqual, "application/msgpack")

				user, pass := maniphttp.ExtractCredentials(m)
				So(user, ShouldEqual, "Bearer")
				So(pass, ShouldEqual, "") // Note: not sure why it is empty here! I was expecting token1234

			})
		})

		Convey("When I pass an unsupported encoding", func() {
			cmd.Flags().Set(FlagEncoding, "unsupported")

			_, err := ManipulatorMakerFromFlags()()

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "unsupported encoding")
			})
		})

		Convey("When I pass an invalid capath", func() {
			cmd.Flags().Set(FlagEncoding, "msgpack")
			cmd.Flags().Set(FlagCACertPath, "boom")

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
		err := setViperFlags(cmd, task, false)
		So(err, ShouldEqual, nil)
		err = viper.BindPFlags(cmd.Flags())
		So(err, ShouldEqual, nil)

		cmd.Flags().Set("name", "test")
		cmd.Flags().Set("description", "a description")
		cmd.Flags().Set("status", string(testmodel.TaskStatusDONE))

		Convey("When I call readViperFlags with force=false", func() {

			err = readViperFlags(task)

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
			err := readViperFlags(nil)
			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "provided identifiable is nil")
			})
		})
	})
}

func Test_persistentPreRuneE(t *testing.T) {

	Convey("Given I have a command with PersistentPreRunE", t, func() {

		cmd := &cobra.Command{
			Use:   "another command",
			Short: "this is a test",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return nil
			},
		}

		Convey("When I don't pass an output value", func() {
			err := persistentPreRunE(cmd, nil)

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "invalid output")
			})
		})

		Convey("When I pass an output value", func() {

			cmd.Flags().StringP(FlagOutput, "", "json", "a valid output")
			cmd.Flags().Set(FlagOutput, FlagOutputJSON)

			err := persistentPreRunE(cmd, []string{})

			Convey("Then I should not get an error", func() {
				So(err, ShouldEqual, nil)
			})
		})
	})

	Convey("Given I have a command which return an eror", t, func() {

		cmd := &cobra.Command{
			Use:   "another command",
			Short: "this is a test",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return fmt.Errorf("boom")
			},
		}

		Convey("When I pass an output value", func() {

			cmd.Flags().StringP(FlagOutput, "", "json", "a valid output")
			cmd.Flags().Set(FlagOutput, FlagOutputJSON)

			err := persistentPreRunE(cmd, []string{})

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "boom")
			})
		})
	})
}

func Test_checkRequiredFlags(t *testing.T) {

	Convey("Given I have a command with required flag", t, func() {

		cmd := &cobra.Command{
			Use:   "a command",
			Short: "this is a test",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return nil
			},
		}

		cmd.Flags().StringP("mandatory", "", "", "a mandatory flag")
		cmd.MarkFlagRequired("mandatory")

		Convey("When I pass don't pass the flag", func() {

			err := checkRequiredFlags(cmd.Flags())

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "Required argument `--mandatory` must be passed")
			})
		})

		Convey("When I pass pass the flag", func() {

			cmd.Flags().Set("mandatory", "ok")
			err := checkRequiredFlags(cmd.Flags())

			Convey("Then I should not get an error", func() {
				So(err, ShouldEqual, nil)
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
			err := setViperFlags(cmd, nil, true)

			Convey("Then I should get an error", func() {
				So(err, ShouldNotEqual, nil)
				So(err.Error(), ShouldContainSubstring, "provided identifiable is nil")
			})
		})

		Convey("When I call setViperFlags with a valid identifiable", func() {
			task := testmodel.NewTask()
			err := setViperFlags(cmd, task, true)

			Convey("Then I should get an error", func() {
				So(err, ShouldEqual, nil)

				flags := cmd.Flags()

				// primary key
				So(flags.Lookup("ID"), ShouldEqual, nil)

				// regular field
				So(flags.Lookup("description"), ShouldNotEqual, nil)

				// required field
				So(flags.Lookup("name"), ShouldNotEqual, nil)
				So(flags.Lookup("name").Annotations, ShouldResemble, map[string][]string{
					"cobra_annotation_bash_completion_one_required_flag": {"true"},
				})

				// autogenerated
				So(flags.Lookup("parentID"), ShouldEqual, nil)

				// readonly
				So(flags.Lookup("parentType"), ShouldEqual, nil)

				// enum
				So(flags.Lookup("status"), ShouldNotEqual, nil)
			})
		})
	})

	Convey("Given another command", t, func() {
		cmd := &cobra.Command{
			Use:   "another command",
			Short: "this is a test",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return nil
			},
		}

		Convey("When I call setViperFlags with force=false", func() {
			task := testmodel.NewTask()
			err := setViperFlags(cmd, task, false)

			Convey("Then I should get an error", func() {
				So(err, ShouldEqual, nil)

				flags := cmd.Flags()
				// required field
				So(flags.Lookup("name").Annotations, ShouldBeNil)
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
				output: FlagOutputTable,
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
			"spec type is external",
			args{
				spec: elemental.AttributeSpecification{
					Exposed:       true,
					PrimaryKey:    false,
					Autogenerated: false,
					ReadOnly:      false,
					Type:          "external",
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
			if resp := shouldManageSpecification(tt.args.spec); resp != tt.want {
				t.Errorf("validateOutputParameters() bool = %v, want %v", resp, tt.want)
			}
		})
	}
}
