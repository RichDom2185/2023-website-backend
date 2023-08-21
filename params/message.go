package params

import (
	"errors"
)

type MessagePostRequest struct {
	Email            string `json:"email"`
	TelegramUsername string `json:"telegram"`
	Message          string `json:"message"`
}

func (p *MessagePostRequest) Validate() error {
	if p.Message == "" {
		return errors.New("Message cannot be empty")
	}
	return nil
}
