package app

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMappings(t *testing.T) {

	assert.EqualValues(t, 0, len(router.Routes()))

	mapUrls()

	routes := router.Routes()

	assert.EqualValues(t, 2, len(routes))

	assert.EqualValues(t, http.MethodGet, routes[0].Method)
	assert.EqualValues(t, "/users/:id", routes[0].Path)

	assert.EqualValues(t, http.MethodPost, routes[1].Method)
	assert.EqualValues(t, "/users", routes[1].Path)
}
