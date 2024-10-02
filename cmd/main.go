package main

import (
	"awesomeProject/internal/routes"
	"awesomeProject/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config := database.NewConfig()
	db, err := database.Connect(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := gin.Default()
	routes.Initialize(router, db)

	if err := router.Run("localhost:9090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
