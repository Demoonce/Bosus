package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"telega/utils"
)

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type Main struct {
	Temp       float64 `json:"temp"`
	Feels_like float64 `json:"feels_like"`
	Pressure   float64 `json:"grnd_level"`
}
type Weather struct {
	// Main        string `json:"main"`
	Description string `json:"description"`
}
type WeatherObj struct {
	Coord   Coord     `json:"coord"`
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
}

func GetWeatherData(lat, lon float64) *WeatherObj {
	request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&lang=ru&units=metric", lat, lon, API_KEY)
	resp, err := http.Get(request)
	utils.Log("Can't get data", err)

	data, err := io.ReadAll(resp.Body)
	utils.Log("Can't get data from response body", err)

	weather := new(WeatherObj)
	json.Unmarshal(data, weather)
	return weather
}
