# sendcloud-go

An API-client for Sendcloud written in Golang.

## Overview [![GoDoc](https://godoc.org/github.com/afosto/sendcloud-go?status.svg)](https://godoc.org/github.com/afosto/sendcloud-go)

This package is under heavy development, but currently supports 
- parcels
- labels
- methods
- addresses
- integrations

## Install

```
go get github.com/afosto/sendcloud-go
```

## Examples

Some examples on how to use the client are found below.

To initialize the client:
```go
api := client.API{}
api.Init("api_key", "api_secret")
```

To list shipping methods:
```go
methods, err := api.Method.GetMethods()
if err != nil {
    log.Fatal(err)
}

for _, m := range methods {
    log.Print(*m)
}
```
Create a parcel:
```go
params := &sendcloud.ParcelParams{
	Name:             "Sendcloud-GO",
	CompanyName:      "Afosto SaaS BV",
	Street:           "Grondzijl",
	HouseNumber:      "16",
	City:             "Groningen",
	PostalCode:       "9731DG",
	PhoneNumber:      "0507119519",
	EmailAddress:     "peter@afosto.io",
	CountryCode:      "NL",
	IsLabelRequested: true,
	Method:           8,
	ExternalID:       uuid.New().String(),
}
parcel, err := api.Parcel.New(params)
```


## Contributing

All contributions / suggestions are welcome. Please raise an issue or submit a PR.


## Thanks

This package was heavily inspired by `stripe-go`.

## Author

This package is developed by [Afosto SaaS BV](https://afosto.com).

## License

MIT.
