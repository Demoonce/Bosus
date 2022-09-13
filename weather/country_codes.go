package weather

import (
	"encoding/gob"
	"net/http"
	"os"

	"telega/utils"

	"github.com/PuerkitoBio/goquery"
)

func GetCodes() map[string]string {
	result := make(map[string]string)
	if _, err := os.Stat("codes.gob"); err == nil {
		file, err := os.Open("codes.gob")
		if err != nil {
			utils.Logger.Fatalln(err)
			return nil
		}
		defer file.Close()
		decoder := gob.NewDecoder(file)
		decoder.Decode(&result)
		return result
	}
	resp, err := http.Get("https://www.borovic.ru/codes.html")
	if err != nil {
		utils.Logger.Println(err)
	}
	file, err := os.Create("codes.gob")
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	encoder := gob.NewEncoder(file)
	document, err := goquery.NewDocumentFromResponse(resp)
	utils.Logger.Println(document.Text())
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	document.Find("td[valign=center]").Each(func(i int, s *goquery.Selection) {
		result[s.Text()] = s.Prev().Text()
	})
	encoder.Encode(result)
	return result
}
