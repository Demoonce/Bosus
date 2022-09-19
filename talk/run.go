package talk

import (
	"strings"

	"telega/utils"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitTalk() {
	Messages = ParseMessageFile("messages.json")
	for _, a := range Messages {
		Chain.Add([]string{a})
	}
}

func RunTalk(message *tg.Message) {
	if !message.IsCommand() {
		Chain.Add(strings.Split(message.Text, " "))
	} else if message.Command() == "g" {
		utils.ReplyTo(message, GenerateMsg())
	}
}
