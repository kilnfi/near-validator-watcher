package near

import "context"

type TxResponse struct {
	ReceiptsOutcome []struct {
		BlockHash string `json:"block_hash"`
		ID        string `json:"id"`
		Outcome   struct {
			ExecutorID string   `json:"executor_id"`
			GasBurnt   int64    `json:"gas_burnt"`
			Logs       []string `json:"logs"`
			Metadata   struct {
				GasProfile []struct {
					Cost         string `json:"cost"`
					CostCategory string `json:"cost_category"`
					GasUsed      string `json:"gas_used"`
				} `json:"gas_profile"`
				Version int `json:"version"`
			} `json:"metadata"`
			ReceiptIds []string `json:"receipt_ids"`
			Status     struct {
				SuccessValue string `json:"SuccessValue"`
			} `json:"status"`
			TokensBurnt string `json:"tokens_burnt"`
		} `json:"outcome"`
		Proof []struct {
			Direction string `json:"direction"`
			Hash      string `json:"hash"`
		} `json:"proof"`
	} `json:"receipts_outcome"`
	Status struct {
		SuccessValue string `json:"SuccessValue"`
	} `json:"status"`
	Transaction struct {
		Actions []struct {
			FunctionCall struct {
				Args       string `json:"args"`
				Deposit    string `json:"deposit"`
				Gas        int64  `json:"gas"`
				MethodName string `json:"method_name"`
			} `json:"FunctionCall"`
		} `json:"actions"`
		Hash       string `json:"hash"`
		Nonce      int64  `json:"nonce"`
		PublicKey  string `json:"public_key"`
		ReceiverID string `json:"receiver_id"`
		Signature  string `json:"signature"`
		SignerID   string `json:"signer_id"`
	} `json:"transaction"`
	TransactionOutcome struct {
		BlockHash string `json:"block_hash"`
		ID        string `json:"id"`
		Outcome   struct {
			ExecutorID string        `json:"executor_id"`
			GasBurnt   int64         `json:"gas_burnt"`
			Logs       []interface{} `json:"logs"`
			Metadata   struct {
				GasProfile interface{} `json:"gas_profile"`
				Version    int         `json:"version"`
			} `json:"metadata"`
			ReceiptIds []string `json:"receipt_ids"`
			Status     struct {
				SuccessReceiptID string `json:"SuccessReceiptId"`
			} `json:"status"`
			TokensBurnt string `json:"tokens_burnt"`
		} `json:"outcome"`
		Proof []struct {
			Direction string `json:"direction"`
			Hash      string `json:"hash"`
		} `json:"proof"`
	} `json:"transaction_outcome"`
}

func (c *Client) Tx(ctx context.Context, params interface{}) (TxResponse, error) {
	var resp TxResponse
	err := c.call(ctx, "tx", params, &resp)
	return resp, err
}
