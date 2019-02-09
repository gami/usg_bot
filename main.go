package main

import (
	"log"
	"usg_bot/config"
	"usg_bot/discord"
)

func main() {
	cfg := config.GetConfig()
	bot := discord.NewBot(cfg)
	err := bot.Run()
	if err != nil {
		log.Fatal(err)
	}
}
