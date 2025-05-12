package sendcloud

import (
	"encoding/json"
)

// OptionParams represents the parameters for fetching shipping options.
// For the simplified client methods (GetDeliveryShippingOptions, GetReturnShippingOptions),
// only Functionalities.Returns will be set.
type OptionParams struct {
	FromCountryCode     *string              `json:"from_country_code,omitempty"`
	ToCountryCode       *string              `json:"to_country_code,omitempty"`
	Functionalities     *FunctionalityFilter `json:"functionalities,omitempty"`
	CarrierCode         *string              `json:"carrier_code,omitempty"`
	ContractID          *int64               `json:"contract_id,omitempty"`
	ShippingProductCode *string              `json:"shipping_product_code,omitempty"`
	Dimensions          *OptionDimension     `json:"dimensions,omitempty"` // For request body
	Weight              *OptionWeight        `json:"weight,omitempty"`     // For request body
	FromPostalCode      *string              `json:"from_postal_code,omitempty"`
	ToPostalCode        *string              `json:"to_postal_code,omitempty"`
	TotalInsurance      *float64             `json:"total_insurance,omitempty"`
	LeadTime            *LeadTimeFilter      `json:"lead_time,omitempty"`
}

// FunctionalityFilter allows specifying functionality-based filters.
// For the simplified client, only 'Returns' will be used.
type FunctionalityFilter struct {
	Returns *bool `json:"returns,omitempty"`
	// You can add other functionality filters here if needed for more advanced use cases
	// e.g., Signature *bool `json:"signature,omitempty"`
}

// OptionDimension specifies dimensions for shipping options in a request.
// Values are strings as per the OpenAPI specification for request bodies.
type OptionDimension struct {
	Length string `json:"length"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Unit   string `json:"unit"` // e.g., "cm", "m"
}

// OptionWeight specifies weight for shipping options in a request.
// Value is a string as per the OpenAPI specification for request bodies.
type OptionWeight struct {
	Value string `json:"value"` // e.g. "2.0"
	Unit  string `json:"unit"`  // e.g., "kg", "g"
}

// LeadTimeFilter allows filtering shipping options based on lead time.
type LeadTimeFilter struct {
	GT  *float64 `json:"gt,omitempty"`
	GTE *float64 `json:"gte,omitempty"`
	EQ  *float64 `json:"eq,omitempty"`
	LT  *float64 `json:"lt,omitempty"`
	LTE *float64 `json:"lte,omitempty"`
}

// GetPayload returns the OptionParams itself, as it directly matches the API request body structure.
func (p *OptionParams) GetPayload() interface{} {
	return p
}

// --- Response Structs (to parse the API response) ---

// OptionResponse is the top-level structure for the shipping options API response.
type OptionResponse struct {
	Data []*ShippingOption `json:"data"` // 'data' can be an array or null according to API spec.
}

// SetResponse unmarshals the API response body into the OptionResponse struct.
func (r *OptionResponse) SetResponse(body []byte) error {
	return json.Unmarshal(body, r)
}

// GetResponse returns the OptionResponse struct itself.
func (r *OptionResponse) GetResponse() interface{} {
	return r
}

// ShippingOption represents a single shipping option with its details, pricing, and functionalities.
type ShippingOption struct {
	Code            string                          `json:"code,omitempty"`
	Carrier         *CarrierOption                  `json:"carrier,omitempty"`
	Product         *ShippingProduct                `json:"product,omitempty"`
	Functionalities *CarrierShippingFunctionalities `json:"functionalities,omitempty"` // Response struct for functionalities
	MaxDimensions   *ResponseDimension              `json:"max_dimensions,omitempty"`  // Response struct for dimensions
	Weight          *ShippingOptionWeightRange      `json:"weight,omitempty"`
	BilledWeight    *BilledWeight                   `json:"billed_weight,omitempty"`
	Contract        *Contract                       `json:"contract,omitempty"`
	Requirements    *Requirements                   `json:"requirements,omitempty"`
	Quotes          []*ShippingQuote                `json:"quotes"` // Can be null from API, unmarshals to nil slice
}

// Carrier represents carrier information.
type CarrierOption struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// ShippingProduct represents product information for a shipping option.
type ShippingProduct struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// CarrierShippingFunctionalities defines various features or attributes of a shipping option from the API response.
// Pointers are used for fields that can be 'null' in the API response.
type CarrierShippingFunctionalities struct {
	AgeCheck             *int    `json:"age_check"`
	B2B                  bool    `json:"b2b"`
	B2C                  bool    `json:"b2c"`
	Boxable              bool    `json:"boxable"`
	BulkyGoods           bool    `json:"bulky_goods"`
	CarrierBillingType   string  `json:"carrier_billing_type"`
	CashOnDelivery       *int    `json:"cash_on_delivery"`
	DangerousGoods       bool    `json:"dangerous_goods"`
	DeliveryAttempts     *int    `json:"delivery_attempts"`
	DeliveryBefore       string  `json:"delivery_before"`
	DeliveryDeadline     string  `json:"delivery_deadline"`
	EcoDelivery          bool    `json:"eco_delivery"`
	ERS                  bool    `json:"ers"`
	FirstMile            string  `json:"first_mile"`
	FlexDelivery         bool    `json:"flex_delivery"`
	FormFactor           *string `json:"form_factor"`
	FragileGoods         bool    `json:"fragile_goods"`
	FreshGoods           bool    `json:"fresh_goods"`
	HarmonizedLabel      bool    `json:"harmonized_label"`
	IDCheck              bool    `json:"id_check"`
	Incoterm             string  `json:"incoterm"`
	Insurance            *int    `json:"insurance"`
	Labelless            bool    `json:"labelless"`
	LastMile             string  `json:"last_mile"`
	Manually             bool    `json:"manually"`
	Multicollo           bool    `json:"multicollo"`
	NeighborDelivery     bool    `json:"neighbor_delivery"`
	NonConveyable        bool    `json:"non_conveyable"`
	PersonalizedDelivery bool    `json:"personalized_delivery"`
	PickUp               bool    `json:"pick_up"`
	Premium              bool    `json:"premium"`
	Priority             string  `json:"priority"`
	RegisteredDelivery   bool    `json:"registered_delivery"`
	Returns              bool    `json:"returns"`
	Segment              string  `json:"segment"`
	ServiceArea          string  `json:"service_area"`
	Signature            bool    `json:"signature"`
	Size                 string  `json:"size"`
	Sorted               bool    `json:"sorted"`
	Surcharge            bool    `json:"surcharge"`
	Tracked              bool    `json:"tracked"`
	Tyres                bool    `json:"tyres"`
	WeekendDelivery      string  `json:"weekend_delivery"`
}

// ResponseDimension specifies dimensions from an API response.
// Values are strings as per the OpenAPI specification.
type ResponseDimension struct {
	Length string `json:"length"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Unit   string `json:"unit"`
}

