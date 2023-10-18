package watcher

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/fatih/color"
	"github.com/kilnfi/near-validator-watcher/pkg/metrics"
	"github.com/kilnfi/near-validator-watcher/pkg/near"
	"github.com/sirupsen/logrus"
)

type Watcher struct {
	config  *Config
	client  *near.Client
	metrics *metrics.Metrics

	isSynced atomic.Bool
}

func New(client *near.Client, metrics *metrics.Metrics, config *Config) *Watcher {
	return &Watcher{
		config:  config,
		client:  client,
		metrics: metrics,
	}
}

func (w *Watcher) IsSynced() bool {
	return w.isSynced.Load()
}

func (w *Watcher) Start(ctx context.Context) error {
	ticker := time.NewTicker(w.config.RefreshRate)

	for {
		retryOpts := []retry.Option{
			retry.Context(ctx),
			retry.Delay(1 * time.Second),
			retry.Attempts(3),
			retry.OnRetry(func(n uint, err error) {
				logrus.WithError(err).Error("failed to collect data, retrying...")
			}),
		}

		retry.Do(func() error {
			return w.collectData(ctx)
		}, retryOpts...)

		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			continue
		}
	}
}

func (w *Watcher) collectData(ctx context.Context) error {
	status, err := w.collectStatus(ctx)
	if err != nil {
		return err
	}
	validators, err := w.collectValidators(ctx)
	if err != nil {
		return err
	}
	_, err = w.collectProtocolConfig(ctx)
	if err != nil {
		return err
	}

	w.printStatusLine(status, validators)

	return nil
}

func (w *Watcher) collectProtocolConfig(ctx context.Context) (near.ProtocolConfigResponse, error) {
	logrus.Debug("collect protocol config")

	config, err := w.client.ProtocolConfig(ctx)
	if err != nil {
		return config, err
	}

	w.metrics.EpochLength.Set(float64(config.EpochLength))
	w.metrics.ProtocolVersion.Set(float64(config.ProtocolVersion))

	return config, nil
}

func (w *Watcher) collectStatus(ctx context.Context) (near.StatusResponse, error) {
	logrus.Debug("collect status")

	status, err := w.client.Status(ctx)
	if err != nil {
		return status, fmt.Errorf("failed to get status: %w", err)
	}

	w.isSynced.Store(!status.SyncInfo.Syncing)

	w.metrics.BlockNumber.Set(float64(status.SyncInfo.LatestBlockHeight))
	w.metrics.ChainID.WithLabelValues(status.ChainID).Set(metrics.StringToFloat64(status.ChainID))
	w.metrics.SyncingDesc.Set(metrics.BoolToFloat64(status.SyncInfo.Syncing))
	w.metrics.VersionBuild.WithLabelValues(status.Version.Version, status.Version.Build).Set(metrics.StringToFloat64(status.Version.Build))

	return status, nil
}

