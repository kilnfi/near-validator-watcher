package collector

import (
	"encoding/json"
	"strconv"

	"github.com/kilnfi/near-exporter/api"
	"github.com/prometheus/client_golang/prometheus"
)

const MetricPrefix = "near_exporter_"

type Collector struct {
	client                      *api.Client
	trackedAccounts             []string
	epochBlockBroducedDesc      *prometheus.Desc
	epochBlockExpectedDesc      *prometheus.Desc
	seatPriceDesc               *prometheus.Desc
	currentStakeDesc            *prometheus.Desc
	epochStartHeightDesc        *prometheus.Desc
	blockNumberDesc             *prometheus.Desc
	syncingDesc                 *prometheus.Desc
	versionBuildDesc            *prometheus.Desc
	validatorStakeDesc          *prometheus.Desc
	validatorSlashedDesc        *prometheus.Desc
	validatorExpectedBlocksDesc *prometheus.Desc
	validatorProducedBlocksDesc *prometheus.Desc
	validatorExpectedChunksDesc *prometheus.Desc
	validatorProducedChunksDesc *prometheus.Desc
	nextValidatorStakeDesc      *prometheus.Desc
	prevEpochKickoutDesc        *prometheus.Desc
	currentProposalsDesc        *prometheus.Desc
	epochLength                 *prometheus.Desc
	chainId                     *prometheus.Desc
	protocolVersion             *prometheus.Desc
}

