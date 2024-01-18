package controller

import (
	"net/http"

	"github.com/eron97/login-authenticator.git/src/controller/request_validate"
	"github.com/eron97/login-authenticator.git/src/models"
	"github.com/eron97/login-authenticator.git/src/service"
	"github.com/gin-gonic/gin"
)

type ControllerInterface interface {
	ReadlAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type useControllerInterface struct {
	service service.DomainService
}

func NewControllerInterface(
	service service.DomainService,
) ControllerInterface {
	return &useControllerInterface{
		service: service,
	}
}

func (controller *useControllerInterface) CreateUser(c *gin.Context) {
	var request models.CreateUser

	err := request_validate.ValidateRequest(c, &request)
	if err != nil {
		return
	}

	resp := controller.service.CreateUser(request)
	c.JSON(http.StatusOK, gin.H{"message": resp})
}

func (controller *useControllerInterface) ReadlAllUsers(c *gin.Context) {

	resp, err := controller.service.ReadAllUsers(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (controller *useControllerInterface) LoginUser(c *gin.Context) {
	var request models.LoginUser

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possível ler os dados de requisição"})
		return
	}

	resp := controller.service.LoginUser(request)
	c.JSON(http.StatusOK, gin.H{"message": resp})

}
