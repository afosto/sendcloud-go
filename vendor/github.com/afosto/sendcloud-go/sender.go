package sendcloud

import "encoding/json"

type SenderResponseContainer struct {
	SenderAddresses []SenderResponse `json:"sender_addresses"`
}

type Sender struct {
	ID          int64
	CompanyName string
	Email       string
	PhoneNumber string
	Street      string
	HouseNumber string
	PostCode    string
	City        string
	CountryCode string
}

type SenderResponse struct {
	ID          int64  `json:"id"`
	CompanyName string `json:"company_name"`
	ContactName string `json:"contact_name"`
	Email       string `json:"email"`
	Telephone   string `json:"telephone"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	PostalBox   string `json:"postal_box"`
	PostalCode  string `json:"postal_code"`
	City        string `json:"city"`
	Country     string `json:"country"`
}

//Get formatted response
func (a *SenderResponseContainer) GetResponse() interface{} {
	var senders []*Sender
	for _, sa := range a.SenderAddresses {
		sender := &Sender{
			ID:          sa.ID,
			CompanyName: sa.CompanyName,
			Email:       sa.Email,
			PhoneNumber: sa.Telephone,
			Street:      sa.Street,
			HouseNumber: sa.HouseNumber,
			PostCode:    sa.PostalCode,
			City:        sa.City,
			CountryCode: sa.Country,
		}
		senders = append(senders, sender)
	}

	return senders
}

//Set the response
func (a *SenderResponseContainer) SetResponse(body []byte) error {
	err := json.Unmarshal(body, &a)
	if err != nil {
		return err
	}
	return nil
}
