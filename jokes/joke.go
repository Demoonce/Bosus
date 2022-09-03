package jokes

import (
	"encoding/gob"
	"math/rand"
	"os"
	"telega/utils"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DecodeJokes() []string {
	file, err := os.Open("jokes.gob")
	jokes := make([]string, 0)
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&jokes)
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	return jokes
}

func SendJoke(chat_id int64) {
	rand.Seed(time.Now().Unix())
	joke := Jokes[rand.Intn(len(Jokes)-1)]
	msg := tg.NewMessage(chat_id, joke)
	utils.Api.Send(msg)
}
