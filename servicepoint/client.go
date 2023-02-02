package servicepoint

import (
	"errors"
	"fmt"
	"github.com/afosto/sendcloud-go"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"net/url"
	"strings"
	"unicode"
)

var (
	NoSuitableServicePointFound = errors.New("could not find a matching service point in SendCloud")
)

type Client struct {
	apiKey    string
	apiSecret string
}

type Matcher struct {
	SPID        string
	Carrier     string
	Country     string
	PostalCode  string
	HouseNumber string
	Latitude    float64
	Longitude   float64
}

func New(apiKey string, apiSecret string) *Client {
	return &Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

// Returns the sendcloud pickup point ID mapped from a SPID ID
func (service Client) GetServicePoint(servicePoint Matcher) (int, error) {
	//prepare bounding box url
	uri, _ := url.Parse("https://servicePoints.sendcloud.sc/api/v2/service-points/")
	params := map[string]string{
		"country":      strings.ToUpper(servicePoint.Country),
		"ne_latitude":  fmt.Sprintf("%.4f", servicePoint.Latitude+0.06),
		"sw_latitude":  fmt.Sprintf("%.4f", servicePoint.Latitude-0.06),
		"ne_longitude": fmt.Sprintf("%.4f", servicePoint.Longitude+0.06),
		"sw_longitude": fmt.Sprintf("%.4f", servicePoint.Longitude-0.06),
		"access_token": service.apiKey,
		"carrier":      servicePoint.Carrier,
	}
	paramsContainer := uri.Query()
	for key, value := range params {
		paramsContainer.Add(key, value)
	}
	uri.RawQuery = paramsContainer.Encode()

	servicePoints := sendcloud.ServicePointList{}
	if err := sendcloud.Request("GET", uri.String(), nil, service.apiKey, service.apiSecret, &servicePoints); err != nil {
		return 0, err
	}

	matching := unaccent(fmt.Sprintf("%s %s", servicePoint.PostalCode, servicePoint.HouseNumber))

	for _, sp := range servicePoints {

		if unaccent(sp.Identifier()) == matching {
			return sp.ID, nil
		}
		if sp.Code == servicePoint.SPID {
			return sp.ID, nil
		}
	}

	return 0, NoSuitableServicePointFound
}

func unaccent(string string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, string)
	return result
}
