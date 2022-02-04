package api

const StatusMethod = "status"

type StatusResult struct {
	ChainId               string `json:"chain_id"`
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

func (c *Client) Status() (*StatusResult, error) {
	var result *StatusResult
	err := c.do(StatusMethod, nil, &result)

	return result, err
}
