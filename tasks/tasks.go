package tasks

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"strings"
	"sync"

	"telega/utils"

	"github.com/PuerkitoBio/goquery"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Task struct {
	Name      string
	Link      string
	Documents []tg.FileBytes
	Contents  string
}

// sends request to the specified course, returns links of the tasks
func sendRequest(course *Course, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequest("GET", course.Link, nil)
	if err != nil {
		utils.Logger.Println(err)
	}
	req.Header.Add("User-Agent", UserAgent)
	resp, err := Client.Do(req)
	if err != nil {
		utils.Logger.Println(err)
	}
	document, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		utils.Logger.Println(err)
	}
	document.Find("li.section.main.section-summary.clearfix").Each(func(i int, s *goquery.Selection) {
		link := s.Find("a")
		if _, ok := link.Attr("onclick"); ok { // checks if link is news forum link
			return
		}
		href, _ := link.Attr("href")
		course.Tasks = append(course.Tasks, &Task{
			link.Text(),
			href,
			nil,
			"",
		})
	})
}

func GetTaskContent(task *Task) {
	req, err := http.NewRequest("GET", task.Link, nil)
	if err != nil {
		utils.Logger.Println(err)
	}
	req.Header.Add("User-Agent", UserAgent)

	resp, err := Client.Do(req)
	if err != nil {
		utils.Logger.Println(err)
	}

	document, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		utils.Logger.Println(err)
	}

	content_block := document.Find("div.single-section").Find("div.content")
	task.Contents += content_block.Find("div.summary").Text()
	content_block.Find("a").Each(func(i int, s *goquery.Selection) {
		if s != nil {
			href, ok := s.Attr("href")
			if ok && s.Text() != task.Name {

				resp, err := LinksClient.Get(href)
				if err != nil {
					log.Println(err)
					return
				}
				redirect_url, err := resp.Location()
				if redirect_url != nil {
					url := redirect_url.String()
					if strings.Contains(url, "pluginfile.php") {
						resp, err := Client.Get(url)
						if err != nil {
							log.Println(err)
							return
						}
						_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
						if err != nil {
							log.Println(err)
						}
						filename := params["filename"]
						file_data, err := io.ReadAll(resp.Body)
						if err != nil {
							log.Println(err)
						}

						task.Documents = append(task.Documents, tg.FileBytes{
							Name:  filename,
							Bytes: file_data,
						})
					}
					href = redirect_url.String()

				}
				task.Contents += "\n" + fmt.Sprintf("<a href='%s'> %s </a>", href, s.Text())
			}
		}
	})
	if task.Contents == "" {
		task.Contents = "Без заданий"
	}
}
