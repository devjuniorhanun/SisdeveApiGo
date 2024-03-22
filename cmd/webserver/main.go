package main

import (
	"fmt"
	"log/slog" // Pacote de Geração de Logs
	"net/http"

	env "github.com/devjuniorhanun/SisdeveApiGo/config/env"
	"github.com/devjuniorhanun/SisdeveApiGo/config/logger" // Pacote de configuração do Log
	"github.com/devjuniorhanun/SisdeveApiGo/internal/database"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/database/sqlc"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/handler"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/handler/routes"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/repository/categoryrepository"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/repository/productrepository"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/repository/userrepository"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/service/categoryservice"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/service/productservice"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/service/userservice"
	"github.com/go-chi/chi"
)

func main() {
	logger.InitLogger()
	slog.Info("starting api")

	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failed to load environment variables", err, slog.String("package", "main"))
		return
	}
	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
		return
	}

	queries := sqlc.New(dbConnection)

	// user
	userRepo := userrepository.NewUserRepository(dbConnection, queries)
	newUserService := userservice.NewUserService(userRepo)

	// category
	categoryRepo := categoryrepository.NewCategoryRepository(dbConnection, queries)
	newCategoryService := categoryservice.NewCategoryService(categoryRepo)

	// product
	productRepo := productrepository.NewProductRepository(dbConnection, queries)
	productsService := productservice.NewProductService(productRepo)

	newHandler := handler.NewHandler(newUserService, newCategoryService, productsService)

	// init routes
	router := chi.NewRouter()
	routes.InitRoutes(router, newHandler)
	routes.InitDocsRoutes(router)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}
}
