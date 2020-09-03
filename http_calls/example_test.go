package http_calls

import (
	"errors"
	"github.com/federicoleon/go-httpclient/gohttp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gohttp.StartMockServer()

	os.Exit(m.Run())
}

func TestGetEndpointsErrorGettingFromApi(t *testing.T) {
	// Initialization:
	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method: http.MethodGet,
		Url:    "https://api.github.com",
		Error:  errors.New("timeout getting the endpoints"),
	})

	// Execution:
	endpoints, err := GetEndpoints()

	// Validation:

	assert.Nil(t, endpoints)
	assert.NotNil(t, err)
	assert.EqualValues(t, "timeout getting the endpoints", err.Error())
}

func TestGetEndpointsNotFound(t *testing.T) {
	// Initialization:
	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		ResponseStatusCode: http.StatusNotFound,
		ResponseBody:       `{"message": "endpoint not found"}`,
	})

	// Execution:
	endpoints, err := GetEndpoints()

	// Validation:
	assert.Nil(t, endpoints)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error when trying to fetch github endpoints", err.Error())
}

func TestGetEndpointsInvalidJsonResponse(t *testing.T) {
	// Initialization:
	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		ResponseStatusCode: http.StatusOK,
		ResponseBody:       `{"events_url": `,
	})

	// Execution:
	endpoints, err := GetEndpoints()

	// Validation:

	assert.Nil(t, endpoints)
	assert.NotNil(t, err)
	assert.EqualValues(t, "unexpected end of JSON input", err.Error())
}

func TestGetEndpointsNoError(t *testing.T) {
	// Initialization:
	gohttp.FlushMocks()
	gohttp.AddMock(gohttp.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		ResponseStatusCode: http.StatusOK,
		ResponseBody:       `{"events_url": "https://api.github.com/events"}`,
	})

	// Execution:
	endpoints, err := GetEndpoints()

	// Validation:

	assert.Nil(t, err)
	assert.NotNil(t, endpoints)
	assert.EqualValues(t, "https://api.github.com/events", endpoints.EventsUrl)
}
