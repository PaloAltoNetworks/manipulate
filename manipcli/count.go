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

// generateCountCommandForIdentity generates the command to count all objects based on its identity.
func generateCountCommandForIdentity(identity elemental.Identity, modelManager elemental.ModelManager, manipulatorMaker ManipulatorMaker) (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:     identity.Name,
		Aliases: []string{identity.Category},
		Short:   "Count " + identity.Category,
		// Aliases: TODO: Missing alias from the spec file -> To be stored in the identity ?,
		PersistentPreRunE: persistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {

			manipulator, err := manipulatorMaker()
			if err != nil {
				return fmt.Errorf("unable to make manipulator: %w", err)
			}

			parameters, err := parametersToURLValues(viper.GetStringSlice("param"))
			if err != nil {
				return fmt.Errorf("unable to convert parameters to url values: %w", err)
			}

			options := []manipulate.ContextOption{
				manipulate.ContextOptionTracking(viper.GetString(flagTrackingID), "cli"),
				manipulate.ContextOptionParameters(parameters),
				manipulate.ContextOptionRecursive(viper.GetBool(flagRecursive)),
				manipulate.ContextOptionReadConsistency(manipulate.ReadConsistencyStrong),
			}

			if viper.IsSet(flagFilter) {
				filter := viper.GetString(flagFilter)
				f, err := elemental.NewFilterFromString(filter)
				if err != nil {
					return fmt.Errorf("unable to parse filter %s: %s", filter, err)
				}

				options = append(options, manipulate.ContextOptionFilter(f))
			}

			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			mctx := manipulate.NewContext(ctx, options...)
			num, err := manipulator.Count(mctx, identity)
			if err != nil {
				return fmt.Errorf("unable to count %s: %w", identity.Category, err)
			}

			output := viper.GetString(flagOutput)
			outputType := output
			if output == flagOutputDefault {
				outputType = flagOutputNone
			}

			result, err := formatMaps(
				prepareOutputFormat(outputType, formatTypeCount, viper.GetStringSlice(formatTypeColumn), viper.GetString(flagOutputTemplate)),
				false,
				[]map[string]interface{}{{formatTypeCount: num}},
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
