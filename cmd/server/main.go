package main

import (
	"log"
	"os"

	"github.com/Chintukr2004/go-ecommerce-backend/internal/config"
	"github.com/Chintukr2004/go-ecommerce-backend/internal/handlers"
	"github.com/Chintukr2004/go-ecommerce-backend/internal/repository"
	"github.com/Chintukr2004/go-ecommerce-backend/internal/routes"
	"github.com/Chintukr2004/go-ecommerce-backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db := config.ConnectDB()
	defer db.Close()

	config.RunMigrations(db)

	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserhandler(userService)

	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	cartRepo := repository.NewCartRepository(db)
	cartService := services.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService)

	routes.SetupRoutes(r, userHandler, productHandler, cartHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("server started on port ", port)

	r.Run(":" + port)

}
