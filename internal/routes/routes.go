package routes

import (
	"awesomeProject/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initialize(router *gin.Engine, db *gorm.DB) {
	userhandler := handlers.UserHandler{DB: db}
	router.GET("/users", userhandler.GetUsers)
	router.GET("/users/:id", userhandler.GetUserByID)
	router.POST("/users", userhandler.PostUsers)
	router.PUT("/users/:id", userhandler.UpdateUser)
}
