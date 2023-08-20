package router

import (
	"github.com/RichDom2185/2023-website-backend/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Setup() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	setupRoutes(r)
	return r
}

func setupRoutes(r chi.Router) {
	r.Get("/", handlers.HandleHealthCheck)
	r.Post("/resume", handlers.HandleResumeForm)
}
