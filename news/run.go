package news

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"telega/utils"
)

func InitNews() {
	News = GetNews()
}

func RunNews(message *tg.Message) {
	if message.Command() == "news" {
		msg := tg.NewMessage(message.Chat.ID, strings.Join(News, "\n\n"))
		_, err := utils.Api.Send(msg)
		if err != nil {
			utils.Logger.Println("NEWS:", err)
		}
	}
}
