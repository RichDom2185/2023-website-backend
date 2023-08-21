package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RichDom2185/2023-website-backend/bot"
	"github.com/RichDom2185/2023-website-backend/messages"
	"github.com/RichDom2185/2023-website-backend/params"
	tgbot "github.com/go-telegram/bot"
)

func HandleMessages(w http.ResponseWriter, r *http.Request) {
	var requestParams params.MessagePostRequest
	err := json.NewDecoder(r.Body).Decode(&requestParams)
	if err != nil {
		log.Fatalln(err)
	}

	err = requestParams.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request: %s", err)
		return
	}

	bot, err := bot.GetBotFrom(r)
	if err != nil {
		log.Fatalln(err)
	}

	chatID, ok := os.LookupEnv("TG_BOT_TARGET_CHAT_ID")
	if !ok {
		log.Fatalln("Cannot find chat ID in environment.")
	}

	msg := messages.ExternalMessageFrom(requestParams)
	bot.SendMessage(r.Context(), &tgbot.SendMessageParams{
		ChatID: chatID,
		Text:   msg.String(),
	})

	fmt.Fprintf(w, "Your message has been sent!")
}
