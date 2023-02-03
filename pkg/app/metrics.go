package app

import "github.com/prometheus/client_golang/prometheus"

const namespace = "near_exporter"

type Metrics struct {
	BlockNumber             prometheus.Gauge
	ChainID                 *prometheus.GaugeVec
	CurrentProposals        *prometheus.GaugeVec
	EpochLength             prometheus.Gauge
	EpochStartHeight        prometheus.Gauge
	NextValidatorStake      *prometheus.GaugeVec
	PrevEpochKickout        *prometheus.GaugeVec
	ProtocolVersion         prometheus.Gauge
	SeatPriceDesc           prometheus.Gauge
	SyncingDesc             prometheus.Gauge
	ValidatorExpectedBlocks *prometheus.GaugeVec
	ValidatorExpectedChunks *prometheus.GaugeVec
	ValidatorProducedBlocks *prometheus.GaugeVec
	ValidatorProducedChunks *prometheus.GaugeVec
	ValidatorSlashed        *prometheus.GaugeVec
	ValidatorStake          *prometheus.GaugeVec
	VersionBuild            *prometheus.GaugeVec
}

func NewMetrics() *Metrics {
	return &Metrics{
		BlockNumber: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "block_number",
			Help:      "The number of most recent block",
		}),
		ChainID: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "chain_id",
			Help:      "Near chain id"},
			[]string{"chain_id"},
		),
		CurrentProposals: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "current_proposals_stake",
			Help:      "Current proposals"},
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
		),
		EpochLength: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "epoch_length",
			Help:      "Near epoch length as specified in the protocol",
		}),
		EpochStartHeight: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "epoch_start_height",
			Help:      "Near epoch start height",
		}),
		NextValidatorStake: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "next_validator_stake",
			Help:      "The next validators"},
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
		),
		PrevEpochKickout: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "prev_epoch_kickout",
			Help:      "Near previous epoch kicked out validators"},
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
		),
		ProtocolVersion: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "protocol_version",
			Help:      "Current protocol version deployed to the blockchain",
		}),
		SeatPriceDesc: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "seat_price",
			Help:      "Validator seat price",
		}),
		SyncingDesc: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "sync_state",
			Help:      "Sync state",
		}),
		ValidatorExpectedBlocks: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "validator_blocks_expected",
			Help:      "Current amount of validator expected blocks"},
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
		),
		ValidatorExpectedChunks: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "validator_chunks_expected",
			Help:      "Current amount of validator expected chunks"},
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
		),
		ValidatorProducedBlocks: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "validator_blocks_produced",
			Help:      "Current amount of validator produced blocks"},
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
		),
		ValidatorProducedChunks: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "validator_chunks_produced",
			Help:      "Current amount of validator produced chunks"},
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
		),
		ValidatorSlashed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "validator_slashed",
			Help:      "Validators slashed"},
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
		),
		ValidatorStake: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "validator_stake",
			Help:      "Current amount of validator stake"},
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
		),
		VersionBuild: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "version_build",
			Help:      "The Near node version build"},
			[]string{"version", "build"},
		),
	}
}

func (m *Metrics) RegisterMetrics(reg prometheus.Registerer) error {
	if err := reg.Register(m.BlockNumber); err != nil {
		return err
	}
	if err := reg.Register(m.ChainID); err != nil {
		return err
	}
	if err := reg.Register(m.CurrentProposals); err != nil {
		return err
	}
	if err := reg.Register(m.EpochLength); err != nil {
		return err
	}
	if err := reg.Register(m.EpochStartHeight); err != nil {
		return err
	}
	if err := reg.Register(m.NextValidatorStake); err != nil {
		return err
	}
	if err := reg.Register(m.PrevEpochKickout); err != nil {
		return err
	}
	if err := reg.Register(m.ProtocolVersion); err != nil {
		return err
	}
	if err := reg.Register(m.SeatPriceDesc); err != nil {
		return err
	}
	if err := reg.Register(m.SyncingDesc); err != nil {
		return err
	}
	if err := reg.Register(m.ValidatorExpectedBlocks); err != nil {
		return err
	}
	if err := reg.Register(m.ValidatorExpectedChunks); err != nil {
		return err
	}
	if err := reg.Register(m.ValidatorProducedBlocks); err != nil {
		return err
	}
	if err := reg.Register(m.ValidatorProducedChunks); err != nil {
		return err
	}
	if err := reg.Register(m.ValidatorSlashed); err != nil {
		return err
	}
	if err := reg.Register(m.ValidatorStake); err != nil {
		return err
	}
	if err := reg.Register(m.VersionBuild); err != nil {
		return err
	}
	return nil
}
