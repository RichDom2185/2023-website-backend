package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RichDom2185/2023-website-backend/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := router.Setup()

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
