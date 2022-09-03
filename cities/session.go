package cities

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CitiesSession struct {
	Username     string
	CalledCities []City
}

// Checks if session has already been started. If not, then the function create new session. Returns gotten/created session
func AddSession(message *tg.Message) *CitiesSession {
	for _, a := range Sessions {
		if a.Username == message.From.UserName {
			RemoveSession(a)
		}
	}
	result := &CitiesSession{
		message.From.UserName,
		make([]City, 0),
	}
	Sessions = append(Sessions, result)
	return result
}

func GetSession(message *tg.Message) *CitiesSession {
	for _, a := range Sessions {
		if message.From.UserName == a.Username {
			return a
		}
	}
	return nil

}

func RemoveSession(session *CitiesSession) {
	for i, a := range Sessions {
		if a == session {
			Sessions = append(Sessions[0:i], Sessions[i+1:]...)
			return
		}
	}
}
