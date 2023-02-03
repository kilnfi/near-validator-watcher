package near

import (
	"context"
)

// CallFunctionResponse holds information about the result of a function call.
type CallFunctionResponse struct {
	QueryResponse
	Logs   []string `json:"logs"`
	Result []byte   `json:"result"`
}

func (c *Client) CallFunction(
	ctx context.Context,
	accountID string,
	methodName string,
	opts ...QueryOption,
) (CallFunctionResponse, error) {
	var resp CallFunctionResponse
	req, err := NewQueryRequest("call_function", accountID, methodName, opts...)
	if err != nil {
		return resp, err
	}
	err = c.call(ctx, "query", req, &resp)
	return resp, err
}
