package main

import (
	"fmt"
	"log/slog" // Pacote de Geração de Logs
	"net/http"

	env "github.com/devjuniorhanun/SisdeveApiGo/config"
	"github.com/devjuniorhanun/SisdeveApiGo/config/logger" // Pacote de configuração do Log
	"github.com/devjuniorhanun/SisdeveApiGo/internal/database"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/database/sqlc"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/handler/userhandler"
	userrepository "github.com/devjuniorhanun/SisdeveApiGo/internal/repository/userepository"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/service/userservice"
	"github.com/go-chi/chi"
)

func main() {
	// Iniciamos nosso logs
	logger.InitLogger()
	slog.Info("Iniciando api")

	// Carregamos nossas envs
	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failed to load environment variables", err, slog.String("package", "main"))
		return
	}
	//  Iniciamos a conexão com o banco de dados
	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
		return
	}

	// Iniciamos as Rotas
	router := chi.NewRouter()
	// Iniciamos as queries do sqlc
	queries := sqlc.New(dbConnection)

	// user
	userRepo := userrepository.NewUserRepository(dbConnection, queries)
	newUserService := userservice.NewUserService(userRepo)
	newUserHandler := userhandler.NewUserHandler(newUserService)

	// init routes
	routes.InitUserRoutes(router, newUserHandler)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}
}
