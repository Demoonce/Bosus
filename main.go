package main

import (
	"log"
	"os"
	"telega/cities"
	"telega/jokes"
	"telega/utils"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func processApp() {
	if utils.Err != nil {
		log.Fatalln(utils.Err)
	}
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	utils.Logger = log.New(file, "BOT: ", log.Flags())
	if err != nil {
		log.Fatalln("Can't start bot api", err)
	}

	u := tg.NewUpdate(-1)
	u.Timeout = 60
	updates := utils.Api.GetUpdatesChan(u)
	cities.BotName = utils.Api.Self.UserName

	cities.InitCities()
	jokes.InitJokes()
	for update := range updates {
		message := update.Message
		if message != nil {
			cities.RunCities(message)
			jokes.RunJokes(message)
		}
	}
}
func main() {
	processApp()
}
