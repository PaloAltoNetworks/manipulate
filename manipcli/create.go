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

// generateCreateCommandForIdentity generates the command to create an object based on its identity.
func generateCreateCommandForIdentity(identity elemental.Identity, modelManager elemental.ModelManager, manipulatorMaker ManipulatorMaker) (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:     identity.Name,
		Aliases: []string{identity.Category},
		Short:   "Create a new " + identity.Name,
		// Aliases: TODO: Missing alias from the spec file -> To be stored in the identity ?,
		RunE: func(cmd *cobra.Command, args []string) error {

			fParam := viper.GetStringSlice("param")
			fTrackingID := viper.GetString(flagTrackingID)
			fOutput := viper.GetString(flagOutput)
			fFormatTypeColumn := viper.GetStringSlice(formatTypeColumn)
			fOutputTemplate := viper.GetString(flagOutputTemplate)

			manipulator, err := manipulatorMaker()
			if err != nil {
				return fmt.Errorf("unable to make manipulator: %w", err)
			}

			parameters, err := parametersToURLValues(fParam)
			if err != nil {
				return fmt.Errorf("unable to convert parameters to url values: %w", err)
			}

			options := []manipulate.ContextOption{
				manipulate.ContextOptionTracking(fTrackingID, "cli"),
				manipulate.ContextOptionParameters(parameters),
				manipulate.ContextOptionFields(fFormatTypeColumn),
			}

			identifiable := modelManager.IdentifiableFromString(identity.Name)
			if err := readViperFlags(identifiable); err != nil {
				return fmt.Errorf("unable to read flags: %w", err)
			}

			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			mctx := manipulate.NewContext(ctx, options...)
			if err := manipulator.Create(mctx, identifiable); err != nil {
				return fmt.Errorf("unable to create %s: %w", identity.Name, err)
			}

			outputType := fOutput
			if fOutput == flagOutputDefault {
				outputType = flagOutputNone
			}

			result, err := formatObjects(
				prepareOutputFormat(outputType, formatTypeHash, fFormatTypeColumn, fOutputTemplate),
				false,
				identifiable,
			)

			if err != nil {
				return fmt.Errorf("unable to format output: %w", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), result)
			return nil
		},
	}

	identifiable := modelManager.IdentifiableFromString(identity.Name)
	if err := setViperFlags(cmd, identifiable, true); err != nil {
		return nil, fmt.Errorf("unable to set flags for %s: %w", identity.Name, err)
	}

	return cmd, nil
}
