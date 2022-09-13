package utils

import (
	"log"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ReplyTo(message *tg.Message, text string) {
	reply := tg.NewMessage(message.Chat.ID, text)
	reply.ReplyToMessageID = message.MessageID
	_, err := Api.Send(reply)
	if err != nil {
		log.Println("Can't reply to message:", err)
	}
}

// Makes the first letter of the word uppercase and others lowercase
func Capitalize(str string) string {
	if len(str) > 1 {
		str := strings.ToLower(str)
		return strings.Title(str)
	} else {
		return str
	}
}

func TrimLower(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}

func Log(msg string, err error) {
	if err != nil {
		Logger.Println(msg)
	}
}

func Fatal(msg string, err error) {
	if err != nil {
		Logger.Fatalln(msg)
	}
}
