package talk

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	random "math/rand"
	"os"
	"strings"
	"time"

	"telega/utils"
)

// parses a message file
func ParseMessageFile(filename string) []string {
	data, err := os.ReadFile(filename)
	result := make([]string, 0)
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	Top := make(map[string]any)
	err = json.Unmarshal(data, &Top)
	if err != nil {
		utils.Logger.Fatalln(err)
	}
	if messages, ok := Top["messages"].([]any); ok {
		for _, message := range messages {
			if msg, ok := message.(map[string]any); ok {
				if message_text, ok := msg["text"].(string); ok {
					result = append(result, strings.Split(message_text, " ")...)
				}
			}
		}
	}
	return result
}

// predicts a next word
func Predict(given string) string {
	for {
		rand_index, err := rand.Int(rand.Reader, big.NewInt(int64(len(Messages))))
		if err != nil {
			utils.Logger.Fatalln(err)
		}
		probability, err := Chain.TransitionProbability(Messages[rand_index.Int64()], []string{given})
		if err != nil {
			utils.Logger.Println(err)
			return ""
		}
		fmt.Println(probability, Messages[rand_index.Int64()], []string{given})
		if probability > 0.5 {
			return Messages[rand_index.Int64()]
		}
	}
}

// Generates message of length 4 to 20
func GenerateMsg() string {
	var Message string
	random.Seed(time.Now().Unix())
	message_length := 4 + random.Intn(20)
	rand_index, err := rand.Int(rand.Reader, big.NewInt(int64(len(Messages))))
	if err != nil {
		log.Fatalln(err)
	}
	rand_word := Messages[rand_index.Int64()]
	for a := 0; a < message_length; a++ {
		Message += Predict(rand_word)
	}
	return Message
}
