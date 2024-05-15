package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariabels() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
