package manipcli

import (
	"github.com/spf13/cobra"
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

	rootCmd := &cobra.Command{
		Use:   "api [command] [flags]",
		Short: "Interact with resources and APIs",
	}
	rootCmd.PersistentFlags().StringP(flagOutput, "o", "default", "Format to used print output. Options are 'table', 'json', 'yaml', 'none', 'template' or 'default'.")
	rootCmd.PersistentFlags().StringSliceP(formatTypeColumn, "c", nil, "Only show the given columns. Only valid when '--output=table'.")
	rootCmd.PersistentFlags().StringSliceP(flagParameters, "p", nil, "Additional parameter to the request, in the form of key=value.")

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new object",
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update an object",
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an object",
	}

	deleteManyCmd := &cobra.Command{
		Use:   "delete-many",
		Short: "Delete multiple objects",
	}
	deleteManyCmd.PersistentFlags().StringP(flagFilter, "f", "", "Query filter.")
	deleteManyCmd.PersistentFlags().BoolP(flagConfirm, "", false, "Confirm deletion of multiple objects")

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get a single object",
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List objects",
	}
	listCmd.PersistentFlags().BoolP(flagRecursive, "r", false, "List all objects from the current namespace and all child namespaces.")
	listCmd.PersistentFlags().IntP(flagPageSize, "S", 0, "Page size to retrieve.")
	listCmd.PersistentFlags().IntP(flagPage, "P", 0, "Page number to retrieve.")
	listCmd.PersistentFlags().StringP(flagFilter, "f", "", "Query filter.")
	listCmd.PersistentFlags().StringSliceP(flagOrder, "O", nil, "Ordering of the result.")

	countCmd := &cobra.Command{
		Use:   "count",
		Short: "Count objects",
	}
	countCmd.PersistentFlags().BoolP(flagRecursive, "r", false, "List all objects from the current namespace and all child namespaces.")
	countCmd.PersistentFlags().StringP(flagFilter, "f", "", "Query filter.")

	// Generate subcommands for each identity
	for _, identity := range modelManager.AllIdentities() {

		if _, ok := cfg.ignoredIdentities[identity.Name]; ok {
			continue
		}

		if identity.Private || identity.Name == "root" {
			continue
		}

		if cmd, err := generateCreateCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			createCmd.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate create command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateUpdateCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			updateCmd.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate update command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateDeleteCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			deleteCmd.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate delete command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateDeleteManyCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			deleteManyCmd.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate delete-many command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateGetCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			getCmd.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate get command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateListCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			listCmd.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate list command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}

		if cmd, err := generateCountCommandForIdentity(identity, modelManager, manipulatorMaker); err == nil {
			countCmd.AddCommand(cmd)
		} else {
			zap.L().Debug("unable to generate count command for identity", zap.String("identity", identity.Name), zap.Error(err))
		}
	}

	rootCmd.AddCommand(
		createCmd,
		updateCmd,
		deleteCmd,
		deleteManyCmd,
		getCmd,
		listCmd,
		countCmd,
	)

	return rootCmd
}
