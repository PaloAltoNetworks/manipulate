package manipcli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.aporeto.io/elemental"
	"go.uber.org/zap"
)

// generateConfig hods the generate configuration.
type generateConfig struct {
	ignoredIdentities map[string]struct{}
}

// GenerateOption represents an option can for the generate command.
type GenerateOption func(generateConfig)

// GenerateOptionIgnoreIdentities sets which non-private identities should be ignored.
func GenerateOptionIgnoreIdentities(identities ...elemental.Identity) GenerateOption {
	return func(g generateConfig) {

		var m = make(map[string]struct{}, len(identities))
		for _, i := range identities {
			m[i.Name] = struct{}{}
		}

		g.ignoredIdentities = m
	}
}

// GenerateCobraCommand generates the API commands and subcommands based on the model manager.
func GenerateCobraCommand(modelManager elemental.ModelManager, manipulatorMaker ManipulatorMaker, options ...GenerateOption) *cobra.Command {

	cfg := generateConfig{}

	for _, opt := range options {
		opt(cfg)
	}

	// API Command
	APICommand := &cobra.Command{
		Use:   "api [command] [flags]",
		Short: "Interact with resources and APIs",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	APICommand.PersistentFlags().StringP(flagOutput, "o", "default", "Format to used print output. Options are 'table', 'json', 'yaml', 'none', 'template' or 'default'.")
	// TODO: Manage output template
	// APICommand.PersistentFlags().String(FlagOutputTemplate, "", "When output is set to 'template', this defines the template to use using Go format.")
	APICommand.PersistentFlags().StringSliceP(formatTypeColumn, "c", nil, "Only show the given columns. Only valid when '--output=table'.")
	APICommand.PersistentFlags().StringSliceP(flagParameters, "p", nil, "Additional parameter to the request, in the form of key=value.")

	// Create command
	createCommands := &cobra.Command{
		Use:   "create",
		Short: "Create a new object",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	// Update command
	updateCommands := &cobra.Command{
		Use:   "update",
		Short: "Update an object",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	// Delete command
	deleteCommands := &cobra.Command{
		Use:   "delete",
		Short: "Delete an object",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	// DeleteMany command
	deleteManyCommands := &cobra.Command{
		Use:   "delete-many",
		Short: "Delete multiple objects",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	deleteManyCommands.PersistentFlags().StringP(flagFilter, "f", "", "Query filter.")
	deleteManyCommands.PersistentFlags().BoolP(flagConfirm, "", false, "Confirm deletion of multiple objects")

	// Get command
	getCommands := &cobra.Command{
		Use:   "get",
		Short: "Get a single object",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	// List command
	listCommands := &cobra.Command{
		Use:   "list",
		Short: "List objects",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	listCommands.PersistentFlags().BoolP(flagRecursive, "r", false, "List all objects from the current namespace and all child namespaces.")
	listCommands.PersistentFlags().IntP(flagPageSize, "S", 0, "Page size to retrieve.")
	listCommands.PersistentFlags().IntP(flagPage, "P", 0, "Page number to retrieve.")
	listCommands.PersistentFlags().StringP(flagFilter, "f", "", "Query filter.")
	listCommands.PersistentFlags().StringSliceP(flagOrder, "O", nil, "Ordering of the result.")

	// Count command
	countCommands := &cobra.Command{
		Use:   "count",
		Short: "Count objects",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	countCommands.PersistentFlags().BoolP(flagRecursive, "r", false, "List all objects from the current namespace and all child namespaces.")
	countCommands.PersistentFlags().StringP(flagFilter, "f", "", "Query filter.")

	// Generate subcommands for each identity
	for _, identity := range modelManager.AllIdentities() {

		if _, ok := cfg.ignoredIdentities[identity.Name]; ok {
			continue
		}

		if identity.Private || identity.Name == "root" {
			continue
		}

		if cmd, err := generateCreateCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			createCommands.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate create command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateUpdateCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			updateCommands.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate update command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateDeleteCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			deleteCommands.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate delete command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateDeleteManyCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			deleteManyCommands.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate delete-many command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateGetCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			getCommands.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate get command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateListCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			listCommands.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate list command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateCountCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			countCommands.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate count command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}
	}

	APICommand.AddCommand(
		createCommands,
		updateCommands,
		deleteCommands,
		deleteManyCommands,
		getCommands,
		listCommands,
		countCommands,
	)

	return APICommand
}
