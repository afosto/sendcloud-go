package sendcloud

import (
	"encoding/json"
	"time"
)

type IntegrationParams struct {
	ID                int64
	Name              string
	URL               string
	IsWebhooksEnabled bool
	WebhookURL        string
}

type IntegrationRequest struct {
	ShopName             string   `json:"shop_name"`
	ShopURL              string   `json:"shop_url"`
	ServicePointEnabled  bool     `json:"service_point_enabled"`
	ServicePointCarriers []string `json:"service_point_carriers"`
	WebhookActive        bool     `json:"webhook_active"`
	WebhookURL           string   `json:"webhook_url"`
}

type IntegrationListResponseContainer []IntegrationResponseContainer

type IntegrationResponseContainer struct {
	ID                   int64     `json:"id"`
	ShopName             string    `json:"shop_name"`
	ShopURL              string    `json:"shop_url"`
	System               string    `json:"system"`
	FailingSince         string    `json:"failing_since"`
	LastFetch            string    `json:"last_fetch"`
	LastUpdatedAt        time.Time `json:"last_updated_at"`
	ServicePointEnabled  bool      `json:"service_point_enabled"`
	ServicePointCarriers []string  `json:"service_point_carriers"`
	WebhookActive        bool      `json:"webhook_active"`
	WebhookURL           string    `json:"webhook_url"`
}

type Integration struct {
	ID                    int64
	Name                  string
	URL                   string
	Type                  string
	IsServicePointEnabled bool
	ServicePointCarriers  []string
	IsWebhooksEnabled     bool
	WebhookURL            string
	UpdatedAt             time.Time
}

func (i *IntegrationParams) GetPayload() interface{} {
	return IntegrationRequest{
		ShopName:             i.Name,
		ShopURL:              i.URL,
		ServicePointEnabled:  false,
		ServicePointCarriers: []string{},
		WebhookActive:        i.IsWebhooksEnabled,
		WebhookURL:           i.WebhookURL,
	}
}

//Get formatted response
func (i *IntegrationListResponseContainer) GetResponse() interface{} {
	var integrations []*Integration
	for _, r := range *i {
		integration := Integration{
			ID:                    r.ID,
			Name:                  r.ShopName,
			URL:                   r.ShopURL,
			Type:                  r.System,
			IsServicePointEnabled: r.ServicePointEnabled,
			ServicePointCarriers:  r.ServicePointCarriers,
			IsWebhooksEnabled:     r.ServicePointEnabled,
			WebhookURL:            r.WebhookURL,
			UpdatedAt:             r.LastUpdatedAt,
		}
		integrations = append(integrations, &integration)
	}

	return integrations
}

//Get formatted response
func (r *IntegrationResponseContainer) GetResponse() interface{} {
	integration := &Integration{
		ID:                    r.ID,
		Name:                  r.ShopName,
		URL:                   r.ShopURL,
		Type:                  r.System,
		IsServicePointEnabled: r.ServicePointEnabled,
		ServicePointCarriers:  r.ServicePointCarriers,
		IsWebhooksEnabled:     r.ServicePointEnabled,
		WebhookURL:            r.WebhookURL,
		UpdatedAt:             r.LastUpdatedAt,
	}

	return integration
}

//Set the response
func (r *IntegrationResponseContainer) SetResponse(body []byte) error {
	err := json.Unmarshal(body, &r)
	if err != nil {
		return err
	}
	return nil
}

//Set the response
func (i *IntegrationListResponseContainer) SetResponse(body []byte) error {
	err := json.Unmarshal(body, &i)
	if err != nil {
		return err
	}
	return nil
}
