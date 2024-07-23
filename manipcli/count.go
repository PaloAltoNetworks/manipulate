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
		RunE: func(cmd *cobra.Command, args []string) error {

			fParam := viper.GetStringSlice("param")
			fTrackingID := viper.GetString(flagTrackingID)
			fRecursive := viper.GetBool(flagRecursive)
			fFilter := viper.GetString(flagFilter)
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
				manipulate.ContextOptionRecursive(fRecursive),
				manipulate.ContextOptionReadConsistency(manipulate.ReadConsistencyStrong),
			}

			if fFilter != "" {
				f, err := elemental.NewFilterFromString(fFilter)
				if err != nil {
					return fmt.Errorf("unable to parse filter %s: %s", fFilter, err)
				}
				options = append(options, manipulate.ContextOptionFilter(f))
			}

			if viper.IsSet(flagParent) {
				parentName, parentID, err := splitParentInfo(viper.GetString(flagParent))
				if err != nil {
					return err
				}

				parent := modelManager.IdentifiableFromString(parentName)
				if parent == nil {
					return fmt.Errorf("unknown identity %s", parentName)
				}
				parent.SetIdentifier(parentID)
				options = append(options, manipulate.ContextOptionParent(parent))
			}

			ctx, cancel := context.WithTimeout(cmd.Context(), 20*time.Second)
			defer cancel()

			mctx := manipulate.NewContext(ctx, options...)
			num, err := manipulator.Count(mctx, identity)
			if err != nil {
				return fmt.Errorf("unable to count %s: %w", identity.Category, err)
			}

			outputType := fOutput
			if fOutput == flagOutputDefault {
				outputType = flagOutputNone
			}

			result, err := formatMaps(
				prepareOutputFormat(outputType, formatTypeCount, fFormatTypeColumn, fOutputTemplate),
				false,
				[]map[string]any{{formatTypeCount: num}},
			)

			if err != nil {
				return fmt.Errorf("unable to format output: %w", err)
			}

			_, _ = fmt.Fprint(cmd.OutOrStdout(), result)

			return nil

		},
	}

	cmd.Flags().BoolP(flagRecursive, "r", false, "List all objects from the current namespace and all child namespaces.")
	cmd.Flags().StringP(flagFilter, "f", "", "Query filter.")
	cmd.Flags().StringP(flagParent, "", "", "Provide information about parent resource. Format `name/ID`")

	return cmd, nil
}
