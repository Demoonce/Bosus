package news

import (
	"net/http"
	"telega/utils"

	"github.com/PuerkitoBio/goquery"
)

func GetNews() []string {
	results := make([]string, 0)
	resp, err := http.Get("https://lenta.ru")
	if err != nil {
		utils.Logger.Fatalln("Can't get news", err)
	}
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	document.Find("span.card-mini__title").Each(func(i int, s *goquery.Selection) {
		if i < 10 {
			results = append(results, s.Text())
		}

	})
	return results
}
