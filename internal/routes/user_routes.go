package routes

import (
	"awesomeProject/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func (r *Routes) setupUserRoutes(rg *gin.RouterGroup) {
	userHandler := handlers.UserHandler{DB: r.db}

	users := rg.Group("/users")
	{
		users.GET("/", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUser)
		users.POST("/", userHandler.CreateUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
}
