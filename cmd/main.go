package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eron97/login-authenticator.git/src/config/connect_db"
	"github.com/eron97/login-authenticator.git/src/controller"
	"github.com/eron97/login-authenticator.git/src/controller/routes"
	"github.com/eron97/login-authenticator.git/src/repository"
	"github.com/eron97/login-authenticator.git/src/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	cfg := connect_db.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}

	db, err := connect_db.NewConnectionDB(cfg)
	if err != nil {
		fmt.Println("Erro ao conectar ao MySQL:", err)
		return
	}

	userController := initDependencies(db)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	router.Run(":8080")

}

func initDependencies(
	database *gorm.DB,
) controller.ControllerInterface {
	repo := repository.NewDatabase(database)
	service := service.NewDomainService(repo)
	return controller.NewControllerInterface(service)
}
