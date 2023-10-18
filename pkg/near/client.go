package near

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	Endpoint   string
}

type Payload struct {
	JsonRPC string      `json:"jsonrpc"`
	Id      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type Response struct {
	JsonRPC string          `json:"jsonrpc"`
	Id      string          `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   struct {
		Name    string `json:"name"`
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
		Cause   struct {
			Info interface{} `json:"info"`
			Name string      `json:"name"`
		} `json:"cause"`
	} `json:"error"`
}

type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func NewClient(endpoint string, options ...Option) *Client {
	client := &Client{
		Endpoint:   endpoint,
		httpClient: &http.Client{},
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func (c *Client) Request(ctx context.Context, method string, params interface{}) (*Response, error) {
	payload, err := json.Marshal(map[string]string{
		"query": method,
	})
	if err != nil {
		log.Println(err)
	}

	if params != "" {
		p := Payload{
			JsonRPC: "2.0",
			Id:      "near_exporter",
			Method:  method,
			Params:  params,
		}

		payload, err = json.Marshal(p)
		if err != nil {
			log.Println(err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.Endpoint, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	r, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var resp *Response
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) call(ctx context.Context, method string, params interface{}, result interface{}) error {
	resp, err := c.Request(ctx, method, params)
	if err != nil {
		return err
	}

	if resp.Error.Name != "" {
		return fmt.Errorf(
			"jsonrpc error(%d): %s %s",
			resp.Error.Code,
			resp.Error.Name,
			resp.Error.Message,
		)
	}

	err = json.Unmarshal(resp.Result, &result)
	if err != nil {
		return err
	}

	return nil
}
