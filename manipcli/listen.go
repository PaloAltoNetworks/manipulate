package manipcli

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.aporeto.io/elemental"
	"go.aporeto.io/manipulate"
	"go.aporeto.io/manipulate/maniphttp"
	"go.uber.org/zap"
)

// generateListenCommandForIdentity generates the command to listen for events based on its identity.
func generateListenCommand(modelManager elemental.ModelManager, manipulatorMaker ManipulatorMaker) (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:   "listen",
		Short: "Listen for events",
		RunE: func(cmd *cobra.Command, args []string) error {

			fRecursive := viper.GetBool(flagRecursive)
			fOutput := viper.GetString(flagOutput)
			fFormatTypeColumn := viper.GetStringSlice(formatTypeColumn)
			fOutputTemplate := viper.GetString(flagOutputTemplate)

			manipulator, err := manipulatorMaker()
			if err != nil {
				return fmt.Errorf("unable to make manipulator: %w", err)
			}

			subscriber := maniphttp.NewSubscriber(
				manipulator,
				maniphttp.SubscriberOptionRecursive(fRecursive),
				maniphttp.SubscriberOptionSupportErrorEvents(),
			)
			if err != nil {
				return err
			}

			var filter *elemental.PushConfig
			filterIdentities := viper.GetStringSlice("identity")

			if len(filterIdentities) > 0 {

				filter = elemental.NewPushConfig()

				for _, i := range filterIdentities {
					identity := modelManager.IdentityFromAny(i)
					if identity.IsEmpty() {
						return fmt.Errorf("unknown identity %s", i)
					}
					filter.FilterIdentity(identity.Name)
				}
			}

			outputType := fOutput
			if fOutput == flagOutputDefault {
				outputType = flagOutputJSON
			}
			outputFormat := prepareOutputFormat(outputType, formatTypeHash, fFormatTypeColumn, fOutputTemplate)

			terminated := make(chan struct{})
			go func() {

				pullErrorIfAny := func() error {
					select {
					case err := <-subscriber.Errors():
						return err
					default:
						return nil
					}
				}

				for {
					select {

					case evt := <-subscriber.Events():
						result, err := formatEvents(outputFormat, false, evt)
						if err != nil {
							zap.L().Error("unable to format event", zap.Error(err))
						}

						fmt.Fprintf(cmd.OutOrStdout(), result)

					case st := <-subscriber.Status():
						switch st {
						case manipulate.SubscriberStatusInitialConnection:
							zap.L().Debug("status update", zap.String("status", "connected"))
						case manipulate.SubscriberStatusInitialConnectionFailure:
							zap.L().Warn("status update", zap.String("status", "connect failed"), zap.Error(pullErrorIfAny()))
						case manipulate.SubscriberStatusDisconnection:
							zap.L().Warn("status update", zap.String("status", "disconnected"), zap.Error(pullErrorIfAny()))
						case manipulate.SubscriberStatusReconnection:
							zap.L().Info("status update", zap.String("status", "reconnected"))
						case manipulate.SubscriberStatusReconnectionFailure:
							zap.L().Debug("status update", zap.String("status", "reconnection failed"), zap.Error(pullErrorIfAny()))
						case manipulate.SubscriberStatusFinalDisconnection:
							zap.L().Debug("status update", zap.String("status", "terminated"))
							close(terminated)
						}

					case err := <-subscriber.Errors():
						zap.L().Error("Error received", zap.Error(err))
					}
				}
			}()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			subscriber.Start(ctx, filter)

			c := make(chan os.Signal, 1)
			signal.Reset(os.Interrupt)
			signal.Notify(c, os.Interrupt)

			<-c

			cancel()
			<-terminated

			return nil
		},
	}

	cmd.Flags().BoolP(flagRecursive, "r", false, "Listen to all events in the current namespace and all child namespaces.")
	cmd.Flags().StringSliceP("identity", "i", []string{}, "Only display events for the given identities.")

	return cmd, nil
}
