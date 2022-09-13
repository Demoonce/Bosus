package jokes

import (
	"os"
	// "time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitJokes() {
	if _, err := os.Stat("jokes.gob"); err != nil {
		ParseJokes()
		WriteJokes()
	} else {
		Jokes = DecodeJokes()
	}
}

func RunJokes(message *tg.Message) {
	switch message.Command() {
	case "joke":
		SendJoke(message.Chat.ID)
	case "parse":
		ParsingMutex.Lock()
		ParseJokes()
		ParsingMutex.Unlock()
	}
}
