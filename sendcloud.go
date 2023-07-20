package sendcloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Payload interface {
	GetPayload() interface{}
}

type Response interface {
	GetResponse() interface{}
	SetResponse(body []byte) error
}

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Request string `json:"request"`
		Message string `json:"message"`
	} `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Request string `json:"request"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("request %s resulted in error code %d: %s", e.Request, e.Code, e.Message)
}

// NewRequest creates and prepares a *http.Request with the given method, url,
// payload and credentials, so it's ready to be sent to Sendcloud.
func NewRequest(ctx context.Context, method, uri string, payload Payload, apiKey, apiSecret string) (*http.Request, error) {
	var request *http.Request
	var err error

	if payload == nil {
		request, err = http.NewRequestWithContext(ctx, method, getUrl(uri), nil)
		if err != nil {
			return nil, err
		}
	} else {
		body, err := json.Marshal(payload.GetPayload())
		if err != nil {
			return nil, err
		}
		request, err = http.NewRequestWithContext(ctx, method, getUrl(uri), bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
	}

	if payload != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("User-Agent", "Sendcloud-Go/0.1 ("+apiKey+")")
	request.SetBasicAuth(apiKey, apiSecret)
	return request, nil
}

// Request sends a request to Sendcloud with given method, path, payload and credentials.
func Request(method, uri string, payload Payload, apiKey, apiSecret string, r Response) error {
	request, err := NewRequest(context.Background(), method, uri, payload, apiKey, apiSecret)
	if err != nil {
		return err
	}

	client := http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err = ValidateResponse(response); err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return r.SetResponse(body)
}

// ValidateResponse validates a received response from Sendcloud. It is valid
// and returns nil when the status code is between 200 and 299.
func ValidateResponse(response *http.Response) error {
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		return nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if !strings.Contains(response.Header.Get("content-type"), "application/json") {
		return &Error{
			Code:    response.StatusCode,
			Request: requestFromResponse(response),
			Message: string(body),
		}
	}

	var errResponse ErrorResponse
	if err = json.Unmarshal(body, &errResponse); err != nil {
		return err
	}
	if errResponse.Error.Request == "" {
		errResponse.Error.Request = requestFromResponse(response)
	}
	if errResponse.Error.Message == "" {
		errResponse.Error.Message = string(body)
	}

	return &Error{
		Code:    response.StatusCode,
		Request: errResponse.Error.Request,
		Message: errResponse.Error.Message,
	}
}

// Return the full URL
func getUrl(uri string) string {
	var url string
	if strings.HasPrefix(uri, "https://") {
		url = uri
	} else {
		url = "https://panel.sendcloud.sc" + uri
	}

	return url
}

func requestFromResponse(resp *http.Response) string {
	if resp.Request != nil && resp.Request.URL != nil {
		return resp.Request.URL.String()
	}
	return ""
}
