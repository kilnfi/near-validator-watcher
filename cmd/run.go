package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kilnfi/near-exporter/pkg/app"
	"github.com/kilnfi/near-exporter/pkg/near"
	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run exporter",
		RunE: func(cmd *cobra.Command, args []string) error {
			rpcAddr := os.Getenv("NEAR_EXPORTER_RPC_ADDR")
			if rpcAddr == "" {
				return fmt.Errorf("missing env var NEAR_EXPORTER_RPC_ADDR")
			}
			trackedAccounts := strings.Split(
				os.Getenv("NEAR_EXPORTER_TRACKED_ACCOUNTS"), ",",
			)

			client := near.NewClient(rpcAddr, &http.Client{
				Timeout: time.Duration(10 * time.Second),
			})

			config := &app.Config{
				TrackedAccounts: trackedAccounts,
				RefreshRate:     15 * time.Second,
			}

			app, err := app.NewApp(config, client)
			if err != nil {
				return err
			}

			app.Logger().SetOutput(cmd.OutOrStderr())

			app.Logger().
				WithField("rpc_addr", rpcAddr).
				WithField("tracked_account", trackedAccounts).
				Info("config")

			return app.Run()
		},
	}

	return cmd
}
