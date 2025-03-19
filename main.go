package main

import (
	"github.com/svnaditya/telegram-x-bot/config"
	"github.com/svnaditya/telegram-x-bot/telegram"
)

func main() {
	config.LoadConfig()
	telegram.StartBot()
}
