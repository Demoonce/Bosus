package utils

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	Api    *tg.BotAPI
	Logger *log.Logger
)
