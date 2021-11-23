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
		Short:   "List all " + identity.Category,
		// Aliases: TODO: Missing alias from the spec file -> To be stored in the identity ?,
		RunE: func(cmd *cobra.Command, args []string) error {

			fParam := viper.GetStringSlice("param")
			fTrackingID := viper.GetString(flagTrackingID)
			fRecursive := viper.GetBool(flagRecursive)
			fPage := viper.GetInt(flagPage)
			fPageSize := viper.GetInt(flagPageSize)
			fOrder := viper.GetStringSlice(flagOrder)
			fFilter := viper.GetString(flagFilter)
			fOutput := viper.GetString(flagOutput)
			fFormatTypeColumn := viper.GetStringSlice(formatTypeColumn)
			fOutputTemplate := viper.GetString(flagOutputTemplate)

			var dest elemental.Identifiables
			if len(fFormatTypeColumn) == 0 {
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

			parameters, err := parametersToURLValues(fParam)
			if err != nil {
				return fmt.Errorf("unable to convert parameters to url values: %w", err)
			}

			options := []manipulate.ContextOption{
				manipulate.ContextOptionTracking(fTrackingID, "cli"),
				manipulate.ContextOptionParameters(parameters),
				manipulate.ContextOptionFields(fFormatTypeColumn),
				manipulate.ContextOptionRecursive(fRecursive),
				manipulate.ContextOptionPage(fPage, fPageSize),
				manipulate.ContextOptionOrder(fOrder...),
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
			if err := manipulator.RetrieveMany(mctx, dest); err != nil {
				return fmt.Errorf("unable to retrieve all %s: %w", identity.Category, err)
			}

			outputType := fOutput
			if fOutput == flagOutputDefault {
				outputType = flagOutputJSON
			}

			result, err := formatObjects(
				prepareOutputFormat(outputType, formatTypeArray, fFormatTypeColumn, fOutputTemplate),
				true,
				dest.List()...,
			)

			if err != nil {
				return fmt.Errorf("unable to format output: %w", err)
			}

			fmt.Fprint(cmd.OutOrStdout(), result)
			return nil
		},
	}

	cmd.Flags().BoolP(flagRecursive, "r", false, "List all objects from the current namespace and all child namespaces.")
	cmd.Flags().IntP(flagPageSize, "S", 0, "Page size to retrieve.")
	cmd.Flags().IntP(flagPage, "P", 0, "Page number to retrieve.")
	cmd.Flags().StringP(flagFilter, "f", "", "Query filter.")
	cmd.Flags().StringSliceP(flagOrder, "O", nil, "Ordering of the result.")
	cmd.Flags().StringP(flagParent, "", "", "Provide information about parent resource. Format `name/ID`")

	return cmd, nil
}
