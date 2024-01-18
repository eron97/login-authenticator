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

	connectionProvider := connect_db.NewConnectionProviderFromEnv(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	db, err := connectionProvider.NewConnectionDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao database:", err)
		return
	}

	userController := initDependencies(db)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	router.Run(":8080")

}

func initDependencies(database *gorm.DB) controller.ControllerInterface {
	repo := repository.NewDatabase(database)
	service := service.NewDomainService(repo)
	return controller.NewControllerInterface(service)
}
