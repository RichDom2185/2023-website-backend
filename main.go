package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RichDom2185/2023-website-backend/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	}))
	r.Get("/", handlers.HandleHealthCheck)
	r.Get("/resume", handlers.HandleResumeForm)

	appMode := os.Getenv("GO_ENV")
	if appMode == "" {
		appMode = "production"
	}
	fmt.Printf("Server started in %s mode!\n", appMode)

	host := ""
	if appMode == "development" {
		host = "127.0.0.1"
	}

	err := http.ListenAndServe(host+":4000", r)
	if err != nil {
		log.Fatalln(err)
	}
}
