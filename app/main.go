package main

import (
	"os"

	"github.com/mark00s/discord-attendance-list/bot"
)

var DISCORD_BOT_TOKEN = os.Getenv("DISCORD_BOT_TOKEN")

func main() {
	bot.BotToken = DISCORD_BOT_TOKEN
	bot.Run()
}
