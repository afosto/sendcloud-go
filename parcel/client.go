package parcel

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	sendcloud "github.com/afosto/sendcloud-go"
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

//Create a new parcel
func (c *Client) New(params *sendcloud.ParcelParams) (*sendcloud.Parcel, error) {
	parcel := sendcloud.ParcelResponseContainer{}
	_, err := sendcloud.Request("POST", "/api/v2/parcels", params, c.apiKey, c.apiSecret, &parcel)

	if err != nil {
		return nil, err
	}
	r := parcel.GetResponse().(*sendcloud.Parcel)
	return r, nil
}

//Validate and read the incoming webhook
func (c *Client) ReadParcelWebhook(payload []byte, signature string) (*sendcloud.Parcel, error) {
	hash := hmac.New(sha256.New, []byte(c.apiSecret))
	hash.Write(payload)

	expectedSignature := hex.EncodeToString(hash.Sum(nil))
	if signature != expectedSignature {
		return nil, errors.New("invalid signature")
	}

	parcelResponse := sendcloud.ParcelResponseContainer{}
	err := json.Unmarshal(payload, &parcelResponse)
	if err != nil {
		return nil, err
	}

	return parcelResponse.GetResponse().(*sendcloud.Parcel), nil
}
