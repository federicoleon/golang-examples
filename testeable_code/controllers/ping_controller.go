package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/federicoleon/golang-examples/testeable_code/services"
)

var (
	PingController = pingController{}
)

type pingController struct{}

func (controller pingController) Ping(c *gin.Context) {
	result, err := services.PingService.HandlePing()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, result)
}
