package app

import (
	"github.com/federicoleon/golang-examples/code_structure/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:id", controllers.UsersController.Get)
	router.POST("/users", controllers.UsersController.Save)
}
