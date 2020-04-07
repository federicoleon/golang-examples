package controllers

import (
	"github.com/federicoleon/golang-examples/code_structure/mvc/domain"
	"github.com/federicoleon/golang-examples/code_structure/mvc/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	UsersController = usersController{}
)

type usersController struct{}

func (controller usersController) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user id"})
		return
	}

	user, err := services.UsersService.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller usersController) Save(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid json body"})
		return
	}

	if err := services.UsersService.Save(&user); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
