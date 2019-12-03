//Sendcloud-GO aims to make interaction with Sendcloud's API's easier.
//
//Currenty supported are integrations, methods, parcels and sender addresses.
//This package is under heavy development and currently depends on github.com/dghubble/sling which is a dependency we would like to remove in the near future.
package sendcloud

import (
	"errors"
	"github.com/dghubble/sling"
)

const (
	ApiURL string = "https://panel.sendcloud.sc"
)

type Payload interface {
	GetPayload() interface{}
}

type Response interface {
	GetResponse() interface{}
}

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Request string `json:"request"`
		Message string `json:"message"`
	} `json:"error"`
}

//Send a request to Sendcloud's API with given method, path, payload and credentials
func Request(method string, path string, payload Payload, apiKey string, apiSecret string, r Response) (*ErrorResponse, error) {
	req := sling.New().Base(ApiURL)
	switch method {
	case "POST":
		req = req.Post(path)
	case "PUT":
		req = req.Put(path)
	case "GET":
		req = req.Get(path)
	}

	if payload != nil {
		req = req.BodyJSON(payload.GetPayload())
	}

	errM := &ErrorResponse{}
	req = req.SetBasicAuth(apiKey, apiSecret)
	resp, err := req.Receive(r, errM)
	if err != nil {
		return errM, err
	}
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return errM, errors.New("got " + resp.Status)
	}
	return nil, nil
}
