package cli

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

// A Namespaceable is an entity that has a namespace.
type Namespaceable interface {
	SetNamespace(string)
	GetNamespace() string
}

// generateDeleteManyCommandForIdentity generates the command to delete many objects based on its identity.
func generateDeleteManyCommandForIdentity(identity elemental.Identity, modelManager elemental.ModelManager) (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:     fmt.Sprintf("%s", identity.Name),
		Aliases: []string{identity.Category},
		Short:   "Allows to delete multiple " + identity.Name,
		// Aliases: TODO: Missing alias from the spec file -> To be stored in the identity ?,
		PersistentPreRunE: persistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {

			manipulator, err := prepareManipulator(
				viper.GetString(FlagAPI),
				viper.GetString(FlagToken),
				viper.GetString(FlagAppCredentials),
				viper.GetString(FlagNamespace),
				viper.GetString(FlagCACertPath),
				viper.GetBool(FlagAPISkipVerify),
				viper.GetString(FlagEncoding),
			)
			if err != nil {
				return fmt.Errorf("unable to prepare manipulator: %w", err)
			}

			parameters, err := parametersToURLValues(viper.GetStringSlice(FlagParameters))
			if err != nil {
				return fmt.Errorf("unable to convert parameters to url values: %w", err)
			}

			options := []manipulate.ContextOption{
				manipulate.ContextOptionTracking(viper.GetString(FlagTrackingID), "cli"),
				manipulate.ContextOptionParameters(parameters),
				manipulate.ContextOptionFields(viper.GetStringSlice(FormatTypeColumn)),
				manipulate.ContextOptionOverride(viper.GetBool(FlagConfirm)),
			}

			if viper.IsSet(FlagFilter) {
				filter := viper.GetString(FlagFilter)
				f, err := elemental.NewFilterFromString(filter)
				if err != nil {
					return fmt.Errorf("unable to parse filter %s: %s", filter, err)
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

			if !viper.IsSet(FlagConfirm) {
				for _, item := range objects {
					zap.L().Debug(fmt.Sprintf("- %s with ID=%s will be removed", identity.Name, item.Identifier()))
				}
				return fmt.Errorf("You are about to delete %d %s. If you are sure, please use --%s option to delete.", len(objects), identity.Category, FlagConfirm)
			}

			var deleted elemental.IdentifiablesList

			errs := elemental.NewErrors()
			for _, o := range objects {
				mctx = mctx.Derive(manipulate.ContextOptionNamespace(o.(Namespaceable).GetNamespace()))
				if err := manipulator.Delete(mctx, o); err != nil {
					errs = errs.Append(err)
					continue
				}

				deleted = append(deleted, o)
			}

			if len(errs) > 0 {
				return fmt.Errorf("some %s were not deleted: %w", identity.Category, errs.Error())
			}

			output := viper.GetString(FlagOutput)
			outputType := output
			if output == FlagOutputDefault {
				outputType = FlagOutputNone
			}

			result, err := FormatObjects(
				prepareOutputFormat(outputType, FormatTypeArray, viper.GetStringSlice(FormatTypeColumn), viper.GetString(FlagOutputTemplate)),
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
