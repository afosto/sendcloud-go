package sendcloud

import (
	"encoding/json"
)

type ReturnParams struct {
	// FROM address details
	FromName         string
	FromCompanyName  string
	FromAddressLine1 string
	FromAddressLine2 string
	FromHouseNumber  string
	FromPostalCode   string
	FromCity         string
	FromCountryCode  string
	FromEmail        string
	FromPhoneNumber  string

	// TO address details
	ToName         string
	ToCompanyName  string
	ToAddressLine1 string
	ToAddressLine2 string
	ToHouseNumber  string
	ToPostalCode   string
	ToCity         string
	ToCountryCode  string
	ToEmail        string
	ToPhoneNumber  string

	// Shipping method details for the returns.
	ShipWithType        string                 // "shipping_option_code" or "shipping_product_code"
	ShippingOptionCode  string                 // Required if ShipWithType is "shipping_option_code"
	ShippingProductCode string                 // Required if ShipWithType is "shipping_product_code"
	Functionalities     map[string]interface{} // Optional shipping functionalities (e.g., {"labelless": true})
	Contract            int64                  // Carrier contract id if more than one active contract exists

	// Parcel specifications
	DimensionLength float64 // Length of a single collo
	DimensionWidth  float64 // Width of a single collo
	DimensionHeight float64 // Height of a single collo
	DimensionUnit   string  // e.g. "cm"

	WeightValue float64 // Total weight in kilograms
	WeightUnit  string  // e.g. "kg"

	// Return-specific details
	ColloCount         int                 // Number of collos (default is 1)
	ParcelItems        []ReturnItemRequest // List of items included in the returns (mandatory for outside-EU returns)
	SendTrackingEmails bool                // When true, Sendcloud sends tracking emails
	BrandID            int64               // ID of the brand for this returns
	TotalInsuredValue  *Price              // Optional insured value
	OrderNumber        string              // Order number associated with the returns
	ExternalReference  string              // Unique user-generated reference
	CustomsInvoiceNr   *string             // Customs invoice number; required for international returns
	DeliveryOption     string              // e.g. "drop_off_point", "in_store", etc.
	CustomsInformation *CustomsInformation // Optional customs information for international returns
	ApplyRules         *bool               // When true, returns rules are applied (rules take precedence)
}

// ReturnItemRequest represents a single returns item.
type ReturnItemRequest struct {
	ItemID         string                 `json:"item_id,omitempty"`
	Description    string                 `json:"description,omitempty"`
	Quantity       int                    `json:"quantity,omitempty"`
	Weight         Weight                 `json:"weight"` // weight as string, e.g. "0.4"
	Price          Price                  `json:"price,omitempty"`
	HSCode         string                 `json:"hs_code,omitempty"` // Harmonized System Code (e.g. "6205.20")
	OriginCountry  string                 `json:"origin_country,omitempty"`
	SKU            string                 `json:"sku,omitempty"`
	ProductID      string                 `json:"product_id,omitempty"`
	ReturnReasonID string                 `json:"return_reason_id,omitempty"`
	ReturnMessage  string                 `json:"return_message,omitempty"`
	Properties     map[string]interface{} `json:"properties,omitempty"`
}

// ReturnRequest is the JSON body structure for creating a returns.
type ReturnRequest struct {
	FromAddress        Address             `json:"from_address"`
	ToAddress          Address             `json:"to_address"`
	ShipWith           ShipWith            `json:"ship_with,omitempty"`
	Dimensions         Dimension           `json:"dimensions"`
	Weight             Weight              `json:"weight,omitempty"`
	ColloCount         int                 `json:"collo_count,omitempty"`
	ParcelItems        []ReturnItemRequest `json:"parcel_items,omitempty"`
	SendTrackingEmails bool                `json:"send_tracking_emails"`
	BrandID            int64               `json:"brand_id,omitempty"`
	TotalInsuredValue  *Price              `json:"total_insured_value,omitempty"`
	OrderNumber        string              `json:"order_number,omitempty"`
	TotalOrderValue    *Price              `json:"total_order_value,omitempty"`
	ExternalReference  *string             `json:"external_reference,omitempty"`
	CustomsInvoiceNr   *string             `json:"customs_invoice_nr,omitempty"`
	DeliveryOption     string              `json:"delivery_option,omitempty"`
	CustomsInformation *CustomsInformation `json:"customs_information,omitempty"`
	ApplyRules         *bool               `json:"apply_rules,omitempty"`
}

