package client

import (
	"github.com/afosto/sendcloud-go/integration"
	"github.com/afosto/sendcloud-go/method"
	"github.com/afosto/sendcloud-go/option"
	"github.com/afosto/sendcloud-go/parcel"
	"github.com/afosto/sendcloud-go/products"
	"github.com/afosto/sendcloud-go/returns"
	"github.com/afosto/sendcloud-go/sender"
	"github.com/afosto/sendcloud-go/servicepoint"
)

type API struct {
	Parcel       *parcel.Client
	Method       *method.Client
	Sender       *sender.Client
	ServicePoint *servicepoint.Client
	Integration  *integration.Client
	Return       *returns.Client
	Product      *products.Client
	Option       *option.Client
}

// Initialize the client
func (a *API) Init(apiKey string, apiSecret string) {
	a.Parcel = parcel.New(apiKey, apiSecret)
	a.Method = method.New(apiKey, apiSecret)
	a.Sender = sender.New(apiKey, apiSecret)
	a.ServicePoint = servicepoint.New(apiKey, apiSecret)
	a.Integration = integration.New(apiKey, apiSecret)
	a.Return = returns.New(apiKey, apiSecret)
	a.Product = products.New(apiKey, apiSecret)
	a.Option = option.New(apiKey, apiSecret)
}
