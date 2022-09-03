package cities

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"telega/utils"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type City struct {
	Name string `json:"name"`
}

// Parses the given file from json and returns a slice of cities that was found
func LoadCities(filename string) ([]City, error) {
	data, err := os.ReadFile(filename)
	result := make(map[string][]City)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result["city"], nil
}

// Sorts cities and makes them lowercase
func PrepareCities() []City {
	Cities_List, err := LoadCities("cities.json")
	sort.Slice(Cities_List, func(i, j int) bool {
		return Cities_List[i].Name < Cities_List[j].Name
	})
	if err != nil {
		log.Fatalln("Unable to unmarshall cities:", err)
	}
	for a := 0; a < len(Cities_List); a++ {
		Cities_List[a].Name = strings.ToLower(Cities_List[a].Name)
	}
	return Cities_List
}

func GenerateCity(last_city City) City {
	rand.Seed(time.Now().Unix())
	last_city_runes := []rune(last_city.Name)
	last_letter := last_city_runes[len(last_city_runes)-1]

	if strings.ContainsAny(string(last_letter), "ьыъ") {
		last_letter = last_city_runes[len(last_city_runes)-2]
	}
	for {
		city_generated := Cities_List[rand.Intn(len(Cities_List))]
		if strings.HasPrefix(city_generated.Name, string(last_letter)) {
			return city_generated
		}
	}
}

// Checks the user's city
func CheckCity(session *CitiesSession, message *tg.Message) bool {
	var message_text = utils.TrimLower(message.Text)
	//checks if city exists
	var loop_broken = false
	for _, city := range Cities_List {
		if city.Name == message_text {
			loop_broken = true
			break
		}
	}
	if !loop_broken {
		utils.ReplyTo(message, "Я не знаю такого города")
		return false
	}
	//checks if city was already called
	for _, a := range session.CalledCities {
		if a.Name == message_text {
			utils.ReplyTo(message, fmt.Sprintf("Вы уже называли город %s", utils.Capitalize(message_text)))
			return false
		}
	}
	//checks if the bot has already called this city
	for _, a := range Sessions[0].CalledCities {
		if a.Name == message_text {
			utils.ReplyTo(message, fmt.Sprintf("Я уже называл город %s", utils.Capitalize(message_text)))
			return false
		}
	}
	//checks if the city is valid (its first letter is the same as the last one of the previous city)
	bot_cities := Sessions[0].CalledCities
	last_city_bot := []rune(bot_cities[len(bot_cities)-1].Name)
	message_text_runes := []rune(message_text)
	check_letter_num := 1
	if strings.ContainsAny(string(last_city_bot[len(last_city_bot)-1]), "ьыъ") {
		check_letter_num++
	}
	if strings.Compare(string(message_text_runes[0]), string(last_city_bot[len(last_city_bot)-check_letter_num])) == 0 {
		session.CalledCities = append(session.CalledCities, City{message_text})
		return true
	} else {
		utils.ReplyTo(message, fmt.Sprintf("Город %s не подходит", utils.Capitalize(message_text)))
		return false
	}
}

// Used tp check and call a city as a bot
func CallCity(session *CitiesSession, message *tg.Message) {
	//if the game just started
	if len(session.CalledCities) == 0 {
		rand.Seed(time.Now().Unix())
		city_called := Cities_List[rand.Intn(len(Cities_List))]
		Sessions[0].CalledCities = append(Sessions[0].CalledCities, city_called)
		utils.ReplyTo(message, utils.Capitalize(city_called.Name))
		return
	}
Pick: //checks if the city was already called
	for {
		city_called := GenerateCity(session.CalledCities[len(session.CalledCities)-1])
		for _, a := range session.CalledCities {
			if a == city_called {
				continue Pick
			}
		}
		Sessions[0].CalledCities = append(Sessions[0].CalledCities, city_called)
		utils.ReplyTo(message, utils.Capitalize(city_called.Name))
		break
	}
}