// ResponseWeight specifies weight from an API response.
// Value is a string as per the OpenAPI specification.
type ResponseWeight struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

// ShippingOptionWeightRange defines the minimum and maximum weight for a shipping option.
type ShippingOptionWeightRange struct {
	Min *ResponseWeight `json:"min,omitempty"`
	Max *ResponseWeight `json:"max,omitempty"`
}

// BilledWeight indicates the weight used for billing, considering volumetric weight.
type BilledWeight struct {
	Value      string `json:"value"`
	Unit       string `json:"unit"`
	Volumetric bool   `json:"volumetric"`
}

// Contract represents a carrier contract.
type Contract struct {
	ID          int64  `json:"id,omitempty"`
	ClientID    string `json:"client_id,omitempty"`
	CarrierCode string `json:"carrier_code,omitempty"`
	Name        string `json:"name,omitempty"`
}

// Requirements indicate necessary fields or documents for a shipping option.
type Requirements struct {
	Fields          []string `json:"fields,omitempty"`
	ExportDocuments bool     `json:"export_documents"`
}

// ShippingQuote provides pricing details for a shipping option within a specific weight range.
type ShippingQuote struct {
	Weight   QuoteWeightRange `json:"weight,omitempty"`
	LeadTime *int             `json:"lead_time,omitempty"` // API spec: integer or null
	Price    QuotePrice       `json:"price,omitempty"`
}

// QuoteWeightRange defines the weight range for a specific quote.
type QuoteWeightRange struct {
	Min ResponseWeight `json:"min,omitempty"`
	Max ResponseWeight `json:"max,omitempty"`
}

// QuotePrice contains the total price and its breakdown for a shipping quote.
type QuotePrice struct {
	Breakdown []*ShippingPriceBreakdownItem `json:"breakdown,omitempty"`
	Total     *ResponsePrice                `json:"total,omitempty"`
}

// ShippingPriceBreakdownItem details a component of the shipping price.
type ShippingPriceBreakdownItem struct {
	Type  string         `json:"type,omitempty"`
	Label string         `json:"label,omitempty"`
	Price *ResponsePrice `json:"price,omitempty"`
}

// ResponsePrice represents a monetary value with currency from an API response.
// Value is a string as per the OpenAPI specification.
type ResponsePrice struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

// Helper functions for creating pointers.
// Useful if constructing OptionParams manually with more specific filters.
func Bool(b bool) *bool          { v := b; return &v }
func String(s string) *string    { v := s; return &v }
func Int(i int) *int             { v := i; return &v }
func Int64(i int64) *int64       { v := i; return &v }
func Float64(f float64) *float64 { v := f; return &v }
