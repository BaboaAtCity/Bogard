package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleScrapeCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {

	response := "Scraping not yet implemented"
	reply := tgbotapi.NewMessage(msg.Chat.ID, response)
	reply.ReplyToMessageID = msg.MessageID
	bot.Send(reply)
}
