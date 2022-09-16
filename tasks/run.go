package tasks

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/PuerkitoBio/goquery"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telega/utils"
)

// Authorizes glazov gov site with the given username and password. Returns cookie of moodle id/session
func authorize(username string, password string) {
	resp, err := Client.PostForm("http://is.glazov-gov.ru/login/index.php", url.Values{
		"username": {username},
		"password": {password},
	})
	if err != nil {
		log.Fatalln("Can't authorize:", err)
	}
	Client.Jar.SetCookies(&url.URL{
		Scheme: "http",
		Host:   "is.glazov-gov.ru",
		Path:   "/",
	}, []*http.Cookie{
		{
			Unparsed: []string{resp.Header.Get("Set-Cookie")},
		},
	})
}

// Gets the main page of the glazov gov site
func InitTasks() *goquery.Document {
	authorize(Username, Password)
	req, err := http.NewRequest("GET", "http://is.glazov-gov.ru/my", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", UserAgent)
	resp, err := Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	document, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
	}
	return document
}

func RunTasks(message *tg.Message, tasks_wg *sync.WaitGroup) {
	if message.Command() == "distant" {
		if AlreadyStarted {
			utils.ReplyTo(message, "Не так часто")
			return
		}
		utils.ReplyTo(message, "Подождите...")
		GetTasks(CourseDocument, tasks_wg) // initializes course slice
		tasks_wg.Wait()
		for _, course := range Courses {
			for _, task := range course.Tasks {
				GetTaskContent(task) // initializes task slices for each course
			}
			var message_text string
			if course.Name == "Человек, природа, мир 11 класс" {
				message_text = fmt.Sprintf("%s\n\n%s\n%s", course.Name, course.Tasks[len(course.Tasks)-2].Name, course.Tasks[len(course.Tasks)-2].Contents)
			} else {
				message_text = fmt.Sprintf("%s\n\n%s\n%s", course.Name, course.Tasks[len(course.Tasks)-1].Name, course.Tasks[len(course.Tasks)-1].Contents)
			}
			reply := tg.NewMessage(message.Chat.ID, message_text)
			reply.ParseMode = "HTML"
			utils.Api.Send(reply)
		}
		AlreadyStarted = false
		Courses = make([]*Course, 0)
	}
}

func Cache(file string) {
	// TODO
}
