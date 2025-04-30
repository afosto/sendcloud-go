package products

import (
	"github.com/afosto/sendcloud-go"
	"net/url"
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

// get shipping products
func (c *Client) GetShippingProducts(fromCountry string) ([]*sendcloud.Product, error) {
	smr := sendcloud.ProductListResponseContainer{}
	values := url.Values{}
	values.Set("from_country", fromCountry)
	reqUrl := "/api/v2/shipping-products?" + values.Encode()
	err := sendcloud.Request("GET", reqUrl, nil, c.apiKey, c.apiSecret, &smr)
	if err != nil {
		return nil, err
	}
	return smr.GetResponse().([]*sendcloud.Product), nil
}

// get return shipping products
func (c *Client) GetReturnShippingProducts(fromCountry string) ([]*sendcloud.Product, error) {
	smr := sendcloud.ProductListResponseContainer{}
	values := url.Values{}
	values.Set("from_country", fromCountry)
	values.Set("returns", "true")
	reqUrl := "/api/v2/shipping-products?" + values.Encode()
	err := sendcloud.Request("GET", reqUrl, nil, c.apiKey, c.apiSecret, &smr)
	if err != nil {
		return nil, err
	}
	return smr.GetResponse().([]*sendcloud.Product), nil
}
