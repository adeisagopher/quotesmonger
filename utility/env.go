package utility

import (
	"log"

	"github.com/joho/godotenv"
)

// Loads in .env variables to be accessed with "godotenv".
func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Unable to load .env file")
	}
}