func NewCollector(client *api.Client) *Collector {
	return &Collector{
		client: client,
		epochBlockBroducedDesc: prometheus.NewDesc(
			MetricPrefix+"epoch_block_produced_number",
			"The number of block produced in epoch",
			nil,
			nil,
		),
		epochBlockExpectedDesc: prometheus.NewDesc(
			MetricPrefix+"epoch_block_expected_number",
			"The number of block expected in epoch",
			nil,
			nil,
		),
		seatPriceDesc: prometheus.NewDesc(
			MetricPrefix+"seat_price",
			"Validator seat price",
			nil,
			nil,
		),
		currentStakeDesc: prometheus.NewDesc(
			MetricPrefix+"current_stake",
			"Current stake of a given account id",
			nil,
			nil,
		),
		epochStartHeightDesc: prometheus.NewDesc(
			MetricPrefix+"epoch_start_height",
			"Near epoch start height",
			nil,
			nil,
		),
		blockNumberDesc: prometheus.NewDesc(
			MetricPrefix+"block_number",
			"The number of most recent block",
			nil,
			nil,
		),
		syncingDesc: prometheus.NewDesc(
			MetricPrefix+"sync_state",
			"Sync state",
			nil,
			nil,
		),
		versionBuildDesc: prometheus.NewDesc(
			MetricPrefix+"version_build",
			"The Near node version build",
			[]string{"version", "build"},
			nil,
		),
		validatorStakeDesc: prometheus.NewDesc(
			MetricPrefix+"validator_stake",
			"Current amount of validator stake",
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
			nil,
		),
		validatorSlashedDesc: prometheus.NewDesc(
			MetricPrefix+"validator_slashed",
			"Validators slashed",
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
			nil,
		),
		validatorExpectedBlocksDesc: prometheus.NewDesc(
			MetricPrefix+"validator_blocks_expected",
			"Current amount of validator expected blocks",
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
			nil,
		),
		validatorProducedBlocksDesc: prometheus.NewDesc(
			MetricPrefix+"validator_blocks_produced",
			"Current amount of validator produced blocks",
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
			nil,
		),
		validatorExpectedChunksDesc: prometheus.NewDesc(
			MetricPrefix+"validator_chunks_expected",
			"Current amount of validator expected chunks",
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
			nil,
		),
		validatorProducedChunksDesc: prometheus.NewDesc(
			MetricPrefix+"validator_chunks_produced",
			"Current amount of validator produced chunks",
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
			nil,
		),
		nextValidatorStakeDesc: prometheus.NewDesc(
			MetricPrefix+"next_validator_stake",
			"The next validators",
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
			nil,
		),
		currentProposalsDesc: prometheus.NewDesc(
			MetricPrefix+"current_proposals_stake",
			"Current proposals",
			[]string{"account_id", "public_key", "epoch_start_height", "tracked"},
			nil,
		),
		prevEpochKickoutDesc: prometheus.NewDesc(
			MetricPrefix+"prev_epoch_kickout",
			"Near previous epoch kicked out validators",
			[]string{"account_id", "reason", "epoch_start_height", "tracked"},
			nil,
		),
		epochLength: prometheus.NewDesc(
			MetricPrefix+"epoch_length",
			"Near epoch length as specified in the protocol",
			nil,
			nil,
		),
		chainId: prometheus.NewDesc(
			MetricPrefix+"chain_id",
			"Near chain id",
			[]string{"chain_id"},
			nil,
		),
		protocolVersion: prometheus.NewDesc(
			MetricPrefix+"protocol_version",
			"Current protocol version deployed to the blockchain",
			nil,
			nil,
		),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.seatPriceDesc
	ch <- c.epochStartHeightDesc
	ch <- c.blockNumberDesc
	ch <- c.syncingDesc
	ch <- c.versionBuildDesc
	ch <- c.validatorStakeDesc
	ch <- c.validatorSlashedDesc
	ch <- c.nextValidatorStakeDesc
	ch <- c.currentProposalsDesc
	ch <- c.prevEpochKickoutDesc
	ch <- c.validatorExpectedBlocksDesc
	ch <- c.validatorProducedBlocksDesc
	ch <- c.validatorExpectedChunksDesc
	ch <- c.validatorProducedChunksDesc
	ch <- c.epochLength
	ch <- c.chainId
	ch <- c.protocolVersion
}

func (c *Collector) WithTrackedAccounts(trackedAccounts []string) *Collector {
	c.trackedAccounts = trackedAccounts
	return c
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.collectStatus(ch)
	c.collectValidators(ch)
	c.collectProtocolConfig(ch)
}

func (c *Collector) isTracked(accountId string) string {
	for _, t := range c.trackedAccounts {
		if accountId == t {
			return "1"
		}
	}
	return "0"
}

// Collect info from Status method
func (c *Collector) collectStatus(ch chan<- prometheus.Metric) {
	status, err := c.client.Status()
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.blockNumberDesc, err)
		ch <- prometheus.NewInvalidMetric(c.syncingDesc, err)
		ch <- prometheus.NewInvalidMetric(c.versionBuildDesc, err)
		return
	}

	blockHeight := status.SyncInfo.LatestBlockHeight
	ch <- prometheus.MustNewConstMetric(c.blockNumberDesc, prometheus.GaugeValue, float64(blockHeight))

	isSyncing := 0
	if status.SyncInfo.Syncing {
		isSyncing = 1
	}
	ch <- prometheus.MustNewConstMetric(c.syncingDesc, prometheus.GaugeValue, float64(isSyncing))

	versionBuildInt := HashString(status.Version.Build)
	ch <- prometheus.MustNewConstMetric(c.versionBuildDesc, prometheus.GaugeValue, float64(versionBuildInt), status.Version.Version, status.Version.Build)
}

