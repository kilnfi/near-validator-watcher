package main

import (
	"context"
	"errors"
	"os"

	"github.com/kilnfi/near-validator-watcher/pkg/app"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var Version = "v0.0.0-dev" // generated at build time

func main() {
	app := &cli.App{
		Name:    "near-validator-watcher",
		Usage:   "NEAR validators monitoring tool",
		Flags:   app.Flags,
		Action:  app.RunFunc,
		Version: Version,
	}

	if err := app.Run(os.Args); err != nil && !errors.Is(err, context.Canceled) {
		logrus.WithError(err).Fatal("application failed")
	}
}
