package tasks

import (
	"sync"

	"telega/utils"

	"github.com/PuerkitoBio/goquery"
)

type Course struct {
	Name  string
	Link  string
	Tasks []*Task
}

func GetTasks(mainpage *goquery.Document, wg *sync.WaitGroup) {
	if mainpage == nil {
		return
	}
	mainpage.Find("div.course_list").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			href, ok := s.Attr("href")
			if !ok {
				utils.Logger.Println("No href in link", s.Text())
			}
			course := Course{s.Text(), href, make([]*Task, 0)}
			Courses = append(Courses, &course)
			wg.Add(1)
			go sendRequest(&course, wg)
		})
	})
}