// Collect info from Validators method
func (c *Collector) collectValidators(ch chan<- prometheus.Metric) {
	validators, err := c.client.Validators()
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.seatPriceDesc, err)
		ch <- prometheus.NewInvalidMetric(c.epochStartHeightDesc, err)
		ch <- prometheus.NewInvalidMetric(c.syncingDesc, err)
		ch <- prometheus.NewInvalidMetric(c.validatorStakeDesc, err)
		ch <- prometheus.NewInvalidMetric(c.validatorSlashedDesc, err)
		ch <- prometheus.NewInvalidMetric(c.validatorExpectedBlocksDesc, err)
		ch <- prometheus.NewInvalidMetric(c.validatorExpectedChunksDesc, err)
		ch <- prometheus.NewInvalidMetric(c.validatorProducedBlocksDesc, err)
		ch <- prometheus.NewInvalidMetric(c.validatorProducedChunksDesc, err)
		ch <- prometheus.NewInvalidMetric(c.nextValidatorStakeDesc, err)
		ch <- prometheus.NewInvalidMetric(c.currentProposalsDesc, err)
		ch <- prometheus.NewInvalidMetric(c.prevEpochKickoutDesc, err)
		return
	}

	ch <- prometheus.MustNewConstMetric(c.epochStartHeightDesc, prometheus.GaugeValue, float64(validators.EpochStartHeight))

	var seatPrice float64
	for _, v := range validators.CurrentValidators {
		isTracked := c.isTracked(v.AccountId)

		ch <- prometheus.MustNewConstMetric(c.validatorExpectedChunksDesc, prometheus.GaugeValue, float64(v.NumExpectedChunks), v.AccountId, v.PublicKey, strconv.FormatInt(validators.EpochStartHeight, 10), isTracked)
		ch <- prometheus.MustNewConstMetric(c.validatorProducedChunksDesc, prometheus.GaugeValue, float64(v.NumProducedChunks), v.AccountId, v.PublicKey, strconv.FormatInt(validators.EpochStartHeight, 10), isTracked)
		ch <- prometheus.MustNewConstMetric(c.validatorExpectedBlocksDesc, prometheus.GaugeValue, float64(v.NumExpectedBlocks), v.AccountId, v.PublicKey, strconv.FormatInt(validators.EpochStartHeight, 10), isTracked)
		ch <- prometheus.MustNewConstMetric(c.validatorProducedBlocksDesc, prometheus.GaugeValue, float64(v.NumProducedBlocks), v.AccountId, v.PublicKey, strconv.FormatInt(validators.EpochStartHeight, 10), isTracked)
		ch <- prometheus.MustNewConstMetric(c.validatorStakeDesc, prometheus.GaugeValue, float64(GetStakeFromString(v.Stake)), v.AccountId, v.PublicKey, strconv.FormatInt(validators.EpochStartHeight, 10), isTracked)

		isSlashed := 0
		if v.IsSlashed {
			isSlashed = 1
		}
		ch <- prometheus.MustNewConstMetric(c.validatorSlashedDesc, prometheus.GaugeValue, float64(isSlashed), v.AccountId, v.PublicKey, strconv.FormatInt(validators.EpochStartHeight, 10), isTracked)

		t := GetStakeFromString(v.Stake)
		if seatPrice == 0 {
			seatPrice = t
		}
		if seatPrice > t {
			seatPrice = t
		}
	}

	ch <- prometheus.MustNewConstMetric(c.seatPriceDesc, prometheus.GaugeValue, seatPrice)

	for _, v := range validators.NextValidators {
		ch <- prometheus.MustNewConstMetric(c.nextValidatorStakeDesc, prometheus.GaugeValue, float64(GetStakeFromString(v.Stake)), v.AccountId, v.PublicKey, strconv.FormatInt(validators.EpochStartHeight, 10), c.isTracked(v.AccountId))
	}

	for _, v := range validators.CurrentProposals {
		ch <- prometheus.MustNewConstMetric(c.currentProposalsDesc, prometheus.GaugeValue, float64(GetStakeFromString(v.Stake)), v.AccountId, v.PublicKey, strconv.FormatInt(validators.EpochStartHeight, 10), c.isTracked(v.AccountId))
	}

	for _, v := range validators.PrevEpochKickOut {
		reason, _ := json.Marshal(v.Reason)
		ch <- prometheus.MustNewConstMetric(c.prevEpochKickoutDesc, prometheus.GaugeValue, 0, v.AccountId, string(reason), strconv.FormatInt(validators.EpochStartHeight, 10), c.isTracked(v.AccountId))
	}
}

// Collect info from ProtocolConfig method
func (c *Collector) collectProtocolConfig(ch chan<- prometheus.Metric) {
	config, err := c.client.ProtocolConfig()
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.epochLength, err)
		ch <- prometheus.NewInvalidMetric(c.chainId, err)
		ch <- prometheus.NewInvalidMetric(c.protocolVersion, err)
		return
	}

	ch <- prometheus.MustNewConstMetric(c.chainId, prometheus.GaugeValue, float64(HashString(config.ChainID)), config.ChainID)
	ch <- prometheus.MustNewConstMetric(c.epochLength, prometheus.GaugeValue, float64(config.EpochLength))
	ch <- prometheus.MustNewConstMetric(c.protocolVersion, prometheus.GaugeValue, float64(config.ProtocolVersion))
}
