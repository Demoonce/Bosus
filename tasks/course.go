package tasks

import (
	"errors"
	"sync"

	"telega/utils"

	"github.com/PuerkitoBio/goquery"
)

type Course struct {
	Name  string
	Link  string
	Tasks []*Task
}

func GetTasks(mainpage *goquery.Document, wg *sync.WaitGroup) error {
	if mainpage == nil {
		return errors.New("Can't get document of the site")
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
	return nil
}
