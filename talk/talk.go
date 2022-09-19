package talk

import (
	"encoding/json"
	"os"
	"strings"

	"telega/utils"

	"github.com/mb-14/gomarkov"
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
func Predict() string {
	message := make([]string, 0)
	for a := 0; a < Order; a++ {
		message = append(message, gomarkov.StartToken)
	}
	for message[len(message)-1] != gomarkov.EndToken {
		word, err := Chain.Generate(message[len(message)-Order:])
		if err != nil {
			utils.Logger.Println(err)
		}
		message = append(message, word)
	}
	return strings.Join(message[Order:len(message)-1], " ")
}
