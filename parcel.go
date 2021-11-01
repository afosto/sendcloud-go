package sendcloud

import (
	"encoding/json"
	"strconv"
	"time"
)

type LabelData []byte

type ParcelParams struct {
	Name             string
	CompanyName      string
	Street           string
	HouseNumber      string
	City             string
	PostalCode       string
	State            string
	CountryCode      string
	AdditionalInfo   string
	IsLabelRequested bool
	Method           int64
	EmailAddress     string
	PhoneNumber      string
	ExternalID       string
	ToServicePointID int64
	Weight           string
	OrderNumber      string
	SenderID         int64
}

type Parcel struct {
	ID             int64       `json:"id"`
	ExternalID     *string     `json:"external_id"`
	Name           string      `json:"name"`
	CompanyName    string      `json:"company_name"`
	Email          string      `json:"email"`
	Street         string      `json:"street"`
	HouseNumber    string      `json:"house_number"`
	Address        string      `json:"address"`
	Address2       string      `json:"address_2"`
	City           string      `json:"city"`
	PostalCode     string      `json:"postal_code"`
	CountryCode    string      `json:"country_code"`
	Method         int64       `json:"method"`
	PhoneNumber    *string     `json:"phone_number"`
	TrackingNumber string      `json:"tracking_number"`
	ServicePointID *int64      `json:"to_service_point"`
	Weight         string       `json:"weight"`
	Label          string      `json:"label"`
	OrderNumber    string      `json:"order_number"`
	IsReturn       bool        `json:"is_return"`
	Note           *string     `json:"note"`
	CarrierCode    string      `json:"carrier"`
	Data           interface{} `json:"data"`
	CreatedAt      time.Time   `json:"created_at"`
}

type ParcelRequestContainer struct {
	Parcel ParcelRequest `json:"parcel"`
}

type ParcelRequest struct {
	Name             string  `json:"name"`
	CompanyName      string  `json:"company_name"`
	Address          string  `json:"address"`
	Address2         string  `json:"address_2"`
	HouseNumber      string  `json:"house_number"`
	City             string  `json:"city"`
	PostalCode       string  `json:"postal_code"`
	CountryState     string  `json:"country_state"`
	Country          string  `json:"country"`
	Telephone        string  `json:"telephone"`
	Email            string  `json:"email"`
	RequestLabel     bool    `json:"request_label"`
	ToServicePointID *int64  `json:"to_service_point,omitempty"`
	OrderNumber      string  `json:"order_number"`
	ExternalID       *string `json:"external_reference,omitempty"`
	SenderID         *int64  `json:"sender_address,omitempty"`
	Shipment         struct {
		ID int64 `json:"id"`
	} `json:"shipment"`
}

type LabelResponseContainer struct {
	Label LabelResponse `json:"label"`
}

type ParcelResponseContainer struct {
	Parcel ParcelResponse `json:"parcel"`
}

type ParcelListResponseContainer struct {
	Parcels []*ParcelResponse `json:"parcel"`
}

type LabelResponse struct {
	NormalPrinter []string `json:"normal_printer"`
	LabelPrinter  string   `json:"label_printer"`
}

type ParcelResponse struct {
	ID                  int64           `json:"id"`
	Address             string          `json:"address"`
	Address2            string          `json:"address_2"`
	AddressDivided      AddressDivided  `json:"address_divided"`
	City                string          `json:"city"`
	CompanyName         string          `json:"company_name"`
	Country             CountryResponse `json:"country"`
	Data                interface{}     `json:"data"`
	DateCreated         string          `json:"date_created"`
	Email               string          `json:"email"`
	Name                string          `json:"name"`
	PostalCode          string          `json:"postal_code"`
	Reference           string          `json:"reference"`
	Shipment            Shipment        `json:"shipment"`
	Status              Status          `json:"status"`
	ToServicePointID    *int64          `json:"to_service_point"`
	Telephone           *string         `json:"telephone"`
	TrackingNumber      string          `json:"tracking_number"`
	Weight              string          `json:"weight"`
	Label               LabelResponse   `json:"label"`
	OrderNumber         string          `json:"order_number"`
	InsuredValue        int64           `json:"insured_value"`
	TotalInsuredValue   int64           `json:"total_insured_value"`
	ToState             interface{}     `json:"to_state"`
	CustomsInvoiceNr    string          `json:"customs_invoice_nr"`
	CustomsShipmentType interface{}     `json:"customs_shipment_type"`
	Type                interface{}     `json:"type"`
	ShipmentUUID        *string         `json:"shipment_uuid"`
	ShippingMethod      int64           `json:"shipping_method"`
	ExternalOrderID     *string         `json:"external_order_id"`
	ExternalShipmentID  *string         `json:"external_shipment_id"`
	ExternalReference   *string         `json:"external_reference"`
	IsReturn            bool            `json:"is_return"`
	Note                *string         `json:"note"`
	Carrier             Carrier         `json:"carrier"`
}

