package news

import (
	"net/http"
	"telega/utils"

	"github.com/PuerkitoBio/goquery"
)

func GetPanoramaNews() []string {
	news := make([]string, 0)
	resp, err := http.Get("https://panorama.pub")
	if err != nil {
		utils.Logger.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		utils.Logger.Println(err)
		return nil
	}
	document.Find("div.pl-2.pr-1.text-sm.basis-auto.self-center").Each(func(i int, s *goquery.Selection) {
		news = append(news, s.Text())
	})
	return news
}
