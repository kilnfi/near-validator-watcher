package near

import "context"

type BlockRequest struct {
	Finality string      `json:"finality,omitempty"`
	BlockID  interface{} `json:"block_id,omitempty"`
}

type BlockResponse struct {
	Author string `json:"author"`
	Header struct {
		Height                int           `json:"height"`
		EpochID               string        `json:"epoch_id"`
		NextEpochID           string        `json:"next_epoch_id"`
		Hash                  string        `json:"hash"`
		PrevHash              string        `json:"prev_hash"`
		PrevStateRoot         string        `json:"prev_state_root"`
		ChunkReceiptsRoot     string        `json:"chunk_receipts_root"`
		ChunkHeadersRoot      string        `json:"chunk_headers_root"`
		ChunkTxRoot           string        `json:"chunk_tx_root"`
		OutcomeRoot           string        `json:"outcome_root"`
		ChunksIncluded        int           `json:"chunks_included"`
		ChallengesRoot        string        `json:"challenges_root"`
		Timestamp             int64         `json:"timestamp"`
		TimestampNanosec      string        `json:"timestamp_nanosec"`
		RandomValue           string        `json:"random_value"`
		ValidatorProposals    []interface{} `json:"validator_proposals"`
		ChunkMask             []bool        `json:"chunk_mask"`
		GasPrice              string        `json:"gas_price"`
		RentPaid              string        `json:"rent_paid"`
		ValidatorReward       string        `json:"validator_reward"`
		TotalSupply           string        `json:"total_supply"`
		ChallengesResult      []interface{} `json:"challenges_result"`
		LastFinalBlock        string        `json:"last_final_block"`
		LastDsFinalBlock      string        `json:"last_ds_final_block"`
		NextBpHash            string        `json:"next_bp_hash"`
		BlockMerkleRoot       string        `json:"block_merkle_root"`
		Approvals             []interface{} `json:"approvals"`
		Signature             string        `json:"signature"`
		LatestProtocolVersion int           `json:"latest_protocol_version"`
	} `json:"header"`
	Chunks []struct {
		ChunkHash            string        `json:"chunk_hash"`
		PrevBlockHash        string        `json:"prev_block_hash"`
		OutcomeRoot          string        `json:"outcome_root"`
		PrevStateRoot        string        `json:"prev_state_root"`
		EncodedMerkleRoot    string        `json:"encoded_merkle_root"`
		EncodedLength        int           `json:"encoded_length"`
		HeightCreated        int           `json:"height_created"`
		HeightIncluded       int           `json:"height_included"`
		ShardID              int           `json:"shard_id"`
		GasUsed              int           `json:"gas_used"`
		GasLimit             int64         `json:"gas_limit"`
		RentPaid             string        `json:"rent_paid"`
		ValidatorReward      string        `json:"validator_reward"`
		BalanceBurnt         string        `json:"balance_burnt"`
		OutgoingReceiptsRoot string        `json:"outgoing_receipts_root"`
		TxRoot               string        `json:"tx_root"`
		ValidatorProposals   []interface{} `json:"validator_proposals"`
		Signature            string        `json:"signature"`
	} `json:"chunks"`
}

func (c *Client) Block(ctx context.Context, blockID uint64) (BlockResponse, error) {
	var req BlockRequest
	var resp BlockResponse

	req.BlockID = blockID
	err := c.call(ctx, "block", req, &resp)
	return resp, err
}
