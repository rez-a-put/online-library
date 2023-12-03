package utils

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func init() {
	projectDirName := "online-library"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(1)
	}
}

func GetEnvByKey(key string) string {
	return os.Getenv(key)
}
