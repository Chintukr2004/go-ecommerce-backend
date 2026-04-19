package routes

import (
	"github.com/Chintukr2004/go-ecommerce-backend/internal/handlers"
	"github.com/Chintukr2004/go-ecommerce-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
	}

	api.GET("profile", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "protected profiles routes",
		})
	},
	)
}