type Carrier struct {
	Code string `json:"code"`
}

type AddressDivided struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
}

type Shipment struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Status struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

//Translate the params into an actual request body
func (p *ParcelParams) GetPayload() interface{} {
	parcel := ParcelRequest{
		Name:         p.Name,
		CompanyName:  p.CompanyName,
		Address:      p.Street,
		Address2:     p.AdditionalInfo,
		HouseNumber:  p.HouseNumber,
		City:         p.City,
		PostalCode:   p.PostalCode,
		CountryState: p.State,
		Country:      p.CountryCode,
		Telephone:    p.PhoneNumber,
		Email:        p.EmailAddress,
		RequestLabel: p.IsLabelRequested,
		Shipment: struct {
			ID int64 `json:"id"`
		}{
			ID: p.Method,
		},
	}
	if p.SenderID != 0 {
		parcel.SenderID = &p.SenderID
	}
	if p.ExternalID != "" {
		parcel.ExternalID = &p.ExternalID
	}
	if p.OrderNumber != "" {
		parcel.OrderNumber = p.OrderNumber
	}
	if p.ToServicePointID != 0 {
		parcel.ToServicePointID = &p.ToServicePointID
	}

	ar := ParcelRequestContainer{Parcel: parcel}
	return ar
}

//Handle the response and return it as a Parcel{}
func (p *ParcelResponseContainer) GetResponse() interface{} {
	parcel := Parcel{
		ID:             p.Parcel.ID,
		ExternalID:     p.Parcel.ExternalReference,
		Name:           p.Parcel.Name,
		CompanyName:    p.Parcel.CompanyName,
		Email:          p.Parcel.Email,
		Street:         p.Parcel.AddressDivided.Street,
		HouseNumber:    p.Parcel.AddressDivided.HouseNumber,
		Address:        p.Parcel.Address,
		Address2:       p.Parcel.Address2,
		City:           p.Parcel.City,
		Method:         p.Parcel.Shipment.ID,
		PostalCode:     p.Parcel.PostalCode,
		CountryCode:    p.Parcel.Country.Iso2,
		PhoneNumber:    p.Parcel.Telephone,
		TrackingNumber: p.Parcel.TrackingNumber,
		ServicePointID: p.Parcel.ToServicePointID,
		Label:          p.Parcel.Label.LabelPrinter,
		OrderNumber:    p.Parcel.OrderNumber,
		IsReturn:       p.Parcel.IsReturn,
		Note:           p.Parcel.Note,
		CarrierCode:    p.Parcel.Carrier.Code,
		Data:           p.Parcel.Data,
		Weight:		p.Parcel.Weight
	}

	layout := "02-01-2006 15:04:05"
	createdAt, _ := time.Parse(layout, p.Parcel.DateCreated)
	parcel.CreatedAt = createdAt

	return &parcel
}

//Set the response
func (p *ParcelResponseContainer) SetResponse(body []byte) error {
	err := json.Unmarshal(body, &p)
	if err != nil {
		return err
	}
	return nil
}

//Get formatted response
func (l LabelData) GetResponse() interface{} {
	return l
}

//Set the response
func (l *LabelData) SetResponse(body []byte) error {
	*l = body
	return nil
}
