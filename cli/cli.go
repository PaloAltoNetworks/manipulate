package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.aporeto.io/elemental"
)

func GenerateCobraCommand(modelManager elemental.ModelManager) *cobra.Command {

	// API Command
	APICommand := &cobra.Command{
		Use:   "api [command] [flags]",
		Short: "Allows to make api calls",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	APICommand.PersistentFlags().StringP(FlagNamespace, "n", "", "Namespace to use.")
	APICommand.PersistentFlags().StringP(FlagOutput, "o", "default", "Format to used print output. Options are 'table', 'json', 'yaml', 'none', 'template' or 'default'.")
	// TODO: Manage output template
	// APICommand.PersistentFlags().String(FlagOutputTemplate, "", "When output is set to 'template', this defines the template to use using Go format.")
	APICommand.PersistentFlags().StringP(FlagToken, "t", "", "Token to use.")
	APICommand.PersistentFlags().StringSliceP(FormatTypeColumn, "c", nil, "Only show the given columns. Only valid when '--output=table'.")
	APICommand.PersistentFlags().StringSliceP(FlagParameters, "p", nil, "Additional parameter to the request, in the form of key=value.")

	// Create command
	createCommands := &cobra.Command{
		Use:   "create",
		Short: "Allows to create a new object",
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
		Short: "Allows to update an object",
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
		Short: "Allows to delete an object",
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
		Short: "Allows to delete multiple objects",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	deleteManyCommands.PersistentFlags().StringP(FlagFilter, "f", "", "Query filter.")
	deleteManyCommands.PersistentFlags().BoolP(FlagConfirm, "", false, "Confirm deletion of multiple objects")

	// Get command
	getCommands := &cobra.Command{
		Use:   "get",
		Short: "Allows to get an object",
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
		Short: "Allows to list all objects",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	listCommands.PersistentFlags().BoolP(FlagRecursive, "r", false, "List all objects from the current namespace and all child namespaces.")
	listCommands.PersistentFlags().IntP(FlagPageSize, "S", 0, "Page size to retrieve.")
	listCommands.PersistentFlags().IntP(FlagPage, "P", 0, "Page number to retrieve.")
	listCommands.PersistentFlags().StringP(FlagFilter, "f", "", "Query filter.")
	listCommands.PersistentFlags().StringSliceP(FlagOrder, "O", nil, "Ordering of the result.")

	// Count command
	countCommands := &cobra.Command{
		Use:   "count",
		Short: "Allows to count objects",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Root().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			return viper.BindPFlags(cmd.PersistentFlags())
		},
	}

	countCommands.PersistentFlags().BoolP(FlagRecursive, "r", false, "List all objects from the current namespace and all child namespaces.")
	countCommands.PersistentFlags().StringP(FlagFilter, "f", "", "Query filter.")

	now := time.Now()

	// Generate subcommands for each identity
	for _, identity := range modelManager.AllIdentities() {

		if identity.Private {
			continue
		}

		if cmd, err := generateCreateCommandForIdentity(identity, modelManager); err == nil {
			createCommands.AddCommand(cmd)
		}

		if cmd, err := generateUpdateCommandForIdentity(identity, modelManager); err == nil {
			updateCommands.AddCommand(cmd)
		}

		if cmd, err := generateDeleteCommandForIdentity(identity, modelManager); err == nil {
			deleteCommands.AddCommand(cmd)
		}

		if cmd, err := generateDeleteManyCommandForIdentity(identity, modelManager); err == nil {
			deleteManyCommands.AddCommand(cmd)
		}

		if cmd, err := generateGetCommandForIdentity(identity, modelManager); err == nil {
			getCommands.AddCommand(cmd)
		}

		if cmd, err := generateListCommandForIdentity(identity, modelManager); err == nil {
			listCommands.AddCommand(cmd)
		}

		if cmd, err := generateCountCommandForIdentity(identity, modelManager); err == nil {
			countCommands.AddCommand(cmd)
		}
	}

	APICommand.AddCommand(createCommands)
	APICommand.AddCommand(updateCommands)
	APICommand.AddCommand(deleteCommands)
	APICommand.AddCommand(deleteManyCommands)
	APICommand.AddCommand(getCommands)
	APICommand.AddCommand(listCommands)
	APICommand.AddCommand(countCommands)

	fmt.Println("time elasped", time.Since(now))

	return APICommand
}
