package parcel

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strconv"

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
	err := sendcloud.Request("POST", "/api/v2/parcels", params, c.apiKey, c.apiSecret, &parcel)

	if err != nil {
		return nil, err
	}
	r := parcel.GetResponse().(*sendcloud.Parcel)
	return r, nil
}

//Return a single parcel
func (c *Client) Get(parcelID int64) (*sendcloud.Parcel, error) {
	parcel := sendcloud.ParcelResponseContainer{}
	err := sendcloud.Request("GET", "/api/v2/parcels/"+strconv.Itoa(int(parcelID)), nil, c.apiKey, c.apiSecret, &parcel)

	if err != nil {
		return nil, err
	}
	r := parcel.GetResponse().(*sendcloud.Parcel)
	return r, nil
}

//Get a label as bytes based on the url that references the PDF
func (c *Client) GetLabel(labelURL string) ([]byte, error) {
	data := &sendcloud.LabelData{}
	err := sendcloud.Request("GET", labelURL, nil, c.apiKey, c.apiSecret, data)
	if err != nil {
		return nil, err
	}
	return *data, nil
}

//Return the portal URL for a parcel
func (c *Client) GetPortalURL(parcelID int64) (string, error) {
	response := sendcloud.PortalURLResponse{}
	err := sendcloud.Request("GET", "/api/v2/parcels/"+strconv.Itoa(int(parcelID))+"/return_portal_url", nil, c.apiKey, c.apiSecret, &response)
	if err != nil {
		return "", err
	}
	if response.URL == nil {
		return "", errors.New("no URL provided by sendcloud")
	}
	return *response.URL, nil
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
