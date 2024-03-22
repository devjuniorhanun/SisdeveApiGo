package rotes

import (
	"github.com/devjuniorhanun/SisdeveApiGo/internal/handler/userhandler"
	"github.com/go-chi/chi"
)

// Função InitUserRoutes()
// Responsável por criar uma estância das rotas para os usuários
func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", h.CreateUser)
	})
}
