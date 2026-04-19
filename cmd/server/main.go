package main

import (
	"log"
	"os"

	"github.com/Chintukr2004/go-ecommerce-backend/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db := config.ConnectDB()
	defer db.Close()

	config.RunMigrations(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "server is runnig",
		})
	})

	log.Println("server started on port ", port)

	r.Run(":" + port)

}
