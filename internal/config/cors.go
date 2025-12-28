package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type CorsConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

func LoadCorsConfig() CorsConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	origins := os.Getenv("CORS_ALLOWED_ORIGINS")

	allowedOrigins := []string{}
	if origins != "" {
		allowedOrigins = strings.Split(origins, ",")
	}

	return CorsConfig{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"Content-Type", "Authorization",
		},
		AllowCredentials: true,
	}
}
