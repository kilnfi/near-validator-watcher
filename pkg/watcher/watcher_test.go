package watcher

import (
	"context"
	"testing"

	"github.com/kilnfi/near-validator-watcher/pkg/metrics"
	"github.com/kilnfi/near-validator-watcher/pkg/near"
	"github.com/kilnfi/near-validator-watcher/pkg/near/testutils"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWatcher(t *testing.T) {
	var (
		ctx     = context.Background()
		resp    = testutils.ExpectedResponse{}
		server  = testutils.NewServer(&resp)
		client  = near.NewClient(server.URL)
		metrics = metrics.New("near_validator_watcher")
		watcher = New(client, metrics, &Config{
			TrackedAccounts: []string{"kiln.pool.f863973.m0"},
		})
	)

	defer server.Close()

	t.Run("Collect Protocol Config", func(t *testing.T) {
		resp.ExpectResponse(200, `
			{
				"id": "dontcare",
				"jsonrpc": "2.0",
				"result": {
					"epoch_length": 43200,
					"protocol_version": 63
				}
		}`)

		_, err := watcher.collectProtocolConfig(ctx)
		require.NoError(t, err)

		assert.Equal(t, float64(43200), testutil.ToFloat64(metrics.EpochLength))
		assert.Equal(t, float64(63), testutil.ToFloat64(metrics.ProtocolVersion))
	})

	t.Run("Collect Status", func(t *testing.T) {
		resp.ExpectResponse(200, `
			{
				"id": "dontcare",
				"jsonrpc": "2.0",
				"result": {
						"chain_id": "testnet",
						"latest_protocol_version": 63,
						"node_key": null,
						"node_public_key": "ed25519:38wj5FVhH5AEUgEPHzcTwc5X7UgwiVQVNqRrrdGu7Bg1",
						"protocol_version": 63,
						"rpc_addr": "127.0.0.1:4040",
						"sync_info": {
								"earliest_block_hash": "4GSGamHhRFXuZxruqUKMz87PxCuwMtN6Dbtmem6QsbJT",
								"earliest_block_height": 142045749,
								"earliest_block_time": "2023-10-17T00:36:29.465562502Z",
								"epoch_id": "CpAZXf3iXAVys8dgB49CWDsi3fzUNWny6w4J78VaADPc",
								"epoch_start_height": 142256359,
								"latest_block_hash": "4N5XL5zR9527vDnoCUEVcBEQMTqRepmfmgwQGrkyrjEM",
								"latest_block_height": 142259035,
								"latest_block_time": "2023-10-18T13:35:36.417320147Z",
								"latest_state_root": "B1do4VKdgL8CToXMd5E4fc56drSEaWqwXRyeSS9bUKY5",
								"syncing": false
						},
						"uptime_sec": 2236367,
						"validator_account_id": null,
						"validator_public_key": null,
						"validators": [
								{
										"account_id": "test1.poolv1.near",
										"is_slashed": false
								},
								{
										"account_id": "test2.poolv1.near",
										"is_slashed": false
								}
						],
						"version": {
								"build": "1.36.0-rc.1",
								"rustc_version": "1.71.0",
								"version": "1.36.0-rc.1"
						}
				}
		}`)

		_, err := watcher.collectStatus(ctx)
		require.NoError(t, err)

		assert.Equal(t, true, watcher.IsSynced())

		assert.Equal(t, float64(142259035), testutil.ToFloat64(metrics.BlockNumber))
		assert.NotEqual(t, float64(0), testutil.ToFloat64(metrics.ChainID.WithLabelValues("testnet")))
	})

	t.Run("Collect Validators", func(t *testing.T) {
		resp.ExpectResponse(200, `
			{
				"id": "dontcare",
				"jsonrpc": "2.0",
				"result": {
						"current_fishermen": [],
						"current_proposals": [
								{
										"account_id": "aurora.pool.f863973.m0",
										"public_key": "ed25519:9c7mczZpNzJz98V1sDeGybfD4gMybP4JKHotH8RrrHTm",
										"stake": "17289725999534933301348791102920",
										"validator_stake_struct_version": "V1"
								},
								{
										"account_id": "chorus-one.pool.f863973.m0",
										"public_key": "ed25519:6LFwyEEsqhuDxorWfsKcPPs324zLWTaoqk4o6RDXN7Qc",
										"stake": "2478635435253716013709224719804",
										"validator_stake_struct_version": "V1"
								},
								{
										"account_id": "kiln.pool.f863973.m0",
										"public_key": "ed25519:Bq8fe1eUgDRexX2CYDMhMMQBiN13j8vTAVFyTNhEfh1W",
										"stake": "6739683523640860052948489127116",
										"validator_stake_struct_version": "V1"
								},
								{
										"account_id": "ni.pool.f863973.m0",
										"public_key": "ed25519:GfCfFkLk2twbAWdsS3tr7C2eaiHN3znSfbshS5e8NqBS",
										"stake": "2842038361242433260299812278899",
										"validator_stake_struct_version": "V1"
								},
								{
										"account_id": "stakely_v2.pool.f863973.m0",
										"public_key": "ed25519:7BanKZKGvFjK5Yy83gfJ71vPhqRwsDDyVHrV2FMJCUWr",
										"stake": "5050537292623241655006223011598",
										"validator_stake_struct_version": "V1"
								}
						],
						"current_validators": [
								{
										"account_id": "node1",
										"is_slashed": false,
										"num_expected_blocks": 640,
										"num_expected_chunks": 2527,
										"num_expected_chunks_per_shard": [
												2527
										],
										"num_produced_blocks": 640,
										"num_produced_chunks": 2527,
										"num_produced_chunks_per_shard": [
												2527
										],
										"public_key": "ed25519:6DSjZ8mvsRZDvFqFxo8tCKePG96omXW7eVYVSySmDk8e",
										"shards": [
												0
										],
										"stake": "47604850844179868792335153947376"
								},
								{
										"account_id": "node2",
										"is_slashed": false,
										"num_expected_blocks": 662,
										"num_expected_chunks": 2593,
										"num_expected_chunks_per_shard": [
												2593
										],
										"num_produced_blocks": 662,
										"num_produced_chunks": 2593,
										"num_produced_chunks_per_shard": [
												2593
										],
										"public_key": "ed25519:GkDv7nSMS3xcqA45cpMvFmfV1o4fRF6zYo1JRR6mNqg5",
										"shards": [
												1
										],
										"stake": "47595986932991225292162473028453"
								},
								{
										"account_id": "node3",
										"is_slashed": false,
										"num_expected_blocks": 650,
										"num_expected_chunks": 2593,
										"num_expected_chunks_per_shard": [
												2593
										],
										"num_produced_blocks": 650,
										"num_produced_chunks": 2593,
										"num_produced_chunks_per_shard": [
												2593
										],
										"public_key": "ed25519:ydgzeXHJ5Xyt7M1gXLxqLBW1Ejx6scNV5Nx2pxFM8su",
										"shards": [
												2
										],
										"stake": "47558522874659976267803051885615"
								},
								{
										"account_id": "kiln.pool.f863973.m0",
										"is_slashed": false,
										"num_expected_blocks": 92,
										"num_expected_chunks": 392,
										"num_expected_chunks_per_shard": [
												392
										],
										"num_produced_blocks": 91,
										"num_produced_chunks": 391,
										"num_produced_chunks_per_shard": [
												391
										],
										"public_key": "ed25519:Bq8fe1eUgDRexX2CYDMhMMQBiN13j8vTAVFyTNhEfh1W",
										"shards": [
												0
										],
										"stake": "6736422258840329637507414885764"
								},
								{
										"account_id": "stakely_v2.pool.f863973.m0",
										"is_slashed": false,
										"num_expected_blocks": 58,
										"num_expected_chunks": 294,
										"num_expected_chunks_per_shard": [
												294
										],
										"num_produced_blocks": 58,
										"num_produced_chunks": 294,
										"num_produced_chunks_per_shard": [
												294
										],
										"public_key": "ed25519:7BanKZKGvFjK5Yy83gfJ71vPhqRwsDDyVHrV2FMJCUWr",
										"shards": [
												0
										],
										"stake": "5048850744401447176504014136424"
								}
						],
						"epoch_height": 2312,
						"epoch_start_height": 142256359,
						"next_fishermen": [],
						"next_validators": [
								{
										"account_id": "node1",
										"public_key": "ed25519:6DSjZ8mvsRZDvFqFxo8tCKePG96omXW7eVYVSySmDk8e",
										"shards": [
												0
										],
										"stake": "47620752994974369049510359713578"
								},
								{
										"account_id": "node2",
										"public_key": "ed25519:GkDv7nSMS3xcqA45cpMvFmfV1o4fRF6zYo1JRR6mNqg5",
										"shards": [
												1
										],
										"stake": "47611886122842673501102298433476"
								},
								{
										"account_id": "node3",
										"public_key": "ed25519:ydgzeXHJ5Xyt7M1gXLxqLBW1Ejx6scNV5Nx2pxFM8su",
										"shards": [
												2
										],
										"stake": "47574409549839192701171919246504"
								},
								{
										"account_id": "kiln.pool.f863973.m0",
										"public_key": "ed25519:Bq8fe1eUgDRexX2CYDMhMMQBiN13j8vTAVFyTNhEfh1W",
										"shards": [
												0
										],
										"stake": "6739683523438532957632289127116"
								},
								{
										"account_id": "stakely_v2.pool.f863973.m0",
										"public_key": "ed25519:7BanKZKGvFjK5Yy83gfJ71vPhqRwsDDyVHrV2FMJCUWr",
										"shards": [
												0
										],
										"stake": "5050537292472590301154723011598"
								}
						],
						"prev_epoch_kickout": [
								{
										"account_id": "example1.pool.f863973.m0",
										"reason": {
												"NotEnoughBlocks": {
														"expected": 16,
														"produced": 0
												}
										}
								},
								{
										"account_id": "example2.pool.f863973.m0",
										"reason": {
												"NotEnoughBlocks": {
														"expected": 63,
														"produced": 0
												}
										}
								}
						]
				}
		}`)

		_, err := watcher.collectValidators(ctx)
		require.NoError(t, err)

		assert.Equal(t, float64(142256359), testutil.ToFloat64(metrics.EpochStartHeight))
		assert.Equal(t, float64(5048850744401447176504014136424), testutil.ToFloat64(metrics.SeatPrice))

		// ValidatorRank
		assert.Equal(t, 5, testutil.CollectAndCount(metrics.ValidatorRank))
		assert.Equal(t, float64(2), testutil.ToFloat64(metrics.ValidatorRank.WithLabelValues(
			"node2",
			`ed25519:GkDv7nSMS3xcqA45cpMvFmfV1o4fRF6zYo1JRR6mNqg5`,
			"142256359",
			"0",
		)))
		assert.Equal(t, float64(4), testutil.ToFloat64(metrics.ValidatorRank.WithLabelValues(
			"kiln.pool.f863973.m0",
			`ed25519:Bq8fe1eUgDRexX2CYDMhMMQBiN13j8vTAVFyTNhEfh1W`,
			"142256359",
			"1",
		)))

		// ExpectedBlocks
		assert.Equal(t, 5, testutil.CollectAndCount(metrics.ValidatorExpectedBlocks))
		assert.Equal(t, float64(662), testutil.ToFloat64(metrics.ValidatorExpectedBlocks.WithLabelValues(
			"node2",
			`ed25519:GkDv7nSMS3xcqA45cpMvFmfV1o4fRF6zYo1JRR6mNqg5`,
			"142256359",
			"0",
		)))
		assert.Equal(t, float64(92), testutil.ToFloat64(metrics.ValidatorExpectedBlocks.WithLabelValues(
			"kiln.pool.f863973.m0",
			`ed25519:Bq8fe1eUgDRexX2CYDMhMMQBiN13j8vTAVFyTNhEfh1W`,
			"142256359",
			"1",
		)))

		// ProducedBlocks
		assert.Equal(t, 5, testutil.CollectAndCount(metrics.ValidatorProducedBlocks))
		assert.Equal(t, float64(662), testutil.ToFloat64(metrics.ValidatorProducedBlocks.WithLabelValues(
			"node2",
			`ed25519:GkDv7nSMS3xcqA45cpMvFmfV1o4fRF6zYo1JRR6mNqg5`,
			"142256359",
			"0",
		)))
		assert.Equal(t, float64(91), testutil.ToFloat64(metrics.ValidatorProducedBlocks.WithLabelValues(
			"kiln.pool.f863973.m0",
			`ed25519:Bq8fe1eUgDRexX2CYDMhMMQBiN13j8vTAVFyTNhEfh1W`,
			"142256359",
			"1",
		)))

		// ExpectedChunks
		assert.Equal(t, 5, testutil.CollectAndCount(metrics.ValidatorExpectedChunks))
		assert.Equal(t, float64(2593), testutil.ToFloat64(metrics.ValidatorExpectedChunks.WithLabelValues(
			"node2",
			`ed25519:GkDv7nSMS3xcqA45cpMvFmfV1o4fRF6zYo1JRR6mNqg5`,
			"142256359",
			"0",
		)))
		assert.Equal(t, float64(392), testutil.ToFloat64(metrics.ValidatorExpectedChunks.WithLabelValues(
			"kiln.pool.f863973.m0",
			`ed25519:Bq8fe1eUgDRexX2CYDMhMMQBiN13j8vTAVFyTNhEfh1W`,
			"142256359",
			"1",
		)))

		// ProducedChunks
		assert.Equal(t, 5, testutil.CollectAndCount(metrics.ValidatorProducedChunks))
		assert.Equal(t, float64(2593), testutil.ToFloat64(metrics.ValidatorProducedChunks.WithLabelValues(
			"node2",
			`ed25519:GkDv7nSMS3xcqA45cpMvFmfV1o4fRF6zYo1JRR6mNqg5`,
			"142256359",
			"0",
		)))
		assert.Equal(t, float64(391), testutil.ToFloat64(metrics.ValidatorProducedChunks.WithLabelValues(
			"kiln.pool.f863973.m0",
			`ed25519:Bq8fe1eUgDRexX2CYDMhMMQBiN13j8vTAVFyTNhEfh1W`,
			"142256359",
			"1",
		)))

		// Slashed
		assert.Equal(t, 5, testutil.CollectAndCount(metrics.ValidatorSlashed))
		assert.Equal(t, float64(0), testutil.ToFloat64(metrics.ValidatorSlashed.WithLabelValues(
			"node2",
			`ed25519:GkDv7nSMS3xcqA45cpMvFmfV1o4fRF6zYo1JRR6mNqg5`,
			"142256359",
			"0",
		)))
		assert.Equal(t, float64(0), testutil.ToFloat64(metrics.ValidatorSlashed.WithLabelValues(
			"kiln.pool.f863973.m0",
			`ed25519:Bq8fe1eUgDRexX2CYDMhMMQBiN13j8vTAVFyTNhEfh1W`,
			"142256359",
			"1",
		)))

		// Stake
		assert.Equal(t, 5, testutil.CollectAndCount(metrics.ValidatorStake))
		assert.Equal(t, float64(47595986932991225292162473028453), testutil.ToFloat64(metrics.ValidatorStake.WithLabelValues(
			"node2",
			`ed25519:GkDv7nSMS3xcqA45cpMvFmfV1o4fRF6zYo1JRR6mNqg5`,
			"142256359",
			"0",
		)))
		assert.Equal(t, float64(6736422258840329637507414885764), testutil.ToFloat64(metrics.ValidatorStake.WithLabelValues(
			"kiln.pool.f863973.m0",
			`ed25519:Bq8fe1eUgDRexX2CYDMhMMQBiN13j8vTAVFyTNhEfh1W`,
			"142256359",
			"1",
		)))

		// assert.Equal(t, float64(43200), testutil.ToFloat64(metrics.NextValidatorStake))
		// assert.Equal(t, float64(43200), testutil.ToFloat64(metrics.CurrentProposals))

		assert.Equal(t, 2, testutil.CollectAndCount(metrics.PrevEpochKickout))
		assert.Equal(t, float64(1), testutil.ToFloat64(metrics.PrevEpochKickout.WithLabelValues(
			"example1.pool.f863973.m0",
			`{"NotEnoughBlocks":{"expected":16,"produced":0}}`,
			"142256359",
			"0",
		)))
		assert.Equal(t, float64(1), testutil.ToFloat64(metrics.PrevEpochKickout.WithLabelValues(
			"example2.pool.f863973.m0",
			`{"NotEnoughBlocks":{"expected":63,"produced":0}}`,
			"142256359",
			"0",
		)))
	})
}
