package utils

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	Api, Err = tg.NewBotAPI("5755474801:AAEb4F1lrCNB7mMsGj6vBucKHAvAN7dIswQ")
	Logger   *log.Logger
)
