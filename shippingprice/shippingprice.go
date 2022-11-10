package shippingprice

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/afosto/sendcloud-go"
)

type WeightUnit string

const (
	Gram     WeightUnit = "gram"
	Kilogram WeightUnit = "kilogram"
)

type ShippingPriceParams struct {
	ShippingMethodId int

	FromCountry string  // FromCountry is an Alpha2 country code, e.g. SE
	ToCountry   *string // ToCountry is an Alpha2 country code, e.g. SE

	Weight     int
	WeightUnit WeightUnit
}

type shippingPriceResponse []shippingPriceResponseItem

func (s shippingPriceResponse) GetResponse() interface{} {
	return s
}

func (s *shippingPriceResponse) SetResponse(body []byte) error {
	fmt.Println(string(body))
	err := json.Unmarshal(body, &s)
	if err != nil {
		return err
	}
	return nil
}

type shippingPriceResponseItem struct {
	Price     *string `json:"price"`
	Currency  *string `json:"currency"`
	ToCountry string  `json:"to_country"`
}

type ShippingPrices []ShippingPriceItem

type ShippingPriceItem struct {
	Price     *float64 `json:"price"`
	Currency  *string  `json:"currency"`
	ToCountry string   `json:"to_country"`
}

type Client struct {
	apiKey    string
	apiSecret string
}

func New(apiKey, apiSecret string) *Client {
	return &Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

// Returns the sendcloud pickup point ID mapped from a SPID ID
func (service Client) GetShippingPrice(params ShippingPriceParams) (ShippingPrices, error) {
	//prepare bounding box url
	uri, _ := url.Parse("https://panel.sendcloud.sc/api/v2/shipping-price/")
	paramsContainer := uri.Query()
	paramsContainer.Set("shipping_method_id", fmt.Sprintf("%d", params.ShippingMethodId))
	paramsContainer.Set("from_country", params.FromCountry)

	if params.ToCountry != nil {
		paramsContainer.Set("to_country", *params.ToCountry)
	}

	paramsContainer.Set("weight", fmt.Sprintf("%d", params.Weight))
	paramsContainer.Set("weight_unit", string(params.WeightUnit))

	uri.RawQuery = paramsContainer.Encode()

	var prices shippingPriceResponse
	if err := sendcloud.Request("GET", uri.String(), nil, service.apiKey, service.apiSecret, &prices); err != nil {
		return nil, err
	}

	formatted := make(ShippingPrices, len(prices))
	for idx, p := range prices {
		if p.Price != nil {
			price, err := strconv.ParseFloat(*p.Price, 64)

			if err != nil {
				return nil, err
			}

			formatted[idx] = ShippingPriceItem{
				Price:     &price,
				Currency:  p.Currency,
				ToCountry: p.ToCountry,
			}
		}
	}

	return formatted, nil
}
