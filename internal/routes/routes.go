package routes

import (
	"github.com/Chintukr2004/go-ecommerce-backend/internal/handlers"
	"github.com/Chintukr2004/go-ecommerce-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler, productHandler *handlers.ProductHandler, cartHandler *handlers.CartHandler) {

	api := r.Group("/api/v1")

	// auth routes
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
	// product routes
	api.POST("/products", productHandler.Create)
	api.GET("/products", productHandler.GetAll)
	api.GET("/products/:id", productHandler.GetByID)
	api.PUT("/products/:id", productHandler.Update)
	api.DELETE("/products/:id", productHandler.Delete)

	// cart routes
	cart := api.Group("/cart")
	cart.Use(middleware.AuthMiddleware())
	{
		cart.POST("/add", cartHandler.Add)
		cart.GET("", cartHandler.Get)
	}
}
