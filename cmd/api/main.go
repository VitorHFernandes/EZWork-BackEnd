package main

import (
	"log"
	"os"

	"github.com/VitorHFernandes/EZWork-BackEnd/internal/config"
	"github.com/VitorHFernandes/EZWork-BackEnd/internal/db"
	"github.com/VitorHFernandes/EZWork-BackEnd/internal/router"
)

func main() {
	config.GetDotEnv()

	database, err := db.NewMySQL(os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	r := router.New()

	if err := r.Run(os.Getenv("API_PORT")); err != nil {
		log.Fatal(err)
	}
}
