package client

import (
	"github.com/afosto/sendcloud-go/integration"
	"github.com/afosto/sendcloud-go/method"
	"github.com/afosto/sendcloud-go/parcel"
	"github.com/afosto/sendcloud-go/sender"
)

type API struct {
	Parcel      *parcel.Client
	Method      *method.Client
	Sender      *sender.Client
	Integration *integration.Client
}

//Initialize the client
func (a *API) Init(apiKey string, apiSecret string) {
	a.Parcel = parcel.New(apiKey, apiSecret)
	a.Method = method.New(apiKey, apiSecret)
	a.Sender = sender.New(apiKey, apiSecret)
	a.Integration = integration.New(apiKey, apiSecret)
}
