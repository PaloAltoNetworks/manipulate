package manipcli

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.uber.org/zap"
)

// generateDeleteManyCommandForIdentity generates the command to delete many objects based on its identity.
func generateDeleteManyCommandForIdentity(identity elemental.Identity, modelManager elemental.ModelManager, manipulatorMaker ManipulatorMaker) (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:     fmt.Sprintf("%s", identity.Name),
		Aliases: []string{identity.Category},
		Short:   "Delete multiple " + identity.Name,
		// Aliases: TODO: Missing alias from the spec file -> To be stored in the identity ?,
		RunE: func(cmd *cobra.Command, args []string) error {

			fParam := viper.GetStringSlice("param")
			fTrackingID := viper.GetString(flagTrackingID)
			fConfirm := viper.GetBool(flagConfirm)
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
				manipulate.ContextOptionFields(fFormatTypeColumn),
				manipulate.ContextOptionOverride(fConfirm),
			}

			if fFilter != "" {
				f, err := elemental.NewFilterFromString(fFilter)
				if err != nil {
					return fmt.Errorf("unable to parse filter %s: %s", fFilter, err)
				}
				options = append(options, manipulate.ContextOptionFilter(f))
			}

			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			mctx := manipulate.NewContext(ctx, options...)

			identifiables := modelManager.Identifiables(identity)
			if err := manipulator.RetrieveMany(mctx, identifiables); err != nil {
				return fmt.Errorf("unable to retrieve %s: %w", identity.Category, err)
			}

			objects := identifiables.List()

			if !fConfirm {
				for _, item := range objects {
					zap.L().Debug(fmt.Sprintf("- %s with ID=%s will be removed", identity.Name, item.Identifier()))
				}
				return fmt.Errorf("you are about to delete %d %s. If you are sure, please use --%s option to delete", len(objects), identity.Category, flagConfirm)
			}

			var deleted elemental.IdentifiablesList

			errs := elemental.NewErrors()
			for _, o := range objects {
				mctx = mctx.Derive(manipulate.ContextOptionNamespace(o.(elemental.Namespaceable).GetNamespace()))
				if err := manipulator.Delete(mctx, o); err != nil {
					errs = errs.Append(err)
					continue
				}

				deleted = append(deleted, o)
			}

			if len(errs) > 0 {
				return fmt.Errorf("some %s were not deleted: %w", identity.Category, errs)
			}

			outputType := fOutput
			if fOutput == flagOutputDefault {
				outputType = flagOutputNone
			}

			result, err := formatObjects(
				prepareOutputFormat(outputType, formatTypeArray, fFormatTypeColumn, fOutputTemplate),
				true,
				deleted...,
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
