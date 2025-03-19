package telegram

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/svnaditya/telegram-x-bot/config"
	"github.com/svnaditya/telegram-x-bot/llm"
	"github.com/svnaditya/telegram-x-bot/x"
)

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(config.Cfg.TgToken)
	if err != nil {
		log.Panic(err)
	}
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command"))
			continue
		}

		if strings.HasPrefix(update.Message.Text, "/post") {
			content := strings.TrimPrefix(update.Message.Text, "/post")
			response := llm.GeneratePost(content)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			bot.Send(msg)

			// Wait for user confirmation
			confirmation := tgbotapi.NewMessage(update.Message.Chat.ID, "Do you want to post this on Twitter? (yes/no)")
			bot.Send(confirmation)

			// Handle user response
			confirmUpdate := <-updates
			if confirmUpdate.Message.Text == "yes" || confirmUpdate.Message.Text == "Yes" {
				x.Post()
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Posted to Twitter!"))
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Post canceled."))
			}
		}
	}
}
