package near

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// QueryRequest is used for RPC query requests.
type QueryRequest struct {
	RequestType  string      `json:"request_type"`
	Finality     string      `json:"finality,omitempty"`
	BlockID      interface{} `json:"block_id,omitempty"`
	AccountID    string      `json:"account_id,omitempty"`
	PrefixBase64 string      `json:"prefix_base64"`
	MethodName   string      `json:"method_name,omitempty"`
	ArgsBase64   string      `json:"args_base64"`
	PublicKey    string      `json:"public_key,omitempty"`
}

func NewQueryRequest(requestType, accountID, methodName string, opts ...QueryOption) (QueryRequest, error) {
	req := QueryRequest{
		RequestType: requestType,
		AccountID:   accountID,
		MethodName:  methodName,
	}

	for _, opt := range opts {
		if err := opt(&req); err != nil {
			return req, err
		}
	}
	return req, nil
}

// QueryResponse is a base type used for responses to query requests.
type QueryResponse struct {
	BlockHash   string `json:"block_hash"`
	BlockHeight int    `json:"block_height"`
	// TODO: this property is undocumented, but appears in some API responses. Is this the right place for it?
	Error string `json:"error"`
}

// QueryOption controls the behavior when calling CallFunction.
type QueryOption func(*QueryRequest) error

// QueryWithFinality specifies the finality to be used when calling the function.
func QueryWithFinality(finality string) QueryOption {
	return func(qr *QueryRequest) error {
		qr.Finality = finality
		return nil
	}
}

// QueryWithBlockHeight specifies the block height to call the function for.
func QueryWithBlockHeight(blockHeight int64) QueryOption {
	return func(qr *QueryRequest) error {
		qr.BlockID = blockHeight
		return nil
	}
}

// QueryWithBlockHash specifies the block hash to call the function for.
func QueryWithBlockHash(blockHash string) QueryOption {
	return func(qr *QueryRequest) error {
		qr.BlockID = blockHash
		return nil
	}
}

// QueryWithArgs specified the args to call the function with.
// Should be a JSON encodable object.
func QueryWithArgs(args interface{}) QueryOption {
	return func(qr *QueryRequest) error {
		if args == nil {
			args = make(map[string]interface{})
		}
		bytes, err := json.Marshal(args)
		if err != nil {
			return err
		}
		qr.ArgsBase64 = base64.StdEncoding.EncodeToString(bytes)
		return nil
	}
}

func (c *Client) Query(ctx context.Context, req QueryRequest, resp interface{}) error {
	if req.BlockID == nil && req.Finality == "" {
		return fmt.Errorf("missing block_id or finality")
	}
	if req.BlockID != nil && req.Finality != "" {
		return fmt.Errorf("you can't use both block_id and finality")
	}
	return c.call(ctx, "query", req, &resp)
}
