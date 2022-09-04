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

func processApp() {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatalln("Unable to get token")
	}
	var err error
	utils.Api, err = tg.NewBotAPI(token)
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

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Bo$$ start"))
	})
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	u := tg.NewUpdate(-1)
	u.Timeout = 60
	updates := utils.Api.ListenForWebhook("/" + token)
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
	processApp()
}
