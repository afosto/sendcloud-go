package sender

import (
	"github.com/itsrever/sendcloud-go"
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

// Get all sender addresses
func (c *Client) GetAddresses() ([]*sendcloud.Sender, error) {
	address := sendcloud.SenderResponseContainer{}
	err := sendcloud.Request("GET", "/api/v2/user/addresses/sender", nil, c.apiKey, c.apiSecret, &address)
	if err != nil {
		return nil, err
	}

	return address.GetResponse().([]*sendcloud.Sender), nil
}
