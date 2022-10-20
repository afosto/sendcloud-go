package integration

import (
	"github.com/itsrever/sendcloud-go"
	"strconv"
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

// List all integrations the the account
func (c *Client) GetIntegrations() ([]*sendcloud.Integration, error) {
	ilrc := sendcloud.IntegrationListResponseContainer{}
	err := sendcloud.Request("GET", "/api/v2/integrations", nil, c.apiKey, c.apiSecret, &ilrc)
	if err != nil {
		return nil, err
	}
	return ilrc.GetResponse().([]*sendcloud.Integration), nil

}

// Update an existing integration
func (c *Client) UpdateIntegration(params *sendcloud.IntegrationParams) (*sendcloud.Integration, error) {
	ilrc := sendcloud.IntegrationResponseContainer{}
	err := sendcloud.Request("PUT", "/api/v2/integrations/"+strconv.Itoa(int(params.ID)), params, c.apiKey, c.apiSecret, &ilrc)
	if err != nil {
		return nil, err
	}
	return ilrc.GetResponse().(*sendcloud.Integration), nil
}
