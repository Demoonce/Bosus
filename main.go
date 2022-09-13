package main

import (
	"log"
	"net/http"
	"os"

	"telega/cities"
	"telega/jokes"
	"telega/news"
	"telega/utils"
	"telega/weather"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func initEnv() {
	utils.Token = os.Getenv("TOKEN")
	if utils.Token == "" {
		log.Fatalln("Unable to get telegram bot token")
		return
	}
	weather.API_KEY = os.Getenv("API_KEY")
	if weather.API_KEY == "" {
		log.Fatalln("Unable to get api key for open weather")
	}
}

func processApp() {
	var err error

	utils.Api, err = tg.NewBotAPI(utils.Token)
	if err != nil {
		utils.Logger.Fatalln(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Bo$$ start"))
	})
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	u := tg.NewUpdate(-1)
	u.Timeout = 60
	// updates := utils.Api.ListenForWebhook("/" + utils.Token)
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
			weather.RunWeather(message)
		}
	}
}

func main() {
	initEnv()
	processApp()
}
