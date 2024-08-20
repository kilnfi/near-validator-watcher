package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Metrics struct {
	BlockNumber                   prometheus.Gauge
	ChainID                       *prometheus.GaugeVec
	CurrentProposals              *prometheus.GaugeVec
	EpochLength                   prometheus.Gauge
	EpochStartHeight              prometheus.Gauge
	NextValidatorStake            *prometheus.GaugeVec
	PrevEpochKickout              *prometheus.GaugeVec
	ProtocolVersion               prometheus.Gauge
	SeatPrice                     prometheus.Gauge
	SyncingDesc                   prometheus.Gauge
	ValidatorExpectedBlocks       *prometheus.GaugeVec
	ValidatorExpectedChunks       *prometheus.GaugeVec
	ValidatorExpectedEndorsements *prometheus.GaugeVec
	ValidatorProducedBlocks       *prometheus.GaugeVec
	ValidatorProducedChunks       *prometheus.GaugeVec
	ValidatorProducedEndorsements *prometheus.GaugeVec
	ValidatorSlashed              *prometheus.GaugeVec
	ValidatorStake                *prometheus.GaugeVec
	ValidatorRank                 *prometheus.GaugeVec
	VersionBuild                  *prometheus.GaugeVec
}

func New(namespace string) *Metrics {
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
			[]string{"account_id", "reason", "epoch_start_height", "tracked"},
		),
		ProtocolVersion: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "protocol_version",
			Help:      "Current protocol version deployed to the blockchain",
		}),
		SeatPrice: prometheus.NewGauge(prometheus.GaugeOpts{
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
		ValidatorExpectedEndorsements: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "validator_endorsements_expected",
			Help:      "Current amount of validator expected endorsements"},
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
		ValidatorProducedEndorsements: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "validator_endorsements_produced",
			Help:      "Current amount of validator produced endorsements"},
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
		ValidatorRank: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "validator_rank",
			Help:      "Current rank of validator based on stake"},
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

func (m *Metrics) Register(reg prometheus.Registerer) {
	reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	reg.MustRegister(collectors.NewGoCollector())

	reg.MustRegister(m.BlockNumber)
	reg.MustRegister(m.ChainID)
	reg.MustRegister(m.CurrentProposals)
	reg.MustRegister(m.EpochLength)
	reg.MustRegister(m.EpochStartHeight)
	reg.MustRegister(m.NextValidatorStake)
	reg.MustRegister(m.PrevEpochKickout)
	reg.MustRegister(m.ProtocolVersion)
	reg.MustRegister(m.SeatPrice)
	reg.MustRegister(m.SyncingDesc)
	reg.MustRegister(m.ValidatorExpectedBlocks)
	reg.MustRegister(m.ValidatorExpectedChunks)
	reg.MustRegister(m.ValidatorExpectedEndorsements)
	reg.MustRegister(m.ValidatorProducedBlocks)
	reg.MustRegister(m.ValidatorProducedChunks)
	reg.MustRegister(m.ValidatorProducedEndorsements)
	reg.MustRegister(m.ValidatorSlashed)
	reg.MustRegister(m.ValidatorStake)
	reg.MustRegister(m.ValidatorRank)
	reg.MustRegister(m.VersionBuild)
}
