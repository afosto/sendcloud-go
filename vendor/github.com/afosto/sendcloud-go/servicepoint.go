package sendcloud

import (
	"encoding/json"
	"fmt"
)

type ServicePoint struct {
	ID                    int         `json:"id"`
	Code                  string      `json:"code"`
	IsActive              bool        `json:"is_active"`
	ExtraData             interface{} `json:"extra_data"`
	Name                  string      `json:"name"`
	Street                string      `json:"street"`
	HouseNumber           string      `json:"house_number"`
	PostalCode            string      `json:"postal_code"`
	City                  string      `json:"city"`
	Latitude              string      `json:"latitude"`
	Longitude             string      `json:"longitude"`
	Email                 string      `json:"email"`
	Phone                 string      `json:"phone"`
	Homepage              string      `json:"homepage"`
	Carrier               string      `json:"carrier"`
	Country               string      `json:"country"`
	FormattedOpeningTimes struct {
		Num0 []string `json:"0"`
		Num1 []string `json:"1"`
		Num2 []string `json:"2"`
		Num3 []string `json:"3"`
		Num4 []string `json:"4"`
		Num5 []string `json:"5"`
		Num6 []string `json:"6"`
	} `json:"formatted_opening_times"`
	OpenTomorrow bool `json:"open_tomorrow"`
}

func (s *ServicePoint) Identifier() string {
	return fmt.Sprintf("%s %s", s.PostalCode, s.HouseNumber)
}

type ServicePointList []ServicePoint

func (s ServicePointList) GetResponse() interface{} {
	return s
}

func (s *ServicePointList) SetResponse(body []byte) error {
	err := json.Unmarshal(body, &s)
	if err != nil {
		return err
	}
	return nil
}
