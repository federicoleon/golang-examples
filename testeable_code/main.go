package main

import (
	"github.com/gin-gonic/gin"
	"github.com/federicoleon/golang-examples/testeable_code/controllers"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/ping", controllers.PingController.Ping)

	router.Run(":8080")
}
