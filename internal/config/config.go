package config

import (
	"os"
	"strconv"
	"strings"
)

// Config holds all configuration for the application
type Config struct {
	Port                string
	MultiGen            bool
	DefaultLength       int
	MaxLength           int
	HIBPEnabled         bool
	CustomWordListURL   string
	AllowedOrigins      []string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		Port:                getEnv("PORT", "8080"),
		MultiGen:            getEnvAsBool("MULTI_GEN", false),
		DefaultLength:       getEnvAsInt("DEFAULT_LENGTH", 16),
		MaxLength:           256, // Hard limit from requirements
		HIBPEnabled:         getEnvAsBool("HIBP_ENABLED", true),
		CustomWordListURL:   getEnv("CUSTOM_WORDLIST_URL", ""),
		AllowedOrigins:      strings.Split(getEnv("ALLOWED_ORIGINS", "*"), ","),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	strValue := getEnv(key, "")
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	strValue := getEnv(key, "")
	if value, err := strconv.ParseBool(strValue); err == nil {
		return value
	}
	return fallback
}
