package cities

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telega/utils"
)

func InitCities() {
	Sessions = append(Sessions, &CitiesSession{ // the session of the bot itself
		Username:     BotName,
		CalledCities: make([]City, 0),
	})
}

func RunCities(message *tg.Message) {
	if message.IsCommand() {
		switch com := message.Command(); com {
		case "start":
			new_player := AddSession(message)
			CallCity(new_player, message)
		case "stop":
			player := GetSession(message)
			if player == nil {
				utils.ReplyTo(message, "Вы не играете")
				break
			}
			RemoveSession(player)
			utils.ReplyTo(message, "Вы проиграли")
		case "help":
			utils.ReplyTo(message, "help - помощь\nstart - начать сессию в города\nstop - закончить сессию в города\njoke - случайный анекдот\nnews - новости\nweather 'город' - показать погоду\n")

		case "sessions":
			for _, a := range Sessions {
				utils.Logger.Println(a.Username, a.CalledCities)
				utils.Logger.Print('\n')
			}
		}
	} else if message.ReplyToMessage != nil {
		if message.ReplyToMessage.From.UserName == utils.Api.Self.UserName {
			player := GetSession(message)
			if player != nil && player.Username != utils.Api.Self.UserName {
				if CheckCity(player, message) {
					CallCity(player, message)
				}
			}
		}
	}
}
