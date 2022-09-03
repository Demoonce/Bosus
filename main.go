package main

import (
	"flag"
	"log"
	"os"
	"telega/cities"
	"telega/jokes"
	"telega/news"
	"telega/utils"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func processApp() {
	token := flag.String("token", "", "Bot token")
	flag.Parse()
	if *token == "" {
		os.Exit(1)
	}
	var err error
	utils.Api, err = tg.NewBotAPI(*token)
	if err != nil {
		utils.Logger.Fatalln(err)
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
	news.InitNews()
	for update := range updates {
		message := update.Message
		if message != nil {
			cities.RunCities(message)
			jokes.RunJokes(message)
			news.RunNews(message)
		}
	}
}
func main() {
	processApp()
}
