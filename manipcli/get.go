package manipcli

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
)

// generateGetCommandForIdentity generates the command to get an object based on its identity.
func generateGetCommandForIdentity(identity elemental.Identity, modelManager elemental.ModelManager, manipulatorMaker ManipulatorMaker) (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:     fmt.Sprintf("%s <id-or-name>", identity.Name),
		Aliases: []string{identity.Category},
		Short:   "Allows to get an existing " + identity.Name,
		Args:    cobra.ExactArgs(1),
		// Aliases: TODO: Missing alias from the spec file -> To be stored in the identity ?,
		PersistentPreRunE: persistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {

			manipulator, err := manipulatorMaker()
			if err != nil {
				return fmt.Errorf("unable to make manipulator: %w", err)
			}

			parameters, err := parametersToURLValues(viper.GetStringSlice(FlagParameters))
			if err != nil {
				return fmt.Errorf("unable to convert parameters to url values: %w", err)
			}

			options := []manipulate.ContextOption{
				manipulate.ContextOptionTracking(viper.GetString(FlagTrackingID), "cli"),
				manipulate.ContextOptionParameters(parameters),
				manipulate.ContextOptionFields(viper.GetStringSlice(FormatTypeColumn)),
				manipulate.ContextOptionRecursive(viper.GetBool(FlagRecursive)),
			}

			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			mctx := manipulate.NewContext(ctx, options...)

			identifiable, err := retrieveObjectByIDOrByName(mctx, manipulator, identity, args[0], modelManager)
			if err != nil {
				return fmt.Errorf("unable to retrieve %s: %w", identity.Name, err)
			}

			output := viper.GetString(FlagOutput)
			outputType := output
			if output == FlagOutputDefault {
				outputType = FlagOutputJSON
			}

			result, err := FormatObjects(
				prepareOutputFormat(outputType, FormatTypeHash, viper.GetStringSlice(FormatTypeColumn), viper.GetString(FlagOutputTemplate)),
				false,
				identifiable,
			)

			if err != nil {
				return fmt.Errorf("unable to format output: %w", err)
			}

			fmt.Println(result)
			return nil
		},
	}

	return cmd, nil
}
