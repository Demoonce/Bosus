package tasks

import (
	"fmt"
	"net/http"
	"sync"

	"telega/utils"

	"github.com/PuerkitoBio/goquery"
)

type Task struct {
	Name string
	Link string

	Contents string
}

// sends request to the specified course, returns links of the last and the prelast tasks. Used internally.
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
			"",
		})
	})
}

func GetTaskContent(task *Task) {
	req, err := http.NewRequest("GET", task.Link, nil)
	if err != nil {
		utils.Logger.Fatal(err)
	}
	req.Header.Add("User-Agent", UserAgent)
	resp, err := Client.Do(req)
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	document, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	content_block := document.Find("div.single-section").Find("div.content")
	task.Contents += content_block.Find("div.summary").Text()
	content_block.Find("a").Each(func(i int, s *goquery.Selection) {
		if s != nil {
			href, ok := s.Attr("href")
			if ok && s.Text() != task.Name {
				task.Contents += "\n" + fmt.Sprintf("<a href='%s'> %s </a>", href, s.Text())
			}
		}
	})
	if task.Contents == "" {
		task.Contents = "Без заданий"
	}
}
