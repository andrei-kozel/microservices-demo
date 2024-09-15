package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}

	return port
}

func getEnvironmentValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}

	return os.Getenv(key)
}
