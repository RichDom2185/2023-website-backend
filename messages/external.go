package messages

import (
	"fmt"

	"github.com/RichDom2185/2023-website-backend/params"
)

// Messages that come from the website and
// are to be forwarded through telegram
type ExternalMessage struct {
	Email            string
	TelegramUsername string
	Text             string
}

const (
	newExternalMessageTemplate = `
New message from %s:
-------------------
%s`
)

func (e ExternalMessage) String() string {
	var sender string
	switch {
	case e.Email == "" && e.TelegramUsername == "":
		sender = "anonymous"
	case e.Email == "" && e.TelegramUsername != "":
		sender = formatUsername(e.TelegramUsername)
	case e.Email != "" && e.TelegramUsername == "":
		sender = e.Email
	case e.Email != "" && e.TelegramUsername != "":
		sender = fmt.Sprintf("%s (%s)", e.Email, formatUsername(e.TelegramUsername))
	}
	return fmt.Sprintf(newExternalMessageTemplate, sender, e.Text)
}

func ExternalMessageFrom(m params.MessagePostRequest) ExternalMessage {
	return ExternalMessage{
		Email:            m.Email,
		TelegramUsername: m.TelegramUsername,
		Text:             m.Message,
	}
}
