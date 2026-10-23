package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rahulreghupdm/tashub-project/internal/config"
	"github.com/rahulreghupdm/tashub-project/internal/database"
	"github.com/rahulreghupdm/tashub-project/internal/middleware"
	"github.com/rahulreghupdm/tashub-project/internal/routes"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Could not load .env file, using system environment variables")
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	if err := database.ConnectPostgres(cfg.DatabaseURL); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create tables
	if err := database.CreateUserTable(); err != nil {
		log.Fatalf("Failed to create user table: %v", err)
	}

	e := gin.New()

	// Public Routes
	auth := e.Group("/auth")
	routes.RegisterAuthRoutes(auth)

	// Protected Routes
	user := e.Group("/user")
	user.Use(middleware.AuthMiddleware())
	user.GET("/profile", func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.JSON(200, gin.H{"user_id": userID})
	})

	port := ":" + cfg.AppPort
	log.Printf("Starting server on port %s", port)
	e.Run(port)
}
