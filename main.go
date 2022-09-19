package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"telega/cities"
	"telega/jokes"
	"telega/news"
	"telega/talk"
	"telega/tasks"
	_ "telega/tasks"
	"telega/utils"
	"telega/weather"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TOKEN - bot token
// API_KEY - api key for open weather
// USERNAME - username for glazov gov
// PASSWORD - password for glazov gov
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

	tasks.Username = os.Getenv("USERNAME")
	if tasks.Username == "" {
		log.Fatalln("Unable to get username for distant website")
	}
	tasks.Password = os.Getenv("PASSWORD")
	if tasks.Password == "" {
		log.Fatalln("Unable to get password for distant website")
	}
}

func processApp() {
	var err error
	var tasks_wg sync.WaitGroup
	signal_chan := make(chan os.Signal)
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
	updates := utils.Api.ListenForWebhook("/" + utils.Token)
	// updates := utils.Api.GetUpdatesChan(u)
	cities.BotName = utils.Api.Self.UserName

	cities.InitCities()
	jokes.InitJokes()
	news.InitNews()
	tasks.CourseDocument = tasks.InitTasks()
	talk.InitTalk()

	signal.Notify(signal_chan, os.Interrupt)
	go func(sig_chan chan os.Signal) {
		<-sig_chan
		talk.SaveChain()
		os.Exit(0)
	}(signal_chan)
	for update := range updates {
		message := update.Message
		if message != nil {
			cities.RunCities(message)
			jokes.RunJokes(message)
			news.RunNews(message)
			weather.RunWeather(message)
			tasks.RunTasks(message, &tasks_wg)
			talk.RunTalk(message)
		}
	}
}

func main() {
	initEnv()
	processApp()
}
