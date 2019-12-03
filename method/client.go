package method

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

func (c *Client) GetMethods() ([]*sendcloud.Method, error) {
	smr := sendcloud.MethodListResponseContainer{}
	_, err := sendcloud.Request("GET", "/api/v2/shipping_methods", nil, c.apiKey, c.apiSecret, &smr)
	if err != nil {
		return nil, err
	}
	return smr.GetResponse().([]*sendcloud.Method), nil
}

func (c *Client) GetMethod(id int64) (*sendcloud.Method, error) {
	mr := sendcloud.MethodResponseContainer{}
	_, err := sendcloud.Request("GET", "/api/v2/shipping_methods/"+strconv.Itoa(int(id)), nil, c.apiKey, c.apiSecret, &mr)
	if err != nil {
		return nil, err
	}
	return mr.GetResponse().(*sendcloud.Method), nil
}
