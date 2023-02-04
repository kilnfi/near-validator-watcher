package app

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/kilnfi/near-exporter/pkg/near"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type Exporter struct {
	config *Config
	client *near.Client

	logger  logrus.FieldLogger
	metrics *Metrics

	errg   *errgroup.Group
	cancel context.CancelFunc
}

func NewExporter(config *Config, client *near.Client) *Exporter {
	return &Exporter{
		config:  config,
		client:  client,
		metrics: NewMetrics(),
		logger:  logrus.StandardLogger(),
	}
}

func (e *Exporter) SetLogger(logger logrus.FieldLogger) {
	e.logger = logger.WithField("service", "exporter")
}

func (e *Exporter) RegisterMetrics(reg prometheus.Registerer) error {
	return e.metrics.RegisterMetrics(reg)
}

func (e *Exporter) Start(ctx context.Context) error {
	ctx, e.cancel = context.WithCancel(context.Background())
	e.errg, ctx = errgroup.WithContext(ctx)

	e.errg.Go(func() error {
		ticker := time.NewTicker(e.config.RefreshRate)

		for {
			tick := func() error {
				return e.tick(ctx)
			}

			backOff := backoff.WithMaxRetries(
				backoff.NewExponentialBackOff(),
				3,
			)
			notify := func(err error, d time.Duration) {
				e.logger.WithError(err).Error("tick failed, retrying...")
			}

			if err := backoff.RetryNotify(tick, backOff, notify); err != nil {
				e.logger.WithError(err).Error("tick run failed")
			}

			select {
			case <-ctx.Done():
				return nil
			case <-ticker.C:
				continue
			}
		}
	})

	return nil
}

func (e *Exporter) Stop(ctx context.Context) error {
	e.cancel()
	return e.errg.Wait()
}

func (e *Exporter) tick(ctx context.Context) error {
	e.logger.Info("collect near data")

	var errg *errgroup.Group

	errg, ctx = errgroup.WithContext(ctx)

	errg.Go(func() error {
		return e.collectProtocolConfig(ctx)
	})
	errg.Go(func() error {
		return e.collectStatus(ctx)
	})
	errg.Go(func() error {
		return e.collectValidators(ctx)
	})

	return errg.Wait()
}

func (e *Exporter) collectProtocolConfig(ctx context.Context) error {
	config, err := e.client.ProtocolConfig(ctx)
	if err != nil {
		return err
	}

	e.metrics.EpochLength.Set(float64(config.EpochLength))
	e.metrics.ProtocolVersion.Set(float64(config.ProtocolVersion))

	return nil
}

func (e *Exporter) collectStatus(ctx context.Context) error {
	status, err := e.client.Status(ctx)
	if err != nil {
		return err
	}

	isSyncing := 0
	if status.SyncInfo.Syncing {
		isSyncing = 1
	}

	e.metrics.BlockNumber.Set(float64(status.SyncInfo.LatestBlockHeight))
	e.metrics.ChainID.WithLabelValues(status.ChainID).Set(float64(HashString(status.ChainID)))
	e.metrics.SyncingDesc.Set(float64(isSyncing))
	e.metrics.VersionBuild.WithLabelValues(status.Version.Version, status.Version.Build).Set(float64(HashString(status.Version.Build)))

	return nil
}

func (e *Exporter) collectValidators(ctx context.Context) error {
	validators, err := e.client.Validators(ctx, "latest")
	if err != nil {
		return err
	}

	// Reset labeled gauge vec
	e.metrics.ValidatorExpectedBlocks.Reset()
	e.metrics.ValidatorExpectedChunks.Reset()
	e.metrics.ValidatorProducedBlocks.Reset()
	e.metrics.ValidatorProducedChunks.Reset()
	e.metrics.ValidatorSlashed.Reset()
	e.metrics.ValidatorStake.Reset()
	e.metrics.NextValidatorStake.Reset()
	e.metrics.CurrentProposals.Reset()
	e.metrics.PrevEpochKickout.Reset()

	labelEpochStartHeight := strconv.FormatInt(validators.EpochStartHeight, 10)

	e.metrics.EpochStartHeight.Set(float64(validators.EpochStartHeight))

	var seatPrice float64

	for _, v := range validators.CurrentValidators {
		isSlashed := 0
		if v.IsSlashed {
			isSlashed = 1
		}

		labels := []string{v.AccountId, v.PublicKey, labelEpochStartHeight, e.isTracked(v.AccountId)}

		e.metrics.ValidatorExpectedBlocks.WithLabelValues(labels...).Set(float64(v.NumExpectedBlocks))
		e.metrics.ValidatorExpectedChunks.WithLabelValues(labels...).Set(float64(v.NumExpectedChunks))
		e.metrics.ValidatorProducedBlocks.WithLabelValues(labels...).Set(float64(v.NumProducedBlocks))
		e.metrics.ValidatorProducedChunks.WithLabelValues(labels...).Set(float64(v.NumProducedChunks))

		e.metrics.ValidatorSlashed.WithLabelValues(labels...).Set(float64(isSlashed))
		e.metrics.ValidatorStake.WithLabelValues(labels...).Set(float64(GetStakeFromString(v.Stake)))

		t := GetStakeFromString(v.Stake)
		if seatPrice == 0 {
			seatPrice = t
		}
		if seatPrice > t {
			seatPrice = t
		}
	}

	e.metrics.SeatPriceDesc.Set(seatPrice)

	for _, v := range validators.NextValidators {
		e.metrics.NextValidatorStake.
			WithLabelValues(v.AccountId, v.PublicKey, labelEpochStartHeight, e.isTracked(v.AccountId)).
			Set(float64(GetStakeFromString(v.Stake)))
	}

	for _, v := range validators.CurrentProposals {
		e.metrics.CurrentProposals.
			WithLabelValues(v.AccountId, v.PublicKey, labelEpochStartHeight, e.isTracked(v.AccountId)).
			Set(float64(GetStakeFromString(v.Stake)))
	}

	for _, v := range validators.PrevEpochKickOut {
		reason, _ := json.Marshal(v.Reason)

		e.metrics.PrevEpochKickout.
			WithLabelValues(v.AccountId, string(reason), labelEpochStartHeight, e.isTracked(v.AccountId)).
			Set(1)
	}

	return nil
}

func (e *Exporter) isTracked(accountId string) string {
	for _, t := range e.config.TrackedAccounts {
		if accountId == t {
			return "1"
		}
	}
	return "0"
}
