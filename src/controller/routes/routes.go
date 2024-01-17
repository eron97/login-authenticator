package routes

import (
	"github.com/eron97/login-authenticator.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, controller controller.ControllerInterface) {
	r.GET("/getAllUsers", controller.ReadlAllUsers)
	r.POST("/createUser", controller.CreateUser)
}
