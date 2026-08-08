package config

import (
	"log"

	"github.com/joho/godotenv"
)

// * GetDotEnv - Responsável por importar o arquivo .env
func GetDotEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error on loading .env file")
	}
}
