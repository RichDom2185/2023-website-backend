package bot

import (
	"context"

	"github.com/go-telegram/bot"
)

func Setup(ctx context.Context, token string) (*bot.Bot, error) {
	b, err := bot.New(token)
	if err != nil {
		return nil, err
	}

	return b, nil
}
