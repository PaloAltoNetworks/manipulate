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

// generateDeleteCommandForIdentity generates the command to delete an object based on its identity.
func generateDeleteCommandForIdentity(identity elemental.Identity, modelManager elemental.ModelManager, manipulatorMaker ManipulatorMaker) (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:     fmt.Sprintf("%s <id-or-name>", identity.Name),
		Aliases: []string{identity.Category},
		Short:   "Delete an existing " + identity.Name,
		Args:    cobra.ExactArgs(1),
		// Aliases: TODO: Missing alias from the spec file -> To be stored in the identity ?,
		RunE: func(cmd *cobra.Command, args []string) error {

			fParam := viper.GetStringSlice("param")
			fTrackingID := viper.GetString(flagTrackingID)
			fOutput := viper.GetString(flagOutput)
			fFormatTypeColumn := viper.GetStringSlice(formatTypeColumn)
			fOutputTemplate := viper.GetString(flagOutputTemplate)
			fForce := viper.GetBool(flagForce)

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
				manipulate.ContextOptionOverride(fForce),
			}

			ctx, cancel := context.WithTimeout(cmd.Context(), 20*time.Second)
			defer cancel()

			mctx := manipulate.NewContext(ctx, options...)

			identifiable, err := retrieveObjectByIDOrByName(mctx, manipulator, identity, args[0], modelManager)
			if err != nil {
				return fmt.Errorf("unable to retrieve %s: %w", identity.Name, err)
			}

			if err := manipulator.Delete(mctx, identifiable); err != nil {
				return fmt.Errorf("unable to delete %s: %w", identity.Name, err)
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

			fmt.Fprint(cmd.OutOrStdout(), result)
			return nil
		},
	}

	cmd.Flags().BoolP(flagForce, "", false, "Force deletion of protected object")

	return cmd, nil
}
