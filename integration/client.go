package integration

import (
	"github.com/afosto/sendcloud-go"
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

func (c *Client) GetIntegrations() ([]*sendcloud.Integration, error) {
	ilrc := sendcloud.IntegrationListResponseContainer{}
	_, err := sendcloud.Request("GET", "/api/v2/integrations", nil, c.apiKey, c.apiSecret, &ilrc)
	if err != nil {
		return nil, err
	}
	return ilrc.GetResponse().([]*sendcloud.Integration), nil

}

func (c *Client) UpdateIntegration(params *sendcloud.IntegrationParams) (*sendcloud.Integration, error) {
	ilrc := sendcloud.IntegrationResponseContainer{}
	_, err := sendcloud.Request("PUT", "/api/v2/integrations/"+strconv.Itoa(int(params.ID)), params, c.apiKey, c.apiSecret, &ilrc)
	if err != nil {
		return nil, err
	}
	return ilrc.GetResponse().(*sendcloud.Integration), nil
}
