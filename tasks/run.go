package tasks

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"

	"telega/utils"

	"github.com/PuerkitoBio/goquery"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Authorizes glazov gov site with the given username and password. Set cookie of moodle id/session to the jar
func authorize(username string, password string) {
	resp, err := Client.PostForm("http://is.glazov-gov.ru/login/index.php", url.Values{
		"username": {username},
		"password": {password},
	})
	if err != nil {
		return
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
		log.Println(err)
		return nil
	}
	req.Header.Add("User-Agent", UserAgent)
	resp, err := Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	document, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Println(err)
		return nil
	}
	return document
}

func RunTasks(message *tg.Message, tasks_wg *sync.WaitGroup) {
	if message.Command() == "distant" {
		utils.ReplyTo(message, "Подождите...")
		GetTasks(CourseDocument, tasks_wg) // initializes course slice
		tasks_wg.Wait()

		for _, course := range Courses {
			for _, task := range course.Tasks {
				GetTaskContent(task) // initializes task slices for each course
			}
			if len(course.Tasks) == 0 {
				Courses = make([]*Course, 0)
				return
			}
			last_task := course.Tasks[len(course.Tasks)-1]
			message_text := fmt.Sprintf("%s\n\n%s\n%s", course.Name, last_task.Name, last_task.Contents)
			reply := tg.NewMessage(message.Chat.ID, message_text)
			reply.ParseMode = "HTML"
			utils.Api.Send(reply)

			for _, a := range last_task.Documents {
				doc := tg.NewDocument(message.Chat.ID, a)
				utils.Api.Send(doc)
			}

		}
		Courses = make([]*Course, 0)
	}
}

func Cache(file string) {
	// TODO
}
