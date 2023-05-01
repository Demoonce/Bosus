package weeb

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"telega/utils"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunWeeb(message *tg.Message) {
	if message.Command() == "g" {
		utils.ReplyTo(message, "Подождите")
		arg := message.CommandArguments()
		Images_Generated++
		data, err := json.Marshal(map[string]string{"inputs": arg})
		if err != nil {
			utils.Logger.Println(err)
		}
		req, err := http.NewRequest("POST", url, bytes.NewReader(data))
		if err != nil {
			utils.Logger.Println("Can't create weeb request", err)
		}
		req.Header.Add("Authorization", api)
		req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
		req.Header.Add("x-use-cache", "false")
		req.Header.Add("x-wait-for-model", "true")
		resp, err := Client.Do(req)
		if err != nil {
			utils.Logger.Println("Error in weeb", err)
		}

		img, err := io.ReadAll(resp.Body)
		if err != nil {
			utils.Logger.Println("Error when reading weeb image", err)
		}
		img_data := tg.FileBytes{Name: "Weeb", Bytes: img}
		photo := tg.NewPhoto(message.Chat.ID, img_data)
		_, err = utils.Api.Send(photo)
		if err != nil {
			utils.ReplyTo(message, "Ошибка")
			utils.Logger.Println("Error whene sending image", err, string(img))
		}
	}
}
