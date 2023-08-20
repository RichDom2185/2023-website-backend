package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/RichDom2185/2023-website-backend/bot"
	"github.com/RichDom2185/2023-website-backend/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	botToken, ok := os.LookupEnv("TG_BOT_TOKEN")
	if !ok {
		log.Fatalln("TG_BOT_TOKEN not found in environment.")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.Setup(ctx, botToken)
	if err != nil {
		log.Fatalln(err)
	}

	r := router.Setup(bot.MakeMiddlewareFrom(b))

	appMode := os.Getenv("GO_ENV")
	if appMode == "" {
		appMode = "production"
	}
	fmt.Printf("Server started in %s mode!\n", appMode)

	host := ""
	if appMode == "development" {
		host = "127.0.0.1"
	}

	err = http.ListenAndServe(host+":4000", r)
	if err != nil {
		log.Fatalln(err)
	}
}
