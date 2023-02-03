package near

import "context"

type StatusResponse struct {
	ChainID               string `json:"chain_id"`
	LatestProtocolVersion int    `json:"latest_protocol_version"`
	ProtocolVersion       int    `json:"protocol_version"`
	RpcAddr               string `json:"rpc_addr"`
	SyncInfo              struct {
		EarliestBlockHash   string `json:"earliest_block_hash"`
		EarliestBlockHeight uint64 `json:"earliest_block_height"`
		EarliestBlockTime   string `json:"earliest_block_time"`
		LatestBlockHash     string `json:"latest_block_hash"`
		LatestBlockHeight   uint64 `json:"latest_block_height"`
		LatestStateRoot     string `json:"latest_state_root"`
		LatestBlockTime     string `json:"latest_block_time"`
		Syncing             bool   `json:"syncing"`
	} `json:"sync_info"`
	//Validators []string `json:"validators"`
	Version struct {
		Version string `json:"version"`
		Build   string `json:"build"`
	} `json:"version"`
}

func (c *Client) Status(ctx context.Context) (StatusResponse, error) {
	var resp StatusResponse
	err := c.call(ctx, "status", nil, &resp)
	return resp, err
}
