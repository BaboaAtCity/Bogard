package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleStartCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	start := tgbotapi.NewMessage(msg.Chat.ID, "Bogard is up and running!")
	start.ReplyToMessageID = msg.MessageID
	bot.Send(start)
}

func HandleHelpCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	helpText := "Current commands list \n/start - just a test command to check if Bogard is awake\n/help - to check what commands are available\n/gpt (msg) - this will send a message to gpt-4o-mini\n/scrape - this is currently being worked on"
	help := tgbotapi.NewMessage(msg.Chat.ID, helpText)
	help.ReplyToMessageID = msg.MessageID
	bot.Send(help)
}
