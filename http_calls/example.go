package http_calls

import (
	"errors"
	"github.com/federicoleon/go-httpclient/gohttp"
)

var (
	httpClient = gohttp.NewBuilder().Build()
)

type Endpoints struct {
	EventsUrl string `json:"events_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() > 299 {
		return nil, errors.New("error when trying to fetch github endpoints")
	}

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}
	return &endpoints, nil
}
