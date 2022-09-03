package jokes

import (
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"telega/utils"
	"time"

	"github.com/PuerkitoBio/goquery"
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
	joke_index, err := rand.Int(rand.Reader, big.NewInt(int64(len(Jokes)-1)))
	if err != nil {
		utils.Logger.Println(err)
	}
	msg := tg.NewMessage(chat_id, Jokes[joke_index.Int64()])
	utils.Api.Send(msg)
}

//Parsing functions

func getJokesPage(url string) {
	resp, err := http.Get(url)
	if err != nil {
		utils.Logger.Println(resp.StatusCode, err)
		return
	}
	jokes, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		utils.Logger.Println(err)
		return
	}
	jokes.Find("div.text").Each(func(i int, s *goquery.Selection) {
		Jokes = append(Jokes, s.Text())
	})
}

func ParseJokes() {
	for a := 1; a < 500; a++ {
		time.Sleep(time.Millisecond * 10)
		go getJokesPage(fmt.Sprintf("https://nekdo.ru/page/%d", a))
	}
loop:
	for {
		select {
		case <-time.NewTimer(time.Second * 5).C:
			close(JokesChan)
			break loop
		case a := <-JokesChan:
			Jokes = append(Jokes, a)
		}
	}
}

func WriteJokes() {
	file, err := os.Create("jokes.gob")
	defer file.Close()
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	encoder := gob.NewEncoder(file)
	encoder.Encode(Jokes)
}
