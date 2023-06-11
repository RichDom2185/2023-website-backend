package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	RESUME_PATH = "./resume-latest.pdf"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Health ok! Welcome to the API!")
	})

	r.Get("/resume", func(w http.ResponseWriter, r *http.Request) {
		resumeFile, err := os.ReadFile(RESUME_PATH)
		if err != nil {
			log.Fatalln(err)
		}

		w.Header().Add("Content-Type", "application/pdf")
		w.Write(resumeFile)
	})

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
