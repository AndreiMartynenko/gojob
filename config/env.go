package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
		return
	}
	log.Println(".env file loaded")
}

// Read the config
type DatabaseConfig struct {
	url string
}

// func to read the config

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		// from .env
		url: os.Getenv("DATABASE_URL"),
	}
}
