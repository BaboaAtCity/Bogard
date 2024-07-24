package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleScrapeCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	url := "https://www.uniqlo.com/on/demandware.store/Sites-GB-Site/en_GB/Product-GetVariants"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		handleError(bot, msg, err)
		return
	}

	// Add query parameters
	q := req.URL.Query()
	q.Add("pid", "473495")
	q.Add("Quantity", "1")
	req.URL.RawQuery = q.Encode()

	// Add headers
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://www.uniqlo.com/uk/en/product/tv-anime-one-piece-25th-ut-graphic-t-shirt-473495.html?dwvar_473495_size=SMA003&dwvar_473495_color=COL00")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		handleError(bot, msg, err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		handleError(bot, msg, err)
		return
	}

	fmt.Println("Raw response:", string(body))

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		handleError(bot, msg, err)
		return
	}

	// Process the result and create a response message
	response := processResult(result)

	reply := tgbotapi.NewMessage(msg.Chat.ID, response)
	reply.ReplyToMessageID = msg.MessageID
	bot.Send(reply)
}

func handleError(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, err error) {
	errorMsg := fmt.Sprintf("An error occurred: %v", err)
	reply := tgbotapi.NewMessage(msg.Chat.ID, errorMsg)
	reply.ReplyToMessageID = msg.MessageID
	bot.Send(reply)
}

func processResult(result map[string]interface{}) string {
	var response strings.Builder

	for key, value := range result {
		if strings.HasPrefix(key, "color-") && strings.Contains(key, "|size-") {
			if productInfo, ok := value.(map[string]interface{}); ok {
				if attributes, ok := productInfo["attributes"].(map[string]interface{}); ok {
					color := attributes["color"].(string)
					size := attributes["size"].(string)

					if availability, ok := productInfo["availability"].(map[string]interface{}); ok {
						available := availability["available"].(bool)
						quantity := int(availability["currentQty"].(float64))

						status := "In stock"
						if !available {
							status = "Out of stock"
						}

						response.WriteString(fmt.Sprintf("Color: %s, Size: %s - %s (Quantity: %d)\n", color, size, status, quantity))
					}
				}
			}
		}
	}

	if response.Len() == 0 {
		return "Unable to retrieve product information."
	}

	return response.String()
}
