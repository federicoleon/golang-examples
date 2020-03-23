package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/federicoleon/golang-examples/gin_microservice/domain/httperrors"
	"net/http"
	"github.com/federicoleon/golang-examples/gin_microservice/domain/users"
	"github.com/federicoleon/golang-examples/gin_microservice/services"
)

var (
	UsersController = usersController{}
)

type usersController struct{}

func respond(c *gin.Context, isXml bool, httpCode int, body interface{}) {
	if isXml {
		c.XML(httpCode, body)
		return
	}
	c.JSON(httpCode, body)
}

func (controller usersController) Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		httpErr := httperrors.NewBadRequestError("invalid json body")
		c.JSON(httpErr.Code, httpErr)
		return
	}
	createdUser, err := services.UsersService.Create(user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	// return created user
	c.JSON(http.StatusCreated, createdUser)
}

func (controller usersController) Get(c *gin.Context) {
	isXml := c.GetHeader("Accept") == "application/xml"

	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httpErr := httperrors.NewBadRequestError("invalid user id")
		respond(c, isXml, httpErr.Code, httpErr)
		return
	}

	user, getErr := services.UsersService.Get(userId)
	if getErr != nil {
		respond(c, isXml, getErr.Code, getErr)
		return
	}
	respond(c, isXml, http.StatusOK, user)
}
