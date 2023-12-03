package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(1)
	}
}

func GetEnvByKey(key string) string {
	return os.Getenv(key)
}
