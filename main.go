package main

import (
	"log"
	"os"
	"payment-service/config"
	"payment-service/middleware"
	"payment-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// 在所有其他init之前加载.env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		log.Println("[Payment Service] JWT_SECRET loaded:", secret)
	} else {
		log.Println("[Payment Service] WARNING: JWT_SECRET not set!")
	}
}

func main() {
	// 初始化JWT
	middleware.InitJWT()
	
	config.InitDB()

	r := gin.Default()
	r.Use(corsMiddleware())

	routes.SetupRoutes(r)

	log.Println("Payment Service starting on :8081")
	r.Run(":8081")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
