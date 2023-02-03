package app

import (
	kilnapp "github.com/kilnfi/go-utils/app"
	"github.com/kilnfi/go-utils/log"
	"github.com/kilnfi/near-exporter/pkg/near"
)

func NewApp(config *Config, client *near.Client) (*kilnapp.App, error) {
	cfg := &kilnapp.Config{
		Logger: &log.Config{
			Level:  "info",
			Format: "json",
		},
	}

	app, err := kilnapp.New(cfg.SetDefault())
	if err != nil {
		return nil, err
	}

	exporter := NewExporter(
		config,
		client,
	)

	app.RegisterService(exporter)

	return app, nil
}
