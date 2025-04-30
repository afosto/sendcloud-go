package sendcloud

import "encoding/json"

type Product struct {
	Name                 string `json:"name"`
	Code                 string `json:"code"`
	Carrier              string `json:"carrier"`
	ServicePointsCarrier string `json:"service_points_carrier"`
	WeightRange          struct {
		MinWeight int `json:"min_weight"`
		MaxWeight int `json:"max_weight"`
	} `json:"weight_range"`
	Methods []struct {
		Id                  int    `json:"id"`
		Name                string `json:"name"`
		ShippingProductCode string `json:"shipping_product_code"`
		Properties          struct {
			MinWeight     int `json:"min_weight"`
			MaxWeight     int `json:"max_weight"`
			MaxDimensions struct {
				Length int    `json:"length"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
				Unit   string `json:"unit"`
			} `json:"max_dimensions"`
		} `json:"properties"`
		LeadTimeHours struct {
			NL struct {
				NL int `json:"NL"`
			} `json:"NL"`
		} `json:"lead_time_hours"`
	} `json:"methods"`
}

type ProductListResponseContainer struct {
	Products []Product `json:"products"`
}

func (p *ProductListResponseContainer) SetResponse(body []byte) error {
	return json.Unmarshal(body, &p.Products)
}

type ProductResponseContainer struct {
	Product Product `json:"product"`
}

func (p *ProductResponseContainer) GetResponse() interface{} {
	return p.Product
}

func (p *ProductListResponseContainer) GetResponse() interface{} {
	var products []*Product
	for i := range p.Products {
		products = append(products, &p.Products[i])
	}

	return products
}
