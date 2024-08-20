package near

import (
	"context"

	"github.com/shopspring/decimal"
)

type ValidatorsResponse struct {
	CurrentValidators []struct {
		Validator
		IsSlashed               bool  `json:"is_slashed"`
		Shards                  []int `json:"shards"`
		NumProducedBlocks       int64 `json:"num_produced_blocks"`
		NumExpectedBlocks       int64 `json:"num_expected_blocks"`
		NumProducedChunks       int64 `json:"num_produced_chunks"`
		NumExpectedChunks       int64 `json:"num_expected_chunks"`
		NumProducedEndorsements int64 `json:"num_produced_endorsements"`
		NumExpectedEndorsements int64 `json:"num_expected_endorsements"`
	} `json:"current_validators"`
	NextValidators []struct {
		Validator
		Shards []int `json:"shards"`
	} `json:"next_validators"`
	CurrentProposals []struct {
		Validator
		StakeStructVersion string `json:"validator_stake_struct_version"`
	} `json:"current_proposals"`
	EpochStartHeight int64 `json:"epoch_start_height"`
	EpochHeight      int64 `json:"epoch_height"`
	PrevEpochKickOut []struct {
		AccountId string      `json:"account_id"`
		Reason    interface{} `json:"reason"`
	} `json:"prev_epoch_kickout"`
}

type Validator struct {
	AccountId string          `json:"account_id"`
	PublicKey string          `json:"public_key"`
	Stake     decimal.Decimal `json:"stake"`
}

func (c *Client) Validators(ctx context.Context, params interface{}) (ValidatorsResponse, error) {
	var resp ValidatorsResponse
	err := c.call(ctx, "validators", params, &resp)
	return resp, err
}
