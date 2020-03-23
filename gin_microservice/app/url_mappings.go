package app

import (
	"github.com/federicoleon/golang-examples/gin_microservice/controllers"
)

func mapUrls() {
	router.POST("/users", controllers.UsersController.Create)
	router.GET("/users/:id", controllers.UsersController.Get)
}
