package sendcloud_test

import (
	"encoding/json"
	"testing"

	"github.com/itsrever/sendcloud-go"
	"github.com/stretchr/testify/assert"
)

func TestGetPayload(t *testing.T) {
	tests := []struct {
		Name   string
		Params sendcloud.ParcelParams
	}{
		{
			Name:   "Should ignore empty weight",
			Params: sendcloud.ParcelParams{},
		},
		{
			Name: "Should include weight in request",
			Params: sendcloud.ParcelParams{
				Weight: "0.040",
			},
		},
	}
	for _, test := range tests {
		payload := test.Params.GetPayload()
		b, _ := json.Marshal(payload)
		var obj map[string]map[string]interface{}

		err := json.Unmarshal(b, &obj)
		assert.NoError(t, err)

		if test.Params.Weight == "" {
			_, ok := obj["parcel"]["weight"]
			assert.False(t, ok)
		}
	}
}

func TestGetResponse(t *testing.T) {
	tests := []struct {
		Name     string
		Response sendcloud.ParcelResponseContainer
		Out      sendcloud.Parcel
	}{
		{
			Name: "Should ignore nil weight",
			Response: sendcloud.ParcelResponseContainer{
				Parcel: sendcloud.ParcelResponse{},
			},
			Out: sendcloud.Parcel{},
		},
		{
			Name: "Should include weight",
			Response: sendcloud.ParcelResponseContainer{
				Parcel: sendcloud.ParcelResponse{
					Weight: "0.05",
				},
			},
			Out: sendcloud.Parcel{
				Weight: "0.05",
			},
		},
	}
	for _, test := range tests {
		res := test.Response.GetResponse()
		if test.Response.Parcel.Weight != "" {
			assert.Equal(t, test.Out.Weight, res.(*sendcloud.Parcel).Weight, test.Name)
		} else {
			assert.Equal(t, res.(*sendcloud.Parcel).Weight, "", test.Name)
		}
	}
}
