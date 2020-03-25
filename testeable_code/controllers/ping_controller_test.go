package controllers

import (
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"net/http"
	"github.com/federicoleon/golang-examples/testeable_code/services"
	"errors"
)

type pingServiceMock struct {
	handlePingFn func() (string, error)
}

func (mock pingServiceMock) HandlePing() (string, error) {
	return mock.handlePingFn()
}

func TestPingWithError(t *testing.T) {
	serviceMock := pingServiceMock{}
	serviceMock.handlePingFn = func() (string, error) {
		return "", errors.New("error executing ping")
	}
	services.PingService = serviceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	PingController.Ping(context)

	if response.Code != http.StatusInternalServerError {
		t.Error("response code should be 500")
	}

	if response.Body.String() != "error executing ping" {
		t.Error("response body should say 'error'")
	}
}

func TestPingNoError(t *testing.T) {
	serviceMock := pingServiceMock{}
	serviceMock.handlePingFn = func() (string, error) {
		return "pong", nil
	}
	services.PingService = serviceMock

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	PingController.Ping(context)

	if response.Code != http.StatusOK {
		t.Error("response code should be 200")
	}

	if response.Body.String() != "pong" {
		t.Error("response body should say 'pong'")
	}
}
