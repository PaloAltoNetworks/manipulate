package manipcli

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"sigs.k8s.io/yaml"
)

// generateUpdateCommandForIdentity generates the command to update an object based on its identity.
func generateUpdateCommandForIdentity(identity elemental.Identity, modelManager elemental.ModelManager, manipulatorMaker ManipulatorMaker) (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:     fmt.Sprintf("%s <id-or-name>", identity.Name),
		Aliases: []string{identity.Category},
		Short:   "Update an existing " + identity.Name,
		Args:    cobra.ExactArgs(1),
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

			retrieveCtx, retrieveCancel := context.WithTimeout(cmd.Context(), 20*time.Second)
			defer retrieveCancel()

			mctx := manipulate.NewContext(retrieveCtx, options...)

			identifiable, err := retrieveObjectByIDOrByName(mctx, manipulator, identity, args[0], modelManager)
			if err != nil {
				return fmt.Errorf("unable to retrieve %s: %w", identity.Name, err)
			}

			if viper.GetBool(flagInteractive) {

				data, err := openInEditor(identifiable, viper.GetString(flagEditor), cmd.Short, true, false, false)
				if err != nil {
					return fmt.Errorf("unable to open editor %s: %w", viper.GetString(flagEditor), err)
				}

				if data == nil {
					return fmt.Errorf("empty data")
				}

				if err := json.Unmarshal(data, identifiable); err != nil {
					return fmt.Errorf("unable to unmarshall: %w", err)
				}

			} else if viper.IsSet(flagInputValues) || viper.IsSet(flagInputData) || viper.IsSet(flagInputFile) || viper.IsSet(flagInputURL) {

				data, err := readData(false)
				if err != nil {
					return fmt.Errorf("unable to read data: %w", err)
				}

				if data != nil {
					data, err = yaml.YAMLToJSON(data)
					if err != nil {
						return err
					}
				}

				if err := json.Unmarshal(data, identifiable); err != nil {
					return fmt.Errorf("unable to unmarshall: %w", err)
				}

			} else {

				if err := readViperFlags(identifiable); err != nil {
					return fmt.Errorf("unable to read flags: %w", err)
				}
			}

			updateCtx, updateCancel := context.WithTimeout(cmd.Context(), 20*time.Second)
			defer updateCancel()

			mctx = manipulate.NewContext(updateCtx, options...)

			if err := manipulator.Update(mctx, identifiable); err != nil {
				return fmt.Errorf("unable to update %s: %w", identity.Name, err)
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

	cmd.Flags().StringP(flagForce, "", "", "Force modification of protected object")
	cmd.Flags().String(flagInputValues, "", "Optional path to file containing templating values")
	cmd.Flags().StringP(flagInputData, "d", "", "Data of the request body in the JSON format.")
	cmd.Flags().StringP(flagInputFile, "f", "", "Optional file to read the data from. Set `-` to read from stdin")
	cmd.Flags().StringP(flagInputURL, "u", "", "Optional url where to read the data from. If you don't set it, stdin or --file will used")
	cmd.Flags().StringSlice(flagInputSet, nil, "Set a value to in the imported data in case it is a Go template.")
	cmd.Flags().Bool(flagPrint, false, "If set will print the raw data. Only works for --file and --url")
	cmd.Flags().Bool(flagRender, false, "If set will render and print the data. Only works for --file and --url")
	cmd.Flags().BoolP(flagInteractive, "i", false, "Set to create the object in the given --editor.")
	cmd.Flags().StringP(flagEditor, "", "vi", "Choose the editor when using --interactive.")

	identifiable := modelManager.IdentifiableFromString(identity.Name)
	if err := setViperFlags(cmd, identifiable, false); err != nil {
		return nil, fmt.Errorf("unable to set flags for %s: %w", identity.Name, err)
	}

	return cmd, nil
}
