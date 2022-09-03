package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"telega/utils"
)

type Local_Names struct {
	Ru string `json:"ru"`
}
type City struct {
	Name        string      `json:"name"`
	Lat         float64     `json:"lat"`
	Lon         float64     `json:"lon"`
	Local_names Local_Names `json:"local_names"`
	Country     string      `json:"country"`
	State       string      `json:"state"`
}

func GetCities(query string) *City {
	var resp, err = http.Get(fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s", query, API_KEY))
	if err != nil {
		panic("Can't get data from API")
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	cities := make([]City, 0)

	if err := json.Unmarshal(data, &cities); err != nil {
		utils.Logger.Fatalln(err)
	}
	if len(cities) > 0 {
		return &cities[0]
	}
	return nil
}
