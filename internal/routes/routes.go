package routes

import (
	"awesomeProject/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initialize(router *gin.Engine, db *gorm.DB) {
	handler := handlers.UserHandler{DB: db}
	router.GET("/users", handler.GetUsers)
	router.GET("/users/:id", handler.GetUserByID)
	router.POST("/users", handler.PostUsers)
	router.PUT("/users/:id", handler.UpdateUser)
}
