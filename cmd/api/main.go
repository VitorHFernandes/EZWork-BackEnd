package main

import (
	"log"
	"os"

	"github.com/VitorHFernandes/EZWork-BackEnd/internal/auth"
	"github.com/VitorHFernandes/EZWork-BackEnd/internal/auth/sessions"
	"github.com/VitorHFernandes/EZWork-BackEnd/internal/config"
	"github.com/VitorHFernandes/EZWork-BackEnd/internal/db"
	"github.com/VitorHFernandes/EZWork-BackEnd/internal/router"
	todolist "github.com/VitorHFernandes/EZWork-BackEnd/internal/todo-list"
)

func main() {
	config.GetDotEnv()

	database, err := db.NewMySQL(os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	//* Métodos de autenticação e sessão.
	authRepository := auth.NewMySQLAuthRepository(database)
	sessionsRepository := sessions.NewMySQLSessionRepository(database)

	authService := auth.NewAuthService(authRepository, sessionsRepository)
	authHandler := auth.NewHandler(authService)

	//* Todo List
	todoRepository := todolist.NewMySQLTodoRepository(database)
	todoService := todolist.NewTodoService(todoRepository)
	todoHandler := todolist.NewHandler(todoService)

	r := router.New(
		todoHandler,
		authHandler,
	)

	if err := r.Run(os.Getenv("API_PORT")); err != nil {
		log.Fatal(err)
	}
}
