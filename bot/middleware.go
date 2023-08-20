package bot

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-telegram/bot"
)

const (
	botContextKey = "bot_context"
)

func MakeMiddlewareFrom(b *bot.Bot) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), botContextKey, b)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetBotFrom(r *http.Request) (*bot.Bot, error) {
	bot, ok := r.Context().Value(botContextKey).(*bot.Bot)
	if !ok {
		return nil, errors.New("Could not get bot from request context")
	}
	return bot, nil
}
