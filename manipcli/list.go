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

// generateListCommandForIdentity generates the command to list all objects based on its identity.
func generateListCommandForIdentity(identity elemental.Identity, modelManager elemental.ModelManager, manipulatorMaker ManipulatorMaker) (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:     identity.Name,
		Aliases: []string{identity.Category},
		Short:   "Allows to list all " + identity.Category,
		// Aliases: TODO: Missing alias from the spec file -> To be stored in the identity ?,
		PersistentPreRunE: persistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {

			fields := viper.GetStringSlice(FormatTypeColumn)

			var dest elemental.Identifiables
			if len(fields) == 0 {
				dest = modelManager.Identifiables(identity)
			} else {
				dest = modelManager.SparseIdentifiables(identity)
			}
			if dest == nil {
				return fmt.Errorf("unable to list %s. unknown identity", identity.Category)
			}

			manipulator, err := manipulatorMaker()
			if err != nil {
				return fmt.Errorf("unable to make manipulator: %w", err)
			}

			parameters, err := parametersToURLValues(viper.GetStringSlice("param"))
			if err != nil {
				return fmt.Errorf("unable to convert parameters to url values: %w", err)
			}

			options := []manipulate.ContextOption{
				manipulate.ContextOptionTracking(viper.GetString(FlagTrackingID), "cli"),
				manipulate.ContextOptionParameters(parameters),
				manipulate.ContextOptionFields(fields),
				manipulate.ContextOptionRecursive(viper.GetBool(FlagRecursive)),
				manipulate.ContextOptionPage(viper.GetInt(FlagPage), viper.GetInt(FlagPageSize)),
				manipulate.ContextOptionOrder(viper.GetStringSlice(FlagOrder)...),
			}

			if viper.IsSet(FlagFilter) {
				filter := viper.GetString(FlagFilter)
				f, err := elemental.NewFilterFromString(filter)
				if err != nil {
					return fmt.Errorf("unable to parse filter %s: %s", filter, err)
				}

				options = append(options, manipulate.ContextOptionFilter(f))
			}

			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			mctx := manipulate.NewContext(ctx, options...)
			if err := manipulator.RetrieveMany(mctx, dest); err != nil {
				return fmt.Errorf("unable to retrieve all %s: %w", identity.Category, err)
			}

			output := viper.GetString(FlagOutput)
			outputType := output
			if output == FlagOutputDefault {
				outputType = FlagOutputJSON
			}

			result, err := FormatObjects(
				prepareOutputFormat(outputType, FormatTypeArray, viper.GetStringSlice(FormatTypeColumn), viper.GetString(FlagOutputTemplate)),
				true,
				dest.List()...,
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
