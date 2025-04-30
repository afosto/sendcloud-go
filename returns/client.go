package returns

import (
	"github.com/afosto/sendcloud-go"
)

type Client struct {
	apiKey    string
	apiSecret string
}

func New(apiKey string, apiSecret string) *Client {
	return &Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

// Create a new returns label

func (c *Client) New(params *sendcloud.ReturnParams) (*sendcloud.ReturnResponse, error) {
	response := sendcloud.ReturnResponse{}
	err := sendcloud.Request("POST", "/api/v3/returns", params, c.apiKey, c.apiSecret, &response)

	if err != nil {
		return nil, err
	}
	return response.GetResponse().(*sendcloud.ReturnResponse), nil
}
