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

// generateUpdateCommandForIdentity generates the command to update an object based on its identity.
func generateUpdateCommandForIdentity(identity elemental.Identity, modelManager elemental.ModelManager, manipulatorMaker ManipulatorMaker) (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:     fmt.Sprintf("%s <id-or-name>", identity.Name),
		Aliases: []string{identity.Category},
		Short:   "Allows to update an existing " + identity.Name,
		Args:    cobra.ExactArgs(1),
		// Aliases: TODO: Missing alias from the spec file -> To be stored in the identity ?,
		PersistentPreRunE: persistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {

			manipulator, err := manipulatorMaker()
			if err != nil {
				return fmt.Errorf("unable to make manipulator: %w", err)
			}

			parameters, err := parametersToURLValues(viper.GetStringSlice(flagParameters))
			if err != nil {
				return fmt.Errorf("unable to convert parameters to url values: %w", err)
			}

			options := []manipulate.ContextOption{
				manipulate.ContextOptionTracking(viper.GetString(flagTrackingID), "cli"),
				manipulate.ContextOptionParameters(parameters),
				manipulate.ContextOptionFields(viper.GetStringSlice(formatTypeColumn)),
			}

			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			mctx := manipulate.NewContext(ctx, options...)

			identifiable, err := retrieveObjectByIDOrByName(mctx, manipulator, identity, args[0], modelManager)
			if err != nil {
				return fmt.Errorf("unable to retrieve %s: %w", identity.Name, err)
			}

			if err := readViperFlags(identifiable); err != nil {
				return fmt.Errorf("unable to read flags: %w", err)
			}

			if err := manipulator.Update(mctx, identifiable); err != nil {
				return fmt.Errorf("unable to create %s: %w", identity.Name, err)
			}

			output := viper.GetString(flagOutput)
			outputType := output
			if output == flagOutputDefault {
				outputType = flagOutputNone
			}

			result, err := formatObjects(
				prepareOutputFormat(outputType, formatTypeHash, viper.GetStringSlice(formatTypeColumn), viper.GetString(flagOutputTemplate)),
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

	identifiable := modelManager.IdentifiableFromString(identity.Name)
	if err := setViperFlags(cmd, identifiable, false); err != nil {
		return nil, fmt.Errorf("unable to set flags for %s: %w", identity.Name, err)
	}

	return cmd, nil
}
