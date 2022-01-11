package manipcli

import (
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	testmodel "go.aporeto.io/elemental/test/model"
)

func assertIdentityCommand(cmd *cobra.Command, expectedSubCommands []string) {

	subCmd := cmd.Commands()
	So(len(subCmd), ShouldEqual, len(expectedSubCommands))
	sort.Slice(subCmd, func(i, j int) bool {
		return subCmd[i].Name() < subCmd[j].Name()
	})
	So(subCmd[0].Name(), ShouldEqual, expectedSubCommands[0])
	So(subCmd[1].Name(), ShouldEqual, expectedSubCommands[1])
}

func assertCommandAndSetFlags(cmd *cobra.Command) {

	So(cmd, ShouldNotEqual, nil)
	cmd.Flags().AddFlagSet(ManipulatorFlagSet())
	cmd.Flags().StringP(flagOutput, "o", "default", "Format to used print output. Options are 'table', 'json', 'yaml', 'none', 'template' or 'default'.")
	err := viper.BindPFlags(cmd.Flags())
	So(err, ShouldEqual, nil)
}

func Test_OptionIgnoreIdentities(t *testing.T) {

	Convey("Given a configuration", t, func() {
		cfg := &cliConfig{}

		Convey("When I call OptionIgnoreIdentities", func() {

			OptionIgnoreIdentities(testmodel.TaskIdentity)(cfg)

			Convey("Then I should have the configuration set", func() {
				So(cfg.ignoredIdentities, ShouldResemble, map[string]struct{}{
					testmodel.TaskIdentity.Name: {},
				})
			})
		})
	})
}

func Test_New(t *testing.T) {

	Convey("Given a valid model manager and a manipulator maker", t, func() {

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

		cmd.Flags().Set(flagAPI, "https://test.com") // nolint
		cmd.Flags().Set(flagNamespace, "/test")      // nolint
		cmd.Flags().Set(flagEncoding, "msgpack")     // nolint
		cmd.Flags().Set(flagToken, "token1234")      // nolint
		cmd.Flags().Set(flagAPISkipVerify, "true")   // nolint

		mmanager := testmodel.Manager()
		mmaker := ManipulatorMakerFromFlags()

		Convey("When I call New", func() {

			genCmd := New(mmanager, mmaker, OptionIgnoreIdentities(testmodel.UserIdentity))

			Convey("Then I should get a generated command", func() {
				So(genCmd, ShouldNotEqual, nil)
				So(genCmd.Use, ShouldEqual, "api [command] [flags]")
				So(genCmd.HasSubCommands(), ShouldEqual, true)

				subCommands := genCmd.Commands()
				So(len(subCommands), ShouldEqual, 8)
				So(subCommands[0].Name(), ShouldEqual, "count")
				So(subCommands[1].Name(), ShouldEqual, "create")
				So(subCommands[2].Name(), ShouldEqual, "delete")
				So(subCommands[3].Name(), ShouldEqual, "delete-many")
				So(subCommands[4].Name(), ShouldEqual, "get")
				So(subCommands[5].Name(), ShouldEqual, "list")
				So(subCommands[6].Name(), ShouldEqual, "listen")
				So(subCommands[7].Name(), ShouldEqual, "update")

				expectedSubCommands := []string{"list", "task"}

				assertIdentityCommand(subCommands[0], expectedSubCommands)
				assertIdentityCommand(subCommands[1], expectedSubCommands)
				assertIdentityCommand(subCommands[2], expectedSubCommands)
				assertIdentityCommand(subCommands[3], expectedSubCommands)
				assertIdentityCommand(subCommands[4], expectedSubCommands)
				assertIdentityCommand(subCommands[5], expectedSubCommands)
				assertIdentityCommand(subCommands[7], expectedSubCommands)

				// Listen
				listenSubCommands := subCommands[6].Commands()
				So(len(listenSubCommands), ShouldEqual, 0)
			})
		})
	})

}
