package app

import (
	"time"

	"github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:  "http-addr",
		Usage: "http server address",
		Value: ":8080",
	},
	&cli.StringFlag{
		Name:  "log-level",
		Usage: "log level (debug, info, warn, error)",
		Value: "info",
	},
	&cli.StringFlag{
		Name:  "namespace",
		Usage: "prefix for Prometheus metrics",
		Value: "near_validator_watcher",
	},
	&cli.BoolFlag{
		Name:  "no-color",
		Usage: "disable colored output",
	},
	&cli.StringFlag{
		Name:  "node",
		Usage: "rpc node endpoint to connect to",
		Value: "https://rpc.mainnet.near.org",
	},
	&cli.DurationFlag{
		Name:  "refresh-rate",
		Usage: "how often to call the rpc endpoint",
		Value: 10 * time.Second,
	},
	&cli.StringSliceFlag{
		Name:  "validator",
		Usage: "validator pool id to track",
	},
}
