package jokes

import (
	"encoding/gob"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"telega/utils"
	"time"
)

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
