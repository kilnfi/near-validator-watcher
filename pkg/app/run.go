package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/kilnfi/near-validator-watcher/pkg/metrics"
	"github.com/kilnfi/near-validator-watcher/pkg/near"
	"github.com/kilnfi/near-validator-watcher/pkg/watcher"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

func RunFunc(cCtx *cli.Context) error {
	var (
		ctx = cCtx.Context

		// Config flags
		httpAddr    = cCtx.String("http-addr")
		logLevel    = cCtx.String("log-level")
		namespace   = cCtx.String("namespace")
		noColor     = cCtx.Bool("no-color")
		node        = cCtx.String("node")
		refreshRate = cCtx.Duration("refresh-rate")
		validators  = cCtx.StringSlice("validator")
	)

	//
	// Setup
	//
	// Logger setup
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logLevelFromString(logLevel))
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: noColor,
	})

	// Disable colored output if requested
	color.NoColor = noColor

	// Handle signals via context
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Create errgroup to manage all goroutines
	errg, ctx := errgroup.WithContext(ctx)

	//
	// Validator Watcher
	//
	logrus.Infof("connecting to node %s", node)
	registry := prometheus.NewRegistry()
	metrics := metrics.New(namespace)
	metrics.Register(registry)
	client := near.NewClient(node, &http.Client{})
	watcher := watcher.New(client, metrics, &watcher.Config{
		Writer:          os.Stdout,
		TrackedAccounts: validators,
		RefreshRate:     refreshRate,
	})
	errg.Go(func() error {
		return watcher.Start(ctx)
	})

	//
	// HTTP server
	//
	logrus.Infof("starting HTTP server on %s", httpAddr)
	readyProbe := func() bool {
		return !watcher.IsSyncing()
	}
	httpServer := NewHTTPServer(
		httpAddr,
		WithReadyProbe(readyProbe),
		WithLiveProbe(upProbe),
		WithMetrics(registry),
	)
	errg.Go(func() error {
		return httpServer.Run()
	})

	//
	// Wait for context to be cancelled (via signals or error from errgroup)
	//
	<-ctx.Done()
	logrus.Info("shutting down")

	//
	// Stop all watchers and exporter
	//
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error().Err(fmt.Errorf("failed to stop http server: %w", err)).Msg("")
	}

	// Wait for all goroutines to finish
	return errg.Wait()
}
