package utils

import (
	"log"
	"os"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Api    *tg.BotAPI
	Logger *log.Logger = log.New(os.Stdout, "BOT: ", log.Flags())
	Token  string
)