// GetPayload converts ReturnParams into a request payload.
func (r *ReturnParams) GetPayload() interface{} {
	req := ReturnRequest{
		FromAddress: Address{
			Name:         r.FromName,
			CompanyName:  r.FromCompanyName,
			AddressLine1: r.FromAddressLine1,
			AddressLine2: r.FromAddressLine2,
			HouseNumber:  r.FromHouseNumber,
			PostalCode:   r.FromPostalCode,
			City:         r.FromCity,
			CountryCode:  r.FromCountryCode,
			Email:        r.FromEmail,
			PhoneNumber:  r.FromPhoneNumber,
		},
		ToAddress: Address{
			Name:         r.ToName,
			CompanyName:  r.ToCompanyName,
			AddressLine1: r.ToAddressLine1,
			AddressLine2: r.ToAddressLine2,
			HouseNumber:  r.ToHouseNumber,
			PostalCode:   r.ToPostalCode,
			City:         r.ToCity,
			CountryCode:  r.ToCountryCode,
			Email:        r.ToEmail,
			PhoneNumber:  r.ToPhoneNumber,
		},
		ShipWith: ShipWith{
			Type:                r.ShipWithType,
			ShippingOptionCode:  r.ShippingOptionCode,
			ShippingProductCode: r.ShippingProductCode,
			Functionalities:     r.Functionalities,
			Contract:            r.Contract,
		},
		Dimensions: Dimension{
			Length: r.DimensionLength,
			Width:  r.DimensionWidth,
			Height: r.DimensionHeight,
			Unit:   r.DimensionUnit,
		},
		Weight: Weight{
			Value: r.WeightValue,
			Unit:  r.WeightUnit,
		},
		ColloCount:         r.ColloCount,
		ParcelItems:        r.ParcelItems,
		SendTrackingEmails: r.SendTrackingEmails,
		BrandID:            r.BrandID,
		TotalInsuredValue:  r.TotalInsuredValue,
		OrderNumber:        r.OrderNumber,
		DeliveryOption:     r.DeliveryOption,
		CustomsInvoiceNr:   r.CustomsInvoiceNr,
		CustomsInformation: r.CustomsInformation,
		ApplyRules:         r.ApplyRules,
	}
	if r.ExternalReference != "" {
		req.ExternalReference = &r.ExternalReference
	}
	return req
}

// ReturnResponse represents the basic response from a create returns call.
type ReturnResponse struct {
	ReturnID      int64   `json:"return_id"`
	ParcelID      int64   `json:"parcel_id"`
	MultiColloIDs []int64 `json:"multi_collo_ids"`
}

// SetResponse unmarshals the API response into ReturnResponseContainer.
func (r *ReturnResponse) SetResponse(body []byte) error {
	return json.Unmarshal(body, r)
}

// GetResponse returns the unmarshaled ReturnResponse.
func (r *ReturnResponse) GetResponse() interface{} {
	return r
}

// Address represents a Sendcloud address.
type Address struct {
	Name         string `json:"name"`
	CompanyName  string `json:"company_name,omitempty"`
	AddressLine1 string `json:"address_line_1,omitempty"`
	AddressLine2 string `json:"address_line_2,omitempty"`
	HouseNumber  string `json:"house_number,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`
	City         string `json:"city,omitempty"`
	CountryCode  string `json:"country_code,omitempty"`
	Email        string `json:"email,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
}

// ShipWith defines the shipping method to be used for a returns.
type ShipWith struct {
	Type                string                 `json:"type,omitempty"` // either "shipping_option_code" or "shipping_product_code"
	ShippingOptionCode  string                 `json:"shipping_option_code,omitempty"`
	ShippingProductCode string                 `json:"shipping_product_code,omitempty"`
	Functionalities     map[string]interface{} `json:"functionalities,omitempty"`
	Contract            int64                  `json:"contract,omitempty"`
}

// Dimension specifies the size of a single collo.
type Dimension struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Unit   string  `json:"unit"`
}

// Weight specifies the weight.
type Weight struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// Price represents a monetary amount.
type Price struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

// CustomsInformation holds customs-related details for international returns.
type CustomsInformation struct {
	InvoiceNumber string `json:"invoice_number"`
	ExportReason  string `json:"export_reason"`
	ExportType    string `json:"export_type"`  // e.g., "commercial_goods"
	InvoiceDate   string `json:"invoice_date"` // ISO 8601 date (e.g., "2023-08-24")
}
