package weather

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"telega/utils"
)

func RunWeather(message *tg.Message) {
	if message.Command() == "weather" {
		commands := message.CommandArguments()
		if commands != "" {
			city := strings.Split(commands, " ")[0]
			if city_data := GetCities(city); city_data != nil {
				curr_weather := GetWeatherData(city_data.Lat, city_data.Lon)
				reply := tg.NewMessage(message.Chat.ID, fmt.Sprintf("Город: %s\nТемпература: %.0fºC\nОщущается как: %.0fºC\nАтмосферное давление(на уровне моря): %.0f мм.рт.ст", utils.Capitalize(city), curr_weather.Main.Temp, curr_weather.Main.Feels_like, curr_weather.Main.Pressure/1.3333))
				utils.Api.Send(reply)
			} else {
				utils.Api.Send(tg.NewMessage(message.Chat.ID, "Не знаю такого города"))
			}
		} else {
			utils.Api.Send(tg.NewMessage(message.Chat.ID, "Укажите город параметром команды"))
		}
	}
}