func (w *Watcher) collectValidators(ctx context.Context) (near.ValidatorsResponse, error) {
	logrus.Debug("collect validators")

	validators, err := w.client.Validators(ctx, "latest")
	if err != nil {
		return validators, err
	}

	// Reset labeled gauge vec
	w.metrics.ValidatorExpectedBlocks.Reset()
	w.metrics.ValidatorExpectedChunks.Reset()
	w.metrics.ValidatorProducedBlocks.Reset()
	w.metrics.ValidatorProducedChunks.Reset()
	w.metrics.ValidatorSlashed.Reset()
	w.metrics.ValidatorStake.Reset()
	w.metrics.NextValidatorStake.Reset()
	w.metrics.CurrentProposals.Reset()
	w.metrics.PrevEpochKickout.Reset()

	labelEpochStartHeight := strconv.FormatInt(validators.EpochStartHeight, 10)

	w.metrics.EpochStartHeight.Set(float64(validators.EpochStartHeight))

	var seatPrice float64

	// Sort validators by stake to be able to calculate their rank
	rankedValidator := validators.CurrentValidators
	sort.SliceStable(rankedValidator, func(i, j int) bool {
		return rankedValidator[i].Stake.GreaterThan(rankedValidator[j].Stake)
	})

	for i, v := range rankedValidator {
		labels := []string{v.AccountId, v.PublicKey, labelEpochStartHeight, w.isTracked(v.AccountId)}

		w.metrics.ValidatorRank.
			WithLabelValues(v.AccountId, v.PublicKey, labelEpochStartHeight, w.isTracked(v.AccountId)).
			Set(float64(i + 1))

		w.metrics.ValidatorExpectedBlocks.WithLabelValues(labels...).Set(float64(v.NumExpectedBlocks))
		w.metrics.ValidatorExpectedChunks.WithLabelValues(labels...).Set(float64(v.NumExpectedChunks))
		w.metrics.ValidatorProducedBlocks.WithLabelValues(labels...).Set(float64(v.NumProducedBlocks))
		w.metrics.ValidatorProducedChunks.WithLabelValues(labels...).Set(float64(v.NumProducedChunks))

		w.metrics.ValidatorSlashed.WithLabelValues(labels...).Set(metrics.BoolToFloat64(v.IsSlashed))
		w.metrics.ValidatorStake.WithLabelValues(labels...).Set(v.Stake.InexactFloat64())

		t := v.Stake.InexactFloat64()
		if seatPrice == 0 {
			seatPrice = t
		}
		if seatPrice > t {
			seatPrice = t
		}
	}

	w.metrics.SeatPrice.Set(seatPrice)

	for _, v := range validators.NextValidators {
		w.metrics.NextValidatorStake.
			WithLabelValues(v.AccountId, v.PublicKey, labelEpochStartHeight, w.isTracked(v.AccountId)).
			Set(v.Stake.InexactFloat64())
	}

	for _, v := range validators.CurrentProposals {
		w.metrics.CurrentProposals.
			WithLabelValues(v.AccountId, v.PublicKey, labelEpochStartHeight, w.isTracked(v.AccountId)).
			Set(v.Stake.InexactFloat64())
	}

	for _, v := range validators.PrevEpochKickOut {
		reason, _ := json.Marshal(v.Reason)

		w.metrics.PrevEpochKickout.
			WithLabelValues(v.AccountId, string(reason), labelEpochStartHeight, w.isTracked(v.AccountId)).
			Set(1)
	}

	return validators, nil
}

func (w *Watcher) isTracked(accountId string) string {
	for _, t := range w.config.TrackedAccounts {
		if accountId == t {
			return "1"
		}
	}
	return "0"
}

func (w *Watcher) printStatusLine(status near.StatusResponse, validators near.ValidatorsResponse) {
	validatorStatus := make([]string, 0)
	for _, account := range w.config.TrackedAccounts {
		for _, validator := range validators.CurrentValidators {
			if validator.AccountId != account {
				continue
			}

			var (
				status               = "✅"
				uptimeBlocks float64 = 100
				uptimeChunks float64 = 100
			)

			if validator.NumExpectedBlocks > 0 {
				uptimeBlocks = 100 * float64(validator.NumProducedBlocks) / float64(validator.NumExpectedBlocks)
			}
			if validator.NumExpectedChunks > 0 {
				uptimeChunks = 100 * float64(validator.NumProducedChunks) / float64(validator.NumExpectedChunks)
			}
			if uptimeBlocks < 90 || uptimeChunks < 90 {
				status = "❌"
			}

			validatorStatus = append(validatorStatus,
				fmt.Sprintf("%s %s (%s%%, %s%%)",
					status,
					prettyPrintAccountID(validator.AccountId),
					prettyPrintFloat(uptimeBlocks),
					prettyPrintFloat(uptimeChunks),
				),
			)
		}
	}

	fmt.Fprintln(
		w.config.Writer,
		color.YellowString(fmt.Sprintf("#%d (%d)", status.SyncInfo.LatestBlockHeight, validators.EpochHeight)),
		color.CyanString(fmt.Sprintf("%d validators", len(validators.CurrentValidators))),
		strings.Join(validatorStatus, " "),
	)
}
