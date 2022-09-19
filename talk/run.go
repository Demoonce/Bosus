package talk

import (
	"os"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	// "telega/jokes"
	"telega/jokes"
	"telega/utils"
)

func SaveChain() {
	file, err := os.Create("model.json")
	defer file.Close()
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	data, err := Chain.MarshalJSON()
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	file.Write(data)
}

func InitTalk() {
	jokes.ParseJokes()
	jokes.WriteJokes()
	Messages = jokes.Jokes
	data, err := os.ReadFile("model.json")
	if err != nil {
		utils.Logger.Println(err)
		for _, a := range Messages {
			Chain.Add(strings.Split(a, " "))
		}
	} else {
		err = Chain.UnmarshalJSON(data)
		if err != nil {
			utils.Logger.Println(err)
		}
	}
}

func RunTalk(message *tg.Message) {
	if !message.IsCommand() {
		Chain.Add(strings.Split(message.Text, " "))
	} else if message.Command() == "g" {
		msg := tg.NewMessage(message.Chat.ID, Predict())
		if msg.Text == "" {
			utils.Api.Send(tg.NewMessage(message.Chat.ID, "Мне нечего сказать"))
		}
		_, err := utils.Api.Send(msg)
		if err != nil {
			utils.Logger.Println(err)
		}

	}
}
