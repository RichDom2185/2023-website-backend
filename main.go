package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/RichDom2185/2023-website-backend/bot"
	"github.com/RichDom2185/2023-website-backend/router"
	"github.com/RichDom2185/2023-website-backend/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot find .env file, skipping...", err)
		err = nil
	}

	botToken, ok := os.LookupEnv("TG_BOT_TOKEN")
	if !ok {
		log.Fatalln("TG_BOT_TOKEN not found in environment.")
	}

	// Set up shared context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Set up bot
	b, err := bot.Setup(ctx, botToken)
	if err != nil {
		log.Fatalln(err)
	}

	// Start bot
	log.Println("Starting bot")
	go func() {
		defer log.Println("Bot stopped.")
		b.Start(ctx)
	}()

	// Set up router
	r := router.Setup(bot.MakeMiddlewareFrom(b))

	appMode := os.Getenv("GO_ENV")
	if appMode == "" {
		appMode = "production"
	}

	host := ""
	if appMode == "development" {
		host = "127.0.0.1"
	}

	s := &server.Server{
		Server: &http.Server{
			Addr:    host + ":4000",
			Handler: r,
		},
	}

	// Start API server
	log.Printf("Starting server in %s mode\n", appMode)
	go func() {
		err = s.ListenAndServe()
		if err != nil {
			log.Fatalf("Unexpected error from ListenAndServe: %v\n", err)
		}
	}()
	go s.WaitForExitingSignal(ctx)

	// Wait for the interrupt signal to gracefully shut down
	<-ctx.Done()

	// Allow the goroutines to exit gracefully
	time.Sleep(2 * time.Second)
	log.Println("Exiting app")
}

// func runConcurrently(items ...func()) {
// 	wg := &sync.WaitGroup{}
// 	wg.Add(len(items))
// 	for _, item := range items {
// 		go func(f func()) {
// 			defer wg.Done()
// 			f()
// 		}(item)
// 	}
// 	wg.Wait()
// }
