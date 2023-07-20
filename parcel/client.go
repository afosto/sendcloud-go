package parcel

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	sendcloud "github.com/afosto/sendcloud-go"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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

// Create a new parcel
func (c *Client) New(params *sendcloud.ParcelParams) (*sendcloud.Parcel, error) {
	parcel := sendcloud.ParcelResponseContainer{}
	err := sendcloud.Request("POST", "/api/v2/parcels", params, c.apiKey, c.apiSecret, &parcel)

	if err != nil {
		return nil, err
	}
	r := parcel.GetResponse().(*sendcloud.Parcel)
	return r, nil
}

// Return a single parcel
func (c *Client) Get(parcelID int64) (*sendcloud.Parcel, error) {
	parcel := sendcloud.ParcelResponseContainer{}
	err := sendcloud.Request("GET", "/api/v2/parcels/"+strconv.Itoa(int(parcelID)), nil, c.apiKey, c.apiSecret, &parcel)

	if err != nil {
		return nil, err
	}
	r := parcel.GetResponse().(*sendcloud.Parcel)
	return r, nil
}

// Get a label as bytes based on the url that references the PDF
func (c *Client) GetLabel(labelURL string) ([]byte, error) {
	data := &sendcloud.LabelData{}
	err := sendcloud.Request("GET", labelURL, nil, c.apiKey, c.apiSecret, data)
	if err != nil {
		return nil, err
	}
	return *data, nil
}

// Validate and read the incoming webhook
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

// GetDocument retrieves the parcel document of parcelID with type docType from the api.
// https://api.sendcloud.dev/docs/sendcloud-public-api/parcel-documents/operations/get-a-parcel-document
func (c *Client) GetDocument(ctx context.Context, parcelID int64, docTyp string, fmt sendcloud.DocumentFormat, dpi int) (*sendcloud.Document, error) {
	uri := "/api/v2/parcels/" + strconv.Itoa(int(parcelID)) + "/documents/" + docTyp
	if dpi > 0 {
		uri += "?dpi=" + strconv.Itoa(dpi)
	}

	req, err := sendcloud.NewRequest(ctx, "GET", uri, nil, c.apiKey, c.apiSecret)
	if err != nil {
		return nil, err
	}

	if fmt != "" {
		req.Header.Set("accept", fmt.String())
	}

	client := http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err = sendcloud.ValidateResponse(response); err != nil {
		return nil, err
	}

	doc := sendcloud.Document{
		Format: sendcloud.DocumentFormat(response.Header.Get("content-type")),
	}
	if doc.Body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}
	return &doc, nil
}
