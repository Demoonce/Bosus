package weather

import (
	"fmt"

	"telega/utils"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunWeather(message *tg.Message) {
	if message.Command() == "weather" {
		commands := message.CommandArguments()
		if commands != "" {
			city := commands
			if city_data := GetCities(city); city_data != nil {
				curr_weather := GetWeatherData(city_data.Lat, city_data.Lon)
				resp := "%s(%s)\n%s\nТемпература: %.0fºC\nОщущается как: %.0fºC\nАтмосферное давление: %.0f мм.рт.ст"
				reply := tg.NewMessage(message.Chat.ID, fmt.Sprintf(resp, utils.Capitalize(city), utils.Capitalize(city_data.Country), utils.Capitalize(curr_weather.Weather[0].Description), curr_weather.Main.Temp, curr_weather.Main.Feels_like, curr_weather.Main.Pressure/1.3333))
				utils.Api.Send(reply)
			} else {
				utils.Api.Send(tg.NewMessage(message.Chat.ID, "Не знаю такого города"))
			}
		} else {
			utils.Api.Send(tg.NewMessage(message.Chat.ID, "Укажите город параметром команды"))
		}
	}
}
