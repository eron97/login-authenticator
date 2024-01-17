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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro de validação"})
		return
	}

	resp := controller.service.CreateUser(request)
	c.JSON(http.StatusOK, gin.H{"message": resp})
}

func (pkg *useControllerInterface) ReadlAllUsers(c *gin.Context) {

	resp, err := pkg.service.ReadAllUsers(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, resp)
}
