package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewRoutes(router *gin.Engine, db *gorm.DB) *Routes {
	return &Routes{
		router: router,
		db:     db,
	}
}

func (r *Routes) Setup() {
	v1 := r.router.Group("/api/v1")
	{
		r.setupUserRoutes(v1)
	}
}
