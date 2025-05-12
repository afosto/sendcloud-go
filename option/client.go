package option

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

func (c *Client) GetShippingOptions(params *sendcloud.OptionParams) (*sendcloud.OptionResponse, error) {
	response := sendcloud.OptionResponse{}
	err := sendcloud.Request("POST", "/api/v3/fetch-shipping-options", params, c.apiKey, c.apiSecret, &response)
	if err != nil {
		return nil, err
	}
	return response.GetResponse().(*sendcloud.OptionResponse), nil
}
