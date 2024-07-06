package configs

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds the configuration values from .env file
type Config struct {
	Port       string
	PublicHost string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

// initConfig loads environment variables and initializes the Config struct
func initConfig() Config {
	_ = godotenv.Load() // Load .env file if present, ignoring errors

	config := Config{
		Port:       os.Getenv("PORT"),
		PublicHost: os.Getenv("PUBLIC_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBAddress:  fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName:     os.Getenv("DB_NAME"),
	}

	validateConfig(config)

	return config
}

// validateConfig checks the Config struct for missing or invalid fields
func validateConfig(config Config) {
	missingFields := []string{}

	// Check for required fields
	if config.Port == "" {
		missingFields = append(missingFields, "PORT")
	}
	if config.PublicHost == "" {
		missingFields = append(missingFields, "PUBLIC_HOST")
	}
	if config.DBUser == "" {
		missingFields = append(missingFields, "DB_USER")
	}
	if config.DBPassword == "" {
		missingFields = append(missingFields, "DB_PASSWORD")
	}
	if config.DBAddress == ":0" || strings.Contains(config.DBAddress, ":") && (strings.Split(config.DBAddress, ":")[0] == "" || strings.Split(config.DBAddress, ":")[1] == "") {
		missingFields = append(missingFields, "DB_HOST or DB_PORT")
	}
	if config.DBName == "" {
		missingFields = append(missingFields, "DB_NAME")
	}

	// If there are missing fields, panic with a descriptive message
	if len(missingFields) > 0 {
		panic(fmt.Sprintf("Missing or invalid required environment variables: %s", strings.Join(missingFields, ", ")))
	}
}
