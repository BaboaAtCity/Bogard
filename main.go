package main

import (
	"Bogard/handlers"
	"Bogard/utils"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	utils.LoadEnv()

	botToken := utils.GetEnv("TELEGRAM_BOT_TOKEN")
	whitelist := utils.GetEnv("WHITELIST_ID")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if handlers.IsAuth(update.Message.From.ID, whitelist) {
				handlers.HandleCommand(bot, update.Message)
			}
		}
	}
}
