package cmd

import (
	"context"

	kilnutils "github.com/kilnfi/go-utils/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRootCommand() *cobra.Command {
	v := viper.New()

	ctx := context.Background()
	ctx = kilnutils.WithViper(ctx, v)

	cmd := &cobra.Command{
		Use: "near-exporter",
	}

	cmd.SetContext(ctx)
	cmd.AddCommand(NewRunCommand())

	return cmd
}
