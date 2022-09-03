package jokes

import (
	"os"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitJokes() {
	if _, err := os.Stat("jokes.gob"); err == nil {
		ParseJokes()
		WriteJokes()
	} else {
		Jokes = DecodeJokes()
	}
}

func RunJokes(message *tg.Message) {
	if message.Command() == "joke" {
		SendJoke(message.Chat.ID)
		return
	}
	if !JokesStarted {
		go func() {
			SendJoke(message.Chat.ID)
			jokes_ticker := time.NewTicker(time.Minute * 1).C
			for {
				select {
				case <-jokes_ticker:
					SendJoke(message.Chat.ID)
				default:
					time.Sleep(time.Minute * 2)
				}
			}
		}()
		JokesStarted = true
	}
}
