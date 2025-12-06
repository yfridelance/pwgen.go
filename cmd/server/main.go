package main

import (
	"log"
	"pwgen/internal/api"
	"pwgen/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Configuration
	cfg := config.LoadConfig()

	// Set Gin Mode
	gin.SetMode(gin.ReleaseMode)
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	// Setup Router
	r := gin.Default()
	
	// CORS Middleware (Simple version)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	handler := api.NewHandler(cfg)

	// Routes
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/generate", handler.Generate)
	}

	log.Printf("Starting server on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
