package routes

import (
	"github.com/devjuniorhanun/SisdeveApiGo/config/env"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/handler"
	"github.com/devjuniorhanun/SisdeveApiGo/internal/handler/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func InitRoutes(router chi.Router, h handler.Handler) {
	router.Use(middleware.LoggerData)

	router.Post("/user", h.CreateUser)
	router.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(env.Env.TokenAuth))
		r.Use(jwtauth.Authenticator)

		//user routes
		r.Patch("/user", h.UpdateUser)
		r.Get("/user", h.GetUserByID)
		r.Delete("/user", h.DeleteUser)
		r.Get("/user/list-all", h.FindManyUsers)
		r.Patch("/user/password", h.UpdateUserPassword)

		// categories routes
		r.Post("/category", h.CreateCategory)

		// products routes
		r.Post("/product", h.CreateProduct)
		r.Patch("/product/{id}", h.UpdateProduct)
		r.Delete("/product/{id}", h.DeleteProduct)
		r.Get("/product", h.FindManyProducts)
	})
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
	})
}