package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	if msg.IsCommand() {
		switch msg.Command() {
		case "gpt":
			HandleGPTCommand(bot, msg)
		case "scrape":
			HandleScrapeCommand(bot, msg)
		case "start":
			HandleStartCommand(bot, msg)
		case "help":
			HandleHelpCommand(bot, msg)
		default:
			msg := tgbotapi.NewMessage(msg.Chat.ID, "Unknown command")
			bot.Send(msg)
		}
	} else {
		defaultMsg := tgbotapi.NewMessage(msg.Chat.ID, "Use my commands, try use /help")
		defaultMsg.ReplyToMessageID = msg.MessageID
		bot.Send(defaultMsg)
	}
}
