package main

import (
	"awesomeProject/internal/routes"
	"awesomeProject/pkg/database"
	"log"

	_ "awesomeProject/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Awesome Project API
// @version 1.0
// @description This is a sample server for Awesome Project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
// @BasePath /api/v1
func main() {
	config := database.NewConfig()
	db, err := database.Connect(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()

	appRoutes := routes.NewRoutes(router, db)
	appRoutes.Setup()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run("localhost:9090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
